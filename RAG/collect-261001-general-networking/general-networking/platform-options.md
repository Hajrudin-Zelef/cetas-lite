---
id: collect-261001-general-networking/general-networking/platform-options
title: "Platform Options"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-options.md
source_anchor: ""
source_lines: [1, 40]
sha256: 242a87470bc487524b68055bbd4ccf4222e7ec2e44d4efab0d9c561e49dff6bd
---

# Platform Options

Some Ansible Network platforms support multiple connection types, privilege escalation (`enable` mode), or other options. The pages in this section offer standardized guides to understanding available options on each network platform. We welcome contributions from community-maintained platforms to this section.

## Settings by Platform

|  |  | `ansible_connection:` settings available |  |  |  | 
|---|---|---|---|---|---|
| Network OS | `ansible_network_os:` | network_cli | netconf | httpapi | local | 
| Arista EOS [†] | `arista.eos.eos` | ✓ |  | ✓ | ✓ | 
| Ciena SAOS6 | `ciena.saos6.saos6` | ✓ |  |  | ✓ | 
| Cisco ASA [†] | `cisco.asa.asa` | ✓ |  |  | ✓ | 
| Cisco IOS [†] | `cisco.ios.ios` | ✓ |  |  | ✓ | 
| Cisco IOS XR [†] | `cisco.iosxr.iosxr` | ✓ |  |  | ✓ | 
| Cisco NX-OS [†] | `cisco.nxos.nxos` | ✓ |  | ✓ | ✓ | 
| Cloudengine OS | `community.network.ce` | ✓ | ✓ |  | ✓ | 
| Dell OS6 | `dellemc.os6.os6` | ✓ |  |  | ✓ | 
| Dell OS9 | `dellemc.os9.os9` | ✓ |  |  | ✓ | 
| Dell OS10 | `dellemc.os10.os10` | ✓ |  |  | ✓ | 
| Ericsson ECCLI | `community.network.eric_eccli` | ✓ |  |  | ✓ | 
| Extreme EXOS | `community.network.exos` | ✓ |  | ✓ |  | 
| Extreme IronWare | `community.network.ironware` | ✓ |  |  | ✓ | 
| Extreme NOS | `community.network.nos` | ✓ |  |  |  | 
| Extreme SLX-OS | `community.network.slxos` | ✓ |  |  |  | 
| Extreme VOSS | `community.network.voss` | ✓ |  |  |  | 
| F5 BIG-IP |  |  |  |  | ✓ | 
| F5 BIG-IQ |  |  |  |  | ✓ | 
| Junos OS [†] | `junipernetworks.junos.junos` | ✓ | ✓ |  | ✓ | 
| Lenovo CNOS | `community.network.cnos` | ✓ |  |  | ✓ | 
| Lenovo ENOS | `community.network.enos` | ✓ |  |  | ✓ | 
| Meraki |  |  |  |  | ✓ | 
| MikroTik RouterOS | `community.network.routeros` | ✓ |  |  |  | 
| Nokia SR OS |  |  |  |  | ✓ | 
| Pluribus Netvisor | `community.network.netvisor` | ✓ |  |  |  | 
| Ruckus ICX | `community.network.icx` | ✓ |  |  |  | 
| VyOS [†] | `vyos.vyos.vyos` | ✓ |  |  | ✓ | 
| Westermo WeOS 4 | `community.network.weos4` | ✓ |  |  |  | 
| OS that supports Netconf [†] | `<network-os>` |  | ✓ |  | ✓ | 

**[†]** Maintained by Ansible Network Team
