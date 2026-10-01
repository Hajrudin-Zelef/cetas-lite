---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9-3
title: "c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9.md
source_anchor: ""
source_lines: [84, 148]
sha256: 91e72763d3e01e28d9ed37c26c0ae1c3c9be090148090800788a628db76785c3
---

# c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9

|  |  | Defines the ports in the group. Enter the command for each port or range of ports. For a list of permitted keywords and well-known port assignments, see the “Protocols and Applications” section. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
| Step 3 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
Example
To create service groups that include DNS (TCP/UDP), LDAP (TCP), and RADIUS (UDP), enter the following commands:
Configuring an ICMP Group
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Adds an ICMP type object group. The grp_id is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: The prompt changes to ICMP type configuration mode. | 
| Step 2 | Add one or more of the following group members: |  | 
|  |  | Defines the ICMP types in the group. Enter the command for each type. For a list of ICMP types, see the“ICMP Types” section. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
| Step 3 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
Example
Create an ICMP type group that includes echo-reply and echo (for controlling ping) by entering the following commands:
Configuring a Protocol Group
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Adds a protocol group. The obj_grp_id is a text string up to 64 characters in length and can be any combination of letters, digits, and the following characters: The prompt changes to protocol configuration mode. | 
| Step 2 | Add one or more of the following group members: |  | 
|  |  | Defines the protocols in the group. Enter the command for each protocol. The protocol is the numeric identifier of the specified IP protocol (1 to 254) or a keyword identifier (for example, icmp , tcp , or udp ). To include all IP protocols, use the keyword ip . For a list of protocols that you can specify, see the “Protocols and Applications” section. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
| Step 3 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
Example
To create a protocol group for TCP, UDP, and ICMP, enter the following commands:
Configuring Local User Groups
You can create local user groups for use in features that support the identity firewall (IDFW) by including the group in an extended ACL, which in turn can be used in an access rule, for example.
The ASA sends an LDAP query to the Active Directory server for user groups globally defined in the Active Directory domain controller. The ASA imports these groups for identity-based rules. However, the ASA might have localized network resources that are not defined globally that require local user groups with localized security policies. Local user groups can contain nested groups and user groups that are imported from Active Directory. The ASA consolidates local and Active Directory groups.
A user can belong to local user groups and user groups imported from Active Directory.
Prerequisites
See “Configuring the Identity Firewall,” to enable IDFW.
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 | object-group user user _ group _name hostname(config)# object-group user users1 | Defines object groups that you can use to control access with the Identity Firewall. | 
| Step 2 | Add one or more of the following group members: |  | 
|  | user domain_NetBIOS_name \ user_name hostname(config-user-object-group)# user SAMPLE\users1 | Specifies the user to add to the access rule. The user_name can contain any character including [a-z], [A-Z], [0-9], [!@#$%^&()-_{}. ]. If domain_NetBIOS_name \ user _ name contains a space, you must enclose the domain name and user name in quotation marks. The user _ name can be part of the LOCAL domain or a user imported by the ASA from Active Directory domain. If the domain_NetBIOS_name is associated with a AAA server, the user _ name must be the Active Directory sAMAccountName, which is unique, instead of the common name (cn), which might not be unique. The domain_NetBIOS_name can be LOCAL or the actual domain name as specified in user - identity domain domain_NetBIOS_name aaa - server aaa _ serve r_ group _ tag command. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
| Step 3 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
Configuring Security Group Object Groups
You can create security group object groups for use in features that support Cisco TrustSec by including the group in an extended ACL, which in turn can be used in an access rule, for example.
When integrated with Cisco TrustSec, the ASA downloads security group information from the ISE. The ISE acts as an identity repository, by providing Cisco TrustSec tag to user identity mapping and Cisco TrustSec tag to server resource mapping. You provision and manage security group ACLs centrally on the ISE.
However, the ASA might have localized network resources that are not defined globally that require local security groups with localized security policies. Local security groups can contain nested security groups that are downloaded from the ISE. The ASA consolidates local and central security groups.
To create local security groups on the ASA, you create a local security object group. A local security object group can contain one or more nested security object groups or Security IDs or security group names. User can also create a new Security ID or security group name that does not exist on the ASA.
You can use the security object groups you create on the ASA to control access to network resources. You can use the security object group as part of an access group or service policy.
Prerequisites
See “Configuring the ASA to Integrate with Cisco TrustSec,” to enable TrustSec.
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 | object-group security objgrp_name ciscoasa(config)# object-group security mktg-sg | Creates a security group object. Where objgrp_name is the name for the group entered as a 32-byte case sensitive string. The objgrp_name can contain any character including [a-z], [A-Z], [0-9], [!@#$%^&()-_{}. ]. | 
| Step 2 | Add one or more of the following group members: |  | 
|  | security-group { tag sgt# \| name sg_name } ciscoasa(config)# security-group name mktg | Specifies the type of security group object as either an inline tag or a named object.  An SGT is assigned to a device through IEEE 802.1X authentication, web authentication, or MAC authentication bypass (MAB) by the ISE. Security group names are created on the ISE and provide user-friendly names for security groups. The security group table maps SGTs to security group names. | 
|  |  | Adds an existing object group under this object group. The nested group must be of the same type. | 
| Step 3 |  | (Optional) Adds a description. The description can be up to 200 characters. | 
Examples
The following example shows how to configure a security group object:
Configuring Regular Expressions
Creating a Regular Expression
A regular expression matches text strings either literally as an exact string, or by using metacharacters so that you can match multiple variants of a text string. You can use a regular expression to match the content of certain application traffic; for example, you can match a URL string inside an HTTP packet.
Guidelines
Use Ctrl+V to escape all of the special characters in the CLI, such as question mark (?) or a tab. For example, type d[Ctrl+V]?g to enter d?g in the configuration.
See the regex command in the command reference for performance impact information when matching a regular expression to packets.
