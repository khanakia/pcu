package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/khanakia/pcu/composer"
	"github.com/khanakia/pcu/jsonupdate"
	"github.com/khanakia/pcu/packagist"
	"github.com/khanakia/pcu/util"
	"github.com/spf13/cobra"
)

var version = "dev"

var (
	update bool
)

func main() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringP("file", "f", "./composer.json", "default to current working directory")
	checkCmd.PersistentFlags().StringP("out", "o", "./composer.json", "output file path")
	checkCmd.PersistentFlags().BoolVarP(&update, "update", "u", false, "udpate file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

var rootCmd = &cobra.Command{
	Use:     "pcu",
	Version: version,
	Short:   "php-composer-update updates your package.json dependencies to their latest versions, disregarding any specified version.",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run command `pcu --help` for more information`")
	},
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check dependencies version",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("file")
		out, _ := cmd.Flags().GetString("out")
		update, _ := cmd.Flags().GetBool("update")
		composerCheck(filepath, out, update)
	},
}

func composerCheck(filepath string, out string, update bool) {
	fileBytes, err := composer.OpenFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonupd := jsonupdate.NewJsonUpdate(string(fileBytes))

	var composerFile composer.ComposerFile
	err = json.Unmarshal(fileBytes, &composerFile)
	if err != nil {
		fmt.Println("unable to parse comopser.json file")
	}

	for key, v := range composerFile.Require {
		if packagist.ShouldSkip(key, v) {
			continue
		}
		// fmt.Println(key)
		versionName := packagist.FetchAndGetLastesVersionName(key)
		newVersionName := util.TagNameConvert(versionName)
		fmt.Printf("%-50.50s  %-15s %s\n", key, v, newVersionName)

		jsonupd.Set("require."+key, newVersionName)
	}

	for key, v := range composerFile.RequireDev {
		if packagist.ShouldSkip(key, v) {
			continue
		}
		// fmt.Println(key)
		versionName := packagist.FetchAndGetLastesVersionName(key)
		newVersionName := util.TagNameConvert(versionName)
		fmt.Printf("%-50.50s  %-15s %s\n", key, v, newVersionName)

		jsonupd.Set("require-dev."+key, newVersionName)
	}

	if update {
		_ = os.WriteFile(out, []byte(jsonupd.String()), 0644)
	}

}
