---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-85
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10387, 10531]
sha256: 38d912f6ee259fb6336e4a018217899b498bcfd2dc51c1e8f5a3d3d4b1dfdad3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Context
                  Each certificate installed on the local device must be authenticated to ensure
                  validity before it is used. Certificate authentication verifies the issuing time, issuer
                  information, and certificate validity. The main point of this is to check the CA
                  signature on the certificate and ensure that the certificate is valid and not revoked.
                  To complete certificate authentication, the local device needs the following
                  information: CA certificate, CRL, local certificate and its private key, and certificate
                  authentication configuration.
                  The local device authenticates a local certificate as follows:
                  1.     Uses the CA certificate's public key to authenticate the CA's signature.
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

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      192
Security Configuration
Security Configuration                                                                                10 PKI Configuration


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

10.7.5 Example for Applying for and Updating a Local
Certificate in Online Mode Using CMPv2

Networking Requirements
                  On an enterprise network shown in Figure 10-14, DeviceA located at the network
                  border functions as the egress gateway, which uses CMPv2 to apply for a local
                  certificate in online mode for the first time from the CA server located on the
                  public network. The local certificate is automatically downloaded to DeviceA's
                  storage medium after being obtained. The local certificate is automatically
                  updated when 80% of the certificate validity period has elapsed.

                  Figure 10-14 Applying for a local certificate for a PKI entity in online mode
                          NOTE

                         In this example, Interface1 and Interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




                          NOTE

                         Ensure that there are reachable routes between devices before the configuration.


Configuration Roadmap
                  The configuration roadmap is as follows:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                           193
Security Configuration
Security Configuration                                                                  10 PKI Configuration


                  1.     Configure IP addresses for interfaces.
                  2.     Create an RSA key pair so that the local certificate application request
                         contains the public key.
                  3.     Configure a PKI entity and its related information to identify the PKI entity.
                  4.     Configure application and automatic update of certificates using CMPv2 and
                         use the message authentication code (MAC) to authenticate messages, so
                         that the device can automatically download the CA and local certificates.
                  5.     Install the local certificate to make it effective. That is, the device can use the
                         certificate to protect communication data.

Data Preparation
                  To complete the configuration, you need the following data:

                  ●      CA name: subject field of the CA certificate.
                  ●      Reference value and secret value of the MAC: obtained from the CMPv2
                         server.
                  ●      CA certificate of the CA server to be imported to the device.

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

                  Create a 3072-bit RSA key pair named rsa_cmp and allow it to be exported from
                  the device.
                  [DeviceA] pki rsa local-key-pair create rsa_cmp exportable
                   Info: The name of the new key-pair will be: rsa_cmp
                   The size of the public key ranges from 2048 to 4096.
                   Input the bits in the modulus:3072
                   Generating key-pairs...
                   Generating key-pairs finished

         Step 3 Configure a PKI entity to identify a certificate applicant.

                  Configure a PKI entity named user01.
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


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                194
Security Configuration
Security Configuration                                                                                   10 PKI Configuration


         Step 4 Configure a CMP session.
                  # Create a CMP session named cmp.
                  [DeviceA] pki cmp session cmp

