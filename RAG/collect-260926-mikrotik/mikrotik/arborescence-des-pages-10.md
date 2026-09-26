---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-10
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-10.md
source_anchor: ""
source_lines: [1, 76]
sha256: d5d804fda7c1e68df504aa2c0e7f6eff0220ffe297a22e939964c9eeb82f9a28
---

# Introduction

A packet sniffer is a tool that can capture and analyze packets that are going to, leaving, or going through the router. Packet sniffing is very useful when you diagnose networks or protect against security attacks over networks.

Unicast traffic between Wireless clients with client-to-client forwarding enabled will not be visible to the sniffer tool.

Packets that are processed with hardware offloading enabled bridge will not be visible (flooded packets like unknown unicast, broadcast, and multicast traffic might be visible).

Sniffer catches packets before they go into the firewall (direction=rx) or after they leave it (direction=tx).

# Packet Sniffer configuration

RouterOS embedded sniffer allows you to capture packets based on various protocols.

In the following example, we will configure the sniffer to match packets going through the ether1 interface:

You can download captured packets from a file section. Then you can use a packet analyzer such as Wireshark to analyze a file:

If you are using packet streaming to PC and are using Wireshark, to ensure you are only viewing the streamed data, you will need to apply a filter that matches the port the sniffer is using, by default 37008 is used. In addition, we recommend using `filter-stream=yes`.

Please note that sniffed packets will be available for 10 minutes, if you need them permanently, set a "file-name" to save them directly or issue a "save" command as described previously.

**Sub-menu:** `/tool sniffer`

| Property | Description | 
|---|---|
| **file-limit** (*integer 10..4294967295[KiB]* ; Default:**1000KiB** ) | File size limit. Sniffer will stop when a limit is reached. | 
| **file-name** (*string* ; Default: ) | Name of the file where sniffed packets will be saved. | 
| **filter-cpu** (*integer* ; Default: ) | CPU core used as a filter. | 
| **filter-ip-address** (*ip/mask[,ip/mask] (max 16 items)* ; Default: ) | Up to 16 IP addresses used as a filter. | 
| **filter-dst-ip-address** (*ip/mask[,ip/mask] (max 16 items)* ; Default: ) | Up to 16 IP destination addresses used as a filter. | 
| **filter-src-ip-address** (*ip/mask[,ip/mask] (max 16 items)* ; Default: ) | Up to 16 IP source addresses used as a filter. | 
| **filter-ipv6-address** (*ipv6/mask[,ipv6/mask] (max 16 items)* ; Default: ) | Up to 16 IPv6 addresses used as a filter. | 
| **filter-dst-ipv6-address** (*ipv6/mask[,ipv6/mask] (max 16 items)* ; Default: ) | Up to 16 IPv6 destination addresses used as a filter. | 
| **filter-src-ipv6-address** (*ipv6/mask[,ipv6/mask] (max 16 items)* ; Default: ) | Up to 16 IPv6 source addresses used as a filter. | 
| **filter-mac-address** (*mac/mask[,mac/mask] (max 16 items)* ; Default: ) | Up to 16 MAC addresses and MAC address masks used as a filter. | 
| **filter-dst-mac-address** (*mac/mask[,mac/mask] (max 16 items)* ; Default: ) | Up to 16 MAC destination addresses and MAC address masks used as a filter. | 
| **filter-src-mac-address** (*mac/mask[,mac/mask] (max 16 items)* ; Default: ) | Up to 16 MAC source addresses and MAC address masks used as a filter. | 
| **filter-port** (*[!]port[,port] (max 16 items)* ; Default: ) | Up to 16 comma-separated ports used as a filter. A list of predefined port names is also available, like ssh and telnet. | 
| **filter-dst-port** (*[!]port[,port] (max 16 items)* ; Default: ) | Up to 16 comma-separated destination ports used as a filter. A list of predefined port names is also available, like ssh and telnet. | 
| **filter-src-port** (*[!]port[,port] (max 16 items)* ; Default: ) | Up to 16 comma-separated source ports used as a filter. A list of predefined port names is also available, like ssh and telnet. | 
| **filter-ip-protocol** (*[!]protocol[,protocol] (max 16 items)* ; Default: ) | Up to 16 comma-separated IP/IPv6 protocols used as a filter. IP protocols (instead of protocol names, protocol numbers can be used):  | 
| **filter-mac-protocol**  (*[!]protocol[,protocol] (max 16 items)* ; Default: ) | Up to 16 comma separated entries used as a filter. Mac protocols (instead of protocol names, protocol number can be used):  | 
| **filter-stream** (*yes \| no* ; Default:**yes** ) | Sniffed packets that are devised for the sniffer server are ignored. | 
| **filter-size** (*integer[-integer]:0..65535* ; Default: ) | Filters packets of specified size or size range in bytes. | 
| **filter-direction**  (*any \| rx \| tx* ; Default: ) | Specifies which direction filtering will be applied. | 
| **filter-interface** (*all \| name* ; Default:**all** ) | Interface name on which sniffer will be running. **all** indicates that the sniffer will sniff packets on all interfaces. | 
| **filter-operator-between-entries** (*and \| or* ; Default:**or** ) | Changes the logic for filters with multiple entries. | 
| **filter-vlan** (*integer[,integer]:0..4095* ; Default: ) | Up to 16 VLAN IDs used as a filter. | 
| **memory-limit**  (*integer 10..4294967295[KiB]* ; Default:**100KiB** ) | Memory amount used to store sniffed data. | 
| **memory-scroll** (*yes \| no* ; Default:**yes** ) | Whether to rewrite older sniffed data when the memory limit is reached. | 
| **only-headers**  (*yes \| no* ; Default:**no** ) | Save in the memory only the packet's headers, not the whole packet. | 
| **show-frame (*yes \| no*; Default: no)** | Whether to see the content of the frame when running quick sniffer in command line. | 
| **streaming-enabled** (*yes \| no* ; Default:**no** ) | Defines whether to send sniffed packets to the streaming server. | 
| **streaming-server** (*IP* ; Default:**0.0.0.0** ) | Tazmen Sniffer Protocol (TZSP) stream receiver. | 
| **streaming-port** (*port* ; Default:**37008** ) | Port to stream the TZSP packets to. | 

The `file-size` limit should not be configured more than available free memory!

## Packet Sniffer Quick Mode

The quick mode will display results as they are filtered out with a limited-size buffer for packets. There are several attributes that can be set up for filtering. If no attributes are set current configuration will be used.

Traffic-Generator packets will not be visible using the packet sniffer on the same interface unless the *fast-path* parameter is set.

## Packet Sniffer Protocols

In this submenu, you can see all sniffed protocols and their share of the whole sniffed amount.

## Packet Sniffer Host

The submenu shows the list of hosts that were participating in the data exchange you've sniffed.

## Packet Sniffer Connections

Here you can get a list of the connections that have been watched during the sniffing time.
