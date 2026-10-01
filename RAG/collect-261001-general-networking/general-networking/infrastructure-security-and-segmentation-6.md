---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-6
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2017-12"]
keywords: []
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [379, 472]
sha256: 6f87ae43594e5a3c0a81609152c95ed856324f5350316cd46874acd317dd140e
---

# infrastructure-security-and-segmentation

The optional **log** keyword at the end of the command causes the router to generate a syslog every time a packet matches that ACL. The log message includes the ACL number, the action taken on the packet, the source IP address of the packet, and the number of matches from a source within a five-minute period. The **log** keyword should be used sparingly to reduce CPU load. Typically, this keyword is used on ACEs that deny certain traffic to monitor for suspicious activity.

The **access-list** command can be repeated with the same *access-list-number* as often as required to add more ACEs in the ACL. Each subsequent command adds the new ACE at the end of the existing list. The order of the ACEs in an ACL is important because the ACL will be evaluated from the top, and the first match will be applied to a given packet. So, if there is a **deny** statement before a **permit** statement for the same source, the packet will be dropped.

The access list can be applied to an interface with the **ip access-group** *access-list-number* {**in**|**out**} command. The **in** and **out** keywords indicate the direction in which the access list is applied.

Example 2-52 shows how an access list is created and applied to deny traffic from the 10.1.1.0/24 subnet and from host 10.2.1.1, while all other traffic is permitted coming into interface Gi1. In the example, note the last ACE that explicitly permits all traffic. If this ACE is not added, all traffic will be dropped because of the implicit deny.

#### **Example 2-52** *Creating and Applying a Standard ACL*

`Router(config)#**access-list 10 deny 10.1.1.0 0.0.0.255**
Router(config)#**access-list 10 deny host 10.2.1.1**
Router(config)#**access-list 10 permit any**
Router(config)#**interface Gi1**
Router(config-if)#**ip access-group 10 in**`

#### Extended ACLs

While standard ACLs are simple, they have limited functionality because they can only filter based on the source address. Most security policies require more granular filtering based on various fields, such as source address, destination address, protocols, and ports. Extended ACLs can be used for such filtering requirements.

Extended ACLs are created and applied in the same way as standard ACLs and follow similar rules of sequential ACEs and implicit deny. The difference begins with the ACL numbers: Extended ACLs are numbered from 100 to 199 and 2000 to 2699. The syntax of extended ACLs is also different, as shown here:

`**access-list** *access-list-number* {**deny**|**permit**} **protocol** *source destination*
[*protocol-options*] [**log|log-input**]`

As with standard ACLs, *source* and *destination* can be defined as a single host with the **host** keyword or a subnet with inverse mask or as all hosts with the **any** keyword.

The optional *protocol-options* part of the command differs based on the protocol specified in the ACE. For example, when TCP or UDP protocols are defined, **eq**, **lt**, and **gt** (equal to, less than, and greater than) keywords are available to specify ports to be matched. In addition to ports, various bits in the TCP headers, such as SYN and ACK, can also be matched with available keywords.

Example 2-53 shows access list 101 created and applied to block all Telnet and RDP traffic coming into interface Gi1. In addition, the access list blocks all communication between hosts 10.1.1.1 and 10.2.1.1. Note the use of the **eq** keyword with TCP in the first two lines to deny Telnet and RDP traffic.

#### **Example 2-53** *Creating and Applying Extended ACLs*

`Router(config)#**access-list 101 deny tcp any any eq 23**
Router(config)#**access-list 101 deny tcp any any eq 3389**
Router(config)#**access-list 101 deny ip host 10.1.1.1 host 10.2.1.1**
Router(config)#**access-list 101 permit ip any any**
Router(config)#**interface Gi1**
Router(config-if)#**ip access-group 101 in**`

#### Named ACLs

During the course of normal operations, it is not uncommon to see tens or hundreds of ACLs created on a device. Eventually, when ACLs are identified with numbers only, it becomes difficult to keep track of the reason an ACL was created. To make administering ACLs easy, you can give them names instead of numbers; in this case, they are called named ACLs. Both standard and extended ACLs can be created as named ACLs, and all the previous discussed rules apply. The difference is in the commands used to create the ACL. Named ACLs are created using the **ip access-list** command, shown here:

`**ip access-list** {**standard**|**extended**} *name*`

This command creates an ACL and brings you to the **nacl** prompt, denoted **config-std-nacl** for a standard ACL or **config-ext-nacl** for an extended ACL. At this prompt, you can add ACEs as usual, starting with a **permit** or **deny** command. The rest of the command follows the syntax discussed earlier for standard and extended ACLs.

The command to apply a named ACL is the same command that is used to apply the standard and extended ACLs except that a name is used instead of a number in the **ip access-group** command, as shown here:

`**ip access-group** *name {in*|*out}*`

To illustrate the differences and similarities between creating numbered and named ACLs, Example 2-54 re-creates the standard ACL from Example 2-52 as a named ACL, and Example 2-55 does the same for the extended ACL shown in Example 2-53.

#### **Example 2-54** *Creating and Applying Standard Named ACLs*

`Router(config)#**ip access-list standard bad-hosts**
Router(config-std-nacl)#**deny 10.1.1.0 0.0.0.255**
Router(config-std-nacl)#**deny host 10.2.1.1**
Router(config-std-nacl)#**permit any**
Router(config-std-nacl)#**exit**
Router(config)#**interface Gi1**
Router(config-if)#**ip access-group bad-hosts in**`

#### **Example 2-55** *Creating and Applying Extended Named ACLs*

`Router(config)#**ip access-list extended bad-traffic**
Router(config-ext-nacl)#**deny tcp any any eq 23**
Router(config-ext-nacl)#**deny tcp any any eq 3389**
Router(config-ext-nacl)#**deny ip host 10.1.1.1 host 10.2.1.1**
Router(config-ext-nacl)#**permit ip any any**
Router(config-ext-nacl)#**exit**
Router(config)#**interface Gi1**
Router(config-if)#**ip access-group bad-traffic in**`

#### Time Based ACLs

ACEs in an extended ACL can be configured to be enforced during certain times only. To do this, you specify a time range at the end of an ACE with the **time-range** *time-range-name* command. *time-range-name* references an object in IOS that defines a time period. The object can be created with the **time-range** *name* global configuration command. Within the object, a time period is defined as recurring or absolute. A recurring, or periodic, range starts and ends at the same time on certain days of the week, while an absolute range starts and ends at a specific date and time.

A recurring, or periodic, range can be defined with the **periodic** {*day-of-the-week*|**daily**|**weekdays**|**weekends**} *start-time* **to** *end-time* command. *start-time* and *end-time* are defined in *hh*:*mm* format.

An absolute range can be defined with the **absolute** {**start**|**end**} *hh*:*mm day month year* command. The *month* option is specified with the first three letters of a month.

In Example 2-56, two time ranges are created. The first range, called *daily*, is a periodic range that starts at 00:00 hours and ends at 02:00 hours. The second range is an absolute range that begins on December 1 at 00:00 hours and ends on December 31 at 21:59 hours.

#### **Example 2-56** *Recurring Time Range*

`Router(config)#**time-range daily**
Router(config-time-range)#**periodic daily 00:00 to 02:00**
Router(config-time-range)#**exit**
Router(config)#**time-range december**
Router(config-time-range)#**absolute start 00:00 01 December 2017 end 21:59  31 December 2017**
Router(config-time-range)#**exit**`

