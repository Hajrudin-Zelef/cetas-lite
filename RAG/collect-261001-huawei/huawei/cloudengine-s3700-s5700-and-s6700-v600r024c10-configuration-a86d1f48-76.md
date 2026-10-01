---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-76
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [9219, 9363]
sha256: f2fe7ad385e60ea53bca4e3df90c385f9ad1b77ee8bee4a85ee73e508eaad5ab
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

10.5.4 Installing a CA Certificate
Context
                  After the CA certificate is downloaded, the device automatically stores the CA
                  certificate in flash:/pki/public.
                  After obtaining a CA certificate in out-of-band mode (for example, by disk or
                  email), you need to upload the CA certificate to the specified storage directory on
                  the device. Before installing a CA certificate, you need to upload the CA certificate
                  to flash:/pki/public.
                  After the CA certificate is saved to the specified directory, you also need to import
                  it to the device memory. After the device restarts, the system automatically loads
                  the certificate.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     172
Security Configuration
Security Configuration                                                                             10 PKI Configuration


                          NOTE

                         To prevent a certificate installation failure, ensure that the CA certificate file size does not
                         exceed 1 MB.
                         By default, the PKI realm named default exists. The default realm can be modified but
                         cannot be deleted.


Procedure
         Step 1 Optional: Enter the user view and download the CA certificate to the flash:/pki/
                public directory.

                  If you obtain the CA certificate in out-of-band mode and upload it to the device
                  storage through FTP or SFTP, perform this step. SFTP is recommended because it is
                  more secure than FTP.
                  cd pki
                  cd public/
                  ftp 172.16.104.110
                  Trying 172.16.104.110...
                  Press CTRL+K to abort
                  Connected to 172.16.104.110.
                  220 FTP service ready.
                  User(172.16.104.110:(none)):ftpuser
                  331 Password required for ftpuser
                  Enter password:
                  230 User logged in.
                  get ca.cer
                  200 Port command okay.
                  150 Opening ASCII mode data connection for temp1.c.
                  226 Transfer complete.
                  FTP: 4 byte(s) received in 8.190 second(s) .48byte(s)/sec.

         Step 2 Enter the system view.
                  system-view

         Step 3 Optional: Import the initial CA certificate to the default realm.
                  pki import-certificate default_ca realm default

                  The initial CA certificate has been saved to the NVRAM before the device is
                  delivered. To use the initial CA certificate, run this command to load the certificate
                  to the default realm. The initial CA certificate can be deleted. After it is deleted,
                  you can import other CA certificates to the default realm. However,
                  default_ca.cer is the name reserved for the initial CA certificate. An imported
                  certificate cannot be named default_ca.cer. To restore the initial CA certificate
                  after it is deleted, run this command to load the certificate to the default realm.

         Step 4 Create a PKI realm.
                  pki realm realm-name
                  quit

         Step 5 Import the CA certificate to the device memory.
                  pki import-certificate ca [ [ realm realm-name ] { der | pkcs12 | pem } ] filename file-name [ cert-name
                  cert-name ] [ no-check-hash-alg ] [ no-check-same-name ]

         Step 6 Optional: Set the number of days in advance you are notified that the CA
                certificate in the memory is about to expire.
                  pki set-certificate expire-prewarning day

         Step 7 Optional: Set the expiration check interval for the CA certificate in the memory.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           173
Security Configuration
Security Configuration                                                                10 PKI Configuration

                  pki certificate expiration-check interval interval-time

                  ----End

Verifying the Configuration
                  Run the display pki certificate ca [ realm realm-name | filename file-name ]
                  command to check the CA certificate that has been loaded on the device.


10.6 Applying for a Certificate in Offline Mode

10.6.1 Understanding Offline Certificate Application
                  Certificate application, also known as certificate enrollment, is a process in which
                  a PKI entity introduces itself to a CA, which then issues it a certificate. To obtain a
                  local certificate offline, you need to configure PKI entity information, configure an
                  RSA key pair, apply for the local certificate, and install the local certificate. Figure
                  10-11 shows the process of applying for a certificate in offline mode.

                  Figure 10-11 Process of applying for a certificate in offline mode




                  1.     Create a public/private key pair on DeviceA. The public key information is
                         required during certificate application.
                  2.     Configure entity information. When applying for a certificate, DeviceA must
                         provide the CA with information that can prove its identity. The entity
                         information represents identity information, including the common name,
                         fully qualified domain name (FQDN), IP address, and email address. The
                         common name is mandatory, while others are optional. After entity
                         information is configured, reference the entity information in the PKI realm.
                  3.     Generate a certificate enrollment request file. The generated certificate
                         enrollment request file is named PKI realm name.req and saved in the storage
                         device of DeviceA.
                  4.     After the certificate enrollment request file is generated, it can be sent to the
                         CA in out-of-band mode (for example, by disk or email).

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            174
Security Configuration
Security Configuration                                                                     10 PKI Configuration


                  5.     After approving the request, the CA creates a certificate based on the
                         certificate enrollment request file.
                  6.     After the certificate is generated, DeviceA's local certificate DeviceA.cer is
                         obtained in out-of-band mode (for example, by disk or email).
                  7.     Download DeviceA.cer to the flash:/pki/public directory on DeviceA.
                  8.     Import DeviceA.cer to the memory of Device A.

10.6.2 Applying for a Local Certificate in Offline Mode
Prerequisites
                  You have completed the preconfiguration for a local certificate application. For
                  details, see 10.5 Preconfiguration for Certificate Application.

Context
                  You can apply for a local certificate offline. To do this, you need to first generate a
                  certificate enrollment request file on the device, and then send the file to the CA
                  in out-of-band mode (for example, by disk or email).

                          NOTE

