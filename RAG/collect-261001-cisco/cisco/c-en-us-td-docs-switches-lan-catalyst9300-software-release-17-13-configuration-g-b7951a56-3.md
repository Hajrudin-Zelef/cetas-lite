---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56.md
source_anchor: ""
source_lines: [124, 183]
sha256: c43bc8e8a81db52b051d4539a09706472097d3bc9f387c8dfdef37cf576af59f
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56

| Step 9 | switchport port-security [mac-address mac-address [vlan {vlan-id \| {access \| voice}}] Example:  DEvice(config-if)# switchport port-security mac-address 00:A0:C7:12:C9:25 vlan 3 voice   | (Optional) Enters a secure MAC address for the interface. You can use this command to enter the maximum number of secure MAC addresses. If you configure fewer secure MAC addresses than the maximum, the remaining MAC addresses are dynamically learned. (Optional) vlan —sets a per-VLAN maximum value. Enter one of these options after you enter the vlan keyword:  | 
| Step 10 | switchport port-security mac-address sticky Example:  Device(config-if)# switchport port-security mac-address sticky  | (Optional) Enables sticky learning on the interface. | 
| Step 11 | switchport port-security mac-address sticky [mac-address \| vlan {vlan-id \| {access \| voice}}] Example:  Device(config-if)# switchport port-security mac-address sticky 00:A0:C7:12:C9:25 vlan voice   | (Optional) Enters a sticky secure MAC address, repeating the command as many times as necessary. If you configure fewer secure MAC addresses than the maximum, the remaining MAC addresses are dynamically learned, are converted to sticky secure MAC addresses, and are added to the running configuration. (Optional) vlan —sets a per-VLAN maximum value. Enter one of these options after you enter the vlan keyword:  | 
| Step 12 | end Example:  Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 13 | show port-security Example:  Device# show port-security   | Displays information about the port-security settings. | 
| Note | Under certain conditions, when port security is enabled on the member ports in a switch stack, the DHCP and ARP packets would be dropped. As a workaround, shutdown the interface and then configure the no shutdown command. | 
| Note | The voice keyword is available only if a voice VLAN is configured on a port and if that port is not the access VLAN. If an interface is configured for voice VLAN, configure a maximum of two secure MAC addresses. | 
| Note | If you enable sticky learning after you enter this command, the secure addresses that were dynamically learned are converted to sticky secure MAC addresses and are added to the running configuration. | 
| Note | If you do not enable sticky learning before this command is entered, an error message appears, and you cannot enter a sticky secure MAC address. | 
Use this feature to remove and add devices on a secure port without manually deleting the existing secure MAC addresses and to still limit the number of secure addresses on a port. You can enable or disable the aging of secure addresses on a per-port basis.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet1/0/1   | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | switchport port-security aging {static \| time time \| type {absolute \| inactivity}} Example:  Device(config-if)# switchport port-security aging time 120   | Enables or disable static aging for the secure port, or set the aging time or type. Enter static to enable aging for statically configured secure addresses on this port. For time , specifies the aging time for this port. The valid range is from 0 to 1440 minutes. For type , select one of these keywords:  | 
| Step 5 | end Example:  Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show port-security [interface interface-id] [address] Example:  Device# show port-security interface gigabitethernet1/0/1  | Displays information about the port-security settings on the specified interface. | 
| Note | The switch does not support port security aging of sticky secure addresses. | 
Follow these steps to configure the dynamic address table aging time:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal | Enters global configuration mode. | 
| Step 3 | mac address-table aging-time [0 \| 10-1000000] [routed-mac \| vlan vlan-id] Example:  Device(config)# mac address-table  aging-time 500 vlan 2   | Sets the length of time that a dynamic entry remains in the MAC address table after the entry is used or updated. The range is 10 to 1000000 seconds. The default is 300. You can also enter 0, which disables aging. Static address entries are never aged or removed from the table. vlan-id— Valid IDs are 1 to 4094. | 
| Step 4 | end Example:  Device(config)# end | Exits global configuration mode and returns to privileged EXEC mode. | 
This table displays port security information.
| Table 6. Commands for Displaying Port Security Status and Configuration |  | 
|---|---|
| Command | Purpose | 
|---|---|
| show port-security [interface interface-id] | Displays port security settings for the device or for the specified interface, including the maximum allowed number of secure MAC addresses for each interface, the number of secure MAC addresses on the interface, the number of security violations that have occurred, and the violation mode. | 
| show port-security [interface interface-id] address | Displays all secure MAC addresses configured on all device interfaces or on a specified interface with aging information for each address. | 
| show port-security interface interface-id vlan | Displays the number of secure MAC addresses configured per VLAN on the specified interface. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable  | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 3 | mgmt-traffic control ipv4 Example:  Device# mgmt-traffic control ipv4 Device(config-mtc-ipv4)# | Enables the management traffic control feature. | 
| Step 4 | interface interface-id Example:  Device(config-mtc-ipv4)# interface gigabitethernet1/0/1 Device(config-mtc-ipv4)# | Defines the interface through which traffic is allowed. | 
| Step 5 | protocol{[ telnet http https netconf scp snmp ssh]} Example:  Device(config-mtc-ipv4)# protocol telnet http https netconf scp snmp ssh Device(config-mtc-ipv4)# | Enables the specified network protocols in the interface. | 
| Step 6 | address ip-address Example:  Device(config-mtc-ipv4)# address 30.30.30.30 Device(config-mtc-ipv4)# | Specifies the destination address of traffic. | 
| Step 7 | end Example:  Device(config-mtc-ipv4)# end  | Exits management traffic control configuration mode and returns to privileged EXEC mode. | 
This example shows how to enable port security on a port and to set the maximum number of secure addresses to 50. The violation mode is the default, no static secure MAC addresses are configured, and sticky learning is enabled.
Device> enable
Device# configure terminal
Device(config)# interface gigabitethernet1/0/1
Device(config-if)# switchport mode access
Device(config-if)# switchport port-security
Device(config-if)# switchport port-security maximum 50
Device(config-if)# switchport port-security mac-address sticky
Device(config-if)# end
This example shows how to configure a static secure MAC address on VLAN 3 on a port:
Device> enable
Device# configure terminal
Device(config)# interface gigabitethernet1/0/2
Device(config-if)# switchport mode trunk
Device(config-if)# switchport port-security
Device(config-if)# switchport port-security mac-address 0000.0200.0004 vlan 3
Device(config-if)# end
