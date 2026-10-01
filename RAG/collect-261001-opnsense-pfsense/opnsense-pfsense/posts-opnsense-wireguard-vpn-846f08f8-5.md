---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8-5
title: "posts-opnsense-wireguard-vpn-846f08f8"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8.md
source_anchor: ""
source_lines: [441, 612]
sha256: 8a6f19ec04a931ba1b83790abc297aaf4d5c590e18731a3eb151f0c83a36285e
---

# posts-opnsense-wireguard-vpn-846f08f8

You can use this method to connect your home to a VPS in the cloud, or you mom’s house to your house.
For reference, link to the table of Site A and Site B LAN addresses used in this writeup.
Site A - Router Pass Traffic - Bunker HQ
Please make sure you do not have overlapping subnets.
Allow traffic between Site A LAN Net and Site B LAN Net
The first firewall rule will make sure our Bunker HQ (172.16.0.0/24) can reach Sarah (192.168.0.0/24).
- Go to OPNsense Site A
- Open Firewall ▸ Rules ▸ LAN andadd a new rule.
  - Note: Change for preference . The network you want to share with WireGuard may be different thanLAN , please modify the name of the Interface to match your network.
- Note: 
- Action - Pass
- Interface - LAN
- Direction - In
- TCP/IP Version - IPv4
- Protocol - Any
- Source - 172.16.0.0/24
- Source port - Any
- Destination - 192.168.0.0/24
- Destination port - Any
- Description - Allow LAN on Site A to remote LAN Site B
- Press Save and Apply.
Allow traffic from WireGuard Site B LAN Net to Site A LAN Net
The second firewall rule is through the WireGuard tunnel. It allows Sarah’s LAN to reach the Bunker HQ’s LAN.
- Go to OPNsense Site A
- Open Firewall ▸ Rules ▸ WireGuard (Group) andadd a new rule.
- Action - Pass
- Interface - WireGuard (Group)
- Direction - In
- TCP/IP Version - IPv4
- Protocol - Any
- Source - 192.168.0.0/24
- Source port - Any
- Destination - 172.16.0.0/24
- Destination port - Any
- Description - Allow LAN on remote Site B to LAN on Site A
- Press Save and Apply.
Site B - Router Pass Traffic - Sarah’s Shop
Please make sure you do not have overlapping subnets.
Allow traffic between Sarah’s Site B LAN Net and Site A LAN Net
- Go to OPNsense Site B
- Open Firewall ▸ Rules ▸ LAN andadd a new rule.
  - Note: Change for preference . The network you want to share with WireGuard may be different thanLAN , please modify the name of the Interface to match your network.
- Note: 
- Action - Pass
- Interface - LAN
- Direction - In
- TCP/IP Version - IPv4
- Protocol - Any
- Source - 192.168.0.0/24
- Source port - Any
- Destination - 172.16.0.0/24
- Destination port - Any
- Description - Allow LAN on Site B to remote LAN Site A
- Press Save and Apply.
Allow traffic from WireGuard Site A LAN Net to Sarah’s Site B LAN Net
- Go to OPNsense Site B
- Open Firewall ▸ Rules ▸ WireGuard (Group) and add a new rule.
- Action - Pass
- Interface - WireGuard (Group)
- Direction - In
- TCP/IP Version - IPv4
- Protocol - Any
- Source - 172.16.0.0/24
- Source port - Any
- Destination - 192.168.0.0/24
- Destination port - Any
- Description - Allow LAN on remote Site A to LAN on Site B
- Press Save and Apply.
Now both sites have full access to the LAN of the other Site through the WireGuard Tunnel. For additional networks just add more Allowed IPs to the WireGuard Endpoints and adjust the firewall rules to allow the traffic.
Alt route - no firewall route
You can try and route without any rules. You want the wireguard subnet to be routed in its entirety to and from that gateway so that access can be established and mapped by your router without the need to add firewall rules.
- Go to System ▸ Gateways ▸ Configuration
- Create a new gateway.
- Assign the interface.
- Make sure you have set the IP of your remote wireguard’s gateway’s IP, the tunnel address, in our example above it was Site B to remote Site A: 10.2.2.1/24 .
- With a new gateway up we can point a route there.
- Head to System ▸ Routes ▸ Configuration .
- Create a new route.
- Make the network address your WireGuard subnet.
- Set the gateway dropdown to the one you defined earlier.
- Open Firewall ▸ Settings ▸ Advanced ▸ Static route filtering .
- Check the Bypass firewall rules for traffic on the same interface checkbox.
The rest of the Firewall stuff you know
- Click on Firewall -> Rules -> WireGuard
- Then on the + ADD button.
- Select Single host or Network assource and
- Enter the IP range of theWireGuard network and its subnet mask.
- Save.
- Looking back..
Firewall -> Rules -> WireGuard, you should see – under Source the IP Address of the Internal subnet you’re using for you’re WireGuard network you made under VPN -> WireGuard -> Local.
Finish - by resetting services
The save and apply are meaningless, as WireGuard never resets the service to load the new configuration. You must be sure to either check and uncheck the Enable WireGuard button in Settings ▸ General or go to the Dashboard and reset services from there.
PART IX - The peer client
Setting up the client software
So the client needs:
- Their Client Config
- Endpoint config of the “server” 
  - Public IP of the Endpoint “server”
  - Public Key of the Endpoint “server”
