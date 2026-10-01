---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-661636-port-forwarding-to-the-router-itself-on-edgerouter-bfb5fa8f
title: "questions-661636-port-forwarding-to-the-router-itself-on-edgerouter-bfb5fa8f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-661636-port-forwarding-to-the-router-itself-on-edgerouter-bfb5fa8f.md
source_anchor: ""
source_lines: [1, 29]
sha256: 0db5583666d6e85f62c7ffe12d8ca05020caa21b1389287aaf89223e3f0f48ac
---

# questions-661636-port-forwarding-to-the-router-itself-on-edgerouter-bfb5fa8f

I recently upgraded one of my customers to a Ubiquiti EdgeRouter Lite, which is a significant improvement over their old ISP-supplied router.
To reduce the frequency of attacks against the router's web interface while still allowing remote administration, one thing that we had done on the old router was to move remote management to a non-standard port, let's say 8642. On the old ISP-supplied router, there was a simple text-box for this, but on the Edgerouter it must be done by hand.
I added a simple Port-Forwarding rule on the Edgerouter to forward PUBLIC_IP:8642 to LOCAL_LAN_IP:443, as well as a corresponding firewall rule:
name WAN_LOCAL {
     default-action drop
     description "WAN to router"
     ...
     rule 2 {
         action accept
         description "Allow remote management"
         destination {
             group {
                 port-group ManagementPorts
             }
         }
         log disable
         protocol tcp
         state {
             established enable
             invalid disable
             new enable
             related enable
         }
     }
     ...
 }
where port-group ManagementPorts contained 8642.
However, I still could not access the web interface.  The only way I could find to resolve the issue was to allow outside access to port 443 as well - then access to port 8642 worked.  However, this means that the web interface is now available from outside on two ports, the default and the one I want.
What is the correct configuration for doing this so that the web interface is available internally on 443 and externally on 8642?
