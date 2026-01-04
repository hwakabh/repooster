package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v80/github"
)

func main() {
	// Parse argument and prechecks
	arg := os.Args
	ParseArguments(arg)
	repo := arg[1]
	ValidateRepoString(repo)

	repoowner := strings.Split(repo, "/")[0]
	reponame := strings.Split(repo, "/")[1]

	// GitHub operations
	fmt.Println("Hello developer!")
	fmt.Println(">>> Fetching GitHub Personal Access Token ...")
	token := FetchAccessToken()

	fmt.Println(">>> Initializing go-github Client instances")
	client := github.NewClient(nil).WithAuthToken(token)

	fmt.Println(">>> Checking repository initial commit")
	if IsInitialCommit(repo, client, repoowner, reponame) != true {
		fmt.Println("Not an initial commit")
		os.Exit(1)
	}
	fmt.Printf("OK\n\n")

	fmt.Println("Now we can start initialize the repository!")

	fmt.Println(">>> Updating workflow permission for GitHub actions")
	UpdateWorkflowPermission(repo, client, repoowner, reponame)
	fmt.Println("OK")

	fmt.Println(">>> Adding branch protection rule to main branch")
	AddBranchProtectionRule(repo, client, repoowner, reponame)
	fmt.Println("OK")

	fmt.Println(">>> Disabling Wiki/Discussions/Projects tabs from repository")
	DisablingRepositoryTabs(repo, client, repoowner, reponame)
	fmt.Println("OK")

	// // print result: Fetch repository instance
	// // https://github.com/google/go-github/blob/master/github/repos.go#L630
	// r, _, fetcherr := client.Repositories.Get(context.Background(), repoowner, reponame)
	// if fetcherr != nil {
	// 	fmt.Printf("Failed to get information of repository: [ %s ]\n", repo)
	// 	fmt.Println(r)
	// 	os.Exit(1)
	// }
	// fmt.Println(r)

	// Git Operation
	const BranchName string = "feature/init"
	if err := CheckoutWithCreateBranch(BranchName); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(">>> Updating placeholder strings on new branch [ %s ] and trigger git-push ...\n", BranchName)

	files := [...]string{
		"README.md",
		"CITATION.cff",
		".github/CONTRIBUTING.md",
	}
	for _, p := range files {
		fmt.Printf("Replacing string in %s\n", p)
		if err := ReplaceStringInFile(p, "GH_REPONAME", reponame); err != nil {
			fmt.Printf("Failed to update GH_REPONAME in %s\n%s", p, err)
			os.Exit(1)
		}
		if err := ReplaceStringInFile(p, "GH_REPONAME", reponame); err != nil {
			fmt.Printf("Failed to update GH_REPONAME in %s\n%s", p, err)
			os.Exit(1)
		}
		if err := ReplaceStringInFile(p, "GH_USERNAME", repoowner); err != nil {
			fmt.Printf("Failed to update GH_USERNAME in %s\n%s", p, err)
			os.Exit(1)
		}
	}
	GitCommit(files[:])
	fmt.Println("OK")

	push_err := GitPush(repoowner, reponame, BranchName, token)
	if push_err != nil {
		fmt.Println(push_err)
		os.Exit(1)
	}
	fmt.Println("OK")

	// Create GitHub PR
	fmt.Println(">>> Creating PR with repository ...")
	RaisePullRequest(client, repoowner, reponame, BranchName)
	fmt.Println("OK")

	// Slack operations
	fmt.Println(">>> Fetching Slack User Token ...")
	user_token := FetchUserToken()

	fmt.Printf(">>> Creating slack channel for [ %s ]\n", reponame)
	channel_id := CreateChannel(user_token, reponame)
	fmt.Println("OK")

	repourl := fmt.Sprintf("https://github.com/%s/%s", repoowner, reponame)
	fmt.Printf(">>> Set GitHub link to Slack channel topics ...\n")
	AddRepositoryLinkToChannel(user_token, channel_id, repourl)
	fmt.Println("OK")

}
