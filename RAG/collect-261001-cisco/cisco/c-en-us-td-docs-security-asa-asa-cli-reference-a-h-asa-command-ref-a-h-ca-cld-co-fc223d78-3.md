---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-3
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [448, 531]
sha256: 5bceea29a4e156ae193368ddcec8b7e3fbda295ab7f967384a8b122950da1bb0
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                           Manually stop or start the packet capture:
                           
                           
                           
                              
                              capture
                              capture_name
                              stop
                              
                               
                           
                           
                           
                              
                              no capture
                              capture_name
                              stop
                              
                               
                           
                           
                           
                        
                        
                        
                           Syntax Description
                           
                              
                              
                                 
                                 
                              
                              
                                 
                                 
                                    
                                    | access-list access_list_name                                                |  (Optional) Captures traffic that matches an access list. In multiple context mode, this is only available within a context. | 
                                 
                                 
                                    
                                    | any  |  Specifies all IPv4 traffic. | 
                                 
                                 
                                    
                                    | any4  |  Specifies all IPv4 traffic. | 
                                 
                                 
                                    
                                    | any6  |  Specifies all IPv6 traffic. | 
                                 
                                 
                                    
                                    | all                                                                                     |  Captures all packets dropped by the accelerated security path. | 
                                 
                                 
                                    
                                    | asa_dataplane  |  Captures packets on the ASA backplane that pass between the ASA and a module that uses the backplane, such as the ASA FirePOWER                                           module.                                         | 
                                 
                                 
                                    
                                    | asp-drop                                               drop-code                                                                                    |  (Optional) Captures packets dropped by the accelerated security path. The drop-code specifies the type of traffic that is dropped by the accelerated security path. See the show                                              asp                                              drop                                              frame  command for a list of drop codes. You can enter this keyword with the packet-length , circular-buffer , and buffer                                               keywords, but not with the interface  or ethernet-type  keyword. In a cluster, dropped forwarded data packets from one unit to another are also captured. In multiple context mode,                                           when this option is issued in the system execution space, all dropped data packets are captured; when this option is issued                                           in a context, only dropped data packets that enter from interfaces belonging to the context are captured.                                         | 
                                 
                                 
                                    
                                    | buffer                                            buf_size  |  (Optional) Defines the buffer size used to store the packet in bytes. Once the byte buffer is full, packet capture stops.                                           When used in a cluster, this is the per-unit size, not the sum of all units.                                         | 
                                 
                                 
                                    
                                    | capture_name  |  Specifies the name of the packet capture. Use the same name on multiple capture  statements to capture multiple types of traffic. When you view the capture configuration using the show                                              capture  command, all options are combined on one line.                                         | 
                                 
                                 
                                    
                                    | circular-buffer  | (Optional) Overwrites the buffer, starting from the beginning, when the buffer is full. | 
                                 
                                 
                                    
                                    | cp-cluster  | (Optional) Capture control packets on cluster interface. | 
                                 
                                 
                                    
                                    | direction  | (Optional. Supported on: Secure Firewall 1220, Secure Firewall 4200, Secure Firewall 6100 devices.) Specifies the direction of the switch traffic to be captured. It can be                                           one of the following:                                                                                                                                                                                                                           both —To capture switch bi-directional traffic                                                                                                                                                                                            egress —To capture switch egressing traffic                                                                                                                                                                                            ingress —To capture switch ingressing traffic                                               | 
                                 
                                 
                                    
