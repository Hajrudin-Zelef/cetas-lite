---
id: collect-260926-mikrotik/mikrotik/questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc-1
title: "questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc.md
source_anchor: ""
source_lines: [1, 60]
sha256: 0bb0ef9f7080894b62f3b5d0d156afc96c0bdb4174d56506ef380bb81f57e801
---

# questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc

Using EAP-TLS certificates to authenticate WiFi clients:
Upsides:
- Granular Access Control: Access can be both granted and restricted on a certificate basis, unlike WPA2 authentication where all users share the same password for the SSID
- Identity Validation: WPA2 password auth only proves a connecting WiFi user knows a password. A Certificate validates the identity of both users and the AP they're connecting to
- Traffic Encrypted between Client & Router: The "TLS" part of "EAP-TLS" ensures traffic between client & router is encrypted. Which is quite nice...
Downsides:
- Higher Administrative Burden: Creating certificates for each WiFi user and configuring their access requires more effort than setting a shared password.
- Unsuitable for Public Networks: Hotels, airports, cafes, etc cannot configure certificate based access for their large, transient WiFi user populations.
HowTo COMPATIBILITY:
- RouterOS Versions: Procedures were documented and developed using ***RouterOS v6.45.1 through v 6.46 ***. Tested & known to work correctly as of 20191214.
- Certificate Creation Procedures: Tested and known to work with IOS 11-13.3 and OSX Mojave and Catalina clients.  YMMV if connecting Windows or other clients using EAP-TLS.
WARNING:
The RouterOS CLI commands offered in this tutorial make excellent templates. But you MUST review and change my place holders and default values including the certificate export passphrase before executing commands in this tutorial.
CONFIGURATION PROCESS OVERVIEW: MikroTik ROUTER
A) CREATE CERTIFICATES: (3) types of certs must be created:
- Server: Create a Certificate for MikroTik side of EAP-TLS connection. This will be used by the wireless interface using EAP-TLS authentication.
- Client(s): Create a certificate for EACH Client connecting to the SSID using the wireless interface we configure for EAP-TLS.
- Certificate Authority: This certificate used to create a Chain of Trust for certs by signing both Client & Server certs.
B) CONFIGURE Wireless Security Profiles: Specify the certificates in Wireless Security Profiles. We will create a wireless Security Profile for the wireless interface using EAP-TLS authentication and for each connecting client.
C) CONFIGURE Wireless Interface: Finally, we will create a wireless interface specifying the wireless Security Profile for the interface.
SECTION 1: CREATE CERTIFICATES
Create CA (Certificate Authority) Certificate:
/certificate add name=CAF1Linux-template common-name=CAF1Linux country=GB days-valid=3650 key-size=4096 locality="Your Town" organization="Your Orgsanization" state=YourCounty trusted=yes unit="Technical Services" subject-alt-name="IP:1.2.3.4" key-usage=digital-signature,key-cert-sign,crl-sign;
/certificate sign CAF1Linux-template ca-crl-host="1.2.3.4" name=CAF1Linux
Export a PEM Format Cert:
/certificate export-certificate CAF1Linux export-passphrase="REPLACE ME WITH YOUR OWN CERTIFICATE PASSPHRASE"
The command /certificate export-certificate creates (2) new certs in files:
 cert_export_CAF1Linux.crt
 cert_export_CAF1Linux.key
