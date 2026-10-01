---
id: collect-261001-general-networking/general-networking/questions-12942-configuring-a-second-mpls-connection-for-our-backup-data-center-9fb40bf0
title: "questions-12942-configuring-a-second-mpls-connection-for-our-backup-data-center--9fb40bf0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "research"]
source: docs/RAG/collect-261001-general-networking/questions-12942-configuring-a-second-mpls-connection-for-our-backup-data-center--9fb40bf0.md
source_anchor: ""
source_lines: [1, 25]
sha256: 66470bebcfed0d2da3f283b8398ab3dfdd292cf00b50e9b643d80e0eaa88837e
---

# questions-12942-configuring-a-second-mpls-connection-for-our-backup-data-center--9fb40bf0

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm a bit new to BGP and need a little assistance with a configuration change. We are currently using our MPLS connection for all traffic to all sites. Going forward We would want to segragate the Datacenter backup link from the other spoke sites and will be adding an additional MPLS leg in our core datacenter to accommodate. See image of current setup and router configs.
I have ordered a new MPLS leg with a Closed User Group(CUG) from the ISP. This new service is for datacenter to datacenter traffic only. The existing MPLS leg will be used exclusivly for our remote sites.(Hub-Spoke)
My question is how do I program the core router with this new CUG that has the same AS number from the ISP (AS207). Do I just add the new network and neighbor to the existing AS 65001? I believe I only need to change the core router because the backup DC MPLS will be integrated into the new CUG.
This is a production network which is why I'm a little apprehensive making this program change.
There a few inconsistencies in your diagram / configurations. In the diagram it is AS 207, in the configuration it is AS 209 and in the final configuration it is AS 207.
If the provider AS is the same in both MPLS networks then:
I'll call your original network CUG-1 and your new one CUG-2.
There won't be any problems with advertising the routes to CUG-2 as
it will be in a different VPN to CUG-1. If you don't do any filtering
then the routes will be automatically advertised to CUG-1 via the
core (provided you aren't separating them with VRFs on that device).
You will need to rewrite the AS number before advertising it out to
the spokes from your core and vice versa.
If you don't do this then:
Backup DC routes will be dropped by the CUG-1 PE router due to the BGP loop prevention mechanism i.e. routers in AS207 will see their own AS in the AS-Path ( e.g. AS 65002 -> AS 207 -> AS 65001 -> AS 207DROP).
Spoke routes will be dropped by the CUG-2 PE router for the same reason (AS 65004 -> AS 207 -> AS 65001 -> AS 207DROP).
If you use BGP AS-override then the AS-Path will look something like this on e.g. Backup DC to Spoke Site: AS 65001 -> AS 65001 -> AS 65001 -> AS 207 -> AS 65004.
Because you will be circumventing the BGP loop prevention mechanism, ensure that the Core DC is the only location that is peering with both CUGs.
Wrote this on mobile so let me know if it is a bit disjointed and I'll clean it up later.
