doc/cipherguard_share_folder.md## cipherguard share folder

Shares a Cipherguard Folder

### Synopsis

Shares a Cipherguard Folder

```
cipherguard share folder [flags]
```

### Options

```
  -g, --group stringArray   Group id's to share with
  -h, --help                help for folder
      --id string           id of Folder to Share
  -t, --type int            Permission Type (1 Read Only, 7 Can Update, 15 Owner) (default 1)
  -u, --user stringArray    User id's to share with
```

### Options inherited from parent commands

```
      --config string               Config File
      --debug                       Enable Debug Logging
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

* [cipherguard share](cipherguard_share)	 - Shares a Cipherguard Entity

