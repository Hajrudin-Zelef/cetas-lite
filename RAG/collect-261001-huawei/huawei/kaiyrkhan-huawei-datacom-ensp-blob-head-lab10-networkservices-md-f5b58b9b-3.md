---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-networkservices-md-f5b58b9b-3
title: "Create VLANs"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-05-03"]
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-networkservices-md-f5b58b9b.md
source_anchor: ""
source_lines: [587, 654]
sha256: a9c2888b0a163fdfb7c3a671d577ea36d29e836c420e4f52f31c60a1313f9795
---

# Create VLANs

student@ubuntu:~$ sudo chmod -R 755 /srv/tftp
student@ubuntu:~$ ls -ld /srv/tftp
drwxr-xr-x 2 tftp nogroup /srv/tftp
Step3: Configure the Firewall
student@ubuntu:~$ sudo ufw enable
student@ubuntu:~$ sudo ufw allow from 172.16.128.0/24 to any port 69 proto udp
student@ubuntu:~$ sudo ufw deny 69/udp
student@ubuntu:~$ sudo ufw reload
student@ubuntu:~$ sudo ufw statusstudent@ubuntu:~$ ss -tulpna
немесе
student@ubuntu:~$ netstat -tulpna
Step4: Testing the TFTP Server
Download and Upload files
get - Download file from TFTP server
put - Upload file from TFTP server
student@ubuntu:~$ sudo touch /srv/tftp/f1.conf
student@ubuntu:~$ tftp 172.16.128.69 -c get f1.conf
student@ubuntu:~$ pwd
/home/student
student@ubuntu:~$ ls -l
-rw-rw-r-- 1 student student f1.conf
Huawei VRP Router/Switch
# Ping from EdgeR1 to TFTP Server
<EdgeR1> ping 172.16.128.69
  Reply from 172.16.128.69: bytes=56 Sequence=2 ttl=64 time=10 ms# Download file from TFTP server
tftp <tftp-server-ip> get <remote-file>
<EdgeR1> tftp 172.16.128.69 get f1.conf
TFTP: Downloading the file successfully
<EdgeR1> dir
  Idx  Attr     Size(Byte)  Date        Time(LMT)  FileName 
    0  -rw-              0  May 03 2026 11:37:46   f1.conf
tftp 172.16.128.10 get f1.conf
tftp 172.16.128.10 get f1.conf f11.cfg
# Upload file from TFTP server
tftp <tftp-server-ip> put <local-file>
<EdgeR1> tftp 172.16.128.69 put vrpcfg.zip
TFTP: Uploading the file successfully
student@ubuntu:~$ ls -lh /srv/tftp/
-rw-r--r-- 1 root root f1.conf
-rw-rw-rw- 1 tftp tftp vrpcfg.zip
NTP серверді конфигурациялау (EdgeR1)
# Уақыт белдеуін (Time Zone) өзгерту
<EdgeR1> clock timezone KZ add 5
<EdgeR1> clock datetime 17:45:00 2026-05-03
<EdgeR1> display clock# NTP қызметін іске қосу
[EdgeR1] ntp-service enable# LOCAL-ды құрылғының уақытын NTP сервер ретінде қолдану
[EdgeR1] ntp-service refclock-master 3                        // NTP сервер болу, stratum 3# NTP аутентификация
ntp-service authentication enable
ntp-service authentication-keyid 1 authentication-mode md5 Datacom@123
ntp-service reliable authentication-keyid 1display cu | include ntp-service# Нәтижені тексеру
display ntp-service status
display ntp-service sessions
display ntp-service sessions verbose
display clock
NTP клиентті конфигурациялау (C1, D1, D2, A1, A2, DHCP Server)
# Уақыт белдеуін (Time Zone) өзгерту
<C1> clock timezone KZ add 5
<C1> display clock# NTP аутентификация
ntp-service authentication enable
ntp-service authentication-keyid 1 authentication-mode md5 Datacom@123
ntp-service reliable authentication-keyid 1# NTP сервермен байланыс орнату
ntp-service unicast-server 50.1.1.1 authentication-keyid 1# Source interface-ті көрсету (сұраныс жіберетін интерфейс)
[C1] ntp-service source-interface Loopback 50
[D1] ntp-service source-interface Loopback 50
[A1] ntp-service source-interface Vlanif50# Нәтижені тексеру
display ntp-service status
display ntp-service sessions
display clock
