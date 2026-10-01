---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-6
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [195, 251]
sha256: 89a70a9e1c61590e8057564016cca77fc913b6737ab0824397858da966774a0d
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

| Step 10 | end Example: Device(config-dhcp-guard)# end | Exits DHCPv6 Guard Policy configuration mode and returns to privileged EXEC mode. | 
| Step 11 | show ipv6 dhcp guard policy policy_name Example: Device# show ipv6 dhcp guard policy example_policy | (Optional) Displays the configuration of the IPv6 DHCP guard policy. Omitting the policy_name variable displays all DHCPv6 policies. | 
| Note | If you configure a trusted port then the device-role option is not available. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface type number Example: Device(config)#  interface gigabitethernet 1/1/4     | Specifies an interface type and identifier, and enters interface configuration mode. | 
| Step 4 | ipv6 dhcp guard [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if)# ipv6 dhcp guard attach-policy example_policyDevice(config-if)# ipv6 dhcp guard attach-policy example_policy vlan 222,223,224Device(config-if)# ipv6 dhcp guard vlan 222, 223,224 | Attaches the DHCP Guard policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 DHCP Guard policy on an EtherChannel interface or VLAN:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface range Interface_name Example: Device(config)#  interface Port-channel 11 | Specify the port-channel interface name assigned when the EtherChannel was created. Enters interface range configuration mode. | 
| Step 4 | ipv6 dhcp guard [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if-range)# ipv6 dhcp guard attach-policy example_policyDevice(config-if-range)# ipv6 dhcp guard attach-policy example_policy vlan 222,223,224Device(config-if-range)# ipv6 dhcp guard vlan 222, 223,224 | Attaches the DHCP Guard policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if-range)# end | Exits interface range configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 DHCP Guard policy to VLANs across multiple interfaces:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | vlan configuration vlan_list Example: Device(config)# vlan configuration 334 | Specifies the VLANs to which the IPv6 Snooping policy will be attached, and enters VLAN interface configuration mode. | 
| Step 4 | ipv6 dhcp guard [attach-policy policy_name] Example: Device(config-vlan-config)# ipv6 dhcp guard attach-policy example_policy | Attaches the IPv6 Neighbor Discovery policy to the specified VLANs across all switch and stack interfaces. The default policy is attached if the attach-policy option is not used. The default policy is, device-role client, no trusted-port. | 
| Step 5 | end Example: Device(config-vlan-config)# end | Exits VLAN interface configuration mode and returns to privileged EXEC mode. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | ipv6 source-guard policy policy_name Example: Device(config)# ipv6 source-guard policy example_policy | Specifies the IPv6 Source Guard policy name and enters IPv6 Source Guard policy configuration mode. | 
| Step 4 | [deny global-autoconf] [permit link-local] [default{. . . }] [exit] [no{. . . }] Example: Device(config-sisf-sourceguard)# deny global-autoconf | (Optional) Defines the IPv6 Source Guard policy.  | 
| Step 5 | end Example: Device(config-sisf-sourceguard)# end | Exits of IPv6 Source Guard policy configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show ipv6 source-guard policy policy_name Example: Device# show ipv6 source-guard policy example_policy | Shows the policy configuration and all the interfaces where the policy is applied. | 
| Note | Trusted option under source guard policy is not supported. | 
Apply the IPv6 Source Guard policy to an interface.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface type number Example: Device(config)# interface gigabitethernet 1/1/4     | Specifies an interface type and identifier; enters interface configuration mode. | 
| Step 4 | ipv6 source-guard [ attach-policy <policy_name> ] Example: Device(config-if)# ipv6 source-guard attach-policy example_policy | Attaches the IPv6 Source Guard policy to the interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show ipv6 source-guard policy policy_name Example: Device#(config)# show ipv6 source-guard policy example_policy | Shows the policy configuration and all the interfaces where the policy is applied. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface port-channel port-channel-number Example: Device(config)# interface Port-channel 4 | Specifies an interface type and port number and places the switch in the port channel configuration mode. | 
| Step 4 | ipv6 source-guard [ attach-policy <policy_name> ] Example: Device(config-if)# ipv6 source-guard attach-policy example_policy | Attaches the IPv6 Source Guard policy to the interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
| Step 6 | show ipv6 source-guard policy policy_name Example: Device# show ipv6 source-guard policy example_policy | Shows the policy configuration and all the interfaces where the policy is applied. | 
| Note | To allow routing protocol control packets sourced by a link-local address when prefix guard is applied, enable the permit link-local command in the source-guard policy configuration mode. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
