---
id: collect-261001-cisco/cisco/enterprise-en-example-for-configuring-stelnet-login-based-on-radius-authenticati-26f523b5
title: "After the preceding configuration, you can log in to Server from Client through Telnet. You cannot log in to Server from other devices."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-example-for-configuring-stelnet-login-based-on-radius-authenticati-26f523b5.md
source_anchor: ""
source_lines: [1, 49]
sha256: 54791a765d066f260a1c2830c8f047a84476bac7fcd985ce37219b58a9709453
---

# After the preceding configuration, you can log in to Server from Client through Telnet. You cannot log in to Server from other devices.

Hello everyone,
Today I will share with you the example for configuring the device as the telnet client to log in to another device.
As shown in Figure 3-9, the PC and Client have reachable routes to each other; Client and Server have reachable routes to each other. The user needs to manage and maintain Server remotely. However, the PC cannot directly login to Server through Telnet because it has no reachable route to Server. The user can log in to Client through Telnet, and then log in to Server from Client. To prevent unauthorized devices from logging in to Server through Telnet, an ACL needs to be configured to allow only the Telnet connection from Client to Server.
Figure 3-9  Networking diagram of configuring the device as the Telnet client to log in to another device 
The Telnet protocol poses a security risk, and therefore the STelnet V2 protocol is recommended.
The configuration roadmap is as follows:
Configure the Telnet authentication mode on Server.
Configure the login user information on Server.
Configure an ACL on Server to allow Client access.
Log in to Server from Client through Telnet.
Configure the Telnet authentication mode and password on Server.
<HUAWEI> system-view 
[HUAWEI] sysname Server 
[Server] telnet server enable 
[Server] user-interface vty 0 4 
[Server-ui-vty0-4] user privilege level 15 
[Server-ui-vty0-4] protocol inbound telnet 
[Server-ui-vty0-4] authentication-mode aaa 
[Server-ui-vty0-4] quit
Configure the login user information.
[Server] aaa 
[Server-aaa] local-user admin1234 password irreversible-cipher Helloworld@6789 
[Server-aaa] local-user admin1234 service-type telnet 
[Server-aaa] local-user admin1234 privilege level 3 
[Server-aaa] quit
Configure an ACL on Switch2 to allow Client access.
[Server] acl 2000 
[Server-acl-basic-2000] rule permit source 10.1.1.1 0 
[Server-acl-basic-2000] quit 
[Server] user-interface vty 0 4 
[Server-ui-vty0-4] acl 2000 inbound 
[Server-ui-vty0-4] quit
NOTE:
It is optional to configure an ACL for Telnet services.
Verify the configuration.
# After the preceding configuration, you can log in to Server from Client through Telnet. You cannot log in to Server from other devices.
<HUAWEI> system-view 
[HUAWEI] sysname Client 
[Client] quit <Client> telnet 10.2.1.1 Trying 10.2.1.1 ... Press CTRL+K to abort Connected to 10.2.1.1 ... Warning: Telnet is not a secure protocol, and it is recommended to use STelnet. Login authentication Username:admin1234 Password: <Server>
Server configuration file
# sysname Server
 # telnet server enable 
 # acl number 2000  rule 5 permit source 10.1.1.1 0 
 # aaa  local-user admin1234 password irreversible-cipher $1a$gRNl~ukoL~0.WU)C2]~2a}Cz/Y0-u8M{j@Ql6/xHryO-Y7m{=A>kWc.-q}>*$ 
local-user admin1234 privilege level 3  local-user admin1234 service-type telnet 
 # user-interface vty 0 4  acl 2000 inbound  authentication-mode aaa  user privilege level 15  protocol inbound telnet  
 # return
See more please click
That is all I want to share with you! Thank you!
