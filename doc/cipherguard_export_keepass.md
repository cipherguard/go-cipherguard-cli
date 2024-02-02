doc/cipherguard_export_keepass.md## cipherguard export keepass

Exports Cipherguard to a Keepass File

### Synopsis

Exports Cipherguard to a Keepass File

```
cipherguard export keepass [flags]
```

### Options

```
  -f, --file string       File name of the Keepass File (default "cipherguard-export.kdbx")
  -h, --help              help for keepass
  -p, --password string   Password for the Keypass File, if empty prompts interactively
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

* [cipherguard export](cipherguard_export)	 - Exports Cipherguard Data

