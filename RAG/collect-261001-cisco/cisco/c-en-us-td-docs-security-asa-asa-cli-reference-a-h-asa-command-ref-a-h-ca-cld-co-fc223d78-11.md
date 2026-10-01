---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-11
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [1350, 1390]
sha256: b8fba98fba6d929e64a8374fa4ad8c2dacd7248410d027fc81a57795fc7962a9
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

  cluster2-asa5585a(config)# cluster exec show capture in | i icmp
  a(LOCAL):*************************************************************
  b:********************************************************************
  cluster2-asa5585a(config)# cluster exec show capture out | i icmp
  a(LOCAL):*************************************************************  
  b:********************************************************************
  cluster2-asa5585a(config)# cluster exec show capture in | i icmp
  a(LOCAL):*************************************************************
     8: 07:22:57.065014       802.1Q vlan#212 P0 211.1.1.1 > 213.1.1.2: icmp: echo request   
  b:********************************************************************
  cluster2-asa5585a(config)# cluster exec show capture out | i icmp
  a(LOCAL):*************************************************************
    10: 07:22:57.068004       802.1Q vlan#214 P0 213.1.1.2 > 211.1.1.1: icmp: echo reply
  b:********************************************************************
  cluster2-asa5585a(config)#
                           The following example shows how to create and start an egress traffic capture for a switch: 
                           
ciscoasa(config)# capture switch_cap switch interface gigabitEthernet0/0 direction ?
exec mode commands/options:
  both     To capture switch bi-directional traffic
  egress   To capture switch egressing traffic
  ingress  To capture switch ingressing traffic
 
ciscoasa(config)# capture switch_cap switch interface gigabitEthernet0/0 direction egress
ciscoasa(config)# no capture switch_cap switch stop
                           
                           If you want to capture packets where the accelerated security path (ASP) drop reason is dispatch-queue-limit, you need to
                              enable debug to delay the packet drops. Otherwise, the packets have been dropped and they cannot be captured. If you have
                              used show asp drop  command and see that dispatch queue tail drops (dispatch-queue-limit) are increasing, enable debug on asp 38 before performing
                              the capture. Do the following: 
                           
                           
(Enable delayed drop for 20 packets. You can specify 0-100. 
When the delayed drop count hits zero, the delayed drop 
behavior stops automatically. You can see how many drops 
are still available using debug menu asp 38 with no number)
ciscoasa(config)# debug menu asp 38 20
(Start the capture)
ciscoasa(config)# capture drop type asp-drop dispatch-queue-limit
(Generate the traffic that has been causing the drops, then view the capture.)
ciscoasa(config)# show capture drop
