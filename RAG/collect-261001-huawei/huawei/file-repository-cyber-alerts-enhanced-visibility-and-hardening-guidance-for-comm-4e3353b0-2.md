---
id: collect-261001-huawei/huawei/file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0-2
title: "file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0"
domain: huawei
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-huawei/file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0.md
source_anchor: ""
source_lines: [108, 194]
sha256: cd5d685da1c09fa97f9ef1077c847cac2c2f416d8f0ba16ca97c59f8c46583b3
---

# file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0

TLP:CLEAR 
CISA | FBI | NSA | ASD’s ACSC | CCCS | NCSC-NZ 
TLP:CLEAR 
4 
Hardening Systems and Devices 
Hardening device and network architecture is a defense-in-depth strategy. Reducing vulnerabilities, 
improving secure configuration habits, and following best practices limit potential entry points for 
PRC-affiliated and other cyber threats. 
Protocols and Management Processes 
Network Engineers 
 Use an out-of-band management network that is physically separate from the operational 
data flow network. Ensure that management of network infrastructure devices can only come 
from the out-of-band management network. In addition, confirm that the out-of-band 
management network does not allow lateral management connections between devices to 
prevent lateral movement in the case that one device becomes compromised. Ensure device 
management is physically isolated from the customer and production networks. When 
properly implemented, out-of-band management can mitigate many threat actor tactics, 
techniques, and procedures (TTPs). 
 Implement a strict, default-deny ACL strategy to control inbound and egressing traffic. Ensure 
all denied traffic is logged. For maximum depth, implement on separate devices from those 
implementing other security controls. 
 Employ strong network segmentation via the use of router ACLs, stateful packet inspection, 
firewall capabilities, and demilitarized zone (DMZ) constructs. Separation via virtual local 
area networks (VLANs) and, if possible, private VLANs (PVLAN) will provide additional 
granular logical separation. This should be done as part of a broader defense-in-depth 
approach that protects and isolates different device groups. 
o Place externally facing services, such as Domain Name System (DNS), web servers, and 
mail servers, in a DMZ to provide segmentation from the internal LAN and backend 
resources. 
o Additionally, as a general strategy, put devices with similar purposes in the same VLAN. 
For example, place all user workstations from a certain team in one VLAN, while putting 
another team with different functions in a separate VLAN. 
o Do not manage devices from the internet. Only allow device management from trusted 
devices on trusted networks. Use dedicated administrative workstations (DAWs) 
connected to dedicated management zones. 
 Harden and secure virtual private network (VPN) gateways by limiting external exposure, if 
possible, and limiting the port exposure to what is minimally required (for example udp/500, 
udp/4500 and protocol type 50 (ESP)). Ensure all VPNs are configured to only use strong 
cryptography for key exchange, authentication, and encryption.[1] 
o Disable unused VPN features and cryptographic algorithms to prevent exploitable 
weaknesses. 
 Ensure that traffic is end-to-end encrypted to the maximum extent possible.

TLP:CLEAR 
CISA | FBI | NSA | ASD’s ACSC | CCCS | NCSC-NZ 
TLP:CLEAR 
5 
 As a management policy, control access to device Virtual Teletype (VTY) lines with an ACL to 
restrict inbound lateral movement connections. 
o Additionally, disable outbound connections to mitigate against lateral movement. Monitor 
for changes as adversaries can modify this configuration on compromised devices to 
allow outbound connections. 
 Ensure all authentication, authorization, and accounting (AAA) logging is securely sent to a 
centralized logging server with modern confidentiality, integrity, and authentication (CIA) 
protections. 
 If using Simple Network Management Protocol (SNMP), ensure only SNMP v3 with encryption 
and authentication is used, along with ACL protections against unnecessary public exposure. 
Ensure configuration with the most secure cryptographic options supported by the hardware. 
 Disable all unnecessary discovery protocols, such as Cisco Discovery Protocol (CDP) or Link 
Layer Discovery Protocol (LLDP). If they are required, only enable on the necessary 
interfaces. 
 Ensure Transport Layer Security (TLS) v1.3 is used on any TLS-capable protocols to secure 
data in transit over a network.[2] Ensure TLS is configured to only use strong cryptographic 
cipher suites.[3] 
o Use Public Key Infrastructure (PKI)-based certificates instead of self-signed certificates. 
o Implement a robust process to renew certificates before they expire. 
 Disable Internet Protocol (IP) source routing. 
 Disable Secure Shell (SSH) version 1. Ensure only SSH version 2.0 is used with the following 
cryptographic considerations.[2] For more information on acceptable algorithms, see NSA’s 
Network Infrastructure Security Guide. 
o Configure with minimally a 3072-bit RSA key. 
o Configure with minimally a 4096 Diffie-Hellman key size (group 16). 
 When possible, apply secure authentication to protocols and services which allow it, such as 
Network Time Protocol (NTP), Terminal Access Controller Access-Control System (TACACS+), 
Open Shortest Path First (OSPF), Border Gateway Protocol (BGP), and Hot Standby Router 
Protocol (HSRP). Similarly, disable any unauthenticated management protocols or functions, 
such as Cisco Smart Install. 
 Use secure cryptographic building blocks when building VPNs such as [3]: 
o Key Exchange: 
 Diffie-Hellman Group 15 with 3072-bit Modular Exponential (MODP) 
 Diffie-Hellman Group 16 with 4096-bit Modular Exponential (MODP) 
 Diffie-Hellman Group 20 with 384-bit Elliptic Curve Group (ECP) 
o Encryption: AES-256 
o Hashing: SHA-384 or SHA-512 
 Ensure that no default passwords are used. 
o Change all default passwords on first use.

