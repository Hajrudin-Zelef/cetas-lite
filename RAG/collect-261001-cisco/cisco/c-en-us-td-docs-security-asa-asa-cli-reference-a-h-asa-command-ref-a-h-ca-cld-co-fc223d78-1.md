---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-1
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [1, 228]
sha256: b16fa89f2e989d1a69aef1645b8b0030433c0578f7e24f411d900b9e89e9a499
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

To enable packet capture capabilities for packet sniffing and network fault isolation, use the capture  command in privileged EXEC mode. To disable packet capture capabilities, use the no  form of this command.
                           
                           
                           
                              Capture network traffic:
                           
                           
                        
                        
                        
                           
                           
                           
                              
                              capture
                              capture_name
                              [ 
                              
                              type
                              { 
                              
                              asp-drop
                              [ 
                              all
                              | 
                              drop-code
                              ] 
                              
                              | 
                              tls-proxy
                              | 
                              raw-data
                              | 
                              
                              isakmp
                              [ 
                              ikev1
                              | 
                              ikev2
                              ] 
                              
                              | 
                              
                              inline-tag
                              [ 
                              tag
                              
                              ] 
                              | 
                              
                              webvpn
                              user
                              webvpn-user
                              
                              } 
                              ] 
                              
                              [ 
                              
                              access-list
                              access_list_name
                              
                              { 
                              
                              interface
                              { 
                              interface_name
                              | 
                              asa_dataplane
                              asa_mgmt_plane
                              | 
                              cplane
                              } 
                              } 
                              
                              [ 
                              
                              buffer
                              buf_size
                              
                              ] 
                              
                              [ 
                              ethernet-type
                              type
                              
                              ] 
                              [ 
                              reeinject-hide
                              ] 
                              [ 
                              
                              packet-length
                              bytes
                              ] 
                              
                              [ 
                              circular-buffer
                              
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
                              
                              
                              [ 
                              switch
                              ] 
                              [ 
                              offload
                              ] 
                              
                              [ 
                              ivlan
                              number
                              ] 
                              
                              
                              [ 
                              ovlan
                              number
                              ] 
                              
                              
                               
                           
                           
                           
                            Capture cluster control-link traffic:
                           
                           
                           
                              
                              capture
                              capture_name
                              {  
                              type lacp
                              interface
                              interface_id
                              [ 
                              buffer
                              buf_size
                              ] 
                              [ 
                              packet-length
                              bytes
                              ] 
                              [ 
                              circular-buffer
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
                              
                               
