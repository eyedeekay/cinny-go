//go:build gen
// +build gen

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	//"runtime"
	"github.com/otiai10/copy"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	if err := Generate("build"); err != nil {
		log.Fatal(err)
	}
}

func Generate(dir string) error {
	Dir, err := filepath.Abs(dir)
	os.MkdirAll(Dir, 0755)
	if err != nil {
		fmt.Errorf("generate: Absolute directory determination failed for %s, %s", dir, err.Error())
	}
	var repoDir string
	if repoDir, err = gitCloneCinnyRepo(Dir); err != nil {
		return fmt.Errorf("generate: gitCloneI2PFirefox failed %ss", err.Error())
	} else {
		if err := npmCi(repoDir); err != nil {
			return fmt.Errorf("generate: npmCi failed %s", err.Error())

		}
	}
	if distDir, err := npmBuild(repoDir); err != nil {
		return fmt.Errorf("generate: npmBuild failed %s", err.Error())
	} else {
		if err := copyDirectory(distDir); err != nil {
			return fmt.Errorf("generate: copy failed %s", err.Error())
		}
	}
	return nil
}

func gitCloneCinnyRepo(Dir string) (string, error) {
	dir := filepath.Join(Dir, "cinny")
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           "https://github.com/ajbura/cinny",
		Progress:      os.Stdout,
		SingleBranch:  true,
		ReferenceName: plumbing.NewBranchReferenceName("dev"),
	})
	if err != nil {
		log.Printf("gitCloneCinnyRepo: git.PlainClone failed: %s", err.Error())
	}
	log.Println("clone complete")
	return dir, nil
}

func npmCi(repoDir string) error {
	cmd := exec.Command("npm", "ci")
	cmd.Dir = repoDir
	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Println("dependency resolution complete")
	return nil
}

func npmBuild(repoDir string) (string, error) {
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = repoDir
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	dd := filepath.Join(repoDir, "dist")
	log.Println("compilation complete", dd)
	return dd, nil
}

func copyDirectory(distDir string) error {
	err := copy.Copy(distDir, "www")
	if err != nil {
		return err
	}
	log.Println("Copied resources")
	return nil
}
