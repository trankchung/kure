# Kure

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://godoc.org/github.com/GGP1/kure)
[![Go Report Card](https://goreportcard.com/badge/github.com/GGP1/kure)](https://goreportcard.com/report/github.com/GGP1/kure)

Kure is a command line password manager written in pure Go.

It also offers storing encrypted cards, crypto wallets and files.

This project is not intended for production yet. Although it is secure and reliable enough, I recommend using other managers like [1Password](https://1password.com/), [Keypass](https://keepass.info/), [gopass](https://www.gopass.pw/) and others.

## Table of contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
    * [Commands flags](#commands-flags)
    * [Subcommands](#subcommands)
- [Documentation](#documentation)
    * [How are records stored?](#how-are-records-stored)
    * [Objects](#objects)
    * [Secret generation](#secret-generation)
    * [Encryption](#encryption)
    * [Backup](#backup)
- [Recommendations](#recommendations)
    * [Secure master password](#secure-master-password)
    * [Secret sharing](#secret-sharing)
    * [Double-blind passwords](#double-blind-passwords)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Cross-Platform:** Android, DragonFly BSD, FreeBSD, iOS, Linux, macOS, NetBSD, OpenBSD, Plan 9, Solaris and Windows supported.
- **Offline:** All the data is handled locally. No connection is made with cloud services nor third parties. 
- **Secure:** Every record is encrypted before hitting the database, using a strong [encryption algorithm](#encryption) and the user's master password SHA-512 hash.
- **Simple and easy to use:** Kure is develop with simplicity in mind, not only to use but to mantain aswell.
- **Dynamic:** Perfect for people that uses passwords frequently.
- **Portable:** Can be easily carried around in an external device.
- **Multiple options**: Kure offers not only storing entries but also: bank cards, crypto wallets and files of any type.

## Installation

No releases yet.

## Configuration

Kure by default will create the database and look for the configuration file in the user home directory, unless the path of it is set in an environment variable called `KURE_CONFIG`. Moreover, for executing a configuration for just one command use `<command> --config path/to/file`.

In this file we can also specify where to save our database, modify its name and set a path to a file with **only** the master password (Kure will read all the bytes in file and use that as the master password), allowing the user to use external hardware and create multiple databases with different passwords.

Finally, use `kure login` and `kure logout` to set and remove the **master password hash** from the configuration file. If you decide not to store it, you will be asked for it everytime it's needed. The master password is the only non-encrypted information, our focus is on giving you secure ways of handling it. Make sure no one reads the hash as they will be able to decrypt your information.

*Formats supported*: JSON, TOML, YAML, HCL, envfile and Java properties.

[YAML example](/config_example.yaml)

## Usage

For detailed information about each command, please visit [docs/commands](/docs/commands) folder.

```bash
Usage:
  kure [command]

Available Commands:
  add         Add a new entry to the database
  backup      Create database backups
  card        Card operations
  clear       Clear clipboard/terminal or both
  config      Read or create the configuration file
  copy        Copy entry password to clipboard
  delete      Delete an entry
  edit        Edit an entry
  file        File operations
  gen         Generate a random password
  help        Help about any command
  list        List an entry or all the entries
  login       Set master password
  logout      Unset master password
  stats       Show database statistics
  wallet      Wallet operations

Flags:
      --config string   config file path
  -h, --help            help for kure

Use "kure [command] --help" for more information about a command.
```

#### Commands flags

```
kure add <name> [-c custom] [-l length] [-f format] [-i include] [-e exclude] [-r repeat]
kure backup [http] [port] [encrypt] [decrypt] [path]
kure card
kure clear [-b both] [-c clipboard] [-t terminal]
kure config [-c create] [-p path]
kure copy <name> [-t timeout]
kure delete <name>
kure edit <name>
kure gen [-l length] [-f format] [-i include] [-e exclude] [-r repeat] [-q qr]
kure help
kure file
kure list <name> [-H hide] [-q qr]
kure login
kure logout
kure stats
kure wallet
```

#### Subcommands

[kure add](/docs/commands/add/add.md): phrase.

[kure card](/docs/commands/card/card.md): add, copy, delete, list.

[kure file](/docs/commands/file/file.md): add, create, delete, list.

[kure gen](/docs/commands/gen/gen.md): phrase.

[kure wallet](/docs/commands/wallet/wallet.md): add, copy, delete, list.

## Documentation

### How are records stored?

[Bolt](https://github.com/etcd-io/bbolt) is a **key-value** store that provides an ordered map, which allows easy access and lookup. All collections of key/value pairs are stored in **buckets** within which all keys must be unique. The keys are stored in byte-sorted order within a bucket.

> We use four buckets, one for each type of object. There can't be two records with the same name, Kure will warn you if you are trying to create a record with an already used name.

For example, adding an entry with the name "Go" will look like:

**Bucket**: kure_entry

**Key**: Go

**Value**: encrypted entry object

#### Objects

```
 ENTRY:                                  CARD:                                  FILE:                                   WALLET:

│   FIELD       │      VALUE       │    │   FIELD       │        VALUE     │    │   FIELD       │        VALUE     │    │   FIELD       │        VALUE     │
│───────────────│──────────────────│    │───────────────│──────────────────│    │───────────────│──────────────────│    │───────────────│──────────────────│
│ Name          │ x                │    │ Name          │ x                │    │ Name          │ x                │    │ Name          │ x                │
│───────────────│──────────────────│    │───────────────│──────────────────│    │───────────────│──────────────────│    │───────────────│──────────────────│
│ Username      │ x                │    │ Type          │ x                │    │ Content       │ x                │    │ Type          │ x                │
│───────────────│──────────────────│    │───────────────│──────────────────│    │───────────────│──────────────────│    │───────────────│──────────────────│
│ Password      │ x                │    │ Number        │ x                │    │ Type          │ x                │    │ Script Type   │ x                │
│───────────────│──────────────────│    │───────────────│──────────────────│                                            │───────────────│──────────────────│
│ URL           │ x                │    │ CVC           │ x                │                                            │ Keystore Type │ x                │
│───────────────│──────────────────│    │───────────────│──────────────────│                                            │───────────────│──────────────────│
│ Notes         │ x                │    │ Expire date   │ x                │                                            │ Seed Phrase   │ x                │
│───────────────│──────────────────│                                                                                    │───────────────│──────────────────│
│ Expires       │ x                │                                                                                    │ Public Key    │ x                │
                                                                                                                        │───────────────│──────────────────│
                                                                                                                        │ Private Key   │ x                │
```

#### Secret generation

For generating secure random secrets we use [Atoll](https://www.github.com/GGP1/atoll) (check repository documentation for further information).

#### Encryption

Kure hashes user records with **SHA-256** that then are encrypted with Bernstein's **XChaCha20** symmetric cipher along with **Poly1305** message authentication code. Detailed information [here](https://tools.ietf.org/html/draft-nir-cfrg-chacha20-poly1305-02).

Also, a **SHA-512** hash of the user master password is utilized to encrypt records before saving them into the database and for decryption aswell.

#### Backups

The user can opt to serve the database file on a local server or doing an encrypted backup of it.

## Recommendations

### Secure master password

> A secure password is one you can't remember.

While this briefly explains why you should use a password manager, we all need at least one password to encrypt/decrypt all the others, this is why it is crucial that you choose a **strong master password** to make it as hard as possible to guess. You should **avoid** choosing one that contains words that can be found in a dictionary. Forget putting names or dates of birth.

A good password is a random combination of upper and lower case letters, numbers and special characters. We recommend choosing a password/passphrase consisting of 20 or more characters (the longer, the better).

### Secret sharing

A way of ensuring the security of a secret is [Shamir secret's sharing](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing), where a secret is divided into parts, giving each participant its own unique part.

To reconstruct the original secret, a minimum number of parts is required. In the threshold scheme this number is less than the total number of parts. Otherwise all participants are needed to reconstruct the original secret. This is useful if you prefer to remember file paths instead of secrets, or to share information securely.

Here is a simple CLI program written in Go: [Horcrux](https://github.com/jesseduffield/horcrux).

### Two-factor authentication

Two-factor authentication is a type, or subset, of multi-factor authentication. It is a method of confirming users' claimed identities by using a combination of **two different factors** (usually 1 and 2): 1. something they know (account credentials), 2. something they have (devices), or 3. something they are. So if an attacker gets access to the secrets, he still need the **constantly refreshing code** to get into the account, making it, not impossible, but much more complicated.

### Double-blind passwords

Money related accounts are, in most cases, the ones with more critical information and we must take all precautions to keep them secure. To achieve this, we encourage our users to add a **non-stored sequence of numbers** after or before the password, for example: notRandomPassword*ID* which could be notRandomPassword*65874*.

What will be stored in the database is *notRandomPassword* but that isn't the complete one so the attacker won't have access to your accounts even being able to manage all your secrets. Of course, this sequence of numbers shall be remembered by the user and must not be stored anywhere.

## Contributing

Any contribution is welcome. We appreciate your time and help. Please follow this steps to do it:

1. **Fork** this repository on Github
2. **Clone** your fork by doing: `git clone github.com/<your_username>/kure.git`
3. **Create** your feature branch (`git checkout -b <your-branch>`)
4. Make changes and **run tests** (`go test ./...`)
5. **Add** them to staging (`git add .`)
6. **Commit** your changes (`git commit -m '<changes>'`)
7. **Push** to the branch (`git push origin <your-branch>`)
8. Create a **Pull request**

## License

Kure is licensed under the Apache-2.0 license. See [LICENSE](/LICENSE).