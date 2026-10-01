---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-8
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [431, 491]
sha256: 9e0c0efb96082641fc8097847879e40fb69435322b4725b1933ee44701bb4358
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

RIP Version 1 does not support authentication. If you are sending and receiving RIP Version 2 packets, you can enable RIP authentication on an interface. The key chain specifies the set of keys that can be used on the interface. If a key chain is not configured, no authentication is performed, not even the default.
The Device supports two modes of authentication on interfaces for which RIP authentication is enabled: plain text and MD5. The default is plain text.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the interface to configure. | 
| Step 4 | ip rip 				  authentication key-chain  				name-of-chain Example:  Device(config-if)# ip rip authentication key-chain trees  | Enables RIP authentication. | 
| Step 5 | ip rip authentication mode {text \|  				md5} Example:  Device(config-if)# ip rip authentication mode md5  | Configures the interface to use plain text authentication (the default) or MD5 digest authentication. | 
| Step 6 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 7 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 8 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Summary Addresses and Split Horizon
| Note | In general, disabling split horizon is not recommended unless you are certain that your application requires it to properly advertise routes. | 
If you want to configure an interface running RIP to advertise a summarized local IP address pool on a network access server for dial-up clients, use the ip summary-address rip interface configuration command.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | ip address  				ip-address 				  subnet-mask Example:  Device(config-if)# ip address 10.1.1.10 255.255.255.0  | Configures the IP address and IP subnet. | 
| Step 5 | ip 				  summary-address rip ip address  				ip-network mask Example:  Device(config-if)# ip summary-address rip ip address 10.1.1.30 255.255.255.0  | Configures the IP address to be summarized and the IP network mask. | 
| Step 6 | no ip split 				  horizon Example:  Device(config-if)# no ip split horizon  | Disables split horizon on the interface. | 
| Step 7 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 8 | show ip 				  interface  				interface-id Example:  Device# show ip interface gigabitethernet 1/0/1  | Verifies your entries. | 
| Step 9 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Split Horizon
Routers connected to broadcast-type IP networks and using distance-vector routing protocols normally use the split-horizon mechanism to reduce the possibility of routing loops. Split horizon blocks information about routes from being advertised by a router on any interface from which that information originated. This feature can optimize communication among multiple routers, especially when links are broken.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the interface to configure. | 
| Step 4 | ip address  				ip-address 				  subnet-mask Example:  Device(config-if)# ip address 10.1.1.10 255.255.255.0  | Configures the IP address and IP subnet. | 
| Step 5 | no ip 				  split-horizon Example:  Device(config-if)# no ip split-horizon  | Disables split horizon on the interface. | 
| Step 6 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 7 | show ip 				  interface  				interface-id Example:  Device# show ip interface gigabitethernet 1/0/1  | Verifies your entries. | 
| Step 8 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuration Example for Summary Addresses and Split Horizon
In this example, the major net is 10.0.0.0. The summary address 10.2.0.0 overrides the autosummary address of 10.0.0.0 so that 10.2.0.0 is advertised out interface Gigabit Ethernet port 2, and 10.0.0.0 is not advertised. In the example, if the interface is still in Layer 2 mode (the default), you must enter a no switchport interface configuration command before entering the ip address interface configuration command.
| Note | If split horizon is enabled, neither autosummary nor interface summary addresses (those configured with the ip summary-address rip router configuration command) are advertised.  Device(config)# router rip Device(config-router)# interface gigabitethernet1/0/2 Device(config-if)# ip address 10.1.5.1 255.255.255.0 Device(config-if)# ip summary-address rip 10.2.0.0 255.255.0.0 Device(config-if)# no ip split-horizon Device(config-if)# exit Device(config)# router rip Device(config-router)# network 10.0.0.0 Device(config-router)# neighbor 2.2.2.2 peer-group mygroup Device(config-router)# end  | 
Information About OSPF
OSPF is an Interior Gateway Protocol (IGP) designed expressly for IP networks, supporting IP subnetting and tagging of externally derived routing information. OSPF also allows packet authentication and uses IP multicast when sending and receiving packets. The Cisco implementation supports RFC 1253, OSPF management information base (MIB).
| Note | OSPF is supported in IP Base. | 
The Cisco implementation conforms to the OSPF Version 2 specifications with these key features:
-  
		  Definition of stub areas is supported.
-  
		  Routes learned through any IP routing protocol can be redistributed into another IP routing protocol. At the intradomain level, this means that OSPF can import routes learned through EIGRP and RIP. OSPF routes can also be exported into RIP.
-  
		  Plain text and MD5 authentication among neighboring routers within an area is supported.
-  
		  Configurable routing interface parameters include interface output cost, retransmission interval, interface transmit delay, router priority, router dead and hello intervals, and authentication key.
-  
		  Virtual links are supported.
-  
		  Not-so-stubby-areas (NSSAs) per RFC 1587are supported.
OSPF typically requires coordination among many internal routers, area border routers (ABRs) connected to multiple areas, and autonomous system boundary routers (ASBRs). The minimum configuration would use all default parameter values, no authentication, and interfaces assigned to areas. If you customize your environment, you must ensure coordinated configuration of all routers.
OSPF Nonstop Forwarding
The Device or switch stack supports two levels of nonstop forwarding (NSF):
OSPF NSF Awareness
