---
id: collect-261001-huawei/huawei/questions-1048189-stop-ubiquiti-icmp-restrictions-two-concurrent-pings-at-once-f-485aa416
title: "questions-1048189-stop-ubiquiti-icmp-restrictions-two-concurrent-pings-at-once-f-485aa416"
domain: huawei
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-1048189-stop-ubiquiti-icmp-restrictions-two-concurrent-pings-at-once-f-485aa416.md
source_anchor: ""
source_lines: [1, 5]
sha256: db1054647033058c2a8f6ff69fa49864343412df1e73d6cb111d6e3ac7dec68e
---

# questions-1048189-stop-ubiquiti-icmp-restrictions-two-concurrent-pings-at-once-f-485aa416

How do I change the Ubiquiti Security Gateway's default icmp restrictions from inside the LAN?
It seems that my Ubiquiti Security Gateway's default settings will drop icmp packets if I'm doing more than one traceroute at a time, but I can't find any setting anywhere in the Ubiquiti Controller's wui nor the Security Gateway's firewall rules that look like they're limiting icmp.
For example, when monitoring network issues, I like to kick-off a few simultaneous traceroutes to popular ping farms. Below I kick one off to Google's 8.8.8.8 and CloudFlare's 1.1.1.1 at the same time -- with the ping to Google in a terminal just below the ping to CloudFlare.
Note that the first traceroute has 100% packet loss from my Ubiquiti Security Gateway (ubnt) while the one below it has 0% packet loss. If I stop the traceroute on the bottom, then the other traceroute immediately goes from 100% packet loss to 0% packet loss. So this seems like some kind of overly sensitive icmp flood protection or rate limiting.
Where is this set in the Ubiquity Controller? How can I tune these icmp limits on the LAN to be something more sane?
