---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-2-remote-aaa-radius-md-3c2c522b-2
title: "IPv4 Client"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab8-2-remote-aaa-radius-md-3c2c522b.md
source_anchor: ""
source_lines: [187, 277]
sha256: 358a9632fe84ac3d622d87ea791d409ad302dcd89618b7167a98a77b2c13cc3b
---

# IPv4 Client

For more information about these matters, see the file named COPYRIGHT
Starting - reading configuration files ... [R1] test-aaa user1 Huawei@123 radius-template LAN1
Info: Account test succeeded!Starting - reading configuration files ...
CTRL+C
student@ubuntu:~$ sudo systemctl start freeradius
student@ubuntu:~$ sudo systemctl status freeradius
[R1]test-aaa user1 Huawei@123 radius-template LAN1
Info: Account test succeeded!
Қосымша ақпарат!
[R1] radius-server test-template LAN1 172.16.128.10 1812 user1 password Huawei@123
SSH server permit interface
ssh server permit interface GigabitEthernet 0/0/2
ssh server permit interface Vlanif 1
ssh server permit interface all
немесе
ssh server-source -i Vlanif 1
ssh server-source all-interface
Enable the SSH Server
stelnet server enable
display ssh server status
rsa local-key-pair create
Warning: Confirm to replace them! Continue? [Y/N] Y
Input the bits in the modulus[default = 2048]: 2048
Configure the VTY User Interface
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound ssh
Configure Local Backup Authentication
aaa
 local-user student password irreversible-cipher Huawei@123
 local-user student service-type terminal ssh
 local-user student privilege level 15
student қолданушыны жүйеден жою: undo local-user student
Troubleshooting Commands
display current-configuration section aaa
display radius-server configuration template LAN1
display domain name LAB.LOCALstudent@ubuntu:~$ sudo tcpdump -i any udp port 1812 -n
Network Engineer (Huawei VRP Router)
[R3] ssh client first-time enable
[R3] stelnet 172.16.128.11
Please input the username: user1
The server is not authenticated. Continue to access it? (y/n)[n]: y
Save the server's public key? (y/n)[n]: y
Enter password: Huawei@123
R1-ге SSH арқылы кіргеннен кейін, төмендегі командаларды орындап көріңіз!
[R1] display privilege state — Privilege Level тексеру
[R1] display users
Network Engineer (Debian Linux)
student@debian:~$ ssh \
-oKexAlgorithms=+diffie-hellman-group1-sha1 \
-oHostKeyAlgorithms=+ssh-rsa \
-oPubkeyAcceptedAlgorithms=+ssh-rsa \
-oCiphers=+aes128-cbc \
user1@172.16.128.11
немесе
student@debian:~$ sudo nano ~/.ssh/config
Host 172.16.128.11
    KexAlgorithms +diffie-hellman-group1-sha1
    HostKeyAlgorithms +ssh-rsa
    PubkeyAcceptedAlgorithms +ssh-rsa
    Ciphers +aes128-cbc
CTRL+O, ENTER, CTRL+X
student@debian:~$ ssh user1@172.16.128.11
Accounting
student@ubuntu:~$ sudo ls -l /var/log/freeradius/radacct/
student@ubuntu:~$ sudo ls -l /var/log/freeradius/radacct/172.16.128.11/
student@ubuntu:~$ tail -f /var/log/freeradius/radacct/172.16.128.11/detail-YYYYMMDD[R1] aaa
      recording-scheme RADIUS
      recording-mode radius LAN1
      quit[R1] aaa
      domain LAB.LOCAL
      command-recording-scheme RADIUS
      quit[R1] command-privilege level 15 recording-scheme RADIUS[R3] stelnet 172.16.128.11
     <R1> system-view
     [R1] display versionstudent@ubuntu:~$ sudo ls -l /var/log/freeradius/radacct/172.16.128.11/
student@ubuntu:~$ tail -f /var/log/freeradius/radacct/172.16.128.11/detail-YYYYMMDD[R1] radius-server template LAN1
     radius-server source interface g0/0/0
     quit
Access Control List (ACL)
[R1] acl 2000
      rule permit source 172.16.128.101 0.0.0.0
      rule permit source 172.16.128.102 0.0.0.0
      rule deny source any
      quit
[R1] user-interface vty 0 4
      acl 2000 inbound
      quit
Idle-Timeout (Автоматты түрде сессияны жабу)
[R1] user-interface vty 0 4
      idle-timeout 10 0  # 10 минут, 0 секунд
      quit
