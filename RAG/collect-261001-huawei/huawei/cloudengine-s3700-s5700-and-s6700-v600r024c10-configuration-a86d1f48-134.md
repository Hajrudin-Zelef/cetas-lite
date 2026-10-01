---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-134
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [17258, 17399]
sha256: 99d54b88ede1d0421e5db3c8d67c77a14deb2f90e93f778c28b00bc1d44ee824
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

17.5 Configuring RA
Context
                  An RA system consists of the RA server and RA client. The administrator needs to
                  configure both the RA server and RA client to implement the RA function. This
                  section describes how to configure the RA client. For details about how to
                  configure the RA server, see the product documentation of related NMS devices.
                  Huawei provides two sets of certificates, each with one Huawei root certificate
                  and one Huawei level-2 CA certificate, for issuing IAK certificates. The two sets of
                  certificates are as follows:
                  ●      Huawei root certificate: Huawei Equipment CA; Huawei level-2 CA certificate:
                         Huawei Enterprise Network Product CA

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                316
Security Configuration
Security Configuration                                                17 Remote Attestation Configuration


                  ●      Huawei root certificate: Huawei RSA Equipment Root CA 2; Huawei level-2 CA
                         certificate: Huawei DataCom RSA Equipment CA 2

                  You are advised to load the two sets of certificates on the RA server. To download
                  Huawei root certificates and Huawei level-2 CA certificates, visit https://
                  support.huawei.com/additionalres/pki.

                  You can perform the following steps to view the IAK certificate currently used by
                  the RA client, so as to obtain the matching Huawei root certificate and Huawei
                  level-2 CA certificate.

                  1.     Run the display remote-attestation iak certificate slot slotid command in
                         the diagnostic view to check IAK certificate information. Copy the content
                         from -----BEGIN CERTIFICATE----- to -----END CERTIFICATE----- in the
                         command output, including -----BEGIN CERTIFICATE----- and -----END
                         CERTIFICATE-----, and save such information to the user PC in .pem format.
                  2.     Install OpenSSL on the user PC.
                  3.     On the CLI of the PC, run the openssl x509 -in filename -text -noout
                         command to check the CN in the Issuer field. filename: specifies that the
                         certificate path contains the certificate name.
                  4.     Download the Huawei level-2 CA certificate at https://support.huawei.com/
                         additionalres/pki based on the queried CN, and then download the Huawei
                         root certificate based on the Huawei level-2 CA certificate.


Prerequisites
                  Before configuring RA, complete the following tasks:

                  ●      Configure the device to communicate with the RA server, and configure the
                         RA function on the RA server.
                  ●      Obtain the baseline file, Huawei root certificate, and Huawei level-2 CA
                         certificate. If the RA client uses the LAK certificate, you also need to obtain
                         the CA certificate chain that issues the LAK certificate. Upload the obtained
                         file to the RA server.
                  ●      Configure NETCONF on the device so that the RA server can establish a
                         NETCONF session with the device.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the trusted management view and configure RA.
                  trustem

                  By default, no trusted management view is created. After you enter the trusted
                  management view, the device automatically enables RA.

         Step 3 (Optional) Return to the user view.
                  return

         Step 4 (Optional) Configure a password for the HTM.
                  set htm password { slot slot-id | all }


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            317
Security Configuration
Security Configuration                                                          17 Remote Attestation Configuration


                  By default, the HTM uses a random password generated before hardware delivery.
                  The device encrypts the configured password and stores it in a secure location to
                  improve the security of the RA function.

         Step 5 (Optional) Apply for a LAK certificate for the HTM.
                  1.     Create a PKI entity and enter the PKI entity view or enter the PKI entity view
                         directly.
                         system-view
                         pki entity entity-name

                         By default, no PKI entity is configured.
                  2.     Configure a common name for the PKI entity.
                         common-name common-name

                         By default, no common name is configured for a PKI entity.
                  3.     Exit the PKI entity view.
                         quit

                  4.     Create a PKI CMP session and enter the PKI CMP session view, or enter the
                         PKI CMP session view directly.
                         pki cmp session session-name

                         By default, no PKI CMP session is created.
                  5.     Specify the PKI entity name used by a device to apply for a certificate through
                         CMPv2.
                         cmp-request entity entity-name

                         entity-name must be the PKI entity name in step Step 5.1.
                  6.     Configure a CA name for the PKI CMP session.
                         cmp-request ca-name ca-name

                         The sequence of each field in the configured CA name must be the same as
                         that in the CA certificate. Otherwise, the CMPv2 server considers the CA name
                         incorrect.
                  7.     Configure a URL for the CMPv2 server.
                         cmp-request server url [ esc ] url-addr

                  8.     Set the authentication mode for a CMPv2-based initialization request (IR) to
                         signature.
                         cmp-request origin-authentication-method signature

                         By default, the authentication mode for a CMPv2-based IR is set to message
                         authentication code. In RA scenarios, the authentication mode for a CMPv2-
                         based IR must be set to signature.
                  9.     Exit the PKI CMP session view.
                         quit

                  10. Obtain the CA certificate chain that issues the LAK certificate as well as the
                      Huawei level-2 CA certificate of the IAK certificate, and upload them to the
                      flash:/pki/public directory of the device through SFTP.
                  11. Import the CA certificate chain of the LAK certificate as well as the Huawei
                      level-2 CA certificate of the IAK certificate.
                         pki import-certificate ca { der | pem } filename file-name


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   318
Security Configuration
Security Configuration                                                          17 Remote Attestation Configuration


                               NOTE

