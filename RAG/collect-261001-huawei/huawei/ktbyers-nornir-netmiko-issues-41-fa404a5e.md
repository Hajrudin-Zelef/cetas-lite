---
id: collect-261001-huawei/huawei/ktbyers-nornir-netmiko-issues-41-fa404a5e
title: "ktbyers-nornir-netmiko-issues-41-fa404a5e"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["energy", "ethernet"]
source: docs/RAG/collect-261001-huawei/ktbyers-nornir-netmiko-issues-41-fa404a5e.md
source_anchor: ""
source_lines: [1, 129]
sha256: 8db37baafd4ba038960f698900ee59b405513b7e85180a6e58cf8f07ba697998
---

# ktbyers-nornir-netmiko-issues-41-fa404a5e

- 
                Notifications
    You must be signed in to change notification settings
- Fork 28
Description
The batch configures behavior is below:
vvvv netmiko_send_config ** changed : True vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv INFO
system-view
Enter system view, return user view with Ctrl+Z.
[SW]interface range Gi 0/0/1 to Gi 0/0/9
[SW-port-group]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/1]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/2]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/3]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/4]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/5]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/6]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/7]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/8]energy-efficient-ethernet enable
[SW-GigabitEthernet0/0/9]energy-efficient-ethernet enable
[SW-port-group]port-auto-sleep enable
[SW-GigabitEthernet0/0/1]port-auto-sleep enablereturn
^^^^ END netmiko_send_config ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
The echo is incomplete, but the configures are successful.
I tried to adjust the delay_factor, max_loops, and set the fast_cli to False, but still didn't work.
How to optimize?
Activity
@sgqjt What should it look like?
Never mind on my question above...you need to avoid using interface range this will generally cause problems for automation.
So convert your configuration changes to be a list of configuration commands that cover all of the relevant interfaces. Also you should full specify the interface names.
Let me know if that helps with your issue.
Remember you can just use a for-loop in Python to construct the list of configuration commands (instead of using interface range).
I want to improve the efficiency of the program.
I tested using interface range and a list of configurations; it takes about 20 seconds when using interface range and 100 seconds when using a list of configurations for nine interfaces.
We need to configure 322 switches and 20300 interfaces; it can consume a lot of time.
@sgqjt The number of switches shouldn't matter much as with Nornir you can just do a whole bunch of them concurrently.
You are free to use interface range, but you are going to have a lot of difficulty to make it work from an automation perspective (and I am just going to say, "don't do that" when it doesn't work as is happening here).
You didn't include your code, but there are likely other ways that performance per switch could be meaningfully improved.
@ktbyers Thanks for your advice.
I used for-loop instead of interface range; it works fine.
Here is my code:
hosts.yaml
SW:
    hostname: 10.0.0.1
    groups:
        - Lab
    data:
        stage: Lab
        eth_trunk: interface Gi 0/0/10
        ap-interfaces:
            - 0.0.1-2
        total-interfaces:
            - 0.0.1-9.36.45-48
config.py
from nornir.core.task import Task, Result
from typing import Any, List, Optional
from nornir_netmiko.connections import CONNECTION_NAME
import getpass
from nornir import InitNornir
from nornir_netmiko import netmiko_save_config, netmiko_send_command, netmiko_send_config
from nornir_utils.plugins.functions import print_result
from nornir.core.filter import F
def config_cmd(
    task: Task, 
    vlan: Optional[str] = None,
    commands: Optional[List[str]] = None,
    interfaces: Optional[str] = None,
    dis_command: Optional[str] = None,
    config_filepath: str = None,
    **kwargs: Any
) -> Result:
    
    # Configure commands
    configs = []
    for i in task.host.data[interfaces]:
        cfg_items = []
        #split ports with '.'
        ports = i[4:]
        multiPort = ports.split('.')
        j = 0
        #ports = []
        while j < len(multiPort):
            if len(multiPort[j]) > 2:
                #split ports with '-'
                temp = multiPort[j].split('-')
                #generate ports list
                port_list = list(range(int(temp[0]),int(temp[1])+1))
                for port in port_list:
                    # append range ports to list
                    cfg_items.append('interface Gi '+i[0]+'/'+i[2]+'/'+str(port))
            else:
                #append single port to list
                cfg_items.append('interface Gi '+i[0]+'/'+i[2]+'/'+multiPort[j])
            j += 1
        for i in cfg_items:
            configs.append(i)
            for k in commands:
                configs.append(k)
    task.run(
        task=netmiko_send_config,
        config_commands = configs,
    )
    # Save Configuration
    task.run(
        task = netmiko_save_config,
        confirm=True, 
        confirm_response="y"
    )
if __name__ == "__main__":
    
    username = input("Username: ")
    password = getpass.getpass("Password: ")
    nr = InitNornir(config_file="inventory/config.yaml")
    nr.inventory.defaults.username = username
    nr.inventory.defaults.password = password
    filter_stage = nr.filter(F(groups__contains = "Lab"))
    powerSave = [
        "energy-efficient-ethernet enable",
        "port-auto-sleep enable "
     ]
    disPower = "display power manage power-information"
    result = filter_stage.run(
        task = config_cmd,
        commands = powerSave,
        dis_command = disPower,
        interfaces = "total-interfaces"
    )
    print_result(result)
Thanks.
