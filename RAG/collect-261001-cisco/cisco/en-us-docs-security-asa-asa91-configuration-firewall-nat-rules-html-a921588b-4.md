---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b-4
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b.md
source_anchor: ""
source_lines: [134, 171]
sha256: d1875aa91bea6f75ff197263dcd123a73291e4e6b547180e3fb975fb5baaff6a
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b

|  |  | (continued) – Mapped—Specify a network object or group, or for static interface NAT with port translation only (routed mode), specify the interface keyword. If you specify ipv6 , then the IPv6 address of the interface is used. If you specify interface , be sure to also configure the service keyword. For this option, you must configure a specific interface for the real_ifc . See the “Static Interface NAT with Port Translation” section for more information. – Real—Specify a network object or group. For identity NAT, simply use the same object or group for both the real and mapped addresses.  | 
Examples
The following example configures interface PAT for inside network 192.168.1.0/24 when accessing outside Telnet server 209.165.201.23, and Dynamic PAT using a PAT pool when accessing any server on the 203.0.113.0/24 network.
ciscoasa(config)# object network INSIDE_NW
ciscoasa(config-network-object)# subnet 192.168.1.0 255.255.255.0
ciscoasa(config)# object network PAT_POOL
ciscoasa(config-network-object)# range 209.165.200.225 209.165.200.254
ciscoasa(config)# object network TELNET_SVR
ciscoasa(config-network-object)# host 209.165.201.23
ciscoasa(config)# object service TELNET
ciscoasa(config-service-object)# service tcp destination eq 23
ciscoasa(config)# object network SERVERS
ciscoasa(config-network-object)# subnet 203.0.113.0 255.255.255.0
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW interface destination static TELNET_SVR TELNET_SVR service TELNET TELNET
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW pat-pool PAT_POOL destination static SERVERS SERVERS
The following example configures interface PAT for inside network 192.168.1.0/24 when accessing outside IPv6 Telnet server 2001:DB8::23, and Dynamic PAT using a PAT pool when accessing any server on the 2001:DB8:AAAA::/96 network.
Configuring Static NAT or Static NAT-with-Port-Translation
This section describes how to configure a static NAT rule using twice NAT. For more information about static NAT, see the “Static NAT” section.
Detailed Steps
| Step 1 | Create network objects or groups for the: | See the “Adding Network Objects for Real and Mapped Addresses” section. If you want to configure source static interface NAT with port translation only, you can skip adding an object for the source mapped addresses, and instead specify the interface keyword in the nat command. If you want to configure destination static interface NAT with port translation only, you can skip adding an object for the destination mapped addresses, and instead specify the interface keyword in the nat command. | 
| Step 2 | (Optional) Create service objects for the: | See the “(Optional) Adding Service Objects for Real and Mapped Ports” section. | 
| Step 3 | nat [ ( real_ifc , mapped_ifc ) ] [ line \| { after-object [ line ]}] source static real_ob [ mapped_obj \| interface [ ipv6 ]] [ destination static { mapped_obj \| interface [ ipv6 ]} real_obj ] [ service real_src_mapped_dest_svc_obj mapped_src_real_dest_svc_obj ][ net-to-net ] [ dns ] [ unidirectional \| no-proxy-arp ] [ inactive ] [ description desc ] ciscoasa(config)# nat (inside,dmz) source static MyInsNet MyInsNet_mapped destination static Server1 Server1 service REAL_SRC_SVC MAPPED_SRC_SVC | Configures static NAT . See the following guidelines:  – Real—Specify a network object or group. – Mapped—Specify a different network object or group. For static interface NAT with port translation only, you can specify the interface keyword (routed mode only). If you specify ipv6 , then the IPv6 address of the interface is used. If you specify interface , be sure to also configure the service keyword (in this case, the service objects should include only the source port). For this option, you must configure a specific interface for the mapped_ifc . See the “Static Interface NAT with Port Translation” section for more information. – Mapped—Specify a network object or group, or for static interface NAT with port translation only, specify the interface keyword. If you specify ipv6 , then the IPv6 address of the interface is used. If you specify interface , be sure to also configure the service keyword (in this case, the service objects should include only the destination port). For this option, you must configure a specific interface for the real_ifc . – Real—Specify a network object or group. For identity NAT, simply use the same object or group for both the real and mapped addresses. | 
|  |  | (Continued)  | 
Examples
The following example shows the use of static interface NAT with port translation. Hosts on the outside access an FTP server on the inside by connecting to the outside interface IP address with destination port 65000 through 65004. The traffic is untranslated to the internal FTP server at 192.168.10.100:6500 through :65004. Note that you specify the source port range in the service object (and not the destination port) because you want to translate the source address and port as identified in the command; the destination port is “any.” Because static NAT is bidirectional, “source” and “destination” refers primarily to the command keywords; the actual source and destination address and port in a packet depends on which host sent the packet. In this example, connections are originated from outside to inside, so the “source” address and port of the FTP server is actually the destination address and port in the originating packet.
ciscoasa(config)# object service FTP_PASV_PORT_RANGE
ciscoasa(config-service-object)# service tcp source range 65000 65004
ciscoasa(config)# object network HOST_FTP_SERVER
ciscoasa(config-network-object)# host 192.168.10.100
ciscoasa(config)# nat (inside,outside) source static HOST_FTP_SERVER interface service FTP_PASV_PORT_RANGE FTP_PASV_PORT_RANGE
The following example shows a static translation of one IPv6 network to another IPv6 when accessing an IPv6 network, and the dynamic PAT translation to an IPv4 PAT pool when accessing the IPv4 network:
ciscoasa(config)# nat (inside,outside) source static INSIDE_NW MAPPED_IPv6_NW destination static OUTSIDE_IPv6_NW OUTSIDE_IPv6_NW
ciscoasa(config)# nat (inside,outside) source dynamic INSIDE_NW pat-pool MAPPED_IPv4_POOL destination static OUTSIDE_IPv4_NW OUTSIDE_IPv4_NW
Configuring Identity NAT
This section describes how to configure an identity NAT rule using twice NAT. For more information about identity NAT, see the “Identity NAT” section.
Detailed Steps
| Step 1 | Create network objects or groups for the: | See the “Adding Network Objects for Real and Mapped Addresses” section. If you want to perform identity NAT for all addresses, you can skip creating an object for the the source real addresses and instead use the keywords any any in the nat command. If you want to configure destination static interface NAT with port translation only, you can skip adding an object for the destination mapped addresses, and instead specify the interface keyword in the nat command. | 
| Step 2 | (Optional) Create service objects for the: | See the “(Optional) Adding Service Objects for Real and Mapped Ports” section. | 
