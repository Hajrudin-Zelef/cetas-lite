---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-95
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright", "ethernet", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [11803, 11936]
sha256: 22cb8112bd9523499a46e7d3792254813584c5f321b9d2a498cb3eb82f1f1943
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Procedure
                  ●      Check whether the configuration for CA certificate download using LDAP is
                         correct. If not, correct the configuration by running the pki ldap-server-
                         template template-name attribute attr-value save-name dn dn-value
                         command.

                  ----End

10.12.2 Failed to Obtain a Local Certificate

Fault Symptom
                  ●      A local certificate has been manually obtained in offline mode but does not
                         exist in the device storage. The possible causes are as follows:
                         –     The PKI entity is incorrect.
                         –     The challenge password is incorrect or not configured.
                         –     The configuration for local certificate download using LDAP is incorrect.
                  ●      A local certificate has been manually obtained using CMPv2 but does not
                         exist in the device storage. The possible causes are as follows:
                         –     No CA certificate exists in the PKI realm.
                         –     The PKI entity is incorrectly configured or not configured.
                         –     The trusted CA name is incorrect or not configured.
                         –     The certificate enrollment server URL is incorrect or not configured.
                         –     The RSA key pair is not configured.
                         –     The source interface for TCP connection is incorrect.
                         –     The digest algorithm used for signing the certificate enrollment request is
                               incorrect.
                         –     The challenge password is incorrect or not configured.
                         –     The reference and secret values of the message authentication code are
                               incorrect or not configured.
                         –     The certificate for identity verification is incorrect.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             216
Security Configuration
Security Configuration                                                                                10 PKI Configuration


Procedure
                  ●      If the local certificate is obtained manually:
                         a.   Check whether the PKI entity is correct.

                              To view the configuration of a PKI entity in a PKI realm, run the display
                              pki entity command.

                              Modify the incorrect configurations, such as the country code.
                         b.   Check whether the challenge password is correct.

                              Check whether the CA server requires a challenge password. If so,
                              configure the challenge password to be the same as that of the CA server.
                              To set the challenge password, run the pki enroll-certificate command.
                         c.   Check whether the configuration for local certificate download using
                              LDAP is correct.

                              If not, correct the configuration by running the pki ldap-server-template
                              template-name attribute attr-value save-name dn dn-value command.
                  ●      If the local certificate is obtained using CMPv2:
                         a.   Check whether the CA certificate has been imported into the device
                              memory.

                              To view the CA certificate in the device memory, run the display pki
                              certificate command.

                              If no CA certificate exists, obtain a CA certificate and run the pki import-
                              certificate command to import the CA certificate into the device
                              memory.
                         b.   Check whether the PKI entity is correct.

                              To view the configuration of a PKI entity in a PKI realm, run the display
                              pki entity command.

                              Modify the incorrect configurations, such as the country code.
                         c.   Check whether the CA certificate application configuration is correct in
                              the PKI realm.

                              Run the display this command in the CMP session view.

                              The following is a sample of CA certificate application configuration:
                              pki cmp session cmp
                               cmp-request ca-name "C=cn,ST=beijing,L=SD,O=BB,OU=BB,CN=BB" //Configure a CA name.
                              The field sequence in the CA name must be the same as that in the actual CA certificate.
                               cmp-request authentication-cert local.cer //Configure a certificate for identity authentication
                              in a CMPv2 request, which is used to update a certificate or apply for a certificate for another
                              device.
                               cmp-request entity user01 //Specify the PKI entity to be used.
                               cmp-request server url http://10.3.0.1:8080 //Configure the URL of the CMPv2 server.
                               cmp-request rsa local-key-pair rsa regenerate //Specify the RSA key pair to be used.
                               cmp-request message-authentication-code 1234 %^%#ZodFBGH[^BkU2(~>[NRBv|#b>se|
                              @I7"'A,llG_B%^%# //Configure the reference value and secret value of the message
                              authentication code to be the same as those on the CA server.

                              Correct the configuration if it is incorrect.

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                217
Security Configuration
Security Configuration                                                      11 PPPoE+ Configuration




                                 11                  PPPoE+ Configuration


                  11.1 Overview of PPPoE+
                  11.2 PPPoE+ Fundamentals
                  11.3 Configuration Precautions for PPPoE+
                  11.4 Default Settings for PPPoE+
                  11.5 Configuring PPPoE+


11.1 Overview of PPPoE+
Definition
                  Point-to-Point Protocol over Ethernet plus (PPPoE+), also known as PPPoE
                  Intermediate Agent, is deployed on an access device (Device in the following
                  figure) that is located between user hosts and a Broadband Remote Access Server
                  (BRAS). The Device sends PPPoE Active Discovery (PAD) packets containing
                  information about the interfaces connected to user hosts (including the slot ID/
                  subcard ID/interface number, VLAN ID, and MAC address) to the PPPoE server. The
                  PPPoE server authenticates the user accounts together with the access interfaces,
                  preventing unauthorized account use.

                  Figure 11-1 PPPoE+ networking




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                        218
Security Configuration
Security Configuration                                                         11 PPPoE+ Configuration


