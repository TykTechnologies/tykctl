## tykctl cloud deployments create edge

will create the edge gateway in a given environment

### Synopsis

 
will create an Edge Gateway.

NOTE: This does not deploy the deployment it just create it.You can use the deploy flag to deploy after create.You can also use the deploy command to deploy the created deployment.

You must pass the organization,team,zone and environment you want deploy this deployment.

If you do not pass the org,zone or environment we will use the ones on your config file.You can set the default org,team and environment by running:

		tykctl cloud init

Sample usage for this command

 		tykctl cloud deployments create edge --name="test deployment"


```
tykctl cloud deployments create edge [flags]
```

### Examples

```
tykctl cloud deployments create edge --name='test deployment'
```

### Options

```
      --control-plane string   control plane to link the edge gateway to.
  -h, --help                   help for edge
      --set strings            set a value for the object using dot-notation
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.tykctl/config/config_default.yaml)
      --deploy          deploy the deployment after create
      --domain string   custom domain for your deployment
      --env string      The environment to use
  -n, --name string     name for the deployment you want to create.
      --org string      The organization to use
      --team string     The team to use
      --zone string     the region you want to deploy into e.g aws-eu-west-2
```

### SEE ALSO

* [tykctl cloud deployments create](tykctl_cloud_deployments_create.md)	 - This is the parent command for creating the edge or home deployment.

###### Auto generated by spf13/cobra on 16-May-2024
