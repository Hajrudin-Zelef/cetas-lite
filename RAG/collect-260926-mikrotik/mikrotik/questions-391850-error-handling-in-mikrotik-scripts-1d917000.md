---
id: collect-260926-mikrotik/mikrotik/questions-391850-error-handling-in-mikrotik-scripts-1d917000
title: "questions-391850-error-handling-in-mikrotik-scripts-1d917000"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-391850-error-handling-in-mikrotik-scripts-1d917000.md
source_anchor: ""
source_lines: [1, 26]
sha256: 4de35b13d8a5fbf638a7d737ad3e6caca27edf19c1bd564b901e13be77a6544a
---

# questions-391850-error-handling-in-mikrotik-scripts-1d917000

I have RouterOS 5.14 on RB493G. I need to write script, that launches
/tool fetch ...
Execution of fetch may result in error, this is OK(URL may be sometimes unavailable). Script hangs on error. Is there any way to ignore it?
Solution:
[admin@Mikrotik] >> /system script
0 name=safe-fetch source=
:global done
:global url
/tool fetch $url
:Set done=true
1 name=test source=
:global done
:global url="google.com"
:set done false
:execute safe-fetch
:local counter 0
:while ( $done != true && $counter < 10 ) do={
    :set counter ($counter+1)
    :delay 0.2
    }
if ($done = "true") do={
   :put "Fetch OK"
   } else={
   :put "Fetch ERROR"
   }
Warning: not documented ":execute" is used.
