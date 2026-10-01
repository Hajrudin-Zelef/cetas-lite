---
id: collect-261001-general-networking/general-networking/questions-13465-internal-bgp-route-exchange-issue-9ac02d58
title: "questions-13465-internal-bgp-route-exchange-issue-9ac02d58"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-13465-internal-bgp-route-exchange-issue-9ac02d58.md
source_anchor: ""
source_lines: [1, 62]
sha256: c8630ca6bca510c4327e05c0769e28d75747ccb1f1731452ef98750e16615a2e
---

# questions-13465-internal-bgp-route-exchange-issue-9ac02d58

I have BGP multihoming with two cisco routers. Each router has its own ISP and announce our AS with our /24 network to the internet.
My issue is between those two routers :
Router A, routing table looks like :
BGP table version is 9622808, local router ID is 10.100.100.2
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal,
              r RIB-failure, S Stale, m multipath, b backup-path, f RT-Filter,
              x best-external, a additional-path, c RIB-compressed,
Origin codes: i - IGP, e - EGP, ? - incomplete
RPKI validation codes: V valid, I invalid, N Not found
     Network          Next Hop            Metric LocPrf Weight Path
 *>i 1.0.0.0/24       10.100.100.3         10041    150      0 175 15169 i
 *                    192.168.200.1              0             0 800 15169 i
Router B, routing table looks like :
router#sh ip bgp
BGP table version is 10261936, local router ID is 192.168.1.2
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal,
              r RIB-failure, S Stale, m multipath, b backup-path, f RT-Filter,
              x best-external, a additional-path, c RIB-compressed,
Origin codes: i - IGP, e - EGP, ? - incomplete
RPKI validation codes: V valid, I invalid, N Not found
     Network          Next Hop            Metric LocPrf Weight Path
 *>  1.0.0.0/24       192.168.1.1        10041    150      0 175 15169 i
Router A
 
    #sh ip bgp 1.0.0.0
    BGP routing table entry for 1.0.0.0/24, version 9124223
    Paths: (2 available, best #1, table default)
    Multipath: eBGP
     Not advertised to any peer
    Refresh Epoch 1
    175 15169
    10.100.100.3 from 10.100.100.3 (192.168.1.1 )
      Origin IGP, metric 10041, localpref 150, valid, internal, best
      rx pathid: 0, tx pathid: 0x0
    Refresh Epoch 1
    8220 15169
    192.168.200.1 from 192.168.200.1 (212.74.90.252)
      Origin IGP, metric 0, localpref 100, valid, external
      Community: 538770940 538771322
      rx pathid: 0, tx pathid: 0
Router B
    #sh ip bgp 1.0.0.0
    BGP routing table entry for 1.0.0.0/24, version 9772009
    Paths: (1 available, best #1, table default)
    Multipath: eBGP
    Advertised to update-groups:
     15
    Refresh Epoch 1
    175 15169
    192.168.1.1 from 192.168.1.1 (38.28.4.10)
      Origin IGP, metric 10041, localpref 150, valid, external, best
      Community: 11424364 11425276
      rx pathid: 0, tx pathid: 0x0
Router B
    #sh ip bgp summary
   Neighbor        V           AS MsgRcvd MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
   10.100.100.2    4        48020  748914 2167278 10705132    0    0 2w4d         3350
   192.168.1.1   4          175 2591251   29738 10705000    0    0 2w4d       509418
My issue is that Router B has only one Next Hop for the 1.0.0.0/24 (this is an example). I am expecting that router B has another Next Hop to router A.
Router A is showing both Next hop as expected.
This issue has serious impact, If we lose router B as the preferred router we lose Internet, since it is not able to route traffic through router A.
sho ip bgp 1.0.0.0? I'm wondering if RouterA is seeing RouterB's path as better than it's external path (due to local pref) and not advertising its other known path over to RouterB. If this is correct, then RouterA will advertise its path automatically if the RouterB path goes away.
