doc/cipherguard_update_user.md## cipherguard update user

Updates a Cipherguard User

### Synopsis

Updates a Cipherguard User

```
cipherguard update user [flags]
```

### Options

```
  -f, --firstname string   User FirstName
  -h, --help               help for user
      --id string          id of User to Update
  -l, --lastname string    User LastName
  -r, --role string        User Role
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

* [cipherguard update](cipherguard_update)	 - Updates a Cipherguard Entity

