---
id: collect-261001-meraki/meraki/questions-60917-meraki-firewall-mx64-how-to-do-two-ip-seperated-inbound-nats-074c8353
title: "questions-60917-meraki-firewall-mx64-how-to-do-two-ip-seperated-inbound-nats-074c8353"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-60917-meraki-firewall-mx64-how-to-do-two-ip-seperated-inbound-nats-074c8353.md
source_anchor: ""
source_lines: [1, 7]
sha256: 6cebcf8cf3e8e8950ed48bfeae620fdd6373f416a28160b70d5e3fa011472b7b
---

# questions-60917-meraki-firewall-mx64-how-to-do-two-ip-seperated-inbound-nats-074c8353

I need to achieve the same result of these two commands which are on Cisco CLI but on Meraki GUI
so we have two valid public IP address(81.1.1.30,31) on outside interface of MX64
both of 'em want to be forwarded to two seperate Webservers
Switch6500(config)#ip nat inside source static 192.168.1.50 tcp 80 81.1.1.30 tcp 80
Switch6500(config)#ip nat inside source static 192.168.1.51 tcp 80 81.1.1.31 tcp 80
How can I do that? I couldn't find a place to specify what outside IP+Port combination I want to be forwarded to this Private(local) IP add+port combination, all they have is either interface(uplink) alone or IP address without port alone
thx
