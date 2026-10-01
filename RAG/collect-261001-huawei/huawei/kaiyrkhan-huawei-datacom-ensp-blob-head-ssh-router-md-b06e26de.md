---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-ssh-router-md-b06e26de
title: "kaiyrkhan-huawei-datacom-ensp-blob-head-ssh-router-md-b06e26de"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-ssh-router-md-b06e26de.md
source_anchor: ""
source_lines: [1, 64]
sha256: 40dbbbd2a7f299f1feef61df1db9da30c49a3247fe98f9196c3ade2d4982d9e0
---

# kaiyrkhan-huawei-datacom-ensp-blob-head-ssh-router-md-b06e26de

ЕСКЕРТУ! Бұл зертханалық жұмыс нақты физикалық құрылғыны (Physical Device) қолданып жасалған!
Құрылғының моделі: Huawei AR6140E-9G-2AC Router
Yellow - Layer 3 Routed Port
Blue - Layer 2 Switch Port
Red - Management (MGMT) Port
Login authentication
Warning: An initial username and password are required for the first login via the console. Set a username and password and keep them safe. Otherwise you will not be able to login via the console.
New Username: admin
Password: P@s$w0rd_&1234
Confirm password: P@s$w0rd_&1234
Info: Configuration console exit, please retry to log on
Login authentication
Username: admin
Password: P@s$w0rd_&1234
Warning: Auto-Config is working. Before configuring the device, stop Auto-Config. If you perform configurations when Auto-Config is running, the DHCP, routing, DNS, and VTY configurations will be lost. Do you want to stop Auto-Config? [y/n]: y
Info: Auto-Config has been stopped.
display version
Төмендегі топологияда көрсетілгендей, PC мен Router-ді Copper кабелмен байланыстырып қосамыз!
display interface briefdisplay ip interface briefping 192.168.1.1
Қосымша ақпарат!
interface GigabitEthernet 0/0/2
portswitch
port link-type access
port default vlan 1
Configure Local User Authentication and Authorization
aaa
 local-user student password irreversible-cipher Huawei@123
 local-user student service-type terminal ssh
 local-user student privilege level 15
Configure VTY Lines
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound ssh
Қосымша ақпарат!
жеке (individual) құқық (privilege) - student қолданушыға ғана тиесілі
aaa
local-user student privilege level 15
жалпы (Global) құқық (privilege) - барлық қолданушыға қатысты
user-interface vty 0 4
user privilege level 15
Generate RSA Key
rsa local-key-pair create
Warning: Confirm to replace them! Continue? [Y/N] Y
Input the bits in the modulus[default = 2048]: 2048
SSH server Permit interface
ssh server permit interface Vlanif 1
немесе
ssh server permit interface GigabitEthernet 0/0/2
немесе
ssh server permit interface all
Қосымша ақпарат!
ssh server-source -i Vlanif 1
ssh server-source all-interface
[Huawei] ssh user student
[Huawei] ssh user student service-type stelnet
[Huawei] ssh user student authentication-type password
Enable SSH
stelnet server enable
Info: Succeeded in starting the STELNET server.
display ssh server status
display current-configuration | include ssh
display current-configuration | include stelnet
Verification
ssh student@192.168.1.1
