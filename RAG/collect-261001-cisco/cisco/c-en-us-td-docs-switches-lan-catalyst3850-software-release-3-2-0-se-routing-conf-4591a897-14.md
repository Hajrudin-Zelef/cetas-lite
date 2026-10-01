---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-14
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: ["1993-01-01"]
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [751, 792]
sha256: 79262bcbe0a8fe9af7cdab7e81d8f1c2c8587f8e2c366d3761433be63744f4ad
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | interface interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 3 | ip 				  bandwidth-percent eigrp  				percent Example:  Device(config-if)# ip bandwidth-percent eigrp 60  | (Optional) Configures the percentage of bandwidth that can be used by EIGRP on an interface. The default is 50 percent. | 
| Step 4 | ip 				  summary-address eigrp  				autonomous-system-number 				  address mask Example:  Device(config-if)# ip summary-address eigrp 109 192.161.0.0 255.255.0.0  | (Optional) Configures a summary aggregate address for a specified interface (not usually necessary if auto-summary is enabled). | 
| Step 5 | ip hello-interval 				  eigrp  				autonomous-system-number 				  seconds Example:  Device(config-if)# ip hello-interval eigrp 109 10  | (Optional) Change the hello time interval for an EIGRP routing process. The range is 1 to 65535 seconds. The default is 60 seconds for low-speed NBMA networks and 5 seconds for all other networks. | 
| Step 6 | ip hold-time 				  eigrp  				autonomous-system-number 				  seconds Example:  Device(config-if)# ip hold-time eigrp 109 40  | (Optional) Change the hold time interval for an EIGRP routing process. The range is 1 to 65535 seconds. The default is 180 seconds for low-speed NBMA networks and 15 seconds for all other networks. | 
| Step 7 | no ip 				  split-horizon eigrp  				autonomous-system-number Example:  Device(config-if)# no ip split-horizon eigrp 109  | (Optional) Disables split horizon to allow route information to be advertised by a router out any interface from which that information originated. | 
| Step 8 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 9 | show ip eigrp 				  interface Example:  Device# show ip eigrp interface  | Displays which interfaces EIGRP is active on and information about EIGRP relating to those interfaces. | 
| Step 10 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Caution | Do not adjust the hold time without consulting Cisco technical support. | 
Configuring EIGRP Route Authentication
EIGRP route authentication provides MD5 authentication of routing updates from the EIGRP routing protocol to prevent the introduction of unauthorized or false routing messages from unapproved sources.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | interface interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 3 | ip authentication 				  mode eigrp  				autonomous-system  				md5 Example:  Device(config-if)# ip authentication mode eigrp 104 md5  | Enables MD5 authentication in IP EIGRP packets. | 
| Step 4 | ip authentication 				  key-chain eigrp autonomous-system 				  key-chain Example:  Device(config-if)# ip authentication key-chain eigrp 105 chain1  | Enables authentication of IP EIGRP packets. | 
| Step 5 | exit Example:  Device(config-if)# exit  | Returns to global configuration mode. | 
| Step 6 | key 				  chain name-of-chain Example:  Device(config)# key chain chain1  | Identify a key chain and enter key-chain configuration mode. Match the name configured in Step 4. | 
| Step 7 | key  				number Example:  Device(config-keychain)# key 1  | In key-chain configuration mode, identify the key number. | 
| Step 8 | key-string  				text Example:  Device(config-keychain-key)# key-string key1  | In key-chain key configuration mode, identify the key string. | 
| Step 9 | accept-lifetime  				start-time {infinite \|  				end-time \|  				duration  				seconds} Example:  Device(config-keychain-key)# accept-lifetime 13:30:00 Jan 25 2011 duration 7200  | (Optional) Specifies the time period during which the key can be received. The start-time and end-time syntax can be either hh:mm:ss Month date year or hh:mm:ss date Month year. The default is forever with the default start-time and the earliest acceptable date as January 1, 1993. The default end-time and duration is infinite. | 
| Step 10 | send-lifetime  				start-time {infinite \|  				end-time \|  				duration  				seconds} Example:  Device(config-keychain-key)# send-lifetime 14:00:00 Jan 25 2011 duration 3600  | (Optional) Specifies the time period during which the key can be sent. The start-time and end-time syntax can be either hh:mm:ss Month date year or hh:mm:ss date Month year. The default is forever with the default start-time and the earliest acceptable date as January 1, 1993. The default end-time and duration is infinite. | 
| Step 11 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 12 | show key 				  chain Example:  Device# show key chain  | Displays authentication key information. | 
| Step 13 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring and Maintaining EIGRP
You can delete neighbors from the neighbor table. You can also display various EIGRP routing statistics. The table given below lists the privileged EXEC commands for deleting neighbors and displaying statistics. For explanations of fields in the resulting display, see the Cisco IOS IP Command Reference, Volume 2 of 3: Routing Protocols, Release 12.4.
| Table 8  IP EIGRP Clear and Show 		  Commands |  | 
|---|---|
| clear ip eigrp neighbors [if-address \| interface] | Deletes neighbors from the neighbor table. | 
| show ip eigrp interface [interface] [as number] | Displays information about interfaces configured for EIGRP. | 
| show ip eigrp neighbors [type-number] | Displays EIGRP discovered neighbors. | 
| show ip eigrp topology [autonomous-system-number] \| [[ip-address] mask]] | Displays the EIGRP topology table for a given process. | 
| show ip eigrp traffic [autonomous-system-number] | Displays the number of packets sent and received for all or a specified EIGRP process. | 
Information About BGP
The Border Gateway Protocol (BGP) is an exterior gateway protocol used to set up an interdomain routing system that guarantees the loop-free exchange of routing information between autonomous systems. Autonomous systems are made up of routers that operate under the same administration and that run Interior Gateway Protocols (IGPs), such as RIP or OSPF, within their boundaries and that interconnect by using an Exterior Gateway Protocol (EGP). BGP Version 4 is the standard EGP for interdomain routing in the Internet. The protocol is defined in RFCs 1163, 1267, and 1771. You can find detailed information about BGP in Internet Routing Architectures, published by Cisco Press, and in the “Configuring BGP” chapter in the Cisco IP and IP Routing Configuration Guide.
For details about BGP commands and keywords, see the “IP Routing Protocols” part of the Cisco IOS IP Command Reference, Volume 2 of 3: Routing Protocols .
BGP Network Topology
Routers that belong to the same autonomous system (AS) and that exchange BGP updates run internal BGP (IBGP), and routers that belong to different autonomous systems and that exchange BGP updates run external BGP (EBGP). Most configuration commands are the same for configuring EBGP and IBGP. The difference is that the routing updates are exchanged either between autonomous systems (EBGP) or within an AS (IBGP). The figure given below shows a network that is running both EBGP and IBGP.
