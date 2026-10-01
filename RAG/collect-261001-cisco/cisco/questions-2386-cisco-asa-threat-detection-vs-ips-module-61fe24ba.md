---
id: collect-261001-cisco/cisco/questions-2386-cisco-asa-threat-detection-vs-ips-module-61fe24ba
title: "questions-2386-cisco-asa-threat-detection-vs-ips-module-61fe24ba"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2021-01-05"]
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/questions-2386-cisco-asa-threat-detection-vs-ips-module-61fe24ba.md
source_anchor: ""
source_lines: [1, 13]
sha256: f68b5292abb9490f3db68e74ca2684ef5580cf529a44523a9ee265733f0300d1
---

# questions-2386-cisco-asa-threat-detection-vs-ips-module-61fe24ba

There is threat-detection mechanism is CISCO ASA. We can configure different rate limits and actions. Also our ASA 5525-X has enabled integrated IPS module. We can setup protection rules on IPS. What is the best practise of security implementation: threat-detection or IPS, or both?
- 
        What is stopping you from just using both now? As long as you have the performance overhead, more protection is better, right?Michael Teeman– Michael Teeman2013-07-17 10:19:00 +00:00Commented Jul 17, 2013 at 10:19
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can post and accept your own answer.Ron Maupin– Ron Maupin ♦2021-01-05 02:19:50 +00:00Commented Jan 5, 2021 at 2:19
2 Answers 2
Whenever possible I would go for IPS module, it s one of the best security products available on the market. As long as you have the license just fine tune it and use it accordingly to your needs.
For any other scenarios when IPS not available, threat-statistics will prove helpful.
- 
        I have to disagree that Cisco's IPS is one of the best security product on the market. It really cannot compete with SourceFire/Snort (ok, ok it's now Cisco...), Radware, F5 and even PaloAlto.Alex– Alex2013-07-26 15:13:34 +00:00Commented Jul 26, 2013 at 15:13
Threat-detection is just gathering statistics (with possibility to shun attacker's ip).
- 
        If I am right, you can block only hosts that produce only scanning activity. All other information - statistical.Эдуард Буремный– Эдуард Буремный2013-07-18 09:17:44 +00:00Commented Jul 18, 2013 at 9:17
