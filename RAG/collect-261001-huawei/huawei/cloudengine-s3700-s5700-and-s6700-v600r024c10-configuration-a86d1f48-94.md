---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-94
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [11636, 11802]
sha256: 0bd7b0b8557e1a55f3c6562da1837e158f57b96b0e0798c4a56617ec7a0a7c98
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

10.11.1 Deleting Certificates
Context
                  When a local certificate expires or a new certificate is required, delete the existing
                  local certificate from the device memory. The following table lists the commands
                  for deleting certificates from the device memory in the system view.

                  Table 10-7 Commands for deleting certificates

                   To...                                        Run...

                   Delete a local certificate from              pki delete-certificate local { realm realm-
                   the device memory                            name | filename file-name }
                   Delete a CA certificate from                 pki delete-certificate ca { realm realm-name |
                   the device memory                            filename file-name }

                   Delete an OCSP server                        pki delete-certificate ocsp { realm realm-
                   certificate from the device                  name | filename file-name }
                   memory




10.11.2 Clearing PKI Information
Context
                  PKI information cannot be restored after it is cleared. Exercise caution when you
                  run the following reset commands in the user view.

Procedure
                  ●      Clear the OCSP response cache.
                         reset pki ocsp response cache

                  ●      Clear the OCSP server down information recorded on the device.
                         reset pki ocsp server down-information [ url [ esc ] url-addr ]


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                  213
Security Configuration
Security Configuration                                                                       10 PKI Configuration


                  ●      Clear the CA certificates, CRLs, local certificates, and OCSP responder
                         certificates that have been imported into the memory.
                         reset pki global-ca

                               NOTE

                              This command will delete all CA certificates, CRLs, local certificates, and OCSP
                              responder certificates that have been imported into the device memory. Exercise
                              caution when running this command.

                  ----End

10.11.3 Moving Overwritten Files to the Recycle Bin
Context
                  Overwritten files are permanently deleted by default and cannot be restored. If
                  you want to restore overwritten files in case new files are unavailable, enable the
                  function of moving these files to the recycle bin.

                  This function applies only to the following scenarios:

                  ●      The existing CRL has been overwritten using the pki get-crl or pki import-crl
                         command.
                  ●      The existing certificates have been overwritten using the pki enroll-
                         certificate, pki create-certificate, pki export-certificate, pki export-
                         certificate default, or pki import-certificate peer command.
                  ●      The existing RSA key pair has been overwritten using the pki import rsa-key-
                         pair or pki export rsa-key-pair command.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable the function of moving overwritten files to the recycle bin.
                  pki recycle-bin enable

                  By default, overwritten files are deleted permanently.

                  ----End

10.11.4 Adding a PKI Realm to a Specified VPN
Context
                  A device needs to communicate with a server (for example, a CA server) to obtain
                  and verify certificates. When the server is in a VPN, add a PKI realm to the VPN.

Procedure
                  ●      PKI realm view
                         a.   Enter the system view.
                              system-view


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     214
Security Configuration
Security Configuration                                                                 10 PKI Configuration


                         b.   Enter the PKI realm view. If no PKI realm exists, create one first.
                              pki realm realm-name

                         c.   Adds a PKI realm to a specified VPN.
                              vpn-instance { vpn-instance-name }

                              By default, a PKI realm does not belong to any VPN.
                              The vpn-instance-name parameter is set using the ip vpn-instance
                              command.
                         d.   Exit the PKI entity view.
                              quit

                  ●      CMP session view
                         In the CMP session view, add a PKI realm to a specified VPN. This operation is
                         used only when CMPv2 is used to apply for and update a certificate in online
                         mode.
                         a.   Enter the system view.
                              system-view

                         b.   Enter the CMP session view. If no CMP session exists, create one first.
                              pki cmp session session-name

                              By default, no CMP session is created.
                              A CMP session is locally available. It is not available to the CA and other
                              devices.
                         c.   Adds a PKI realm to a specified VPN.
                              vpn-instance vpn-name { vpn-instance-name }
                              quit

                              By default, a PKI realm does not belong to any VPN.
                              The vpn-instance-name parameter is set using the ip vpn-instance
                              command.
                  ----End

10.11.5 Verifying and Viewing Initial Certificates

Context
                  Before a device is delivered, the CA and local certificates are loaded on the device
                  and stored in the NVRAM. These certificates cannot be deleted or modified. The
                  initial certificates, which function as the device identity, are imported to the
                  default realm to ensure the security of the device and external communication.
                  By default, the validity of initial certificates has been verified before a device is
                  delivered. Generally, you do not need to verify initial certificates.

Procedure
                  ●      Verify the validity of initial certificates.
                         pki validate-certificate device slot slot-id

                  ●      View the contents of initial certificates.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                            215
Security Configuration
Security Configuration                                                                   10 PKI Configuration

                         display pki certificate device slot slot-id

                  ----End


10.12 Troubleshooting PKI

10.12.1 Failed to Obtain a CA Certificate

Fault Symptom
                  A CA certificate has been manually obtained but does not exist in the device
                  storage. The cause is that the configuration for CA certificate download using
                  LDAP is incorrect.


