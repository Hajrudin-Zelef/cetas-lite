---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-99
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [12375, 12509]
sha256: 9b27ee723b9048db009dbfc90e35527a38f4d545bdf76f7379fa56150b53c7e4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Procedure
         Step 1 Check whether the network-side interface connected to the PPPoE server is a
                trusted interface.
                  If the interface is not trusted, the device discards PPPoE packets, and PPPoE users
                  cannot go online.
                  Enter the view of the network-side interface, and run the display this command
                  to check whether the pppoe uplink-port trusted command is configured on the
                  interface.
                  ●      If it is not, the network-side interface is an untrusted interface. Run the pppoe
                         uplink-port trusted command to configure it as a trusted interface.
                  ●      If it is, the network-side interface is a trusted interface. Go to the next step.
         Step 2 Check whether the policy for processing original information fields in user-side
                PPPoE packets meets service requirements.
                  In the system view and the view of each PPPoE user-side interface, run the display
                  this command to check whether the pppoe intermediate-agent information
                  policy command is configured.
                  ●      If the policy is configured in both the system view and the interface view, the
                         policy configured in the interface view takes effect preferentially.
                  ●      If the policy is configured in neither the system view nor the interface view,
                         the device uses the default policy replace for processing the packets.
                  Check whether the processing policy meets service requirements.
                  ●      If it does not, run the pppoe intermediate-agent information policy { drop |
                         keep | replace } command to configure a suitable processing policy.
                  ●      If it does, go to the next step.
         Step 3 Verify the formats of the information fields added to PPPoE packets.
                  Run the display pppoe intermediate-agent information format command to
                  check whether the formats of the information fields added to PPPoE packets are
                  the same as those required by the PPPoE server.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               227
Security Configuration
Security Configuration                                                      11 PPPoE+ Configuration


                  If they are not, run the pppoe intermediate-agent information [ vlan vlan-id ]
                  [ ce-vlan cevlan-id ] format { circuit-id | remote-id } { common | extend | user-
                  defined text } command to configure a suitable format.

                  ----End




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                         228
Security Configuration
Security Configuration                                                           12 SSL Configuration




                                           12                SSL Configuration


                  12.1 Overview of SSL
                  12.2 Configuration Precautions for SSL
                  12.3 Understanding SSL
                  12.4 Default Settings for SSL
                  12.5 Configuring SSL


12.1 Overview of SSL
Definition
                  Secure Sockets Layer (SSL) is a cryptographic protocol that provides
                  communication security over the Internet. SSL prevents eavesdropping on
                  communications between the client and server, and verifies the identities of
                  communicating parties, guaranteeing secure data transmission over the Internet.

Purpose
                  New Internet-based applications such as e-commerce and online banking have
                  increased due to the convenience they offer in people's daily lives. As these
                  applications require online transactions on the Internet, they pose higher
                  requirements on network communication security. However, the traditional
                  Hypertext Transfer Protocol (HTTP) for web-based communication does not have
                  a security mechanism. This protocol transmits data in clear text. It is unable to
                  verify the identities of communicating parties or prevent transmitted data from
                  being tampered with. As a result, HTTP does not meet the security requirements
                  of new applications. In this case, SSL, developed by Netscape, provides data
                  encryption, identity authentication, and message integrity verification
                  mechanisms, improving data transmission security. Specifically, SSL provides secure
                  connections for HTTP, significantly improving web-based communication security.

                  In addition, SSL can secure data transmission for any application layer protocol
                  based on Transmission Control Protocol (TCP) connections, because it functions
                  between the application and transport layers.

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                              229
Security Configuration
Security Configuration                                                             12 SSL Configuration




12.2 Configuration Precautions for SSL

12.3 Understanding SSL

12.3.1 Security Mechanisms
                  SSL provides the following security mechanisms:

                  ●      Identity Authentication: Digital signing is used to authenticate a server and
                         client that attempt to communicate with each other. Authenticating the
                         client's identity is optional.
                  ●      Data Encryption: The symmetric key algorithm is used to encrypt the
                         transmitted data.
                  ●      Message Integrity Check: The message authentication code (MAC)
                         algorithm is used to verify message integrity during message transmission.


Identity Authentication
                  The client verifies the identity of the SSL server to ensure that important
                  information is not being stolen. SSL uses digital signatures to authenticate the
                  communicating parties.

                  Digital signatures are implemented using an asymmetric key algorithm. Data
                  encrypted by a private key can only be decrypted by the matching public key.
                  Therefore, if the receiver successfully decrypts the data, the receiver considers the
                  sender valid. For example, Alice uses her own private key to encrypt a segment of
                  fixed information, and sends the encrypted information to Bob, who uses Alice's
                  public key to decrypt the information. If the decrypted segment is the same as the
                  fixed information, Bob considers that the information sender is Alice.

                  In this process, the authenticity of the sender's public key must be guaranteed.
                  Otherwise, unauthorized users may leverage a fake public key to communicate
                  with the receiver. SSL guarantees the authenticity of the sender's public key by
                  publishing the public key in a digital certificate.

                  A digital certificate (certificate for short) is a digitally signed file that includes
                  information about a public key and identity of its owner, proving the ownership of
                  a public key. Digital certificates are issued by a Certificate Authority (CA). When
                  issuing a digital certificate, the CA also provides a trusted-CA file to prove its own
                  identity and ensure the authenticity of the issued certificate.

