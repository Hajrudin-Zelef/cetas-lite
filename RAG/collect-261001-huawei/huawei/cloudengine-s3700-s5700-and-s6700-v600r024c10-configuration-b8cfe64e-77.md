---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-77
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "license"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [11034, 11182]
sha256: 35e8b674f757b7ef7db3d9972b9682d405ea480bd4c884dec0bf4e9a856c74bd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    ●   CE4 at the branch egress of enterprise B
                        #
                        sysname CE4
                        #
                        vlan batch 50
                        #
                        interface 10ge1/0/1
                         undo portswitch
                         ip address 10.4.1.1 255.255.255.0
                        #
                        bgp 65440
                         peer 10.4.1.2 as-number 100
                         #
                         ipv4-family unicast
                          undo synchronization
                          import-route direct
                          peer 10.4.1.2 enable
                        #
                        return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        199
QoS Configuration
QoS Configuration                                                  13 Experience Assurance Configuration




  13                    Experience Assurance Configuration


                    13.1 Overview of Experience Assurance
                    13.2 Understanding Experience Assurance
                    13.3 Configuration Precautions for Experience Assurance
                    13.4 Default Settings for Experience Assurance
                    13.5 Configuring Experience Assurance


13.1 Overview of Experience Assurance
Definition
                    Experience assurance uses application identification technology to quickly and
                    accurately identify the application to which packets belong and manage
                    application traffic based on the identification results.

Purpose
                    As network and multimedia technologies develop rapidly, network applications are
                    becoming more and more diversified, audio and video conferences are widely used
                    for collaboration at work, and audio and video conference traffic on the network
                    increases rapidly. Other types of services preempt bandwidth for audio and video
                    services, affecting users' audio and video experience severely. When both key and
                    non-key application traffic is transmitted together, non-key services occupy a lot of
                    bandwidth resources, packets of key services are discarded, the delay and jitter are
                    uncontrollable, and the service quality cannot be guaranteed. If the traffic of key
                    and non-key applications cannot be accurately identified, differentiated policy
                    control cannot be implemented. After experience assurance is configured, the
                    device can identify applications and manage application traffic based on the
                    identification results.


13.2 Understanding Experience Assurance

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            200
QoS Configuration
QoS Configuration                                                  13 Experience Assurance Configuration


                    Experience assurance uses application identification technology to identify traffic
                    and uses packet re-marking technology to set the priority field of specified
                    application packets. In this way, these packets can be scheduled or forwarded
                    based on the re-marked priority, making it possible to manage application traffic
                    and ensure users' Internet access experience.
                    For example, packets identified as belonging to a video conference application
                    that has high requirements on delay and service quality can be re-marked with a
                    higher priority. As such, these packets have a higher scheduling weight in
                    subsequent forwarding, thereby ensuring video conference experience. Similarly,
                    for applications that do not have special requirements on delay or service quality,
                    the priority of their packets can be reduced so that sufficient network resources
                    are provided for applications that do have such requirements.

Implementation Process
                    Figure 13-1 shows the implementation process of experience assurance.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            201
QoS Configuration
QoS Configuration                                                   13 Experience Assurance Configuration


                    Figure 13-1 Implementation process of experience assurance




                    The following phases are involved in the implementation process of experience
                    assurance:
                    1.   Updating a service awareness signature database (SA-SDB)
                         Because different applications have different signatures, obtaining the unique
                         signatures of application protocols is the key for devices to accurately identify
                         various applications. Huawei periodically analyzes the signatures of common
                         applications on networks and stores the signatures in an SA-SDB. After being
                         loaded on a device, an SA-SDB exists as a predefined application. It can be

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             202
QoS Configuration
QoS Configuration                                                  13 Experience Assurance Configuration


                         referenced by other signatures as a matching condition, but cannot be
                         modified or deleted.
                         Huawei Security Center (isecurity.huawei.com) periodically releases the
                         latest SA-SDB to keep up with the ever-changing application software. With
                         an up-to-date SA-SDB, devices are fully capable of identifying applications in
                         traffic. It is recommended that the SA-SDB be updated once a week.
                    2.   Configuring MQC-based experience assurance
                         The experience assurance function re-marks packets of identified application
                         traffic. Currently, this function supports the following packet re-marking
                         modes:
                         –    Re-marking 802.1p values of VLAN packets
                         –    Re-marking DSCP values of IP packets


13.3 Configuration Precautions for Experience
Assurance
Licensing Requirements
                    Experience Assurance is not under license control.


Hardware Requirements

                    Table 13-1 Hardware requirements

                     Series                         Models

                     S5755E-H                       S5755E-H24HB2Y2CZ

                     S6730E-H-V2                    S6730E-H6FX4Y2CZ-V2

                     S6750E-S                       S6750E-S24T16X8Y2CZ, S6750E-S16X10Y2CZ

                     S5735E-S-V2                    S5735E-S48HS4XE-V2, S5735E-S24HS4XE-V2

