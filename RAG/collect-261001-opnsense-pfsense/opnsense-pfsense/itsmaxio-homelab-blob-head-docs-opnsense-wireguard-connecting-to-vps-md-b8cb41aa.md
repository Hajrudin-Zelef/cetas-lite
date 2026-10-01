---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/itsmaxio-homelab-blob-head-docs-opnsense-wireguard-connecting-to-vps-md-b8cb41aa
title: "itsmaxio-homelab-blob-head-docs-opnsense-wireguard-connecting-to-vps-md-b8cb41aa"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Oracle"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/itsmaxio-homelab-blob-head-docs-opnsense-wireguard-connecting-to-vps-md-b8cb41aa.md
source_anchor: ""
source_lines: [1, 74]
sha256: 2d43e7e763226ee5063c91d44cc4a8e04d1f884b8f2ae2d106fc4b356e65ae28
---

# itsmaxio-homelab-blob-head-docs-opnsense-wireguard-connecting-to-vps-md-b8cb41aa

OPNsense version: 25.1.2
This guide will help you configure WireGuard on an OPNsense firewall as a peer and a VPS as the server. The goal is to allow the VPS to access an server (10.10.20.2/24) behind OPNsense.
Run the following command on your OPNsense shell (SSH or console):
wg genkey | tee privatekey | wg pubkey > publickey
To view the keys:
cat privatekeycat publickey
Save the public key for the VPS configuration.
- Navigate to VPN > WireGuard.
- Go to the Local tab and click Add.
- Configure as follows:
  - Enabled: ✅
  - Name: VPS_Instance
  - Private Key: <OPNSENSE_PRIVATE_KEY> (fromprivatekey )
  - Public Key: <OPNSENSE_PUBLIC_KEY> (frompublickey )
  - Tunnel Address: 10.10.10.2/24
  - Click Save and Apply Changes.
- Go to the Peers tab and click Add.
- Configure as follows:
  - Enabled: ✅
  - Name: VPS
  - Public Key: <VPS_PUBLIC_KEY> (from the VPS generation below)
  - Allowed IPs: 10.10.10.1/32, 10.10.10.0/24
  - Endpoint Address: <VPS_PUBLIC_IP>
  - Endpoint Port: 51820
  - Persistent Keepalive: 25
  - Click Save and Apply Changes.
sudo apt update && sudo apt install wireguard -ywg genkey | tee privatekey | wg pubkey > publickey
Save the public key for the OPNsense peer configuration.
Edit the WireGuard config file:
sudo nano /etc/wireguard/wg0.conf
Add the following:
[Interface]
PrivateKey = <VPS_PRIVATE_KEY> #This is private key of VPS
Address = 10.10.10.1/24
ListenPort = 51820
[Peer]
PublicKey = <OPNSENSE_PUBLIC_KEY> #This is public key of OPNsense
AllowedIPs = 10.10.10.2/32, 10.10.20.0/24
#AllowedIPs = 10.10.10.2/32, 10.10.20.0/24, 10.10.30.0/24
#If you want you can add more subnets then VPS will be able to communicate with them. Of course you have to change the rules in OPNsense
PersistentKeepalive = 25sudo systemctl start wg-quick@wg0
sudo systemctl enable wg-quick@wg0
- Go to Firewall > Rules > WireGuard.
- Click Add and configure:
  - Action: Pass
  - Interface: WireGuard
  - Protocol: Any
  - Source: WireGuard Net
  - Destination: <WHERE YOU WANT>
  - Click Save and Apply Changes.
If your VPS has a firewall (UFW or cloud provider firewall), allow the WireGuard port:
sudo ufw allow 51820/udp
- Log into your Oracle Cloud Console.
- Navigate to Networking > Virtual Cloud Networks (VCN).
- Select your VCN and go to the Security Lists.
- Click on the default security list.
- Click Add Ingress Rule and configure:
  - Source: 0.0.0.0/0
  - Protocol: UDP
  - Port Range: 51820
  - Click Save Changes.
- Source: 
To allow the VPS to route traffic to the Ubuntu server behind OPNsense, enable IP forwarding:
echo "net.ipv4.ip_forward=1" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
- 
On OPNsense, check the WireGuard status: wg show
- 
From the VPS, ping the OPNsense WireGuard IP: ping 10.10.10.2
- 
From the VPS, ping the another subnet behind OPNsense: ping 10.10.20.2
- 
From the example server behind OPNsense, ping the VPS: ping 10.10.10.1
If you see a handshake in wg show and successful pings, the connection is working!
