---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-139
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15396, 15508]
sha256: 69a7dae1aa286350fc2067a02a94fd30894e92d24d9a71dd186c1e39ef550120
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     Configure the alarm    sa flow-table           By default, the lower and     To
                     percentage of the      usage-alarm             upper alarm percentage        adj
                     application            percentage lower-       thresholds for the            ust
                     identification flow    percent upper-          application identification    the
                     table usage.           percent                 flow table usage are 50%      alar
                                                                    and 80%, respectively.        m
                                                                                                  thr
                                                                                                  esh
                                                                                                  old
                                                                                                  for
                                                                                                  the
                                                                                                  app
                                                                                                  lica
                                                                                                  tion
                                                                                                  ide
                                                                                                  ntifi
                                                                                                  cati
                                                                                                  on
                                                                                                  flo
                                                                                                  w
                                                                                                  tabl
                                                                                                  e
                                                                                                  usa
                                                                                                  ge,
                                                                                                  use
                                                                                                  this
                                                                                                  fun
                                                                                                  ctio
                                                                                                  n.



                    ----End

13.5.4 Verifying the Configuration

Procedure
                     Operation                                 Command

                     Check the enabling status of the SA       display interface sa-configuration
                     function on an interface.                 [ interface-name | interface-type
                                                               interface-number ]
                     Check global SA configurations.           display sa global configuration


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          283
QoS Configuration
QoS Configuration                                                       13 Experience Assurance Configuration


                     Operation                                      Command

                     Check information about the                    display sa flow-table [ statistic ] { all
                     application identification flow table.         | { application application-name |
                                                                    source-ip src-ipv4-address |
                                                                    destination-ip dest-ipv4-address |
                                                                    source-port src-port | destination-
                                                                    port dest-port-number | protocol { tcp
                                                                    | udp } | vpn-instance vpn-instance }
                                                                    *} slot slotid
                                                                    NOTE
                                                                     Before you specify a VPN instance in this
                                                                     command, the VPN instance must have
                                                                     been created.

                     Check the configuration status of the          display sa whitelist applied-record
                     application identification whitelist.


13.5.5 Example for Configuring MQC-based Experience
Assurance

Networking Requirements
                    On an enterprise network shown in the figure, packets sent from Host1 and Host2
                    to DeviceB are identified by different VLAN IDs. The VLAN ID for Host1 is 10, and
                    that for Host2 is 20. To ensure the voice and video conference experience of the
                    enterprise, you need to configure experience assurance on DeviceB so that voice
                    and video application traffic is preferentially forwarded.

                    Figure 13-2 Network diagram for configuring experience assurance
                         NOTE

                        In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                        10GE1/0/3, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       284
QoS Configuration
QoS Configuration                                                       13 Experience Assurance Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:
                    1. Create VLANs and configure interfaces so that DeviceB can communicate with
                    Host1, Host2, and DeviceA.
                    2. Enable application identification on inbound interfaces.
                    3. Configure a traffic policy to match a specific application, and configure a rule of
                    re-marking the packet priority in the traffic behavior.
                    4. Apply the traffic policy to interfaces to ensure that the corresponding
                    application traffic is preferentially forwarded.

Procedure
         Step 1 Create VLANs and configure interfaces so that DeviceB can communicate with
                Host1, Host2, and DeviceA.
                    # Create VLAN 10, VLAN 20, and VLAN 30 on DeviceB.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20 30

