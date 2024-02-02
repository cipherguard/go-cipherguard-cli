doc/cipherguard_create_user.md## cipherguard create user

Creates a Cipherguard User

### Synopsis

Creates a Cipherguard User and Returns the Users ID

```
cipherguard create user [flags]
```

### Options

```
  -f, --firstname string   First Name
  -h, --help               help for user
  -l, --lastname string    Last Name
  -r, --role string        Role of User.
                           Possible: user, admin (default "user")
  -u, --username string    Username (needs to be a email address)
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

* [cipherguard create](cipherguard_create)	 - Creates a Cipherguard Entity

