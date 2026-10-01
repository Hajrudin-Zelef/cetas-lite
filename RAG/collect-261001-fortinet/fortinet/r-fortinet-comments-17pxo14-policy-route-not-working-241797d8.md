---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-17pxo14-policy-route-not-working-241797d8
title: "r-fortinet-comments-17pxo14-policy-route-not-working-241797d8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-17pxo14-policy-route-not-working-241797d8.md
source_anchor: ""
source_lines: [1, 36]
sha256: 9d620001c1432d436149b4d36d5ed9ac40b485dfc9552bebce9dc826fc6e9ed4
---

# r-fortinet-comments-17pxo14-policy-route-not-working-241797d8

Policy route not working 
        
    We just replaced an ASA with a Fortigate 100f.
We have 3 VDOMs.
      INT - 10Gb interface with a bunch of VLANs
WAN - only the Wan1 interface
Root - with an interface to the WAN
    
The reason we do this is the WAN VDOM sends traffic to a CATO Networks device that does additional filtering including TLS Inspection.
In the INT VDOM, I have a static route that sends 0.0.0.0/0 to the WAN VDOM.
I also have a policy route that sends all traffic from a guest VLAN to the Root VDOM.
The reason we do this, is I don't want guest traffic to go through TLS inspection.
I see no traffic on the policy route.
My understanding is that policy routes are evaluated first.
The policy route is as follows:
      Oncoming Interface: VLAN410
Source Address: 410 address
Destination Address: All
Action: Forward Traffic
Outgoing interface: Root VDOM link.
    
What am I missing?
Section des commentaires
There was a route missing. Thank you everybody.
What is the default gateway for the Root VDOM?
ISP.
I added the interVDOM link IP as the gateway to the policy. I see traffic hitting the policy, but not returning. A tracert to 8.8.8.8 still shows the traffic going to the WAN VDOM
Do you have a route back from the root VDOM to the guest vlan over the VDOM link?
Do you have a valid route in the firewall's routing table to the Cato box?
Yes. But I'm trying to not route that traffic to CATO.
I'm seeing traffic hit the policy route. I'm seeing traffic hit the FW rule from that VLAN to the INT-Root link. I'm not seeing traffic in the Root VDOM FW rule that says from Any interface to the Wan interface, all traffic
Run a debug flow in the root VDOM to see if traffic is arriving correctly.
Diagnose the flow, my money is on URPF blocking it. URPF only looks at the routing table not policy routes.
Kindly refer to the below and make sure you have the correct route:
Ref:
https://docs.fortinet.com/document/fortigate/6.2.15/cookbook/335646/inter-vdom-routing
