---
id: collect-260926-mikrotik/mikrotik/ipsec-tunnel
title: "ipsec-tunnel"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/ipsec-tunnel.md
source_anchor: ""
source_lines: [1, 43]
sha256: c87a7ee311963f400d74ab86c93d18fd620e96f8a203f0d1d37f1d4da63db3d0
---

# ipsec-tunnel

1.Can someone tell me how many tunnels can i configure that would  work at the same time in rb750.

1. in the ipsec configuration in the wiki there are all the steps to establish a tunnel between 2 networks. but what are the rules in the firewall that i have to configure in order to make the tunnel work.???

lets say that first network is: public 1.1.1.1 lan 2.2.2.2

and second network is: public: 3.3.3.3 lan:4.4.4.4

thanks.

             
            
           
          
            
            
              I have been able to run up to 10 tunnels with out a problem on RB450G’s and the RB1100

You will want your “/ip ipsec policy” to look like this for each tunnel

0   src-address=2.2.2.0/24 src-port=any dst-address=4.4.4.0/24

dst-port=any protocol=all action=encrypt level=require

ipsec-protocols=esp tunnel=yes sa-src-address=1.1.1.1

sa-dst-address=3.3.3.3 proposal=default priority=2

When you set up the other end you can just have it generate the policy

The side you want to make initial contact you will also need to have your  “/ip firewall nat” to look something like this and make sure it has a high priority



0   chain=srcnat action=accept src-address=2.2.2.0/24

dst-address=4.4.4.0/24 out-interface=“your WAN interface”

Then make sure you are not filtering this traffic in your /ip firewall filter

Here is a link with a nice instructional video on how to set up tunnels  http://gregsowell.com/?p=787

Hope this helps
