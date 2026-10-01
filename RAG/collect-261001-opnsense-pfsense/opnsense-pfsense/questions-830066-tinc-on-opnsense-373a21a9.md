---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-830066-tinc-on-opnsense-373a21a9
title: "questions-830066-tinc-on-opnsense-373a21a9"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-830066-tinc-on-opnsense-373a21a9.md
source_anchor: ""
source_lines: [1, 78]
sha256: cc869b78d370698cf751288f6be098b63cfb5d83120725d71cbb844bc4e64abb
---

# questions-830066-tinc-on-opnsense-373a21a9

Having this kind of network layout https://i.stack.imgur.com/NQXDl.jpg I can get it working using mode=switch in with tinc, but not using mode=routed

## Those are the network details

**On quantum**
those are the routes/config (hint, external IP has been obfuscated to 1.1.1.1)

```
netstat -rn4
Routing tables
Internet:
Destination        Gateway            Flags     Netif Expire
default            1.1.1.1     UGS      vtnet0
8.8.8.8            1.1.1.1     UGHS     vtnet0
10.1.5.0/24        link#2             U        vtnet1
10.1.5.1           link#2             UHS         lo0
10.1.6.0/24        tinc0              US        tinc0
127.0.0.1          link#4             UH          lo0
172.16.0.0/24      link#7             U         tinc0
172.16.0.5         link#7             UHS         lo0
1.1.1.0/29  link#1             U        vtnet0
1.1.1.1     link#1             UHS         lo0
ifconfig vtnet1 | grep inet
    inet 10.1.5.1 netmask 0xffffff00 broadcast 10.1.5.255
ifconfig tinc0 | grep inet
    inet 172.16.0.5 netmask 0xffffff00 broadcast 172.16.0.255
```
with this tinc host entry

```
Address=1.1.1.1 655
Subnet=10.1.5.0/24
Cipher=aes-256-cbc
Digest=sha256
-----BEGIN RSA PUBLIC KEY-----
...
-----END RSA PUBLIC KEY-----
```
**on Rocket** (hint, external ip has been obfuscated to 2.2.2.2)

```
netstat -rn4
Routing tables
Internet:
Destination        Gateway            Flags     Netif Expire
default            2.2.2.2      UGS      vtnet0
8.8.8.8            2.2.2.2      UGHS     vtnet0
10.1.5.0/24        tinc0              US        tinc0
10.1.6.0/24        link#2             U        vtnet1
10.1.6.1           link#2             UHS         lo0
2.2.2.0/29   link#1             U        vtnet0
2.2.2.2      link#1             UHS         lo0
127.0.0.1          link#4             UH          lo0
172.16.0.0/24      link#7             U         tinc0
172.16.0.6         link#7             UHS         lo0
ifconfig vtnet1 | grep inet
    inet 10.1.6.1 netmask 0xffffff00 broadcast 10.1.6.255
ifconfig tinc0 | grep inet
    inet 172.16.0.6 netmask 0xffffff00 broadcast 172.16.0.255
```
with this tinc host entry

```
Address=2.2.2.2 655
Subnet=10.1.6.0/24
Cipher=aes-256-cbc
Digest=sha256
-----BEGIN RSA PUBLIC KEY-----
...
-----END RSA PUBLIC KEY-----
```
## Further details

When doing a `ping -S 10.1.5.1 10.1.6.1` from quantum to rocket, I see on the FW that the ICMP package passes the firewall, but it never reaches rocket, no block or pass.

## Question

What is wrong with my setup, why is mode=routed not working?
