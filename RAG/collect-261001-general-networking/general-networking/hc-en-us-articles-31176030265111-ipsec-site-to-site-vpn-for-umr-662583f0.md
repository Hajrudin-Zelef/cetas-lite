---
id: collect-261001-general-networking/general-networking/hc-en-us-articles-31176030265111-ipsec-site-to-site-vpn-for-umr-662583f0
title: "hc-en-us-articles-31176030265111-ipsec-site-to-site-vpn-for-umr-662583f0"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/hc-en-us-articles-31176030265111-ipsec-site-to-site-vpn-for-umr-662583f0.md
source_anchor: ""
source_lines: [1, 130]
sha256: 66d3ed193e0dd608f34fef79c9fb64d85b0dfc41c179e68f3a011b1e80ebfdec
---

# hc-en-us-articles-31176030265111-ipsec-site-to-site-vpn-for-umr-662583f0

IPSec Site-to-Site VPN for UMR
IPSec Site-to-Site VPN on UMR enables secure, encrypted communication between remote networks using standard IPSec protocols. This is especially useful when integrating with third-party infrastructure such as firewalls, routers, or cloud VPN gateways (e.g., AWS, Azure), as well as legacy enterprise systems.
Important: IPSec Site-to-Site VPN (as well as in Wireguard Site-to-Site VPN) supports bidirectional communication between networks. Devices on the remote LAN can access the UMR LAN and vice versa
Note: IPSec is not recommended for UniFi device onboarding or camera adoption. WireGuard is better suited for such use cases.
If a device has a private IP address due to being NAT’d, follow the corresponding setup scenario below. Port forwarding may be required.
|  | UMR address | Remote address | Notes | 
| Scenario 1 | Public | Public | Simplest case. No port forwarding or ID changes needed | 
| Scenario 2 | CG-NAT | Public | Remote must allow connections from any IP. | 
| Scenario 3 | Public | NAT | Remote must have fixed external IP. Remote must accept connections from the UMR address. | 
| Scenario 4 | CG-NAT | NAT | Remote must forward UDP 500/4500. IDs must be set manually. Requires remote to accept connections from unknown addresses. | 
IPSec Site-to-Site VPN Configuration on UMR
UMR supports IPSec VPN in multiple NAT/public IP setups. Choose the scenario that matches your network and make the UMR firmware version is at least 1.15.0 or later
Scenario 1
Scenario 2
Scenario 3
Scenario 4
Scenario 1: Both Devices Have Public IPs
Prerequisites:
- Static public IPs on both UMR and remote peer.
Settings
- Local ID: Auto*
- Remote ID: Auto*
*Use Manual ID if connection with Auto is not successful or not supported by remote end. See more in How to configure Manual ID
Scenario 2: UMR behind CG-NAT, Remote Peer has Public IP
Use Case: This is a common deployment for mobile or SIM-based UMR devices. The remote peer must have a real public IP and accept incoming connections from any IP address.
Prerequisites:
- UMR is behind CG-NAT and cannot accept incoming connections.
- Remote peer must have a public IP and allow connections from unknown IPs.
UMR (CG-NAT side) Settings
- Remote IP: Public IP of the remote peer
- Local ID: Auto*
- Remote ID: Auto*
*Use Manual ID if connection with Auto is not successful or not supported by remote end. See more in How to configure Manual ID
Remote Peer (Public IP side) Settings
- Local ID: Auto
- Remote ID: Auto
- Must accept IPSec connections from unknown (dynamic) IPs.
Note: Port forwarding is not possible on the CG-NAT UMR side. This scenario relies entirely on UMR initiating the tunnel.
Scenario 3: UMR Has Public IP, Remote Peer Behind NAT (RARE)
Use Case: As an example in some rare cases there is a headquarters VPN server behind NAT. A remote UMR device wants to connect to the headquarters network. At the same time, UMR has a public IP address.
Prerequisites:
- UMR must have a static public IP
- Remote peer is behind NAT
- Remote peer should have a fixed external IP address or use Dynamic DNS
UMR (Public IP side) Settings
- Local ID: Auto*
- Remote ID: Auto*
- Remote IP: Set to external remote IP
*Use Manual ID if connection with Auto is not successful or not supported by remote end. See more in How to configure Manual ID
Remote Peer (NAT-side) Settings
- Port forwarding for UDP 500 and 4500 to internal IPSec service is mandatory if remote end cannot initiate IPSec connection
- Local ID: Auto
- Remote ID: Auto
- External IP address must remain stable and known to UMR (static or DDNS)
Note: The VPN connection may fail if the remote peer does not support initiating the tunnel or if port forwarding is not properly configured for UDP ports 500 and 4500
Scenario 4: UMR behind CG-NAT, Remote peer behind NAT
Prerequisites:
- Remote peer (NAT side) must have port forwarding for UDP 500 & 4500.
- Manually configure Local and Remote IDs on both ends (e.g. 1.2.3.4 ↔ 4.3.2.1).
UMR (CG-NAT side) Settings
- Local ID: Manual
- Remote ID: Manual
- Remote IP: Public IP of the remote peer
* Manual ID required. See more in How to configure Manual ID
Remote Peer (NAT-side) Settings
- Port forwarding for UDP 500 and 4500 to internal IPSec service is mandatory
- Local ID: Manual
- Remote ID: Manual
- External IP address must remain stable and known to UMR (static or DDNS)
- NAT-T: Enable if available
- Must accept IPSec connections from unknown (dynamic) IPs.
How to configure Manual ID?
- Pick two random IDs looking like IP addresses. For example:
  - 1.2.3.4 and 2.3.4.5
  - 192.168.1.1 and 192.168.1.2
  - May or may not coincide with the actual IP addresses in use
