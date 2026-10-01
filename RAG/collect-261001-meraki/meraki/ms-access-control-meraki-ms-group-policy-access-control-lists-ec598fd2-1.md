---
id: collect-261001-meraki/meraki/ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2-1
title: "ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2.md
source_anchor: ""
source_lines: [1, 40]
sha256: ac3f042a1b570d468ac2f6440ee98f9b5168dfaf0e2fb19c12e0d770dc7e16ad
---

# ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2

Meraki MS Group Policy Access Control Lists
Overview
Group policies on MS switches allow users to define sets of Access Control Entries that can be applied to devices in order to control what they can access on the network. MS Group Policy ACLs can be applied to clients directly connected to an MS switch on access switchports . This article provides the details of the MS switch platforms that support Group Policy ACLs and explains how Group Policy ACLs can be created and how they are applied to clients.
Requirements, guidelines and limitations
- Hardware and software requirements: Group Policy ACLs for MS Switches are supported on the following platforms and firmware versions
    MS Switch Family MS Switch Model Minimum Firmware Required MS100 Series MS130 MS18 MS150 MS18 MS200 series MS210 MS 14.5 MS225 MS 14.5 MS250 MS 14.5 MS300 series MS350 MS 14.5 MS355 MS 14.5 MS390 MS 15.8 MS400 series MS410 MS 14.5 MS425 MS 14.5 MS450 MS 14.5 Catalyst 9K-M C9300/L/X -M CS16+ Catalyst 9K-M Cloud Native IOS XE IOS XE 17.15.2+
- 
    The recommended limits for Group Policy ACLs on the MS150/210/225/250/350/355 are : Configuration Upper limit Active groups per switch 20 Total number of active rules with layer-4 port ranges per switch 32 The per-switch limit of 32 rules with layer-4 port ranges is shared between QoS and Group Policy ACL rules (on MS150/210/225/250/350/355/410/425/450). However, while every QoS rule with a port range counts towards the limit, a Group Policy ACL rule with port range is counted only if a client device in that group is connected to the switch.
 NOTE: The MS130/MS150/210/225/250/350/355 should not be configured with more than 200 concurrent ACL entries actively assigned to 802.1X / MAB sessions on each switch. Overrunning the TCAM will result in silent failures that will cause connectivity issues to any clients attached to an active Group Policy on the switch.
 NOTE: the MS390 and C9300-M are capable of supporting up to 4900 Access Control Entries (ACL entries) active (assigned to MAB/802.1X sessions) when Group Policy ACL is used alone. When combined with Adaptive Policy the limits will decrease based on TCAM availability . Each ACL for the MS390/C9300-M is limited to a maximum of 1000 entries. This is NOT a suggested number to use, but is an upper bound of what is permitted to configure. The recommended maximum number of Group Policy ACLs defined and intended on being active concurrently should not exceed 50.
 NOTE: the C9200L-M is capable of supporting up to 1000 Access Control Entries (ACL entries) active (assigned to MAB/802.1X sessions) when Group Policy ACL is used alone. When combined with Adaptive Policy the limits will decrease based on TCAM availability. The recommended maximum number of Group Policy ACLs defined and intended on being active concurrently should not exceed 50.
 
To validate TCAM resources on Cloud Native IOS XE network devices, the live tool for "TCAM Utilization" will provide the current usage.
- 
    Group Policy ACLs enable the application of the Layer 3 Firewall rules in a group policy on the MS switches within the network. The other configuration sections of the group policy will not apply to the MS switches, but will continue to be pushed to the devices in the network, such as the MX appliance and MR access-points, to which they are relevant.
- 
    Only IP or CIDR based rules are supported. Groups containing rules using FQDNs will not be supported by MS switches.
- 
    RADIUS authentication: Group Policy ACLs on MS are applied through client authentication Access Policies and, therefore, require a RADIUS server. Static assignment of a group to a client for Group Policy ACL application is not possible on MS switches.
- 
    Access-Policy host-modes supported by Group Policy ACLs include single-host, multi-auth and multi-domain; Application of Group Policy ACL to a client authenticated by an access-policy using multi-host mode is not supported.
- 
    Group Policy ACLs on MS cannot be applied to clients connecting on trunk ports.
- 
    Group Policy ACLs on MS switches are implemented as stateless access control entires.
- 
    Group Policy ACL rules will take precedence over Switch ACL rules (configured from the Switch > ACL section) on and only on the switch where the client has been authenticated.
- 
    Group Policy ACLs on MS switches must begin with an alphanumeric character and can only be followed by alphanumeric, underscores, or hyphens characters
How it works
Group Policy ACL on MS switches are designed to work with RADIUS authentication, to allow access control lists to be dynamically applied to client traffic based on the role the RADIUS server associates with the client. The illustration below summarises the functional process.
Here is a more detailed look into the Group Policy ACL implementation shown in the illustration above.
- We start by classifying the client devices in the network into Groups and defining the Access Control List rules to allow or deny traffic from these groups to specific destinations. We then configure these groups and the rules associated with them on the Meraki dashboard from the Network Wide > Group Policies page (see Create a User Group and ACL rules for details on the configuration process).
 
 Group Policy ACLs use the Filter-Id attribute in RADIUS authentication to identify the group a client device belongs to and so, next, we enable our access-policies to use Filter-Id as the criterion for classifying devices into groups. (see Configure Access Policy to use Filter-Id for details).
- The configuration changes made in step 1 are uploaded to each MS switch in that network on the next configuration sync. This ensure that all the MS switches in your network have the same understanding of who connects to our network and how their traffic should be regulated, thus, allowing these devices to move between RADIUS-authenticated switch-ports without the need for changes to any ACL related configuration on the Dashboard.
- Each switch maintains a copy of all the groups in the network along with their ACL rules. However, a group is considered active on a switch only if at least one authenticated client device in that group exists on that switch. Conversely, when the last client device belonging to a group active on a switch disconnects or de-authenticates, the group is marked inactive on that switch. The total number of groups active on a switch at any point in times should be within the limits stated in the Guidelines and Limitations section.
- As mentioned earlier, it is the RADIUS server which determines which group a client device belongs and it communicates this information to a switch using the Filter-Id attribute. Therefore, we also configure the RADIUS server to send back a case-sensitive Filter-Id value identical to the name of the group the device being authenticated belongs to.
- When a new device is detected on a switch-port configured with an access-policy, the switch initiates communication with this device to collect credential required to verify its identity with the RADIUS servers.
 
