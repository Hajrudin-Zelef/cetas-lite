---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [31, 69]
sha256: ecfd23e908168cc7a06f33098c2fb315fc2eb2ff88f60e4ff5901cfe5617587e
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

It involves a hardware-programmed (TCAM table) filter which allows or denies traffic based on its source address. For the filter to work this way, an entry (of the source address) in the binding table is required. If the source address is in the binding table, the filter allows the packet into the network; if the address is not in the binding table, entry is denied and the packet is dropped. When an entry is removed from the binding table, the filter is also removed, and subsequent packets with that source address are dropped.
When configuring this feature, consider the following:
The IPv6 Source Guard and Prefix Guard features are supported only in the ingress direction and not supported in the egress direction.
You cannot use IPv6 Source Guard and Prefix Guard together. When you attach the policy to an interface, it should be "validate address" or "validate prefix" but not both.
PVLAN and Source or Prefix Guard cannot be applied together.
IPv6 Source Guard and Prefix Guard is supported on EtherChannels
An IPv6 source guard policy cannot be attached to a VLAN. It is supported only at the interface level.
When you configure IPv4 and IPv6 source guard together on an interface, it is recommended to use ip verify source mac-check command instead of ip verify source tracking mac-check command. IPv4 connectivity on a given port might break due to two different filtering rules set: one for IPv4 (IP-filter) and the other for IPv6 (IP-MAC filter).
When IPv6 source guard is enabled on a switch port, NDP or DHCP snooping must be enabled on the interface to which the switch port belongs. Otherwise, all data traffic from this port will be blocked.
Binding information is normally gleaned from IPv6 NDP traffic and DHCP packets. If you rely only on a DHCP server for source addresses of hosts, ensure that you also configure a data-glean recovery function to counteract a situation where entries are prematurely removed from the binding table (for various reasons) before the DHCP lease timer expires. This way, the recovery function restores binding entries of valid hosts and you can be sure that that the IPv6 Source Guard feature allows only packets with a DHCP server-assigned source address. See Example: Using the Data-Glean Recovery Function.
To use this feature, you must configure an IPv6 Source Guard policy and attach it to a target. See Configuring IPv6 Source Guard .
To debug source-guard packets, use the debug ipv6 snooping source-guard privileged EXEC command.
The IPv6 Prefix Guard feature works within the IPv6 Source Guard feature to enable the device to deny traffic originated from non-topologically correct addresses. IPv6 Prefix Guard is often used when IPv6 prefixes are delegated to devices (for example, home gateways) using DHCP prefix delegation. The feature discovers ranges of addresses assigned to the link and blocks any traffic sourced with an address outside this range.
In order to use this feature, you must configure an IPv6 Prefix Guard policy and attach it to a target. See Configuring IPv6 Prefix Guard.
| Note | Ensure that you have read the configuration considerations listed in the IPv6 Source Guard section above - some of them apply to the IPv6 Prefix Guard feature as well. | 
The IPv6 Destination Guard feature works with IPv6 neighbor discovery to ensure that the device performs address resolution only for those addresses that are known to be active on the link. It relies on the address glean functionality to populate all destinations active on the link into the binding table and then blocks resolutions before they happen when the destination is not found in the binding table.
| Note | We recommend that you apply an IPv6 Destination Guard policy on all Layer 2 VLANs with an SVI configured. | 
In order to use this feature, you must configure an IPv6 Destination Guard policy and attach it to a target. See Configuring an IPv6 Destination Guard Policy.
| Note | The IPv6 Snooping Policy feature has been deprecated. Although the commands are visible on the CLI and you can configure them, we recommend that you use the Switch Integrated Security Feature (SISF)-based Device Tracking feature instead. | 
Beginning in privileged EXEC mode, follow these steps to configure IPv6 Snooping Policy :
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | ipv6 snooping policy policy-name Example: Device(config)# ipv6 snooping policy example_policy | Creates a snooping policy and enters IPv6 snooping policy configuration mode. | 
| Step 4 | {[default ] \| [device-role {node \| switch}] \| [limit address-count value] \| [no] \| [protocol {dhcp \| ndp} ] \| [security-level {glean \| guard \| inspect} ] \| [tracking {disable [stale-lifetime [seconds \| infinite] \| enable [reachable-lifetime [seconds \| infinite] } ] \| [trusted-port ] } Example: Device(config-ipv6-snooping)# security-level inspect Example: Device(config-ipv6-snooping)# trusted-port | Enables data address gleaning, validates messages against various criteria, specifies the security level for messages.  | 
| Step 5 | end Example: Device(config-ipv6-snooping)# end | Exits IPv6 snooping policy configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show ipv6 snooping policy policy-name Example: Device#show ipv6 snooping policy example_policy | Displays the snooping policy configuration. | 
Attach an IPv6 Snooping policy to interfaces or VLANs.
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Snooping policy on an interface or VLAN:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface interface_type stack/module/port Example: Device(config)#  interface gigabitethernet 1/1/4     | Specifies an interface type and identifier and enters the interface configuration mode. | 
| Step 4 | switchport Example: Device(config-if)# switchport | Enters the Switchport mode. | 
| Step 5 | ipv6 snooping [attach-policy policy_name [ vlan {vlan_id \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids}] \| vlan {vlan_id \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if)# ipv6 snooping Device(config-if)# ipv6 snooping attach-policy example_policyDevice(config-if)# ipv6 snooping vlan 111,112Device(config-if)# ipv6 snooping attach-policy example_policy vlan 111,112 | Attaches a custom IPv6 snooping policy to the interface or the specified VLANs on the interface. To attach the default policy to the interface, use the ipv6 snooping command without the attach-policy keyword. To attach the default policy to VLANs on the interface, use the ipv6 snooping vlan command. The default policy is, security-level guard, device-role node, protocol ndp and dhcp. | 
| Step 6 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 7 | show running-config Example: Device# show running-config  | Verifies that the policy is attached to the specified interface without exiting the interface configuration mode. | 
