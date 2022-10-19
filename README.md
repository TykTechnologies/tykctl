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

 #### Install our prebuilt binaries
  - We do have prebuilt [ binaries here](https://github.com/TykTechnologies/tykctl/releases).Download the latest binary for your OS unzip it and store in `$GOPATH/bin` directory
    - Binaries offered:
       - Linux
       - MacOS 
    
  #### Build from source (Linux,macOS)
 
   ###### Requirements
   - Working go environment- Some libraries use generics hence you will need go version 1.18 or later.
   ```
     git clone git@github.com:TykTechnologies/tykctl.git
     go build 
  ```

   
 ### Setting Up Autocompletion
   - The cli generate shell completions for:
     - Bash
     - Zsh
     - Fish
     - PowerShell
   - To know the shell you are using run:
     `echo $0` in your terminal.
   - To get the instruction on how to enable autocompletion run:
     `tykctl completion <you shell name> --help`
     

  ### Commands and usage
   - The cli is divided into multiple tyk services.
   - At the moment the tyk cloud is the only supported service.
   - Commands should be prefixed by the services you are accessing:
   - For example to use the team command in tyk cloud you would write:
        ```tykctl cloud team list```

   #### tykctl cloud login
   - Login to tyk cloud and it will return  a token that will be saved in your configuration file.
   - Usage:
      
      `tykctl cloud login [flags]`
   - Flags :
     ```
     --ba-pass string     Basic auth password to be used in staging server
     --ba-user string     Basic auth user to be used in the staging server
     --dashboard string   Url to connect to the dashboard (default "https://dashboard.cloud-ara.tyk.io")
     --email string       The email address you used to sign up for a cloud account
     --password string    The password you used to sigh up for a cloud account.
     

     
     ```
       


  ### Repository structure


