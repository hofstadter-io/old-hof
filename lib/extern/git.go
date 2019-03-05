package extern

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing"
  "gopkg.in/src-d/go-git.v4/storage/memory"
)


func CloneAndWrite(srcUrl, srcVer string) error {


	co := &git.CloneOptions{
			URL: srcUrl,
			// SingleBranch: true,
	}

	if srcVer != "" {
		co.ReferenceName = plumbing.ReferenceName(srcVer)
	}

	err := os.MkdirAll("design-vendor", 0755)
	if err != nil {
		return err
	}

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	_, err = git.PlainClone("design-vendor", false, co)
	if err != nil {
		return err
	}

	return nil
}

func CloneAndRender(srcUrl, srcVer, destBasePath string, data interface{}) error {


	co := &git.CloneOptions{
			URL: srcUrl,
			// SingleBranch: true,
	}

	if srcVer != "" {
		co.ReferenceName = plumbing.ReferenceName(srcVer)
	}

	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	r, err := git.Clone(storer, fs, co)
	if err != nil {
		return err
	}

	t, err := r.Worktree()
	if err != nil {
		return err
	}

	err = walkRepo(t.Filesystem, destBasePath, data)
	if err != nil {
		return err
	}

	return nil
}

func walkRepo(fs billy.Filesystem, destBasePath string, data interface{}) error {

	srcBasePath := "."

	files, err := fs.ReadDir(srcBasePath)
	if err != nil {
		return err
	}

	return walkFiles(fs, files, srcBasePath, destBasePath, data)
}

func walkFiles(fs billy.Filesystem, files []os.FileInfo, srcBasePath, destBasePath string, data interface{}) error {

	for _, file := range files {
		fullPath := filepath.Join(srcBasePath, file.Name())

		if file.IsDir() {
			ff, err := fs.ReadDir(fullPath)
			if err != nil {
				return err
			}

			err = walkFiles(fs, ff, fullPath, destBasePath, data)
			if err != nil {
				return err
			}

		} else {
			err := handleFile(fs, file, srcBasePath, destBasePath, data)
			if err != nil {
				return err
			}

		}


	}

	return nil
}

func handleFile(fs billy.Filesystem, file os.FileInfo, srcBasePath, destBasePath string, data interface{}) error {

	src := filepath.Join(srcBasePath, file.Name())
	dest := filepath.Join(destBasePath, srcBasePath, file.Name())

	fmt.Println(src, "->", dest)

	return nil
}

func renderFile(template string, data interface{}) (string, error) {
	return raymond.Render(template, data)
}

func writeFile(content string, filename string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
