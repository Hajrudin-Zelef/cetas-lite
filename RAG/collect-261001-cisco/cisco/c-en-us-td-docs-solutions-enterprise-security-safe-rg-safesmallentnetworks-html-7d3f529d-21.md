---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-21
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [712, 745]
sha256: cce7293a1ee5b61866821b835e0a59a61bc498e82e9542c3c06a85ec7808cb89
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Time synchronization is critical for forensic analysis and troubleshooting, therefore enabling NTP is highly recommended. Figure 38 shows how the WSA is configured to synchronize its clock with an NTP server located on the OOB management network.
Figure 38 WSA NTP Configuration
Working with Upstream Proxies
If Internet access is provided by an upstream proxy, then the WSA must be configured to use the proxy for component updates and system upgrades. This is illustrated in Figure 39 and Figure 40.
Figure 39 WSA Upgrade Settings
Figure 40 WSA Component Updates
WCCP Transparent Web Proxy
The configuration of the WCCP Transparent Web Proxy involves:
1. Defining WSA WCCP Service Group.
2. Enabling WSA Transparent Redirection.
3. Enabling WCCP redirection on the Cisco ASA.
4. Enabling WSA HTTPS scanning.
5. Working with upstream proxy (if present).
Defining WSA WCCP Service Group
Web Proxy settings are configured as part of an initial setup using the System Setup Wizard and can be later modified with the WSA Web-based GUI. The Web Proxy settings include:
•HTTP Ports to Proxy—List the ports to be proxied; default is 80 and 3128.
•Caching—Defines whether or not the WSA should cache response and requests. Caching helps reduce latency and the load on the Internet links; default is enabled.
•IP Spoofing—Defines whether or not the Web Proxy should spoof IP addresses when forwarding requests to upstream proxies and servers. The Cisco ASA does not support source address spoofing.
Figure 41 illustrates the Web Proxy settings.
Figure 41 WSA Proxy Settings
Enabling WSA Transparent Redirection
Configuring WCCP Transparent Redirection requires the definition of a WCCP service profile in the WSA. If redirecting HTTP and HTTPS, define a dynamic service ID to be used with the Cisco ASA. Use MD5 authentication to protect the WCCP communication between the WSA and Cisco ASA. Figure 42 shows an example.
Figure 42 WSA Transparent Proxy
Enabling WCCP Redirection on Cisco ASA
The configuration of WCCP on the Cisco ASA appliance requires:
•A group-list indicating the IP addresses of the appliances member of the service group. In the example provided below, the group-list is named wsa-farm.
•A redirect-list indicating the ports and subnets of traffic to be redirected. In the example, the ACL named proxylist is configured to redirect any HTTP and HTTPS traffic coming from the 10.0.0.0/8 subnet. It is critical to ensure traffic from the WSA(s) bypasses redirection. To that end, add an entry to the redirect-list explicitly denying traffic sourced from the WSA(s).
•WCCP service indicating the service ID. Make sure you use the same ID as defined on the WSAs. Use a password for MD5 authentication.
•Enabling WCCP redirection on an interface. Apply the WCCP service on the inside interface of the Cisco ASA.
Cisco ASA WCCP configuration example:
! Group-list defining the IP addresses of all WSAsaccess-list wsa-farm extended permit ip host 10.125.33.8 any!! Redirect-list defining what ports and hosts/subnets should be redirectedaccess-list proxylist extended deny ip host 10.125.33.8 anyaccess-list proxylist extended permit tcp 10.0.0.0 255.0.0.0 any eq wwwaccess-list proxylist extended permit tcp 10.0.0.0 255.0.0.0 any eq https!! WCCP servicewccp 10 redirect-list proxylist group-list wsa-farm password cisco!! Applies WCCP on an interfacewccp interface inside 10 redirect in
The WCCP connection status and configuration can be monitored on the Cisco ASA with the show wccp command. An example is provided below:
cr26-asa5520-do# show wccp
 
