---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [115, 153]
sha256: b7868f0b02872fe8cec29b2712a975b9a338062d60e5aeb1a4521d3c14ac2a12
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

Starting with Cisco IOS XE Amsterdam 17.1.1 the IPv6 ND Inspection feature is deprecated and the SISF- based device tracking feature replaces it and offers the same capabilities. For the corresponding replacement task, see Attaching a Device Tracking Policy to an Interface under the Configuring SISF-Based Device Tracking chapter in this document.
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 ND Inspection policy to an interface or VLANs on an interface :
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface interface-type interface-number Example: Device(config)#  interface gigabitethernet 1/1/4     | Specifies an interface type and identifier; enters the interface configuration mode. | 
| Step 4 | ipv6 nd inspection [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if)# ipv6 nd inspection attach-policy example_policyDevice(config-if)# ipv6 nd inspection attach-policy example_policy vlan 222,223,2Device(config-if)# ipv6 nd inspection vlan 222, 223,224 | Attaches the Neighbor Discovery Inspection policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Neighbor Discovery Inspection policy on an EtherChannel interface or VLAN:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface range interface_name Example: Device(config)# interface range Port-channel 11     | Specifies the port-channel interface name assigned when the EtherChannel was created. Enters interface range configuration mode. | 
| Step 4 | ipv6 nd inspection [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if-range)# ipv6 nd inspection attach-policy example_policyDevice(config-if-range)# ipv6 nd inspection vlan 222, 223,224Device(config-if-range)# ipv6 nd inspection attach-policy example_policy vlan 222,223,224 | Attaches the ND Inspection policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if-range)# end | Exits interface range configuration mode and returns to privileged EXEC mode. | 
Starting with Cisco IOS XE Amsterdam 17.1.1 the IPv6 ND Inspection feature is deprecated and the SISF- based device tracking feature replaces it and offers the same capabilities. For the corresponding replacement task, see Attaching a Device Tracking Policy to a VLAN under the Configuring SISF-Based Device Tracking chapter in this document.
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 ND Inspection policy to VLANs across multiple interfaces:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | vlan configuration vlan_list Example: Device(config)# vlan configuration 334     | Specifies the VLANs to which the IPv6 Snooping policy will be attached, and enters VLAN interface configuration mode. | 
| Step 4 | ipv6 nd inspection [attach-policy policy_name] Example: Device(config-vlan-config)#ipv6 nd inspection attach-policy example_policy | Attaches the IPv6 Neighbor Discovery policy to the specified VLANs across all switch and stack interfaces. The default policy is attached if the attach-policy option is not used. The default policy is, device-role host, no drop-unsecure, limit address-count disabled, sec-level minimum is disabled, tracking is disabled, no trusted-port, no validate source-mac. | 
| Step 5 | end Example: Device(config-vlan-config)# end | Exits VLAN interface configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to configure an IPv6 Router Advertisement policy :
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | ipv6 nd raguard policy policy-name Example: Device(config)# ipv6 nd raguard policy example_policy | Specifies the RA guard policy name and enters RA guard policy configuration mode. | 
| Step 4 | [no]device-role {host \| monitor \| router \| switch} Example: Device(config-nd-raguard)# device-role switch | Specifies the role of the device attached to the port. The default is host. | 
| Step 5 | hop-limit {maximum \| minimum} value Example: Device(config-nd-raguard)# hop-limit maximum 33 | Enables filtering of Router Advertisement messages by the Hop Limit value. A rogue RA message may have a low Hop Limit value (equivalent to the IPv4 Time to Live) that when accepted by the host, prevents the host from generating traffic to destinations beyond the rogue RA message generator. An RA message with an unspecified Hop Limit value is blocked. (1–255) Range for Maximum and Minimum Hop Limit values. If not configured, this filter is disabled. Configure minimum to block RA messages with Hop Limit values lower than the value you specify. Configure maximumto block RA messages with Hop Limit values greater than the value you specify. | 
| Step 6 | managed-config-flag {off \| on} Example: Device(config-nd-raguard)# managed-config-flag on | Enables filtering of Router Advertisement messages by the managed address configuration, or "M" flag field. A rouge RA message with an M field of 1 can cause a host to use a rogue DHCPv6 server. If not configured, this filter is disabled. On: Accepts and forwards RA messages with an M value of 1, blocks those with 0. Off: Accepts and forwards RA messages with an M value of 0, blocks those with 1. | 
| Step 7 | match {ipv6 access-list list \| ra prefix-list list} Example: Device(config-nd-raguard)# match ipv6 access-list example_list | Matches a specified prefix list or access list. | 
| Step 8 | other-config-flag {on \| off} Example: Device(config-nd-raguard)# other-config-flag on  | Enables filtering of Router Advertisement messages by the Other Configuration, or "O" flag field. A rouge RA message with an O field of 1 can cause a host to use a rogue DHCPv6 server. If not configured, this filter is disabled. On: Accepts and forwards RA messages with an O value of 1, blocks those with 0. Off: Accepts and forwards RA messages with an O value of 0, blocks those with 1. | 
| Step 9 | [no]router-preference maximum {high \| medium \| low} Example: Device(config-nd-raguard)# router-preference maximum high  | Enables filtering of Router Advertisement messages by the router preference flag. If not configured, this filter is disabled.  | 
| Step 10 | trusted-port Example: Device(config-nd-raguard)# trusted-port | When configured as a trusted port, all attached devices are trusted, and no further message verification is performed. | 
