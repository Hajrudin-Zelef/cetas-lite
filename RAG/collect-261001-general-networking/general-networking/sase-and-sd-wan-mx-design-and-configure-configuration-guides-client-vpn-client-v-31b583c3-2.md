---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-31b583c3-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-31b583c3"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2024-07"]
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-31b583c3.md
source_anchor: ""
source_lines: [86, 132]
sha256: 52011ee156b3e2bdd55ced393bc9737043ed60e4b4a225769d6d41dc15974372
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-31b583c3

  - Description: Username
  - Email: User email address
  - Password: Enter a password for the user or click "Generate" to automatically generate a password
  - Authorized: Select whether this user is authorized to use the client VPN
To edit an existing user, click on the username in the User Management section.
To delete a user, click the X next to the username on the right side of the user list.
The user email address is the username used for authentication in Meraki Cloud Authentication.
RADIUS
Use this option to authenticate users on a RADIUS server.
- In Security & SD-WAN > Configure > Client VPN click Add a RADIUS server to configure the server(s) to use
- In the Host field, enter the IP address of the RADIUS server
- In the Port field, enter the port to be used for RADIUS communication
- In the Secret field, enter the shared secret for the RADIUS server
For more information on how to configure RADIUS authentication server for client VPN, see Configuring RADIUS Authentication with Client VPN.
If multiple RADIUS servers are configured, RADIUS traffic will not be load-balanced. Server IP addresses are sorted in ascending order from lowest IP to highest IP. RADIUS traffic towards the servers are prioritized in the same manner, starting with the lowest IP address first. This behaviour cannot be modified.
When the IP address of the RADIUS authentication server is reachable over site-to-site VPN, the source IP used for the RADIUS traffic will be that of the highest-numbered VLAN that is VPN-enabled.
If there is no VPN-enabled VLAN or the MX is in passthrough/concentrator mode, the IP address will be within 6.0.0.0/8 (6.x.y.z).
This is described in more detail here.
For details about the RADIUS Message-Authenticator verification option, see RADIUS Protocol Spoofing Vulnerability (Blast-RADIUS): July 2024.
Active Directory
Use this option if user authentication should be done with Active Directory domain credentials.
In Security & SD-WAN > Configure > Active Directory > Active Directory servers, the following information is required for configuration:
- Short domain: The short name of the Active Directory domain
- Server IP: The IP address of an Active Directory server on the MX LAN or a remote subnet routable through Auto VPN
- Domain admin: The domain administrator account the MX should use to query the server
- Password: The password for the domain administrator account
For more information on how to configure the Active Directory authentication server for client VPN, see Configuring Active Directory with MX Security Appliances.
Security appliances do not support mapping group policies via Active Directory for users connecting through the client VPN.
When the IP address of the Active Directory authentication server is reachable over site-to-site VPN, the source IP used for the LDAP traffic will be that of the highest-numbered VLAN that is VPN-enabled.
If there is no VPN-enabled VLAN or the MX is in passthrough/concentrator mode, the IP address will be within 6.0.0.0/8 (6.x.y.z).
This is described in more detail here.
Systems Manager Sentry VPN Security
When using Meraki Cloud Authentication, Systems Manager Sentry VPN security can be configured if your dashboard organization contains one or more Mobile Device Management (MDM) networks. Systems Manager Sentry VPN security allows for devices enrolled in Systems Manager to receive the configuration to connect to the client VPN through the Systems Manager profile on the device.
To enable Systems Manager Sentry VPN security, go to Security & SD-WAN > Configure > Client VPN > IPsec Settings page. Under the Client VPN Server drop-down menu, select Enabled. You can configure the following options:
- Install scope: The install scope allows for a selection of Systems Manager tags for a particular MDM network. Devices with these tags applied in a Systems Manager network will receive a configuration to connect to this network's client VPN server through their Systems Manager profile.
    
  - Send all traffic: Select whether all client traffic should be sent to the MX.
- Proxy: Whether a proxy should be used for this VPN connection. This can be set to automatic, manual, or disabled.
When using Systems Manager Sentry VPN security, the username and password used to connect to the client VPN are generated by the Meraki cloud.
Usernames are generated based on a hash of a unique identifier on the device and the username of that device. Passwords are randomly generated.
Client VPN Connections
After configuring the client VPN and users starting to connect, it may be useful to see how many and which client devices are connected to your network. To see connected client VPN devices, navigate to Network-wide > Monitor > Clients. Click on the Search drop-down menu and select the following options: Status = Online, offline (or both), Client type = Client VPN.
Group Policies
It is possible to manually apply group policies to clients connected via client VPN. A group policy applied to a client VPN user is associated with the username and not the device. Different devices that connect to client VPN with the same username will receive the same group policy. For more help on assigning or removing group policies applied to a client, refer to the Creating and Applying Group Policies document.
It is not possible to assign group policies automatically once a user connects to client VPN.
Traceroute
The following is expected behavior: When you run a traceroute from a client who is connected to Client VPN to other destinations, one of the hops will appear as IP 192.168.100.1.
