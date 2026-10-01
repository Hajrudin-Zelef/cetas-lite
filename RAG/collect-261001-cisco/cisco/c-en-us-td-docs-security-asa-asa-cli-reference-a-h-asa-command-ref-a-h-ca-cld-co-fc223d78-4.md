---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-4
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [532, 579]
sha256: 986c6b70e914b7739e5fdb04c1a6f641fe106fac3a72f64e591befe3cf8722b1
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                                    | drop  | Specifies the packet capture configuration of the mac-filter drop:                                                                                          | Note |                                                                                                                                                                                                                                   For the Secure Firewall 3100, drop  is available when you                                                             select the interface.                                                                                                                                                                                For the Secure Firewall 4200, the drop  keyword is available                                                             only when you select the direction. However, the mac-filter                                                              option is supported only for the ingress packet capture direction.                                                                                                                                                                               For the Secure Firewall 6100, the drop  keyword is available                                                             only for the ingress packet capture direction.                                                           |  | 
                                 
                                 
                                    
                                    | ethernet-type                                               type  |  (Optional) Selects an Ethernet type to capture. Supported Ethernet types include 8021Q, ARP, IP, IP6, LACP, PPPOED, PPPOES,                                           RARP, and VLAN. An exception occurs with the 802.1Q or VLAN type. The 802.1Q tag is automatically skipped and the inner Ethernet                                           type is used for matching.                                         | 
                                 
                                 
                                    
                                    | host                                            ip                                                                                    |  Specifies the single IP address of the host to which the packet is being sent. | 
                                 
                                 
                                    
                                    | include-decrypted  | (Optional) Captures decrypted IPsec packets which contain both normal and decrypted traffic once they enter the firewall device.                                           It also captures packets of SSL decrypted traffic. However, the capture does not include the decrypted packets from VTI because                                           they are available only on the VTI interface and not on the outside interface.                                         | 
                                 
                                 
                                    
                                    | inline-tag                                            tag                                                                                    |  Specifies a tag for a particular SGT value or leaves it unspecified to capture a tagged packet with any SGT value. | 
                                 
                                 
                                    
                                    | interface                                               interface_name                                                                                    |  Sets the name of the interface on which to use packet capture. You must configure an interface for any packets to be captured                                           except for type                                              asp-drop . You can configure multiple interfaces using multiple capture  commands with the same name. To capture packets on the dataplane, management plane, or control plane of an ASA, you can use                                           the interface  keyword with asa_dataplane , asa_mgmt_plane , or cplane                                               as the interface name.You can specify cluster  as the interface name to capture the traffic on the cluster control link interface. If the type lacp  capture is configured, the interface name is the physical name.                                         | 
                                 
                                 
                                    
                                    | ikev1  or ikev2  |  Captures only IKEv1 or IKEv2 protocol information. | 
                                 
                                 
                                    
                                    | isakmp  |  (Optional) Captures ISAKMP traffic for VPN connections. The ISAKMP subsystem does not have access to the upper layer protocols.                                           The capture is a pseudo capture, with the physical, IP, and UDP layers combined together to satisfy a PCAP parser. The peer                                           addresses are obtained from the SA exchange and are stored in the IP layer.                                          | 
                                 
                                 
                                    
                                    | lacp  |  (Optional) Captures LACP traffic. If configured, the interface name is the physical interface name. | 
                                 
                                 
                                    
                                    | mask  |  The subnet mask for the IP address. When you specify a network mask, the method is different from the Cisco IOS software                                           access-list command. The ASA uses a network mask (for example, 255.255.255.0 for a Class C mask). The Cisco IOS mask uses                                           wildcard bits (for example, 0.0.0.255).                                         | 
                                 
                                 
                                    
                                    | match protocol  |  Specifies the packets that match the five-tuple to allow filtering of those packets to be captured. You can use this keyword                                           up to three times on one line.                                         | 
                                 
                                 
                                    
                                    |  operator                                                                                     |  (Optional) Matches the port numbers used by the source or destination. The permitted operators are as follows:                                                                                                                                                                                    lt—less than                                                                                              gt—greater than                                                                                              eq—equal to                                                                                             neq —not equal to                                                                                                                                           range —range                                               | 
                                 
                                 
                                    
