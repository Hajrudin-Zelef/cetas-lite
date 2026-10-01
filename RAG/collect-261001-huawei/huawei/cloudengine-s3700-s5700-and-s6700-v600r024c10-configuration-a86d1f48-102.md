---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-102
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [12712, 12853]
sha256: 5b3614f7456c4de37f19d50c99c025b34298ebce548c73b11722ae44abea2e14
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  1.     The SSL client sends a ClientHello message to the SSL server. The session ID in
                         this message is set to the ID of the session to be resumed.
                  2.     If the server allows this session to be resumed, it replies with a ServerHello
                         message with the same session ID. After that, the client and server can use
                         the key and cipher suite of the resumed session without further negotiation.
                  3.     The client sends a ChangeCipherSpec message to notify the server that
                         subsequent messages will be encrypted and MACs will be computed based on
                         the key and cipher suite negotiated for the resumed session.
                  4.     The client computes a hash value over exchanged handshake messages, uses
                         the key and cipher suite negotiated for the resumed session to process the
                         hash value, and then sends a Finished message to the server so that it can
                         check whether the key and cipher suite are correct.
                  5.     Similarly, the server sends a ChangeCipherSpec message to notify the client
                         that subsequent messages will be encrypted and MACs will be computed
                         based on the key and cipher suite negotiated for the resumed session.
                  6.     The server computes a hash value over exchanged handshake messages, uses
                         the key and cipher suite negotiated for the resumed session to process the
                         hash value, and then sends a Finished message to the client so that it can
                         check whether the key and cipher suite are correct.

Data Transmission Process
                  After the handshake is complete, the client and server can exchange application
                  layer data. Data transmission is implemented using the SSL record protocol.
                  Figure 12-4 shows the data transmission process. The SSL record protocol
                  fragments the application data to be transmitted into manageable blocks,
                  compresses the data (optional), adds a MAC to the end of each block, encrypts
                  the block using an encryption algorithm, and then adds an SSL record header to
                  the encrypted block. In reverse order, the receiver decrypts, verifies, decompresses,
                  and reassembles the received message to obtain clear-text data.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           235
Security Configuration
Security Configuration                                                             12 SSL Configuration


                  Figure 12-4 Data transmission process




12.4 Default Settings for SSL
                  Table 12-1 describes the default settings for SSL.

                  Table 12-1 Default settings for SSL
                   Parameter                                  Default Setting

                   Cipher suite in the SSL policy             Not configured

                   Encryption algorithms supported by         Not configured
                   the cipher suite bound to the SSL
                   policy

                   SSL policy                                 Not configured

                   Minimum protocol version in the SSL        TLS1.2
                   policy

                   Certificate loaded for the SSL policy      No digital certificate is loaded for the
                                                              SSL policy.

                   CRL loaded for the SSL policy              No CRL is loaded for the SSL policy.

                   Trusted-CA file loaded for the SSL         No trusted-CA file is loaded for the
                   policy                                     SSL policy.

                   Cipher suite bound to the SSL policy       No cipher suite is bound to the SSL
                                                              policy.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            236
Security Configuration
Security Configuration                                                              12 SSL Configuration




12.5 Configuring SSL

12.5.1 (Optional) Configuring a Cipher Suite for an SSL Policy

Context
                  A cipher suite is a set of encryption algorithms used by a server and client during
                  SSL communication. In the initial SSL handshake process, the client sends a cipher
                  suite containing its supported algorithms to the server. The server then selects an
                  algorithm from the received cipher suite based on its own configurations. This
                  algorithm is used in subsequent communications.

                  Each encryption algorithm in the cipher suite contains the following information:

                  ●      Key exchange algorithm: determines how the server and client authenticate
                         each other. An asymmetric encryption algorithm is used to generate a session
                         key because it does not transmit significant data. Key exchange algorithms
                         include RSA, Diffie-Hellman, and ECDHE.
                  ●      Signature algorithm: used to sign the CA certificate. Signature algorithms
                         include RSA and DSS.
                  ●      Encryption algorithm: used to encrypt data to be transmitted. Both symmetric
                         and asymmetric encryption algorithms are available. However, asymmetric
                         encryption algorithms are rarely used as they consume too many resources
                         and are limited in terms of the length of data that can be transmitted.
                         Usually, an encryption algorithm name contains the key length and
                         encryption mode, such as GCM and CBC. Encryption algorithms include
                         AES_128, AES_256, AES_128_CBC, AES_256_CBC, AES_128_GCM,
                         AES_256_GCM, and ChaCha20-Poly1305.
                  ●      Message integrity check algorithm: used to check the integrity of messages.
                         Message integrity check algorithms include SHA, SHA256, and SHA384.

                  For example, a cipher suite supports the encryption algorithm
                  tls12_ck_rsa_aes_128_cbc_sha. This algorithm is based on the TLS protocol, and
                  uses RSA for key exchange, AES_128_CBC (with a key length of 128 bits and
                  encryption mode of CBC) for data encryption, and SHA for message integrity
                  check.


                  Table 12-2 Encryption algorithms supported by a cipher suite

                   TLS Version      Encryption Algorithm         Description
                                    Supported by a Cipher
                                    Suite

                   TLS1.1,          tls1_ck_rsa_with_aes_256     In this algorithm, RSA is used for key
                   TLS1.2, and      _sha                         exchange and signature, AES_256 is
                   TLS1.3                                        used for data encryption, and SHA is
                                                                 used for message integrity check.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             237
Security Configuration
Security Configuration                                                             12 SSL Configuration


                   TLS Version     Encryption Algorithm         Description
                                   Supported by a Cipher
                                   Suite

