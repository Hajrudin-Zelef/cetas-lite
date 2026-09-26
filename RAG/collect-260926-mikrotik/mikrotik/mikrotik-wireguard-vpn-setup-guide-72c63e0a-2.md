---
id: collect-260926-mikrotik/mikrotik/mikrotik-wireguard-vpn-setup-guide-72c63e0a-2
title: "MikroTik WireGuard VPN Setup Guide"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["lean"]
source: docs/RAG/lot-mikrotik/forum/wireguard/mikrotik-wireguard-vpn-setup-guide-72c63e0a.md
source_anchor: ""
source_lines: [130, 181]
sha256: d419f56be02364d1a240e0b349011a799d7d4dab12c94ed4d1f1569e1d11d323
---

# MikroTik WireGuard VPN Setup Guide

Configuring your MikroTik router to connect through a VPN enhances privacy and security. Setting up a WireGuard VPN client on MikroTik RouterOS 7 can be straightforward if you follow precise instructions. WireGuard offers a simpler and faster alternative to traditional VPNs like OpenVPN. Its lean design enables users to establish secure point-to-point connections for remote offices and roadwarrior scenarios such as connecting from hotels or coffee shops.

### Configuring WireGuard on Client Devices

When setting up WireGuard as a VPN client, you must first ensure your MikroTik router runs on RouterOS 7 or later. It's crucial to gain access to the router's admin interface via Winbox or a similar management tool. Create a virtual interface for the WireGuard connection on your MikroTik router. This involves specifying **Allowed IPs** which dictates the traffic allowed through the VPN.

For instance, if you're configuring a roadwarrior setup, create a peer representing the VPN server. In this peer's settings, you'll need to add the server's public key and set the **Endpoint** to the server's IP address and port.

On client devices such as Windows, Linux, Android, or iOS, installing the WireGuard application available from their respective app stores or websites is needed. After installing the application, import the client configuration file provided by your MikroTik router or manually input the necessary **Peer** information, including public and private keys, and set **Allowed IPs**.

During setup, Windows users might prefer to use port 443 as it's typically not blocked by firewalls, whereas MikroTik's default port is 13231. Remember that WireGuard operates over UDP protocol, renowned for its speed.

Finally, ensure DNS servers are configured on the client devices for proper name resolution when connected to the VPN.

### Troubleshooting Connection Issues

If you encounter connection problems after setting up your WireGuard VPN client, go through a systematic process to identify and resolve them.

1. **Ping Test**: Check if your MikroTik router can ping the WireGuard server's IP address.
2. **Cryptography Check**: Verify the cryptographic keys (public and private keys) to ensure they are correctly configured on both the client and server ends.
3. **Persistent Keepalive**: If your client is behind a NAT, setting the **persistent-keepalive** may be necessary to maintain the connection.
4. **Firewall Settings**: Inspect any firewall rules that may be blocking VPN traffic and adjust them as needed.
5. **Client Configuration**: Re-examine the client config to ensure the endpoint IP addresses, allowed subnets, and keys match those expected by the server.
6. **Subnet Overlap**: Confirm there's no overlap between your local LAN subnets and the VPN's internal subnets which might cause routing conflicts.

If issues persist, consulting the detailed MikroTik logs can provide further clues. It's also beneficial to seek support from the MikroTik community or professional networking consultants if you're in over your head.

## Frequently Asked Questions

### How do I configure a WireGuard server on MikroTik for secure VPN connections?

To set up a WireGuard server on MikroTik, ensure you're running RouterOS 7 or newer. Enable the WireGuard package and create a WireGuard interface. Assign it an IP address and configure peers with their public keys and allowed IPs.

### What are the steps for setting up a MikroTik router as a WireGuard VPN client?

For a MikroTik router to act as a WireGuard client, first update RouterOS to version 7.6 or higher. Next, establish a WireGuard interface and specify the VPN server details. Add the server's public key to the peer list and set up the allowed IPs to secure the connection.

### Can I establish a site-to-site VPN with WireGuard on MikroTik routers, and if so, how?

Yes, site-to-site VPNs are viable using WireGuard on MikroTik routers. Start by configuring WireGuard interfaces on both routers. Set up each as a peer on the other, ensuring proper routing of IP addresses between sites for a secure direct connection.

### What is the process for configuring WireGuard VPN on MikroTik for mobile devices such as Android and iOS?

To configure WireGuard VPN for mobile devices, create a dedicated WireGuard interface on your MikroTik router. Provide mobile devices with configuration details, including the server's public key and assigned IP. Use the respective device's WireGuard app to finalize the setup.

### Is there a tutorial for implementing a WireGuard road warrior VPN setup on MikroTik?

For a road warrior setup, which allows remote users to connect to a network, create a WireGuard interface on your MikroTik router. Configure peers for each remote user, providing them with necessary credentials. Detailed guides are available on the MikroTik documentation site.

### Where can I find the official download for the MikroTik WireGuard package?

The official MikroTik WireGuard package can be downloaded from the MikroTik website. Navigate to the download section and select the relevant RouterOS version. The package is typically enabled by default in RouterOS version 7 and above.
