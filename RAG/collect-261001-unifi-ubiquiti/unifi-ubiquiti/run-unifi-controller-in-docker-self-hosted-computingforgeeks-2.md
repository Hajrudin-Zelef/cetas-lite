---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/run-unifi-controller-in-docker-self-hosted-computingforgeeks-2
title: "run-unifi-controller-in-docker-self-hosted-computingforgeeks"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/run-unifi-controller-in-docker-self-hosted-computingforgeeks.md
source_anchor: ""
source_lines: [120, 305]
sha256: abe180d07404fe2389ede855303e5ac175a546f90e2fa0eaa28ec511e46bc260
---

# run-unifi-controller-in-docker-self-hosted-computingforgeeks

```
docker run -d --name=unifi-network-application --net unifi-net 
  -e PUID=1000 
  -e PGID=1000 
  -e TZ=Etc/UTC 
  -e MONGO_USER=unifi 
  -e MONGO_PASS=Str0ngUnifiPass 
  -e MONGO_HOST=unifi-db 
  -e MONGO_PORT=27017 
  -e MONGO_DBNAME=unifi 
  -e MONGO_AUTHSOURCE=admin 
  -e MEM_LIMIT=1024 
  -e MEM_STARTUP=1024 
  -p 8443:8443 
  -p 3478:3478/udp 
  -p 10001:10001/udp 
  -p 8080:8080 
  -v /unifi_data/config:/config 
  --restart unless-stopped 
  lscr.io/linuxserver/unifi-network-application:latest
```
First startup takes a few minutes while the Java application initialises. Poll the status endpoint until it reports up:

`curl -sk https://localhost:8443/status`
A fully started controller answers with its version:

`{"meta":{"rc":"ok","up":true,"server_version":"10.4.57","uuid":"016f58b1-2e5c-4fe4-97ed-9bd3229ac209"},"data":[]}`
## Run UniFi Controller using Docker Compose (Recommended)

You can also run the UniFi Controller using Docker Compose. First, ensure that Docker compose is installed on your system.

The Compose file wires the same two containers together and reads the secrets from an `.env` file sitting next to it, so create that first. Make a project directory and open the environment file:

```
mkdir -p ~/unifi && cd ~/unifi
vim .env
```
Define the two database passwords in it, using your own values:

```
MONGO_ROOT_PASS=Str0ngRootPass
MONGO_PASS=Str0ngUnifiPass
```
Then create the Compose file:

`vim docker-compose.yml`
It defines the database, with the init script mounted, and the controller:

```
---
services:
  unifi-db:
    image: mongo:8.0
    container_name: unifi-db
    environment:
      - MONGO_INITDB_ROOT_USERNAME=root
      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_ROOT_PASS}
      - MONGO_USER=unifi
      - MONGO_PASS=${MONGO_PASS}
      - MONGO_DBNAME=unifi
      - MONGO_AUTHSOURCE=admin
    volumes:
      - /unifi_data/db:/data/db
      - /unifi_data/init-mongo.sh:/docker-entrypoint-initdb.d/init-mongo.sh:ro
    restart: unless-stopped
  unifi-network-application:
    image: lscr.io/linuxserver/unifi-network-application:latest
    container_name: unifi-network-application
    depends_on:
      - unifi-db
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Etc/UTC
      - MONGO_USER=unifi
      - MONGO_PASS=${MONGO_PASS}
      - MONGO_HOST=unifi-db
      - MONGO_PORT=27017
      - MONGO_DBNAME=unifi
      - MONGO_AUTHSOURCE=admin
      - MEM_LIMIT=1024
      - MEM_STARTUP=1024
    volumes:
      - /unifi_data/config:/config
    ports:
      - 8443:8443
      - 3478:3478/udp
      - 10001:10001/udp
      - 8080:8080
    restart: unless-stopped
```
Save the file and start the stack. Recent Docker ships Compose as a plugin, so the command is `docker compose`, not the old `docker-compose` binary:

`docker compose up -d`
Both containers come up:

```
 Container unifi-db  Started
 Container unifi-network-application  Started
```
Whichever way you started it, verify both containers are running:

