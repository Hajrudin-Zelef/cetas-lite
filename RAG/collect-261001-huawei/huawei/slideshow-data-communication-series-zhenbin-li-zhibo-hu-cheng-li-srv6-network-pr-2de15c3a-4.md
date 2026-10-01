---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-4
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: ["distribution", "ethernet", "latency", "research"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [505, 644]
sha256: 3fae6ba4ad1e149587def130bc165295046887ac6e4d96585562acb4367f0d53
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

IPv6 innovation, and boasts the following two key characteristics:
• SRv6 is an IPv6-native network technology designed to simplify
the transport network. In this era where IPv4, IPv6, MPLS, and
SR-MPLS coexist, SRv6 will achieve optimal simplifcation at the
transport layer under the minimalism and centralization guiding
32.
Foreword II ◾xxvii
principles. Te real benefts of SRv6 include the support of VPN,
FRR, TE, network slicing, and BIERv6, enabling application-driven
path programming, providing diferentiated SLA assurance for
users, and helping carriers transform from “ofering bandwidth” to
“ofering services.”
• SRv6 is a rare project that the IETF has focused on over the past
decade. It provides a programmable network architecture for 5G and
cloud transport, while also heralding the era of IPv6+ network sys-
tem innovation. Datacom industry players, such as Huawei, Cisco,
China Telecom, China Mobile, China Unicom, and SofBank, have
invested a tremendous amount of efort into jointly promoting SRv6
maturity, and the SRv6 transport solution is being applied to an
increasing number of networks.
Huawei is the frst in the industry to launch SRv6-capable data communi-
cation products. In the initial phase of SRv6’s history, the company actively
participated in innovation and standardization, and achieved a plethora
of signifcant milestones. To add to that, Zhenbin Li was elected as an
IETF Internet Architecture Board (IAB) member in early 2019. Tis book
is a summary of the SRv6 research conducted by him and his team, as well
as their corresponding experiences. Te book systematically interprets the
innovative technologies and standards of SRv6, especially with regard to
SRv6 implementation practices, which is not only inspiring and helpful
but also benefcial to the continuous promotion of IPv6 innovation.
Latif Ladid
Founder and President, IPv6 Forum; Member of 3GPP PCG (MRP),
Co-chair, IEEE FNI 5G, World Forum/5G Summits.
Chair, ETSI IPv6 Integration ISG.
34.
xxix
Preface
Segment Routing overIPv6 (SRv6) is an emerging IP technology.
Te development of 5G and cloud services creates many new require-
ments on network service deployment and automated O&M. SRv6 pro-
vides comprehensive network programming capabilities to better meet
the requirements of new network services and is compatible with IPv6 to
simplify network service deployment.
IP transport networks are built around connections. 5G changes the
attributes of connections, and cloud changes their scope. Tese changes
bring big opportunities for SRv6 development. Te development of 5G ser-
vices poses higher requirements on network connections, such as stronger
Service Level Agreement (SLA) guarantee and deterministic latency. As
such, packets need to carry more information. Tese requirements can
be well met through SRv6 extensions. Te development of cloud services
makes service processing locations more fexible. Some cloud services
(such as telco cloud) further break down the boundary between physical
and virtual network devices, integrating services and transport networks.
All of this changes the scope of network connections. Te unifed pro-
gramming capability of SRv6 in service and transport, as well as the native
IP attribute, enables the rapid establishment of connections and meets the
requirements for fexible adjustment of the connection scope. As men-
tioned in this book, SRv6 network programming has ushered in a new
network era, redefning the development of IP technologies signifcantly.
Technical experts from Huawei Data Communication Product Line
have been conducting long-term and in-depth research and development
in IP and SRv6 felds, and contributing to numerous SRv6 standards in
the IETF. In addition, many experts have participated in SRv6 network
deployments, accumulating extensive experience in the feld. We compiled
this book based on comprehensive research, development, and network
35.
xxx ◾ Preface
O&Mexperience to provide a complete overview of SRv6 in the hope of
better understanding its fundamentals and the new network technologies
that derive from it. We also hope it will spark participation in the research,
application, and deployment of these new technologies, promoting the
development of communications networks.
OVERVIEW
Tis book begins with the challenges services face regarding IP technol-
ogy development. It describes the background and mission of SRv6 and
presents a comprehensive overview of the technical principles, service
applications, planning and design, network deployment, and industry
development of next-generation IP networks. Tis book consists of 13
chapters and is divided into four parts:
Part I: Introduction
Chapter 1 focuses on the development of IP technologies and reveals
why SRv6 is developing so quickly.
Part II: SRv6 1.0
Chapters 2 through 8 introduce SRv6 1.0, covering basic SRv6 capa-
bilities and demonstrating how SRv6 supports existing services in a
simple and efcient manner.
Part III: SRv6 2.0
Chapters 9 through 12 introduce SRv6 2.0, covering new SRv6-based
network technologies targeted at 5G and cloud services and dem-
onstrating service model innovation brought by SRv6 network
programming.
Part IV: Summary and Future Developments
Chapter 13 summarizes the development of the SRv6 industry and
forecasts developmental trends from SRv6 to IPv6+, which stretches
people’s imagination about future networks.
CHAPTER 1: SRv6 BACKGROUND
Tis chapter comprehensively describes the development of IP technolo-
gies, introduces SRv6 based on historical experience and requirements for
36.
Preface ◾ xxxi
IPtechnology development, and summarizes the value and signifcance of
SRv6 from a macro perspective.
CHAPTER 2: SRv6 FUNDAMENTALS
Tis chapter describes SRv6 fundamentals to demonstrate how SRv6 is used
to implement network programming, and summarizes the advantages of
SRv6 network programming and its mission from a micro perspective.
CHAPTER 3: BASIC PROTOCOLS FOR SRv6
Tis chapter describes the fundamentals of and extensions to Intermediate
System to Intermediate System (IS-IS) and Open Shortest Path First version 3
(OSPFv3) for SRv6. Protocols such as Resource Reservation Protocol-Trafc
Engineering (RSVP-TE) and Label Distribution Protocol (LDP) are not
required on an SRv6 network, thereby simplifying the network control plane.
CHAPTER 4: SRv6 TE
Tis chapter describes the fundamentals and extensions of SRv6 Trafc
Engineering (TE), which is a basic feature of SRv6. SRv6 Policy is the main
mechanism for implementing SRv6 trafc engineering and draws on the
source routing mechanism of Segment Routing to encapsulate an ordered
list of instructions on the headend, guiding packets through the network.
CHAPTER 5: SRv6 VPN
Tis chapter describes the principles and protocol extensions of Virtual
Private Network (VPN), a basic SRv6 feature. SRv6 supports existing
Layer 2 VPN (L2VPN), Layer 3 VPN (L3VPN), and Ethernet VPN (EVPN)
services, which can be deployed as long as edge nodes are upgraded to
support SRv6, shortening the VPN service provisioning time.
CHAPTER 6: SRv6 RELIABILITY
Tis chapter describes the fundamentals and protocol extensions of SRv6
reliability technologies, including FRR, TI-LFA, midpoint protection,
egress protection, and microloop avoidance. Tese technologies ensure
E2E local protection switching within 50 ms on an SRv6 network.
CHAPTER 7: SRv6 NETWORK EVOLUTION
Tis chapter describes the challenges and technical solutions of SRv6 net-
work evolution, that is, how to evolve an existing IP/MPLS network to
37.
xxxii ◾ Preface
anSRv6 network. SRv6 supports incremental deployment, which protects
existing investments and allows for new services.
CHAPTER 8: SRv6 NETWORK DEPLOYMENT
Tis chapter describes SRv6 network deployment, including SRv6 applica-
tion scenarios and how to design and confgure features such as TE, VPN,
and reliability. Network deployment practices show that SRv6 is advanta-
