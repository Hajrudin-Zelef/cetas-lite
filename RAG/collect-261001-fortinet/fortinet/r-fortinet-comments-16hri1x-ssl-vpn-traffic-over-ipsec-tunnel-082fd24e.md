---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-16hri1x-ssl-vpn-traffic-over-ipsec-tunnel-082fd24e
title: "r-fortinet-comments-16hri1x-ssl-vpn-traffic-over-ipsec-tunnel-082fd24e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-16hri1x-ssl-vpn-traffic-over-ipsec-tunnel-082fd24e.md
source_anchor: ""
source_lines: [1, 26]
sha256: b26b818682a2784e6c9724531adfc556d9a1c920f00f2a2791179a02569ebfb0
---

# r-fortinet-comments-16hri1x-ssl-vpn-traffic-over-ipsec-tunnel-082fd24e

SSL VPN traffic over IPSEC tunnel 
        
    After some help, I have the following:
IPSEC tunnel between two 60F’s traffic flowing both ways no issues there.
SSL VPN to site one is working with no issues there. Can connet and access service on site one LAN.
What I need to achieve is to allow SSL VPN traffic to continue to access the site one LAN but also flow across the IPSEC tunnel to access the site two LAN.
I’ve tried adding the site two (remote) LAN address range to the SSL VPN policy, but that did not work. I also tried creating a secondary SSL VPN policy to take all traffic from the SSL VPN interface to the IPSEC Interface with the SSL VPN range as the source and the remote LAN range as the destination and that has not worked either.
I’ve run through this guide SSL VPN to IPsec VPN | FortiGate / FortiOS 7.2.5 | Fortinet Document Library although this does appear to be just to send SSL VPN traffic into site one send across to site two, but that has not helped either. I may be missing something, but I can't work it out.
Just after some thoughts and pointers. Both firewalls are on 7.2.5
Section des commentaires
Do you have matching phase 2s set up on both sides for the SSL VPN?
Do you have the SSL VPN in your firewall policies on both sides?
Do you have the routes built on both sides for the SSL VPN subnet?
Not all of that gets built in the wizard.
When you say SSL VPN do you mean the ipsec tunnel. Yes the selectors are correct on both sides. I set up the tunnel manually and also tore it down again to build with the wizard to see if that helped. Same result each time. Ipsec traffic is fine and SSL VPN traffic is fine. Just SSL VPN not flowing across the Ipsec tunnel.
No, I mean your client VPN but do you have that interface/subnet in your S2S VPN policies?
Like someone else mentioned, you probably just missed the route.
On the site FortiGate that you can't get to via the VPN, you need to add a static route pointing towards the SSL-VPN traffic on the IPSec tunnel.
Also, if it's split tunnel makes sure you have the subnets advertised on the VPN settings portal and the necessary firewall policies
Thanks going to re-check all this shortly. I'm sure I have missed something simple somewhere.
I have exactly the same situation - did you manage to resolve your issue (like you I think I'm missing something simple)? Just wondering if there's any wisdom/learning you're willing to share - thanks in advance
Me too I have the same problem. anybody accomplish this ?
Did you ever find a solution for this?
Hello,
I have had the same scenario and to provide these accesses to the output interface I have set it to any and at the routing level it is in charge of routing the traffic to where it wants to go.
Thanks will look into this.
