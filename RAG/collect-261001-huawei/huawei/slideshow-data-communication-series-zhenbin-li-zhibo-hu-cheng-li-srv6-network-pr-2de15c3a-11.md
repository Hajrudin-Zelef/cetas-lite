---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-11
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["revenue"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1463, 1600]
sha256: 86432130d4fcc2812c0c13572c1fe91ff99f0a5f6790c23c6a60cb0b74759d4e
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

ing the services supported by IPv6, as opposed to IPv4. Driven by business
benefts, carriers will upgrade their networks to IPv6. But what drives the
need for IPv6 over IPv4? Te answer lies in the network programmability
of IPv6, which allows new services to be deployed quickly and easily to
create revenue for carriers.
Although IPv4 also provides the programmable Options feld, this feld
is not ofen used. IPv6, however, takes the extensibility of packet headers
into account from the very beginning. A number of extension headers,
including Hop-by-Hop Options, Destination Options, and Routing
headers,[19] were designed to support further extension.
However, afer 20-plus years, IPv6 extension headers are still rarely
applied to their full potential. With the rise of new services such as 5G
and cloud, and the development of network programming technologies,
services require the forwarding plane of a network to provide stronger
programming capabilities and a simpler converged network solution.
Tis is where SRv6 comes into play.
SRv6 is an SR network paradigm based on IPv6 data plane by making
use of a new IPv6 Routing Header, called Segment Routing Header (SRH),
allowing the ingress to insert forwarding instructions to guide data packet
forwarding. As shown in Figure 1.11, SRv6 combines the advantages of
69.
22 ◾ SRv6Network Programming
FIGURE 1.11 SR+IPv6=SRv6.
SR-MPLS’s programmability and IPv6 headers’ extensibility, giving IPv6
an edge.
As of now, SRv6 has been commercially deployed by multiple carri-
ers around the world just afer 2years since the draf SRv6 Network
Programming[23] was submitted to the IETF. Such rapid development is
uncommon among IP technologies. During these 2-plus years when we
promoted SRv6 innovation and standardization, we have communicated
extensively with industry experts and made a number of refections on the
experience and lessons learned in the development of Internet technolo-
gies. Tis has given us a further understanding of the value and signif-
cance of SRv6.
To sum up, the MPLS-based All IP 1.0 era has achieved great success
but has also brought some problems and challenges:
1. Isolated IP transport network islands: Although MPLS unifed the
technologies for transport networks, the IP backbone, metro, and
mobile transport networks are separated and need to be intercon-
nected using complex technologies such as inter-AS VPN, making
E2E service deployment difcult.
2. Limited programming space in IPv4 and MPLS encapsulation: Many
new services require more forwarding information to be added to
packets. However, the IETF announced that it has stopped formulat-
ing further standards for IPv4. In addition, the format of the MPLS
Label feld is fxed and lacks extensibility. Tese reasons make it dif-
fcult for IPv4 and MPLS to meet the requirements of new services
for network programming.
70.
SRv6 Background ◾23
3. Decoupling of applications and transport networks: Tis makes it
difcult to optimize networks and improve the value of networks.
Many carriers fnd themselves stuck as a provider of pipes and cannot
beneft from value-added applications. Moreover, the lack of applica-
tion information means that carriers can only implement network
adjustment and optimization in a coarse-granularity way, result-
ing in wasting resources. Attempts have been made over the years
to apply network technologies to user terminals, but all have failed.
An example of such attempts is ATM-to-desktop. Other attempts
have been made to deploy MPLS closer to hosts and applications, for
example, deploying MPLS for the cloud. But the fact is, deploying
MPLS in data centers is very difcult while VXLAN becomes the de
facto standard of data centers.
SRv6 technology is the answer to these problems.
1. SRv6 is compatible with IPv6 forwarding and can implement inter-
connection of diferent network domains easily through IPv6 reach-
ability. Unlike MPLS, SRv6 does not require additional signaling or
networkwide upgrades.
2. SRv6, based on SRHs, supports encapsulation of more information
into packets, meeting diversifed requirements of new services.
3. SRv6’s afnity to IPv6 enables it to seamlessly integrate IP transport
networks with IPv6-capable applications and provide more poten-
tial value-added services for carriers through application-aware
networks.
Te development of IPv6 over the past 20-plus years proves that the
demand for address space alone cannot promote the large-scale deploy-
ment of IPv6. But the rapid development of SRv6 indicates that IPv6
development can be boosted by requirements on new services. As shown
in Figure 1.12, along with the development of services such as 5G, cloud,
and Internet of Tings (IoT), the increasing number of network devices
require more addresses and network programmability. SRv6 can better
meet the requirements of these services, promote the development of
network services, and drive networks into a new All IP era, that is, an
71.
24 ◾ SRv6Network Programming
FIGURE 1.12 IP technology development generations.
intelligent IP era where all things are connected based on IPv6. In this
book, we refer to this era as the All IP 2.0 era.
1.6 STORIES BEHIND SRv6 DESIGN
I. SRv6 and SDN
Since its proposal in 2007, SDN has exerted a lasting impact on
industries. Te concept of SDN is also made more applicable over
diferent scenarios. From the originally revolutionary OpenFlow
and POF that require complete separation of forwarding and con-
trol, SDN has undergone gradual evolution amid a heated debate in
the industry. Te most important driving force behind this is the
IETF. From 2013 to 2016, an important task of the IETF was the
standardization of southbound protocols of SDN controllers, such as
BGP, Path Computation Element Communication Protocol (PCEP),
and Network Confguration Protocol (NETCONF)/Yet Another
Next Generation (YANG). Afer more than 4years of efort, the IETF
completed the main work of SDN transition in the control plane and
started to work on SRv6 Network Programming in 2017. Compared
with SR-MPLS, SRv6 has stronger network programming capabili-
ties. Programming, once implemented based on an OpenFlow- or
POF-capable forwarding plane, is now implemented through SRv6,
which provides better compatibility. Tis refects the SDN transition
in the forwarding plane. In other words, the SDN transition is still
ongoing, and the focus has shifed from the control plane to the for-
warding plane.
72.
SRv6 Background ◾25
II. Rethinking the Value and Signifcance of SR
While exploring this technology and discussing it with others, I
was ofen struck with some excellent thoughts, which were of great
help for me to understand the essence of technologies. Section 1.5 of
this chapter summarizes the problems faced by MPLS and the value
and signifcance of SRv6. Later I discussed with experts in the indus-
try and acquired more enlightening viewpoints which are shared in
the following for your reference:
1. MPLS is also essentially an extension of IP functions. However,
due to the limitations in the past, it was implemented by using
the Shim layer, which requires networkwide upgrades to sup-
port the extended functions. Afer more than 20years of sof-
ware and hardware development, many of these limitations
have been removed. SRv6 uses a new method to integrate IP and
MPLS functions (SRv6 SIDs refect both IP-like and MPLS-like
identifers), which better complies with the trend of technology
development.
2. SRv6 enables SDN on carrier IP networks. VXLAN is an impor-
tant foundation for the development of SDN on data center
networks. However, no VXLAN-like technology was available
to boost the development of SDN on carrier IP networks. SRv6
solves this problem.
During the development of SDN, there was a bias toward the
construction of SDN controller capabilities, but the impact of
network infrastructure was largely ignored. Carrier IP networks
are a lot more complex than data center networks, and this has
