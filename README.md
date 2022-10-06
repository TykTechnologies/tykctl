# tykctl
Tykctl is a cli that can be used to interact with all tyk components (tyk cloud,tyk gateway and tyk dashboard).

The cli is grouped into services.
For example to use the tyk cloud services you should prefix all your subcommands with:
tykcli cloud <subcommand here>

## Commands and usage

### tyk cloud login - login  into the tyk cloud dashboard
This command will login into your cloud account and set the token in your config file.

Note: The token will only last for 30 minute you will need to login again after 30 minutes.

You will be prompted to provide your email and  password to login.

When using the cloud service you should always run this command first as each command will require a token.

For the staging server you will also need to provide nginx basic auth.

Sample usage:
tykctl cloud login --ba-pass=<use this only is staging> --ba-pass=<use this in staging>

`


### tyk cloud init
This command will initialize the cli and set default in the config file.

Before using this command you will need to login with:
tykctl cloud login

Use this command to:
1. Set the default organization.
2. Set the default team
3. Set the default environment.
4. Set your zone and home region.

This command should ideally be run after the login command.
`



  


  
