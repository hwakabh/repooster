package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v68/github"
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
	token := FetchAccessToken()

	fmt.Println(">>> Initializing go-github Client instances")
	client := github.NewClient(nil).WithAuthToken(token)

	if IsInitialCommit(repo, client, repoowner, reponame) != true {
		fmt.Println("Not initial commit")
		os.Exit(1)
	}

	fmt.Println("Now we can start initialize the repository!")

	fmt.Println(">>> Updating workflow permission for GitHub actions")
	default_workflow_permission := "write"
	can_approve_pull_request_reviews := true
	permissions := &github.DefaultWorkflowPermissionRepository{
		DefaultWorkflowPermissions:   &default_workflow_permission,
		CanApprovePullRequestReviews: &can_approve_pull_request_reviews,
	}
	_, _, editerr := client.Repositories.EditDefaultWorkflowPermissions(context.Background(), repoowner, reponame, *permissions)
	if editerr != nil {
		fmt.Printf("Failed to update default workflow permissions in %s\n", repo)
		fmt.Println(editerr)
		os.Exit(1)
	}
	fmt.Println("OK")

	fmt.Println(">>> Adding branch protection rule to main branch")
	rules := &github.ProtectionRequest{
		RequiredStatusChecks: &github.RequiredStatusChecks{
			Strict: true,
			// https://github.com/google/go-github/issues/2467#issuecomment-1250072559
			Checks: &([]*github.RequiredStatusCheck{}),
		},
		RequiredPullRequestReviews: &github.PullRequestReviewsEnforcementRequest{
			DismissStaleReviews:          false,
			DismissalRestrictionsRequest: nil,
			RequireCodeOwnerReviews:      false,
			RequiredApprovingReviewCount: 0,
		},
		EnforceAdmins: false,
		Restrictions:  nil,
	}
	_, _, updateerr := client.Repositories.UpdateBranchProtection(context.Background(), repoowner, reponame, "main", rules)
	if updateerr != nil {
		fmt.Printf("Failed to update branch protection rule in %s\n", repo)
		fmt.Println(updateerr)
		os.Exit(1)
	}
	fmt.Println("OK")

	fmt.Println(">>> Disabling Wiki/Discussions/Projects tabs from repository")
	// Fetch repository instance
	// https://github.com/google/go-github/blob/master/github/repos.go#L630
	r, _, fetcherr := client.Repositories.Get(context.Background(), repoowner, reponame)
	if fetcherr != nil {
		fmt.Printf("Failed to get information of repository: [ %s ]\n", repo)
		fmt.Println(r)
		os.Exit(1)
	}
	// fmt.Println(r)
	rbody := struct {
		has_discussions bool
		has_projects    bool
		has_wiki        bool
	}{
		false,
		false,
		false,
	}
	_, _, err := client.Repositories.Edit(context.Background(), repoowner, reponame, &github.Repository{
		HasDiscussions: &rbody.has_discussions,
		HasProjects:    &rbody.has_projects,
		HasWiki:        &rbody.has_wiki,
	})
	if err != nil {
		fmt.Printf("Failed to update project tabs in %s\n", repo)
		os.Exit(1)
	}
	fmt.Println("OK")

	// Git Operation

	// Slack operations

}
