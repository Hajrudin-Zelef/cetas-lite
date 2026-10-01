---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-5
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["EU"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [434, 499]
sha256: bbac75b5db7567b39ab6d615c4d6d7ee414d8875beef3e78a1640088b5ffd83e
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

With the release of OPNsense 24.1, the WireGuard plugin (os-wireguard) has been integrated into the core system. As a result, the separate plugin is no longer required or visible under installed plugins.
If you're seeing an error message such as os-wireguard (missing), there's no need to reinstall the plugin manually.
To resolve this issue, follow the steps below.
To resolve this issue and remove outdated plugin references, follow the steps below:
- 
Navigate to System → Firmware → Status on the OPNsense Web UI.
- 
Click Check for Updates, and wait for the list to refresh.
- 
If any system updates are available, apply them.
- 
(Optional) Reboot the system to ensure all integrated components are correctly loaded.
How to Resolve WireGuard Handshake Issue?
If you are encountering issues with the WireGuard handshake in OPNsense, the following checks and actions may help you diagnose and resolve the problem:
- 
Check Firewall Rules: Ensure that your firewall is configured to allow UDP traffic on port 51820, which is the default port used by WireGuard. Missing or misconfigured rules may prevent handshake packets from reaching the server.
- 
Verify VPN Keys: Double-check that the public and private key pairs are correctly configured on both the server and client sides. Each device should have the correct peer public key, and the keys must not be mismatched.
- 
Check IP Address and DNS Configuration: If you're using a dynamic DNS (DDNS) service, make sure it is correctly set up and actively resolving to your current public IP address. An outdated or incorrect DNS record can prevent a successful handshake.
- 
Restart Services and Network Devices: Sometimes, a simple restart can fix the issue. Try restarting the WireGuard service on both the server and the client. Additionally, rebooting your router, modem, or firewall may help refresh network routes and fix handshake failures.
Use the Status page in OPNsense’s WireGuard interface to verify if packets are being received. If the "Handshake" field is empty or doesn't update, it’s an indication that something is blocking the connection.
How do I Verify the Installation of WireGuard on OPNsense?
After updating OPNsense or setting up WireGuard, it’s essential to verify that the plugin is correctly installed and operational. The steps below outline how to perform this verification via the web interface and, optionally, using the CLI.
- 
Log in to the OPNsense Web UI: Open a browser and navigate to the IP address of your OPNsense firewall. For example: https://192.168.56.2. Log in using your administrative credentials (e.g., root user).
- 
Check WireGuard Menu Availability: On the left-hand menu, go to VPN → WireGuard. Verify that the following tabs are visible. 
  - Instances
  - Peers
  - Peer Generator
  - Status
  - Log File
 The presence of these menu items confirms that the WireGuard component is integrated into the system.
- 
Confirm WireGuard Status on Web UI: Navigate to VPN → WireGuard → Status page. You should see information about running WireGuard instances and active peers. 
  - If a tunnel is properly configured and active, the instance status will appear as UP.
 tipThe Status tab also shows connection statistics such as handshake time and data transfer details.
- 
Peer Generator Verification: Navigate to VPN → WireGuard → Peer Generator. Confirm that the configuration fields, public/private key details, and generated QR code are displayed correctly. This feature allows quick peer configuration for mobile or desktop clients.
If all the above components are visible and operational, your WireGuard installation on OPNsense is correctly set up. If you cannot see the menu options or the system behaves unexpectedly, consider rebooting the firewall or reviewing system logs via System → Log Files.
What are Best Practices for Managing WireGuard Keys and Configurations?
Managing WireGuard VPN keys and configurations effectively is essential for ensuring the security, stability, and scalability of your VPN setup. Below are industry-recognized best practices you should follow for WireGuard VPN key management.
- 
Generate Unique Key Pairs for Each Peer: Each peer (device) should have its own unique public and private key pair to maintain proper identification and avoid potential security risks. Sharing or duplicating key pairs between peers weakens traceability and increases vulnerability. tipUse wg genkey andwg pubkey on each client device to generate independent credentials.
- 
Securely Store Private Keys: Private keys are extremely sensitive and must be stored in secure locations with restricted access. They should never be shared or transmitted through unsecured channels. Store them in secure file systems with limited access rights ( chmod 600 ) or within secure vault solutions (e.g., HashiCorp Vault, Bitwarden).noteNever paste private keys into unsecured platforms (email, chat apps, etc.)
- 
Rotate Keys Regularly: Regular key rotation is essential to minimize risks due to key exposure, outdated configurations, or compromised devices. Rotating keys ensures your WireGuard network remains secure and up-to-date. Rotate keys every 90 days and re-share updated configuration files or QR codes with clients. tipAfter rotation, be sure to update configuration files or QR codes on both the server and client sides to reflect the new key pairs.
- 
Use Pre-Shared Keys (PSKs) for Extra Security: In addition to using public and private key pairs, it is highly recommended to configure a Pre-Shared Key (PSK) to add an extra layer of encryption and protect against potential brute-force or replay attacks. In OPNsense, the PSK can be generated and configured directly within the VPN → WireGuard → Peers section. Simply create or edit a peer, and use the Generate button next to the Pre-shared key field. Once generated, this PSK must be applied on both the server and the client to ensure mutual authentication.
- 
Backup Configuration Files Securely: To prevent data loss and ensure fast recovery in case of system failure or misconfiguration, it's crucial to maintain encrypted backups of both server and peer configuration files. Sensitive configuration files should be backed up using encrypted ZIP archives or GPG encryption, and these backups must be stored in secure locations such as zero-knowledge cloud storage (e.g., Tresorit, Sync.com), local NAS systems with restricted access, or offline USB drives.
- 
Restrict Allowed IPs for Each Peer: Limiting each peer’s access using the Allowed IPs field is essential for minimizing network exposure and enforcing strict access controls within your VPN infrastructure. For instance, specifying Allowed IPs: 192.168.100.0/24 grants the client access only to the local LAN segment.Avoid using 0.0.0.0/0 unless full tunneling (redirecting all traffic through the VPN) is explicitly required, as it may lead to unintended access or traffic leakage.
- 
Label and Document Each Peer Configuration Clearly: Using meaningful and descriptive labels for each peer configuration significantly simplifies management, especially in larger environments. Clearly labeled peers make debugging and monitoring easier, improve team communication, and help track device ownership and intended purpose over time. Examples of good labeling practices are visible below. 
  - 
MarketingLaptop_Anna: A laptop used by the marketing team member Anna
  - 
DevServer_EU: A development server located in the EU region
  - 
HomeTablet_Merve: A personal device used for remote access
- 
- 
