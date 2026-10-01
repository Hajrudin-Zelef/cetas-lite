---
id: collect-261001-general-networking/general-networking/manual-how-tos-ntopng-html-5f127a63
title: "manual-how-tos-ntopng-html-5f127a63"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-ntopng-html-5f127a63.md
source_anchor: ""
source_lines: [1, 13]
sha256: 029be79bb62d85e0f6fb39802076b750c12b92e3707715f4c7edf3bae0a9a4d0
---

# manual-how-tos-ntopng-html-5f127a63

ntopng
Installation
First of all, you have to install the ntopng plugin (os-ntopng) from the plugins view reachable via .
After a page reload you will get a new menu entry under Services for ntopng. If you don’t have Redis plugin installed, you’ll receive a warning in ntopng main menu. Please go back to , install os-redis, change to and just enable the service. That’s enough to run ntopng.
General Settings
- Enable ntopng
- Enable and start ntopng.
- Interfaces
- Here you set the interfaces ntopng should listen on. If you don’t select any interface it listens to the first in the system, e.g. em0, but you can change the interfaces within ntopng’s UI on demand; while setting an explicit interface you will not get any other interface presented in its own UI.
- HTTP Port
- The port ntopng’s UI should listen on. When you leave it on the default just open a browser and go to your Firewall IP with port 3000 and HTTP. If you want to secure the connection feel free to setup HAProxy or Nginx as a reverse proxy (SSL offloading).
- DNS Mode
- Here you can choose if ntopng should try to resolve IPs to host names.
