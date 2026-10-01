---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98.md
source_anchor: ""
source_lines: [52, 87]
sha256: e53d459372346b12d8fde39d2e8e069faab03326eef1043903c015893529834b
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98

                                          Symmetric active mode
The client and the symmetric active modes should be used when NTP is required to provide a high level of time accuracy and reliability.
When a networking device is operating in the client mode, it polls its assigned time-serving hosts for the current time. The networking device will then pick a host from among all the polled time servers to synchronize with. Because the relationship that is established in this case is a client-host relationship, the host will not capture or use any time information sent by the local client device. This mode is most suited for file-server and workstation clients that are not required to provide any form of time synchronization to other local clients. Use the ntp server command to individually specify the time server that you want your networking device to consider synchronizing with and to set your networking device to operate in the client mode.
When a networking device is operating in the symmetric active mode, it polls its assigned time-serving hosts for the current time and it responds to polls by its hosts. Because this is a peer-to-peer relationship, the host will also retain time-related information of the local networking device that it is communicating with. This mode should be used when a number of mutually redundant servers are interconnected via diverse network paths. Most stratum 1 and stratum 2 servers on the Internet adopt this form of network setup. Use the ntp peer command to individually specify the time serving hosts that you want your networking device to consider synchronizing with and to set your networking device to operate in the symmetric active mode.
The specific mode that you should set for each of your networking devices depends primarily on the role that you want them to assume as a timekeeping device (server or client) and the device’s proximity to a stratum 1 timekeeping server.
A networking device engages in polling when it is operating as a client or a host in the client mode or when it is acting as a peer in the symmetric active mode. Although polling does not usually place a burden on memory and CPU resources such as bandwidth, an exceedingly large number of ongoing and simultaneous polls on a system can seriously impact the performance of a system or slow the performance of a given network. To avoid having an excessive number of ongoing polls on a network, you should limit the number of direct, peer-to-peer or client-to-server associations. Instead, you should consider using NTP broadcasts to propagate time information within a localized network.
Broadcast-Based NTP Associations
Broadcast-based NTP associations should be used when time accuracy and reliability requirements are modest and if your network is localized and has more than 20 clients. Broadcast-based NTP associations are also recommended for use on networks that have limited bandwidth, system memory, or CPU resources.
A networking device operating in the broadcast client mode does not engage in any polling. Instead, it listens for NTP broadcast packets that are transmitted by broadcast time servers. Consequently, time accuracy can be marginally reduced because time information flows only one way.
Use the ntp broadcast client command to set your networking device to listen for NTP broadcast packets propagated through a network. For broadcast client mode to work, the broadcast server and its clients must be located on the same subnet. You must enable the time server that transmits NTP broadcast packets on the interface of the given device by using the ntp broadcast command.
Authoritative NTP Server
An authoritative NTP server is a time server that can distribute time in the network. Other devices can configure it as a time server. You can configure a Cisco Catalyst 9000 Series Switch to act as an authoritative NTP server, enabling it to distribute time even when it is not synchronized to an outside time source. Use the ntp master command, in global configuration mode, to configure the device to be an authoritative NTP server.
| Caution | Use the ntp master command with caution. Usage of this command can override valid time sources, especially if a low stratum number is configured. Configuring multiple devices in the same network with the ntp master command can cause instability in timekeeping if the devices do not agree on the time. | 
NTP Security
The time kept on a device is a critical resource; you should use the security features of NTP to avoid the accidental or malicious setting of an incorrect time. Two mechanisms are available: an access list-based restriction scheme and an encrypted authentication mechanism.
| Note | We do not recommend configuring Message Direct 5 (MD5) authentication. You can use other supported authentication methods for stronger encryption. | 
NTP Access Group
The access list-based restriction scheme allows you to grant or deny certain access privileges to an entire network, a subnet within a network, or a host within a subnet. To define an NTP access group, use the ntp access-group command in global configuration mode.
The access group options are scanned in the following order, from least restrictive to the most restrictive:
- 
                                          
                                          ipv4 —Configures IPv4 access lists.
- 
                                          
                                          ipv6 —Configures IPv6 access lists.
- 
                                          
                                          peer —Allows time requests and NTP control queries, and allows the system to synchronize itself to a system whose address passes the access list criteria.
- 
                                          
                                          serve —Allows time requests and NTP control queries, but does not allow the system to synchronize itself to a system whose address passes the access list criteria.
- 
                                          
                                          serve-only —Allows only time requests from a system whose address passes the access list criteria.
- 
                                          
