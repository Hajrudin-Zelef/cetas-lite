---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6-5
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6.md
source_anchor: ""
source_lines: [125, 165]
sha256: cb6a3eb738450d4ad6d9608d0398ccc7ab3d1f666c0561e75eae32188300fecc
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6

After a user authenticates for a given connection, the ASA can use authorization to further control traffic from the user.
This section includes the following topics:
Configuring TACACS+ Authorization
You can configure the ASA to perform network access authorization with TACACS+. You identify the traffic to be authorized by specifying ACLs that authorization rules must match. Alternatively, you can identify the traffic directly in authorization rules themselves.
Tip Using ACLs to identify traffic to be authorized can greatly reduced the number of authorization commands that you must enter. This is because each authorization rule that you enter can specify only one source and destination subnet and service, whereas an ACL can include many entries.
Authentication and authorization statements are independent; however, any unauthenticated traffic matched by an authorization rule will be denied. For authorization to succeed:
1. A user must first authenticate with the ASA.
Because a user at a given IP address only needs to authenticate one time for all rules and types, if the authentication session has not expired, authorization can occur even if the traffic is not matched by an authentication rule.
2. After a user authenticates, the ASA checks the authorization rules for matching traffic.
3. If the traffic matches the authorization rule, the ASA sends the username to the TACACS+ server.
4. The TACACS+ server responds to the ASA with a permit or a deny for that traffic, based on the user profile.
5. The ASA enforces the authorization rule in the response.
See the documentation for your TACACS+ server for information about configuring network access authorizations for a user.
To configure TACACS+ authorization, perform the following steps:
|  |  |  | 
|---|---|---|
| Step 1 | ciscoasa(config)# aaa-server AuthOutbound protocol tacacs+ | Identifies your AAA servers. If you have already identified them, continue to the next step. | 
| Step 2 | ciscoasa(config)# access-list MAIL_AUTH extended permit tcp any any eq smtp | Creates an ACL that identifies the source addresses and destination addresses of traffic you want to authenticate. For details, see the general operations configuration guide. The permit ACEs mark matching traffic for authentication, while deny entries exclude matching traffic from authentication. Be sure to include the destination ports for either HTTP, HTTPS, Telnet, or FTP in the ACL, because the user must authenticate with one of these services before other services are allowed through the ASA. | 
| Step 3 | ciscoasa(config)# aaa authentication match MAIL_AUTH inside AuthOutbound | Configures authentication. The acl_name argument is the name of the ACL that you created in Step 2., The interface_name argument is the name of the interface specified with the nameif command, and the server_group argument is the AAA server group that you created in Step 1. Note You can alternatively use the aaa authentication include command (which identifies traffic within the command). However, you cannot use both methods in the same configuration. See the command reference for more information. | 
| Step 4 | ciscoasa(config)# aaa authentication listener http inside redirect | (Optional) Enables the redirection method of authentication for HTTP or HTTPS connections. The interface_name argument is the interface on which you want to enable listening ports. The port portnum argument specifies the port number on which the ASA listens; the defaults are 80 (HTTP) and 443 (HTTPS). You can use any port number and retain the same functionality, but be sure your direct authentication users know the port number; redirected traffic is sent to the correct port number automatically, but direct authenticators must specify the port number manually. Enter this command separately for HTTP and for HTTPS. | 
| Step 5 |  | (Optional) Uses the local database for network access authentication and limits the number of consecutive failed login attempts that the ASA allows any given user account (with the exception of users with a privilege level of 15. This feature does not affect level 15 users). The number argument value is between 1 and 16. | 
| Step 6 |  | Create an ACL that identifies the source addresses and destination addresses of traffic that you want to authorize. For instructions, see the general operations configuration guide. The permit ACEs mark matching traffic for authorization, while deny entries exclude matching traffic from authorization. The ACL that you use for authorization matching should include rules that are equal to or a subset of the rules in the ACL used for authentication matching. Note If you have configured authentication and want to authorize all the traffic being authenticated, you can use the same ACL that you created for use with the aaa authentication match command. | 
| Step 7 |  | Enables authorization. The acl_name argument is the name of the ACL you created in Step 6, the interface_name argument is the name of the interface as specified with the nameif command or by default, and the server_group argument is the AAA server group that you created when you enabled authentication. Note Alternatively, you can use the aaa authorization include command (which identifies traffic within the command) but you cannot use both methods in the same configuration. See the command reference for more information. | 
Examples
The following example authenticates and authorizes inside Telnet traffic. Telnet traffic to servers other than 209.165.201.5 can be authenticated alone, but traffic to 209.165.201.5 requires authorization.
Configuring RADIUS Authorization
When authentication succeeds, the RADIUS protocol returns user authorizations in the access-accept message sent by a RADIUS server. For more information about configuring authentication, see the “Configuring Network Access Authentication” section.
When you configure the ASA to authenticate users for network access, you are also implicitly enabling RADIUS authorizations; therefore, this section contains no information about configuring RADIUS authorization on the ASA. It does provide information about how the ASA handles ACL information received from RADIUS servers.
You can configure a RADIUS server to download an ACL to the ASA or an ACL name at the time of authentication. The user is authorized to do only what is permitted in the user-specific ACL.
Note If you have used the access-group command to apply ACLs to interfaces, be aware of the following effects of the per-user-override keyword on authorization by user-specific ACLs:
- Without the per-user-override keyword, traffic for a user session must be permitted by both the interface ACL and the user-specific ACL.
- With the per-user-override keyword, the user-specific ACL determines what is permitted.
For more information, see the access-group command entry in the command reference.
Configuring a RADIUS Server to Send Downloadable Access Control Lists
This section describes how to configure Cisco Secure ACS or a third-party RADIUS server and includes the following topics:
About the Downloadable ACL Feature and Cisco Secure ACS
Downloadable ACLs is the most scalable means of using Cisco Secure ACS to provide the appropriate ACLs for each user. It provides the following capabilities:
- Unlimited ACL size—Downloadable ACLs are sent using as many RADIUS packets as required to transport the full ACL from Cisco Secure ACS to the ASA.
- Simplified and centralized management of ACLs—Downloadable ACLs enable you to write a set of ACLs once and apply it to many user or group profiles and distribute it to many ASAs.
This approach is most useful when you have very large ACL sets that you want to apply to more than one Cisco Secure ACS user or group; however, its ability to simplify Cisco Secure ACS user and group management makes it useful for ACLs of any size.
The ASA receives downloadable ACLs from Cisco Secure ACS using the following process:
