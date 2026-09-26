---
id: collect-260926-mikrotik/mikrotik/best-practice-for-multiple-offices-interconnection
title: "best-practice-for-multiple-offices-interconnection"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/misc/best-practice-for-multiple-offices-interconnection.md
source_anchor: ""
source_lines: [1, 49]
sha256: 31ec3c0c54e5a9a81acec93ba90da1a9b853341f60a82735a2c182310dd4939f
---

# best-practice-for-multiple-offices-interconnection

Hello, I know there are many ways to connect sites over network, my question is more about the correct way..

At this moment, I’m using packet marking with list of addresses and routing to VPN network, it is good from one side, but connection from other side is unstable..

I am also updating the stability of the sites so each office has redundant WAN connection, in case primary connection is broken, it is switched to secondary backup WAN.. for that I used recursive failover found here https://www.prinmath.com/ham/mikrotik-failover.htm

Problem is that there is change of the public IP address and other VPNs disconnects from the network.. so the solution as I thought is to use external stable root point which is cloud hosted router on external stable site. For that I’ve installed mikrotik CHR on bought virtual machine with static public IP where I want other sites to connect to. I’ve found relevant video, but with no exact configuration to be set up.. https://www.youtube.com/watch?v=m4xavyO4sok

To explain my situation, I have 3 offices, where local networks are 192.168.10.0/24, 192.168.15.0/24 and 192.168.20.0/24.. Now I need to be able to connect for example from client computer with obtained ip address 192.168.10.8 connect to NAS which is on 192.168.20.99 and to be able to connect client computer on 192.168.20.6 to 192.168.10.88 webserver.

What is the best VPN service to do such interconnection? How should the configuration of mikrotik routers looks like?

Thank you!

L.

             
            
           
          
            
            
              Hey. My advise is to use EoIP tunnels over IPsec(do not merge them in a hub) and run OSPF on loopback interfaces on each office router. Then configure iBGP from each loopback and make server’s traffic exchange via iBGP with even prefix filtering from wherever point you want.

             
            
           
          
            
            
              I am presuming all sites have Mikrotiks as gateways.

I would do something simple.

- Connect all 3 sites between them (imagine a triangle) with GRE+IPSEC
- Enable OSPF on the GRE interfaces and on the bridge interfaces

you will have:

- encryption between sites ( choose IPSEC parameters acording to the wiki so you keep the hardware encryption)
- automatic rerouting(example: if connection from site 1 to site 2 fails, site 1 can still talk to site 2 via site 3)
- only one dynamic protocol to take care of

Now regarding you Dynamic WAN IPs - you can enable IP/Cloud service on each mikrotik router. This will basically generate a random DNS name hosted by Mikrotik which will update your WAN IP every minute. From there you can:

1. Do a script that updates the GRE endpoints if the DNS gets updated with a new IP. Example: Site 1 router checks the dns name of sites 2 and 3 every minute and changes the GRE endpoint related to those sites if needed
2. Target the GRE tunnel directly at the DNS name. I have never tested this but ROS seems to allow it in the configuration. Not sure how often it checks the DNS. Maybe someone has input.

PS: Regarding your idea of a cloud hosted router with stable Public IP. The only way that works is if you use client-server VPNs such as OPENVPN or L2TP/IPSEC. You make the cloud hosted one a server and the sites clients. This way it does not matter if the client (your sites) WAN IP changes as long as the “hub” remains unchanged.
