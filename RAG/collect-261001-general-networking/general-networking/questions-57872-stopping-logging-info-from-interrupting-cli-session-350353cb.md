---
id: collect-261001-general-networking/general-networking/questions-57872-stopping-logging-info-from-interrupting-cli-session-350353cb
title: "questions-57872-stopping-logging-info-from-interrupting-cli-session-350353cb"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2019-12-14"]
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-57872-stopping-logging-info-from-interrupting-cli-session-350353cb.md
source_anchor: ""
source_lines: [1, 5]
sha256: 3dfc9c2d88411de585cd47bd1267d1f83e74bf8790218960dcf9002e670bffff
---

# questions-57872-stopping-logging-info-from-interrupting-cli-session-350353cb

With cisco devices the logging synchronous command will stop logging output from interrupting the the CLI session. Is there an equivalent command for the fortigate firewalls? i.e. When running a debug on the firewall is it possible to keep the cli prompt below of the debug output at all times so that typing/input is never interrupted?
- 
        I don't think so, but you can filter the debug console: kb.fortinet.com/kb/documentLink.do?externalID=FD33882 -- or perhaps start a second SSH session - one for config changes and one for debugging?PSaul– PSaul2019-03-22 15:42:52 +00:00Commented Mar 22, 2019 at 15:42
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can provide and accept your own answer.Ron Maupin– Ron Maupin ♦2019-12-14 21:39:46 +00:00Commented Dec 14, 2019 at 21:39
