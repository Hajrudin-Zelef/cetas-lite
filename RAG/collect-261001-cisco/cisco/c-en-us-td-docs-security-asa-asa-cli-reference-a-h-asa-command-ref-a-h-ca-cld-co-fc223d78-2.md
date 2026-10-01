---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-2
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [229, 447]
sha256: f25f7568440d81a56e206c41b7fc25152ff7682821eb14e7ae3c0abae7a2779a
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                           
                           
                           
                              
                              capture
                              capture_name
                              interface cluster
                              [ 
                              buffer
                              buf_size
                              ] 
                              [ 
                              
                              ethernet-type
                              type
                              
                              ] 
                              
                              [ 
                              packet-length
                              bytes
                              
                              ] 
                              [ 
                              circular-buffer
                              ] 
                              [ 
                              cp-cluster
                              ] 
                              
                              [ 
                              trace
                              [ 
                              
                              trace-count
                              number
                              ] 
                              ] 
                              
                              
                              [ 
                              
                              real-time
                              [ 
                              dump
                              ] 
                              [ 
                              detail
                              ] 
                              ] 
                              
                              [ 
                              
                              match
                              protocol
                              { 
                              
                              host
                              source-ip
                              | 
                              
                              
                              source-ip
                              mask
                              | 
                              
                              any
                              | 
                              any4
                              | 
                              any6
                              } 
                              
                              [ 
                              operator
                              src_port
                              ] 
                              
                              { 
                              
                              host
                              dest_ip
                              
                              | 
                              
                              dest_ip mask
                              | 
                              
                              | 
                              any
                              | 
                              any4
                              | 
                              any6
                              } 
                              
                              [ 
                              operator dest_port
                              ] 
                              ] 
                              
                              
                              
                               
                           
                           
                           
                           Ingress switch capture packets  for Secure Firewall 3100 model devices:
                           
                           
                           
                              
                              capture
                              capture_name
                              switch
                              interface
                              interface_name
                              
                              [ 
                              drop
                              { 
                              disable
                              | 
                              mac-filter
                              }] 
                              
                              
                               
                           
                           
                           
                           Bi-directional switch packet capture for 1200/4200/6100 model devices:
                           
                           
                           
                           
                              
                              capture
                              capture_name
                              switch
                              interface
                              interface_name
                              [ 
                              direction
                              
                              
                              { { 
                              both
                              | 
                              egress
                              } 
                              [ 
                              drop disable
                              ] | 
                              ingress
                              [ 
                              drop
                              { 
                              disable
                              | 
                              mac-filter
                              }] }] 
                              
                              
                               
                           
                           
                           
                           
                              
                                 |  Note |  For Secure Firewall 4200 model devices, the mac-filter  option is supported only for the ingress  direction.                                          | 
                           
                              
                                 |  Note |  For Secure Firewall 6100 model devices, the drop  option is                                           supported only for the ingress direction.                                          | 
                           
 Capture packets cluster-wide:
                           
                           
                           
                              
                              cluster exec capture
                              capture_name
                              [ 
                              persist
                              ] 
                              [ 
                              include-decrypted
                              ] 
                              
                               
                           
                           
                           
                           Clear persistent packet traces cluster-wide:
                           
                           
                           
                              
                              cluster exec clear packet-trace
                              
                               
                           
                           
                           
                           Remove the packet capture:
                           
                           
                           
                              
                              no capture
                              capture_name
                              [ 
                              arguments
                              ] 
                              
                               
                           
                           
                           
