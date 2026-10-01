---
id: collect-261001-huawei/huawei/resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd-1
title: "resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd"
domain: huawei
role: reference
task: reference
actors: ["CISA", "China"]
dates: []
keywords: ["cyber", "cybersecurity"]
source: docs/RAG/collect-261001-huawei/resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd.md
source_anchor: ""
source_lines: [1, 32]
sha256: a4c202c2c171a35528e329c6f7d6a536669c3ff555710355d30a3e73de71efb3
---

# resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd

Enhanced Visibility and Hardening Guidance for Communications Infrastructure
Introduction
The Cybersecurity and Infrastructure Security Agency (CISA), National Security Agency (NSA), Federal Bureau of Investigation (FBI), Australian Signals Directorate’s (ASD’s) Australian Cyber Security Centre (ACSC), Canadian Centre for Cyber Security (CCCS), and New Zealand’s National Cyber Security Centre (NCSC-NZ) warn that People’s Republic of China (PRC)-affiliated threat actors compromised networks of major global telecommunications providers to conduct a broad and significant cyber espionage campaign. The authoring agencies are releasing this guide to highlight this threat and provide network engineers and defenders of communications infrastructure with best practices to strengthen their visibility and harden their network devices against successful exploitation carried out by PRC-affiliated and other malicious cyber actors. Although tailored to network defenders and engineers of communications infrastructure, this guide may also apply to organizations with on-premises enterprise equipment. The authoring agencies encourage telecommunications and other critical infrastructure organizations to apply the best practices in this guide.
As of this release date, identified exploitations or compromises associated with these threat actors’ activity align with existing weaknesses associated with victim infrastructure; no novel activity has been observed. Patching vulnerable devices and services, as well as generally securing environments, will reduce opportunities for intrusion and mitigate the actors’ activity.
Strengthening Visibility
In the context of this guide, visibility refers to organizations’ abilities to monitor, detect, and understand activity within their networks. High visibility means having detailed insight into network traffic, user activity, and data flow, allowing network defenders to quickly identify threats, anomalous behavior, and vulnerabilities. Visibility is critical for network engineers and defenders, particularly when identifying and responding to incidents.
Monitoring
Network Engineers
- Closely scrutinize and investigate any configuration modifications or alterations to network devices such as switches, routers, and firewalls outside of the change management process. Implement comprehensive alerting mechanisms to detect unauthorized changes to the network, including unusual route updates, enabled weak protocols, and configuration changes (i.e., changes to users and Access Control Lists [ACLs]).
  - Store configurations centrally and push to devices. Do not allow devices to be the trusted source of truth for their configuration. Monitor configuration and, if feasible, test and override on a frequent basis.
- Implement a strong network flow monitoring solution. This solution should allow for network flow data exporters and the associated collectors to be strategically centered around key ingress and egress locations that provide visibility into inter-customer traffic.
- If feasible, limit exposure of management traffic to the Internet. Only allow management via a limited and enforced network path, ideally only directly from dedicated administrative workstations.
- Monitor user and service account logins for anomalies that could indicate potential malicious activity. Validate all accounts and disable inactive accounts to reduce the attack surface. Monitor logins occurring internally and externally from the management environment.
- Implement secure, centralized logging with the ability to analyze and correlate large amounts of data from different sources. Encrypt any logging traffic destined for a remote destination via IPsec, TLS, or any other available encrypted transport options. Additionally, store copies of logs off-site to ensure they cannot be modified or deleted. Enable logging and auditing on devices and ensure logs can be offloaded from the device.
  - If possible, implement a Security Information and Event Management (SIEM) tool to analyze and correlate logs and alerts from the routers for rapid identification of security incidents.
  - Ensure logging takes place at all levels of the environment, network operating system, application, and software levels, as it pertains to network devices.
  - Establish a baseline of normal network behavior and define rules on security appliances to alert on abnormal behavior.
- Ensure the inventory of devices and firmware in the environment are up to date to enable effective visibility and monitoring.
Network Defenders
- Implement a monitoring and network management capability that, at a minimum, enforces configuration management, automates routine administrative functions, and alerts on changes detected within the environment, such as connections and user and account activity.
  - Establish understanding of the architecture of infrastructure and production enclaves, as well as where the two environments meet or are segregated. Map and understand boundary and ingress/egress points of the network management enclave.
  - Understand which assets should be forward facing and remove those that should not be forward facing. Closely monitor all devices that accept external connections from outside the corporate network and investigate any configurations that do not comply with known good configurations, such as open ports, services, or unexpected Generic Routing Encapsulation (GRE) or IPsec tunnel usage. Threat actors have been observed taking advantage of external-facing vulnerable services and features; therefore, proper visibility of network and security operations is vital.
  - If appropriate, implement a packet capture capability as part of the broader visibility effort for the enterprise. Determine capture location(s) and retention policies based on organizational demands.
Hardening Systems and Devices
Hardening device and network architecture is a defense-in-depth strategy. Reducing vulnerabilities, improving secure configuration habits, and following best practices limit potential entry points for PRC-affiliated and other cyber threats.
Protocols and Management Processes
Network Engineers
- Use an out-of-band management network that is physically separate from the operational data flow network. Ensure that management of network infrastructure devices can only come from the out-of-band management network. In addition, confirm that the out-of-band management network does not allow lateral management connections between devices to prevent lateral movement in the case that one device becomes compromised. Ensure device management is physically isolated from the customer and production networks. When properly implemented, out-of-band management can mitigate many threat actor tactics, techniques, and procedures (TTPs).
- Implement a strict, default-deny ACL strategy to control inbound and egressing traffic. Ensure all denied traffic is logged. For maximum depth, implement on separate devices from those implementing other security controls.
- Employ strong network segmentation via the use of router ACLs, stateful packet inspection, firewall capabilities, and demilitarized zone (DMZ) constructs. Separation via virtual local area networks (VLANs) and, if possible, private VLANs (PVLAN) will provide additional granular logical separation. This should be done as part of a broader defense-in-depth approach that protects and isolates different device groups.
  - Place externally facing services, such as Domain Name System (DNS), web servers, and mail servers, in a DMZ to provide segmentation from the internal LAN and backend resources.
  - Additionally, as a general strategy, put devices with similar purposes in the same VLAN. For example, place all user workstations from a certain team in one VLAN, while putting another team with different functions in a separate VLAN.
