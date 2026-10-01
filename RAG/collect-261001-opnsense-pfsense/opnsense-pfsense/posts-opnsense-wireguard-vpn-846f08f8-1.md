---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8-1
title: "posts-opnsense-wireguard-vpn-846f08f8"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8.md
source_anchor: ""
source_lines: [1, 117]
sha256: 50d35e549c2bbf948c5bb22c8ef20ea62fd51d54a428b2e99327b126b4f1af97
---

# posts-opnsense-wireguard-vpn-846f08f8

WireGuard on OPNsense
There are multiple scenarios where WireGuard can be used, but they all require different configs for that setup.
This writeup should enable a single user (Roadwarrior) and further along the line, a Site-to-Site VPN.
WireGuard then OPNsense
This tutorial discusses the setup of WireGuard first.
Please hang on till we’re done with WireGuard.
None of this will work without the client, peer setup, and also changing the OPNsense router and firewall.
Tutorial Steps
- This Intro
- Client Device Setup
- WireGuard
- OPNsense
- Site-to-Site VPN Router/Firewall
- Getting the Peer Client’s Device Connected
- Adding DNScrypt to Unbound on OPNsense
WireGuard intro
To setup WireGuard first, you must understand it’s conceptual overview.
- It’s not a client/server VPN setup.
- Both sides are a peer.
- So, in essence - you need to setup private and public keys for each side. (Technically you will derive the public key from the private key.)
- No dynamic IP assignment (or very little), each client has a fixed IP.
- WireGuard associates tunnel IP addresses with public keys and remote endpoints.
Cryptokey Routing
At the heart of WireGuard is a concept called Cryptokey Routing, which works by associating public keys with a list of tunnel IP addresses that are allowed inside the WireGuard tunnel.
Each WireGuard network interface has a private key and a list of peers.
WireGuard conceptual overview example
Someone sending a packet
- When WireGuard needs tosend a packet, it looks for the IP address in its Local Addresses Subnet.
- Let’s say, this packet is meant for 192.168.30.8 .
- Your computer’s WireGuard looks for which peer that is.
- Okay, it’s for peer ABCDEFGH . (Or if it’s not for any configured peer, drop the packet.)
- Encrypt entire IP packet using peer ABCDEFGH’spublic key .
- Now where do I send peer ABCDEFGH’s encrypted packet?
- What remote endpoint was listed for peer ABCDEFGH?
- The endpoint is listed as 216.58.211.110 port 53133.
- Send encrypted data over the Internet to 216.58.211.110:53133.
WireGuard receiving a packet
- When an interface for WireGuardreceives a packet, this could be from port forwarding or an open interface, it attempts to identify it.
- The WireGuard interface checks the source IP address and port to determine whichpeer the packet is from.
- Once the peer is identified, WireGuard looks up the corresponding key associated with that peer from its internal configuration.
- It then begins to decrypt the information that was sent.
- WireGuard then uses ABCDEFGH’s key to verify the authenticity of the packet.
- With the accepted, verified packet - WireGuard will now remember that peer’s most recent Internetendpoint .
- Once all that has been completed, WireGuard will look at the packet.
- It can see the plain-text packet from someone on the 192.168.30.X subnet.
- A final verification takes place as WireGuard looks at the packet from 192.168.30.X to verify that IP is evenallowed to be sending us packets.
- If the association is successful, the packets areallowed to pass through the VPN tunnel.
Summarize WireGuard routing
1
Some App -> WireGuard -> Destination IP in tunnel -> Public key for peer holding that IP -> peer's most recent Internet endpoint
PART II - The client device setup
WireGuard on the peer’s client machine
WireGuard is an exchange of keys. Your OPNsense firewall’s WireGuard cannot connect with a peer it doesnt have a key for. Go ahead and skip this step if you’re not using WireGuard as a Roadwarrior setup. This step is for the client machine’s peer setup. If you’re doing a WireGuard Site-to-Site VPN you may proceed to the interface creation.
There are many ways to do this, we’ll use COPY/PASTE. For ease of use, you can copy/paste a generated config from OPNsense in the WireGuard Peer Generator instead of using the clients below to generate one.
What is your machine?
This is the WireGuard peer client’s software connecting back to OPNsense. Here’s a few that I can mention:
Android
- The official WireGuard app for Android - The official app includes an auto-updater, this is against F-Droid policy, and you will find this app at the IzzyOnDroid Repository.
- WireGuard Tunnel - An alternative client app for WireGuard with additional features, available on F-Droid.
iOS
- The official WireGuard iOS App Store app - iOS’s official WireGuard app for iOS 15.0 or later. Mostly feature parity with Android.
MacOS
- The official WireGuard MacOS App Store app - Apple’s Mac App Store’s WireGuard app for macOS 12.0 or later. This app allows users to manage and use WireGuard tunnels.
Windows
- The official WireGuard for Windows app - This installer is the only official and recommended way of using WireGuard on Windows.
Linux
- KDE - Since Plasma 5.15, Plasma support WireGuard VPN tunnels, when the appropriate Network Manager plugin is installed.
- Ubuntu - WireGuard gtk gui for linux
- Ubuntu Server - quick forum guide with official ppa
- Debian Server - official quickstart documentation
- Arch - arch wiki
- Raspbian - WireGuard.how’s guide for Raspbian OS Bullseye
- RHEL - Red Hat’s official documentation on WireGuard
WireGuard client key generation
There are many different clients listed above. The concept below remains the same for each of them.
Generate a private key for this peer’s client
Use your GUI
You can use any of the GUI clients to hit a button to generate a Private Key and a Public Key.
Copy the Public Key.
Windows - Example
In this example we’ll be using the Official Windows WireGuard client.
- Heading over to the client machine.
- Open the WireGuard client application.
- Use the Add Tunnel drop down arrow and selectAdd Empty Tunnel .
- You will have an [Interface] with aPrivate Key = …. DONT MESS WITH IT
- Copy the Public Key at the top, including the equals sign.
- Dont copy anything else.
More information can be found here.
Use the terminal
You can also use any terminal client to do basically the same with wg genkey.
Linux Terminal - Example
This quick script will generate the keys needed to /etc/wireguard, and print them to the screen:
1
if [ "$EUID" -ne 0 ]; then echo "Please re-run as root." && sleep 2 && echo "Application Exiting..." && sleep 30 && exit 1; fi && echo -e "\n\nSetting up public & private key in /etc/wireguard\n" && wg genkey | tee /etc/wireguard/$HOSTNAME.private.key | wg pubkey > /etc/wireguard/$HOSTNAME.public.key && chmod 600 /etc/wireguard/$HOSTNAME.private.key /etc/wireguard/$HOSTNAME.public.key && echo "Private Key:" && cat /etc/wireguard/$HOSTNAME.private.key && echo -e "\nPublic Key (copy this):" && cat /etc/wireguard/$HOSTNAME.public.key;
WireGuard peer client - review
- You will need to have the public key from your client copied.
At this point we should have a public and private key generated for our client.
PART III - WireGuard Configuration
WireGuard on OPNsense
Install WireGuard on OPNsense
WireGuard is now Kernel level in OPNsense. There is no need to download a package anymore.
WireGuard tunnel subnet interface
This is the configuration for the tunnel address of the OPNsense endpoint, the “server”.
Instance is the WireGuard interface’s subnet.
- VPN ▸ WireGuard ▸ Settings ▸ Instances
- New (Click on the + symbol)
- Enabled
- Name - wgopn1-memestorThese names wont be seen anywhere outside of this config screen. BUT, you will see an interface name when you go to assign an interface.
- Public Key - Hit the cog. You will see a series of characters with an equals sign that will always appear at the end.
- Keep hitting the cog until a public key with a series of numbers and letters appears without any special characters (except the = indicates the end of the key).
- Listen Port - 51820
- Tunnel Address - 10.2.2.1/24
- Peers - Blank for now
- Save
- Apply
WireGuard client endpoint as a peer
WireGuard peer info
