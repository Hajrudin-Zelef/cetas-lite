---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-8
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2019-11-25"]
keywords: ["cost", "revenue"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1051, 1193]
sha256: 862c06f06b55c094e2dc517f7d3bb33d65fd762568b4b77c2fd9dd95a012546a
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

which VXLAN is a typical example. At the same time, quite a few attempts
55.
8 ◾ SRv6Network Programming
were made to provide VPN services by introducing MPLS to data cen-
ters. However, these attempts all wound up in failure due to multiple fac-
tors, including numerous network boundaries, complex management, and
insufcient scalability.
In Figure 1.2, trafc from end users to the cloud data center needs to
travel multiple network domains. It passes through the MPLS-based Fixed
Mobile Convergence (FMC) transport network. Te trafc then enters
the MPLS-based IP backbone network through the native IP network,
accesses the data center’s IP network at the edge, and reaches the VXLAN
gateway. From there, it travels along the VXLAN tunnel, arrives at the
Top of Rack (TOR) switch at the egress of the VXLAN tunnel, and fnally
accesses the Virtual Network Function (VNF) device. We can therefore
envision how complex the service access process is due to an excessive
number of network domains.
Other major factors hindering the development of MPLS are scalability
and extensibility, which involves two aspects: scalability of the label space
and extensibility of encapsulation.
In the MPLS label space, as shown in Figure 1.3, there are 20 bits for the
MPLS label, which equates to a 220 label space.
FIGURE 1.2 Isolated MPLS network islands.
FIGURE 1.3 MPLS label encapsulation format.
56.
SRv6 Background ◾9
As the network scale expands, the label space is no longer sufcient.
Moreover, due to the limitation of the RSVP-TE protocol, the control
plane of MPLS networks also faces challenges such as complexity and a
lack of scalability. Tese make it difcult for MPLS to satisfy the require-
ments of network development.
On the other hand, the encoding of felds of MPLS encapsulation is
fxed. Although the MPLS label stack provides certain extensibility, as the
new network services develop and require more fexible encapsulation
in the forwarding plane (e.g., carrying metadata[17] in Service Function
Chaining (SFC) or In-situ Operations, Administration, and Maintenance
(IOAM)[18] packets), the extensibility of MPLS encapsulation faces more
challenges.
1.3.2 IPv4 Dilemma
One of the biggest problems regarding IPv4 is its insufcient address
resources.Sincethe1980s,IPv4addresseswereconsumedatanunexpectedly
fast pace. Put diferently, the Internet Assigned Numbers Authority (IANA)
announced that the last fve IPv4 address blocks were allocated on February
3, 2011. At 15:35 (UTC+1) on November 25, 2019, the fnal/22 IPv4 address
block was allocated in Europe, thereby signifying the depletion of global
IPv4 public addresses. Although technologies such as Network Address
Translation (NAT) help alleviate this issue by reusing private network
address blocks, this is by no means the ultimate solution.
As shown in Figure 1.4, NAT not only requires extra network confgu-
rations but also needs the maintenance of network state mappings, which
further complicates network deployment. On top of that, NAT does not
support source tracing of IPv4 addresses because the actual addresses are
hidden, and this generates management risks.
IPv4 also faces another dilemma: the insufcient extensibility of packet
headers results in inadequate programmability. Given this, it is difcult
for IPv4 networks to support many new services that require more exten-
sions of the header, such as source routing, SFC, and IOAM. Although
IPv4 defnes the Options feld for extension, it is rarely implemented and
used. Tis means that the insufcient extensibility of the IPv4 header will
constrain IPv4 development to a certain extent. Taking this into account,
the Internet Architecture Board (IAB) stated in 2016 that formulating
standards about new features based on IPv4 will not be considered in the
future.
57.
10 ◾ SRv6Network Programming
FIGURE 1.4 NAT.
To fnd a solution against IPv4 address exhaustion and poor extensi-
bility, the industry designed the next-generation upgrade solution for
IPv4 — IPv6.[2]
1.3.3 Challenges for IPv6
As the next-generation IP protocol, IPv6 aims to solve two major prob-
lems of IPv4: limited address space and insufcient extensibility.[19] To this
end, IPv6 ofers certain improvements to IPv4.
One such improvement is the expanded address space, or stated dif-
ferently, the increase from 32 bits for IPv4 addresses to 128 bits for IPv6
addresses. Te magnitude of this expansion can be compared to allocating
an IPv6 address to every grain of sand on the Earth, efectively solving the
problem of insufcient IPv4 addresses.
Te extension header mechanism counts as another noteworthy
enhancement. Requirement For Comments (RFC) 8200[19] defnes the fol-
lowing IPv6 extension headers, which ideally should appear in a packet in
the following order (Appendix A expands on the details of IPv6):
1. IPv6 header
2. Hop-by-Hop Options header
58.
SRv6 Background ◾11
3. Destination Options header
4. Routing header
5. Fragment header
6. Authentication header
7. Encapsulating Security Payload header
8. Destination Options header
9. Upper-Layer header
Figure 1.5 shows the encapsulation format of a common IPv6 extension
header that carries a TCP message.
As far as IPv6 is concerned, extension headers provide good extensibil-
ity and programmability. For example, the Hop-by-Hop Options header
can be used to implement hop-by-hop IPv6 data processing, and the
Routing header to implement source routing.
Over 20years have passed since IPv6 was proposed, yet IPv6 develop-
ment is still quite slow. Only within the last few years have technology
development and policies propelled IPv6 deployment. In retrospect, the
slow IPv6 development can mainly be attributed to the following factors:
1. Incompatibility with IPv4 and high costs for network upgrades:
Although IPv6 ofers an address space of 128 bits compared to
32 bits in IPv4, it is incompatible with IPv4, meaning that hosts
using IPv6 addresses cannot directly communicate with those using
IPv4 addresses. As such, a transition solution is required, resulting
in high network upgrade costs.
2. Insufcient service driving force and low network upgrade ben-
efts: IPv6 advocates have been promoting the 128-bit address space,
FIGURE 1.5 Encapsulation format of an IPv6 extension header carrying a TCP
message.
59.
12 ◾ SRv6Network Programming
which is viewed as a solution to IPv4 address exhaustion. However,
technologies such as NAT can also solve this problem, and NAT does
indeed serve as the main remedy. Te specifc process involves lever-
aging private network addresses and address translation technolo-
gies to temporarily alleviate the problem which may hinder network
service development.
A key advantage of NAT is the fact that it can be deployed at a lower cost
than that needed to upgrade IPv4 networks to IPv6. Considering the
fact that existing services run well on IPv4 networks, there is no point in
upgrading them, especially as this does not yield new revenue and, on the
contrary, would lead to higher costs. Tis is the major reason why carriers
are unwilling to upgrade.
In light of this, the key to solving slow IPv6 deployment lies in fnding
more attractive services supported by IPv6, as opposed to IPv4. Tis way,
business benefts can drive carriers to upgrade their networks to IPv6.
1.4 OPPORTUNITIES FOR ALL IP 1.0: SDN
AND NETWORK PROGRAMMING
In addition to the insufcient extensibility and programmability of
IPv4 and MPLS data planes, the All IP 1.0 era also faces the following
challenges[20]:
• Lack of a global network view and trafc visualization capability.
As a consequence of this, it is difcult to make optimal decisions
from a global perspective of the network, or quickly respond to TE
requirements.
• Lack of a unifed abstract model in the data plane prevents the con-
trol plane from supporting new functions by programming with
data-plane Application Programming Interfaces (APIs).
• Lack of automation tools and a long service rollout period.
