---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6-4
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6.md
source_anchor: ""
source_lines: [95, 179]
sha256: 246221287e6e933d867ceb5a8cb0cd816484f7b91fdf62a9b1813344c519c52e
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6

- Cisco NX-OS removes all Layer 3 configurations on an interface when you change the interface VRF membership, port channel membership, or the port mode to Layer 2.
- If you configure virtual MAC addresses with vPC, you must configure the same virtual MAC address on both vPC peers.
- You cannot use the HSRP MAC address burned-in option on a VLAN interface that is a vPC member.
- Cisco NX-OS does not support having the same HSRP groups on all nodes in a double-sided vPC.
- If you have not configured authentication, the show hsrp command displays the following string:
The default behavior of HSRP is as defined in RFC 2281 :
value is 0x63 0x69 0x73 0x63 0x6F 0x00 0x00 0x00.
Default Settings
Table 17-2 lists the default settings for HSRP parameters.
Configuring HSRP
This section includes the following topics:
Note If you are familiar with the Cisco IOS CLI, be aware that the Cisco NX-OS commands for this feature might differ from the Cisco IOS commands that you would use.
Enabling HSRP
You must globally enable HSRP before you can configure and enable any HSRP groups.
DETAILED STEPS
To enable the HSRP feature, use the following command in global configuration mode:
To disable the HSRP feature and remove all associated configurations, use the following command in global configuration mode:
Configuring the HSRP Version
You can configure the HSRP version. If you change the version for existing groups, Cisco NX-OS reinitializes HSRP for those groups because the virtual MAC address changes. The HSRP version applies to all groups on the interface.
Note IPv6 HSRP groups must be configured as HSRP version 2.
To configure the HSRP version, use the following command in interface configuration mode:
Configuring an HSRP Group for IPv4
You can configure an HSRP group on an IPv4 interface and configure the virtual IP address and virtual MAC address for the HSRP group.
BEFORE YOU BEGIN
Ensure that you have enabled the HSRP feature (see the “Enabling HSRP” section).
Cisco NX-OS enables an HSRP group once you configure its virtual IP address. You should configure HSRP attributes such as authentication, timers, and priority before you enable the HSRP group.
SUMMARY STEPS
3. ip address ip-address/length
5. ip [ ip-address [ secondary ]]
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | interface type number  Example: switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 3 | ip address ip-address/length  Example : switch(config-if)# ip 192.0.2.2/8 | Configures the IPv4 address of the interface. | 
| Step 4 | hsrp group-number [ ipv4 ]  Example : switch(config-if)# hsrp 2 switch(config-if-hsrp)# | Creates an HSRP group and enters HSRP configuration mode. The range for HSRP version 1 is from 0 to 255. The range is for HSRP version 2 is from 0 to 4095. The default value is 0. | 
| Step 5 | ip [ ip-address [ secondary ]]  Example : switch(config-if-hsrp)# ip 192.0.2.1 | Configures the virtual IP address for the HSRP group and enables the group. This address should be in the same subnet as the IPv4 address of the interface. | 
| Step 6 | exit  Example : switch(config-if-hsrp)# exit | Exits HSRP configuration mode. | 
| Step 7 | no shutdown  Example : switch(config-if)# no shutdown | Enables the interface. | 
| Step 8 | show hsrp [ group group-number ] [ ipv4 ]  Example : switch(config-if)# show hsrp group 2 | (Optional) Displays HSRP information. | 
| Step 9 | copy running-config startup-config  Example: switch(config-if)# copy running-config startup-config | (Optional) Copies the running configuration to the startup configuration. | 
Note You should use the no shutdown command to enable the interface after you finish the configuration.
This example shows how to configure an HSRP group on Ethernet 1/2:
switch(config)# interface ethernet 1/2
switch(config-if)# ip 192.0.2.2/8
switch(config-if-hsrp)# ip 192.0.2.1
Configuring an HSRP Group for IPv6
You can configure an HSRP group on an IPv6 interface and configure the virtual MAC address for the HSRP group.
When you configure an HSRP group for IPv6, HSRP generates a link-local address from the link-local prefix. HSRP also generates a modified EUI-64 format interface identifier in which the EUI-64 interface identifier is created from the relevant HSRP virtual MAC address.
BEFORE YOU BEGIN
You must enable HSRP (see the “Enabling HSRP” section).
Ensure that you have enabled HSRP version 2 on the interface that you want to configure an IPv6 HSRP group on.
Ensure that you have configured HSRP attributes such as authentication, timers, and priority before you enable the HSRP group.
SUMMARY STEPS
3. ipv6 address ipv6-address/length
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | interface type number  Example: switch(config)# interface ethernet 3/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 3 | ipv6 address ipv6-address/length  Example : switch(config-if)# ipv6 address 2001:0DB8::0001:0001/64 | Configures the IPv6 address of the interface. | 
| Step 4 | hsrp version 2  Example : switch(config-if-hsrp)# hsrp version 2 | Configures this group for HSRP version 2. | 
| Step 5 | hsrp group-number ipv6  Example : switch(config-if)# hsrp 10 ipv6 switch(config-if-hsrp)# | Creates an IPv6 HSRP group and enters HSRP configuration mode. The range for HSRP version 2 is from 0 to 4095. The default value is 0. | 
| Step 6 | ip ipv6-address  Example : switch(config-if-hsrp)# ip 2001:DB8::1 | Configures the virtual IPv6 address for the HSRP group and enables the group. | 
| Step 7 | ip autoconfig  Example : switch(config-if-hsrp)# ip autoconfig | Autoconfigures the virtual IPv6 address for the HSRP group from the calculated link-local virtual IPv6 address and enables the group. | 
| Step 8 | exit  Example : switch(config-if-hsrp)# exit switch(config-if)# | Exits HSRP configuration mode. | 
| Step 9 | no shutdown  Example : switch(config-if)# no shutdown | Enables the interface. | 
| Step 10 | show hsrp [ group group-number ] [ ipv6 ]  Example : switch(config-if)# show hsrp group 10 | (Optional) Displays HSRP information. | 
| Step 11 | copy running-config startup-config  Example: switch(config-if)# copy running-config startup-config | (Optional) Copies the running configuration to the startup configuration. | 
Note You should use the no shutdown command to enable the interface after you finish the configuration.
This example shows how to configure an IPv6 HSRP group on Ethernet 3/2:
switch(config)# interface ethernet 3/2
switch(config-if)# ipv6 address 2001:0DB8::0001:0001/64
switch(config-if-hsrp)# hsrp version 2
switch(config-if)# hsrp 2 ipv6
switch(config-if-hsrp)# ip 2001:DB8::1
Configuring the HSRP Virtual MAC Address
You can override the default virtual MAC address that HSRP derives from the configured group number.
Note You must configure the same virtual MAC address on both vPC peers of a vPC link.
To manually configure the virtual MAC address for an HSRP group, use the following command in hsrp configuration mode:
To configure HSRP to use the burned-in MAC address of the interface for the virtual MAC address, use the following command in interface configuration mode:
|  |  | 
|---|---|
| hsrp use-bia [ scope interface ]  Example : switch(config-if)# hsrp use-bia | Configures HSRP to use the burned-in MAC address of the interface for the HSRP virtual MAC address. You can optionally configure HSRP to use the burned-in MAC address for all groups on this interface by using the scope interface keyword. | 
Authenticating HSRP
