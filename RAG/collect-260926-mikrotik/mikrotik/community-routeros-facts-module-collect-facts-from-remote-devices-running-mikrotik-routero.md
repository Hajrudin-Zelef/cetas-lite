---
id: collect-260926-mikrotik/mikrotik/community-routeros-facts-module-collect-facts-from-remote-devices-running-mikrotik-routero
title: "community.routeros.facts module – Collect facts from remote devices running MikroTik RouterOS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/community-routeros-facts-module-collect-facts-from-remote-devices-running-mikrotik-routeros.md
source_anchor: ""
source_lines: [1, 61]
sha256: 111ce0d911c092a8bc9efd22f128243576737892f6e162233b716e3a2825f68e
---

# community.routeros.facts module – Collect facts from remote devices running MikroTik RouterOS

Note

This module is part of the community.routeros collection (version 2.1.0).

You might already have this collection installed if you are using the `ansible` package.
It is not included in `ansible-core`.
To check whether it is installed, run `ansible-galaxy collection list`.

To install it, use: `ansible-galaxy collection install community.routeros`.

To use it in a playbook, specify: `community.routeros.facts`.

- Collects a base set of device facts from a remote device that is running RouterOS. This module prepends all of the base network fact keys with `ansible_net_<fact>` .  The facts module will always collect a base set of facts from the device and can enable or disable collection of additional facts.

| Parameter | Comments | 
|---|---|
|  | When supplied, this argument will restrict the facts collected to a given subset.  Possible values for this argument include `all` ,`hardware` ,`config` ,`interfaces` , and`routing` . Can specify a list of values to include a larger subset. Values can also be used with an initial `!` to specify that a specific subset should not be collected. Default: “!config” | 

```
- name: Collect all facts from the device
  community.routeros.facts:
    gather_subset: all
- name: Collect only the config and default facts
  community.routeros.facts:
    gather_subset:
      - config
- name: Do not collect hardware facts
  community.routeros.facts:
    gather_subset:
      - "!hardware"
```
Facts returned by this module are added/updated in the `hostvars` host facts and can be referenced by name just like any other host fact. They do not need to be registered in order to use them.

| Key | Description | 
|---|---|
|  | All IPv4 addresses configured on the device. Returned: *gather_subset* contains`interfaces` | 
|  | All IPv6 addresses configured on the device. Returned: *gather_subset* contains`interfaces` | 
|  | The CPU architecture of the device. Returned: *gather_subset* contains`default` | 
|  | A dictionary with BGP instance information. Returned: *gather_subset* contains`routing` | 
|  | A dictionary with BGP peer information. Returned: *gather_subset* contains`routing` | 
|  | A dictionary with BGP vpnv4 route information. Returned: *gather_subset* contains`routing` | 
|  | The current active config from the device. Returned: *gather_subset* contains`config` | 
|  | The current active config from the device in minimal form. This value is idempotent in the sense that if the facts module is run twice and the device’s config was not changed between the runs, the value is identical. This is achieved by running `/export` and stripping the timestamp from the comment in the first line. Returned: *gather_subset* contains`config` | 
|  | Current CPU load. Returned: *gather_subset* contains`default` | 
|  | The list of fact subsets collected from the device. Returned: always | 
|  | The configured hostname of the device. Returned: *gather_subset* contains`default` | 
|  | A hash of all interfaces running on the system. Returned: *gather_subset* contains`interfaces` | 
|  | The available free memory on the remote device in MiB. Returned: *gather_subset* contains`hardware` | 
|  | The total memory on the remote device in MiB. Returned: *gather_subset* contains`hardware` | 
|  | The model name returned from the device. Returned: *gather_subset* contains`default` | 
|  | The list of neighbors from the remote device. Returned: *gather_subset* contains`interfaces` | 
|  | A dictionary with OSPF instances. Returned: *gather_subset* contains`routing` | 
|  | A dictionary with OSPF neighbors. Returned: *gather_subset* contains`routing` | 
|  | A dictionary for routes in all routing tables. Returned: *gather_subset* contains`routing` | 
|  | The serial number of the remote device. Returned: *gather_subset* contains`default` | 
|  | The available disk space on the remote device in MiB. Returned: *gather_subset* contains`hardware` | 
|  | The total disk space on the remote device in MiB. Returned: *gather_subset* contains`hardware` | 
|  | The uptime of the device. Returned: *gather_subset* contains`default` | 
|  | The operating system version running on the remote device. Returned: *gather_subset* contains`default` |
