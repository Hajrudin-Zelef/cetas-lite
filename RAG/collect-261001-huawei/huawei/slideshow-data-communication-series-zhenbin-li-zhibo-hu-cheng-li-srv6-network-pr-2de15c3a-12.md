---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-12
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["China"]
dates: ["2007-08-31", "2008-03-31", "2013-03-02", "2013-08-16", "2014-07-28", "2015-10-14", "2018-12-19", "2019-11-25", "2019-12-05", "2019-12-06", "2020-01-21", "2020-02-04", "2020-03-09", "2020-03-14", "2020-03-25"]
keywords: ["distribution", "latency"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1601, 1736]
sha256: b9e2dc5d70d59b8390962eeb1c2fe79dbdaeb06c90463b5a8001993ca78c7084
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

much to do with MPLS as a basic transport technology. In turn,
this has resulted in complex functions and difcult deployment
of the SDN controller. Te native IP attribute of SRv6 greatly
simplifes basic transport technologies. Furthermore, SRv6 can
be deployed from the edge to the core, in a staggered manner,
from SRv6 VPN to loose TE, strict TE, and more. Tis also
enables SDN to be gradually rolled out on carrier IP networks in
a simple-to-complex manner.
3. 5G changes the attributes of connections, and cloud changes the
scope of connections. Tese changes bring big opportunities for
SRv6 development.
73.
26 ◾ SRv6Network Programming
Network connections are the core of IP transport networks,
and the development of 5G services poses more requirements on
network connections, such as stronger Service Level Agreement
(SLA) assurance and deterministic latency. 5G also changes or,
more precisely, enhances the attributes of connections, requiring
more information to be carried in packets. In addition, with the
development of cloud services, service processing locations have
become more fexible. Some cloud services (such as telco cloud)
furtherbreaktheboundarybetweenphysicalandvirtualnetwork
devices, integrating services and transport networks. All these
have changed the scope of network connections. SRv6 enables
the programming of services and transport networks based on
a single data plane. Also, thanks to its native IP attribute, SRv6
allows rapid setup of connections and satisfes requirements of
fexible adjustment of the connection scope.
III. Development of IP Generations
We used to believe that IP was developed in an incremental and
compatible manner. Tat is, IP does not have clear defnitions of gen-
erations such as 2G, 3G, 4G, and 5G like wireless. However, when we
look back upon the development of IP over the past few decades, we
fnd some intergenerational characteristics.
First is the rise and decline of network protocols. In the 1990s,
the competition between ATM and IP led to the decline of telecom-
munications. IP transport networks unifed networks by replac-
ing independent networks such as ATM, FR, and Time Division
Multiplexing (TDM). With the rise of SR, traditional MPLS signaling
Label Distribution Protocol (LDP) and RSVP-TE are declining, and
the entire MPLS will gradually decline as its data plane is replaced by
IPv6 extensions due to the increasing popularity of SRv6.
Second, IP has the habit of expanding application scenarios. IP was
frst applied to the Internet. Later, IP-based MPLS was applied to IP
backbone, metro, and mobile transport networks. With the success
of SDN, IP is widely used in data centers to replace the traditional
Layer 2 networking. SRv6 is in the process of continuing this trend,
as it can well meet the requirements of new scenarios such as 5G and
cloud. To give another analogy of this trend, the development of ring
roads in Beijing is similar to the development of IP. As the city keeps
74.
SRv6 Background ◾27
expanding, new ring roads are built around the city, going from two
to three, and even to six rings. Construction of each ring road requires
the development and improvement of a new solution, along with the
reconstruction of areas that have already been encircled within the
ring roads. Just like SRv6 is used for new 5G and cloud scenarios,
MPLS will be replaced by SRv6 on existing IP transport networks.
Tese signifcant changes in network development act as a
reference for defning IP generations. Summarizing the history of
network development and defning IP generations help us better
seize opportunities for the future.
REFERENCES
[1] Postel J. Internet Protocol[EB/OL]. (2013-03-02)[2020-03-25]. RFC 791.
[2] Deering S, Hinden R. Internet Protocol Version 6 (IPv6) Specifcation[EB/
OL]. (2013-03-02)[2020-03-25]. RFC 2460.
[3] Rosen E, Viswanathan A, Callon R. Multiprotocol Label Switching
Architecture[EB/OL]. (2020-01-21)[2020-03-25]. RFC 3031.
[4] Casado M, Freedman M J, Pettit J, Luo J, Mckeown N, Shenker S. Ethane:
Taking Control of the Enterprise[EB/OL]. (2007-08-31)[2020-03-25]. ACM
SIGCOMM Computer Communication Review, 2007.
[5] Mckeown N, Anderson T, Balakrishnan H, Parulkar G, Peterson L,
Rexford J, Shenker S, Turner J. OpenFlow: Enabling Innovation in Campus
Networks[EB/OL]. (2008-03-31)[2020-03-25]. ACM SIGCOMM Computer
Communication Review, 2008.
[6] Filsfls C, Previdi S, Insberg L, Decraene B, Litkowski S, Shakir R. Segment
Routing Architecture[EB/OL]. (2018-12-19)[2020-03-25]. RFC 8402.
[7] Bashandy A, Filsfls C, Previdi S, Decraene B, Litkowski S, Shakir R.
Segment Routing with MPLS Data Plane[EB/OL]. (2019-12-06)[2020-03-
25]. draf-ietf-spring-segment-routing-mpls-22.
[8] Filsfls C, Dukes D, Previdi S, Leddy J, Matsushima S, Voyer D. IPv6 Segment
Routing Header (SRH)[EB/OL]. (2020-03-14)[2020-03-25]. RFC 8754.
[9] MahalingamM,DuttD,DudaK,AgarwalP,KreegerL,SridharT,BursellM,
Wright C. Virtual eXtensible Local Area Network (VXLAN): A Framework
for Overlaying Virtualized Layer 2 Networks over Layer 3 Networks[EB/
OL]. (2020-01-21)[2020-03-25]. RFC 7348.
[10] Réseaux IP Européens Network Coordination Centre. Te RIPE NCC has
run out of IPv4 Addresses[EB/OL]. (2019-11-25)[2020-03-25].
[11] Huang S, Liu J. Computer Network Tutorial Problem Solving and Experiment
Guide[M]. Beijing: Tsinghua University Press, 2006.
[12] RekhterY,DavieB,KatzD,RosenE,SwallowG.CiscoSystems’TagSwitching
Architecture - Overview[EB/OL]. (2013-03-02)[2020-03-25]. RFC 2105.
75.
28 ◾ SRv6Network Programming
[13] Pan P, Swallow G, Atlas A. Fast Reroute Extensions to RSVP-TE for LSP
Tunnels[EB/OL]. (2020-01-21)[2020-03-25]. RFC 4090.
[14] Awduche D, Berger L, Gan D, Li T, Srinivasan V, Swallow G. RSVP-TE:
Extensions to RSVP for LSP Tunnels[EB/OL]. (2020-01-21)[2020-03-25].
RFC 3209.
[15] Rosen E, Rekhter Y. BGP/MPLS IP Virtual Private Networks (VPNs)[EB/
OL]. (2020-01-21)[2020-03-25]. RFC 4364.
[16] Leymann N, Decraene B, Filsfls C, Konstantynowicz M, Steinberg D.
Seamless MPLS Architecture[EB/OL]. (2015-10-14)[2020-03-25]. draf-ietf-
mpls-seamless-mpls-07.
[17] Halpern J, Pignataro C. Service Function Chaining (SFC) Architecture[EB/
OL]. (2020-01-21)[2020-03-25]. RFC 7665.
[18] Brockners F, Bhandari S, Pignataro C, Gredler H, Leddy J, Youell S, Mizrahi
T, Mozes D, Lapukhov P, Chang R, Bernier D. Data Fields for In-situ
OAM[EB/OL]. (2020-03-09)[2020-03-25]. draf-ietf-ippm-ioam-data-09.
[19] Deering S, Hinden R. Internet Protocol Version 6 (IPv6) Specifcation[EB/
OL]. (2020-02-04)[2020-03-25]. RFC 8200.
[20] Yang Z, Li C. Network Reconstruction: SDN Architecture and Implemen-
tation[M]. Beijing: Electronic Industry Press, 2017.
[21] Song H. Protocol-Oblivious Forwarding: Unleash the Power of SDN
through a Future-Proof Forwarding Plane[EB/OL]. (2013-08-16)[2020-03-
25]. Proceedings of the Second ACM SIGCOMM Workshop on Hot Topics in
Sofware Defned networking, 2013.
[22] Bosshart P, Daly D, Gibb G, Izzard M, Mckeown N, Rexford J, Schlesinger C,
Talayco A, Vahdat A, Varghese G. P4: Programming Protocol-Independent
Packet Processors[EB/OL]. (2014-07-28)[2020-03-25]. ACM SIGCOMM
Computer Communication Review, 2014.
[23] Filsfls C, Camarillo P, Leddy J, Voyer D, Matsushima S, Li Z. SRv6 Network
Programming[EB/OL]. (2019-12-05)[2020-03-25]. draf-ietf-spring-srv6-
network-programming-05.
C H AP T E R 2
SRv6 Fundamentals
This chapter describes SRv6 fundamentals, including the basic
concepts, SRv6 extension header, instruction sets, packet forwarding
processes, and technical advantages of SRv6. Inheriting the advantages of
both IPv6 and source routing technology, SRv6 is more suitable for cross-
domain deployment than MPLS as it supports incremental evolution
and ensures better scalability, extensibility, and programmability. Tese
advantages mean that SRv6 is strategically important for future technical
evolution.
2.1 SRv6 OVERVIEW
SRv6 provides outstanding network programming capabilities.[1]
