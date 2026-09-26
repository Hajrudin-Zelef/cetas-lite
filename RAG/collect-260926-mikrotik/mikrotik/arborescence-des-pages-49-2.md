---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-2
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [106, 238]
sha256: 2655e9705ae7fdc25955f54236e68536aa136422b831682f77f89cd94e1d2847
---

# DHCP Client

In some cases, administrators tend to set the 'router' option which cannot be resolved with offered IP's subnet. For example, the DHCP server offers 192.168.88.100/24 to the client, and option 3 is set to 172.16.1.1. This will result in an unresolved default route:

To fix this we need to add /32 route to resolve the gateway over ether1, which can be done by the running script below each time the DHCP client gets an address

Now we can further extend the script, to check if the address already exists, and remove the old one if changes are needed

# DHCPv6 Client

## Summary

**Sub-menu:** `/ipv6 dhcp-client`

DHCP-client in RouterOS is capable of being a DHCPv6-client and DHCP-PD client. So it is able to get a prefix from the DHCP-PD server as well as the DHCPv6 stateful address from the DHCPv6 server.

## Properties

| Property | Description | 
|---|---|
| **add-default-route** (*yes \| no* ; Default:**no** ) | Whether to add default IPv6 route after a client connects. | 
| **allow-reconfigure** (*yes \| no* ; Default:**no** )t | Allows to receive Reconfigure (forcerenew) message from DHCP server. | 
| **default-route-tables** (*table:distance* ; Default: main) | List of routing tables to which default route must be added. Table name can be proceeded with ":x" where x would be the distance for the route to be installed with. | 
| **comment** (*string* ; Default: ) | Short description of the client | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **interface** (*string* ; Default: ) | The interface on which the DHCPv6 client will be running. | 
| **pool-name** (*string* ; Default: ) | Name of the IPv6 pool in which received IPv6 prefix will be added | 
| **pool-prefix-length** (*integer* ; Default: ) | Prefix length parameter that will be set for IPv6 pool in which received IPv6 prefix is added. Prefix length must be greater or equal as the length of received prefix. If value is unset, then prefix will be added with the length as received from the server. | 
| **prefix-address-lists**  (*string* ; Default: ) | Names of the firewall address lists to which received prefix will be added. | 
| **prefix-hint** (*string* ; Default: ) | Include a preferred prefix length. | 
| **request** (*prefix, address* ; Default: ) | to choose if the DHCPv6 request will ask for the address or the IPv6 prefix, or both. | 
| **script** (*string* ; Default: ) | Run this script on the DHCP-client status change. Available variables:  | 
| **use-peer-dns** (*yes \| no* ; Default:**yes** ) | Whether to accept the DNS settings advertised by the IPv6 DHCP Server. | 
| **custom-duid** (*hex string* ; Default: ) | Allow to specify custom DUID. | 
| **use-interface-duid**  (*yes \| no* ; Default:**no** ) | According to RFC DHCPv6 client generate DUID based on the router first interface MAC address, not the interface on which client is configured. With this option enabled you will override this requirement and MAC address from "interface" specified on client will be used. | 
| **validate-server-duid** (*yes \| no* ; Default:**yes** ) | Allow to ignore incorrectly formed DUID provided by DHCPv6 server. Still checks that minimal DUID length is correct. | 
| **custom-iapd-id**  (integer; Default: ) | Allow to specify custom IAPD ID. | 
| **custom-iana-id**  (*integer* ; Default: ) | Allow to specify custom IANA ID. | 

**Read-only properties**

| Property | Description | 
|---|---|
| **duid** (*string* ) | Auto-generated DUID that is sent to the server. DUID is generated using one of the MAC addresses available on the router. | 
| **request** (*list* ) | specifies what was requested - prefix, address, or both. | 
| **dynamic** (*yes \| no* ) |  | 
| **expires-after** (*time* ) | A time when the IPv6 prefix expires (specified by the DHCPv6 server). | 
| **invalid** (*yes \| no* ) | Shows whether a configuration is invalid. | 
| **prefix** (*IPv6 prefix* ) | Shows received IPv6 prefix from DHCPv6-PD server | 
| **status** (*stopped \| searching \| requesting... \| bound \| renewing \| rebinding \| error \| stopping* ) | Shows the status of DHCPv6 Client:  | 
| **reconfigure-key**  (string) | Reconfiguration authentication key | 
| **reconfigure-last-counter** (integer) | Count of recieved forcerenew messages | 

**Menu specific commands**

| Property | Description | 
|---|---|
| **release** (*numbers* ) | Release current binding and restart DHCPv6 client | 
| **renew** (*numbers* ) | Renew current leases. If the renewal operation was not successful, the client tries to reinitialize the lease (i.e. it starts the lease request procedure (rebind) as if it had not received an IP address yet) | 

## Script

It is possible to add a script that will be executed when a prefix or an address is acquired and applied or expires and is removed using the DHCP client. There are separated sets of variables that will have the value set by the client depending on prefix or address status change as the client can acquire both and each of them can have a different effect on the router configuration.

Available variables for dhcp-client

- pd-valid - value - 1 or 0 - if prefix is acquired and it is applied or not
- pd-prefix - value ipv6/num (ipv6 prefix with mask) - the prefix inself
- na-valid - value - 1 or 0 - if address is acquired and it is applied or not
- na-address - value - ipv6 address - the address

## IAID

To determine what IAID will be used, convert the internal ID of an interface on which the DHCP client is running from hex to decimal.

For example, the DHCP client is running on interface PPPoE-out1. To get internal ID use the following command:

Now convert hex value 15 to decimal and you get IAID=21

## Configuration Examples

### Simple DHCPv6 client

This simple example demonstrates how to enable dhcp client to receive IPv6 prefix and add it to the pool.

Detailed print should show status of the client and we can verify if prefix is received

Notice that server gave us prefix 2a02:610:7501:ff04::/62 . And it should be also added to ipv6 pools

It works! Now you can use this pool, for example, for pppoe clients.

### Use received prefix for local RA

Consider following setup:

- ISP is routing prefix 2001:DB8::/62 to the router R1
- Router R1 runs DHCPv6 server to delegate /64 prefixes to the customer routers CE1 CE2
- DHCP client on routers CE1 and CE2 receives delegated /64 prefix from the DHCP server (R1).
- Client routers uses received prefix to set up RA on the local interface

**Configuration**

**R1**

**CE1**

**CE2**

**Check the status**

After configuration is complete we can verify that each CE router received its own prefix

On server:

On client:

We can also see that IPv6 address was automatically added from the prefix pool:

And pool usage shows that 'Address' is allocating the pool

# DHCP Server

## Summary

The DHCP (Dynamic Host Configuration Protocol) is used for the easy distribution of IP addresses in a network. The MikroTik RouterOS implementation includes both server and client parts and is compliant with RFC 2131.

The router supports an individual server for each Ethernet-like interface. The MikroTik RouterOS DHCP server supports the basic functions of giving each requesting client an IP address/netmask lease, default gateway, domain name, DNS-server(s) and WINS-server(s) (for Windows clients) information (set up in the DHCP networks submenu)

In order for the DHCP server to work, IP pools must also be configured (do not include the DHCP server's own IP address into the pool range) and the DHCP networks.

It is also possible to hand out leases for DHCP clients using the RADIUS server; the supported parameters for a RADIUS server are as follows:


Access-Request:

