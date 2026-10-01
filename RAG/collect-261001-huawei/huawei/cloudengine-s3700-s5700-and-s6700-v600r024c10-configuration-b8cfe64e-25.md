---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-25
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [3194, 3355]
sha256: 69ca058a6e6dbc0982f9e4d2eacb696765fb8b23c8821eeec56fa5c912a1734c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    Check TCAM delivery failures.                       display system tcam fail-record [ slot
                                                                        slot-id ]




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                           56
QoS Configuration
QoS Configuration                                                                     7 Redirection Configuration


                    Operation                                        Command

                    Check the group indexes and rule                 display system tcam service brief
                    counts occupied by different services.           [ slot slot-id ]
                                                                     display system tcam service { cpcar
                                                                     slot slot-id | service-name slot slot-id
                                                                     [ chip chip-id ] }

                    Check the traffic policy application             display system tcam service traffic-
                    records.                                         policy

                    Check information about matched                  display system tcam match-rules slot
                    rules.                                           slot-id

                    Check statistics on packets that match           display traffic-policy statistics
                    a traffic policy.

                    Check matching fields and actions                display system tcam acl group-
                    supported by a traffic policy in each            information
                    view.

                    Check information about the resources display traffic-policy pre-state
                    occupied by the traffic policy to be
                    applied to determine whether the
                    traffic policy can be successfully applied
                    after the configuration is committed.




7.4 Example for Configuring Redirection to an Interface
Networking Requirements
                    In Figure 7-1, the server connects to the Internet through DeviceA, DeviceB, and
                    DeviceD. All traffic from the Internet needs to be redirected to DeviceC for filtering
                    to ensure the security of traffic to the server.

Figure 7-1 Network diagram of redirecting packets to an interface
     NOTE

   In this example, interface 1, interface 2, interface 3, and interface 4 represent 10GE 1/0/1, 10GE 1/0/2, 10GE
   1/0/3, and 10GE 1/0/4, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         57
QoS Configuration
QoS Configuration                                                             7 Redirection Configuration




Procedure
         Step 1 Create VLANs and configure interfaces to ensure Layer 2 connectivity.

                    # Create VLAN 200 and VLAN 300 on DeviceB.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 200 300

                    # Configure 10GE 1/0/1 on DeviceB as a trunk interface and add it to VLAN 200
                    and VLAN 300. Configure 10GE 1/0/2 and 10GE 1/0/3 on DeviceB as access
                    interfaces, and add 10GE 1/0/2 to VLAN 200 and 10GE 1/0/3 to VLAN 300.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 200 300
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type access
                    [DeviceB-10GE1/0/2] port default vlan 200
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] portswitch
                    [DeviceB-10GE1/0/3] port link-type access
                    [DeviceB-10GE1/0/3] port default vlan 300
                    [DeviceB-10GE1/0/3] quit

                    # Create VLAN 200 and VLAN 300 on DeviceA.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           58
QoS Configuration
QoS Configuration                                                               7 Redirection Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 200 300

                    # Configure 10GE 1/0/1, 10GE 1/0/2, 10GE 1/0/3, and 10GE 1/0/4 on DeviceA as
                    trunk interfaces and add them to VLAN 200 and VLAN 300. To prevent loops, add
                    10GE 1/0/3 and 10GE 1/0/4 to the same port isolation group and disable MAC
                    address learning on 10GE 1/0/4 to prevent MAC address flapping.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] portswitch
                    [DeviceA-10GE1/0/1] port link-type trunk
                    [DeviceA-10GE1/0/1] port trunk allow-pass vlan 200 300
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface 10ge 1/0/2
                    [DeviceA-10GE1/0/2] portswitch
                    [DeviceA-10GE1/0/2] port link-type trunk
                    [DeviceA-10GE1/0/2] port trunk allow-pass vlan 200 300
                    [DeviceA-10GE1/0/2] quit
                    [DeviceA] interface 10ge 1/0/3
                    [DeviceA-10GE1/0/3] portswitch
                    [DeviceA-10GE1/0/3] port link-type trunk
                    [DeviceA-10GE1/0/3] port trunk allow-pass vlan 200 300
                    [DeviceA-10GE1/0/3] port-isolate enable group 1
                    [DeviceA-10GE1/0/3] quit
                    [DeviceA] interface 10ge 1/0/4
                    [DeviceA-10GE1/0/4] portswitch
                    [DeviceA-10GE1/0/4] port link-type trunk
                    [DeviceA-10GE1/0/4] port trunk allow-pass vlan 200 300
                    [DeviceA-10GE1/0/4] port-isolate enable group 1
                    [DeviceA-10GE1/0/4] mac-address learning disable
                    [DeviceA-10GE1/0/4] quit

         Step 2 Configure redirection to an interface on DeviceA.
                    # Configure a traffic classifier. Configure a matching rule based on all data packets
                    in the traffic classifier c1.
                    [DeviceA] traffic classifier c1
                    [DeviceA-classifier-c1] if-match any
                    [DeviceA-classifier-c1] quit

                    # Configure a traffic behavior. Define redirection to a specified interface in the
                    traffic behavior b1.
                    [DeviceA] traffic behavior b1
                    [DeviceA-behavior-b1] redirect interface 10ge 1/0/3
                    [DeviceA-behavior-b1] quit

                    # Create a traffic policy p1, and bind the traffic classifier c1 and traffic behavior
                    b1 to the traffic policy.
                    [DeviceA] traffic policy p1
                    [DeviceA-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceA-trafficpolicy-p1] quit

                    # Apply the traffic policy to the inbound direction of 10GE 1/0/1.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceA-10GE1/0/1] quit

                    ----End

Verifying the Configuration
                    # Check the traffic classifier configuration.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                               59

