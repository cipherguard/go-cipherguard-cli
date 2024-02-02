doc/cipherguard_verify.md## cipherguard verify

Verify Setup the Server Verification

### Synopsis

Verify Setup the Server Verification. You need to run this once after that the Server will always be verified if the same config is used

```
cipherguard verify [flags]
```

### Options

```
  -h, --help   help for verify
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

* [cipherguard](cipherguard)	 - A CLI tool to interact with Cipherguard.

