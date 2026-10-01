---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-11
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [346, 381]
sha256: e279bd2640b4589987ef3eca3cf26aaf9d4e81a43947f8bcccaaaf0ad664cdae
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

•Storm Control—Prevents broadcast and multicast storms by monitoring packets passing from an interface to the switching bus and determines whether the packet is unicast, multicast, or broadcast. The switch counts the number of packets of a specified type received within the one-second time interval and compares the measurement with a predefined suppression-level threshold. When the suppression-level threshold is reached, the port blocks traffic until the traffic falls below the threshold level.
Cisco Unified Wireless Network (CUWN) Integrated Security Features
The Cisco Unified Wireless Network adds to the 802.11 security standards by providing additional security features. Some of these are the WLAN equivalent of CiSF features, such as Dynamic Host Configuration Protocol (DHCP) and Address Resolution Protocol (ARP) protection, peer-to-peer blocking, and access control list and firewall features. Additionally other more WLAN-specific features are provided, including Enhanced WLAN security options, wireless intrusion detection system (IDS), client exclusion, rogue AP detection, management frame protection, dynamic radio frequency management, and network IDS integration.
The Cisco Unified Wireless Network solutions are discussed in the Wireless and Network Security Integration Solution Design Guide at: http://www.cisco.com/en/US/solutions/ns340/ns414/ns742/ns820/landing_sec_wireless.html.
Cisco NAC Appliance
Cisco Network Admission Control (NAC) Appliance (formerly known as Cisco Clean Access) uses the network infrastructure to enforce security policy compliance on all devices seeking to access network computing resources. With Cisco NAC Appliance, network administrators can authenticate, authorize, evaluate, and remediate wired, wireless, and remote users and their machines before network access. The NAC Appliance identifies whether networked devices such as laptops, IP phones, or game consoles are compliant with your network security policies and can repair any vulnerability before permitting access to the network.
When deployed, the Cisco NAC Appliance provides the following benefits:
•Recognizes users, their devices, and their roles in the network. This first step occurs at the point of authentication, before malicious code can cause damage.
•Evaluates whether machines are compliant with security policies. Security policies can include specific anti-virus or anti-spyware software, OS updates, or patches. Cisco NAC Appliance supports policies that vary by user type, device type, or operating system.
•Enforces security policies by blocking, isolating, and repairing non-compliant machines.
•Non-compliant machines are redirected to a quarantine network, where remediation occurs at the discretion of the administrator.
The NAC solution provides the following four functions, as shown in Figure 14:
Figure 14 The Four Functions of the NAC Framework
For more details of the NAC Appliance Solution, see: http://www.cisco.com/go/nacappliance.
Cisco NAC Appliance is a network-centric integrated solution administered from the Cisco Clean Access Manager (NAC Manager) Web console and enforced through the Clean Access Server (NAC Server) and (optionally) the Clean Access Agent or Cisco NAC Web Agent. Cisco NAC Appliance checks client systems, enforces network requirements, distributes patches and antivirus software, and quarantines vulnerable or infected clients for remediation before clients access the network. Cisco NAC Appliance consists of the components shown in Figure 15.
Figure 15 NAC Appliance Components
Clean Access Manager (CAM)
The Cisco CAM (a.k.a NAC Manager) is the administration server for Clean Access deployment. The secure Web console of the Clean Access Manager is the single point of management for up to 20 Clean Access Servers in a deployment (or 40 CASs if installing a SuperCAM). For Out-of-Band (OOB) deployment, the Web admin console allows you to control switches and VLAN assignment of user ports through the use of SNMP. In the small enterprise network design the CAM would be located at the main site.
Clean Access Server (CAS)
The Cisco CAS (a.k.a NAC Server) is the enforcement server between the untrusted (managed) network and the trusted network. The CAS enforces the policies you have defined in the CAM Web admin console, including network access privileges, authentication requirements, bandwidth restrictions, and Clean Access system requirements. You can install a CAS as either a standalone appliance (like the Cisco NAC-3300 Series) or as a network module (Cisco NME-NAC-K9) in a Cisco ISR chassis and deploy it in-band (always inline with user traffic) or OOB (inline with user traffic only during authentication/posture assessment).
The CAS can also be deployed in Layer 2 mode (users are Layer 2 adjacent to the CAS) or Layer 3 mode (users are multiple Layer 3 hops away from the CAS). You can also deploy several CASs of varying size/capacity to fit the needs of varying network segments. You can install Cisco NAC-3300 Series appliances in your company headquarters core, for example, to handle thousands of users and simultaneously install one or more Cisco NAC network modules in ISR platforms to accommodate smaller groups of users at a satellite office.
In the small enterprise network design, the CAS would be located at the main site and the remote locations and it would be used to provide Layer 2 or Layer 3 OOB authentication/posture assessment.
Clean Access Agent (CAA)
CAA is an optional read-only agent that resides on Windows clients. It checks applications, files, services, or registry keys to ensure that clients meet your specified network and software requirements prior to gaining access to the network.
Note There is no client firewall restriction with CAA posture assessment. The agent can check the client registry, services, and applications even if a personal firewall is installed and running.
If NAC is implemented as part of the small enterprise network design, it is recommended that the CAA be used.
Cisco NAC Web Agent
The Cisco NAC Web Agent provides temporal posture assessment for client machines. Users launch the Cisco NAC Web Agent executable, which installs the Web Agent files in a temporary directory on the client machine via ActiveX control or Java applet. When the user terminates the Web Agent session, the Web Agent logs the user off of the network and their user ID disappears from the Online Users list.
Clean Access Policy Updates
Regular updates of prepackaged policies/rules that can be used to check the up-to-date status of operating systems, antivirus (AV), antispyware (AS), and other client software. It also provides built-in support for 24 AV vendors and 17 AS vendors.
NAC Appliance Modes and Positioning
NAC Appliance allows multiple deployment options and may be placed at different points in the network. The modes of operation can be generally defined as:
•Out-of-band (OOB) virtual gateway
Out-of-Band Modes
Out-of-Band (OOB) deployments require user traffic to traverse through the NAC Appliance only during authentication, posture assessment, and remediation. When a user is authenticated and passes all policy checks, their traffic is switched normally through the network and bypasses the appliance. See Figure 16.
Figure 16 Layer 2 OOB Topology
