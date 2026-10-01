---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-10
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [512, 607]
sha256: 4b91271c3abcfb615ef76e8740e2703879e85fc9c2df1ec092a5f6c26e39151b
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

Step 1 Configuring network interfaces
Step 2 Adding routes
Step 3 Configuring DNS
Step 4 Setting time
Step 5 Working with upstream proxy (if present)
These settings are configured as part of an initial setup using the System Setup Wizard, but can be later modified by using the WSA Web-based GUI.
Configuring Network Interfaces
Independently from the model, all Cisco IronPort WSA appliances are equipped with six Ethernet interfaces as shown in Figure 4-16.
Figure 4-16 WSA Interfaces
The WSA interfaces are grouped for the following functions:
•Management—Interfaces M1 and M2 are out-of-band (OOB) management interfaces. However, only M1 is enabled. In the school architecture, interface M1 connects to the out-of-band management network. Interface M1 can optionally be used to handle data traffic in case the school does not have an out-band management network.
•Web Proxy—Interfaces P1 and P2 are Web Proxy interfaces used for data traffic. Only the P1 interface is used in the school architecture. P1 connects to the inside subnet of the firewall.
•L4 Traffic Monitor (L4TM)—T1 and T2 are the L4TM interfaces. The school design uses only the T1 interfaces. The T1 interface connects to a core/distribution switch port configured as the destination of the SPAN session used to capture traffic bound to the Internet.
Figure 4-17 illustrates the network topology around the WSA used in the Cisco validation lab.
Figure 4-17 WSA Network Topology
Figure 4-18 illustrate the IP address and hostname configurations for the interfaces used. In this case, an out-of-band management network is used; therefore the M1 port is configured with an IP address in the management subnet. In addition, the WSA is configured to maintain a separate routing instance for the M1 management interface. This allows the definition of a default route for management traffic separate from the default route used for data traffic.
Figure 4-18 WSA Interface Configuration
Adding Routes
A default route is defined for management traffic pointing to the OOB management default gateway (172.26.191.1). A separate default route is defined for the data traffic pointing to the inside IP address of the firewall (10.125.33.10). As all internal networks are reachable throughout the core/distribution switch, a route to 10.0.0.0/8 is defined pointing to the switch IP address (10.125.33.9) to allow the WSA to communicate with the clients directly without having to go to the firewall first. These settings are illustrated in Figure 4-19.
Figure 4-19 WSA Route Configuration
Configuring DNS
The initial setup requires the configuration of a host name for the WSA appliance, and listing the DNS servers. Figure 4-20 shows the DNS configuration.
Figure 4-20 WSA DNS Configuration
Time Settings
Time synchronization is critical for forensic analysis and troubleshooting, therefore enabling NTP is highly recommended. Figure 4-21 shows how the WSA is configured to synchronize its clock with an NTP server located on the OOB management network.
Figure 4-21 WSA NTP Configuration
Working with Upstream Proxies
If Internet access is provided by an upstream proxy, then the WSA must be configured to use the proxy for component updates and system upgrades. This is illustrated in Figure 4-22 and Figure 4-23.
Figure 4-22 WSA Upgrade Settings
Figure 4-23 WSA Component Updates
WCCP Transparent Web Proxy
The configuration of the WCCP Transparent Web Proxy includes the following:
Step 1 Defining WSA WCCP Service Group
Step 2 Enabling WSA Transparent Redirection
Step 3 Enabling WCCP redirection on the Cisco ASA
Step 4 Enabling WSA HTTPS scanning
Step 5 Working with upstream proxy (if present)
Defining WSA WCCP Service Group
Web Proxy settings are configured as part of an initial setup using the System Setup Wizard and can be later modified with the WSA Web-based GUI. The Web Proxy setting include the following:
•HTTP Ports to Proxy—List the ports to be proxied. Default is 80 and 3128.
•Caching—Defines whether or not the WSA should cache response and requests. Caching helps reduce latency and the load on the Internet links. Default is enabled.
•IP Spoofing — Defines whether or not the Web Proxy should spoof IP addresses when forwarding requests to upstream proxies and servers. The Cisco ASA does not support source address spoofing.
Figure 4-24 illustrates the Web Proxy settings.
Figure 4-24 WSA Proxy Settings
Enabling WSA Transparent Redirection
Configuring WCCP Transparent Redirection requires the definition of a WCCP service profile in the WSA. If redirecting HTTP and HTTPS, define a dynamic service ID to be used with the Cisco ASA. Use MD5 authentication to protect the WCCP communication between the WSA and Cisco ASA. Figure 4-25 shows an example.
Figure 4-25 WSA Transparent Proxy
Enabling WCCP Redirection on Cisco ASA
The configuration of WCCP on the Cisco ASA appliance requires:
•A group-list indicating the IP addresses of the appliances member of the service group. In the example provided below the group-list is called wsa-farm.
•A redirect-list indicating the ports and subnets of traffic to be redirected. In the example, the ACL named proxylist is configured to redirect any HTTP and HTTPS traffic coming from the 10.0.0.0/8 subnet. It is critical to ensure traffic from the WSA(s) bypasses redirection. To that end, add an entry to the redirect-list explicitly denying traffic sourced from the WSA(s).
•WCCP service indicating the service ID. Make sure you use the same ID as defined on the WSAs. Use a password for MD5 authentication.
•Enabling WCCP redirection on an interface. Apply the WCCP service on the inside interface of the Cisco ASA.
Cisco ASA WCCP configuration example:
! Group-list defining the IP addresses of all WSAs
access-list wsa-farm extended permit ip host 10.125.33.8 any 
!
! Redirect-list defining what ports and hosts/subnets should be redirected
access-list proxylist extended deny ip host 10.125.33.8 any 
access-list proxylist extended permit tcp 10.0.0.0 255.0.0.0 any eq www 
access-list proxylist extended permit tcp 10.0.0.0 255.0.0.0 any eq https 
!
! WCCP service
wccp 10 redirect-list proxylist group-list wsa-farm password cisco
!
! Applies WCCP on an interface
wccp interface inside 10 redirect in
The WCCP connection status and configuration can be monitored on the Cisco ASA with the show wccp command. An example is provided below:
cr26-asa5520-do# show wccp
Global WCCP information:
    Router information:
	Router Identifier:                   198.133.219.5
	Protocol Version:                    2.0
    Service Identifier: 10
	Number of Cache Engines:             1
	Number of routers:                   1
	Total Packets Redirected:            428617
	Redirect access-list:                proxylist
	Total Connections Denied Redirect:   0
	Total Packets Unassigned:            4
	Group access-list:                   wsa-farm
	Total Messages Denied to Group:      0
	Total Authentication failures:       0
	Total Bypassed Packets Received:     0
cr26-asa5520-do# 
Enabling WSA HTTPS Scanning
To monitor and decrypt HTTPS traffic, you must enable HTTPS scanning on the WSA. The HTTPS Proxy configuration is illustrated in Figure 4-26.
Figure 4-26 WSA HTTPS Proxy
Working with Upstream Proxies
In case Internet traffic is handled by one or more upstream proxies, follow these guidelines:
•Add an Upstream Proxy Group
•Define a routing policy to direct traffic to the upstream proxies
The Upstream Proxy Group lists the IP addresses or domain names of the proxies to be used for traffic sent to the Internet. When multiple proxies are available, the WSA can be configured for failover or load balancing.
The following are the options available:
•None (failover)—The first proxy in the list is used. If one proxy cannot be reached, the Web Proxy attempts to connect to the next one in the list.
•Fewest connections—Transactions are directed to the proxy servicing the fewest number of connections.
