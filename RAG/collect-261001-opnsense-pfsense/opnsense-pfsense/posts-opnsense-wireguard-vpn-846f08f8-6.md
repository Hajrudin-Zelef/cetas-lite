---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8-6
title: "posts-opnsense-wireguard-vpn-846f08f8"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-wireguard-vpn-846f08f8.md
source_anchor: ""
source_lines: [613, 745]
sha256: 51b773619014c02145c14271e33f3c25bfda9bd249892d2e8feba7131f5ffcb0
---

# posts-opnsense-wireguard-vpn-846f08f8

Using query forwarding in Unbound with DNSCrypt ensures encrypted DNS queries for privacy and security, while still leveraging Unbound’s caching and filtering capabilities. It prevents ISPs or third parties from snooping on DNS traffic, supports DNSSEC for authenticity, and allows anonymization through DNSCrypt. This setup is ideal if you prioritize local domain resolution while securely forwarding external queries to DNSCrypt-Proxy for encryption.
This ensures that local devices can communicate using their hostnames, while external traffic is securely encrypted.
Keeping Unbound
In this section of the tutorial we need to establish the upstream (DNScrypt) server.
- Services > Unbound DNS > General
- Ensure Enable Unbound DNS is checked.
- Under Network Interfaces, select LAN/VLANs.
- Enable DNSSEC Support.
- Optionaly: 
  - Register ISC DHCP4 Leases.
  - Register DHCP Static Mappings
- Services > Unbound DNS > Query Forwarding
- Uncheck Use System Nameservers
- Add Server:
- Server IP: 127.0.0.1
- Server Port: 15353
- Description: Forward DNS to DNS-Crypt
- Disable/Delete any DNS over TLS records
- DNS over TLS will be preferred. To use the forwarding server as your primary DNS, you must disable any local DNS resolution.
- Click Save & Apply
Setting up DNSCrypt-Proxy
- Services > DNSCrypt-Proxy > Configuration
- Listen Address: 127.0.0.1:15353
- Uncheck: Allow Privileged Ports
- Based on your requirements for your DNS server you may check: 
  - Require DNSSEC
  - Require NoLog
  - Require NoFilter
  - Block IPv6
- Set the Server List, these are the servers you’re reaching out to for DNS information. They all have different features. Please see this address for a list: https://dnscrypt.info/public-servers 
  - Based on your location pick the closest dnscry.pt- to you.
- Based on your location pick the closest 
- Save
(Optionaly) Enabling Encrypted Client Hello
When using The DNSCrypt setup with DoH you have the option of using ECH. This will loose the Domain name with Encrypted Client Hello (ECH). By loose I mean your SNI wont be seen.
Pretty much this isnt implimented, so your only real option is to use Cloudflare and send all your DNS info to them, just like with a VPN, they are on the other end and can see the requests:
https://github.com/DNSCrypt/dnscrypt-proxy/wiki/Local-DoH
https://tls-ech.dev/
https://defo.ie/ech-check.php
Multiple Servers on Same Domain
Query forwarding in OPNsense’s Unbound DNS can be particularly useful when you have multiple OPNsense instances in your network, especially when one OPNsense router is behind another. This feature helps ensure smooth DNS resolution across your network hierarchy. Here’s a detailed explanation and tutorial on how to set it up:
Why Use Query Forwarding?
Query forwarding allows you to:
- Resolve local hostnames behind a secondary OPNsense router
- Forward DNS queries upstream to your primary network
- Maintain consistent DNS resolution across multiple OPNsense instances
Tutorial: Setting Up Query Forwarding
- Log into your OPNsense web interface
- Navigate to Services > Unbound DNS > Query Forwarding
- Click the “+” button to add a new forwarding rule
- Configure the following settings:
  - Enabled: Check this box
  - Domain: Leave blank to forward all queries (or specify a domain for selective forwarding)
  - Server IP: Enter the IP address of your upstream DNS server (e.g., your primary OPNsense router’s IP)
  - Server Port: Use 53 for standard DNS or 853 for DNS over TLS
  - Forward TCP upstream: Check if you want to use TCP for upstream queries
