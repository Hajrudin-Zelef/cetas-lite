---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-3-remote-aaa-hwtacacs-md-09b3925f
title: "kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-3-remote-aaa-hwtacacs-md-09b3925f"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-3-remote-aaa-hwtacacs-md-09b3925f.md
source_anchor: ""
source_lines: [1, 142]
sha256: b4865b4300517c784e4c3673c1f6b51f26bb56ad025bafa59cbcf02784ac51b2
---

# kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-3-remote-aaa-hwtacacs-md-09b3925f

AAA (Authentication, Authorization, Accounting):
Authentication - User identity (қолданушының жеке басын растау)
Authorization - User Permissions (қолданушының жүйеде жұмыс жасау қолжетімділігі немесе құқығы)
Accounting - User Actions Logging (қолданушының жүйеде жасаған әрекеттері)
Негізгі қолдану аясы:
HWTACACS - желілік құрылғыларды басқару (SSH, Telnet)
RADIUS - қолданушылардың желіге (WiFi, VPN) кіруі
| Device Name | Role | Operating System | IP Address | 
|---|---|---|---|
| ubuntu | HWTACACS Server | Linux | 172.16.128.10/24 | 
| R1 | HWTACACS Client | Huawei VRP | 172.16.128.11/24 | 
| R2 | HWTACACS Client | Huawei VRP | 172.16.128.12/24 | 
| R3 | Network Engineer | Huawei VRP | 172.16.128.101/24 | 
| debian | Network Engineer | Linux | 172.16.128.102/24 | 
| Host Machine | Bridge | Windows | 172.16.128.254/24 | 
- Basic Device Configuration
  - Configure IP Address (Linux and Router)
- Create HWTACACS Server Template
- Configure AAA Scheme
- Configure AAA Domain
- Enable SSH Server
- Configure VTY User Interface
- Configure Local Backup Authentication
- Verify the Configuration
student@ubuntu:~$ sudo nano /etc/netplan/50-cloud-init.yaml
network:
  version: 2
  ethernets:
    ens32:
      dhcp4: true
    ens34:
      dhcp4: false
      addresses: [172.16.128.10/24]
CTRL+O, ENTER, CTRL+X
student@ubuntu:~$ sudo netplan apply
student@ubuntu:~$ ip address
ens32: DHCP Assigned
ens34: 172.16.128.10/24student@ubuntu:~$ sudo apt update
student@ubuntu:~$ sudo apt install -y build-essential flex bison libwrap0-dev libpcre2-dev libssl-dev zlib1g-devstudent@ubuntu:~$ git clone https://github.com/MarcJHuber/event-driven-servers.gitstudent@ubuntu:~$ cd event-driven-servers
student@ubuntu:~/event-driven-servers$ ./configure tac_plus
student@ubuntu:~/event-driven-servers$ make
student@ubuntu:~/event-driven-servers$ sudo make install
student@ubuntu:~$ which tac_plus
/usr/local/sbin/tac_plusstudent@ubuntu:~$ sudo nano /etc/tac_plus.conf
id = spawnd {
    listen = { port = 49 }
}
id = tacacs {
    key = Datacom@123
    user = user1 {
        password = cleartext Huawei@123
        member = admin
    }
	
    user = user2 {
        password = cleartext Huawei@123
        member = operator
    }	
    user = user3 {
        password = cleartext Huawei@123
        member = readonly
    }
    group = admin {
        service = exec {
            priv-lvl = 15
        }
    }
    group = operator {
        service = exec {
            priv-lvl = 5
        }
    }
    group = readonly {
        service = exec {
            priv-lvl = 1
        }
    }
}
CTRL+O, ENTER, CTRL+Xstudent@ubuntu:~$ sudo tac_plus -b /etc/tac_plus.conf
student@ubuntu:~$ ss -lntp | grep 49
немесе
student@ubuntu:~$ sudo apt install net-tools
student@ubuntu:~$ netstat -an | grep 49
Create Daemon Service (systemd) File
Creating a systemd Service Unit
student@ubuntu:~$ sudo nano /etc/systemd/system/tac_plus.service
[Unit]
Description=TACACS+ Server
After=network.target
[Service]
ExecStart=/usr/local/sbin/tac_plus /etc/tac_plus.conf
Restart=always
RestartSec=3
[Install]
WantedBy=multi-user.targetstudent@ubuntu:~$ sudo systemctl daemon-reload
student@ubuntu:~$ sudo systemctl start tac_plus
student@ubuntu:~$ sudo systemctl enable tac_plus
student@ubuntu:~$ sudo systemctl status tac_plusstudent@ubuntu:~$ sudo radtest user1 Huawei@123 127.0.0.1 0 testing123
Access-Accept
Create HWTACACS Server Template
[R1] hwtacacs enable
hwtacacs-server template LAN2
 hwtacacs-server authentication 172.16.128.10 49
 hwtacacs-server authorization 172.16.128.10 49
 hwtacacs-server accounting 172.16.128.10 49
 hwtacacs-server shared-key cipher Datacom@123
Configure AAA Scheme
aaa
authentication-scheme HWTACACS
 authentication-mode hwtacacs local
authorization-scheme HWTACACS
 authorization-mode hwtacacs local
accounting-scheme HWTACACS
 accounting-mode hwtacacs
 accounting start-fail online
 accounting realtime 3
Configure AAA Domain
aaa
domain LAB.LOCAL
authentication-scheme HWTACACS
authorization-scheme HWTACACS
accounting-scheme HWTACACS
hwtacacs-server LAN2
Configure the global default domain for administrations
[R1] domain LAB.LOCAL admindisplay hwtacacs-server template LAN2
display domain name LAB.LOCAL
Enable SSH Server
stelnet server enable
display ssh server status
rsa local-key-pair create
Configure VTY User Interface
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound ssh
Configure Local Backup Authentication
aaa
 local-user student password irreversible-cipher Huawei@123
 local-user student service-type terminal ssh
 local-user student privilege level 15
Verify the Configuration
[R1] test-aaa user1 Huawei@123 hwtacacs-server LAN2
[R3] ssh user1@172.16.128.11
