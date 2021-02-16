# Your Tray (WIP)
Add your own app indicator by just setting a yaml file

## Install
```bash
wget https://github.com/ArcaneDiver/your-tray/releases/download/1.0.0/your-tray-v1.0.0-linux-amd64.tar.gz -O your-tray.tar.gz
tar -zxvf your-tray.tar.gz
sudo mv your-tray /usr/local/bin/your-tray
```

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

## Service
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
```
sudo systemctl --user start your-tray.service
```
and keep it running between reboots:
```
sudo systemctl --user enable your-tray.service
```

## Usage
#### Cli arguments
```bash
-config string
    Path to the configuration (default "/etc/your-tray/config.yaml")
-level string
    Log levels: error, warn, info, debug (default "debug")
```
#### YAML Config
```yaml
tray:
  name: My tray # Optional
  tooltip: My wonderfull tray # Optional
  items:
    - text: "Ip: {{ .output }}" # .output is the output of the command
      command: ifconfig <my vpn interface name> | awk '/inet / {print $2}'
      type: data
    - text: Start
      command: sudo systemctl start my-vpn.service
      tooltip: Start the service
      type: command
    - text: Stop
      command: sudo systemctl stop my-vpn.service
      tooltip: Stop the service
      type: command
  icon: /etc/your-tray/my-icon.png # Icon path of the tray
updateRate: 1 # Update rate of type data items
```