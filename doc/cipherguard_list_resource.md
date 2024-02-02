doc/cipherguard_list_resource.md## cipherguard list resource

Lists Cipherguard Resources

### Synopsis

Lists Cipherguard Resources

```
cipherguard list resource [flags]
```

### Options

```
  -c, --column stringArray   Columns to return, possible Columns:
                             ID, FolderParentID, Name, Username, URI, Password, Description, CreatedTimestamp, ModifiedTimestamp (default [ID,FolderParentID,Name,Username,URI])
      --favorite             Resources that are marked as favorite
  -f, --folder stringArray   Resources that are in folder
  -g, --group string         Resources that are shared with group
  -h, --help                 help for resource
      --own                  Resources that are owned by me
```

### Options inherited from parent commands

```
      --config string               Config File
      --debug                       Enable Debug Logging
      --filter string               Define a CEl expression as filter for any list commands. In the expression, all available columns of subcommand can be used (see -c/--column).
                                    See also CEl specifications under https://github.com/google/cel-spec.
                                    Examples:
                                    	--filter '(Name == "SomeName" || matches(Name, "RegExpr")) && URI.startsWith("https://auth.")'
                                    	--filter 'Username == "User" && CreatedTimestamp > timestamp("2022-06-10T00:00:00.000-00:00")'
  -j, --json                        Output JSON
      --mfaDelay duration           Delay between MFA Attempts, only used in noninteractive modes (default 10s)
      --mfaMode string              How to Handle MFA, the following Modes exist: none, interactive-totp and noninteractive-totp (default "interactive-totp")
      --mfaRetrys uint              How often to retry TOTP Auth, only used in nointeractive modes (default 3)
      --serverAddress string        Cipherguard Server Address (https://cipherguard.example.com)
      --timeout duration            Timeout for the Context (default 1m0s)
      --totpOffset duration         TOTP Generation offset only used in noninteractive-totp mode
      --totpToken string            Token to generate TOTP's, only used in nointeractive-totp mode
      --userPassword string         Cipherguard User Password
      --userPrivateKey string       Cipherguard User Private Key
      --userPrivateKeyFile string   Cipherguard User Private Key File, if set then the userPrivateKey will be Overwritten with the File Content
```

### SEE ALSO

* [cipherguard list](cipherguard_list)	 - Lists Cipherguard Entitys

