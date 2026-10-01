---
id: collect-261001-huawei/huawei/resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd-2
title: "resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd"
domain: huawei
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd.md
source_anchor: ""
source_lines: [33, 79]
sha256: e66b8beb715c55472241683302f618d8c699b8b9161d3f1adcbfc8dc890422e7
---

# resources-tools-resources-enhanced-visibility-and-hardening-guidance-communicati-1615d2dd

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
Network Defenders
- Disable any unnecessary, unused, exploitable, or plaintext services and protocols, such as Telnet, File Transfer Protocol (FTP), Trivial FTP (TFTP), SSH v1, Hypertext Transfer Protocol (HTTP) servers, and SNMP v1/v2c. Ensure any required internet-exposed services are adequately protected by ACLs and are fully patched.
- Conduct port-scanning and scanning of known internet-facing infrastructure to ensure no additional services are accessible across the network or from the internet. Remove unnecessary internet-facing infrastructure, monitor necessary internet-facing infrastructure, and continuously validate the architecture.
  - Routers with an active shell environment—even if they have not been tampered with—have significantly more listeners running at the operating system (OS) level compared to the software level.
Network defenders and network engineers should ensure close collaboration and open communication to accomplish the following:
- Ensure all networking configurations are stored, tracked, and regularly audited for compliance with security policies and best practices.
  - Whenever networking configurations are transmitted for storage, tracking, and troubleshooting, confirm that they are sent using encrypted protocols. Additionally, be sure they are not attached to plaintext emails or sent via FTP or TFTP.
- Monitor for vendor end-of-life (EOL) announcements for hardware devices, operating system versions, and software, and upgrade as soon as possible.
- Implement a change management system that anticipates both routine and emergency patching. Continuously monitor for vendor vulnerability and patch announcements and ensure patches are applied in a timely manner. Ensure use of vendor recommended version of the operating system for the features and capabilities required.
  - Test and validate patches as part of the change and patch management processes.
- As part of a broader password policy, store passwords with secure hashing algorithms. Passwords should meet complexity requirements and should be stored using one-way hashing algorithms or, if available, unique keys. Follow National Institute of Standards and Technologies guidelines when creating password policies.
- Require phishing-resistant multi-factor authentication (MFA) for all accounts that access company systems, networks, and applications, including sensitive administrative access to routers. MFA should use a combination of credentials and a phishing-resistant secondary verification method, such as hardware-based PKI or FIDO authentication, to ensure secure access and prevent unauthorized entry.
- As part of a broader identity and access management policy, use local accounts only for emergencies and change the passwords after each use. Verify that each use was authorized and expected. For everyday management of network infrastructure, use a centralized AAA server that supports multi-factor authentication requirements; however, ensure the AAA server is not linked to the primary corporate identity store.
- Limit session token durations and require users to reauthenticate when the session expires. Conduct audits to determine the standard session duration for each role to implement session expirations.
- Implement a Role-Based Access Control (RBAC) strategy that assigns users to a specific role with defined and inherited permissions to better control and manage what users can do.
- Remove any unnecessary accounts and periodically review accounts to verify that they continue to be needed. Apply the principle of least privilege to make sure accounts only have the minimum permissions necessary to complete their tasks. Additionally, continuously monitor accounts in use.
Cisco-Specific Guidance
