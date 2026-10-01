---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-24
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1105, 1185]
sha256: 2bbc64df2c4b91e657a633b034374380affef96c7f3ea6a533889ec52f378c4d
---

# ms-meraki-campus-lan-5d88fe48

    When VLAN profiles are disabled you can still configure and assign profiles, but they won't take effect until you enable named VLAN profiles for the network (This also allows the feature to be temporarily removed from the switches and switch stacks without losing the existing configurations in dashboard)
- 
    It is recommended to create your own profiles otherwise any switch or stack that doesn't have a profile assigned will use the default profile (Removing a profile from a switch or stack will reapply the default profile automatically)
- 
    With MS15+, Named VLAN Profiles is supported on the following MS platforms: MS120, MS125, MS210, MS250, MS350, MS355, MS390, MS410
Named VLAN Profiles is not supported on MS420, MS425 and MS450
MS390 Specific Guidance
- MS390s supports Named VLAN profiles with MS15+
- When using multi-auth mode, MS390 has a different behavior; When multiple hosts authenticate to a single port on the MS390, each host may be assigned a unique VLAN to their session (e.g. the first host to authenticate on a switch port might be assigned to VLAN 3, and a subsequently authenticated host may be assigned to VLAN 5)
Access Control Lists (ACLs)
General Guidance
- 
    MS ACLs configured on Meraki switches are stateless (i.e. each packet is evaluated individually)
- 
    Remember to create rules that allow desired traffic in both directions where desired
- 
    All traffic traversing the switch (even non-routed traffic) will be evaluated
- 
    As traffic is evaluated in sequence down the list, it will only use the first rule that matches. Any traffic that doesn't match a specific allow or deny rule will be permitted by the default allow rule at the end of the list
- 
    Summarize IP addresses as much as possible to reduce ACL entries and improve overall performance
Configuration Guidelines for ACLs:
In a single rule, MS ACLs currently do not support following inputs:
- port ranges (e.g. '20000-30000')
- port lists (e.g. '80,443,3389')
- subnet lists (e.g. '192.168.1.0/24, 10.1.0.0/23')
- Review user and application traffic profiles and other permissible network traffic to determine the protocols and applications that should be granted access to the network.
- Please ensure traffic to the Meraki dashboard is permitted
- It may take 1-2 minutes for the changes to the ACL to propagate from the Meraki dashboard to the switches in your network
- MS platforms (except MS390) support a maximum of 128 access control entries (ACEs) per network
It is recommended to keep the number of ACLs for MS220-8P and MS220-24P platforms below 80
- To use IPv6 ACLs, please ensure to upgrade firmware to 10.0+
IPv6 ACL is not supported on MS220 and MS320
MS390 Specific Guidance
- You need to specify the IP address information (as opposed to just the VLAN ID like other MS platforms)
The VLAN qualifier is not supported on the MS390. For the MS390, ACL rules with non-empty VLAN fields will be ignored.
Group Policy Access Control Lists
General Guidance
- Group policies on MS switches allow users to define sets of Access Control Entries that can be applied to devices in order to control what they can access on the network.
- MS Group Policy ACLs can be applied to clients directly connected to an MS switch on access switchports
- This enables the application of the Layer 3 Firewall rules in a group policy on the MS switches within the network.
- When configuring this on dashboard, please note that the other configuration sections of the group policy will not apply to the MS switches, but will continue to be pushed to the devices in the network, such as the MX appliance and MR access-points, to which they are relevant.
- Only IP or CIDR based rules are supported. Groups containing rules using FQDNs are not be supported by MS switches
- Group Policy ACLs on MS are applied through client authentication Access Policies and, therefore, require a RADIUS server. Static assignment of a group to a client for Group Policy ACL application is not possible on MS switches.
- Access-Policy host-modes supported by Group Policy ACLs include single-host, multi-auth and multi-domain; Application of Group Policy ACL to a client authenticated by an access-policy using multi-host mode is not supported
- Do not use Group Policy ACLs for connecting on trunks as it will not be applied
- Group Policy ACLs on MS switches are implemented as stateless access control entries
- Also please note that Group Policy ACL rules will take precedence over Switch ACL rules (configured from the Switch > ACL section) on and only on the switch where the client has been authenticated
- Please refer to the below table for a compatibility matrix for Group Policy Access Control Lists:
| MS Switch Family | MS Switch Model | Minimum Firmware Required | 
| MS200 series | MS210 | MS 14.5 | 
|  | MS225 | MS 14.5 | 
|  | MS250 | MS 14.5 | 
| MS300 series | MS350 | MS 14.5 | 
|  | MS355 | MS 14.5 | 
|  | MS390 | MS 15.8 | 
| MS400 series | MS410 | MS 14.5 | 
|  | MS425 | MS 14.5 | 
|  | MS450 | MS 14.5 | 
Scaling Considerations for GP ACLs
Active Groups per switch = 20
Total number of active rules with layer-4 port ranges per switch = 32*
* The per-switch limit of 32 rules with layer-4 ports is shared between QoS and Group Policy ACL rules. However, while every QoS rule with a port range counts towards the limit, a Group Policy ACL rule with port range is counted only if a client device in that group is connected to the switch
GP ACLs are not supported on the following MS platforms: MS120, MS125, MS220, MS320
MS390 Specific Guidance
- 
    As of MS 15-8, MS390s support GP-ACLs and use the same Filter-Id attribute to process the policy as classic MS
- 
    If a valid Filter-Id is received from the RADIUS server during a client's authentication, the MS390 will apply the associated Group Policy ACL to the client's traffic regardless of the configuration explained in this section
- 
    For using Group Policy ACLs in networks where Access Policies are shared by MS390 and non-MS390 switches, please set RADIUS attribute specifying group policy name to Filter-Id
- 
    You need to specify the IP address information (as opposed to just the VLAN ID like other MS platforms)
Secure Connect
In relation to IP addressing, SecureConnect is a feature that is used to automate the process of securely provisioning Meraki MR Access Points when directly connected to switch-ports on Meraki MS Switches, without the requirement of a per-port configuration on the switch. With SecureConnect, connecting an MR access point to a switch-port on an MS switch triggers the switch-port to be configured to allow the MR to connect to the Meraki cloud and obtain a security certificate. The MR, subsequently, uses the certificate to identify itself at the switch-port via 802.1X and is allowed access to the network upon successful authentication.
General Guidance
- SecureConnect automates the process of securely provisioning Meraki MR Access Points when directly connected to switch-ports on Meraki MS Switches, without the requirement of a per-port configuration on the switch
- With SecureConnect, connecting an MR access point to a switch-port on an MS switch triggers the switch-port to be configured to allow the MR to connect to the Meraki cloud and obtain a security certificate
- The MR, subsequently, uses the certificate to identify itself at the switch-port via 802.1X and is allowed access to the network upon successful authentication
- For seamless operation of Secure Connect, it is recommended to have the same management VLAN configured for both the MR and the MS switch (i.e. Either configure it as the native VLAN on the switchport connecting to MR or change this manually on dashboard and ensure that the VLAN is also allowed on the trunk connecting the MR)
- For more information on the supported switch models and firmware, please refer to the following guide. Details provided in the below table:
