---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-2-remote-aaa-radius-md-3c2c522b-1
title: "IPv4 Client"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "license"]
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-2-remote-aaa-radius-md-3c2c522b.md
source_anchor: ""
source_lines: [1, 186]
sha256: 20c9ebeccfdc4bf07950aea353326d527d3468ca6683d87b815e8a34be84d528
---

# IPv4 Client

AAA (Authentication, Authorization, Accounting):
Authentication - User identity (қолданушының жеке басын растау)
Authorization - User Permissions (қолданушының жүйеде жұмыс жасау қолжетімділігі немесе құқығы)
Accounting - User Actions Logging (қолданушының жүйеде жасаған әрекеттері)
Негізгі қолдану аясы:
HWTACACS - желілік құрылғыларды басқару (SSH, Telnet)
RADIUS - қолданушылардың желіге (WiFi, VPN) кіруі
| Device | Role | interface | IP Address / Prefix | Operating System | 
|---|---|---|---|---|
| Ubuntu | RADIUS Server | ens34 | 172.16.128.10 /24 | Linux | 
|  |  | ens32 | DHCP Assigned |  | 
| R1 | RADIUS Client | g0/0/0 | 172.16.128.11 /24 | Huawei VRP | 
| R2 | RADIUS Client | g0/0/0 | 172.16.128.12 /24 | Huawei VRP | 
| R3 | Network Engineer | g0/0/0 | 172.16.128.101 /24 | Huawei VRP | 
| Debian | Network Engineer | ens34 | 172.16.128.102 /24 | Linux | 
|  |  | ens32 | DHCP Assigned |  | 
| Host Machine | Bridge | Loopback1 | 172.16.128.254 /24 | Windows | 
- Basic Device Configuration
  - Configure IP Address
- Create RADIUS Server Template
- Configure AAA Scheme
- Configure AAA Domain
- Enable SSH Server
- Configure VTY User Interface
- Configure Local Backup Authentication
- Verify the Configuration
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
student@ubuntu:~$ sudo netplan apply
немесе
student@ubuntu:~$ sudo netplan trystudent@ubuntu:~$ ip addressstudent@ubuntu:~$ networkctl status
Ping from Ubuntu to Host Machine (Loopback 1)
student@ubuntu:~$ ping -c4 172.16.128.254
Windows+R ➜ Turn off Windows Defender Firewall
FreeRADIUS пакетін (package) орнату
student@ubuntu:~$ sudo apt update
student@ubuntu:~$ sudo apt install -y freeradius freeradius-utils
Daemon-ды жүктеу және автожүктеу қызметін қосу
student@ubuntu:~$ sudo systemctl status freeradius
student@ubuntu:~$ sudo systemctl start freeradius
student@ubuntu:~$ sudo systemctl is-enabled freeradius
student@ubuntu:~$ sudo systemctl enable freeradius
FreeRADIUS пакетінің конфигурациялық файлдар тізімі
student@ubuntu:~$ sudo ls -l /etc/freeradius/3.0/
RADIUS клиенттерді қосу
student@ubuntu:~$ sudo nano /etc/freeradius/3.0/clients.conf
# IPv4 Client
client 172.16.128.0/24 {
    ipaddr = 172.16.128.0/24
    secret = Datacom@123
    nastype = other
}
CTRL+O, ENTER, CTRL+X
немесе
# IPv4 Client
client RADIUS_Client1 {
    ipaddr = 172.16.128.11
    secret = Datacom@123
    shortname = R1
    require_message_authenticator = no
    nastype = other
}
Қолданушыларды қосу
student@ubuntu:~$ sudo nano /etc/freeradius/3.0/users
user1   Cleartext-Password := "Huawei@123"
CTRL+O, ENTER, CTRL+X
ЕСКЕРТУ! Production ортада міндетті түрде MySQL/MariaDB сияқты мәліметтер қорын қолданып, құпиясөзді хэштеу керек!
Huawei Vendor-Specific Attributes (VSA) қосу
student@ubuntu:~$ sudo nano /etc/freeradius/3.0/users
user1   Cleartext-Password := "Huawei@123"
        Service-Type = Login-User,
        Huawei-Exec-Privilege = 15
CTRL+O, ENTER, CTRL+X
немесе
student@ubuntu:~$ sudo nano /etc/freeradius/3.0/users
user1   Cleartext-Password := "Huawei@123"
        Huawei-Exec-Privilege = 15,
        Service-Type = NAS-Prompt-User
CTRL+O, ENTER, CTRL+X
Конфигурациялық файлдың қатесін тексеру
student@ubuntu:~$ sudo freeradius -CX
"Configuration appears to be OK" деген хабарлама шықса, қате жоқ!
Daemon-ды қайта жүктеу
student@ubuntu:~$ sudo systemctl restart freeradius
student@ubuntu:~$ sudo systemctl reload freeradius
UFW конфигурациясы
student@ubuntu:~$ sudo ufw status
student@ubuntu:~$ sudo ufw enable
RADIUS порттарын ашу
student@ubuntu:~$ sudo ufw allow from 172.16.128.0/24 to any port 1812,1813 proto udp
student@ubuntu:~$ sudo ufw allow from 172.16.128.0/24 to any port 22 proto tcp
student@ubuntu:~$ sudo ufw reload
student@ubuntu:~$ sudo ufw status
1812 - Authentication Port Number
1813 - Accounting Port Number
Verify the Configuration
RADIUS Server (Ubuntu)
student@ubuntu:~$ sudo radtest user1 Huawei@123 127.0.0.1 0 testing123
Received Access-Acceptstudent@ubuntu:~$ sudo radtest user1 Huawei@123 172.16.128.10 0 Datacom@123
Received Access-Acceptstudent@debian:~$ ip addressstudent@debian:~$ sudo nano /etc/network/interfaces
# allow-hotplug ens32
auto ens32
iface ens32 inet dhcp
auto ens34
iface ens34 inet static
        address 172.16.128.102
        netmask 255.255.255.0student@debian:~$ ip addressPing from Debian to Ubuntu
student@debian:~$ ping -c2 172.16.128.10
64 bytes from 172.16.128.10: icmp_seq=1 ttl=64 time=1.56 ms
64 bytes from 172.16.128.10: icmp_seq=2 ttl=64 time=0.508 ms
Verify the Configuration
student@debian:~$ sudo apt install -y freeradius-utils
student@debian:~$ sudo radtest user1 Huawei@123 172.16.128.10 0 Datacom@123
Received Access-Accept
Configure the IP Address
int g0/0/0
 ip address 172.16.128.11 24
display ip int brief
Interface                         IP Address/Mask      Physical   Protocol  
GigabitEthernet0/0/0              172.16.128.11/24     up         up        
GigabitEthernet0/0/1              unassigned           down       down      
GigabitEthernet0/0/2              unassigned           down       down      
Ping from Router to Ubuntu
[R1] ping 172.16.128.10
Reply from 172.16.128.10: bytes=56 Sequence=1 ttl=64 time=40 ms
Reply from 172.16.128.10: bytes=56 Sequence=2 ttl=64 time=20 ms
Create a RADIUS Server Template
radius-server template LAN1
 radius-server authentication 172.16.128.10 1812
 radius-server accounting 172.16.128.10 1813
 radius-server shared-key cipher Datacom@123
 quit
Қосымша ақпарат!
radius-server template LAN1
radius-server authentication 172.16.128.10 1812 weight 80
radius-server accounting 172.16.128.10 1813 weight 80
radius-server authentication 172.16.128.9 1812 weight 20
radius-server accounting 172.16.128.9 1813 weight 20
Configure the AAA Scheme
aaa
 authentication-scheme RADIUS
  authentication-mode radius local
  quit
 accounting-scheme RADIUS
  accounting-mode radius
  quit
Configure the AAA Domain
aaa
 domain LAB.LOCAL
  authentication-scheme RADIUS
  accounting-scheme RADIUS
  radius-server LAN1
  quit
 quitdomain LAB.LOCAL admin
Configure the Global default Domain for administrations
domain default_admin admin
domain LAB.LOCAL admin
domain default domain LAB.LOCAL
Verify the Configuration
[R1] test-aaa user1 Huawei@123 radius-template LAN1
Info: Account test time out!
RADIUS серверді "Debug" режимге қосу
student@ubuntu:~$ sudo systemctl stop freeradius
student@ubuntu:~$ sudo freeradius -X
FreeRADIUS Version 3.2.5
Copyright (C) 1999-2023 The FreeRADIUS server project and contributors
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE
You may redistribute copies of FreeRADIUS under the terms of the
GNU General Public License
