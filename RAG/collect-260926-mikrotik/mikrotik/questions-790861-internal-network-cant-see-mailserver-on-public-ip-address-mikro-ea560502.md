---
id: collect-260926-mikrotik/mikrotik/questions-790861-internal-network-cant-see-mailserver-on-public-ip-address-mikro-ea560502
title: "questions-790861-internal-network-cant-see-mailserver-on-public-ip-address-mikro-ea560502"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-790861-internal-network-cant-see-mailserver-on-public-ip-address-mikro-ea560502.md
source_anchor: ""
source_lines: [1, 29]
sha256: 98674724b8f0f0b46e640636097895fdf2cb61a8ca437b2c7e6d50eed7cc3eb2
---

# questions-790861-internal-network-cant-see-mailserver-on-public-ip-address-mikro-ea560502

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Please esxcuse my english as it is not my home language
I have mailserver in my network behind a mikrotik router and firewall
the problem that I'm having is that my mailserver is on a public ip address
and that im not able to connect to it internaly. I have no problem connecting from outside the network. The only way I'm able to connect at this point is to assign the public ip to the mikrotik router and setup these two NAT rules
IP Address I have used are only examples
- Internal Network = 172.162.30.0/12
- Public IP mailserver = 42.20.16.18
- Private IP mailserver = 172.162.30.65
**Internal clients Note:** Source address is my whole network this has a massive affect on my internet speed because
I think all traffic is being routed to mail server
chain=srcnat action=src-nat to-addresses=42.20.16.18
src-address=172.162.30.0/24 log=no log-prefix=""
For clients connecting from outside the network
chain=dstnat action=dst-nat to-addresses=172.162.30.65
dst-address=42.20.16.18 log=no log-prefix=""
I tried almost every solution out there this is my last resort
my main goal is that I physically can assign my public ip to the machine and connect internally without using the private ip.
Either you implement NAT Loopback or NAT Hairpin or NAT Reflection (it's the same thing with different names) or you modify your internal DNS so that your mailserver's hostname does not resolve to the public IP but to your local IP (only from inside your private network).
With your src-nat rule you are already doing the first solution (NAT Hairpin) which as you mentioned works. The problem with this approach is that all your connections to the mailserver are being changed to src-address 42.20.16.18 so the mailserver only logs this IP instead of each user's real internal IP.
The other approach is that you either use Mikrotik's DNS server or any other local DNS server of your choice. There you set up a dns record with the mailserver's hostname and you set it to resolve to the internal IP of the mailserver.
Then you must configure all your PCs on your network to use this dns server and when hitting your mailserver's hostname they will resolve the internal IP instead of the public allowing them to connect to it directly instead of having to go through the router via NAT.
For more information about NAT Hairpin you can check the official Mikrotik Documentation here and for more information about the Mikrotik's DNS Server you can check here
For those interested, I encountered an issue with this solution whenever I sent mail to somebody that receivers mail server would detect a different source ip, a source ip that is blacklisted I think that it has to do with the fact that my ISP has masquerading setup somewhere down the line. I don't know, my work around was I created the dst-st and src-nat again, I kept the dst-nat the same as before and just changed the src-nat source address from my whole network to just my mail servers private ip. This worked 100% see below
