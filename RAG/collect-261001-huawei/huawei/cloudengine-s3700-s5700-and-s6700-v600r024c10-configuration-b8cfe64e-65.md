---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-65
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "latency", "parameters", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9195, 9375]
sha256: e9d04443e970b8a81d708cc6b55aaa3e8343530185438fc0a4a68f8793cf3538
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                                                   11 Congestion Management Configuration

                    8021p-inbound 6 phb ef green
                    8021p-inbound 7 phb cs7 green
                    8021p-outbound be green map 0
                    8021p-outbound be yellow map 0
                    8021p-outbound be red map 0
                    ...

                    In the DiffServ domain, 802.1p values 6, 5, and 2 are mapped to CoS values EF,
                    AF3, and AF1, respectively, and packets are colored green, yellow, and red,
                    respectively.
                    # Check the configuration of 10GE 1/0/3. You can see the scheduling parameters
                    of queues with different CoS values.
                    <DeviceB> display qos configuration interface 10ge 1/0/3
                     interface 10GE1/0/3
                     --------------------------------------------------------------------------
                     trust flag        : outer 8021p
                     diffserv domain : default
                     dei enable          : disable
                     port priority      :0
                     phb marking 8021p : disable
                     phb marking dscp : disable
                     phb marking exp : -
                     port wred           :-
                     port lr         : cir = -, cbs = -
                     port car inbound : -
                     port car outbound : -
                     schedule profile : -
                     --------------------------------------------------------------------------
                     queue          shaping        schedule      wred
                               cir      pir
                               cbs        pbs
                     --------------------------------------------------------------------------
                     0            -        - drr        -
                                 -        - weight = 1

                    1           -        -   drr       -
                               -        -    weight = 50

                    2           -        -   drr       -
                               -        -    weight = 1

                    3           -        -   drr       -
                               -        -    weight = 100

                    4           -        -   drr       -
                               -        -    weight = 1

                    5           -        -   pq         -
                               -        -

                    6           -        -   pq         -
                               -        -

                    7           -        -   pq         -
                               -        -

                    --------------------------------------------------------------------------


Configuration Scripts
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 100 200


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       167
QoS Configuration
QoS Configuration                                                11 Congestion Management Configuration

                    #
                    diffserv domain ds1
                     8021p-inbound 2 phb af1 red
                     8021p-inbound 5 phb af3 yellow
                     8021p-inbound 6 phb ef green
                    #
                    interface 10GE1/0/1
                     port link-type access
                     port default vlan 100
                     trust upstream ds1
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 200
                     trust upstream ds1
                    #
                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 100 200
                     qos drr 0 to 4
                     qos queue 1 drr weight 50
                     qos queue 3 drr weight 100
                    #
                    return



11.10 Example for Configuring Congestion Avoidance
and Congestion Management (PQ+WDRR Scheduling
and WRED Profile)
Networking Requirements
                    DeviceB is connected to DeviceA through interface 1. The 802.1p priorities of
                    voice, video, and data service packets from the Internet are 6, 5, and 2,
                    respectively. Packets of these services can reach users through DeviceA and
                    DeviceB, as shown in Figure 11-5. Because the rate of the inbound interface
                    interface 1 on DeviceB is higher than the rates of outbound interfaces interface 2
                    and interface 3, congestion may occur on the two outbound interfaces.
                    To reduce the impact of network congestion and guarantee high-priority and
                    latency-sensitive services, set congestion avoidance and congestion management
                    parameters according to Table 11-4 and Table 11-5.

                    Table 11-4 Congestion avoidance parameter settings
                     Service Type         Color           Lower           Upper         Drop
                                                          Threshold       Threshold     Probability
                                                          (%)             (%)

                     Voice                Green           80              100           10

                     Video                Yellow          60              80            20

                     Data                 Red             40              60            40




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          168
QoS Configuration
QoS Configuration                                                      11 Congestion Management Configuration


                    Table 11-5 Congestion management parameter settings

                     Service Type                      CoS Value                      WDRR

                     Voice                             EF                             0

                     Video                             AF3                            100

                     Data                              AF1                            50




                    Figure 11-5 Network diagram for configuring congestion avoidance and
                    congestion management
                          NOTE

                        In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                        and 10GE 1/0/3, respectively.




Procedure
         Step 1 Configure VLANs for each interface so that devices can communicate with each
                other at the link layer.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 2 5 6
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 2 5 6
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type trunk
                    [DeviceB-10GE1/0/2] port trunk allow-pass vlan 2 5 6


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     169
QoS Configuration
QoS Configuration                                                         11 Congestion Management Configuration

