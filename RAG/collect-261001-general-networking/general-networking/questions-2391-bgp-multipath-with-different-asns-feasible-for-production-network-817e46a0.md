---
id: collect-261001-general-networking/general-networking/questions-2391-bgp-multipath-with-different-asns-feasible-for-production-network-817e46a0
title: "questions-2391-bgp-multipath-with-different-asns-feasible-for-production-network-817e46a0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "research"]
source: docs/RAG/collect-261001-general-networking/questions-2391-bgp-multipath-with-different-asns-feasible-for-production-network-817e46a0.md
source_anchor: ""
source_lines: [1, 55]
sha256: f88d4fe473cba399520c18c40f632706bc29e90853547dd688f383c3a7395309
---

# questions-2391-bgp-multipath-with-different-asns-feasible-for-production-network-817e46a0

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
16
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
on Cisco (this command is hidden for some reason)
#bgp bestpath as-path multipath-relax
The default BGP behavior only installs only routes with exact the same AS_PATH into RIB. With multipath-relax, the AS_PATH only needs to be of the same length.
What problems can it potentially cause? Why isn't is used more often?
As a transit provider, does this feature complicate troubleshooting (I am thinking about end-user complaints about network performance)? Does it make it more difficult to know the path specific traffic took at a given time? Is there something else to that can assist troubleshooting. I am not sure about scalability and cost for NetFlow in SP network.
bgp bestpath as-path multipath-relax was introduced by CSCea19918. Normally eBGP load-balancing requires the candidate routes to be equal-cost paths; i.e. identical BGP attributes:
same Weight
same Local-Pref
same AS-Path (both the AS numbers, and the AS path length)
same Origin
same MED
different next-hop
As you mentioned, this command relaxes the same AS-Path requirement so any candidate eBGP AS-Path with the same AS-path length could be used for eBGP load-balancing (this will not load-balance between eBGP and iBGP paths). If you run BGP between multiple ISPs and you are looking for better egress load-balancing between your upstream connections this may help you out.
What problems can it potentially cause?
There isn't much danger as long as you're an enterprise customer that doesn't give transit service to another ASN; for a transit provider it might be perfectly safe, but I can't be sure there aren't routing loops if a transit ASN uses this feature. At first, I thought there would easily be a loop in transit ASN cases, on more reflection I can't find a real problem.
Why will it is rarely used?
Good question, it's been around since at least 2005.
The basic issue is that the BGP speaker configured with
"multipath-relax" gets into a control plane <-> data plane
inconsistency; i.e. it advertises only the best path, but installs
multiple paths in the forwarding that have different ASPATHs than the
best. This breaks the basic tool BGP has to detect loops - ASPATH loop
check.
A (distorted) scenario below. I am sure you can come up with
a better example with a bit more time at hand.
...............
: R4 AS1 (10/8)
/:..............
..../......
: R5 AS2
:....\.....
/ \ ...............
/ --:--R1
R6 AS4 : \ AS3
\--------:--- R2
: /
: R3 (10/8)
:..............
In this example,
- R3 in AS3 and R4 in AS1 announce a prefix 10/8. R5 in AS2 receives
the prefix from R1(AS3) and R4(AS1).
- AS2 is configured with 'multipath-relax' and chooses both paths
for multipath forwarding, though it selects AS1's path as best.
- R5 advertises the prefix with AS_PATH "2 1" to R6, and R6 in turn
to R2.
- Because of some specific policy, it is possible that R2(AS3)
chooses R6's path as best. If it happens, there is a loop.
Note that R1-R2-R3 represents the physical connectivity of
the routers in AS3.
