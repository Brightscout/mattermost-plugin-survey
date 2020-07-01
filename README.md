# mattermost-plugin-survey

A mattermost plugin to send surveys for Riff Edu meetings.

## Installation and setup

### Platform & tools

- Make sure you have following components installed:

  - Go - v1.12 - [Getting Started](https://golang.org/doc/install)
    > **Note:** If you have installed Go to a custom location, make sure the `$GOROOT` variable is set properly. Refer [Installing to a custom location](https://golang.org/doc/install#install).

  - NodeJS - v10.11 and NPM - v6.4.1 - [Downloading and installing Node.js and npm](https://docs.npmjs.com/getting-started/installing-node).

  - Make

## Building the plugins

- Run the following commands to prepare a compiled, distributable plugin zip:

```bash
$ mkdir -p ${GOPATH}/src/github.com/rifflearning
$ cd ${GOPATH}/src/github.com/rifflearning
$ git clone git@github.com:rifflearning/mattermost-plugin-survey.git
$ cd mattermost-plugin-survey
$ make dist
```

- This will produce three tar.gz files in `/dist` directory corresponding to various platforms:

| Flavor  | Distribution |
|-------- | ------------ |
| Linux   | `mattermost-plugin-survey-v<X.Y.Z>-linux-amd64.tar.gz`   |
| MacOS   | `mattermost-plugin-survey-v<X.Y.Z>-darwin-amd64.tar.gz`  |
| Windows | `mattermost-plugin-survey-v<X.Y.Z>-windows-amd64.tar.gz` |

This will also install, **Glide** - the Go package manager.

## Setting up CircleCI

Set up CircleCI to run the build job for each branch and build-and-release for each tag.

1. Go to [CircleCI Dashboard](https://circleci.com/dashboard).
2. In the top left, you will find the Org switcher. Select your Organisation.
3. If this is your first project on CircleCI, go to the Projects page, click the **Add Projects** button, then click the **Set Up Project** button next to your project. You may also click **Start Building** to manually trigger your first build.
4. To manage GitHub releases using CircleCI, you need to add your github personal access token to your project's environment variables.
   - Follow the instructions [here](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line) to create a personal access token. For CircleCI releases, you would need the `repo` scope.
   - Add the environment variable to your project as `GITHUB_TOKEN` by following the instructions [here](https://circleci.com/docs/2.0/env-vars/#setting-an-environment-variable-in-a-project).

## Installation

1. Go to the [releases page of this GitHub repository](https://github.com/rifflearning/mattermost-plugin-survey/releases/latest) and download the latest release for your Mattermost server.
2. Upload this file in the Mattermost **System Console > Plugins > Management** page to install the plugin. To learn more about how to upload a plugin, [see the documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).
3. You should set **Enable integrations to override usernames** and **Enable integrations to override profile picture icons** in **System Console > Custom Integrations** to `true`.
4. You can configure the Plugin from **System Console > Plugins > Survey**.

## Documentation

- [Overview](docs/overview.md)
- [System Console Settings](docs/system_console_settings.md)
- [DB Schema](docs/db_schema.md)
