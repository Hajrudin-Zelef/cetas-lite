---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-1
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1, 160]
sha256: 5241be02660507504815d7118f9465872fafc89cf11700dc018a3a46d765fd98
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

(Data Communication Series) Zhenbin Li, Zhibo Hu, Cheng Li - SRv6 Network Programming_ Ushering in a New .pdf
Ouvre dans une nouvelle fenêtreOuvre un site Web externeOuvre un site Web externe dans une nouvelle fenêtre
Ce site Web utilise des technologies telles que les cookies pour activer les fonctionnalités essentielles du site, ainsi que pour analyses, personnalisation et publicité ciblée. Pour en savoir plus, consultez le lien suivant : Politique de confidentialité
Préférences en matière de conservation des données
(Data Communication Series) Zhenbin Li, Zhibo Hu, Cheng Li - SRv6 Network Programming_ Ushering in a New .pdf
Description améliorée par l'IA
This document provides advance praise for the book "SRv6 Network Programming: Ushering in a New Era of IP Networks" from several experts in the field. They praise the book for providing a comprehensive overview of SRv6 including its background, fundamentals, benefits, and applications. They believe the book will help promote the deployment and development of SRv6.
(Data Communication Series) Zhenbin Li, Zhibo Hu, Cheng Li - SRv6 Network Programming_ Ushering in a New .pdf
2.
Advance Praise
“Tis bookprovides a comprehensive overview of SRv6, including its back-
ground, fundamentals, benefts, and applications. I frmly believe that the
publication of this book will further promote the large-scale deployment
and development of IPv6.”
— Hequan Wu, Academician of Chinese Academy of Engineering
“Te book SRv6 Network Programming: Ushering in a New Era of IP
Networks, written by Huawei’s Zhenbin Li and his team, provides a com-
plete description of SRv6 innovations and standards, and systematically
summarizes their experience in SRv6 R&D. It gives a holistic view into
SRv6. I sincerely hope that the publication of this book can positively
promote the development of IPv6 core technologies in China and make a
strong impact on the world.”
— Xing Li, Professor of Department of Electronic Engineering,
Tsinghua University
“Tis book takes a much-needed step-by-step approach to explaining
the SR technologies that apply to IPv6 networks. In doing so, it clarifes
and builds on the standardization work of the IETF’s SPRING working
group that has successfully brought together all of the large networking
equipment vendors with a number of future-looking network operators to
‘make the Internet work better.’ Te future of the Internet is notoriously
hard to predict, but this book will assist readers to embrace its potential.”
— Adrian Farrel, Former IETF Routing Area Director
“In the past 20+ years, MPLS played an important role in the development
of IP networks. As the development of 5G and cloud progresses, SRv6 is
3.
winning much attentionin the IP industry as the promising technology.
In the process, IP experts from Huawei make great contributions to the
innovation and standardization work in IETF. Te book written by the
team details the principles of SRv6 and corresponding applications for 5G
and cloud according to years of extensive experience. I believe that it can
be of much help to readers to master SRv6 and will facilitate the develop-
ment of SRv6 technology and industry.”
— Loa Andersson, MPLS WG Chair of IETF
“In this, Zhenbin Li and his team walk the reader through the history
leading to SRv6 and introduce the basic concepts as well as more advanced
ideas for the technology. Contributions for this book are written by
technologists who have been actively involved in the development and
standardization of SRv6 in the IETF from the beginning and have person-
ally helped to evolve not only the base technology but also many more
advanced features that are critical for successful service creation to meet
the demands placed upon today’s networks. Te reader will gain much
insight into how SRv6 has and continues to evolve, and in addition will
read about many new concepts that are actively being worked on in the
standards bodies.”
— James N. Guichard, SPRING WG Chair of IETF
“Tis book is a comprehensive introduction and tutorial to SRv6. It com-
prehensively covers all aspects from background to SRv6 principles and
basics, and to services built upon SRv6, such as trafc engineering, reac-
tions to network failures, VPN, and multicast. It concludes with a preview
of future evolution directions such as 5G and SRv6 header compression.
In addition to the technology, it includes personal refections on the back-
ground and reasons for the technical choices.”
— Bruno Decraene, Orange expert and senior network architect at
Orange/SPRING WG Co-chair of IETF
“SRv6 is maturing, and the family of standards defning all aspects of
the solution keeps growing — greatly supported by the authors’ endur-
ing eforts in the IETF. Our lab and other organizations have successfully
completed multi-vendor interoperability tests in the past three years, with
a growing number of participants confrming that the technology becomes
adopted. Te authors systematically and comprehensively describe the
4.
evolution and applicationof SRv6 as well as the benefts that it brings to
the industry. I truly believe that this book provides excellent value in guid-
ing SRv6 application and deployment.”
— Carsten Rossenhovel, Managing Director of EANTC (European
Advanced Networking Test Center)
Data Communication Series
CloudData Center Network Architectures and Technologies
Lei Zhang and Le Chen
Campus Network Architectures and Technologies
Ningguo Shen, Bin Yu, Mingxiang Huang, and Hailin Xu
Enterprise Wireless Local Area Network Architectures and Technologies
Rihai Wu, Xun Yang, Xia Zhou, and Yibo Wang
Sofware-Defned Wide Area Network Architectures and Technologies
Cheng Sheng, Jie Bai, and Qi Sun
SRv6 Network Programming
Ushering in a New Era of IP Networks
Zhenbin Li, Zhibo Hu, and Cheng Li
For more information on this series, please visit: https://www.routledge.com/
Data-Communication-Series/book-series/DCSHW
v
Contents
Foreword I, xxi
ForewordII, xxv
Preface, xxix
Teams, xxxv
Acknowledgments, xxxix
Authors, xli
PART I Introduction
CHAPTER 1 ◾ SRV6 Background 3
1.1 OVERVIEW OF INTERNET DEVELOPMENT 3
1.2 START OF ALL IP 1.0: A COMPLETE VICTORY FOR IP 4
1.2.1 Competition between ATM and IP 4
1.2.2 MPLS: Te Key to All IP 1.0 5
1.3 CHALLENGES FACING ALL IP 1.0: IP/MPLS DILEMMA 7
1.3.1 MPLS Dilemma 7
1.3.2 IPv4 Dilemma 9
1.3.3 Challenges for IPv6 10
1.4 OPPORTUNITIES FOR ALL IP 1.0: SDN AND
NETWORK PROGRAMMING 12
1.4.1 OpenFlow 14
1.4.2 POF 16
Contents ◾ xi
5.3.1Principles of L3VPN over SRv6 BE 170
5.3.1.1 Workfow of L3VPN over SRv6 BE
in the Control Plane 170
5.3.1.2 Workfow of L3VPN over SRv6 BE
in the Forwarding Plane 171
5.3.2 Principles of L3VPN over SRv6 TE 172
5.3.2.1 Workfow of L3VPN over SRv6 Policy
in the Control Plane 173
5.3.2.2 Workfow of L3VPN over SRv6 Policy
in the Forwarding Plane 175
5.4 SRv6 EVPN 176
5.4.1 Principles of EVPN E-LAN over SRv6 179
5.4.1.1 MAC Address Learning and
Unicast Forwarding 179
5.4.1.2 Replication List Establishment and
BUM Trafc Forwarding 182
5.4.2 Principles of EVPN E-Line over SRv6 186
5.4.3 Principles of EVPN L3VPN over SRv6 188
5.4.4 SRv6 EVPN Protocol Extensions 189
5.4.4.1 Ethernet A-D Route 189
5.4.4.2 MAC/IP Advertisement Route 191
5.4.4.3 IMET Route 193
5.4.4.4 ES Route 195
5.4.4.5 IP Prefx Route 195
5.5 STORIES BEHIND SRv6 DESIGN 197
REFERENCES 198
CHAPTER 6 ◾ SRv6 Reliability 201
6.1 IP FRR AND E2E PROTECTION 201
6.1.1 TI-LFA Protection 202
6.1.1.1 LFA 202
6.1.1.2 RLFA 204
17.
xii ◾ Contents
6.1.1.3TI-LFA 209
6.1.2 SRv6 Midpoint Protection 213
6.1.3 Egress Protection 217
6.1.3.1 Anycast FRR 218
6.1.3.2 Mirror Protection 219
6.2 MICROLOOP AVOIDANCE 223
6.2.1 Microloop Cause 223
6.2.2 SRv6 Local Microloop Avoidance in a Trafc
Switchover Scenario 225
6.2.3 SRv6 Microloop Avoidance in a Trafc Switchback
Scenario 227
6.2.4 SRv6 Remote Microloop Avoidance in a Trafc
Switchover Scenario 231
6.3 STORIES BEHIND SRv6 DESIGN 233
