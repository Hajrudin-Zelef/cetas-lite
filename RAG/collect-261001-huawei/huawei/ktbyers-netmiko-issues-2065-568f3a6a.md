---
id: collect-261001-huawei/huawei/ktbyers-netmiko-issues-2065-568f3a6a
title: "ktbyers-netmiko-issues-2065-568f3a6a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/ktbyers-netmiko-issues-2065-568f3a6a.md
source_anchor: ""
source_lines: [1, 56]
sha256: 8a5e31b565297711b4d5b1422b2b2755b54feff46df0e45d7916a4144d063699
---

# ktbyers-netmiko-issues-2065-568f3a6a

- 
                Notifications
    You must be signed in to change notification settings
- Fork 1.4k
Closed
Description
Hello,
i am trying to configure and commit ne40e
device - HUAWEI NE40E-F1A-14H24Q
system - VRP (R) software, Version 8.190
netmiko 3.2.0
testing script:
from netmiko import ConnectHandler
device = {
'username': 'xxx',
'allow_agent': 'True',
'use_keys': 'True',
'key_file': '/home/xxx/.ssh/id_rsa.pub',
'host': 'xxx',
'device_type': 'huawei'
}
commands = ['int 100ge0/1/43', 'description asbr-test-script']
ssh_connection = ConnectHandler(**device)
print(ssh_connection.send_config_set(commands, exit_config_mode=False)) 
print(ssh_connection.commit())
send config set works ok, but commit not:
In [27]: print(ssh_connection.commit())                                                                               
---------------------------------------------------------------------------
AttributeError                            Traceback (most recent call last)
<ipython-input-27-f94a48d946e3> in <module>
----> 1 print(ssh_connection.commit())
~/.local/lib/python3.8/site-packages/netmiko/base_connection.py in commit(self)
   1924     def commit(self):
   1925         """Commit method for platforms that support this."""
-> 1926         raise AttributeError("Network device does not support 'commit()' method")
   1927 
   1928     def save_config(self, *args, **kwargs):
AttributeError: Network device does not support 'commit()' method
this workaround works:
commands = ['int 100ge0/1/43','description asbr-test-script', 'commit']
print(ssh_connection.send_config_set(commands))  
In [31]: print(ssh_connection.send_config_set(commands))                                                              
system-view
Enter system view, return user view with return command.
[~hua.test]int 100ge0/1/43
[~hua.test-100GE0/1/43]description asbr-test-script
[*hua.test-100GE0/1/43]commit
[~hua.test-100GE0/1/43]return
<hua.test>
In [32]: 
Activity
Is this a huawei_vrpv8 or something else?
There is no commit method in the standard huawei driver. The huawei_vrpv8 device_type has a commit method, though.
Is this a huawei_vrpv8 or something else?
There is no commit method in the standard huawei driver. The huawei_vrpv8 device_type has a commit method, though.
sorry for late reply, that solved issue, thanks!
