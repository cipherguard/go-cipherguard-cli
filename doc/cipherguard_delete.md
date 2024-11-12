doc/cipherguard_delete.md## cipherguard delete

Deletes a Cipherguard Entity

### Synopsis

Deletes a Cipherguard Entity

### Options

```
  -h, --help        help for delete
      --id string   ID of the Entity to Delete
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
* [cipherguard delete folder](cipherguard_delete_folder)	 - Deletes a Cipherguard Folder
* [cipherguard delete group](cipherguard_delete_group)	 - Deletes a Cipherguard Group
* [cipherguard delete resource](cipherguard_delete_resource)	 - Deletes a Cipherguard Resource
* [cipherguard delete user](cipherguard_delete_user)	 - Deletes a Cipherguard User

