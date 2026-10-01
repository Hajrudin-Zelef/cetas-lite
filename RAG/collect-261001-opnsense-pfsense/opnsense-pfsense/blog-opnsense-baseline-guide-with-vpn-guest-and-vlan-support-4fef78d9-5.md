---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9-5
title: "blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9.md
source_anchor: ""
source_lines: [599, 726]
sha256: fc7426d35e6a2375c993f3bd25f4c36fb269321c36d72b4cc8841a7c4a37970d
---

# blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9

Finally, we enable unrestricted internet access on Guest networks.
| Action | Pass | 
| Interface | VLAN40_GUEST | 
| Protocol | TCP/UDP | 
| Source | VLAN40_GUEST net | 
| Destination / Invert | checked | 
| Destination | IG_LOCAL net | 
| Description | Unrestricted internet access | 
LAN Network For Testing And Debugging
I just keep the pre-defined “LAN to any” rules. I periodically reconfigure this network for testing and debugging and don’t use it for anything else.
Redirect Outbound DNS Traffic
To prevent clients from explicitly querying outbound DNS and leaking information to the outside, we redirect any outbound DNS traffic to Unbound or Dnsmasq.
Navigate to Firewall → NAT → Port Forward and add the following rules.
| Interface | IG_DNS_FORWARD | 
| Protocol | TCP/UDP | 
| Source | IG_DNS_FORWARD net | 
| Destination | any | 
| Destination port range | DNS | 
| Redirect target IP | 127.0.0.1 | 
| Redirect target port | 5335 | 
| Description | Redirect any DNS traffic to Dnsmasq | 
| Interface | IG_DNS_RESOLVE | 
| Protocol | TCP/UDP | 
| Source | IG_DNS_RESOLVE net | 
| Destination / Invert | checked | 
| Destination | IG_DNS_RESOLVE net | 
| Destination port range | DNS | 
| Redirect target IP | 127.0.0.1 | 
| Redirect target port | DNS | 
| Description | Redirect outbound DNS traffic to Unbound | 
Redirect Outbound NTP Traffic
To sync the time of all our devices on the network to OPNsense, we redirect all NTP traffic.
Navigate to Firewall → NAT → Port Forward and add the following rule.
| Interface | IG_NTP | 
| Protocol | UDP | 
| Source | IG_NTP net | 
| Destination / Invert | checked | 
| Destination | IG_NTP net | 
| Destination port range | NTP | 
| Redirect target IP | 127.0.0.1 | 
| Redirect target port | NTP | 
| Description | Redirect outbound NTP traffic to OPNsense | 
Test
Now would be a could time to reboot OPNsense to make sure all settings are applied.
Test DHCP
Connect to a host in each VLAN and verify it receives an IP inside the specified
DHCP range. Here is the output of the ip -4 addr show eth0 command from a
Ubuntu host connected to the VPN VLAN.
Test DNS
We have to verify the following functionality of our DNS architecture:
- VLAN20_VPN
  - Unbound resolves remote and local hostname lookups
  - Redirect outbound DNS traffic to Unbound
  - Reverse lookups of private IPs
  - Don’t leak lookups for the private corp.example.com domain
- VL30_CLEAR
  - Dnsmasq forwards remote hostname lookups to the system DNS servers like Quad9 and Unbound
  - Forward local hostname lookups to Unbound
  - Redirect outbound DNS traffic to Dnsmasq
  - Forward local reverse lookups of private IPs to Unbound
  - Don’t leak lookups for the private corp.example.com domain and forward
them to Unbound
- VL40_GUEST
  - Use external DNS resolvers
  - Allow for clients to override DNS
  - OPNsense lookups are blocked
We’ll use the dig tool and the firewall logs under
Firewall → Log Files → Live View for testing.
I’ll also skip the Management network because it requires the same testing as the VPN network.
VLAN20_VPN: Test DNS
Connect to VLAN20_VPN.
VLAN20_VPN: Remote Hostname Lookups
Run dig california.gov:
Here are the firewall logs showing the iterative DNS requests Unbound sends.
VLAN20_VPN: Local Hostname Lookups
Run dig opnsense.corp.example.com:
VLAN20_VPN: Redirect Outbound DNS Traffic
Run dig opnsense.org @8.8.8.8:
dig can’t tell that OPNsense hijacked the request and thus displays an
incorrect SERVER value. If you check the firewall logs, you shouldn’t see any
requests to 8.8.8.8. Instead, you should see iterative root server requests.
VLAN20_VPN: Reverse Lookups of Private IPs
Run dig -x 192.168.20.1:
If you want, additionally reverse-lookup an IP that doesn’t exist. The firewall logs mustn’t contain requests to external DNS servers.
VLAN20_VPN: Verify corp.example.com Is Private
To test whether OPNsense is the authoritative server for corp.example.com, we
lookup a non-existent hostname in that domain. dig should return an
authoritative NXDOMAIN response with the SOA record we earlier defined earlier
in the AUTHORITY SECTION.
Run dig nowhere.corp.example.com:
VLAN20_VPN: DNS Leak Test
In your browser, navigate to dnsleaktest.com or mullvad.net/check. We expect the “leaked” DNS server to match our Mullvad public Mullvad IP. The second leak is from the Outgoing Interface we configured for Unbound:
VLAN30_CLEAR: Test DNS
Connect to VLAN30_CLEAR.
VLAN30_CLEAR: Remote Hostname Lookups
Run dig opnsense.org:
Check the firewall logs. Enable logging for the port forward rule if you want it to show up.
You can see that Dnsmasq forwards to the DNS servers defined under System → Settings → General and Unbound.
VLAN30_CLEAR: Local Hostname Lookups
Run dig opnsense.corp.example.com:
VLAN30_CLEAR: Redirect Outbound DNS Traffic
Run dig opnsense.org @8.8.8.8:
We confirm it works by looking at the firewall logs again:
VLAN30_CLEAR: Forward Reverse Lookups of Private IPs to Unbound
Run dig -x 192.168.20.1:
The firewall logs confirm it works.
VLAN30_CLEAR: Verify corp.example.com Is Private
Run dig nowhere.corp.example.com:
This time, requests will only be forwarded to Unbound, but not external DNS resolvers.
VLAN30_CLEAR: DNS Leak Test
As we saw earlier, we expect the Quad9 and the Mullvad public IPs to leak. Here is the result of an extended test from dnsleaktest.com:
VLAN40_GUEST: Test DNS
Connect to VLAN40_GUEST.
Verify that dig opnsense.org @192.168.40.1 times out.
The Cloudflare DNS servers you configured in the DHCP settings of the Guest VLAN should show up when running the leak test:
Thanks For Reading ❤️
If you’re here, I thank you for reading all this! Any feedback is highly appreciated.
When I decided to write this guide, I didn’t think it would take soooo long. I worked on it intensively for the better part of a month. I spent like a week configuring Unbound to use WireGuard tunnels, only to find out that I couldn’t get it working due to a bug. But it was all worth it and a fulfilling journey — I have learned so much OPNsense and networking. But I’m also happy to be able to put this aside for a while. 🎉
So what’s next?
Let’s Encrypt certificates and HAProxy to secure self-hosted services. I imagine configuring it should be pretty straightforward.
Me and possibly others want to be able to access my home network from the
outside via WireGuard. I have a dynamic IP, so I thought I’d have to resort to
Dynamic DNS. But I think
port forwarding with Mullvad
is the better solution and doesn’t require me to associate my public IP address
with a public DNS record.
Traffic shaping and intrusion prevention is something I want to look into, too.
So yeah, OPNsense and I will be friends for a while. 👫
