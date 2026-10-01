---
id: collect-261001-general-networking/general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc-2
title: "Add IPv4 route"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Nvidia"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-ipsec-swanctl-rw-ikev2-eap-mschapv2-html-81aa2fcc.md
source_anchor: ""
source_lines: [212, 414]
sha256: 15135c4f01c8bd7bfddb1ad530d96842e09ac4b446726dbaf7fbca7b114be8e1
---

# Add IPv4 route

Save to reveal the next options:
Local Authentication:
Round:
0
Authentication:
Public Key
Id:
vpn1.example.com
Certificates:
vpn1.example.com
Description:
local-vpn1.example.com
Remote Authentication:
Round:
0
Authentication:
EAP-MSCHAPv2
EAP Id:
laura@vpn1.example.com
Description:
remote-laura-eap-mschapv2
Children:
Press + to add a new Child, enable advanced mode with the toggle.
Start action:
None
ESP proposals:
aes256-sha256-modp2048 (Disable default!)
Local:
0.0.0.0/0 ::/0
Rekey time (s):
600 for most clients - Or 0 when using Windows native VPN client
Description:
roadwarrior-laura-eap-mschapv2-p2
Save and Apply the configuration.
Note
You have to repeat this workflow for each additional roadwarrior you create. They all need new pools and new connections.
Client configuration
In this section there are a few example configurations of different clients. All configurations here are tuned to the exact settings above. If you change anything in the server configuration, make sure you change it here too.
All clients are configured to use the Configuration Payload for virtual IP address, traffic selectors and DNS Server(s). They are pushed by the VPN server to the client.
Note
Import the CA certificate to clients, not the server certificate.
Windows 10/11 native VPN client
Note
- Windows 10/11 native VPN client works best with Method 1, which connects right away on the first authentication round.
- If you use Method 2 you should rather use the NCP client. The Windows VPN client does not send its local ID on the first authentication round. That means that users have to type their passwords twice before the connection establishes. You can mitigate one authentication round by saving the username and password into the vpn profile. Attention: If they press cancel or click outside of the authentication window, it will vanish and trying to connect again will fail until the PC is rebooted!
- Open Powershell as user (for userspace import) or as admin (for computer wide import) and apply the following commands:
Add-VpnConnection -Name "vpn1.example.com" -ServerAddress "vpn1.example.com" -TunnelType "Ikev2"
Set-VpnConnectionIPsecConfiguration -ConnectionName "vpn1.example.com" -AuthenticationTransformConstants SHA256 -CipherTransformConstants AES256 -EncryptionMethod AES256 -IntegrityCheckMethod SHA256 -PfsGroup PFS2048 -DHGroup Group14 -PassThru -Force
- Only set this parameter if you want a split tunnel:
Set-VpnConnection -Name "vpn1.example.com" -SplitTunneling $true
If you use Split Tunneling, you must set routes manually. You can use the Powershell command Add-VpnConnectionRoute to add routes:
# Add IPv4 route
Add-VpnConnectionRoute -ConnectionName 'vpn1.example.com' -DestinationPrefix '192.168.1.0/24' -PassThru
# Add IPv6 route
Add-VpnConnectionRoute -ConnectionName 'vpn1.example.com' -DestinationPrefix 'fe0d:abcd:1234:cafe::/64' -PassThru
# Get corresponding route with VPN connection
(Get-VpnConnection -ConnectionName 'vpn1.example.com').routes
# Remove associate route
Remove-VpnConnectionRoute -ConnectionName 'vpn1.example.com' -DestinationPrefix '192.168.1.0/24' -PassThru
- Import the CA certificate into the Windows certificate store, please note that you have to be admin for this action: 
  - Open MMC: Windows + R > Type mmc > Enter.
  - Add Certificates Snap-In: File > Add/Remove Snap-in > Certificates > Add > Computer account > Local computer > Finish.
  - Install Certificate: Go to Trusted Root Certification Authorities > Certificates > Right-click > All Tasks > Import > Select your CA certificate > Ensure it is set to Trusted - Root Certification Authorities > Finish.
  - Confirm: Check the certificate appears under Trusted Root Certification Authorities.
  - Close MMC. Choose ‘No’ if asked to save console settings.
- Connect the new VPN connection and use the following credentials, you can also save them prior to connecting: 
  - Username: john@vpn1.example.com
  - Password: 48o72g3h4ro8123g8r
