---
id: collect-260926-mikrotik/mikrotik/vlan-for-wlan-and-ethernet-in-crs1xx
title: "Recommended reading"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["asic", "ethernet", "throughput"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/vlan-for-wlan-and-ethernet-in-crs1xx.md
source_anchor: ""
source_lines: [1, 155]
sha256: 9f72c3f58560803fd16d99edcc5fe7ba2da44ca8aeadedaad691175e9269624e
---

# Recommended reading

Hello everyone,

First of all, I am a self-taught Mikrotik user so forgive me for any silly mistakes, comment or description.

I have the following configuration of CRS109 and CRS125 at my disposal, with the following configuration.


I would like to segregate the traffic using VLAN, which described as red VLAN (including WLAN in the CRS109), green VLAN and purple VLAN with blue as Trunk port and yellow as admin port. I implemented the VLAN successfully, using the configuration greatly described by the following guide. http://forum.mikrotik.com/t/using-routeros-to-vlan-your-network/126489/1

However, reading various guides on CRS1xx, I understand that implement VLAN inside the CRS1xx is best by utilizing switch VLAN rather than bridge VLAN as the bridge VLAN consumes higher CPU resources and might result in lower throughput. I have implemented the switch VLAN successfully if only for the Ethernet ports, but unfortunately I cannot make it work with the WLAN. I understand that the ethernet ports and the WLAN are on different hardware chips, but I would like to see if I can create VLAN in the above configuration by combining between bridge VLAN and switch VLAN? I tried to search but cannot find any resource that perfectly describe my use case. Will this configuration of bridge and switch VLAN be possible? Will it gives a better throughput and lower CPU utilization, or will it be very minimum improvement? Or is it because both hardware are not designed to be used in such configuration?

Appreciate very much for any input and advise given.

             
            
           
          
            
            
              What is the purpose of the trunk link between the two switches if the only vlan being used on the CRS-109 is the “RED” vlan?

In other words, for your application, would a “dumb vlan-transparent switch” work in place of the CRS-109?  I don’t have one, but I would guess that the CRS109 can be configured as a vlan transparent switch.

The block diagram for the CRS109 shows that the SFP (what I assume you mean by the WAN port) is connected to the same 26 port non-blocking switch ASIC, so there isn’t any obvious to me reason that you shouldn’t be able to switch traffic at wire speed between the SFP and the rest of the copper RJ45 ports.  In fact, it may be the same ASIC as in the CRS125 (see CRS125 block diagram).  This seems to confirm that; both use the QCA-8513L.  So what you can do on the CRS125 you should be able to do on the CRS109, unless MikroTik has done something to limit what can be done, and that seems unlikely to me.

Have you watched Mikrotik VLANs - CRS1xx & CRS2xx - Mikrotik Tutorial ?

And see this @anav post.  You have provided a diagram already, but no configs or use case.

Also see section P of @anav’s New User Pathway To Config Success for more links to documentation.

             
            
           
          
            
            
              I never added CRS1xx support to the article, but if I did, it would look something like this.

**CRS1xx VLAN Example**




```
###############################################################################
# Recommended reading
# https://wiki.mikrotik.com/wiki/Manual:Basic_VLAN_switching
#
# Notes: Start with a reset (/system reset-configuration)
#
# Based on: http://forum.mikrotik.com/t/using-routeros-to-vlan-your-network/126489/1
###############################################################################
#######################################
# Naming
#######################################
# name the device being configured
/system identity set name="CRS1xx_Switch"
#######################################
# VLAN Overview
#######################################
# 10 = BLUE
# 20 = GREEN
# 30 = RED
# 99 = BASE (MGMT) VLAN
#######################################
# Bridge
#######################################
# create one bridge
/interface bridge add name=BR1 protocol-mode=none
# add "all" ports to this one bridge
/interface bridge port
add bridge=BR1 interface=ether1
add bridge=BR1 interface=ether2
add bridge=BR1 interface=ether3
add bridge=BR1 interface=ether4
# and so on until you get to 24 ...
#######################################
#
# -- Access Ports --
#
#######################################
# ingress behavior, egress dynamically handled
/interface ethernet switch ingress-vlan-translation
add customer-vid=0 new-customer-vid=10 ports=ether2
add customer-vid=0 new-customer-vid=20 ports=ether3
add customer-vid=0 new-customer-vid=30 ports=ether4
#######################################
#
# -- Trunk Ports --
#
#######################################
# ingress behavior
# L2 switching only, Bridge (aka switch1-cpu) not needed as tagged member (except for BASE_VLAN)
/interface ethernet switch vlan
add ports=ether1,ether2 vlan-id=10
add ports=ether1,ether3 vlan-id=20
add ports=ether1,ether4 vlan-id=30
add ports=switch1-cpu,ether1 vlan-id=99
# egress behavior
# L2 switching only, Bridge (aka switch1-cpu) not needed as tagged member (except for BASE_VLAN)
/interface ethernet switch egress-vlan-tag
add tagged-ports=ether1 vlan-id=10
add tagged-ports=ether1 vlan-id=20
add tagged-ports=ether1 vlan-id=30
add tagged-ports=switch1-cpu,ether1 vlan-id=99
#######################################
# VLAN Security
#######################################
# drop traffic that does not follow the above port layout
/interface ethernet switch set forward-unknown-vlan=no
#######################################
# IP Addressing & Routing
#######################################
# LAN facing Switch's IP address on a BASE_VLAN
/interface vlan add interface=BR1 name=BASE_VLAN vlan-id=99
/ip address add address=192.168.0.2/24 interface=BASE_VLAN network=192.168.0.0
# The Router's IP this switch will use
/ip route add distance=1 gateway=192.168.0.1
#######################################
# MAC Server settings
#######################################
# Ensure only visibility and availability from BASE_VLAN, the MGMT network
/interface list add name=BASE
/interface list member add interface=BASE_VLAN list=BASE
/ip neighbor discovery-settings set discover-interface-list=BASE
/tool mac-server mac-winbox set allowed-interface-list=BASE
/tool mac-server set allowed-interface-list=BASE
```

             
            
           
          
            
            
              
Since you have now done most of the work, why not add it to Using RouterOS to VLAN your network, or at least add a link to your post so someone could more easily find it in the future?

             
            
           
          
            
            
              Because He knows Ive already added it to my switch chip section  para P…

             
            
           
          
            
            
              
Well, I don’t really recommend using CRS1xx switches, but I understand some may have to.
