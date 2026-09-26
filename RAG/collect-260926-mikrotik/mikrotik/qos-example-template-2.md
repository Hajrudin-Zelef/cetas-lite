---
id: collect-260926-mikrotik/mikrotik/qos-example-template-2
title: "qos-example-template"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/lot-mikrotik/forum/qos/qos-example-template.md
source_anchor: ""
source_lines: [146, 272]
sha256: c97ef7806805b5bac38049ee5dc975868f9a30585c7a4a1023b14d8ce07bf7e7
---

# qos-example-template

              What I’ve done is to prioritise small TCP packets on the upload direction.  That way if a big upload is taking place the the small TCP “ack” packets relied on for download will still get through.  For streaming services you’ll need to look at what upstream control packets are send and do something similar.

             
            
           
          
            
            
              
DSCP=0 is the lowest possible priority.

             
            
           
          
            
            
              
Originally that was the case, but as 0 is also the default DSCP value that made it impossible to have below-normal priority e.g. for large transfers.

Therefore in most systems the DSCP values 8 and 16 are used to indicate lowest and one-but-lowest priority, both below the default 0.

When looking at the top 3 bits of DSCP/TOS the priority order is:

7 <= highest

6

5 <= e.g. EF (expedited forwarding), DSCP 46, often used for voice

4

3

0 <= default

2

1 <= lowest

             
            
           
          
            
            
              Oops, looks like I’m a little out of date.  However doing a quick catch up it looks like that convention is subject to change, with the latest RFC recommending 000001 for LE traffic, this being a value not previously classified.  Do you know which convention Mikrotik follows, specifically in the “new-priority=from-dscp-high-3-bits” action?   Does that prioritise CS1 above or below default?

             
            
           
          
            
            
              MikroTik follows no particular convention at all, I think.  Everything RouterOS does is just copying certain fields into others, it is up to the user to assign meaning to that.

The “priority” field is just a field assigned to each packet, it does not change the handling of the packet by itself.

There are some places where it is used:

- it is copied into the 802.1q VLAN header where it may be used by a switch.  usually in the switch it will be used to determine priority according to the scheme above.
- it can be used by Wireless when WMM is enabled.  effectively the top 2 bits will be used to select one of 4 queues, again using that scheme.

When you want to use it in a queue, you are on your own.  There is no direct way of using the priority field to determine priority in a queue tree.

Linux does support that, and it has a mapping table for it which you would then again fill with that strange sequence.

In RouterOS you need to use packet marks to accomplish that, like this:

```
/ip firewall mangleadd action=set-priority chain=postrouting comment="From dscp high 3 bits" \
    new-priority=from-dscp-high-3-bits passthrough=yes
add action=mark-packet chain=postrouting comment="Priority 0" \
    new-packet-mark=prio0 passthrough=no priority=0
add action=mark-packet chain=postrouting comment="Priority 1" \
    new-packet-mark=prio1 passthrough=no priority=1
add action=mark-packet chain=postrouting comment="Priority 2" \
    new-packet-mark=prio2 passthrough=no priority=2
add action=mark-packet chain=postrouting comment="Priority 3" \
    new-packet-mark=prio3 passthrough=no priority=3
add action=mark-packet chain=postrouting comment="Priority 4" \
    new-packet-mark=prio4 passthrough=no priority=4
add action=mark-packet chain=postrouting comment="Priority 5" \
    new-packet-mark=prio5 passthrough=no priority=5
add action=mark-packet chain=postrouting comment="Priority 6" \
    new-packet-mark=prio6 passthrough=no priority=6
add action=mark-packet chain=postrouting comment="Priority 7" \
    new-packet-mark=prio7 passthrough=no priority=7
/queue tree
add limit-at=30M max-limit=30M name=inet parent=pppoe-inet
    wireless-default
add limit-at=8M max-limit=28M name=inet-p1 packet-mark=prio7 parent=inet \
    priority=1 queue=default
add limit-at=8M max-limit=28M name=inet-p2 packet-mark=prio6 parent=inet \
    priority=2 queue=default
add limit-at=8M max-limit=28M name=inet-p3 packet-mark=prio5 parent=inet \
    priority=3 queue=default
add limit-at=8M max-limit=28M name=inet-p4 packet-mark=prio4 parent=inet \
    priority=4 queue=default
add limit-at=8M max-limit=28M name=inet-p5 packet-mark=prio3 parent=inet \
    priority=5 queue=default
add limit-at=8M max-limit=28M name=inet-p6 packet-mark=prio0 parent=inet \
    priority=6 queue=pcq-upload-default
add limit-at=8M max-limit=28M name=inet-p7 packet-mark=prio2 parent=inet \
    priority=7 queue=default
add limit-at=8M max-limit=28M name=inet-p8 packet-mark=prio1 parent=inet \
    queue=default
```

Note that the mark “prio2” refers to “top 3 bits of DSCP have value 002” here and that this is then translated into priority 7 for the queues.

It is a bit confusing that in the packet priority field higher value generally indicates higher priority, while in the queue tree priority field lower values is higher priority.

Here they are mapped by a series of rules, while in Linux itself you could do that with a single “priomap” table.

             
            
           
          
            
            
              
pardon me if a bit out of main topic, may i ask about bandwith allocation on above queue,

if we sum total limit at then they will exceed max limit, will the exceed will be dropped since we have set the max limit 30M only?

thank you
