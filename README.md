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
See [build](build/)

Build time without Ccache on:
```
c5ad.xlarge:  
    4 vCPU, 
    8GB RAM, 
    x86_64, 
    150 GB SSD, 
    Up to 10GiB Network, 
    ~0.172  USD/hour

All deps installed: 12~14 min.
Building (estimated): 3.5 ~ 4.5 h ~ 0.7 USD
CPU: 95% ~ 99% during compillation
Network: ~ 700 MBytes/s during deps
Disk max: Read:
```


## Developers
Integrated as a school-project \
[@kaliiiiiiiiii](https://github.com/kaliiiiiiiiii) aka Steve (Single dev) \
[@The-An0nym](https://github.com/The-An0nym)

## Sponsors
None yet:(