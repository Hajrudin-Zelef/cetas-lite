---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-2022-09-opnsense-wireguard-site-to-site-67f0bc3c-1
title: "/etc/wireguard/wg0.conf"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/blog-2022-09-opnsense-wireguard-site-to-site-67f0bc3c.md
source_anchor: ""
source_lines: [1, 97]
sha256: f3e16575dbf111f385fa3265ca0a858f38b84073c1c48a7b43d27c43f9521ce0
---

# /etc/wireguard/wg0.conf

OPNsense WireGuard Site to Site
This article will show you how to set up an OPNsense router with a WireGuard site-to-site topology. For our example scenario, on one one side of the WireGuard connection we’ll use a generic Linux router, Router A, and on the other side, we’ll have our OPNsense router, Router B:
This connection will allow hosts in Site A (behind the Linux Router A) to connect to hosts in Site B (behind the OPNsense Router B) as if they were in adjacent LANs (Local Area Networks). In particular, we’ll allow Endpoint A (with an IP address of 192.168.1.11) in the Site A LAN (using a subnet of 192.168.1.0/24) to connect to a webserver running on Endpoint B (TCP port 80 on 192.168.200.22) in the Site B LAN (192.168.200.0/24).
Here are the steps we’ll follow:
We’ll show you how you’d accomplish each step on the Linux router first, for comparison with the OPNsense router. The steps on the Linux router will be pretty much the same as you’ll find in the WireGuard Site-to-Site Configuration guide.
Install WireGuard on Linux
To install WireGuard on Router A, install WireGuard through its OS (Operating System) package manager (as described on the WireGuard Installation page).
Install WireGuard on OPNsense
To install WireGuard on Router B, navigate to the System > Firmware > Plugins page of the OPNsense GUI (Graphical User Interface). Search for the os-wireguard package in the plugins list, and click the Add icon for it:
Then navigate to VPN > WireGuard page. Select the Enable WireGuard checkbox, and click the Apply button:
Configure Local Interface on Linux
First, run the following commands on Router A to generate a new WireGuard key pair for it:
$ wg genkey > router-a.key
$ wg pubkey < router-a.key > router-a.pub
This will generate two files: router-a.key and router-a.pub. The first is the private key, and the second is the public key:
$ cat router-a.key
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEE=
$ cat router-a.pub
/TOE4TKtAqVsePRVR+5AA43HkAK5DSntkOCO7nYq5xU=
Then create a WireGuard config file on Router A at /etc/wireguard/wg0.conf with the following content:
# /etc/wireguard/wg0.conf
# local settings for Router A
[Interface]
PrivateKey = AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEE=
Address = 10.0.0.1
ListenPort = 51821
Be sure to replace the PrivateKey setting above with the private key you generated in the router-a.key file (and you can now delete the router-a.key file; you won’t need it again). See the Configure WireGuard section of the WireGuard Site-to-Site Configuration guide for a detailed explanation of each setting.
Configure Local Interface on OPNsense
To set up a WireGuard interface on Router B, in the OPNsense GUI, switch to the Local tab of the VPN > WireGuard page. Click the Add icon this page:
In the Edit Local Configuration dialog box, fill in the following fields:
- Name
- 
Display name for the WireGuard interface, like RouterB .
- Listen Port
- 
Publicly exposed UDP port on Router B to which Router A will connect, like 51822 .
- Tunnel Address
- 
Virtual IP address of this WireGuard interface, like 10.0.0.2 . Make sure this doesn’t conflict with any other networks to which Router B can connect.
Leave the Public Key and Private Key fields blank, so that OPNsense will generate a new WireGuard key pair for you. Then click the Save button:
Now click the Edit icon for the WireGuard interface you just created to view the public key pair:
Copy the generated value from the Public Key field to Router A; you’ll need it for the next step:
| Tip | If you leave the “Disable Routes” checkbox unchecked as shown in this guide, WireGuard will automatically add a route through the WireGuard interface for each network listed in the “Allowed IPs” field of the endpoints used by this interface — so you won’t have to manually make any route changes to enable access to those networks from the OPNsense LAN. If you check the “Disabled Routes” checkbox, you will have to add these routes manually. | 
Configure Remote Endpoint on Linux
Back on Router A, edit the /etc/wireguard/wg0.conf file, and add the following [Peer] section to it:
# /etc/wireguard/wg0.conf
# local settings for Router A
[Interface]
PrivateKey = AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEE=
Address = 10.0.0.1/32
ListenPort = 51821
# remote settings for Router B
[Peer]
PublicKey = fE/wdxzl0klVp/IR8UcaoGUMjqaWi3jAd7KzHKFS6Ds=
Endpoint = 203.0.113.2:51822
AllowedIPs = 192.168.200.0/24
Replace the PublicKey setting above with the public key you copied from Router B. Replace the Endpoint setting with the WAN (Wide Area Network) address of Router B (or the IP address of Router B from the perspective of Router A, if not the WAN address), plus the “Listen Port” setting for Router B that you chose in the step above. Include an AllowedIPs setting for each network in Site B that you want to be able to access from Site A through this WireGuard connection.
See the Configure WireGuard section of the WireGuard Site-to-Site Configuration guide for a detailed explanation of each setting.
Configure Remote Endpoint on OPNsense
To enable Router B to connect to Router A, in the OPNsense GUI, switch to the Endpoints tab of the VPN > WireGuard page. Click the Add icon on it:
In the Edit Remote Endpoint dialog box, fill in the following fields:
- Name
- 
Display name for the remote endpoint, like RouterA .
- Public Key
- 
Public key you generated on Router A. Copy the content of the router-a.pub file you generated above and paste it into this field.
- Allowed IPs
- 
List of networks in Site A that you want to be able to access from Site B through this WireGuard connection. In our example, we just want to be able to access the Site A LAN, 192.168.1.0/24 .
- Endpoint Address
- 
The WAN (Wide Area Network) address of Router A (or the IP address of Router A from the perspective of Router B, if not the WAN address). In our example, this is 198.51.100.1 .
- Endpoint Port
- 
The “ListenPort” setting for Router A from above; in our example, this is 51821 .
Then click the Save button:
| Tip | If the WAN interface for Router B is itself behind NAT (Network Address Translation) — for example, the ISP (Internet Service Provider) for Site B uses CGNAT (Carrier Grade NAT) — Site A would normally be blocked from initiating connections to Site B. To work around this, you can set the Keepalive field for this endpoint to 25 . This will direct WireGuard on Router B to send a keepalive packet to the endpoint on Router A every 25 seconds — opening a hole in the NAT that will allow Site A to initiate connections through the WireGuard tunnel to Site B. | 
Next, switch back to the Local tab, and click the Edit icon for the WireGuard interface:
Select the endpoint you just created (RouterA) as the value of the Peers field. Then click the Save button:
Finally, to apply your changes, click the Apply button:
This will restart the WireGuard service on OPNsense with your new WireGuard configuration settings. You’ll need to click the Apply button (on any of the WireGuard tabs) every time you want to apply any WireGuard changes you’ve made through the GUI.
If you switch to the List Configuration tab, you’ll see the active WireGuard configuration displayed (it may take a moment for it to appear). It should look like this:
interface: wg1
  public key: fE/wdxzl0klVp/IR8UcaoGUMjqaWi3jAd7KzHKFS6Ds=
  private key: (hidden)
  listening port: 51822
peer: /TOE4TKtAqVsePRVR+5AA43HkAK5DSntkOCO7nYq5xU=
  endpoint: 198.51.100.1:51821
  allowed ips: 192.168.1.0/24
You can use this tab to check that your changes to the OPNsense WireGuard configuration has been applied.
Configure Firewall on Linux
Back on Router A, you need to make sure that your firewall allows UDP port 51821 access from Router B.
If you’re using nftables on the router, the following minimal ruleset would allow unrestricted access between Site A and Site B:
#!/usr/sbin/nft -f
flush ruleset
define wan_iface = "eth0"
