# mattermost-plugin-riff-survey

A mattermost plugin to send surveys for Riff Edu meetings.

## Installation and setup

### Platform & tools

- Make sure you have following components installed:

  - Go - v1.14 - [Getting Started](https://golang.org/doc/install)
    > **Note:** If you have installed Go to a custom location, make sure the `$GOROOT` variable is set properly. Refer [Installing to a custom location](https://golang.org/doc/install#install).

  - NodeJS - v12.18 and NPM - [Downloading and installing Node.js and npm](https://docs.npmjs.com/getting-started/installing-node).

  - Make

## Building the plugins

- Run the following commands to prepare a compiled, distributable plugin zip:

```bash
$ mkdir -p ${GOPATH}/src/github.com/rifflearning
$ cd ${GOPATH}/src/github.com/rifflearning
$ git clone git@github.com:rifflearning/mattermost-plugin-riff-survey.git
$ cd mattermost-plugin-riff-survey
$ make dist
```

- This will produce a `.tar.gz` file in `/dist` directory that can be uploaded to mattermost.

## Installation

1. Go to the [releases page of this GitHub repository](https://github.com/rifflearning/mattermost-plugin-survey/releases/latest) and download the latest release for your Mattermost server.
2. Upload this file in the Mattermost **System Console > Plugins > Management** page to install the plugin. To learn more about how to upload a plugin, [see the documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).
3. You should set **Enable integrations to override usernames** and **Enable integrations to override profile picture icons** in **System Console > Custom Integrations** to `true`.
4. You can configure the Plugin from **System Console > Plugins > Survey**.

## Documentation

Check the `docs` directory for the documentation.

- [Overview](docs/overview.md)
- [System Console Settings](docs/system_console_settings.md)
- [DB Schema](docs/db_schema.md)
