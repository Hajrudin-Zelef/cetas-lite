---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-46
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5956, 6101]
sha256: 90e8d522a316ab03f2f8b5ef4cf968c2681073b4b3bca5e0a14a8f044dd0b71c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  ----End

Verifying the Configuration
                  Run the display traffic behavior [ behavior-name ] command to check the traffic
                  behavior configuration.




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                    107
Security Configuration
Security Configuration                                                       8 MACsec Configuration




                                         8          MACsec Configuration


                  8.1 Overview of MACsec
                  8.2 Understanding MACsec
                  8.3 Configuration Precautions for MACsec
                  8.4 Default Settings for MACsec
                  8.5 Enabling the MACsec Function
                  8.6 Modifying MACsec Parameters
                  8.7 Verifying the Configuration
                  8.8 Example for Configuring Device-to-Device MACsec
                  8.9 Maintaining MACsec


8.1 Overview of MACsec

Definition
                  Media Access Control Security (MACsec) is a secure communication method based
                  on 802.1AE and 802.1X that ensures device-to-device security protection for
                  Ethernet links. It provides users with secure MAC-layer data sending and receiving
                  services through functions such as data encryption, integrity check, and replay
                  protection, ensuring the security of Ethernet frames.


Purpose
                  When data is transmitted in plaintext, there are many security risks, for example,
                  bank account information is stolen and tampered with, or malicious network
                  attacks occur. MACsec is a Layer 2 encryption technology that provides hop-by-
                  hop secure data transmission. It protects transmitted Ethernet frames to reduce
                  the risk of information leakage and malicious network attacks. MACsec is suitable
                  for meeting high requirements on data confidentiality.


Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                          108
Security Configuration
Security Configuration                                                         8 MACsec Configuration




8.2 Understanding MACsec
                  MACsec is used on a point-to-point link between interfaces of two devices. The
                  local and remote ends use security keys to encrypt and decrypt data packets. The
                  MACsec Key Agreement (MKA) protocol provides key negotiation as well as
                  establishment and management of secure channels. The MKA protocol defines a
                  complex key generation system to ensure the security of MACsec data
                  transmission. A Connectivity Association Key (CAK) is configured on a device but
                  not directly used to encrypt data packets. Data encryption is performed using a
                  Secure Association Key (SAK) derived from the CAK and other parameters. For
                  details about MACsec key derivation relationships, see the MACsec key system.

MACsec Implementation
                  The implementation of MACsec between two devices comprises three stages:
                  session negotiation, secure communication, and session keepalive.

                  Figure 8-1 MACsec implementation




                  1.     Session negotiation
                         After the connected interfaces of the two devices are enabled with MACsec
                         and configured with the same CAK, the two devices elect a key server through
                         the MKA protocol. The key server generates an SAK for encrypting data
                         packets based on the CAK and sends the SAK to the remote device.
                  2.     Secure communication
                         The sender uses the SAK to encrypt data packets, and the receiver uses the
                         SAK to decrypt data packets. Both devices can function as the sender or
                         receiver, and their communication is protected by MACsec.
                  3.     Session keepalive
                         The MKA protocol defines an MKA session keepalive timer that specifies the
                         timeout period of an MKA session. After MKA session negotiation is
                         successful, the two devices exchange MKA protocol packets to ensure that the

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             109
Security Configuration
Security Configuration                                                             8 MACsec Configuration


                         session is alive. The local device starts the timer after receiving MKA protocol
                         packets from the remote device.
                         –   If the local device receives subsequent MKA protocol packets within the
                             timeout period, it restarts the timer.
                         –   If the local device does not receive subsequent MKA protocol packets
                             within the timeout period, it considers the session insecure, deletes the
                             session, and performs MKA session negotiation again.

MACsec Key System
                  After the same CAK is configured on two devices, the key server generates an SAK
                  and sends it to the remote device, as shown in Figure 8-2. The related concepts
                  are as follows:
                  ●      CAK: The same CAK must be configured on the connected interfaces of the
                         two devices. The CAK is not directly used to encrypt data packets. Instead, it is
                         used together with a Connectivity Association Key Name (CKN) to derive the
                         SAK for data encryption.
                  ●      CKN: A CKN is the name of a CAK. The same CKN must be configured on the
                         connected interfaces of the two devices.
                  ●      SAK: An SAK is generated by the key server based on the CAK and CKN. The
                         SAK is used to encrypt and decrypt data packets.
                  ●      Key Encrypting Key (KEK): The two devices generate the same KEK based on
                         the same CAK and CKN. The KEK is used to encrypt and decrypt the SAK,
                         ensuring secure transmission of the SAK.
                  ●      Integrity Check Value (ICV): The sender calculates an ICV based on a packet
                         to be sent and adds the ICV to the end of the packet. The receiver uses the
                         same algorithm to calculate an ICV and compares it with the ICV carried in
                         the packet. If the two ICVs are the same, the packet is not modified and
                         therefore passes the verification. If they are different, the packet is considered
                         modified and is therefore discarded.
                  ●      ICV Key (ICK): The two devices generate the same ICK based on the same CAK
                         and CKN to calculate the ICV of MKA protocol packets. The ICK is used only
                         for the ICV calculation of MKA protocol packets. The ICV calculation of data
                         packets does not require an ICK.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              110
Security Configuration
Security Configuration                                                           8 MACsec Configuration


                  Figure 8-2 MACsec key derivation relationships




