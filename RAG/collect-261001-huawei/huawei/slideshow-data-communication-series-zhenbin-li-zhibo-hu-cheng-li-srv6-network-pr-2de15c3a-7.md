---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-7
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2019-11-25"]
keywords: ["cost", "ethernet", "revenue"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [923, 1050]
sha256: 68aee6bc42b8f4793821ce80d05a43963a1b4dbe8dcf56199644e80bb4f5d677
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

In 2014, Virtual eXtensible Local Area Network (VXLAN)[9] was
released.
On November 25, 2019, IPv4 addresses were exhausted.[10]
1.2 START OF ALL IP 1.0: A COMPLETE VICTORY FOR IP
1.2.1 Competition between ATM and IP
In the initial stage of network development, multiple types of networks,
such as X.25, Frame Relay (FR), ATM, and IP, coexisted to meet difer-
ent service requirements. Tese networks could not interwork with each
other, and also competed, with mainly ATM and IP networks taking cen-
ter stage.
ATM is a transmission mode that uses fxed-length cell switching. It
establishes paths in connection-oriented mode and can provide bet-
ter Quality of Service (QoS) capabilities than IP. Its design philosophy
involves centering on networks and providing reliable transmission, and
its design concepts refect the reliability and manageability requirements
52.
SRv6 Background ◾5
of telecommunications networks. Tis is the reason why ATM was widely
deployed on early telecommunications networks.
Te design concepts of IP difer greatly from those of ATM. To be more
precise, IP is a connectionless communication mechanism that provides
the best-efort forwarding capability, and the packet length is not fxed.
On top of that, IP networks mainly rely on the transport-layer protocols
(e.g., TCP) to ensure transmission reliability, and the requirement for the
network layer involves ease of use. To add on to this, the design concept of
IP networks embodies the “terminal-centric and best-efort” notion of the
computer network. We can therefore say that IP is widely used on com-
puter networks because it meets the corresponding service requirements.
Te competition between ATM and IP networks can essentially be
represented as a competition between telecommunications and computer
networks. In other words, telecommunications practitioners sought to use
ATM for network interconnection to protect network investments. On the
fip side, computer practitioners aimed at using ATM as only a link-layer
technology to provide QoS guarantee for IP networks, while setting aside
the task of establishing network connections for IP.
Computer networks subsequently evolved toward broadband, intelli-
gence, and integration, with mainly burst services. Despite this, the QoS
requirements that trafc places on computer networks are not as high
as those on telecommunications networks, and the length of packets is
not fxed. As such, the advantages of ATM — fxed-length cell switching
and good QoS capabilities — cannot be brought into full play on com-
puter networks. Not only that, the QoS capabilities of ATM are based
on connection-oriented control with a certain packet header overhead.
Terefore, ATM is inefcient in carrying computer network trafc, and it
yields high transmission and switching costs.
To sum up, as network scale expanded and network services increased
in number, ATM networks became more complex than IP networks, while
also bearing higher management costs. Within the context of costs versus
benefts, ATM networks exited the arena as they were gradually replaced
by IP networks.
1.2.2 MPLS: The Key to All IP 1.0
Although with relation to the development of computer networks, the
IP network is more ftting than the ATM network, a certain level of QoS
guarantee is still required. To compensate for the IP network’s insufcient
53.
6 ◾ SRv6Network Programming
QoS capabilities, numerous technologies integrating IP and ATM, such
as Local Area Network Emulation (LANE), IP over ATM (IPoA),[11] and
tag switching,[12] have been proposed. However, these technologies only
addressed part of the issue, until 1996 when MPLS technology was pro-
posed[3] to provide a better solution to this issue.
MPLS is considered as a Layer 2.5 technology that runs between Layer
2 and Layer 3. It supports multiple network-layer protocols, such as IPv4
and IPv6, and is compatible with multiple link-layer technologies, such
as ATM and Ethernet. Some of its other highlights include the fact that
it incorporates ATM’s Virtual Channel Identifer (VCI) and Virtual Path
Identifer (VPI) switching concepts, combines the fexibility of IP routing
and simplicity of label switching, and adds connection-oriented attributes
to connectionless IP networks. By establishing virtual connections, MPLS
provides better QoS capabilities for IP networks.
However, this is not the only reason why it was initially proposed. Point
in case being that MPLS also forwards data based on the switching of fxed-
length 32-bit labels, and therefore it features a higher forwarding efciency
than IP, which forwards data based on the Longest Prefx Match (LPM).
Tat said, as hardware capabilities have and continue to improve, MPLS no
longer features distinct advantages in forwarding efciency. Nevertheless,
MPLS provides a good QoS guarantee for IP through connection-oriented
label forwarding and also supports Trafc Engineering (TE), Virtual
Private Network (VPN), and Fast Reroute (FRR).[13] Tese advantages play
a key role in the continuous expansion of IP networks, while also catapult-
ing the IP transformation of telecom networks.
In general, the success of MPLS depends mainly on its three important
features: TE, VPN, and FRR.
• TE: Based on Resource Reservation Protocol-Trafc Engineering
(RSVP-TE),[14] MPLS labels can be allocated and distributed along
the MPLS TE path, and TE features (such as resource guarantee and
explicit path forwarding) can be implemented. Tis overcomes IP
networks’ lack of support for TE.
• VPN: MPLS labels can be used to identify VPNs[15] for isolation
of VPN services. As one of the major application scenarios of MPLS,
VPN is a key technology for enterprise interconnection and multi-
service transport as well as an important revenue source for carriers.
54.
SRv6 Background ◾7
• FRR: Te IP network cannot provide complete FRR protection,
which in turn means that it is unable to meet the high-reliability
requirements of carrier-grade services. MPLS improves the FRR
capabilities of IP networks and supports 50 ms carrier-grade protec-
tion switching in most failure scenarios.
Because IP networks are cost-efective and MPLS provides good TE, VPN,
and FRR capabilities, IP/MPLS networks gradually replaced dedicated
networks, such as ATM, FR, and X.25. Ultimately, MPLS was applied to
various networks, including IP backbone, metro, and mobile transport, to
support multiservice transport and implement the Internet’s All IP trans-
formation. In this book, we refer to the IP/MPLS multiservice transport
era as the All IP 1.0 era.
1.3 CHALLENGES FACING ALL IP 1.0: IP/MPLS DILEMMA
Although IP/MPLS drove networks into the All IP 1.0 era, the IPv4 and
MPLS combination has also set forth numerous challenges, which are
becoming more prominent as network scale expands and cloud services
develop, and are thereby hindering the further development of networks.
1.3.1 MPLS Dilemma
From one perspective, MPLS plays an important role in All IP transport,
while from another perspective, it complicates inter-domain network
interconnection by causing isolated network islands.
To put it more precisely, consider the fact that on the one hand, MPLS
is deployed in diferent network domains, such as IP backbone, metro,
and mobile transport networks, forming independent MPLS domains
and creating new network boundaries. However, many services require
E2E deployment, and this means that services need to be deployed across
multiple MPLS domains, which in turn results in complex inter-domain
MPLS solutions. In that regard, multiple inter-Autonomous System (AS)
solutions, such as Option A, Option B, and Option C,[15,16] have been pro-
posed for inter-AS MPLS VPN, and each one involves relatively complex
service deployment.
On the other hand, as the Internet and cloud computing develop, more
and more cloud data centers are built. To meet the requirements of multi-
tenant networking, multiple overlay technologies were proposed, among
