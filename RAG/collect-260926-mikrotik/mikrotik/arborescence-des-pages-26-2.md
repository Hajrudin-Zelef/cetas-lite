---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [106, 202]
sha256: 26cbabc061d759b632a947b60e2c7eceb26e6d72b1998df6b85f62e01318b1d4
---

# Overview

CAPsMAN has the following certificate settings:

- **certificate** - this is CAPsMAN certificate, private key must be available for this certificate. If set to**none** , CAPsMAN will operate in no-certificate mode and none of certificate requiring features will work. If set to**auto** , CAPsMAN will attempt to issue certificate to itself using CA certificate (see**ca-certificate** description). Note that CommonName automatically issued certificate will be "CAPsMAN-<mac address>" and validity period for will be the same as for CA certificate.
- **ca-certificate** - this is CA certificate that CAPsMAN will use when issuing certificate for itself if necessary (see**certificate** description) and when signing certificate requests from CAPs. If set to**none** , CAPsMAN will not be able to issue certificate to itself or sign certificate requests from CAPs. If set to**auto** , CAPsMAN will generate self-signed CA certificate to use as CA certificate. CommonName for this certificate will take form "CAPsMAN-CA-<mac address>" and validity period will be from jan/01/1970 until jan/18/2038.

When CAPsMAN will auto-generate certificates, this will be reflected like this:

```
[admin@CM] /caps-man manager> pr
                   enabled: yes
               certificate: auto
            ca-certificate: auto
  require-peer-certificate: no
     generated-certificate: CAPsMAN-000C424C30F3
  generated-ca-certificate: CAPsMAN-CA-000C424C30F3
```
And certificates:

```
[admin@CM] /certificate> print detail
Flags: K - private-key, D - dsa, L - crl, C - smart-card-key, 
A - authority, I - issued, R - revoked, E - expired, T - trusted 
 0 K   A T name="CAPsMAN-CA-000C424C30F3" common-name="CAPsMAN-CA-000C424C30F3" key-size=2048 
           days-valid=24854 trusted=yes 
           key-usage=digital-signature,key-encipherment,data-encipherment,key-cert-sign,crl-sign 
           serial-number="1" fingerprint="69d77bbb45c50afd2d6c1785c2a3d72596b8a5f6" 
           invalid-before=jan/01/1970 00:00:01 invalid-after=jan/18/2038 03:14:07 
 1 K   I   name="CAPsMAN-000C424C30F3" common-name="CAPsMAN-000C424C30F3" key-size=2048 
           days-valid=24854 trusted=no key-usage=digital-signature,key-encipherment 
           ca=CAPsMAN-CA-000C424C30F3 serial-number="1" 
           fingerprint="e853ddb9d41fc139083a176ab164331bc24bc5ed" 
           invalid-before=jan/01/1970 00:00:01 invalid-after=jan/18/2038 03:14:07 
```
CAP can be configured to request certificate from CAPsMAN. In order for this to work, CAP must be configured with setting **certificate=request** and CAPsMAN must have CA certificate available (either specified in **ca-certificate** setting or auto-generated).

CAP will initially generate private key and certificate request with CommonName of form "CAP-<mac address>". When CAP will establish connection with CAPsMAN, CAP will request CAPsMAN to sign its certificate request. If this will succeed, CAPsMAN will send CA certificate and newly issued certificate to CAP. CAP will import these certificates in its certificate store:

```
[admin@CAP] > /interface wireless cap print
...
              requested-certificate: cert_2
        locked-caps-man-common-name: CAPsMAN-000C424C30F3
[admin@CAP] > /certificate print detail 
Flags: K - private-key, D - dsa, L - crl, C - smart-card-key, 
A - authority, I - issued, R - revoked, E - expired, T - trusted 
 0       T name="cert_1" issuer=CN=CAPsMAN-CA-000C424C30F3 common-name="CAPsMAN-CA-000C424C30F3" 
           key-size=2048 days-valid=24837 trusted=yes 
           key-usage=digital-signature,key-encipherment,data-encipherment,key-cert-sign,crl-sign 
           serial-number="1" fingerprint="69d77bbb45c50afd2d6c1785c2a3d72596b8a5f6" 
           invalid-before=jan/01/1970 00:00:01 invalid-after=jan/01/2038 03:14:07 
 1 K     T name="cert_2" issuer=CN=CAPsMAN-CA-000C424C30F3 common-name="CAP-000C4200C032" 
           key-size=2048 days-valid=24837 trusted=yes 
           key-usage=digital-signature,key-encipherment serial-number="2" 
           fingerprint="2c85bf2fbc9fc0832e47cd2773a6f4b6af35ef65" 
           invalid-before=jan/01/1970 00:00:01 invalid-after=jan/01/2038 03:14:07 
```
On subsequent connections to CAPsMAN, CAP will use generated certificate.

# **CAP Configuration**

When an AP is configured to be controlled by CAPsMAN, configuration of the managed wireless interfaces on the AP is ignored *(exceptions: antenna-gain,antenna-mode)*. Instead, AP accepts configuration for the managed interfaces from CAPsMAN.

The CAP wireless interfaces that are managed by CAPsMAN and whose traffic is being forwarded to CAPsMAN (ie. they are not in *local forwarding* mode), are shown as *disabled*, with the note **Managed by CAPsMAN**. Those interfaces that are in *local forwarding* mode (traffic is locally managed by CAP, and only management is done by CAPsMAN) are not shown disabled, but the note **Managed by CAPsMAN** is shown

CAP behavior of AP is configured in **/interface wireless cap** menu. From there you can:

- Disable or enable CAP feature on the device
- Set list of wireless interfaces to be controlled by Manager
- Set list of interfaces over which CAP should attempt to discover Manager
- Set list of Manager IP addresses that CAP will attempt to contact during discovery
- Set list of Manager names that CAP will attempt to connect
- Set list of Manager certificate CommonNames that CAP will connect to
- Set bridge to which interfaces should be added when local forwarding mode is used


Each wireless interface on a CAP that is under CAPsMAN control appears as a virtual interface on the CAPsMAN. This provides maximum flexibility in data forwarding control using regular RouterOS features, such as routing, bridging, firewall, etc.CAPsMAN Configuration Concepts

Many wireless interface settings are able to be grouped together into named groups ('profiles') that simplifies the reuse of configuration - for example, common configuration settings can be configured in a 'configuration profile' and multiple interfaces can then refer to that profile. At the same time any profile setting can be overridden directly in an interface configuration for maximum flexibility.

Currently there are the following setting groups:

- channel - channel related settings, such as frequency and width
- datapath - data forwarding related settings, such as bridge to which particular interface should be automatically added as port
- security - security related settings, such as allowed authentication types or passphrase
- configuration - main wireless settings group, includes settings such as SSID, and additionally binds together other setting groups - that is, configuration profile can refer to channel, security, etc. named setting groups. Additionally any setting can be overridden directly in configuration profile.

Interface settings bind together all setting groups, but additionally any setting can be overridden directly in interface settings.

By means of setting groups, configuration is organized in hierarchical structure with interface (actual user of configuration) as the root. In order to figure out the effective value of some setting this structure is consulted in a fashion where a higher level setting value overrides a lower level value.

For example, when WPA2 passphrase to be used by a particular interface needs to be found, the following places are consulted and the first place with WPA2 passphrase configured specifies effective passphrase. "->" denotes referring to setting profile (if configured):

- interface passphrase
- interface->security passphrase
- interface->configuration passphrase
- interface->configuration->security passphrase

