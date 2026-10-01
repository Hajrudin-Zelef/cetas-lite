---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e-2
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e.md
source_anchor: ""
source_lines: [128, 235]
sha256: 606f88c5b270217819130e4de6fbe72138ac904bfee3cdfd2d0d6e3a7d2910dd
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-security-cisco-nexus-9000-9f50e45e

                                          Stream Control Transmission Protocol (SCTP)
  -  
                                          				
                                          SCTP, TCP, and UDP ports
  -  
                                          				
                                          ICMP types and codes
  -  
                                          				
                                          DSCP value
  -  
                                          				
                                          TCP packets with the ACK, FIN, PSH, RST, SYN, or URG bit set
  -  
                                          				
                                          Established TCP connections
  -  
                                          				
                                          Packet length
-  
                                          				
                                          
-  
                                    		  
                                    MAC ACLs support the following additional filtering options: 
  -  
                                          				
                                          Layer 3 protocol (Ethertype)
  -  
                                          				
                                          VLAN ID
  -  
                                          				
                                          Class of Service (CoS)
-  
                                          				
                                          
- 
                                    				
                                    Beginning Cisco NX-OS Release 9.2(4), IPv4 ACLs and IPv6 in Cisco Nexus 9500 platform switches with N9K-X96136YC-R, N9K-X9636C-R, and N9K-X9636C-RX line cards and N9K-C9504-FM-R fabric module support the following additional filtering options: 
  - 
                                          						
                                          TCP packets with the ACK, FIN, PSH, RST, SYN, or URG bit set
  - 
                                          						
                                          Established TCP connections
- 
                                          						
                                          
| Note |  | 
Sequence Numbers
The device supports sequence numbers for rules. Every rule that you enter receives a sequence number, either assigned by you or assigned automatically by the device. Sequence numbers simplify the following ACL tasks:
- Adding new rules between existing rules
- 
                                    			 
                                    By specifying the sequence number, you specify where in the ACL a new rule should be positioned. For example, if you need to insert a rule between rules numbered 100 and 110, you could assign a sequence number of 105 to the new rule.
- Removing a rule
- 
                                    			 
                                    Without using a sequence number, removing a rule requires that you enter the whole rule, as follows: switch(config-acl)# no permit tcp 10.0.0.0/8 any
However, if the same rule had a sequence number of 101, removing the rule requires only the following command: switch(config-acl)# no 101
- Moving a rule
- 
                                    			 
                                    With sequence numbers, if you need to move a rule to a different position within an ACL, you can add a second instance of the rule using the sequence number that positions it correctly, and then you can remove the original instance of the rule. This action allows you to move the rule without disrupting traffic.
If you enter a rule without a sequence number, the device adds the rule to the end of the ACL and assigns a sequence number that is 10 greater than the sequence number of the preceding rule to the rule. For example, if the last rule in an ACL has a sequence number of 225 and you add a rule without a sequence number, the device assigns the sequence number 235 to the new rule.
In addition, Cisco NX-OS allows you to reassign sequence numbers to rules in an ACL. Resequencing is useful when an ACL has rules numbered contiguously, such as 100 and 101, and you need to insert one or more rules between those rules.
Logical Operators and Logical Operation Units
IP ACL rules for TCP and UDP traffic can use logical operators to filter traffic based on port numbers. Cisco NX-OS supports logical operators in only the ingress direction.
The device stores operator-operand couples in registers called logical operator units (LOUs). The LOU usage for each type of operator is as follows:
- eq
- Is never stored in an LOU
- gt
- Uses 1 LOU
- lt
- Uses 1 LOU
- neq
- Uses 1 LOU
- range
- Uses 1 LOU
| Note | For range operators, LOU threshold configuration is used to control how the port range is expanded when configuring an ACL entry. If you want to use the LOU operator when the number of the ACL rules exceed the configured threshold value, run the following command: hardware access-list lou resource threshold <x>, wherein <x> denotes the number of ACL rules to be used before the LOU threshold is reached. The range value for <x> is 1 to 50, and the default value for LOU threshold is 5. | 
ACL Logging
The ACL logging feature monitors ACL flows and logs statistics.
A flow is defined by the source interface, protocol, source IP address, source port, destination IP address, and destination port values. The statistics maintained for a flow include the number of forwarded packets (for each flow that matches the permit conditions of the ACL entry) and dropped packets (for each flow that matches the deny conditions of the ACL entry).
Time Ranges
You can use time ranges to control when an ACL rule is in effect. For example, if the device determines that a particular ACL applies to traffic arriving on an interface, and a rule in the ACL uses a time range that is not in effect, the device does not compare the traffic to that rule. The device evaluates time ranges based on its clock.
When you apply an ACL that uses time ranges, the device updates the affected I/O module whenever a time range referenced in the ACL starts or ends. Updates that are initiated by time ranges occur on a best-effort priority. If the device is especially busy when a time range causes an update, the device may delay the update by up to a few seconds.
IPv4, IPv6, and MAC ACLs support time ranges. When the device applies an ACL to traffic, the rules in effect are as follows:
-  
                                 		  
                                 All rules without a time range specified
-  
                                 		  
                                 Rules with a time range that includes the second when the device applies the ACL to traffic
The device supports named, reusable time ranges, which allows you to configure a time range once and specify it by name when you configure many ACL rules. Time range names have a maximum length of 64 alphanumeric characters.
A time range contains one or more rules. The two types of rules are as follows:
- Absolute
-  
                                 			 
                                 A rule with a specific start date and time, specific end date and time, both, or neither. The following items describe how the presence or absence of a start or end date and time affect whether an absolute time range rule is active: 
  -  
                                       				  
                                       Start and end date and time both specified—The time range rule is active when the current time is later than the start date and time and earlier than the end date and time.
  -  
                                       				  
                                       Start date and time specified with no end date and time—The time range rule is active when the current time is later than the start date and time.
  -  
                                       				  
