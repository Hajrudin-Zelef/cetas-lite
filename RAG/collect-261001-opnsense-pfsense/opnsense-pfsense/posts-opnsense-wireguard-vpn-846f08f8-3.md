---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8-3
title: "posts-opnsense-wireguard-vpn-846f08f8"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8.md
source_anchor: ""
source_lines: [200, 305]
sha256: 95e2041a7d1f5466ce93c0cb6ec601132afa377a9d0d9d8f94dca2b7cbc9b7fb
---

# posts-opnsense-wireguard-vpn-846f08f8

- You need to have a Wireguard firewall rule set, for those interfaces to allow traffic in both directions across their respective interfaces. 🔥🧱
Setting up WireGuard on each Instance of OPNsense for Site-to-Site
The following example covers an IPv4 Site to Site WireGuard Tunnel between two OPNsense Firewalls with public IPv4 addresses on their WAN interfaces. You will connect the (Site A LAN) 172.16.0.0/24 to the (Site B LAN) 192.168.0.0/24 using the (WireGuard Transfer Network) 10.2.2.0/24. (Site A Public IP) is 203.0.113.1 and (Site B Public IP) is 203.0.113.2. On the (WireGuard Tunnel Network) the tunnel address for (Site A WireGuard) is 10.2.2.1/24 and the tunnel address for (Site B WireGuard) is 10.2.2.2/24.
Table of all addresses and interfaces used
There are a lot of confusing segments in this tutorial. I have adapted this table to the information being used.
` T A B L E __ O F __ A D D R E S S E S`
| Address | IP | 
|---|---|
| WireGuard Network | 10.2.2.0/24 | 
| Site A - WireGuard | 10.2.2.1/24 | 
| Site B - WireGuard | 10.2.2.2/24 | 
| Site A - LAN | 172.16.0.0/24 | 
| Site B - LAN | 192.168.0.0/24 | 
| Site A - Public WAN | 203.0.113.1 | 
| Site B - Public WAN | 203.0.113.2 | 
“Future-me” is ashamed to have written 10.2.2.0/24. We have (2) IP addresses required. We should use Classless Inter-Domain Routing (CIRD) to promote Variable-Length Subnet Masking (VLSM) and minimize the size of our subnets. An improved subnet size would be - /31.
Assume the network block is 10.2.2.0, and we are allocating a subnet for point-to-point links:
- Network: 10.2.2.0/31
- Usable IPs: 10.2.2.0 and 10.2.2.1
Site A - Meme Storage Bunker HQ’s Server - Instance setup
This is, presumably, the OPNsense router you’ve already been configuring from above - as this is the meme bunker. We can ignore these steps if you’ve already got the first Instance from above already setup, as these steps are primarily the same.
- VPN ▸ WireGuard ▸ Settings ▸ Instances
- New (Click on the + symbol)
- Enable the advanced mode toggle in the upper corner.
- Name - wgopn1-memestor
- Public Key - Generate with “Generate new keypair” cog looking button.
- Copy (and label) this public key somewhere, as we will need it shortly.
- Listen Port - 51820
- DNS servers - Give the address of the router, on the Tunnel address’ peer’s Allowed IPs subnet.
- Tunnel Address - 10.2.2.1/24
- Save
- Apply
Site B - Sarah’s Flower Shop Server - Instance setup
Back in another Instance of OPNsense, we are going to follow mostly the same steps.
- Get off the Airplane in Detroit.
- Call Sarah and see where her Flower Shop is at.
- Arrive at Flower Shop, and login to OPNsense server.
- Visit: VPN ▸ WireGuard ▸ Settings ▸ Instances
- New (Click on the + symbol)
- Enable the advanced mode toggle in the upper corner.
- Name - wgopn2-flwrstor
- Public Key - Generate with “Generate new keypair” cog looking button.
- Copy this public key somewhere, as we will need it shortly.
- Listen Port - 51820
- DNS servers - Give the address of the router, on the Tunnel address’ peer’s Allowed IPs subnet.
- Tunnel Address - 10.2.2.2/24
- Save
- Apply
Site A - MSB HQ’s Server - Peer server setup
This OPNsense is back in the original WireGuard Instance that was made.
The Meme Storage Bunker HQ’s Server, wgopn1-memestor. Site A Instance with Tunnel Address of 10.2.2.1/24.
This part is where you will setup each individual key connecting to WireGuard, and is why the public key of each site’s Instance was copied somewhere handy.
- VPN ▸ WireGuard ▸ Settings ▸ Peers
- New (Click on the + symbol)
- Name - wgopn2-flwrstor
- Public Key - Insert the public key of the instance from wgopn2-flwrstor .Remember, when you were in Detroit? You setup an OPNsense WireGuard Instance at Sarah’s Flower Shop.
- Allowed IPs - 10.2.2.2/24 192.168.0.0/24You are allowing the subnet IPs of Sarah’s Flower Shop Server’s WireGuard Instance’s tunnel (10.2.2.2/24) and the subnet for Site B LAN (Sarah’s Flower Shop’s Detroit based server’s subnet).
- Endpoint Address - This is set to the public IP of the WireGuard Instance we’re connecting to, 203.0.113.2 .
- Instances - Select the Instance to connect this Peer to, wgopn1-memestor .
- Save
- Apply
Site B - SFS’ Server - Peer server setup
Back in Detroit, at our other OPNsense server in Sarah’s Flower Shop - setup is primarily the same.
Sarah’s Flower Shop Server, wgopn2-flwrstor. The Site B Instance with Tunnel Address of 10.2.2.2/24.
Again, setup the key, IPs, and Instance connecting to MSB HQ’s WireGuard.
- VPN ▸ WireGuard ▸ Settings ▸ Peers
- New (Click on the + symbol)
- Name - wgopn1-memestor
- Public Key - Insert the public key of the instance from wgopn1-memestor .The public key you have printed and locked in a firesafe back at the Meme Storage Bunker Headquarters.
- Allowed IPs - 10.2.2.1/32 172.16.0.0/24The is for the subnet of IPs for home’s MSB WireGuard Instance tunnel (10.2.2.1/24) and the subnet for Site A LAN (Meme Storage Bunker at home server’s subnet).
- Endpoint Address - This is set to the public IP of the WireGuard Instance we’re connecting to, 203.0.113.1 .
- Instances - Select the Instance to connect this Peer to, wgopn2-flwrstor .
- Save
- Apply
WireGuard site-to-site setup review
A lot of this was the same as the inital Roadwarrior setup in the beginning. The difference with site-to-site it’s between two OPNsense servers and not a peer client.
- You should have an Instance setup on both OPNsense servers.
- Each WireGuard Instance should have a unique tunnel address on the same subnet.
- A Peer was added with thepublic key that was generated from the server we’re going to connect to’s Instance.
- The Peer needs to have each subnet from the otherSite listed in theAllowed IPs .
- On the same Peer , theEndpoint Address needs to point to the other server’s connectable IP.
- Peers andInstances must alsobelong to each other with thedrop down selecting each, respectively.
Finish - by restarting services
The save and apply are meaningless, as WireGuard never resets the service to load the new configuration. You must be sure to either check and uncheck the Enable WireGuard button in Settings ▸ General or go to the Dashboard and reset services from there.
PART V - WireGuard Interface
Configure a WireGuard Interface on OPNsense
This allows separation of the firewall rules for each WireGuard instance (wgX device).
Assign an interface to WireGuard
Please note, if you have not enabled the WireGuard service the interface creation will fail.
- Go to Interfaces ▸ Assignments
- At the bottom is the Assign a new interface section.
- In the dropdown next to Device , select the WireGuard device you created.
- Add a description (eg wgopn1memestor)This is what will be visible under Interfaces on the menu.
- Click the Add button, then clickSave .
Enable new WireGuard interface
- Click on your new interface ’s description under theInterfaces menu .
- Once on this new screen, we need to Enable the interface.
- Enable - Checked
- Lock - Checked
- IPv4 Configuration Type - NoneThere is no need to configure IPs on the interface. The tunnel address(es) specified in the Instance configuration for your server will be automatically assigned to the interface once WireGuard is restarted
- MTU - 1420
- MSS - 1420
- Save
- Apply
When assigning interfaces, make sure to edit the MTU and MSS. Setting the correct MTU can help avoid packet fragmentation and improve overall throughput. Otherwise you could get working ICMP and UDP, but some encrypted TCP sessions will refuse to work.
The maximum packet size within a WireGuard tunnel is 40 bytes less than the WireGuard MTU. This is because WireGuard adds a 40-byte overhead to each packet for its own headers. Therefore, if your WireGuard MTU is set to 1420 bytes, the maximum packet size that can be transmitted without fragmentation would be 1380 bytes (1420 - 40)
