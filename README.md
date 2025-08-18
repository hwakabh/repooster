# repooster
GitHub Repository kickstarter \
`repooster` is coming from "repo" + "booster", and this application will intend to speeding up your GitHub repository setup with single commands.


<!-- *** -->
## Features
For boosting up the scaffolding repository setup to developement, `repooster` will do sequencially:
1. Precheck of initial commit
2. GitHub Operations (updating repository setting)
3. Slack Operations (creating notification channel for repository)
4. File Operations (editing README.md or any other templated files and pushing to remote)
5. GitHub Operations (raising PR for initialization)

Please note that above features are based on the repository, which have been created from [specific GitHub template](https://github.com/hwakabh/.github), so if you would like to customize more, you could add features to the repository template as well as `repooster`

### GitHub Operations
With using [`go-github`](https://github.com/google/go-github), `repooster` will do the following configurations:

- [Workflow Permissions](https://docs.github.com/en/enterprise-cloud@latest/rest/actions/permissions?apiVersion=2022-11-28#set-default-workflow-permissions-for-a-repository)
- [`main` branch protections](https://docs.github.com/en/rest/branches/branch-protection?apiVersion=2022-11-28#update-branch-protection)
- [Disabling `Discussions`, `Projects`, and `Wiki` tabs](https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#update-a-repository)
- [Creating initial PR](https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#create-a-pull-request) once after complete all File operations below

### File Operations
With using [`go-git`](https://github.com/go-git/go-git), `repooster` will do:
1. Create branch called `feature/init` with checkout
2. Replace placeholder texts in templated files
3. Stage all changes to working tree and create commit object
4. Push the commit to remote repository

### Slack Operations
For fetching updates on your new repository, Slack notification integrated with GitHub repository is one of the best way. \
By interacting with [Slack Web APIs](https://api.slack.com/methods), `repooster` will do:
- create new dedicated channel for your repo by [`conversations.create` endpoint](https://api.slack.com/methods/conversations.create)
- set link of your GitHub repository as channel topics by [`conversations.setTopic` endpoint](https://api.slack.com/methods/conversations.setTopic)

Also please note that you have to set `SLACK_USER_TOKEN` beforehand. \
As Slack APIs can handle [several types of its tokens](https://api.slack.com/concepts/token-types), but `repooster` will expect to use OAuth User Token in general.


<!-- *** -->
## Configurations

### `TOKEN`
As this repository will invoke and update configurations of GitHub repository, we need to set GitHub Token, which has permissions of:
- `Read` for Commits
- `Read and write` for Administration

While GitHub has several types of token, we generally expect to use Fine-grained Token. \
For generating fine-grained tokens, please refer [the official documents](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token) for futher information.

### `SLACK_USERTOKEN`
TBA


<!-- *** -->
## Distributions
Since `repooster` application has been built on top of CLI driven, the CLI has been available with the several form-factors:
- Single Binary
- Container images, stored in [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

### Single Binary
As the `repooster` has been developed with Go, which can easily build artifacts for several platforms, we have uploaded various binaries for general platforms. \
You can download the binary from [GitHub Releases](https://github.com/hwakabh/repooster/releases) of this repo, and currently we are supporting the following platforms:
- `darwin/arm64`
- `darwin/amd64`
- `linux/arm64`
- `linux/amd64`
- `windows/amd64`

After downloading the binary, you can easily start using of it, once you exported GitHub tokens (described below).

```shell
% export TOKEN='...'
% export SLACK_USER_TOKEN='...'
% repooster hwakabh/repooster
```

### Container image
As we are using [`ko`](https://ko.build) for building containerized Go application with OCI format, the container images has been also available from [GitHub Packages](https://github.com/hwakabh/repooster/pkgs/container/repooster). \
Same as binaries, you can start using CLI after downloading image onto your environment.

```shell
% export TOKEN='...'
% export SLACK_USER_TOKEN='...'
% docker image pull ghcr.io/hwakabh/repooster:main
% docker run -e TOKEN=$TOKEN -e SLACK_USER_TOKEN=$SLACK_USER_TOKEN ghcr.io/hwakabh/repooster:latest hwakabh/repooster
```

<!-- *** -->
## Local Setup
For building application on your local environment, please source this repository and run general go build processes like:

```shell
% git clone git@github.com:hwakabh/repooster.git
% cd repooster

% go build .
```

In case you need to build OCI image on your local environment, please install `ko` first, and run:

```shell
% ko version
0.17.1

# Build locally
% ko build -L .
2025/04/30 02:40:16 Using base cgr.dev/chainguard/static:latest@sha256:2e3db1641bb4fe4e85d2210f4aadb79252e90d5fa745f53a3ffed6a1aab4f73b for github.com/hwakabh/repooster
2025/04/30 02:40:17 Building github.com/hwakabh/repooster for linux/amd64
2025/04/30 02:40:18 Loading ko.local/repooster-ee3247f55ae92694152f961e9a3e01e8:210a86a0fd8c178b47b0826f1e3c8593913560c62bfff222d6f56b2f4c58d94a
2025/04/30 02:40:19 Loaded ko.local/repooster-ee3247f55ae92694152f961e9a3e01e8:210a86a0fd8c178b47b0826f1e3c8593913560c62bfff222d6f56b2f4c58d94a
2025/04/30 02:40:19 Adding tag latest
2025/04/30 02:40:19 Added tag latest
ko.local/repooster-ee3247f55ae92694152f961e9a3e01e8:210a86a0fd8c178b47b0826f1e3c8593913560c62bfff222d6f56b2f4c58d94a

% docker image ls
REPOSITORY                                            TAG                                                                IMAGE ID       CREATED         SIZE
ko.local/repooster-ee3247f55ae92694152f961e9a3e01e8   210a86a0fd8c178b47b0826f1e3c8593913560c62bfff222d6f56b2f4c58d94a   9c7460230cb9   4 weeks ago     6.23MB
ko.local/repooster-ee3247f55ae92694152f961e9a3e01e8   latest                                                             9c7460230cb9   4 weeks ago     6.23MB
```

Please note that currently we do not have any customized configurations with ko, so the default base image is [`cgr.dev/chainguard/static`](https://images.chainguard.dev/directory/image/static/versions), which is generally minimal distroless image, and it does not have any shells. \
If you would like to use shells for debugging purpose, please note that you need to override the base image with using `KO_DEFAULTBASEIMAGE` variables with `ko build` command in your local development. \
Please refer [`ko` official documents](https://ko.build/configuration/) for more details.

For contribuing this project if you would like, please check [the docs](./CONTRIBUTING.md) first. \
We are always welcome for any contributions.
