---
id: collect-240926-tomshardware/tomshardware/elon-musk-receives-ftc-greenlight-to-buy-mesh-optical-as-interconnects-emerge-as-1
title: "elon-musk-receives-ftc-greenlight-to-buy-mesh-optical-as-interconnects-emerge-as"
domain: tomshardware
role: reference
task: reference
actors: ["Cohere", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["amd", "antitrust", "asic", "blackwell", "compute", "gpus", "humanoid", "inference", "memory", "nvidia", "optics", "packaging"]
source: docs/RAG/clean_en/tomshardware/elon-musk-receives-ftc-greenlight-to-buy-mesh-optical-as-interconnects-emerge-as.md
source_anchor: ""
source_lines: [1, 42]
sha256: 5bb5321f100017204191c24a008a85d799ec31f9b88aecf9fd94f095dff16198
---

# elon-musk-receives-ftc-greenlight-to-buy-mesh-optical-as-interconnects-emerge-as

<!-- source: https://www.tomshardware.com/tech-industry/big-tech/elon-musk-receives-ftc-greenlight-to-buy-mesh-optical-as-interconnects-emerge-as-ais-tightest-bottleneck-the-move-will-expand-musks-growing-stack-of-critical-ai-infrastructure -->

Elon Musk has received the go-ahead from the Federal Trade Commission (FTC) to acquire Mesh Optical Technologies, an AI infrastructure startup that develops light-based networking hardware for data centers. Records published by the FTC on June 25 show that the regulatory body granted early termination of its antitrust review of the transaction, permitting Musk to procure Mesh. While the deal is yet to be finalized, with no official statement from either party, the government's green light indicates it’s all but done, as this was the last hurdle.

Interestingly, Mesh was founded by three former SpaceX employees who helped develop the Starlink optical communication links that keep thousands of satellites interconnected. So, why is Musk — who is simultaneously building the world's largest multibillion-dollar semiconductor manufacturing facility and an 11-million-square-foot orbital data center factory — seeking to own a company founded by his former employees? The answer appears to be optical interconnects, a critical technology that connects all three.

## The connection problem: AI's latest bottleneck

As AI continues to grow in capability and user base, so do the enabling AI clusters, many of which now comprise tens to hundreds of thousands of processors. The hardest problem in scaling an AI cluster has evolved beyond making the chips faster to moving data between them. Training and inference tasks on frontier AI models are split across thousands of GPUs using parallel-computing techniques, requiring the processors to exchange enormous volumes of data every fraction of a second.

While per-chip compute capacity has raced ahead, the bandwidth linking those chips has not kept pace, a mismatch the industry refers to as the "I/O wall." The processors mostly communicate via copper interconnects, which currently dominate AI clusters. However, copper presents inherent limitations. As per-lane signaling climbs toward 200 gigabits per second (Gbps), attenuation, crosstalk, and the skin effect all worsen at higher frequencies, driving up power and corrupting the signal until passive copper becomes impractical beyond a meter or two.

To overcome these constraints, the industry is increasingly turning to optical networking, bringing the technology closer to the processor. Optical links use transceivers to convert a chip's electrical signals into light for transmission over fiber, then convert them back into electrical signals at the receiving end. They can carry far more data over much longer distances while consuming less power than equivalent high-speed copper connections, making them increasingly essential as AI clusters grow larger. Chipmakers and networking vendors are racing to deliver faster 800G and 1.6T optical transceivers while shortening electrical paths with co-packaged optics, which place the optical engine alongside the switch ASIC (application-specific integrated circuit).

This shift has transformed optical interconnects from a supporting technology into one of the industry's most strategically important AI infrastructure markets, attracting billions of dollars in investments and resulting in major partnerships for new and existing industry players. One such player is Mesh, the optical hardware startup that has drawn the interest of the world’s richest man.

## A mesh solution to Musk’s ambition?

Elon Musk has been one of the most aggressive players in the AI industry. After co-founding OpenAI, he went on to launch a proprietary company, xAI, before turning his focus to building data centers. In less than two years, xAI deployed the Colossus supercomputer with over 200,000 Nvidia Hopper- and Blackwell-generation accelerators. Colossus 2, with a long-term target of 1 million GPUs, is already operational. For Musk, however, buying the chips was not enough. Why not build them, too?

Characteristic of the world's richest man’s preference for complete vertical integration, SpaceX — in collaboration with Tesla and xAI — is now building Terafab, a vertically integrated, multi-billion-dollar semiconductor manufacturing facility aimed at producing chips capable of delivering an unprecedented over 1 terawatt of AI compute capacity annually. Located in Austin, Texas, the colossal facility aims to consolidate every stage of chip production under one roof, handling everything from logic and memory fabrication to advanced packaging and testing. An ambitious project that we've also analyzed for its feasibility.

The facility's output will serve to meet the chip needs of the broader AI industry, as well as those of Musk’s xAI, self-driving vehicles, Optimus humanoid robots, and SpaceX's orbital AI data center plans. Musk says 80% of Terafab's total compute output is ultimately destined for Earth orbit to support SpaceX's orbital data centers.

“But there aren't any data centers floating around in space,” observers may point out. Introducing Gigasat, Musk's 11-million-square-foot fix for that reality. Gigasat is yet another massive facility under construction, this time for manufacturing everything needed for SpaceX’s AI1 satellite, the company's most likely world-first orbital data center with 150 kW of compute.

At first glance, everything seems in place for the next generation of Ultra-capable AI infrastructure. However, there is one critical missing piece in this stack, one that we've established earlier. Hundreds of gigawatts of extremely powerful silicon are not particularly useful if the data can't move between the dies fast enough in AI clusters, whether on the ground or in space. The industry-prevalent copper hits a wall long before you reach the scale Musk is chasing. Hence, the need for the missing piece: optical interconnects.

这就引出了Mesh——一家制造正是那块缺失拼图的厂商。Mesh Optical Technologies是一家美国光通信初创公司，开发高速光互连硬件——光收发器，将芯片的电信号转换为光信号，以便通过光纤进行高速传输——面向AI数据中心和太空通信。

其旗舰产品Alpha C1支持800G和1.6T数据速率，据称功耗仅为竞争模块的三分之一，采用了一种倒装芯片键合工艺，该公司表示该工艺使光引擎能够在AI集群所需的规模上实现可重复制造——潜在可达数百万条链路。

这些正是无缝互连下一代地面AI超级计算机以及潜在未来天基计算平台所需的特性，而Terafab正是旨在实现这些目标。一个额外的好处是Mesh三位创始人的太空相关经验，他们恰好是前SpaceX员工，曾帮助构建连接Starlink星座的激光星间链路。

同样，以典型的马斯克风格，他不是简单地购买硬件，而是着手收购整家公司，获得其研发和供应链的完全控制权。如果这笔交易——几乎已经敲定——顺利完成，马斯克将拥有推动AI行业未来所需的关键基础设施的完整技术栈。

## 聪明钱正流向光互连

SpaceX生态系统只是众多认识到光网络在AI中巨大技术和经济重要性的实体之一。AI芯片制造商正在积极投资光供应链，以确保制造产能并防止硬件瓶颈。

仅英伟达一家就已承诺向组件制造商Coherent和Lumentum投入据报道40亿美元以锁定供应。在其他地方，包括微软、Meta和OpenAI在内的多家超大规模企业已与硬件巨头博通、AMD和英伟达联手，成立了光学计算互连（OCI）多源协议（MSA）小组，目标是开发面向AI集群的协议无关的纵向扩展互连技术。

