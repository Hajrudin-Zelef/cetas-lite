---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-7
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [252, 340]
sha256: 009b16098aa3ebabca1dcda6e2fc3ae1d1f05fd7e5b75f8a0984f2f9185c0c27
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

| Step 3 | ipv6 source-guard policy source-guard-policy Example: Device(config)# ipv6 source-guard policy my_snooping_policy | Defines an IPv6 source-guard policy name and enters switch integrated security features source-guard policy configuration mode. | 
| Step 4 | validate address Example: Device(config-sisf-sourceguard)# no validate address | Disables the validate address feature and enables the IPv6 prefix guard feature to be configured. | 
| Step 5 | validate prefix Example: Device(config-sisf-sourceguard)# validate prefix | Enables IPv6 source guard to perform the IPv6 prefix-guard operation. | 
| Step 6 | exit Example: Device(config-sisf-sourceguard)# exit | Exits switch integrated security features source-guard policy configuration mode and returns to privileged EXEC mode. | 
| Step 7 | show ipv6 source-guard policy [ source-guard-policy] Example: Device# show ipv6 source-guard policy policy1 | Displays the IPv6 source-guard policy configuration. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface type number Example: Device(config)# interface gigabitethernet 1/1/4     | Specifies an interface type and identifier, and enters interface configuration mode. | 
| Step 4 | ipv6 source-guard attach-policy policy_name Example: Device(config-if)# ipv6 source-guard attach-policy example_policy | Attaches the IPv6 Source Guard policy to the interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show ipv6 source-guard policy policy_name Example: Device(config-if)# show ipv6 source-guard policy example_policy | Shows the policy configuration and all the interfaces where the policy is applied. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface port-channel port-channel-number Example: Device(config)# interface Port-channel 4 | Specifies an interface type and port number and places the switch in the port channel configuration mode. | 
| Step 4 | ipv6 source-guard [ attach-policy <policy_name> ] Example: Device(config-if)# ipv6 source-guard attach-policy example_policy | Attaches the IPv6 Source Guard policy to the interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show ipv6 source-guard policy policy_name Example: Device(config)# show ipv6 source-guard policy example_policy | Shows the policy configuration and all the interfaces where the policy is applied. | 
Beginning in privileged EXEC mode, follow these steps to configure an IPv6 destination guard policy:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | ipv6 destination-guard policy policy-name Example: Device(config)# ipv6 destination-guard policy pol1 | Defines the destination guard policy name and enters destination-guard configuration mode. | 
| Step 4 | enforcement {always \| stressed} Example: Device(config-destguard)# enforcement always | Sets the enforcement level for the target address. | 
| Step 5 | exit Example: Device(config-destguard)# exit | Exits destination-guard configuration mode and returns to global configuration mode. | 
| Step 6 | interface type number Example: Device(config)# interface GigabitEthernet 0/0/1 | Enters interface configuration mode. | 
| Step 7 | ipv6 destination-guard attach-policy [policy-name] Example: Device(config-if)# ipv6 destination-guard attach-policy pol1 | Attaches a destination guard policy to an interface. | 
| Step 8 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC configuration mode. | 
| Step 9 | show ipv6 destination-guard policy [policy-name] Example: Device# show ipv6 destination-guard policy pol1 | (Optional) Displays the policy configuration and all interfaces where the policy is applied. | 
Device> enable
Device# configure terminal
Device(config)# ipv6 access-list acl1
Device(config-ipv6-acl)# permit host 2001:DB8:0000:
0000:0000:0000:0000:0001 any
Device(config-ipv6-acl)# exit
Device(config)# ipv6 prefix-list abc permit 2001:0DB8::/64 le 128	
Device(config)# ipv6 dhcp guard policy pol1
Device(config-dhcp-guard)# device-role server
Device(config-dhcp-guard)# match server access-list acl1
Device(config-dhcp-guard)# match reply prefix-list abc
Device(config-dhcp-guard)# preference min 0
Device(config-dhcp-guard)# preference max 255
Device(config-dhcp-guard)# trusted-port
Device(config-dhcp-guard)# exit
Device(config)# interface GigabitEthernet 0/2/0
Device(config-if)# switchport
Device(config-if)# ipv6 dhcp guard attach-policy pol1 vlan add 1
Device(config-if)# exit
Device(config)# vlan 1
Device(config-vlan)# ipv6 dhcp guard attach-policy pol1
Device(config-vlan)# end
The following example shows how to attach an IPv6 Source Guard Policy to a Layer 2 EtherChannel Interface:
Device> enable
Device# configure terminal
Device(config)# ipv6 source-guard policy POL
Device(config-sisf-sourceguard) # validate address
Device(config-sisf-sourceguard)# exit
Device(config)# interface Port-Channel 4
Device(config-if)# ipv6 snooping 
Device(config-if)# ipv6 source-guard attach-policy POL
Device(config-if)# end
Device#
The following example shows how to attach an IPv6 Prefix Guard Policy to a Layer 2 EtherChannel Interface:
Device> enable
Device# configure terminal
Device(config)# ipv6 source-guard policy POL
Device (config-sisf-sourceguard)# no validate address
Device((config-sisf-sourceguard)# validate prefix
Device(config-sisf-sourceguard)# exit
Device(config)# interface Po4
Device(config-if)# ipv6 snooping
Device(config-if)# ipv6 source-guard attach-policy POL
Device(config-if)# end
Binding entries can be removed from the binding table for various reasons: the switch may have reset, or you may have used the clear commands, and so on. The following example shows how you can use the data-glean recovery function to restore valid binding entries in the binding table.
The scenario used in this example involves interaction between the IPv6 Source Guard, IEEE 802.1x authentication, and SISF-based device-tracking features. Described below is the set-up we are using for this example, along with sample configuration, followed by a description of situations that can cause premature removal of valid entries from the binding table, and finally, the configuration that you must have in-place, for such entries to be restored.
The key aspects of this example set-up are outlined below:
An IPv6 Source Guard policy is configured and attached to an interface.
Device# show ipv6 source-guard policy src-guard-policy 
Source guard policy src-guard-policy configuration: 
  validate address
Policy src-guard-policy is applied on the following targets: 
Target               Type  Policy               Feature        Target range
Gi1/0/1              PORT  src-guard-policy     Source guard   vlan all
A custom SISF-based device-tracking policy, which allows gleaning of only DHCP packets and not NDP packets is attached to the same interface as the source guard policy.
This means that any host in the network can use only a DHCP-assigned IP address to communicate.
