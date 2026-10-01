---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade.md
source_anchor: ""
source_lines: [1, 58]
sha256: b3b1ea39afe4da39044612f6a37936353a00129ebc4e8127a2dfe20823c224df
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade

Overview of Port-Based Traffic Control
Port-based traffic control is a set of Layer 2 features on the Cisco Catalyst switches used to filter or block packets at the port level in response to specific traffic conditions. The following port-based traffic control features are supported:
- 
                              				
                              Storm Control
- 
                              		  
                              Protected Ports
- 
                              		  
                              Port Blocking
Finding Feature Information
Your software release may not support all the features documented in this module. For the latest caveats and feature information, see Bug Search Tool and the release notes for your platform and software release. To find information about the features documented in this module, and to see a list of the releases in which each feature is supported, see the feature information table at the end of this module.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature Navigator, go to http://www.cisco.com/go/cfn. An account on Cisco.com is not required.
Information About Storm Control
Storm Control
Storm control prevents traffic on a LAN from being disrupted by a broadcast, multicast, or unicast storm on one of the physical interfaces. A LAN storm occurs when packets flood the LAN, creating excessive traffic and degrading network performance. Errors in the protocol-stack implementation, mistakes in network configurations, or users issuing a denial-of-service attack can cause a storm.
Storm control (or traffic suppression) monitors packets passing from an interface to the switching bus and determines if the packet is unicast, multicast, or broadcast. The switch counts the number of packets of a specified type received within the 1-second time interval and compares the measurement with a predefined suppression-level threshold.
How Traffic Activity is Measured
Storm control uses one of these methods to measure traffic activity:
- 
                                    Bandwidth as a percentage of the total available bandwidth of the port that can be used by the broadcast, multicast, or unicast traffic
- 
                                    Traffic rate in packets per second at which broadcast, multicast, or unicast packets are received
- 
                                    Traffic rate in bits per second at which broadcast, multicast, or unicast packets are received
With each method, the port blocks traffic when the rising threshold is reached. The port remains blocked until the traffic rate drops below the falling threshold (if one is specified) and then resumes normal forwarding. If the falling suppression level is not specified, the switch blocks all traffic until the traffic rate drops below the rising suppression level. In general, the higher the level, the less effective the protection against broadcast storms.
| Note | When the storm control threshold for multicast traffic is reached, all multicast traffic except control traffic, such as bridge protocol data unit (BDPU) and Cisco Discovery Protocol frames, are blocked. However, the switch does not differentiate between routing updates, such as OSPF, and regular multicast data traffic, so both types of traffic are blocked. | 
Traffic Patterns
Broadcast traffic being forwarded exceeded the configured threshold between time intervals T1 and T2 and between T4 and T5. When the amount of specified traffic exceeds the threshold, all traffic of that kind is dropped for the next time period. Therefore, broadcast traffic is blocked during the intervals following T2 and T5. At the next time interval (for example, T3), if broadcast traffic does not exceed the threshold, it is again forwarded.
The combination of the storm-control suppression level and the 1-second time interval controls the way the storm control algorithm works. A higher threshold allows more packets to pass through. A threshold value of 100 percent means that no limit is placed on the traffic. A value of 0.0 means that all broadcast, multicast, or unicast traffic on that port is blocked.
| Note | Because packets do not arrive at uniform intervals, the 1-second time interval during which traffic activity is measured can affect the behavior of storm control. | 
You use the storm-control interface configuration commands to set the threshold value for each traffic type.
How to Configure Storm Control
Configuring Storm Control and Threshold Levels
You configure storm control on a port and enter the threshold level that you want to be used for a particular type of traffic.
However, because of hardware limitations and the way in which packets of different sizes are counted, threshold percentages are approximations. Depending on the sizes of the packets making up the incoming traffic, the actual enforced threshold might differ from the configured level by several percentage points.
| Note | Storm control is supported on physical interfaces. You can also configure storm control on an EtherChannel. When storm control is configured on an EtherChannel, the storm control settings propagate to the EtherChannel physical interfaces. | 
Follow these steps to storm control and threshold levels:
Before you begin
Storm control is supported on physical interfaces. You can also configure storm control on an EtherChannel. When storm control is configured on an EtherChannel, the storm control settings propagate to the EtherChannel physical interfaces.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable   | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config)# interface gigabitethernet1/0/1  | Specifies the interface to be configured, and enter interface configuration mode. | 
| Step 4 | storm-control {broadcast \| multicast \| unicast} level {level [level-low] \| bps bps [bps-low] \| pps pps [pps-low]} Example:  Device(config-if)# storm-control unicast level 87 65  | Configures broadcast, multicast, or unicast storm control. By default, storm control is disabled. The keywords have these meanings:  For BPS and PPS settings, you can use metric suffixes such as k, m, and g for large number thresholds. | 
| Step 5 | storm-control action {shutdown \| trap} Example:  Device(config-if)# storm-control action trap  | Specifies the action to be taken when a storm is detected. The default is to filter out the traffic and not to send traps.  | 
| Step 6 | end Example:  Device(config-if)# end  | Returns to privileged EXEC mode. | 
| Step 7 | show storm-control [interface-id] [broadcast \| multicast \| unicast] Example:  Device# show storm-control gigabitethernet1/0/1 unicast  | Verifies the storm control suppression levels set on the interface for the specified traffic type. If you do not enter a traffic type, details for all traffic types (broadcast, multicast and unicast) are displayed. | 
| Step 8 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Information About Protected Ports
Protected Ports
Some applications require that no traffic be forwarded at Layer 2 between ports on the same switch so that one neighbor does not see the traffic generated by another neighbor. In such an environment, the use of protected ports ensures that there is no exchange of unicast, broadcast, or multicast traffic between these ports on the switch.
Protected ports have these features:
-  
                                    		  
