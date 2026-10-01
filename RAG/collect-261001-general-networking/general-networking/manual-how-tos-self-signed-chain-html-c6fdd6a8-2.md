---
id: collect-261001-general-networking/general-networking/manual-how-tos-self-signed-chain-html-c6fdd6a8-2
title: "manual-how-tos-self-signed-chain-html-c6fdd6a8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-self-signed-chain-html-c6fdd6a8.md
source_anchor: ""
source_lines: [83, 204]
sha256: b8a8eaeb85aae46bcf6731f6225f4a4297adb8bc2af7452ecbdfa1103fd46526
---

# manual-how-tos-self-signed-chain-html-c6fdd6a8

| General | General information fields for the certificate’s Distinguished Name (DN). | 
| Country Code | The two-letter ISO code representing the country. | 
| State or Province | The full name of the state or province. | 
| City | The locality or city name. | 
| Organization | The legal name of the organization. | 
| Organizational Unit | A subdivision or department within the organization. | 
| Email Address | Contact email for the certificate subject. | 
| Common Name | The primary identifier for the certificate, often the FQDN of the server or the name of the individual. | 
| OCSP URI | The URL where the OCSP responder can be reached for certificate status checking. | 
| Alternative Names | Additional identifiers for the certificate subject. | 
| DNS Domain Names | Alternative domain names covered by the certificate. | 
| IP Addresses | IP addresses associated with the certificate subject. | 
| URIs | Uniform Resource Identifiers associated with the certificate subject. | 
| Email Addresses | Additional email addresses associated with the certificate subject. | 
| Output (PEM format) | The resulting certificate and key data in PEM format. | 
| Certificate Data | The leaf certificate’s public certificate in PEM format. | 
| Private Key Data | The leaf certificate’s private key in PEM format; handle with strict security measures. | 
| Certificate Signing Request | A CSR containing the public key and Distinguished Name to be signed by a CA. | 
Note
This is where leaf certificates signed by intermediate or root certificate authorities are created or imported.
| Options | Description | 
|---|---|
| Method | The operation to perform, such as creating a new CRL or importing an existing one. | 
| CA Reference | The CA associated with this CRL. | 
| Description | A brief description or identifier for this CRL. | 
| CRL Data | Contains the actual CRL in PEM format, listing revoked certificates and their statuses. | 
| Serial | The unique serial number identifying the certificate to be revoked. | 
| Lifetime (days) | Specifies how long the CRL is valid before it needs to be regenerated. | 
| Revocations per type | Specifies reasons for revocation as categories (e.g., Unspecified, Key Compromise, CA Compromise). | 
| Unspecified | Indicates a certificate was revoked without a specified reason. | 
| Key Compromise | Indicates the private key associated with the certificate was compromised. | 
| CA Compromise | Indicates that the issuing CA’s private key was compromised. | 
| Affiliation Changed | Indicates that the certificate subject’s affiliation with the organization has changed. | 
| Superseded | Indicates that the certificate was replaced by another. | 
| Cessation of Operation | Indicates that the entity associated with the certificate no longer operates. | 
| Certificate Hold | Temporarily revokes a certificate, which may be reinstated later. | 
Note
This is where certificate revocation lists for certificates signed by intermediate or root certificate authorities are created or imported.
| Options | Description | 
|---|---|
| Store intermediate | Allow local defined intermediate certificate authorities to be used in the local trust store. We advise to only store root certificates to prevent cross signed ones causing breakage when included but expired later in the chain. | 
| Store CRL’s | Store all configured CRL’s in the default trust store. | 
| Auto fetch CRL’s | Schedule an hourly job to download CRLs using the defined Distributionpoints in the CAs deployed in our trust store. | 
| Enable legacy | Enable Legacy Providers. | 
| Enable | Enable custom constraints. | 
| CipherString | Sets the ciphersuite list for TLSv1.2 and below. | 
| Ciphersuites | Sets the available ciphersuites for TLSv1.3. | 
| SignatureAlgorithms | Sets the available SignatureAlgorithms. | 
| DHGroups / Curves | Limit the default set of built-in curves to be used when using the standard openssl configuration. | 
| MinProtocol | Sets the minimum supported SSL or TLS version. | 
| MinProtocol (DTLS) | Sets the minimum supported DTLS version. When configuring MinProtocol and leaving this empty, DTLS will be disabled. | 
Step 1: Creating the Root CA
Go to
Press + to create a new authority, it will become your root certificate authority.
| Options | Description | 
|---|---|
| Method | Create an internal Certificate Authority | 
| Description | Root CA (or a custom description) | 
| Key |  | 
| Key Type | RSA-2048 (or higher) | 
| Digest Algorithm | SHA256 (or higher) | 
| Issuer | self-signed (root CA is always self-signed) | 
| Lifetime (days) | 3650 (after this expires the root CA, all its issued intermediate CAs and their issued leaf certificates must be recreated) | 
| General |  | 
| Country Code | Netherlands (your country) | 
| State or Province | Zuid-Holland (your state or empty) | 
| City | Middelharnis (your city or empty) | 
| Organization | Deciso B.V. (your organization name or empty) | 
| Organizational Unit | IT (your organizational unit or leave empty) | 
| Email Address | info@example.com (your email address, it is best practice to use a real existing one) | 
| Common Name | root-ca (or a custom name) | 
| OCSP URI | leave empty | 
Press Save and the root CA has been created. The private and public key are saved on the OPNsense.
Step 2: Issuing the Intermediate CA
Go to
Press + to create a new authority, it will become your intermediate certificate authority.
| Options | Description | 
|---|---|
| Method | Create an internal Certificate Authority | 
| Description | Intermediate CA (or a custom description) | 
| Key |  | 
| Key Type | RSA-2048 (or higher) | 
| Digest Algorithm | SHA256 (or higher) | 
| Issuer | Root CA (The intermediate CA is always signed by the Root CA) | 
| Lifetime (days) | 1095 (after this expires the intermediate CA and all its issued leaf certificates must be recreated) | 
| General |  | 
| Country Code | Netherlands (your country) | 
| State or Province | Zuid-Holland (your state or empty) | 
| City | Middelharnis (your city or empty) | 
| Organization | Deciso B.V. (your organization name or empty) | 
| Organizational Unit | IT (your organizational unit or leave empty) | 
| Email Address | info@example.com (your email address, it is best practice to use a real existing one) | 
| Common Name | intermediate-ca (or a custom name) | 
| OCSP URI | leave empty | 
Press Save and the intermediate CA has been created. The private and public key are saved on the OPNsense.
Step 3: Issuing a Leaf Certificate
Go to
Press + to create a new certificate, it will become your leaf certificate (end-entity certificate). It can be used on a server, user, or both; depending on the type.
| Options | Description | 
|---|---|
| Method | Create an internal Certificate | 
| Description | leaf-certificate.example.com (or a custom description, like user or server name) | 
| Key |  | 
| Type | Server Certificate (or client certificate for a user) | 
| Private Key Location | Save on this firewall | 
| Key Type | RSA-2048 (or higher) | 
| Digest Algorithm | SHA256 (or higher) | 
| Issuer | Intermediate CA | 
| Lifetime (days) | 365 (after this expires the leaf certificate must be recreated) | 
| General |  | 
| Country Code | Netherlands (your country) | 
| State or Province | Zuid-Holland (your state or empty) | 
| City | Middelharnis (your city or empty) | 
| Organization | Deciso B.V. (your organization name or empty) | 
| Organizational Unit | IT (your organizational unit or leave empty) | 
| Email Address | info@example.com (your email address, it is best practice to use a real existing one) | 
| Common Name | leaf-certificate.example.com (or a custom name) | 
| OCSP URI | leave empty | 
| Alternative Names |  | 
| DNS Domain Names | leaf-certificate.example.com (or a custom name) | 
Press Save and the leaf certificate has been created. The private and public key are saved on the OPNsense.
Note
