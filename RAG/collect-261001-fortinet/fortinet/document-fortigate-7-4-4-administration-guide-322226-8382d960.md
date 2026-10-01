---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-322226-8382d960
title: "Generate a new certificate"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-322226-8382d960.md
source_anchor: ""
source_lines: [1, 26]
sha256: afddf432eecad9625ba95c88a5019a5b944d55f72f860f078b13729b28694603
---

# Generate a new certificate

# Generate a new certificate

The FortiGate can generate a certificate using a pre-loaded, self-signed CA certificate: *Fortinet_CA_SSL*, instead of generating a CSR and providing it to a CA for signing. It is recommended that a server certificate from a well-known and trusted CA is used.

###### To generate a new certificate:

1. 
                                                    Go to *System > Certificates* and select*Create/Import > Certificate* .
2. 
                                                    Click *Generate Certificate* .
3. 
                                                    Set *Certificate name* to the name of the certificate. This is what is referenced when using the certificate in FortiGate configurations.
4. 
                                                    Set the *Common name* (CN) for the certificate. The common name should match the FQDN or IP of the primary SSL-VPN interface.
5. 
                                                    Optionally, set the *Subject alternative name* .
6. 
                                                    Click *Download CA Certificate* to download the CA certificate so that it can be installed or imported to all the machines that need to trust this certificate.
7. 
                                                    Click *Create* .
8. 
                                                    After the certificate is created, click *Download Certificate* to download the certificate. Click*View Details* to review the certificate details.
9. 
                                                    Click *OK* .
