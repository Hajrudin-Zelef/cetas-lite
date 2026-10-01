---
id: collect-261001-meraki/meraki/r-meraki-comments-6t1ecd-mx-autovpn-mtu-issues-5bdef7a7
title: "r-meraki-comments-6t1ecd-mx-autovpn-mtu-issues-5bdef7a7"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-6t1ecd-mx-autovpn-mtu-issues-5bdef7a7.md
source_anchor: ""
source_lines: [1, 18]
sha256: 8cdae11875f4c7f07aade2e608fd792daeb2bffa41043d568fd550b20a5dab38
---

# r-meraki-comments-6t1ecd-mx-autovpn-mtu-issues-5bdef7a7

MX AutoVPN MTU issues 
        
    Am I the only one who faces MTU issues from time to time on MX AutoVPN setups?
Sometimes the PathMTU Discovery does not work properly for whatever reasons (like the upstream provider router is set to "no ip unreachable" so the ICMP packets are never sent to the MX)
This always kills the performance completely, after calling in to support I let them change the MTU to 1400 on the internet ports and it starts working as a charm.
I mean why don't we get the option to configure it ourselves, maybe on the local device configuration page? And yes, I already made a wish :)
Section des commentaires
Yes we had this exact same issue and discovered that the only way to change is to call support. Been fine ever since we called them
Ok good so I'm not alone, may I ask where you're located? Also europe? Also a lot of PPPoE lines (so 8 Byte PPP overhead = 1492 MTU on wan link)
Maybe you can also send them a "wish" to give us the ability to change it via the dashboard?
Rural Canada with shitty Internet
Sadly, autovpn doesn't handle the case where an intermediate hop (ie a hop via a network segment non-local to either VPN peer) between two peers has a lower MTU than either of the peer's uplinks.
I appreciate your input but that's not completely true. If path MTU works as it's supposed to, the MX will change the MTU and MSS accordingly. I tested that behaviour in the lab and it was working fine.
The problem is that path mtu discovery relies on ICMP messages and we all know that ICMP could be leveraged for some dirty attacks, so it's not realiable at all.
I prefer to go for 1400 MTU and 1360 MSS which works usually
The issue is that if an encrypted packet gets dropped for being too big, the needs-frag doesn't propagate to the end host nor does the MX learn anything from it.
Are you doubting the Meraki way ?!? How dare you !
/s
