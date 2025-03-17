# repooster
GitHub Repository kickstarter

<!-- *** -->
## What is repooster
With using [`go-github`](https://github.com/google/go-github), `repooster` will do the following configurations:

1. Workflow Permissions
<https://docs.github.com/en/enterprise-cloud@latest/rest/actions/permissions?apiVersion=2022-11-28#set-default-workflow-permissions-for-a-repository>

2. `main` branch protections
<https://docs.github.com/en/rest/branches/branch-protection?apiVersion=2022-11-28#update-branch-protection>

3. Disabling `Discussions`, `Projects`, and `Wiki` tabs
<https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#update-a-repository>

<!-- *** -->
## Distributions
Since `repooster` application has been built on top of CLI driven, the CLI has been available with the several form-factors:
- [GitHub Actions](https://docs.github.com/en/actions)
- Container images, stored in [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

### repooster-actions


#### Usage


#### Inputs


#### Outputs


### Container image

```shell
% docker image pull ghcr.io/hwakabh/repooster:main
```

<!-- *** -->
## Local Setup
Environmental variables, Makefile, docker-compose, ...etc


<!-- *** -->
## Configurations

### The Fine-grained Token permission for this repository.
As this repository will invoke and update configurations of GitHub repository, we need to set GitHub Token, which has permissions of:
- `Read` for Commits
- `Read and write` for Administration

For generating fine-grained tokens, please refer [the official documents](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token) for futher information.

