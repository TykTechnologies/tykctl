## tykctl cloud deployments restart

restart a home or edge gateway deployment given its uuid

### Synopsis


This command will restart a Home or edge gateway given its uuid.

The org,team,environment where the deployment was created has to be provided.

If org,team and environment are not set we will use the default set on your config file. 

Sample usage of this command:

tykctl cloud dep restart --org=<org here> --team=<team here> --env=<environment here> --uid=<deployment id>



```
tykctl cloud deployments restart [flags]
```

### Examples

```
tykctl cloud dep restart <deployment id> --org=<org here> --team=<team here> --env=<environment here> 
```

### Options

```
  -h, --help   help for restart
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.tykctl/config/config_default.yaml)
      --env string      The environment to use
      --org string      The organization to use
      --team string     The team to use
```

### SEE ALSO

* [tykctl cloud deployments](tykctl_cloud_deployments.md)	 - Parent command for all actions you can perform in a deployment.

###### Auto generated by spf13/cobra on 16-May-2024
