---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-7
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [591, 695]
sha256: 6a7d85f955c15fb5e052f7e253ca7302e3f5134f9a4cdb63f3cfb4298c3a8dd4
---

# DHCP Client

| Property | Description | 
|---|---|
| **address** (*IPv6 prefix* ; Default: ) | IPv6 prefix that will be assigned to the client | 
| **allow-dual-stack-queue** (*yes \| no* ; Default:**yes** ) | Creates a single simple queue entry for both IPv4 and IPv6 addresses, uses the MAC address and DUID for identification. Requires IPv4 DHCP Server to have this option enabled as well to work properly. | 
| **comment** (*string* ; Default: ) | Short description of an item. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Whether an item is disabled | 
| **dhcp-option** (*string* ; Default: ) | Add additional DHCP options from the option list. | 
| **dhcp-option-set** (*string* ; Default: ) | Add an additional set of DHCP options. | 
| **life-time** (*time* ; Default:**3d** ) | The time period after which binding expires. | 
| **duid** (*hex string* ; Default: ) | DUID value. Should be specified only in hexadecimal format. | 
| **iaid** (*integer [0..4294967295]* ; Default: ) | Identity Association Identifier, part of the Client ID. | 
| **prefix-pool** (*string* ; Default: ) | Prefix pool that is being advertised to the DHCPv6 Client. | 
| **rate-limit** (*integer[/integer] [integer[/integer] [integer[/integer] [integer[/integer]]]]* ; Default: ) | Adds a dynamic simple queue to limit IP's bandwidth to a specified rate. Requires the lease to be static. Format is: rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time]]]]. All rates should be numbers with optional 'k' (1,000s) or 'M' (1,000,000s). If tx-rate is not specified, rx-rate is as tx-rate too. Same goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold are not specified (but burst-rate is specified), rx-rate and tx-rate is used as burst thresholds. If both rx-burst-time and tx-burst-time are not specified, 1s is used as default. | 
| **server** (*string \| all* ; Default:**all** ) | Name of the server. If set to **all** , then binding applies to all created DHCP-PD servers. | 

**Read-only properties**

| Property | Description | 
|---|---|
| **dynamic** (*yes \| no* ) | Whether an item is dynamically created. | 
| **expires-after** (*time* ) | The time period after which binding expires. | 
| **last-seen** (*time* ) | Time period since the client was last seen. | 
| **status** (*waiting \| offered \| bound* ) | Three status values are possible:  | 
| **reconfigure-key** (string) | Reconfiguration authentication key | 
| **reconfigure-last-sent**  (integer) | Count of sent Reconfigure ( *forcerenew* ) messages | 
| **reconfigure-status**  |  | 

For example, dynamically assigned /62 prefix

**Menu specific commands**

| Property | Description | 
|---|---|
| **make-static** () | Set dynamic binding as static. | 
| **send-reconfigure** (*id* ) | Send Reconfigure ( *forcerenew* ) message | 

### Rate limiting

It is possible to set the bandwidth to a specific IPv6 address by using DHCPv6 bindings. This can be done by setting a rate limit on the DHCPv6 binding itself, by doing this a dynamic simple queue rule will be added for the IPv6 address that corresponds to the DHCPv6 binding. By using the `rate-limit` the parameter you can conveniently limit a user's bandwidth.

For any queues to work properly, the traffic must not be FastTracked, make sure your Firewall does not FastTrack traffic that you want to limit.

First, make the DHCPv6 binding static, otherwise, it will not be possible to set a rate limit to a DHCPv6 binding:

Then you need can set a rate to a DHCPv6 binding that will create a new dynamic simple queue entry:

By default `allow-dual-stack-queue` is enabled, this will add a single dynamic simple queue entry for both DHCPv6 binding and DHCPv4 lease, without this option enabled separate dynamic simple queue entries will be added for IPv6 and IPv4.

If `allow-dual-stack-queue` is enabled, then a single dynamic simple queue entry will be created containing both IPv4 and IPv6 addresses:

## RADIUS Support

Since RouterOS v6.43 it is possible to use RADIUS to assign a rate-limit per DHCPv6 binding, to do so you need to pass the Mikrotik-Rate-Limit attribute from your RADIUS Server for your DHCPv6 binding. To achieve this you first need to set your DHCPv6 Server to use RADIUS for assigning bindings. Below is an example of how to set it up:

After that, you need to tell your RADIUS Server to pass the Mikrotik-Rate-Limit attribute. In case you are using FreeRADIUS with MySQL, then you need to add appropriate entries into **radcheck** and **radreply** tables for a MAC address, that is being used for your DHCPv6 Client. Below is an example for table entries:

By default allow-dual-stack-queue is enabled and will add a single dynamic queue entry if the MAC address from the IPv4 lease (or DUID, if the DHCPv4 Client supports `Node-specific Client Identifiers` from RFC4361), but DUID from DHCPv6 Client is not always based on the MAC address from the interface on which the DHCPv6 client is running on, DUID is generated on a per-device basis. For this reason, a single dynamic queue entry might not be created, separate dynamic queue entries might be created instead.

## Configuration Example

### Enabling IPv6 Prefix delegation

Let's consider that we already have a running DHCP server.

To enable IPv6 prefix delegation, first, we need to create an address pool:

Notice that prefix-length is 62 bits, which means that clients will receive /62 prefixes from the /60 pool.

The next step is to enable DHCP-PD:

To test our server we will set up wide-dhcpv6 on an ubuntu machine:

- install wide-dhcpv6-client
- edit "/etc/wide-dhcpv6/dhcp6c.conf" as above

You can use also RouterOS as a DHCP-PD client.

- Run DHCP-PD client:

- Verify that prefix was added to the:

- You can make binding to specific client static so that it always receives the same prefix:

- DHCP-PD also installs a route to assigned prefix into IPv6 routing table:

### Enabling IPv6 Address delegation

Address delegation on DHCPv6 server side works almost in the exact same way as when you configure prefix server. Only difference is that you must specify in configuration address-pool instead of prefix-pool and the pool used for this server must be defined to use /128 prefix-length. Of course, you can create server which only assigns static addresses and skip using the pool.

This configuration is already enough to work with DHCPv6 clients such as, for example, RouterOS client.

However, usually end-devices as computers do not know if their network is managed by DHCP server or not. That is why DHCPv6 server configuration is combined with SLAAC functionality. You can even avoid using SLAAC in order to advertise prefix for local network device, all you need to do is advertise "managed-address-configuration" option to your network devices.

Now, for example, your computer which will be connected to router ether2 interface will receive advertisement message from RouterOS ND configuration stating that this network is using "managed-address-configuration" which normally on end user devices will enable DHCPv6 client requesting IPv6 address.

Full configuration backup from the server with several comments is provided here.

Server configuration might vary based on operating systems used by clients. For example, macOS will get an address on initialisation with such configuration but might not renew lease after sleep, if prefix is set to "none", since macOS does not use DHCPv6 client without SLAAC address. Other clients might need also "autonomous" option to be set to "no" in order to trigger DHCPv6 client usage. This is just a configuration example - settings might need adjustments depending on client devices.

# DHCP Relay

## Summary

**Sub-menu:** `/ip dhcp-relay`

