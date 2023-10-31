package packagist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/ubgo/goutil"
)

var developmentalTags = []string{
	"alpha",
	"beta",
	"dev",
	"develop",
	"development",
	"master",
	"rc",
	"rc1",
	"untagged",
	"wip",
}

// "doctrine/dbal"
func FetchAndGetLastesVersionName(repoName string) string {
	repo, _ := FetchRepo(repoName)
	if repo == nil {
		return ""
	}
	pkgs := GetPackagesFromList(repo.PackagesMap)
	pkg := GetLatestPackage(pkgs)

	return pkg.VersionNormalized
}

func FetchRepo(repoName string) (*Repo, error) {
	if len(repoName) == 0 {
		return nil, errors.New("package name is empty")
	}

	url := fmt.Sprintf("https://repo.packagist.org/p2/%s.json", repoName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("unable to fetch package info")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("unable to fetch package info")
	}

	var repo *Repo
	err = json.Unmarshal(body, &repo)
	if err != nil {
		return nil, errors.New("unable to parse json")
	}

	return repo, nil
}

// get repo from file
func OpenRepoFile(filepath string) (*Repo, error) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("unable to read file")
		return nil, err
	}
	var repo *Repo
	err = json.Unmarshal(fileBytes, &repo)
	if err != nil {
		return nil, errors.New("unable to parse json")
	}

	return repo, nil
}

// extract the latest pkg
func GetLatestPackage(packages []Package) *Package {
	var latestPackage Package

	for _, v := range packages {
		if IsDevelopmentPackage(v.VersionNormalized) {
			continue
		}

		if len(latestPackage.VersionNormalized) == 0 {
			latestPackage = v
			continue
		}

		v1, _ := version.NewVersion(latestPackage.VersionNormalized)
		v2, _ := version.NewVersion(v.VersionNormalized)

		if v2.GreaterThan(v1) {
			latestPackage = v
		}
	}
	return &latestPackage
}

// check if the version name is in dev mode e.g. 4.0.0-beta1
func IsDevelopmentPackage(version string) bool {
	for _, v := range developmentalTags {
		ok := strings.Contains(strings.ToLower(version), v)
		if ok {
			return true
		}
	}
	return false
}

// packagist retusn the key > []Packages so extract the Packages Array from the repo
func GetPackagesFromList(pkgMap map[string][]Package) []Package {
	var list []Package
	for _, p := range pkgMap {
		list = p
	}

	return list
}

func ShouldSkip(name, version string) bool {
	skip := false
	_, skip = goutil.StringIndex([]string{"php"}, name)
	if skip {
		return skip
	}

	_, skip = goutil.StringIndex([]string{"dev-main"}, version)

	return skip
}

type Repo struct {
	Minified           string               `json:"minified"`
	PackagesMap        map[string][]Package `json:"packages"`
	SecurityAdvisories []any                `json:"security-advisories"`
}

// type Packages []Package

type Source struct {
	URL       string `json:"url"`
	Type      string `json:"type"`
	Reference string `json:"reference"`
}

type Dist struct {
	URL       string `json:"url"`
	Type      string `json:"type"`
	Shasum    string `json:"shasum"`
	Reference string `json:"reference"`
}

type Package struct {
	Name              string    `json:"name"`
	Version           string    `json:"version"`
	VersionNormalized string    `json:"version_normalized"`
	Source            Source    `json:"source"`
	Dist              Dist      `json:"dist"`
	Time              time.Time `json:"time"`
}
