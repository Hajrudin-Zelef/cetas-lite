---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-75
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [9077, 9218]
sha256: 47cf24f0df512f19ef96597aef36697450aab2d61bfc7901c71ac054398dd1f9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                                  pki ecc local-key-pair destroy When ECC key pairs are leaked,
                                  key-name                       damaged, lost, or unused, run this
                                                                 command in the system view to
                                                                 destroy a specified ECC key pair.
                                                                 After this command is executed,
                                                                 the system deletes the specified
                                                                 ECC key pair.

                   Search for     pki match-rsa-key                To check the RSA key pair
                   the            certificate-filename file-       corresponding to a certificate, run
                   RSA/SM2/       name                             this command in the system view
                   ECC key                                         to configure a device to search for
                   pair                                            the RSA key pair associated with a
                   associate                                       specific certificate.
                   d with a
                   specific       pki match-sm2-key                To check the SM2 key pair
                   certificate    certificate-filename file-       corresponding to a certificate, run
                                  name                             this command in the system view
                                                                   to configure a device to search for
                                                                   the SM2 key pair associated with a
                                                                   specific certificate.

                                  pki match-ecc-key                To check the ECC key pair
                                  certificate-filename file-       corresponding to a certificate, run
                                  name                             this command in the system view
                                                                   to configure a device to search for
                                                                   the ECC key pair associated with a
                                                                   specific certificate.



10.5.2 Configuring a PKI Entity
Context
                  Local certificates are signed and issued by the CA. A local certificate is a bundle of
                  a public key and a PKI entity. PKI entity information contains the identity
                  information of a PKI entity, based on which the CA identifies a certificate
                  applicant. As such, when applying for a local certificate, the PKI entity must send
                  PKI entity information to the CA.
                  The entity information includes the common name, FQDN, IP address, and email
                  address. The common name is mandatory, while others are optional. The
                  preceding information is contained in the certificate.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           170
Security Configuration
Security Configuration                                                               10 PKI Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the PKI entity view. If no PKI entity exists, create one first.
                  pki entity entity-name

         Step 3 Configure a common name for the PKI entity.
                  common-name common-name

         Step 4 Optional: Set other parameters for the PKI entity.

                  To uniquely identify an applicant, you can run the following optional commands
                  to configure the alias name for the PKI entity. If you do not configure alias names
                  for the PKI entities that have the same common name, these PKI entities will fail
                  to apply for a certificate.

                   To...                                       Run...

                   Configure an IP address for the PKI         ip-address { ipv4-address | ipv6-
                   entity                                      address | interface-type interface-
                                                               number [ ipv6 ] }
                   Configure an FQDN for the PKI entity        fqdn fqdn-name

                   Configure an email address for the PKI      email email-address
                   entity

                   Configure a country code for the PKI        country country-code
                   entity

                   Configure a locality name for the PKI       locality locality-name
                   entity

                   Configure a state name for the PKI          state state-name
                   entity

                   Configure an organization name for          organization organization-name
                   the PKI entity

                   Configure an organizational unit            organization-unit organization-unit-
                   name for the PKI entity                     name



         Step 5 Exit the PKI entity view.
                  quit

                  ----End


Verifying the Configuration
                  Run the display pki entity [ entity-name ] command to check the PKI entity
                  information.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            171
Security Configuration
Security Configuration                                                                          10 PKI Configuration


10.5.3 Downloading a CA Certificate
Context
                  When applying for a local certificate, the PKI entity sends the certificate
                  enrollment request to the CA. To improve transmission security, the PKI entity
                  must use the CA's public key to encrypt the certificate enrollment request.
                  Therefore, the PKI entity must download and obtain the CA certificate and then
                  obtain the CA's public key.
                  A CA certificate can be downloaded using the following methods, which differ in
                  the service types provided by the CA:
                  ●      Download the CA certificate from the CMPv2 server to the device storage
                         through CMPv2.
                  ●      Download the CA certificate from the server where the certificate is stored to
                         the device storage through LDAP.
                  ●      Obtain the CA certificate in out-of-band mode (for example, by disk or email)
                         and then upload it to the device storage.

Procedure
                  ●      Download a CA certificate through CMPv2.
                         For details about how to download a CA certificate through CMPv2, see
                         10.7.2 Applying for and Updating a Local Certificate in Online Mode
                         Using CMPv2.
                  ●      Download a CA certificate through LDAP.
                         system-view
                         pki ldap-server-template template-name attribute attr-value save-name dn dn-value
                  ●      Obtain a CA certificate in out-of-band mode.
                         After you obtain a CA certificate in out-of-band mode (for example, by disk or
                         email), manually upload it to the device storage. You can also download a CA
                         certificate through the administrator's PC and then upload it to the device
                         storage through FTP or SFTP. SFTP is recommended because it is more secure
                         than FTP.
                  ----End

