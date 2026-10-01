---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6-3
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6.md
source_anchor: ""
source_lines: [56, 95]
sha256: bb2b4c91f0ab799394ab0d1f932b4f39dcf1f99f9007a8da3b53cf8c45de8b20
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6

An authentication rule (also known as “cut-through proxy”) controls network access based on the user. Because this function is very similar to an access rule plus an identity firewall, AAA rules can now be used as a backup method of authentication if a user AD login expires or a valid user has not yet logged into AD. For example, for any user without a valid login, you can trigger a AAA rule. To ensure that the AAA rule is only triggered for users that do not have valid logins, you can specify special usernames in the extended ACL that are used for the access rule and for the AAA rule: None (users without a valid login) and Any (users with a valid login). In the access rule, configure your policy as usual for users and groups, but then include a rule that permits all None users before deny any any; you must permit these users so they can later trigger a AAA rule. Then, configure a AAA rule that does not match Any users (these users are not subject to the AAA rule, and were handled already by the access rule), but matches all None users only to trigger AAA authentication for these users. After the user has successfully logged in via cut-through proxy, the traffic will flow normally again.
Static PAT and HTTP
For HTTP authentication, the ASA checks real ports when static PAT is configured. If it detects traffic destined for real port 80, regardless of the mapped port, the ASA intercepts the HTTP connection and enforces authentication.
For example, assume that outside TCP port 889 is translated to port 80 and that any relevant ACLs permit the traffic:
Then when users try to access 10.48.66.155 on port 889, the ASA intercepts the traffic and enforces HTTP authentication. Users see the HTTP authentication page in their web browsers before the ASA allows HTTP connection to complete.
If the local port is different than port 80, as in the following example:
Then users do not see the authentication page. Instead, the ASA sends an error message to the web browser, indicating that the user must be authenticated before using the requested service.
When a mapped address is used for static PAT, it is automatically placed into the dynamic PAT pool.
For instance, this configuration,
The second line ensures that all PAT bindings are accounted for.This accounting is necessary to avoid connection failure from port collision.
As the the mapped address is placed under dynamic PAT, any additional service that is to be accessed through the mapped address, must also be explicitly configured.
For example, the following is the correct configuration for three services through address 192.150.49.10. Additionally, the SMTP and HTTP services also reside at a host with the same address as the mapped address, 192.150.49.10.
Configuring Network Access Authentication
To configure network access authentication, perform the following steps:
|  |  |  | 
|---|---|---|
| Step 1 | ciscoasa(config)# aaa-server AuthOutbound protocol tacacs+ | Identifies your AAA servers. If you have already identified them, continue to the next step. | 
| Step 2 | ciscoasa(config)# access-list MAIL_AUTH extended permit tcp any any eq smtp | Creates an ACL that identifies the source addresses and destination addresses of traffic that you want to authenticate. The syntax shown here is just an example. For details, see the general operations configuration guide. If you specify identity firewall arguments in the ACL, then the following keywords in the ACL are specifically relevant to AAA rules. The keywords user - group any and user - group none can be specified to support cut-through proxy authentication. | 
| Step 3 | ciscoasa(config)# aaa authentication match MAIL_AUTH inside AuthOutbound | Configures authentication. The acl_name argument is the name of the ACL that you created in Step 2. The interface_name argument is the name of the interface specified with the nameif command. The server_group argument is the AAA server group that you created in Step 1. Note You can alternatively use the aaa authentication include command (which identifies traffic within the command). However, you cannot use both methods in the same configuration. See the command reference for more information. The user-identity keyword matches authentication to the identity firewall. | 
| Step 4 | ciscoasa(config)# aaa authentication listener http inside redirect | (Optional) Enables the redirection method of authentication for HTTP or HTTPS connections. The interface_name argument is the interface on which you want to enable listening ports. The port portnum argument specifies the port number on which the ASA listens; the defaults are 80 (HTTP) and 443 (HTTPS). You can use any port number and retain the same functionality, but be sure your direct authentication users know the port number; redirected traffic is sent to the correct port number automatically, but direct authenticators must specify the port number manually. Enter this command separately for HTTP and for HTTPS. | 
| Step 5 |  | (Optional) Uses the local database for network access authentication and limits the number of consecutive failed login attempts that the ASA allows any given user account (with the exception of users with a privilege level of 15. This feature does not affect level 15 users). The number argument value is between 1 and 16. | 
Examples
The following example authenticates all inside HTTP traffic and SMTP traffic:
The following example authenticates Telnet traffic from the outside interface to a particular server (209.165.201.5):
The following example shows a typical cut-through proxy configuration to allow a user to log in through the ASA. In this example, the following conditions apply:
- The ASA IP address is 192.168.123.10.
- The Active Directory domain controller has the IP address 10.1.2.10.
- The end user client has the IP address 192.168.123.10 and uses HTTPS to log in through a web portal.
- The user is authenticated by the Active Directory domain controller via LDAP.
- The ASA uses the inside interface to connect to the Active Directory domain controller on the corporate network.
In this example, the following guidelines apply:
- In access - list commands, you should configure permit user NONE rules before entering the access-list 100 ex deny any any command to allow unauthenticated incoming users to trigger AAA cut-through proxy.
- In access-list AUTH commands, permit user NONE rules specify that only unauthenticated users can trigger AAA cut-through proxy.
The following example shows how you can use AAA rules plus identity firewall (cut-through proxy) to authenticate successfully:
For more information about authentication, see the “Information About Authentication” section.
Enabling Secure Authentication of Web Clients
If you use HTTP authentication, by default the username and password are sent from the client to the ASA in clear text; in addition, the username and password are sent to the destination web server as well.
The ASA provides the following methods for securing HTTP authentication:
- Enable the redirection method of authentication for HTTP—Use the aaa authentication listener command with the redirect keyword. This method prevents the authentication credentials from continuing to the destination server. See the “ASA Authentication Prompts” section for more information about the redirection method compared to the basic method.
- Enable virtual HTTP—Use the virtual http command to authenticateseparately with the ASA and with the HTTP server. Even if the HTTP server does not need a second authentication, this command achieves the effect of stripping the basic authentication credentials from the HTTP GET request. See the “Authenticating HTTP(S) Connections with a Virtual Server” section for more information.
