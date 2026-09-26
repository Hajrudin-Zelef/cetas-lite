---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-75-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "lora"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-75.md
source_anchor: ""
source_lines: [118, 188]
sha256: 072f3f4d9da5a468e2d46fe25be887908d705b2582bb93cb559636529c2cee05
---

# Overview

| Property | Description | 
|---|---|
| **builtin-trust-store** (*all* \|*default* \|*capsman* \|*dns* \|*email* \|*ipsec* \|*mqtt* \|*openflow* \|*radius* \| \|*sstp* \|*userman* \|*www* \|*api* \|*container* \|*dot1x* \|*fetch* \|*lora* \|*netwatch* \|*ovpn* \|*tr069* \|*wpa-eap \|**untrusted* Default:***default*** ) | Services that can use built-in trust store authorities for certificate verification. The current defaults:  | 
| **crl-download** (*yes \| no* Default:**no** ) | Whether to automatically download/update CRL | 
| **crl-store** (*ram \| sytem* Default:**ram** ) | Where to store downloaded CRL information CRL will be automatically renewed every hour for certificates which have "trusted=yes" using http protocol (ldap and ftp is currently unsupported) | 
| **crl-use** (*yes \| no* Default: no) | Whether to use CRL | 

If */certificate/settings/set crl-use* is set to *yes,* RouterOS will check CRL for each certificate in a certificate chain, therefore, an entire certificate chain should be installed into a device - starting from Root CA, intermediate CA (if there are such), and certificate that is used for specific service.

An example on importing a root certificate.

# ACME client

The ACME client automates the acquisition and renewal of multiple TLS certificates via ACME.

To add a new ACME client via CLI, use the command /certificate add-acme.

Existing ACME clients appear in the Certificates view and are marked with the *a* (acme-manage) flag.

Domain names must resolve to the router, and TCP port 80 must be accessible from the WAN (HTTP-01 challange is used). For example.sn.mynetname.net domain name, DNS-01 challange is used.

Certificates are automatically renewed when 80% of their validity period has elapsed.

If the certificate is not retrieved during the initial setup, a new ACME client must be added.

## Properties

| Property | Description | 
|---|---|
| **directory-url** (*string* ) | ACME directory URL | 
| **domain-names** (*string* ) | comma separated list of domain names | 
| **eab-hmac-key** (*string* ) | HMAC key for ACME External Account Binding | 
| **eab-kid** (*string* ) | Key identifier | 
| **name** (*string* ) | ACME client name | 

## Let's Encrypt certificate

To retrieve Let's Encrypt certificate with automatic certificate renewal, must manually provide ACME directory URL (https://acme-v02.api.letsencrypt.org/directory) and domain-name.

To generate Let's Encrypt certificate for /*ip cloud* name (ie. example.sn.mynetname.net), as domain-name provide *dns-name* from */ip/cloud* menu or use "*[/ip/cloud/get dns-name]*"

# SCEP

SCEP is using HTTP protocol and base64 encoded GET requests. Most of the requests are without authentication and cipher, however, important ones can be protected if necessary (ciphered or signed using a received public key).

SCEP client in RouterOS will:

- get CA certificate from CA server or RA (if used);
- user should compare the fingerprint of the CA certificate or if it comes from the right server;
- generate a self-signed certificate with a temporary key;
- send a certificate request to the server;
- if the server responds with status x, then the client keeps requesting until the server sends an error or approval.

The SCEP server supports the issuance of one certificate only. RouterOS supports also renew and next-ca options:

- renew - the possibility to renew the old certificate automatically with the same CA.
- next-ca - possibility to change the current CA certificate to the new one.

The client polls the server for any changes, if the server advertises that the next-ca is available, then the client may request the next CA or wait until CA almost expires and then request the next-ca.

The RouterOS client by default will try to use POST, AES, and SHA256 if the server advertises that. If the above algorithms are not supported, then the client will try to use 3DES, DES and SHA1, MD5.

SCEP certificates are renewed when 3/4 of their validity time has passed.

RouterOS contains list of built-in root certificate authorities that specific services can use for host certificate verification.

List of services that can use built-in root certificate authorities can be found here.

It is possible to use DoH, download Adlist from URL or use fetch tool with certificate validation without the need to manually import the relevant root certificate.

The list of built-in root certificate authorities is accessible in System → Certificates → Built In CA
