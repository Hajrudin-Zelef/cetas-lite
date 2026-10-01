---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-30
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [1168, 1223]
sha256: 07f7bf077794070d6ff6520068d740609f1a4fc758f4f030f1b127021b2ebdc2
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

                                                                                                                                                                       used together to implement single sign-on (SSO)
                                                                                                           2.7.2.2 Compliant Access
                                                                                                                                                                       for network access, service access, and O&M
                                                                                                           Compliant access: Only compliant, authorized                operations, eliminating privilege escalation risks.
                                                                                                           terminals are allowed to access the campus
                                                                                                           network, and different network permissions are              Microsegmentat ion: is a net work securit y
                                                                                                           assigned to terminals based on the identity and             architecture approach that focuses on dividing
                                                                                                           compliance status. The terminal compliance                  the network into granular security zones (micro-
                                                                                                           mechanism includes host security check, secure              segments) and enforcing stringent access control
                                                                                                           access authentication (such as 802.1X and Portal            policies between these zones to limit attacker
                                                                                                           authentication), dumb terminal identification and           lateral movement, thereby enhancing overall
                                                                                                           anti-spoofing, and unauthorized access prevention.          security.
                                                                                                           RADIUS and TACACS+ authentication protocols are




                                  Figure 2-19 EDR security protection


The campus physical space is also secured.              intrusion analysis, special event monitoring,
Technologies such as Wi-Fi 7 and mmWave radar           and indoor unauthorized device detection. With
implement personnel and object sensing in areas,        technological maturity and industry development,
which can be used in scenarios such as personnel        more scenario-specific solutions will emerge.                                              Figure 2-21 Security groups


                                                                                                                                                                       Air interface signal scrambling technology draws
                                                                                                           2.7.2.3 Link Security
                                                                                                                                                                       on multi-user multiple-input multiple-output
                                                                                                           Encryption algorithms such as WPA2 and WPA3                 (MU-MIMO) and leverages idle antennas of APs to
                                                                                                           are used to encrypt the transmission over                   send extra electromagnetic wave noise, protecting
                                                                                                           the air interface, ensuring that data captured              the communication path of STAs. Within the
                                                                                                           by unauthorized users is encrypted and can                  location range of a target STA, the actual data is
                                                                                                           be parsed into valid information only after                 not affected by interference signals, ensuring that
                                                                                                           decryption. In wireless scenarios with high security        the data can be correctly demodulated. However,
                                                                                                           requirements, scrambling is performed for air               outside the location range of a target STA, the
                                                                                                           interface transmission, making it impossible for            actual data is affected by the interference signals,
                                                                                                           unauthorized users to obtain any valid information          meaning unauthorized users cannot demodulate
                                                                                                           over the air interface. In terms of wired links,            Wi-Fi signals. As shown in the following figure,
                                                                                                           MACsec is used to implement physical layer                  outside the location range of a target STA ,
                                                                                                           encryption.                                                 only meaningless noises can be captured when
                                                                                                                                                                       someone attempts to eavesdrop or capture
                                                                                                           Wi-Fi Protected Access (WPA), which includes                wireless packets.
                                 Figure 2-20 Spatial security detection                                    WPA , WPA2, and WPA3, is a set of security
                                                                                                           standards developed by the Wi-Fi Alliance to                When the target STA moves, the information can
                                                                                                           safeguard wireless network access, and supports             be quickly updated to ensure the accuracy of air
                                                                                                           authentication modes such as EAP-PEAP and EAP-              interface scrambling.
                                                                                                           TLS.




                                                   40                                                                                                             41
                                                                                                                                   2.7.3            Key Metrics

