---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1761737-openvpn-on-opnsense-fails-to-start-after-restore-1d00ea8b
title: "OpenVPN on OPNsense fails to start after restore"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1761737-openvpn-on-opnsense-fails-to-start-after-restore-1d00ea8b.md
source_anchor: ""
source_lines: [1, 26]
sha256: e95783251c96e19f6f8cb48b1dda230dc286dd1df264d5011210f6c0a9cd5556
---

# OpenVPN on OPNsense fails to start after restore

*Score : 0 | Source : https://superuser.com/questions/1761737/openvpn-on-opnsense-fails-to-start-after-restore*

After restoring my OPNsense system, the OpenVPN server fails to start.
Here’s how I restored the system:
However, the OpenVPN server I have configured fails to start. The error message I get is:
Options error: --dh fails with '/usr/local/etc/dh-parameters..sample': No such file or directory (errno=2)
Options error: Please correct these errors.
Indeed, when open a shell, the file does not exist. However, a bunch of similarly named files are there:
dh-parameters.1024.sample
dh-parameters.2048.sample
dh-parameters.4096.sample
as well as some ending in .rfc7919 instead of .sample, and some missing the last extension altogether.
I figure the digits should correspond to what is set as DH Parameters Length in the web GUI (1024 in my case).
How do I get back to a working OpenVPN config?

---

### Reponse (acceptee) — score 0

The DH parameter length has been dropped in more recent versions of OPNsense. As it turns out, I had originally set up OPNsense on the 22.1 series and, after restring the SD card, upgraded to 22.1.10, which was the last update of the 22.1 series. However, 22.7 had been out for a while and I had updated earlier. So essentially I was restoring a configuration from a newer version to an older version.
I didn’t notice that at first and, after looking at the XML config I restored, concluded that the DH parameter length did not get exported. Presumably the parameter was empty internally, which would explain the invalid filename. The GUI, on the other hand, just fell back to its default.
Editing the server config once, (re-)selecting the default DH parameter length and saving the config fixed things for me. The server starts up again and I can connect. Looking at the config history, I see that a new element /opnsense/openvpn/openvpn-server/dh_length got added.
I have since updated to 22.7.10 – doing so before restoring the configuration would have saved me the hassle.
Lesson learned: check twice you are really on the latest version before you restore a previously saved configuration.
