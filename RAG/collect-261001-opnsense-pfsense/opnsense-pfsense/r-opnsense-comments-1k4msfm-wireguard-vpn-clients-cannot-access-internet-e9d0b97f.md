---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-1k4msfm-wireguard-vpn-clients-cannot-access-internet-e9d0b97f
title: "r-opnsense-comments-1k4msfm-wireguard-vpn-clients-cannot-access-internet-e9d0b97f"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-1k4msfm-wireguard-vpn-clients-cannot-access-internet-e9d0b97f.md
source_anchor: ""
source_lines: [1, 124]
sha256: 853ff74d1db84911b7099c37467fbefda95a3634e43caef710f0cf811b428d46
---

# 
       WireGuard VPN Clients Cannot Access Internet (Quantum Fiber in Bridge Mode) 

      
    Hello OPNsense Community,

I'm experiencing an issue where clients connected to my WireGuard VPN server on OPNsense cannot access the internet. My setup involves:

- 
    **Internet:** Quantum Fiber, their provided modem/ONT is configured in transparent bridge mode.
- 
    **OPNsense:** Running the latest stable version. My WAN interface is receiving a DHCP address from the bridged modem.
- 
    **LAN:** Standard`192.168.1.0/24` network for local devices (which have full internet access).
- 
    **WireGuard VPN:** Server configured on OPNsense with the`wg1` interface, using the`10.0.0.0/24` subnet for clients. The server's tunnel address is`10.0.0.1/24` . "Disable routes auto-add" is unchecked.
- 
    **VPN Client (Example):** My laptop is configured with the address`10.0.0.3/32` and DNS server`192.168.1.1` . Allowed IPs are`192.168.1.0/24, 10.0.0.0/24, 0.0.0.0/0` . The VPN connection shows as active.
- 
    **DNS:** I have AdGuard Home running on OPNsense (`192.168.1.1` ), listening on the standard DNS port. It is configured to forward queries to Unbound, also running on OPNsense (listening on port 53530). Unbound has Cloudflare (`1.1.1.1` ,`1.0.0.1` ) and Google (`8.8.8.8` ) DNS servers configured as forwarders. I have tried disabling DNSSEC and "Agressive NSEC" in Unbound. I have also tried setting the system DNS servers in OPNsense (System > Settings > General) directly to`1.1.1.1` and`1.0.0.1` with "Allow DNS server list to be overridden by DHCP/PPP/RADVD on WAN" unchecked.
- 
    **Firewall Rules:**  - 
    **WG1:** A "pass all" rule is in place for IPv4 from`wg1 net` to any destination.
  - 
    **LAN:** Rules are in place to allow LAN clients internet access and to allow OPNsense to communicate with external DNS servers.
  - 
    **WAN:** I have reviewed the WAN rules and do not see any explicit block rules for outbound traffic on ports 80 or 443 originating from my WAN IP.
- 
    
- 
    **Outbound NAT:** A rule exists on the WAN interface with source`10.0.0.0/24` , protocol "any", source port "any", destination "any", destination port "any", NAT address "Interface address".

    **Problem:** While connected to the VPN, my laptop can resolve internal LAN addresses (e.g., ping `192.168.1.1`) and DNS queries appear to be reaching OPNsense (based on AdGuard Home logs when system DNS was set to `192.168.1.1`). However, I cannot access any websites (e.g., `cloudflare.com`). The browser indicates "address could not be found".
  

    **Troubleshooting Steps Taken:**
  

- 
    Verified Quantum Fiber modem is in transparent bridge mode.
- 
    Rebooted both the modem and OPNsense.
- 
    Checked firewall rules on all interfaces multiple times.
- 
    Confirmed Outbound NAT rule for the VPN subnet is in place.
- 
    Tried different DNS configurations (Unbound forwarders, direct system DNS).
- 
    Disabled DNSSEC and Agressive NSEC in Unbound.
- 
    Verified WireGuard server and client configurations.
- 
    Used the Firewall Live View to monitor traffic. I see traffic from the VPN client ( `10.0.0.3` ) going to`192.168.1.1:53` (DNS), but I do not see any traffic originating from`10.0.0.3` with a destination of public IPs on ports 80 or 443. Interestingly, I did see traffic on the LAN interface with the VPN client as the source and a public IP as the destination, which seems incorrect.

I am at a loss as to why internet traffic from my VPN clients is not reaching the internet. Any insights or suggestions for further troubleshooting would be greatly appreciated.

    Thank you in advance for your help! <sup>1</sup> 
  

Where did you configure this?

    
    Allowed IPs are `192.168.1.0/24, 10.0.0.0/24, 0.0.0.0/0`
  


Your rulset on WG1 pass rule wg1 net to any; is wg1 net 10.0.0.0/24?

What does a traceroute say to 1.1.1.1? Can you only try with an allowed route of 0.0.0.0/0?

I configured the allowed IPs on the wireguard client configuration. On the Peer configuration, I have 192.168.1.0/24, 10.0.0.3/32 (the ip of the wireguard client when connected), and 0.0.0.0/24

wg1 net is 10.0.0.0/24, but I've also tried typing it out instead of using the alias to no avail.

What do the live logs say, if you do a traceroute? From 10.0.0.3 to 1.1.1.1 ?

And change Tunnel Address in your wireguard Instance to 10.0.0.0/24 instead of 10.0.0.1/24

Only try allowed IPs 0.0.0.0/0 -> remove 192.168.1.0/24 and 10.0.0.0/24 doesnt make sense, when you have 0.0.0.0/0 (everything) allowed

DNS config is fine

DNS = 192.168.1.1

Here are the live logs when i filter source = 10.0.0.3

I changed tunnel address to 10.0.0.0/24

I also left only 0.0.0.0/0 in Allowed IPs for both the client config and the Peer config in Opnsense

    I just tried navigating to cloudflare.com and it gives me `cloudflare.com’s server IP address could not be found.`
  

    EDIT:

For what its worth, here is my Unbound DNS log
  

What are you using as Upstream DNS Server in Unbound, is DNSSEC active? Could you try to disable it?

You need to make an adjustment to the config. In the Peer Generator’s DNS Servers field, specify the server IPs without CIDR notation.

In my wireguard client config, I have

DNS = 192.168.1.1

This should be fine, no?

The opnsense wireguard Address ip should be configured as 10.0.0.1/24,peer allowed ip is configured as 10.0.0.2/32 and it is allowed to route. Client Address ip is configured as 10.0.0.2/32, peer allowed ip is 0.0.0.0/0 (via opnsense internet) and 192.168.1.0/24 (opnsense LAN)

My usual guess when it comes to this is NAT.

Can you share your nat picture?

Also show a packet that was heading out, and click the lil i button, and paste a picture of that?

After hours of troubleshooting, I realized that the issue is not just limited to my wireguard clients, but to my whole LAN as well. I think there is some trouble with the way Opnsense is communicating with the Quantum Fiber network.

To test it, I took my Opnsense router to my work and tested it behind my work firewall. Besides the double NAT, it worked perfectly without issue. Then I took it back home and tried it there again, and I started getting "no route to host" for every traceroute that I tried, despite getting a public IP and gateway on my WAN interface.

Based on that, there must be something wrong with how it is communicating through the Quantum Fiber modem/ONT/router that is set to bridge mode. I think I'll make a post in r/centurylink and see what they think.

Here is a link to the traceroute.
