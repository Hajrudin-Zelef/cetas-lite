---
id: collect-260926-mikrotik/mikrotik/questions-397320-routeros-on-hyper-v-v3-2012-any-way-to-get-it-working-8a453bb4
title: "questions-397320-routeros-on-hyper-v-v3-2012-any-way-to-get-it-working-8a453bb4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-397320-routeros-on-hyper-v-v3-2012-any-way-to-get-it-working-8a453bb4.md
source_anchor: ""
source_lines: [1, 4]
sha256: 780b6fdc075525f445b1065e4406298d93e1f26c594827e130bf86b45dac7bc3
---

# questions-397320-routeros-on-hyper-v-v3-2012-any-way-to-get-it-working-8a453bb4

Trying to set up a small VPN point to connect into a remote Hyper-V cluster using ROuterOS. Anyone got it working ON Hyper-V with the latest builds of RouterOS? It seems the legacy network adapter is not supported anymore either (or just broken). The platform is a Windows Server 2012 RC.
This is not a high performance setup - the RouterOS wont do the routing for more than the backend administrative access, and the only real traffic we will see there is when ISO images for new operating systems are uploaded. Otherwise we will have possibly RDP traffic as well as web / http traffioc, but this is internal only (dashboards, some control panel). The server has no public business. So the price for non virtualized network cards is ok for me.
After hooking up - ping just does not work. After some time I see in windows (arp -a on the command line), so I know that the Hyper-V side is set up properly. Just no packets arrived. I have turned off all protection on Hyper-V (or : not turned them on), so no MAC spoofing protection etc. in the Advanced page for the legacy adapters.
Unless I can get it work I will have to resort to using a windows install as router / VPN endpoint, which introduces another OS into the fabric (we run all routers etc. so far on mikrotik in hardware, which is why I want this one to be RouterOS, too). And no, putting hardware there is NOT an option - the cost would be significant.
