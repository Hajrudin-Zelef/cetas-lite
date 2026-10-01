---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-115012240067-enabling-unifi-remote-management-623ec7b4
title: "hc-en-us-articles-115012240067-enabling-unifi-remote-management-623ec7b4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-115012240067-enabling-unifi-remote-management-623ec7b4.md
source_anchor: ""
source_lines: [1, 18]
sha256: 9e4dcdf850890e31739ce07b562d10091dc45f83d709417757567b1a33e1b2da
---

# hc-en-us-articles-115012240067-enabling-unifi-remote-management-623ec7b4

Enabling UniFi Remote Management
Remote Management allows you to manage all your UniFi deployments through the UniFi Site Manager, available at unifi.ui.com. This serves as a central hub, especially convenient for Managed Service Providers or Enterprise organizations with a vast geographical footprint. To learn more, see our article on UniFi Remote Management via Site Manager.
Configuring Remote Management
Remote management is enabled by default during the setup process for all UniFi deployments. If it is not enabled, follow these steps to manually activate it:
- Create a UI Account: Sign up at UI Account or log in with your existing account.
- Connect Locally to Your UniFi Site: For instructions, refer to UniFi Local Management.
- Navigate to Settings > Control Plane > Console.
- Check Remote Management.
Troubleshooting
Here are common issues that may prevent remote management and their solutions:
- 
Blocked Outbound Traffic on TCP 443 and 8883:
 UniFi Remote Management requires outbound connections on TCP ports 443 and 8883. You do not need to open these ports for inbound traffic. However, if outbound traffic is blocked—by local firewalls (e.g., Windows Firewall), routers, or ISP equipment—Site Manager will not be able to reach your UniFi Console, including self-hosted instances.
 Most networks allow this traffic by default, but issues can occur if outbound traffic is restricted or filtered. To ensure connectivity, allow outbound connections on TCP ports 443 and 8883.
For a complete list of required ports see here. If necessary, contact your ISP for assistance.
- DNS Resolution Issues: Use a public DNS server like 8.8.8.8 or 1.1.1.1 for reliable resolution. If you’re using internal DNS servers, ensure UDP Port 53 is open and not blocked by any firewalls or ISP modems.
- Incorrect Date/Time Settings: Ensure your console’s date and time are accurate. Incorrect settings are often caused by blocked UDP Port 123. Contact your ISP if you cannot resolve this.
- Outdated UniFi OS Version: Update to the latest UniFi OS version for improved stability and performance. Check Community Releases for the latest updates.
