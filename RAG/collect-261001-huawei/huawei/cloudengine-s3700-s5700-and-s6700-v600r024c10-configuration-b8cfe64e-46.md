---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-46
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6318, 6483]
sha256: 9a57e76fdc9f416213b6e40484e1b4aca89b9f3e9012d746fd6a89da877afe38
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    The rates of traffic from tenants must be limited within proper ranges on DeviceB.
                    Table 9-7 lists the required CIR values for uplink traffic from tenants.

                    Table 9-7 CIR values for uplink traffic from tenants on DeviceB
                     Host                                         CIR (kbit/s)

                     Host1                                        2000

                     Host2                                        4000

                     Host3                                        8000




Procedure
         Step 1 Create VLANs and configure interfaces so that hosts can access the network
                through DeviceB.
                    # Create VLANs 10, 20, and 30 on DeviceA and add interfaces to the
                    corresponding VLANs.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 10 20 30
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] portswitch
                    [DeviceA-10GE1/0/1] port link-type access
                    [DeviceA-10GE1/0/1] port default vlan 10
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface 10ge 1/0/2
                    [DeviceA-10GE1/0/2] portswitch
                    [DeviceA-10GE1/0/2] port link-type access
                    [DeviceA-10GE1/0/2] port default vlan 20
                    [DeviceA-10GE1/0/2] quit
                    [DeviceA] interface 10ge 1/0/3
                    [DeviceA-10GE1/0/3] portswitch
                    [DeviceA-10GE1/0/3] port link-type access
                    [DeviceA-10GE1/0/3] port default vlan 30
                    [DeviceA-10GE1/0/3] quit
                    [DeviceA] interface 10ge 1/0/4
                    [DeviceA-10GE1/0/4] portswitch
                    [DeviceA-10GE1/0/4] port link-type trunk
                    [DeviceA-10GE1/0/4] port trunk allow-pass vlan 10 20 30
                    [DeviceA-10GE1/0/4] quit

                    # Create VLANs 10, 20, and 30 on DeviceB and add 10GE 1/0/1 to the VLANs.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20 30
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 10 20 30
                    [DeviceB-10GE1/0/1] quit

         Step 2 Configure traffic classifiers.
                    # On DeviceB, configure traffic classifiers c1, c2, and c3 to match service flows
                    from different hosts based on VLAN IDs.
                    [DeviceB] traffic classifier c1
                    [DeviceB-classifier-c1] if-match vlan 10


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                   115
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration

                    [DeviceB-classifier-c1] quit
                    [DeviceB] traffic classifier c2
                    [DeviceB-classifier-c2] if-match vlan 20
                    [DeviceB-classifier-c2] quit
                    [DeviceB] traffic classifier c3
                    [DeviceB-classifier-c3] if-match vlan 30
                    [DeviceB-classifier-c3] quit

         Step 3 Configure traffic behaviors and define traffic policing.

                    # On DeviceB, create traffic behaviors b1, b2, and b3 to police packets from
                    different hosts.
                    [DeviceB] traffic behavior b1
                    [DeviceB-behavior-b1] car cir 2000
                    [DeviceB-behavior-b1] statistics enable
                    [DeviceB-behavior-b1] quit
                    [DeviceB] traffic behavior b2
                    [DeviceB-behavior-b2] car cir 4000
                    [DeviceB-behavior-b2] statistics enable
                    [DeviceB-behavior-b2] quit
                    [DeviceB] traffic behavior b3
                    [DeviceB-behavior-b3] car cir 8000
                    [DeviceB-behavior-b3] statistics enable
                    [DeviceB-behavior-b3] quit

         Step 4 Configure a traffic policy and apply it to an inbound interface.

                    # On DeviceB, create traffic policy p1, bind configured traffic behaviors and traffic
                    classifiers to this traffic policy, and apply the traffic policy to the inbound direction
                    of 10GE 1/0/1 to police packets from hosts.
                    [DeviceB] traffic policy p1
                    [DeviceB-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceB-trafficpolicy-p1] classifier c2 behavior b2
                    [DeviceB-trafficpolicy-p1] classifier c3 behavior b3
                    [DeviceB-trafficpolicy-p1] quit
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceB-10GE1/0/1] quit

                    ----End

Verifying the Configuration
                    After the configuration is complete, check the traffic policing configuration on
                    DeviceB.

                    # Check the traffic classifier configuration.
                    [DeviceB] display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match vlan 10

                      Classifier: c2
                       Type: OR
                       Rule(s):
                         if-match vlan 20

                      Classifier: c3
                       Type: OR
                       Rule(s):
                         if-match vlan 30


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       116
QoS Configuration                                                      9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                      based Rate Limiting Configuration


                    Total classifier number is 3

                    # Check the traffic policy configuration.
                    [DeviceB] display traffic policy p1
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Committed Access Rate:
                           CIR 2000 (Kbps), PIR 2000 (Kbps), CBS 16000 (Bytes), PBS16000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard
                          Share car:
                           Car car1 share
                          Statistics: enable

                        Classifier: c2
                         Type: OR
                        Behavior: b2
                         Committed Access Rate:
                           CIR 4000 (Kbps), PIR 4000 (Kbps), CBS 32000 (Bytes), PBS 32000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard
                         Statistics: enable

