---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-79
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [9644, 9758]
sha256: d11e94491af2acede52c6865481aaad7ecc1c8328f67a37343947f6df94e6577
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         To authenticate a certificate, a PKI entity must obtain the public key of the
                         certificate-issuing CA from the CA's certificate, allowing it to check the CA's
                         signature on the certificate. An upper-level CA authenticates the certificates of
                         lower-level CAs. The authentication is performed along the certificate chain,
                         and terminated at a trustpoint (the root CA holding a self-signed certificate or
                         a subordinate CA trusted by the PKI entity).
                         PKI entities that share the same root or subordinate CA and possess CA
                         certificates can authenticate certificates of each other (peer certificates).
                         In short, certificate chain authentication starts at the target certificate (PKI
                         entity's certificate to be authenticated) and ends at a trustpoint.
                         Authentication of a peer certificate chain generally ends at the first trusted
                         certificate or CA.
                  2.     Checks whether the certificate has expired.
                  3.     Checks whether the certificate has been revoked in CRL, OCSP, or None mode.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Check the validity of the CA or local certificate.
                  pki validate-certificate { ca | local } { realm realm-name | filename file-name }

                  The pki validate-certificate ca command allows you to check the validity of only
                  the root CA certificate, not subordinate CA certificates. When multiple CA
                  certificates are imported on a device, you can use only the pki validate-
                  certificate local command to check the validity of subordinate certificates.

                  ----End

10.6.6 Example for Applying for a Local Certificate for a PKI
Entity in Offline Mode
Networking Requirements
                  As shown in Figure 10-12, the device applies for a local certificate from the CA
                  server on the public network in offline mode.
                          NOTE

                         In this example, Interface1 and Interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                           180
Security Configuration
Security Configuration                                                                10 PKI Configuration


                  Figure 10-12 Applying for a local certificate for a PKI entity in offline mode




Configuration Roadmap
                  The configuration roadmap is as follows:
                  1.     Create an RSA key pair so that the local certificate application request
                         contains the public key.
                  2.     Configure a PKI entity and its related information to identify the PKI entity.
                  3.     Configure local certificate application for the PKI entity in offline mode and
                         generate a local certificate request file.
                  4.     Send the local certificate request file in out-of-band mode and download the
                         local certificate.
                  5.     Install the local certificate so that the device can protect communication data.

Procedure
         Step 1 Configure IP addresses for interfaces.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10GE1/0/1] undo portswitch
                  [DeviceA-10GE1/0/1] ip address 10.2.0.2 24
                  [DeviceA-10GE1/0/1] quit
                  [DeviceA] interface 10ge 1/0/2
                  [DeviceA-10GE1/0/2] undo portswitch
                  [DeviceA-10GE1/0/2] ip address 10.1.0.2 24
                  [DeviceA-10GE1/0/2] quit

         Step 2 Create an RSA key pair.
                  Create a 3072-bit RSA key pair named rsakey and allow it to be exported from
                  the device.
                  [DeviceA] pki rsa local-key-pair create rsakey exportable
                   Info: The name of the new key-pair will be: rsakey
                   The size of the public key ranges from 2048 to 4096.
                   Input the bits in the modulus:3072
                   Generating key-pairs...
                   Generating key-pairs finished

         Step 3 Configure a PKI entity to identify a certificate applicant.
                  Configure a PKI entity named user01.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                181
Security Configuration
Security Configuration                                                                            10 PKI Configuration

                  [DeviceA] pki entity user01
                  [DeviceA-pki-entity-user01] common-name hello
                  [DeviceA-pki-entity-user01] country cn
                  [DeviceA-pki-entity-user01] email user@test.abc.com
                  [DeviceA-pki-entity-user01] fqdn test.abc.com
                  [DeviceA-pki-entity-user01] ip-address 10.2.0.2
                  [DeviceA-pki-entity-user01] state jiangsu
                  [DeviceA-pki-entity-user01] organization huawei
                  [DeviceA-pki-entity-user01] organization-unit info
                  [DeviceA-pki-entity-user01] quit

         Step 4 Apply for a local certificate for the PKI entity in offline mode.
                  [DeviceA] pki realm abc
                  [DeviceA-pki-realm-abc] entity user01
                  [DeviceA-pki-realm-abc] rsa local-key-pair rsakey
                  [DeviceA-pki-realm-abc] quit
                  [DeviceA] pki enroll-certificate realm abc pkcs10 filename cer_req
                   Info: Creating certificate request file...
                   Info: Create certificate request file successfully.


