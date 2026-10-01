---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-137
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17741, 17914]
sha256: b0a20b624bdf1a37528f712dd958ff33b5f227a4df7c85c81a6a5208608272a6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  ●      If a protocol packet matches a GTSM policy, the device checks whether the
                         TTL value in the packet is within the specified range. If so, the device sends
                         the packet to the control plane. If not, the device drops the packet.
                  ●      If a protocol packet does not match the GTSM policy, the device sends the
                         protocol packet to the control plane if the default action of "pass" is used or
                         drops the packet if the action is set to "drop." For detailed configurations, see
                         20.4 (Optional) Configuring an Action to Process Packets That Do Not
                         Match a GTSM Policy.

                  Table 20-1 describes the protocols that support GTSM.

                  Table 20-1 Protocols supporting GTSM

                   Protocol       Matching                        Configuration Reference
                                  Granularity of a
                                  GTSM Policy

                   RIP            Public network                  IP Routing Configuration > RIP
                                  instance or VPN                 Configuration > Improving RIP Network
                                  instance                        Security > Configuring RIP GTSM

                   OSPF/          Public network                  IP Route Configuration > OSPF
                   OSPFv3         instance or VPN                 Configuration > Configuring OSPF GTSM
                                  instance                        IP Route Configuration > OSPFv3
                                                                  Configuration > Configuring OSPFv3 GTSM

                   BGP/           Public network                  IP Route Configuration > BGP Configuration
                   BGP4+          instance or VPN                 > Configuring BGP GTSM
                                  instance                        IP Route Configuration > BGP4+
                                  Peer IP address or              Configuration > Configuring BGP4+ GTSM
                                  peer group name




                  The following uses RIP as an example to describe how to enable RIP GTSM.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable RIP GTSM and configure a TTL value range and GTSM policy.
                  rip valid-ttl-hops valid-ttl-hops-value [ vpn-instance vpn-instance-name ]


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      327
Security Configuration
Security Configuration                                                                     20 GTSM Configuration


                          NOTE

                         The value range is [255 – valid-ttl-hops-value + 1, 255].

                  ----End

Verifying the Configuration
                  ●      Run the display gtsm statistics { slot-id | all } command to check GTSM
                         statistics.
                  ●      Run the reset gtsm statistics { slot-id | all } command to clear GTSM
                         statistics.


20.4 (Optional) Configuring an Action to Process
Packets That Do Not Match a GTSM Policy
Context
                  After GTSM is enabled, you can enable a device either to allow the packets that
                  do not match a GTSM policy to pass or to drop such packets.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable the device either to allow packets that do not match a GTSM policy to pass
                or to drop such packets.
                  gtsm default-action { drop | pass }

                          NOTE

                         ● The configured action takes effect only if GTSM is enabled.
                         ● Check whether the granularity specified in the GTSM policy is too fine. If an over fine-
                           grained policy is configured, you are not advised to set the action to drop, preventing a
                           large number of unmatched packets from being incorrectly dropped.
                         ● If the action is set to drop, you can determine whether to log the dropped packet
                           information. Enabling the log function facilitates fault locating.

                  ----End




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     328
Security Configuration
Security Configuration                                                             21 Security Risk Query Configuration




      21                    Security Risk Query Configuration


                  21.1 Querying Security Risks
                  21.2 Querying Security Configurations


21.1 Querying Security Risks
Context
                  Due to variations in security performance between protocols, some protocols may
                  pose security risks. You can run the display security risk command to check
                  security risks in the system, and eliminate the risks according to the recommended
                  solutions. For example, if SNMPv1 is configured, the display security risk
                  command output will prompt for the use of SNMPv3.


Procedure
         Step 1 Check security risks in the system and recommended solutions for the risks.
                  display security risk [ [ feature feature-name ] | [ level level-para ] | [ type type-para ] ] *

                  ----End


Example
                  Run the display security risk command to view security risks in the system.
                  <HUAWEI> display security risk
                  Risk level    : high
                  Feature name       : SNMP
                  Risk Type      : insecure-protocol
                  Risk information : SNMP V1/V2c is enabled.
                  Repair action : Disable SNMP V1/V2c and enable SNMP V3 only.

                  Risk Level    : medium
                  Feature Name       : FTPS
                  Risk Type      : insecure-protocol
                  Risk Information : FTP is not a secure protocol.
                  Repair Action : It is recommended to use SFTP



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          329
Security Configuration
Security Configuration                                                             21 Security Risk Query Configuration




21.2 Querying Security Configurations
Context
                  Due to variations in security performance between protocols, some protocols may
                  pose security risks. You can run the display security configuration command to
                  check security configurations in the system.

Procedure
         Step 1 Check current security configurations in the system.
                  display security configuration [ feature feature-name ]

                  ----End

Example
                  Run the display security configuration command to check security configurations
                  in the system. This example uses only some fields in the command output.
                  <HUAWEI> display security configuration
                  Feature Name : FTPS
                  Security Item : ftp security configuration
                  Item content : Ftp server is disabled.Ftp Ipv6 server is disabled.IP block feature is disabled.The FTP server
                  does not bind all interface.

                  Feature Name : TELNET
                  Security Item : telnet security configuration
                  Item content : The Telnet server function is used.The TELNET server bind all interface.

