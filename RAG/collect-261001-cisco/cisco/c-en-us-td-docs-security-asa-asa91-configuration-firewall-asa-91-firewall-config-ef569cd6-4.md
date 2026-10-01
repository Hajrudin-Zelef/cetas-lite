---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6-4
title: "c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6.md
source_anchor: ""
source_lines: [96, 124]
sha256: e19a2c13ee085d4b85ee94cffb1000c3d9e64beed2ffe0300397a3dd4f5d1aad
---

# c-en-us-td-docs-security-asa-asa91-configuration-firewall-asa-91-firewall-config-ef569cd6

Enable the exchange of usernames and passwords between a web client and the ASA with HTTPS—Use the aaa authentication secure-http-client command to enable the exchange of usernames and passwords between a web client and the ASA with HTTPS. This is the only method that protects credentials between the client and the ASA, as well as between the ASA and the destination server. You can use this method alone, or in conjunction with either of the other methods so you can maximize your security.
After enabling this feature, when a user requires authentication when using HTTP, the ASA redirects the HTTP user to an HTTPS prompt. After you authenticate correctly, the ASA redirects you to the original HTTP URL.
Secured, web-client authentication has the following limitations:
– A maximum of 64 concurrent HTTPS authentication sessions are allowed. If all 64 HTTPS authentication processes are running, a new connection requiring authentication will not succeed.
– When uauth timeout 0 is configured (the uauth timeout is set to 0),HTTPS authentication might not work. If a browser initiates multiple TCP connections to load a web page after HTTPS authentication, the first connection is let through, but the subsequent connections trigger authentication. As a result, users are continuously presented with an authentication page, even if the correct username and password are entered each time. To work around this, set the uauth timeout to 1 second with the timeout uauth 0:0:1 command. However, this workaround opens a 1-second window of opportunity that might allow unauthenticated users to go through the firewall if they are coming from the same source IP address.
Because HTTPS authentication occurs on the SSL port 443, users must not configure an access-list command statement to block traffic from the HTTP client to the HTTP server on port 443. Furthermore, if static PAT is configured for web traffic on port 80, it must also be configured for the SSL port.
– In the following example, the first set of commands configures static PAT for web traffic, and the second set of commands must be added to support the HTTPS authentication configuration:
Authenticating Directly with the ASA
If you do not want to allow HTTP, HTTPS, Telnet, or FTP through the ASA but want to authenticate other types of traffic, you can authenticate with the ASA directly using HTTP, HTTPS, or Telnet.
Authenticating HTTP(S) Connections with a Virtual Server
If you enabled the redirection method of HTTP and HTTPS authentication in the “Configuring Network Access Authentication” section, then you have also automatically enabled direct authentication.
When you use HTTP authentication on the ASA (see the“Configuring Network Access Authentication” section), the ASA uses basic HTTP authentication by default.
To continue to use basic HTTP authentication, and to enable direct authentication for HTTP and HTTPS, enter the following command:
|  |  | 
|---|---|
| ciscoasa(config)# aaa authentication listener http inside redirect | (Optional) Enables the redirection method of authentication for HTTP or HTTPS connections. The interface_name argument is the interface on which you want to enable listening ports. The port portnum argument specifies the port number on which the ASA listens; the defaults are 80 (HTTP) and 443 (HTTPS). You can use any port number and retain the same functionality, but be sure your direct authentication users know the port number; redirected traffic is sent to the correct port number automatically, but direct authenticators must specify the port number manually. Enter this command separately for HTTP and for HTTPS. | 
If the destination HTTP server requires authentication in addition to the ASA, then to authenticate separately with the ASA (via a AAA server) and with the HTTP server, enter the following command:
|  |  | 
|---|---|
| ciscoasa(config)# virtual http | Redirects all HTTP connections that require AAA authentication to the virtual HTTP server on the ASA. The ASA prompts for the AAA server username and password. After the AAA server authenticates the user, the ASA redirects the HTTP connection back to the original server, but it does not include the AAA server username and password. Because the username and password are not included in the HTTP packet, the HTTP server prompts the user separately for the HTTP server username and password. For inbound users (from lower security to higher security), you must also include the virtual HTTP address as a destination interface in the ACL applied to the source interface. In addition, you must add a static NAT command for the virtual HTTP IP address, even if NAT is not required. An identity NAT command is typically used (where you translate the address to itself). For outbound users, there is an explicit permit for traffic, but if you apply an ACL to an inside interface, be sure to allow access to the virtual HTTP address. A static statement is not required. Note Do not set the timeout uauth command duration to 0 seconds when using the virtual http command, because this setting prevents HTTP connections to the actual web server. You can authenticate directly with the ASA at the following URLs when you enable AAA for the interface: Without virtual HTTP, the same username and password that you used to authenticate with the ASA are sent to the HTTP server; you are not prompted separately for the HTTP server username and password. Assuming the username and password are not the same for the AAA and HTTP servers, then the HTTP authentication fails. | 
Authenticating Telnet Connections with a Virtual Server
Although you can configure network access authentication for any protocol or service (see the aaa authentication match or aaa authentication include command), you can authenticate directly with HTTP, Telnet, or FTP only. A user must first authenticate with one of these services before other traffic that requires authentication is allowed through. If you do not want to allow HTTP, Telnet, or FTP traffic through the ASA, but want to authenticate other types of traffic, you can configure virtual Telnet; the user Telnets to a given IP address configured on the ASA, and the ASA issues a Telnet prompt.
To configure a virtual Telnet server, enter the following command:
|  |  | 
|---|---|
| ciscoasa(config)# virtual telnet 209.165.202.129 | Configures a virtual Telnet server. The ip_address argument sets the IP address for the virtual Telnet server. Make sure this address is an unused address that is routed to the ASA. You must configure authentication for Telnet access to the virtual Telnet address as well as the other services that you want to authenticate using the authentication match or aaa authentication include command. When an unauthenticated user connects to the virtual Telnet IP address, the user is challenged for a username and password, and then authenticated by the AAA server. Once authenticated, the user sees the message “Authentication Successful.” Then, the user can successfully access other services that require authentication. For inbound users (from lower security to higher security), you must also include the virtual Telnet address as a destination interface in the ACL applied to the source interface. In addition, you must add a static NAT command for the virtual Telnet IP address, even if NAT is not required. An identity NAT command is typically used (where you translate the address to itself). For outbound users, there is an explicit permit for traffic, but if you apply an ACL to an inside interface, be sure to allow access to the virtual Telnet address. A static statement is not required. To log out from the ASA, reconnect to the virtual Telnet IP address; you are then prompted to log out. | 
Examples
The following example shows how to enable virtual Telnet together with AAA authentication for other services:
Configuring Authorization for Network Access
