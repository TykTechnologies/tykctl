# Tykctl

Tykctl is a command line tool that can be used to interact with tyk cloud.

## Overview

Some features the tykctl provides include:

1. Ability to login to tyk cloud.
2. Fetching all your organizations from tyk cloud.
3. Creating teams in you tyk cloud organization.
4. Create environments in your team.
5. Deploy control plane and edge gateways.

## Install

#### With Homebrew (recommended for macOS)

 ```shell
 
  brew tap TykTechnologies/tykctl https://github.com/TykTechnologies/tykctl
  brew install tykctl 
  ```

For instructions on how to get
the [github access token please follow this guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
.

#### Install our prebuilt binaries

- We do have prebuilt [ binaries here](https://github.com/TykTechnologies/tykctl/releases).Download the latest binary
  for your OS unzip it and store in `$GOPATH/bin` directory
    - Binaries offered:
        - Linux
        - MacOS - Note for MacOS it recommended that you use Homebrew for easier updates.

#### Build from source (Linux,macOS)

If you want to test the latest changes this is the best way to install tykctl.

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
- [ ] Adding verbose flag
- [ ] Adding a loading indicator
- [ ] Tyk Gateway
- [ ] Tyk dashboard

### Project structure

*NOTE: To add a new tyk service to this repo create a new package with your cli code, then add it as a subcommand of the
rootcmd in the sharedCmd package.*

1. cloudcmd - This package contains all the code related to the cloud. It performs all the cloud operations.
2. gatewaycmd - This package should contain code related to the gateway.
3. sharedCmd - This package contains the RootCmd.Here is where you should add a service to the tykctl. For example to
   add
   the cloud service you should add:
   `rootCmd.AddCommand(cloudcmd.NewCloudCommand())`.
4. testutil - contains shared utility that can be used for testing the tykctl.
5. util - contains the utility functions that can be shared by all the tyk service(e.g email validation).
6. docs - contains the generated documentation for the cli.
7. internal - contains mocks,cloud http clients and all the common functions that will be used within the cli.The file
   labeled `command.go` contains a builder to build the cmd.

### License

Tykctl is released under the MPL v2.0; please see [LICENSE.md](./LICENSE) for a full version of the license.
