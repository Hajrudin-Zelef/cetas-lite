---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-20680072882967-unifi-site-manager-e2d7b62c
title: "hc-en-us-articles-20680072882967-unifi-site-manager-e2d7b62c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-20680072882967-unifi-site-manager-e2d7b62c.md
source_anchor: ""
source_lines: [1, 33]
sha256: e4d322980d0358026b366f91d7b102e8614f663e77534ce021f25b2cc210bde9
---

# hc-en-us-articles-20680072882967-unifi-site-manager-e2d7b62c

UniFi Remote Management via Site Manager
UniFi Site Manager, available at unifi.ui.com, provides a centralized interface for remotely managing all UniFi sites you own or have administrative access to. It is designed to simplify operations across multiple locations, making it well suited for managed service providers and organizations with distributed deployments.
Getting Started
UniFi Site Manager is your personalized dashboard associated with your UI Account. All UniFi sites you own or have been granted administrative access to appear automatically, with no additional configuration required.
Remote management is enabled by default during UniFi setup, allowing sites to be securely accessed through Site Manager as soon as they come online. Once signed in, you can immediately view, manage, and monitor all of your deployments from a single interface.
Accessing UniFi Site Manager
To access Site Manager:
- Go to unifi.ui.com or download the UniFi Mobile App (iOS / Android).
- Sign in using your UI Account.
- Click the desired site to begin managing it.
Site Manager Features
Site Manager offers several advanced tools to streamline multi-site management:
- Site Magic SD WAN: Easily establish scalable, high-performance VPN connections between UniFi Gateways without the hassle of complex configurations or subnet management. Learn more here.
- ISP Viewer: Analyze key internet performance metrics, including latency, packet loss, and uptime, across all your deployments.
- Update Manager: Manage and initiate all updates across your sites from one location. Learn more here.
- Admin Management: Easily assign admin roles and permissions with this scalable solution.
- API Integration: The Site Manager API enables developers to monitor and manage UniFi deployments programmatically. Visit https://developer.ui.com/ to get started.
UniFi Fabrics
UniFi Fabrics is a management layer built on top of Site Manager that allow multiple UniFi sites to be grouped under a shared administrative and identity model. This facilitates management at scale through key features including centralized role-based access control (RBAC), Identity Provider (IdP) integration for zero-trust network access, policy & device templates, and more.
To learn more about UniFi Fabrics, check out this Academy topic.
Hybrid Cloud Architecture
Site Manager is part of UniFi’s Hybrid Cloud Architecture, which combines secure, license-free remote management with full local control. This architecture is designed for high availability: even if cloud connectivity or the internet is unavailable, sites remain fully manageable through direct local access.
UniFi automatically selects the most efficient connection method available. When possible, Site Manager establishes a direct connection to the UniFi Console for improved performance, while seamlessly falling back to cloud-assisted access when required.
For more information, see UniFi Architecture Overview.
Direct and Direct Remote Connections
When accessing a site through Site Manager while connected to the same local network, UniFi automatically establishes a Direct Connection to the UniFi Console. This provides a faster, low-latency experience for configuration changes and tasks such as viewing Protect video footage.
This same high-performance experience can be extended to remote access using Direct Remote Connection, which allows Site Manager to connect directly to the UniFi Console over the internet rather than relaying traffic through the cloud.
To enable Direct Remote Connection:
- Go to Control Plane > Console.
- Enable Direct Remote Connection.
Note: This requires a UniFi Gateway with a public IP and that TCP Port 443 that is not assigned to a port forwarding rule.
Privacy and Security
UniFi Site Manager is designed with privacy and security in mind. Your data is encrypted in transit, and remote management connections are protected by industry-standard protocols. For more details, refer to our Privacy Policy.
