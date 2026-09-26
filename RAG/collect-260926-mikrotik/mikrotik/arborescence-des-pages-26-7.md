---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-7
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [570, 735]
sha256: bc179c55710c8a26d8d3c9dbf2dd2de7bee0198b7c4ea9a36b0c388a7c44bbaa
---

# Overview

Print output:

```
[admin@CAPsMAN] /caps-man manager print 
                   enabled: yes
               certificate: auto
            ca-certificate: auto
              package-path: 
            upgrade-policy: none
  require-peer-certificate: no
     generated-certificate: CAPsMAN-D4CA6D987C26
  generated-ca-certificate: CAPsMAN-CA-D4CA6D987C26
```
CAPsMAN device first will generate *CA-Certificate* and then it will generate *Certificate* which depends on *CA-Certificate*.

**CAP device:**

Set in CAP configuration to *request* certificate:

/interface wireless cap
set certificate=request

CAP will connect to CAPsMAN and request certificate. CAP will receive *CA-Certificate* form CAPsMAN and another certificate will be created for use on CAP.

**In Result**

On CAP device in CAP menu *Requested Certificate* is set:

```
[admin@CAP] /interface wireless cap print
                            enabled: yes
                         interfaces: wlan1
                        certificate: request
                   lock-to-caps-man: no
               discovery-interfaces: ether1
                 caps-man-addresses: 
                     caps-man-names: 
  caps-man-certificate-common-names: 
                             bridge: none
                     static-virtual: no
         -->  requested-certificate: CAP-D4CA6D7F45BA  <--
```
Also, two certificates are gained and are seen in *Certificate* menu:

[admin@CAP] > /certificate print 
Flags: K - private-key, D - dsa, L - crl, C - smart-card-key, A - authority, I - issued, R - revoked, E - expired, T - trusted 
 #          NAME              COMMON-NAME              SUBJECT-ALT-NAME                                           FINGERPRINT             
 0     A  T _0                CAPsMAN-CA-D4CA6D987C26                                                             383e63d7b...
 1 K        CAP-D4CA6D7F45BA  CAP-D4CA6D7F45BA                                                                    d495d1a94...

On CAPsMAN device in Certificate menu three certificates are created. CAPsMAN and CAPsMAN-CA certificates, as well as a certificate which is issued to CAP:

[admin@CAPsMAN] > /certificate print 
Flags: K - private-key, D - dsa, L - crl, C - smart-card-key, A - authority, I - issued, R - revoked, E - expired, T - trusted 
 #          NAME                     COMMON-NAME              SUBJECT-ALT-NAME                                    FINGERPRINT         
 0 K   A  T CAPsMAN-CA-D4CA6D987C26  CAPsMAN-CA-D4CA6D987C26                                                      383e63d7b...
 1 K    I   CAPsMAN-D4CA6D987C26     CAPsMAN-D4CA6D987C26                                                         02b0f7ff4...
 2      I   issued_1                 CAP-D4CA6D7F45BA                                                             d495d1a94...

**Additionally**

If you want to allow only CAPs with a valid certificate to connect to this CAPsMAN you can set *Require Peer Certificate* to *yes* on CAPsMAN device:

/caps-man manager
set require-peer-certificate=yes

However, when you will want to add new CAP devices to your CAPsMAN network you will have to set this option to *no* and then back to *yes* after CAP has gained certificates. Every time you change this option CAPsMAN will drop all dynamic interfaces and CAPs will try to connect again.

If you want to lock CAP to specific CAPsMAN and be sure it won't connect to other CAPsMANs you should set option *Lock To CAPsMAN* to *yes*. Additionally, you can specify CAPsMAN to lock to by setting *CAPsMAN Certificate Common Names* on CAP device:

/interface wireless cap
set lock-to-caps-man=yes
set caps-man-certificate-common-names=CAPsMAN-D4CA6D987C26

### Manual certificates and issuing with SCEP

