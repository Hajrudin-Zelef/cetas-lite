---
id: collect-260926-mikrotik/mikrotik/questions-1839170-mikrotik-port-forwarding-c349cae7
title: "questions-1839170-mikrotik-port-forwarding-c349cae7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1839170-mikrotik-port-forwarding-c349cae7.md
source_anchor: ""
source_lines: [1, 17]
sha256: 40a71a884e3d3b0e35bd75693c592ca0797a633d149745af7ff99789893aedf3
---

# questions-1839170-mikrotik-port-forwarding-c349cae7

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
First, the DNAT rule only performs address rewriting – you also need a filter rule that allows the packets to be forwarded. Add one under IP > Firewall > Filter to the forward chain.
"Port forwarding" in home routers usually combines both rules into one configuration, but under the hood they're two separate things and RouterOS also keeps them separate (very much like iptables does on Linux).
'Forward' filtering is checked after DNAT (but before SNAT), so if you're adding individual rules for each forwarded port, the filter rule should check for the new destination address and port, e.g.: dst-address=192.168.1.236 protocol=tcp dst-port=5252 action=accept.
Alternatively, you can add a catch-all rule that automatically allows everything that has already matched a DNAT rule by checking conntrack, e.g.: connection-nat-state=dstnat action=accept
Second, look at your rule counters to make sure the rule is actually being matched; it could be that your ISP (or something else along the way) prevents the inbound packets from arriving at your router in the first place. If needed, enable logging on your rules or even add some additional action=log, and use the packet capture tool (/tool/sniffer/quick) if you have any doubts about packets arriving.
Finally, external tests can't really distinguish between "rejected by th erouter" and "rejected by the internal host", as the internal host borrows the router's IP address so they both look the same from the outside. That is to say, even if your port-forwarding rules are all good, the server itself needs to accept the connection – it needs to have something listening on that port and its firewall rules need to accept it on input – for the port to appear "open".
So the next step is troubleshooting why it isn't working.
If you have a laptop, connect your wifi with your mobile hotspot to place yourself outside of the network.
Assuming you are on windows, download paping.
After you saved it, open a command prompt and navigate to the location you saved paping to.
run paping and ping to the WAN IP and port: eg paping 123.45.67.89 -p 5252
Paping will now test the connection every second. Either it times out (with red text) or it pings succesfully (with green text).
If its green, everything works.
While the paping runs, go to your NATRule, statistics page, and make sure you see spikes in the graph. If you see them, the mikrotik side is setup correctly. If you don't, then the NAT rule is not even reached.
If the NAT rule works, make sure that the server you're connecting to is not blocking the connection or that the service is down. also double check the ip address.
If the NAT rule doesn't work, ensure that you are testing from OUTSIDE and that internal routing and firewall rules are not blocking it. Also, you may need to add a mascerade rule if you are indeed inside.
