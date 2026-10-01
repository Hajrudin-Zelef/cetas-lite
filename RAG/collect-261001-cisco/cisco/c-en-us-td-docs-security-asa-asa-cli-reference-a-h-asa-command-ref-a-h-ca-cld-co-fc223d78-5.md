---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-5
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [580, 727]
sha256: d34120adc2abb680ac80cf287a6a1f304236e4ca8893f3b54d3e38f38d9e5cd7
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                                    | packet-length                                            bytes  |  (Optional) Sets the maximum number of bytes of each packet to store in the capture buffer. | 
                                 
                                 
                                    
                                    | persit                                                                                     |  (Optional) Captures persistent packets on cluter units. | 
                                 
                                 
                                    
                                    | port  |  (Optional) If you set the protocol to tcp or udp, specifies the integer or name of a TCP or UDP port.  | 
                                 
                                 
                                    
                                    | raw-data  |  (Optional) Captures inbound and outbound packets on one or more interfaces.  | 
                                 
                                 
                                    
                                    | real-time  |  Displays the captured packets continuously in real-time. To terminate real-time packet capture, enter Ctrl                                              +                                              c.                                               To permanently remove the capture, use the no  form of this command. This option applies only to raw-data , switch ,  and asp-drop  captures. This option is not supported when you use the cluster                                              exec                                              capture  command.                                         | 
                                 
                                 
                                    
                                    | reinject-hide  |  (Optional) Specifies that no reinjected packets will be captured. Applies only in a clustering environment. | 
                                 
                                 
                                    
                                    | stop  | (Optional) Manually stops the capture without removing it. Use the no  form of this command to start the capture.                                         | 
                                 
                                 
                                    
                                    | tls-proxy  | (Optional) Captures decrypted inbound and outbound data from TLS proxy on one or more interfaces. | 
                                 
                                 
                                    
                                    |                                            trace trace_count  | (Optional) Captures packet trace information, and the number of packets to capture. This option is used with an access list                                           to insert trace packets into the data path to determine whether or not the packet has been processed as expected.                                         | 
                                 
                                 
                                    
                                    | type  |  (Optional) Specifies the type of data captured. | 
                                 
                                 
                                    
                                    | user                                            webvpn-user  |  (Optional) Specifies a username for a WebVPN capture. | 
                                 
                                 
                                    
                                    | webvpn  |  (Optional) Captures WebVPN data for a specific WebVPN connection. | 
                                 
                              
                           
Command Default
                           
                           
                              The defaults are as follows:
                           
                           
                           
                              
                              - 
                                 
                                    The default type  is raw-data 
- 
                                 
                                    The default 
                                    buffer size  is 
                                    512 KB
                                    .
                                    
                                 
- 
                                 
                                    The default Ethernet type is IP packets.
                                 
- 
                                 
                                    The default packet-length  
                                    is 
                                    1518
                                    bytes.
                                    
                                 
- 
                                 
                                 The default direction  is ingress.
                                 
Command Modes
                           
                           
The following table shows the modes in which you can enter the command:
                           
                           
                           
                              
                              
                                 
                                 
                                 
                                 
                                 
                                 
                              
                              
                                 
                                 
                                    
                                    | Command Mode | Firewall Mode | Security Context | 
                                 
                                 
                                    
                                    | Routed | Transparent | Single | Multiple | 
                                 
                                 
                                    
                                    | Context | System | 
                                 
                              
                              
                                 
                                 
                                    
                                    | Privileged EXEC |  |  |  |  |  | 
                                 
                              
                           
Command History
                           
                              
                              
                                 
                                 
                              
                              
                                 
                                 
                                    
                                    |  Release |  Modification | 
                                 
                              
                              
                                 
                                 
                                    
                                    |  6.2(1) |  This command was added. | 
                                 
                                 
                                    
                                    |  7.0(1) |  This command was modified to include the following keywords: type                                              asp-drop , type                                              isakmp , type                                              raw-data , and type                                              webvpn .                                         | 
                                 
                                 
                                    
