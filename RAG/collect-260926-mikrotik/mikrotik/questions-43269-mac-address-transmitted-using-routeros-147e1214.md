---
id: collect-260926-mikrotik/mikrotik/questions-43269-mac-address-transmitted-using-routeros-147e1214
title: "questions-43269-mac-address-transmitted-using-routeros-147e1214"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2017-08-15"]
keywords: ["consumer"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-43269-mac-address-transmitted-using-routeros-147e1214.md
source_anchor: ""
source_lines: [1, 9]
sha256: 31db9fdb160725f378b8d8cca8794d7cb45037d8582af0eecfc2f4ad311e91b0
---

# questions-43269-mac-address-transmitted-using-routeros-147e1214

Using routerOS connected to another Access Point what MAC address will be sent to the remote AP ? The MAC address of the mikrotik router or of the device connected to this router ?
- 
        Which client mode are you using? What protocol are you using (802.11, NV2 etc)?MerlinTheMagic– MerlinTheMagic2017-08-15 11:15:11 +00:00Commented Aug 15, 2017 at 11:15
- 
        Unfortunately, questions about consumer-grade devices are explicitly off-topic here. You could try to ask this question on Server Fault for a business network, or on Super User for a personal network.Ron Maupin– Ron Maupin ♦2017-08-15 15:26:01 +00:00Commented Aug 15, 2017 at 15:26
1 Answer 1
It depend on the router configuration :
- if the router configured to bridge the traffic between the router client interface and the AP interface the AP will receive the whole clients MAC address
- but if the router configured as normal L3 router which is the most case , the AP will see the router interface MAC address where it is connected
