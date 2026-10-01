---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-20
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [664, 711]
sha256: ec5b75326c927bf80bbf9a6e924c12b0fa635297309dd6a873d05a7325511ba7
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

access-list IPS permit ip any anyclass-map my-ips-classmatch access-list IPSpolicy-map my-ips-policyclass my-ips-classips inline fail-closeservice-policy my-ips-policy global
IPS Global Correlation Deployment
There are a number of aspects that need to be understood prior to the deployment of the IPS Global Correlation feature. To start with, before configuration be sure that you are using Cisco IPS Version 7.0 with the latest patch and signature updates and that Cisco IPS is configured for network connectivity in either IDS or IPS mode.
The configuration of Global Correlation can be performed using the command line interface (CLI), Cisco IDS Device Manager (IDM), Cisco IME, or the Cisco Security Manager. The figures are screen shots from Cisco IME that illustrate the basic steps in the configuration of IPS Global Correlation.
The first step in configuring the IPS sensor (module or appliance) to use Global Correlation is to add either a DNS address and/or the proxy server setup. This step, illustrated in Figure 27, enables connection to Cisco SensorBase. After you configure the DNS and proxy settings, these settings will go into effect as soon as the sensor has downloaded the latest Global Correlation updates.
Figure 27 DNS and HTTP Proxy Within the Network Setting Configuration Screen
By default, a sensor runs Global Correlation Inspection in Standard mode and enables the reputation filters, both illustrated in Figure 28. A good practice is to configure Global Correlation Inspection initially in permissive mode while monitoring the effects and later change the configuration into Standard or Aggressive mode.
Figure 28 Global Correlation Inspection Settings
By default network participation is disabled, which means the sensor does not share any event data back to the Cisco SensorBase network. The event data provided by all devices participating in the Cisco SensorBase network is a key element that provides real-time and worldwide visibility into the threat activity, accelerating the identification and mitigation of threats propagating throughout the Internet. For this reason Cisco recommends to configure the IPS sensors with partial or full network participation.
Figure 29 Network Participation Settings (Off by Default)
Event Monitoring with Global Correlation
Event monitoring with Global Correlation is similar to event monitoring with signature-only IPS. The primary difference is the potential addition of reputation scores representing the Global Correlation data. Figure 30 shows Cisco IPS events with reputation scores in Cisco IME.
Figure 30 Event Monitoring with Global Correlation in Cisco IME
Figure 30 shows several TCP SYN Port Sweep and ICMP Network Sweep attacks that were seen by the sensor. The first three events had no reputation and the event's risk ratings were 52 and 60, which did not meet the threshold for the packets to be dropped. The next three events were identical except that the attacker had a negative reputation of -1.8, which elevated the risk ratings to 70 and 75, but still did not meet the thresholds to be dropped in Standard Mode. The last events were also identical, except this time the attacker has a negative reputation if -4.3, which elevated the risk ratings to 86 and 89. This time the risk rating was high enough for the packets to be dropped.
Figure 31 illustrates a detailed view of the TCP SYN Port Sweep event coming from the attacker with a negative reputation of -4.3.
Figure 31 Detailed View of a TCP SYN Port Sweep from an Attacker with a Negative Reputation Score
Web Security Deployment
The small enterprise network design implements a Cisco IronPort WSA at the core/distribution layer of the main site, as illustrated in Figure 32. The WSA is located at the inside of the Cisco ASA acting as the Internet firewall. That ensures that clients and WSA are reachable over the same inside interface of the firewall and that the WSA can communicate with them without going through the firewall. At the same time, deploying the WSA at the core/distribution layer gives complete visibility to the WSA on the traffic before getting out to the Internet through the firewall.
Figure 32 Cisco IronPort WSA deployment
The guidelines for WSA configuration and deployment are presented below.
Initial System Setup Wizard
The WSA provides a browser-based system setup wizard that must be executed the first time the appliance is installed. The System Setup Wizard guides the user through initial system configuration, such as network and security settings. It should be noted that running the initial System Setup Wizard completely reconfigures the WSA appliance and resets the administrator password. Only use the System Setup Wizard the first time you install the appliance or if you want to completely overwrite the existing configuration.
The following are some of the default settings when running the System Setup Wizard:
•Web Proxy is deployed in transparent mode.
•The L4 Traffic Monitor is active and set to monitor traffic on all ports.
All these settings can be changed any time after the initial configuration by running the WSA Web-based configuration GUI.
Interface and Network Configuration
The steps need to be completed as part of the initial setup of the WSA appliance:
1. Configuring network interfaces.
5. Working with upstream proxy (if present).
Configuring Network Interfaces
Independently from the model, all Cisco IronPort WSA appliances are equipped with six Ethernet interfaces as shown in Figure 33.
Figure 33 WSA Interfaces
The WSA interfaces are grouped for the following functions:
•Management—Interfaces M1 and M2 are out-of-band (OOB) management interfaces. However, only M1 is enabled. In the Enterprise Design Profile for Small Enterprise Networks, interface M1 connects to the out-of-band management network. Interface M1 can optionally be used to handle data traffic in case the enterprise does not have an out-band management network.
•Web Proxy—Interfaces P1 and P2 are Web Proxy interfaces used for data traffic. Only the P1 interface is used in the Enterprise Design Profile for Small Enterprise Networks. P1 connects to the inside subnet of the firewall.
•L4 Traffic Monitor (L4TM)—T1 and T2 are the L4TM interfaces. L4TM is not used in the Enterprise Design Profile for Small Enterprise Networks; Cisco IPS Global Correlation and the Cisco ASA Botnet Filter features are used instead.
Figure 34 illustrates the network topology around the WSA used in the Cisco validation lab.
Figure 34 WSA Network Topology
Figure 35 illustrates the IP address and hostname configurations for the interfaces used. In this case, an out-of-band management network is used; therefore the M1 port is configured with an IP address in the management subnet. In addition, the WSA is configured to maintain a separate routing instance for the M1 management interface. This allows the definition of a default route for management traffic separate from the default route used for data traffic.
Figure 35 WSA Interface Configuration
Adding Routes
A default route is defined for management traffic pointing to the OOB management default gateway (172.26.191.1). A separate default route is defined for the data traffic pointing to the inside IP address of the firewall (10.125.33.10). As all internal networks are reachable throughout the core/distribution switch, a route to 10.0.0.0/8 is defined pointing to the switch IP address (10.125.33.9) to allow the WSA to communicate with the clients directly without having to go to the firewall first. These settings are illustrated in Figure 36.
Figure 36 WSA Route Configuration
Configuring DNS
The initial setup requires the configuration of a host name for the WSA appliance and listing the DNS servers. Figure 37 shows the DNS configuration.
Figure 37 WSA DNS Configuration
Setting Time
