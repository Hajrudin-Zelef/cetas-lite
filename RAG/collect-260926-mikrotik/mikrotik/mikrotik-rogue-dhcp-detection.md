---
id: collect-260926-mikrotik/mikrotik/mikrotik-rogue-dhcp-detection
title: "mikrotik-rogue-dhcp-detection"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/mikrotik-rogue-dhcp-detection.md
source_anchor: ""
source_lines: [1, 53]
sha256: 3121288774725ade8d80967a69246bef17ce7e5bfab05c3c67f47d892d5a847e
---

# mikrotik-rogue-dhcp-detection

Mikrotik actually has a rogue detection service you can configure, but as we have found running apartment complexes, it can give you false positives. To combat this, I’ve come up with a new method.

First configure syslog exporting with the DHCP service dumping. Next enable the dhcp client to run on all your inside facing interfaces. I’ve got my Cacti syslog server set to match “dhcp,info,debug dhcp-client%got ip address”, which is the message sent when the dhcp client receives an IP address. Once the syslog server processes the message it sends us an alert.

Add dhcp-client to each interface, be sure to disable default route, peer dns and peer ntp.

```
/ ip dhcp-client 
add interface=vlan10 comment="" disabled=no 
add interface=ether2 comment="" disabled=no 
```
This script will need to be scheduled to run around every hour or so. It will release the dhcp reservation on your interfaces. Otherwise it won’t attempt to pull a new address until the old allocation has expired, which can be up to a year. It loops through releasing from all your interfaces.

```
:log info ("dhcp detect release")
:for e from=0 to=40 do={
/ip dhcp-client release ($e)
}
```
Once you get an alert from your syslog server, you log into the Mikrotik and issue the:

```
 
/ip dhcp-client print detail
```
You’ll get the following:

Flags: X – disabled, I – invalid

 0   interface=vlan10 status=bound address=192.168.1.100/24

     gateway=192.168.1.1 dhcp-server=192.168.1.1

     primary-dns=209.189.224.45 secondary-dns=209.189.224.40

     expires-after=21h13m2s


Take the dhcp-server address and use it below:

`/ip arp print where address="192.168.1.1"`
You will get the following result:

Flags: X – disabled, I – invalid, H – DHCP, D – dynamic

 #   ADDRESS         MAC-ADDRESS       INTERFACE

 0 D 192.168.0.253   00:08:74:4B:7F:BC ether1


Now you track this guy down and shut him off at the switch port, or if you are using mac-track in cacti, you simply look him up, connect to the proper switch and kill him. You could also use this in conjunction with the standard rogue detection service to more quickly find the MAC address.

## Leave a Reply
