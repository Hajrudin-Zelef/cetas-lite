---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1825446-unifi-dream-machine-special-edition-udm-se-iptables-nat-prerou-7e55953b
title: "questions-1825446-unifi-dream-machine-special-edition-udm-se-iptables-nat-prerou-7e55953b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1825446-unifi-dream-machine-special-edition-udm-se-iptables-nat-prerou-7e55953b.md
source_anchor: ""
source_lines: [1, 19]
sha256: 2555597a2bd23574cc6be23e094e79c64fd16d6be2221b3c334a8a4f1e3d4dfa
---

# questions-1825446-unifi-dream-machine-special-edition-udm-se-iptables-nat-prerou-7e55953b

This is a very confusing situation for me and I feel like there isn't an answer to this. But I'm at a loss after hours of testing, figuring out how things are currently working or not working as I would expect. If you're brave enough to read this post, thank you in advance.
Gateway 192.168.1.1 (UDM SE)
DNS server resides on 192.168.1.2
Initial Firewall Rules:
iptables -t nat -A PREROUTING ! -s 192.168.1.2 -p tcp --dport 53 -j DNAT --to 192.168.1.2
iptables -t nat -A PREROUTING ! -s 192.168.1.2 -p udp --dport 53 -j DNAT --to 192.168.1.2
iptables -t nat -A POSTROUTING -m iprange --src-range 192.168.1.4-192.168.4.254 -j MASQUERADE
This works great. It forces anyone doing a DNS request like 'nslookup site.tld 8.8.8.8' to get a response from 192.168.1.2 instead of 8.8.8.8 on all 4 subnets. This is the expected and desired result. For reasons I won't bore you with I need to change things around and discovered unexpected and unwanted results.
Change One Delta (Post Routing Range):
iptables -t nat -A PREROUTING ! -s 192.168.1.2 -p tcp --dport 53 -j DNAT --to 192.168.1.2
iptables -t nat -A PREROUTING ! -s 192.168.1.2 -p udp --dport 53 -j DNAT --to 192.168.1.2
iptables -t nat -A POSTROUTING -m iprange --src-range 192.168.2.1-192.168.4.254 -j MASQUERADE
With this configuration devices on the 192.168.1.x subnet can 'nslookup site.tld 192.168.1.2' but will timeout on 'nslookup site.tld 8.8.8.8'. The other three subnets ( 192.168.2.1-192.168.4.1 ) get results from 192.168.1.2 doing lookups either way. If I delete the 2 PREROUTING entries, 'nslookup site.tld 8.8.8.8' resolves from 8.8.8.8 with no timeout on the 192.168.1.x devices and again on the other 3 subnets. Why is the MASQUERADE rule required in order for the 192.168.1.x devices to resolve when the PREROUTING rules were present?
Change Two Delta (DNAT --to)
iptables -t nat -A PREROUTING ! -s 192.168.1.2 -p tcp --dport 53 -j DNAT --to 208.67.222.123
iptables -t nat -A PREROUTING ! -s 192.168.1.2 -p udp --dport 53 -j DNAT --to 208.67.222.123
iptables -t nat -A POSTROUTING -m iprange --src-range 192.168.1.4-192.168.4.254 -j MASQUERADE
Doing an 'nslookup welcome.opends.com 192.168.1.2' or 'nslookup welcome.opends.com 8.8.8.8' from the three subnets ( 192.168.2.1-192.168.4.1 ) gives me the expected results of using 208.7.222.123 as the DNS and correctly returns 146.112.59.8. From the 192.168.1.x devices 'nslookup welcome.opends.com 8.8.8.8' returns 146.112.59.8 as I would expect because of the PREROUTING rules. However, doing a 'nslookup welcome.opends.com 192.168.1.2' returns 146.112.59.9 meaning that the query went through 192.168.1.2 instead of 208.67.222.123 as defined by the PREROUTING rules. Why aren't these 192.168.1.x devices getting pre-routed regardless of the DNS server specified during nslookup? Clearly they get intercepted when using 8.8.8.8 with nslookup.
Thank you for looking at this monster post.
