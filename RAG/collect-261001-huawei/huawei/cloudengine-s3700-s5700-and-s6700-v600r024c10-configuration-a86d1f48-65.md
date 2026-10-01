---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-65
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "license", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7793, 7947]
sha256: 10b5c7d0b1b1605b21e62f764fef95b286e7919d6f75252117cd51a1b1b336ad
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  DeviceB
                  #
                  sysname DeviceB
                  #
                  mac-security-profile name test2
                   mka keyserver priority 2
                   macsec mode normal
                  #
                  interface 10GE1/0/1
                   mka cak-mode static ckn f1c3b2a4d6d9a7c5b4e1ab56dc21ed79ac97be533671dcab2678ac55cf71aced cak
                  %^%#W5_!'~9]i>47d&X^Vro#S!z<4s+/N5\Ek*#27i_Wz-U3/"3tJM1.6++,nP+Z%^%#
                   mac-security-profile test2
                  #
                  return



8.9 Maintaining MACsec

Procedure
         Step 1 Run the reset mka statistics interface { interface-name | interface-type interface-
                number } command in the user view to clear MKA protocol packet statistics on a
                specified interface.
         Step 2 Run the reset macsec statistics interface { interface-name | interface-type
                interface-number } command in the user view to clear statistics about data
                packets protected by MACsec on a specified interface.

                  ----End




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    144
Security Configuration
Security Configuration                                     9 Service and Management Isolation Configuration




              9            Service and Management Isolation
                                             Configuration

Context
                  The purpose of service and management isolation is to minimize the attack
                  surface and ensure network security. The device supports such isolation.
                  Specifically, it complies with the three-layer and three-plane security isolation
                  mechanism of X.805. The three planes refer to:
                  ●      Management plane (or O&M plane): carries O&M data flows of the device.
                  ●      Control plane (or signaling plane): carries protocol interaction data flows of
                         the device.
                  ●      Service plane (or forwarding/user plane): carries information forwarding data
                         flows of the device.

                  After the three planes are isolated, if one plane is attacked, the operations and
                  security of other planes are not affected. For example, if the service plane is
                  subject to a DoS attack, the management plane will be unaffected. The
                  administrator can then log in to the management plane to eliminate the DoS
                  attacks. If, on the other hand, the planes are not isolated, such an attack would
                  cause the processing tasks on the service plane to further occupy resources such
                  as CPU and memory resources until the resources are fully exhausted. In such
                  cases, the administrator would be unable to manage the device.

                  Isolating the service plane from the management plane is to isolate service
                  interface traffic from the traffic of the management interface, and is implemented
                  as follows:

                  ●      Management data is prevented from being sent from service interfaces
                         (physical isolation). That is, users on the service network cannot use a device's
                         service interface to access the management network that is connected to the
                         of the device.
                  ●      Service interfaces and the are bound to different VPNs (logical isolation). As
                         such, data cannot be transmitted between any service interface and the .
                          NOTE

                         This function is supported only on the S6780-H, S6750-S, S6750-H, S6730-H-V2, S5755-H,
                         S6750E-S, S6730E-H-V2, S5755E-H, S5755-S, and S5732-H-V2.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  145
Security Configuration
Security Configuration                              9 Service and Management Isolation Configuration


Licensing Requirements
                  Service and management isolation is not under license control.


Hardware Requirements

                  Table 9-1 Hardware requirements

                   Series                        Models

                   S6780-H                       S6780-H4Z

                   S6750-H                       S6750-H36C/S6750-H48Y8C

                   S6750-S                       S6750-S16X8YZ/S6750-S16X10Y2CZ/S6750-
                                                 S24T16X8Y2CZ

                   S6750E-S                      S6750E-S16X10Y2CZ/S6750E-S16X10Y2CZ/
                                                 S6750E-S24T16X8Y2CZ

                   S6730-H-V2                    S6730-H48X6CZ-V2/S6730-H48X6CZ-TV2/S6730-
                                                 H28X6CZ-V2/S6730-H28X6CZ-TV2/S6730-
                                                 H48Y6C-V2/S6730-H48Y6C-TV2/S6730-
                                                 H6FX4Y2CZ-V2

                   S6730E-H-V2                   S6730E-H6FX4Y2CZ-V2

                   S5732-H-V2                    S5732-H48UM4Y2CZ-V2/S5732-H48UM4Y2CZ-
                                                 TV2/S5732-H24UM4Y2CZ-V2/S5732-
                                                 H24UM4Y2CZ-TV2/S5732-H44S4X6QZ-V2/S5732-
                                                 H44S4X6QZ-TV2/S5732-H24S4X6QZ-V2/S5732-
                                                 H24S4X6QZ-TV2

                   S5755-H                       S5755-H24N4Y-A/S5755-H24P4Y2CZ/S5755-
                                                 H24T4Y2CZ/S5755-H24U4Y2CZ/S5755-
                                                 H24UN4Y2CZ/S5755-H24UTM4X4Y2C/S5755-
                                                 H24UTM4X4Y2C-T/S5755-H48N4Y-A/S5755-
                                                 H48P4Y2CZ/S5755-H48T4Y2CZ/S5755-
                                                 H48U4Y2CZ/S5755-H48UN4Y2CZ/S5755-
                                                 H48UTM4X4Y2C/S5755-H48UTM4X4Y2C-T/
                                                 S5755-H48T4Y2CZ-B/S5755-H24HB2Y2CZ/S5755-
                                                 H24UM4Y2CZ/S5755-H24UM4Y2CZ-T/S5755-
                                                 H48UM4Y2CZ/S5755-H48UM4Y2CZ-T

                   S5755E-H                      S5755E-H24HB2Y2CZ

                   S5755-S                       S5755-S24P8J8YZ/S5755-S24P8Y/S5755-
                                                 S24T8J8YZ/S5755-S24T8Y/S5755-S24U8J8YZ/
                                                 S5755-S24U8Y/S5755-S48P8Y/S5755-S48P8YZ/
                                                 S5755-S48T8Y/S5755-S48T8YZ/S5755-S48U8Y/
                                                 S5755-S48U8YZ




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                         146
Security Configuration
Security Configuration                                        9 Service and Management Isolation Configuration


Feature Requirements
                  None

Procedure
                  ●      Enable service plane and management plane isolation to prevent
                         management data from being sent through the service plane.
                         system-view
                         undo management-plane isolate disable

