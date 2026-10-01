---
id: collect-261001-fortinet/fortinet/questions-23487-fortigate-redirecting-subdomains-only-1af885e8
title: "questions-23487-fortigate-redirecting-subdomains-only-1af885e8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-23487-fortigate-redirecting-subdomains-only-1af885e8.md
source_anchor: ""
source_lines: [1, 15]
sha256: 3394c93170267d2f783a9d3d96f266de0e96cff2642dce2f8bda1f0a5c1e0b24
---

# questions-23487-fortigate-redirecting-subdomains-only-1af885e8

I'm trying to figure out a way to have our Fortigate 60C redirect certain subdomains but not block complete access to the domain.
E.G. I want to block update.microsoft.com but I don't want microsoft.com or office.microsoft.com to be blocked.
My current settings are blocking the complete domain. I do find this strange though as I specified a hostname.
config system dns-database 
edit "test" 
config dns-entry 
edit 2 
set hostname "update" 
set ip 1.1.1.1 
next
end
set domain "microsoft.com"
My goals is to prevent certain software from accessing certain urls to prevent devices from downloading updates or phoning home.
1.1.1.1 is set as a loopback.
Blocking access on devices itself is not an option.
