---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-run-unifi-controller-in-docker-container-a59cd191-3
title: "How To Run UniFi Controller in Docker Container"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-run-unifi-controller-in-docker-container-a59cd191.md
source_anchor: ""
source_lines: [139, 191]
sha256: 41c779d3321e2440ff31b4acf3cb8ab0fc75fb3f75b9859d1c602d4c3a782475
---

# How To Run UniFi Controller in Docker Container

      - 3478:3478/udp
      - 10001:10001/udp
      - 8080:8080
    restart: unless-stopped
Save the file and start the stack. Recent Docker ships Compose as a plugin, so the command is docker compose, not the old docker-compose binary:
docker compose up -d
Both containers come up:
 Container unifi-db  Started
 Container unifi-network-application  Started
Whichever way you started it, verify both containers are running:
docker ps
The controller and its database both show Up:
NAMES                       IMAGE                                                  STATUS          PORTS
unifi-network-application   lscr.io/linuxserver/unifi-network-application:latest   Up 2 minutes    0.0.0.0:8443->8443/tcp, 0.0.0.0:8080->8080/tcp, 0.0.0.0:3478->3478/udp, 0.0.0.0:10001->10001/udp
unifi-db                    mongo:8.0                                              Up 2 minutes    27017/tcp
Both of these came up in testing and in reader reports, and both leave the web UI unreachable.
If the init script does not exist as a file when the containers first start, Docker silently creates an empty directory at that path. The database then comes up without the unifi user and the controller can never connect. Check what you have:
ls -la /unifi_data/init-mongo.sh
If that lists a directory, tear the stack down, remove the directory together with the half-initialised database, recreate the script as a file per the persistent volumes section, and start again:
docker compose down
sudo rm -rf /unifi_data/init-mongo.sh /unifi_data/db
sudo mkdir -p /unifi_data/db
This appears in /unifi_data/config/logs/server.log when the config volume still holds data from the retired unifi-controller image, which bundled its own database. The new image sees the old configuration and tries to launch a mongod binary it does not ship:
ERROR system - unable to exec
java.io.IOException: Cannot run program "bin/mongod" (in directory "/usr/lib/unifi"): Exec failed, error: 2 (No such file or directory)
Export a backup from the old controller first (Settings > System > Backups) if you need the configuration, then stop the container, clear /unifi_data/config, start fresh, and restore the backup through the setup wizard.
Now access the UniFi Controller web UI using the URL https://10.0.1.50:8443
Set the name of the application and proceed to sign in using your Ubiquiti account.
Configure the network.
The devices are not available since the application is running in a docker container. So skip and configure this later.
This too can be skipped and set later.
Review the configurations.
The configurations will be made as shown.
Once complete you will see the below dashboard.
Adoption is where a self-hosted controller behaves differently from a UniFi console. A factory access point looks for a controller on its local layer 2 segment, then keeps reporting in to whatever inform URL it is handed. A controller in a container answers on the Docker host’s LAN address, not on the internal bridge IP a device would otherwise learn, so the fix is to advertise an address the devices can actually reach.
Set that address once. Open Settings, then System, scroll to Advanced, switch Inform Host to manual, and enter the LAN IP or DNS name of the Docker host (here 10.0.1.50):
Apply the change. Any access point on the same subnet as the host is now discovered automatically and shows up under Devices ready to adopt.
When a device never appears, point it at the controller by hand over SSH. This is the reliable path when the AP sits on a different subnet from the container:
ssh ubnt@10.0.1.60
set-inform http://10.0.1.50:8080/inform
Replace 10.0.1.60 with the device address and 10.0.1.50 with the Docker host. A factory device signs in with ubnt / ubnt. Run set-inform once so the device shows as Pending Adoption, click Adopt in the controller, then run it a second time after the device provisions so it locks onto the controller for good.
Layer 2 discovery does not cross a router, so a controller in one VLAN and access points in another need one of the standard hand-off methods:
01:04:<controller-ip-in-hex>.unifi that resolves to the controller, which factory devices query on boot.set-inform command above, run once per device.
The inform host set earlier is what makes all three work: the device reaches the controller, and the controller replies with an inform URL it can keep using.
The payoff of running the controller yourself is the same segmented network a hardware UniFi console would build: a VLAN for a class of devices and a wireless network whose traffic is tagged onto it. Both are defined once in the controller and pushed to every adopted gateway, switch, and access point.
Create the VLAN first. Open Settings, then Networks, and add a network. Name it, set the VLAN ID (here 20 for an IoT segment), and either let the controller run DHCP for the subnet or mark it VLAN-only when an upstream router already serves it:
With the network saved, create the SSID that rides it. Open Settings, then WiFi, add a WiFi network, set the name and passphrase, then under the network option choose the VLAN you just created so client traffic is tagged with that VLAN ID:
Every access point the controller has adopted now broadcasts the SSID and tags its clients onto the VLAN, exactly as a UniFi gateway would. The tag only carries end to end when the switch port feeding the AP trunks that VLAN, so confirm the uplink profile allows it. Measuring real throughput through the tagged SSID needs a physical access point on the wire; the controller-side configuration is complete here.
A self-hosted controller is only as safe as its last backup, and the container makes both jobs quick. Pull new images on a schedule and recreate the stack:
cd ~/unifi
docker compose pull
docker compose up -d
The LinuxServer image tracks the current stable Network Application while MongoDB stays pinned to the 8.0 line, so upgrades are predictable. Before any upgrade, take a controller backup from Settings, then System, then Backups, and keep a copy off the host. Because the whole state lives in the config volume plus the database under /unifi_data, recovering after a host rebuild is a matter of bringing the same stack up against that directory and importing the backup through the setup wizard.
