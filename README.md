# What is Tykctl

Tykctl is a Work In Progress CLI. The plan is to have one CLI to interact with all Tyk components and services - Tyk Cloud, Tyk Gateway, Tyk Dashboard, etc.

We decided to start with Tyk Cloud, as this is the only service tykctl supports at the moment.

The CLI is grouped into services. For example to use the Tyk Cloud service options you should prefix all your subcommands with: 

`tykctl cloud <subcommand and arguments go here>`


### Installation
   
   #### With Homebrew (recommended for macOS ) 
 - This is a private repo hence you will need to set the `HOMEBREW_GITHUB_API_TOKEN` environment variable with a GitHub personal access token before running `brew install`.
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
       - MacOS - Note for MacOS it recommended that you use Homebrew for easier updates.
    
   #### Build from source (Linux,macOS)
  
  If you want to test the latest changes this is the best way to install tykctl.
   
   ###### Requirements

   - A working Go environment- Some libraries use generics hence you will need Go version 1.18 or later.
   ```
     git clone git@github.com:TykTechnologies/tykctl.git
     go build 
  ```
### Setting Up Autocompletion
   - The cli generate shell completions for:
     - Bash
     - Zsh
     - [Fish](https://fishshell.com/)
     - [PowerShell](https://learn.microsoft.com/en-gb/powershell/)
   - To know the shell you are using run:
     `echo $0` in your terminal.
   - To get the instruction on how to enable autocompletion run:
   
     `tykctl completion <you shell name> --help`

  ## Tyk Cloud Commands and usage
   - The cli is divided into multiple Tyk services.
   - At the moment the Tyk Cloud is the only supported service.
   - Commands should be prefixed by the services you are accessing:
   - For example to use the team command in Tyk Cloud you would enter:
        
     ```tykctl cloud team list```
  
  - Note you have to login first (`tykctl cloud login`) before you run any other commands to get a token that will be used to access the Tyk Cloud services.
  - For tykctl we store the token in your configuration file. The default location of the config file is ($HOME/.tykctl.yaml).However you can pass a different location for your file.
   #### tykctl cloud login
   - Login to Tyk Cloud and it will return a token that will be saved in your configuration file.
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
   #### tykctl cloud init
     
   - `init` initialises the cli and sets your home region, your default organisation , default team and default environment and sets them in your config file.
   - You need to login before you run this command.
   - All commands that require org, team and env flag will use the default set in this step if you do not pass any of them.
   - Usage:
           
      `tykctl cloud init [flag]`
     - Flags:

         ```
          -h, --help          help for the init
          
          --config string     config file to store the (default is $HOME/.tykctl.yaml)
       ```
        
     #### Org commands :
       - These are the actions you can perform on an organisation.

       ##### tykctl cloud org list
     - Lists all of your organisations. The user has only one organisation so in this case only one organisation will be listed.
     - Usage: 
         
         `tykctl cloud org list --output<json/table>`
     - Flags :
        ```
         -o, --output string     Format you want to use. It can be table or json (default is "table")
         
         -h, --help              help for this command
         
         --config string        config file  to use (default is $HOME/.tykctl.yaml)
       ```
        
     #### Team commands
     - These are the actions you can take on a team.
      
      ##### tykctl cloud team create
       - This will create a team with a given name.
       - Usage:
     
         `tyckctl cloud team create --name="first team" --org=<org uuid>`
          
         ```
         --org string           The organisation you want to create the team in. The default organisation in your config file will be used. (required)
          
         -h, --help             help for this command
         
         -n, --name string      name you want to give to the organization you are creating
         
         -o, --output string    Format you want to use. It can be table or json (default is "table")

         --config string        config file (default is $HOME/.tykctl.yaml)
          ```
       ##### tykctl cloud team list
       - This will list the teams in an organisation.
       - Usage:

         `tykctl cloud team list [flags]`
         
       - Flags: 
       ```
       --org string            The organisation you belong to. If not passsed default org in your config will be used.
        
        -h, --help             help for this command
     
       -o, --output string     Format you want to use. It can be table or json (default is "table")
       
       --config string         config file (default is $HOME/.tykctl.yaml)
     ```
    
   #### Environment commands
   - This are the actions you can take on an environment.
   - ##### tykctl cloud environment create
    
      - Create and environment in a specified team and organisation.
      - Usage: 
   
           `tykctl cloud environment create [flags]`
      - 
      - Flags: 
        ```
         -n, --name string     name to give the new environment (required)
         
         --org string          The organization to create this environment in. If omited it will use the one in the configuration file.(required)
        
         --team string         The team you want to create this environment in. (required)
        
         -h, --help            help for this command
         
          --config string      config file to use (default is $HOME/.tykctl.yaml)
         
         -o, --output string   Format you want to use. It can be table or json (default is "table")
       ```
     
   - ##### tykctl cloud environment list
     
     - This will list the environment in your organisation.
     - Usage:

           `tykctl cloud environment list [flags]`
     - Flags :
       ```
       --org string             The organisation whose enviroment you want to list (required)
       
       --team string            The team whose enviroments you want to list (required)
      
        -o, --output string     Format you want to use. It can be table or json (default is "table") 
       
        --config string         config file (default is $HOME/.tykctl.yaml)
       
       -h, --help     help for this command
       ```
   
   #### Deployment commands
       
   This are commands that help you interact with your home and gateway deployments.
       
   ##### tykctl cloud deployment create.
   - Create a deployment in a specified environment.
   - Usage: 
     
      `tykctl cloud deployment create [flags]`
   - Flags:
      ```
      -d, --deploy          deploy the deployment after creation. By default the deployment is not deployed automaticaly
      
      -k, --kind string     The type of deployment you want to create (default "Home")
     
      --team string         The team in which you want to create the deployment (required)
      
      -n, --name string     name to give the new deployment (required)
      
      -z, --zone string     zone you want to deploy into (required)
     
      --org string          The organization you want the deployment to reside
     
      -o, --output string   Format you want to use. It can be table or json (default is "table")
     
      -h, --help     help for this command.
      ```
   ##### tykctl cloud deployment list.
      
   - This will list all the deployments in a given environment.
   - Usage:
    
      `tykctl cloud deployment list [flags]`
   
   - Flags :

      ```
      -o, --output string     Format you want to use. It can be table or json (default is "table")
     
      --config string         config file (default is $HOME/.tykctl.yaml)
     
     --env string             The environment whose deployment you want to list
     
     --org string             The organisation whose deployment you want to list (required)
     
     --team string            The team whose deployment you want to list (required)
     
     -h, --help               help for this command
     ```
      
   ### Repository structure
   
   #### I will work on this once we decide on a structure.


