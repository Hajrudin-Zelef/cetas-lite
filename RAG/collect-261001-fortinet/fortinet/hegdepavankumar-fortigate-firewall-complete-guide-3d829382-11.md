---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-11
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses", "parameters"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [821, 892]
sha256: e7d94770217ef156b16fe691e2e2bc1bab8eae6be98e691b968f911dff6b42ca
---

# Example SSH command to connect to the FortiGate firewall

- Purpose: Key Management Protocols are used to negotiate and manage encryption keys required for establishing secure communication channels between IPsec peers.
- Protocols: Common Key Management Protocols include Internet Key Exchange (IKE) and IKEv2, which automate the negotiation and exchange of encryption keys, authentication parameters, and security policies.
- Purpose: ESP provides confidentiality, integrity, and authentication for IP packets by encapsulating the payload in a secure envelope, encrypting it, and adding authentication data.
- Functions: ESP encrypts the IP payload using encryption algorithms such as AES or 3DES, adds authentication data to ensure packet integrity, and optionally provides anti-replay protection.
- Purpose: AH provides data integrity, authentication, and anti-replay protection for IP packets by including a hash-based message authentication code (HMAC) in each packet.
- Functions: AH calculates a hash-based authentication code over the IP packet's entire contents, including the IP header, ensuring the integrity of the packet and detecting any modifications.
IKE (Internet Key Exchange) is a key management protocol used in IPsec VPNs to establish and manage secure communication channels between peers. IKE operates in two phases, known as Phase 1 and Phase 2, each serving distinct purposes in the VPN establishment process. Let's delve into the details of each phase:
- Purpose: Phase 1 establishes a secure channel for negotiating the parameters required for subsequent communications, including encryption algorithms, authentication methods, and session keys.
- Main Mode Exchange: Phase 1 typically employs Main Mode exchange, where peers authenticate each other and exchange encryption keys securely.
- Components:
  - Authentication: Peers authenticate each other using pre-shared keys (PSK) or digital certificates to ensure mutual trust.
  - Diffie-Hellman (DH) Exchange: Peers perform a DH key exchange to generate a shared secret used to derive encryption keys.
  - Encryption and Integrity: Phase 1 negotiates encryption and integrity algorithms to protect subsequent IKE and IPsec communications.
- Parameters Negotiated: IKE Phase 1 negotiates parameters such as encryption algorithm, authentication method, DH group, and session lifetime.
Want to Know more about IKEv1: Click Here
- Purpose: Phase 2 establishes security associations (SAs) for IPsec traffic, defining the parameters for encrypting and authenticating data packets exchanged between peers.
- Quick Mode Exchange: Phase 2 typically uses Quick Mode exchange, which focuses on negotiating IPsec parameters efficiently.
- Components:
  - Encryption and Authentication: Phase 2 negotiates encryption and authentication algorithms specifically for IPsec traffic protection.
  - IPsec SA Establishment: Peers establish one or more IPsec SAs, each defining the parameters for encrypting and authenticating data packets.
  - Traffic Selector Negotiation: Peers agree on which traffic will be protected by IPsec, defining source and destination IP addresses and protocols.
- Parameters Negotiated: IKE Phase 2 negotiates parameters such as encryption algorithm, authentication method, IPsec mode (tunnel or transport), and IPsec SAs.
IPsec VPNs provide secure communication between networks or devices over the internet by encrypting and authenticating data traffic. Configuring IPsec between FortiGate firewalls involves several steps to establish a secure VPN tunnel. Here's a detailed guide:
- Ensure both FortiGate firewalls have valid licenses for IPsec VPN.
- Determine the public IP addresses of both FortiGate firewalls.
- Establish connectivity between the public IP addresses of the two FortiGate firewalls.
- Navigate to VPN > IPsec > Wizard.
- Select Custom VPN Tunnel (No Template) and click Next.
- Enter a Name for the VPN tunnel and click Next.
- Configure Phase 1 parameters:
  - Remote Gateway: Public IP address of FortiGate B.
  - Authentication Method: Pre-shared Key or Digital Certificate.
  - Pre-shared Key: Enter a shared secret for authentication.
  - Encryption Algorithm: Select desired encryption algorithm.
  - Authentication Algorithm: Select desired authentication algorithm.
  - Diffie-Hellman Group: Select DH group for key exchange.
- Click Next and then Finish.
- Repeat the same steps as above, configuring Phase 1 parameters with the public IP address of FortiGate A.
- Navigate to VPN > IPsec > Wizard.
- Select Custom VPN Tunnel (No Template) and click Next.
- Enter the same Name used for Phase 1 and click Next.
- Configure Phase 2 parameters:
  - Remote Gateway: Public IP address of FortiGate B.
  - Local Interface: Select the outgoing interface.
  - Encryption Algorithm: Select the encryption algorithm.
  - Authentication Algorithm: Select the authentication algorithm.
  - PFS: Enable Perfect Forward Secrecy if required.
- Click Next and then Finish.
- Repeat the same steps as above, configuring Phase 2 parameters with the public IP address of FortiGate A.
- Once Phase 1 and Phase 2 configurations are completed on both FortiGate firewalls, the VPN tunnel should be established automatically.
- Navigate to VPN > Monitor > IPsec Monitor to verify the status of the VPN tunnel.
- Monitor traffic and logs to ensure the proper functioning of the IPsec VPN tunnel.
Sample Topology:
Site-A to Site-B [Site-to-Site] VPN. Go to VPN —> IPsec Wizard
Enter the Remote Site Public IP along with PSK(pre-shared key).
Now it is time to add the interesting network (private subnet of both sides)
Configure the Tunnel both the sides:(once both side tunnel configured and both the phase negotiation complete tunnel come up or you need to manually bring up the tunnel)
once you passed the traffic between two peers the tunnel comes up:
Automatically Policy added: to allowing traffic from two different sites.
Static Routes between two Tunnel:
To Monitoring the VPN Tunnel traffic and User actions:
- Covers the architecture of IPsec, detailing its components and protocols for securing internet communications.
- Explains IKE Phase 1 & 2, which are crucial for establishing and managing secure VPN connections between peers.
- Provides a step-by-step guide on configuring IPsec VPN between FortiGate firewalls, ensuring secure communication channels between networks or devices.
This module equips administrators with essential knowledge and practical skills for setting up IPsec VPNs using FortiGate firewalls, enhancing network security, and enabling secure communication over the internet.
- Module 1: Introduction to FortiGate Firewall: Provided an overview of FortiGate firewall features, platform design, CLI access, management GUI access, and administration profiles.
- Module 2: Interface Configurations & Firewall Policies: Covered basic interface configurations, routing configurations, DHCP setup, firewall policies, and Network Address Translation (NAT).
- Module 3: High Availability: Explored high availability concepts including active-standby and active-active failover setups.
- Module 4: Firewall Authentication: Discussed user and policy creation, authentication policies, and user monitoring, including captive portal setups.
- Module 5: Security Profiles: Detailed security profiles such as application control, web filtering, antivirus, intrusion prevention, and SSL/SSH inspection to enforce security policies.
- Module 6: Logging and Monitoring: Covered log severity levels, log types, log structures, configuring log settings, and redirecting logs to external systems like Syslog and SNMP.
- Module 7: Basic IPSEC VPN: Provided an understanding of IPsec architecture, IKE Phase 1 & 2, and practical guidance on configuring IPsec VPNs between FortiGate firewalls.
