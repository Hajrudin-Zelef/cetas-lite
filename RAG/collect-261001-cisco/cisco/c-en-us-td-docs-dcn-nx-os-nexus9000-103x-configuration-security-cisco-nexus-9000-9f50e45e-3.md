---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e-3
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e.md
source_anchor: ""
source_lines: [236, 325]
sha256: 1cc7df1ef5ffd5118738758fab222dc7b391b201caad81225263cae3f2934f2d
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e

                                       No start date and time with end date and time specified—The time range rule is active when the current time is earlier than the end date and time.
  -  
                                       				  
                                       No start or end date and time specified—The time range rule is always active.
 For example, you could prepare your network to allow access to a new subnet by specifying a time range that allows access beginning at midnight of the day that you plan to place the subnet online. You can use that time range in ACL rules that apply to the subnet. After the start time and date have passed, the device automatically begins applying the rules that use this time range when it applies the ACLs that contain the rules.
-  
                                       				  
                                       
- Periodic
-  
                                 			 
                                 A rule that is active one or more times per week. For example, you could use a periodic time range to allow access to a lab subnet only during work hours on weekdays. The device automatically applies ACL rules that use this time range only when the range is active and when it applies the ACLs that contain the rules.
| Note | The order of rules in a time range does not affect how a device evaluates whether a time range is active. Cisco NX-OS includes sequence numbers in time ranges to make editing the time range easier. | 
Time ranges also allow you to include remarks, which you can use to insert comments into a time range. Remarks have a maximum length of 100 alphanumeric characters.
The device determines whether a time range is active as follows:
-  
                                 		  
                                 The time range contains one or more absolute rules—The time range is active if the current time is within one or more absolute rules.
-  
                                 		  
                                 The time range contains one or more periodic rules—The time range is active if the current time is within one or more periodic rules.
-  
                                 		  
                                 The time range contains both absolute and periodic rules—The time range is active if the current time is within one or more absolute rules and within one or more periodic rules.
When a time range contains both absolute and periodic rules, the periodic rules can only be active when at least one absolute rule is active.
Policy-Based ACLs
The device supports policy-based ACLs (PBACLs), which allow you to apply access control policies across object groups. An object group is a group of IP addresses or a group of TCP or UDP ports. When you create a rule, you specify the object groups rather than specifying IP addresses or ports.
Using object groups when you configure IPv4 or IPv6 ACLs can help reduce the complexity of updating ACLs when you need to add or remove addresses or ports from the source or destination of rules. For example, if three rules reference the same IP address group object, you can add an IP address to the object instead of changing all three rules.
PBACLs do not reduce the resources required by an ACL when you apply it to an interface. When you apply a PBACL or update a PBACL that is already applied, the device expands each rule that refers to object groups into one ACL entry per object within the group. If a rule specifies the source and destination both with object groups, the number of ACL entries created on the I/O module when you apply the PBACL is equal to the number of objects in the source group multiplied by the number of objects in the destination group.
The following object group types apply to port, router, policy-based routing (PBR), and VLAN ACLs:
- IPv4 Address Object Groups
-  
                                 			 
                                 Can be used with IPv4 ACL rules to specify source or destination addresses. When you use the permit or deny command to configure a rule, the addrgroup keyword allows you to specify an object group for the source or destination.
- IPv6 Address Object Groups
-  
                                 			 
                                 Can be used with IPv6 ACL rules to specify source or destination addresses. When you use the permit or deny command to configure a rule, the addrgroup keyword allows you to specify an object group for the source or destination.
- Protocol Port Object Groups
-  
                                 			 
                                 Can be used with IPv4 and IPv6 TCP and UDP rules to specify source or destination ports. When you use the permit or deny command to configure a rule, the portgroup keyword allows you to specify an object group for the source or destination.
| Note | Policy-based routing (PBR) ACLs do not support deny access control entries (ACEs) or deny commands to configure a rule. | 
Kernel Stack ACL
The Kernel Stack ACL is a common CLI infrastructure to configure ACLs for management of inband and outband components.
The Kernel Stack ACL uses NX-OS ACL CLI to secure management applications on management and front panel ports. Configuring a single ACL must be able to secure all management applications on NX-OS.
Kernel Stack ACL is the component that fixes the manual intervention of the user and automatically programs iptable entries when the ACL is applied to mgmt0 interface.
The following is an example for configuring Kernel Stack ACL:
swtich# conf t
Enter configuration commands, one per line. End with CNTL/Z.
switch(config)# ip access-list kacl1
switch(config-acl)# statistics per-entry
switch(config-acl)# 10 deny tcp any any eq 443
switch(config-acl)# 20 permit ip any any
switch(config-acl)# end
switch#
switch(config-if)# interface mgmt0
switch(config-if)#   ip access-group acl1 in
switch(config-if)#   ipv6 traffic-filter acl6 in
switch(config-if)#
switch# sh ip access-lists kacl1
IP access list kacl1
statistics per-entry
10 deny tcp any any eq 443 [match=136]
20 permit ip any any [match=44952]
switch(config)#
The following is the Kernel Stack filtering for iptables entries based on the configuration:
bash-4.4# ip netns exec management iptables -L -n -v --line-numbers
Chain INPUT (policy ACCEPT 0 packets, 0 bytes)
num pkts bytes target prot opt in out source destination
1 9 576 DROP tcp -- * * 0.0.0.0/0 0.0.0.0/0 tcp dpt:443
2 0 0 ACCEPT all -- * * 0.0.0.0/0 0.0.0.0/0
3 0 0 DROP all -- * * 0.0.0.0/0 0.0.0.0/0
Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
num pkts bytes target prot opt in out source destination
Chain OUTPUT (policy ACCEPT 0 packets, 0 bytes)
num pkts bytes target prot opt in out source destination
bash-4.4#
The following are the limitations for the Kernel Stack ACL support:
- 
                                 
                                 This feature is supported only on mgmt0 interface and not on other inband interfaces.
- 
                                 
                                 Five tuples (protocol, source-ip, destination-ip, source-port, and destination-port) of the ACL entry are programmed in the iptables. Rest of the options provided in the ACL entry are not programmed in the iptables and throws a warning syslog in such instances. For example, "WARNING: Some ACL options are not supported in kstack. Only partial rule will be installed".
- 
                                 
                                 If the device user has host bash access, then the user can manually update the iptables. This update could potentially corrupt the iptable rules for which they are programmed.
- 
                                 
