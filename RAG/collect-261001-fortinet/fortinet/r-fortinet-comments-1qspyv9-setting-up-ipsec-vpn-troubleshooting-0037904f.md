---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1qspyv9-setting-up-ipsec-vpn-troubleshooting-0037904f
title: "r-fortinet-comments-1qspyv9-setting-up-ipsec-vpn-troubleshooting-0037904f"
domain: fortinet
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["cost", "incident"]
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1qspyv9-setting-up-ipsec-vpn-troubleshooting-0037904f.md
source_anchor: ""
source_lines: [1, 40]
sha256: 699e55f0f50632c0b0890f8413022b7c8f9ef515e279a1f561ee01e51776e304
---

# r-fortinet-comments-1qspyv9-setting-up-ipsec-vpn-troubleshooting-0037904f

Setting up IPsec VPN & troubleshooting 
        
    Let's say you need to setup a VPN option via a FortiGate firewall.
The SSL VPN path is easy, but if you don't have static IPs to narrow down the allowed remote sites, it can get abused pretty hard and the SSL service reloads when overwhelmed (and all current sessions dropped).
The next idea is for an IPsec option, which leads you to needing to set an authentication source [Local, LDAP, RADIUS, SAML].
- Local sounds pretty good, but then there's the administrative cost of maintaining user accounts, and combating drift.
- LDAP is an oldie and a goodie, but comes with some problems. Like the free FortiClient (non-EMS) client missing the EAP-TTLS configuration to credential against LDAP. If you want to use IKEv2 for the tunnel, you are going to be dealing with this. If you can edit the registry, you can work around this, but Android & IOS FortiClient support varies, so this is might be rough too.
- RADIUS might be your best path, as it plays nice with all FortiClient VPN clients, and reduces Admin cost.
- SSO/SAML authenticating to your cloud provider, like Microsoft 365 or FortiCloud could be great for ease of administration and centralizing where authentication occurs, but Fortinet is currenting working to make it secure due to incident that have taken place.
If you decided that your best option was IPsec VPN using RADIUS authentication & IKEv2 for the VPN tunnel then you are probably in the majority.
Some things I would recommend you stay aware of.
- IPsec VPN using the default UDP (500) port is not like SSL VPN, Customer Premise Equipment (CPE) or ISPs may block that traffic, so TEST TEST TEST to verify that.
- If you want to change the port you are listening for the traffic (like 501, instead of 500), do it at the beginning because the FortiGate views this as a global setting. If you change it, it will impact all IPsec VPN tunnels, not just the new one. Site-to-Site, Dialup groups. All of them.
- When you configure the IPsec Dial-up VPN tunnel on the firewall, make sure you set a Peer ID (called 'Local ID' on the FortiClient) using alphanumeric to identify the Dial-up group and set the "Peer ID" to Specific. If you don't, the client will try to authenticate to the first IPsec tunnel, regardless if it's the right one.
- There are more consistent and probably severe things that I am not remembering, but hopefully others will chime in with their $0.02 for helping build out a VPN on Fortinet hardware.
Section des commentaires
A few things to note (IMO):
FortiCloud vuln doesn't affect other SSO providers
Even then, FortiCloud vuln is talking about exposed admin ports, not VPN
IPsec TCP on port 443 as a backup is a suitable workaround for the UDP 500 port problem
I always recommend using EMS these days
I appreciate the insight! IPsec via TCP on 443 does sound like a good workaround. Is that also a global setting that will impact other IPsec tunnels on that firewall?
I’m interested as well. A response I have from a Fortinet professional is that my current IPsec tunnels will not be changed.
The setting will only affect the tunnels that we change/set the TCP port. The current tunnels are all using UDP ports and we are not changing those. I reached out to my contacts at Fortinet for verification and this is their reply.
The cmd below is a global setting and it will take effect for all tunnels that are set (inside that particular tunnel settings) with the Set Transport to TCP. But it will only affect those tunnels.
Here's a useful guide -
https://community.fortinet.com/t5/FortiGate/Technical-Tip-How-to-use-TCP-as-transport-for-IKE-IPsec-traffic/ta-p/300834
The option to use TCP is per-tunnel, however the custom TCP port used for any such tunnels is a global setting.
For IKEv2 you should use network ID.
Apart from that I don't see any issues with using saml.
Is there a benefit of using Network ID over Peer ID/Local ID for the Phase1-interface?
Network ID is part of the v2 negotiation and recommended for v2 tunnels
https://docs.fortinet.com/document/fortigate/7.6.0/ssl-vpn-to-ipsec-vpn-migration/690046/customizing-ipsec-tunnel-settings#:~:text=However%2C%20peer%20ID%20functionality%20is,configured%20on%20different%20IPsec%20tunnels
Isn't this only if you have multiple remote access VPN tunnels on the same interface (i.e. site-to-site VPNs are not impacted by this?)
I can tell from what testing I've done that this is going to be a problem. What's the workaround for this? That was one benefit of SSL VPN, it could punch through restrictive firewalls or crappy hotel or conference wi-fi that is restrictive on outbound connections.
I see there's a "Fall back to 443" option in FortiClient, anyone set that up?
That is going to be helpful, I've seen a couple places that are blocking just in testing. It won't be a problem for the remote users I have in their spec-ed roles but sooner than later someone is going to need connectivity at a convention or another restrictive firewall.
Knowing I can configure for a fallback port using TCP will be a proactive way to save the day.
set transport tcp
It works. FortiClient will first try connecting via UDP and after a timeout fall back to whatever port you have configured for TCP.
