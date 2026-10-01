---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107-6
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107.md
source_anchor: ""
source_lines: [341, 416]
sha256: eb139420be5707237d008204c18f0cab890dd87fdb229c2a880be09e1d39a3d0
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107

                                          Dynamic Host Configuration Protocol (DHCP) snooping
- 
                                          						
                                          Network Access Control (NAC)
- 
                                          						
                                          Dynamic ARP Inspection (DAI)
- 
                                          						
                                          IP source guard (IPSG)
- 
                                          						
                                          Internet Group Management Protocol (IGMP) snooping
- 
                                          						
                                          Hot Standby Routing Protocol (HSRP)
- 
                                          						
                                          Protocol Independent Multicast (PIM)
- 
                                          						
                                          All routing protocol configurations
After configuring the vPC, display the configurations for each vPC peer device to ensure all parameters are compatible.
Consequences of Parameter Mismatches
You can configure the graceful consistency check feature, which suspends only the links on the secondary peer device when a mismatch is introduced in a working vPC. This feature is configurable only in the CLI and is enabled by default.
Consistency Check Behavior
The graceful consistency-check command is configured by default.
As part of the consistency check of all parameters from the list of parameters that must be identical, the system checks the consistency of all VLANs.
The vPC remains operational, and only the inconsistent VLANs are brought down. This per-VLAN consistency check feature cannot be disabled and does not apply to Multiple Spanning Tree (MST) VLANs.
Deleting the vPC port-channel on the switch results in the suspension of the allowed VLANs on the corresponding vPC port-channel on the peer switch, regardless of the vPC role.
vPC Numbers
Once you have created the vPC domain ID and the vPC Peer-Link, you create port channels to attach the downstream device to each vPC peer device. That is, you create one port channel to the downstream device from the primary vPC peer device and you create another port channel to the downstream device from the secondary peer device.
| Note | We recommend that you configure the ports on the downstream devices that connect to a host or a network device that is not functioning as a switch or a bridge as STP edge ports. | 
On each vPC peer device, you assign a vPC number to the port channel that connects to the downstream device. You will experience minimal traffic disruption when you are creating vPCs. To simplify the configuration, you can assign the vPC ID number to every port channel to be the same as the port channel itself (that is, vPC ID 10 for port channel 10).
| Note | The vPC number that you assign to the port channel that connects to the downstream device from the vPC peer device must be identical on both vPC peer devices. | 
Moving Other Port Channels into vPCs
Moving Other Port Channels into vPCs — concept overview.
| Note | You must attach a downstream device using a port channel to both vPC peer devices. | 
To connect to the downstream device, you create a port channel to the downstream device from the primary vPC peer device and you create another port channel to the downstream device from the secondary peer device. On each vPC peer device, you assign a vPC number to the port channel that connects to the downstream device. You will experience minimal traffic disruption when you are creating vPCs.
vPC Object Trackings
vPC Object Trackings — concept overview.
vPC Object Tracking
| Note | We recommend that you configure the vPC Peer-Links on dedicated ports of different modules on Cisco Nexus 9500 devices. This is recommended to reduce the possibility of a failure. For the best resiliency scenario, use at least two modules. | 
vPC object tracking is used to prevent traffic black-holing in case of failure of a module where both vPC Peer-Link and uplinks to the core resides. By tracking interface feature can suspend vPC on affected switch and prevent traffic black-holing.
If you must configure all the vPC Peer-Links and core-facing interfaces on a single module, you should configure, using the command-line interface, a track object and a track list that is associated with the Layer 3 link to the core and on all vPC Peer-Links on both vPC peer devices. You use this configuration to avoid dropping traffic if that particular module goes down because when all the tracked objects on the track list go down, the system does the following:
- 
                                    					
                                    Stops the vPC primary peer device sending peer-keepalive messages, which forces the vPC secondary peer device to take over.
- 
                                    					
                                    Brings down all the downstream vPCs on that vPC peer device, which forces all the traffic to be rerouted in the access switch toward the other vPC peer device.
Once you configure this feature and if the module fails, the system automatically suspends all the vPC links on the primary vPC peer device and stops the peer-keepalive messages. This action forces the vPC secondary device to take over the primary role and all the vPC traffic to go to this new vPC primary device until the system stabilizes.
You should create a track list that contains all the links to the core and all the vPC Peer-Links as its object. Enable tracking for the specified vPC domain for this track list. Apply this same configuration to the other vPC peer device. See the Cisco Nexus 9000 Series NX-OS Unicast Routing Configuration Guide for information about configuring object tracking and track lists.
| Note | This example uses Boolean OR in the track list and forces all traffic to the vPC peer device only for a complete module failure. If you want to trigger a switchover when any core interface or vPC Peer-Link goes down, use a Boolean AND in the torack list below. | 
To configure a track list to switch over a vPC to the remote peer when all related interfaces on a single module fail, follow these steps:
- 
                                    					
                                    Configure track objects on an interface (Layer 3 to core) and on a port channel (vPC Peer-Link). 
switch(config-if)# track 35 interface ethernet 8/35 line-protocol
switch(config-track)# track 23 interface ethernet 8/33 line-protocol
switch(config)# track 55 interface port-channel 100 line-protocol
- 
                                    					
                                    Create a track list that contains all the interfaces in the track list using the Boolean OR to trigger when all objects fail. 
switch(config)# track 44 list boolean OR
switch(config-track)# object 23
switch(config-track)# object 35
switch(config-track)# object 55
switch(config-track)# end
- 
                                    					
                                    Add this track object to the vPC domain: 
switch(config)# vpc domain 1
switch(config-vpc-domain)# track 44
- 
                                    					
