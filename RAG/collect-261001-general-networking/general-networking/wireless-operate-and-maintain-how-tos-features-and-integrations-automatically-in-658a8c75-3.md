---
id: collect-261001-general-networking/general-networking/wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75-3
title: "wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75.md
source_anchor: ""
source_lines: [388, 467]
sha256: 0aea5f3040bd8ff26267451b31878a3ededb8202fef1ccf72f866dcb142a20fa
---

# wireless-operate-and-maintain-how-tos-features-and-integrations-automatically-in-658a8c75

Enabling Automatic Umbrella Integration via Group Policy
- Navigate to Network-wide > Configure > Group Policies > Add a group
- Type a name for the Group Policy.
- Go to Firewall and traffic shaping section and select Custom SSID firewall and shaping rules.
4. Save changes on the bottom of the page.
- 
    Navigate back to the newly created policy.
- 
    Select Link Umbrella to [Group Policy name] under DNS layer protection (Cisco Umbrella).
8. Select a desired predefined policy from the dropdown.
9. Save changes on the bottom of the page
Note: Group Policies do not support umbrella DNS exclusions.
Applying Group Policy with Umbrella Protection to Wireless Clients
- Navigate to Network-wide > Monitor > Clients
- Click on the client and select Different policies by SSID
- Select your Group Policy name. Click Save under the Device policy drop-down.
- Click Save
Disabling Automatic Umbrella Integration in a Group Policy
- Navigate to Network-wide > Configure > Group Policies and select the desired group policy
- Click Disable Umbrella protection under DNS layer protection.
3. Click Yes in the dialog box to confirm that you wish to enable Umbrella integration.
4. Save changes on the bottom of the page
Umbrella DNS Protection Traffic Flow
This section describes the expected DNS traffic flow after Umbrella policies have been applied to a wireless client on either SSID-level or group policy-level.
- 
    A client sends a DNS query for the desired domain name (e.g., twitter.com).
- 
    Upstream MR access point intercepts the DNS query and attaches an Umbrella identifier to it, allowing Umbrella to determine which policy to enforce.
- 
    MR then encrypts the DNS query using DNSCrypt, source NATs the packet to the MR management IP, and redirects it to the appropriate Umbrella resolver.
- 
    Once received, the Umbrella resolver decrypts the DNS query and enforces the appropriate Umbrella policy (based on the attached identifier).
- 
    Umbrella returns an encrypted DNS response with the appropriate IP if the request is allowed per configured policy.
- 
    If the request should be blocked, Umbrella returns an encrypted DNS response pointing to the Umbrella block page IP address.
- 
    The client is sent to the desired domain name (e.g. twitter.com) or Umbrella Block page based on the applied policy.
Cisco Umbrella DNS filtering has the following general limitations:
- 
    If a client machine is programmed to reach out to an IP address of a remote server directly, a DNS-filtering solution like Umbrella will not prevent this communication, since there is no DNS query that MR can intercept.
- 
    If a client is using some form of end-to-end encryption (e.g. VPN solution) that encrypts traffic between the client and a remote server (including DNS queries), MR will not be able to intercept those queries and forward them to an Umbrella resolver.
DNSCrypt Compatibility
Access points that do not support 802.11ac, such as the MR18, will still be able to utilize Umbrella DNS services, but do not support the use of DNSCrypt when communicating to the Umbrella servers. All access points that are capable of 802.11ac or newer fully support the use of DNSCrypt with Umbrella DNS.
Blocking HTTPS Websites
Some websites (e.g., twitter.com, facebook.com, instagram.com, dropbox.com) have an HTTP Strict Transport Security (HSTS) security policy. This policy means that a website does not allow HTTP connections. So, for example, if a user types “dropbox.com” in the web browser address bar, which usually implies http://dropbox.com, the browser will automatically connect to https://dropbox.com.
As explained below, this presents a unique challenge for displaying an Umbrella block page.
If a user tries to visit an HTTPS website by typing "https://example.com," or if the website has an HSTS security policy (e.g., dropbox.com), the following will happen:
- 
    The client machine sends a DNS query for dropbox.com asking for an IP address.
- 
    Upstream MR access point intercepts the DNS query and attaches an identifier to it, allowing Umbrella to determine which policy to enforce.
- 
    MR then encrypts the DNS query using DNSCrypt, source NATs the packet to the MR management IP, and redirects it to the appropriate Umbrella resolver.
- 
    Once received, the Umbrella resolver decrypts the DNS query and enforces the appropriate Umbrella policy (based on the attached identifier).
- 
    Let’s say that “Full Appropriate Use Filtering” is applied to the SSID and, therefore, https://dropbox.com (File Storage) should be blocked.
- 
    Umbrella will return an encrypted DNS response pointing to the Umbrella block page IP address (e.g., 146.112.61.106). This response will be decrypted by the MR and sent to the client.
- 
    The client will try to establish a TLS session with 146.112.61.106 (Umbrella server) IP address received in the DNS response.
- 
    Umbrella’s block page presents an SSL certificate to browsers that make connections to HTTPS sites. The certificate will match the requested site name (Common Name - CN) but will be signed by the Cisco Umbrella Root Certificate Authority (CA). If this CA is not trusted by your browser, an error may be displayed. For example, you can see the TLS certificate presented by OpenDNS, Inc. for dropbox.com here:
As you can see, the main issue here is that “Cisco Umbrella Root CA” is not trusted.
- 
    Here is an example of what we would see in the Wireshark packet capture taken on the client machine:
10. Most modern browsers (like Chrome, Firefox, Safari) will prevent users from accessing a website with an untrusted/unexpected TLS certificate. Typical errors include: "Your connection is not private" (Chrome), "Did Not Connect: Potential Security Issue" (Firefox), or “Safari Can’t Open the page” (Safari). Although the error is expected, the messages displayed can confuse end-users.
We recommend installing Cisco Root Certificate on client machines to avoid these errors. This certificate can be installed per browser or device for personal use or small deployments. In addition, an automatic installation through GPO can be done for larger deployments. Note that the automated installation will only work for Internet Explorer, Edge, or Chrome users on Windows systems. If your network includes users on Firefox or Safari browsers, the manual installation procedures must be followed.
Once you have installed the Cisco Root Certificate, users will be presented with Umbrella Block Page even for HTTPS and HSTS websites.
.
NOTE: Cisco Umbrella's resolvers live at 208.67.222.222/32 and 208.67.220.220/32; Meraki sends DNS traffic to either one. Make sure any upstream devices allow bi-directional UDP 443 to these addresses.
NOTE: The instructions above presume that you have access to an existing Umbrella dashboard. If you do not have such access you can download the Cisco Umbrella Root CA certificate.
Using the Security Center to View MR DNS Events
The Meraki Security Center provides reporting functionality for MR DNS events for all networks in the organization. To view these reports, navigate to Organization > Monitor > Security Center > MR DNS Events.
It’s possible to search for a particular blocked website by adding "uri:" to the search string (e.g. “uri:exampleadult.com”) or a particular blocked client by adding "client:" to the search string (e.g. client:192.168.219.13).
You can also use the Filter option to filter DNS events by type (Content or Security) and/or action (Allowed or Blocked).
Note: Content categories are shown in gray in the Categories column while Security categories are shown in red.
Note: If the network where Umbrella is provisioned is bound to a configuration template, DNS events reported for that network will have the template name in the "Network" column.
