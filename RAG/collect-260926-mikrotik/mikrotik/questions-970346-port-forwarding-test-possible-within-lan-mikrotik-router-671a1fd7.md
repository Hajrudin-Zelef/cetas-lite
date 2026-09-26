---
id: collect-260926-mikrotik/mikrotik/questions-970346-port-forwarding-test-possible-within-lan-mikrotik-router-671a1fd7
title: "questions-970346-port-forwarding-test-possible-within-lan-mikrotik-router-671a1fd7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-970346-port-forwarding-test-possible-within-lan-mikrotik-router-671a1fd7.md
source_anchor: ""
source_lines: [1, 45]
sha256: 6677a1460a9c3d95a25f2cbe1082f91002dd56d974aceae6d639c9dc4f2ed0c7
---

# questions-970346-port-forwarding-test-possible-within-lan-mikrotik-router-671a1fd7

Mikrotik RouterOS based on the Linux kernel and has inherited most of the conceptions of networking.
So, I'll describe the setup of the port forwarding in the Mikrotik routers. And I'll try to describe your issue.
If you want to understand it more deeply, you can read the iptables tutorial. It's pretty cool documentation with detailed explanations.
- Let's draw the network topology:
- Create the port forwarding rule. Better way specify the original destination address instead the interface. This allows use single rule to rewrite destination address for both for packets from outside and for packets from LAN. So, your rule is correct:
- Chain: dstnat
- Protocol: 6 (tcp)
- Dst Address: <wan-ip>
- Dst Port: 8000
- In. Interface: all Ethernet (for testing purposes)
- Action: dst-nat
- Log: yes (for testing purposes)
- To Addresses: 192.168.1.33
- To Ports: 8000
- When the packets arrive from outside (wan interface), the case is a trivial. But, when the packets, which should be port forwarded, arrive from LAN, something interesting happens. 
  - Let's guess the user of the LAN host tries to browse the http://<wan-ip>:8000 web page. The TCP packet in form192.168.1.Z:Y -> <wan-ip>:8000 TCP [SYN] is originated on the LAN host and is sent to the default gateway (Mikrotik router).
  - The mikrotik router receives this packet.
  - Then rewrites the destination address due your dst-nat rule. After this action the packet will looks like192.168.1.Z:Y -> 192.168.1.33:8000 TCP [SYN]
  - The router looks up the further path and sends the packet into the LAN to the 192.168.1.33 host.
  - The 192.168.1.33 host receives the TCP packet in form ofTCP SYN 192.168.1.Z:Y -> 192.168.1.33:8000 , create the reply in form of192.168.1.33:8000 -> 192.168.1.Z:Y TCP [SYN-ACK] and sends it to LAN host directly.
  - The LAN host receives the packet 192.168.1.33:8000 -> 192.168.1.Z:Y TCP [SYN-ACK] , but it isn't what the host expects! And this packet will be dropped.
  - As can you see, to make all work, the 192.168.1.33 host should send replies to the mikrotik router, not to the LAN host directly. To do it, you can add additionalsrc-nat rule to the Mikrotik.
- Chain: srcnat
- Protocol: 6 (tcp)
- Dst Address: 192.168.1.33
- Dst Port: 8000
- Src Address: 192.168.1.0/24
- Action: src-nat
- Log: yes (for testing purposes)
- To Addresses: 192.168.1.1
It makes the Mikrotik rewrites the source address in the port forwarded packets, originated in the LAN. After it the 192.168.1.33 host will see these packets as 192.168.1.1:X -> 192.168.1.33:8000 TCP [SYN], and send the reply to the Mikrotik. The mikrotik will do the reverse addresse translation and all will work.
- Another thing, what you should configure, is the firewall filter rules. In simple case the default rules are enough. If you prefer configure such thing youself, then you need a couple of rules in the filter/FORWARD chain:
- Chain: filter/FORWARD
- Protocol: 6 (tcp)
- Connection state: new
- Dst address: 192.168.1.33
- Dst port: 8000
- Action: accept
- Chain: filter/FORWARD
- Connection state: established,related
- Action: accept
Obviously, there are enough ways to improve, but the main concept is should be clear now.
     
    
src-natrule. Also, the Mikrotiks have built-in packet sniffer. You can use it to troubleshoot the issue.
