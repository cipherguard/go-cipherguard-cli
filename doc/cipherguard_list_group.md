doc/cipherguard_list_group.md## cipherguard list group

Lists Cipherguard Groups

### Synopsis

Lists Cipherguard Groups

```
cipherguard list group [flags]
```

### Options

```
  -c, --column stringArray    Columns to return, possible Columns:
                              ID, Name, CreatedTimestamp, ModifiedTimestamp (default [ID,Name])
  -h, --help                  help for group
  -m, --manager stringArray   Groups that are in folder
  -u, --user stringArray      Groups that are shared with group
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
      --mfaTotpOffset duration      TOTP Generation offset only used in noninteractive-totp mode
      --mfaTotpToken string         Token to generate TOTP's, only used in nointeractive-totp mode
      --serverAddress string        Cipherguard Server Address (https://cipherguard.example.com)
      --timeout duration            Timeout for the Context (default 1m0s)
      --userPassword string         Cipherguard User Password
      --userPrivateKey string       Cipherguard User Private Key
      --userPrivateKeyFile string   Cipherguard User Private Key File, if set then the userPrivateKey will be Overwritten with the File Content
```

### SEE ALSO

* [cipherguard list](cipherguard_list)	 - Lists Cipherguard Entitys

