# Your Tray (WIP)
Add your own app indicator by just setting a yaml file

## Dependencies

#### Build essential
```bash
sudo apt-get install build-essential
```

#### GTK-3-dev
```bash
sudo apt-get install libgtk-3-dev
```

#### Appindicator 3.0.1
```bash
sudo apt-get install libappindicator3-dev gir1.2-appindicator3-0.1
```

## Install
#### Service
Create a service:
```bash
sudo nano /etc/systemd/user/your-tray.service
```
Use this configuration:
```
[Unit]
Description=Your tray
After=default.target

[Service]
Type=simple
ExecStart=<path to the bin>

[Install]
WantedBy=multi-user.target

```
You can start the service with:
```bash
sudo systemctl --user start your-tray.service
```
and keep it running between reboots:
```bash
sudo systemctl --user enable your-tray.service
```