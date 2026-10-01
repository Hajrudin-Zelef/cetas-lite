---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/unifi-layer-3-adoption-access-points-bd5278b7-2
title: "unifi-layer-3-adoption-access-points-bd5278b7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/unifi-layer-3-adoption-access-points-bd5278b7.md
source_anchor: ""
source_lines: [61, 71]
sha256: 4615c7bff8f975e914b79d89ea5e2db2b1d5b0da20a8e915af91db4b75c3dbe7
---

# unifi-layer-3-adoption-access-points-bd5278b7

With Layer 2 adoption, the device and the controller are on the same subnet, so the device normally appears automatically as “Pending Adoption”. With Layer 3 adoption there is a routing boundary in between – another VLAN, another site, or a hosted controller – so the device has to be told explicitly where the controller is.
What is the set-inform command?
set-inform http://<ip-or-hostname>:8080/inform, issued over SSH directly on the UniFi device. The device then appears in the remote controller and can be adopted there. Note that it is HTTP, not HTTPS.
Which SSH credentials apply before adoption?
Prior to setup and adoption, the default credentials are ui / ui, or ubnt / ubnt on older devices. SSH is enabled by default on UniFi network devices. After adoption, the controller assigns a random string of characters as the credentials.
Which ports do I need to open for adoption?
TCP 8080 and UDP 10001 must be open between the UniFi host and the UniFi devices, across all gateways, firewalls and antivirus software. For Layer 3 adoption, unrestricted connectivity over TCP 8080 is the critical part. UniFi additionally uses TCP 8443 for the interface and API, and UDP 3478 for STUN.
Which devices support zero-touch provisioning?
According to Ubiquiti, ZTP is currently supported by U7 Pro Max access points, with more devices to follow. Note that the ZTP code is only preserved if the device is not reset with the hardware reset button – always factory reset through the UniFi management interface instead.
Why do my access points lose the connection after adoption?
A common cause is an inform host the devices cannot resolve or reach from their network. The UniFi application lets you override the inform host with a specific hostname or IP, and name resolution for it has to be reliable. Editing the configuration files on the devices directly does not help – the controller overwrites those changes.
