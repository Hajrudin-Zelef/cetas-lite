---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-22
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license", "licenses"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [746, 790]
sha256: c26d99aae60e6be0a2e197b7044d5db1bfcce9d0a6a9c3c4464ac33b9faf6b09
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

   Global WCCP information:Router information:Router Identifier: 198.133.219.5Protocol Version: 2.0Service Identifier: 10Number of Cache Engines: 1Number of routers: 1Total Packets Redirected: 428617Redirect access-list: proxylistTotal Connections Denied Redirect: 0Total Packets Unassigned: 4Group access-list: wsa-farmTotal Messages Denied to Group: 0Total Authentication failures: 0Total Bypassed Packets Received: 0cr26-asa5520-do#
Enabling WSA HTTPS Scanning
To monitor and decrypt HTTPS traffic, you must enable HTTPS scanning on the WSA. The HTTPS Proxy configuration is illustrated in Figure 43.
Figure 43 WSA HTTPS Proxy
Working with Upstream Proxies
In case Internet traffic is handled by one or more upstream proxies, follow these guidelines:
•Define a routing policy to direct traffic to the upstream proxies
The Upstream Proxy Group lists the IP addresses or domain names of the proxies to be used for traffic sent to the Internet. When multiple proxies are available, the WSA can be configured for failover or load balancing.
The following are the options available:
•None (failover)—The first proxy in the list is used. If one proxy cannot be reached, the Web Proxy attempts to connect to the next one in the list.
•Fewest connections—Transactions are directed to the proxy servicing the fewest number of connections.
•Hash-based—Requests are distributed using a hash function. The function uses the proxy ID and URL as inputs so that requests for the same URL are always directed to the same upstream proxy.
•Least recently used—Transactions are directed to the proxy that least recently received a transaction if all proxies are currently active.
•Round robin—The Web Proxy cycles transactions equally among all proxies in the group in the listed order.
Figure 44 illustrates the upstream proxy group configuration. Two upstream proxies are used and transactions are forwarded to the proxy servicing the fewest number of connections.
Figure 44 WSA Upstream Proxy Group
Next, a routing rule needs to be defined to indicate when and how to direct transactions to the upstream proxy group. Use the Global Routing Policy if all traffic is to be handled by the upstream proxies. If no proxies are present, leave the routing destination of the Global Routing Policy configured as Direct Connection. Figure 45 presents an example where all traffic is directed to the proxies in the Upstream-Lab_proxy group.
Figure 45 WSA Routing Policies
Web Access Policies
Access policies define how the Web Proxy handles HTTP requests and decrypted HTTPS connections for network users. By configuring access policies, the enterprise can control user acces to Internet applications (instant messaging clients, peer-to-peer file-sharing, Web browsers, Internet phone services, etc.) and URL categories. In addition, access policies can be used to block file downloads based on file characteristics, such as file size and file type.
The WSA comes with a default Global Policy that applies to all users. However, multiple policies can be defined when different policies need to be applied to different groups of users. Figure 46 shows the global policy.
Figure 46 Global Access Policy
URL categories corresponding to non-business-related content should be blocked in compliance with the company's Internet access policies. Figure 47 provides an example of how the "Adult/Sexually Explicit" category is blocked.
Figure 47 URL Categories
CISF Protected Ports
Catalyst Integrated Security Features (CISF) is a set of native security features available on Cisco Catalyst Switches that protect the network against attacks such as man-in-the-middle, spoofing, and infrastructure denial-of-services (DoS) attacks. CISF includes:
The following is a sample CISF port configuration:
! configure port security parametersswitchport port-security maximum 2switchport port-securityswitchport port-security aging time 2switchport port-security violation restrictswitchport port-security aging type inactivity! configure arp inspection rate limitingip arp inspection limit rate 100! configure dhcp snooping rate limitingip dhcp snooping limit rate 100! configure storm control parametersstorm-control broadcast level 20.00 10.00storm-control multicast level 50.00 30.00
NAC Appliance Deployment
In the small enterprise network design, a NAC Appliance solution is deployed at all locations, the main site, and each of the remote offices. To that end, a centralized CAM is deployed at the main site, likely within the serverfarm. A CAS is deployed at the main site and each remote site, each of which directly connects to the core/distribution layer at each of the locations.
Deploying the NAC appliance solution requires the configuration of the CAM and CAS components. The initial CAM and CAS configuration is performed via console access (described in the Cisco NAC Appliance Hardware Installation guide). During the configuration stage, multiple steps must be followed in configuring the NAC Appliance with IP addresses, VLANs. passwords, etc. The installation guide contains worksheets that assist in gathering and preparing this information for both CAM and CAS.
After performing the initial setup through the console, the rest of the configuration of the CAM and CAS is performed using the CAM Web-based GUI. The first task on the CAM, before beginning configuration, is the installation the licenses for the solution. A license must be installed for the CAM and the CAS servers controlled by the CAM. The Cisco NAC Appliance Ordering Guide provides information on the ordering options.
Licenses can be entered via the CAM Web-based GUI, as shown in Figure 48.
Figure 48 NAC Appliance Licensing
Adding a CAS to the CAM
For a CAS server to be managed by the CAM, it must be first added to the list of managed servers on the CAM. To do this the CAM needs to know the IP address of the CAS and the Server Type (its role in the network) of the CAS. In addition, the CAS and the CAM must have the same shared secret. The shared secret is configured during the server installation. An example of adding a CAS to the CAM is shown in Figure 49.
Figure 49 Adding a new CAS to the CAM
Once the CAS has been added to the CAM, it appears in the list of servers on the CAM. From this point, it can be managed directly from the CAM for almost all tasks. An example of a list of servers is shown in Figure 50.
Figure 50 List of CAS Servers
Managing the CAS
Once the CAS is in the list of servers managed by the CAM, it can be configured for its role in the network. To manage the server, click the icon under the Manage heading in the server list, which connects you to the CAS server and shows you the summary menu shown in Figure 51.
Figure 51 CAS Management Menu
Under the CAS Network setting tab, shown in Figure 52, the basic network settings for the CAS can be viewed and altered, if needed. In this example, we are keeping the network configuration set up during the server installation. The primary dialog under the Network Tab is the IP dialog, shown in Figure 52; the other dialogs allow the DHCP options to the configured—our example uses the default of DHCP passthrough—and the DNS options where host name, domain name, and DNS server information are added, as shown in Figure 53.
Figure 52 CAS Network Settings
Figure 53 CAS DNS Settings
