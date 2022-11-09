# Introduction

This application comes on the device by default.

- Creates `/data/rubix-registry/device_info.json` file if it doesn't exist --it is like a MAC address of each device
- Can upload, delete zip, unzip, move files and dirs in `/data` `/bios` `/home/user` and `/etc/system/systemd`
- Can start, stop, enable, disable, systemctl-reload a service
- BIOS will be installed in the `/bios` dir

### How to Run

- `go build -o rubix-edge-bios main.go && sudo ./rubix-edge-bios server --auth=false --device-type=amd64`
- `go build -o rubix-edge-bios main.go && sudo ./rubix-edge-bios server --auth --device-type=amd64`

## How to Install

- Download the artifacts from the release as per the device type
- Unzip it
- Hit command: `./rubix-edge-bios install --device-type=amd64`