Export a PKCS12 Format Cert:
/certificate export-certificate CAF1Linux export-passphrase="REPLACE ME WITH YOUR OWN CERTIFICATE PASSPHRASE" type=pkcs12
Appending type=pkcs12 to the /certificate export-certificate command will produce the following in Files":
cert_export_CAF1Linux.p12    
Create SERVER Certificate:
This is the certificate used by the MikroTik's wireless interface offering EAP-TLS authentication.
NOTE: You'll remark in key-usage I additionally specify ipsec-tunnel,ipsec-end-system. I use the same certs for both WiFi and VPN access, making it easy to centrally revoke a cert for both services if a server is compromised or to revoke a user's access.
/certificate add name=F1LinuxServer-template common-name="F1LinuxServer" country=GB days-valid=3650 key-size=4096 locality="Your City or Town" organization="Your Company" state=YourStateOrCounty trusted=yes unit="Technical Services" subject-alt-name="IP:1.2.3.4" key-usage=digital-signature,data-encipherment,key-agreement,ipsec-tunnel,ipsec-end-system,tls-server,tls-client;
/certificate sign F1LinuxServer-template ca=CAF1Linux name=F1LinuxServer
/certificate set F1LinuxServer trusted=yes
Export a PEM Format Cert:
/certificate export-certificate F1LinuxServer export-passphrase="REPLACE ME WITH YOUR OWN CERTIFICATE PASSPHRASE"
Export a PKCS12 Format Cert: Apple Clients require pkcs12 certs, so we will export all Client certs additionally in pkcs12 format
/certificate export-certificate F1LinuxServer export-passphrase="REPLACE ME WITH YOUR OWN CERTIFICATE PASSPHRASE" type=pkcs12
Repeat above process to create a unique certificate for EACH connecting client.
Create CLIENT Certificates:
An example for a MacBook is shown below, however the process is the same for any device supporting EAP-TLS auth.
/certificate add name=F1LinuxClientMacBook-template common-name=F1LinuxClientIpadPro country=GB days-valid=3650 key-size=4096 locality="Your City or Town" organization="Your Company" state=YourStateOrCounty trusted=yes unit="Technical Services" subject-alt-name="" key-usage=digital-signature,data-encipherment,key-agreement,ipsec-tunnel,ipsec-end-system,tls-client;
/certificate sign F1LinuxClientMacbook-template ca=CAF1Linux name=F1LinuxClientMacBook
/certificate set F1LinuxClientMacBook trusted=yes
Export a PEM Format Cert:
/certificate export-certificate F1LinuxClientMacBook export-passphrase="REPLACE ME WITH A DIFFERENT PASSPHRASE FOR EACH CLIENT CERTIFICATE"
Export a PKCS12 Format Cert:
/certificate export-certificate F1LinuxClientMacBook export-passphrase="REPLACE ME WITH A DIFFERENT PASSPHRASE FOR EACH CLIENT CERTIFICATE" type=pkcs12
SECTION 2: CONFIGURE WIRELESS SECURITY PROFILES
After creating certificates for each connecting wireless client using EAP-TLS authentication, you can use these certificates to create Wireless Security Profile's.
Unlike standard WPA2 password encryption which sets a single password for all connecting clients, since each wireless client will have a unique certificate, we must create a Security Profile for each device and the wireless interface itself.
SERVER Security Profile
/interface wireless security-profiles add name="24083_F1_EAP_TLS_Server" mode=dynamic-keys authentication-types=wpa2-eap unicast-ciphers=aes-ccm group-ciphers=aes-ccm wpa-pre-shared-key="" wpa2-pre-shared-key="" supplicant-identity="" eap-methods=eap-tls tls-mode=verify-certificate tls-certificate=F1LinuxServer mschapv2-username="" mschapv2-password="" disable-pmkid=no static-algo-0=none static-key-0="" static-algo-1=none static-key-1="" static-algo-2=none static-key-2="" static-algo-3=none static-key-3="" static-transmit-key=key-0 static-sta-private-algo=none static-sta-private-key="" radius-mac-authentication=yes radius-mac-accounting=yes radius-eap-accounting=yes interim-update=0s radius-mac-format=XX:XX:XX:XX:XX:XX radius-mac-mode=as-username radius-called-format=mac:ssid radius-mac-caching=disabled group-key-update=5m management-protection=allowed management-protection-key=""
CLIENT Security Profile
/interface wireless security-profiles add name="24083_F1_EAP_TLS_MacBook" mode=dynamic-keys authentication-types=wpa2-eap unicast-ciphers=aes-ccm group-ciphers=aes-ccm wpa-pre-shared-key="" wpa2-pre-shared-key="" supplicant-identity="" eap-methods=eap-tls tls-mode=verify-certificate tls-certificate=F1LinuxClientMacBook mschapv2-username="" mschapv2-password="" disable-pmkid=no static-algo-0=none static-key-0="" static-algo-1=none static-key-1="" static-algo-2=none static-key-2="" static-algo-3=none static-key-3="" static-transmit-key=key-0 static-sta-private-algo=none static-sta-private-key="" radius-mac-authentication=yes radius-mac-accounting=yes radius-eap-accounting=yes interim-update=0s radius-mac-format=XX:XX:XX:XX:XX:XX radius-mac-mode=as-username radius-called-format=mac:ssid radius-mac-caching=disabled group-key-update=5m management-protection=allowed management-protection-key="" 
