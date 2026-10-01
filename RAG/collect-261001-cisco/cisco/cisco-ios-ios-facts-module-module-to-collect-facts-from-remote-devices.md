---
id: collect-261001-cisco/cisco/cisco-ios-ios-facts-module-module-to-collect-facts-from-remote-devices
title: "cisco.ios.ios_facts module – Module to collect facts from remote devices."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/cisco-ios-ios-facts-module-module-to-collect-facts-from-remote-devices.md
source_anchor: ""
source_lines: [1, 96]
sha256: 80940e57191256f52fd87507bee0d36917394aa51eecd3a63bec44f76b0c2fa7
---

# cisco.ios.ios_facts module – Module to collect facts from remote devices.

Note

This module is part of the cisco.ios collection (version 11.5.1).

You might already have this collection installed if you are using the `ansible` package.
It is not included in `ansible-core`.
To check whether it is installed, run `ansible-galaxy collection list`.

To install it, use: `ansible-galaxy collection install cisco.ios`.

To use it in a playbook, specify: `cisco.ios.ios_facts`.

New in cisco.ios 1.0.0

- Collects a base set of device facts from a remote device that is running IOS. This module prepends all of the base network fact keys with `ansible_net_<fact>` . The facts module will always collect a base set of facts from the device and can enable or disable collection of additional facts.

Aliases: facts

| Parameter | Comments | 
|---|---|
|  | When ‘True’ a list of network resources for which resource modules are available will be provided. **Choices:**  | 
|  | When supplied, this argument will restrict the facts collected to a given subset. Possible values for this argument include all and the resources like interfaces, vlans etc. Can specify a list of values to include a larger subset. Values can also be used with an initial `!` to specify that a specific subset should not be collected. Valid subsets are ‘hsrp_interfaces’, ‘bgp_global’, ‘l3_interfaces’, ‘lag_interfaces’, ‘ntp_global’, ‘acls’, ‘hostname’, ‘interfaces’, ‘lldp_interfaces’, ‘logging_global’, ‘ospf_interfaces’, ‘ospfv2’, ‘prefix_lists’, ‘static_routes’, ‘acl_interfaces’, ‘all’, ‘bgp_address_family’, ‘l2_interfaces’, ‘lacp’, ‘lacp_interfaces’, ‘lldp_global’, ‘ospfv3’, ‘snmp_server’, ‘vlans’, ‘service’. | 
|  | When supplied, this argument restricts the facts collected to a given subset. Possible values for this argument include `all` ,`min` ,`default` ,`hardware` ,`config` , and`interfaces` . Specify a list of values to include a larger subset. Use a value with an initial `!` to collect all facts except that subset. **Default:**`["min"]` | 

Note

- Tested against Cisco IOSXE Version 17.3 on CML and IOS 15.6 for L2 specific resource.
- Facts gathering for L3 devices are supposed to produce blank output for unsupported resources like vlan.
- This module works with connection `network_cli` . See https://docs.ansible.com/ansible/latest/network/user_guide/platform_ios.html
- For more information on using Ansible to manage network devices see the :ref:`Ansible Network Guide <network_guide>`
- For more information on using Ansible to manage Cisco devices see the `Cisco integration page <https://www.ansible.com/integrations/networks/cisco>`_.

```
- name: Gather all legacy facts
  cisco.ios.ios_facts:
    gather_subset: all
- name: Gather only the config and default facts
  cisco.ios.ios_facts:
    gather_subset:
      - config
- name: Do not gather hardware facts
  cisco.ios.ios_facts:
    gather_subset:
      - "!hardware"
- name: Gather legacy and resource facts
  cisco.ios.ios_facts:
    gather_subset: all
    gather_network_resources: all
- name: Gather only the interfaces resource facts and no legacy facts
  cisco.ios.ios_facts:
    gather_subset:
      - "!all"
      - "!min"
    gather_network_resources:
      - interfaces
- name: Gather interfaces resource and minimal legacy facts
  cisco.ios.ios_facts:
    gather_subset: min
    gather_network_resources: interfaces
- name: Gather L2 interfaces resource and minimal legacy facts
  cisco.ios.ios_facts:
    gather_subset: min
    gather_network_resources: l2_interfaces
- name: Gather L3 interfaces resource and minimal legacy facts
  cisco.ios.ios_facts:
    gather_subset: min
    gather_network_resources: l3_interfaces
```
Common return values are documented here, the following are the fields unique to this module:

| Key | Description | 
|---|---|
|  | All IPv4 addresses configured on the device **Returned:** when interfaces is configured | 
|  | All IPv6 addresses configured on the device **Returned:** when interfaces is configured | 
|  | The name of the transport **Returned:** always | 
|  | The current active config from the device **Returned:** when config is configured | 
|  | The current CPU utilization of the device **Returned:** when hardware is configured | 
|  | All file system names available on the device **Returned:** when hardware is configured | 
|  | A hash of all file systems containing info about each file system (e.g. free and total space) **Returned:** when hardware is configured | 
|  | The list of fact for network resource subsets collected from the device **Returned:** when the resource is configured | 
|  | The list of fact subsets collected from the device **Returned:** always | 
|  | The configured hostname of the device **Returned:** always | 
|  | The image file the device is running **Returned:** always | 
|  | A hash of all interfaces running on the system **Returned:** when interfaces is configured | 
|  | The operating system type (IOS or IOS-XE) running on the remote device **Returned:** always | 
|  | The available free memory on the remote device in MiB **Returned:** when hardware is configured | 
|  | The total memory on the remote device in MiB **Returned:** when hardware is configured | 
|  | The model name returned from the device **Returned:** always | 
|  | The list of CDP and LLDP neighbors from the remote device. If both, CDP and LLDP neighbor data is present on one port, CDP is preferred. **Returned:** when interfaces is configured | 
|  | The Python version Ansible controller is using **Returned:** always | 
|  | The serial number of the remote device **Returned:** always | 
|  | The model names of each device in the stack **Returned:** when multiple devices are configured in a stack | 
|  | The serial numbers of each device in the stack **Returned:** when multiple devices are configured in a stack | 
|  | The operating system version running on the remote device **Returned:** always |
