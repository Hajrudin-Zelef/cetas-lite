---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-82
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10023, 10147]
sha256: 3a38a6a0390d5703d4f084b71d62427ebee37032969fe20e84405bb3a9a575ae
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                             ▪    Signature: When the device has an external identity certificate (local
                                  certificate issued by another CA) and sends a certificate enrollment
                                  request to the CA, the device uses the private key of this external
                                  identity certificate for signature.
                         –   Signature-based non-initial local certificate application using a CR
                             This method is applicable to the scenario where the device has an
                             external identity certificate and needs to apply for a local certificate.
                             When sending a certificate enrollment request to the CA, the device uses
                             the private key of the external identity certificate for signature.
                  5.     The CA creates a certificate based on the certificate enrollment request, which
                         is DeviceA.cer for DeviceA.
                  6.     The CA automatically uploads DeviceA.cer to the flash:/pki/public directory
                         on Device A.
                  7.     Manually import the certificate to the memory of DeviceA.

Certificate Update
                  When a PKI entity's certificate expires or the certificate key is disclosed, the PKI
                  entity must replace the certificate. In this case, a new application is required for
                  updating the certificate. Two methods are available for updating the local
                  certificate using CMPv2:

                  ●      Manual certificate update using a key update request (KUR)
                         A KUR, also called a certificate update request, is used to update the device's
                         existing certificate that is not expired and not revoked. During the update, the
                         device uses the existing certificate for identity authentication, and can use the
                         new or previous public key to update the local certificate.
                  ●      Automatic certificate update

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 186
Security Configuration
Security Configuration                                                                 10 PKI Configuration


                         To prevent service interruptions, you must apply for a new certificate before
                         the existing certificate expires. If manually updating the certificate, the user
                         may forget to do so. To avoid this problem, the device supports automatic
                         certificate update. With this function enabled, the device initiates a certificate
                         update request to the CMPv2 server when it detects that the certificate
                         automatic update time expires. The newly obtained certificate will replace
                         both the certificate file in the device storage and the certificate in the device
                         memory, without interrupting services.
                         In this mode, the local certificate requested using an IR or updated using a
                         KUR can be automatically updated.

10.7.2 Applying for and Updating a Local Certificate in Online
Mode Using CMPv2
Prerequisites
                  You have completed the preconfiguration for a certificate application. For details,
                  see 10.5 Preconfiguration for Certificate Application.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the file format in which the device stores the certificate.
                  pki file-format { der | pem }

         Step 3 Enter the CMP session view. If no CMP session exists, create one first.
                  pki cmp session session-name

                  A CMP session is locally available. It is not available to the CA and other devices.
         Step 4 Configure the PKI entity name used for CMPv2-based certificate application.
                  cmp-request entity entity-name

         Step 5 Configure a CA name for the CMP session.
                  cmp-request ca-name ca-name

                  The field sequence in the CA name must be the same as that in the CA certificate;
                  otherwise, the CMPv2 server considers the CA name invalid.
         Step 6 Configure the CMPv2 server URL.
                  cmp-request server url [ esc ] url-addr

         Step 7 Configure the RSA key pair used for CMPv2-based certificate application.
                  cmp-request rsa local-key-pair key-name [ regenerate [ key-bit ] ]

                  If the regenerate parameter is specified, the system generates a new RSA key pair
                  to apply for a new certificate and uses the new certificate and RSA key pair to
                  replace the previous ones during automatic certificate update. If regenerate is not
                  specified, the system uses the original RSA key pair during automatic certificate
                  update.
         Step 8 Optional: Configure the source IP address used to establish a TCP connection.
                  source { interface interface-type interface-number | ip-address }


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             187
Security Configuration
Security Configuration                                                                          10 PKI Configuration


                  If the interface parameter is specified, ensure that the interface is a Layer 3
                  interface and has an IP address configured.

         Step 9 Optional: Configure the encryption mode of CMPv2-based certificate application
                packets.
                  cmp-request integrity-algorithm { hmac-sha256 | hmac-sha1 }

                  When CMPv2 is used to apply for a certificate, packets need to be encrypted using
                  the hash algorithm. By default, the encryption algorithm SHA256 is used when
                  CMPv2 is used to apply for a certificate.

                          NOTE

                         For security purpose,you are not advised to use the weak security algorithm or weak
                         security protocols provided by this feature. If you need to use the weak security algorithm
                         or protocols, run the install feature-software WEAKEA command to install the weak
                         security algorithm or protocol feature package WEAKEA. By default, the device provides the
                         weak security algorithm or protocol feature package WEAKEA. For details about how to
                         install or uninstall the feature package, see "Upgrade Maintenance Configuration" in CLI
                         Configuration Guide > System Management Configuration.

        Step 10 Optional: Configure the certificate file used to verify the CA response signature.

                  Perform this step when applying for a local certificate in signature mode, in which
                  case the device needs to check whether the local certificate is issued by a valid CA.
                  Skip this step if the message authentication code is used.
                  cmp-request verification-cert cert-file-name

                  ●      If this command is configured and the CMPv2 server signs its certificate
                         response, the device uses the certificate (cert-file-name) configured using this
                         command to verify the CMPv2 server's response signature. The configured
                         certificate is a CA certificate used to verify a CA's identity.
                               NOTE

