---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-6
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [486, 590]
sha256: fe24a3f4ac8520dd060c03b0ee22f56277b955793d6ec34b3d2e97f0ffefcd20
---

# DHCP Client

It's possible to define two substring matches for the same option, for example, one that matches "ABC" and another one that matches "ABCDE", if there is a DHCP option with the value of "ABCDEF", both entries would match it, but the matcher that will get applied will be selected randomly.

Clients with a static lease will continue to receive their static address, even when matched by the option matcher.

### **Properties**

| Property | Description | 
|---|---|
| **address-pool** (*string \| static-only* ; Default:**static-only** ) | IP pool, from which to take IP addresses for the clients. If set to **static-only** , then only clients that have a static lease (added in the lease submenu) will be allowed. | 
| **code** (*integer:1..254* ; Default: ) | DHCP option code. All codes are available at http://www.iana.org/assignments/bootp-dhcp-parameters | 
| **comment** (*string* ; Default: ) | Short description of option matcher. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Whether an item is disabled | 
| **matching-type** (*exact \| substring* ; Default: ) | Matching method:  | 
| **name** (*string* ; Default: ) | Descriptive name for option matcher. | 
| **option-set** (*name \| none* ; Default:**none** ) | A custom set of DHCP options defined in the Option Sets menu. | 
| **server** (*string* \|*all* ; Default: ) | Server name which serves option matcher. | 
| **value** (*string* ; Default: ) | A value that will be searched for in option. Available data types for value are:  | 

### Option matcher examples

Match *dhcp1* server clients by *exact* Vendor class identifier (DHCP option 60) and assign address from the *pool1*:

Match clients on all DHCP servers by *exact* Client Id (DHCP option 61) configured as hex value and assign address from the *pool2*:

Match *dhcp2* server clients *partially* by Hostname (DHCP option 12) and assign address from the *pool3*:

## Configuration Examples

### Setup

To simply configure DHCP server you can use a `setup` command.

First, you configure an IP address on the interface:

Then you use `setup` a command which will automatically ask necessary parameters:

That is all. You have configured an active DHCP server.

### Manual configuration

To configure the DHCP server manually to respond to local requests you have to configure the following:

- An **IP pool** for addresses to be given out, make sure that your gateway/DHCP server address is not part of the pool.

- A **network**  indicating subnets that DHCP-server will lease addresses from, among other information, like a gateway, DNS-server, NTP-server, DHCP options, etc.

- In our case, the device itself is serving as the gateway, so we'll add the **address** to the bridge interface:

- And finally, add **DHCP Server** , here we will add the previously created address **pool** , and specify on which**interface** the DHCP server should work on

# DHCPv6 Server

## Summary

**Standards:** `RFC 3315, RFC 3633`

Single DUID is used for client and server identification, only IAID will vary between clients corresponding to their assigned interface.

Client binding creates a dynamic pool with a timeout set to binding's expiration time (note that now dynamic pools can have a timeout), which will be updated every time binding gets renewed.

When a client is bound to a prefix, the DHCP server adds routing information to know how to reach the assigned prefix.

## General

**Sub-menu:** `/ipv6 dhcp-server`

This sub-menu lists and allows to configure DHCP-PD servers.

## DHCPv6 Server Properties

| Property | Description | 
|---|---|
| **address-pool** (*enum \| static-only* ; Default:**static-only** ) | IPv6 pool, from which to take IPv6 address for the clients, pool prefix-length must be specified as /128. | 
| **prefix-pool (*enum \| static-only*; Default: static-only)** | IPv6 pool, from which to take IPv6 prefxies for the clients. | 
| **allow-dual-stack-queue** (*yes \| no* ; Default: **yes** ) | Creates a single simple queue entry for both IPv4 and IPv6 addresses, and uses the MAC address and DUID for identification. Requires IPv6 DHCP Server to have this option enabled as well to work properly. | 
| **address-lists** (*string* ; Default:) | Comma seperated list of address-lists. Address or prefix issued by the server will be added to these lists. | 
| **binding-script** (*string* ; Default: ) | A script that will be executed after binding is assigned or de-assigned. Internal "global" variables that can be used in the script:  | 
| **dhcp-option** (*string* ; Default:**none** ) | Add additional DHCP options from option list. | 
| **insert-queue-before** (*bottom \| first \| name* ; Default:**first** ) | Specify where to place dynamic simple queue entries for static DHCP leases with a rate-limit parameter set. | 
| **parent-queue** (*string \| none* ; Default:**none** ) | A dynamically created queue for this lease will be configured as a child queue of the specified parent queue. | 
| **preference** (*integer [0..255]* ; Default: 255) | Defines server priority level in client selection when multiple servers respond. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Whether DHCP-PD server participates in the prefix assignment process. | 
| **interface** (*string* ; Default: ) | The interface on which server will be running. | 
| **lease-time** (*time* ; Default:**3d** ) | The time that a client may use the assigned address. The client will try to renew this address after half of this time and will request a new address after the time limit expires. | 
| **rapid-commit (** yes \| no; Default: **yes)** | Enables a two-message exchange (Solicit and Reply) for quicker client configuration by skipping the standard four-message process. | 
| **route-distance**  (*integer [0..255]* ; Default:**1** ) | Specify distance to set for dynamically installed routes towards DHCPv6 clients. | 
| **use-radius** (*yes \| no \| accounting* ; Default:**no** ) | Whether to use RADIUS server:  | 
| **use-reconfigure**  (*yes \| no* ; Default:**no** ) | Allow the server to send Reconfigure messages to clients, prompting them to renew or update their configuration without waiting for their lease to expire. | 
| **name** (*string* ; Default: ) | Reference name | 
| **address-list**  (*string* ; Default:**none** ) | Address list to which address will be added if the lease is bound. | 
| **ignore-ia-na-bindings** (*yes \| no* ; Default:**no** ) | Do not reply to DHCPv6 address requests and process only prefixes. Without this setting even if server does not have address-pool configured, it has to respond to client that there is no address available for the client. That can lead up to the situation when DHCPv6 client requests address and prefix in a loop. | 

**Read-only Properties**

| Property | Description | 
|---|---|
| **dynamic** (*yes \| no* ) |  | 
| **invalid** (*yes \| no* ) |  | 

## Bindings

**Sub-menu:** `/ipv6 dhcp-server binding`

DUID is used only for dynamic bindings, so if it changes then the client will receive a different prefix than previously.

