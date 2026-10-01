---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-ntp-linux-md-ba174a99
title: "Kazakhstan NTP pool"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-ntp-linux-md-ba174a99.md
source_anchor: ""
source_lines: [1, 143]
sha256: 9303e55d674a0bfbe48ae305e6019ea3e92348f6215dba000a645cffade4b787
---

# Kazakhstan NTP pool

| Device | Role | interface | IP Address / Prefix | Operating System | 
|---|---|---|---|---|
| Ubuntu | NTP Server | ens34 | 172.16.128.10 /24 | Ubuntu Server | 
|  |  | ens32 | DHCP |  | 
| R1 | NTP Client | g0/0/0 | 172.16.128.11 /24 | Huawei VRP | 
| Debian | NTP Client | ens34 | 172.16.128.12 /24 | Debian Linux | 
|  |  | ens32 | DHCP |  | 
| Host Machine | Bridge | Loopback1 | 172.16.128.254 /24 | Microsoft Windows | 
- Configure NTP Server on Ubuntu;
- Configure NTP Client on Huawei VRP;
- Configure NTP Client on Debian.
VMware Workstation Pro ➜ Virtual Machine Settings ➜ Add Hardware Wizard ➜ ...
VMware Workstation Pro ➜ Edit ➜ Virtual Network Editor ➜ Change Settings
student@ubuntu:~$ lsb_release -a
Ubuntu 24.04.4 LTS
student@ubuntu:~$ uname -rs
Linux 6.8.0-101-generic x86_64 GNU/Linuxstudent@ubuntu:~$ ip addressstudent@ubuntu:~$ sudo nano /etc/netplan/50-cloud-init.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    ens32:
      dhcp4: true
    ens34:
      dhcp4: false
      addresses:
        - 172.16.128.10/24
CTRL+O, ENTER, CTRL+X
ЕСКЕРТУ: YAML файлында бос орындар (indentation) өте маңызды. Әр қатарда 2 бос орын қолдануды ұмытпаңыз! (Tab пернесін қолданбаған дұрыс)
student@ubuntu:~$ sudo netplan try
немесе
student@ubuntu:~$ sudo netplan applystudent@ubuntu:~$ ip addressstudent@ubuntu:~$ networkctl status
Ping from Ubuntu to Host Machine (Loopback 1)
student@ubuntu:~$ ping -c4 172.16.128.254
 4 packets transmitted, 0 received, 100% packet loss, time 3058ms
Windows+R ➜ Turn off Windows Defender Firewall
student@ubuntu:~$ ping -c4 172.16.128.254
 64 bytes from 172.16.128.254: icmp_seq=1 ttl=128 time=0.230 ms
Chrony пакетін (package) орнату
Package атауы: chrony
Daemon/Service атауы: chrony немесе chronyd
chronyd – the actual daemon to sync and serve via the Network Time Protocol
chronyc – command-line interface for the chrony daemon
$ sudo apt update 
$ sudo apt install -y chrony$ sudo systemctl status chronyd
Уақыт белдеуін (Time Zone) өзгерту
$ sudo timedatectl set-timezone Asia/Almaty
$ timedatectl status
Time Zones in Kazakhstan https://www.timeanddate.com/time/zone/kazakhstan
NTP серверді конфигурациялау
NTP Pool Time Servers Link: https://www.ntppool.org/zone/kz
артық DNS атауларды "#" comment-ге алып, төменгі қатарға Қазақстанға ең жақын NTP серверлердің DNS атауын енгіземіз!
$ sudo nano /etc/chrony/chrony.conf
#pool ntp.ubuntu.com        iburst maxsources 4
#pool 0.ubuntu.pool.ntp.org iburst maxsources 1
#pool 1.ubuntu.pool.ntp.org iburst maxsources 1
#pool 2.ubuntu.pool.ntp.org iburst maxsources 2
# Kazakhstan NTP pool
server ntp.nic.kz iburst
pool 2.kz.pool.ntp.org iburst
pool 1.kz.pool.ntp.org iburst
# Global NTP pool
pool time.google.com iburst
pool time.cloudflare.com iburst
# Listen on all interfaces
bindcmdaddress 0.0.0.0
bindcmdaddress ::
# Allow NTP client access from Local Network
allow 172.16.128.0/24
# NTP authentication
keyfile /etc/chrony/chrony.keys
# Log files location
logdir /var/log/chrony
log measurements statistics tracking
# Hardware clock synchronization
rtcsync
# Time adjustment settings (уақыт дәлдігін реттеу)
makestep 1 3
NTP аутентификация
$ sudo nano /etc/chrony/chrony.keys
# Huawei VRP <key_id> <algorithm> <secret_key>
1 MD5 Datacom@123
CTRL+O, ENTER, CTRL+X
ЕСКЕРТУ: мұндағы, "MD5" бас әріппен жазылуы міндетті!
Firewall конфигурациялау
$ sudo ufw status
$ sudo ufw enable
NTP портына (123/UDP) рұқсат ету
$ sudo ufw allow from 172.16.128.0/24 to any port 123 proto udp
$ sudo ufw allow from 172.16.128.0/24 to any port 22 proto tcp
$ sudo ufw reload
$ sudo ufw status verbose
Daemon-ды қайта жүктеу
$ sudo systemctl restart chronyd
немесе
$ sudo systemctl reload chronyd
$ sudo systemctl status chronyd$ ss -tulpn
Netid  State    Local Address:Port    Peer Address:Port
udp    -        0.0.0.0:123           0.0.0.0:*$ sudo apt install -y net-tools
$ netstat -tulpn
Proto  Local Address  Foreign Address   State
udp    0.0.0.0:123    0.0.0.0:*         -
Нәтижені тексеру
$ sudo chronyc sources -v$ sudo chronyc tracking
$ sudo chronyc activity$ sudo apt install ntpdate
$ sudo ntpdate -q 80.241.0.72
Configure the IP Address
<Huawei> system-view
[Huawei] sysname R1
[R1]
int g0/0/0
 ip address 172.16.128.11 24
 quit
display ip int brief[R1] ping 172.16.128.10
 Reply from 172.16.128.10: bytes=56 Sequence=3 ttl=64 time=10 ms<Huawei> system-view
[Huawei] sysname S1
[S1]
int Vlanif 1
 ip address 172.16.128.12 24
 quit
display ip int brief[S1] ping 172.16.128.10
 Reply from 172.16.128.10: bytes=56 Sequence=4 ttl=64 time=40 ms
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
ntp-service unicast-server 172.16.128.10 authentication-keyid 1
Source interface-ті көрсету (сұраныс жіберетін интерфейс)
[R1] ntp-service source-interface g0/0/0
[S1] ntp-service source-interface Vlanif1display cu | include ntp-service
Нәтижені тексеру (Check the Configuration)
display clockdisplay ntp-service statusdisplay ntp-service sessions
мұндағы, "reach" мәні 255 және 377 көрсету керек!
