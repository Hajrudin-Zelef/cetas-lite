---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-5
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["reasoning", "research"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [645, 791]
sha256: 72e322884426e33d60dbb8a1ea67997155a1de78dc20894f7449807d572027a0
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

geous over traditional solutions in terms of inter-AS, scalability, protocol
simplifcation, and incremental deployment.
CHAPTER 9: SRv6 OAM AND ON-PATH
NETWORK TELEMETRY
Tis chapter describes the fundamentals and protocol extensions of SRv6
OAM and on-path network telemetry. SRv6 OAM and on-path network
telemetry provide a solid basis for SRv6 network quality guarantee, allow-
ing for extensive use of SRv6 on carrier networks.
CHAPTER 10: SRv6 FOR 5G
Tis chapter describes new SRv6 technologies applicable to 5G services,
including VPN+ network slicing, Deterministic Networking (DetNet),
and SRv6-to-mobile core solutions and relevant protocol extensions.
CHAPTER 11: SRv6 FOR CLOUD SERVICES
Tis chapter describes the application of SRv6 in cloud services, including
the concept, challenges, and SRv6-based solution for telco cloud as well as
the fundamentals and protocol extensions of SRv6 for SFC and SRv6 for
SD-WAN.
CHAPTER 12: SRv6 MULTICAST/BIERv6
Tis chapter describes SRv6-based multicast technologies, focusing on the
fundamentals and protocol extensions of Bit Index Explicit Replication
(BIER) and Bit Index Explicit Replication IPv6 Encapsulation (BIERv6).
Working with SRv6, BIERv6 allows both unicast and multicast services to
be transmitted on a unifed IPv6 data plane based on an explicit forward-
ing path programmed on the ingress.
CHAPTER 13: SRv6 INDUSTRY AND FUTURE
Tis chapter summarizes the development of the SRv6 industry and
provides forecasts for its future, focusing on SRv6 extension header
38.
Preface ◾ xxxiii
compression,Application-aware IPv6 Networking (APN6), and the three
phases that may exist during its development from SRv6 to IPv6+.
As SRv6 involves a signifcant number of IPv6 basics, this book
expands on them in Appendix A to help better understand SRv6. In addi-
tion, this book describes IS-IS and OSPFv3 TLVs in Appendices B and
C, respectively, as SRv6 is implemented based on these IS-IS and OSPFv3
extensions. Tis will help to demonstrate how SRv6 is implemented as
well as the relationships between IS-IS and SRv6, and between OSPFv3
and SRv6.
In the Postface “SRv6 Path,” Zhenbin Li draws on his personal
experience to summarize the development history of SRv6 and Huawei’s
participation in promoting SRv6 innovation and standardization. In addi-
tion, “Stories behind SRv6 Design” is provided at the end of each chapter,
touching on the experience and philosophy behind the protocol design,
interpreting its technical nature, or describing the experience and lessons
learned during the development of technologies and standards. Trough
this content, it is hoped that readers can obtain a greater understanding of
SRv6 and the reasoning behind its design principles. Some of the content
constitutes the author’s opinion and should be used for reference only.
Te editorial board for this book gathered the technical experts from
Huawei’s data communication research team, standards and patents team,
protocol development team, solution team, technical documentation
team, and translation team. Team members include the developers and
promoters of SRv6 standards, R&D staf on SRv6 design and implemen-
tation, and solution experts who helped customers successfully complete
SRv6 network design and deployment. Teir achievements and experi-
ence are systematically summed up in this book. Information department
staf, including Lanjun Luo, edited the text and worked on the fgures.
Te publication of this book is the collaborative result of team efort and
collective wisdom. Heartfelt thanks go to every member of the editorial
board. Tough the process is hard, we really enjoyed the teamwork in
which everyone learned so much from each other.
40.
xxxv
Teams
TECHNICAL COMMITTEE
Chair
Kewen Hu(Kevin Hu), President of Huawei Data Communication Product
Line
Members
Shaowei Liu, Director of Huawei Data Communication R&D Mgmt Dept
Zhipeng Zhao, Director of Huawei Data Communication Marketing Dept
Chenxi Wang, Director of Huawei Data Communication Strategy &
Business Development Dept
Wei Hu, Director of Huawei Carrier IP Marketing & Solution Sales Dept
Jinzhu Chen, President of Service Router Domain, Huawei Data
Communication Product Line
Meng Zuo, President of Core Router Domain, Huawei Data
Communication Product Line
Banghua Chen, Director of Huawei Data Communication Marketing
Execution Dept
Suning Ye, Director of Router PDU, Huawei Data Communication
Product Line
Xiao Qian, Director of Huawei Data Communication Research Dept
Jianbing Wang, Director of Huawei Data Communication Architecture &
Design Dept
41.
xxxvi ◾ Teams
ZhaokunDing, Director of Huawei Data Communication Protocol
Development Dept
Minwei Jin, Director of IP Technology Research Dept, NW, Huawei Data
Communication Product Line
Dawei Fan, Director of Huawei Data Communication Standard & Patent
Dept
Jianping Sun, Director of Huawei Data Communication Solutions Dept
Mingzhen Xie, Director of Huawei Information Digitalization and
Experience Assurance (IDEA) Dept, DC
EDITORIAL BOARD
Editor-in-Chief: Zhenbin Li (Robin Li)
Deputy Editors-in-Chief: Zhibo Hu, Cheng Li
Members: Lanjun Luo, Ting Liao, Shunwan Zhuang, Haibo Wang, Huizhi
Wen, Yaqun Xiao, Guoyi Chen, Tianran Zhou, Jie Dong, Xuesong Geng,
Shuping Peng, Lei Li, Jingrong Xie, Jianwei Mao
Technical Reviewers: Rui Gu, Gang Yan
Translators: Yanqing Zhao, Wei Li, Xue Zhou, Yufu Yang, Zhaodi Zhang,
Huan Liu, Yadong Deng, Chen Jiang, Ruijuan Li, Junjie Guan, Jun Peng,
Samuel Luke Winfeld-D'Arcy, George Fahy, Rene Okech, Evan Reeves
TECHNICAL REVIEWERS
Rui Gu: Chief Solution Architect of Huawei Data Communication
Product Line. Mr. Gu joined Huawei in 2007 and now leads the Data
Communication Solutions Design Dept. He has extensive experience
working in the Versatile Routing Platform (VRP) department and has
conducted in-depth research on IP/MPLS protocols. He once led the R&D
team of Huawei’s next-generation backbone routers and also has a wealth
of experience in implementing end-to-end data communication products
and solutions. He worked in Europe from 2012 to 2017, led the develop-
ment and innovation of data communication solutions across Europe,
and took the lead in the solution design of numerous data communication
projects for leading European carriers.
42.
Teams ◾ xxxvii
GangYan: Chief IGP expert of Huawei Data Communication Product
Line. Mr. Yan joined Huawei in 2000 and has been working in the VRP
department. From 2000 to 2010, he was responsible for the architecture
design of the VRP IGP sub-system and completed the design and delivery
of fast convergence, Loop-Free Alternate (LFA) Fast Reroute (FRR)/multi-
source FRR, and Non-Stop Routing (NSR). He has extensive experience
in protocol design, and since 2011, he has been responsible for improving
the O&M capabilities of VRP protocols, building the VRP compatibility
management system, and building and delivering the YANG model base-
line. He is currently leading the team to innovate IP technologies, includ-
ing SRv6 slicing, E2E 50 ms protection, and Bit Index Explicit Replication
(BIER)-related protocols.
44.
xxxix
Acknowledgments
While promoting theinnovation of SRv6 standards, numer-
ous people have contributed heavily throughout the process from
the initial concept to actual products and solutions, including strategic
decision-making, technical research, standards promotion, product devel-
opment, commercial deployment, and industry ecosystem construction.
We have received extensive support and help from both inside and outside
Huawei. We would like to take this opportunity to express our heartfelt
thanks to Kewen Hu, Shaowei Liu, Chenxi Wang, Jinzhu Chen, Meng
Zuo, Suning Ye, Xiao Qian, Jianbing Wang, Zhaokun Ding, Minwei Jin,
Dawei Fan, Jianping Sun, Mingzhen Xie, Jiandong Zhang, Nierui Yang,
Yue Liu, Zhiqiang Fan, Yuan Zhang, Xinbing Tang, Shucheng Liu, Juhua
Xu, Xinjun Chen, Lei Bao, Naiwen Wei, Xiaofei Wang, Shuying Liu, Wei
