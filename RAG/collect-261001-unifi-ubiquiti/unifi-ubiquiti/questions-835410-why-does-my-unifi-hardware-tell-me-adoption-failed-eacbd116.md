---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-835410-why-does-my-unifi-hardware-tell-me-adoption-failed-eacbd116
title: "questions-835410-why-does-my-unifi-hardware-tell-me-adoption-failed-eacbd116"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-835410-why-does-my-unifi-hardware-tell-me-adoption-failed-eacbd116.md
source_anchor: ""
source_lines: [1, 11]
sha256: baea98c97ed3cd644ca2c3acde97467b42d4d84aaa6732593b828fe7cd14ae94
---

# questions-835410-why-does-my-unifi-hardware-tell-me-adoption-failed-eacbd116

A common issue is that the controller and the target device don't share the same network. During adoption, the controller will attempt to establish a direct connection to the device, this may fail.
Connecting to the device via SSH and checking /var/log/messages can also be very helpful.
Also, avoid using the root account for anything. Especially device authentication. It can really mess up your network, as the controller might try to create the named user on new devices. This will fail, as root already exists.
Configuring a different user name for Device Authentication in the controller settings can magically resolve weird adoption errors.
When you're deploying a USG in your network, you can end up in a situation where your controller tries to configure the site for a certain network, but your USG still provides an outdated configuration via DHCP. The result is that the USG is in a different network than the rest of your devices. This will make it impossible to adopt the USG.
In such a scenario, you will have to connect to the USG via SSH and add an additional IP address to the LAN interface you're using. This will allow the controller to communicate with your USG, adopt and deploy as configured.
The key command sequence here is:
configure
set interfaces ethernet eth0 address 192.168.1.2/24
commit
exit
