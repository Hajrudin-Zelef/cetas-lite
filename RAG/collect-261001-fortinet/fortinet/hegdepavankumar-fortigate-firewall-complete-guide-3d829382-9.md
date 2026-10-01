---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-9
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "distribution", "license"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [661, 726]
sha256: ee198f78a952a08bfa61bfc7174303a1e7c021ee35801ef4e4db5db42daeace2
---

# Example SSH command to connect to the FortiGate firewall

To check the User activity on Firewall, Go to Monitor → Firewall User Monitor
Module 4 of the FortiGate firewall course focuses on firewall authentication. Here's a summary of the topics covered:
- Creating User and Policies: This topic covers the process of creating users and policies on the FortiGate firewall. It includes creating local users with authentication privileges and configuring firewall policies to control traffic flow based on defined criteria.
- Create Authentication Policies (Captive Portal): Captive Portal authentication allows administrators to authenticate users before granting access to network resources. This topic explains how to configure authentication policies using Captive Portal to enforce user authentication requirements.
- Monitor Firewall Users: Monitoring firewall users is crucial for maintaining network security and performance. This topic covers how to monitor and track user activities on the FortiGate firewall, including viewing logged-in users, session details, and authentication logs.
By covering these topics, Module 4 provides administrators with the knowledge and skills to effectively manage firewall authentication, create user and policy configurations, and monitor user activities to ensure network security and compliance.
Table of contents:
- Application Control
- Web Filtering
- File Filter
- DNS Filter
- Antivirus
- Intrusion Prevention
- Video Filter
- SSL/SSH Inspection Profile
Security profiles on FortiGate firewall provide advanced threat protection and content filtering capabilities to safeguard networks from various cyber threats. Here's a detailed overview of each security profile:
- Purpose: Application Control allows administrators to monitor and control applications' usage within the network, including identifying and blocking unauthorized or risky applications.
- Features: Application identification, application-based policies, application usage monitoring, and application-based traffic shaping.
To configure the Application control Go to Security Profiles —> Application Control
Select the Application you want block and apply action as Block.
Apply the Policy which we have created on Firewall & Objects as per requirements.
- Purpose: Web Filtering enables administrators to control access to websites based on categories, URLs, or specific content types, providing protection against malicious or inappropriate web content.
- Features: URL filtering, content filtering, antivirus scanning of web traffic, safe search enforcement, and SSL inspection for HTTPS websites.
- Purpose: File Filter scans file transfers for malware, malicious content, and unauthorized file types, helping to prevent the spread of malware and enforce data loss prevention policies.
- Features: File type filtering, antivirus scanning of file transfers, quarantine and blocking of infected files, and granular control over allowed file types.
- Purpose: DNS Filter blocks access to malicious or inappropriate domains by filtering DNS requests, providing an additional layer of security against phishing attacks, malware distribution, and access to undesirable content.
- Features: Domain filtering, domain reputation-based blocking, DNS sinkholing, and integration with threat intelligence feeds.
- Purpose: Antivirus scans network traffic for known malware, viruses, and other malicious content, protecting endpoints and networks from infection.
- Features: Real-time antivirus scanning, heuristic analysis, signature-based detection, quarantine and removal of infected files, and automatic updates of antivirus definitions.
For Antivirus Filtering we need valid License.
- Purpose: Intrusion Prevention System (IPS) identifies and blocks network-based attacks, including exploits, vulnerabilities, and malicious traffic patterns, to prevent unauthorized access and data breaches.
- Features: Signature-based detection, anomaly-based detection, protocol inspection, traffic anomaly detection, and blocking of known attack vectors.
Apply it, as previous we added to required Policy or Zone.
- Purpose: Video Filter controls access to streaming video content based on categories, URLs, or specific content types, helping to optimize bandwidth usage and enforce acceptable use policies.
- Features: Video streaming control, bandwidth management for video content, content filtering based on video categories, and visibility into video traffic usage.
- Purpose: SSL/SSH Inspection decrypts and inspects encrypted SSL/TLS or SSH traffic to detect and prevent threats hidden within encrypted communications, providing visibility into encrypted traffic and enforcing security policies.
- Features: SSL/TLS decryption, certificate inspection, SSL handshake validation, encryption protocol enforcement, and application control for decrypted traffic.
The Policy has configured LAN_WAN traffic.
Module 5 of the FortiGate firewall course focuses on security profiles, which are essential components for safeguarding networks against various cyber threats and enforcing security policies. Here's a summary of the security profiles covered:
- Application Control: Allows administrators to monitor and control the usage of applications within the network, helping to identify and block unauthorized or risky applications.
- Web Filtering: Enables administrators to control access to websites based on categories, URLs, or specific content types, providing protection against malicious or inappropriate web content.
- File Filter: Scans file transfers for malware, malicious content, and unauthorized file types, helping to prevent the spread of malware and enforce data loss prevention policies.
- DNS Filter: Blocks access to malicious or inappropriate domains by filtering DNS requests, providing an additional layer of security against phishing attacks, malware distribution, and access to undesirable content.
- Antivirus: Scans network traffic for known malware, viruses, and other malicious content, protecting endpoints and networks from infection.
- Intrusion Prevention: Identifies and blocks network-based attacks, including exploits, vulnerabilities, and malicious traffic patterns, to prevent unauthorized access and data breaches.
- Video Filter: Controls access to streaming video content based on categories, URLs, or specific content types, helping to optimize bandwidth usage and enforce acceptable use policies.
- SSL/SSH Inspection Profile: Decrypts and inspects encrypted SSL/TLS or SSH traffic to detect and prevent threats hidden within encrypted communications, providing visibility into encrypted traffic and enforcing security policies.
By understanding and configuring these security profiles, administrators can implement comprehensive threat protection measures and enforce security policies to safeguard their networks effectively against cyber threats.
Table of contents:
- Understanding Log severity levels
- Understanding Logs & Sublog types
- Understanding Log structures
- Configuring log settings
- Redirect logs to Syslog & SNMP
- Fortigate One-Arm Sniffer & SPAN
Logging and monitoring are essential components of network security, providing visibility into network activity, detecting threats, and troubleshooting issues. FortiGate firewall offers comprehensive logging and monitoring capabilities to help administrators effectively manage their network environments.
Log severity levels indicate the importance or severity of logged events. FortiGate firewall categorizes logs into different severity levels, including:
- Emergency: System is unusable.
- Alert: Immediate action is required.
- Critical: Critical conditions.
- Error: Error conditions.
- Warning: Warning conditions.
- Notice: Normal but significant condition.
- Informational: Informational messages.
- Debug: Debug-level messages.
FortiGate firewall generates various types of logs to capture different aspects of network activity and security events. Common log types include:
