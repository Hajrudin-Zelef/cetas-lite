---
id: collect-261001-general-networking/general-networking/manual-certificates-html-6ad89efd-2
title: "cat ca_chain.crt ca_crl.crl > my_chain.pem"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-certificates-html-6ad89efd.md
source_anchor: ""
source_lines: [101, 116]
sha256: 6a281211fb7846d819869d8a517de436eeb7f7a4ee72b5e3b13b12d594e0fa42
---

# cat ca_chain.crt ca_crl.crl > my_chain.pem

below.
| Path | Topic | 
|---|---|
| /etc/ssl/certs | Directory with links (and files) to certificates named by their hash. (for example ef954a4e.0 oref954a4e.r0 for a CRL) Primary location for the base system. | 
| /etc/ssl/untrusted | Same as above, but untrusted (ignored) | 
| /usr/local/openssl/certs | Same as /etc/ssl/certs , default location for applications build from the “ports” tree | 
| /usr/local/etc/ssl/cert.pem | Combined bundle file for applications that require a single file. | 
Each target link (or file) contains a single certificate or revocation list, which OpenSSL can easily locate using the certificate subject.
Note
If either /etc/ssl/cert.pem  and/or /usr/local/etc/ssl/cert.pem exists, they will have preference above
the hashed links insode the target directories. Staring with OPNsense version 25.1, these files will be removed
when they exist.
Default settings for OpenSSL are saved into /usr/local/openssl/openssl.cnf (ports) and /etc/ssl/openssl.cnf
(base), both are managed by OPNsense.
Usage examples
In Setup Self-Signed Certificate Chains you will find examples of how to setup certificate chains yourself.