With this example, you can create your own certificates for CAPsMAN and take control over issuing certificates to CAPs. This configuration can be useful in big, growing CAPsMAN networks. Many segments of this example can be done differently depending on your situation and needs. At this point, some knowledge about Certificates and their application can be useful.

**CAPsMAN device:**

In *Certificate* menu add certificate templates for CA certificate and CAPsMAN server certificate:

/certificate
add name=CA-temp common-name=CA
add name=CAPsMAN-temp common-name=CAPsMAN

Now *Sign* the certifiace templates. First *Sign* the CA certificate and use CAPsMAN device IP as *CA CRL Host*:

/certificate
sign CA-temp ca-crl-host=10.5.138.157 name=CA
sign CAPsMAN-temp ca=CA name=CAPsMAN

Alternatively, previous two steps can be done with auto setting in *Certificate* and *CA-Certificate* option in CAPsMAN Manager menu, see the Fast and easy configuration.

*Export* CA certificate. You will have to *Import* it on CAP device. You can use *Download -> Drag&Drop* to CAP device, in this example *fetch* command is used later from CAP device. Using long passphrase is advisable - longer passphrase will take longer to crack if it gets into the wrong hands:

/certificate
export-certificate CA export-passphrase=thelongerthebetterpassphrase

Create *SCEP server* which will be used to issue and grant certificates to CAP devices:

/certificate scep-server
add ca-cert=CA path=/scep/CAPsMAN

Set certificates in CAPsMAN Manager menu and set *Require Peer Certificate* to yes:

/caps-man manager
set ca-certificate=CA certificate=CAPsMAN
set require-peer-certificate=yes

At this point, only CAPs with a valid certificate will be able to connect.

**CAP device**

Download export of CA certificate from CAPsMAN device to CAP device. In this example *fetch* is used, however, there are multiple other ways:

/tool fetch address=10.5.138.157 src-path=cert_export_CA.crt user=admin password="123" mode=ftp

Import CA certificate from CAPsMAN device in *Certificate* menu:

/certificate> import file-name=cert_export_CA.crt passphrase=thelongerthebetterpassphrase

Add certificate template for CAP:

/certificate
add name=CAP1 common-name=CAP1

Ask CAPsMAN device to grant this certificate with a key using SCEP:

/certificate
add-scep template=CAP1 scep-url="https://10.5.138.157/scep/CAPsMAN"

You will have to return to CAPsMAN device to grant key to this certificate.

In CAP menu set just created certificate:

/interface wireless cap
set certificate=CAP1

**CAPsMAN device:**

Return to CAPsMAN device to grant a key to CAP certificate in *Certificate Request* menu:

/certificate scep-server requests
grant numbers=0

**In Result**

Now CAP should be able to connect to CAPsMAN, see in *CAPsMAN interfaces* if it connects. In CAPsMAN device *Certificate* menu three certificates can be seen: CA, CAPsMAN, and the one which is issued to CAP:

[admin@CAPsMAN] /certificate print 
Flags: K - private-key, D - dsa, L - crl, C - smart-card-key, A - authority, I - issued, R - revoked, 
E - expired, T - trusted 
 #          NAME        COMMON-NAME      SUBJECT-ALT-NAME                                   FINGERPRINT     
 0 K L A  T CA          CA                                                                  752775b457a37...
 1 K   A    CAPsMAN     CAPsMAN                                                             12911ba445b3b...
 2      I   issued_1    CAP1                                                                5b9a52b6ce3fb...

In CAP devices *Certificate* menu two acquired certificates can be seen:

[admin@CAP1] /interface wireless> /certificate print 
Flags: K - private-key, D - dsa, L - crl, C - smart-card-key, A - authority, I - issued, R - revoked, 
E - expired, T - trusted 
 #          NAME        COMMON-NAME      SUBJECT-ALT-NAME                                   FINGERPRINT     
 0   L A  T cert_exp... CA                                                                  752775b457a37...
 1 K      T CAP1        CAP1                                                                5b9a52b6ce3fb...
