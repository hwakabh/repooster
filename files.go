package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func ReplaceStringInFile(fname string, placeholder string, replacement string) error {
	// fmt.Printf("replace [ %s ] in [ %s ]\n", placeholder, fname)
	bytes, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println()
		return fmt.Errorf("could not open the file [ %s ]\n", fname)
	}
	contents := string(bytes)
	contents = strings.ReplaceAll(contents, placeholder, replacement)
	err = os.WriteFile(fname, []byte(contents), os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not write the file [ %s ]\n", fname)
	}
	return nil
}

func CheckoutWithCreateBranch(branch_name string) error {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Failed to git-checkout\n")
	}

	// create branch during checkout
	w, _ := repo.Worktree()
	return w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch_name),
		Create: true,
	})
}

// this function will run `git add $filename` & `git commit`
func GitCommit(fnames []string) error {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// git add
	w, _ := repo.Worktree()
	for _, p := range fnames {
		if _, err := w.Add(p); err != nil {
			return fmt.Errorf("failed to git-add. path: %s", p)
		}
	}

	// git commit
	author := &object.Signature{
		Name:  "hwakabh",
		Email: "hrykwkbys1024@gmail.com",
		When:  time.Now(),
	}
	message := "feat: project initialization."
	_, err = w.Commit(message, &git.CommitOptions{Author: author})
	if err != nil {
		return fmt.Errorf("failed to git-commit")
	}

	return nil
}

func GitPush(repoowner string, reponame string, branch_name string, token string) error {
	remote_url := fmt.Sprintf("https://github.com/%s/%s", repoowner, reponame)

	repo, _ := git.PlainOpen(".")
	remote, err := repo.Remote("origin")
	if err != nil {
		return fmt.Errorf("failed to get remote. name: origin")
	}
	fmt.Printf("Uploading changes in [ %s ] to [ %s ] ...\n", branch_name, remote_url)
	return remote.Push(&git.PushOptions{
		RefSpecs:  []config.RefSpec{config.RefSpec(fmt.Sprintf("+refs/heads/%s:refs/heads/%s", branch_name, branch_name))},
		Progress:  os.Stdout,
		RemoteURL: remote_url,
		Auth: &http.BasicAuth{
			Username: repoowner,
			Password: token,
		},
	})

}
