---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-68-3
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "preemption"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-68.md
source_anchor: ""
source_lines: [187, 222]
sha256: 672bb98e561875dc2e6a9c16d1cc59c0f594d753ade47e7318102da3c250e962
---

# Summary

| Property | Description | 
|---|---|
| **arp** (*disabled \| enabled \| proxy-arp \| reply-only* ; Default:**enabled** ) | ARP resolution protocol mode. | 
| **arp-timeout** (*integer;* Default:**auto** ) | How long the ARP record is kept in the ARP table after no packets are received from IP. Value `auto` equals to the value of`arp-timeout` in IP/Settings, default is 3. | 
| **authentication** (*ah \| none \| simple* ; Default:**none** ) | Authentication method to use for VRRP advertisement packets.  | 
| **comment** (*string* ; Default: ) | Short description of the interface. | 
| **connection-tracking-mode** (*active-active \| passive-active* ; Default:**passive-active** ) | Specifies the mode for connection tracking synchronization. This setting is only relevant when `sync-connection-tracking=yes` is enabled.  Using multiple VRRP groups with `passive-active` mode may lead to unsynchronized connection tracking tables, since only one master handles synchronization, and the others do not exchange tracking data. Example configuration: | 
| **connection-tracking-port** (*integer;* Default:**8275** ) | Specifies UDP port for connection tracking synchronization. This setting is only relevant when `sync-connection-tracking=yes` is enabled. | 
| **group-authority** (*none \| self \| vrrp-interface;* Default:**none** ) | Allows multiple VRRP interfaces to be grouped so they share the same VRRP state. Within a group, a single group authority interface is selected - it controls the state of the other group members and is the only interface that sends VRRP advertisements. When the group-authority VRRP interface transitions to the backup state, all group members also transition to the backup state. If a failure is detected on any group member (not only the group-authority interface), for example due to a link-down on its parent interface, all group members will transition to the failure state.  For example, VRRP instances run on LAN and WAN networks with NAT in-between. If one VRRP instance is Master and the other is Backup on the same device, the entire network malfunctions due to NAT failure. Grouping LAN and WAN VRRP interfaces ensures that both are either VRRP Master or Backup. In a VRRP group, VRRP advertisements are sent only by the group authority. That's why in a typical WAN+LAN setup, it is recommended to use the LAN network as the group authority to keep VRRP control traffic in the internal network. | 
| **interface** (*string* ; Default: ) | Interface name on which VRRP instance will be running. | 
| **interval** (*time [10ms..4m15s]* ; Default:**1s** ) | The VRRP interval defines how often the VRRP master router sends Advertisement packets to backup routers. This interval directly determines the frequency at which backups receive keepalive information confirming that the master is operational. A shorter interval increases the rate of Advertisement packets, allowing faster detection of master failure, but also increases sensitivity to packet loss, processing delays, and timer inaccuracies. Longer intervals reduce control traffic and improve stability, at the cost of slower failover detection. This Master Down interval is derived from the configured VRRP interval and the router’s priority, and is calculated to allow multiple missed Advertisements before triggering failover. Configuring VRRP intervals below 1 second may lead to unpredictable behavior and unintended master role changes. | 
| **mtu** (*read-only* ; Default: ) | Layer3 MTU size. Since RouterOS v7.7, the VRRP interface always uses slave interface MTU. | 
| **name** (*string* ; Default: ) | VRRP interface name. | 
| **on-backup** (*string* ; Default: ) | Script to execute when the node is switched to the backup state. | 
| **on-master** (*string* ; Default: ) | Script to execute when the node is switched to master state. | 
| **on-fail** (*string* ; Default: ) | Script to execute when the node fails. | 
| **password** (*string* ; Default: )*sensitive* | Password required for authentication. Can be ignored if authentication is not used. | 
| **preemption-mode** (*yes \| no* ; Default:**yes** ) | Whether the master node always has the priority. When set to 'no' the backup node will not be elected to be a master until the current master fails, even if the backup node has higher priority than the current master. This setting is ignored if the owner router becomes available. | 
| **priority** (*integer: 1..254* ; Default:**100** ) | Priority of VRRP node used in Master election algorithm. A higher number means higher priority. '255' is reserved for the router that owns VR IP and '0' is reserved for the Master router to indicate that it is releasing responsibility. | 
| **remote-address** (*IPv4;* Default: ) | Specifies the remote address of the other VRRP router for syncing connection tracking. If not set, the system autodetects the remote address via VRRP. The remote address is used only if sync-connection-tracking=yes. Explicitly setting a remote address has the following benefits:  Sync connection tracking uses UDP port 8275. | 
| **v3-protocol** (*ipv4 \| ipv6* ; Default:**ipv4** ) | A protocol that will be used by VRRPv3. Valid only if the **version** is 3. | 
| **version** (*integer [2, 3]* ; Default:**3** ) | Which VRRP version to use. | 
| **vrid** (*integer: 1..255* ; Default:**1** ) | Virtual Router identifier. Each Virtual router must have a unique id number. | 
| **sync-connection-tracking** (*string* ; Default:**no** ) | Synchronize connection tracking entries from Master to Backup device. The VRRP connection tracking synchronization requires that RouterOS connection tracking is running. | 

#### Read-only flags

| Property | Description | 
|---|---|
| **backup** | The VRRP interface is in the backup state. | 
| **disabled** | The VRRP interface is disabled by the user. | 
| **failure** | The VRRP interface is in the failure state, for example due to a link-down on its parent interface. | 
| **grp-authority** | The VRRP interface is `group-authority` . It controls the state of the other group members and is the only interface that sends VRRP advertisements. | 
| **grp-member** | The VRRP interface is group member. Its state machine follows the state of the specified `group-authority` interface. | 
| **invalid** | The VRRP interface is in the invalid state, for example due to configuration error. | 
| **master** | The VRRP interface is in the master state. |
