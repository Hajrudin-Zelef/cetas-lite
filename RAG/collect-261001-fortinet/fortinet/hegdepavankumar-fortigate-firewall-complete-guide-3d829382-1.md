---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-1
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "research", "sandbox", "zero-day"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [1, 50]
sha256: 396f8a6b1bc99fd82b7bca8ff9550b0de0acce5301383cf3d9380976c86078b2
---

# Example SSH command to connect to the FortiGate firewall

"Welcome to the FortiGate Firewall Complete Guide! This comprehensive GitHub repository offers detailed lab guides and invaluable information for mastering FortiGate firewalls. Created with countless days and nights of dedication, this guide is entirely free for your benefit. If you find it helpful, please consider showing your support with a coffee ☕️ or your best wishes. Your encouragement fuels the continuation of this resource. Thank you for being part of this journey! 🚀🔥 #FortiGate #Firewall #Networking"
- All the labs are simulated on EVE-NG
- Fortigate Firewall Image Used Version: v7.0.9
- Used Cisco Routers and Switches
- Windows Active Directory Used [Windows Server 2019]
- Eve-NG Hosted on VMware Workstation Pro 17.
- You Will Get Fortigate Image: Click here to Download
- All the Images for EVE-NG and GNS3 available here also: Cisco-Images-for-GNS3-and-EVE-NG
- Free VMware Workstation PRO 17 Practice License: VMware-Workstation-Pro-17-Licence-Keys
- Free VMware ESXI License Keys Lab License: VMware-ESXi-License-Keys
- Let's start : click here
| No | Name | 
|---|---|
| Module 1 | Introduction to Fortigate Firewall | 
| Module 2 | Interface Configurations and Firewall Policies | 
| Module 3 | High availability | 
| Module 4 | Firewall Authentication | 
| Module 5 | Security Profiles | 
| Module 6 | Logging and Monitoring | 
| Module 7 | Basic IPSEC VPN | 
| Module 8 | SSL VPN Upcoming.. | 
| Module 9 | **** Upcoming.. | 
Table of Contents:
- Understanding Features of Fortigate
- Fortigaurd Queries & Packages
- UTM firewalls futures
- Platform Design and Architecture
- About CLI
- Getting Mgmt GUI Access
- About Administration Profiles
FortiGate is a family of network security appliances developed by Fortinet, designed to provide a wide range of security features to protect networks from various threats. Understanding the features of FortiGate involves grasping its capabilities in different areas of network security. Here's a detailed breakdown:
- Firewall: FortiGate operates as a firewall, providing traditional packet filtering capabilities to monitor and control the traffic passing through the network based on predefined rules. It can inspect packets at the application layer for more granular control.
- Intrusion Prevention System (IPS): FortiGate includes an IPS module that identifies and blocks malicious activities within the network. It analyzes traffic patterns and signatures to detect and prevent known attacks, such as SQL injection, buffer overflow, and denial-of-service (DoS) attacks.
- Virtual Private Network (VPN): FortiGate supports VPN technologies, allowing secure communication between remote sites or individual users and the corporate network over untrusted networks like the Internet. It offers various VPN types such as SSL VPN, IPsec VPN, and L2TP.
- Antivirus and Antimalware: FortiGate includes antivirus and antimalware functionalities to detect and block malicious software, including viruses, worms, Trojans, and spyware. It can inspect files and URLs in real time to prevent the spread of malware within the network.
- Web Filtering: FortiGate can enforce web filtering policies to control access to websites based on categories, URLs, or specific keywords. It helps organizations enforce acceptable use policies, improve productivity, and mitigate security risks associated with malicious or inappropriate web content.
- Application Control: FortiGate offers application control features to identify and control the usage of various applications within the network. It can classify applications based on their behavior and characteristics, allowing administrators to define policies to permit, deny, or limit access to specific applications.
- Data Loss Prevention (DLP): FortiGate includes DLP capabilities to prevent the unauthorized transmission of sensitive data outside the network. It can inspect outgoing traffic for predefined data patterns such as credit card numbers, social security numbers, or intellectual property, and enforce policies to prevent data leakage.
- Advanced Threat Protection (ATP): FortiGate integrates advanced threat protection mechanisms such as sandboxing and behavior-based analysis to detect and block sophisticated threats like zero-day exploits and targeted attacks. It isolates suspicious files in a sandbox environment to observe their behavior before allowing them into the network.
- Traffic Shaping and Quality of Service (QoS): FortiGate allows administrators to prioritize and control network traffic based on predefined policies. It can allocate bandwidth, enforce traffic shaping rules, and ensure quality of service for critical applications to optimize network performance and user experience.
- Logging and Reporting: FortiGate provides extensive logging and reporting capabilities to track network activity, security events, and policy violations. It generates detailed reports and alerts for administrators to analyze security incidents, troubleshoot issues, and maintain compliance with regulatory requirements.
FortiGuard is a comprehensive security intelligence service provided by Fortinet that offers real-time updates and protection against emerging threats for Fortinet products, including FortiGate. FortiGuard queries and packages play a crucial role in keeping security solutions up-to-date and effective. Here's a detailed explanation:
FortiGuard queries are requests made by Fortinet security products, such as FortiGate firewalls, to the FortiGuard service infrastructure. These queries are initiated to retrieve the latest threat intelligence, security updates, and other relevant information needed to enhance the security posture of the network. Key aspects of FortiGuard queries include:
- Real-time Threat Intelligence: FortiGuard queries fetch real-time threat intelligence data from FortiGuard Labs, Fortinet's global threat research team. This includes information on new malware signatures, vulnerabilities, zero-day exploits, and other security threats.
- URL Filtering Updates: FortiGuard queries also update URL filtering databases to ensure accurate categorization of websites and protection against web-based threats. This helps in enforcing web filtering policies and blocking access to malicious or inappropriate websites.
- Antivirus and Antimalware Definitions: FortiGuard queries update antivirus and antimalware definitions to detect and block the latest malware strains, including viruses, worms, Trojans, and spyware. This ensures that the security solution can identify and neutralize evolving threats effectively.
- IPS Signatures: FortiGuard queries update intrusion prevention system (IPS) signatures to detect and prevent known network attacks, including vulnerabilities, exploits, and protocol anomalies. Keeping IPS signatures up to date is crucial for protecting against emerging threats and vulnerabilities.
- Application Control Updates: FortiGuard queries fetch updates related to application control, allowing FortiGate firewalls to identify and control the usage of various applications within the network. This includes new application signatures, behavior patterns, and categorization updates.
- DLP Definitions: FortiGuard queries update data loss prevention (DLP) definitions to prevent the unauthorized transmission of sensitive data outside the network. This includes patterns for detecting sensitive information such as credit card numbers, social security numbers, and intellectual property.
FortiGuard packages are bundles of security updates, intelligence feeds, and definitions provided by Fortinet as part of the FortiGuard subscription service. These packages contain the latest threat intelligence and updates necessary to keep Fortinet security solutions, including FortiGate firewalls, up to date and protected against evolving threats. Key aspects of FortiGuard packages include:
