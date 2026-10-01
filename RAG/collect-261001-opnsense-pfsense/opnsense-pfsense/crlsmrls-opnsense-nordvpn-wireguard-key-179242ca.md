---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/crlsmrls-opnsense-nordvpn-wireguard-key-179242ca
title: "crlsmrls-opnsense-nordvpn-wireguard-key-179242ca"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/crlsmrls-opnsense-nordvpn-wireguard-key-179242ca.md
source_anchor: ""
source_lines: [1, 65]
sha256: e7d40d9e48614887c37e475d169757be6be7b8ed2c1836d70f721306eeca9c8e
---

# crlsmrls-opnsense-nordvpn-wireguard-key-179242ca

This document provides a comprehensive guide on configuring a WireGuard (NordLynx) interface in OPNsense using a NordVPN subscription, including steps to obtain necessary credentials and configure the connection.
Versions:
- OPNsense 25
- NordVPN Active Subscription (2026)
Neither NordVPN nor OPNsense provide an official unified guide for this configuration.
- 
NordVPN offers WireGuard under the name "NordLynx," their custom implementation for enhanced speed and privacy. However, unlike other providers, they do not provide .conf configuration files nor display the Private Key in the user dashboard, which complicates manual configuration on third-party routers.
- 
OPNsense has native WireGuard support but lacks a specific guide for NordVPN. Furthermore, the official WireGuard documentation is often outdated regarding recent user interface changes (such as the terminology shift to "Instances" and "Peers" in version 25.7).
Before configuring OPNsense, you need to obtain your credentials and keys via a terminal.
For some reason, NordVPN hides this; it is possible they only want us to use their own applications.
Navigate to "Access Tokens" and create a new one with permissions to "Get service credentials".
- NordAccount >
- NordVPN >
- Access token >
- You will need to verify your email address.
- Create a new token with the "Get service credentials" permission. Let's assume the generated token is e9f2___X___66 .
- The next step is to use that token to obtain your WireGuard Private Key. Run this command in your terminal (replace YOUR_TOKEN_HERE with your actual token):
curl -s -u token:YOUR_TOKEN_HERE https://api.nordvpn.com/v1/users/services/credentials | jq -r .nordlynx_private_key
Let's say it returns: Vnn______________fMQ= for future reference.
Note: You must have jq installed to process JSON. https://jqlang.org/download/
The command above returns your WireGuard Private Key.
Run this to obtain the IP and Public Key of the best server, in my case in Spain (ID 202):
curl -s "https://api.nordvpn.com/v1/servers/recommendations?filters\[country_id\]=202&limit=1" \
  | jq -r '.[0] | {hostname: .hostname, ip: .station, public_key: (.technologies[] | select(.identifier=="wireguard_udp") | .metadata[] | select(.name=="public_key") | .value)}'
It returns something like this:
{
  "hostname": "es225.nordvpn.com",
  "ip": "185.214.97.110",
  "public_key": "OaSGa___________AXY="
}
To create a WireGuard connection, it is necessary to create an "Instance" (your connection) and a "Peer" (the remote server).
Go to VPN > WireGuard > Instances and add a new one:
- Name: NordVPN_Local
- Private Key: The key Vnn______________fMQ= that you obtained in step 1.1.
- Public Key: Leave empty; it calculates automatically.
- Listen Port: 51820
- DNS Server: 103.86.96.100 (NordVPN DNS).
- Tunnel Address: 10.5.0.2/32 (Internal IP to be used as NordLynx Gateway).
- Disable Routes: CHECKED ✅ In my case, I do not want all router traffic to pass through the VPN, only that of a specific VLAN. If you want ALL router traffic to pass through the VPN, leave unchecked.
- Save
- Apply
Go to VPN > WireGuard > Peers and add a new one:
- Name: NordVPN_ES_225 (In this case, I want to identify the Spain server 225).
- Public Key: The server's public_key obtained in step 1.2 (e.g.,OaSGa___________AXY= ).
- Endpoint Address: The numeric ip of the server obtained in step 1.2 (e.g.,185.214.9.110 ).
- Endpoint Port: 51820
- Allowed IPs: 0.0.0.0/0 (Allow all internet traffic).
- Keepalive Interval: 25 (Set persistent keepalive interval in seconds).
- Save
- Apply
Go to VPN > WireGuard > Instances and link the remote server to your connection:
- Edit the NordVPN_Local instance.
- In the Peers tab, add the peer NordVPN_ES_225 .
- Save
- Apply
Go to VPN > WireGuard > Peers:
- Edit the peer NordVPN_ES_225 , click the pencil icon.
- Enable: Check ✅ This will enable the peer.
- Save
- Apply
Go to VPN > WireGuard > Status:
It must show the interface and the peer active. Both with a green check ✅.
OPNsense should have created the wg0 interface automatically; the interface and the peer use this interface. If everything went well, the peer shows TX/RX traffic and the duration since the last "Handshake".
In case of issues, check the logs in VPN > WireGuard > Log File.
