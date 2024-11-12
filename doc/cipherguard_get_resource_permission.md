doc/cipherguard_get_resource_permission.md## cipherguard get resource permission

Gets Permissions for a Cipherguard Resource

### Synopsis

Gets Permissions for a Cipherguard Resource

```
cipherguard get resource permission [flags]
```

### Options

```
  -c, --column stringArray   Columns to return, possible Columns:
                             ID, Aco, AcoForeignKey, Aro, AroForeignKey, Type, CreatedTimestamp, ModifiedTimestamp (default [ID,Aco,AcoForeignKey,Aro,AroForeignKey,Type])
  -h, --help                 help for permission
      --id string            id of Resource to Get
```

### Options inherited from parent commands

```
      --config string               Config File
      --debug                       Enable Debug Logging
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

* [cipherguard get resource](cipherguard_get_resource)	 - Gets a Cipherguard Resource

