---
id: collect-261001-automatisation-infra/automatisation-infra/github-baltah666-netbox-pyats-2
title: "Description: This script generates a testbed file based on the Netbox data"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/github-baltah666-netbox-pyats.md
source_anchor: ""
source_lines: [185, 297]
sha256: e8135bdfb57af2b2424a0e6b23673870212e5d1610ea1898e34300571d13acd4
---

# Description: This script generates a testbed file based on the Netbox data

```
{
  "interface": {
    "GigabitEthernet1": {
      "interface_is_ok": "YES",
      "ip_address": "10.0.0.15",
      "method": "manual",
      "protocol": "up",
      "status": "up"
    },
    "GigabitEthernet2": {
      "interface_is_ok": "YES",
      "ip_address": "192.168.1.1",
      "method": "manual",
      "protocol": "up",
      "status": "up"
    },
    "Loopback0": {
      "interface_is_ok": "YES",
      "ip_address": "1.1.1.1",
      "method": "manual",
      "protocol": "up",
      "status": "up"
    }
  }
}
```
There are parsers available for a large number of network OS's (not just Cisco), and you can view the complete list here
In this example we will run the `genie parse` command to parse the output of the `show ip ospf neighbor` command and limit it to just device `CSR1`:

```
genie parse 'show ip ospf neighbor' --testbed-file testbed.yaml --device CSR1
  0%|                                                                                                                                                                | 0/1 [00:00<?, ?it/s]{
  "interfaces": {
    "GigabitEthernet2": {
      "neighbors": {
        "2.2.2.2": {
          "address": "192.168.1.2",
          "dead_time": "00:00:34",
          "priority": 1,
          "state": "FULL/DR"
        }
      }
    }
  }
}
```
This will output the parsed data to your terminal, but if you want to save the output just append the directory you wish to save it to with the `--output` switch. This will save both the `_console` file (the unstructured data) and the `_parsed` file (the structured data), along with the `connection_` log (the raw output of the full connection process) into the chosen directory:

```
genie parse 'show ip ospf neighbor' --testbed-file testbed.yaml --device CSR1 --output csr1
100%|████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 1/1 [00:01<00:00,  1.01s/it]
+==============================================================================+
| Genie Parse Summary for CSR1                                                 |
+==============================================================================+
|  Connected to CSR1                                                           |
|  -  Log: csr1/connection_CSR1.txt                                            |
|------------------------------------------------------------------------------|
|  Parsed command 'show ip ospf neighbor'                                      |
|  -  Parsed structure: csr1/CSR1_show-ip-ospf-neighbor_parsed.txt             |
|  -  Device Console:   csr1/CSR1_show-ip-ospf-neighbor_console.txt            |
|------------------------------------------------------------------------------|
```
In this example we will run the `genie learn` command to learn all about `routing` and `ospf` for both devices in our testbed file. We will also save the output into directory called `pre-change`. Genie will automagically create directories for us if they don't already exist.

For reference if you run the `genie learn all` command then for IOS-XE devices the list of features learned is:

```
acl, arp, bgp, device, dot1x, eigrp, fdb, hsrp, igmp, interface, isis, lag, lisp, lldp, mcast, mld, msdp, nd, ntp, ospf, pim, platform, prefix_list, rip, route_policy, routing, static_routing, stp, terminal, utils, vlan, vrf, vxlan, config
```
```
genie learn routing ospf --testbed-file testbed.yaml --output pre-change  
Learning '['routing', 'ospf']' on devices '['CSR1', 'CSR2']'
100%|████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 2/2 [00:11<00:00,  5.85s/it]
+==============================================================================+
| Genie Learn Summary for device CSR1                                          |
+==============================================================================+
|  Connected to CSR1                                                           |
|  -   Log: pre-change/connection_CSR1.txt                                     |
|------------------------------------------------------------------------------|
|  Learnt feature 'routing'                                                    |
|  -  Ops structure:  pre-change/routing_iosxe_CSR1_ops.txt                    |
|  -  Device Console: pre-change/routing_iosxe_CSR1_console.txt                |
|------------------------------------------------------------------------------|
|  Learnt feature 'ospf'                                                       |
|  -  Ops structure:  pre-change/ospf_iosxe_CSR1_ops.txt                       |
|  -  Device Console: pre-change/ospf_iosxe_CSR1_console.txt                   |
|==============================================================================|
+==============================================================================+
| Genie Learn Summary for device CSR2                                          |
+==============================================================================+
|  Connected to CSR2                                                           |
|  -   Log: pre-change/connection_CSR2.txt                                     |
|------------------------------------------------------------------------------|
|  Learnt feature 'routing'                                                    |
|  -  Ops structure:  pre-change/routing_iosxe_CSR2_ops.txt                    |
|  -  Device Console: pre-change/routing_iosxe_CSR2_console.txt                |
|------------------------------------------------------------------------------|
|  Learnt feature 'ospf'                                                       |
|  -  Ops structure:  pre-change/ospf_iosxe_CSR2_ops.txt                       |
|  -  Device Console: pre-change/ospf_iosxe_CSR2_console.txt                   |
|==============================================================================|
```
Let's make a small change to the `CSR2` router in our test network, and remove the `network 2.2.2.2 0.0.0.0 area 0` statement from the OSPF configuration:

```
CSR2#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
CSR2(config)#router ospf 1
CSR2(config-router)#no network 2.2.2.2 0.0.0.0 area 0
```
Next lets re-learn `routing` and `ospf` for both devices in our testbed file. We will also save the output into directory called `post-change`:

