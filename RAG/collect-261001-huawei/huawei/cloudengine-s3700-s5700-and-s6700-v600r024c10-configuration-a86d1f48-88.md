---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-88
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10806, 10930]
sha256: 0ecb6e56a42b099467e22293853072083de0b76d2814bdc46031f84dd1e58621
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  ----End


10.9 Authenticating the Peer Entity's Certificate

10.9.1 Configuring Certificate Revocation Status Check
Context
                  During an attempt to establish a secure connection between two PKI entities, the
                  two entities must check each other's local certificate. If the certificate is invalid, a
                  secure connection cannot be established. However, the CA sometimes needs to
                  unbind a public key from related PKI entities due to factors such as user name
                  change, private key leak, or service interruption. A PKI entity must obtain the
                  certificate status of the peer entity in a timely manner to ensure communication
                  security.
                  The device provides the following certificate status check modes: CRL and None.
                  If multiple methods are configured, the system checks the certificate revocation
                  status in the configured sequence. The latter mode is used only when the current
                  mode is unavailable (for example, the server cannot be connected). When the
                  configured CRL mode is unavailable and None is configured, the certificate is
                  considered valid. For example, if the certificate-check crl ocsp none command is
                  configured, the device checks whether a certificate is valid using the CRL mode. If
                  the CRL mode is unavailable, the device considers the certificate valid.
                  You can select a mode for checking the certificate revocation status as required.
                  ●      CRL
                         The status of a certificate is determined by checking whether the certificate is
                         included in the CRL that is saved in the CRL database. After a certificate is
                         revoked, the SN of the certificate is recorded in the CRL. When a PKI entity
                         authenticates the local certificate of the peer entity, it searches for the
                         certificate in the CRL stored in local memory. If the certificate is included in
                         the CRL, it indicates that the certificate has been revoked. If no CRL is
                         available in local memory, the CRL needs to be downloaded and installed into
                         the local memory.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             199
Security Configuration
Security Configuration                                                                       10 PKI Configuration


                         Figure 10-15 Certificate status check in the CRL mode




                               NOTE

                              A PKI entity can download a CRL using an LDAPv3 template. A PKI entity must
                              frequently download the CRL to keep it up to date. By default, the device allocates
                              about 5 KB memory space for processing and caching the CRL. If the reserved memory
                              space is insufficient, new certificate revocation data cannot be imported. If you want
                              to import new certificate revocation data, delete the old data first.
                  ●      None
                         If no CRL is available for the PKI entity, or the local certificate status of the
                         PKI entity does not need to be checked, you can use the None mode, which
                         does not check whether the certificate is revoked.
                          NOTE

                         The global certificate revocation status can be checked in CRL mode or None mode.
                         The certificate revocation status in a PKI realm can be checked in CRL or None mode.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a PKI realm and enter its view, or enter the view of an existing PKI realm.
                  pki realm realm-name

                  By default, the system has a PKI realm named default. This realm can be modified
                  but cannot be deleted.
         Step 3 Configure the mode of checking certificate revocation status in the PKI realm.
                  certificate-check { { crl | ocsp } * [ none ] | none }

                  By default, the mode of checking certificate revocation status is not configured in
                  a PKI realm.
                  If this command is not configured, the global certificate revocation status check
                  mode is used. That is, the configuration of the pki certificate-check crl [ none ],
                  pki certificate-check none, or undo pki certificate-check command in the
                  system view takes effect.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    200
Security Configuration
Security Configuration                                                                     10 PKI Configuration


         Step 4 Select a mode to check peer certificate status according to the service types
                provided by the CA:
                  ●      Automatic CRL Update
                         a.   Return to the system view.
                              quit

                         b.   Configure the file format in which the device saves the CRL.
                              pki file-format { der |pem }

                                      NOTE

                                     By default, the device saves the CRL in PEM format.
                         c.   Optional: Enable the CRL mode for global certificate revocation status
                              check.
                              pki certificate-check crl [ none ]

                              By default, global certificate revocation status check in CRL mode is
                              enabled, and a certificate is regarded valid if the CRL mode is unavailable.
                         d.   Enter the PKI realm view and enable the automatic CRL update function.
                              pki realm realm-name
                              crl auto-update enable

                              By default, automatic CRL update is disabled.
                         e.   Set the interval for automatic CRL update.
                              crl update-period interval

                              By default, the automatic CRL update interval is 8 hours.
                         f.   Configure the attribute and identifier that the system uses to obtain CRLs
                              from the LDAP server.
                              crl ldap [ attribute attr-value ] dn dn-value

                              By default, no attribute or identifier is configured for the system to obtain
                              CRLs from the LDAP server.
                         g.   Configure automatic CRL update using an LDAPv3 template.
                              ldap-server-template template-name

