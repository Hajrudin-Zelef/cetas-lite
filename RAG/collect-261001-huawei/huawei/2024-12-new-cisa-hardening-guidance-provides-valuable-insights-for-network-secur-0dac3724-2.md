---
id: collect-261001-huawei/huawei/2024-12-new-cisa-hardening-guidance-provides-valuable-insights-for-network-secur-0dac3724-2
title: "2024-12-new-cisa-hardening-guidance-provides-valuable-insights-for-network-secur-0dac3724"
domain: huawei
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["benchmark", "benchmarks", "exploit", "parameters", "research"]
source: docs/RAG/collect-261001-huawei/2024-12-new-cisa-hardening-guidance-provides-valuable-insights-for-network-secur-0dac3724.md
source_anchor: ""
source_lines: [24, 74]
sha256: 1c42b60134db3fe6804e9e59f8db8bf43afbdf4b51d80aadb381ed17c70a1d65
---

# 2024-12-new-cisa-hardening-guidance-provides-valuable-insights-for-network-secur-0dac3724

The secure-by-design concept helps introduce the security conversation earlier in the development lifecycle. This approach helps ensure that security considerations are addressed at the beginning of the product lifecycle. Customers should make sure that products they plan to buy adhere to this principle. CISA has more information on its “Secure by Design” site. Tenable has committed to a secure-by-design approach, as can be seen in a recent initiative reported on here and here.
How Tenable can help
This overview is meant to help give network and security engineers a summary of the best practices, as well as provide insight on how CIS Benchmarks cover many of the guidance’s topics. Still, engineers should read the guidance to ensure they fully understand the material and how it relates to their own networks. It’s equally important to map out the network and understand what devices exist and where they are placed. However, this is only a first step in securing the network.
Tenable has several products, such as Tenable Vulnerability Management, Tenable Security Center, and Nessus that support auditing a wide array of devices and operating systems using CIS Benchmarks. These products could help with maintaining control over risk factors that threat actors often attempt to exploit. Tenable audits are written to test for the criteria of each automated recommendation in CIS Benchmarks. After an evaluation is run against the target, a result is provided as well as remediation text from the CIS Benchmark so that engineers can remediate and harden the device or operating system.
Tenable provides audit files for the following CIS Benchmarks to help organizations assess device configurations:
- CIS Check Point Firewall Benchmark v1.1.0 – Level 1, Level 2
- CIS Cisco ASA 9.x Firewall Benchmark v1.1.0 – Level 1, Level 2
- CIS Cisco Firewall v8.x Benchmark v4.2.0 – Level 1
- CIS Cisco IOS XE 16.x Benchmark v2.1.0 – Level 1, Level 2
- CIS Cisco IOS XE 17.x Benchmark v2.1.1 – Level 1, Level 2
- CIS Cisco IOS XR 7.x v1.0.0 – Level 1, Level 2
- CIS Cisco NX-OS Benchmark v1.1.0 – Level 1, Level 2
- CIS Fortigate 7.0.x Benchmark v1.3.0 – Level 1, Level 2
- CIS Juniper OS Benchmark v2.1.0 – Level 1, Level 2
- CIS Palo Alto Firewall 10 Benchmark v1.2.0 – Level 1, Level 2
- CIS Palo Alto Firewall 11 Benchmark v1.1.0 – Level 1, Level 2
These CIS Benchmarks align with the intent of the CISA hardening guidance. The example below highlights the CIS Cisco IOS XE 17.x v2.1.1 CIS Benchmark, and how it relates to the CISA hardening guidance:
Section 1.1 – Authentication, Authorization and Accounting (AAA) configuration
- Strengthening visibility as AAA logging supports user account login monitoring, and tracking changes
- Hardening systems and devices by providing identity management and policy enforcement
Section 1.2 – Access Rules for device administration
- Hardening systems and devices by restricting device management, and ensuring sessions are limited
Section 1.3 – Banner Rules to communicate legal rights to users
- Strengthening visibility by informing users they are subject to monitoring, and the event logs can support prosecution
Section 1.4 – Password Rules to enforce secure credentials and password lifecycle
- Hardening systems and devices by ensuring strong passwords are utilized, and passwords are securely stored
Section 1.5 – SNMP Rules provides guidance for secure configuration parameters
- Hardening systems and devices by ensuring SNMP is disabled, or is configured with secure parameters
Section 2.1 – Global Service Rules to reduce attack surface and disable unnecessary services
- Hardening systems and devices by disabling unnecessary, unused, exploitable, or plaintext services and protocols
Section 2.2 – Logging Rules configures log collection and forwarding
- Strengthening visibility by collecting event logs, and forwarding to a central log collection source
- Hardening systems and devices by forwarding logs to a central log collection source
Section 2.3 – NTP Rules ensures system time is provided by a single, consistent source
- Strengthening visibility by ensuring a consistent time source for event logs
- Hardening systems and devices by requiring that NTP is authenticated
Section 2.4 – Lookback Rules for configuring device initiated connections to supporting services such as AAA, SYSLOG, or NTP
- Hardening systems and devices by ensuring that traffic is initiated from a specific source, which can be used to set ACLs/filtering
Section 3.1 – Routing Rules to disable unneeded services
- Hardening systems and devices by disabling unneeded services such as source routing
Section 3.2 – Border Router Filtering defines filtering between internal and external networks
- Hardening systems and devices by implementing a strategy to control inbound and egress traffic
Section 3.3 – Neighbor Authentication configures routing protocol authentication
- Hardening systems and devices by requiring routing protocols are authenticated
Learn more
- Tenable Nessus
- Tenable Security Center
- Tenable Vulnerability Management
- Tenable Research audits
- Who is CIS?
- “Enhanced Visibility and Hardening Guidance for Communications Infrastructure” joint publication from various U.S. and international government agencies
