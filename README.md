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

### tyckctl cloud team create --name="first team" --org=<org uuid>
This command will create a team.

You have to pass the name you want to give the team and org in which you want to create the team.

If the org is not provided we will use the one you set in the config file.

To set a default team in the config file run:

tykctl cloud init

Sample usage for this command:

tyckctl cloud team create --name="first team" --org=<org uuid>

### tykctl team list --org=<orgID> --output=<json/table>

This command will fetch and list all the teams in an organization.

You must pass the --org flag.If it is not passed we will use the default one set in the config file.

The output can be either json or table. Default is table.
To change the format use --output=<json/table> flag.

Sample usage:

tykctl team list --org=<orgID> --output=<json/table>

### tyk cloud environment create --name="staging"
This command create an environment in a team.

You must pass the name of the environment.

You must also set the org and team you want to create this environment in.

If you don't pass the org and team we will use the one set in the config file.

Sample usage:

tyk cloud environment create --name="staging"

### tckctl cloud environment list --org=<orgID> --output=<json/table>
This command will fetch all the environment in an organization.

You must pass the --org.If it is not passed we will use the default org set in your config file.

We support json and table as the output format.To set the output format use the --output<json/table> flag.

Sample usage of this command:

tckctl cloud environment list --org=<orgID> --output=<json/table>

### tykctl cloud deployment create --name="test deployment" --kind="Home"

This command creates a Home or a Edge Gateway.

NOTE: This does not deploy the deployment it just create it.You can use the deploy command to deploy the created deployment.

You must pass the organization,team,zone and environment you want deploy this deployment.

NOTE: For the home deployment you have to select the you home zone as the deployment zone.

If you do not pass the org,zone or environment we will use the ones on your config file.You can set the default org,team and environment by running:

tykctl cloud init

Sample usage for this command

tykctl cloud deployment create --name="test deployment" --kind="Home"




  


  
