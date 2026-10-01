---
id: collect-261001-huawei/huawei/2020-07-03-basic-configurasi-snmplogssh-untuk-switch-huawei-s5700-series-3c993189
title: "View traps output to the NMS station menggunakan SNMP agent."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/2020-07-03-basic-configurasi-snmplogssh-untuk-switch-huawei-s5700-series-3c993189.md
source_anchor: ""
source_lines: [1, 71]
sha256: fa6140cf20674b7b1664d070588a03cefea65c712bc90325ccd4563a16b4a413
---

# View traps output to the NMS station menggunakan SNMP agent.

Configure ip address dan ssh
[huawei]system-view
[huawei]interface Vlanif 1
[huawei]ip address 169.254.1.5 16
#Enable http & https
[huawei]system-view
[huawei]http  server load s5700si-v200r003c00.001.web.7z
[huawei]http secure-server enable
[huawei]http server enable
#Enable stelnet for user telnet
[huawei]stelnet server enable
[huawei]user-interface vty 0 4
[huawei]authentication-mode aaa
[huawei]protocol inbound ssh
#Add user and password for telnet
[huawei]aaa
[huawei-aaa]local-user admin password cipher password123
[huawei-aaa]local-user admin privilege level 15
[huawei-aaa]quit
#add service-type for user
[huawei]ssh user admin authentication-type password
[huawei]ssh user admin service-type all
[huawei]quit
<huawei>save
————————————————————————–
Reset ke factory setting
#<huawei>reset saved-configuration
The configuration will be erased to reconfigure. Continue? [Y/N]:y
#<huawei>rebootWarning: All the configuration will be saved to the configuration file for the next startup:, Continue?[Y/N]:n  
pilih n, jika memilih y konfigurasi lama akan tersimpan, dan setelah direstart configurasi akan tetap sama.
Turn off dan Turn on kembali switch
—————————————————————————–
Set SNMP Community Name and password
-go to configure mode
<huawei> system-view
[huawei]snmp-agent
#set community name public (encrypt name when display current configuration)
[huawei]snmp-agent community read cipher public
#set community version( sekarang sudah support untuk semua versi)
[huawei]snmp-agent sys-info version all  
*jika kamu menginginkan hanya versi tertentu, ganti “all” dengan v3,v1 atau v2c
#configure tujuan alamat IPv4 address untuk traps dan error code terkirim dari device dan encript password cipher
[huawei]snmp-agent target-host trap address udp-domain 10.105.9.28 udp-port 161 params securityname cipher public
#enable snmp trap
[huawei]snmp-agent trap enable
quit and save.
#verify configuration:
# View traps output to the NMS station menggunakan SNMP agent.
<huawei> display snmp-agent target-host
#View SNMP trap version
<huawei> display snmp-agent system-info
——————————————————————-
Configure send log ke linux Server
1. enable command di huawei switch
<Switch>system-view
[Switch]info-center enable
[Switch]info-center loghost 203.217.168.48 facility local7 language English log-counter enable
[Switch]undo info-center source default channel loghost
2. Di sisi linux, kamu bisa mendapatkan syslog atau syslog-ng jika dibutuhkan
configuras bisa seperti berikut :
# mkdir /var/log/Switch
# touch /var/log/Switch/information
# Switch configuration messages
local7.info /var/log/Switch/information
– restart syslog daemon
=======================================
artikel lain:
–http://huaweioverview.blog.fc2.com/blog-entry-11.html
–http://huaweiswitchesrouters.blogspot.com/2013/11/how-to-configure-basic-snmpv3-functions.html
–http://huaweis3700.weebly.com/blog/huawei-s5700-snmp-configuration
–http://www.kuncar.net/blog/huawei-equipment-and-syslog-daemon/2008/
