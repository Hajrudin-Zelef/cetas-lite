---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [70, 114]
sha256: 2f846ad3946f526ae444ee98d5defabd14e125a04dbef038605b5421435fd11d
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

| Note | To configure Layer 2 parameters, if the interface is in Layer 3 mode, you must enter the switchport interface configuration command without any parameters to change the interface into Layer 2 mode. This shuts down the interface and then re-enables it, which might generate messages on the device to which the interface is connected. When change the interface mode from Layer 3 to Layer 2 mode, the previous configuration information related to the affected interface might be lost, and the interface is returned to its default configuration. The command prompt displays as (config-if)# in Switchport configuration mode. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Snooping policy on an EtherChannel interface or VLAN:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface range interface_name Example: Device(config)#  interface range Port-channel 11  | Specifies the port-channel interface name assigned when the EtherChannel was created. Enters the interface range configuration mode. | 
| Step 4 | ipv6 snooping [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if-range)# ipv6 snooping attach-policy example_policyDevice(config-if-range)# ipv6 snooping attach-policy example_policy vlan 222,223,224Device(config-if-range)# ipv6 snooping vlan 222, 223,224 | Attaches the IPv6 Snooping policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if-range)# end | Exits interface range configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show running-config interfaceportchannel_interface_name Example: Device# show running-config interface portchannel 11 | Confirms that the policy is attached to the specified interface. | 
| Tip | Enter the show interfaces summary command for quick reference to interface names and types. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Snooping Policy to VLANs across multiple interfaces:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | vlan configuration vlan_list Example: Device(config)# vlan configuration 333     | Specifies the VLANs to which the IPv6 Snooping policy will be attached, and enters the VLAN interface configuration mode. | 
| Step 4 | ipv6 snooping [attach-policy policy_name] Example: Device(config-vlan-config)#ipv6 snooping attach-policy example_policy | Attaches the IPv6 Snooping policy to the specified VLANs across all device interfaces. The default policy is attached if the attach-policy option is not used. The default policy is, security-level guard, device-role node, protocol ndp and dhcp. | 
| Step 5 | end Example: Device(config-vlan-config)# end | Exits VLAN interface configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to configure IPv6 Binding Table Content :
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | [no] ipv6 neighbor binding [vlan vlan-id {ipv6-address interface interface_type stack/module/port hw_address [reachable-lifetimevalue [seconds \| default \| infinite] \| [tracking{ [default \| disable] [ reachable-lifetimevalue [seconds \| default \| infinite] \| [enable [reachable-lifetimevalue [seconds \| default \| infinite] \| [retry-interval {seconds\| default [reachable-lifetimevalue [seconds \| default \| infinite] } ] Example: Device(config)# ipv6 neighbor binding  | Adds a static entry to the binding table database. | 
| Step 4 | [no] ipv6 neighbor binding max-entries number [mac-limit number \| port-limit number [mac-limit number] \| vlan-limit number [ [mac-limit number] \| [port-limit number [mac-limitnumber] ] ] ] Example: Device(config)# ipv6 neighbor binding max-entries 30000 | Specifies the maximum number of entries that are allowed to be inserted in the binding table cache. | 
| Step 5 | ipv6 neighbor binding logging Example: Device(config)# ipv6 neighbor binding logging   | Enables the logging of binding table main events. | 
| Step 6 | exit Example: Device(config)# exit | Exits global configuration mode and returns to privileged EXEC mode. | 
| Step 7 | show ipv6 neighbor binding Example: Device# show ipv6 neighbor binding | Displays contents of a binding table. | 
Starting with Cisco IOS XE Amsterdam 17.1.1 the IPv6 ND Inspection feature is deprecated and the SISF- based device tracking feature replaces it and offers the same capabilities. For the corresponding replacement task, see Creating a Custom Device Tracking Policy with Custom Settings under the Configuring SISF-Based Device Tracking chapter in this document.
Beginning in privileged EXEC mode, follow these steps to configure an IPv6 ND Inspection Policy:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | ipv6 nd inspection policy policy-name Example: Device(config)# ipv6 nd inspection policy example_policy | Specifies the ND inspection policy name and enters ND Inspection Policy configuration mode. | 
| Step 4 | device-role {host \| switch} Example: Device(config-nd-inspection)# device-role switch | Specifies the role of the device attached to the port. The default is host. | 
| Step 5 | limit address-count value Example: Device(config-nd-inspection)# limit address-count 1000 | Limits the number of IPv6 addresses allowed to be used on the port. | 
| Step 6 | tracking {enable [reachable-lifetime {value \| infinite}] \| disable [stale-lifetime {value \| infinite}]} Example: Device(config-nd-inspection)# tracking disable stale-lifetime infinite | Overrides the default tracking policy on a port. | 
| Step 7 | trusted-port Example: Device(config-nd-inspection)# trusted-port | Configures a port to become a trusted port. | 
| Step 8 | validate source-mac Example: Device(config-nd-inspection)# validate source-mac | Checks the source media access control (MAC) address against the link-layer address. | 
| Step 9 | no {device-role \| limit address-count \| tracking \| trusted-port \| validate source-mac} Example: Device(config-nd-inspection)# no validate source-mac | Removes the current configuration of a parameter with the no form of the command. | 
| Step 10 | default {device-role \| limit address-count \| tracking \| trusted-port \| validate source-mac} Example: Device(config-nd-inspection)# default limit address-count | Restores configuration to the default values. | 
| Step 11 | end Example: Device(config-nd-inspection)# end | Exits ND Inspection Policy configuration mode and returns to privileged EXEC mode. | 
| Step 12 | show ipv6 nd inspection policy policy_name Example: Device# show ipv6 nd inspection policy example_policy | Verifies the ND inspection configuration. | 
