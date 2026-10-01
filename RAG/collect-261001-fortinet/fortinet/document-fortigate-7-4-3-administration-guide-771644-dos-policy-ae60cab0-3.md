---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-771644-dos-policy-ae60cab0-3
title: "DoS policy"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2020-11-20"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-771644-dos-policy-ae60cab0.md
source_anchor: ""
source_lines: [151, 200]
sha256: 8d6dc7ca121222befc1cafff2f88c84c502157a992cc3af0d05fc91a3d516422
---

# DoS policy

1. 
                                                    From the Attacker, launch an icmp_flood with 50pps lasting for 3000 packets.
2. 
                                                    On the FortiGate, configure continuous mode and create a DoS policy with an icmp_flood threshold of 30pps: ```
config firewall DoS-policy
    edit 1
        set name icmpFlood
        set interface "port1"
        set srcaddr "all"
        set dstaddr "all"
        set service "ALL"
        config anomaly
            edit "icmp_flood"
                set status enable
                set log enable
                set action block
                set threshold 30
            next
        end
    next
end
```
3. 
                                                    Configure the debugging filter: # diagnose ips anomaly config DoS sensors in kernel vd 0: DoS id 1 proxy 0 0 tcp_syn_flood status 0 log 0 nac 0 action 0 threshold 2000 ... 7 udp_dst_session status 0 log 0 nac 0 action 0 threshold 5000 **8 icmp_flood status 1 log 1 nac 0 action 7 threshold 30** 9 icmp_sweep status 0 log 0 nac 0 action 0 threshold 100
  ...
total # DoS sensors: 1.# diagnose ips anomaly filter id 8
4. 
                                                    Launch the icmp_flood from a Linux machine. This example uses Nmap: $ sudo nping --icmp --rate 50 -c 3000 192.168.2.50 SENT (0.0522s) ICMP [192.168.2.205 > 192.168.2.50 Echo request (type=8/code=0) id=8597 seq=1] IP [ttl=64 id=47459 iplen=28 ] ... Max rtt: 11.096ms | Min rtt: 0.028ms | Avg rtt: 1.665ms Raw packets sent: 3000 (84.000KB) | Rcvd: 30 (840B) | Lost: 2970 (99.00%) Nping done: 1 IP address pinged in 60.35 seconds
5. 
                                                    During the attack, check the anomaly list on the FortiGate: # diagnose ips anomaly list list nids meter: **id=icmp_flood         ip=192.168.2.50 dos_id=1 exp=998 pps=46 freq=50** total # of nids meters: 1.id=icmp_flood The anomaly name. ip=192.168.2.50 The IP address of the host that triggered the anomaly. It can be either the client or the server. For icmp_flood, the IP address is the destination IP address. For icmp_sweep, it would be the source IP address. dos_id=1 The DoS policy ID. exp=998 The time to be expired, in jiffies (one jiffy = 0.01 seconds). pps=46 The number of packets that had been received when the diagnose command was executed. freq=50 For session based anomalies, freq is the number of sessions. For packet rate based anomalies (flood, scan): 
  - 
                                                                            In continuous mode: freq is the greater of pps, or the number of packets received in the last second.
  - 
                                                                            In periodic mode: freq is the pps.
6. 
                                                                            
7. 
                                                    Go to *Log & Report > Security Events* and download the*Anomaly* logs:date=2020-11-20 time=14:38:39 eventtime=1605911919824184594 tz="-0800" logid="0720018433" type="utm" subtype="anomaly" eventtype="anomaly" level="alert" vd="root" severity="critical" **srcip=192.168.2.205** srccountry="Reserved"**dstip=192.168.2.50** srcintf="port1" srcintfrole="undefined" sessionid=0**action="clear_session"** proto=1 service="PING" count=1307 attack="icmp_flood" icmpid="0x2195" icmptype="0x08" icmpcode="0x00" attackid=16777316 policyid=1 policytype="DoS-policy" ref="http://www.fortinet.com/ids/VID16777316"**msg="anomaly: icmp_flood, 31 > threshold 30, repeats 28 times"** crscore=50 craction=4096 crlevel="critical"date=2020-11-20 time=14:39:09 eventtime=1605911949826224056 tz="-0800" logid="0720018433" type="utm" subtype="anomaly" eventtype="anomaly" level="alert" vd="root" severity="critical" **srcip=192.168.2.205** srccountry="Reserved"**dstip=192.168.2.50** srcintf="port1" srcintfrole="undefined" sessionid=0**action="clear_session"** proto=1 service="PING" count=1497 attack="icmp_flood" icmpid="0x2195" icmptype="0x08" icmpcode="0x00" attackid=16777316 policyid=1 policytype="DoS-policy" ref="http://www.fortinet.com/ids/VID16777316"**msg="anomaly: icmp_flood, 50 > threshold 30, repeats 1497 times"** crscore=50 craction=4096 crlevel="critical"###### AnalysisIn the first log message: msg="anomaly: icmp_flood, 31 > threshold 30 At the beginning of the attack, a log is recorded when the threshold of 30pps is broken. repeats 28 times The number of packets that has exceeded the threshold since the last time a log was recorded. srcip=192.168.2.205 dstip=192.168.2.50 The source and destination IP addresses of the attack. action="clear_session" Equivalent to block. If `action` was set to`monitor` and logging was enabled, this would be`action="detected"` .In the second log message: 
  - 
                                                            Because it is an ongoing attack, the FortiGate generates one log message for multiple packets every 30 seconds..
  - 
                                                            It will not generate a log message if: 
    - 
                                                                    The same attack ID happened more than once in a five second period, or
    - 
                                                                    The same attack ID happened more than once in a 30 second period and the actions are the same and have the same source and destination IP addresses.
  - 
                                                                    
 msg="anomaly: icmp_flood, 50 > threshold 30 In the second before the log was recorded, 50 packets were detected, exceeding the configured threshold. repeats 1497 times The number of packets that has exceeded the threshold since the last time a log was recorded
8.
