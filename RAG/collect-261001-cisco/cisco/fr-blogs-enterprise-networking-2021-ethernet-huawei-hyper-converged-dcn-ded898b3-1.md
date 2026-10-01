---
id: collect-261001-cisco/cisco/fr-blogs-enterprise-networking-2021-ethernet-huawei-hyper-converged-dcn-ded898b3-1
title: "fr-blogs-enterprise-networking-2021-ethernet-huawei-hyper-converged-dcn-ded898b3"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2021-09"]
keywords: ["ethernet", "capex", "datacenter", "gpus", "latency", "memory", "research", "throughput", "training"]
source: docs/RAG/collect-261001-cisco/fr-blogs-enterprise-networking-2021-ethernet-huawei-hyper-converged-dcn-ded898b3.md
source_anchor: ""
source_lines: [1, 31]
sha256: aafd19d1a6b45295c07090dce03a8af0f3940f19bd287b4616d6f86307706f52
---

# fr-blogs-enterprise-networking-2021-ethernet-huawei-hyper-converged-dcn-ded898b3

Produits, solutions et services pour les entreprises
Cloud products, solutions & services
Produits, solutions et services pour les opérateurs
Smartphones, PC, tablettes, montres et plus
A propos de Huawei, presse, événements et plus
Digital transformation and the growing need for digital resiliency have made it clear to enterprises that their Data Center Networks (DCNs) must be extensively automated and modernized in order to support strategic business outcomes.
Currently, DCNs involve separate technologies and networks that satisfy different requirements. Enterprises have an Ethernet network to support general-purpose computing, a Fibre Channel (FC) network to support storage networking, and an InfiniBand network to support High-Performance Computing (HPC). All three networks are needed because their requirements and use cases are different.
General-purpose computing: Ethernet is widely used for general-purpose computing. Servers host virtualization technologies such as Virtual Machines (VMs) or containers to provide services for employees and external users through general-purpose computing networks, sometimes called application networks, service networks, and front-end networks.
Storage networking: FC has been adopted for storage networking. Storage nodes are interconnected through storage networks, including Storage Area Networks (SANs). SANs support high-performance Solid-State Drives (SSDs) by delivering increased Input/Output Operations Per Second (IOPS) and low-latency connections between server hosts and storage arrays.
HPC: InfiniBand is often used for HPC networks. These environments run very computationally intensive, latency-sensitive workloads, including big data analytics and advanced simulations. In these environments, where low latency is prized, applications tend to run on Bare Metal (BM) and servers are rarely virtualized. Computing units such as Central Processing Units (CPUs) and Graphics Processing Units (GPUs) are configured for HPC or Artificial Intelligence (AI) training, and server nodes are interconnected through HPC networks.
The problem with having three different networks is higher costs. Since the networks involve three different sets of procurement, deployment, and operations, separate Capital Expenditure (CAPEX) outlays are required. In addition, there are high requirements on operations specialization for personnel as they need to be knowledgeable about different protocols and have technology-specific expertise associated with each technology. Tools are also separate in this model, resulting in redundant costs.
Enterprises increasingly demand networking between data centers and clouds to cope with the requirements of hybrid Information Technology (IT), multicloud, and digital resiliency. Although every network technology has played a crucial role, only Ethernet is well placed to become a foundational technology for hyper-converged networking across all domains of the DCN, namely, general-purpose computing, storage networking, and networking for HPC.
Ethernet switching is commonplace in the market. As a result, the technology benefits from economies of scale, and the innovations that occur within its thriving open ecosystem. Based on the Ethernet ecosystem, Ethernet's continuing evolution makes it a great candidate for hyper-converged networks. It not only makes such networks possible, but also makes them viable for many use cases. Ethernet also has a significant bandwidth advantage over the other technologies. It has already evolved to 400GE, and will evolve to 800GE soon. Ethernet extends its advantage over the other technologies in terms of bandwidth and throughput, too. Impressive advances in achieving low latency and losslessness have also been made.
Meanwhile, Ethernet alternatives such as Remote Direct Memory Access over Converged Ethernet (RoCE) and various Non-Volatile Memory express (NVMe) over Fabrics (NoF) are now entering the market, and will become more common in the years to come. As such, Ethernet is well placed to address SAN and HPC use cases.
According to market research firm IDC's worldwide survey of enterprises across several major industry verticals, the respondents working in the finance, banking, media, entertainment, and gaming sectors, as well as those with three or more data centers, have the highest interest in a converged Internet Protocol (IP) data and storage network that is lossless, low latency, and high performance.
Figure 1: Interest in a Converged IP Data and Storage Network That is Lossless, Low Latency, and High Performance (by Industry)
Figure 2: Interest in a Converged IP Data and Storage Network That is Lossless, Low Latency, and High Performance (by Number of Data Centers)
N = 205
Source: Autonomous Driving Datacenter Networks Survey for Huawei, IDC, August, 2020
According to survey results, more than 55% of finance and banking respondents claimed that they were likely or extremely likely to deploy a converged IP data and storage network, compared to 46% of media, entertainment, and gaming respondents. Approximately 50% of those with a single data center said that they were likely or extremely likely to deploy such a network within 12 to 24 months, compared to 46% of those with two data centers, and 43% of those with three or more data centers.
Figure 3: Likelihood of Deploying a Converged IP Data and Storage Network That is Lossless, Low Latency, and High Performance in 12 to 24 Months (by Industry)
Figure 4: Likelihood of Deploying a Converged IP Data and Storage Network That is Lossless, Low Latency, and High Performance in 12 to 24 Months (by Number of Data Centers)
Source: IDC White Paper, sponsored by Huawei, The Hyper-converged Datacenter Network: How Ethernet Addresses the Need for Low Latency, Losslessness, and Full Life-Cycle Network Automation, doc #US48217221, September 2021
Hyper-converged DCNs offer a wide range of benefits.
Network consolidation with a single Ethernet network dramatically reduces the costs associated with redundant network infrastructure, including procurement, design, build, and operations costs.
Eliminating redundant networks also simplifies and improves the efficiency of network operations and management.
Automated network operations across the full lifecycle of a hyper-converged Ethernet DCN extend from Day 0 planning and design to maintenance and closed-loop change management and optimizations.
Root-cause analysis can be achieved quickly, leading to faster troubleshooting and remediation, which in turn yields highly available applications and services.
To cope with the challenges of providing an all-Ethernet hyper-converged DCN, Huawei has introduced its CloudFabric 3.0 Hyper-Converged DCN Solution, which consists of CloudEngine switches and iMaster-NCE, an intelligent network management and control system. The solution offers a wide range of leading features.
Fully lossless Ethernet: Ethernet is natively prone to packet loss, which has gone unresolved for 40 years. To overcome this, Huawei created an innovative iLossless algorithm that enables real-time and precise speed control, eliminating the heavy dependency on expert experience. The algorithm ensures zero packet loss on Ethernet and helps unleash 100% of computing power.
Full-lifecycle automated management: Huawei's CloudFabric 3.0 enables the industry's first L3 Autonomous Driving Network (ADN) that innovatively adopts digital twin, knowledge graph, and big data modeling technologies for network management. The resulting benefits include service provisioning in seconds, fault location in minutes, and proactive detection of 90% of network risks.
