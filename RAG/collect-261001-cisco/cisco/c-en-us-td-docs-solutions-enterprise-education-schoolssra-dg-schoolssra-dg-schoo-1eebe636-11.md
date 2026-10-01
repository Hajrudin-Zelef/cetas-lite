---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-11
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [608, 650]
sha256: 5c88fa05b27362bed562f933e7bf1eda871b2a6d4339eeb50eb0b91734172817
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

•Hash-based—Requests are distributed using a hush function. The function uses the proxy ID and URL as inputs so that requests for the same URL are always directed to the same upstream proxy.
•Least recently used—Transactions are directed to the proxy that least recently received a transaction if all proxies are currently active.
•Round robin—The Web Proxy cycles transactions equally among all proxies in the group in the listed order.
Figure 4-27 illustrates the upstream proxy group configuration. Two upstream proxies are used, and transactions are forwarded to the proxy servicing the fewest number of connections.
Figure 4-27 WSA Upstream Proxy Group
Next, a routing rule needs to be defined to indicate when and how to direct transactions to the upstream proxy group. Use the Global Routing Policy if all traffic is to be handled by the upstream proxies. If no proxies are present, then leave the routing destination of the Global Routing Policy configured as Direct Connection. Figure 4-28 presents an example where all traffic is directed to the proxies in the Upstream-Lab_proxy group.
Figure 4-28 WSA Routing Policies
Web Access Policies
The access policies define how the Web Proxy handles HTTP requests and decrypted HTTPS connections for network users. By configuring access policies the school can control what Internet applications (instant messaging clients, peer-to-peer file-sharing, web browsers, Internet phone services, etc.) and URL categories students, staff and faculty may access. In addition, access policies can be used to block file downloads based on file characteristics, such as file size and file type.
The WSA comes with a default Global Policy that applies to all users. However, multiple policies can be defined when different policies need to be applied to different group of users. Figure 4-29 shows the global policy.
Figure 4-29 Global Access Policy
URL categories corresponding to content inappropriate for minors should be blocked in compliance with the school's Internet access policies. Figure 4-30 provides an example on how the "Adult/Sexually Explicit" category is blocked.
Figure 4-30 URL Categories
Layer-4 Traffic Monitoring
(L4TM)
L4TM can be implemented in the school environment to identify rogue traffic across all network ports and detect malware attempts to bypass port 80. Additionally, L4TM is capable of identifying internal clients with malware and that attempt to phone-home across non-standard ports and protocols.
Implementing L4TM requires the following:
Step 1 Configuring L4TM interfaces
Step 2 Configuring WSA L4TM global settings
Step 3 Configuring traffic monitoring
Configuring L4TM Interfaces
The wiring type depends on how traffic is directed to the WSA appliance. Network taps and SPAN can be either configured in simplex or duplex mode. If using a hub, only duplex mode can be used. The wiring type configuration is typically done during the initial setup as described earlier in this chapter. Figure 4-31 show the wiring options.
Figure 4-31 L4TM Wiring Type
Configuring WSA L4TM Global Settings
The ports to be monitored can be specified in the L4TM Global Settings. Options are:
•All ports—Monitors all 65535 TCP ports for rogue activity.
•All ports except proxy ports—Monitors all TCP ports except HTTP and HTTPS proxy ports.
Note The Cisco ASA in the Internet perimeter is configured to allow only permitted ports, so any connection attempts on rogue ports should be blocked by the firewall.
Figure 4-32 shows the options.
Figure 4-32 L4TM Global Settings
Configuring Traffic Monitoring
While a hub or a network tap could be used, using SPAN port mirroring provides the greatest flexibility. SPAN allows the monitoring of port traffic based on VLANs or source interfaces and it can easily be reconfigured.
The following is a configuration example of SPAN to replicate traffic:
! Enables port mirroring on the switch ports connecting to the firewall inside interfaces
monitor session 10 source interface Gi4/4
monitor session 10 source interface Gi5/3
!
! Sets the interface connecting to the WSA as the destination
monitor session 10 destination interface Gi6/3
The SPAN configuration can be seen on the switch with the show monitor session command.
Activity monitored by the L4TM feature can be seen in the L4 Traffic Monitor page of the WSA web-based GUI. Figure 4-33 shows client activity with a website known to be the source of malware.
Figure 4-33 L4 Traffic Monitor
Back to Top
