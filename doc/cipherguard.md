doc/cipherguard.md## cipherguard

A CLI tool to interact with Cipherguard.

### Synopsis

A CLI tool to interact with Cipherguard.

### Options

```
      --config string               Config File
      --debug                       Enable Debug Logging
  -h, --help                        help for cipherguard
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

* [cipherguard configure](cipherguard_configure)	 - Configure saves the provided global flags to the Config File
* [cipherguard create](cipherguard_create)	 - Creates a Cipherguard Entity
* [cipherguard delete](cipherguard_delete)	 - Deletes a Cipherguard Entity
* [cipherguard export](cipherguard_export)	 - Exports Cipherguard Data
* [cipherguard get](cipherguard_get)	 - Gets a Cipherguard Entity
* [cipherguard list](cipherguard_list)	 - Lists Cipherguard Entitys
* [cipherguard move](cipherguard_move)	 - Moves a Cipherguard Entity
* [cipherguard share](cipherguard_share)	 - Shares a Cipherguard Entity
* [cipherguard update](cipherguard_update)	 - Updates a Cipherguard Entity
* [cipherguard verify](cipherguard_verify)	 - Verify Setup the Server Verification

