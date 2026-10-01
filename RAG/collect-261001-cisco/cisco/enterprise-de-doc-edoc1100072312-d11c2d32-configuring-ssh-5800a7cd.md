---
id: collect-261001-cisco/cisco/enterprise-de-doc-edoc1100072312-d11c2d32-configuring-ssh-5800a7cd
title: "enterprise-de-doc-edoc1100072312-d11c2d32-configuring-ssh-5800a7cd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-de-doc-edoc1100072312-d11c2d32-configuring-ssh-5800a7cd.md
source_anchor: ""
source_lines: [1, 42]
sha256: 58726d95657aa5d57c0e99fd97d8f9ddb3f34363e521eda34ffc349a09d8cb61
---

# enterprise-de-doc-edoc1100072312-d11c2d32-configuring-ssh-5800a7cd

Unternehmen
SSH commands must be configured on all switches on the network. SSH is used for communication between CE switches and VMware vRNI. The configuration on a CE switch is used as an example. The configurations on other CE switches are similar.
<HUAWEI> system-view
[~HUAWEI] user-interface maximum-vty 21
[~HUAWEI] user-interface vty 0 20
[*HUAWEI-hi-vty0-20] idle-timeout 0
[*HUAWEI-ui-vty0-20] authentication-mode aaa
[*HUAWEI-ui-vty0-20] protocol inbound all
[*HUAWEI-ui-vty0-20] commit
[~HUAWEI-ui-vty0-20] quit
[~HUAWEI] interface meth 0/0/0
[*HUAWEI-MEth0/0/0] ip address 192.105.146.33 24
[*HUAWEI-MEth0/0/0] commit
[~HUAWEI-MEth0/0/0] quit
[~HUAWEI] aaa
[~HUAWEI-aaa] local
[~HUAWEI-aaa] undo local-user policy security-enhance
[*HUAWEI-aaa] commit
[~HUAWEI-aaa] local-user netconftest password irreversible-cipher huaweiDC
[*HUAWEI-aaa] local-user netconftest service-type telnet ssh
[*HUAWEI-aaa] local-user netconftest level 3
[*HUAWEI-aaa] local-user netconftest user-group manage-ug
[*HUAWEI-aaa] commit
[*HUAWEI-aaa] quit
[~HUAWEI] ssh user netconftest
[*HUAWEI] ssh user netconftest authentication-type password
[*HUAWEI] ssh user netconftest service-type all
[*HUAWEI] ssh authorization-type default aaa
[*HUAWEI] commit
[~HUAWEI] stelnet server enable
[*HUAWEI] snetconf server enbale
[*HUAWEI] commit
[~HUAWEI] rsa local-key-pair create
The key name will be:HUAWEI_Host
% RSA keys defined for HUAWEI_Host already exist.
Confirm to replace them？ Please select [Y/N]: y
The rangge of public key size is (2048 ~ 2048).
NOTE: Key pair generation will take a short while.
[*HUAWEI] commit
[~HUAWEI] undo telnet server disable
[*HUAWEI] commit
Select the content with the mouse pointer to quickly report the problem.
