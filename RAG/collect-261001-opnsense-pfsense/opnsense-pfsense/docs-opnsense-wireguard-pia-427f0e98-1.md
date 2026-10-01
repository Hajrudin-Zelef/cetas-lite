---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-opnsense-wireguard-pia-427f0e98-1
title: "docs-opnsense-wireguard-pia-427f0e98"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["kill switch"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-opnsense-wireguard-pia-427f0e98.md
source_anchor: ""
source_lines: [1, 177]
sha256: 676695ed21090bf998eee5ad9d208f065e26874bba80231e374cdc18a7fd5565
---

# docs-opnsense-wireguard-pia-427f0e98

WireGuard Private Internet Access
Thanks to https://forum.gl-inet.com/t/configure-wireguard-client-to-connect-to-nordvpn-servers/10422/27 and https://imgur.com/gallery/JBf2RF6, this is how I made a whole subnet go out on the internet through a NordVPN WireGuard connection.
Table of contents
Prerequisites
- OPNsense 20.7.8_4-amd64
- Ubuntu 20.04
- NordVPN user account
Getting started
Even though WireGuard requires static IPs, NordVPN has deviced a solution which makes this connection as secure as any other available protocol for both authentication and session data, namely double NAT’ing. Read more here: https://nordvpn.com/blog/nordlynx-protocol-wireguard/.
We need to fetch our private key and the corresponding site’s public key. Since there are no documentation for 3rd party apps yet, we’ll have to use what is available to do this for now, using the linux cli, WireGuard itself and software from NordVPN.
Subnet which will travel through NordVPN WireGuard interface is 192.168.10.0/24 - named 10_VPN.
WireGuard tunnel interface is named WG_NordVPN_FR.
NordVPN WireGuard gateway is GW_WG_NordVPN_FR.
Linux
WireGuard
Install WireGuard on a linux machine. Check tutorial here; https://www.wireguard.com/install/.
sudo apt install wireguard
NordLynx
Install NordVPN. Check tutorial here; https://support.nordvpn.com/Connectivity/Linux/1325531132/Installing-and-using-NordVPN-on-Debian-Ubuntu-Raspberry-Pi-Elementary-OS-and-Linux-Mint.htm.
sudo sh <(curl -sSf https://downloads.nordcdn.com/apps/linux/install.sh)
Check internet IP address before you start:
curl ifconfig.me
NordVPN login:
sudo nordvpn login
Please enter your login details.
Email / Username: user@name.com
Password: 
Welcome to NordVPN! You can now connect to VPN by using 'nordvpn connect'.
Change from default VPN protocol OpenVPN to NordLynx (WireGuard):
sudo nordvpn set technology NordLynx
Technology is successfully set to 'NordLynx'.
Connect with NordVPN:
sudo nordvpn connect
Connecting to France #111 fr111.nordvpn.com
You are connected to France #111 (fr111.nordvpn.com)!
You’ll notice that your public IP has changed.
curl ifconfig.me
After a successfull connection, figure out the IP scheme of this particular connection:
sudo wg
interface: nordlynx
  public key: UTZ4PHmX5zAOSvdhqp0Ed8q4z0OHgMk8ztalChHaPU=
  private key: (hidden)
  listening port: 39069
  fwmark: 0xca6c
peer: 21dz9Y6HFRzaXKLpFpcZHjcI5AJQopJW/JZShKjTKkZ=
  endpoint: 11.112.192.11:51820
  allowed ips: 0.0.0.0/0
  latest handshake: 39 seconds ago
  transfer: 3.09 KiB received, 3.46 KiB sent
  persistent keepalive: every 25 seconds
(These are not valid keys by the way).
What about tunnel address?
ip address show nordlynx
8: tun0: <POINTOPINT,NOARP,UP,LOWER_UP> mtu 1420 qdisc noqueue state UNKNOWN group default qlen 1000
    link/none
    inet 10.5.0.2/16 scope global nordlynx
        valid_lft forever preferred_lft forever
Allright. Whats the opposite side’s address?
ping 10.5.0.1
PING 10.5.0.1 (10.5.0.1) 56(84) bytes of data. 
64 bytes from 10.5.0.1: icmp_seq=1 ttl=64 time=6.86 ms
Let’s assume this is the gateway address :)
Private key
Now, figure out which private key you have for your user:
sudo wg show nordlynx private-key
FSzJDH1171AJKldoqohndlakO3918djals/jkdjkfl0=
(This is not a valid key by the way).
Now you have everything you need. Your private key, your public key, servers public key, the endpoint address and the port.
OPNSense configuration
Allright, we have what we need to get things going regards to configuring our OPNsense firewall.
WireGuard
Local
Add a server by pressing the little + icon
-  MAKE SURE TO SELECT “SHOW ADVANCED”
- Enabled: [-]
- Name: NordVPN
- Public Key: insert public key from sudo wg :UTZ4PHmX5zAOSvdhqp0Ed8q4z0OHgMk8ztalChHaPU= )
- Private Key: insert private key from sudo wg show nordlynx private-key :FSzJDH1171AJKldoqohndlakO3918djals/jkdjkfl0=
- Listen Port: 51822 (use a random port which is not in use on the system)
- DNS Server: 103.86.96.100, 103.86.99.100 (https://support.nordvpn.com/General-info/1047409702/What-are-your-DNS-server-addresses.htm
- Tunnel Address: insert inet address from ip addr show nordlynx :10.5.0.2/16
- Peers: Nothing selected, leave blank for now
- Disable Routes: Check
- Gateway: 10.5.0.1
Click Save.
Endpoints
Create a new Endpoint by hitting the + icon. Here you will copy the information from the [peer] section from sudo wg.
- Name: fr111.nordvpn.com
- Public Key: insert public key from sudo wg (21dz9Y6HFRzaXKLpFpcZHjcI5AJQopJW/JZShKjTKkZ= )
- Shared Secret:
- Allowed IPs: 0.0.0.0/0
- Endpoint Address: 11.112.192.11
- Endpoint Port: 51820
- Keepalive: 25
Click Save.
Now, go back to Local. Select the NordVPN WireGuard instance. Hit Edit (the little pencil).
- Under Peers, select the newly created fr111.nordvpn.com peer.
Click Save.
General
[-] Enable WireGuard
Hit Save.
After you have selected Save- go to List Configuration.
Because of our persistent keepalive - you should see the received and sent transfer is steadily increasing. You’ll also notice you have a successfull Handshake with the specific interface whenever this is > 0 (wg0).
Assignments
Now go to Interfaces > Assignments. You’ll have a new interface you can assign (wg0).
Assign this interface. After assignment, click the name of the interface (OPT5 or something similar).
[x] Enable Interface
- Description: WAN_WG_NordVPN_FR
Leave rest of the configuration as is. Click Save.
Apply the changes.
Gateways
Go System > Gateways Click +Add gateway.
- Name: GW_WG_NordVPN_FR
- Description: PIA through WG_NordVPN_FR
-  Interface: WG_NordVPN_FR
-  IP adress: 10.5.0.1
- Far Gateway [x]
Set rest to default.
Click Save, Apply.
Rules
Go to Rules.
Select the designated interface for your subnet which you would like to go out on internet through this WireGuard VPN.
Add Rule.
Allow any any IPv4, but be sure to select
- Gateway: GW_WG_NordVPN - 10.5.0.1 as your gateway under Advanced settings.
While we are at it, do this to enable a kill-switch for your traffic should the WireGuard interface go down:
- Set local tag: NO_WAN_EGRESS
Click Add.
Add another rule on the same interface, but this time - make sure to select Block. Leave the rest as default. This will be our “kill switch”. Make sure the reject rule is below the allow rule.
Kill switch
Firewall > Rules > Floating > +Add
- Action: Block
- Interface: WAN
- Direction: out
- Description: NO_WAN_EGRESS match
- Match local tag: NO_WAN_EGRESS
NAT
Firewall > NAT > Outbound Select Hybrid outbound NAT rule generation. Click Save and Apply Rules.
Then click +Add.
Interface: WG_NordVPN_FR Source adress: 10_VPN net Translation / target: Interface address Description: 10_VPN to WG_NordVPN_FR
Save. Apply changes.
DNS
Let us provide some security regards to DNS leaks on this 10_VPN interface of ours.
Services > DHCPv4
- 10_VPN
Add DNS servers provided by NordVPN here, so that DHCP offers DNS servers provided by NordVPN:
- DNS servers : 103.86.96.100, 103.86.99.100
If you have devices that have hardcoded DNS servers, you want to redirect those requests to NordVPN’ DNS servers. We’ll define an ALIAS and use NAT port forwarding to achieve this.
Firewall > Aliases. Hit the + icon.
- Name: ALIAS_HOSTS_NordVPN_DNS
- Type: Host(s)
- Content: 103.86.99.100, 103.86.96.100
- Description: NordVPN DNS servers
Now, go to Firewall > NAT > Port Forward.
+Add
- Interface: 10_VPN
- Protocol: TPC / UDP
- Source: 10_VPN net
- Destination / Invert: checked
- Destination: ALIAS_HOSTS_NordVPN_DNS
- Destination port range: DNS
- Redirect target IP: ALIAS_HOSTS_NordVPN_DNS
- Redirect target port: DNS
Troubleshooting
DNS leaktest
resolvectl | grep 'DNS'
Current DNS Server: 103.86.96.100
Download dnsleaktest.sh from https://github.com/macvk/dnsleaktest
./dnsleaktest.sh
Your IP:
123.123.123.123 [France BE1800 K19]
You use 1 DNS server:
123.123.123.123 [France BE1800 K19]
Public IP
Are you connected through a VPN?
curl ifconfig.me
Recommended server
