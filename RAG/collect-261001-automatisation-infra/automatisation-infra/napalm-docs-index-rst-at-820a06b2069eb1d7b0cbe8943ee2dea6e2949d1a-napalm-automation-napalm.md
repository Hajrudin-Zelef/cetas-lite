---
id: collect-261001-automatisation-infra/automatisation-infra/napalm-docs-index-rst-at-820a06b2069eb1d7b0cbe8943ee2dea6e2949d1a-napalm-automation-napalm
title: "napalm-docs-index-rst-at-820a06b2069eb1d7b0cbe8943ee2dea6e2949d1a-napalm-automation-napalm"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/napalm-docs-index-rst-at-820a06b2069eb1d7b0cbe8943ee2dea6e2949d1a-napalm-automation-napalm.md
source_anchor: ""
source_lines: [1, 43]
sha256: 17d72b6fae6f2fc82b31a99fd4cd2704dcc25c7c51f4bd3f239f1bbc68e7fbea
---

# napalm-docs-index-rst-at-820a06b2069eb1d7b0cbe8943ee2dea6e2949d1a-napalm-automation-napalm

NAPALM (Network Automation and Programmability Abstraction Layer with Multivendor support) is a Python library that implements a set of functions to interact with different network device Operating Systems using a unified API.

NAPALM supports several methods to connect to the devices, to manipulate configurations or to retrieve data.

- Arista EOS
- Cisco IOS
- Cisco IOS-XR
- Cisco NX-OS
- Juniper JunOS

In addition to the core drivers napalm also supports community driven drivers. You can find more information about them here: :ref:`contributing-drivers`

You can select the driver you need by doing the following:

```
>>> from napalm import get_network_driver
>>> get_network_driver('eos')
<class napalm.eos.eos.EOSDriver at 0x10ebad6d0>
>>> get_network_driver('iosxr_netconf')
<class napalm.iosxr_netconf.iosxr_netconf.IOSXRNETCONFDriver at 0x10ad170f0>
>>> get_network_driver('iosxr')
<class napalm.iosxr.iosxr.IOSXRDriver at 0x10ec90050>
>>> get_network_driver('junos')
<class napalm.junos.junos.JunOSDriver at 0x10f8f61f0>
>>> get_network_driver('nxos')
<class napalm.nxos.nxos.NXOSDriver at 0x10f9304c8>
>>> get_network_driver('ios')
<class napalm.ios.ios.IOSDriver at 0x10f9b0738>
```
.. toctree::
   :maxdepth: 2
   installation/index
   tutorials/index
   validate/index
   support/index
   cli
   base
   yang
   logs
   integrations/index
   contributing/index
   development/index
   hackathons/index
