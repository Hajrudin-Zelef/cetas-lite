---
id: collect-261001-huawei/huawei/enhanced-visibility-and-hardening-guidance-for-communications-infrastructure-cisa-2
title: "Enhanced Visibility and Hardening Guidance for Communications Infrastructure"
domain: huawei
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/enhanced-visibility-and-hardening-guidance-for-communications-infrastructure-cisa.md
source_anchor: ""
source_lines: [43, 86]
sha256: 07ffd9d26314e91e5a42b34caa114c6e881c57857653ec48b52c6a2a27d64663
---

# Enhanced Visibility and Hardening Guidance for Communications Infrastructure

- Use an out-of-band management network that is physically separate from the operational data flow network. Ensure that management of network infrastructure devices can only come from the out-of-band management network. In addition, confirm that the out-of-band management network does not allow lateral management connections between devices to prevent lateral movement in the case that one device becomes compromised. Ensure device management is physically isolated from the customer and production networks. When properly implemented, out-of-band management can mitigate many threat actor tactics, techniques, and procedures (TTPs).
- Implement a strict, default-deny ACL strategy to control inbound and egressing traffic. Ensure all denied traffic is logged. For maximum depth, implement on separate devices from those implementing other security controls.
- Employ strong network segmentation via the use of router ACLs, stateful packet inspection, firewall capabilities, and demilitarized zone (DMZ) constructs. Separation via virtual local area networks (VLANs) and, if possible, private VLANs (PVLAN) will provide additional granular logical separation. This should be done as part of a broader defense-in-depth approach that protects and isolates different device groups.
  - Place externally facing services, such as Domain Name System (DNS), web servers, and mail servers, in a DMZ to provide segmentation from the internal LAN and backend resources.
  - Additionally, as a general strategy, put devices with similar purposes in the same VLAN. For example, place all user workstations from a certain team in one VLAN, while putting another team with different functions in a separate VLAN.
  - Do not manage devices from the internet. Only allow device management from trusted devices on trusted networks. Use dedicated administrative workstations (DAWs) connected to dedicated management zones.
- Harden and secure virtual private network (VPN) gateways by limiting external exposure, if possible, and limiting the port exposure to what is minimally required (for example udp/500, udp/4500 and protocol type 50 (ESP)). Ensure all VPNs are configured to only use strong cryptography for key exchange, authentication, and encryption. [1]
  - Disable unused VPN features and cryptographic algorithms to prevent exploitable weaknesses.
- Ensure that traffic is end-to-end encrypted to the maximum extent possible.
- As a management policy, control access to device Virtual Teletype (VTY) lines with an ACL to restrict inbound lateral movement connections.
  - Additionally, disable outbound connections to mitigate against lateral movement. Monitor for changes as adversaries can modify this configuration on compromised devices to allow outbound connections.
- Ensure all authentication, authorization, and accounting (AAA) logging is securely sent to a centralized logging server with modern confidentiality, integrity, and authentication (CIA) protections.
- If using Simple Network Management Protocol (SNMP), ensure only SNMP v3 with encryption and authentication is used, along with ACL protections against unnecessary public exposure. Ensure configuration with the most secure cryptographic options supported by the hardware.
- Disable all unnecessary discovery protocols, such as Cisco Discovery Protocol (CDP) or Link Layer Discovery Protocol (LLDP). If they are required, only enable on the necessary interfaces.
- Ensure Transport Layer Security (TLS) v1.3 is used on any TLS-capable protocols to secure data in transit over a network. [2] Ensure TLS is configured to only use strong cryptographic cipher suites. [3]
  - Use Public Key Infrastructure (PKI)-based certificates instead of self-signed certificates.
  - Implement a robust process to renew certificates before they expire.
- Disable Internet Protocol (IP) source routing.
- Disable Secure Shell (SSH) version 1. Ensure only SSH version 2.0 is used with the following cryptographic considerations [2]. For more information on acceptable algorithms, see NSA’s Network Infrastructure Security Guide.
  - Configure with minimally a 3072-bit RSA key.
  - Configure with minimally a 4096 Diffie-Hellman key size (group 16).
- When possible, apply secure authentication to protocols and services which allow it, such as Network Time Protocol (NTP), Terminal Access Controller Access-Control System (TACACS+), Open Shortest Path First (OSPF), Border Gateway Protocol (BGP), and Hot Standby Router Protocol (HSRP). Similarly, disable any unauthenticated management protocols or functions, such as Cisco Smart Install.
- Use secure cryptographic building blocks when building VPNs such as [3]:
  - Key Exchange:
    - Diffie-Hellman Group 15 with 3072-bit Modular Exponential (MODP)
    - Diffie-Hellman Group 16 with 4096-bit Modular Exponential (MODP)
    - Diffie-Hellman Group 20 with 384-bit Elliptic Curve Group (ECP)
  - Encryption: AES-256
  - Hashing: SHA-384 or SHA-512
- Key Exchange:
- Ensure that no default passwords are used.
  - Change all default passwords on first use.
  - Ensure no passwords are reset back to the default.
- Confirm the integrity of the software image in use by using a trusted hashing calculation utility, if available.
  - If a utility is unavailable, calculate a hash of the software image on a trusted administration workstation and compare against the vendor’s published hashes on an authenticated site as a trusted source of truth. This may require engaging the device’s maintenance contract to access source of truth hash values. For additional security, copy the image to a forensic workstation and calculate the hash value to compare against the vendor’s published hashes.

#### **Network Defenders**

- Disable any unnecessary, unused, exploitable, or plaintext services and protocols, such as Telnet, File Transfer Protocol (FTP), Trivial FTP (TFTP), SSH v1, Hypertext Transfer Protocol (HTTP) servers, and SNMP v1/v2c. Ensure any required internet-exposed services are adequately protected by ACLs and are fully patched.
- Conduct port-scanning and scanning of known internet-facing infrastructure to ensure no additional services are accessible across the network or from the internet. Remove unnecessary internet-facing infrastructure, monitor necessary internet-facing infrastructure, and continuously validate the architecture.
  - Routers with an active shell environment—even if they have not been tampered with—have significantly more listeners running at the operating system (OS) level compared to the software level.

Network defenders and network engineers should ensure close collaboration and open communication to accomplish the following:

