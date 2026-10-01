---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-3
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei", "United States"]
dates: ["1998-12", "2000-05-10", "2011-02-01"]
keywords: ["cost", "research"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [379, 504]
sha256: fd02613773ff9175196d95d427e7ffa9cdc92249e9a5cf605e9b538496c80b55
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

From there on, the ability to implement specifc policies for diferent ser-
vices and data fows became more and more a necessity for most network
operators. Te era of SDN started. However, the way policies were imple-
mented using MPLS required state in the network and we arrived at a
scalability boundary.
Segment Routing (RFC8402) fnally brought a powerful way to intro-
duce source routing into modern networks running either MPLS or IPv6
data plane. Why suddenly has source routing become so popular? Because
of timing! Now is the time when you want to control data fows with-
out incurring heavy scale state machinery (that would not work anyway).
Source routing brings the state in the packet and leaves the control plane
of the network very light and very easy to manage.
Now the times are mature for the introduction of source routing but
revisited with a modern technology that copes with both MPLS and IPv6
data planes. In addition, Segment Routing applied to IPv6 data plane
(SRv6) allows the leveraging of the two main strong points of IPv6 archi-
tecture. First, SRv6 leverages the IPv6 128 bit address space that allows the
defnition of a node with enough bits so that an augmented semantic can
be defned (e.g., a node identity and a node location). Second, IPv6 defnes
the concept of Extension Headers (EH) which can be added to the packet
for diferent purposes. One of these purposes is source routing, and espe-
cially with the Segment Routing Header (SRH, RFC8754), you can now
specify the path the packet should follow.
Te innovation SRv6 brought to the industry has multiple aspects:
it allows the implementation of a straightforward, highly scalable, back-
ward compatible mechanism for source routing. Tis allows the imple-
mentation of large-scale policies addressing the variety of use cases
network operators have to address these days (5G, IoT, Connected Objects,
Virtualization, etc.). Another advantage is the ability to express not only
a path (that the packet should take) but also the process a packet should
go through during its journey. A typical example is the Service Chain use
case where a packet has to traverse service application elements before
reaching its destination.
SRv6 introduces the ability to combine forwarding decisions with
application/processing decisions, and this is called SRv6 Network
Programming. Finally, and this is the timing factor, we are able to lever-
age a mechanism that was invented many years ago (source routing, IPv6)
28.
Foreword I ◾xxiii
and combine it with modern capabilities, allowing to address current use
cases and requirements and taking into account the scale and perfor-
mance dimensions.
SR and SRv6 began in 2012 and 2013, and at that time, we were a very
small team. Rapidly, the industry understood the potential, and large col-
laboration between vendors and operators took place. Te standardization
of SR and SRv6 took place in IETF, and now we have a large community
participating in the SR and SRv6 efort.
Zhenbin Li and the Huawei team he leads are part of this community,
and their contribution to the evolution and standardization of SR tech-
nologies became substantial over the last years. Huawei SRv6 team has
been instrumental in the way SRv6 has been defned and standardized in
IETF, adopted by the industry and still extended to cope with increasing
and ever-emerging new requirements. Also, at the time of this writing,
Zhenbin Li is also a member of IAB.
Tis book gives an exhaustive view of the SRv6 technology including
the network programming capability, and I’m sure it will help the reader
to understand SRv6 technology, the use cases it applies to, and eventually,
to appreciate SRv6 technology as much as I have appreciated participating
in its invention.
Stefano Previdi
SR-MPLS and SRv6 Pioneer and Original Contributor
30.
xxv
Foreword II
The InternetProtocol version 6 (IPv6) was designed by the IPv6
Task Force within the IETF under the co-chairmanship of Steve
Deering and Robert Hinden and a small group of 40 engineers. Steve is
also the designer of Multicast. Te last draf standard of IPv6 (RFC 2460)
was released in December 1998, winning against IPv7, IPv8, and IPv9 pro-
posals. Tis release was basically the last efort of the IPv6 Task Force as
it had to close down its working group at the February 2–5, 1999, meeting
in Grenoble. In this meeting, I proposed the formation of the IPv6 Forum
to promote IPv6 to industry, ISPs/ MNOs, governments, academia, and
research ecosystem to start large-scale pilots and initial deployments.
Tere was a nice coincidence with the creation of 3GPP in December 1998
within ETSI in Sophia Antipolis. It was quite clear to the core IPv6 team
that 3G is the frst innovative driver that needs plenty of IP addresses. I
contacted Karl-Heinz Rosenbruck, Director General of ETSI and at that
time he was the chairman of 3GPP and proposed to him to adopt IPv6
for 3G instead of WAP. Afer 6 months of discussions with the various
3GPP WGs, 3GPP announced in May 10, 2000, the adoption of IPv6 for
3G. However, the 3G MNOs did not have the capacity building and the
skills to adopt IPv6 but preferred to use IPv4 and especially the Network
Address Translation (NAT) for 3G. Tis was in itself a great achievement
to get the wireless world to adopt the Internet for the telecom world as they
were at that time not that enthusiastic about the open Internet. Te client/
server model became the norm and the wireless Internet especially with
4G boomed in our hands and the rest is history. Te Internet has reached
5 billion users with a kind of economy class service.
Te public IPv4 address space managed by IANA (http://www.iana.org)
was completely depleted back in February 1, 2011. Tis creates by itself a
critical challenge when adding new IoT networks and enabling machine
31.
xxvi ◾ ForewordII
learning services on the Internet. Without publicly routable IP addressing,
the Internet of Tings, and anything that’s part of Machine Learning ser-
vices on the Internet, would be greatly reduced in its capabilities and then
limited in its potential success. Most discussions of IP over everything
have been based on the illusionary assumption that the IP address space is
an unlimited resource or it’s even taken for granted that IP is like oxygen
produced for free by nature.
Te introduction of IPv6 provides enhanced features that were not
tightly designed or scalable in IPv4 like IP mobility, end-to-end (e2e) con-
nectivity, ad hoc services, etc. IPv6 will be addressing the extreme sce-
narios where IP becomes a commodity service. Tis new address platform
will enable lower cost network deployment of large-scale sensor networks,
RFID, IP in the car, to any imaginable scenario where networking adds
value to commodity.
IPv6 deployment is now in full swing with some countries, such as
Belgium, achieving over 50% penetration. India has taken the lead by hav-
ing over 350 million IPv6 users. China has over 200 million IPv6 users
while the US has over 100 million users using IPv6 without the users even
knowing it.
Tere are many infections happening this decade to infuence the
design of the frst tangible IoT, 5G, and Smart Cities. It will take a com-
bination of IoT, SDN-NFV, Cloud Computing, Edge Computing, Big IoT
Data, and 5G, to sif through to realize the paradigm shif from current
research-based work to advanced IoT, 5G, and Smart Cities.
However, the move to NAT has basically killed the end-to-end model
that IPv6 restored so that applications and devices can peer directly with
each other. Te IPv6 deployment has continued its growth and reached
today over 1 billion IPv6 users. Tese users can be qualifed as business
class benefciaries. Te restoration of the end-to-end model has created
new opportunities for designing new end-to-end protocols such as IPv6-
based Segment (SRv6) which uses the prime and clean state design of IPv6
of end-to-end services. SRv6 is the frst big attempt to usher in the era of
