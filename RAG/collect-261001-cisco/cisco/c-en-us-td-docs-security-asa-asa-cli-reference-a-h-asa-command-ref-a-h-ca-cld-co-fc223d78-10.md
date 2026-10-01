---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-10
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [1211, 1349]
sha256: 9b2c7bce8d8318d52cd38cefb7ef782a0b3631b2909c2c8eee6919ec8110b2e4
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                                 |  Note |  In multiple context mode, the access-list  and match  options are not available in the system context.                                          | 
                           
Examples
                           
                           
                           
                              Capture for Clustering
                           
                           
                           
                              To enable capture on all units in the cluster, you can add the cluster exec  keywords in front of each of these commands. 
                           
                           
                           
                              The following example shows how to create an LACP capture for the clustering environment:
                           
                           
ciscoasa (config)# capture lacp type lacp interface gigabitEthernet0/0
                           
                              The following example shows how to create a capture for control path packets in the clustering link:
                           
                           
ciscoasa (config)# cap cp interface cluster match udp any eq 49495 any 
ciscoasa (config)# cap cp interface cluster match udp any any eq 49495
                           
                              The following example shows how to create a capture for data path packets in the clustering link:
                           
                           
ciscoasa (config)# access-list cc1 extended permit udp any any eq 4193
ciscoasa (config)# access-list cc1 extended permit udp any eq 4193 any
ciscoasa (config)# capture dp interface cluster access-list ccl
                           
                              The following example shows how to capture data path traffic through the cluster:
                           
                           
ciscoasa (config)# capture abc interface inside match tcp host 1.1.1.1 host 2.2.2.2 eq www
ciscoasa (config)# capture abc interface inside match dup host 1.1.1.1 any
ciscoasa (config)# capture abc interface inside access-list xxx
                           
                              The following example shows how to capture logical update messages for flows that match the real source to the real destination,
                              and capture packets forwarded over CCL that match the real source to the real destination:
                           
                           
ciscoasa (config)# access-list dp permit
 real src real dst
                           
                              The following example shows how to capture a certain type of data plane message, such as icmp echo request/response, that
                              is forwarded from one ASA to another ASA using the match  keyword or the access list for the message type: 
                           
                           
ciscoasa (config)# capture capture_name interface cluster access-list match icmp any any 
                           
                              The following example shows how to create a capture by using access list 103 on a cluster control link in a clustering environment:
                           
                           
ciscoasa (config)# access-list 103 permit ip A B
ciscoasa (config)# capture example1 interface cluster 
                           
                              In the previous example, if A and B are IP addresses for the CCL interface, only the packets that are sent between these two
                              units are captured.
                           
                           
                           
                              If A and B are IP addresses for through-device traffic, then the following is true:
                           
                           
                           
                              
                              - 
                                 
                                    Forwarded packets are captured as usual, provided the source and destination IP addresses are matched with the access list.
                                 
- 
                                 
                                    The data path logic update message is captured provided it is for the flow between A and B or for an access list (for example,
                                    access-list 103). The capture matches the five-tuple of the embedded flow.
                                 
- 
                                 
                                    Although the source and destination addresses in the UDP packet are CCL addresses, if this packet is to update a flow that
                                    is associated with addresses A and B, it is also captured. That is, as long as addresses A and B that are embedded in the
                                    packet are matched, it is also captured.
                                 
                              The following example shows how to configure capture with persistent option:
                           
                           
cluster2-asa5585a(config)# cluster exec capture test interface outside trace persist 
  a(LOCAL):*************************************************************
  cluster2-asa5585a(config)#
                           
                              Now, you can send some traffic.
                           
                           
cluster2-asa5585a(config)# cluster exec show packet-tracer
 
  a(LOCAL):*************************************************************
  tracer 29/25 (allocate/freed), handle 29/25 (allocated/freed), error 0
  =======  Tracer origin-id a:23, hop 0 =======
  packet-id: Protocol: 0 src-port: 0 dst-port: 0
  Phase: 1
  Type: CAPTURE
  Subtype: 
  Result: ALLOW
  Config:
  Additional Information:
  MAC Access list
  Phase: 2
  Type: ACCESS-LIST
  Subtype: 
  Result: DROP
  Config:
  Implicit Rule
  Additional Information:
  MAC Access list
  Result:
  input-interface: outside
  input-status: up
  input-line-status: up
  Action: drop
  Drop-reason: (l2_acl) FP L2 rule drop
                           
                              The following example shows that, to free up some memory you must clear the captured persistent traces from the box.
                           
                           
ciscoasa# cluster exec clear packet-trace
                           
                              The following example displays how to configure the capture with include-decrypted option:
                           
                           
  cluster2-asa5585a(config)# cluster exec show capture
 
  a(LOCAL):*************************************************************
  capture in type raw-data trace interface outside include-decrypted [Capturing – 588 bytes] 
  capture out type raw-data trace interface outside include-decrypted [Capturing - 420 bytes] 
  cluster2-asa5585a(config)#
                           
                              Now, you can send some ICMP traffic through IPSec tunnel. The capture command obtains the decrypted ICMP packets as outlined:
                           
                           
