# Tykctl

Tykctl is a command line tool to interact with Tyk Cloud, Tyk Self-Managed and Tyk OSS (open source).

## Overview

Some features `tykctl cloud` provides include:

1. Ability to login to Tyk Cloud.
2. Fetching all your organizations from Tyk Cloud.
3. Creating teams in your Tyk Cloud organization.
4. Create environments for your team.
5. Deploy Tyk Cloud Control Plane and Cloud Data Planes (which are gateways in Tyk Cloud).

## Install

#### With Homebrew (recommended for macOS)

 ```shell
 
  brew tap TykTechnologies/tykctl https://github.com/TykTechnologies/tykctl
  brew install tykctl 
  ```

#### Install our prebuilt binaries

- We do have prebuilt [binaries here](https://github.com/TykTechnologies/tykctl/releases). Download the latest binary
  for your OS unzip it and store it in `$GOPATH/bin` directory
    - Binaries offered:
        - Linux
        - MacOS - Note for MacOS it is recommended that you use Homebrew for easier updates.

#### Build from source (Linux,macOS)

If you want to test the latest changes this is the best way to install `tykctl`.

##### Requirements

A working Go environment- Some libraries use generics hence you will need Go version 1.18 or later.

   ```
     git clone git@github.com:TykTechnologies/tykctl.git
     go build 
  ```

### Docs

Check the full tykctl [documentation here](./docs/tykctl.md).

### Roadmap

- [x] Tyk cloud
- [ ] Adding a verbose flag
- [ ] Adding a loading indicator
- [ ] Tyk Gateway
- [ ] Tyk dashboard

### Project structure

*NOTE: To add a new tyk service to this repo create a new package with your cli code, then add it as a subcommand of the
rootcmd in the sharedCmd package.*

1. **cloudcmd** - This package contains all the code related to the cloud. It performs all the cloud operations.
2. **gatewaycmd** - This package should contain code related to the gateway.
3. **sharedCmd** - This package contains the RootCmd. Here is where you should add a service to the tykctl. For example
   to add the cloud service you should add: `rootCmd.AddCommand(cloudcmd.NewCloudCommand())`.
4. **testutil** - contains shared utility that can be used for testing the tykctl.
5. **util** - contains the utility functions that can be shared by all the tyk service(e.g email validation).
6. **docs** - contains the generated documentation for the cli.
7. **internal** - contains mocks, cloud http clients and all the common functions that will be used within the cli. The
   file labeled `command.go` contains a builder to build the cmd.

### License

Tykctl is released under the MPL v2.0; please see [LICENSE.md](./LICENSE) for a full version of the license.
