---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-70-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-70.md
source_anchor: ""
source_lines: [85, 146]
sha256: 1813269513f99c3b136ab91fe6ce7f9d7a0608f0d3e0c0cd429ef8f54c19e510
---

# Overview

Available read only properties:

| Property | Description | 
|---|---|
| **ac-mac** (*MAC address* ) | MAC address of the access concentrator (AC) the client is connected to | 
| **ac-name** (*string* ) | name of the Access Concentrator | 
| **active-links** (*integer* ) | Number of bonded MLPPP connections, ('1' if not using MLPPP) | 
| **encoding** (*string* ) | encryption and encoding (if asymmetric, separated with '/') being used in this connection | 
| **local-address** (*IP Address* ) | IP Address allocated to client | 
| **remote-address** (*IP Address* ) | Remote IP Address allocated to server (ie gateway address) | 
| **mru** (*integer* ) | effective MRU of the link | 
| **mtu** (*integer* ) | effective MTU of the link | 
| **service-name** (*string* ) | used service name | 
| **status** (*string* ) | current link status. Available values are:  | 
| **uptime** (*time* ) | connection time displayed in days, hours, minutes and seconds | 

## Scanner

PPPoE Scanner allows scanning all active PPPoE servers in the layer2 broadcast domain. Command to run scanner is as follows:

Available read only properties:

| Property | Description | 
|---|---|
| **service** (*string* ) | Service name configured on server | 
| **mac-address** (*MAC* ) | Mac address of detected server | 
| **ac-name** (*string* ) | name of the Access Concentrator | 

For Windows, some connection instructions may use the form where the "phone number", such as "MikroTik_AC\mt1", is specified to indicate that "MikroTik_AC" is the access concentrator name and "mt1" is the service name.

Specifying MRRU means enabling MP (Multilink PPP) over a single link. This protocol is used to split big packets into smaller ones. Under Windows, it can be enabled in the Networking tab, Settings button, "Negotiate multi-link for single link connections". MRRU is hardcoded to 1614 on Windows. This setting is useful to overcome PathMTU discovery failures. The MP setting should be enabled on both peers.

# PPPoE Server

There are two types of interface (tunnel) items in PPPoE server configuration - static users and dynamic connections. An interface is created for each tunnel established to the given server. Static interfaces are added administratively if there is a need to reference the particular interface name (in firewall rules or elsewhere) created for the particular user. Dynamic interfaces are added to this list automatically whenever a user is connected and its username does not match any existing static entry (or in case the entry is active already, as there can not be two separate tunnel interfaces referenced by the same name - set *one-session-per-host* value if this is a problem). Dynamic interfaces appear when a user connects and disappear once the user disconnects, so it is impossible to reference the tunnel created for that use in router configuration (for example, in firewall), so if you need a persistent rule for that user, create a static entry for him/her. Otherwise, it is safe to use a dynamic configuration. 

In both cases PPP users must be configured properly - static entries do not replace PPP configuration.

## Access concentrator

### Properties

| Property | Description | 
|---|---|
| **accept-untagged** (*yes \| no* ; Default:**yes** ) | This setting controls whether the PPPoE server will accept untagged (non-VLAN) PPPoE packets on its interface, when `pppoe-over-vlan-range` is specified. By default, untagged PPPoE packets are accepted. If you are using the `pppoe-over-vlan-range` property (which enabled PPPoE over 802.1Q VLANs), this option lets you decide whether to still allow untagged clients on the same interface. If you are not using the`pppoe-over-vlan-range` , this setting do not have any effect. | 
| **authentication** ( *mschap2 \| mschap1 \| chap \| pap* ; Default:**"mschap2, mschap1, chap, pap"** ) | Authentication algorithm. | 
| **default-profile** (*string* ; Default:**"default"** ) |  | 
| **interface** (*string* ; Default:**""** ) | Interface that the clients are connected to. | 
| **keepalive-timeout** (*time* ; Default:**"10", or disabled** ) | Defines the time period (in seconds) after which the router is starting to send keepalive packets every second. If there is no traffic and no keepalive responses arrive for that period of time (i.e. 2 * keepalive-timeout), the non responding client is proclaimed disconnected. After a successful LCP handshake, the client sends LCP echo packets to verify MTU forwarding; if no reply is received, it falls back to a backup MTU of 1480. A new option, `keepalive-timeout=disabled` , disables sending echo packets, effectively turning off the MTU test. | 
| **max-mru** (*integer* ; Default:**"1480"** ) | Maximum Receive Unit. The optimal value is the MTU of the interface the tunnel is working over reduced by 20 (so, for 1500-byte Ethernet link, set the MTU to 1480 to avoid fragmentation of packets) | 
| **max-mtu** (*integer* ; Default:**"1480"** ) | Maximum Transmission Unit. The optimal value is the MTU of the interface the tunnel is working over reduced by 20 (so, for 1500-byte Ethernet link, set the MTU to 1480 to avoid fragmentation of packets) | 
| **max-sessions** (*integer* ; Default:**"0"** ) | Maximum number of clients that the AC can serve. '0' = no limitations. | 
| **mrru** (*integer: 512..65535 \| disabled* ; Default:**"disabled"** ) | Maximum packet size that can be received on the link. If a packet is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet packets to be sent over the tunnel. | 
| **one-session-per-host** (*yes \| no* ; Default:**"no"** ) | Allow only one session per host (determined by MAC address). If a host tries to establish a new session, the old one will be closed. | 
| **pppoe-over-vlan-range** (*integer 1..4094* ; Default: "") | This setting allows a PPPoE server to operate over 802.1Q VLANs. By default, a PPPoE server only accepts untagged packets on its interface. However, in scenarios where clients are on separate VLANs, instead of creating multiple 802.1Q VLAN interfaces and bridging them together or configuring individual PPPoE servers for each VLAN, you can specify the necessary VLANs directly in the PPPoE server settings. When you specify the VLAN IDs, the PPPoE server will accept 802.1Q tagged packets from clients, and it will reply using the same VLAN. You then have an option to accept or drop untagged PPoE clients on the same interface using the `accept-untagged` property. You can configure the PPPoE server with `pppoe-over-vlan-range` setting even on VLAN interface enabling the QinQ setups as well. But keep in mind that the inner VLAN tag should be 802.1Q. The setting supports a range of VLAN IDs, as well as individual VLANs specified using comma-separated values. For example: pppoe-over-vlan-range=100-115,120,122,128-130. Avoid configuring a server with `pppoe-over-vlan-range` on an interface while also creating a VLAN interface using a VLAN ID that falls within that range. For example: If you need this type of setup, remove the overlapping VLAN ID from `pppoe-over-vlan-rang` and create a separate PPPoE server instance directly on the VLAN interface, like this: | 
| **service-name** (*string* ; Default:**""** ) | The PPPoE service name. Server will accept clients which sends PADI message with service-names that matches this setting or if service-name field in PADI message is not set. | 

The PPPoE server (access concentrator) supports multiple servers for each interface - with differing service names. The access concentrator name and PPPoE service name are used by clients to identify the access concentrator to register with. The access concentrator name is the same as the identity of the router displayed before the command prompt. The identity may be set within the */system identity* submenu.

Do not assign an IP address to the interface you will be receiving the PPPoE requests on.


