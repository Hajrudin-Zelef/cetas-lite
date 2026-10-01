---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-101
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["compute", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [12622, 12711]
sha256: 0226025dc02c63fc3f4b005eea803332a476226a1e10dae58f58d27171ab307e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  1.     The SSL client sends a ClientHello message carrying information, such as the
                         supported SSL versions and cipher suites, to the SSL server beginning a
                         handshake.
                  2.     The SSL server sends a ServerHello message to the SSL client with the
                         selected SSL version and cipher suite. If the SSL server allows the SSL client to
                         reuse the current session in subsequent communication, the SSL server
                         allocates a session ID to this session.
                  3.     The SSL server sends a digital certificate carrying its public key to the SSL
                         client so that the client can authenticate the server.
                  4.     (Optional) The SSL server requires the SSL client to provide a certificate for
                         identity authentication.
                  5.     The server sends a ServerHelloDone message, which means the SSL version
                         and cipher suite negotiation has finished and key information exchange can
                         begin.
                  6.     (Optional) The SSL client sends its own certificate to the SSL server.
                  7.     The SSL client verifies the authenticity of the SSL server certificate, encrypts
                         the randomly generated key using the public key in the certificate, and sends
                         the encrypted key to the SSL server.
                         The randomly generated key (premaster secret) cannot be directly used to
                         encrypt data or compute MACs. It is used to compute the symmetric key for
                         encryption and decryption and MACs for data integrity verification. The SSL
                         client and server use the premaster secret to compute the same master secret,
                         and then based on this, use a symmetric key algorithm to compute the
                         symmetric key and MACs. Therefore, the premaster secret plays a crucial role
                         in calculating the symmetric key and MACs.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               233
Security Configuration
Security Configuration                                                                  12 SSL Configuration


                  8.     (Optional) The SSL client sends its certificate to the server for identity
                         authentication.
                         The client computes a hash value for the master secret over exchanged
                         handshake messages, encrypts the hash value using its private key, and then
                         sends a CertificateVerify message to the server. The server computes a hash
                         value for the master secret over exchanged handshake messages, decrypts the
                         received CertificateVerify message using the public key in the client's
                         certificate, and compares the decrypted result with the computed hash value.
                         If the two values are the same, client authentication succeeds.
                  9.     The SSL client notifies the SSL server that subsequent packets will be
                         encrypted and MACs will be computed using the negotiated key (generated
                         based on the master secret) and cipher suite.
                  10. The client instructs the server to verify that the SSL negotiation has been
                      successful.
                         The client computes a hash value over exchanged handshake messages, uses
                         the negotiated key and cipher suite to process the hash value, and sends a
                         Finished message containing the hash value and MACs to the server. The
                         server computes a hash value in the same way, decrypts the received Finished
                         message, and verifies the hash value and MACs. If the verification succeeds,
                         the key and cipher suite negotiation is successful.
                              NOTE

                             Computing a hash value means that a hash algorithm (MD5 or SHA) is used to
                             convert an arbitrary-length message to a fixed-length message.
                  11. The SSL server notifies the SSL client that subsequent packets will be
                      encrypted and MACs will be computed using the negotiated key (generated
                      based on the master secret) and cipher suite.
                  12. The server instructs the client to verify that the SSL negotiation has been
                      successful.
                         The server computes a hash value over exchanged handshake messages, uses
                         the negotiated key and cipher suite to process the hash value, and sends a
                         Finished message containing the hash value and MACs to the client. The
                         client computes a hash value in the same way, decrypts the received Finished
                         message, and verifies the hash value and MACs. If the verification succeeds,
                         the key and cipher suite negotiation is successful.
                  Successful SSL negotiation indicates that the SSL server passes the SSL client's
                  identity authentication. The SSL server can decrypt the ClientKeyExchange
                  message to obtain the premaster secret only after it obtains the private key of the
                  client.
                  In the SSL handshake process, an asymmetric key algorithm is used to encrypt
                  keys and authenticate the identities of the communicating parties. It involves
                  heavy computation workload and consumes a lot of system resources. To simplify
                  the SSL handshake process, SSL allows resumed sessions, as shown in Figure 12-3.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                234
Security Configuration
Security Configuration                                                              12 SSL Configuration


                  Figure 12-3 SSL handshake process for resuming a session




