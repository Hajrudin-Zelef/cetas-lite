---
id: collect-260926-mikrotik/mikrotik/questions-549070-mikrotik-beep-on-rdp-c2030b14
title: "questions-549070-mikrotik-beep-on-rdp-c2030b14"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-549070-mikrotik-beep-on-rdp-c2030b14.md
source_anchor: ""
source_lines: [1, 13]
sha256: 2b610f3f27f72f5c13a04cf9a5572de110ad02246c179c44f0ee99e1fa7b4364
---

# questions-549070-mikrotik-beep-on-rdp-c2030b14

I am a newbie in MikroTik RouterOS.
How can I add a rule or script to beep when some one tries to connect to a special port on one of interfaces?
For example I need to know if someone is trying to ftp to my server from outside world.
As far as I know, there is no way to do that directly.
Workaround could be using firewall rule to add source ip to addresslist:
/ip firewall filter add action=add-src-to-address-list address-list=beeplist
    address-list-timeout=1m1s chain=input disabled=no dst-port=21 protocol=tcp
and then run script to check'n'beep:
:if ([ :len [ip firewall address-list find where list=beeplist]]>0) do={:beep}
periodically via scheduler:
/system scheduler add interval=60s name=beeplist on-event=":if ([ :len [ip firewall address-list find where list=beeplist]]>0) do={:beep}"
Entries in address-list are going to be deleted via address-list-timeout settings.
Still, you are going to hear a lot of beeps...
