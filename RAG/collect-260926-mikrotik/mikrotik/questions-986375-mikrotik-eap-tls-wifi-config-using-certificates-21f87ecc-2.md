---
id: collect-260926-mikrotik/mikrotik/questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc-2
title: "questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc.md
source_anchor: ""
source_lines: [61, 70]
sha256: eefe177d26143800fa1a9a44f592796d05d89cff676e36b174873d0a861fb7c2
---

# questions-986375-mikrotik-eap-tls-wifi-config-using-certificates-21f87ecc

Repeat the above command for each device authenticating with EAP-TLS, remembering to change the name of the certificate.
SECTION 3: CONFIGURE VIRTUAL AP WITH EAP AUTH
Finally configure a wireless interface to use the server's EAP-TLS Security Profile:
/interface wireless name="wlan1010" mtu=1500 l2mtu=1600 mac-address=74:4D:28:XX:XX:XX arp=proxy-arp interface-type=virtual master-interface=wlan5ghz mode=ap-bridge ssid="240E83_F1_EAP" vlan-mode=no-tag vlan-id=1 wds-mode=disabled wds-default-bridge=none wds-ignore-ssid=no bridge-mode=enabled default-authentication=no default-forwarding=yes default-ap-tx-limit=0 default-client-tx-limit=0 hide-ssid=yes security-profile=24083_F1_EAP_TLS_Server
BACKUP CONFIG AND CERTIFICATES:
Now that you've gone to all this effort, ensure you backup the config.  I use the naming convvention for my backups RBmodelNumber-YearMonthDay-Time_ROSversionNumber.rsc:
/export compact file=RB4011-20191214-1644_ROSv6.46.0.rsc
PLEASE NOTE*: Although you can restore the configuration to a new MikroTik, those backups won't capture your certificates.  Open a web browser and connect to the MikroTik via a WebGUI. Go to the "Files" menu and you can download each of them to your laptop and then punt them somewhere secure & sensible for long-term storage.
CLIENT CONFIGURATION
To learn how to configure IOS & OSX Clients to use EAP-TLS authentication, go HERE
