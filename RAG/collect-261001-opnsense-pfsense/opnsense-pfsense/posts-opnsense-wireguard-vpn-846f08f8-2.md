---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8-2
title: "posts-opnsense-wireguard-vpn-846f08f8"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8.md
source_anchor: ""
source_lines: [118, 199]
sha256: 70c5aa41a8f45951c79c26b55c9ee120d1ef5d50eec4393dc2abdc037216f543
---

# posts-opnsense-wireguard-vpn-846f08f8

VPN ▸ WireGuard ▸ Settings ▸ Peers ▸ Allowed IPs - These are what IP addresses are going to be permitted over the tunnel.
You can send and the server will receive it, but it will do nothing and send nothing back… UNLESS you have the SUBNET of the Endpoint’s routable network (in it’s [Interface] Address =  section) on the Allowed IPs list for that endpoint.
This is why you have to configure every client that wants to connect to this firewall/WireGuardserver. (Unless they’re sharing certificates.)
That client should have created the private/public key pair, and you will paste the public key… or, try the new and easy method of The Peer Generator.
WireGuard peer creation - Peer Generator
Using the generator, you will not need the public key set earlier, it is defined in the generator. The peer generator will also load in the correct Address for you, but the rest needs set. If you’re doing a WireGuard Site-to-Site VPN go ahead and skip this step, and head to Instance/Peer creation.
- VPN ▸ WireGuard ▸ Settings ▸ Peer Generator
- Select the Instance you would like to use, in this example it waswgopn1-memestor .
- Endpoint - Specify how to reach the instance, usually the public address of this firewall. (e.g. my.endpoint.local:51820)
- Name - thinkpad
- Public key - Set by the generator and copied out of this page in the ‘config’ below.
- Address - Should be automatically set and incremented for each peer in the ‘instance’ subnet.If you’re using the same certificate for multiple clients, then you will need to increase the CIDR subnet for more IP addresses.
- Allowed IPs - List the networks on the other side of WireGuard that we’re allowed to pass through. (e.g. - 10.2.2.0/24, 172.23.10.0/24)
- DNS Servers - Give the address of the router, on that network’s ‘Allowed IPs’ subnet.
- Config - Here lies the generated WireGuard config. Copy this text and save it to for your client, ‘WireGuard-memestor.conf’You will NOT get a chance to copy the ‘Private Key’ again, as it will only appear on here this screen, now. Refresh, clears it - Save, clears it.
WireGuard peer creation - Manual Creation
If you’re using the ‘peer generator’ instructions above, feel free to skip this section. This page is where you setup each individual key connecting to WireGuard, and is why we were required to setup the client, for the key generation earlier. Again, if you’re doing a WireGuard Site-to-Site VPN go ahead and skip this step and head below.
- VPN ▸ WireGuard ▸ Settings ▸ Peers
- New (Click on the + symbol)
- Name - thinkpad
- Public key - Paste in thePublic Key from the client machine you generated earlier.If you forgot, you will need to go to the client you’re using and copy the Public Key.
- Pre-shared Key —Optional , and may be omitted. This option adds an additional layer of symmetric-key cryptography for post-quantum resistance. You will need to take this key from OPNsense to the client. Write it down or copy it.
- Allowed IPs -10.2.2.0/24, 172.23.10.0/24Allowed IPs adds a route inside of opnsense to the “allowed IPs” subnet over WireGuard’s local profile’sTunnel Address .
- Endpoint address - coolserver.dyndns.netHow do you reach the server you’re connecting to?
- Endpoint port - 51820
- Instances - Should have the name of the tunnel subnet we made earlier, wgopn1-memestor .
- Keep Alive interval - 15
- Save
- Apply
If a site/instance/peer is behind NAT, a Keep Alive has to be set on the site behind the NAT. The Keep Alive should be set above, if you followed the tutorial, to 25 seconds as stated in the official WireGuard docs. It keeps the UDP session open when no traffic flows, preventing the WireGuard tunnel from becoming stale because the outbound port changes. Tailscale, Zerotier, Netbird all do the same thing.
Behind a NAT? Probably. Set keep Alive!
- Does your server switch IP addresses?
- Is your server restricted to no open ports?
- Dont worry, they make a Diet Soda special for you, it’s called PersistentKeepalive .
One side of the Wireguard connection can be hidden
The server with the open port doesn’t need to have a PersistentKeepalive, because it’s already accessible from the outside.
The client behind NAT needs to send periodic keepalive packets to maintain the NAT mapping and keep the “hole” open.
The PersistentKeepalive setting on the client serves two purposes:
- It keeps the NAT mapping alive on the client’s side, allowing incoming packets from the server to reach the client.
- It ensures that the stateful firewall or NAT mapping remains valid, allowing bi-directional communication.
A typical configuration would set PersistentKeepalive to a value between 25 seconds and 15 seconds on the client side. This interval is usually sufficient to keep most NAT mappings alive without generating excessive traffic.
Finish WireGuard initial config - resetting services
The save and apply are meaningless, as WireGuard never resets the service to load the new configuration.
You must be sure to either check and uncheck the Enable WireGuard button in Settings ▸ General or go to the Dashboard and reset services from there.
WireGuard configuration - review
Reviewing what should be completed at this point.
- You should have an Instance setup with atunnel address that you can use.
- There should also be a Peer with thepublic key that was generated from your client, client being the remote machine you’re using to connect back to OPNsense.
- On the same Peer , you should also have anIP set for your OPNsense’s peer within the Instance’ssubnet `.
Caveats
Allowed IPs
Generally, it is important to keep the subnets small on the endpoints, especially when using multiple endpoints.
- Allowed IPs sets routes for you, using WireGuard.If you want to talk to that network after you’re connected, list the subnet here.
- Route all traffic over the connection (including the internet):AllowedIPs = 0.0.0.0/0
Allowed IP Explained
In other words, when sending packets, the list of Allowed IPs behaves as a sort of routing table, and when receiving packets, the list of Allowed IPs behaves as a sort of access control list.
- In sending direction this list behaves like a routing table.
- In receiving direction it serves as Access Control List.
One Side Needs a Static IP
If you use hostnames in the Endpoint Address, WireGuard will only resolve them once when you start the tunnel. If both sites have dynamic Endpoint Addresses set, the tunnel will stop working if a site receive a new WAN IP lease from the ISP.
To mitigate this you’d need to check DNS resolve the IP addresses for the system, then restart WireGuard if there was a new IP.
IP addressing with WireGuard on OPNsense
OPNsense’s subnetting example:
1
Tunnel Address  >   Allowed IPs      >   OPNsense [Interface] address
WireGuard subnetting example:
1
2
3
4
   Instances    >     Peers          >   Address on the Client's Interface
       ^                ^                           ^
     Entire       Any Subnet to           Single IP that resides on both
     Subnet       connect between         the Peer's and Instance's subnet
Each WireGuard network interface has a private key and a list of peers.
Each WireGuard peer on OPNsense must have the client’s public key and match the tunnel’s Allowed IPs
PART IV - WireGuard Site-to-Site VPN
You may skip this section if you do not require a site-to-site WireGuard VPN, and you are strictly using this as a roadwarrior way to remote devices into your OPNsense router.
Doing an OPNsense Site-to-Site WireGuard VPN
The peer generator didnt help site-to-site WireGuard VPN config generation at all.
- You need a seperate Instance on BOTH locations for a Site-to-Site VPN over WireGuard. 💾👍
- You need to connect those Instances with a respective peer on both sites. 😀-😀
