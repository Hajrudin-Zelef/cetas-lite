---
id: collect-261001-general-networking/general-networking/questions-1714372-registering-openvpn-client-addresses-with-dns-d921ad4b
title: "Registering OpenVPN client addresses with DNS"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1714372-registering-openvpn-client-addresses-with-dns-d921ad4b.md
source_anchor: ""
source_lines: [1, 38]
sha256: 11a2a72db41423231ad9568d7a6afc3ce628489cb24af99b778c455832088773
---

# Registering OpenVPN client addresses with DNS

*Score : 2 | Source : https://superuser.com/questions/1714372/registering-openvpn-client-addresses-with-dns*

I am running OPNSense on my home router and have configured OpenVPN on the device, allowing me to connect to my home network from anywhere in the world.
I am also running a BIND DNS server on my home hetwork, with a dedicated zone for all the systems on that network. So far, all RRs are static and maintained by hand.
I would now like to resolve my client VPN addresses through my internal DNS (the clients in question run a Debian-based Linux distro). As I understand it, I have two options:
The constraint is that OpenVPN is running on the OPNsense box, which limits my ability to install some cutting-edge extension server-side (I have to work with whatever is available officially from the OPNsense repo).
For option 2, there is an article on the OpenVPN wiki, but it refers to a feature under development that is 8 years old at the time of this writing, and appears to require some extra server-side packages which might not be available for my use case.
Any tried-and-true recipes to get my internal DNS to resolve my clients’ addresses, given the constraints I mentioned?

---

### Reponse — score 1

Configure BIND to accept dynamic updates for the "VPN clients" zone. (Note that this is mostly incompatible with hand-maintained zonefiles – either it's dynamic or not – but the nsdiff tool can help with maintaining the "manual" parts of a dynamic zone, or you could manually CNAME each host from your main zone to the dynamic zone.)
Use an OpenVPN --client-connect script to invoke nsupdate to insert new A and AAAA records. (The nsupdate tool comes with BIND.)
Use --ifconfig-pool-persist to make client IP addresses "sticky" after first connection. Even though client-connect scripts will be invoked every time, having a sticky IP address is still useful as it allows the dynamic records to have longer TTLs.
The "DynamicDNS" page talks about the exact same approach – although it doesn't explicitly say so, the update mechanism is just an external Perl script that runs as a "client-connect" script, gets the client info from environment and submits a DNS update.
The constraint is that OpenVPN is running on the OPNsense box, which limits my ability to install some cutting-edge extension server-side (I have to work with whatever is available officially from the OPNsense repo).
In pfSense you could add the standard FreeBSD package repository and install anything from it using pkg add. I'm not sure if that works in OPNSense, but it should.

---

### Reponse (acceptee) — score 0

I have found a different approach which turned out to be easier than setting up dynamic DNS: configure client-specific overrides so that certain clients will be assigned a static address, then register the (now static) client address in DNS.
With the net30 topology (default in OPNsense), each client connection gets its own subnet with a 30-bit mask. Of the two addresses (not counting the network and broadcast address), the first is the server and the second is the client.
Addresses used in client-specific overrides must be outside the IPv4 tunnel network configured for the server.
Steps to configure:
- Increase the subnet mask for the IPv4 tunnel network by one. This means that only half of the range will still be used for dynamically assigned addresses, the rest is now free for client-specific ones. This way, all rules and other configurations based on IP ranges (including elsewhere on your network) will continue to work for clients with static addresses. Optionally, change the network address so only the second half of the former range will be used.
- Add the old IPv4 tunnel network address (with the shorter subnet mask) as an IPv4 remote network. This way, traffic to the entire range (both static and dynamic addresses) will be routed through OpenVPN. (Without this setting, OpenVPN will not be able to reach a client with an address set through a client-specific override.)
- Add a client override for each client to which you wish to assign a static address:
  - Server: choose your OpenVPN server instance
  - Common name: common name as used by the client certificate
  - IPv4 Tunnel Network: assign a portion of the network you freed up in the previous step; it must have a subnet mask of 30 and the last octet must be a multiple of 4.
- If you need communication between VPN clients, add the full range (static and non-static) to the IPv4 Local Networks in your server config. (This will also apply to clients with specific overrides).
- Finally, register the client address (subnet address for that client, plus 2 in the last octet) in DNS.
