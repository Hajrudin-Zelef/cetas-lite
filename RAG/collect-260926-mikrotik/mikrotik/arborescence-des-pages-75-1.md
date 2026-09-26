---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-75-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["lora", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-75.md
source_anchor: ""
source_lines: [1, 117]
sha256: b6e751f2eab572cf514c932003f53c4ee2a4ebfba1b6ad81b188b11c6048cda8
---

# Overview

The general menu is used to manage certificates, add templates, issue certificates, and manage CRL and SCEP Clients.

## Certificate Template

Certificate templates are used to prepare a desired certificate for signing.

Certificate template is deleted right after a certificate is signed or a certificate request command is executed

To print out certificates:

### Certificate template properties

During the certificate template creation process, it is possible define and configure multiple parameters to meet specific requirements.

| Property | Description | 
|---|---|
| **common-name** (*string* ) | Certificate common name | 
| **copy-from**  (*name* ) | Certificate name from which to copy general settings | 
| **country** (*string* ) | Certificate issuer country | 
| **days-valid**  (days Default: 365) | Days certificate will be valid after signing | 
| **digest-algorithm** (*md5 \| sha1 \| sha256 \| sha384 \| sha512* Default:**sha256** ) | Certificate public key algorithm | 
| **key-size** (1024 \| 1536 \| 2048 \| 4096 \| 8192 \| prime256v1 \| secp384r1 \| secp521r1 Default:**2048** ) | Certificate public key size | 
| **key-usage** (code-sign \| crl-sign \| decipher-only \| dvcs \| encipher-only     key-cert-sign \| ocsp-sign \| tls-client \| content-commitment \| data-encipherment \| digital-signature \| email-protect \| key-agreement \| key-encipherment \| timestamp \| tls-server Default: **digital-signature,key-encipherment,data-encipherment,key-cert-sign,crl-sign,tls-server,tls-client)** | Certificate usage | 
| **locality** (*string* ) | Certificate issuer locality | 
| **name** (*string* ) | Certificate name | 
| **organization** (*string* ) | Certificate issuer organization | 
| **state** (*string* ) | Certificate issuer state | 
| **subject-alt-name** (*DNS: \| IP: \| email:* ) | Certificate subject alternative name | 
| **trusted** (*no \| yes*  ) | Wherever to trust certificate. If *yes,*  certificate will be used for host certificate verification. | 
| **trust-store**  (*all* \|*capsman* \|*dns* \|*email* \|*ipsec* \|*mqtt* \|*openflow* \|*radius* \|*sstp* \|*userman* \|*www* \|*api* \|*container* \|*dot1x* \|*fetch* \|*lora* \|*netwatch* \|*ovpn* \|*tr069* \|*wpa-eap* Default:**all** ) | Specify service which can use a specific certificate for certificate verification or trust-chain creation (www, sstp). | 
| **unit** (*string* ) | Certificate issuer organizational unit | 

## Certificate properties

For a signed certificate, most properties are read-only, with the exception of *name, trusted*, and *trust-store*.

| Property | Description | 
|---|---|
| **acme-status***(string)* | ACME client status | 
| **common-name** (*string* ) | Certificate common name | 
| **copy-from**  (*name* ) | Certificate name from which to copy general settings | 
| **country** (*string* ) | Certificate issuer country | 
| **days-valid**  (days) | Days certificate will be valid after signing | 
| **digest-algorithm** (*md5 \| sha1 \| sha256 \| sha384 \| sha512*  ) | Certificate public key algorithm | 
| **directory-url** *(string)* | ACME client directory URL | 
| **domain-names** *(string)* | ACME client used domain names | 
| **key-size** (1024 \| 1536 \| 2048 \| 4096 \| 8192 \| prime256v1 \| secp384r1 \| secp521r1) | Certificate public key size | 
| **key-usage** (code-sign \| crl-sign \| decipher-only \| dvcs \| encipher-only     key-cert-sign \| ocsp-sign \| tls-client \| content-commitment \| data-encipherment \| digital-signature \| email-protect \| key-agreement \| key-encipherment \| timestamp \| tls-server) | Certificate usage | 
| **locality** (*string* ) | Certificate issuer locality | 
| **organization** (*string* ) | Certificate issuer organization | 
| **revoked** *(date)* | Certificate revoke time (only for certificates that are signed and revoked in specific device) | 
| **state** (*string* ) | Certificate issuer state | 
| **subject-alt-name** (*DNS \| IP \| email* ) | Certificate subject alternative name | 
| **trusted** (*no \| yes* ) | Wherever to trust certificate. If *yes,*  certificate will be used for host certificate verification. | 
| **trust-store**  (*all* \|*capsman* \|*dns* \|*email* \|*ipsec* \|*mqtt* \|*openflow* \|*radius*  \|*sstp* \|*userman* \|*www* \|*api* \|*container* \|*dot1x* \|*fetch* \|*lora* \|*netwatch* \|*ovpn* \|*tr069* \|*wpa-eap* ) | Specify service which can use a specific certificate for certificate verification or trust-chain creation (www, sstp). | 
| **unit** (*string* ) | Certificate issuer organizational unit | 
| **serial-number** (*string* ) | Certificate serial number | 
| **fingerprint** (*string* ) | Certificate fingerprint | 
| **akid** (*string* ) | Certificate authority ID | 
| **skid** (*string* ) | Certificate subject ID | 
| **issuer** (*string* ) | Certificate Authority | 
| **invalid-before***(date)* | Date and time before which a certificate expired | 
| **invalid-after** *(date)* | Date and time after which a certificate expired | 
| **expires-after** *(time)* | Time left before expiration | 
| **key-type** (string) | Private key ype | 
| **ca** *(string)* | CA certificate name (shown only for certificates that are signed in specific device) | 

If the CA certificate is removed, all issued certificates in the chain are also removed.

## Sign Certificate

Certificates should be signed. In the following example, we will sign certificates and add CRL URL for the server certificate:

Let`s check is the certificates are signed:

For a video example click here.

The time of the key signing process depends on the key size of a specific certificate. With values of 4k and higher, it might take a substantial time to sign this specific certificate on less powerful CPU-based devices.

## Export Certificate

It is possible to export client certificates with keys and CA certificates in two formats - PEM or PCKS12.

| Property | Description | 
|---|---|
| **export-passphrase** (*string* Default: none)*sensitive* | Passphrase that will be used for exported certificate private key encryption. | 
| **file-name** (*string*  Default: cert_export_[Certificate name].crt/key/pkcs12) | Exported certificate file name. | 
| **type** (*pem \| pkcs12* Default: pem) | Exported certificate type. In case of PEM, certificate will be exported with CRT extension, if export-passphrase is specified, also encrypted private KEY file will be exported. In case of PKCS12, certificate will be exported with P12 extension, if export-passphrase is specified, exported certificate will contain encryted private key. | 

Exported certificates are available under the */file* section:

Exporting certificates requires "sensitive" user policy.

## Import Certificate

To import certificates, certificates must be uploaded to a device using one of the file upload methods.

Certificates must be imported as a file.

Supported are PEM, DER, CRT, PKCS12 formats.

| Property | Description | 
|---|---|
| **name** (*string* Default: file-name_number) | A certificate name that will be shown in the certificate manager | 
| **file-name** (*string* ) | A file name that will be imported | 
| **passphrase** (*string* Default: none)*sensitive* | File passphrase if there is such | 
| **trusted** (*yes \| no* Default: yes) | Adds *trusted* flag for imported certificate | 
| **trust-store**  (*all* \|*capsman* \|*dns* \|*email* \|*ipsec* \|*mqtt* \|*openflow* \|*radius* \| \|*sstp* \|*userman* \|*www* \|*api* \|*container* \|*dot1x* \|*fetch* \|*lora* \|*netwatch* \|*ovpn* \|*tr069* \|*wpa-eap* Default:**all** ) | Specify service which can use a specific certificate for certificate verification or trust-chain creation (www, sstp). | 

## Settings

*/certificate settings* allows configuring Certificate Revocation List (CRL) settings.

By default, CRL is not utilized, and certificates are not verified for revocation status.

