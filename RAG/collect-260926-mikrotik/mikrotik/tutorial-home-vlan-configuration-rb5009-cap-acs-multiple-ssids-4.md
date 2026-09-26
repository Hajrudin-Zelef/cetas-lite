---
id: collect-260926-mikrotik/mikrotik/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids-4
title: "1. Motivation"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/tutorial-home-vlan-configuration-rb5009-cap-acs-multiple-ssids.md
source_anchor: ""
source_lines: [444, 482]
sha256: 2a3e863719a7b9058267b1a53588aa53985a75e114a34e2747f3304453187807
---

# 1. Motivation

```
/ip firewall
# Configure NAT for internet access
nat add chain=srcnat out-interface=pppoe-out action=masquerade comment="NAT for internet access"
/ip firewall filter
# Traffic into the router
add chain=input   action=accept connection-state=established,related,untracked                                       comment="Allow established connections"
add chain=input   action=drop   connection-state=invalid                                                             comment="Drop invalid packets"
add chain=input   action=accept protocol=icmp                                                                        comment="Allow ICMP"
add chain=input   action=accept dst-address=127.0.0.1                                                                comment="Allow local loopback for CAPsMAN"
add chain=input   action=accept in-interface-list=!wan dst-port=53 protocol=udp                                      comment="Allow LAN DNS queries"
add chain=input   action=accept in-interface-list=!wan dst-port=53 protocol=tcp                                      comment="Allow LAN DNS queries"
add chain=input   action=accept in-interface-list=management                                                         comment="Allow full access to the management interfaces"
add chain=input   action=accept in-interface=vlan10-owner src-address-list=owner-webfig dst-port=80,443 protocol=tcp comment="Allow restricted webfig access to the router"
add chain=input   action=drop                                                                                        comment="Drop all other inputs"
# Traffic forwarded through the router
# Basics
add chain=forward action=fasttrack-connection connection-state=established,related hw-offload=yes comment="Fasttrack established connections"
add chain=forward action=accept               connection-state=established,related,untracked      comment="Allow established connections"
add chain=forward action=drop                 connection-state=invalid                            comment="Drop invalid packets"
# Inter-VLAN communication
add chain=forward action=accept connection-state=new in-interface=vlan10-owner out-interface=vlan20-iot        comment="Allow Owner -> IoT traffic"
add chain=forward action=accept connection-state=new in-interface=vlan10-owner out-interface=vlan99-management comment="Allow Owner -> Management traffic"
# Internet access
add chain=forward action=accept connection-state=new in-interface-list=management out-interface-list=wan comment="Allow internet access for management VLAN"
add chain=forward action=accept connection-state=new in-interface-list=vlan       out-interface-list=wan comment="Allow internet access for all other VLANs"
# Port forwarding
add chain=forward action=accept connection-nat-state=dstnat comment=“Allow port forwarding”
# Drop all other forwarded traffic
add chain=forward action=drop comment="Drop all other forwarded traffic"
```

 Customization

- The firewall configuration above is opinionated and may not suit your needs, so please review it carefully and adjust it to suit your needs.
- You may for example not plan to use port forwarding, in which case you can remove the "Allow port forwarding" filter.


*The next post will continue the guide at "5. Wi-Fi Access Point Configuration".*
