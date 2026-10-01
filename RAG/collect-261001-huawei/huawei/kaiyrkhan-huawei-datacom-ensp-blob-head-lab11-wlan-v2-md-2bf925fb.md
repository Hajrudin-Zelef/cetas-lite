---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab11-wlan-v2-md-2bf925fb
title: "kaiyrkhan-huawei-datacom-ensp-blob-head-lab11-wlan-v2-md-2bf925fb"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab11-wlan-v2-md-2bf925fb.md
source_anchor: ""
source_lines: [1, 209]
sha256: 619d793dde953bb5d1e4b9e5e6d6afeed9b30ebfdc53a207f75d9675c428b713
---

# kaiyrkhan-huawei-datacom-ensp-blob-head-lab11-wlan-v2-md-2bf925fb

Download Link for eNSP Topology File
Table1 - WLAN Data Plan
| Item | Value | 
|---|---|
| Management VLAN for APs | VLAN 43 | 
| Service VLAN for STAs | SSID Staff: VLAN 100, SSID Guest: VLAN 200 | 
| Default Gateway for AP | 10.1.43.254 | 
| DHCP Pool for APs | 10.1.43.101 - 10.1.43.200/24 | 
| Default Gateway for Staff | 192.168.100.254 | 
| DHCP Pool for Staff | 192.168.100.11 - 192.168.100.250/24 | 
| Default Gateway for Guest | 192.168.200.254 | 
| DHCP Pool for Guest | 192.168.200.11 - 192.168.200.250/24 | 
| AP Name | AP1, AP2 | 
| AP Group | Name: ap-group1 | 
|  | Referenced profiles: VAP profile VAP-Staff and VAP-Guest, Regulatory domain profile default | 
| Regulatory Domain Profile | Name: default | 
|  | Country code: KZ | 
| SSID Profile (Staff) | Name: WLAN-Staff | 
|  | SSID name: Staff-WiFi | 
| Security Profile (Staff) | Name: WLAN-Staff | 
|  | Security policy: WPA-WPA2+PSK+AES | 
|  | Password: Huawei@123 | 
| SSID Profile (Guest) | Name: WLAN-Guest | 
|  | SSID name: Guest-WiFi | 
| Security Profile (Guest) | Name: WLAN-Guest | 
|  | Security policy: WPA-WPA2+PSK+AES | 
|  | Password: Huawei@123 | 
| VAP Profile (Staff) | Name: VAP-Staff | 
|  | Forwarding mode: Direct forwarding | 
|  | Service VLAN: 100 | 
|  | Referenced profiles: SSID profile WLAN-Staff and Security profile WLAN-Staff | 
| VAP Profile (Guest) | Name: VAP-Guest | 
|  | Forwarding mode: Direct forwarding | 
|  | Service VLAN: 200 | 
|  | Referenced profiles: SSID profile WLAN-Guest and Security profile WLAN-Guest | 
Wireless STA (Station) — клиент құрылғы
Configure Hostname
system-view
sysname A1
Create VLANs
vlan batch 43 100 200
vlan 43
 description MGMT VLAN
vlan 100
 description Service VLAN
vlan 200
 description Service VLAN
display vlan
Configure Trunk Port and Allowed VLANs
interface g0/0/1
 port link-type trunk
 port trunk allow-pass vlan 43 100 200
interface g0/0/4
 port link-type trunk
 port trunk pvid vlan 43
 port trunk allow-pass vlan 43 100 200
display port vlan
PVID (Port VLAN ID) — Switch кіріс трафик үшін VLAN43 tag-ін алады, нәтижесінде untagged Frame-нен tagged Frame-ге өзгереді. Ал шығыс трафик үшін tag-ті алып тастап, AP-ға untagged Frame жібереді. Бұл Cisco әлеміндегі Native VLAN ұғымына толық сәйкес келеді. Default жағдайда кез келген Trunk портында бұл мән VLAN 1 болады!
Configure Hostname
system-view
sysname D1
Create VLANs
vlan batch 43 100 200
vlan 43
 description MGMT VLAN
vlan 100
 description Service VLAN
vlan 200
 description Service VLAN
display vlan
Configure Trunk Port and Allowed VLANs
interface g0/0/10
 port link-type trunk
 port trunk allow-pass vlan 43 100 200
interface g0/0/13
 port link-type trunk
 port trunk allow-pass vlan 43 100 200
