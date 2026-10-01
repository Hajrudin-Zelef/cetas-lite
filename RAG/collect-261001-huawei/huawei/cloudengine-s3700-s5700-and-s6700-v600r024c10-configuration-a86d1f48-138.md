---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-138
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17915, 18075]
sha256: 0d20b20b99b2843570ef1b8ae255a25128db96d5d341a4b3543ab36e320e86b7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  Feature Name : VTY
                  Security Item : Protocol used by VTY
                  Item content : SSH




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                   330
Security Configuration                                      22 Weak Password Dictionary Maintenance
Security Configuration                                                                 Configuration




                         22             Weak Password Dictionary
                                        Maintenance Configuration

                  22.1 Overview of Weak Password Dictionary Maintenance
                  22.2 Configuration Precautions for Weak Password Dictionary Maintenance
                  22.3 Configuring a Weak Password Dictionary


22.1 Overview of Weak Password Dictionary
Maintenance
Definition
                  Weak passwords are simple passwords that can be easily guessed or cracked
                  within a short period of time. You can configure a weak password dictionary
                  containing the prohibited passwords in advance and load the dictionary to the
                  device. In this case, when a new user is added or a user password is changed, the
                  system prevents the passwords in this dictionary from being used.
                  The weak password dictionary is stored as a text file and supports only the .txt
                  format. Each line stores a password. The following is an example of the weak
                  password dictionary pwd_dict.txt.
                  Abcd@123
                  Huawei@123
                  Aaabb@321
                  Raatr@321

                  After a weak password dictionary is loaded, the password settings in the device
                  login CLI, configuration file management, SNMP configuration, AAA configuration,
                  NAC configuration, and system's master key configuration will be affected. For
                  details, see the configuration precautions of these features and related command
                  reference.

Purpose
                  A device provides the weak password dictionary maintenance function to prevent
                  security problems caused by simple passwords.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             331
Security Configuration                                                22 Weak Password Dictionary Maintenance
Security Configuration                                                                           Configuration




22.2 Configuration Precautions for Weak Password
Dictionary Maintenance

22.3 Configuring a Weak Password Dictionary
Context
                  A device provides the function of weak password dictionary maintenance. With
                  this function, you can configure a weak password dictionary containing prohibited
                  passwords in advance and load the dictionary to the device.
                  The weak password dictionary to be loaded has been generated in .txt format and
                  uploaded to the device.

Procedure
                  ●      Load a weak password dictionary.
                         load security weak-password-dictionary filePath

                         Before the loading, ensure that the weak password dictionary in .txt format
                         has been uploaded to the device.
                  ●      Unload the weak password dictionary.
                         unload security weak-password-dictionary

                  ----End

Verifying the Configuration
                  ●      Run the display fips-mode command in the system view to check whether
                         the FIPS mode is enabled on the device.
                  ●      Run the display fips-mode algorithm self-check command in the system
                         view to check whether the algorithms provided by cryptographic modules of
                         the device comply with FIPS.
                  ●      Run the display fips-mode finite-state command in the system view to
                         check the historical records of FIPS mode status changes.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               332
Security Configuration
Security Configuration                                        23 Layer 2 Traffic Suppression Configuration




                         23                 Layer 2 Traffic Suppression
                                                          Configuration


                  23.1 Overview of Layer 2 Traffic Suppression
                  Layer 2 traffic suppression limits the bandwidth for forwarding unknown unknown
                  multicast traffic, which ensures bandwidth for forwarding unicast traffic.
                  23.2 Configuration Precautions for Layer 2 Traffic Suppression
                  23.3 Configuring BD-based Layer 2 Traffic Suppression
                  Disabling an EVC Layer 2 sub-interface from broadcasting packets to other EVC
                  Layer 2 sub-interfaces in the same bridge domain (BD) prevents devices from
                  being attacked and helps improve network security.
                  23.4 Configuring VLAN-based Layer 2 Traffic Suppression
                  VLAN-based traffic suppression prevents excessive traffic from burdening the
                  network.


23.1 Overview of Layer 2 Traffic Suppression
                  Layer 2 traffic suppression limits the bandwidth for forwarding unknown unknown
                  multicast traffic, which ensures bandwidth for forwarding unicast traffic.

Definition
                  Traffic on a Layer 2 network is classified into the following types:
                  ●      Unicast traffic: consists of unicast packets whose destination MAC addresses
                         are in the MAC table. The device forwards these packets in unicast mode
                         according to the information in the MAC address table.
                  ●      Unknown unicast traffic: consists of unicast packets whose destination MAC
                         addresses are not in the MAC address table. The device broadcasts these
                         packets.
                  ●      Multicast traffic: consists of packets that use multicast addresses as
                         destination MAC addresses. The device broadcasts these packets.
                  ●      Unknown multicast traffic: consists of packets that use multicast addresses as
                         destination MAC or IP addresses and do not have matching multicast

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            333
Security Configuration
Security Configuration                                            23 Layer 2 Traffic Suppression Configuration


                         forwarding entries. After Layer 2 multicast is enabled, the device broadcasts
                         these packets.


Purpose
                  Layer 2 traffic suppression limits the bandwidth for forwarding unknown multicast
                  traffic, which ensures bandwidth for forwarding unicast traffic.


23.2 Configuration Precautions for Layer 2 Traffic
Suppression

23.3 Configuring BD-based Layer 2 Traffic Suppression
                  Disabling an EVC Layer 2 sub-interface from broadcasting packets to other EVC
                  Layer 2 sub-interfaces in the same bridge domain (BD) prevents devices from
                  being attacked and helps improve network security.


Context
                  When an EVC Layer 2 sub-interface in a BD receives unknown multicast packets, it
                  broadcasts the packets to other EVC Layer 2 sub-interfaces in the BD.