- Click “Save”
- Click “Apply” to activate the changes
Example Configuration
For an OPNsense router behind another OPNsense router:
- Domain: (leave blank)
- Server IP: 192.168.1.1 (IP of primary OPNsense router)
- Server Port: 53
- Forward TCP upstream: Checked
Additional Tips
- If using DNS over TLS, ensure the upstream server supports it and use port 853
- For specific domain forwarding, enter the domain in the “Domain” field (e.g., “home.arpa”)
- You can add multiple forwarding rules for different domains or upstream servers
By implementing query forwarding, you ensure that DNS queries from devices behind your secondary OPNsense router are properly resolved, either locally or by forwarding to the primary router. This maintains consistent name resolution across your entire network infrastructure.
(Deprecated) Unbound DNS Overrides
Dont do this. Do the method above. This method is now officially deprecated. It is only here for reference.
Overview
We’ll set up a system where:
- Each site has its own OPNsense router with Unbound DNS.
- Each site can resolve hostnames for all other sites.
- DNS queries for remote sites are forwarded through the VPN.
Let’s assume we have three sites:
- Site 1: 10.1.0.0/24 (VPN IP: 10.0.0.1)
- Site 2: 10.2.0.0/24 (VPN IP: 10.0.0.2)
- Site 3: 10.3.0.0/24 (VPN IP: 10.0.0.3)
Step 1: Configure Unbound DNS on Each OPNsense Router
On each OPNsense router:
- Navigate to Services > Unbound DNS > General
- Check “Enable Unbound”
- Set “Network Interfaces” to LAN and VPN interfaces
- Enable DNSSEC
- Check “Register DHCP leases” and “Register DHCP static mappings”
- Click “Apply”
Step 2: Set Up Domain Overrides
On each OPNsense router, set up domain overrides for the other sites. This tells Unbound to forward queries for specific domains to other DNS servers.
For Site 1 OPNsense:
- Go to Services > Unbound DNS > Overrides
- Add two overrides: a. For Site 2:
  - Domain: site2.domain.com
  - IP Address: 10.0.0.2 (Site 2’s VPN IP)
  - Description: Forward to Site 2 DNS b. For Site 3:
  - Domain: site3.domain.com
  - IP Address: 10.0.0.3 (Site 3’s VPN IP)
  - Description: Forward to Site 3 DNS
Repeat this process on Site 2 and Site 3 OPNsense routers, adjusting the domains and IP addresses accordingly.
Step 3: Configure Host Overrides for Local Applications
On each OPNsense router, add host overrides for local applications.
For Site 1 OPNsense:
- Go to Services > Unbound DNS > Overrides
- Add a host override:
  - Host: app
  - Domain: site1.domain.com
  - Type: A
  - IP: 10.1.0.10 (local IP of the app server)
  - Description: Site 1 App Server
Repeat for Site 2 (anotherapp.site2.domain.com) and Site 3 (someotherapp.site3.domain.com), using their respective local IPs.
Step 4: Configure Firewall Rules
On each OPNsense router:
- Go to Firewall > Rules
- On the VPN interface, add a rule:
  - Action: Pass
  - Interface: VPN
  - Protocol: TCP/UDP
  - Source: VPN net
  - Destination: This Firewall
  - Destination port range: DNS (53)
  - Description: Allow DNS queries over VPN
Step 5: Configure DHCP to Use Local DNS
On each OPNsense router:
- Go to Services > DHCPv4 > [LAN interface]
- In “DNS servers”, enter the local IP of the OPNsense router (e.g., 10.1.0.1 for Site 1)
- Click “Save” and “Apply Changes”
Future Updates - help
If any of this is out of date, send me a pull request. I would love the chance to update this if there’s a change.
At the original time of writing OPNsense is at version 24.1.4
The updates that appear were done for OPNsense at version 24.7.1
Having trouble?
For a one-time donation you can get one-on-one troubleshooting support for any of my guides/projects. I’ll help you fix any issue you may have encountered regarding usage/deployment of one of my guides or projects. More info in my Github Sponsors profile.
