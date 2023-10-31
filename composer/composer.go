package composer

import (
	"errors"
	"fmt"
	"os"
)

type ComposerFile struct {
	Name             string            `json:"name"`
	Type             string            `json:"type"`
	Description      string            `json:"description"`
	Keywords         []string          `json:"keywords"`
	License          string            `json:"license"`
	Require          map[string]string `json:"require"`
	RequireDev       map[string]string `json:"require-dev"`
	Autoload         Autoload          `json:"autoload"`
	AutoloadDev      Autoload          `json:"autoload-dev"`
	Scripts          Scripts           `json:"scripts"`
	Repositories     []Repositories    `json:"repositories"`
	Extra            Extra             `json:"extra"`
	Config           Config            `json:"config"`
	MinimumStability string            `json:"minimum-stability"`
	PreferStable     bool              `json:"prefer-stable"`
}

type Autoload struct {
	Psr4 map[string]string `json:"psr-4"`
}

type Scripts struct {
	PostAutoloadDump       []string `json:"post-autoload-dump"`
	PostUpdateCmd          []string `json:"post-update-cmd"`
	PostRootPackageInstall []string `json:"post-root-package-install"`
	PostCreateProjectCmd   []string `json:"post-create-project-cmd"`
}

type Repositories struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Extra struct {
	Laravel map[string]interface{} `json:"laravel"`
}

type Config struct {
	OptimizeAutoloader bool            `json:"optimize-autoloader"`
	PreferredInstall   string          `json:"preferred-install"`
	SortPackages       bool            `json:"sort-packages"`
	AllowPlugins       map[string]bool `json:"allow-plugins"`
}

func OpenFile(filepath string) ([]byte, error) {
	info, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) || info.IsDir() {
		fmt.Println("file does not exist")
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("unable to read file")
		return nil, err
	}

	return fileBytes, nil
}
