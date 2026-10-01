---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-6
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [297, 368]
sha256: 1fce970d42516372622b98f4b1061c5a608ac641b02670ca0965468238849a03
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 4 | ip directed-broadcast 				[access-list-number] Example:  Device(config-if)# ip directed-broadcast 103  | Enables directed broadcast-to-physical broadcast translation on the interface. You can include an access list to control which broadcasts are forwarded. When an access list, only IP packets permitted by the access list can be translated. | 
| Step 5 | exit Example:  Device(config-if)# exit  | Returns to global configuration mode. | 
| Step 6 | ip forward-protocol 				{udp [port] \|  				nd \|  				sdns} Example:  Device(config)# ip forward-protocol nd  | Specifies which protocols and ports the router forwards when forwarding broadcast packets. | 
| Step 7 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 8 | show ip interface [interface-id] Example:  Device# show ip interface  | Verifies the configuration on the interface or all interfaces | 
| Step 9 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 10 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Forwarding UDP Broadcast Packets and Protocols
If you do not specify any UDP ports when you configure the forwarding of UDP broadcasts, you are configuring the router to act as a BOOTP forwarding agent. BOOTP packets carry DHCP information.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | ip 				  helper-address  				address Example:  Device(config-if)# ip helper address 10.1.10.1  | Enables forwarding and specifies the destination address for forwarding UDP broadcast packets, including BOOTP. | 
| Step 5 | exit Example:  Device(config-if)# exit  | Returns to global configuration mode. | 
| Step 6 | ip forward-protocol {udp [port] \|  				nd \|  				sdns} Example:  Device(config)# ip forward-protocol sdns  | Specifies which protocols the router forwards when forwarding broadcast packets. | 
| Step 7 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 8 | show ip interface [interface-id] Example:  Device# show ip interface gigabitethernet 1/0/1  | Verifies the configuration on the interface or all interfaces. | 
| Step 9 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 10 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Establishing an IP Broadcast Address
The most popular IP broadcast address (and the default) is an address consisting of all ones (255.255.255.255). However, the Device can be configured to generate any form of IP broadcast address.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the interface to configure. | 
| Step 4 | ip 				  broadcast-address  				ip-address Example:  Device(config-if)# ip broadcast-address 128.1.255.255  | Enters a broadcast address different from the default, for example 128.1.255.255. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show ip interface [interface-id] Example:  Device# show ip interface  | Verifies the broadcast address on the interface or all interfaces. | 
| Step 7 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Flooding IP Broadcasts
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | ip 				  forward-protocol spanning-tree Example:  Device(config)# ip forward-protocol spanning-tree  | Uses the bridging spanning-tree database to flood UDP datagrams. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Step 7 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 8 | ip 				  forward-protocol turbo-flood Example:  Device(config)# ip forward-protocol turbo-flood  | Uses the spanning-tree database to speed up flooding of UDP datagrams. | 
| Step 9 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 10 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 11 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring and Maintaining IP Addressing
When the contents of a particular cache, table, or database have become or are suspected to be invalid, you can remove all its contents by using the clear privileged EXEC commands. The Table lists the commands for clearing contents.
| Table 2 Commands to Clear Caches, 		  Tables, and Databases |  | 
|---|---|
| clear arp-cache | Clears the IP ARP cache and the fast-switching cache. | 
| clear host {name \| *} | Removes one or all entries from the hostname and the address cache. | 
| clear ip route {network [mask] \| *} | Removes one or more routes from the IP routing table. | 
You can display specific statistics, such as the contents of IP routing tables, caches, and databases; the reachability of nodes; and the routing path that packets are taking through the network. The Table lists the privileged EXEC commands for displaying IP statistics.
| Table 3 Commands to Display Caches, 		  Tables, and Databases |  | 
|---|---|
| show arp | Displays the entries in the ARP table. | 
| show hosts | Displays the default domain name, style of lookup service, name server hosts, and the cached list of hostnames and addresses. | 
| show ip aliases | Displays IP addresses mapped to TCP ports (aliases). | 
| show ip arp | Displays the IP ARP cache. | 
| show ip interface [interface-id] | Displays the IP status of interfaces. | 
| show ip irdp | Displays IRDP values. | 
| show ip masks address | Displays the masks used for network addresses and the number of subnets using each mask. | 
| show ip redirects | Displays the address of a default gateway. | 
| show ip route [address [mask]] \| [protocol] | Displays the current state of the routing table. | 
| show ip route summary | Displays the current state of the routing table in summary form. | 
How to Configure IP Unicast Routing
Enabling IP Unicast Routing
By default, the Device is in Layer 2 switching mode and IP routing is disabled. To use the Layer 3 capabilities of the Device, you must enable IP routing.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
