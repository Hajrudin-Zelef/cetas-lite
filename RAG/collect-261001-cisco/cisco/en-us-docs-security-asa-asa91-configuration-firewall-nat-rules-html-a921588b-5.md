---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b-5
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b.md
source_anchor: ""
source_lines: [172, 220]
sha256: d567b8979a3faa8976f6f81470d0761bf0a761202389650696bbc1749afb2a86
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-rules-html-a921588b

| Step 3 | nat [ ( real_ifc , mapped_ifc ) ] [ line \| { after-object [ line ]}] source static { nw_obj nw_obj \| any any } [ destination static { mapped_obj \| interface [ ipv6 ]} real_obj ] [ service real_src_mapped_dest_svc_obj mapped_src_real_dest_svc_obj ] [ no-proxy-arp ] [ route-lookup ] [ inactive ] [ description desc ] ciscoasa(config)# nat (inside,outside) source static MyInsNet MyInsNet destination static Server1 Server1 | Configures identity NAT . See the following guidelines:  – Mapped—Specify a network object or group, or for static interface NAT with port translation only, specify the interface keyword (routed mode only).If you specify ipv6 , then the IPv6 address of the interface is used. If you specify interface , be sure to also configure the service keyword (in this case, the service objects should include only the destination port). For this option, you must configure a specific interface for the real_ifc . See the “Static Interface NAT with Port Translation” section for more information. – Real—Specify a network object or group. For identity NAT, simply use the same object or group for both the real and mapped addresses.  | 
|  |  | (Continued)  | 
Configuring Per-Session PAT Rules
By default, all TCP PAT traffic and all UDP DNS traffic uses per-session PAT. To use multi-session PAT for traffic, you can configure per-session PAT rules: a permit rule uses per-session PAT, and a deny rule uses multi-session PAT. For more information about per-session vs. multi-session PAT, see the “Per-Session PAT vs. Multi-Session PAT” section.
Detailed Steps
To configure a per-session PAT rule, see the “Configuring Per-Session PAT Rules” section.
Monitoring Twice NAT
To monitor twice NAT, enter one of the following commands:
| show nat | Shows NAT statistics, including hits for each NAT rule. | 
| show nat pool | Shows NAT pool statistics, including the addresses and ports allocated, and how many times they were allocated. | 
| show xlate | Shows current NAT session information. | 
| show nat divert-table | All NAT rules build an entry in the NAT divert table. If the NAT divert field is set to ignore=yes NAT on the matching rule, the ASA stops the lookup and does a route lookup based on the destination IP to determine the egress interface. If the NAT divert field is set to ignore=no on the matching rule, walk the NAT table based on the found input_ifc and output_ifc and do the necessary translation. Egress interface will be output_ifc. | 
Configuration Examples for Twice NAT
This section includes the following configuration examples:
- Different Translation Depending on the Destination (Dynamic PAT)
- Different Translation Depending on the Destination Address and Port (Dynamic PAT)
Different Translation Depending on the Destination (Dynamic PAT)
Figure 5-1 shows a host on the 10.1.2.0/24 network accessing two different servers. When the host accesses the server at 209.165.201.11, the real address is translated to 209.165.202.129: port . When the host accesses the server at 209.165.200.225, the real address is translated to 209.165.202.130: port .
Figure 5-1 Twice NAT with Different Destination Addresses
Step 1 Add a network object for the inside network:
Step 2 Add a network object for the DMZ network 1:
Step 3 Add a network object for the PAT address:
Step 4 Configure the first twice NAT rule:
ciscoasa(config)# nat (inside,dmz) source dynamic myInsideNetwork PATaddress1 destination static DMZnetwork1 DMZnetwork1
Because you do not want to translate the destination address, you need to configure identity NAT for it by specifying the same address for the real and mapped destination addresses.
By default, the NAT rule is added to the end of section 1 of the NAT table, See the “Configuring Dynamic PAT (Hide)” section for more information about specifying the section and line number for the NAT rule.
Step 5 Add a network object for the DMZ network 2:
Step 6 Add a network object for the PAT address:
Step 7 Configure the second twice NAT rule:
ciscoasa(config)# nat (inside,dmz) source dynamic myInsideNetwork PATaddress2 destination static DMZnetwork2 DMZnetwork2
Different Translation Depending on the Destination Address and Port (Dynamic PAT)
Figure 5-2 shows the use of source and destination ports. The host on the 10.1.2.0/24 network accesses a single host for both web services and Telnet services. When the host accesses the server for Telnet services, the real address is translated to 209.165.202.129: port . When the host accesses the same server for web services, the real address is translated to 209.165.202.130: port .
Figure 5-2 Twice NAT with Different Destination Ports
Step 1 Add a network object for the inside network:
Step 2 Add a network object for the Telnet/Web server:
Step 3 Add a network object for the PAT address when using Telnet:
Step 4 Add a service object for Telnet:
Step 5 Configure the first twice NAT rule:
ciscoasa(config)# nat (inside,outside) source dynamic myInsideNetwork PATaddress1 destination static TelnetWebServer TelnetWebServer service TelnetObj TelnetObj
Because you do not want to translate the destination address or port, you need to configure identity NAT for them by specifying the same address for the real and mapped destination addresses, and the same port for the real and mapped service.
By default, the NAT rule is added to the end of section 1 of the NAT table, See the “Configuring Dynamic PAT (Hide)” section for more information about specifying the section and line number for the NAT rule.
Step 6 Add a network object for the PAT address when using HTTP:
Step 7 Add a service object for HTTP:
Step 8 Configure the second twice NAT rule:
ciscoasa(config)# nat (inside,outside) source dynamic myInsideNetwork PATaddress2 destination static TelnetWebServer TelnetWebServer service HTTPObj HTTPObj
Feature History for Twice NAT
Table 5-1 lists each feature change and the platform release in which it was implemented.
| Twice NAT | 8.3(1) | Twice NAT lets you identify both the source and destination address in a single rule. We modified or introduced the following commands: nat , show nat , show xlate , show nat pool .  | 
| Identity NAT configurable proxy ARP and route lookup | 8.4(2)/8.5(1) | In earlier releases for identity NAT, proxy ARP was disabled, and a route lookup was always used to determine the egress interface. You could not configure these settings. In 8.4(2) and later, the default behavior for identity NAT was changed to match the behavior of other static NAT configurations: proxy ARP is enabled, and the NAT configuration determines the egress interface (if specified) by default. You can leave these settings as is, or you can enable or disable them discretely. Note that you can now also disable proxy ARP for regular static NAT. For pre-8.3 configurations, the migration of NAT exempt rules (the nat 0 access-list command) to 8.4(2) and later now includes the following keywords to disable proxy ARP and to use a route lookup: no-proxy-arp and route-lookup . The unidirectional keyword that was used for migrating to 8.3(2) and 8.4(1) is no longer used for migration. When upgrading to 8.4(2) from 8.3(1), 8.3(2), and 8.4(1), all identity NAT configurations will now include the no-proxy-arp and route-lookup keywords, to maintain existing functionality. The unidirectional keyword is removed. We modified the following command: nat source static [ no-proxy-arp ] [ route-lookup ].  | 
