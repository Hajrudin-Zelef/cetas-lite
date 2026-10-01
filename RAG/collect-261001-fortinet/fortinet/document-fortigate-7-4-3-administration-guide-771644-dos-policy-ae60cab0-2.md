---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-771644-dos-policy-ae60cab0-2
title: "DoS policy"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-771644-dos-policy-ae60cab0.md
source_anchor: ""
source_lines: [71, 150]
sha256: c9e3703f893017c788892f762f9965ebc539c125589e75b36a03f856f0f4a98e
---

# DoS policy

1. 
                                                    Go to *Policy & Objects > IPv4 DoS Policy* or*Policy & Objects > IPv6 DoS Policy* and click*Create New* .If the option is not visible, enable *DoS Policy* in*Feature Visibility* . See Feature visibility for details.
2. 
                                                    Configure the following: Name Enter a name for the policy. Incoming Interface Enter the interface that the policy applies to. Source Address Enter the source address. Destination Address Enter the destination address. This is the address that the traffic is addressed to. In this case, it must be an address that is associated with the firewall interface. For example, it could be an interface address, a secondary IP address, or the address assigned to a VIP address. Service Select the services or service groups. The ALL service can used or, to optimize the firewall resources, only the services that will be answered on an interface can be used. L3 Anomalies L4 Anomalies Configure the anomalies: 
  - 
                                                                            *Logging* : Enable/disable logging for specific anomalies or all of them. Anomalous traffic will be logged when the action is*Block* or*Monitor* .
  - 
                                                                            *Action* : Select the action to take when the threshold is reached:
    - 
                                                                                    *Disable* : Do not scan for the anomaly.
    - 
                                                                                    *Block* : Block the anomalous traffic.
    - 
                                                                                    *Monitor* : Allow the anomalous traffic, but record a log message if logging is enabled.
  - 
                                                                                    
  - 
                                                                            *Threshold* : The number of detected instances that triggers the anomaly action.
 Comments Optionally, enter a comment.
3. 
                                                                            
4. 
                                                    Enable the policy, then click *OK* .

|  | The quarantine option is only available in the CLI. See Quarantine for information. | 

###### To configure a DoS policy in the GUI:

```
config firewall DoS-policy
    edit 1
        set name "Flood"
        set interface "port1"
        set srcaddr "all"
        set dstaddr "all"
        set service "ALL"
        config anomaly
            edit "icmp_flood"
                set status enable
                set log enable
                set action block
                set quarantine attacker
                set quarantine-expiry 1d1h1m
                set quarantine-log enable
                set threshold 100
            next
        end
    next
end
```
                                            | name <string> | Enter a name for the policy. | 
| interface <string> | Enter the interface that the policy applies to. | 
| srcaddr <string> | Enter the source address. | 
| dstaddr <string> | Enter the destination address. This is the address that the traffic is addressed to. In this case, it must be an address that is associated with the firewall interface. For example, it could be an interface address, a secondary IP address, or the address assigned to a VIP address. | 
| service <string> | Enter the services or service groups. The `ALL` service can used or, to optimize the firewall resources, only the services that will be answered on an interface can be used. | 
| status {enable \| disable} | Enable/disable this anomaly. | 
| log {enable \| disable} | Enable/disable anomaly logging. When enabled, a log is generated whenever the anomaly action is triggered, regardless of which action is configured. | 
| action {pass \| block} | Set the action to take when the threshold is reached:  | 
| quarantine {none \| attacker} | Set the quarantine method (see Quarantine):  | 
| quarantine-expiry <###d##h##m> | Set the duration of the quarantine, in days, hours, and minutes (###d##h##m) (1m - 364d23h59m, default = 5m). This option is available  if `quarantine` is set`attacker` . | 
| quarantine-log {enable \| disable} | Enable/disable quarantine logging (default = disable).  This option is available  if `quarantine` is set`attacker` . | 
| threshold <integer> | The number of detected instances - packets per second or concurrent session number - that triggers the anomaly action. | 

Quarantine is used to block any further traffic from a source IP address that is considered a malicious actor or a source of traffic that is dangerous to the network. Traffic from the source IP address is blocked for the duration of the quarantine, and the source IP address is added to the banned user list.

The banned user list is kept in the kernel, and used by Antivirus, Data Loss Prevention (DLP), DoS, and Intrusion Prevention System (IPS). Any policies that use any of these features will block traffic from the attacker's IP address.

###### To view the quarantined user list:

# diagnose user banned-ip list
src-ip-addr       created                  expires                  cause            
192.168.2.205     Wed Nov 25 12:47:54 2020 Wed Nov 25 12:57:54 2020 DOS   

## Troubleshooting DoS attacks

The best way to troubleshoot DoS attacks is with Anomaly logs and IPS anomaly debug messages.


###### To test an icmp_flood attack:

