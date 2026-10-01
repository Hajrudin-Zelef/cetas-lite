---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-6
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [792, 922]
sha256: 13e262d57ea9a6873059bff3e0ebbd40301be9424a3a82fd0b774a38368a3fe0
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

Hu, Banghua Chen, Su Feng, Jianwu Hao, Dahe Zhao, Jianming Cao,
Yahao Zhang, Minhu Zhang, Yi Zeng, Jiandong Jin, Weiwei Jin, Xing
Huang, Xiaohui Li, Ming Yang, Wenjiang Shi, Xiaoxing Xu, Yiguang Cao,
Qingjun Li, Gang Zhao, Xiaoqi Gao, Jialing Li, Peng Zheng, Peng Wu,
Xiaoliang Wang, Fu Miao, Fuyou Miao, Min Liu, Xia Chen, Yunan Gu,
Jie Hou, Yuezhong Song, Ling Xu, Bing Liu, Ning Zong, Qin Wu, Liang
Xia, Bo Wu, Zitao Wang, Yuefeng Qiu, Huaru Yang, Xipeng Xiao, Dhruv
Dhody, Shucheng Liu, Hongjie Yang, Shuanglong Chen, Ruizhao Hu, Xin
Yan, Hong Shen, Wenxia Dong, Guanjun Zhou, Yuanyi Sun, Leyan Wang,
Shuhui Wang, Xiaohui Tong, Mingyan Xi, Xiaoling Wang, Yonghua Mao,
Lu Huang, Kaichun Wang, Chen Li, Jiangshan Chen, Guoyou Sun, Xiuren
Dou, Ruoyu Li, Cheng Yang, Ying Liu, Hongkun Li, Mengling Xu, Taixu
Tian, Jun Gong, Yang Xia, Fenghua Zhao, Pingan Yang, Yongkang Zhang,
Guangying Zheng, Sheng Fang, Chuang Chen, Ka Zhang, Yu Jiang, Hanlin
Li, Ren Tan, Yunxiang Zheng, Yanmiao Wang, Weidong Li, Zhidong Yin,
Yuanbin Yin, Zhong Chen, Chun Liu, Xinzong Zeng, Mingliang Yin,
45.
xl ◾ Acknowledgments
FengqingYu, Weiguo Hao, Shuguang Pan, Haotao Pan, Wei Li, and other
leaders and colleagues from Huawei. We sincerely thank Hui Tian, Feng
Zhao, Yunqing Chen, Huiling Zhao, Chongfeng Xie, Fan Shi, Bo Lei,
Qiong Sun, Aijun Wang, Yongqing Zhu, Huanan Chen, Xiaodong Duan,
Weiqiang Cheng, Fengwei Qin, Zhenqiang Li, Liang Geng, Peng Liu,
Xiongyan Tang, Chang Cao, Ran Pang, Xiaohu Xu, Ying Liu, Zhonghui
Li, Suogang Li, Tao Huang, Jiang Liu, Bin Yang, Fengyu Zhang, Dong
Liu, Dujuan Gu, and other technical experts in China’s IP feld who have
innovated in technology and promoted the formulation of standards for
a long time now. We also sincerely thank Adrian Farrel, Loa Andersson,
James Guichard, Bruno Decraene, Carsten Rossenhovel, Clarence Filsfls,
Keyur Patel, Pablo Camarillo Garvia, Daniel Voyer, Satoru Matsushima,
Alvaro Retana, Martin Vigoureux, Deborah Brungard, Susan Hares,
John Scudder, Joel Halpern, Matthew Bocci, Stephane Litkowski, Lou
Berger, Christian Hopps, Acee Lindem, Gunter Van de Velde, Daniel
King, Nicolai Leymann, Tarek Saad, Sam Aldrin, Stewart Bryant, Andy
Malis, Julien Meuric, Mike McBride, Jef Tantsura, Chris Bowers, Zhaohui
Zhang, Jefrey Haas, Reshad Rahman, Tony Przygienda, Zafar Ali, Darren
Dukes, Francois Clad, Kamran Raza, Ketan Talaulikar, Daniel Bernier,
Daisuke Koshiro, Gaurav Dawra, Jaehwan Jin, Jongyoon Shin, Ahmed
Shawky, Linda Dunbar, Donald Eastlake, Yingzhen Qu, Haoyu Song,
Parviz Yegani, Huaimo Chen, Young Lee, Ifekhar Hussain, Himanshu
Shah, Wim Henderickx, Ali Sajassi, Patrice Brissette, Jorge Rabadan, and
other IP experts in the industry who helped us with work on SRv6. Finally,
we would like to express our special thanks to Stefano Previdi and Latif
Ladid, both of whom wrote Forewords for this book.
We hope that this book can, as comprehensively as possible, present
basic SRv6 technologies and emerging technologies targeted at 5G and
cloud services to help readers comprehensively understand the fundamen-
tals of SRv6, its benefts to the industry, and its far-reaching impact on
communications networks. As SRv6 is an emerging technology and our
capabilities are limited, errors and omissions are inevitable. We will be
grateful for any information that helps to rectify these. Please send any
comments and feedback by email to lizhenbin@huawei.com.
46.
xli
Authors
Zhenbin Li isthe Chief Protocol Expert of Huawei and member of the
IETF IAB, responsible for IP protocol research and standards promotion
at Huawei. Zhenbin Li joined Huawei in 2000. For more than a decade,
he had been responsible for the architecture, design, and development
of Huawei’s IP operating system—Versatile Routing Platform (VRP)—
and MPLS sub-system as an architect and system engineer. From 2015
to 2017, Zhenbin Li worked as an SDN architect and was responsible for
the research, architecture design, and development of network control-
lers. Since 2009, Zhenbin Li has been actively participating in the innova-
tion and standardization work of the IETF, and in the past 6 years, he has
been continuously promoting the innovation and standardization of BGP,
PCEP, and NETCONF/YANG protocols for SDN transition. His research
currently focuses on Segment Routing over IPv6 (SRv6), 5G transport,
telemetry, and network intelligence, and he has led and participated in
more than 100 IETF RFCs/drafs as well as applied for more than 110
patents. In 2019, he was elected a member of the IETF IAB to undertake
Internet architecture management from 2019 to 2021.
Zhibo Hu is a Senior Huawei Expert in SR and IGP, responsible for SR
and IGP planning and innovation. Currently, Zhibo Hu is mainly engaged
in the research of SR-MPLS/SRv6 protocols and 5G network slicing tech-
nologies. Since 2017, Zhibo Hu has been actively participating in the inno-
vation and standardization work of the IETF and leading in standards
related to SRv6 reliability, SRv6 YANG, 5G network slicing, and IGP, and
is committed to leveraging SRv6 innovation to support network evolution
toward 5G and cloudifcation.
47.
xlii ◾ Authors
ChengLi is a Huawei Senior Pre-research Engineer and IP standards rep-
resentative, responsible for Huawei’s SRv6 research and standardization,
involving SRv6 extension header compression, SRv6 OAM/path segments,
SFC, and security. Cheng Li started to participate in IETF meetings from
2018. Up to now, Cheng Li has submitted more than 30 individual drafs,
more than ten of which have been promoted to working group drafs.
C H AP T E R 1
SRv6 Background
In this chapter, we expand on the history of Internet technology
development, from the competition between Asynchronous Transfer
Mode (ATM) and IP to the emergence of Multiprotocol Label Switching
(MPLS), the dawning of the All IP 1.0 era and the consequent challenges,
as well as the rise of Sofware Defned Networking (SDN). We then sum
up by introducing Segment Routing over IPv6 (SRv6), which is a key tech-
nology to enable the All IP 2.0 era.
1.1 OVERVIEW OF INTERNET DEVELOPMENT
Humans develop in tandem, as opposed to individually, through com-
munication and collaboration. Over our extensive history, various com-
munications modes have emerged, ranging from beacon fres to emails,
and from messenger birds to quantum entanglement. Te scope of com-
munications has also expanded from people within a certain vicinity to
people in distant localities, across a whole country, around the globe, and
even in outer space. Tat said, we as a people have never stopped our pur-
suit of communication technologies, which in turn have boosted human
prosperity. It is, therefore, safe to say that we are no longer content with
just human-to-human communication, or put diferently, we now aim at
achieving the connectivity of everything anytime and anywhere using the
Internet. It is within this context that the development of the Internet has
profoundly impacted the path of human society.
Afer years of development, the Internet has become almost as essen-
tial as water and electricity. Although the Internet has made life more
3
51.
4 ◾ SRv6Network Programming
FIGURE 1.1 Internet development milestones.
convenient through the information age, few people know the ins and outs
of its technological development history. With this in mind, we briefy
summarize the history of Internet technology development in Figure 1.1.
In 1969, Advanced Research Projects Agency Network (ARPANET) —
the major predecessor of the Internet — came into existence.
In 1981, IPv4[1] was defned.
In 1986, the Internet Engineering Task Force (IETF), dedicated to
formulating Internet standards, was founded.
In 1995, IPv6,[2] the next generation of IPv4, was standardized.
In 1996, MPLS[3] was proposed.
In 2007, SDN[4] debuted.
In 2008, the OpenFlow[5] protocol was introduced.
In 2013, SR[6] was proposed, including Segment Routing over MPLS
(SR-MPLS)[7] and SRv6.[8]
