package util

import (
	"io/ioutil"

  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing"
)


func CloneRepo(srcUrl, srcVer string) (string, error) {

	co := &git.CloneOptions{
			URL: srcUrl,
			// SingleBranch: true,
	}

	if srcVer != "" {
		co.ReferenceName = plumbing.ReferenceName(srcVer)
	}

	// temp dir to clone to
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	_, err = git.PlainClone(dir, false, co)
	if err != nil {
		return dir, err
	}

	return dir, nil
}
