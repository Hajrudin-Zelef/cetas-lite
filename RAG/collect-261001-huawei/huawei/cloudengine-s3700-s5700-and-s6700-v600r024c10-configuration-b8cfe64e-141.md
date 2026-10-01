---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-141
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15681, 15822]
sha256: 6351251addbe7735c6ea1965bdd8d0eaabc688cf02787526ebd6c8fc2252cf20
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                                       13 Experience Assurance Configuration


Networking Requirements
                    In Figure 13-3, packets sent from Host1 and Host2 to DeviceB are identified by
                    different VLAN IDs. The VLAN ID for Host1 is 10, and that for Host2 is 20. To
                    record application traffic statistics on terminals, enable application identification
                    and application identification statistics collection on DeviceB.

                         NOTE

                        Only the S6730E-H-V2, S6730-H-V2 and S5732-H-V2 series support this example.


                    Figure 13-3 Network diagram for configuring experience assurance
                         NOTE

                        In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                        10GE1/0/3, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1. Create VLANs and configure interfaces so that DeviceB can communicate with
                    Host1, Host2, and DeviceA.

                    2. Enable application identification on inbound interfaces.

                    3. Enable the application identification statistics collection function.

                    4. Configure NetStream on DeviceB because application identification statistics
                    collection depends on NetStream.

Procedure
         Step 1 Create VLANs and configure interfaces so that DeviceB can communicate with
                Host1, Host2, and DeviceA.

                    # Create VLAN 10, VLAN 20, and VLAN 30 on DeviceB.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     288
QoS Configuration
QoS Configuration                                                             13 Experience Assurance Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20 30

                    # Configure 10GE1/0/1 as a trunk interface and add it to VLAN 30. Configure
                    10GE1/0/2 and 10GE1/0/3 as access interfaces and add them to VLAN 10 and
                    VLAN 20, respectively.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 30
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type access
                    [DeviceB-10GE1/0/2] port default vlan 10
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] portswitch
                    [DeviceB-10GE1/0/3] port link-type access
                    [DeviceB-10GE1/0/3] port default vlan 20
                    [DeviceB-10GE1/0/3] quit

                    # Create VLANIF 10, VLANIF 20, and VLANIF 30, and configure IP addresses for
                    them.
                    [DeviceB] interface vlanif 10
                    [DeviceB-Vlanif10] ip address 192.168.10.1 24
                    [DeviceB-Vlanif10] quit
                    [DeviceB] interface vlanif 20
                    [DeviceB-Vlanif20] ip address 192.168.20.1 24
                    [DeviceB-Vlanif20] quit
                    [DeviceB] interface vlanif 30
                    [DeviceB-Vlanif30] ip address 192.168.100.1 24
                    [DeviceB-Vlanif30] quit

         Step 2 Enable application identification on inbound interfaces.
                    # Enable application identification on access-side interfaces of DeviceB.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] sa enable
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] sa enable
                    [DeviceB-10GE1/0/3] quit

         Step 3 Enable the application identification statistics collection function.
                    # Enable application identification statistics collection on DeviceB.
                    [DeviceB] sa application-statistic enable

         Step 4 Configure NetStream on which traffic statistics collection depends.
                    # Switch to the NetStream resource mode on DeviceB.
                    [DeviceB] assign forward enp netstream enable

                    # Configure NetStream globally on DeviceB. The recommended configurations are
                    as follows:
                    [DeviceB] netstream timeout ip active 300
                    [DeviceB] netstream timeout vxlan inner-ip active 300
                    [DeviceB] netstream timeout ip inactive 180
                    [DeviceB] netstream timeout vxlan inner-ip inactive 180
                    [DeviceB] netstream timeout ip tcp-session
                    [DeviceB] netstream timeout vxlan inner-ip tcp-session


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   289
QoS Configuration
QoS Configuration                                                           13 Experience Assurance Configuration


                    # Configure a flexible flow statistics template on DeviceB, including 5-tuple,
                    statistics collection, and sampling rate information. The recommended
                    configurations are as follows. You are advised to perform the configurations
                    according to the recommended configurations. Otherwise, traffic statistics may be
                    incorrect.
                    [DeviceB] netstream record sac_v4 ip
                    [DeviceB-netstream-record-ipv4-sac_v4] collect counter bytes
                    [DeviceB-netstream-record-ipv4-sac_v4] collect counter packets
                    [DeviceB-netstream-record-ipv4-sac_v4] collect interface sampler-info
                    [DeviceB-netstream-record-ipv4-sac_v4] match ip destination-address
                    [DeviceB-netstream-record-ipv4-sac_v4] match ip destination-port
                    [DeviceB-netstream-record-ipv4-sac_v4] match ip protocol
                    [DeviceB-netstream-record-ipv4-sac_v4] match ip source-address
                    [DeviceB-netstream-record-ipv4-sac_v4] match ip source-port
                    [DeviceB-netstream-record-ipv4-sac_v4] quit
                    [DeviceB] netstream record sac_vxlan vxlan inner-ip
                    [DeviceB-netstream-record-vxlan-sac_vxlan] collect counter bytes
                    [DeviceB-netstream-record-vxlan-sac_vxlan] collect counter packets
                    [DeviceB-netstream-record-vxlan-sac_vxlan] collect interface sampler-info
                    [DeviceB-netstream-record-vxlan-sac_vxlan] match inner-ip destination-address
                    [DeviceB-netstream-record-vxlan-sac_vxlan] match inner-ip destination-port
                    [DeviceB-netstream-record-vxlan-sac_vxlan] match inner-ip protocol
                    [DeviceB-netstream-record-vxlan-sac_vxlan] match inner-ip source-address
                    [DeviceB-netstream-record-vxlan-sac_vxlan] match inner-ip source-port
                    [DeviceB-netstream-record-vxlan-sac_vxlan] quit

