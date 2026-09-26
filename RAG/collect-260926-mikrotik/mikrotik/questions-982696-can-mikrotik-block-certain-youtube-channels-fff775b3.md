---
id: collect-260926-mikrotik/mikrotik/questions-982696-can-mikrotik-block-certain-youtube-channels-fff775b3
title: "questions-982696-can-mikrotik-block-certain-youtube-channels-fff775b3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-982696-can-mikrotik-block-certain-youtube-channels-fff775b3.md
source_anchor: ""
source_lines: [1, 8]
sha256: 4d42051c88f2fb2ec71e943290ee00a5a7d6ecae254926d5d179882f9b4fab3b
---

# questions-982696-can-mikrotik-block-certain-youtube-channels-fff775b3

I want to block some adult content from YouTube.
Can MikroTik routers block certain Youtube Channels? Or certain content from YouTube?
Well, yes and no. You can block specific videos from YouTube, but as far as I know, not a certain YouTube channel due the format of YouTube URLs.
Through RouterOS, you can accomplish this using one of these two methods:
1) You can create Layer 7 Protocol specifications in the Firewall section, and then firewall rules to block them;
2) You can enable RouterOS built-in web-proxy and block them there.
Hope that helps.
Best regards.
