---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-1epr2rm-wireguard-issues-after-upgraded-to-247-70bb8660
title: "r-opnsense-comments-1epr2rm-wireguard-issues-after-upgraded-to-247-70bb8660"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-1epr2rm-wireguard-issues-after-upgraded-to-247-70bb8660.md
source_anchor: ""
source_lines: [1, 87]
sha256: 41747bb51d35cc4b1278b9523180fc8543ed89efef6ba1804743658e70da11de
---

# 
       WireGuard issues after Upgraded to 24.7 

      
    Hi,

I upgraded to 24.7 (now on 24.7.1) and my WireGuard VPN seems to have issues. I'm using MullVad, have tried recreating the Instance and Peer and in the WireGuard settings shows as Up.

Have assigned the wg0 interface, and added to Gateway. The Gateway is showing as Online. But, no clients on the can connect.

Anybody else had similar issues? What info would I need to share to help troubleshoot this issue?

UPDATE:

I contacted Nullvad and they suggested following this guide:

I followed this, and the section on Selective Routing and its running again. I hope that this helps any others that have issues getting this working. Also, make sure you get the Outbound NAT rules set-up correctly.

I've been trying to setup Mullvad for the first time today. Still haven't got it all working, but I did notice that no guides mention you need to set the Keepalive in the Peer or it won't work at all.

I couldn't figure out wtf was going on, so I just randomly tried stuff and found that without a number in Keepalive, I can't get a handshake, as soon as I add a number, I can.

I’ve had various issues with setting up WireGuard on OPNsense. Here’s a few things to check out

- 
    have you setup firewall rules to allow vpn traffic?
- 
    have you set a keepalive interval on the peer?
- 
    have you tried explicitly setting the dns servers on the peer?

- 
    have you setup firewall rules to allow vpn traffic?

Yes. WireGuard used to work fine. It seems to have stopped in one of the latest updates.

- 
    have you set a keepalive interval on the peer

Yes

- 
    have you tried explicitly setting the dns servers on the peer?

No, how does one do that? I can not see any such options on the Peer?

Just to add:

* VPN / Status - I seem to be getting a Handshake for the Peer

* VPN / Status - Instance Status = up

* System / Gateway - WireGuardInterface status is Online

But, when I try to route traffic through the VPN nothing.

Any tips on troubleshooting / working out where the block / issue is?

Please note that this was working about a month ago. So, something, perhaps an update, has changed a requirement / setting perhaps?

In the [Interface] block on your peer's configuration, add a DNS server that points to something that's accessible from the WireGuard peer (I have unbound setup, so I use 192.168.1.1). For some reason that helped my peer's connect.

Thank you for your reply.

Can you please elaborate - where do you mean re "[Interface] block on your peers config"?

I just setup everything for the first time and I was having issues. Didn’t realize you have to restart the WireGuard service after every new peer gets created/generated so you can try that if you haven’t. I’m only using WireGuard though.

THank you. I had / have tried restating the service and the server. Alas, not fixed the issue.

I did have WireGuard working before, following this:

It was working fine before. But, for some reason, perhaps an upgrade, has caused it to stop. Not sure if it was 24.7 or 24.3 but used to work great. Now, can not connect clients to the WG / VPN Gateway.

I followed the same tutorial recently to set it up on 24.7 but I did get some stuff wrong at first and troubleshooting it was a challenge, what helped was doing a packet capture on both the client and wireguard interfaces, and turning on logging for all the rules involved including the outbound NAT (though logging the NAT rules seemed to suppress further logging of firewall rules on the same packets - turn logging back off on NAT rules when you know they work), then work from there.

Edit: one pitfall I did run into and did not see mentioned in the thread nor the tutorial was the MTU of the hosts routed through WG being set too high for the WG interface and OPNsense refusing to fragment the packets, which the capture showed me.

What is your setup like?

for me i had to port forward my vlan20 DNS to the same VLAN address then redirect it to 192.168.1.1 to get it to work.

https://i.imgur.com/ADHPcXD.png

I initially tried to set it up all myself but ultimately didn’t get it working until following this guide: https://docs.opnsense.org/manual/how-tos/wireguard-client.html

https://docs.opnsense.org/manual/how-tos/wireguard-selective-routing.html
