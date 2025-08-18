package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v74/github"
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

	// Slack operations
	user_token := FetchUserToken()
	fmt.Printf(">>> Creating slack channel for [ %s ]\n", reponame)
	channel_id := CreateChannel(user_token, reponame)
	fmt.Println("OK")

	repourl := fmt.Sprintf("https://github.com/%s/%s", repoowner, reponame)
	fmt.Printf(">>> Set GitHub link to Slack channel topics ...\n")
	AddRepositoryLinkToChannel(user_token, channel_id, repourl)
	fmt.Println("OK")

}
