---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-6
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [407, 484]
sha256: 6242664b6f601b79c88d1f546a946d0c6aa3607fae552ad194f52a2b13bfb8c6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    ●   Reliability design declaration
                        Network planning and site design must comply with reliability design
                        principles and provide device- and solution-level protection. Device-level
                        protection includes planning principles of dual-network and inter-card dual-
                        link to avoid single point or single link of failure. Solution-level protection
                        refers to fast convergence protection mechanisms such as FRR and VRRP. If
                        solution-level protection is used, ensure that the primary and backup paths do
                        not share links or transmission devices. Otherwise, solution-level protection
                        may fail to take effect.

Reference Standards and Protocols
                    To obtain reference standards and protocols, log in to Huawei official website,
                    search for "standard and protocol compliance list", and download the Huawei S-
                    Series Switch Standard and Protocol Compliance List. If you have not obtained the
                    access permission of the document, see Help on the website to find out how to
                    obtain it.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            4
QoS Configuration
QoS Configuration                                                                     2 Overview of QoS




                                                        2         Overview of QoS

Introduction to QoS
                    As networks grow in both scope and service diversity, the resulting sharp increase
                    in traffic exacerbates network congestion and increases the forwarding delay,
                    while in some cases even causing packet loss. Any one of these situations will
                    cause services to experience deterioration in quality or even total interruption. To
                    support increasingly diversified services such as gaming, video, and live streaming,
                    it is critical that network congestion be resolved. One answer is to increase
                    network bandwidth, but this comes with increased costs. A better, more cost-
                    effective solution is to use an "assured" policy to manage traffic.

                    Quality of service (QoS) is able to offer such a solution. It does this by providing
                    end-to-end (E2E) service quality guarantee to meet diversified service
                    requirements. It is important to note that QoS does not increase the network
                    bandwidth. Instead, it improves network resource utilization and allows different
                    types of traffic to compete for network resources based on their priorities, so that
                    voice, video, and important data applications are processed preferentially on
                    network devices. QoS is now widely used on, and vital to, Internet applications.


DiffServ
                    QoS is an overall solution, instead of being merely a single function. When two
                    hosts on a network communicate with each other, traffic between them may
                    traverse a large number of devices. QoS can guarantee E2E service quality only
                    when all devices on the network use a unified QoS service model, of which
                    Differentiated Services (DiffServ) is the most commonly used.

                    DiffServ classifies packets on a network into multiple classes and provides
                    differentiated processing for each class. In this way, when congestion occurs,
                    classes with a higher priority are given preference. Packets of the same class are
                    aggregated and sent as a whole to ensure the same delay, jitter, and packet loss
                    rate.

                    In the DiffServ model, a device can flexibly classify packets based on a
                    combination of conditions and re-mark packets with different priorities. The other
                    devices simply need to identify the priorities carried in packets to allocate
                    resources and control traffic.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  5
QoS Configuration
QoS Configuration                                                                         2 Overview of QoS


                    QoS technologies supported by the device are based on the DiffServ service
                    model. In this chapter, we describe QoS technologies in terms of packet
                    classification methods and QoS service technologies.

