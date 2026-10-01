---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-5
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["cost", "memory", "throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [166, 199]
sha256: a7bc976bebec2c96f3ed7ecf2a8bfc0caa28ddedc4d9aca6abcb60bd2ab4cc93
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Note The SensorBase network is an extensive network that monitors global E-mail and Web traffic for anomalies, viruses, malware, and other abnormal behavior. The network is composed of the Cisco IronPort appliances, Cisco ASA, and Cisco IPS appliances and modules installed in more than 100,000 organizations worldwide, providing a large and diverse sample of Internet traffic patterns.
The Cisco ASA uses the dynamic database as follows:
1. When the domain name in a DNS reply matches a name in the dynamic database, the Botnet Traffic Filter adds the name and IP address to the DNS reverse lookup cache.
2. When the infected host starts a connection to the IP address of the malware site, the Cisco ASA sends a syslog message reporting the suspicious activity and optionally drops the traffic if the Cisco ASA is configured to do so.
3. In some cases, the IP address itself is supplied in the dynamic database and the Botnet Traffic Filter logs or drops any traffic to that IP address without having to inspect DNS requests.
The database files are stored in running memory rather than Flash memory. The database can be deleted by disabling and purging the database through the configuration.
Note To use the database, be sure to configure a domain name server for the Cisco ASA so that it can access the URL of the update server. To use the domain names in the dynamic database, DNS packet inspection with Botnet Traffic Filter snooping needs to be enabled; the Cisco ASA looks inside the DNS packets for the domain name and associated IP address.
In addition to the dynamic database, a static database can be used by manually entering domain names or IP addresses (host or subnet) that you want to tag as bad names in a blacklist. Static blacklist entries are always designated with a Very High threat level. Domain names or IP addresses can also be entered in a whitelist.
When a domain name is added to the static database, the Cisco ASA waits one minute and then sends a DNS request for that domain name and adds the domain name/IP address pairing to the DNS host cache. This action is a background process and does not affect your ability to continue configuring the Cisco ASA. Cisco also recommends that DNS packet inspection be enabled with Botnet Traffic Filter snooping. When enabled, the Cisco ASA uses Botnet Traffic Filter snooping instead of the regular DNS lookup to resolve static blacklist domain names in the following circumstances:
•The Cisco ASA DNS server is unavailable.
•A connection is initiated during the one minute waiting period before the Cisco ASA sends the regular DNS request.
If DNS snooping is used, when an infected host sends a DNS request for a name on the static database, the Cisco ASA looks inside the DNS packets for the domain name and associated IP address and adds the name and IP address to the DNS reverse lookup cache.
If Botnet Traffic Filter snooping is not enabled and one of the above circumstances occurs, that traffic is not monitored by the Botnet Traffic Filter.
Note It is important to realize that a comprehensive security deployment should include Cisco Intrusion Prevention Systems (IPS) with its reputation-based Global Correlation service and IPS signatures in conjunction with the security services provided by the Cisco ASA security appliance, such as Botnet Traffic Filter.
For more information on the Cisco ASA Botnet Traffic Filter feature, see: http://www.cisco.com/en/US/prod/vpndevc/ps6032/ps6094/ps6120/botnet_index.html.
Intrusion Prevention Guidelines
IPS is responsible for identifying and blocking anomalous traffic and packets recognized as well-known attacks. An AIP SSM module on the Cisco ASA Internet firewall or a separate IPS appliance can be implemented in the Internet perimeter for enhanced threat detection and mitigation. IPS may also be configured to help block certain Internet applications such as AOL Messenger, BitTorrent, Skype, etc.
Integrating IPS on a Cisco ASA appliance using an AIP SSM provides a cost-effective solution for small enterprise networks. The AIP SSM is supported on Cisco ASA 5510 and higher platforms. The AIP SSM runs advanced IPS software providing proactive, full-featured intrusion prevention services to stop malicious traffic before it can affect the enterprise network.
The AIP SSM may be deployed in inline or promiscuous mode:
•Inline mode—The AIP SSM is placed directly in the traffic flow (see the left side of Figure 5.). Traffic identified for IPS inspection cannot continue through the ASA without first passing through and being inspected by the AIP SSM. This mode is the most secure because every packet that has been identified for inspection is analyzed before being allowed through. Also, the AIP SSM can implement a blocking policy on a packet-by-packet basis. This mode, however, can affect throughput if not designed or sized appropriately.
•Promiscuous mode—A duplicate stream of traffic is sent to the AIP SSM. This mode is less secure, but has little impact on traffic throughput. Unlike inline mode, in promiscuous mode the AIP SSM can block traffic only by instructing the ASA to shun the traffic or by resetting a connection on the ASA. Also, while the AIP SSM is analyzing the traffic, a small amount of traffic might pass through the ASA before the AIP SSM can shun it. The right side of Figure 5 shows the AIP SSM in promiscuous mode.
Figure 5 IPS Inline and Promiscuous Modes
The recommended IPS deployment mode depends on the goals and policies of the enterprise. IPS inline mode is more secure because of its ability to stop malicious traffic in real-time, however it may impact traffic throughput if not properly designed or sized. Conversely, IPS promiscuous mode has less impact on traffic throughput but is less secure because there may be a delay in reacting to the malicious traffic.
Although the AIP SSM runs as a separate application within the Cisco ASA, it is integrated into the traffic flow. The AIP SSM contains no external interfaces itself, except for the management interface on the SSM itself. When traffic is identified for IPS inspection on the ASA, traffic flows through the ASA and the AIP SSM in the following sequence:
1. Traffic enters the adaptive security appliance.
2. Firewall policies are applied.
3. Traffic is sent to the AIP SSM over the backplane.
4. The AIP SSM applies its security policy to the traffic and takes appropriate actions.
5. Valid traffic (for inline mode only) is sent back to the adaptive security appliance over the backplane; the AIP SSM might block some traffic according to its security policy and that traffic is not passed on.
6. VPN policies are applied (if configured).
7. Traffic exits the adaptive security appliance.
The AIP SSM card may be configured to fail open or close when the module becomes unavailable. When configured to fail open, the ASA allows all traffic through, uninspected, if the AIP SSM becomes unavailable. Conversely, when configured to fail close, the ASA blocks all traffic in case of an AIP SSM failure.
Cisco IPS Global Correlation
If desired, the AIP SSM module on the Cisco ASA (or IPS appliance) may participate in Cisco IPS Global Correlation for further threat visibility and control. Once enabled, the participating IPS sensor receives threat updates from the Cisco SensorBase Network at regular intervals. The Cisco SensorBase Network contains detailed information about known threats on the Internet, including serial attackers, botnet harvesters, malware outbreaks, and dark nets. The IPS uses this information to filter out the worst attackers before they have a chance to attack critical assets. It then incorporates the global threat data into its system to detect and prevent malicious activity even earlier.
