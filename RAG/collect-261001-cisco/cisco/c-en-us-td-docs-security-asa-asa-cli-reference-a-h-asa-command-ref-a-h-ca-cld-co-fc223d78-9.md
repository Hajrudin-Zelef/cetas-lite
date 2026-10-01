---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-9
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [1070, 1210]
sha256: be707975fc3f0cdce2c200f11867000ab45ab30322086db87922633742c09fb5
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                                       The Marvell 98CX8540 switch used in 4200 and 6100 Series device has a hardware limitation. Due to this limitation IPv6 packets
                                          with extension headers cannot be captured using egress switch packet capture when filters are based on IP protocol or TCP/UDP
                                          port ranges.
                                        This is because the egress processing is unaware of the Layer 4 protocol type.
Examples
                           
                           
                              To capture a packet, enter the following command:
                           
                           
ciscoasa# capture captest interface inside
ciscoasa# capture captest interface outside
                           
                              On a web browser, you can view the content of the capture  command that was issued, named “captest,” at the following location:
                           
                           
https://171.69.38.95/admin/capture/captest
                           
                              To download a libpcap file (that web browsers use) to a local machine, enter the following command:
                           
                           
https://171.69.38.95/capture/http/pcap
                           
                              The following example shows how to capture a packet in the single-mode when the ASA box crashes:
                           
                           
ciscoasa# capture 123 interface inside
                           
                              The contents of capture ‘123’ is saved as 123.pcap
                              file.
                           
                           
                           
                              The following example shows how to capture a packet in the multi-mode when the ASA box crashes:
                           
                           
ciscoasa# capture 456 interface inside
                           
                              The contents of capture ‘456’ in ‘admin’ context is saved as admin.456.pcap
                              file.
                           
                           
                           
                              The following example shows that the traffic is captured from an outside host at 171.71.69.234 to an inside HTTP server:
                           
                           
ciscoasa# access-list http permit tcp host 10.120.56.15 eq http host 171.71.69.234
ciscoasa# access-list http permit tcp host 171.71.69.234 host 10.120.56.15 eq http
ciscoasa# capture http access-list http packet-length 74 interface inside
                           
                              The following example shows how to capture ARP packets:
                           
                           
ciscoasa# capture arp ethernet-type arp interface outside
                           
                              The following example inserts five tracer packets into the data stream, where access-list 101  defines traffic that matches TCP protocol FTP:
                           
                           
hostname# capture ftptrace interface outside access-list 101 trace 5
                           
                              To view the traced packets and information about packet processing in an easily readable manner, use the show capture ftptrace  command.
                           
                           
                           
                              The following example shows how to display captured packets in real-time:
                           
                           
ciscoasa# capture test interface outside real-time
Warning: Using this option with a slow console connection may result in an excess amount of non-displayed packets due to performance limitations.
Use ctrl-c to terminate real-time capture.
10 packets displayed
12 packets not displayed due to performance limitations
                           
                              The following example shows how to configure an extended access list that matches the IPv4 traffic that needs to be captured:
                           
                           
ciscoasa (config)# access-list capture extended permit ip any any
                           
                              The following examples shows how to configure the capture:
                           
                           
ciscoasa (config)# capture name access-list acl_name interface interface_name
                           
                              By default, configuring a capture creates a linear capture buffer of size 512 KB. You can optionally configure a circular
                              buffer. By default, only 68 bytes of the packets are captured in the buffer. You can optionally change this value.
                           
                           
                           
                              The following example creates a capture called “ip-capture” using the capture access list previously configured that is applied
                              to the outside interface: 
                           
                           
ciscoasa (config)# capture ip-capture access-list capture interface outside
                            The following example creates a capture called “switch-capture” on outside interface for Secure Firewall 3100: 
                           
ciscoasa (config)# capture switch-capture switch interface outside drop ?
exec mode commands/options:
  disable Disable capturing dropped packets from switch
  mac-filter To capture switch mac-filter drop
ciscoasa(config)# capture switch-capture switch interface outside drop mac-filter
                           
                              The following example shows how to view the capture:
                           
                           
ciscoasa (config)# show capture name
                           
                              The following example shows how to end the capture, but retain the buffer:
                           
                           
ciscoasa (config)# no capture name access-list acl_name interface interface_name
                           
                              The following example shows how to end the capture and delete the buffer:
                           
                           
ciscoasa (config)# no capture name
                           
                              The following example shows how to filter traffic captured on the backplane in single mode:
                           
                           
ciscoasa# capture x interface asa_dataplane access-list any4
ciscoasa# capture y interface asa_dataplane match ip any any
                           
                              
                                 |  Note |  Control packets are captured in the single mode even though you have specified the access list.  | 
                           
                              The following examples show how to filter traffic captured on the backplane in multiple context mode:
                           
                           
                           
                              Usage in user context:
                           
                           
ciscoasa (contextA)# capture x interface asa_dataplane access-list any4
ciscoasa (contextA)# capture y interface asa_dataplane match ip any any
                           
                              Usage in system context:
                           
                           
ciscoasa# capture z interface asa_dataplane
                           
                              
