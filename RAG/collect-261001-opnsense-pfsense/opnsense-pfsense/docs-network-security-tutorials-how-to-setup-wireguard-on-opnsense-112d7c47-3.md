---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-3
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["kill switch"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [180, 316]
sha256: d653a0e82bd6e6fdae05941f75c501d54687f2afb19bb02b531e850f67af1af2
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

This automatically generates a public and private key pair for the client and displays them on the screen. tipNote down the Public Key. You will need it later to add the Windows client as a peer on your OPNsense server. Figure 10. Generated key pair for the new tunnel
- 
Enter a name with alphanumeric characters only (no spaces or punctuation) for the tunnel, such as MyWireGuard, and edit the configuration as follows. [Interface]
 PrivateKey = CLIENT_PRIVATE_KEY
 Address = 10.0.0.11/24
 DNS = 10.0.0.1
 [Peer]
 PublicKey = SERVER_PUBLIC_KEY
 Endpoint = SERVER_IP_ADDRESS:51820
 AllowedIPs = 0.0.0.0/0
Explanations of the fields in the interface section are given below. [Interface] 
  - 
PrivateKey: The private key of the Windows client (auto-generated)
  - 
Address: VPN IP address for this client. It must be unique (e.g., 10.0.0.11/24)
  - 
DNS: The DNS server used inside the tunnel, typically the WireGuard server (e.g., 10.0.0.1)
 Explanations of the fields in the peer section are given below. [Peer] 
  - 
PublicKey: The public key of the WireGuard server on OPNsense (e.g., fyKJ4c6sXTVRTJla6zQ9wi4okRPRd/GsMbTMszjhAgA=)
  - 
Endpoint: Server’s public IP address and WireGuard port (e.g., 203.0.113.1:51820)
  - 
AllowedIPs – 0.0.0.0/0 means route all traffic through the VPN tunnel
- 
3.3. Block untunneled traffic (kill switch) option
In the Edit Tunnel window of the WireGuard client, you can enable the Block untunneled traffic (kill-switch) option. This setting adds Windows Firewall rules to block all traffic that is not routed through the VPN tunnel.
You should enable this option only if:
- There is exactly one [Peer] section in your configuration.
- The AllowedIPs field is set to a catch-all address such as 0.0.0.0/0.
When enabled, this option ensures that no network traffic bypasses the VPN tunnel, preventing accidental IP leaks and protecting your privacy.
Figure 11. WireGuard Tunnel configuration on Windows client
Once completed, click the Save button to apply the settings.
Step 4. Adding WireGuard Endpoint (Client Peer) Configuration to the Server
To add the client’s public key and allowed IP to the WireGuard server on OPNsense, follow these steps below.
- 
Navigate to VPN → WireGuard → Peers on the OPNsense Web UI.Figure 12. Adding WireGuard endpoint configuration on OPNsense
- 
Click on the "+" button to add a new peer.
- 
In the Edit peer dialog that appears, enter the following details. 
  - 
Enabled: Ensure this box is checked.
  - 
Name: Choose a descriptive name for the peer, such as MyWindows.
  - 
Public Key: Enter the public key generated during the WireGuard client setup on Windows. (Example: ZzBdrHD0m3vmCYkhZz80ujXJD8vsnqnw34tHqNO0SkQ=)
  - 
Pre-shared Key: Leave this field empty unless you’ve configured a pre-shared key on both sides.
  - 
Allowed IPs: Enter the client IP address you assigned in the WireGuard client configuration (use CIDR notation). (Example: 10.0.0.11/32)
  - 
Endpoint Address: Optional. You may leave this empty for dynamic clients.
  - 
Endpoint Port: Optional. Default is usually 51820, or match what’s configured on the client.
  - 
Instances: Select the WireGuard instance you want to associate this peer with.
  - 
Keepalive Interval: Optional. Recommended to set this to 25 seconds if the client is behind NAT.
- 
- 
Click Save to apply the peer configuration. Figure 13. Setting WireGuard Endpoint(Windows) configuration on OPNsense
Step 5. Configuring WireGuard Android Client
You can easily configure the WireGuard application on your Android device by following the steps below.
5.1. Download and install WireGuard Application on Android device
You can download and install the official WireGuard app from the Google Play Store.
Figure 14. Installing the WireGuard Android application from the Play Store
5.2. Configuring WireGuard Client on Android
Once the WireGuard application is installed, follow these steps to create a VPN tunnel:
- 
Tap the blue + icon in the bottom-right corner. Figure 15. Adding a new WireGuard tunnel
- 
Select Create from Scratch to manually configure a new tunnel. Figure 16. Creating a tunnel configuration from scratch
- 
Set a Name for your tunnel, e.g., MyWireGuard.
- 
Tap the recycle icon next to the "Private Key" field to auto-generate a key pair.
- 
Set the Address, for example: 10.0.0.12.
- 
Set the DNS, for example: 10.0.0.1. Figure 17. Android WireGuard client interface configuration
- 
Tap Add Peer at the bottom to configure the peer settings. Figure 18. Adding peer information to the Android client
- 
Copy the OPNsense WireGuard Server's Public Key from the configuration on OPNsense and paste it into the Public Key field in the peer section.
- 
Set the Endpoint to the public IP address of the OPNsense server, followed by a colon : and the WireGuard port number. Example: 198.51.100.1:51820
- 
Set Allowed IPs to 0.0.0.0/0. This means all traffic will be routed through the VPN tunnel.
- 
Tap the floppy disk icon (top-right corner) to save the tunnel configuration.
Explanations of the fields in the interface section are given below:
- 
PrivateKey: The private key for this Android client.
- 
PublicKey: Automatically generated. This must be copied into the Peer configuration on the OPNsense server.
- 
Address: The unique IP address for this client. It should match the AllowedIPs configured in OPNsense.
- 
DNS: The IP address of a DNS server (typically the WireGuard server itself, e.g., 10.0.0.1).
Explanations of the fields in the peer section are given below:
- 
PublicKey: The public key of the OPNsense WireGuard server.
- 
Endpoint: The public IP address of the OPNsense server, followed by a colon : and the WireGuard listen port (e.g., 51820).
- 
Allowed Ips: Use 0.0.0.0/0 to route all traffic through the VPN tunnel.
If you are using a Ubuntu desktop that needs a VPN connection, refer to the WireGuard Installation Guide for configuration steps specific to Linux clients.
Step 6. Adding WireGuard Endpoint (Android Client Peer) Configuration to the Server
To allow the Android client to connect to the VPN, its public key and IP address must be added as a new endpoint on the OPNsense WireGuard server. Follow these steps:
- 
Go to VPN → WireGuard → Peers on the OPNsense web interface.
- 
Click the + button to add a new endpoint.
- 
Enter a descriptive Name, such as MyAndroid.
- 
In the Public Key field, paste the public key generated on the Android WireGuard client. Example: rQdjEcn7UMbIverQ4D0FKfz+fkGLxClArwDsXCNf+DE=
- 
Set Allowed IPs to the IP address assigned to the Android client in CIDR format. Example: 10.0.0.12/32
- 
Click Save to apply the settings.
Once saved, you can view all configured WireGuard VPN endpoints (e.g., MyWindows, MyAndroid) under the VPN → WireGuard → MyAndroid tab on OPNsense.
Figure 19. Setting WireGuard Endpoint(Android) configuration on OPNsense
Step 7. Adding Peers(VPN Clients) to Server Configuration on OPNsense
Once you have defined the VPN clients as Endpoints in OPNsense (e.g., MyWindows and MyAndroid), the next step is to add them as Peers to the WireGuard server configuration.
This step is essential because defining an Endpoint alone only registers the client's public key and IP address, it does not yet associate the client with the server instance.
By adding the clients as Peers to the configuration, you explicitly tell the WireGuard server which endpoints are allowed to connect and exchange traffic. This enables mutual communication between the server and each configured VPN client.
Follow the steps below to add the peers to the server's configuration.
- 
Navigate to VPN → WireGuard → Instances in the OPNsense web interface.
- 
Locate your WireGuard server instance (e.g., MyWireGuard) and click the Edit (pencil) icon.
- 
In the **Peers dropdown menu, select the endpoints you have previously created, in this example, select MyWindows and MyAndroid.
- 
Click Save to update the configuration.
