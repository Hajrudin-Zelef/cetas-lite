---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972-1
title: "c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses", "memory", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972.md
source_anchor: ""
source_lines: [1, 62]
sha256: f8b9f24eb42935ce86f4193c24d71fd5d9dbe2374ee7ccdc2773254710b4df1b
---

# c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972

LAN-to-LAN IPsec VPNs
A LAN-to-LAN VPN connects networks in different geographic locations.
The ASA supports LAN-to-LAN VPN connections to Cisco or third-party peers when the two peers have IPv4 inside and outside networks (IPv4 addresses on the inside and outside interfaces).
For LAN-to-LAN connections using mixed IPv4 and IPv6 addressing, or all IPv6 addressing, the security appliance supports VPN tunnels if both peers are ASAs and if both inside networks have matching addressing schemes (both IPv4 or both IPv6).
Specifically, the following topologies are supported when both peers are ASAs:
- The ASAs have IPv4 inside networks and the outside network is IPv6 (IPv4 addresses on the inside interfaces and IPv6 addresses on the outside interfaces).
- The ASAs have IPv6 inside networks and the outside network is IPv4 (IPv6 addresses on the inside interface and IPv4 addresses on the outside interfaces).
- The ASAs have IPv6 inside networks and the outside network is IPv6 (IPv6 addresses on the inside and outside interfaces).
Note The ASA supports LAN-to-LAN IPsec connections with Cisco peers, and with third-party peers that comply with all relevant standards.
This chapter describes how to build a LAN-to-LAN VPN connection. It includes the following sections:
Summary of the Configuration
This section provides a summary of the example LAN-to-LAN configuration this chapter describes. Later sections provide step-by-step instructions.
Configuring Site-to-Site VPN in Multi-Context Mode
Follow these steps to allow site-to-site support in multi-mode for all platforms except the 5505. By performing these steps, you can see how resource allocation breaks down.
Step 1 To configure the VPN in multi-mode, configure a resource class and choose VPN licenses as part of the allowed resource. The "Configuring a Class for Resource Management" provides these configuration steps. The following is an example configuration:
Step 2 Configure a context and make it a member of the configured class that allows VPN licenses. The "Configuring a Security Contextt" provides these configuration steps. The following is an example configuration:
Step 3 Configure connection profiles, policies, crypto maps, and so on, just as would with single context VPN configuration of site-to-site VPN.
Configuring Interfaces
An ASA has at least two interfaces, referred to here as outside and inside. Typically, the outside interface is connected to the public Internet, while the inside interface is connected to a private network and is protected from public access.
To begin, configure and enable two interfaces on the ASA. Then, assign a name, IP address and subnet mask. Optionally, configure its security level, speed, and duplex operation on the security appliance.
Note The ASA’s outside interface address (for both IPv4/IPv6) cannot overlap with the private side address space.
To configure interfaces, perform the following steps, using the command syntax in the examples:
Step 1 To enter Interface configuration mode, in global configuration mode enter the interface command with the default name of the interface to configure. In the following example the interface is ethernet0.
Step 2 To set the IP address and subnet mask for the interface, enter the ip address command. In the following example the IP address is 10.10.4.100 and the subnet mask is 255.255.0.0.
Step 3 To name the interface, enter the nameif command, maximum of 48 characters. You cannot change this name after you set it. In the following example the name of the ethernet0 interface is outside.
Step 4 To enable the interface, enter the no version of the shutdown command. By default, interfaces are disabled.
Step 5 To save your changes, enter the write memory command:
Step 6 To configure a second interface, use the same procedure.
Configuring ISAKMP Policy and Enabling ISAKMP on the Outside Interface
ISAKMP is the negotiation protocol that lets two hosts agree on how to build an IPsec security association (SA). It provides a common framework for agreeing on the format of SA attributes. This includes negotiating with the peer about the SA, and modifying or deleting the SA. ISAKMP separates negotiation into two phases: Phase 1 and Phase 2. Phase 1 creates the first tunnel, which protects later ISAKMP negotiation messages. Phase 2 creates the tunnel that protects data.
IKE uses ISAKMP to setup the SA for IPsec to use. IKE creates the cryptographic keys used to authenticate peers.
The ASA supports IKEv1 for connections from the legacy Cisco VPN client, and IKEv2 for the AnyConnect VPN client.
To set the terms of the ISAKMP negotiations, you create an IKE policy, which includes the following:
- The authentication type required of the IKEv1 peer, either RSA signature using certificates or preshared key (PSK).
- An encryption method, to protect the data and ensure privacy.
- A Hashed Message Authentication Codes (HMAC) method to ensure the identity of the sender, and to ensure that the message has not been modified in transit.
- A Diffie-Hellman group to determine the strength of the encryption-key-determination algorithm. The ASA uses this algorithm to derive the encryption and hash keys.
- For IKEv2, a separate pseudo-random function (PRF) used as the algorithm to derive keying material and hashing operations required for the IKEv2 tunnel encryption.
- A limit to the time the ASA uses an encryption key before replacing it.
With IKEv1 policies, for each parameter, you set one value. For IKEv2, you can configure multiple encryption and authentication types, and multiple integrity algorithms for a single policy. The ASA orders the settings from the most secure to the least secure and negotiates with the peer using that order. This allows you to potentially send a single proposal to convey all the allowed transforms instead of the need to send each allowed combination as with IKEv1.
The following sections provide procedures for creating IKEv1 and IKEv2 policies and enabling them on an interface:
Configuring ISAKMP Policies for IKEv1 Connections
To configure ISAKMP policies for IKEv1 connections, use the crypto ikev1 policy priority command to enter IKEv1 policy configuration mode where you can configure the IKEv1 parameters.
Perform the following steps and use the command syntax in the following examples as a guide.
Step 1 Enter IPsec IKEv1 policy configuration mode. For example:
Step 2 Set the authentication method. The following example configures a preshared key:
Step 3 Set the encryption method. The following example configures 3DES:
Step 4 Set the HMAC method. The following example configures SHA-1:
Step 5 Set the Diffie-Hellman group. The following example configures Group 2:
Step 6 Set the encryption key lifetime. The following example configures 43,200 seconds (12 hours):
Step 7 Enable IKEv1 on the interface named outside in either single or multiple context mode:
Step 8 To save your changes, enter the write memory command:
Configuring ISAKMP Policies for IKEv2 Connections
To configure ISAKMP policies for IKEv2 connections, use the crypto ikev2 policy priority command to enter IKEv2 policy configuration mode where you can configure the IKEv2 parameters.
Step 1 Enter IPsec IKEv2 policy configuration mode. For example:
Step 2 Set the encryption method. The following example configures 3DES:
Step 3 Set the Diffie-Hellman group. The following example configures Group 2:
Step 4 Set the pseudo-random function (PRF) used as the algorithm to derive keying material and hashing operations required for the IKEv2 tunnel encryption. The following example configures SHA-1 (an HMAC variant):
Step 5 Set the encryption key lifetime. The following example configures 43,200 seconds (12 hours):
Step 6 Enable IKEv2 on the interface named outside:
Step 7 To save your changes, enter the write memory command:
Creating an IKEv1 Transform Set
