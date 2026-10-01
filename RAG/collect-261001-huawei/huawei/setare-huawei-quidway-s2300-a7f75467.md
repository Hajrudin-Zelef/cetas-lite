---
id: collect-261001-huawei/huawei/setare-huawei-quidway-s2300-a7f75467
title: "setare-huawei-quidway-s2300-a7f75467"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-huawei/setare-huawei-quidway-s2300-a7f75467.md
source_anchor: ""
source_lines: [1, 52]
sha256: ca1b3e758f730c6443b408caa3cda2b552793a94111b121bbba2557a67dede3a
---

# setare-huawei-quidway-s2300-a7f75467

setare huawei quidway s2300
Setare Huawei Quidway IP manual
Setare Huawei Quidway IP automat (dhcp)
Setări Huawei Quidway standard la prima pornire
Setare Huawei Quidway user si parola
Configurare Huawei Quidway server ssh
Setare Huawei Quidway timp automat din NTP
De reținut, după ce faceți toate modificările dorite, dați din save all:
<Quidway>save all The current configuration will be written to the device. Are you sure to continue?[Y/N]y Now saving the current configuration to the slot 0 . Info: Save the configuration successfully.
– dacă dorim sa punem ip-ul 192.168.0.23, cu gateway la 192.168.0.1, rulam:
system-view interface Vlanif 1 ip address 192.168.0.23 255.255.255.0 quit ip route-static 0.0.0.0 0.0.0.0 192.168.0.1
– ma întâi scoate ip dacă este pus un ip:
system-view dhcp enable interface vlanif 1 quit
– după care trecem interfața în dhcp și verificam și ce ip s-a alocat:
system-view dhcp enable interface vlanif 1 ip address dhcp-alloc quit display ip interface Vlanif 1 | include Internet
– și va arata ceva de genul:
Internet Address is allocated by DHCP, 192.168.0.83/24
Setări standard la prima pornire:
system-view loopback-detect enable sysname switch-ul-meu quit system-view header login information /Vezi ce draq mai strici pe aici .../ quit
system-view
rsa local-key-pair create
aaa
local-user root password cipher parola
local-user root level 15
local-user root service-type ssh telnet
quit
– verificam dacă s-a și făcut modificarea:
system-view display current-configuration configuration aaa#
aaa
authentication-scheme default
authorization-scheme default
accounting-scheme default
domain default
domain default_admin
local-user admin password simple admin
local-user admin service-type http
local-user root password cipher JBN;QMAF4<1!!
local-user root privilege level 15
local-user root service-type telnet ssh
#
return
– atenție, pentru ca suntem disperați de securitate și sa evitam pe cat putem, schimbam portul ssh default în 2222 (sau în ce vreți noi și dacă vreti… altfel puneți 22 la port sa fie default):
system-view stelnet server enable ssh user root authentication-type password ssh user root service-type stelnet user-interface vty 0 4 authentication-mode aaa protocol inbound ssh ssh server port 2222 quit
– puneți la ip un server de NTP:
ntp-service unicast-server 193.192.42.18
– pentru verificare tastați ulterior:
display ntp-service status
Ultimele știri huawei:
No feed items found.
Leave a Reply
Want to join the discussion?
Feel free to contribute!
