---
id: collect-261001-general-networking/general-networking/manual-ips-html-7b5f5386-1
title: "manual-ips-html-7b5f5386"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-ips-html-7b5f5386.md
source_anchor: ""
source_lines: [1, 53]
sha256: ff31bde5727b091c151a91bbbadd82a4c4441efd21a8505d1ff9273b29c520d9
---

# manual-ips-html-7b5f5386

Intrusion Prevention System
The Intrusion Prevention System (IPS) system of OPNsense is based on Suricata and utilizes Netmap to enhance performance and minimize CPU utilization. This deep packet inspection system is very powerful and can be used to detect and mitigate security threats at wire speed.
IDS and IPS
It is important to define the terms used in this document. An Intrustion Detection System (IDS) watches network traffic for suspicious patterns and can alert operators when a pattern matches a database of known behaviors. An Intrusion Prevention System (IPS) goes a step further by inspecting each packet as it traverses a network interface to determine if the packet is suspicious in some way. If it matches a known pattern the system can drop the packet in an attempt to mitigate a threat.
The Suricata software can operate as both an IDS and IPS system.
Choosing an interface
You can configure the system on different interfaces. One of the most commonly asked questions is which interface to choose. Considering the continued use IPv4, usually combined with Network Address Translation, it is quite important to use the correct interface. If you are capturing traffic on a WAN interface you will see only traffic after address translation. This means all the traffic is originating from your firewall and not from the actual machine behind it that is likely triggering the alert.
Rules for an IDS/IPS system usually need to have a clear understanding about the internal network; this information is lost when capturing packets behind NAT.
Without trying to explain all the details of an IDS rule (the people at Suricata are way better in doing that), a small example of one of the ET-Open rules usually helps understanding the importance of your home network.
alert tls $HOME_NET any -> $EXTERNAL_NET any (msg:"ET TROJAN Observed Glupteba CnC Domain in TLS SNI"; flow:established,to_server; tls_sni; content:"myinfoart.xyz"; depth:13; isdataat:!1,relative; metadata: former_category MALWARE; reference:md5,4cc43c345aa4d6e8fd2d0b6747c3d996; classtype:trojan-activity; sid:2029751; rev:2; metadata:affected_product Windows_XP_Vista_7_8_10_Server_32_64_Bit, attack_target Client_Endpoint, deployment Perimeter, signature_severity Major, created_at 2020_03_30, updated_at 2020_03_30;)
The $HOME_NET can be configured, but usually it is a static net defined
in RFC 1918. Using advanced mode you can choose an external address, but
bear in mind you will not know which machine was really involved in the attack
and it should really be a static address or network.
$EXTERNAL_NET is defined as being not the home net, which explains why
you should not select all traffic as home since likely none of the rules will
match.
Since the firewall is dropping inbound packets by default it usually does not improve security to use the WAN interface when in IPS mode because it would drop the packet that would have also been dropped by the firewall.
Note
IDS mode is available on almost all (virtual) network types.
When your network card is not (fully) supported, you can set the tunable dev.netmap.admode to the value 2
in which case emulated mode will be enforced (Configurable in ). A list of natively
supported physical adapters is available in the FreeBSD man page.
General setup
The settings page contains the standard options to get your IDS/IPS system up and running.
| Enabled | Enable Suricata | 
| Capture mode | Choose between “PCAP live move (IDS)” for alerts only, “Netmap (IPS)” for alerts and discards via netmap driver, or “Divert (IPS)” to redirect packets via firewall rules. The action for a rule needs to be “drop” in order to discard the packet, this can be configured per rule or ruleset (using an input filter) | 
| Listeners | When “Divert (IPS)” mode is used, specify the number of listeners to initiate, usually this equals the number of CPUs in your system. | 
| Promiscuous mode | Listen to traffic in promiscuous mode. (all packets instead of only the ones addressed to this network interface) | 
| Enable syslog alerts | Send alerts to syslog, using fast log format | 
| Enable eve syslog output | Send alerts in EVE format to syslog, using log level info. This will not change the alert logging used by the product itself. Drop logs will only be send to the internal logger, due to restrictions in suricata. | 
| Pattern matcher | Controls the pattern matcher algorithm. Aho–Corasick is the default. On supported platforms, Hyperscan is the best option. On commodity hardware if Hyperscan is not available the suggested setting is “Aho–Corasick Ken Steele variant” as it performs better than “Aho–Corasick”. | 
| Interfaces | Interfaces to protect. When in IPS mode, this needs to be real interfaces supporting netmap. (when using VLANs, enable IPS on the parent) | 
| Rotate log | Log rotating frequency, also used for the internal event logging (see Alert tab) | 
| Save logs | Number of logs to keep | 
Note
To use the “Divert (IPS)” mode, you must use and create firewall rules that contain the “Divert-to” setting. Check the Rules manual for more information.
Tip
When using an external reporting tool, you can use syslog to ship your EVE log easily. Just enable “Enable EVE syslog output” and create a target in . (filter application “suricata” and level “info”)
Note
When using IPS mode make sure all hardware offloading features are disabled in the interface settings (). Prior to version 20.7, “VLAN Hardware Filtering” was not disabled which may cause issues for some network cards.
Advanced options
Some less frequently used options are hidden under the “advanced” toggle.
| Home networks | Define custom home networks, when different than an RFC1918 network. In some cases, people tend to enable IDPS on a wan interface behind NAT (Network Address Translation), in which case Suricata would only see translated addresses instead of internal ones. Using this option, you can define which addresses Suricata should consider local. | 
| default packet size | With this option, you can set the size of the packets on your network. It is possible that bigger packets have to be processed sometimes. The engine can still process these bigger packets, but processing it will lower the performance. | 
Download rulesets
When enabling IDS/IPS for the first time the system is active without any rules to detect or block malicious traffic. The download tab contains all rulesets available on the system (which can be expanded using plugins).
In this section you will find a list of rulesets provided by different parties and when (if installed) they where last downloaded on the system. In previous versions (prior to 21.1) you could select a “filter” here to alter the default behavior of installed rules from alert to block. As of 21.1 this functionality will be covered by Policies, a separate function within the IDS/IPS module, which offers more fine grained control over the rulesets.
Note
When migrating from a version before 21.1 the filters from the download rulesets page will automatically be migrated to policies.
Policies
The policy menu item contains a grid where you can define policies to apply to installed rules. Here you can add, update or remove policies as well as disabling them. Policies help control which rules you want to use in which manner and are the preferred method to change behaviour. Although you can still update separate rules in the rules tab, adding a lot of custom overwrites there is more sensitive to change and has the risk of slowing down the user-interface.
A policy entry contains 3 different sections. First some general information, such as the description and if the rule is enabled as well as a priority. Overlapping policies are taken care of in sequence, the first match with the lowest priority number is the one to use.
