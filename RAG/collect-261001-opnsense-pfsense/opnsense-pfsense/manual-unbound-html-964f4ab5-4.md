---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-unbound-html-964f4ab5-4
title: "manual-unbound-html-964f4ab5"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/manual-unbound-html-964f4ab5.md
source_anchor: ""
source_lines: [150, 193]
sha256: 0f87c83633a4e614c322a96481119a9c36f6e0595e82e48b6f2047d79c8b77e2
---

# manual-unbound-html-964f4ab5

| Verify CN | The name to use for certificate verification, e.g. “445b9e.dns.nextdns.io”. Used by Unbound to check the TLS authentication certificates. It is strongly discouraged to omit this field since man-in-the-middle attacks will still be possible. | 
Tip
To ensure a validated environment, it is a good idea to block all outbound DNS traffic on port 53 using a firewall rule when using DNS over TLS. Should clients query other nameservers directly themselves, a NAT redirect rule to 127.0.0.1:53 (the local Unbound service) can be used to force these requests over TLS.
Public Resolvers
| Hosted by | Server IP | Server Port | Verify CN | 
|---|---|---|---|
| Cloudflare | 1.1.1.1 | 853 | cloudflare-dns.com | 
|  | 1.0.0.1 |  |  | 
|  | 2606:4700:4700::1111 |  |  | 
|  | 2606:4700:4700::1001 |  |  | 
|  | 8.8.8.8 | 853 | dns.google | 
|  | 8.8.4.4 |  |  | 
|  | 2001:4860:4860::8888 |  |  | 
|  | 2001:4860:4860::8844 |  |  | 
| Quad9 | 9.9.9.9 | 853 | dns.quad9.net | 
|  | 149.112.112.112 |  |  | 
|  | 2620:fe::fe |  |  | 
|  | 2620:fe::9 |  |  | 
Statistics
The statistics page provides some insights into the running server, such as the number of queries executed, cache usage and uptime.
Advanced Configurations
Some installations require configuration settings that are not accessible in the UI.
To support these, individual configuration files with a .conf extension can be put into the
/usr/local/etc/unbound.opnsense.d directory. These files will be automatically included by
the UI generated configuration. Multiple configuration files can be placed there. But note that
- As it cannot be predicted in which clause the configuration currently takes place, you must prefix the configuration with the required clause. For the concept of “clause” see the unbound.conf(5) documentation.
- The wildcard include processing in Unbound is based on glob(7) . So the order in which the files are included is in ascending ASCII order.
- Name collisions with plugin code, which use this extension point e. g. dnsbl.conf , may occur. So be sure to use a unique filename.
- It is a good idea to check the complete configuration via: # check if the resulting configuration is valid configctl unbound check This will report errors that prevent Unbound from starting and also list warnings that may give hints as to why a particular configuration is not working or how it could be improved.
This is a sample configuration file to add an option in the server clause:
server:
  private-domain: xip.io
Note
As a more permanent solution the template system (“Using Templates”) can be used to automatically generate these files.
To get the same effect as placing the file in the sample above directly in /usr/local/etc/unbound.opnsense.d follow these steps:
- Create a +TARGETS file in/usr/local/opnsense/service/templates/sampleuser/Unbound :sampleuser_additional_options.conf:/usr/local/etc/unbound.opnsense.d/sampleuser_additional_options.conf
- Place the template file as sampleuser_additional_options.conf in the same directory:server: private-domain: xip.io
- Test the template generation by issuing the following command: # generate template configctl template reload sampleuser/Unbound
- Check the output in the target directory: # show generated file cat /usr/local/etc/unbound.opnsense.d/sampleuser_additional_options.conf # check if configuration is valid configctl unbound check
Warning
It is the sole responsibility of the administrator which places a file in the extension directory to ensure that the configuration is valid.
Note
This method replaces the Custom options settings in the General page of the Unbound configuration,
which was removed in version 21.7.
