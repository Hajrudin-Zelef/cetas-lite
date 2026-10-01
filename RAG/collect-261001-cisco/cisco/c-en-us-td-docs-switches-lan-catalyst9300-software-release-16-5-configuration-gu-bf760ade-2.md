---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade.md
source_anchor: ""
source_lines: [59, 123]
sha256: 045522175025dd4036907627aa1046036a49fe9be9e23e4f734358f91f4fb674
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade

                                    A protected port does not forward any traffic (unicast, multicast, or broadcast) to any other port that is also a protected port. Data traffic cannot be forwarded between protected ports at Layer 2; only control traffic, such as PIM packets, is forwarded because these packets are processed by the CPU and forwarded in software. All data traffic passing between protected ports must be forwarded through a Layer 3 device.
-  
                                    		  
                                    Forwarding behavior between a protected port and a nonprotected port proceeds as usual.
Because a switch stack represents a single logical switch, Layer 2 traffic is not forwarded between any protected ports in the switch stack, whether they are on the same or different switches in the stack.
Default Protected Port Configuration
The default is to have no protected ports defined.
Protected Ports Guidelines
You can configure protected ports on a physical interface (for example, Gigabit Ethernet port 1) or an EtherChannel group (for example, port-channel 5). When you enable protected ports for a port channel, it is enabled for all ports in the port-channel group.
How to Configure Protected Ports
Configuring a Protected Port
Before you begin
Protected ports are not pre-defined. This is the task to configure one.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable   | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet1/0/1   | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | switchport protected Example:  Device(config-if)# switchport protected   | Configures the interface to be a protected port. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show interfaces interface-id switchport Example:  Device# show interfaces gigabitethernet1/0/1 switchport   | Verifies your entries. | 
| Step 7 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 8 | copy running-config startup-config Example:  Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring Protected Ports
| Table 1. Commands for Displaying Protected Port Settings |  | 
|---|---|
| Command | Purpose | 
|---|---|
| show interfaces [interface-id] switchport | Displays the administrative and operational status of all switching (nonrouting) ports or the specified port, including port blocking and port protection settings. | 
Information About Port Blocking
Port Blocking
By default, the switch floods packets with unknown destination MAC addresses out of all ports. If unknown unicast and multicast traffic is forwarded to a protected port, there could be security issues. To prevent unknown unicast or multicast traffic from being forwarded from one port to another, you can block a port (protected or nonprotected) from flooding unknown unicast or multicast packets to other ports.
| Note | With multicast traffic, the port blocking feature blocks only pure Layer 2 packets. Multicast packets that contain IPv4 or IPv6 information in the header are not blocked. | 
How to Configure Port Blocking
Blocking Flooded Traffic on an Interface
Before you begin
The interface can be a physical interface or an EtherChannel group. When you block multicast or unicast traffic for a port channel, it is blocked on all ports in the port-channel group.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable   | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet1/0/1   | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | switchport block multicast Example:  Device(config-if)# switchport block multicast   | Blocks unknown multicast forwarding out of the port. | 
| Step 5 | switchport block unicast Example:  Device(config-if)# switchport block unicast   | Blocks unknown unicast forwarding out of the port. | 
| Step 6 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 7 | show interfaces interface-id switchport Example:  Device# show interfaces gigabitethernet1/0/1 switchport   | Verifies your entries. | 
| Step 8 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 9 | copy running-config startup-config Example:  Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring Port Blocking
| Table 2. Commands for Displaying Port Blocking Settings |  | 
|---|---|
| Command | Purpose | 
|---|---|
| show interfaces [interface-id] switchport | Displays the administrative and operational status of all switching (nonrouting) ports or the specified port, including port blocking and port protection settings. | 
Prerequisites for Port Security
| Note | If you try to set the maximum value to a number less than the number of secure addresses already configured on an interface, the command is rejected. | 
Restrictions for Port Security
- 
                                    					
                                    The maximum number of secure MAC addresses that you can configure on a switch is set by the maximum number of available MAC addresses allowed in the system. This number is the total of available MAC addresses, including those used for other Layer 2 functions and any other secure MAC addresses configured on interfaces.
- 
                                    					
                                    					
