See [arduino on arch](https://wiki.archlinux.org/title/arduino)

Open `/etc/udev/rules.d/01-ttyusb.rules` and add:
```
SUBSYSTEMS=="usb-serial", TAG+="uaccess"
```

Reload udev rules

```sh
sudo udevadm control --reload
```

Reattach Arduino's USB connection. 

Run ...
```sh
stat /dev/ttyACM0
```

... and it should return with code 0.

## Adding user to group uucp

Arch has a group `uucp` and members can access `/dev/tty...` devices without requiring to use the security privileges of root (using sudo).  It can be checked with `stat`, which lists the Gid. It's called `dialout` on Debian/Ubuntu.

```sh
stat /dev/ttyACM0 
  File: /dev/ttyACM0
  Size: 0         	Blocks: 0          IO Block: 4096   character special file
Device: 0,5	Inode: 1181        Links: 1     Device type: 166,0
Access: (0660/crw-rw----)  Uid: (    0/    root)   Gid: (  986/    uucp)
Access: 2024-02-21 08:42:25.579156457 +0100
Modify: 2024-02-21 08:42:25.579156457 +0100
Change: 2024-02-21 08:42:25.579156457 +0100
 Birth: 2024-02-21 08:42:24.869149891 +0100
```

Your user should be a member:
```sh
groups
docker wheel fabian
```

Add your user if `uucp` is missing and reboot.
```sh
sudo usermod -a -G uucp $USER
```

After a reboot, your user should be a member:
```sh
groups
docker uucp wheel fabian
```