# Safe-chrx
A chromium build with an improved password manager and prototype implementation of [chrx-design](https://github.com/chrxer/chrx-design/tree/main/chrx).

## Why?
Chromium's native password manager is **poorly implemented** (see [password_harvester.py](https://github.com/kaliiiiiiiiii/PublicDuckyChallenger/blob/master/pass_harvester/password_harvester.py)) with the following excuse [source](https://chromium.googlesource.com/chromium/src/+/HEAD/docs/security/faq.md#why-arent-physically_local-attacks-in-chromes-threat-model)

> We consider these attacks outside Chrome's threat model, because there is no way for Chrome (or any application) to defend against a malicious user who has managed to log into your device as you, or who can run software with the privileges of your operating system user account.

## How?
Our password manager requires a **master password** when autofilling the first time after startup. All credentials are stored encrypted on the user's disk over that password.

## Architecture
![Architecture](bridge.drawio.svg)

## Roadmap
See [roadmap](roadmap/)

## Building

see [CONTRIBUTING.md](CONTRIBUTING.md)

<details>

<summary>on EC2</summary>

Build time without existing Ccache on EC2:
```
c5ad.2xlarge: 0.344 USD/h
Compillation start: ~ 9 min. -> 0.05 USD
CPU: ~ 99% during compilation
~7 h -> 2.4 USD
```

with ccache:
```
Hits: 99.97%
Errors: 0.03% (Could not read or parse input file:)
Deps installed: ~ 11 min.
50 min -> 0.29 USD
```

File sizes (131.0.6778.139)
```
Sources: 24G (downloaded from google server)
Ccache:  3.4G (tar.gz archive => S3)
```

</details>

## Developers
Integrated as a school-project \
[@kaliiiiiiiiii](https://github.com/kaliiiiiiiiii) aka Steve (Single dev) \
[@The-An0nym](https://github.com/The-An0nym)

## Sponsors
None yet:(
