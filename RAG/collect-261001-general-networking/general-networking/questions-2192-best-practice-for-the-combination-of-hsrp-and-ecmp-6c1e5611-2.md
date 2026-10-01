---
id: collect-261001-general-networking/general-networking/questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611-2
title: "questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611.md
source_anchor: ""
source_lines: [76, 105]
sha256: 8125fec701f2fbc102462bfe6bc1d69072dc1b6244b27ea111ef0a1136d796a1
---

# questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611

Enable HSRP to periodically send Gratuitous ARP packets. Granted, this is similar to altering timers, but it's a much more graceful alteration than manipulating the CAM table and ARP timers. (Note though that this depends on your hardware and software combination, not all HSRP implementations offer this.)
By default, HSRP sends 3 GARPs, at 0, 2, and 4 seconds after the router becomes the forwarding gateway. However, there is a configuration parameter that allows you to choose the number of GARPs (including "infinite") and the interval.
I use MC-LAG pretty extensively, particularly VSS, VPC, and Clustering (I'm not a fan of stacking).
Where I can't use MC-LAG or GLBP, this is what I apply to my campus L2/L3 boundary routers (I have a 350-building campus so I use Cat6k pretty heavily):
Cat6k-v15(config)#interface vlan 100
Cat6k-v15(config-if)#standby arp ?
gratuitous Setup gratuitous ARP interval and count
Cat6k-v15(config-if)#standby arp gratuitous ?
count Set HSRP gratuitous ARP count
interval Set HSRP gratuitous ARP interval
<cr>
Cat6k-v15(config-if)#standby arp gratuitous count ?
<0-60> Number of gratuitous ARPs to send after group is activated (0 for continuous)
Cat6k-v15(config-if)#standby arp gratuitous count 0 ?
count Set HSRP gratuitous ARP count
interval Set HSRP gratuitous ARP interval
<cr>
Cat6k-v15(config-if)#standby arp gratuitous count 0 interval ?
<3-1800> Gratuitous ARP Interval (sec)
Cat6k-v15(config-if)#standby arp gratuitous count 0 interval 60 ?
count Set HSRP gratuitous ARP count
interval Set HSRP gratuitous ARP interval
<cr>
Cat6k-v15(config-if)#standby arp gratuitous count 0 interval 60
(I would post references to all these, but I don't have a high-enough "reputation" on this site to post more than two URLs.)
I just realized my original comment is valid - but woefully incomplete. Vendor-neutral design recommendation is to build in triangles, not rectangles. So:
Not just MC-LAG, but MC-LAG at both layers. Then you're dealing with a shared CAM table at both the the switch level and the router level.
If you can't do that, MC-LAG either the router or switch, and MC-LAG to the other layer with additional link (i.e. full-mesh between routers and switches). STP will ensure loop-free topology.
If you can't do that, still full-mesh the routers and switches. STP will ensure loop-free topology, and the switch CAM tables will still know all the appropriate MAC forwarding rules. The server will always send it's MAC, and if you configure the HSRP GARPs on 1-min intervals the switches will also not forget the HSRP vMAC.
Preferred options are in that order. But at the very least, install that extra pair of links.
