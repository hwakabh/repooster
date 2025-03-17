# repooster
GitHub Repository kickstarter

## What is repooster
With using [`go-github`](https://github.com/google/go-github), `repooster` will do the following configurations:

1. Workflow Permissions
<https://docs.github.com/en/enterprise-cloud@latest/rest/actions/permissions?apiVersion=2022-11-28#set-default-workflow-permissions-for-a-repository>

2. `main` branch protections
<https://docs.github.com/en/rest/branches/branch-protection?apiVersion=2022-11-28#update-branch-protection>

3. Disabling `Discussions`, `Projects`, and `Wiki` tabs
<https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#update-a-repository>

## Distributions
Since `repooster` application has been built on top of CLI driven, the CLI has been available with the several form-factors:
- [GitHub Actions](https://docs.github.com/en/actions)
- Container images, stored in [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

### repooster-actions

### Container image


## Local Setup
Environmental variables, Makefile, docker-compose, ...etc

## Good to know / Caveats
Anything if you have

## License
Choose licenses for your project, see more details in [GitHub Official Docs](https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/adding-a-license-to-a-repository)
