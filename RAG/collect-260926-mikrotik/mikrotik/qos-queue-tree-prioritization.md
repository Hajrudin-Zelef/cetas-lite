---
id: collect-260926-mikrotik/mikrotik/qos-queue-tree-prioritization
title: "qos-queue-tree-prioritization"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/qos-queue-tree-prioritization.md
source_anchor: ""
source_lines: [1, 39]
sha256: a2538b073c6352ad010d368e520bef8e884a77d37c1c66c136f94265197cd3ec
---

# qos-queue-tree-prioritization

Hi,

i am having some issued regarding QOS on several vlans.

My ISP is providing 500/150 and my idea was to balance the bandwith between 3 vlans.

vlan10 - Private vlan (should have unlimited bandwith and highest priority, in case guests are using 400 i would like to take bandwith from them if i need it) or do i need to set a Max Limit for my private vlan?

vlan20 - Guest vlan (should be on PCQ0 and limited to max 400/100

vlan 30 - TV vlan (should be on PCQ0 and limited to max 100/20

I dont want to use pcq 0 on private network due to a lot of different hubs for smart home which are using very low bandwith. (could i keep on “default small”?)

Does my priotiry and queue type for my private network make sense ? (screenshot attached)

Thanks for your help!

 
            
           
          
            
            
              Start here 

Most likely the example config you’re looking for and answer is in this thread:

http://forum.mikrotik.com/t/using-routeros-to-qos-your-network-2020-edition/66683/1

             
            
           
          
            
            
              Thanks!

Will give it a try
