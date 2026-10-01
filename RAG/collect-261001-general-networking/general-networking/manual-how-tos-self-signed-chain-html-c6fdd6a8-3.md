---
id: collect-261001-general-networking/general-networking/manual-how-tos-self-signed-chain-html-c6fdd6a8-3
title: "manual-how-tos-self-signed-chain-html-c6fdd6a8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-self-signed-chain-html-c6fdd6a8.md
source_anchor: ""
source_lines: [205, 249]
sha256: 2a8310e83ce54f373fe5e857732a441615327a63712e8b495af1acf9a94c4dbd
---

# manual-how-tos-self-signed-chain-html-c6fdd6a8

The private key does not have to be stored. For higher security, only download it by setting the Private Key Location to the alternative option. This way, the private key can be installed on the required server without a copy in the OPNsense trust store. To issue leaf certificates, only the Intermediate CAs need their private keys stored in the trust store.
Tip
Additional leaf certificates can be issued with the same intermediate CA. Each level in the chain is a one-to-many relationship. One root CA can issue many intermediate CAs, one intermediate CA can issue many leaf certificates.
The next step will be exporting the certificate chain, and using the leaf certificate on an external webserver.
Step 4: Exporting the Certificate Chain
Now that we have the whole certificate chain, we create a certificate bundle for a generic linux apache web server.
For that, we need two files:
- A certificate bundle populated in the correct order in PEM format: 
  - Root CA
  - Intermediate CA
  - Leaf Certificate
- The private key of the leaf certificate
- Export Root CA public key: 
  - Go to
  - Press the download button in the Commands column of the Root CA row
  - Choose File type: Certificate and press Download
- Export Intermediate CA public key: 
  - Go to
  - Press the download button in the Commands column of the Intermediate CA row
  - Choose File type: Certificate and press Download
- Export Leaf Certificate public key: 
  - Go to
  - Press the download button in the Commands column of the Leaf Certificate row
  - Choose File type: Certificate and press Download
- Export Leaf Certificate private key: 
  - Go to
  - Press the download button in the Commands column of the Leaf Certificate row
  - Choose File type: Private Key and press Download
Open a text editor and create a file with all 3 public keys:
certificate-bundle.pem
-----BEGIN CERTIFICATE-----
Root CA public key data
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
Intermediate CA public key data
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
Leaf Certificate public key data
-----END CERTIFICATE-----
The key of the certificate bundle is only the leaf certificate’s private key. Private keys of root and intermediate CAs are not required.
certificate-bundle.key
-----BEGIN PRIVATE KEY-----
Leaf Certificate private key data
-----END PRIVATE KEY-----
Implement these into your webserver, and the website will be secured with this certificate bundle. For automatic trust, you must install the intermediate and/or root CA certificate public keys into any client that connects to the webserver. Since the webserver offers a full certificate chain to the connecting client, manual trust can be established if a user decides to install the public key themselves.
