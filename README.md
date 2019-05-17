# Passgen

Password manager that can generate non and re-generatable passwords and stores their generation credential or existing passwords using military grade (AES-256) encryption.

## WTH is re-generatable password?

It is password that is generated using some credentials. If you supply same credentials you can always generate same password.

## So what makes it more secure than normal?

Password generation requires another password and if you give something different, it generates completely different password. So attacker never know if it is true password or not until trying it. So basically this method redirects attackers to different ways to keep you secure long enough if you feel insecure until you change your passwords, credentials and etc.

## Wait a minute. Can attackers easily access my credentials?

No. If you want to use wallet feature of passgen, it will keeps your credentials in wallet file with military grade AES-256 encryption. So it is nearly impossible to crack it with our current technology if you keep it safe with secure password.

## How do I know my password is secure enough?

Passgen has safety check functionality for your existing passwords that you can use it with this command;

```sh
passgen safety check
```

## Can I use passgen without wallet or storage?

Yes you can. If you believe you can keep your credentials on your mind easily (if you create yourself a pattern for it, it is quite possible) go for it; because it is going to be much more secure than keeping it in wallet.

## What is difference between wallet and storage?

Wallet keeps your re-generatable password's credentials except generation password.

Storage keeps your current existing user names and passwords.

## Can I create single use password with it?

Yes you can. It uses secure random numbers to be sure it is unique. Example usage;

```sh
passgen gen --one-time

# OR

passgen generate --one-time
```

## What is difference wallet/storage password and master password?

Wallet and storage passwords used for securing wallet or storage files.

Master password is used for generating your re-generatable password. So it is a part of re-generation credentials and affects of output of password generation process.

## Can I use different paths for wallet and storage files?

Of course. You can pass `--wallet-path` or `--storage-path` argument to wallet or storage command to change path of it.

## Does it have any detailed help for passgen and it's sub commands?

Passgen uses cobra to handle sub commands and arguments. So it has a phenomenal helper. In case you can't find right command just pass `--help` argument to your command or sub command for help. For example;

```sh
passgen --help
passgen wallet --help
passgen storage --help
passgen safety --help
passgen safety check --help
# and many more
```

---

## TODOs

- [ ] Better TUI