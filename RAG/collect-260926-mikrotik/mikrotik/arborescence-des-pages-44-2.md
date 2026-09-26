---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-44-2
title: "Types of NAT:"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-44.md
source_anchor: ""
source_lines: [79, 180]
sha256: 3fa97957f0b6d5508082bc346f29da429c608bf8e5d5dea7494ff9a2ab309153
---

# Types of NAT:

Packets with Shared Address Space source or destination addresses MUST NOT be forwarded across Service Provider boundaries. Service Providers MUST filter such packets on ingress links. In RouterOS this can be easily done with firewall filters on edge routers:

Service providers may be required to log of MAPed addresses, in a large CGN deployed network which may be a problem. Fortunately, RFC 7422 suggests a way to manage CGN translations in such a way as to significantly reduce the amount of logging required while providing traceability for abuse response.

RFC states that instead of logging each connection, CGNs could deterministically map customer private addresses (received on the customer-facing interface of the CGN, a.k.a., internal side) to public addresses extended with port ranges.

That means that separate NAT rules have to be added to achieve individual mappings such as the ones seen in the below example:

| **Inside IP** | **Outside IP/Port range** | 
| 100.64.0.1 | 2.2.2.2:5000-5199 | 
| 100.64.0.2 | 2.2.2.2:5200-5399 | 
| 100.64.0.3 | 2.2.2.2:5400-5599 | 
| 100.64.0.4 | 2.2.2.2:5600-5799 | 
| 100.64.0.5 | 2.2.2.2:5800-5999 | 

Instead of writing the rules by hand, it is suggested to use a script instead. The following example could be adapted to any requirements of your setup.

The six local values can be adjusted and the script can be either simply pasted in the terminal or it can be stored in the system script section, in case the configuration needs to be regenerated later.

After execution, you should get a set of rules:

### 

Hairpin NAT

Hairpin network address translation (*NAT Loopback*) is where the device on the LAN can access another machine on the LAN via the public IP address of the gateway router. 

In the above example, the gateway router has the following `dst-nat` configuration rule:

When a user from the PC at home establishes a connection to the web server, the router performs DST NAT as configured:

1. the client sends a packet with a source IP address of 192.168.88.1 to a destination IP address of 172.16.16.1 on port 443 to request some web resources;
2. the router destination NAT`s the packet to 10.0.0.3 and replaces the destination IP address in the packet accordingly. The source IP address stays the same: 192.168.88.1;
3. the server replies to the client's request and the reply packet has a source IP address of 10.0.0.3 and a destination IP address of 192.168.88.1.
4. the router determines that the packet is part of a previous connection and undoes the destination NAT, and puts the original destination IP address into the source IP address field. The destination IP address is 192.168.88.1, and the source IP address is 172.16.16.1;
5. The client receives the reply packet it expects, and the connection is established;

But, there will be a **problem**, when a client on the same network as the web server requests a connection to the web server's **public** IP address: 

1. the client sends a packet with a source IP address of 10.0.0.2 to a destination IP address of 172.16.16.1 on port 443 to request some web resources;
2. the router destination NATs the packet to 10.0.0.3 and replaces the destination IP address in the packet accordingly. The source IP address stays the same: 10.0.0.2;
3. the server replies to the client's request. However, the source IP address of the request is on the same subnet as the web server. The web server does not send the reply back to the router but sends it back directly to 10.0.0.2 with a source IP address in the reply of 10.0.0.3;
4. The client receives the reply packet, but it discards it because it expects a packet back from 172.16.16.1, and not from 10.0.0.3;

To resolve this issue, we will configure a new *src-nat* rule (the hairpin NAT rule) as follows:

After configuring the rule above:

1. the client sends a packet with a source IP address of 10.0.0.2 to a destination IP address of 172.16.16.1 on port 443 to request some web resources;
2. the router destination NATs the packet to 10.0.0.3 and replaces the destination IP address in the packet accordingly. It also source NATs the packet and replaces the source IP address in the packet with the IP address on its LAN interface. The destination IP address is 10.0.0.3, and the source IP address is 10.0.0.1;
3. the web server replies to the request and sends the reply with a source IP address of 10.0.0.3 back to the router's LAN interface IP address of 10.0.0.1;
4. the router determines that the packet is part of a previous connection and undoes both the source and destination NAT, and puts the original destination IP address of 10.0.0.3 into the source IP address field, and the original source IP address of 172.16.16.1 into the destination IP address field

## Endpoint-Independent NAT

Endpoint-independent NAT creates mapping in the source NAT and uses the same mapping for all subsequent packets with the same source IP and port. This mapping is created with the following rule:

This mapping allows running source-independent filtering, which allows forwarding packets from any source from WAN to mapped internal IP and port. The following rule enables filtering:

Endpoint-independent NAT works only with UDP protocol.

Additionally, endpoint-independent-nat can take a few other parameters:

**randomize-port**

More info https://www.ietf.org/rfc/rfc5128.txt section 2.2.3 and 2.2.5

# NAT Helpers

Hosts behind a NAT-enabled router do not have true end-to-end connectivity. Therefore some Internet protocols might not work in scenarios with NAT. To overcome these limitations RouterOS includes a number of NAT helpers, that enable NAT traversal for various protocols.

Nat helpers can be managed from `/ip firewall service-ports` menu.

List of available nat helpers:

| Helper | Description | 
|---|---|
| **FTP** | FTP service helper | 
| **H323** | H323 service helper | 
| **IRC** | IRC service helper | 
| **PPTP** | PPTP (GRE) tunneling helper | 
| **UDPLITE** | UDP-Lite service helper | 
| **DCCP** | DCCP service helper | 
| **SCTP** | SCTP service helper | 
| **SIP** | SIP helper. Additional options:  | 
| **TFTP** | TFTP service helper | 
| **RSTP** | RTSP service helper | 

If connection tracking is not enabled then firewall service ports will be shown as inactive

**udplite**, **dccp**, and **sctp** are built-in services of the connection tracking. Since these are not separately loaded modules, they cannot be disabled separately, they got disabled together with the connection tracking.

# NAT Actions

Table lists NAT actions and their associated properties. Other actions are listed here.

| Property | Description | 
|---|---|
| **action** (*action name* ; Default:**accept** ) |  | 
| **same-not-by-dst** (*yes \| no* ; Default: ) | Specifies whether to take into account or not the destination IP address when selecting a new source IP address. Applicable if `action=same` | 
| **to-addresses** (*IP address[-IP address]* ; Default:**0.0.0.0** ) | Replace the original address with the specified one. Applicable if action is `dst-nat` ,`netmap` ,`same` ,`src-nat` | 
| **to-ports** (*integer[-integer]: 0..65535* ; Default: ) | Replace the original port with the specified one. Applicable if action is `dst-nat` ,`redirect` ,`masquerade` ,`netmap` ,`same` ,`src-nat` |
