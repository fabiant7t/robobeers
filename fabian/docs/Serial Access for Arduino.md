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
test -e /dev/ttyACM0
```

... and it should return with code 0.