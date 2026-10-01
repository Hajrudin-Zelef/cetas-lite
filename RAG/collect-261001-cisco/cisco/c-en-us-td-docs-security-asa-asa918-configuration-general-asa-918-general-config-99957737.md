---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa918-configuration-general-asa-918-general-config-99957737
title: "c-en-us-td-docs-security-asa-asa918-configuration-general-asa-918-general-config-99957737"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa918-configuration-general-asa-918-general-config-99957737.md
source_anchor: ""
source_lines: [1, 64]
sha256: cea1962bb7383fb163a358e8db72f565c4eb0ced9349803bd408333d237a4400
---

# c-en-us-td-docs-security-asa-asa918-configuration-general-asa-918-general-config-99957737

About DHCP and DDNS services
This topic describes the DHCP server, DHCP relay agent, and DDNS update.
DHCPv4 servers
A DHCPv4 server provides network configuration parameters, such as IP addresses, to DHCP clients.
- 
                                    
                                    delivers configuration parameters directly to DHCP clients attached to ASA interfaces, and
- 
                                    
                                    uses UDP port communication for client-server messaging.
DHCPv4 server communication details
The DHCP communication uses specific network protocols and port assignments:
- 
                                    
                                    IPv4 DHCP client: Uses broadcast rather than multicast address to reach the server and listens for messages on UDP port 68
- 
                                    
                                    DHCP server: Listens for messages on UDP port 67
DHCP options
DHCP provides a framework for passing configuration information to hosts on a TCP/IP network. The configuration parameters are carried in tagged items that are stored in the Options field of the DHCP message and the data are also called options. Vendor information is also stored in Options, and all of the vendor information extensions can be used as DHCP options.
Common DHCP options
Cisco IP Phones download their configuration from a TFTP server. When a Cisco IP Phone starts, if it does not have both the IP address and TFTP server IP address preconfigured, it sends a request with option 150 or 66 to the DHCP server to obtain this information.
- 
                                       
                                       DHCP option 150 provides the IP addresses of a list of TFTP servers.
- 
                                       
                                       DHCP option 66 gives the IP address or the hostname of a single TFTP server.
- 
                                       
                                       DHCP option 3 sets the default route.
A single request might include both options 150 and 66. In this case, the DHCP server provides values for both options in the response if they are already configured on the ASA.
You can use advanced DHCP options to provide DNS, WINS, and domain name parameters to DHCP clients. DHCP option 15 is used for the DNS domain suffix. You can also use the DHCP automatic configuration setting to obtain these values or define them manually. When you use more than one method to define this information, it is passed to DHCP clients in the following sequence:
- 
                                       
                                       Manually configured settings.
- 
                                       
                                       Advanced DHCP options settings.
- 
                                       
                                       DHCP automatic configuration settings.
Domain name configuration precedence
You can manually define the domain name that you want the DHCP clients to receive and then enable DHCP automatic configuration. Although DHCP automatic configuration discovers the domain together with the DNS and WINS servers, the manually defined domain name is passed to DHCP clients with the discovered DNS and WINS server names, because the domain name discovered by the DHCP automatic configuration process is superseded by the manually defined domain name.
DHCPv6 stateless server
A DHCPv6 stateless server is a DHCP IPv6 configuration that
- 
                                    
                                    provides information such as DNS server or domain name when clients send Information Request (IR) packets,
- 
                                    
                                    accepts IR packets but does not assign addresses to clients, and
- 
                                    
                                    works in conjunction with StateLess Address Auto Configuration (SLAAC) and Prefix Delegation features.
DHCPv6 stateless server operation
For clients that use StateLess Address Auto Configuration (SLAAC) in conjunction with the Prefix Delegation feature (Enable the IPv6 Prefix Delegation Client), you can configure the ASA to provide information such as the DNS server or domain name when they send Information Request (IR) packets to the ASA. The ASA only accepts IR packets and does not assign addresses to the clients. You will configure the client to generate its own IPv6 address by enabling IPv6 autoconfiguration on the client. Enabling stateless autoconfiguration on a client configures IPv6 addresses based on prefixes received in Router Advertisement messages; in other words, based on the prefix that the ASA received using Prefix Delegation.
DHCP relay agents
A DHCP relay agent forwards DHCP requests received on an interface to one or more DHCP servers.
DHCP relay agent operation
DHCP clients use UDP broadcasts to send their initial DHCPDISCOVER messages because they do not have information about the network to which they are attached. If the client is on a network segment that does not include a server, UDP broadcasts normally are not forwarded by the ASA because it does not forward broadcast traffic. The DHCP relay agent lets you configure the interface of the ASA that is receiving the broadcasts to forward DHCP requests to a DHCP server on another device.
DHCP Relay Server Support on VTI
You can configure DHCP relay agent on an ASA interface to receive and forward DHCP messages between a DHCP client and a DHCP server. However, a DHCP relay server to forward messages through a logical interface was not supported.
The same procedure is followed for a DHCPREQUEST and DHCPACK/NACK requirements.
