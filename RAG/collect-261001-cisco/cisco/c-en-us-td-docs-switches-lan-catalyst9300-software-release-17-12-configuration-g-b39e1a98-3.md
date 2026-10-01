---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98.md
source_anchor: ""
source_lines: [88, 127]
sha256: 5c76109c6ed635ac7e2d83a08e38772b8561a4da06aa88c9870232c0ab214e98
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98

                                          query-only —Allows only NTP control queries from a system whose address passes the access list criteria.
If the source IP address matches the access lists for more than one access type, the first type is granted access. If no access groups are specified, all access types are granted access to all systems. If any access groups are specified, only the specified access types will be granted access.
For details on NTP control queries, see RFC 1305.
The encrypted NTP authentication scheme should be used when a reliable form of access control is required. Unlike the access list-based restriction scheme that is based on IP addresses, the encrypted authentication scheme uses authentication keys and an authentication process to determine if NTP synchronization packets sent by designated peers or servers on a local network are deemed as trusted before the time information that they carry along with them is accepted.
The authentication process begins from the moment an NTP packet is created. Cryptographic checksum keys are generated using the message digest algorithm 5 (MD5) and are embedded into the NTP synchronization packet that is sent to a receiving client. Once a packet is received by a client, its cryptographic checksum key is decrypted and checked against a list of trusted keys. If the packet contains a matching authentication key, the time-stamp information that is contained within the packet is accepted by the receiving client. NTP synchronization packets that do not contain a matching authenticator key are ignored.
| Note | In large networks, where many trusted keys must be configured, the Range of Trusted Key Configuration feature enables configuring multiple keys simultaneously. | 
It is important to note that the encryption and decryption processes used in NTP authentication can be very CPU-intensive and can seriously degrade the accuracy of the time that is propagated within a network. If your network setup permits a more comprehensive model of access control, you should consider the use of the access list-based form of control.
After NTP authentication is properly configured, your networking device will synchronize with and provide synchronization only to trusted time sources.
NTP Services on a Specific Interface
Network Time Protocol (NTP) services are disabled on all interfaces by default. NTP is enabled globally when any NTP commands are entered. You can selectively prevent NTP packets from being received through a specific interface by using the ntp disable command in interface configuration mode.
Source IP Address for NTP Packets
When the system sends an NTP packet, the source IP address is normally set to the address of the interface through which the NTP packet is sent. Use the ntp source interface command in global configuration mode to configure a specific interface from which the IP source address will be taken.
This interface will be used for the source address for all packets sent to all destinations. If a source address is to be used for a specific association, use the source keyword in the ntp peer or ntp server command.
NTP Implementation
Implementation of NTP does not support stratum 1 service; it is not possible to connect to a radio or atomic clock. We recommend that the time service for your network be derived from the public NTP servers available on the IP Internet.
If the network is isolated from the Internet, NTP allows a device to act as if it is synchronized through NTP, when in fact it has learned the time by using other means. Other devices then synchronize to that device through NTP.
When multiple sources of time are available, NTP is always considered to be more authoritative. NTP time overrides the time set by any other method.
Several manufacturers include NTP software for their host systems, and a publicly available version for systems running UNIX and its various derivatives is also available. This software allows host systems to be time-synchronized as well.
DNS
The DNS protocol controls the Domain Name System (DNS), a distributed database with which you can map hostnames to IP addresses. When you configure DNS on your device, you can substitute the hostname for the IP address with all IP commands, such as ping, telnet, connect , and related Telnet support operations.
IP defines a hierarchical naming scheme that allows a device to be identified by its location or domain. Domain names are pieced together with periods (.) as the delimiting characters. For example, Cisco Systems is a commercial organization that IP identifies by a com domain name, so its domain name is cisco.com. A specific device in this domain, for example, the File Transfer Protocol (FTP) system is identified as ftp.cisco.com.
To keep track of domain names, IP has defined the concept of a domain name server, which holds a cache (or database) of names mapped to IP addresses. To map domain names to IP addresses, you must first identify the hostnames, specify the name server that is present on your network, and enable the DNS.
Default DNS Settings
| Table 1. Default DNS Settings |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| DNS enable state | Enabled. | 
| DNS default domain name | None configured. | 
| DNS servers | No name server addresses are configured. | 
Login Banners
| Note | For complete syntax and usage information for the commands used in this section, see the Cisco IOS Configuration Fundamentals Command Reference, Release 12.4. | 
Default Banner Configuration
MAC Address Table
The MAC address table contains address information that the device uses to forward traffic between ports. All MAC addresses in the address table are associated with one or more ports. The address table includes these types of addresses:
- 
                                 				
                                 Dynamic address—A source MAC address that the device learns and then ages when it is not in use.
- 
                                 				
