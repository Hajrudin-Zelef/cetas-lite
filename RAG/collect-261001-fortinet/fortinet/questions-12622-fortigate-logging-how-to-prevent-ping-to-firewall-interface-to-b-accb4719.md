---
id: collect-261001-fortinet/fortinet/questions-12622-fortigate-logging-how-to-prevent-ping-to-firewall-interface-to-b-accb4719
title: "questions-12622-fortigate-logging-how-to-prevent-ping-to-firewall-interface-to-b-accb4719"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-12622-fortigate-logging-how-to-prevent-ping-to-firewall-interface-to-b-accb4719.md
source_anchor: ""
source_lines: [1, 23]
sha256: b3e9f9a6d8496193acf3e9069e98cd6c4951609ed9c7564ad3906151eb1eff86
---

# questions-12622-fortigate-logging-how-to-prevent-ping-to-firewall-interface-to-b-accb4719

Fortigate logging: how to prevent ping to firewall interface to be logged while still getting logs for implicitly denied traffic - Network Engineering Stack Exchange
Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Is there a way to configure a fortinet firewall (e.g, fortigate600 running FortiOS 5 or FortiOS 4) so it do not generate log entries for pings that are directed to the firewall's own interfaces but still do generate log entries for implicitly denied traffic?
In both cases, the log entries specifies the policy with id '0' as the policy generating the log message.
In the case of successful pings, 'status' is set to 'accept' in the log and the VDOM name is set as the 'dstintf'.
I have tried to create firewall rules that match the ping traffic directed to local firewall interfaces, with the intent to explicitly disable logging, but I failed to come up with a rule that manage to match the traffic. Also, there is the option to disable the logging for implicit rule 0 (the implicit 'deny' rule at the bottom of the policy) but that also disables the logging of denied traffic, which is not what I want.
Pinging firewall interfaces (for determining that the firewall interface is available) is relied upon in certain situations and can not always be designed away. (E.g is some setups using load balancers). Also, being able to configure network equipment to avoid unwanted logging messages from being generated is always desirable, to keep down the amount of "noise" that is sent to external logging servers (Splunk etc), and in our case, logs about those 'heartbeat pings' is just considered noise.
For Fortigate firewalls running FortiOS 5.0 or newer, it is possible to use the CLI to specifically disable logs for accepted traffic directed to the firewall itself:
Log on to firewall using SSH, then run the following commands (assuming the firewall has a VDOM named 'root')
config vdom
edit root
config log settings
set local-in-allow disable
This has to be done on a per VDOM basis.
Once this is done, the firewall keeps logging all denied traffic, without logging accepted pings, SNMP monitoring queries etc.
For Fortigate firewalls running FortiOS older than 5.0, I suppose the best advice is to upgrade to 5.0 or newer and then apply the setting suggested above. It seems like the feature 'set local-in-allow disable' is not available before FortiOS 5.0.
A policy allowing ping only from specific addresses following by a policy that denies ping from any source. havent tested it, but if it falls on a policy, it shouldnt get to the implicit rule.
Another option is to restrict admin login from specific host. you can restrict the admin login to all addresses that you need ping from and addresses that needs access to the fortigate. if anyone else tries to ping, it will be blocked before it gets to the policies.
