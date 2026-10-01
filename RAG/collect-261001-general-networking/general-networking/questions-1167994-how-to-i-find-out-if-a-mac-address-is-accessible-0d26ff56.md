---
id: collect-261001-general-networking/general-networking/questions-1167994-how-to-i-find-out-if-a-mac-address-is-accessible-0d26ff56
title: "How to I find out if a mac address is accessible?"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-1167994-how-to-i-find-out-if-a-mac-address-is-accessible-0d26ff56.md
source_anchor: ""
source_lines: [1, 21]
sha256: 2c90abd6882e948164a1aa253d7c0b732256854cc037e8e0d08ac59236219885
---

# How to I find out if a mac address is accessible?

*Score : 0 | Source : https://serverfault.com/questions/1167994/how-to-i-find-out-if-a-mac-address-is-accessible*

I have a printer that has suddenly become inaccessible, and I can no longer connect to it. Suspecting DHCP, I configured it for a static IP address, so I know both the IP address and the MAC address.
Whenever I ping the IP address, on my Windows box, I get a Destination Host Unreachable.
I did a Ping from my OpnSense router, and got successful response. Then I did a Ping from my Windows box, and it returned host unreachable. At the same time, the OpnSense box started returning host unreachable too.
Both machines and the printer are plugged into ethernet ports on the same network switch.
While it's possible that the JetDirect card is failing, that seems unlikely.
When I do an arp -a on the OpnSense box, I see neither the mac address nor the IP address listed.
How do I confirm that the printer network interface is accessible on the network, and what should I consider when trying to work out why it is no longer accessible?
My current suspicion is some sort of security measure is kicking in, but I have no idea where to start diagnosing this problem.

---

### Reponse — score 1

I would do it in following run tcpdump -v -n ether host [mac_address] and arp and try to see any what are the IP address request for an ARP. Example:
ARP, Request who-has 192.168.0.11 tell 192.168.0.244, length 46
ARP, Reply 192.168.0.11 is-at 11:22:33:44:55:68, length 28
...or Wireshark (can't remember the mac filter syntax)
