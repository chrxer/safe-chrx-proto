> [!WARNING]  
> This project currently does not accept any contribution since it will be graded for school.

### Chromium version

See [chromium.version](chromium.version)

## Initial setup

Currently for linux-only

clone this repo

```bash
clone --recurse-submodules https://github.com/chrxer/safe-chrx-proto.git
cd safe-chrx-proto
```

(optional) open in VSC

```bash
code ./chrxer.code-workspace
```

Install deps

```bash
sudo scripts/deps.bat
```

## Patching guidelines

1. Move as much as possible to `chromium/src/chrxer/*`
2. Minimize modifying files
3. Don't delete any files within `chromium/src`

## Build

Apply patches

```
scripts/patch.py
```

build

```
scipts/build.py
```

## Develop patches

Run (after your edits)

```
scripts/diff.py
```

The diff for modified files can currently be found at [os_crypt.patch](os_crypt.patch) and added files (tree-preserving) at [patch](patch/)

## Building on debian Live USB

On less than 32GB RAM, the build might fail due to OOM (Out-of-memory) and freeze the OS.

Workaround: Create a partition `swap` using `gparted` on the USB-storage device.

> **Note** \
> You might have to resize other partirions to create some space. If this partition has to be mounted whilst running the live usb (such as for `persistence`), you might have to resize it from another Live-USB.

Find the correct partitions

```bash
lsblk -I 8
# NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
# sda      8:0    0 931.5G  0 disk
# ├─sda1   8:1    0   4.5G  0 part /usr/lib/live/mount/medium
# │                                /run/live/medium
# ├─sda2   8:2    0   911G  0 part /usr/lib/live/mount/persistence/sda2
# │                                /run/live/persistence/sda2
# └─sda3   8:3    0    16G  0 part
```

run each of the following commands

```bash
PART=/dev/sda3 # point to your newly created swap partition
sudo mkswap $PART
FS_UUID=$(sudo blkid -o value -s UUID $PART)
sudo swapon -U $FS_UUID
echo "\nUUID=$FS_UUID    none    swap    sw      0   0" | sudo tee -a /etc/fstab
free -h
```
