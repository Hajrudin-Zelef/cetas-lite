---
id: collect-261001-general-networking/general-networking/ms-port-and-vlan-configuration-switch-ports-19ab0597-2
title: "ms-port-and-vlan-configuration-switch-ports-19ab0597"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["energy", "ethernet", "voice"]
source: docs/RAG/collect-261001-general-networking/ms-port-and-vlan-configuration-switch-ports-19ab0597.md
source_anchor: ""
source_lines: [75, 129]
sha256: 6ed1d849193de2bf07c78e060814c3987f14c8f1e77d6f133d7673298811052a
---

# ms-port-and-vlan-configuration-switch-ports-19ab0597

  - VLAN: All traffic will be placed on this VLAN.
  - Voice VLAN: CDP/LLDP capable voice devices will be able to use this VLAN.
The option for MAC allow list and Sticky MAC allow list on MS390 and C9300-M switches requires CS 16+ firmware version. If the switches are running prior firmware versions (MS 15.5 and/or CS 15.5 or older), please contact Meraki Support in order to enable this feature set.
Note: In Catalyst IOS-XE you won't be able to configure the same MAC address in two different interfaces.
- Energy Efficient Ethernet (EEE): Enable/disable Energy Efficient Ethernet (EEE)
Searching for ports
The virtual stack allows an administrator to view all switch ports in one easy-to-navigate page. To further simplify switch port management, a dynamic search bar is available at the top to allow for quick searching of ports.
Search terms
- Enter any value in to the search omnibox for an instant search result
- Use conditional operators to separate multiple search queries (AND, OR)
- Use a wildcard to search for more general results ( * )
- Use a dash to exclude a search value ( - )
- Enter specific search terms to find a particular port:
Meraki is committed to providing an inclusive experience for our customers. The following section contains language that does not adhere to our standards for inclusivity. We are working with our partners/teams to replace it.
| Search Type | Search Value | Result | Example | 
|---|---|---|---|
| Port | port:value | return all specified ports or port ranges | port:1-10 | 
| Module (MS390 & C9300-M) | module:value  | return or exclude (using -) module model types  | module:8x10 (only 8x10 modules) -module:8x10 (all except 8x10 modules) -module:0 (excludes all modules models) | 
| Name | name:value | return all ports with the specified switch name | name:"joe's desktop" | 
| Switch | switch:value | return all ports for the designated switch(es) | switch:"1st floor" | 
| Detected Uplink  | is:uplink  | return interface(s) detected as uplink to Meraki Cloud  | is:uplink not:uplink | 
| Tags | tag:value | return all ports with the specified tag | tag:"blue 132" | 
| VLAN | vlan:value vlan:native vlan:voice | return all ports with the specified vlan return all ports with a native vlan return all ports with a voice vlan | vlan:"60" vlan:"native 60" vlan:"voice 20" | 
| LLDP | lldp:value | return all ports containing matching LLDP information | lldp:"MR24" | 
| Type | is:value | will return all ports with type "trunk" or type "access" | is:trunk | 
| Link | link:value | return all ports with the link type set to specified speed/duplex | link:"100 mbps" link:"10 gbps" | 
| Link Aggregate | is:aggregated | return only link aggregated (LACP) ports | is:"aggregated" | 
| Access Policy | ap:value | return all ports with the specified access policy applied (wildcard supported) | ap:* | 
| Port Schedule | schedule:value | return all ports with the specified port schedule (wildcard supported) | schedule:* | 
| Group  | group:value  | return all ports belonging to a common group (the virtual stack automatically categorizes the 3 most common configuration types into groups 1,2 and 3) | group:1 group:2 group:3 | 
| MAC Allow list  | mac_whitelist:*  | return all ports with a mac-allowlist enabled (you can substitute the * with a mac address value using colons as separators) | mac_whitelist:aa:bb:cc:dd:ee:ff mac_whitelist:* | 
The search tool is also capable of intelligently combining multiple search queries. See a few examples below.
Search: name:"joe's port" AND switch:"2nd floor POE"
Result: returns all port(s) with the name "joe's port" on the switch named "2nd floor POE"
Search: port:1-15 link:"10 gbps" switch:"2nd floor IDF"
Result: Returns all ports configured for 10gbit from the port range of 1-15 on the switch named "2nd floor IDF"
Link Aggregation
The MS switches support Link Aggregation (LACP) groups of up to 8 ports on the same switch or physical stack. A "Link Aggregate" is a combination of ports that act as one logical link. This is often referred to as Link Bonding, Link Aggregation, or EtherChannel. A link aggregate will load balance across the different physical links for additional performance, and will also give higher reliability because the link aggregate will continue to function as long as at least one of the physical links is working.
To configure an aggregate, simply choose the ports to be aggregated by checking their respective boxes (under Switching > Monitor > Switch Ports page) and then select the Aggregate option at the top of the page.
By default, link-aggregation groups are configured to run in LACP active mode.
Switches on MS firmware use an adaptive LACP active implementation, in which, the link-aggregation group is not completely suspended if the LACPDUs from the connected device are not received on the links. Instead, the switch allows one link in the aggregation to retain connectivity, to ensure that mis-configurations do not lead to either of the connected devices becoming isolated.
If you are setting up a link-aggregation with a device that does not support LACP, you can disable the Enforce LACP active option from the port configuration UI. 
Please note: This is NOT a supported configuration. For LAG to work properly, both devices should be configured to use LACP negotiation.
The option for controlling LACP enforcement mode is currently available on MS firmware only. To enable the UI for this feature on your Dashboard Organization, please reach out to Meraki Support.
When LACP active enforcement is disabled, the ports will continue to send out LACPDUs, to support compatibility with LACP passive. However, the links in the aggregation will not be suspended if LACPDUs are not received on the links.
MS390 and C9300 switches use a strict LACP active implementation. If the connected device is not configured for LACP, the link-aggregation is suspended as a whole. The LACP active configuration cannot be disabled on these switches.
Note: Disabling LACP enforcement can lead to link-state and traffic-forwarding inconsistencies, and is not recommended.
Note: It is generally recommended that ports are first aggregated and then physically connect the aggregated ports. Be sure to configure the aggregate (or have LACP enabled) on both ends of the link. Configure the downlink device first, wait for the config to state up to date, before configuring the aggregation (uplink) device. If the process is performed in the uplink side first, there may be an outage depending on the models of switches used. For c9300-M/MS390s, the process described must be followed to ensure the aggregation forms correctly.
Make sure both switch ports share the same configuration, including tags, prior to aggregating.
c9300-M / MS390 are limited to 128 LACP groups per standalone switch or switch stack (a stack of 8 switches is still limited to 128 LACP groups). There is no limit to the number of LACP groups on other platforms, provided there are enough member ports available in the switch or stack.
For example, in order to create 8 LACP groups, each including the maximum of 8 ports per group, the switch stack must have at least 64 ports.
Note: (applicable to any non c9300-M/MS390s) By default, prior to configuring LACP, the MS series runs an LACP Passive instance per port. This is to prevent loops when a bonded link is connected to a switch running the default configuration. Once LACP is configured, the MS will run an Active LACP instance with a 30-second update interval and will always send LACP frames along the configured links.
Note: For additional information on Link Aggregation and Load Balancing, please refer to this article.
Note: When configuring LACP between Meraki MS and Catalyst, it may be advantageous on the Catalyst switch to disable the feature "spanning-tree etherchannel guard misconfig" if there are issues with getting the LACP aggregate established.
Selecting Aggregate ports
