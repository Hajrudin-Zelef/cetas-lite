---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-84
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10262, 10386]
sha256: abf325dd73fd639c2339d628a09516f4d272c8d2b8e0536f7ee3275de42eb430
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                              When the device requests to update the key with the CMPv2 server, it
                              also applies for a new local certificate.
                              After this command is configured, the system first checks the CMP
                              session configuration to determine whether the certificate update
                              condition is met. If the condition is not met, an error message is
                              displayed. If the condition is met, the system initiates a certificate update
                              request according to the configuration. The certificate obtained is saved
                              in a file to the device storage without being imported to the memory.
                  ●      Automatic update
                         a.   Enter the CMP session view and configure the certificate for identity
                              authentication in a CMPv2 request.
                              pki cmp session session-name
                              cmp-request authentication-cert cert-name

                              This certificate is the local certificate that the CA has issued to the device
                              and needs to be replaced by a new local certificate.
                         b.   Enable CMPv2-based automatic certificate update.
                              certificate auto-update enable

                         c.   Configure the time when the local certificate is updated automatically.
                              The value is expressed as the percentage of the certificate validity period.
                              certificate update expire-time valid-percent
                              quit

                              By default, the certificate update time is 50% of the certificate validity
                              period.
                              After this command is configured, the system initiates a certificate update
                              request and determines whether to create an RSA key pair according to
                              the cmp-request rsa local-key-pair command configuration when
                              finding that the automatic certificate update time reaches the value
                              specified by valid-percent. After the new certificate is obtained, the
                              system replaces the previous certificate and RSA key pair with the new
                              ones.

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                190
Security Configuration
Security Configuration                                                                               10 PKI Configuration


Verifying the Configuration
                  Run the display pki cmp statistics [ session session-name ] command to check
                  CMP session statistics.

10.7.3 Installing a Local Certificate

Context
                  The system automatically saves the local certificate that has been obtained online
                  using CMPv2 to flash:/pki/public.

                  After the local certificate is stored in the preceding directory, you need to
                  manually import the certificate into the device memory. After the device restarts,
                  the system automatically loads the certificate.

                          NOTE

                         To prevent a certificate installation failure, ensure that the local certificate file size does not
                         exceed 1 MB.
                         The initial local certificate that identifies a Huawei device provides certificate authentication
                         for user login services of the device by default.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Import the local certificate into the device memory.
                  ●      When applying for a local certificate from the CA based on the PKI entity
                         information and RSA key pair created in the local PKI entity, you only need to
                         run the following command to import the local certificate into the device
                         memory. The RSA key pair has been imported into the device memory by
                         default during the creation of the RSA key pair.
                         pki import-certificate local [ [ realm realm-name ] { der | pkcs12 | pem } ] filename file-name
                         [ cert-name cert-name ] [ no-check-hash-alg ] [ no-check-same-name ]

                  ●      If you need to use the key pair generated by another PKI entity and the
                         certificate of another PKI entity, you need to import both the certificate and
                         key pair. A key pair file can either be included in a certificate file or exist
                         independently of the certificate file. The methods of importing a key pair file
                         vary accordingly.
                         Select a method of importing a key pair file accordingly.
                         –     If the key pair file is included in a certificate file:
                               pki import rsa-key-pair keyname { pem | pkcs12 } filename [ exportable ] [ password
                               password ]

                         –     If the key pair file exists independently of the certificate file:
                               # Import the certificate file.
                               pki import-certificate local [ [ realm realm-name ] { der | pkcs12 | pem } ] filename file-
                               name [ cert-name cert-name ] [ no-check-hash-alg ] [ no-check-same-name ]

                               # Import the key pair file.
                               pki import rsa-key-pair keyname exclude-cert { pem | pkcs12 } filename [ exportable ]
                               [ password password ]


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                 191
Security Configuration
Security Configuration                                                                        10 PKI Configuration


                         NOTE

                  If no certificate format is specified, the system automatically detects the certificate format and
                  imports the certificate.

         Step 3 Optional: Set the number of days in advance you are notified that the local
                certificate in the memory is about to expire.
                  pki set-certificate expire-prewarning day

         Step 4 Optional: Set the expiration check interval for the local certificate in the memory.
                  pki certificate expiration-check interval interval-time

                  ----End

Verifying the Configuration
                  Run the display pki certificate local [ realm realm-name | filename filename ]
                  command to check the local certificate that has been loaded on the device.

10.7.4 Checking the Validity of a Certificate
Prerequisites
                  The CA certificate and local certificate have been installed on the device.

