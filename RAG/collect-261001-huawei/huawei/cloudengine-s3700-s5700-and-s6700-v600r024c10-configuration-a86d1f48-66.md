---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-66
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7948, 8103]
sha256: 1d3aa601bbb957a26245e579a5b4f9443f77f267eea3c82d6670b51e240e90e7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         By default, service plane and management plane isolation is enabled.
                  ●      Bind service interfaces and the to different VPNs to ensure that they cannot
                         communicate with each other.
                         a.   Create and configure a VPN instance named management.
                              system-view
                              ip vpn-instance management
                               ipv4-family
                               quit

                              In this example, the instance name management represents the VPN of
                              the management plane. You can assign a specific instance name
                              according to your requirements.
                         b.   Create and configure a VPN instance named service.
                              ip vpn-instance service
                               ipv4-family
                               quit

                              In this example, the instance name service represents the VPN of the
                              service plane. You can assign a specific instance name according to your
                              requirements.
                         c.   Bind the to the management VPN instance and service interfaces to the
                              service VPN instance.
                               interface meth 0/0/0
                               ip binding vpn-instance management
                               quit
                              interface interface-type interface-number
                               ip binding vpn-instance service
                               quit

                               NOTE

                              ● You are advised to bind service interfaces to different VPN instances based on your
                                service requirements. That is, bind specific services only to necessary service
                                interfaces to implement fine-grained service isolation.
                              ● The loopback interface used for managing the device can also be bound to the
                                management VPN instance.
                              ● To isolate the service plane and management plane on an IPv6 network, first run
                                the ipv6-family [ unicast ] command in the VPN instance view to enable the IPv6
                                address family of the VPN instance.

                  ----End

Example
                  In Figure 9-1, the service network on the 192.168.20.0/24 network segment is
                  connected to interface 1 (service interface) of the device; the management
                  network on the 192.168.10.0/24 network segment is connected to interface 2 () of

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    147
Security Configuration
Security Configuration                                       9 Service and Management Isolation Configuration


                  the device. VPN-based logical isolation is configured on the device to isolate the
                  service plane and management plane. This is to prevent 192.168.20.0/24 from
                  communicating with 192.168.10.0/24, ultimately protecting the device against
                  attacks caused by address leakage on the .

                  Figure 9-1 Networking diagram of service and management isolation
                          NOTE

                         In this example, Interface 1 refers to 10GE1/0/1, and Interface 2 refers to MEth0/0/0.




                  Configuration script:
                  #
                  ip vpn-instance management
                   ipv4-family
                  #
                  ip vpn-instance service
                   ipv4-family
                  #
                  interface MEth0/0/0
                   ip binding vpn-instance management
                   ip address 192.168.10.1 255.255.255.0
                  #
                  interface 10GE1/0/1
                   ip binding vpn-instance service
                   ip address 192.168.20.1 255.255.255.0
                  #
                  return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       148
Security Configuration
Security Configuration                                                               10 PKI Configuration




                                              10                 PKI Configuration


                  10.1 Overview of PKI
                  10.2 Understanding PKI
                  10.3 Configuration Precautions for PKI
                  10.4 Default Settings for PKI
                  10.5 Preconfiguration for Certificate Application
                  10.6 Applying for a Certificate in Offline Mode
                  10.7 Applying for and Updating Certificates in Online Mode Using CMPv2
                  10.8 Configuring a Self-signed Certificate
                  10.9 Authenticating the Peer Entity's Certificate
                  10.10 Importing and Exporting a Certificate
                  10.11 Maintaining PKI
                  10.12 Troubleshooting PKI


10.1 Overview of PKI
Definition
                  Public Key Infrastructure (PKI) provides certificate management in compliance
                  with established standards. It uses public keys to provide security services for all
                  network applications. PKI is the core of information security and basis of e-
                  commerce.


Purpose
                  As the network and information technologies develop, e-commerce is widely used
                  and accepted. However, e-commerce has the following problems:

                  ●      The transaction parties cannot verify the identities of each other.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           149
Security Configuration
Security Configuration                                                                10 PKI Configuration


                  ●      Data may be eavesdropped and tampered with during transmission.
                         Information is not secure.
                  ●      No paper receipt is used in transaction, making arbitration difficult.
                  PKI uses public keys to implement identity verification, confidentiality, data
                  integrity, and non-repudiation of transactions. Therefore, PKI is widely used in
                  network communication and transactions, especially e-government and e-
                  commerce.

Benefits
                  ●      Certificate authentication allows users to authenticate network devices to
                         which they connect, ensuring that users connect to secure and legal networks.
                  ●      Cryptography protects data against eavesdropping so that data is securely
                         transmitted.
                  ●      Digital signature protects data against tampering so that data is securely
                         transmitted.
                  ●      PKI prevents unauthorized users from connecting to enterprise networks.
                  ●      PKI establishes secure connections between enterprise branches to ensure
                         data security.


10.2 Understanding PKI

10.2.1 Basic Concepts of PKI

