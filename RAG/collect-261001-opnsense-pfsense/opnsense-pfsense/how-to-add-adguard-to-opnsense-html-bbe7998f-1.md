---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f-1
title: "how-to-add-adguard-to-opnsense-html-bbe7998f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f.md
source_anchor: ""
source_lines: [1, 37]
sha256: 538911dca1b7d66e3ecbcb52729a9365bfb72b7a6b10dddefac947bdf5f6da53
---

# how-to-add-adguard-to-opnsense-html-bbe7998f

Free tools Windows power users keep installed
One-click scans. No signup required.
AdGuard Home is a powerful network-wide ad blocker that enhances privacy and security by blocking unwanted content at the DNS level. Integrating AdGuard with OPNsense, an open-source firewall platform, allows you to leverage its filtering capabilities directly within your network infrastructure. This setup not only improves browsing experience by eliminating ads but also protects against malicious domains and trackers. Implementing AdGuard on OPNsense is a straightforward process that involves configuring DNS forwarding and setting up AdGuard as the primary DNS resolver for your network.
Before beginning, ensure that you have a working installation of OPNsense and have administrative access to its web interface. You will also need to install AdGuard Home on a dedicated server or a virtual machine within your network. AdGuard can run on various platforms, including Linux, Windows, or Docker, providing flexibility depending on your environment. Once installed, AdGuard will operate as a DNS resolver, intercepting DNS queries and filtering out unwanted content based on predefined or custom blocklists.
The core idea behind adding AdGuard to OPNsense is to configure DNS forwarding so that all DNS requests from your network clients are routed through the AdGuard server. This involves setting up a static DNS resolver or DNS forwarder on OPNsense, pointing it to the AdGuard instance. Once the DNS traffic is directed correctly, you can access the AdGuard web interface to customize filtering rules, manage blocklists, and monitor DNS activity. This integration provides a transparent and efficient way to enhance your network’s security and browsing experience without affecting user devices directly.
By following this guide, you’ll be able to seamlessly add AdGuard to your OPNsense setup, creating a robust, network-wide ad-blocking solution that’s easy to manage and update. Proper configuration ensures minimal disruption and maximum filtering efficiency, making your network safer and more enjoyable for all users. Whether for home or small business use, this combination offers a powerful toolset for maintaining privacy and reducing unwanted content at the network level.
Overview of OPNsense and AdGuard
OPNsense is an open-source, FreeBSD-based firewall and routing platform designed to provide robust security, networking, and routing features. It is widely used by individual users and organizations to manage network traffic, implement security policies, and monitor network activity. Its modular architecture and extensive plugin support make it highly customizable and adaptable to various network configurations.
AdGuard, on the other hand, is a comprehensive ad-blocking solution that can be deployed across multiple platforms. It offers ad and tracker blocking, privacy protection, and malware filtering, enhancing online security and browsing experience. While traditionally used as a browser extension or standalone application, AdGuard also provides network-wide filtering capabilities through its DNS filtering service, making it suitable for integration into network firewalls like OPNsense.
Integrating AdGuard into OPNsense allows network administrators to enforce ad blocking and privacy policies at the network level, benefiting all connected devices without requiring individual configuration. This setup leverages AdGuard’s DNS filtering service, which intercepts DNS queries and blocks unwanted content before it reaches the client devices.
To achieve this integration, OPNsense can be configured to route DNS traffic through AdGuard’s DNS server or local instances of AdGuard Home, a DNS-based ad and tracker blocker. This configuration not only enhances network security but also reduces bandwidth consumption and improves user privacy. Proper setup involves configuring DNS forwarding on OPNsense to point to AdGuard’s servers and ensuring that DNS traffic is correctly routed through this service for all devices on the network.
What’s actually slowing this PC down?
Pick the symptom - the matching free tool is one click away.
Understanding the basic roles of OPNsense as a firewall and router, and AdGuard as a network-wide ad blocker, lays the foundation for effective integration. This process enhances your network’s security, privacy, and overall browsing performance by blocking unwanted content at a foundational level.
Benefits of Integrating AdGuard with OPNsense
Combining AdGuard with OPNsense offers a robust solution for enhancing your network’s security and privacy. This integration allows you to effectively block unwanted content, malware, and trackers at the network level, providing a seamless and comprehensive filtering experience. Here are the key advantages:
- Centralized Filtering Management: By integrating AdGuard with OPNsense, you manage ad blocking and content filtering across all devices connected to your network from a single interface. This eliminates the need for individual device configurations, ensuring uniform protection.
- Enhanced Privacy and Security: AdGuard blocks malicious domains, trackers, and unwanted advertisements, reducing exposure to malware, phishing sites, and intrusive ads. When combined with OPNsense’s firewall capabilities, it creates a multi-layered defense system.
- Network Performance Optimization: Blocking ads and trackers at the network level decreases unnecessary bandwidth consumption and speeds up browsing. This results in a smoother user experience for all connected devices.
- Customizable Filtering Rules: Both AdGuard and OPNsense support customizable rules, allowing you to tailor content filtering according to your specific needs. You can whitelist trusted domains or block specific categories of content.
- Reduced Device Resource Usage: Since filtering occurs at the network gateway, individual devices, especially low-powered ones like IoT gadgets or smartphones, conserve their resources by not running resource-intensive ad-blocking software.
- Comprehensive Logging and Monitoring: Integrating these tools offers better insight into network traffic and blocked threats. You can monitor, analyze, and respond proactively to potential security issues.
Overall, integrating AdGuard with OPNsense enhances your network’s security, privacy, and performance, creating a safer and more efficient digital environment for your users and devices.
Prerequisites and Requirements for Adding AdGuard to OPNsense
Before integrating AdGuard with OPNsense, ensure you have the necessary prerequisites in place. This will facilitate a smooth setup process and prevent potential issues.
Hardware Requirements
- OPNsense Device: A device running OPNsense, preferably with sufficient CPU and RAM to handle DNS filtering and additional load. Minimum specifications depend on network size.
- Network Interface: At least one network interface configured for your LAN and WAN. Additional interfaces may be required for complex setups.
Software Requirements
- OPNsense Version: Ensure you are running OPNsense version 21.1 or newer. Update if necessary to access the latest features and security patches.
- AdGuard Home Docker Container or Binary: Choose the appropriate installation method. For Docker, your system should support container deployment; for binary, download from the official source.
Network Configuration
- Static IP Address: Allocate a static IP for the AdGuard service to ensure consistent DNS resolution.
- Firewall Rules: Configure rules to allow DNS traffic between clients and AdGuard, typically on port 53 (UDP and TCP).
Additional Requirements
- Root or Sudo Access: Administrative privileges on the OPNsense device are necessary for configuration and package installation.
- Knowledge Base: Familiarity with OPNsense interfaces, DNS concepts, and basic networking will streamline the setup process.
