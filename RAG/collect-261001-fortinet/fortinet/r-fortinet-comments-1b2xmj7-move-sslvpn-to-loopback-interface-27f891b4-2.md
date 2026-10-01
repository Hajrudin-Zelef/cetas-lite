---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4-2
title: "r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4.md
source_anchor: ""
source_lines: [54, 88]
sha256: dc36a66f8b3ea6f8d33d0976f9fdb7cc5b20f00c27dcd7e6d46f91687e0599b3
---

# r-fortinet-comments-1b2xmj7-move-sslvpn-to-loopback-interface-27f891b4

Thank you for your input. Yes I would use the same public IP for the SSL-VPN as for the IPsec tunnels. Did you map the specific SSL-VPN port on your VIP object? And which FW train did you test this on?
u/extreme_questions : I tested it on FortiGate 80E and I remember it was 6.4.x.
I lost by notepad++ how to but I didn't map the port.
If I remember it correct.
I created loopback interface and assigned it a IP address
I created Virtual IP and mapped by Public IP to Loop back interface.
I changed the Listen on interface to Loopback interface in SSL VPN settings.
Created a policy with WAN as my source and VIP as my destination with all service.
Then you did it wrong and forwarded all ports to the loopback or something else. You should only forward the necesssary ports. I have this setup on lots of installations and have zero IPsec issues.
u/HappyVlane Yes may be as I didnt map any port.
This is what I wrote down in the comment.
If I remember it correct.
I created loopback interface and assigned it a IP address
I created Virtual IP and mapped by Public IP to Loop back interface.
I changed the Listen on interface to Loopback interface in SSL VPN settings.
Created a policy with WAN as my source and VIP as my destination with all service.
Let me know the necessary step I missed.
I was in the process of configuring our 201F for SSL Azure SAML VPN and saw this post. So I went ahead and tried it with success today. I kept seeing people knocking on my door trying to authenticate and I wanted an easier way to control this traffic. So it was perfect timing.
I am doing central SNAT and that was different from the guide I was following here: https://www.youtube.com/watch?v=T_l-do_oci8
Even though I was doing central SNAT I still configured my NAT under the DNAT/Virtual IP area. I picked one of our open public IP's and did the DNAT/Virtual IP to the loopback I created. Also note, if you want ping to work you need to include that in the allow firewall policy in addition to HTTPS. Even if you enable ping on the interface itself, you needed to add it to the policy.
I tested a block policy and added the public IP of my test machine and confirmed it blocked the connection as expected. So I will probably go through the other reddit post and start blocking some AS ranges to tighten this up. That was located here: https://www.reddit.com/r/fortinet/comments/1b2ewwo/using_sslvpn_reduce_your_security_footprint_block/
I had some issues that ended up working themselves out. But I had Fortinet support on and he said this wasn't really documented on their site at all. So he took a backup of my config and said he will probably write up a official guide on this which is cool.
So thank you reddit once again for giving me ideas!
Here from Fortinet - it shows also how to block Internet Services than on this interface.
https://community.fortinet.com/t5/FortiGate/Technical-Tip-Prevent-TOR-IP-addresses-from-accessing-SSL-VPN/ta-p/269785
Can I just create the policy to block ip from wan interface to lan interface (to deny certain IPs to connect sslvpn) instead of making a loopback interface? What is the difference between these two methods?
I dont see why anything you said makes a difference?
Create a loopback interface. Give it an IP.
Create a VIP that forwards a specific port number from your WAN interface to the loopback IP.
Set the SSLVPN service to listen on the loopback interface.
Create policies.
Those of you using this , are you getting LOG entries
I have had this setup and not getting log , which i would like to do and confirm my geo-block and malicious and tor policy
My Log level is set to ALL
It will hit your Local traffic log now
