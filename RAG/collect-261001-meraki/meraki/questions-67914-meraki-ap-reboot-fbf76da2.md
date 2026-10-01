---
id: collect-261001-meraki/meraki/questions-67914-meraki-ap-reboot-fbf76da2
title: "questions-67914-meraki-ap-reboot-fbf76da2"
domain: meraki
role: reference
task: reference
actors: []
dates: ["2020-12-17"]
keywords: []
source: docs/RAG/collect-261001-meraki/questions-67914-meraki-ap-reboot-fbf76da2.md
source_anchor: ""
source_lines: [1, 5]
sha256: c2f9c9f5042ffa4a555493572b63b0a3997875c12ccdbd977cee833f1ba3cfb2
---

# questions-67914-meraki-ap-reboot-fbf76da2

Does Meraki WiFi Access Points reboot after 4 Hours of not getting reach-ability with Cloud Controller?
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can post and accept your own answer.Ron Maupin– Ron Maupin ♦2020-12-17 17:28:10 +00:00Commented Dec 17, 2020 at 17:28
1 Answer 1
The Meraki MR series access points will only reboot if they have sustained 4 hours of continuous loss to the cloud and all of the SSIDs are configured for NAT mode otherwise they will continue to operate normally albeit without telemetry streaming to the dashboard or being able to configure.
