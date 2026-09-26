---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-5
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [366, 485]
sha256: 40bdc1eba7501f47a7f8ce7440a6cdb9cb5c6d657d1e3d6aecd5f7af846186c7
---

# DHCP Client

If *allow-dual-stack-queue* is enabled, then a single dynamic simple queue entry will be created containing both IPv4 and IPv6 addresses:

## Network

**Sub-menu:** `/ip dhcp-server network`

**Properties**

| Property | Description | 
|---|---|
| **address** (*IP/netmask* ; Default: ) | the network DHCP server(s) will lease addresses from | 
| **boot-file-name** (*string* ; Default: ) | Boot filename | 
| **caps-manager** (*string* ; Default: ) | A comma-separated list of IP addresses for one or more CAPsMAN system managers. DHCP Option 138 (capwap) will be used. | 
| **dhcp-option** (*string* ; Default: ) | Add additional DHCP options from the option list. | 
| **dhcp-option-set** (*string* ; Default: ) | Add an additional set of DHCP options. | 
| **dns-none** (*yes \| no* ; Default:**no** ) | If set, then DHCP Server will not pass DNS servers configured on the router to the DHCP clients. | 
| **dns-server** (*string* ; Default: ) | DNS servers that will be passed to DHCP clients. Two comma-separated DNS servers can be specified to be used by the DHCP client as primary and secondary DNS servers By default, if there are no DNS servers configured, then the router dynamic DNS Servers from IP>DNS will be passed to DHCP clients, if there are no dynamic DNS server configured, the router static DNS servers from IP>DNS will be passed to DHCP clients. | 
| **domain** (*string* ; Default: ) | The DHCP client will use this as the 'DNS domain' setting for the network adapter. | 
| **gateway** (*IP* ; Default:**0.0.0.0** ) | The default gateway to be used by DHCP Client. | 
| **netmask** (*integer: 0..32* ; Default:**0** ) | The actual network mask is to be used by the DHCP client. If set to '0' - netmask from network address will be used. | 
| **next-server** (*IP* ; Default: ) | The IP address of the next server to use in bootstrap. | 
| **ntp-none** (*yes \| no* ; Default:**no** ) | If set, then DHCP Server will not pass NTP servers configured on the router to the DHCP clients. | 
| **ntp-server** (*IP* ; Default: ) | the DHCP client will use these as the default NTP servers. Two comma-separated NTP servers can be specified to be used by the DHCP client as primary and secondary NTP servers | 
| **wins-server** (*IP* ; Default: ) | The Windows DHCP client will use these as the default WINS servers. Two comma-separated WINS servers can be specified to be used by the DHCP client as primary and secondary WINS servers | 

## RADIUS Support

Since RouterOS v6.43 it is possible to use RADIUS to assign a rate limit per lease, to do so you need to pass the Mikrotik-Rate-Limit attribute from your RADIUS Server for your lease. To achieve this you first need to set your DHCPv4 Server to use RADIUS for assigning leases. Below is an example of how to set it up:

After that, you need to tell your RADIUS Server to pass the Mikrotik-Rate-Limit attribute. In case you are using FreeRADIUS with MySQL, then you need to add appropriate entries into **radcheck** and **radreply** tables for a MAC address, that is being used for your DHCPv4 Client. Below is an example for table entries:

## Alerts

To find any rogue DHCP servers as soon as they appear in your network, the DHCP Alert tool can be used. It will monitor the interface for all DHCP replies and check if this reply comes from a valid DHCP server. If a reply from an unknown DHCP server is detected, an alert gets triggered:

When the system alerts about a rogue DHCP server, it can execute a custom script.

As DHCP replies can be unicast, the rogue DHCP detector may not receive any offer to other DHCP clients at all. To deal with this, the rogue DHCP detector acts as a DHCP client as well - it sends out DHCP discover requests once a minute.

The DHCP alert is not recommended on devices that are configured as DHCP clients. Since the alert itself generates DHCP discovery packets, it can affect the operation of the DHCP client itself. Use this feature only on devices that are DHCP servers or using a static IP address.

**Sub-menu:** `/ip dhcp-server alert`

**Properties**

| Property | Description | 
|---|---|
| **alert-timeout**  (none \| time; Default: 1h) | Time after which the alert will be forgotten. If after that time the same server is detected, a new alert will be generated. If set to **none** timeout will never expire. | 
| **interface** (*string* ; Default: ) | Interface, on which to run rogue DHCP server finder. | 
| **on-alert** (*string* ; Default: ) | Script to run, when an unknown DHCP server is detected. | 
| **valid-server**  (*string* ; Default: ) | List of MAC addresses of valid DHCP servers. | 

**Read-only properties**

| Property | Description | 
|---|---|
| **unknown-server**  (*string* ) | List of MAC addresses of detected unknown DHCP servers. The server is removed from this list after alert-timeout | 

**Menu specific commands**

| Property | Description | 
|---|---|
| **reset-alert**  (*id* ) | Clear all alerts on an interface | 

## DHCP Options

**Sub-menu:** `/ip dhcp-server option`

With the help of the DHCP Option list, it is possible to define additional custom options for DHCP Server to advertise. Option precedence is as follows:

- radius,
- lease,
- server,
- network.

This is the order in which the client option request will be filled in.

According to the DHCP protocol, a parameter is returned to the DHCP client only if it requests this parameter, specifying the respective code in the DHCP request Parameter-List (code 55) attribute. If the code is not included in the Parameter-List attribute, the DHCP server will not send it to the DHCP client, but **since RouterOS v7.1rc5 it is possible to force the DHCP option** from the server-side even if the DHCP-client does not request such parameter:

**Properties**

| Property | Description | 
|---|---|
| **code** (*integer:1..254* ; Default: ) | dhcp option code. All codes are available at http://www.iana.org/assignments/bootp-dhcp-parameters | 
| **name** (*string* ; Default: ) | Descriptive name of the option | 
| **value** (*string* ; Default: ) | Parameter's value. Available data types for options are:  RouterOS has predefined variables that can be used:  Now it is also possible to combine data types into one, for example: "0x01'vards'$(HOSTNAME)" For example if HOSTNAME is 'kvm', then raw value will be 0x0176617264736b766d. | 
| **raw-value** (*HEX string* ) | Read-only field which shows raw DHCP option value (the format actually sent out) | 

### DHCP Option Sets

**Sub-menu:** `/ip dhcp-server option sets`

This menu allows combining multiple options in option sets, which later can be used to override the default DHCP server option set.

### Example

**Classless Route**

A classless route adds a specified route in the clients routing table. In our example, it will add

- dst-address=160.0.0.0/24 gateway=10.1.101.1
- dst-address=0.0.0.0/0 gateway=10.1.101.1


According to RFC 3442: The first part is the netmask ("18" = netmask /24). Second part is significant part of destination network ("A00000" = 160.0.0). Third part is IP address of gateway ("0A016501" = 10.1.101.1). Then There are parts of the default route, destination netmask (0x00 = 0.0.0.0/0) followed by default route (0x0A016501 = 10.1.101.1)

Result:

A much more robust way would be to use built-in variables, the previous example can be rewritten as:

**Auto proxy config**

## Option matcher

The Option matcher allows to identify DHCP clients by any of DHCP options and assign IP address from specific IP pool.

It is possible to perform `exact` (provided *value* should match exactly ) or  `substring` matching (will look for *value* match anywhere in the option string — it can match values at the start, middle, or end).

Substring matching is useful in cases where value can change depending on the end device, for example, if the class-identifier sent by the device contains not only vendor information, but also exact MAC or other additional information.

