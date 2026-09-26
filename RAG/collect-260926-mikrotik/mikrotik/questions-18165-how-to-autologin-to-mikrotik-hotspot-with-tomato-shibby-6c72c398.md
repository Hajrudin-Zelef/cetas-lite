---
id: collect-260926-mikrotik/mikrotik/questions-18165-how-to-autologin-to-mikrotik-hotspot-with-tomato-shibby-6c72c398
title: "questions-18165-how-to-autologin-to-mikrotik-hotspot-with-tomato-shibby-6c72c398"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-18165-how-to-autologin-to-mikrotik-hotspot-with-tomato-shibby-6c72c398.md
source_anchor: ""
source_lines: [1, 15]
sha256: 76aa0a86150667e4ae97ac17d212f88d0863cdf8f35f0d02223e30e26bf23e10
---

# questions-18165-how-to-autologin-to-mikrotik-hotspot-with-tomato-shibby-6c72c398

I would like to to login automatically on Mikrotik's hotspot using a router with Tomato Shibby firmware.
- 
        this is off topic here, please ask tomato firmware questions on Super UserMike Pennington– Mike Pennington2015-04-28 12:17:07 +00:00Commented Apr 28, 2015 at 12:17
- 
        ok sorry, didn't know that :(brauliobo– brauliobo2015-04-28 13:39:30 +00:00Commented Apr 28, 2015 at 13:39
1 Answer 1
Schedule the following script on Administration/Scheduler
/bin/sh -c 'echo "POST /login HTTP/1.1
User-Agent: Wget/1.16.1 (linux-gnu)
Accept: */*
Accept-Encoding: identity
Host: 10.22.0.1
Content-Type: application/x-www-form-urlencoded
Content-Length: 35
username=adrianoimba2&password=5149"| nc 10.22.0.1 80'
