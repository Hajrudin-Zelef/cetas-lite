---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-4
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [173, 236]
sha256: b7121e85bdfafe381145dcd8f336b4c3993932d0a81a8101acee04f64d877af6
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

An interface can have one primary IP address. A mask identifies the bits that denote the network number in an IP address. When you use the mask to subnet a network, the mask is referred to as a subnet mask. To receive an assigned network number, contact your Internet service provider.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | no 				  switchport Example:  Device(config-if)# no switchport  | Removes the interface from Layer 2 configuration mode (if it is a physical interface). | 
| Step 5 | ip address  				ip-address 				  subnet-mask Example:  Device(config-if)# ip address 10.1.5.1 255.255.255.0  | Configures the IP address and IP subnet mask. | 
| Step 6 | no 				  shutdown Example:  Device(config-if)# no shutdown  | Enables the physical interface. | 
| Step 7 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 8 | show ip 				  route Example:  Device# show ip route  | Verifies your entries. | 
| Step 9 | show ip interface [interface-id] Example:  Device# show ip interface gigabitethernet 1/0/1  | Verifies your entries. | 
| Step 10 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 11 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Using Subnet Zero
Subnetting with a subnet address of zero is strongly discouraged because of the problems that can arise if a network and a subnet have the same addresses. For example, if network 131.108.0.0 is subnetted as 255.255.255.0, subnet zero would be written as 131.108.0.0, which is the same as the network address.
You can use the all ones subnet (131.108.255.0) and even though it is discouraged, you can enable the use of subnet zero if you need the entire subnet space for your IP address.
Use the no ip subnet-zero global configuration command to restore the default and disable the use of subnet zero.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | ip 				  subnet-zero Example:  Device(config)# ip subnet-zero  | Enables the use of subnet zero for interface addresses and routing updates. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Disabling Classless Routing
To prevent the Device from forwarding packets destined for unrecognized subnets to the best supernet route possible, you can disable classless routing behavior.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | no ip 				  classless Example:  Device(config)#no ip classless  | Disables classless routing behavior. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Address Resolution Methods
You can perform the following tasks to configure address resolution.
Defining a Static ARP Cache
ARP and other address resolution protocols provide dynamic mapping between IP addresses and MAC addresses. Because most hosts support dynamic address resolution, you usually do not need to specify static ARP cache entries. If you must define a static ARP cache entry, you can do so globally, which installs a permanent entry in the ARP cache that the Device uses to translate IP addresses into MAC addresses. Optionally, you can also specify that the Device respond to ARP requests as if it were the owner of the specified IP address. If you do not want the ARP entry to be permanent, you can specify a timeout period for the ARP entry.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | arp  				ip-address hardware-address 				  type Example:  Device(config)# ip 10.1.5.1 c2f3.220a.12f4 arpa  | Associates an IP address with a MAC (hardware) address in the ARP cache, and specifies encapsulation type as one of these: | 
| Step 4 | arp  				ip-address hardware-address 				  type [alias] Example:  Device(config)# ip 10.1.5.3 d7f3.220d.12f5 arpa alias  | (Optional) Specifies that the switch respond to ARP requests as if it were the owner of the specified IP address. | 
| Step 5 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the interface to configure. | 
| Step 6 | arp  				timeout seconds Example:  Device(config-if)# arp 20000  | (Optional) Sets the length of time an ARP cache entry will stay in the cache. The default is 14400 seconds (4 hours). The range is 0 to 2147483 seconds. | 
| Step 7 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 8 | show interfaces [interface-id] Example:  Device# show interfaces gigabitethernet 1/0/1  | Verifies the type of ARP and the timeout value used on all interfaces or a specific interface. | 
| Step 9 | show 				  arp Example:  Device# show arp  | Views the contents of the ARP cache. | 
| Step 10 | show ip 				  arp Example:  Device# show ip arp  | Views the contents of the ARP cache. | 
| Step 11 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Setting ARP Encapsulation
By default, Ethernet ARP encapsulation (represented by the arpa keyword) is enabled on an IP interface. You can change the encapsulation methods to SNAP if required by your network.
To disable an encapsulation type, use the no arp arpa or no arp snap interface configuration command.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/2  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | arp {arpa \|  				snap} Example:  Device(config-if)# arp arpa  | Specifies the ARP encapsulation method: | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show interfaces [interface-id] Example:  Device# show interfaces  | Verifies ARP encapsulation configuration on all interfaces or the specified interface. | 
