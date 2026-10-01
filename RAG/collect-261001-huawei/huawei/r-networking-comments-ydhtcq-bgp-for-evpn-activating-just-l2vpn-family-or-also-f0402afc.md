---
id: collect-261001-huawei/huawei/r-networking-comments-ydhtcq-bgp-for-evpn-activating-just-l2vpn-family-or-also-f0402afc
title: "r-networking-comments-ydhtcq-bgp-for-evpn-activating-just-l2vpn-family-or-also-f0402afc"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-ydhtcq-bgp-for-evpn-activating-just-l2vpn-family-or-also-f0402afc.md
source_anchor: ""
source_lines: [1, 20]
sha256: 84772ae7d95c5ad0065ff6667c3c1e2b23d9eb6e16997f07201bc7eb81480d9b
---

# r-networking-comments-ydhtcq-bgp-for-evpn-activating-just-l2vpn-family-or-also-f0402afc

BGP for EVPN, activating just l2vpn family or also ipv4 family? Huawei S5700 
        
        
        
    
    
    Background: we use OSPF as IGP protocol for the underlay and then BGP sourced by loopbacks in each of the VTEPs/pods.
For the BGP process I thought I could only activate the l2vpn family because I need to distribute just MAC addresses, no IPV4 routes. On Huawei S5700 switches it seems that if the ipv4 family is not active everything stops working though. I thought I may generalize and just say "activate the family you need and nothing else"; if you need not IPv4 then shut it down.
Is really like I descrived above or is there an explanation for that? Or might it a bug of the OS on the S5700?
      TIA,
Pan
    
Section des commentaires
With Huawei you need to define the peering in the main BGP config. By default only IPv4 unicast routes will be exchanged. You can disable the IPv4 exchange under ipv4-family unicast by adding the command undo peer ipaddress enable. You then use the command peer ipaddress enable under any supported address family that you require. You must always leave the primary peer configuration active.
At the beginning I disabled, or even removed the IPV4 family . Disabling peers in the IPV4 family is different than disabling it, maybe that's the culprit?
Are you defining
l2vpn-family evpnin the BGP config, then specifying the EVPN peers under that?
Yes,
l2vpn-family evpnpolicy vpn-targetpeer 192.0.2.2 enablepeer 192.0.2.3 enablepeer 192.0.2.5 enablepeer 192.0.2.7 enablepeer 192.0.2.8 enable
I’m not familiar with how Huawei does it, but where do you define the router-I’d for bgp process? Is it in the ipv4 address family? What happens if you define that but not activate any neighbors in ipv4.
