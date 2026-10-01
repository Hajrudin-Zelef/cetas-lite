---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-ac339467
title: "c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-ac339467"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-ac339467.md
source_anchor: ""
source_lines: [1, 56]
sha256: 31259f459d1ef21a31138c89bde049d668085d6baa38cd9ad0f3bd6ea51c6f5c
---

# c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-ac339467

Prerequisites for Configuring DHCP
The following prerequisites apply to DHCP Snooping and Option 82:
- 
                                 					
                                 You must globally enable DHCP snooping on the switch.
- 
                                 					
                                 Before globally enabling DHCP snooping on the switch, make sure that the devices acting as the DHCP server and the DHCP relay agent are configured and enabled.
- 
                                 					
                                 If you want the switch to respond to DHCP requests, it must be configured as a DHCP server.
- 
                                 					
                                 Before configuring the DHCP snooping information option on your switch, be sure to configure the device that is acting as the DHCP server. You must specify the IP addresses that the DHCP server can assign or exclude, or you must configure DHCP options for these devices.
- 
                                 					
                                 For DHCP snooping to function properly, all DHCP servers must be connected to the switch through trusted interfaces, as untrusted DHCP messages will be forwarded only to trusted interfaces. In a service-provider network, a trusted interface is connected to a port on a device in the same network.
- 
                                 					
                                 You must configure the switch to use the Cisco IOS DHCP server binding database to use it for DHCP snooping.
- 
                                 					
                                 To use the DHCP snooping option of accepting packets on untrusted inputs, the switch must be an aggregation switch that receives packets with option-82 information from an edge switch.
- 
                                 					
                                 The following prerequisites apply to DHCP snooping binding database configuration: 
  - 
                                       							
                                       You must configure a destination on the DHCP snooping binding database to use the switch for DHCP snooping.
  - 
                                       							
                                       Because both NVRAM and the flash memory have limited storage capacity, we recommend that you store the binding file on a TFTP server.
  - 
                                       							
                                       For network-based URLs (such as TFTP and FTP), you must create an empty file at the configured URL before the switch can write bindings to the binding file at that URL. See the documentation for your TFTP server to determine whether you must first create an empty file on the server; some TFTP servers cannot be configured this way.
  - 
                                       							
                                       To ensure that the lease time in the database is accurate, we recommend that you enable and configure Network Time Protocol (NTP).
  - 
                                       							
                                       If NTP is configured, the switch writes binding changes to the binding file only when the switch system clock is synchronized with NTP.
- 
                                       							
                                       
- 
                                 					
                                 Before configuring the DHCP relay agent on your switch, make sure to configure the device that is acting as the DHCP server. You must specify the IP addresses that the DHCP server can assign or exclude, configure DHCP options for devices, or set up the DHCP database agent.
- 
                                 					
                                 If you want the switch to relay DHCP packets, the IP address of the DHCP server must be configured on the switch virtual interface (SVI) of the DHCP client.
- 
                                 					
                                 If a switch port is connected to a DHCP server, configure a port as trusted by entering the ip dhcp snooping trust interface configuration command.
- 
                                 					
                                 If a switch port is connected to a DHCP client, configure a port as untrusted by entering the no ip dhcp snooping trust interface configuration command.
