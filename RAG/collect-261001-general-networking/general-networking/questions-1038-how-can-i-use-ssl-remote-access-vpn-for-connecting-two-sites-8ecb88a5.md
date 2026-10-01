---
id: collect-261001-general-networking/general-networking/questions-1038-how-can-i-use-ssl-remote-access-vpn-for-connecting-two-sites-8ecb88a5
title: "questions-1038-how-can-i-use-ssl-remote-access-vpn-for-connecting-two-sites-8ecb88a5"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "reasoning", "research"]
source: docs/RAG/collect-261001-general-networking/questions-1038-how-can-i-use-ssl-remote-access-vpn-for-connecting-two-sites-8ecb88a5.md
source_anchor: ""
source_lines: [1, 37]
sha256: e91ad83dbaefa7dbe952b74b183e53175fef7f63d4e1d2ffe9d2136fe9646cae
---

# questions-1038-how-can-i-use-ssl-remote-access-vpn-for-connecting-two-sites-8ecb88a5

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Fortigate 110C on my central site. I have a remote site office too. The users at the remote site need to access the central site via a VPN. Can I run SSL client software on one of the PC's or any device on the remote site, and redirect all the VPN traffic via it?
The purpose is to avoid both:
Placing a new VPN concentrator or firewall on the remote site
We don't want each user to install an SSL client on their PC at the remote office and to dial separately.
You can do this in theory, but you will need a good client machine to do that and by good, I mean a good Windows installation.
This is because the operating system on the machine you want to use SSL VPN client will have to deal with all the traffic, and that machine will have to somehow prove router&firewall capabilities.
Step 1: you connect that machine (from remote office) to the headquarter. You will receive an IP address from the SSL_VPN_pool.
Step 2: you will add a static (persistent route) on all stations (from the remote office), that for the HQ destination would have to reach through the machine connected at Step 1.
Step 3: you will have to enable Routing&Remote Service on the machine you use SSL VPN client (I assume you will use a Windows platform, although Linux will work better for this), so that traffic from that location will be routed from lan interface to the VPN_interface. Here is a catch: you will either NAT this traffic with a source of your SSL_IP_pool or you can let it this way.
Step 4: if you don't NAT you have to add on Fortigate static routes for the remote office network and also firewall rule on the ssl.root interface-->to-->HQ_internal.
You can do this, but that extra_vpn_equipment_money you don't want to spend would be NAT-ed into some workstation_configuration_sweat.
This is a bit of an out of the box answer, but depending on your wireless vendor, several of them offer a "remote AP" configuration that allows an AP to create a VPN tunnel back to the controller when it boots.
Your remote (and temporary) users would then connect to the AP and back to the controller over VPN. You would have to come up with other solutions for things like printing (wireless printer?), but it would be possible especially for a 2 month deployment.
OpenVPN is an SSL VPN, that you could deploy on commodity hardware,
it can do SSL based Site-to-Site VPN, or Remot access VPN.
contrary to popular belief, SSL VPN can do Site-to-Site VPN just as well or better than IPSEC.
I don't know if specific recommendations are allowed but.
I would suggest get an old desktop PC, install 1 additional NIC, and install pfsense firewall.
which supports IPSEC as well as OpenVPN (SSL/TLS based) (which gets around firwall and NAT issues)
this solution costs nothing and very small learning curve, it even handles Dynamic DNS if you do not have a public IP
You can probably pick up a used Fortigate-60 on ebay for less than $50 ...I just saw some there.
Stick that in the remote office and create an IPSEC tunnel back to the central office. That's what I do. You don't need any subscriptions on it.
The general idea of an SSL VPN is to allow individual remote teleworkers access to central resources. Your best bet in this case is to deploy a small VPN device at the remote site and tunnel all the traffic back to the central site.
As Tim has said, the idea of the SSL VPN is for it to be used by individual users, not to create a site-to-site VPN.
Is there any reasoning that you don't want to put a VPN capable device at the remote site, if its only a small site there are small/cheap consumer grade devices that could do it or even one of the low end firewalls like the Cisco ASA 5505.
The best solution would be to add a VPN Router/Firewall in the remote office.
But since you dont want that, I'd go with running Linux on a physical or virtual machine in the remote office, and use this Linux box to establish a VPN link to your Fortigate VPN concentrator at the Central Office. You'll also have to configure the routers at both offices to route traffic over the VPN link as per your policy (Selective traffic, or all internet traffic etc)
You can probably do the same thing with a Windows machine, but I've found Linux or FreeBSD works the best for such computing tasks.
Please make sure that the physical/virtual machine has enough computing power to handle the encryption and decryption without slowing down the traffic.
I would have to second the answer from hyussuf, as I was just about to respond the same. If you absolutely cannot purchase any additional (new) hardware, running openvpn is probably the right answer for you. I haven't used pfsense personally, but I have used LEAF as a dedicated firewall instance (running shorewall).
I'm not specifically familiar with the Fortigate appliance, but, if I'm not mistaken, the appliance you reference supports IPSec VPN tunnels.
