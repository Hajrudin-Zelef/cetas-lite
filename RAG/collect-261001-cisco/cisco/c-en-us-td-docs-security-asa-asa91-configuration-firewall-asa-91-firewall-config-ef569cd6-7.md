---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6-7
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6.md
source_anchor: ""
source_lines: [207, 233]
sha256: f219d5be9f9f549e6dfbcbf24f7a95f1b83fdd5480790d59baa364928ffe6f01
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6

The ASA can send accounting information to a RADIUS or TACACS+ server about any TCP or UDP traffic that passes through the ASA. If that traffic is also authenticated, then the AAA server can maintain accounting information by username. If the traffic is not authenticated, the AAA server can maintain accounting information by IP address. Accounting information includes session start and stop times, username, the number of bytes that pass through the ASA for the session, the service used, and the duration of each session.
To configure accounting, perform the following steps:
|  |  |  | 
|---|---|---|
| Step 1 | ciscoasa(config)# access-list TELNET_AUTH extended permit tcp any any eq telnet | If you want the ASA to provide accounting data per user, you must enable authentication. For more information, see the “Configuring Network Access Authentication” section. If you want the ASA to provide accounting data per IP address, enabling authentication is not necessary. Creates an ACL that identifies the source addresses and destination addresses of traffic for which you want accounting data. For instructions, see the general operations configuration guide. The permit ACEs mark matching traffic for accounting, while deny entries exclude matching traffic from accounting. Note If you have configured authentication and want accounting data for all the traffic being authenticated, you can use the same ACL that you created for use with the aaa authentication match command. | 
| Step 2 |  | Enables accounting. The acl_name argument is the ACL name set in the access-list command. The interface_name argument is the interface name set in the nameif command. The server_group argument is the server group name set in the aaa-server command. Note Alternatively, you can use the aaa accounting include command (which identifies traffic within the command), but you cannot use both methods in the same configuration. See the command reference for more information. | 
Examples
The following example authenticates, authorizes, and accounts for inside Telnet traffic. Telnet traffic to servers other than 209.165.201.5 can be authenticated alone, but traffic to 209.165.201.5 requires authorization and accounting.
AAA provides an extra level of protection and control for user access than using ACLs alone. For example, you can create an ACL allowing all outside users to access Telnet on a server on the DMZ network. If you want only some users to access the server and you might not always know IP addresses of these users, you can enable AAA to allow only authenticated and/or authorized users to connect through the ASA. (The Telnet server enforces authentication, too; the ASA prevents unauthorized users from attempting to access the server.)
Using MAC Addresses to Exempt Traffic from Authentication and Authorization
The ASA can exempt from authentication and authorization any traffic from specific MAC addresses. For example, if the ASA authenticates TCP traffic originating on a particular network, but you want to allow unauthenticated TCP connections from a specific server, you would use a MAC exempt rule to exempt from authentication and authorization any traffic from the server specified by the rule.
This feature is particularly useful to exempt devices such as IP phones that cannot respond to authentication prompts.
To use MAC addresses to exempt traffic from authentication and authorization, perform the following steps:
|  |  |  | 
|---|---|---|
| Step 1 | ciscoasa(config)# mac-list abc permit 00a0.c95d.0282 ffff.ffff.ffff | Configures a MAC list. The id argument is the hexadecimal number that you assign to the MAC list. To group a set of MAC addresses, enter the mac-list command as many times as needed with the same ID value. Because you can only use one MAC list for AAA exemption, be sure that your MAC list includes all the MAC addresses that you want to exempt. You can create multiple MAC lists, but you can only use one at a time. The order of entries matters, because the packet uses the first entry it matches, instead of a best match scenario. If you have a permit entry, and you want to deny an address that is allowed by the permit entry, be sure to enter the deny entry before the permit entry. The mac argument specifies the source MAC address in 12-digit hexadecimal form; that is, nnnn.nnnn.nnnn. The macmask argument specifies the portion of the MAC address that should be used for matching. For example, ffff.ffff.ffff matches the MAC address exactly. ffff.ffff.0000 matches only the first 8 digits. | 
| Step 2 |  | Exempts traffic for the MAC addresses specified in a particular MAC list. The id argument is the string identifying the MAC list that includes the MAC addresses whose traffic is to be exempt from authentication and authorization. You can only enter one instance of the aaa mac-exempt match command. | 
Examples
The following example bypasses authentication for a single MAC address:
The following example bypasses authentication for all Cisco IP Phones, which have the hardware ID 0003.E3:
The following example bypasses authentication for a a group of MAC addresses except for 00a0.c95d.02b2. Enter the deny statement before the permit statement, because 00a0.c95d.02b2 matches the permit statement as well, and if it is first, the deny statement will never be matched.
Feature History for AAA Rules
Table 7-1 lists each feature change and the platform release in which it was implemented.
|  |  |  | 
|---|---|---|
| AAA Rules | 7.0(1) | AAA Rules describe how to enable AAA for network access. We introduced the following commands: aaa authentication match, aaa authentication include \| exclude, aaa authentication listener http[s], aaa local authentication attempts max-fail, virtual http, virtual telnet, aaa authentication secure-http-client, aaa authorization match, aaa accounting match, aaa mac-exempt match. | 
| Authentication using Cut-Through Proxy | 9.0(1) | You can authenticate using AAA rules in conjunction with the Identity Firewall feature. We modified the following command: aaa authentication match. |
