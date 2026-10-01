---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-2f711dbf
title: "c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-2f711dbf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-2f711dbf.md
source_anchor: ""
source_lines: [1, 51]
sha256: 30ea3665b8aef1aec27593cdc66e1be2c1dfef6dc5d5023d2d7501b155b8a276
---

# c-en-us-td-docs-switches-lan-catalyst9200-software-release-17-8-configuration-gu-2f711dbf

Information About Enhanced Object Tracking
The following sections provide information about enhanced object tracking.
Enhanced Object Tracking Overview
Before the introduction of the Enhanced Object Tracking feature, Hot Standby Router Protocol (HSRP) had a simple tracking mechanism that allowed you to track the interface line-protocol state only. If the line-protocol state of the interface went down, the HSRP priority of the router was reduced, allowing another HSRP router with a higher priority to become active.
The Enhanced Object Tracking feature separates the tracking mechanism from HSRP and creates a separate standalone tracking process that can be used by processes other than HSRP. This feature allows the tracking of other objects in addition to the interface line-protocol state.
A client process such as HSRP, Virtual Router Redundancy Protocol (VRRP), or Gateway Load Balancing Protocol (GLBP), can register its interest in tracking objects and then be notified when the tracked object changes state.
Each tracked object has a unique number that is specified in the tracking command-line interface (CLI). Client processes use this number to track a specific object. The tracking process periodically polls the tracked object for value changes and sends any changes (as up or down values) to interested client processes, either immediately or after a specified delay. Several clients can track the same object, and can take different actions when the object changes state.
You can also track a combination of objects in a list by using either a weight threshold or a percentage threshold to measure the state of the list. You can combine objects using Boolean logic. A tracked list with a Boolean “AND” function requires that each object in the list be in an up state for the tracked object to be up. A tracked list with a Boolean “OR” function needs only one object in the list to be in the up state for the tracked object to be up.
Tracking Interface Line-Protocol or IP Routing State
You can track either the interface line protocol state or the interface IP routing state. When you track the IP routing state, these three conditions are required for the object to be up:
- 
                                 				
                                 IP routing must be enabled and active on the interface.
- 
                                 				
                                 The interface line-protocol state must be up.
- 
                                 				
                                 The interface IP address must be known.
If all three of these conditions are not met, the IP routing state is down.
Tracked Lists
You can configure a tracked list of objects with a Boolean expression, a weight threshold, or a percentage threshold. A tracked list contains one or more objects. An object must exist before it can be added to the tracked list.
- 
                                 				
                                 You configure a Boolean expression to specify calculation by using either “AND” or “OR” operators.
- 
                                 				
                                 When you measure the tracked list state by a weight threshold, you assign a weight number to each object in the tracked list. The state of the tracked list is determined by whether or not the threshold was met. The state of each object is determined by comparing the total weight of all objects against a threshold weight for each object.
- 
                                 				
                                 When you measure the tracked list by a percentage threshold, you assign a percentage threshold to all objects in the tracked list. The state of each object is determined by comparing the assigned percentages of each object to the list.
Tracking Other Characteristics
You can also use the enhanced object tracking for tracking other characteristics.
- 
                                 				
                                 You can track the reachability of an IP route by using the track ip route reachability global configuration command.
- 
                                 				
                                 You can use the track ip route metric threshold global configuration command to determine if a route is above or below threshold.
- 
                                 				
                                 You can use the track resolution global configuration command to change the metric resolution default values for routing protocols.
- 
                                 				
                                 You can use the track timer tracking configuration command to configure the tracking process to periodically poll tracked objects.
Use the show track privileged EXEC command to verify enhanced object tracking configuration.
IP SLAs Object Tracking
Cisco IOS IP Service Level Agreements (IP SLAs) is a network performance measurement and diagnostics tool that uses active monitoring by generating traffic to measure network performance. Cisco IP SLAs operations collects real-time metrics that you can use for network troubleshooting, design, and analysis.
Object tracking of IP SLAs operations allows clients to track the output from IP SLAs objects and use this information to trigger an action. Every IP SLAs operation maintains an SNMP operation return-code value, such as OK or OverThreshold, that can be interpreted by the tracking process. You can track two aspects of IP SLAs operation: state and reachability. For state, if the return code is OK, the track state is up; if the return code is not OK, the track state is down. For reachability, if the return code is OK or OverThreshold, reachability is up; if not OK, reachability is down.
Static Route Object Tracking
Static routing support using enhanced object tracking provides the ability for the device to use ICMP pings to identify when a pre-configured static route or a DHCP route goes down. When tracking is enabled, the system tracks the state of the route and informs the client when that state changes. Static route object tracking uses Cisco IP SLAs to generate ICMP pings to monitor the state of the connection to the primary gateway.
