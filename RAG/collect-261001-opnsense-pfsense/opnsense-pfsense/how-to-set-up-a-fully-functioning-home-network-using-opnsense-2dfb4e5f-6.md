---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-6
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [330, 386]
sha256: 244c73a994710e71422d87a85ddae86e08f6d5a3c1d76e43d6b7a9caba994484
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

Although I have learned a lot about how to configure IPv6 in OPNsense, I still do not consider myself to be an IPv6 expert, which is why I am providing a more general IPv6 configuration.
Go to each interface’s page under the “Services > Router Advertisements” section. The first option is “Router Advertisements”. Simply choose “Assisted” and click “Save”. This step is the same for all of the local networks.
DNS Configuration
I think configuring the DNS options in OPNsense can be a bit confusing for new users (I struggled at first too) primarily because there are a couple of places where you may specify DNS information. There are various approaches to how you may configure DNS so depending on the approach taken is where you need to enter the DNS configuration.
You may simply use the ISP’s DNS servers (that is the default DNS configuration in OPNsense), external DNS over TLS (DoT) servers, or another internal DNS server such as Pi-hole or AdGuard Home (which then may use external DoT/DoH servers). In this guide, I am going to configure Unbound DNS as the primary local DNS server which then uses DNS over TLS as the upstream, external DNS servers.
As I briefly mentioned earlier in this guide, I am going to assume you have no DNS servers configured on the “System > Settings > General " page and that you have the 2 “DNS server options” unchecked (to ensure the ISP DNS servers are not being used as the upstream/external DNS servers). Since DNS over TLS is going to be configured in the Unbound DNS settings, the system DNS servers do not need to be specified.
Unbound DNS: General
On the “Services > Unbound DNS > General” page, set the following configuration values:
| Option | Value | 
|---|---|
| Enable | Check “Enable Unbound” (if not enabled already) | 
| Listen Port | Leave as default 53 | 
| Network Interfaces | Choose “All (recommended)” (should be default) | 
| DNSSEC | Check “Enable DNSSEC Support” | 
| DHCP Registration | Check “Register DHCP leases” (to use hostnames of DHCP clients) | 
| DHCP Static Mappings | Check “Register DHCP static mappings” (to use hostnames of static DHCP clients) | 
| IPv6 Link-local | Check “Register IPv6 link-local addresses” | 
| DNS Cache | Check “Flush DNS cache during reload” (to clear the cache after making changes to Unbound) | 
| Local Zone Type | transparent (the default value) | 
Click the “Save” button at the bottom of the page and then click the “Apply changes” button at the top of the page to reload the Unbound service to apply configuration changes.
Unbound DNS: Advanced
You may want to set a few of the “advanced” options especially if you are interested in looking at DNS logs in more detail.
| Option | Value | 
|---|---|
| Prefetch Support | Checked (increases DNS queries and CPU load slightly so consider your hardware resources) | 
| Prefetch DNS Key Support | Checked (increases CPU load slightly so consider your hardware resources) | 
| Log Queries | Checked (if you want to troubleshoot DNS lookups or verify DoT is working by seeing DNS servers/ports being queried) | 
| Log Level Verbosity | Level 2 (to see more log detail such as the DNS servers being used which helps verify DoT configuration) | 
For the logging of queries and the log levels, you may only want to temporarily change these settings to verify DNS is operating as expected since this can be slightly detrimental to performance on slower hardware. It would also add more wear to the disks due to increased logging, which may be important depending on the hardware you are using. I find the DNS query logging useful when verifying DNS over TLS configurations since I can see that it is reaching out to the Cloudflare DNS servers on port 853, which is the port used by DoT for encrypted DNS queries.
Click “Apply” to persist the changes.
Unbound DNS: DNS over TLS
I prefer to use encrypted DNS queries for all outbound DNS queries since it is designed to not only enhance the privacy by limiting tracking of your DNS lookups, but it also offers improved security by reducing the likelihood your DNS queries will be hijacked via man in the middle attack. With a DNS over TLS (DoT) connection, you essentially have an encrypted tunnel to the configured DNS servers. I will be using Cloudflare in my example.
Navigate to the “Services > Unbound DNS > DNS over TLS” page to begin configuring DoT. Click on the “+” button to add a new DoT server. Cloudflare provides several different servers some of which provide additional filtering. Since I am using Zenarmor, I prefer to use the unfiltered Cloudflare servers because it is one less area I need to troubleshoot in case something is inadvertently blocked. The visibility of what Cloudflare blocks is not seen in OPNsense because Cloudflare does the filtering on their end unlike Zenarmor, Pi-hole, or AdGuard where you can see the queries which are being blocked.
Below is the list of DNS over TLS servers you may use for both IPv4 and IPv6. I recommend you include at least 2 DNS servers for IPv4 and at least 2 DNS servers for IPv6 for redundancy:
| DNS Filter | Server IP | Server Port | Verify CN | 
|---|---|---|---|
| No filter | 1.1.1.1 | 853 | cloudflare-dns.com | 
| No filter | 1.0.0.1 | 853 | cloudflare-dns.com | 
| No filter | 2606:4700:4700::1111 | 853 | cloudflare-dns.com | 
| No filter | 2606:4700:4700::1001 | 853 | cloudflare-dns.com | 
| Malware | 1.1.1.2 | 853 | security.cloudflare-dns.com | 
| Malware | 1.0.0.2 | 853 | security.cloudflare-dns.com | 
| Malware | 2606:4700:4700::1112 | 853 | security.cloudflare-dns.com | 
| Malware | 2606:4700:4700::1002 | 853 | security.cloudflare-dns.com | 
| Malware + Adult | 1.1.1.3 | 853 | family.cloudflare-dns.com | 
| Malware + Adult | 1.0.0.3 | 853 | family.cloudflare-dns.com | 
| Malware + Adult | 2606:4700:4700::1113 | 853 | family.cloudflare-dns.com | 
| Malware + Adult | 2606:4700:4700::1003 | 853 | family.cloudflare-dns.com | 
Use the information above to fill in for each DoT server you wish to use. You may leave the “Domain” box empty so all lookups go through the same DoT servers since that option is only used if you want a specific domain name (such as homenetworkguy.com) to use a specific DoT server when performing DNS lookups.
Firewall Configuration
The time has come to create firewall rules. Firewall rules are critical for providing increased security among the devices in your network. Having a solid understanding in this area will be crucial in helping you lock down your network tighter.
As you likely know, no software or hardware is fully impenetrable, which is why it is important to have several layers of defense when protecting your network.
I decided to include several VLANs in this guide in order to demonstrate different types of networks you may want to implement in your home network. They should provide some good use cases for how the firewall rules may look different for the types of access desired. Of course, you do not need to implement all of these VLANs in your network.
Note
I will be making use of firewall aliases in the firewall rules below, so if you see a name instead of an IP address for the “Source” or “Destination” it means I am using either a built-in firewall alias or the custom firewall aliases I describe in the “Firewall: Aliases” section.
The names in the firewall rules are not hostnames of devices on the network because hostnames can only be used in firewall aliases. Firewall rules only allow you to enter a single IP/network address or a single firewall alias (aliases may contain more than one value). If you want to use multiple IP addresses or network addresses in a single firewall rule, you have to create an alias containing those addresses and use that alias in the firewall rule.
Firewall: Aliases
