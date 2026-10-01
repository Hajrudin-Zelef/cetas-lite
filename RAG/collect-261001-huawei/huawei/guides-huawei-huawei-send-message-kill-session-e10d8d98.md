---
id: collect-261001-huawei/huawei/guides-huawei-huawei-send-message-kill-session-e10d8d98
title: "guides-huawei-huawei-send-message-kill-session-e10d8d98"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/guides-huawei-huawei-send-message-kill-session-e10d8d98.md
source_anchor: ""
source_lines: [1, 8]
sha256: fc9393c5e6c9faf0d0936d1aec613ad8f2199c505135db4856df23a7d5b530d5
---

# guides-huawei-huawei-send-message-kill-session-e10d8d98

You can send messages from one VTY session to another (e.g. SSH). This can be helpful to inform other users about important information. If you want to kill a VTY session this can also be done from the CLI. Both features are described here.
If you are connected to a Huawei device (switch, router, WLC, etc.) you can check with display users if somebody else is connected to this device. You can then send a message to this session, which will be displayed on the VTY line.
| 1. | Display all users connected to the device. In this case two users (philip and another-user) are connected on VTY line 0 and 1. The + in front shows my own VTY line (line 0 in this case): <HUAWEI>display users   User-Intf     Delay     Type     Network Address        AuthenStatus     AuthorcmdFlag+ 34 VTY 0    00:00:00     SSH      192.168.1.10           pass                   no        Username : philip  35 VTY 1    00:00:02     SSH      192.168.1.11           pass                   no        Username : another-user | 
| 2. | If you want to send a message to VTY line 1 where user "another-user" is connected, then you can do it with the following command (press enter to send): <HUAWEI>send vty 1Enter message, end with CTRL+Z or Enter; abort with CTRL+C:Please leave my device!Warning: Send the message? [Y/N]:y | 
| 3. | The message on VTY line 1 will look like this: <HUAWEI>Info: Receive a message from VTY0:Please leave my device! | 
If you want to kill the session on VTY line 1 then you can do it with the following command - it will immediately close the other session:
| <HUAWEI>free user-interface vty 1Warning: User interface VTY1 will be freed. Continue? [Y/N]:yInfo: User interface VTY1 is free. | 
Leave a Comment
