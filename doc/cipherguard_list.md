doc/cipherguard_list.md## cipherguard list

Lists Cipherguard Entitys

### Synopsis

Lists Cipherguard Entitys

### Options

```
      --filter string   Define a CEl expression as filter for any list commands. In the expression, all available columns of subcommand can be used (see -c/--column).
                        See also CEl specifications under https://github.com/google/cel-spec.
                        Examples:
                        	--filter '(Name == "SomeName" || matches(Name, "RegExpr")) && URI.startsWith("https://auth.")'
                        	--filter 'Username == "User" && CreatedTimestamp > timestamp("2022-06-10T00:00:00.000-00:00")'
  -h, --help            help for list
  -j, --json            Output JSON
```

### Options inherited from parent commands

```
      --config string               Config File
      --debug                       Enable Debug Logging
      --mfaDelay duration           Delay between MFA Attempts, only used in noninteractive modes (default 10s)
      --mfaMode string              How to Handle MFA, the following Modes exist: none, interactive-totp and noninteractive-totp (default "interactive-totp")
      --mfaRetrys uint              How often to retry TOTP Auth, only used in nointeractive modes (default 3)
      --mfaTotpOffset duration      TOTP Generation offset only used in noninteractive-totp mode
      --mfaTotpToken string         Token to generate TOTP's, only used in nointeractive-totp mode
      --serverAddress string        Cipherguard Server Address (https://cipherguard.example.com)
      --timeout duration            Timeout for the Context (default 1m0s)
      --userPassword string         Cipherguard User Password
      --userPrivateKey string       Cipherguard User Private Key
      --userPrivateKeyFile string   Cipherguard User Private Key File, if set then the userPrivateKey will be Overwritten with the File Content
```

### SEE ALSO

* [cipherguard](cipherguard)	 - A CLI tool to interact with Cipherguard.
* [cipherguard list folder](cipherguard_list_folder)	 - Lists Cipherguard Folders
* [cipherguard list group](cipherguard_list_group)	 - Lists Cipherguard Groups
* [cipherguard list resource](cipherguard_list_resource)	 - Lists Cipherguard Resources
* [cipherguard list user](cipherguard_list_user)	 - Lists Cipherguard Users

