---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-ntp-md-48e3378e
title: "kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-ntp-md-48e3378e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-03-23"]
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-ntp-md-48e3378e.md
source_anchor: ""
source_lines: [1, 130]
sha256: 5ed3024aa832c9d2a3ba445e64487b770ba34c59003edfd00d1a274cfd5a1967
---

# kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-ntp-md-48e3378e

Download Link for eNSP Topology File
| Device | Role | interface | IP Address /Prefix | default Gateway | 
|---|---|---|---|---|
| EdgeR1 | NTP Server | g0/0/0 | 192.168.137.254 /24 | 192.168.137.1 | 
|  |  | g0/0/1 | 10.1.77.1 /24 |  | 
| R1 | NTP Client | g0/0/0 | 10.1.77.101 /24 | 10.1.77.1 | 
| S1 | NTP Client | Vlanif1 | 10.1.77.102 /24 | 10.1.77.1 | 
- Configure the IP Address;
- Configure Single-Area OSPF;
- Configure NAT (Easy IP);
- Configure NTP Server;
- Configure NTP Client.
<Huawei> system-view
[Huawei] sysname EdgeR1
[EdgeR1]
int g0/0/0
 ip address 192.168.137.254 24
 quit
int g0/0/1
 ip address 10.1.77.1 24
 quit
display ip int brief[EdgeR1] ping 192.168.137.1
 Request time out
Windows+R ➜ Turn off Windows Defender Firewall
[EdgeR1] ping 192.168.137.1
 Reply from 192.168.137.1: bytes=56 Sequence=1 ttl=128 time=10 ms<Huawei> system-view
[Huawei] sysname R1
[R1]
int g0/0/0
 ip address 10.1.77.101 24
 quit
display ip int brief[R1] ping 10.1.77.1
 Reply from 10.1.77.1: bytes=56 Sequence=1 ttl=255 time=40 ms<Huawei> system-view
[Huawei] sysname S1
[S1]
int Vlanif 1
 ip address 10.1.77.102 24
 quit
display ip int brief[S1] ping 10.1.77.1
 Reply from 10.1.77.1: bytes=56 Sequence=1 ttl=255 time=60 ms[EdgeR1] display ip int brief
[EdgeR1] ospf 1 router-id 1.1.1.1
          area 0
           network 10.1.77.0 0.0.0.255
Configure the Default Static Route
[EdgeR1] ip route-static 0.0.0.0 0.0.0.0 192.168.137.1
Advertise the Default Route
[EdgeR1] ospf 1
          default-route-advertise[R1] display ip int brief
[R1] ospf 1 router-id 2.2.2.2
      area 0
       network 10.1.77.0 0.0.0.255[S1] display ip int brief
[S1] ospf 1 router-id 3.3.3.3
      area 0
       network 10.1.77.0 0.0.0.255display ospf peer
display ospf peer brief
display ip routing-table
display cu section ospf
display cu | begin ospf[EdgeR1] ping 8.8.8.8
 Reply from 8.8.8.8: bytes=56 Sequence=1 ttl=108 time=90 ms[R1] ping 8.8.8.8
 Request time out[S1] ping 8.8.8.8
 Request time out
Configure NAT (Easy IP)
[EdgeR1] acl 2000
          rule permit source 10.1.77.0 0.0.0.255
[EdgeR1] int g0/0/0
          nat outbound 2000
Verify the Configuration
[R1] ping 8.8.8.8
 Reply from 8.8.8.8: bytes=56 Sequence=1 ttl=107 time=130 ms[S1] ping 8.8.8.8
 Reply from 8.8.8.8: bytes=56 Sequence=1 ttl=107 time=120 ms
Уақыт белдеуін (Time Zone) өзгерту
<EdgeR1> clock timezone Almaty add 05:00:00
немесе
<EdgeR1> clock timezone KZ add 5
<EdgeR1> clock datetime 14:26:00 2026-03-23
<EdgeR1> display clock
NTP қызметін іске қосу
ntp-service enable
NTP серверін құру
1-әдіс: LOCAL-ды құрылғының уақытын NTP сервер ретінде қолдану
ntp-service refclock-master 2                   // NTP сервер болу, stratum 2
2-әдіс: Сыртқы NTP сервер уақытын қолдану
undo ntp-service refclock-master               // LOCAL-ды уақытты өшіру
ntp-service unicast-server 80.241.0.72
3-әдіс: Сыртқы NTP сервер уақытымен бірге LOCAL-ды құрылғының уақытын (резервті NTP сервер ретінде) қолдану
ntp-service unicast-server 80.241.0.72
ntp-service refclock-master 5                 // тек қосымша (резерв) NTP сервер ретінде қолданылады
Best practice бойынша Production ортада 2 немесе 3-әдісті қолдану ұсынылады!
NTP аутентификация
ntp-service authentication enable
ntp-service authentication-keyid 1 authentication-mode md5 Datacom@123
ntp-service reliable authentication-keyid 1
Нақты физикалық құрылғыда "hmac-sha256" аутентификация режимін қолдану ұсынылады!
Мысалы: ntp-service authentication-keyid 1 authentication-mode hmac-sha256 cipher Datacom@123
Access Control List (ACL)
acl 2001
 rule 5 permit source 10.1.77.0 0.0.0.255
 rule 10 deny
 quit
ntp-service access peer 2001display cu | include ntp-service
Нәтижені тексеру
display ntp-service status
display ntp-service sessions
display ntp-service sessions verbose
display clock
offset - сервер мен клиент арасындағы уақыт айырмашылығы
Уақыт белдеуін өзгерту (міндетті емес, ұсынылады)
<Huawei> clock timezone Almaty add 05:00:00
немесе
<Huawei> clock timezone KZ add 5
<Huawei> display clock
NTP қызметін іске қосу
ntp-service enable
NTP аутентификация
ntp-service authentication enable
ntp-service authentication-keyid 1 authentication-mode md5 Datacom@123
ntp-service reliable authentication-keyid 1
Нақты физикалық құрылғыда "hmac-sha256" аутентификация режимін қолдану ұсынылады!
Мысалы: ntp-service authentication-keyid 1 authentication-mode hmac-sha256 cipher Datacom@123
NTP сервермен байланыс орнату
ntp-service unicast-server 10.1.77.1 authentication-keyid 1
NTP аутентификация қолданбаған жағдайда NTP сервермен байланыс орнату
ntp-service unicast-server 10.1.77.1
Source interface-ті көрсету (сұраныс жіберетін интерфейс)
[R1] ntp-service source-interface g0/0/0
[S1] ntp-service source-interface Vlanif1display cu | include ntp-service
Нәтижені тексеру
display ntp-service status
display ntp-service sessions
display clock
