---
id: collect-261001-meraki/meraki/questions-74054-meraki-and-syslog-ng-ose-360cfd87
title: "questions-74054-meraki-and-syslog-ng-ose-360cfd87"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-74054-meraki-and-syslog-ng-ose-360cfd87.md
source_anchor: ""
source_lines: [1, 3]
sha256: 95068d45bea2ff91ccb3d953c075cd04077585267330075b52de4ea234ee6c6b
---

# questions-74054-meraki-and-syslog-ng-ose-360cfd87

I've been struggling epically to export legible logs from my Meraki devices to a server running Syslog-NG OSE 3.30. No matter what source driver I use on the server, I see errors like this (identifying details changed):
May 28 15:56:23  syslog-ng[32734]: Error processing log message: <134>1>@< 1622231783.857281611 HOSTNAME2 flows allow src=10.1.1.2 dst=10.2.1.2 mac=BLAH protocol=icmp type=0
It seems apparent that Meraki doesn't send messages that conform to RFC3164 or RFC5424. Is that true? Why on earth not? And if so, does that mean I have to parse them specially on my Syslog-NG server with an XML file in patterndb? Can anyone point to an example of one that I can look at?
