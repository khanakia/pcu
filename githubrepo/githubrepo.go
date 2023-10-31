package githubrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetRepoLatestRelease(pkgname string) (*Repo, error) {
	if len(pkgname) == 0 {
		return nil, errors.New("package name is empty")
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", pkgname)
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

type Repo struct {
	TargetCommitish string `json:"target_commitish"`
	TagName         string `json:"tag_name"`
	Name            string `json:"name"`
}
