---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d-4
title: "setup-adguard-home-opnsense-adblocker-c464187d"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/setup-adguard-home-opnsense-adblocker-c464187d.md
source_anchor: ""
source_lines: [234, 247]
sha256: 897a0f7fa59be83e454462dc051e16886ea2fc6e0783a6226904a6da7e5f1349
---

# setup-adguard-home-opnsense-adblocker-c464187d

What if I have AdGuard running on a different server? Or want to keep using Unbound DNS?
How do I change the interface / port for the Web UI or DNS?
So perhaps we mis-typed something when configuring AdGuard. Or just wanted to change the interface IP address AdGuard listens on. No problem!
Unfortunately, since this is a community plugin – there is no configuration for the plugin within the OPNsense interface.
We’ll need to reconnect to the OPNsense command line to make some additional configuration changes. This can be done via SSH or the device console.
Once there, we can use the command edit /usr/local/AdGuardHome/AdGuardHome.yaml.
At the top, bind_host & bind_port pertains to the admin web interface. A little below there, under the dns section – you’ll see another bind_hosts and port config. Those ones are specific to the DNS server side of things.
Once done, save the config file by pressing Esc then selecting to quit the editor & save the file.
Lastly – Go back into the OPNsense web UI & restart the AdGuard Home service for the changes to take effect.
Src:
https://samuelsson.dev/install-adguard-home-on-an-opnsense-router/
https://0x2142.com/how-to-set-up-adguard-on-opnsense/
Maciej Zytowiecki
Network security expert with a deep passion for wireless networks, networking and data security. When I'm not working, you'll find me diving into hobby projects, contributing to open-source initiatives, or enjoying hands-on experiments with cutting-edge tech. My goal is to bridge the gap between complex concepts and accessible knowledge, making the world of network security both intriguing and approachable for all.
