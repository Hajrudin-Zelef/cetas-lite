---
id: collect-260926-rattrapage/rattrapage/questions-840402-multicast-causes-link-down-on-mikrotik-rb2011uias-rm-switch2-779e5bfe
title: "questions-840402-multicast-causes-link-down-on-mikrotik-rb2011uias-rm-switch2-779e5bfe"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/servers-reviews/questions-840402-multicast-causes-link-down-on-mikrotik-rb2011uias-rm-switch2-779e5bfe.md
source_anchor: ""
source_lines: [1, 51]
sha256: 41aaca3903ae3a540d4f42f03adfce490832479d56c33db7377a50656a25b4ff
---

# questions-840402-multicast-causes-link-down-on-mikrotik-rb2011uias-rm-switch2-779e5bfe

I am trying to simply push multicast traffic through the second switch of the router... for now I connected swtich 2 port 8 to my ISP router iPTV port and port9 to SagemonUHD88 IPTV decoder. I set master port port 6, and added slaves port 8,9.
I removed all ports in switch2 from default bridge (no bridges beside). After few second of multicast traffic the log clamin link down... and goes on an on every few second. I tried firmware 6.38.5 and 6.37.5. Both the same issue.
Just to be sure, I got connector between two RJ-45 plugs so when I disconnect the cables from mikrotik and connect it directly it works without any issues (I mean, the video on TV goes smoothly, no breaks).
I wanted it to just have wirespeed. When I will replace the default router (the Livebox 2.0 fro my ISP) i will try push tagged vlan with that multicast and untag on the port to the IPTV decoded. However for now I am now able to do that and as far as I see RB2011 can only tag frames on the way out from switch (egress), so no way for internal tagging the frames from specific port.
mar/24 00:54:55 route,debug,calc Begin calculation 
mar/24 00:54:55 route,debug,event Link up 
mar/24 00:54:55 route,debug,event     interface=ether8 
mar/24 00:54:55 route,debug,event Update 
mar/24 00:54:55 route,debug,event     interface=ether8 
mar/24 00:54:55 route,debug,event Link up 
mar/24 00:54:55 route,debug,event     interface=ether9 
mar/24 00:54:55 route,debug,event Update 
mar/24 00:54:55 route,debug,event     interface=ether9 
mar/24 00:54:55 route,debug,calc End calculation 
mar/24 00:54:56 interface,info ether6-master link up 
mar/24 00:54:56 route,debug,event Interface change 
mar/24 00:54:56 route,debug,event     interface=ether6-master 
mar/24 00:54:56 route,debug,event     status=UP,RUNNING 
mar/24 00:54:56 route,debug,event     mtu=1500 
mar/24 00:54:56 route,debug,calc Begin calculation 
mar/24 00:54:56 route,debug,event Link up 
mar/24 00:54:56 route,debug,event     interface=ether6-master 
mar/24 00:54:56 route,debug,event Update 
mar/24 00:54:56 route,debug,event     interface=ether6-master 
mar/24 00:54:56 route,debug,calc End calculation 
mar/24 00:55:01 interface,info ether6-master link down 
mar/24 00:55:01 interface,info ether8 link down 
mar/24 00:55:01 route,debug,event Interface change 
mar/24 00:55:01 route,debug,event     interface=ether6-master 
mar/24 00:55:01 route,debug,event     status=UP 
mar/24 00:55:01 route,debug,event     mtu=1500 
mar/24 00:55:01 route,debug,event Interface change 
mar/24 00:55:01 route,debug,event     interface=ether8 
mar/24 00:55:01 route,debug,event     status=UP 
mar/24 00:55:01 route,debug,event     mtu=1500 
mar/24 00:55:01 interface,info ether9 link down 
mar/24 00:55:01 route,debug,event Interface change 
mar/24 00:55:01 route,debug,event     interface=ether9 
mar/24 00:55:01 route,debug,event     status=UP 
mar/24 00:55:01 route,debug,event     mtu=1500 
mar/24 00:55:01 route,debug,calc Begin calculation 
mar/24 00:55:01 route,debug,event Link down 
mar/24 00:55:01 route,debug,event     interface=ether6-master 
mar/24 00:55:01 route,debug,event Update 
mar/24 00:55:01 route,debug,event     interface=ether6-master 
mar/24 00:55:01 route,debug,event Link down 
mar/24 00:55:01 route,debug,event     interface=ether8 
mar/24 00:55:01 route,debug,event Update 
mar/24 00:55:01 route,debug,event     interface=ether8 
mar/24 00:55:01 route,debug,event Link down 
mar/24 00:55:01 route,debug,event     interface=ether9
