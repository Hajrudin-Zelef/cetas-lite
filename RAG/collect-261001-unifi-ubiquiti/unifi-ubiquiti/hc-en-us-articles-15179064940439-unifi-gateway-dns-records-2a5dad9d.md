---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-15179064940439-unifi-gateway-dns-records-2a5dad9d
title: "hc-en-us-articles-15179064940439-unifi-gateway-dns-records-2a5dad9d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-15179064940439-unifi-gateway-dns-records-2a5dad9d.md
source_anchor: ""
source_lines: [1, 43]
sha256: ef0f6da2064856e77ea12f382370aef4d95c3b317ca80fbffc578d825afe49d2
---

# hc-en-us-articles-15179064940439-unifi-gateway-dns-records-2a5dad9d

UniFi DNS Records and Local Hostnames
UniFi Gateways give you powerful tools to manage DNS resolution within your network—whether you’re creating custom DNS records, assigning local hostnames, or integrating with a dedicated internal DNS server. These features allow you to support internal services, simplify access to local resources, and maintain control over how clients resolve domains.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Creating DNS Records
To create DNS records on your UniFi gateway:
- Navigate to DNS Records:
  - Network 9.4: Settings > Policy Table > Create New Policy > DNS
  - Network 9.3: Settings > Policy Engine > DNS > Create DNS Record
- Select the Type of record:
  - Host (A): Maps a domain name to an IPv4 address.
  - Host (AAAA): Maps a domain name to an IPv6 address.
  - Alias (CNAME): Aliases one domain name to another domain name.
  - Mail (MX): Specifies the mail server for a domain.
  - Text (TXT): Stores human-readable data like SPF/DKIM values.
  - Service (SRV): Defines service-specific details (e.g., LDAP, SIP).
  - 
Forward Domain: Redirects DNS queries for a domain to another DNS server.
    - Note: The UniFi Gateway needs to be able to reach the specified DNS server. If the DNS server is reachable over a VPN Client connection, add a static route for the DNS server's IP address / subnet with the VPN server's IP address or VPN Client interface as the next hop.
- Specify the Domain Name, and IP Address.
- (Optional) Manually specify the TTL.
- Click Add.
Note: CNAME records are not yet supported, but are supported with UniFi OS 4.3+ and Network 9.3+.
Creating Local Hostnames in UniFi
If your goal is to assign a simple, local hostname to a known client device (e.g., printer.local or camera.lab), UniFi allows you to do this directly from the Client Devices page. This is essentially a shortcut for creating a Host (A) record tied to a specific IP.
To do so:
- Go to Client Devices.
- Select the device and open its Settings.
- Enable Fixed IP Address to set a DHCP reservation for this or another IP.
- Enable Local DNS Record and enter a hostname.
- Click Apply Changes.
Remember: These hostnames are stored in the Gateway’s DNS cache and only resolve for clients using the Gateway as their DNS server.
Note: For clients with fixed IP addresses outside of DHCP scope, we recommend creating a Host (A) record separately.
Assigning a Local DNS Server via DHCP
If you’re using a dedicated internal DNS server—such as for Active Directory—you can configure UniFi to direct clients to it automatically via DHCP:
- Go to Settings > Networks.
- Select the network that you want to use the DNS server.
- Scroll down to DNS Server and uncheck the toggle.
- Enter the IP address of your internal DNS server.
- Click Apply Changes.
Local DNS Servers With Content Filtering
If you’ve assigned a custom local DNS server and are using Content or Domain Filtering, additional steps are required to preserve both policy enforcement and local DNS resolution. By default, content filtering redirects DNS traffic to the UniFi gateway for inspection—this can bypass custom local DNS entries unless properly configured.
For full setup guidance, see Support for Local DNS Resolution in Content and Domain Filtering.
