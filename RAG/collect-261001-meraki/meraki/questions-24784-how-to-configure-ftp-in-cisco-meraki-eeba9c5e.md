---
id: collect-261001-meraki/meraki/questions-24784-how-to-configure-ftp-in-cisco-meraki-eeba9c5e
title: "How to configure FTP in Cisco Meraki?"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-24784-how-to-configure-ftp-in-cisco-meraki-eeba9c5e.md
source_anchor: ""
source_lines: [1, 24]
sha256: e6d4b7da41df3b170b77e454a5c19851429cafe3b4cafa49ba9489b154114eaf
---

# How to configure FTP in Cisco Meraki?

*Score : 3 | Source : https://networkengineering.stackexchange.com/questions/24784/how-to-configure-ftp-in-cisco-meraki*

Has anyone configured FTP rules in the Meraki MX family of network devices?
Locally, inside the LAN, I can login to the FTP server (IIS Windows 2012 server) without problems; but, I cannot login from outside (via WAN).
The FTP Server uses Passive mode.
Later I'm going to attached links from Meraki and YouTube websites that I have already tried.

---

### Reponse — score 4

Active and Passive FTP Overview and Configuration
Under Security Appliance > Firewall, configure a 1:1 NAT with the allowed inbound connections.
  Two firewall rules are necessary for passive FTP to function properly: - The firewall must allow connections on port 21. - The firewall must allow connections to the ephemeral ports used by the FTP application.

---

### Reponse — score 0

Allow port forward on TCP 21 to your server
Same for TCP 50000 - 50100 following advice of https://wiki.filezilla-project.org/Network_Configuration#Active_mode
Worked for me on Meraki MX64w
