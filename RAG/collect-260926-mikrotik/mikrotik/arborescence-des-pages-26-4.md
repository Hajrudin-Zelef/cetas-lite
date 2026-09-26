---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-4
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [296, 344]
sha256: 031528674e3cf73029f2a13c682cb867097adec0040bea4f5021c96c1548024f
---

# Overview

| Property | Description | 
|---|---|
| **channel** (*list* ; Default: ) | User defined list taken from Channel names ( **/caps-man channels** ) | 
| **channel.band** (*2ghz-b \| 2ghz-b/g \| 2ghz-b/g/n \| 2ghz-onlyg \| 2ghz-onlyn \| 5ghz-a \| 5ghz-a/n \| 5ghz-onlyn \| 5ghz-a/n/ac \| 5ghz-only-ac* ; Default: ) | Defines set of used channels. | 
| **channel.control-channel-width** (*40mhz-turbo \| 20mhz \| 10mhz \| 5mhz* ; Default: ) | Defines set of used channel widths. | 
| **channel.extension-channel** (*Ce \| Ceee \| eC \| eCee \| eeCe \| eeeC \| xx \| xxxx \| disabled* ; Default: ) | Extension channel configuration. (E.g. Ce = extension channel is above Control channel, eC = extension channel is below Control channel) | 
| **channel.frequency** (*integer [0..4294967295]* ; Default: ) | Channel frequency value in MHz on which AP will operate. If left blank, CAPsMAN will automatically determine the best frequency that is least occupied. | 
| **channel.reselect-interval** (*time [00:00:00]* ;*[00:00:00..00:00:00];* Default: ) | The interval after which the least occupied frequency is chosen, can be defined as a random interval, ex. as "30m..60m". Works only if **channel.frequency** is left blank. | 
| **channel.save-selected** (*yes \| no* ; Default:**no** ) | If channel frequency is chosen automatically and **channel.reselect-interval** is used, then saves the last picked frequency. | 
| **channel.secondary-frequency** (*integer [0..4294967295]* ; Default:**auto** ) | Specifies the second frequency that will be used for 80+80MHz configuration. Set it to **Disabled** in order to disable 80+80MHz capability. | 
| **channel.skip-dfs-channels** (*yes \| no* ; Default:**no** ) | If **channel.frequency** is left blank, the selection will skip DFS channels | 
| **channel.tx-power** (*integer [-30..40]* ; Default: ) | TX Power for CAP interface (for the whole interface not for individual chains) in dBm. It is not possible to set higher than allowed by country regulations or interface. By default max allowed by country or interface is used. | 
| **channel.width** (; Default: ) | Sets Channel Width in MHz. | 
| **comment** (*string* ; Default: ) | Short description of the Configuration profile | 
| **country** (*name of the country \| no_country_set* ; Default:**no_country_set** ) | Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default value of **scan-list** . Value*no_country_set* is an FCC compliant set of channels. | 
| **datapath** (*list* ; Default: ) | User defined list taken from Datapath names ( **/caps-man datapath** ) | 
| **datapath.bridge** (*list* ; Default: ) | Bridge to which particular interface should be automatically added as port | 
| **datapath.bridge-cost** (*integer [1..*200000000*]* ; Default: ) | bridge port cost to use when adding as bridge port | 
| **datapath.bridge-horizon** (*integer [0..4294967295]* ; Default: ) | bridge horizon to use when adding as bridge port | 
| **datapath.client-to-client-forwarding** (*yes \| no* ; Default:**no** ) | controls if client-to-client forwarding between wireless clients connected to interface should be allowed, in local forwarding mode this function is performed by CAP, otherwise it is performed by CAPsMAN | 
| **datapath.interface-list** (; Default: ) |  | 
| **datapath.l2mtu** (; Default: ) | set Layer2 MTU size | 
| **datapath.local-forwarding** (*yes \| no* ; Default:**no** ) | controls forwarding mode | 
| **datapath.mtu** (; Default: ) | set MTU size | 
| **datapath.openflow-switch** (; Default: ) | OpenFlow switch port (when enabled) to add interface to | 
| **datapath.vlan-id** (*integer [1..4095]* ; Default: ) | VLAN ID to assign to interface if vlan-mode enables use of VLAN tagging | 
| **datapath.vlan-mode** (*use-service-tag \| use-tag* ; Default: ) | Enables and specifies the type of VLAN tag to be assigned to the interface (causes all received data to get tagged with VLAN tag and allows the interface to only send out data tagged with given tag) | 
| **disconnect-timeout** (; Default: ) |  | 
| **distance** (; Default: ) |  | 
| **frame-lifetime** (; Default: ) |  | 
| **guard-interval** (*any \| long* ; Default:**any** ) | Whether to allow the use of short guard interval (refer to 802.11n MCS specification to see how this may affect throughput). "any" will use either short or long, depending on data rate, "long" will use long only. | 
| **hide-ssid** (*yes \| no* ; Default: ) |   *yes* can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP. | 
| **hw-protection-mode** (; Default: ) |  | 
| **hw-retries** (; Default: ) |  | 
| **installation** (*any \| indoor \| outdoor* ; Default:**any** ) |  | 
| **keepalive-frames** (*enabled \| disabled* ; Default:**enabled** ) |  | 
| **load-balancing-group** (*string* ; Default: ) | Tags the interface to the load balancing group. For a client to connect to interface in this group, the interface should have the same number of already connected clients as all other interfaces in the group or smaller. Useful in setups where ranges of CAPs mostly overlap. | 
| **max-sta-count** (*integer [1..2007]* ; Default: ) | Maximum number of associated clients. | 
| **mode** (; Default:**ap** ) | Set operational mode. Only ap currently supported. | 
| **multicast-helper** (*default \| disabled \| full* ; Default:**default** ) | When set to full multicast packets will be sent with unicast destination MAC address, resolving  multicast problem on a wireless link. This option should be enabled only on the access point, clients should be configured in **station-bridge** mode. Available starting from v5.15. | 
| **name** (*string* ; Default: ) | Descriptive name for the Configuration Profile | 
| **rates** (; Default: ) | User defined list taken from Rates names ( **/caps-man rates** ) | 
| **rates.basic** (*1Mbps \| 2Mbps \| 5.5Mbps \| 6Mbps \| 11Mbps \| 11Mbps \| 12Mbps \| 18Mbps \| 24Mbps \| 36Mbps \| 48Mbps \| 54Mbps* ; Default: ) |  | 
| **rates.supported** (*1Mbps \| 2Mbps \| 5.5Mbps \| 6Mbps \| 11Mbps \| 11Mbps \| 12Mbps \| 18Mbps \| 24Mbps \| 36Mbps \| 48Mbps \| 54Mbps* ; Default: ) |  | 
| **rates.ht-basic-mcs** (*list of (mcs-0 \| mcs-1 \| mcs-2 \| mcs-3 \| mcs-4 \| mcs-5 \| mcs-6 \| mcs-7 \| mcs-8 \| mcs-9 \| mcs-10 \| mcs-11 \| mcs-12 \| mcs-13 \| mcs-14 \| mcs-15 \| mcs-16 \| mcs-17 \| mcs-18 \| mcs-19 \| mcs-20 \| mcs-21 \| mcs-22 \| mcs-23)* ; Default:**mcs-0; mcs-1; mcs-2; mcs-3; mcs-4; mcs-5; mcs-6; mcs-7** ) | Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification. | 
| **rates.ht-supported-mcs** (*list of (mcs-0 \| mcs-1 \| mcs-2 \| mcs-3 \| mcs-4 \| mcs-5 \| mcs-6 \| mcs-7 \| mcs-8 \| mcs-9 \| mcs-10 \| mcs-11 \| mcs-12 \| mcs-13 \| mcs-14 \| mcs-15 \| mcs-16 \| mcs-17 \| mcs-18 \| mcs-19 \| mcs-20 \| mcs-21 \| mcs-22 \| mcs-23)* ; Default:**mcs-0; mcs-1; mcs-2; mcs-3; mcs-4; mcs-5; mcs-6; mcs-7; mcs-8; mcs-9; mcs-10; mcs-11; mcs-12; mcs-13; mcs-14; mcs-15; mcs-16; mcs-17; mcs-18; mcs-19; mcs-20; mcs-21; mcs-22; mcs-23** ) | Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification. | 
| **rates.vht-basic-mcs** (*none \| MCS 0-7 \| MCS 0-8 \| MCS 0-9* ; Default:**none** ) | Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream  | 
| **rates.vht-supported-mcs** (*none \| MCS 0-7 \| MCS 0-8 \| MCS 0-9* ; Default:**none** ) | Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream  | 
| **rx-chains** (*list of integer [0..2]* ; Default:**0** ) | Which antennas to use for receive. | 
