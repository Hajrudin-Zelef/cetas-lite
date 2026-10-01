---
id: collect-261001-huawei/huawei/ktbyers-netmiko-issues-517-0d9b0396
title: "ktbyers-netmiko-issues-517-0d9b0396"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ktbyers-netmiko-issues-517-0d9b0396.md
source_anchor: ""
source_lines: [1, 48]
sha256: 88f5b43cf79feaa2e31f9f7a6a7740a6af584b06d8ff53edd5cb95dade77a436
---

# ktbyers-netmiko-issues-517-0d9b0396

- 
                Notifications
    You must be signed in to change notification settings
- Fork 1.4k
Description
Hello,
First of all - a real useful library.
Now,my problem: I am trying to configure some Huawei routers (ATN,PTN,NE40E - all with VRP). I can run some commands as "display version" but I am stuck with configuration.
Version - 1.4.1
import netmiko
from netmiko import ConnectHandler
import time
device = ConnectHandler(
device_type='huawei',
ip='10.255.82.136',
username='aaaaa',
password='bbbb'
)
#output = device.send_command('disp version')
output = device.send_command('screen-length 0 temporary')
time.sleep(1)
print output
output = device.config_mode()
print output
output = device.send_command('interface G0/2/10\n')
At the last command the script gets stuck - below is the output (I stop it via CTRL+C,after 10-15 seconds of waiting).
python config_etn_conf.py
Info: The configuration takes effect on the current user terminal interface only.
system-view
Enter system view, return user view with Ctrl+Z.
[eTN_BU_03316_AC_09]
^CTraceback (most recent call last):
File "config_etn_conf.py", line 19, in 
output = device.send_command('interface G0/2/10\n')
File "/usr/lib/python2.7/site-packages/netmiko/base_connection.py", line 811, in send_command
time.sleep(delay_factor * .2)
KeyboardInterrupt
I've tried to figure out how to solve it (checked base_connection and huawei_ssh....no luck. Somebody can help me ?
Thank you,
Bogdan T
Activity
@BogdanTomoiu Use the send_config_set() or send_config_from_file() methods:
config_commands = [ 'logging buffered 20000',
                    'logging buffered 20010',
                    'no logging console' ]
output = net_connect.send_config_set(config_commands)
print(output)
Thank you very much. Works very nice - problem solved !
