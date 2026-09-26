---
id: collect-260926-mikrotik/mikrotik/questions-1504136-configuring-routeros-on-a-crs312-4c8xg-aa43c43b
title: "questions-1504136-configuring-routeros-on-a-crs312-4c8xg-aa43c43b"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1504136-configuring-routeros-on-a-crs312-4c8xg-aa43c43b.md
source_anchor: ""
source_lines: [1, 24]
sha256: 96b1317c9962f5b9af9a9dec798768f2a184e29999936a9e7aad6d4e6af8995a
---

# questions-1504136-configuring-routeros-on-a-crs312-4c8xg-aa43c43b

Home 10 Gig Network Upgrade for CHEAP encouraged me to buy a CRS312-4C+8XG. I want to do basically the same thing Linus did (mainly set it up as a router) but this device, now that I have it, seems more enterprise-y than I had anticipated and the video skipped over the configuration. Mikrotik has a wiki but more about examples rather than being a manual.
Most of my network is GbE but I have a file server (Linux) connected to my desktop (Windows 10) via 10GbE over just a cat6A cable. I'd like to expand this 10GbE network and so I got the CRS312 to help with that.
By default, the CRS312 came configured as a bridge. Each of the ports seem to be individually configurable. It also seems that you can connect these ports together however you like; by default all the ports, including the management port, seem to be bridged together.
What I'd like to do:
- Separate the management port from the rest of the ports such that I can access the console from any port but the management port can't be used for regular traffic.
- Increase the MTU on the 10GbE ports to something compatible with my NICs (Thecus C10GTR). (The NIC Windows driver has the option "Jumbo frames" with a dropdown that lists Disabled, 4088, 9014, and 16128. I originally has this set to 16128 at both ends and that worked well, giving much better transfer speeds.)
Could somebody point me in the right direction?
EDIT: Thanks to grawity's answer, I was able to pin down more info about getting the MTU configured.
The bridge MTU values can't be set higher than the smallest MTU of the associated ports. The console port seems to be ether9. I found its number with /interface bridge port print (which told me it was item 12). I was able to remove that port from bridge with /interface bridge port remove 12. (I suspect there's a way to do this by name.)
Then I updated mtu and l2mtu on the relevant interfaces and bridge with:
/interface set [find max-l2mtu>9000] l2mtu=9092
/interface set [find max-l2mtu>9000] mtu=9000
/interface bridge set bridge l2mtu=9092
/interface bridge set bridge mtu=9000
92 is the default overhead the device uses so I decided to stay with that. I'm not sure what impact this really has. As far as I can tell, MTU is usually the L3 packet size and L2 adds 18 bytes. See Jumbo Frame/Bandwidth Efficiency.
How to Change MTU value on Linux helped me figure out what the different devices and system mean when they list MTU. The process uses ping to detect a usable packet size by disabling packet fragmentation and changing the packet size of each ping in a process of elimination, then add 28 to get the L3 MTU.
Linux: ping -c 3 -M do –s 8972 192.168.88.100
Windows: ping -f -l 8972 -4 192.168.88.101
Setting the MTU:
Linux: sudo ifconfig eth1 mtu 9000
or: sudo ip link set eth1 mtu 9000
Windows: netsh interface ipv4 set subinterface "Interface name" mtu=9000 store=persistent
In Windows, you can list your interfaces with: netsh interface ipv4 show subinterfaces. You want the name in the "Interface" column.
Doing this via command line in Windows is necessary as "Jumbo Frames", in the advanced adapter properties, for the TN9210 driver, only lists a few values and they aren't L3 MTU values.
