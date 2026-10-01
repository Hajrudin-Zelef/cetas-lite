---
id: collect-261001-meraki/meraki/questions-9856-raw-api-data-for-cisco-meraki-cmx-dashboard-3dc09aea
title: "questions-9856-raw-api-data-for-cisco-meraki-cmx-dashboard-3dc09aea"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-9856-raw-api-data-for-cisco-meraki-cmx-dashboard-3dc09aea.md
source_anchor: ""
source_lines: [1, 12]
sha256: 753c05e2353a15d2cdfdc7b83c9c965ab564c607701248f1ee02ea1724b0fc97
---

# questions-9856-raw-api-data-for-cisco-meraki-cmx-dashboard-3dc09aea

Has anyone worked with Meraki's CMX raw data feed and can help answer some questions?
I came here to ask because I have found no conversations about Meraki's CMX API elsewhere and I imagined that network engineers are still more likely to know about this service than BI teams or developers.
We have recently set up Meraki access points in a few of our retail stores to report on foot traffic. We set up the data feed to point to a SQL server database with almost live results. Our plan is to have a nightly routine to summarize the data to report on similar metrics as the Meraki dashboard. I have combed through the data we currently have and read through the documentations but notice that not everything is explained. Here are some points I am struggling with.
1. RSSI not fully explained
- an RSSI value of 0 and 102 are disproportionaly higher than any other RSSI value. I assume that a 0 value indicates connected but I am not sure what 102 is for? Max strength or an error?
2. Documentation does not support multiple cases when there are an odd amount of pings or a small number of them
- I have noticed instances where there may be two or three pings in a given 20 minute slot that are about 3 minutes apart. This would suggest that this device was a visitor based on Meraki's documentation but could also be read as a slow passerby, window watcher or someone sitting parking their car out front. Anyone have an idea how Meraki handles these?
- Another case would be when there are only two pings and both are above the RSSI thresholds and greater than 5 minutes apart but less than 20 minutes. Again, this would suggest another visitor but I would argue to be a another passerby.
- My other point of confusion is with strong single pings. I imagine if a phone has a strong ping and only when then they are a passerby but it surprises me not to see multiple pings.
3. Documentation does not explain methodology of tracking visitors that stay longer than 20 minutes
- Are all segments broken up in 20 minute intervals then summed together to create time in store?
Let me know if you want me to add the links to the documentation I have read. Any help would be super helpful.
