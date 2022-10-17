# What is Tykctl

Tykctl is a WIP CLI. The plan is to have one CLI to interact with all Tyk components and services, tyk cloud, tyk gateway, tyk dashboard for start.

We decided to start with Tyk cloud, as such this is the only service tykctl supports at the moment.

The CLI is grouped into services. For example to use the tyk cloud options you should prefix all your subcommands with: tykctl cloud <subcommand and arguments go here>


### Installation
  #### From Homebrew (macOS)
- This is a private repo hence you will need to set HOMEBREW_GITHUB_API_TOKEN environment variable with a GitHub access token before running brew install.
  ```shell
  export HOMEBREW_GITHUB_API_TOKEN=<Github access token here>
  brew tap TykTechnologies/tykctl https://github.com/TykTechnologies/tykctl
  brew install tykctl 
  ```
  For instructions on how to get the [github access token please follow this guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).
 
  #### Install on Linux 
  
  #### Install our prebuilt binaries
   
  #### Build from source

 ### Commands and usage
 ### Repository structure


