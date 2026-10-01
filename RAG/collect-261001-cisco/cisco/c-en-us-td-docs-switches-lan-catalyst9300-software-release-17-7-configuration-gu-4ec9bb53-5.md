---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53.md
source_anchor: ""
source_lines: [154, 194]
sha256: 7c6de1a85b9d63ee6a336bd449b9f0486daa95c23e971fb64a1c13493eb9f17a
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-7-configuration-gu-4ec9bb53

| Step 11 | default {device-role \| hop-limit {maximum \| minimum} \| managed-config-flag \| match {ipv6 access-list \| ra prefix-list } \| other-config-flag \| router-preference maximum\| trusted-port} Example: Device(config-nd-raguard)# default hop-limit | Restores a command to its default value. | 
| Step 12 | end Example: Device(config-nd-raguard)# end | Exits RA Guard policy configuration mode and returns to privileged EXEC mode. | 
| Step 13 | show ipv6 nd raguard policy policy_name Example: Device# show ipv6 nd raguard policy example_policy | (Optional) Displays the ND guard policy configuration. | 
| Note | For a network with both host-facing ports and router-facing ports, along with a RA guard policy configured with device-role host on host-facing ports or vlan, it is mandatory to configure a RA guard policy with device-role router on router-facing ports to allow the RA Guard feature to work properly. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Router Advertisement policy to an interface or to VLANs on the interface :
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface type number Example: Device(config)# interface gigabitethernet 1/1/4 | Specifies an interface type and identifier; enters the interface configuration mode. | 
| Step 4 | ipv6 nd raguard [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if)# ipv6 nd raguard attach-policy example_policyDevice(config-if)# ipv6 nd raguard attach-policy example_policy vlan 222,223,224Device(config-if)# ipv6 nd raguard vlan 222, 223,224 | Attaches the Neighbor Discovery Inspection policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Router Advertisement Guard Policy on an EtherChannel interface or VLAN:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface range type number Example: Device(config)# interface Port-channel 11 | Specifies the port-channel interface name assigned when the EtherChannel was created. Enters interface range configuration mode. | 
| Step 4 | ipv6 nd raguard [attach-policy policy_name [ vlan {vlan_ids \| add vlan_ids \| except vlan_ids \| none \| remove vlan_ids \| all} ] \| vlan [ {vlan_ids \| add vlan_ids \| exceptvlan_ids \| none \| remove vlan_ids \| all} ] Example: Device(config-if-range)# ipv6 nd raguard attach-policy example_policyDevice(config-if-range)# ipv6 nd raguard attach-policy example_policy vlan 222,223,224Device(config-if-range)# ipv6 nd raguard vlan 222, 223,224 | Attaches the RA Guard policy to the interface or the specified VLANs on that interface. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-if-range)# end | Exits interface range configuration mode and returns to privileged EXEC mode. | 
| Tip | Enter the show interfaces summary command in privileged EXEC mode for quick reference to interface names and types. | 
Beginning in privileged EXEC mode, follow these steps to attach an IPv6 Router Advertisement policy to VLANs regardless of interface:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | vlan configuration vlan_list Example: Device(config)# vlan configuration 335 | Specifies the VLANs to which the IPv6 RA Guard policy will be attached, and enters VLAN interface configuration mode. | 
| Step 4 | ipv6 dhcp guard [attach-policy policy_name] Example: Device(config-vlan-config)# ipv6 nd raguard attach-policy example_policy | Attaches the IPv6 RA Guard policy to the specified VLANs across all switch and stack interfaces. The default policy is attached if the attach-policy option is not used. | 
| Step 5 | end Example: Device(config-vlan-config)# end | Exits VLAN interface configuration mode and returns to privileged EXEC mode. | 
Beginning in privileged EXEC mode, follow these steps to configure an IPv6 DHCP (DHCPv6) Guard policy:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password, if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | ipv6 dhcp guard policy policy-name Example: Device(config)# ipv6 dhcp guard policy example_policy | Specifies the DHCPv6 Guard policy name and enters DHCPv6 Guard Policy configuration mode. | 
| Step 4 | device-role {client \| server} Example: Device(config-dhcp-guard)# device-role server | (Optional) Filters out DHCPv6 replies and DHCPv6 advertisements on the port that are not from a device of the specified role. Default is client.  | 
| Step 5 | match server access-list ipv6-access-list-name Example: ;;Assume a preconfigured IPv6 Access List as follows: Device(config)# ipv6 access-list my_acls Device(config-ipv6-acl)# permit host 2001:BD8:::1 any   ;;configure DCHPv6 Guard to match approved access list. Device(config-dhcp-guard)# match server access-list my_acls   | (Optional). Enables verification that the advertised DHCPv6 server or relay address is from an authorized server access list (The destination address in the access list is 'any'). If not configured, this check will be bypassed. An empty access list is treated as a permit all. | 
| Step 6 | match reply prefix-list ipv6-prefix-list-name Example:  ;;Assume a preconfigured IPv6 prefix list as follows: Device(config)# ipv6 prefix-list my_prefix permit 2001:DB8::/64 le 128  ;; Configure DCHPv6 Guard to match prefix Device(config-dhcp-guard)#  match reply prefix-list my_prefix  | (Optional) Enables verification of the advertised prefixes in DHCPv6 reply messages from the configured authorized prefix list. If not configured, this check will be bypassed. An empty prefix list is treated as a permit. | 
| Step 7 | preference{ max limit \| min limit } Example: Device(config-dhcp-guard)# preference max 250 Device(config-dhcp-guard)#preference min 150 | Configure max and min when device-role is serverto filter DCHPv6 server advertisements by the server preference value. The defaults permit all advertisements. max limit—(0 to 255) (Optional) Enables verification that the advertised preference (in preference option) is less than the specified limit. Default is 255. If not specified, this check will be bypassed. min limit—(0 to 255) (Optional) Enables verification that the advertised preference (in preference option) is greater than the specified limit. Default is 0. If not specified, this check will be bypassed. | 
| Step 8 | trusted-port Example: Device(config-dhcp-guard)# trusted-port | (Optional) trusted-port—Sets the port to a trusted mode. No further policing takes place on the port. | 
| Step 9 | default {device-role \| trusted-port} Example: Device(config-dhcp-guard)# default device-role | (Optional) default—Sets a command to its defaults. | 
