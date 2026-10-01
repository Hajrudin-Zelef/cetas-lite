---
id: collect-261001-huawei/huawei/ai-campus-white-paper-26h1-0f8c526c-1
title: "ai-campus-white-paper-26h1-0f8c526c"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: ["alignment", "research"]
source: docs/RAG/collect-261001-huawei/ai-campus-white-paper-26h1-0f8c526c.md
source_anchor: ""
source_lines: [1, 174]
sha256: df76604b812aaa316fb15c36ae2a84c863573bb26fdda1db4a4d9ae692954ed9
---

# ai-campus-white-paper-26h1-0f8c526c

Building a Fully Connected,
          Intelligent World
Chief editing organization

Huawei Technologies Co., Ltd.




Participating organizations

Beijing Bridata Technology Co., Ltd.
Beijing Digital Hail Information Technology Co., Ltd.
Chinasoft International Co., Ltd.
Guangdong Flying Enterprise Internet Technology Co., Ltd.
Hangzhou Zhongling Future Technology Co., Ltd.
Hi-Think Technology, Corp.
HMN Smart Co., Ltd.
iSoftStone Smart Technologies Co., Ltd.
Lianfa Group Co., Ltd.
East China Architectural Design & Research Institute Co., Ltd.
Shandong Bittel Intelligent Technology Co., Ltd.
Shanghai Yuankong Automation Technology Co., Ltd.
Shenzhen Webuild Technology Co., Ltd.
SHENZHEN HUAZHAO TECHNOLOGY CO., LTD.
Sichuan Timwave Yunze Intelligent Technology Co., Ltd.
Wanchow Qizhi (Qingdao) Information Technology Co., Ltd.
Zhuhai Xiangyi Aviation Technology Co., Ltd.
Editorial advisors (in alphabetical order)

Chen Zhiwei, Deng Xiao, Dong Shaohua, Gu Rui, Guo Hongfu,
Guo Wanzhou, Han Bing, He Ji, Hu Yanxin, Liu Jiajia, Liu Shuqing,
Liu Song, Liu Xi, Liu Yue, Su Yufeng, Tan Jianxin, Wang Gang,
Wang Jiankang, Wang Jingji, Wang Xiaoan, Wang Zhong,
Wei Zequn, Xiao Guangrui, Xu Bo, Xue Long, Yang Xi, Yi Yangao,
Yu Xunxun, Yue Xuefeng, Zeng Eryang, Zhang Jirong, Zhao Shaoqi,
Zhao Xianghui




Chief editors

Wang Zhiming, Liu Yaoyao, Niu Kun, Zheng Rongjian, Song Junjing,
Zhuang Guanglong, Chen Dan, Jiang Lin, Chen Zhongzhou, Zhang
Jiabao, Xu Yong, Liu Zhipiao, Jiang Yinqing, Pan Feng, Zhang Juan




Contributors (in alphabetical order)

Chen Suixiao, Ge Feng, Guo Jiani, He Wenxi, Hu Xiumin, Lei Shifei,
Li Honggang, Liu Bi, Liu Guanjie, Ouyang Wanru, Sun Shengbai,
Sun Yang, Wang Binbin, Wang Bo, Wang Haijiao, Wang Liangsheng,
Xia Wei, Zhou Yang, Zhu Huizhang
Foreword I




Yang Chaobin

Huawei's Director of the Board and CEO of the ICT Business Group




Since we first defined intelligent campuses alongside our customers, partners, and industry
stakeholders in 2019, Huawei has delivered over 1,500 projects worldwide, driving rapid digital
transformation across various industries.

Today, at a pivotal moment in intelligent industrial transformation, we recognize that AI's true
value emerges from the integration of applications, data, and ICT infrastructure in real-world
settings. Campuses not only provide the physical foundation for these scenarios but also act as
essential testbeds for translating AI potential into tangible industrial advancement.

AI Campus represents a new model for campuses, evolving in alignment with the prevailing trends
of intelligence. Centered on the innovative "intelligent spaces," it combines computing power
and data flows to create an open, collaborative, and co-evolving industrial ecosystem. Working
closely with dozens of leading customers, partners, and research institutions, we have compiled
industry-wide insights into one paper. Drawing on this collective expertise and hands-on campus
experience, we developed the AI Campus concept, defined its core features, and established a
reference architecture and implementation roadmap. This white paper aims to be a practical
guide for industries planning and building intelligent campuses in the AI era; its core value lies in
this blueprint.
Foreword II




Michael Ma

Vice President of Huawei and President of
Huawei's ICT Product Portfolio Mgmt & Solutions Dept




Every advancement in intelligent campuses stems from the integration of technological innovation
and industry practices. Huawei goes beyond offering products and solutions; we drive campus
transformation. Capitalizing on our full-stack cloud, network, edge, and device capabilities, we
collaborate with partners to innovate together and consistently create value for our customers.

AI is now advancing campuses from isolated to all-scenario intelligence. To support this
transition, Huawei has partnered with dozens of customers, partners, and experts from
research institutions to hold more than 20 specialized workshops. Combining industry insights
with Huawei's hands-on experience, we jointly authored the AI Campus white paper. This
paper explores 17 intelligent spaces across 11 business segments, such as intelligent office,
production, intelligent guest rooms, and commercial consumption. It also defines the six
technical features of AI campuses for the first time: HI-AI coexistence, data-AI convergence,
all-domain connectivity, multi-dimensional sensing, resilient security, and eco-sustainability.
Additionally, the white paper outlines the reference architecture of AI campuses and provides
suggestions on integrated planning, construction, and operations.

We recognize that AI evolves rapidly and campus scenarios will continue to change. This paper
summarizes and reviews current technological innovation and industry practices. Looking ahead,
we will remain dedicated to ongoing exploration and reflection and welcome your insights to help
shape the future of AI campuses.
Foreword III




Deng Yiou

Secretary-General of NIDA




Since its establishment, the Network Innovation and Development Alliance (NIDA) has remained
committed to defining the evolution path and network construction standards of global Internet
infrastructure, and driving the upgrade of the data communication industry. As AI deeply
integrates into industries, fixed networks are transitioning from connecting people to connecting
intelligence. Innovative network architectures are now fundamental to maximizing AI's computing
power and enabling intelligent applications.

As core units of urban and industrial development, campuses serve as essential carriers for
translating the value of AI into real-world impact, and as ideal proving grounds for Net5.5G. The
AI Campus white paper outlines future intelligent campus scenarios centered on the concept of
"intelligent spaces." Its key technical features, including HI-AI coexistence, all-domain connectivity,
and multi-dimensional sensing, closely align with NIDA's Net5.5G vision. This alignment offers a
clear architectural blueprint and capability framework for advancing campus networks.

We believe the release of this white paper will strengthen industry consensus and deliver a
definitive roadmap for planning and developing AI campuses. Moving forward, NIDA will continue
to collaborate with industry partners to advance campus networks toward a smarter, greener, and
more sustainable future. Together, we are dedicated to strengthening the network foundation
that supports intelligent upgrades across diverse industries.
CONTENTS



01
Initiating with AI: Driving Forces and the Evolution of AI Campuses                                                                                                                         02

1.1 Four Driving Forces of AI Campuses������������������������������������������������������������������������������������������������������������������������������ 03
1.2 Four Evolution Directions of AI Campuses������������������������������������������������������������������������������������������������������������������ 10
1.3 AI Campus Evolution: From Traditional Campus to AI Campus�������������������������������������������������������������������������� 22




02
Innovating with AI: Faster Evolution Toward Intelligent Spaces for AI Campuses 26