Optional if DNS Server provisioning via Configuration Payload does not work: - Set up DNS for the VPN:
Open Network Connections: Windows + R > Type ncpa.cpl > Enter.
Locate VPN adapter (e.g. “vpn1.example.com”).
Right-click VPN adapter > Properties.
- For IPv4:
Select Internet Protocol Version 4 (TCP/IPv4) > Properties.
Set DNS: 192.168.1.1
- For IPv6:
Select Internet Protocol Version 6 (TCP/IPv6) > Properties.
Set DNS: 2001:db8:1234:1::1
Click OK to apply changes.
iOS native VPN client
- Import the self-signed CA certificate into the iOS certificate store.
- Go to Settings > General > VPN.
- Tap on Add VPN Configuration….
- Select the type of VPN you are using. For this example, it is IKEv2.
- In the fields provided, enter: 
  - Description: vpn1.example.com
  - Server: vpn1.example.com
  - Remote ID: vpn1.example.com
  - Local ID: john@vpn1.example.com
- In the Authentication section, select Username. 
  - Username: john@vpn1.example.com
  - Password: 48o72g3h4ro8123g8r
- Tap Done in the top right corner.
- To connect to the VPN, go back to Settings > VPN, then turn the VPN toggle switch to the ON position next to the profile you just created.
Note
iOS does not allow setting a DNS Server for the VPN, and it ignores the DNS Configuration Payload. The only workaround would be to change the DNS Server manually in the Wi-Fi settings each time the tunnel is brought up, and change them back when it is turned off.
Android StrongSwan VPN client
- Import the self-signed CA certificate into the Android certificate store.
- Install the StrongSwan app from the Google Play Store
- Open the StrongSwan app and create a new VPN profile. 
  - Server: vpn1.example.com
  - VPN Typ: IKEv2 EAP
  - Username: john@vpn1.example.com
  - Password: 48o72g3h4ro8123g8r
  - CA-Certificate: choose the imported CA certificate
  - Activate advanced mode:
  - IKEv2 Algorithms: aes256-sha256-modp2048
  - IPsec/ESP Algorithms: aes256-sha256-modp2048
- You can start the new profile and it should connect. If not, check the Logfile for the error message.
Android native VPN client
- Import the self-signed CA certificate into the Android certificate store.
- Create a new VPN network. 
  - Name: vpn1.example.com
  - Type: IKEv2/IPSec MSCHAPv2`
  - Server address: vpn1.example.com
  - IPSec identifier: leave empty
  - IPSec CA certificate: CN=IPsec CA ...
  - IPSec server certificate: Received from server
  - Username: john@vpn1.example.com
  - Password: 48o72g3h4ro8123g8r
Note
On the IPsec server, the local EAP identifier must be the username, and the remote EAP identifier must be the server address. Otherwise the authentication round will fail.
Windows/macOS NCP Secure Entry client
Attention
This is a commercial client and needs to be licensed. It is not affiliated with Deciso B.V. or OPNsense®.
- Install the NCP Secure Entry Client
- Save the following code as example.ini
[GENERAL]
Export=1
Product=NCP Secure Entry Client
Version=13.14 Build 29669
Date=11.09.2023 09:30:42
[PROFILE1]
Name=vpn1.example.com
ConnMedia=21
UseForAuto=0
SeamRoaming=1
NotKeepVpn=0
BootProfile=0
UseRAS=0
SavePw=0
PhoneNumber=
DialerPhone=
ScriptFile=
HttpName=
HttpPw=
HttpScript=
Modem=
ComPort=1
Baudrate=57600
RelComPort=1
InitStr=
DialPrefix=
3GApnSrc=2
3GProvider=
APN=
3GPhone=
3GAuth=0
GprsATCmd=AT+CPIN=
GprsPin=""
BiometricAuth=0
PreAuthEap=0
PreAuthHttp=0
ConnMode=0
Timeout=0
TunnelTrafficMonitoring=0
TunnelTrafficMonitoringAddr=0.0.0.0
QoS=none
PkiConfig=
ExchMode=34
TunnelIpVersion=1
IKEv2Auth=3
IKE-Policy=automatic mode
IKEv2Policy=aes256-sha256
IkeDhGroup=14
IkeLTSec=000:00:40:00
IPSec-Policy=aes256-sha256
PFS=14
IPSecLTType=1
IpsecLTSec=000:00:10:00
IPSecLTKb=50000
UseComp=0
IkeIdType=3
IkeIdStr=john@vpn1.example.com
Gateway=vpn1.example.com
ConnType=1
UsePreShKey=0
XAUTH-Src=0
SplitOptionV4=1
UseTunnel=1
SplitOptionV6=1
VpnBypass=none
UseXAUTH=1
UseUdpEnc=500
UseUdpEncTmp=4500
DisDPD=0
DPDInterval=30
DPDRetrys=8
AntiReplay=0
PathFinder=0
UseRFC7427=1
RFC7427Padding=2
Ikev2AuthPrf=0
