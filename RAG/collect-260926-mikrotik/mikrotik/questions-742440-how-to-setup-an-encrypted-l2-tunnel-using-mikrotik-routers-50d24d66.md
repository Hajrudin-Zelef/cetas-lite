---
id: collect-260926-mikrotik/mikrotik/questions-742440-how-to-setup-an-encrypted-l2-tunnel-using-mikrotik-routers-50d24d66
title: "questions-742440-how-to-setup-an-encrypted-l2-tunnel-using-mikrotik-routers-50d24d66"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-742440-how-to-setup-an-encrypted-l2-tunnel-using-mikrotik-routers-50d24d66.md
source_anchor: ""
source_lines: [1, 20]
sha256: 535f166d4a9bfa32f31d0c6980de5f836c6a51c5e9027117b7d8182fd36f42f2
---

# questions-742440-how-to-setup-an-encrypted-l2-tunnel-using-mikrotik-routers-50d24d66

What I would like to achieve
I want to securely spread an existing internal subnet over multiple buildings. That means that I have two locations with virtual machines that need to be within the same subnet. The idea is that the virtual machines (having a static IP) can be migrated from one location to the other.
The (physical) host machines are connected to a switch at each location. So, if there wasn't any security or cost problem I would simply connect both switches with a network cable:
[Machines]---[Switch A] <---- LONG CABLE ---> [Switch B]---[Machines]
What I would like, is to replace this long cable by an encrypted tunnel using two gateways that don't need to care about IP adresses or routing and just take any incoming packets encrypt them and send them to the other gateway via an encrypted tunnel. The other gateway then decrypts the packets and sends them to the remote switch. This would physically look like this:
[Machines]---[Switch A]---[GATEWAY A] <-- INTERNET --> [GATEWAY B]--[Switch B]---[Machines]
I would like to avoid that the gateways need any IP adresses within the subnet. The rules shall be completly port-based:
- Incoming data at port 1: route through tunnel interface
- Incoming data at tunnel interface: route via port 1
The two Gateways would have a static, routeable IP address to establish the tunnel. The encryption shall be strong (at least AES128, SHA256, DH2048; shared secret is fine), which simple PPP type tunnels don't support. So an additional/seperate encryption layer might be needed.
I've only MikroTik Routers available. So I would prefer to use them. However, I'm mostly looking for the 'magic words' (protocol names and the like) and the right combination of technology that allows me to do that. So, if you know how to do it with Cisco routers or HP routers, it would probably also help, if you explained how you would just do it with that ones...
Questions/Attempts
What kind of firewall filters and protocols can I use to achieve this?
My first idea was to use IPsec to span the encrypted tunnel. But, then I would need to define IPsec Policy that is physical-port based. But there is only an option to define that data from/to a special IP-adress / IP-port combination.
So IPsec would just work as encryption layer for another tunnel type (PPTP, SSTP, L2TP and OVPN are currently supported by the MikroTik RouterOS). As PPP-Tunnels typically don't support strong encryption, I would let IPsec do this job and span the unencrypted-PPP-Tunnel through the encrypted IPsec-tunnel.
Ok, now we had at least some tunnel interface, that we can use like a outgoing port. However, I'm kind of lost here. I don't find that possibility to say: "a frame with incoming at has to be sent out via interface " and "a frame incoming at has to be sent out via interface ".
I'm not often working at Layer2... so I'm actually looking for the right 'term' or 'category'. I could imagine finding it at the IP-Firewall (mangle->prerouting) or something like that, but I assume that's already Layer 3 stuff...
Do I just need to setup a bridge? If so, how can I add the tunnel-interface to the bridge (preferably using the winbox-Interface)? Does the bridge need a MAC adress?
Just in case this attempt is a dead end: I also found "EoIP", "IP Tunnel" and "GRE Tunnel" at the "Interface" setting. But I've no real idea what they can do... So just in case, let me know which of them is worth investigating...
Also, If there is a more easy-and-clean solution don't mind just telling me your solution... You don't have to continue my above attempts, if there is just an easier way!
