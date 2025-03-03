## Build release

TBD

## Build tests

TBD


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
