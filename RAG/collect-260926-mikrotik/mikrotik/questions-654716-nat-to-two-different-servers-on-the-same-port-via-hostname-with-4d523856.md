---
id: collect-260926-mikrotik/mikrotik/questions-654716-nat-to-two-different-servers-on-the-same-port-via-hostname-with-4d523856
title: "questions-654716-nat-to-two-different-servers-on-the-same-port-via-hostname-with-4d523856"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-654716-nat-to-two-different-servers-on-the-same-port-via-hostname-with-4d523856.md
source_anchor: ""
source_lines: [1, 36]
sha256: 4596651f923a462ee102ca06b8452f0f1a7292fc9d415547e146ed4b7a8e6437
---

# questions-654716-nat-to-two-different-servers-on-the-same-port-via-hostname-with-4d523856

I have a Mikrotik RB2011 router, running RouterOS which connects to the internet via a static IP. In my lan I have two different servers, one that is on IP 192.168.89.11 and another on 192.168.89.12
My DNS (on cloudflare) resolves both myfirstserver.com and mysecondserver.com to my router's static IP.
Now, what I want to do is to somehow separate the traffic so all traffic for myfirstserver.com goes to 192.168.89.11 and traffic for mysecondserver.com goes to 192.168.89.12 (and both on port 80. I know I could just change ports but if these are servers publicly available, no user would set a different port than 80)
What I have tried so far, is to somehow mark the packets through mangle and then use that mark on NAT to do the proper dst-nat forward.
I try marking packets through either content or Layer 7 protocol regex (they work properly if the action is log. I can see them being logged correctly).
The thing is that after I mark them, it seems that NAT just ignores them and forwards the connection to the server that accepts the non-marked packets.
I think I have mixed up the order of filtering and the chains somehow.
Would someone be able to provide some pointers/assistance on how to accomplish this?
Thanks!
EDIT: My question is very specific, about marking packets in RouterOS on a Mikrotik router and checking the mark in NAT. it is not about whether i need a reverse proxy or whatever. Thanks
EDIT 2: @Cha0s asked my /ip firewall export so here it is:
/ip firewall layer7-protocol
add name=haf1a regexp=haf1a
/ip firewall address-list
add address=192.168.89.0/24 list=local
add address=192.168.88.0/24 list=local
add address=www.xxx.yyy.zzz list=local
add address=192.168.87.0/24 list=local
/ip firewall filter
add chain=input comment="default configuration" protocol=icmp
add chain=input comment="default configuration" connection-state=established
add chain=input comment="default configuration" connection-state=related
add action=drop chain=virus comment="Drop 80 DoS attack" src-address-list=spammer
add action=add-src-to-address-list address-list=spammer address-list-timeout=1d chain=input connection-limit=10,32 dst-address-list=!local dst-port=53 protocol=udp
add action=drop chain=input dst-port=53 protocol=udp src-address-list=!local
add action=add-src-to-address-list address-list=spammer address-list-timeout=1d chain=input connection-limit=30,32 protocol=tcp
/ip firewall mangle
add action=change-mss chain=forward new-mss=1422 out-interface=all-ppp protocol=tcp tcp-flags=syn tcp-mss=1423-65535
/ip firewall nat
add action=masquerade chain=srcnat comment="default configuration" out-interface=ether1-gateway
add action=masquerade chain=srcnat out-interface=pppoe-out1
add action=dst-nat chain=dstnat comment="No HAF mark" dst-address=www.xxx.yyy.zzz dst-port=80 packet-mark=!haf1 port="" protocol=tcp src-port="" to-addresses=192.168.89.41 to-ports=80
add action=dst-nat chain=dstnat comment="HAF mark" dst-address=www.xxx.yyy.zzz dst-port=80 packet-mark=haf1 port="" protocol=tcp src-port="" to-addresses=192.168.89.31 to-ports=80
where www.xxx.yyy.zzz is my external ip (static) and 192.168.89.31 and 192.168.89.41 my two servers inside my local network
My goal is to redirect any connections with haf1a in them (eg haf1a.lala.com) to the relevant server (192.168.89.31)
/ip firewall exporthere so we can understand your setup. Your described method (L7 regex + NAT rules) sounds right. But if the NAT rules don't match the packets properly there probably is an error in your configuration.
