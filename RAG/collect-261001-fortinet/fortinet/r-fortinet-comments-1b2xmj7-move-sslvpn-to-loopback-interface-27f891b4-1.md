---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4-1
title: "r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4.md
source_anchor: ""
source_lines: [1, 53]
sha256: 5820bbd3a43ad764480f9927cf8c8c0a29a14c4c2b51b8860f58db25fc259b41
---

# r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4

Move SSL-VPN to Loopback Interface
I want to move the SSL-VPN Service to a Loopback Interface, to be able to use the same functions like on a normal firewall policy, as the local-in-policies are too limited.
I have a /28 public IP range, but it is all used up, so I would like to use the same IP as the primary IP of the WAN interface and create a virtual IP pointing to the IP of the loopback interface.
The problem that I see that could cause issues is, that the primary WAN IP is used for the source NAT for the internet access of the users, and also for some IPsec tunnels. The users working with SSL-VPN don't have split-tunnel enabled and all the traffic is routed to our fortigate. They use the same IP as the primary WAN IP when they are connected with SSL-VPN to browse the internet.
I'm on 7.2.7 and have already seen that some functionality like using Internet Service Databases in local-in-policies was added in 7.4, but as that train is not mature, I don't want to use it.
Can this configuration work or do I have to think about a different solution?
Section des commentaires
I've got this set up in my home lab and at several customers.
You create a Virtual IP that NATs TCP 443 (or your preferred port) to the loopback interface. Create a firewall policy from WAN to the loopback interface, and set your preferred rules on it. Then you switch the interface in the SSL-VPN Settings, and it should work.
Great, thank you very much. So just to be sure, I use the SSL-VPN tunnel interface as source/destination interface in my firewall rules. As I would only change the listening interface on the SSL-VPN settings, this should not affect my firewall rules. I'll just have to create a new rule allowing access to the SSL-VPN loopback interface and that should be it?
I have this firewall policy set up:
edit 72
set srcintf "SD-WAN"
set dstintf "LB-SSL-VPN"
set action accept
set srcaddr "geo-Allowed Countries"
set dstaddr "vip-x.x.255.255-SSL-VPN" (VIP is from the dynamic IP on the wan1 interface to the loopback)
set schedule "always"
set service "HTTPS"
set logtraffic all
next
end
The LB-SSL-VPN is the loopback interface and the only interface selected in the SSL-VPN Settings.
To answer your direct question, yes this should be possible as long as port 443 is not in use elsewhere (for DNATs/virtual servers/etc.).
A VIP can change the mapped outbound NAT address so check this - if it does then use an IP Pool to set the specific public IP to use for the outbound SNAT policy.
We recently made this change as well.
Yes, you can use your primary external IP that is used for outbound SNAT for the inbound connection as well. The dynamic ports used by NAT are generally in the high range (by default) and typically won't veer downwards unless you are getting into SNAT port exhaustion.
Here's pretty much how we did it.
Interface Setup Create a loopback interface in the same VDOM as you want SSL VPN functionality. Specify an IP that is unique internally. The CIDR range can be small - we typically assign a /30 but you can do a /32, if you want. Do not enable any admin services on the interface.
Services Setup We set up a Services entry to represent the ports used by SSL VPN. If you're going to listen on TCP/443, include both TCP/443 and UDP/443. Essentially, include TCP and UDP for the same port number whichever port you choose. That allows DTLS to work for the SSL VPN connection.
Virtual IP Setup Create a new VIP. The external IP is the public IP address you want to use. The internal IP is the IP of the loopback interface. Enable the optional filters and enable Services. Select the service you just defined above. If you want to support HTTP to HTTPS redirection, also include the HTTP service.
SSL VPN Setup Under SSL-VPN Settings, set up to listen on the loopback interface using the port number you want to listen on. Enable that HTTP redirection if desired (totally optional). Other SSL-VPN settings are set how you need them (portal/tunnel/etc.).
Address Setup To protect your solution, add Address definitions for any geographies that you wish to exclude or include. Create an Address Group that groups together the geographies to be blocked and/or create an Address Group that groups together the geographies that you want to allow.
Firewall Policy Setup Create some rules to facilitate the traffic management:
Block Known Bad IPs Set the incoming interface to the external interface (or SD-WAN) that contains the external IP of the VIP. Set the outgoing interface to be the loopback interface. Set the Source to be selections from the Internet Services Database such as Malicious-Malicious.Server. Set the Destination to be the VIP you defined. Set the Service to the Service you created earlier (include HTTP if doing redirection). Set the action to DENY.
Block Geographies Set the incoming and outgoing interfaces as above. Set the Source to your address group containing geographies to be blocked. Set the Destination to the VIP. Set the Service to the service you created (and HTTP if doing redirection). Set the action to DENY.
Allow Access Set the incoming and outgoing interfaces as above. Set the Source to All and the destination to your VIP. Set the Service to the service you created (and HTTP if doing redirection). Set the action to ALLOW.
Optional Block all and only allow a geography Instead of the Block Geographies and Allow rule, create a rule that sets the Source to the Address Group you created of allowed geographies; everything else is the same as the Allow Access rule above. Your default deny rule should block everything else.
Once you have everything deployed, you can monitor the logs and look for sign in attempts from other systems that are acting against you. You can then block those as well using the above rules or new ones above the Allow Access rule in your policy set.
Good luck!
Thank you very much for the detailed answer. I will definitely consider your points when doing this move.
Thanks for the detailed walk-through! I tried this step by step and unfortunately, I cannot get it to work. I am using DHCP on my WAN1 interface, and it is the sole member of my SDWAN zone. So, in my VIP, I put 0.0.0.0, mapped it to my loopback with port forwarding on 10443 (which is not in use anywhere), selected the loopback in my SSL-VPN Settings (and deselecting my WAN1 interface) so it's listening, and created my firewall policy accordingly: SDWAN > loopback, All > VIP, service TCP-10443, allow, no NAT. I'm on 7.0.14. Not sure why it's not working! I have other DNATs with port forwarding (on other ports) working that are mapped to actual internal resources and those work fine!
I use ssl vpn settings to set allowed geo locations, is there any issues with that?, it works but I might be missing something which is why you prefer it done on the policies?
Configuring access at the SSL VPN itself is definitely an option for you. There's nothing wrong with that.
We've elected to do everything in the firewall policy rules so that it is all managed in one place and we can be a bit more flexible and predictable in how things work.
Have you tested it yet?
Nope, but I'm inclined to announce a maintenance window and test it one evening. I was hoping someone had already done something similar, before I try it and find out it doesn't work.
Yea that’s what I would do.
I have it working on my side with the loopback interface with full tunnel and there is no problem.
You should be fine.
I don't see a problem. Forward the correct ports to the VIP and you're good.
If you are using the same public IP for SSL VPN and IPSEC Tunnel and want to move this same IP to Loopback setup then you will have issue with IPSEC tunnel. Your IPSEC will drop regularly.
I tested it. I tried to use the same Public IP and map it loopback for SSL VPN and I started notification about IPSEC going down.
