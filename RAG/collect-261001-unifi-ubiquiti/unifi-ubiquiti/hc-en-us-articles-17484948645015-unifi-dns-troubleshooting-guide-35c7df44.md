---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-17484948645015-unifi-dns-troubleshooting-guide-35c7df44
title: "hc-en-us-articles-17484948645015-unifi-dns-troubleshooting-guide-35c7df44"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-17484948645015-unifi-dns-troubleshooting-guide-35c7df44.md
source_anchor: ""
source_lines: [1, 17]
sha256: 5e2835b0353b77f4755900fac1c8be263da53b36cec680c890ce599f3c1abe49
---

# hc-en-us-articles-17484948645015-unifi-dns-troubleshooting-guide-35c7df44

UniFi - DNS Troubleshooting Guide
The DNS server is responsible for translating a url (ex., google.com) into the IP address where the host is located. There's a number of UniFi services (updating, remote management, etc) that require you to sufficiently resolve ui.com domains. Incorrect configuration could impact these services.
Check Firewalls & ISP Restrictions
DNS works using UDP Port 53. Ensure that this is not being blocked by any upstream firewalls, gateways or ISP modems. If it is, DNS resolution will fail.
Selecting a Reliable DNS Server
By default, UniFi will use the DNS Server provided by your Internet service provider. This is usually okay, but if you are unsure, we recommend using a public DNS Server such as 1.1.1.1 or 8.8.8.8.
We caution against the use of custom, internal DNS servers because they may not be updated to resolve the required ui.com domains, or they may block them entirely.
Re-Configuring Your DNS Server
UniFi Cloud Gateways
Navigate to UniFi Network > Settings > Internet > DNS Server and enter the new DNS Server.
CloudKeys, Network Video Recorders & Other Non-Gateway Consoles
Navigate to UniFi OS > Console Settings and check if the IP Configuration is set to DHCP or Static.
- If it is Static, enter the new DNS Server and select Apply Changes.
- If it is DHCP, you will need to modify the DNS Server directly from your DHCP server.
 
  - If you have a UniFi gateway, this is found in UniFi Network > Settings > Networks > [Network Name] > DHCP Service Management > DNS Server.
  - Note: This change will take effect after your device’s DHCP lease expires, usually after 24 hours. To expedite this, turn your device off and back on again.
