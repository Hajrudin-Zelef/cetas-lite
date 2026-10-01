---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-23
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [791, 828]
sha256: 0026719be665a9e461ee6423c7aaf8dd760862e376e1d1b0671c9de2a4b71976
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

The next tab in CAS configuration is the Filter tab (see Figure 54). For our example, the important dialog is Roles, where network traffic filters may be applied to different user Roles. The role of interest is the default Unauthenticated Role, which by default blocks all traffic. In this example, we are allowing the Unauthenticated Role to pass Active Directory client authentication traffic to the Active Directory Sever. This allows a Windows client to join the active Directory Domain and Windows users to authenticate to the domain, although they have not been through the NAC process. This is often important to allow printer and drive mapping information to be sent to Windows users. As the user has already authenticated to the Active Directory Domain, the user authentication information may be learned from Active Directory and the user does not have to reauthenticate for the NAC server.
Note The creation of Roles and their associated filters is performed in the CAM User Management -> User Roles menu.
Figure 54 CAS Filter Settings
The next tab that requires configuration is the Advanced tab, which has multiple dialogs that require configuration. The first of these is the Managed Subnet dialog, where each of the trusted VLAN subnets is added to the CAS for management. An example of this shown in Figure 55.
Figure 55 CAS Managed Subnet
The next dialog of the Advanced Tab is the VLAN Mapping dialog, which tells the CAS which trusted VLAN are mapped to an untrusted VLAN; an example of this is shown in Figure 56. In our example, VLAN Prunning and VLAN Mapping are also enabled.
Figure 56 VLAN Mapping
The next tab of interest is the Authentication Tab (see Figure 57), which has multiple dialogs for configuring different authentication options. The first dialog is the Login Page Dialog, which allows the configuration of different Web login pages depending upon the untrusted subnet being used for authenticating clients.
Figure 57 Authentication Login Page
The other Authentication dialog of interest is the Windows Auth dialog, as Windows Single Sign On (SSO) is used in this example. To perform Windows SSO, the CAS must be able to communicate with Active Directory to determine the authentication state of the Windows user. If Active Directory confirms that the user has authenticated to Active Directory, the user does not need to perform additional authentication to the CAS. An example of this configuration is shown in Figure 58. There are a number of steps required to configure Active Directory SSO, as these are described in the Cisco NAC Appliance -Clean Access Server Installation and Configuration Guide. The key components in this configuration are:
•The creation of a Active Directory client account for the CAS
•Using the KTPass Application on Active Directory to convert the account encryption to DES encryption
Figure 58 CAS Windows Authentication
Clean Access Roles
The unauthenticated role is common to all clients, but once the client has been authenticated, a different role may be applied based upon the identity of the client. Different roles may be assigned for admin staff, users, etc.
User roles allow you to aggregate various policies into a user role. These policies include:
Note If bandwidth policies are to be enforced by the Clean Access Server, it must be operating in band.
•Clean Access network port scanning plugins
•Clean Access Agent/Cisco NAC Web Agent client system requirements
For example, the Admin and User roles could each have different traffic polices and VLANs; in addition, the User role may enforce bandwidth policies by keeping the user traffic in band.
For more information on roles, refer to the Cisco NAC Appliance - Clean Access Manager Installation and Configuration Guide at: http://www.cisco.com/en/US/docs/security/nac/appliance/configuration_guide/45/cam/45cam-book.html.
Layer 2 OOB Example
Figure 59 shows an example of a Layer 2 OOB deployment where a wired client connected to an access switch is originally on the untrusted VLAN 264 and is switched to a trusted VLAN 64 once it has completed the NAC functions. The first NAC function is the authentication and authorization function, which is the first design decision in implementing the NAC solution, i.e, how will authentication and authorization be achieved and what will the user experience be.
This example is focused on the virtual gateway example, as a virtual gateway provides the simplest deployment. In the virtual gateway example, the original IP addressing, interfaces, VLANs, and normal traffic flows are maintained. The only changes are the addition of the untrusted VLANs that carry client traffic during NAC Authentication and Authorization, Scanning and evaluation, remediation, and quarantine modes.
Figure 59 Layer 2 OOB Example
NAC Authentication Options
The authentication option in the NAC solution can be broadly categorized as NAC Authentication or NAC Single Sign On
•NAC Authentication—NAC authentication gives the NAC system the role of authenticating users against a user database, either local to the NAC system or a separate system such as RADIUS or LDAP.
•NAC Single Sign On—NAC SSO address systems perform authentication as part of their normal operation (e.g., 802.1X, VPN access, or Active Directory). NAC SSO learns the authentication state of clients through RADIUS accounting or Active Director and therefore does not require the user to reenter authentication.
Topology Considerations
The Layer 2 OOB solution relies upon there being a Layer 2 network connection available between the client devices and the Cisco CAS. In Figure 59, a trunk connects the access switch to the core/distribution switch. The Cisco CAS is connected to the core/distribution switch through two interfaces—trusted and untrusted. In such a simple network, it is relatively easy to provide a Layer 2 connection between the client and the Cisco CAS; for larger networks, Layer 3 OOB may be a better choice.
The roles of the untrusted and trusted interfaces:
•Untrusted interface—The untrusted interface connects the client to the to the Cisco CAS during the NAC Authentication and Authorization, Scanning and Evaluation, Remediation, and Quarantine modes
•Trusted interface—The trusted interface connects the NAC CAS to the "normal" network interface. This makes a network connection available to the CAS while it is sitting between the client and the network, thus allowing client access to services such as DHCP and DNS—and user-defined services. Once a client has successfully completed its authentication and scanning phases, the CAM uses SNMP to change the client VLAN, on the access switch, from the untrusted VLAN to the trusted VLAN, thus providing a direct connection to the network that was on the other side of the CAS (the trusted network).
Availability Considerations
The CAS and CAM are both highly involved in client network access and consideration must be given to the impact on clients if either a CAS or CAM should fail or need to be taken out of service.
The CAS is inline with client devices during the authentication, authorization, and posture assessment phases of NAC and, if "In Band NAC" is being used, it may be inline at all times. A CAS outage in an OOB deployment would not impact already-connected clients, but would prevent network access for new clients. A CAS outage for "In Line" clients prevents access for all clients.
In situations where availability of a CAS is critical, a high availability CAS solution may be implemented, where a pair of CAS servers are installed using a primary CAS and a secondary in hot standby. For more information, refer to the Cisco NAC Appliance - Clean Access Server Installation and Configuration Guide.
