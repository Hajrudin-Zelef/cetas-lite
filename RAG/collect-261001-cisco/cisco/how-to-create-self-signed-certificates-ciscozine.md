---
id: collect-261001-cisco/cisco/how-to-create-self-signed-certificates-ciscozine
title: "how-to-create-self-signed-certificates-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/how-to-create-self-signed-certificates-ciscozine.md
source_anchor: ""
source_lines: [1, 129]
sha256: 9c3e69c7342ae57fa9d909aa09f28b490706f5fab2587206c701ee42ee3edbba
---

# how-to-create-self-signed-certificates-ciscozine

A digital certificate or identity certificate is an electronic document which uses a **digital signature** to **bind a public key with an identity**, information such as the name of a person or an organization, their address, and so forth. The certificate can be used to verify that a public key belongs to an individual.

In a typical public key infrastructure (PKI) scheme, the signature will be of a certificate authority (CA). However, there are situations where it is not possible use a CA, so the only solutions is to use a self-signed certificate, an identity certificate that is signed by the same entity whose identity it certifies.

The **standard** used by Cisco is **X.509**, an ITU-T standard for a public key infrastructure (PKI) and Privilege Management Infrastructure (PMI). X.509 specifies, amongst other things, standard formats for public key certificates, certificate revocation lists, attribute certificates, and a certification path validation algorithm.

**The most common use** of certificates is for **HTTPS**-based web sites. A web browser validates that an SSL web server is authentic, so that the user can feel secure that his/her interaction with the web site has no eavesdroppers and that the web site is who it claims to be. **Other uses are VPN lan2lan, GetVPN and so on.**

**two** different **approaches** to create a self-signed certificate: **automatic** or **manual**.

To automatically create an rsa key pairs and a certificate, enable the https server:

Ciscozine(config)#ip http secure-server
% Generating 1024 bit RSA keys, keys will be non-exportable...
[OK] (elapsed time was 3 seconds)
Ciscozine(config)#
*May 23 21:35:01.855: %SSH-5-ENABLED: SSH 1.99 has been enabled
Ciscozine(config)#
*May 23 21:35:02.035: %PKI-4-NOAUTOSAVE: Configuration was modified.  Issue "write memory" to save new certificate
Ciscozine(config)#

Where is the certificate?

```
Ciscozine#dir nvram:
Directory of nvram:/
  124  -rw-          85                      startup-config
  125  ----           0                      private-config
  126  -rw-          85                      underlying-config
    1  ----          57                      persistent-data
129016 bytes total (127855 bytes free)
Ciscozine#
```
**Remember**: Save the configuration to save the certificate!

Ciscozine#wr
Building configuration...
[OK]
Ciscozine#

```
Ciscozine#dir nvram:
Directory of nvram:/
  122  -rw-        1069                      startup-config
  123  ----        1914                      private-config
  124  -rw-        1069                      underlying-config
    1  ----          57                      persistent-data
    2  -rw-         559                      IOS-Self-Sig#1.cer
129016 bytes total (123933 bytes free)
Ciscozine#
```
The second method requires three steps: create an rsa key pairs, create a self signed trust point and enroll the certificate.

Create an RSA keys:

Ciscozine(config)#crypto key generate rsa label ciscozine-rsa modulus 2048
The name for the keys will be: ciscozine-rsa
% The key modulus size is 2048 bits
% Generating 2048 bit RSA keys, keys will be non-exportable...
[OK] (elapsed time was 10 seconds)
Ciscozine(config)#

Create a local PKI:

Ciscozine(config)#crypto pki trustpoint my-ciscozine-ca
Ciscozine(ca-trustpoint)#enrollment selfsigned
Ciscozine(ca-trustpoint)#subject-name O=Test,CN=www.ciscozine.com
Ciscozine(ca-trustpoint)#rsakeypair ciscozine-rsa

**Remembers:** The “subject-name” is the name of the entity whose public key the certificate identifies. For instance, “O” identify the “Organization” and “CN” the Common name

Obtain the certificate from the local certificate authority:

Ciscozine(config)#crypto pki enroll my-ciscozine-ca
% Include the router serial number in the subject name? [yes/no]: no
% Include an IP address in the subject name? [no]: no
Generate Self Signed Router Certificate? [yes/no]: yes
Router Self Signed Certificate successfully created
Ciscozine(config)#

```
Ciscozine#dir nvram:
Directory of nvram:/
  124  -rw-          85                      startup-config
  125  ----           0                      private-config
  126  -rw-          85                      underlying-config
    1  ----          57                      persistent-data
129016 bytes total (127855 bytes free)
Ciscozine#
```
**Remember:** You must save the configuration to save the certificate!

Ciscozine#wr
Building configuration...
[OK]
Ciscozine#

```
Ciscozine#dir nvram:
Directory of nvram:/
  118  -rw-        1019                      startup-config
  119  ----        5468                      private-config
  120  -rw-        1019                      underlying-config
    1  ----          57                      persistent-data
    2  -rw-         860                      Ciscozine#1.cer
129016 bytes total (120429 bytes free)
Ciscozine#
```
**To see more information** about the certificates and the rsa keys are, you can use these commands:

- show crypto pki certificates
- show crypto pki trustpoints
- show crypto key mypubkey rsa

**References:**

Is it possible to create SHA-256 self signed certificates?

Check this link

http://www.cisco.com/c/en/us/td/docs/ios-xml/ios/sec_conn_pki/configuration/15-mt/sec-pki-15-mt-book/sec-deploy-rsa-pki.html

I have followed this guide, thank you very much. However my SSL VPN now results in a redirect loop after changing the self signed cert from the default 1024 RSA one to the 2048bit one from this guide. Is the subject-name required? I think this could be causing the redirect loop

how can i verify that the certificate working ?

is it possible to export the certificate from the Cisco router, if possible what is the procedure?

Yes, check this link https://supportforums.cisco.com/t5/security-documents/certificate-backup-and-installation-trustpoints/ta-p/3109377
