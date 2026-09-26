---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-z3lmx1-eoip-tunnel-with-only-one-public-ip-address-912501a4
title: "r-mikrotik-comments-z3lmx1-eoip-tunnel-with-only-one-public-ip-address-912501a4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/r-mikrotik-comments-z3lmx1-eoip-tunnel-with-only-one-public-ip-address-912501a4.md
source_anchor: ""
source_lines: [1, 40]
sha256: 7c90f2f068e54da534c8e82673ea089068f096064dc7af406c6f23d4dabeb701
---

# r-mikrotik-comments-z3lmx1-eoip-tunnel-with-only-one-public-ip-address-912501a4

EoIP tunnel with only one public IP address 
        
    Is it possible to establish an EoIP tunnel between two mikrotik devices, of which only one has a public IP address? All the examples/guides seem to show the simple configuration where each of the routers specifies the IP of the other end of the tunnel.
What I wanted to achieve is simply to connect two private LANs so that they would be bridged over EoIP, but in my case only one of the gateways has a public IP address. I imagine it could be possible to set it up so that the "public" gateway works like a "server" listening to connections, while the private gateway connects to the public IP address to establish and maintain a two-way tunnel.
Is it achievable on mikrotik devices via EoIP or maybe similar protocol?
Section des commentaires
Hi!
I'd recommend WireGuard instead of EoIP. WG is carried over UDP and the endpoint behind NAT can have keepalive enabled in order to maintain contact with the public endpoint.
WG however requires RouterOS 7.6
Commentaire supprimé par le membre
Ah, meant 7, thanks 😅
Wireguard is your best option.
I would advice against running EoIP over a public network. Set up a VPN and run your tunnel on that instead.
MikroTik supports EoIP over IPsec.
That's near enough to qualify as a VPN. But running it like that behind NAT may give you a headache.
For extending your subnet use EoIP over IPsec.
For connecting two subnets in VPN use site to site IPsec or GRE over IPsec.
Use IP Cloud in case you don't have static IP on peer/peers
Eoip over an l2tp/ipsec tunnel?
Would one of you mind explaining how 2 sites can be connected by any of the solutions suggested here with one site not having a public ip?
I’m not being facetious — it’s an honest question.
Thanks.
Protocols like L2TP and Wireguard work with only one public IP because of the way connections are NATed (Wireguard does require an additional setting for this, see the NAT section in Wireguard quick start
Once you've had a connection dial out that gets mapped in the connections table (IP>Firewall>Connections) for the duration of the connection.
For example connecting to my office L2TP server from my PC behind my ISP router on being NATed again by the ISP to a shared address results in the following:
Example IP addresses:
My PC: 192.168.1.100
My Router: 100.64.0.2
ISP NAT Router: 198.51.100.10
My L2TP server: 203.0.113.22
My PC initiates the connection to 203.0.113.22
My L2TP server gets an incoming connection from 198.51.100.10
My L2TP server replies to 198.51.100.10
The ISP NAT router checks the connection table. An established connection exists and the ports match so it forwards the packet to 100.64.0.2
My router receives the packet and again matches it in the connection table to my PC and forwards it there.
The L2TP tunnel is established and small amounts of traffic back and forth keep the connection state current in all the routers.
In ROS 7 you've also got zerotier which uses another server to establish the connections meaning you can build a VPN solution without either end having a public address.
Thank you very much for the explanation.
But, I still don’t get it.
What circumstances or situation would have a site connected to the internet but without a public ip address?
