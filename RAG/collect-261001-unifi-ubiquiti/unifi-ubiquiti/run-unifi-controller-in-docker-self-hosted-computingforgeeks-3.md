---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/run-unifi-controller-in-docker-self-hosted-computingforgeeks-3
title: "run-unifi-controller-in-docker-self-hosted-computingforgeeks"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/run-unifi-controller-in-docker-self-hosted-computingforgeeks.md
source_anchor: ""
source_lines: [306, 348]
sha256: c3137069d432c5d4f400b6cf0b3cff44121a89f73358ab477c77cfc1393360de
---

# run-unifi-controller-in-docker-self-hosted-computingforgeeks

Create the VLAN first. Open **Settings**, then **Networks**, and add a network. Name it, set the **VLAN ID** (here `20` for an IoT segment), and either let the controller run DHCP for the subnet or mark it VLAN-only when an upstream router already serves it:

With the network saved, create the SSID that rides it. Open **Settings**, then **WiFi**, add a WiFi network, set the name and passphrase, then under the network option choose the VLAN you just created so client traffic is tagged with that VLAN ID:

Every access point the controller has adopted now broadcasts the SSID and tags its clients onto the VLAN, exactly as a UniFi gateway would. The tag only carries end to end when the switch port feeding the AP trunks that VLAN, so confirm the uplink profile allows it. Measuring real throughput through the tagged SSID needs a physical access point on the wire; the controller-side configuration is complete here.

## Keep the controller updated and backed up

A self-hosted controller is only as safe as its last backup, and the container makes both jobs quick. Pull new images on a schedule and recreate the stack:

```
cd ~/unifi
docker compose pull
docker compose up -d
```
The LinuxServer image tracks the current stable Network Application while MongoDB stays pinned to the 8.0 line, so upgrades are predictable. Before any upgrade, take a controller backup from **Settings**, then **System**, then **Backups**, and keep a copy off the host. Because the whole state lives in the config volume plus the database under `/unifi_data`, recovering after a host rebuild is a matter of bringing the same stack up against that directory and importing the backup through the setup wizard.

Telling people to set SELinux to permissive just to provide directory access is a very insecure approach. Why have you suggested this?

From experience most users are not that technical and troubleshooting SELinux related issues needs a bit of experience working with Linux.

For some reason when I try to start the unifi docker using docker-compose, I can’t get to the web interface like it is not using the ports and I followed the steps above exactly and it works fine if I use docker run.

Hello DP,

I have tested both on my end, it works fine, please share your

`docker ps`output. In case you are trying both methods, remember to stop and remove the container first.
What hardware setup is used to handle this? I mean, you need a router to receive the raw internet signal, and then send it to this computer, right? How about the outbound connections?

You just need UniFi devices to be on the same network as the controller deployment host.

where is the docker file ?

The compose file is one you create yourself: run vim docker-compose.yml and paste the contents shown in the guide, which has since been updated for the supported unifi-network-application image.

The docker compose file didn’t work for me either.

I see that you are using .env, but there is no step about creating the .env file.

Also, after bring up the container, it created a directory init-mongo.sh, but the directory is empty. Should this be a file and not a directory?

You hit two real problems and the guide has been rebuilt around the supported unifi-network-application image to fix both (the old unifi-controller image it used was retired by LinuxServer.io). The .env file now has its own creation step, and init-mongo.sh became a directory because Docker creates an empty directory when a mounted file does not exist yet. To recover on your machine: docker compose down, then sudo rm -rf init-mongo.sh and the MongoDB data directory, create init-mongo.sh as a regular file with the contents now shown in the guide, and run docker compose up -d again. The init script only executes against an empty database directory, which is why the data dir must go too.
