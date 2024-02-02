doc/cipherguard_get_group.md## cipherguard get group

Gets a Cipherguard Group

### Synopsis

Gets a Cipherguard Group

```
cipherguard get group [flags]
```

### Options

```
  -c, --column stringArray   Membership Columns to return, possible Columns:
                             UserID, Username, UserFirstName, UserLastName, IsGroupManager (default [UserID,Username,UserFirstName,UserLastName,IsGroupManager])
  -h, --help                 help for group
      --id string            id of Group to Get
```

### Options inherited from parent commands

```
      --config string               Config File
      --debug                       Enable Debug Logging
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

* [cipherguard get](cipherguard_get)	 - Gets a Cipherguard Entity

