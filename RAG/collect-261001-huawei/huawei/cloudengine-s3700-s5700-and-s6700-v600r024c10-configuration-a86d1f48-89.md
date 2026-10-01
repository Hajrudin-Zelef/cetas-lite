---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-89
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10931, 11069]
sha256: 9b7d824129272c84df143616bca22dd2cb38ba09f28f0be5709b5a00dd9b6df7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                              The ldap-server-template template-name command references an LDAP
                              server template in a PKI realm. For details about how to configure an
                              LDAP server template, see Configuring an LDAP Server Connected to
                              the Device under "AAA Configuration" in CLI Configuration Guide > User
                              Access and Authentication Configuration.
                         h.   Optional: Update the CRL immediately and import the CRL to the device
                              memory.
                              The CRL can be automatically updated only when the time for automatic
                              CRL update arrives. To update the CRL immediately, you can use
                              immediate CRL update function.
                              pki get-crl realm realm-name
                              pki import-crl realm realm-name filename file-name

                              After the CRL is updated immediately, the new CRL replaces the old CRL
                              in the device storage, and is automatically imported to the device
                              memory to replace the old one.
                         i.   Optional: Enable CRL expiration check, and configure the CRL expiration
                              check period and the prewarning percentage of remaining CRL validity
                              period.
                              pki crl expiration-check enable
                              pki crl expiration-check interval interval-time
                              pki crl expiration-check prewarning remain-percent percent


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                             201
Security Configuration
Security Configuration                                                                          10 PKI Configuration


                  ●      Manual CRL update
                         a.   Return to the system view.
                              quit

                         b.   Configure the file format in which the device saves the CRL.
                              pki file-format { der | pem }

                         c.   Configure CRL download through LDAP.
                              pki ldap-server-template template-name attribute attr-value save-name dn dn-value

                         d.   Import the CRL to the device memory.
                              pki import-crl [ realm realm-name ] filename file-name

                  ----End


Verifying the Configuration
                  ●      Run the display pki crl [ realm realm-name | filename file-name ] command
                         to check the CRL content on the device.
                  ●      Run the display pki certificate ocsp [ realm realm-name | filename file-
                         name ] command to check the OCSP server certificate loaded on the device.

Follow-up Procedure
                  If a CRL expires or is not used, run the pki delete-crl { realm realm-name |
                  filename filename } command to delete the CRL from the device memory.

10.9.2 Configuring Certificate Attribute-based Filtering to
Implement Access Control

Context
                  Certificate attribute-based filtering is a method of certificate authentication.
                  Configuring a certificate attribute-based access control policy allows only the
                  certificates that meet specific attribute conditions to pass authentication, thereby
                  implementing refined access control.

                  A certificate attribute-based access control policy consists of one or more
                  certificate attribute groups, certificate attribute conditions, and certificate
                  attribute-based control rules. The certificate attribute conditions are defined in the
                  certificate attribute group. When a certificate matches all certificate attribute
                  conditions, the configured certificate attribute-based control rule determines
                  whether to permit the certificate.

                  The following table lists the certificate attribute conditions.

                   Certificate Attribute Condition                    Description

                   Start time and end time of the                     Start time and end time of the validity
                   certificate validity period                        period for the PKI entity's local
                                                                      certificate




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                    202
Security Configuration
Security Configuration                                                                        10 PKI Configuration


                   Certificate Attribute Condition                      Description

                   FQDN                                                 FQDN of the PKI entity's local
                                                                        certificate
                                                                        An FQDN consists of a host name and
                                                                        a domain name, for example,
                                                                        www.example.com.

                   Certificate IP address                               IP address of the PKI entity's local
                                                                        certificate

                   Certificate issuer name                              Name of the issuer of the PKI entity's
                                                                        local certificate

                   Certificate subject name                             Subject name of the PKI entity's local
                                                                        certificate




                  A certificate attribute-based control rule contains two actions: permit and deny.
                  This rule determines whether to permit or block certificates that meet the
                  certificate attribute conditions.

                  The matching principles for access control through certificate attribute-based
                  filtering are as follows:

                  ●      If a service has a specified certificate attribute-based access control policy, the
                         specified certificate attribute-based access control policy is used. Otherwise,
                         the default certificate attribute-based access control policy is used. By default,
                         the action in the default policy is permit. That is, the certificate is allowed to
                         pass authentication.
                  ●      If a certificate attribute-based access control policy contains multiple control
                         rules with the OR relationship, the action in the policy is taken as long as the
                         certificate to be authenticated matches one rule.
                  ●      If a certificate attribute group contains multiple certificate attribute conditions
                         with the AND relationship, the action in the corresponding control rule is
                         taken when the certificate to be authenticated matches all conditions.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the default certificate attribute-based access control policy.
                  pki certificate access-control-policy default { deny | permit }

                  By default, the action in the default policy is permit. That is, the certificate is
                  allowed to pass authentication.