interface g0/0/14
 port link-type trunk
 port trunk allow-pass vlan 43 100 200
display port vlan
Create VLANIF interface
interface vlanif 100
 ip address 192.168.100.254 24
 description Default Gateway for VLAN100
interface vlanif 200
 ip address 192.168.200.254 24
 description Default Gateway for VLAN200
display ip int brief
DHCP Pool for STAs
dhcp enable
ip pool VLAN100
 network 192.168.100.0 mask 24
 gateway-list 192.168.100.254
 dns-list 8.8.8.8
 excluded-ip-address 192.168.100.1 192.168.100.10
 excluded-ip-address 192.168.100.251 192.168.100.253
 lease day 5
interface vlanif 100
 dhcp select globalip pool VLAN200
 network 192.168.200.0 mask 24
 gateway-list 192.168.200.254
 dns-list 8.8.8.8
 excluded-ip-address 192.168.200.1 192.168.200.10
 excluded-ip-address 192.168.200.251 192.168.200.253
 lease day 5
interface vlanif 200
 dhcp select globaldisplay ip pool
Configure Hostname
system-view
sysname AC1
Create VLANs
vlan batch 43 100 200
vlan 43
 description MGMT VLAN
vlan 100
 description Service VLAN
vlan 200
 description Service VLAN
display vlan brief
Configure Trunk Port and Allowed VLANs
interface g0/0/10
 port link-type trunk
 port trunk allow-pass vlan 43 100 200
display port vlan
Create VLANIF interface
interface vlanif 43
 ip address 10.1.43.254 24
 description Default Gateway for APs
display ip int brief
CAPWAP Tunnel
capwap source interface Vlanif 43
DHCP Pool for APs
dhcp enable
 ip pool AP
 network 10.1.43.0 mask 24
 gateway-list 10.1.43.254
 option 43 sub-option 2 ip-address 10.1.43.254
 excluded-ip-address 10.1.43.1 10.1.43.100
 excluded-ip-address 10.1.43.201 10.1.43.253
 lease day 5
interface vlanif 43
 dhcp select global
display ip pool
1-қадам: WLAN mode
system-view
wlan
2-қадам: Create a Regulatory Domain Profile
regulatory-domain-profile name default
 country-code kz
quit
3-қадам: Create Security Profiles
security-profile name WLAN-Staff
 security wpa-wpa2 psk pass-phrase Huawei@123 aes
quitsecurity-profile name WLAN-Guest
 security wpa-wpa2 psk pass-phrase Huawei@123 aes
quit
4-қадам: Create SSID Profiles
ssid-profile name WLAN-Staff
 ssid Staff-WiFi
quitssid-profile name WLAN-Guest
 ssid Guest-WiFi
quit
5-қадам: Create VAP Profiles
vap-profile name VAP-Staff
 forward-mode direct-forward
 service-vlan vlan-id 100
 ssid-profile WLAN-Staff
 security-profile WLAN-Staff
quitvap-profile name VAP-Guest
 forward-mode direct-forward
 service-vlan vlan-id 200
 ssid-profile WLAN-Guest
 security-profile WLAN-Guest
quit
Forwarding Mode — трафикті бағыттау режимі
6-қадам: AP Group
AP Group-ға барлық Profile-дерді байланыстыру
ap-group name ap-group1
 regulatory-domain-profile default
 vap-profile VAP-Staff wlan 1 radio all
 vap-profile VAP-Guest wlan 2 radio all
quit
Import APs to the AC
wlan
 ap auth-mode mac-auth
 ap-id 0 ap-mac 00E0-FC84-1B70
 ap-name AP1
 ap-group ap-group1
 quit
 ap-id 1 ap-mac 00E0-FCDA-5BF0
 ap-name AP2
 ap-group ap-group1
 quit
Access Point Model: AirEngine 6761-21
Access Point Type: AP2050DN
AP1 MAC Address: 90F9-B722-2000
AP2 MAC Address: 90F9-B722-17C0
display ap all
7-қадам: Verify the Configuration
General Status
<AC1> display ap all
State: "Normal"
State: "Fault немесе idle" болса, AP-мен байланыс жоқ дегенді білдіреді!STA1> ipconfig
STA2> ipconfig<AC1> display station allChecking Profiles
<AC1> display wlan vap-profile all
<AC1> display wlan ssid-profile all
<AC1> display wlan security-profile all<AC1> display wlan vap allChecking AP Group
<AC1> display ap-group name ap-group1
