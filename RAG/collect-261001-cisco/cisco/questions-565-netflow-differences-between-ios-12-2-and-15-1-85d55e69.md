---
id: collect-261001-cisco/cisco/questions-565-netflow-differences-between-ios-12-2-and-15-1-85d55e69
title: "questions-565-netflow-differences-between-ios-12-2-and-15-1-85d55e69"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2013-05-19", "2021-01-04"]
keywords: []
source: docs/RAG/collect-261001-cisco/questions-565-netflow-differences-between-ios-12-2-and-15-1-85d55e69.md
source_anchor: ""
source_lines: [1, 10]
sha256: 664f84a544fcfb4e231301aeb90d9371053ab40b7be087bad0537367621aecac
---

# questions-565-netflow-differences-between-ios-12-2-and-15-1-85d55e69

Has anyone upgraded their cisco routers from 12.2 to 15.1? I'm curious to know what kind of changes were made to netflow. Most specifically, I'd like to know how (if at all) the CPU utilization changed. The router in question is a 6506-E with a Sup720-3BXL
- 
        Did any answer help you? if so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you could post and accept your own answer.Ron Maupin– Ron Maupin ♦2021-01-04 02:25:35 +00:00Commented Jan 4, 2021 at 2:25
2 Answers 2
There should not be obvious difference if you don't have any cfg changed.
Unless you want to use Netflow v9 or IPFIX in new template format in new IOS versions.
I'm not sure, but does this give you any useful insights? Though doubtless a professional Network Engineer would have no need to guess, I'm guessing that your references to 12.2 and 15.1 are Cisco IOS version numbers which I don't see mentioned in the reference, but I thought I'd offer it just in case it was useful anyway?
  ...One reason you may see high CPU on the DFC is due to Netflow Data Export. Typically CPU from NDE is expected, but in rare instances it can become high enough to disrupt other processes...
- 
        That's something pulled from Cisco documentation, it says nothing of the differences between IOS 12 and 15 on a cat 6500Olipro– Olipro2013-05-19 18:47:35 +00:00Commented May 19, 2013 at 18:47
