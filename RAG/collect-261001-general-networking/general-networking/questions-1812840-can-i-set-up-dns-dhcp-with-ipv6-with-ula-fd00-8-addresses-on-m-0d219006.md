---
id: collect-261001-general-networking/general-networking/questions-1812840-can-i-set-up-dns-dhcp-with-ipv6-with-ula-fd00-8-addresses-on-m-0d219006
title: "questions-1812840-can-i-set-up-dns-dhcp-with-ipv6-with-ula-fd00-8-addresses-on-m-0d219006"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1812840-can-i-set-up-dns-dhcp-with-ipv6-with-ula-fd00-8-addresses-on-m-0d219006.md
source_anchor: ""
source_lines: [1, 17]
sha256: 3e8bdd70ea6efcc0e098cfe8b636b985a89a50f824d0a17ce8910f8b401c04a6
---

# questions-1812840-can-i-set-up-dns-dhcp-with-ipv6-with-ula-fd00-8-addresses-on-m-0d219006

So I have written up a post to answer something like this in a thread over at Ubiquiti forums. In your case since you do not have a globally-routable IPv6 prefix from your ISP, a bunch of problems just don't exist for you.
I use dnsmasq as my DHCP and DNS server for IPv4 devices. Hence I could not use ISC DHCP server for DHCPv6 either (?). If you want to use ISC DHCP, you could :
- Enable the ipv6 > router-advert section of the selected interface ubnt@ubnt# set interfaces switch switch0 ipv6 router-advert 
- Suitably configure the router-advert to enable DHCPv6 stateful addressing: ubnt@ubnt# edit interfaces switch switch0 ipv6 router-advert
[edit interfaces switch switch0 ipv6 router-advert]
ubnt@ubnt# set managed-flag true
ubnt@ubnt# set other-config-flag true
ubnt@ubnt# set prefix fdXX::/64 
ubnt@ubnt# set prefix fdXX::/64 autonomous-flag false
ubnt@ubnt# exit
[edit]
- Set up the ULA for the interface in question of the router (ethX or switch0) : which you do by set interfaces switch switch0 address fdXX::1/64
But you said you want DNS integration of hostnames, which you can only do with dnsmasq - note that I have not yet succeeded in this part, although this is how dnsmasq is supposed to work. I do however have dnsmasq acting as a DHCPv6 server:
- Use the link above to activate dnsmasq as your DHCP server (obviously)
- Set up the range of addresses you want to hand out (adapting this for the edgeOS framework) : set service dns forwarding options dhcp-range=fd09::9,fd09::ffff,12h ,set service dns forwarding options enable-ra
- Of course, you must still assign a suitable address within this subnet to the router interface set interface switch switch0 address fdXX::1/64 .
In either case, you obviously have to commit and save. Or do all this in the Config Tree GUI on the web interface.
