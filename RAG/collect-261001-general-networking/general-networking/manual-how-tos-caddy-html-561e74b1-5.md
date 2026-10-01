---
id: collect-261001-general-networking/general-networking/manual-how-tos-caddy-html-561e74b1-5
title: "Reverse Proxy Domain: \"531e7877-0b58-4f93-a9f0-54beee58bdea\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-how-tos-caddy-html-561e74b1.md
source_anchor: ""
source_lines: [496, 567]
sha256: 6449cd0be5960111a9726ed3af591013be06aa038c54fabdd07524ac40d94d20
---

# Reverse Proxy Domain: "531e7877-0b58-4f93-a9f0-54beee58bdea"

- Caddy upgrades all connections automatically from HTTP to HTTPS. When cookies do not have thesecure flag set by the application serving them, they can still be transmitted unencrypted before the connection is upgraded. If these cookies contain very sensitive information, it might be a good choice to close port 80.
- There is optional Layer4 TCP/UDP routing support. In the scope of this plugin, only traffic that looks like TLS and has SNI can be routed. The HTTP App and Layer4 App can work together at the same time.
- There is no WAF (Web Application Firewall) support in this plugin. For a business grade Reverse Proxy with WAF functionality, useos-OPNWAF .
Help, Nothing Works!
Note
Even though Caddy itself is quite easy to configure in the plugin, setting the infrastructure up correctly poses the real challenge. If you feel stumped, the best approach is knowledge about what should happen. This section tries to explain that and gives examples how to resolve issues.
Tip
Most errors happen because the infrastructure is not set up correctly, or wrong options for the Handler have been set.
Attention
Do not use the Layer4 module without knowing the implications of it. It is for very advanced usecases. Better deactivate it if things do not work as expected.
This is what should happen if Caddy works correctly:
- A Web Browser is opened and an URL is put into the address bar: https://example.com
- The underlying Operating System of the Web Browser sends a request to its default DNS Server, and asks where to find example.com. The DNS Server will try to find the requested A- and/or AAAA-Record for that domain, and will answer with e.g. 203.0.113.1.
- The Web Browser now sends a HTTPS request to 203.0.113.1. This request contains a Client Hello in the TLS handshake, that contains example.com.
- This HTTPS request hits port 443 of the OPNsense’s WAN, LAN (or VPN) interface, determined by the network location of the Web Browser.
- There is a Firewall rule that allows destination port 443 to access This Firewall. The request will then be received by Caddy, because it listens on This Firewall on port 443.
- In Caddy, there is a domain for example.com set up. It has a valid Let’s Encrypt or ZeroSSL certificate. Since the Client Hello contains example.com, Caddy will match it with the domain, and the Web Browser shows a certificate next to https://example.com in the address bar.
- Caddy takes the HTTPS request and terminates the TLS connection. That means, it will convert the HTTPS into HTTP, so it can be processed by the Handler.
- Caddy checks if there is a matching Handler set up. It will be used to reverse proxy the HTTP request to an internal service.
- Inside the Handler, the domain example.com and an Upstream Domain e.g. 192.168.10.1 and Upstream Port e.g. 8080 point the request to the internal service. Caddy then sends the HTTP request directly to the internal service.
- The HTTP response from the internal service is received by Caddy, wrapped back into TLS, and sent back to the Web Browser as HTTPS response.
- The website of the internal service shows up in the Web Browser, secured by HTTPS.
Attention
If that does not work, it means that one or multiple steps in that chain of events fail. Please check the following steps for initial troubleshooting.
1. Check the Infrastructure:
- Do A- and/or AAAA-Record for all Domains and Subdomains exist?
- In case of activated Dynamic DNS, check that the correct A- and/or AAAA-Records have been set automatically with Cloudflare.
- Do they point to one of the external IPv4 or IPv6 addresses of the OPNsense Firewall? Check that with commands like nslookup example.com
- Do the OPNsense Firewall Rules allow connections from any source to destination ports 80 and 443 to the destination This Firewall?
- Is the Caddy service running?
2. Check if the Domain is set up correctly:
- Open the Domain in a Web Browser. Inspect the certificate by clicking on the 🔒 in the address bar. It should be a Let’s Encrypt, ZeroSSL or custom certificate (if chosen).
- Activate the HTTP Access Log in a Domain, and check the Log File. Are there any log entries that show connections?
- If nothing shows up, go back to Step 1 and check the infrastructure.
3. Check the functionality of the internal webserver:
- Does the service accept HTTP or HTTPS connections? It is recommended to connect via HTTP, since it removes complexity.
- Open the internal service via IP address and port in a Web Browser, e.g. http://192.168.10.1:8080 . Validate that it shows the website on either HTTP or HTTPS ports.
- Does the internal service actually use the HTTP or HTTPS protocol? Other protocols will not work, e.g. SSH.
- If the Web Browser can not connect, it is a good idea to troubleshoot the internal webserver before continuing.
4. Check the setup of the Handler:
- Is the correct Domain chosen?
- Are Upstream Domain and Upstream Port correct? Do they point to the internal service, e.g 192.168.10.1:8080 ?
- If the internal service only accepts HTTPS connections, is https:// chosen and TLS insecure skip verify checked?
Attention
If the configuration is still not working, it is time to continue with logs and Caddyfile syntax checks.
Get Help from the Caddy Community
Sometimes, things do not work as expected. Caddy provides a few powerful debugging tools to analyze issues.
This section explains how to obtain the required files to get help from the Caddy Community.
- Change the global Log Level to DEBUG. This will log everything the reverse_proxy directive handles.
Go to
- Set the Log Level to DEBUG
- Press Apply
Go to
- Change the dropdown from INFORMATIONAL to DEBUG
Now the reverse_proxy debug logs will be visible and can be downloaded.
- Validate and download the Caddyfile.
Go to
- Press the Validate Caddyfile button to make sure the current Caddyfile is valid. Refresh the page afterwards to ensure the Caddyfile is correctly formatted.
- Press the Download button to get this current Caddyfile.
- If there are custom imports in/usr/local/etc/caddy/caddy.d/ , download the JSON configuration.
Attention
Rarely, a performance profile might be requested. For this, a special admin endpoint can be activated. This admin endpoint is deactivated by default. To enable it and access it on the OPNsense, follow these additional steps. Do not forget to deactivate it after use. Anybody with network access to the admin endpoint can use REST API to change the running configuration of Caddy, without authentication.
- SSH into the OPNsense shell
- Stop Caddy withconfigctl caddy stop
- Go to/usr/local/etc/caddy/caddy.d/
- Create a new file calledadmin.global and put the following content into it:admin :2019
- After saving the file, go to/usr/local/etc/caddy and runcaddy validate to ensure the configuration is valid.
- Start Caddy withconfigctl caddy start
- Use sockstat to see if the admin endpoint has been created.sockstat -l | grep -i caddy - it should show the endpoint*:2019 .
- Create a firewall rule onLAN that allowsTCP to destinationThis Firewall and destination port2019 .
- Open the admin endpoint:http://YOUR_LAN_IP:2019/debug/pprof/
- Follow the instructions on Profiling Caddy.
