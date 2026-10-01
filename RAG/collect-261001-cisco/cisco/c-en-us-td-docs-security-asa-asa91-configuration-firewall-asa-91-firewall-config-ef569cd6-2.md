---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6-2
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6.md
source_anchor: ""
source_lines: [7, 55]
sha256: a81b3a5675d7a809f414d833b4d6e9ca35a106f9ff6c5c308bcd1740cd1b57f5
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6

  This chapter describes how to enable AAA (pronounced “triple A”) for network access.
For information about AAA for management access, see the general operations configuration guide.
AAA Performance
The ASA uses “cut-through proxy” to significantly improve performance compared to a traditional proxy server. The performance of a traditional proxy server suffers because it analyzes every packet at the application layer of the OSI model. The ASA cut-through proxy challenges a user initially at the application layer and then authenticates with standard AAA servers or the local database. After the ASA authenticates the user, it shifts the session flow, and all traffic flows directly and quickly between the source and destination while maintaining session state information.
Licensing Requirements for AAA Rules
The following table shows the licensing requirements for this feature:
Guidelines and Limitations
This section includes the guidelines and limitations for this feature.
Supported in single and multiple context mode.
Supported in routed and transparent firewall mode.
In clustering, this feature is only supported on the master unit.
Configuring Authentication for Network Access
This section includes the following topics:
Information About Authentication
The ASA lets you configure network access authentication using AAA servers. This section includes the following topics:
One-Time Authentication
A user at a given IP address only needs to authenticate one time for all rules and types, until the authentication session expires. (See the timeout uauth command in the command reference for timeout values.) For example, if you configure the ASA to authenticate Telnet and FTP, and a user first successfully authenticates for Telnet, then as long as the authentication session exists, the user does not also have to authenticate for FTP.
Applications Required to Receive an Authentication Challenge
Although you can configure the ASA to require authentication for network access to any protocol or service, users can authenticate directly with HTTP, HTTPS, Telnet, or FTP only. A user must first authenticate with one of these services before the ASA allows other traffic requiring authentication.
The authentication ports that the ASA supports for AAA are fixed as follows:
ASA Authentication Prompts
For Telnet and FTP, the ASA generates an authentication prompt.
For HTTP, the ASA uses basic HTTP authentication by default, and provides an authentication prompt. You can optionally configure the ASA to redirect users to an internal web page where they can enter their username and password (configured with the aaa authentication listener command).
For HTTPS, the ASA generates a custom login screen. You can optionally configure the ASA to redirect users to an internal web page where they can enter their username and password (configured with the aaa authentication listener command).
Redirection is an improvement over the basic method because it provides an improved user experience during authentication, and an identical user experience for HTTP and HTTPS in both Easy VPN and firewall modes. It also supports authentication directly with the ASA.
You might want to continue to use basic HTTP authentication for the following reasons:
- You do not want the ASA to open listening ports.
- You use NAT on a router and you do not want to create a translation rule for the web page served by the ASA.
- Basic HTTP authentication might work better with your network.
For example non-browser applications, as when a URL is embedded in e-mail, might be more compatible with basic authentication.
After you authenticate correctly, the ASA redirects you to your original destination. If the destination server also has its own authentication, the user enters another username and password. If you use basic HTTP authentication and need to enter another username and password for the destination server, then you need to configure the virtual http command.
Note If you use HTTP authentication, by default the username and password are sent from the client to the ASA in clear text; in addition, the username and password are sent on to the destination web server as well. See the “Enabling Secure Authentication of Web Clients” section for information to secure your credentials.
For FTP, a user has the option of entering the ASA username followed by an at sign (@) and then the FTP username (name1@name2). For the password, the user enters the ASA password followed by an at sign (@) and then the FTP password (password1@password2). For example, enter the following text:
This feature is useful when you have cascaded firewalls that require multiple logins. You can separate several names and passwords by multiple at signs (@).
AAA Prompts and Identity Firewall
In an enterprise, some users log into the network by using other authentication mechanisms, such as authenticating with a web portal (cut-through proxy). For example, users with a Mac and Linux client might log into a web portal (cut-through proxy). Therefore, you must configure the identity firewall to allow these types of authentication in connection with identity-based access policies.
Figure 7-1 shows a deployment to support a cut-through proxy authentication captive portal. Active Directory servers and the AD Agent are installed on the main site LAN. However, the identity firewall is configured to support authentication of clients that are not part of the Active Directory domain.
Figure 7-1 Deployment Supporting Cut-through Proxy Authentication
The ASA designates users logging in through a web portal (cut-through proxy) as belonging to the Active Directory domain with which they authenticated.
The ASA reports users logging in through a web portal (cut-through proxy) to the AD Agent, which distributes the user information to all registered ASA devices. In this case, the identity firewall can associate the users with their Active Directory domain. Specifically, the user identity-IP address mappings of authenticated users are forwarded to all ASA contexts that contain the input interface where packets are received and authenticated.
Users can log in by using HTTP/HTTPS, FTP, Telnet, or SSH. When users log in with these authentication methods, the following guidelines apply:
- For HTTP/HTTPS traffic, an authentication window appears for unauthenticated users.
- For Telnet and FTP traffic, users must log in through the cut-through proxy server and again to the Telnet and FTP servers.
- A user can specify an Active Directory domain while providing login credentials (in the format, domain\username). The ASA automatically selects the associated AAA server group for the specified domain.
- If a user specifies an Active Directory domain while providing login credentials (in the format, domain\username), the ASA parses the domain and uses it to select an authentication server from the AAA servers that have been configured for the identity firewall. Only the username is passed to the AAA server.
- If the backslash (\) delimiter is not found in the login credentials, the ASA does not parse the domain and authentication is conducted with the AAA server that corresponds to the default domain configured for the identity firewall.
- If a default domain or a server group is not configured for that default domain, the ASA rejects the authentication.
- If the domain is not specified, the ASA selects the AAA server group for the default domain that is configured for the identity firewall.
AAA Rules as a Backup Authentication Method
