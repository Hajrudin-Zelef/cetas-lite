---
id: collect-261001-automatisation-infra/automatisation-infra/github-baltah666-netbox-pyats-3
title: "Description: This script generates a testbed file based on the Netbox data"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-automatisation-infra/github-baltah666-netbox-pyats.md
source_anchor: ""
source_lines: [298, 444]
sha256: 43cf92774daf3f36248b41dd19ff7de077c31aa323738c08f3aae9cdb825c982
---

# Description: This script generates a testbed file based on the Netbox data

```
genie learn routing ospf --testbed-file testbed.yaml --output post-change                  
Learning '['routing', 'ospf']' on devices '['CSR1', 'CSR2']'
100%|████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 2/2 [00:11<00:00,  5.84s/it]
+==============================================================================+
| Genie Learn Summary for device CSR1                                          |
+==============================================================================+
|  Connected to CSR1                                                           |
|  -   Log: post-change/connection_CSR1.txt                                    |
|------------------------------------------------------------------------------|
|  Learnt feature 'routing'                                                    |
|  -  Ops structure:  post-change/routing_iosxe_CSR1_ops.txt                   |
|  -  Device Console: post-change/routing_iosxe_CSR1_console.txt               |
|------------------------------------------------------------------------------|
|  Learnt feature 'ospf'                                                       |
|  -  Ops structure:  post-change/ospf_iosxe_CSR1_ops.txt                      |
|  -  Device Console: post-change/ospf_iosxe_CSR1_console.txt                  |
|==============================================================================|
+==============================================================================+
| Genie Learn Summary for device CSR2                                          |
+==============================================================================+
|  Connected to CSR2                                                           |
|  -   Log: post-change/connection_CSR2.txt                                    |
|------------------------------------------------------------------------------|
|  Learnt feature 'routing'                                                    |
|  -  Ops structure:  post-change/routing_iosxe_CSR2_ops.txt                   |
|  -  Device Console: post-change/routing_iosxe_CSR2_console.txt               |
|------------------------------------------------------------------------------|
|  Learnt feature 'ospf'                                                       |
|  -  Ops structure:  post-change/ospf_iosxe_CSR2_ops.txt                      |
|  -  Device Console: post-change/ospf_iosxe_CSR2_console.txt                  |
|==============================================================================|
```
Now let's use the `genie diff` command to tell us what changed in our network, and we will store the output in a directory called `routing-diff`:

```
genie diff pre-change post-change --output diff 
1it [00:00, 333.44it/s]
+==============================================================================+
| Genie Diff Summary between directories pre-change/ and post-change/          |
+==============================================================================+
|  File: routing_iosxe_CSR2_ops.txt                                            |
|   - Identical                                                                |
|------------------------------------------------------------------------------|
|  File: ospf_iosxe_CSR1_ops.txt                                               |
|   - Identical                                                                |
|------------------------------------------------------------------------------|
|  File: routing_iosxe_CSR1_ops.txt                                            |
|   - Diff can be found at diff/diff_routing_iosxe_CSR1_ops.txt                |
|------------------------------------------------------------------------------|
|  File: ospf_iosxe_CSR2_ops.txt                                               |
|   - Diff can be found at diff/diff_ospf_iosxe_CSR2_ops.txt                   |
|------------------------------------------------------------------------------|
```
There are differences in the outputs of both the `routing_iosxe_CSR1_ops.txt` and `ospf_iosxe_CSR2_ops.txt` files and we can view the contents of those files to see what has changed:

```
cat diff/diff_routing_iosxe_CSR1_ops.txt 
--- pre-change/routing_iosxe_CSR1_ops.txt
+++ post-change/routing_iosxe_CSR1_ops.txt
 info:
  vrf:
   default:
    address_family:
     ipv4:
      routes:
-      2.2.2.2/32:
-       active: True
-       metric: 2
-       next_hop:
-        next_hop_list:
-         1:
-          index: 1
-          next_hop: 192.168.1.2
-          outgoing_interface: GigabitEthernet2
-          updated: 03:32:43
-       route: 2.2.2.2/32
-       route_preference: 110
-       source_protocol: ospf
-       source_protocol_codes: O%  
```
```
cat diff/diff_ospf_iosxe_CSR2_ops.txt   
--- pre-change/ospf_iosxe_CSR2_ops.txt
+++ post-change/ospf_iosxe_CSR2_ops.txt
 info:
  vrf:
   default:
    address_family:
     ipv4:
      instance:
       1:
        areas:
         0.0.0.0:
          interfaces:
-          Loopback0:
-           bfd:
-            enable: False
-           cost: 1
-           demand_circuit: False
-           enable: True
-           interface_type: loopback
-           name: Loopback0% 
```
As we can see, CSR1 has lost the the route to the `2.2.2.2/32` prefix and on CSR2 the interface `lo0` has been removed from `OSPF area 0`.

As you would expect, it is also very easy to work with pyATS and python. Type `python` in your terminal, and then start by importing the `load` module from the `genie.testbed` library and load the testbed file:

```
from genie.testbed import load
testbed = load('testbed.yaml') 
```
Then connect to the CSR1 device:

```
device = testbed.devices['CSR1']
device.connect()
```
Then parse the command `show ip interface brief` and print the output nicely using the `pprint` library:

```
parsed_data = device.parse('show ip interface brief')
from pprint import pprint 
pprint(parsed_data)
{'interface': {'GigabitEthernet1': {'interface_is_ok': 'YES',
                                    'ip_address': '10.0.0.15',
                                    'method': 'manual',
                                    'protocol': 'up',
                                    'status': 'up'},
               'GigabitEthernet2': {'interface_is_ok': 'YES',
                                    'ip_address': '192.168.1.1',
                                    'method': 'manual',
                                    'protocol': 'up',
                                    'status': 'up'},
               'Loopback0': {'interface_is_ok': 'YES',
                             'ip_address': '1.1.1.1',
                             'method': 'manual',
                             'protocol': 'up',
                             'status': 'up'}}}
```
Lastly, close the connection to the device:

```
device.disconnect()
```
As you can imagine, the possibilities are endless and you can see an example of a script that gets and prints OSPF Neighbor table for a given device in the file `ospf_neighbor_table.py`. Run the script and you will get a nicely formatted table of the OSPF neighbors for CRS1:

