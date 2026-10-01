---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade.md
source_anchor: ""
source_lines: [241, 293]
sha256: 601fa17c91a35ecce1b6281240fc37ace64efaac66597751debb39a6b46be175
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade

                                    Port Blocking
How to Configure Port Security
Enabling and Configuring Port Security
Before you begin
This task restricts input to an interface by limiting and identifying MAC addresses of the stations allowed to access the port:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable   | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet1/0/1   | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | switchport mode {access \| trunk} Example:  Device(config-if)# switchport mode access   | Sets the interface switchport mode as access or trunk; an interface in the default mode (dynamic auto) cannot be configured as a secure port. | 
| Step 5 | switchport voice vlan vlan-id Example:  Device(config-if)# switchport voice vlan 22   | Enables voice VLAN on a port. vlan-id—Specifies the VLAN to be used for voice traffic. | 
| Step 6 | switchport port-security Example:  Device(config-if)# switchport port-security   | Enable port security on the interface. | 
| Step 7 | switchport port-security [maximum value [vlan {vlan-list \| {access \| voice}}]] Example:  Device(config-if)# switchport port-security maximum 20   | (Optional) Sets the maximum number of secure MAC addresses for the interface. The maximum number of secure MAC addresses that you can configure on a switch or switch stack is set by the maximum number of available MAC addresses allowed in the system. This number is the total of available MAC addresses, including those used for other Layer 2 functions and any other secure MAC addresses configured on interfaces. (Optional) vlan —sets a per-VLAN maximum value Enter one of these options after you enter the vlan keyword:  | 
| Step 8 | switchport port-security violation {protect \| restrict \| shutdown \| shutdown vlan} Example:  Device(config-if)# switchport port-security violation restrict   | (Optional) Sets the violation mode, the action to be taken when a security violation is detected, as one of these:  | 
| Step 9 | switchport port-security [mac-address mac-address [vlan {vlan-id \| {access \| voice}}] Example:  Device(config-if)# switchport port-security mac-address 00:A0:C7:12:C9:25 vlan 3 voice   | (Optional) Enters a secure MAC address for the interface. You can use this command to enter the maximum number of secure MAC addresses. If you configure fewer secure MAC addresses than the maximum, the remaining MAC addresses are dynamically learned. (Optional) vlan —sets a per-VLAN maximum value. Enter one of these options after you enter the vlan keyword:  | 
| Step 10 | switchport port-security mac-address sticky Example:  Device(config-if)# switchport port-security mac-address sticky  | (Optional) Enables sticky learning on the interface. | 
| Step 11 | switchport port-security mac-address sticky [mac-address \| vlan {vlan-id \| {access \| voice}}] Example:  Device(config-if)# switchport port-security mac-address sticky 00:A0:C7:12:C9:25 vlan voice   | (Optional) Enters a sticky secure MAC address, repeating the command as many times as necessary. If you configure fewer secure MAC addresses than the maximum, the remaining MAC addresses are dynamically learned, are converted to sticky secure MAC addresses, and are added to the running configuration. (Optional) vlan —sets a per-VLAN maximum value. Enter one of these options after you enter the vlan keyword:  | 
| Step 12 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 13 | show port-security Example:  Device# show port-security   | Verifies your entries. | 
| Step 14 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 15 | copy running-config startup-config Example:  Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | Under certain conditions, when port security is enabled on the member ports in a switch stack, the DHCP and ARP packets would be dropped. To resolve this, configure a shut and no shut on the interface. | 
| Note | The voice keyword is available only if a voice VLAN is configured on a port and if that port is not the access VLAN. If an interface is configured for voice VLAN, configure a maximum of two secure MAC addresses. | 
| Note | If you enable sticky learning after you enter this command, the secure addresses that were dynamically learned are converted to sticky secure MAC addresses and are added to the running configuration. | 
| Note | The voice keyword is available only if a voice VLAN is configured on a port and if that port is not the access VLAN. If an interface is configured for voice VLAN, configure a maximum of two secure MAC addresses. | 
| Note | If you do not enable sticky learning before this command is entered, an error message appears, and you cannot enter a sticky secure MAC address. | 
| Note | The voice keyword is available only if a voice VLAN is configured on a port and if that port is not the access VLAN. | 
Enabling and Configuring Port Security Aging
Use this feature to remove and add devices on a secure port without manually deleting the existing secure MAC addresses and to still limit the number of secure addresses on a port. You can enable or disable the aging of secure addresses on a per-port basis.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable   | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet1/0/1   | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | switchport port-security aging {static \| time time \| type {absolute \| inactivity}} Example:  Device(config-if)# switchport port-security aging time 120   | Enables or disable static aging for the secure port, or set the aging time or type. Enter static to enable aging for statically configured secure addresses on this port. For time , specifies the aging time for this port. The valid range is from 0 to 1440 minutes. For type , select one of these keywords:  | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show port-security [interface interface-id] [address] Example:  Device# show port-security interface gigabitethernet1/0/1  | Verifies your entries. | 
| Step 7 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 8 | copy running-config startup-config Example:  Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | The switch does not support port security aging of sticky secure addresses. | 
Monitoring Port Security
This table displays port security information.
| Table 6. Commands for Displaying Port Security Status and Configuration |  | 
|---|---|
| Command | Purpose | 
|---|---|
| show port-security [interface interface-id] | Displays port security settings for the switch or for the specified interface, including the maximum allowed number of secure MAC addresses for each interface, the number of secure MAC addresses on the interface, the number of security violations that have occurred, and the violation mode. | 
| show port-security [interface interface-id] address | Displays all secure MAC addresses configured on all switch interfaces or on a specified interface with aging information for each address. | 
| show port-security interface interface-id vlan | Displays the number of secure MAC addresses configured per VLAN on the specified interface. | 
Configuration Examples for Port Security
