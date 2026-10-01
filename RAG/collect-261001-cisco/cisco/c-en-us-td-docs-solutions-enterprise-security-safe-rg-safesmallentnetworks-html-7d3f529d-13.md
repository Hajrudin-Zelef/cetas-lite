---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-13
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["Microsoft", "Oracle"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [416, 444]
sha256: 18719c06c142345c6fe96a82a395df146b0b07d6c9165628e8cec58443a42daa
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

NAC Appliance Deployment in Small Enterprise Networks
In the small enterprise network design, various components of the NAC Appliance solution are deployed throughout all locations, including the main site and each of the remote sites. More precisely, a centralized CAM is deployed at the main site, likely residing at the serverfarm, while a CAS is deployed at the main site and each remote site. Each CAS is connected directly to the core/distribution layer at each of those locations.
For simplicity, the Small Enterprise Network design uses the same design principles for all locations, providing a common NAC configuration that works regardless of whether the sites use a Layer 2 or Layer 3 access model. This common configuration consists in the use of an untrusted VLAN that extends from the access layer all the way to the untrusted interface of the CAS appliance. This requires the configuration of a trunk between the access switch and the core/distribution layer carrying a pair of VLANs. One VLAN of the trunk is the untrusted VLAN, used for authentication and posture assessment, and the second VLAN is the trusted VLAN used to carry all traffic once the client is properly authenticated and authorized. The link between the access switch and the core/distribution layer is configured as a trunk, regardless of whether the sites use a Layer 2 or Laye 3 access model.
This is illustrated in Figure 19 and Figure 20. In Figure 19, there is a simple Layer 2 NAC OOB connection where a client device upon initial connection to the network is given VLAN 264, which connects them directly to the untrusted interface of the NAS. VLAN 264 on the untrusted interface is mapped to VLAN 64 on the trusted interface within the NAC appliance, which allows the client to obtain an IP address that belongs on VLAN 64. Upon successful completion of the NAC authentication and validation functions, the access switch is instructed, via SNMP from the CAM, to change the client VLAN to VLAN 64. Even though the client has changed Layer 2 VLANs, its Layer 3 network connections are unchanged and the traffic from the client no longer passes through the NAC appliance.
Figure 19 Layer 2 OOB Topology
In Figure 20, the same processes are followed when the client is untrusted, but once the client has successfully completed its NAC functions, the access switch is instructed via SNMP to change the client VLAN to VLAN 67—a subnet local to the access switch. As the Layer 3 information for the client has changed, the switch is also instructed to "bounce" the client switch port to initiate a new DHCP request for an IP address appropriate to VLAN 67.
Figure 20 Layer 3 OOB Topology
Cisco Identity-Based Network Networking Services (IBNS)
The best and most secure solution to vulnerability at the access edge is to leverage the intelligence of the network. The Cisco Identity-Based Network Networking Services solution (IBNS) is a set of Cisco IOS software services designed to enable secure user and host access to enterprise networks powered by Cisco Catalyst switches and wireless LANs. It provides standards-based network access control at the access layer by using the 802.1X protocol to secure the physical ports where end users connect. 802.1X is an IEEE standard for media-level (Layer 2) access control, offering the capability to permit or deny network connectivity based on the identity of the end user or device. 802.1X is well-known as a way to secure wireless network access. It is equally essential in securing wired network access.
IEEE 802.1X Protocol
The IEEE 802.1X protocol allows Cisco Catalyst switches to offer network access control at the port level. Every port on the switch is individually enabled or disabled based on the identity of the user or device connecting to it. When 802.1X is first enabled on a port, the switch automatically drops all traffic received on that port. There is one exception to this rule. The only traffic a switch will accept is a request to start 802.1X authentication. Only after the 802.1X authentication has successfully completed will the switch accept any other kind of traffic on the port.
The high-level message exchange in Figure 21 illustrates how port-based access control works within an identity-based system.
Figure 21 Port-Based Access Control
The following steps describe the port-based access control flow shown in Figure 21:
1. A client, such as a laptop with an 802.1X supplicant, connects to an IEEE 802.1X-enabled network and sends a start message to the LAN switch (the authenticator).
2. When the start message is received, the LAN switch sends a login request to the client.
3. The client replies with a login response.
4. The switch forwards the response to the policy database (authentication server).
5. The authentication server authenticates the user.
6. After the user identity is confirmed, the policy database authorizes network access for the user and informs the LAN switch.
7. The LAN switch then enables the port connected to the client.
The user and device credentials are processed by an AAA server. The AAA server is able to reference user or device policy profile information either internally, using the integrated user database, or externally, using database sources such as Microsoft Active Directory, LDAP, Novell NDS, or Oracle databases. This enables the integration of the system into exiting user management structures and schemes, thereby simplifying overall deployment.
802.1X and EAP
When authenticating users for the purposes of network access control, the system must provide user and/or device identification using strong authentication technologies known to be secure and reliable. IEEE 802.1X does not by itself dictate how this is achieved. Rather, the 802.1X protocol defines an encapsulation for the transport of the Extensible Authentication Protocol (EAP) from the client to the switch. The 802.1X encapsulation is sometimes referred to as EAP over LAN (EAPoL). The switch in turn relays the EAP information to the authentication server using the RADIUS protocol (EAP over RADIUS).
EAP, which is defined by RFC 3748, is itself a framework and not a specific authentication method. EAP provides a way for the client and the authentication server to negotiate an authentication method that they both support. There are many EAP methods but the ones used more frequently for 802.1X wired authentication include EAP-TLS, EAP-PEAP, and EAP-FAST.
Impacts of 802.1X on the Network
Before enabling 802.1X in the network, it is essential to review the default security posture of a port enabled for 802.1X authentication: all traffic is dropped except 802.1X EAPoL packets. This is a fundamental change from the traditional model in which the port is enabled and all traffic is allowed from the moment that a device plugs into the port. Ports that were traditionally open will now be closed by default. This is one of the cornerstones of the strong security and network access control provided by 802.1X. However, this change in the default network access model can have a profound impact on network devices and applications. Understanding and providing for the impacts of this change will make for a smooth deployment of 802.1X network access control.
Non-802.1X-Enabled Devices
802.1X must be enabled on both the host device and on the switch to which the device connects. If a device without an 802.1X supplicant attempts to connect to a port that is enabled for 802.1X, it will be subjected to the default security policy. The default security policy says that 802.1X authentication must succeed before access to the network is granted. Therefore, by default, non-802.1X-capable devices cannot get access to an 802.1X-protected network.
