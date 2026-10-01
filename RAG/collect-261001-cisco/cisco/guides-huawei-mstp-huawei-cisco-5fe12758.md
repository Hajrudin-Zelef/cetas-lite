---
id: collect-261001-cisco/cisco/guides-huawei-mstp-huawei-cisco-5fe12758
title: "guides-huawei-mstp-huawei-cisco-5fe12758"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/guides-huawei-mstp-huawei-cisco-5fe12758.md
source_anchor: ""
source_lines: [1, 12]
sha256: 98457861aa23bad411b7e7c129b389727178a11813db05235ae35c48bd472308
---

# guides-huawei-mstp-huawei-cisco-5fe12758

In some cases you need to create Layer-2 domains with switches from different vendors. This article gives you information how to connect Huawei and Cisco switches together with MSTP (IEEE 802.1s).
If you have connected Cisco and Huawei switches in the same Layer-2 network and you would like to have them in the same MSTP region then the following three parts needs to be the same:
- Region name
- Revision level
- VLAN to instance mapping
Region name and revision level are part of the MSTP BPDU frames in the MST Extension. As the VLAN to instance mapping table is too big to be included the same way in the BPDU frames they need to exchanged differently. The switch is calculating a 128-bit MD5 hash digest of the VLAN to instance mapping table which will be added to the BPDU frames. You can easily see that when capturing a MSTP BPDU frame using Wirehark:
To configure Cisco and Huawei switches in the same MSTP region the configuration is straight forward.
| 1. | Cisco MSTP configuration: Switch> enableSwitch# configure terminalSwitch(config)# spanning-tree mst configurationSwitch(config-mst)# name NETCAMPSwitch(config-mst)# revision 10Switch(config-mst)# instance 1 vlan 1-1000Switch(config-mst)# instance 2 vlan 1001-2000Switch(config-mst)# instance 3 vlan 2001-3000Switch(config-mst)# instance 4 vlan 3001-4094Switch(config-mst)# exit Note: MSTP configuration is immediately active | 
| 2. | Huawei MSTP configuration: <HUAWEI> system-view[HUAWEI] stp region-configuration[HUAWEI-mst-region] region-name NETCAMP[HUAWEI-mst-region] revision-level 10[HUAWEI-mst-region] instance 1 vlan 1 to 1000[HUAWEI-mst-region] instance 2 vlan 1001 to 2000[HUAWEI-mst-region] instance 3 vlan 2001 to 3000[HUAWEI-mst-region] instance 4 vlan 3001 to 4094[HUAWEI-mst-region] active region-configuration Note: MSTP configuration must be activated with the last command | 
| 3. | Check if the MSTP region digest is the same on Cisco and Huawei switches. Huawei: display stp region-configuration digest Cisco: show spanning-tree mst configuration digest If both MD5 digest are the same you can be sure that the VLAN to instance mapping is the same for both switches. If the region name and revision level is also the same you are sure that they are part of the same MSTP region. In case that the digest are not the same and you have checked that the VLAN to instance mapping is configured correctly on both switches then it's possible that the Cisco switch has an old IOS release which is using the prestandard digest calculation. In this case you can configure the digest snooping on the Huawei uplink connecting to the Cisco switch: <HUAWEI> system-view [HUAWEI] interface 10ge 1/0/10 [HUAWEI-10GE1/0/10] stp config-digest-snoop [HUAWEI-10GE1/0/10] | 
➡️ Huawei switches are using their MAC address as the default region name (Cisco has no default region name). If you don't change the default settings on Huawei switches then they are all belonging to different MSTP regions.
Leave a Comment