`docker ps`
The controller and its database both show Up:

```
NAMES                       IMAGE                                                  STATUS          PORTS
unifi-network-application   lscr.io/linuxserver/unifi-network-application:latest   Up 2 minutes    0.0.0.0:8443->8443/tcp, 0.0.0.0:8080->8080/tcp, 0.0.0.0:3478->3478/udp, 0.0.0.0:10001->10001/udp
unifi-db                    mongo:8.0                                              Up 2 minutes    27017/tcp
```
## Fix the two common startup failures

Both of these came up in testing and in reader reports, and both leave the web UI unreachable.

### init-mongo.sh was created as an empty directory

If the init script does not exist as a file when the containers first start, Docker silently creates an empty directory at that path. The database then comes up without the `unifi` user and the controller can never connect. Check what you have:

`ls -la /unifi_data/init-mongo.sh`
If that lists a directory, tear the stack down, remove the directory together with the half-initialised database, recreate the script as a file per the persistent volumes section, and start again:

```
docker compose down
sudo rm -rf /unifi_data/init-mongo.sh /unifi_data/db
sudo mkdir -p /unifi_data/db
```
### Error: Cannot run program “bin/mongod”

This appears in `/unifi_data/config/logs/server.log` when the config volume still holds data from the retired `unifi-controller` image, which bundled its own database. The new image sees the old configuration and tries to launch a mongod binary it does not ship:

```
ERROR system - unable to exec
java.io.IOException: Cannot run program "bin/mongod" (in directory "/usr/lib/unifi"): Exec failed, error: 2 (No such file or directory)
```
Export a backup from the old controller first (Settings > System > Backups) if you need the configuration, then stop the container, clear `/unifi_data/config`, start fresh, and restore the backup through the setup wizard.

## Access UniFi Controller Web UI

Now access the UniFi Controller web UI using the URL https://10.0.1.50:8443

Set the name of the application and proceed to sign in using your Ubiquiti account.

Configure the network.

The devices are not available since the application is running in a docker container. So skip and configure this later.

This too can be skipped and set later.

Review the configurations.

The configurations will be made as shown.

Once complete you will see the below dashboard.

## Adopt a UniFi access point into the self-hosted controller

Adoption is where a self-hosted controller behaves differently from a UniFi console. A factory access point looks for a controller on its local layer 2 segment, then keeps reporting in to whatever inform URL it is handed. A controller in a container answers on the Docker host’s LAN address, not on the internal bridge IP a device would otherwise learn, so the fix is to advertise an address the devices can actually reach.

Set that address once. Open **Settings**, then **System**, scroll to **Advanced**, switch **Inform Host** to manual, and enter the LAN IP or DNS name of the Docker host (here `10.0.1.50`):

Apply the change. Any access point on the same subnet as the host is now discovered automatically and shows up under **Devices** ready to adopt.

### Adopt manually with set-inform

When a device never appears, point it at the controller by hand over SSH. This is the reliable path when the AP sits on a different subnet from the container:

```
ssh [email protected]
set-inform http://10.0.1.50:8080/inform
```
Replace `10.0.1.60` with the device address and `10.0.1.50` with the Docker host. A factory device signs in with `ubnt` / `ubnt`. Run `set-inform` once so the device shows as *Pending Adoption*, click **Adopt** in the controller, then run it a second time after the device provisions so it locks onto the controller for good.

### Adopt across subnets (layer 3)

Layer 2 discovery does not cross a router, so a controller in one VLAN and access points in another need one of the standard hand-off methods:

- **DHCP option 43** on the device subnet, encoding the controller IP as`01:04:<controller-ip-in-hex>` .
- A **DNS record named `unifi`** that resolves to the controller, which factory devices query on boot.
- The `set-inform` command above, run once per device.

The inform host set earlier is what makes all three work: the device reaches the controller, and the controller replies with an inform URL it can keep using.

## Create a VLAN network and a tagged SSID

The payoff of running the controller yourself is the same segmented network a hardware UniFi console would build: a VLAN for a class of devices and a wireless network whose traffic is tagged onto it. Both are defined once in the controller and pushed to every adopted gateway, switch, and access point.

