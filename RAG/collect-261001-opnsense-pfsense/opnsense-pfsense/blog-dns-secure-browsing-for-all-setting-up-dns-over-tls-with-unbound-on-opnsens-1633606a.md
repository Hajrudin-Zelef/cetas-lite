---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-dns-secure-browsing-for-all-setting-up-dns-over-tls-with-unbound-on-opnsens-1633606a
title: "blog-dns-secure-browsing-for-all-setting-up-dns-over-tls-with-unbound-on-opnsens-1633606a"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/blog-dns-secure-browsing-for-all-setting-up-dns-over-tls-with-unbound-on-opnsens-1633606a.md
source_anchor: ""
source_lines: [1, 58]
sha256: 5574dac139c370f0d73859c311f81ec3ae88d0913bdf08778abf36b042363c46
---

# blog-dns-secure-browsing-for-all-setting-up-dns-over-tls-with-unbound-on-opnsens-1633606a

Zero Trust can fail long before an attacker appears. Treating...
Read More
DNS is the phone book of the internet, but traditional DNS requests aren’t encrypted, leaving them open to eavesdropping or manipulation. That’s where DNS over TLS (DoT) comes in, encrypting DNS queries for enhanced privacy and security.
If you use OPNsense, a powerful open source firewall, you can enable DNS over TLS with the built in Unbound DNS resolver. Even better: by leveraging Cloudflare’s DNS blacklists, you can block malware and filter content to keep your family or organization safe.
This guide walks you step by step through configuring DNS over TLS on OPNsense using Unbound, and explores how Cloudflare’s free DNS filtering adds another crucial security layer.
DNS over TLS on OPNsense - Why Use DNS Over TLS (DoT)?
- Privacy: Your ISP or network snoops can’t see which sites you’re visiting.
- Security: Prevents DNS spoofing and man in the middle attacks.
- Peace of Mind: When paired with threat blocking lists, you reduce risks from malware, phishing, and inappropriate content.
DNS over TLS on OPNsense - Prerequisites
- OPNsense firewall (latest version recommended)
- Unbound DNS enabled (default on OPNsense)
- Admin access to the OPNsense web interface
Step-by-Step: Configuring DNS Over TLS on OPNsense
Enable Unbound DNS (if not already active)
- Go to Services → Unbound DNS → General
- Ensure “Enable Unbound” is checked
Configure DNS Over TLS Servers
- Still in the Unbound DNS settings, find the DNS over TLS section
- Enter Cloudflare’s DNS servers: 
  - 1.1.1.1 and 1.0.0.1 (for standard privacy)
  - For family filtering: 1.1.1.3 and 1.0.0.3 (Cloudflare’s malware and adult content blocking)
- Set the appropriate port (usually 853 for DoT)
Adjust Advanced Settings
- Go to the Advanced tab
- Add or ensure these options are set: 
  - Enable DNSSEC Support (for authenticating responses)
  - TLS Upstream (ensures upstream queries are encrypted)
- Save and apply changes
Update Firewall Rules
- Ensure outbound traffic on port 853 (TCP) is allowed for Unbound to communicate with external DNS servers
Test Your Configuration
- Use tools like Cloudflare’s Browsing Experience Security Check or dig anddrill commands to confirm encrypted DNS queries
About Cloudflare DNS Blacklists
Cloudflare offers free DNS filtering to block malware and protect families or businesses from unwanted content:
- 1.1.1.2 / 1.0.0.2: Blocks malware and known threats
- 1.1.1.3 / 1.0.0.3: Blocks both malware and adult content
By using these servers with DNS over TLS, OPNsense users can combine privacy, security, and content filtering in one step, no need for extra plugins or paid subscriptions.
DNS over TLS on OPNsense - Best Practices & Final Tips
- Keep OPNsense and Unbound updated for security patches
- Periodically test that DNS queries remain encrypted
- Adjust Cloudflare filtering as needed (malware only vs. malware + adult content)
- Consider DNSSEC for added integrity
Call to Action
Have you secured your home or business network with DNS over TLS?
What DNS filtering tools do you use?
Share your tips or questions in the comments, and subscribe for more practical security guides!
Zero Trust vs Traditional Security: Full Breakdown
Traditional security trusted the internal network. Zero Trust challenges that...
Read More
Smaller, Tougher and More Efficient: Why New Radiation Tolerant Computer Chips Matter
Space systems need more computing power, but radiation can flip...
Read More
Blocking QUIC for Advertisement Filtering in OPNsense While Allowing Facebook Mobile
Blocking QUIC can support a wider advertisement filtering policy, but...
Read More
1 reply on “Secure Browsing for All: Setting Up DNS Over TLS with Unbound on OPNsense (and Why Cloudflare’s DNS Blacklists Matter)”
[…] Secure Browsing for All: Setting Up DNS Over TLS with Unbound on OPNsense and Why Cloudflare’s DNS… — Configure encrypted DNS and filtering through Unbound on OPNsense. […]
