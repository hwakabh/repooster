# repooster
GitHub Repository kickstarter \
`repooster` is coming from "repo" + "booster", and this application will intend to speeding up your GitHub repository setup with single commands.

<!-- *** -->
## What is repooster
With using [`go-github`](https://github.com/google/go-github), `repooster` will do the following configurations:

1. [Workflow Permissions](https://docs.github.com/en/enterprise-cloud@latest/rest/actions/permissions?apiVersion=2022-11-28#set-default-workflow-permissions-for-a-repository)

2. [`main` branch protections](https://docs.github.com/en/rest/branches/branch-protection?apiVersion=2022-11-28#update-branch-protection)

3. [Disabling `Discussions`, `Projects`, and `Wiki` tabs](https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#update-a-repository)

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
% repooster hwakabh/repooster
```

### Container image
As we are using [`ko`](https://ko.build) for building containerized Go application with OCI format, the container images has been also available from [GitHub Packages](https://github.com/hwakabh/repooster/pkgs/container/repooster). \
Same as binaries, you can start using CLI after downloading image onto your environment.

```shell
% export TOKEN='...'
% docker image pull ghcr.io/hwakabh/repooster:main
% docker run -e TOKEN=$TOKEN ghcr.io/hwakabh/repooster:latest hwakabh/repooster
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

<!-- *** -->
## Configurations

### `TOKEN`
As this repository will invoke and update configurations of GitHub repository, we need to set GitHub Token, which has permissions of:
- `Read` for Commits
- `Read and write` for Administration

While GitHub has several types of token, we generally expect to use Fine-grained Token. \
For generating fine-grained tokens, please refer [the official documents](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token) for futher information.

