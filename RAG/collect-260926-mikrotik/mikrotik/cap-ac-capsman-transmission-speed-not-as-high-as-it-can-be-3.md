---
id: collect-260926-mikrotik/mikrotik/cap-ac-capsman-transmission-speed-not-as-high-as-it-can-be-3
title: "nov/16/2022 11:00:52 by RouterOS 7.6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/cap-ac-capsman-transmission-speed-not-as-high-as-it-can-be.md
source_anchor: ""
source_lines: [298, 372]
sha256: 9711c97f16e4d061de581c12ce9af9d7b556275e61bfce5cc2a11f1a3f778acf
---

# nov/16/2022 11:00:52 by RouterOS 7.6

hAP ac3 (172.22.99.254) as router, DHCP and CAPsMAN server.

hEX (172.22.99.253) as switch #1

cAP ac (172.22.99.252) as AP #1

RB951G (172.22.99.251) as switch #2

wAP R ac (172.22.99.250)  as AP #2


Access points are currently configured to connect to CAPsMAN via L3. The “local forwarding” mode does not work with this network scheme. Switching access points to search for the CAPsMAN server via L2 does not change anything and the “local forwarding” mode also remains inoperative.

172.22.99.253_19-11-2022_v7.rsc (2.21 KB)

172.22.99.254_19-11-2022_v7.rsc (9.64 KB)

172.22.99.250_19-11-2022_v7.rsc (4 KB)

172.22.99.251_19-11-2022_v7.rsc (2.32 KB)

172.22.99.252_19-11-2022_v7.rsc (2.98 KB)

             
            
           
          
            
            
              @BrateloSlava

1. CAPsMAN forwarding
On a CAP, the wlan interfaces are not in the bridge. Otherwise loop messages on the capsman controller and RSTP triggering disabling the ports. Since the same MAC comes from different interfaces.
2. Local forwarding
 On the CAP, the wlan interfaces are in the bridge.
 You need to look in the logs sometimes false RSTP triggering happens and you have to disable it on the bridge of the capsman manager
For experiments it is better to use provisioning add action=create-enabled instead of action=create-dynamic-enabled

PS If you want we can contact you, I will help you tweak some wifi settings.

PSS When the access points are powered from the CRS328 then it is more convenient to configure the capsman on it.

             
            
           
          
            
            
              @**Ca6ko**

As I wrote earlier, in the case when all access points are connected to one router or one switch, there are no problems.

The problem arises when there is a cascade of switches. Moreover, this chain of switches has branchings. And somewhere in these chains, access points are connected.

A wireless network client roams between points on such a network. From the point of view of the CAPsMAN server - the MAC address of the device of such a client appears in “completely unexpected” places on the network. It is possible, that I am explaining my point of view in terms that are not entirely correct, but as long as the client is near one point and talking through the messenger, there are no problems. Moved to another access point - the connection was interrupted. Rather, there is a reconnection in the messenger.

In the case when the forwarding mode via CAPsMAN is used, this situation is not observed.

That’s why today I put together the “simplest network” at home, in which there is a problem when using “local forwarding”. A router to which access points are connected via separate switches.

In this scheme, there are no difficulties, the “simplest” bridges, a single address space, no VLAN.

I gave the configuration of all devices. RTSP on bridges I turned on and off. I tried adding wireless interfaces to bridges before connecting to CAPsMAN. No difference. In the “local forwarding” mode, wireless interfaces are dynamically added with the necessary parameters to the bridges on the access points. When using CAPsMAN forwarding, these local wireless interfaces on access points are automatically disabled from local bridges.

P.S. Thank you for your willingness to help with a personal consultation. But for me it is important to understand my mistakes myself. Perhaps there is some step-by-step instruction for configuring CAPsMAN in a situation where access points are “scattered” over the network. At work, I very often have to make networks “work fine”, that others have designed and installed.

PSS. About action=create-enabled instead of action=create-dynamic-enabled. In my current configuration, access points are generally described statically, with binding to MAC addresses. Therefore, new elements in the “CAP interfaces” section are not created at all.

             
            
           
          
            
            
              @BrateloSlava: I’ve had a (quick) look at config files … nothing really hits me as wrong, the only thing which might interfere are arp settings on DHCP server. There have been some questions on this forum about what exactly “add-arp=yes” option does and it seems to have some security implications … which can kick in when a device needs to move between bridge ports. I’d set that option to default (which is “add-arp=no”) and see if it helps.
