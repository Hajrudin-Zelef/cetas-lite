---
id: collect-261001-general-networking/general-networking/manual-how-tos-self-signed-chain-html-c6fdd6a8-1
title: "manual-how-tos-self-signed-chain-html-c6fdd6a8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-self-signed-chain-html-c6fdd6a8.md
source_anchor: ""
source_lines: [1, 82]
sha256: 12ac803c2ec7d3b7dc83d58828e2fb42d3ca86d8c0b8b4cbc6269957e8339b09
---

# manual-how-tos-self-signed-chain-html-c6fdd6a8

Setup Self-Signed Certificate Chains
This how-to describes the process of creating self-signed certificate chains with OPNsense as PKI (Public Key Infrastructure).
The following section starts with a small overview of PKI, digital certificates, and trust.
Introduction to Public Key Infrastructure (PKI)
A Public Key Infrastructure (PKI) is a framework used to manage digital keys and certificates. PKI supports asymmetric encryption, where two related keys – a public key and a private key – are used together to secure communications and authenticate users.
OPNsense provides its own PKI in .
Components of PKI
- Certificate Authority (CA): The trusted entity that issues and verifies digital certificates.
- Registration Authority (RA): Assists the CA by verifying entities before certificates are issued.
- Digital Certificates: Bind public keys to individuals or entities, confirming identity.
- Certificate Revocation List (CRL): A list of certificates that have been revoked, used to manage trust.
When creating and using your own PKI, the above components depend on compliance. When you are a home user or manage a small company network, these can all be fulfilled by a single person or a small team. In large infrastructures, PKI management needs careful considerations.
How PKI Works
- The CA issues a digital certificate containing the entity’s public key.
- A sender uses the public key to encrypt information sent to the certificate holder.
- The certificate holder uses their private key to decrypt the information.
- Digital signatures can also be created with the private key and verified by others using the public key.
Together, PKI enables secure data exchange, authentication, and digital identity verification.
Public and Private Keys
In PKI, each entity has a unique pair of cryptographic keys:
- Public Key: The public key is shared openly. It is used to encrypt data that can only be decrypted by the matching private key. It is also used to verify digital signatures created by the private key. Public keys are included in digital certificates to enable secure exchanges.
- Private Key: The private key is kept secret by the owner. It decrypts data that was encrypted with the matching public key and creates digital signatures that can be verified by the public key. Protecting the private key is essential.
Understanding Trust
Trust in certificates refers to the confidence that the certificate’s issuer has verified identity and authenticity. This is less of a technical term, but trust between either real people or organizations who issue and verify certificates. The actual action of trust is when you install a CA certificate in your operating systems trust store. If the CA certificate has been pre-installed by a third party organization, you trust that organization by keeping it in your trust store.
Trust in CA-Issued Certificates
When using a CA-issued certificate, trust is based on identity verification, as the CA verifies the identity of certificate requesters; compliance, as CAs follow industry standards and are regularly audited; and widespread recognition, since CA root certificates are pre-trusted in most systems.
Trust in Self-Signed Certificates
For self-signed certificates, trust relies on the issuer’s security practices and if you want to trust their self-signed CA. Users must accept or install the certificate in their operating system trust store. Compromise of a self-signed certificate has higher risks, as there is likely no automatic revocation process. Using them for publicly accessible services introduces constraints due to lack of third-party validation and automatic trust.
Certificate Chains
Certificate chains enable verification of multiple levels of certificates, rather than relying on the validation of a single, isolated certificate.
Should you consider using self-signed certificate chains when free and widely recognized certificates are readily available?
- Self-signed certificates can be just as secure as publicly trusted certificates.
- They establish a chain of trust, where each party is well-known and directly controlled.
- You have full control of the PKI and the validity of each certificate. No external party can invalidate the certificates.
For a functional certificate chain, you will need at least a root CA and a leaf certificate (end-entity certificate). An intermediate CA is not required but adds a layer of security: if compromised, only the intermediate CA is affected, protecting the root CA key.
Attention
To enhance security:
- Store the root CA private key offline, such as on a USB or secure drive, kept in a safe location inaccessible to potential malware or unauthorized access.
- Keep multiple copies of the private key in different locations.
- The intermediate CA, generally with a shorter lifecycle, may be stored on a firewall host.
Creating a Certificate Chain
We will create a certificate chain using the following components:
- Root CA: Self-signed → Signs intermediate certificates
- Intermediate CA: Signed by the Root CA → Signs leaf certificates
- Leaf Certificate: Signed by the Intermediate CA → Server or user certificate
Important: Ensure all data is backed up before proceeding.
Here is an overview of the available GUI options:
| Options | Description | 
|---|---|
| Method | The operation to perform, such as creating a new CA, importing an existing one, or generating a Certificate Signing Request (CSR). | 
| Description | A brief identifier or label for the CA to distinguish it from others. | 
| Key | Configuration settings related to the cryptographic key associated with the CA. | 
| Key Type | The algorithm and size of the key pair (e.g., RSA 2048-bit, ECDSA). | 
| Digest Algorithm | The hash function used in the digital signature process (e.g., SHA-256). | 
| Issuer | The entity that signs the CA certificate; for a root CA, this is self-signed. | 
| Lifetime (days) | The validity period of the CA certificate, specified in days. | 
| General | General information fields for the CA’s Distinguished Name (DN). | 
| Country Code | The two-letter ISO code representing the country (e.g., ‘DE’ for Germany). | 
| State or Province | The full name of the state or province. | 
| City | The locality or city name. | 
| Organization | The legal name of the organization. | 
| Organizational Unit | A subdivision or department within the organization. | 
| Email Address | Contact email for the CA administrator. | 
| Common Name | The primary identifier for the CA, often its fully qualified domain name (FQDN). | 
| OCSP URI | The URL where the Online Certificate Status Protocol responder can be reached for certificate status checking. | 
| Output (PEM format) | The resulting certificate and key data in PEM format. | 
| Certificate Data | The CA’s public certificate in PEM format. | 
| Private Key Data | The CA’s private key in PEM format; handle with strict security measures. | 
| Serial for Next Certificate | The starting serial number for certificates issued by this CA; it increments with each issued certificate. | 
Note
This is where root or intermediate certificate authorities are created or imported.
| Options | Description | 
|---|---|
| Method | The action to perform, such as creating, importing, or signing a CSR for a leaf certificate. | 
| Description | A label or identifier for the certificate to distinguish it from others. | 
| Key | Settings related to the cryptographic key for the certificate. | 
| Type | The purpose of the certificate (e.g., server authentication, client authentication). | 
| Private Key Location | Specifies where the private key is stored or generated. | 
| Key Type | The algorithm and size of the key pair (e.g., RSA 2048-bit, ECDSA). | 
| Digest Algorithm | The hash function used in the digital signature process (e.g., SHA-256). | 
| Issuer | The CA that signs and issues the certificate. | 
| Lifetime (days) | The validity period of the certificate, specified in days. | 
