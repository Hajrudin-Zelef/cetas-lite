---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-771644-dos-policy-ae60cab0-1
title: "DoS policy"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-771644-dos-policy-ae60cab0.md
source_anchor: ""
source_lines: [1, 70]
sha256: ae47f52e287c3c1f83f8263acaaa76d0577220b2182418f96231ff980a7d0fd5
---

# DoS policy

# DoS policy

A Denial of Service (DoS) policy examines network traffic arriving at a FortiGate interface for anomalous patterns, which usually indicates an attack.

A denial of service occurs when an attacking system starts an abnormally large number of sessions with a target system. The large number of sessions slows down or disables the target system, preventing legitimate users from using it.

DoS policies are checked before security policies, preventing attacks from triggering more resource intensive security protection and slowing down the FortiGate.

## DoS anomalies

Predefined sensors are setup for specific anomalous traffic patterns. New DoS anomalies cannot be added by the user.

The predefined anomalies that can be used in DoS policies are:

| Anomaly | Description | Recommended Threshold | 
|---|---|---|
| tcp_syn_flood | If the SYN packet rate of new TCP connections, including retransmission, to one destination IP address exceeds the configured threshold value, the action is executed. | 2000 packets per second. | 
| tcp_port_scan | If the SYN packet rate of new TCP connections, including retransmission, from one source IP address exceeds the configured threshold value, the action is executed. | 1000 packets per second. | 
| tcp_src_session | If the number of concurrent TCP connections from one source IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions. | 
| tcp_dst_session | If the number of concurrent TCP connections to one destination IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions. | 
| udp_flood | If the UDP traffic to one destination IP address exceeds the configured threshold value, the action is executed. | 2000 packets per second. | 
| udp_scan | If the UDP sessions setup rate originating from one source IP address exceeds the configured threshold value, the action is executed. | 2000 sessions per second. | 
| udp_src_session | If the number of concurrent UDP connections from one source IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions. | 
| udp_dst_session | If the number of concurrent UDP connections to one destination IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions. | 
| icmp_flood | If the number of ICMP packets sent to one destination IP address exceeds the configured threshold value, the action is executed. | 250 packets per second. | 
| icmp_sweep | If the ICMP sessions setup rate originating from one source IP address exceeds the configured threshold value, the action is executed. | 100 sessions per second. | 
| icmp_src_session | If the number of concurrent ICMP connections from one source IP address exceeds the configured threshold value, the action is executed. | 300 concurrent sessions | 
| icmp_dst_session | If the number of concurrent ICMP connections to one destination IP address exceeds the configured threshold value, the action is executed. | 1000 concurrent sessions | 
| ip_src_session | If the number of concurrent IP connections from one source IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions. | 
| ip_dst_session | If the number of concurrent IP connections to one destination IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions. | 
| sctp_flood | If the number of SCTP packets sent to one destination IP address exceeds the configured threshold value, the action is executed. | 2000 packets per second | 
| sctp_scan | If the number of SCTP sessions originating from one source IP address exceeds the configured threshold value, the action is executed. | 1000 packets per second | 
| sctp_src_session | If the number of concurrent SCTP connections from one source IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions | 
| sctp_dst_session | If the number of concurrent SCTP connections to one destination IP address exceeds the configured threshold value, the action is executed. | 5000 concurrent sessions | 

For thresholds based on the number of concurrent sessions, blocking the anomaly will not allow more than the number of concurrent sessions to be set as the threshold.

For example, if the period for a particular anomaly is 60 seconds, such as those where the threshold is measured in concurrent sessions, after the 60 second timer has expired the number of allowed sessions that match the anomaly criteria is reset to zero. This means that, if you allow 10 sessions through before blocking, after the 60 seconds has elapsed, another 10 sessions will be allowed. The attrition of sessions from expiration should keep the allowed sessions from reaching the maximum.

For rate based thresholds, where the threshold is measured in packets per second, the *Block* action prevents anomalous traffic from overwhelming the firewall in two ways:

- 
                                                    continuous: Block packets once an anomaly is detected, and continue to block packets while the rate is above the threshold. This is the default setting.
- 
                                                    periodical: After an anomaly is detected, allow the configured number of packets per second.

For example, if a DoS policy is configured to block icmp_flood with a threshold of 10pps, and a continuous ping is started at a rate of 20pps for 1000 packets:

- 
                                                    In continuous mode, the first 10 packets are passed before the DoS sensor if triggered, and then the remaining 990 packets are blocked.
- 
                                                    In periodical mode, 10 packets are allowed to pass per second, so 500 packets are blocked in the 50 seconds during which the ping is occuring.

|  | The actual numbers of passed and blocked packets may not be exact, as fluctuations in the rates can occur, but the numbers should be close to the defined threshold. | 

###### To configure the block action for rate based anomaly sensors:

```
config ips global
    set anomaly-mode {continuous | periodical}
end
```
                                            ## DoS policies

A DoS policy can be configured to use one or more anomalies.

###### To configure a DoS policy in the GUI:

