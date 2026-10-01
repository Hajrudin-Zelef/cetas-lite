---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-11444786290071-unifi-remote-management-via-site-manager-e12b67a8
title: "hc-en-us-articles-11444786290071-unifi-remote-management-via-site-manager-e12b67a8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-11444786290071-unifi-remote-management-via-site-manager-e12b67a8.md
source_anchor: ""
source_lines: [1, 29]
sha256: 28dd780aa5cfbc0bc66d57086f3e49068bd5ae70c4ad2acfc80315eec8754d60
---

# hc-en-us-articles-11444786290071-unifi-remote-management-via-site-manager-e12b67a8

UniFi Remote Management via Site Manager
The UniFi Site Manager, located at unifi.ui.com, provides a centralized platform for managing all your deployments remotely. With Site Manager, you can access and administer all sites you own or have been granted administrative permissions to from a single interface. This feature is ideal for scaling management across multiple locations or ensuring seamless control, even during routine operations. This is designed around managed service providers and large enterprises with globally-dispersed locations.
Accessing UniFi Site Manager
To access Site Manager:
- Go to unifi.ui.com or download the UniFi Mobile App (iOS / Android).
- Sign in using your UI Account.
- Click the desired site to begin managing it.
Note: Remote Management is enabled by default during initial setup. If you need to enable it manually, click here.
Site Manager Features
Site Manager offers several advanced tools to streamline multi-site management:
- Site Magic SD WAN: Easily establish scalable, high-performance VPN connections between UniFi Gateways without the hassle of complex configurations or subnet management. Learn more here.
- ISP Viewer: Analyze key internet performance metrics, including latency, packet loss, and uptime, across all your deployments.
- Update Manager: Manage and initiate all updates across your sites from one location. Learn more here.
- Admin Management: Easily assign admin roles and permissions with this scalable solution.
- API Integration: The Site Manager API enables developers to monitor and manage UniFi deployments programmatically. Visit https://developer.ui.com/ to get started.
UniFi Fabrics
UniFi Fabrics are a management layer built on top of Site Manager that allow multiple UniFi sites to be grouped under a shared administrative and identity model. Fabrics enable centralized role-based access control (RBAC), Identity Provider (IdP) integration for zero-trust network access, and policy templatization, allowing administrators to define identity, access, and configuration policies once and apply them consistently across all associated sites.
For more information, check out this UniFi Academy Topic.
Direct Connections
A Direct Connection automatically activates when accessing a site via Site Manager while connected to the same local network. This provides a faster, low-latency experience for managing settings and viewing Protect video footage.
Direct Remote Connections
To achieve the same high-performance experience for remote connections:
- Go to Control Plane > Console.
- Enable Direct Remote Connection.
Note: This requires a UniFi Gateway with a public IP and that TCP Port 443 that is not assigned to a port forwarding rule.
Privacy and Security
UniFi Site Manager is designed with privacy and security in mind. Your data is encrypted in transit, and remote management connections are protected by industry-standard protocols. For more details, refer to our Privacy Policy.
Offline (Local-Only) Management
For guidance on managing UniFi deployments locally—whether during internet outages or for air-gapped setups—visit UniFi Local Management.