Example: Official Windows WireGuard client
That client software we started up and then left out in the cold, hungry for input. Let’s go do something with that client now.
- Make sure the SAME client software is open, with the same Public Key that was copied earlier. No exit and starting over.
- Copy this into your client config for tunnel generation:
1
2
3
4
5
6
7
8
9
[Interface]
PrivateKey = #somenumber#
Address = 10.10.10.2/32
DNS = #Internally Routed DNS. This is on the subnet you're VPNing into, example 172.23.55.254#
[Peer]
PublicKey = #HEYWAIT!---WeDontHaveThisYet---#
AllowedIPs = 0.0.0.0/0 #This is the subnet youre VPNing into, so this would be, example 172.23.55.0/24#
Endpoint = edge.sub.domain.com:51820
- That’s right! We still have to get the Public Key of our other confidant, the Server we’re connecting to.
- Moving over to the web interface, under VPN ▸ WireGuard ▸ Settings ▸ Local
- Edit the Local Profile's Configuration that we created earlier, namedwgopn1-memestor and copy thePublic Key making sure to get THE ENTIRE THING including the=
- Paste this into the config section of the client, under [peer]
- Save
Example commented wg0.conf
THERE ARE A BUNCH OF STEPS FOR CREATING CONFIGS.
THERE ARE SEVERAL AUTO GENERATING ONES ONLINE
HERE IS ANOTHER EXAMPLE:
1
2
3
4
5
6
7
8
9
10
11
12
13
[Interface]
Address = <Configured client IP>/<Netmask> // For example the IP "10.11.0.20/32"
PrivateKey = <Private Key of the client>
[Peer]
PublicKey = <Public Key of the OPNsense WireGuard instance>
AllowedIPs = <Networks to which this client should have access>/<Netmask>
             // For example "10.11.0.0/24, 192.168.1.0/24"
             //               |             |
             //               +--> The network area of the OPNsense WireGuard VPNs
             //                             |
             //                             +--> Network behind the firewall
Endpoint = <Public IP of the OPNsense firewall>:<WireGuard Port>
Adding DNScrypt to Unbound on OPNsense
DSNCrypt is significantly faster than DoT which is faster than DoH.
DoT is faster than DoH because it works directly at the transport layer, while DoH has additional overhead due to the HTTP layer.
DNSCrypt uses UDP as it’s main performance increase.
Install DNSCrypt-Proxy for DNS encryption
Head over to your firmware section of OPNsense and install DNSCrypt-Proxy
1
2
3
4
5
New packages to be INSTALLED:
	dnscrypt-proxy2
   os-dnscrypt-proxy
Number of packages to be installed: 2
Unbound -> Query forwarding -> DNSCrypt-Proxy
Query forwarding in OPNsense’s Unbound DNS can be particularly useful when you want to access an upstream service, and provide both secure and smooth DNS resolution across your network. Here’s a detailed explanation and tutorial on how to set it up:
Why Use Query Forwarding?