- Configure UMR:
  - Local ID: for example, 1.2.3.4
  - Remote ID: for example, 2.3.4.5
- Configure remote end:
  - Local ID: for example, 2.3.4.5
  - Remote ID: for example, 1.2.3.4
Step-by-Step IPSec VPN Setup on UMR
Step 1: Access VPN Settings
- Log in to Mobility Manager.
- Go to Mobile Routing > Settings > VPN > Site-to-Site VPN.
- Select VPN Type: IPSec.
- Enter a VPN Name.
- Under Device Assignment, select the UMR (only one device per IPSec profile).
Step 2: Configure VPN Settings
- Pre-shared Key: Used for authentication.
- Remote IP: Public IP of the remote gateway.
- Remote Network(s): List remote subnets (e.g., 192.168.1.0/24)
- Restriction: Subnet masks from /8 to /32 are allowed.
Step 3: (Optional) Advanced IPSec Settings
IKE Settings
- Key Exchange Version: IKEv1 or IKEv2
- Encryption: AES-128 / AES-256
- DH Group: 14 (recommended)
- Lifetime: Default 28,800 sec
ESP Settings
- Encryption: AES-128 / AES-256
- Hash: SHA1 or SHA2
- DH Group: 14
- Lifetime: Default 3,600 sec
- Perfect Forward Secrecy (PFS): Enable
Other Settings
- Local Authentication ID – Keep "Auto" in most cases. Only manually configure if necessary, ensuring it matches the Remote Authentication ID set on the remote site.
  - Allowed values: IPv4 address.
- Remote Authentication ID – Keep "Auto" in most cases. Only manually configure if necessary, ensuring it matches the Local Authentication ID set on the remote site.
  - Allowed values: IPv4 address.
- Maximum Transmission Unit (MTU) – Defines the maximum packet size allowed over the VPN tunnel.
  - Recommended: Keep the default 1419 bytes unless specific changes are required.
- Ensure consistency: MTU should match the value configured on the remote site to prevent connectivity loss or excessive fragmentation.
Step 4: Apply & Verify
- Click "Apply Changes" to establish the VPN connection.
- Monitor the VPN status from: Devices > [Select UMR] > Settings > VPN
- In order to set up a successful VPN, the following information needs to match between the gateways:
  - VPN Protocol
  - Pre-shared Key
  - Remote and local server IP address
  - Remote and local subnets
  - Key Exchange Version, Encryption, Hash, and DH Groups
- Perfect Forward Secrecy
Best Practices
- Ensure both VPN endpoints use matching IKE and ESP settings.
- Use at least AES-256 with SHA256 for enhanced security in enterprise environments.
- Choose the strongest DH group supported by both UMR and the remote site.
- IKEv2 and PFS are recommended.
- Monitor VPN status in Mobile Routing to check for disconnections or key negotiation failures.
