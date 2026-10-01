---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1ago3w5-inbound-traffic-issues-1e609270
title: "r-fortinet-comments-1ago3w5-inbound-traffic-issues-1e609270"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1ago3w5-inbound-traffic-issues-1e609270.md
source_anchor: ""
source_lines: [1, 24]
sha256: 01b39f7a2961c055d1079b2b4e9170046a3fe5ea5d41f9ef4ef9176599fedad7
---

# r-fortinet-comments-1ago3w5-inbound-traffic-issues-1e609270

inbound traffic issues 
        
    Recently we had to move from Comcast fiber to a BCI while waiting on the fiber to be moved. We purchased a static block of 5 and when I first cut the client over ping/SSLVPN(444) worked then after a few minutes it stopped responding. I checked with Fortinet support and we did a debug and do not see any traffic even hitting the Fortigate when I attempt to connect from my remote computer. I spent hours with a Comcast support tech on site, his boss, inhouse support, and tier 2 support and they said that when they load our static IP to the modem it is pingable but when I move it back to my firewall it stops. I did not make many configuration changes on my Fortigate besides editing the interface, route, and policy to reflect the new IP changes. I am becoming very frustrated because I cannot escalate further with Comcast and no clue what is going on.
Has anyone ever experienced an issue like this?
      *UPDATE: SOLVED*
I did a packet capture from the WAN interface and actually saw traffic. I looked and saw a DNAT for a local Citrix server that I would not expect to see for VPN traffic and found out that we had a 1:1 static NAT and policy combo that was forwarding all traffic to this server. I disabled the policy and now we can connect on the VPN and ping the public IP.
    
Section des commentaires
When you perform a packet capture on the FortiGate, do you see the traffic on its interface?
https://community.fortinet.com/t5/FortiGate/Troubleshooting-Tip-Packet-Capture-on-FortiOS-GUI/ta-p/194444
I am not doing a packet capture but a debug. Here's the CLI commands:
diagnose debug application sslvpn -1
diagnose debug enable
When I do this I see nothing. I did a Wireshark capture from my computer when trying to connect and I do think there's a response coming back.
Run a pcap on the entire incoming interface.
Also what's the route table say?
This could also be when you get both supports on a call together.
Hey man this suggestion was spot on! I did a packet capture and debug and noticed a DNAT that looked off so I investigated and found we had a Virtual IP/policy combo that was routing ALL inbound traffic for any port to a Citrix server on their LAN. I disabled that policy because it is actually not in use(no hits) and it works now. I'm going to make this as solved.
Sweet amigo glad you got it to work out.
As a test can you route traffic out via the new IP block? If not then it will probably be a routing issue with comcast
When you say route outside the new IP block do you just mean like internet access and pinging outwards and such? If so yes that all works thankfully.
Yep that's spot on. It most likely that Comcast are most likely not forwarding the traffic to your firewall then.
If you add an IP from your new block as a secondary IP can you ping the IP?
Note that you will need to enable ping on the wan interface.
