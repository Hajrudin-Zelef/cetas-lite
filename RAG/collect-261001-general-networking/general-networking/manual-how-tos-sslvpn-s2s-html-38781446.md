---
id: collect-261001-general-networking/general-networking/manual-how-tos-sslvpn-s2s-html-38781446
title: "manual-how-tos-sslvpn-s2s-html-38781446"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-sslvpn-s2s-html-38781446.md
source_anchor: ""
source_lines: [1, 72]
sha256: 8945a1ae513485921e06ef1bd39a64d4d67addc93ac3e6758b2151ecf072f9b5
---

# manual-how-tos-sslvpn-s2s-html-38781446

Setup SSL VPN site to site tunnel
Site to site VPNs connect two locations with static public IP addresses and allow traffic to be routed between the two networks. This is most commonly used to connect an organization’s branch offices back to its main office, so branch users can access network resources in the main office.
Note
When using the site to site example with SSL/TLS instead of a shared key, make sure to configure “client specific overrides”
as well to correctly bind the remote networks to the correct client.
Before you start
Before starting with the configuration of an OpenVPN SSL tunnel you need to have a working OPNsense installation with a unique LAN IP subnet for each side of your connection (your local network needs to be different than that of the remote network).
Note
For the sample we will use a private IP for our WAN connection. This requires us to disable the default block rule on WAN to allow private traffic. To do so, go to and uncheck “Block private networks”. (Don’t forget to save and apply)
Sample Setup
For the sample configuration we use two OPNsense boxes to simulate a site to site tunnel, with the following configuration:
Site A - Server
| Hostname | fw1 | 
| WAN IP | 172.10.1.1/16 | 
| LAN IP | 192.168.1.1/24 | 
| LAN DHCP Range | 192.168.1.100-192.168.1.200 | 
| Tunnel Network | 10.10.0.0/24 | 
Site B - Client
| Hostname | fw2 | 
| WAN IP | 172.10.2.1/16 | 
| LAN Net | 192.168.2.0/24 | 
| LAN DHCP Range | 192.168.2.100-192.168.2.200 | 
| Tunnel Network | 10.10.0.0/24 | 
Full Network Diagram Including SSL VPN Tunnel
Step 1 - Add SSL Server
Adding a new SSL VPN server is relatively simple. We’ll start by adding a server that uses a shared key. This setup offers a good protection and it is easy to setup.
Go to and click on click Add in the top right corner of the form.
For our example will use the following settings (leave everything else on its default):
| Server Mode | Peer to Peer (Shared Key) | 
| Protocol | UDP | 
| Device Mode | tun | 
| Interface | WAN | 
| Local port | 1194 | 
| Description | SSL VPN Server | 
| Shared Key | Leave on enabled (checked) to create a new key | 
| DH Parameters Length | 4096 | 
| Encryption algorithm | AES-256-CBC (256-bit) | 
| Auth Digest Algorithm | SHA512 (512-bit) | 
| IPv4 Tunnel Network | 10.10.0.0/24 | 
| IPv4 Local Network/s | 192.168.1.0/24 | 
| IPv4 Remote Network/s | 192.168.2.0/24 | 
| Compression | Enabled with Adaptive Compression | 
Step 3 - Server Firewall Rules
To allow SSL VPN client connections, we should allow access to the OpenVPN server port on the WAN interface. When using multiple servers we need to open up each port.
For our configuration we only use one server accessible on UDP port 1194.
Next we also need to allow traffic from the VPN client network (192.168.2.0/24). For our example we will allow client to access anything on our local network(s), however you may decide just to allow traffic to one or more IPs.
You are done configuring Site A.
Step 4 - Site B Client
Now we will have to setup the client. Login to the second firewall, go to and click on add client in the upper right corner of the form.
Now enter the following into the form (and leave everything else default):
| Server Mode | Peer to Peer (Shared Key) | 
| Protocol | UDP | 
| Device Mode | tun | 
| Interface | WAN | 
| Server host or address | 172.10.1.1 | 
| Server port | 1194 | 
| Description | SSL VPN Client | 
| Shared Key | Uncheck to paste the shared key | 
| … | Paste your shared key | 
| Server Certificate | SSLVPN Server Certificate (CA: SSL VPN CA) | 
| DH Parameters Length | 4096 | 
| Encryption algorithm | AES-256-CBC (256-bit) | 
| Auth Digest Algorithm | SHA512 (512-bit) | 
| Hardware Crypto | No Hardware Crypto Acceleration | 
| IPv4 Tunnel Network | 10.10.0.0/24 | 
| IPv4 Remote Network/s | 192.168.1.0/24 | 
| Compression | Enabled with Adaptive Compression | 
Now click on Save to apply your settings.
The Connection Status can be viewed under
Step 5 - Client Firewall Rules
To allow traffic from the remote network just add a rule under OpenVPN tab.
Done
