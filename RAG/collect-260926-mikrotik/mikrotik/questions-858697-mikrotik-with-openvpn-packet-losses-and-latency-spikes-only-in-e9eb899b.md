---
id: collect-260926-mikrotik/mikrotik/questions-858697-mikrotik-with-openvpn-packet-losses-and-latency-spikes-only-in-e9eb899b
title: "questions-858697-mikrotik-with-openvpn-packet-losses-and-latency-spikes-only-in--e9eb899b"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "research"]
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-858697-mikrotik-with-openvpn-packet-losses-and-latency-spikes-only-in--e9eb899b.md
source_anchor: ""
source_lines: [1, 50]
sha256: 2500928c5d0ce87b59a7333d75a79bedd1fa52caf5a552a49262467d238a3b33
---

# questions-858697-mikrotik-with-openvpn-packet-losses-and-latency-spikes-only-in--e9eb899b

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I apologize, if this question had already been asked and answered, but I was looking for an answer to this for a while without any luck.
The background: we have several Mikrotik routers with RouterOS v6.39.2 that are located all over the city; the routers are behind NAT and connected to our OpenVPN server that we use for management and monitoring. The protocol we use for OpenVPN is TCP, since UDP for OpenVPN is not supported by Mikrotik.
The problem: from time to time we observe response delays that occur randomly. There is no specific router or time when the issue occurs; we are constantly monitoring the CPU load and traffic and there are no anomalies on the graphs at the time this is happening, but when we simultaneously ping the OpenVPN server from the Mikrotik via public and private IP address using the Mikrotik ping utility we see that packets drop only inside the tunnel.
Here is the /etc/openvpn/server.conf
dev tun
port 1194
proto tcp
dh .key/dh1024.pem
keepalive 10 120
user nobody
group nogroup
persist-key
persist-tun
status openvpn-status.log
log-append openvpn.log
verb 3
username-as-common-name
client-config-dir /etc/openvpn/ccd
ifconfig-pool-persist ipp.txt
ca .key/ca.crt
cert .key/server.crt
key .key/server.key
server 10.48.0.0 255.255.128.0
plugin /usr/lib/openvpn/openvpn-plugin-auth-pam.so "login login USERNAME password PASSWORD"
client-cert-not-required
client-to-client
push "route 10.48.0.0 255.255.128.0"
push "route-gateway 10.48.0.1"
The OpenVPN implementation on MikroTik is crippled. It can barely work and its performance is terrible (either on bandwidth or latency), especially on non x86 routers.
People have been asking for UDP and LZO support for ages and MikroTik simply refuses to implement those.
Here's Normis' (MikroTik staff) reply on the subject:
OpenVPN is very very buggy and hard to implement. Our developers
almost all committed suicide trying to make it work. It's a big mess,
so we can't continue to implement it 100%
This reply is from 2010 and they still haven't implemented those features. So you can draw your own conclusions about MikroTik's OpenVPN implementation.
Here's the whole thread if you are interested https://forum.mikrotik.com/viewtopic.php?f=1&t=26499
If you need better stability/performance (and security) you should look into some other solution like GRE over IPsec, or EoIP over IPsec.
I've been using these protocols for many years and they work flawlessly.
But if you can't get the routers to work without NAT then you could try SSTP, but since this is also TCP based, the performance and latency are not the best. There's also L2TP over IPsec.
In any case, your problem is an implementation one at the MikroTik side, so you can't actually fix it. Only MikroTik can. You can only try other protocols and see what fits your needs better.
Of course there's always the possibility that your configuration is bad (as mentioned in the comments already, FastTrack rules can mess up a lot of stuff in MikroTIk if you don't know what you are doing).
It is possible to use UDP on mikrotik hardware. You can install a WRT instance in the metarouter. Its a buggy implementation. But it is possible to use a meta to get OpenVPN UDP.
I had used it before with iPECS phones in "VPN" mode.
I will add to this solution a little more when i dig up my notes, might help someone else in the future.
