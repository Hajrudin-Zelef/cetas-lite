---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-3
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "ascend", "ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [239, 297]
sha256: e25201c16ed687a382d3c91908e0e4cf537716f1c7fe3349a574848116534b09
---

# DHCP Client

- NAS-Identifier - router identity
- NAS-IP-Address - IP address of the router itself
- NAS-Port - the ID of the interface where the DHCP server is configured. This value is the same as the IF-MIB::ifIndex
- NAS-Port-Id - the name of the interface where the DHCP server is configured
- NAS-Port-Type - Ethernet
- Calling-Station-Id - client identifier (active-client-id)
- Framed-IP-Address - IP address of the client (active-address)
- Called-Station-Id - the name of DHCP server
- User-Name - MAC address of the client (active-mac-address)
- Password - " "

Access-Accept:

- Framed-IP-Address - IP address that will be assigned to a client
- Framed-Pool - IP pool from which to assign an IP address to a client
- Rate-Limit - Datarate limitation for DHCP clients. Format is: rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time][priority] [rx-rate-min[/tx-rate-min]]]]. All rates should be numbers with optional 'k' (1,000s) or 'M' (1,000,000s). If tx-rate is not specified, rx-rate is as tx-rate too. Same goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold are not specified (but burst-rate is specified), rx-rate and tx-rate are used as burst thresholds. If both rx-burst-time and tx-burst-time are not specified, 1s is used as default. Priority takes values 1..8, where 1 implies the highest priority, but 8 - the lowest. If rx-rate-min and tx-rate-min are not specified rx-rate and tx-rate values are used. The rx-rate-min and tx-rate-min values can not exceed rx-rate and tx-rate values.
- Ascend-Data-Rate - TX/RX data rate limitation if multiple attributes are provided, first limits tx data rate, second - RX data rate. If used together with Ascend-Xmit-Rate, specifies RX rate. 0 if unlimited
- Ascend-Xmit-Rate - tx data rate limitation. It may be used to specify the TX limit only instead of sending two sequential Ascend-Data-Rate attributes (in that case Ascend-Data-Rate will specify the receive rate). 0 if unlimited
- Session-Timeout - max lease time (lease-time)

DHCP server requires a real interface to receive raw ethernet packets. If the interface is a Bridge interface, then the Bridge must have a real interface attached as a port to that bridge which will receive the raw ethernet packets. It cannot function correctly on a dummy (empty bridge) interface.

## DHCP Server Properties

| Property | Description | 
|---|---|
| **add-arp** (*yes \| no* ; Default:**no** ) | Whether to add dynamic ARP entry. If set to **no** either ARP mode should be enabled on that interface or static ARP entries should be administratively defined in*/ip arp* submenu. | 
| **address-pool** (*string \| static-only* ; Default:**static-only** ) | IP pool, from which to take IP addresses for the clients. If set to **static-only** , then only the clients that have a static lease (added in the lease submenu) will be allowed. | 
| **allow-dual-stack-queue** (*yes \| no* ; Default:**yes** ) | Creates a single simple queue entry for both IPv4 and IPv6 addresses, and uses the MAC address and DUID for identification. Requires IPv6 DHCP Server to have this option enabled as well to work properly. | 
| **always-broadcast** (*yes \| no* ; Default:**no** ) | Changes whether to force broadcast DHCP replies:  | 
| **authoritative** (*after-10sec-delay \| after-2sec-delay \| yes \| no* ; Default:**yes** ) | Option changes the way how a server responds to DHCP requests:    **delay-threshold=x** setting should be used. | 
| **bootp-lease-time** (*forever \| lease-time \| time* ; Default:**forever** ) | Accepts two predefined options or time value:  | 
| **bootp-support** (*none \| static \| dynamic* ; Default:**static** ) | Support for BOOTP clients:  | 
| **client-mac-limit** (*integer \| unlimited* ; Default:**unlimited** ) | Specifies whether to limit a specific number of clients per single MAC address or leave unlimited. Note that this setting should not be used in relay setups. | 
| **conflict-detection** (*yes \| no* ; Default:**yes** ) | Allows disabling/enabling conflict detection. If the option is enabled, then whenever the server tries to assign a lease it will send ICMP and ARP messages to detect whether such an address in the network already exists. If any of the above get a reply address is considered already used. | 
| **delay-threshold** (*time \| none* ; Default:**none** ) | If the sec's field in the DHCP packet is smaller than the delay threshold, then this packet is ignored. If set to **none** - there is no threshold (all DHCP packets are processed) | 
| **dhcp-option-set** (*name \| none* ; Default:**none** ) | Use a custom set of DHCP options defined in the option sets menu. | 
| **dynamic-lease-identifiers** (*list of client-id, client-mac. opt-82* ; Default:**client-id,client-mac** ) | Specify which parameters to use and store when generating a dynamic DHCP lease. | 
| **insert-queue-before** (*bottom \| first \| name* ; Default:**first** ) | Specify where to place dynamic simple queue entries for static DHCP leases with a rate-limit parameter set. | 
| **interface** (*string* ; Default: ) | The interface on which the DHCP server will be running. | 
| **lease-script** (*string* ; Default:**""** ) | A script that will be executed after a lease is assigned or de-assigned. Internal "global" variables that can be used in the script:  | 
| **lease-time** (*time* ; Default:**3****0m** ) | The time that a client may use the assigned address. The client will try to renew this address after half of this time and will request a new address after the time limit expires. | 
| **name** (*string* ; Default: ) | Reference name | 
| **parent-queue** (*string \| none* ; Default:**none** ) | A dynamically created queue for this lease will be configured as a child queue of the specified parent queue. | 
| **relay** (*IP* ; Default:**0.0.0.0** ) | The IP address of the relay this DHCP server should process requests from:  | 
| **server-address** (*IP* ; Default:**0.0.0.0** ) | The IP address of the server to use in the next step of the client's bootstrap process (For example, to assign a specific server address in case several addresses are assigned to the interface) | 
| **support-broadband-tr101** (*yes \| no* ; Default:**no** ) | Enables or disables the inclusion of additional Option 82 suboptions (e.g. 0x81 actual upstream, 0x82 actual downstream) in RADIUS Access-Request and Accounting-Request messages as described in RFC 4679 and Broadband Forum TR-101. When enabled, the DHCP server includes specific suboptions under DHCP Option 82 (Relay Agent Information) that are used by RADIUS servers to identify subscriber line parameters in broadband access networks, especially those based on DSL infrastructure. This property only has an effect when `use-radius` is set to`yes` or`accounting` . | 
| **use-framed-as-classless** (*yes \| no* ; Default:**yes** ) | Forward RADIUS Framed-Route as a DHCP Classless-Static-Route to DHCP-client. Whenever both Framed-Route and Classless-Static-Route are received Classless-Static-Route is preferred. | 
| **use-radius** (*yes \| no \| accounting* ; Default:**no** ) | Whether to use RADIUS server:  | 
| **use-reconfigure**  (*yes \| no* ; Default:**no** ) | Allow the server to send Reconfigure ( *forcerenew)* messages to clients, prompting them to renew configuration without waiting for their lease to expire. | 

## Leases

**Sub-menu:** `/ip dhcp-server lease`

DHCP server lease submenu is used to monitor and manage server leases. The issued leases are shown here as dynamic entries. You can also add static leases to issue a specific IP address to a particular client (identified by MAC address).

Generally, the DHCP lease is allocated as follows:

