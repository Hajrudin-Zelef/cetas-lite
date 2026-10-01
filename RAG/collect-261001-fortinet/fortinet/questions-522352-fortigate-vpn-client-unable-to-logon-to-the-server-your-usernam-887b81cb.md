---
id: collect-261001-fortinet/fortinet/questions-522352-fortigate-vpn-client-unable-to-logon-to-the-server-your-usernam-887b81cb
title: "questions-522352-fortigate-vpn-client-unable-to-logon-to-the-server-your-usernam-887b81cb"
domain: fortinet
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-522352-fortigate-vpn-client-unable-to-logon-to-the-server-your-usernam-887b81cb.md
source_anchor: ""
source_lines: [1, 23]
sha256: 8c87b744303a9dd80b96fd3a0a5f416b5a6f5d2e5d916578224f047d3aaaf6a8
---

# questions-522352-fortigate-vpn-client-unable-to-logon-to-the-server-your-usernam-887b81cb

fortinet - Fortigate VPN client "Unable to logon to the server. Your username or password may not be properly configured for this connection. (-12)" - Server Fault
Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We're running a Fortigate 100D, and having some trouble with the SSL VPN via FortiClient.
I installed FortiClient on an external Windows 7 PC a few days pack and the SSL VPN connected and worked.
I uninstalled it from that PC and installed it on a different external Windows 7 PC, and now cannot connect to the VPN. I get this error when it attempts to connect:
Unable to logon to the server. Your username or password may not be properly configured for this connection. (-12)
Google was pretty unhelpful.
We have tried it on two external Windows systems now, and get the same error on both systems.
There does not appear to be any information in the firewall's event logs/vpn logs.
Not sure what is preventing the VPN from connecting.
Edit: We have reset the password for the user - and are 100% sure that we have a correct username and password. In fact it is happening with two different accounts, both of which worked previously.
Discovered that the problem was that I had special characters in my password. There was never any indication that special characters were not permitted, but sure enough, when I reset the password to something alphanumeric, it works.
Also... if you do not have the Tunnel Mode allowed in the SSL Portal configuration for that particular Portal. As soon as I added it in I could connect with the full SSL VPN client.
As the error states itself the most common problem is that either the username or the password isn't matching the one of the device.
Other problems might be:
the user is not in the correct user group that has VPN access (either the local firewall group or the LDAP server group if you’re using one)
there isn't a corresponding firewall policy rule that allows access for the user group to any of the internal networks. You need to have the rule from the wan interface to one of the internal interfaces with action SSL-VPN and select the group of users which will have access, check if your user is in correct group.
you might be trying to connect to VPN from the wrong side of the interface (from one of your internal networks or from the network of one of the sites you already have a site to site connection
