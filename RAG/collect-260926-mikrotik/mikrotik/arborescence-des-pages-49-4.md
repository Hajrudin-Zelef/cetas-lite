---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-4
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [298, 365]
sha256: d43d95e5b86dd7fa96a3dc2bfb8f3e92271e14caa5c04d0b1801a23c466c9b0b
---

# DHCP Client

- an unused lease is in the "waiting" state
- if a client asks for an IP address, the server chooses one
- if the client receives a statically assigned address, the lease becomes offered, and then bound with the respective lease time
- if the client receives a dynamic address (taken from an IP address pool), the router sends a ping packet and waits for an answer for 0.5 seconds. During this time, the lease is marked testing
- in the case where the address does not respond, the lease becomes offered and then bound with the respective lease time
- in other cases, the lease becomes busy for the lease time (there is a command to retest all busy addresses), and the client's request remains unanswered (the client will try again shortly)

A client may free the leased address. The dynamic lease is removed, and the allocated address is returned to the address pool. But the static lease becomes busy until the client reacquires the address.

IP addresses assigned statically are not probed!

| Property | Description | 
|---|---|
| **address** (*IP* ; Default:**0.0.0.0** ) | Specify IP address (or ip pool) for static lease. If set to **0.0.0.0** - a pool from the DHCP server will be used | 
| **address-list** (*string* ; Default:**none** ) | Address list to which address will be added if the lease is bound. | 
| **agent-circuit-id** (*hex string;*  Default**none***)* | If specified, value must match the Option 82 Agent Circuit ID suboption of the request. | 
| **agent-remote-id** (*hex string;*  Default**none***)* | If specified, value must match the Option 82 Agent Remote ID suboption of the request. | 
| **allow-dual-stack-queue** (*yes \| no* ; Default:**yes** ) | Creates a single simple queue entry for both IPv4 and IPv6 addresses, and uses the MAC address and DUID for identification. Requires IPv6 DHCP Server to have this option enabled as well to work properly. | 
| **always-broadcast** (*yes \| no* ; Default:**no** ) | Changes whether to force broadcast DHCP replies:  | 
| **block-access** (*yes \| no* ; Default:**no** ) | Block access for this client | 
| **client-id** (*string* ; Default:**none** ) | If specified, must match the DHCP 'client identifier' option of the request | 
| **dhcp-option** (*string* ; Default:**none** ) | Add additional DHCP options from option list. | 
| **dhcp-option-set** (*string* ; Default:**none** ) | Add an additional set of DHCP options. | 
| **insert-queue-before** (*bottom \| first \| name* ; Default:**first** ) | Specify where to place dynamic simple queue entries for static DHCP leases with rate-limit parameter set. | 
| **lease-time** (*time* ; Default:**0s** ) | Time that the client may use the address. If set to **0s** lease will never expire. | 
| **mac-address** (*MAC* ; Default:**00:00:00:00:00:00** ) | If specified, must match the MAC address of the client | 
| **parent-queue** (*string \| none* ; Default:**none** ) | A dynamically created queue for this lease will be configured as a child queue of the specified parent queue. | 
| **queue-type** (*default, ethernet-default, multi-queue-ethernet-default, pcq-download-default, synchronous-default, default-small, hotspot-default, only-hardware-queue, pcq-upload-default, wireless-default* ) | Queue type that can be assigned to the specific lease | 
| **rate-limit** (*integer[/integer] [integer[/integer] [integer[/integer] [integer[/integer]]]];* ; Default: ) | Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static. Format is: rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time]]]]. All rates should be numbers with optional 'k' (1,000s) or 'M' (1,000,000s). If tx-rate is not specified, rx-rate is as tx-rate too. Same goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold are not specified (but burst-rate is specified), rx-rate and tx-rate is used as burst thresholds. If both rx-burst-time and tx-burst-time are not specified, 1s is used as default. | 
| ***routes*** ([dst-address/mask] [gateway] [distance]; Default:***none*** ) | Routes that appear on the server when the client is connected. It is possible to specify multiple routes separated by commas. This setting will be ignored for OpenVPN. | 
| **server** (*string* ) | Server name which serves this client | 
| **use-src-mac**  (*yes \| no* ; Default:**no** ) | When this option is set server uses the source MAC address instead of the received CHADDR to assign the address. | 
| **status** (waiting \| testing \| declined \| offered \| bound \| authorizing \| conflict) | Shows the status of DHCP `lease` : | 

### Menu specific commands

| **check-status** (*id* ) | Check the status of a given busy (status is conflict or declined) dynamic lease, and free it in case of no response | 
| **make-static** (*id* ) | Convert a dynamic lease to a static one | 
| **send-reconfigure** (*id* ) | Send Reconfigure ( *forcerenew* ) message | 

### Store Configuration

**Sub-menu:** `/ip dhcp-server config`

**Store Leases On Disk:** The configuration of how often the DHCP leases will be stored on disk. If they would be saved on a disk on every lease change, a lot of disk writes would happen which is very bad for Compact Flash (especially, if lease times are very short). To minimize writes on disk, all changes are saved on disk every store-leases-disk seconds. Additionally, leases are always stored on disk on graceful shutdown and reboot.

Manual changes to leases - addition/removal of a static lease, removal of a dynamic lease will cause changes to be pushed for this lease to storage.

**Accounting:** The accounting parameter in the DHCP server configuration enables or disables accounting for DHCP leases. When accounting is enabled, the DHCP server logs information about IP address assignments and lease renewals. This information can be useful for tracking and monitoring network usage, analyzing traffic patterns, or generating reports on IP address allocations.

**Interim-update:** The interim-update parameter determines whether the DHCP server sends periodic updates to the accounting server during a lease. These updates provide information about the lease duration, usage, and other relevant details. Enabling interim updates allows for more accurate tracking of lease activity.

**Radius-password:** The radius-password parameter is used to set the password for the RADIUS (Remote Authentication Dial-In User Service) server. RADIUS is a networking protocol commonly used for providing centralized authentication, authorization, and accounting for network access. When configuring the DHCP server to communicate with a RADIUS server for authentication or accounting purposes, you need to specify the correct password to establish a secure connection. This parameter ensures that the DHCP server can authenticate with the RADIUS server using the specified password.

### Rate limiting

It is possible to set the bandwidth to a specific IPv4 address by using DHCPv4 leases. This can be done by setting a rate limit on the DHCPv4 lease itself, by doing this a dynamic simple queue rule will be added for the IPv4 address that corresponds to the DHCPv4 lease. By using the *rate-limit* parameter you can conveniently limit a user's bandwidth.

For any queues to work properly, the traffic must not be FastTracked, make sure your Firewall does not FastTrack traffic that you want to limit.


First, make the DHCPv4 lease static, otherwise, it will not be possible to set a rate limit to a DHCPv4 lease:


Then you can set a rate to a DHCPv4 lease that will create a new dynamic simple queue entry:

By default allow-dual-stack-queue is enabled, this will add a single dynamic simple queue entry for both DHCPv6 binding and DHCPv4 lease, without this option enabled separate dynamic simple queue entries will be added for IPv6 and IPv4.

