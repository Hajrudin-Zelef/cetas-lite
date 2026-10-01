---
id: collect-261001-huawei/huawei/guides-huawei-huawei-ssh-activation-44ec4464
title: "guides-huawei-huawei-ssh-activation-44ec4464"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/guides-huawei-huawei-ssh-activation-44ec4464.md
source_anchor: ""
source_lines: [1, 15]
sha256: 52ed0eb0ae7843c5197495a5ad0e3c37292493f14a985f3210a8642378b2800e
---

# guides-huawei-huawei-ssh-activation-44ec4464

Instead of using Telnet you should use SSH to get access to a Huawei device (switch, router, wlan controller, etc.). This can be done in just four steps, both for the VRP operating system and for the new YunShan, and is described here.
SSH (or stelnet in Huawei language) access is not activated per default. Only a few configurations are necessary to activate it, but the commands are slightly different depending on the operating system (VRP or YunShan). Which operating system is used can be checked with  the command display version.
Activate SSH on Huawei VRP
| 1. | Generate RSA key pair <HUAWEI> system-view[HUAWEI] rsa local-key-pair create | 
| 2. | Create local user <HUAWEI > system-view [ HUAWEI ] aaa [ HUAWEI -aaa] local-user netcamp password cipher Netcamp2023! [ HUAWEI -aaa] local-user netcamp privilege level 3 [ HUAWEI -aaa] local-user netcamp service-type ssh | 
| 3. | Start SSH server and give user SSH access <HUAWEI> system-view[HUAWEI] stelnet server enable[HUAWEI] ssh user netcamp[HUAWEI] ssh user netcamp authentication-type password[HUAWEI] ssh user netcamp service-type stelnet | 
| 4. | Just allow SSH access on VTY lines <HUAWEI> system-view[HUAWEI] user-interface vty 0 4[HUAWEI-ui-vty0-4] authentication-mode aaa[HUAWEI-ui-vty0-4] protocol inbound ssh | 
Activate SSH on Huawei YunShan
| 1. | Generate RSA key pair <HUAWEI> system-view[HUAWEI] rsa local-key-pair create | 
| 2. | Create local user <HUAWEI> system-view[HUAWEI] aaa[HUAWEI-aaa] local-user netcamp password irreversible-cipher Netcamp2023!Info: The initial password of netcamp must be changed during login.[HUAWEI-aaa] local-user netcamp privilege level 3Warning: This operation may affect online users and will change the user privilege level, Continue? [Y/N]:y[HUAWEI-aaa] local-user netcamp service-type ssh | 
| 3. | Start SSH server, configure source interface and give user SSH access <HUAWEI> system-view[HUAWEI] stelnet server enable[HUAWEI] ssh server-source all-interfaceWarning: SSH server source configuration will take effect in the next login. Continue? [Y/N]:y[HUAWEI] ssh user netcamp[HUAWEI] ssh user netcamp authentication-type password[HUAWEI] ssh user netcamp service-type stelnet | 
| 4. | Just allow SSH access on VTY lines <HUAWEI> system-view[HUAWEI] user-interface vty 0 4[HUAWEI-ui-vty0-4] authentication-mode aaa[HUAWEI-ui-vty0-4] protocol inbound ssh | 
➡️ The password must be changed the first time you log in. If this is not desired, it can be deactivated, which is described in guide Disable password change prompt on Huawei device.
➡️ If you try to access the device from another Huawei device as a client you need to activate it first with ssh client first-time enable.
Leave a Comment
