---
id: collect-261001-cisco/cisco/r-cisco-comments-qvgmua-eigrp-neighbor-relation-wont-form-across-vzn-epl-227be147
title: "r-cisco-comments-qvgmua-eigrp-neighbor-relation-wont-form-across-vzn-epl-227be147"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-qvgmua-eigrp-neighbor-relation-wont-form-across-vzn-epl-227be147.md
source_anchor: ""
source_lines: [1, 20]
sha256: 88aa149ba8057eb885c0851f96f366db72c46b8b98ae5f4f913da94f1f8ec682
---

# r-cisco-comments-qvgmua-eigrp-neighbor-relation-wont-form-across-vzn-epl-227be147

EIGRP Neighbor Relation won't form across VZN EPL 
        
    EIGRP keeps timing out on the C9300 side and won't even form on the ASR side. Both side show as Up/Up on the ports. The system mtu on the C9300 side is set for 9000. Link was working fine until Saturday when it went down. Called Verizon and they told me "No Trouble Found" on their link between their equipment. Nothing changed on our equipment in that time.
Not exactly sure what could have cause this link to go down and I've tried many things like adjusting the speed, mtu, negotiations, rebooted the switch, changed the port. Did a debug EIGRP packet on the switch side and I see it sending the hello packets but not getting anything from the router and vice versa.
Is there something I should try or should I keep pressing Verizon?
Section des commentaires
eigrp configs
C9300router eigrp 1no default-information innetwork 10.10.90.244 0.0.0.3passive-interface defaultno passive-interface TenGigabitEthernet1/0/46eigrp router-id 10.102.222.1ASR 9000router eigrp 1address-family ipv4router-id 10.160.1.3log-neighbor-changesredistribute static route-policy static-redinterface Loopback0interface GigabitEthernet0/0/1/10route-policy OUT outprefix-set WAN-OUT### SUMMARY ONLY ###10.10.0.0/16route-policy WAN-OUTif destination in WAN-OUT thenpasselsedropendifend-policy
You can try adding static EIGRP neighbors on both ends, which will transition to unicast coms for the relevant interface(s) as opposed to default-coms (multicast).
This should let you confirm whether it's multicast packets that are being dropped, assuming that you haven't already ruled that out.
If you could also share your EIGRP configs, it may be helpful for myself or others to give you more input.
A lot of these virtual circuits block or throttle multicast, so it's pretty likely that this is the issue.
this ^ ^ ^ ^
share the specifics for your EIGRP config---both sides please.
Possibly much related if a change happened in verizons end. Ping from one side to the other with an extended ping. See if it works at 1500 and also at 9000.
You should probably check and see if your traffic is truly being sent, and truly arriving between the two routers first....
What do you see in a packet capture?
Check spanning tree to make sure it's not hosing you.
No spanning tree on L3 routed ports.
No but if you do so?
