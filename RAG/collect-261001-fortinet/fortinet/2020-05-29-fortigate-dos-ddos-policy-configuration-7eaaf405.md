---
id: collect-261001-fortinet/fortinet/2020-05-29-fortigate-dos-ddos-policy-configuration-7eaaf405
title: "2020-05-29-fortigate-dos-ddos-policy-configuration-7eaaf405"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/2020-05-29-fortigate-dos-ddos-policy-configuration-7eaaf405.md
source_anchor: ""
source_lines: [1, 184]
sha256: ba83e9040ed07a77f797a390be5a2ddce14b14bd7b9bb1e7efc575ca1afd6e20
---

# 2020-05-29-fortigate-dos-ddos-policy-configuration-7eaaf405

Facts to know:

- You use Dos protection by creating Dos policy (**Policy & Objects -> IPv4/Ipv6 DoS Policy** ) in which you enable/modify  anomalies.
- The list of anomalies is pre-set in any policy you create. You only have the choice which ones to enable and which ones not to.
- All anomalies are set by default to **Pass** the offending traffic and are disabled, so make sure under the given anomaly to set`status  enable` and action to`block` . On Fortigates with hardware NP modules,  you also have**Proxy** as an action in`tcp_syn_flood` protection to enable, which makes Fortigate to proxy SYN connections.
- You can (actually must) specify: source/destination IPs to match the DoS policy (`all` can be used), service (`ALL` can be used), and incoming interface to apply the DoS policy  to.
- Thresholds for anomalies are configurable and do what they say - once traffic matched by this policy exceeds the threshold, it gets blocked. No learning or adaptive thresholds here.
- By default, only exceeding the threshold packets get blocked. To block the sender IP completely, you can use `set qurantine` parameter under the specific anomaly.
- Dos sensor/policy protects against INCOMING traffic for the specified interface.
- For smarter anti-DDoS solution Fortinet have FortiDDoS physical appliance.
- Fortigate applies Dos protection early in the policy matching, before the Security policy is checked, so it consumes less resources than blocking the same traffic in Security rules. This means, though, that even if some security rule allows traffic, if such traffic exceeds DoS thresholds it may be blocked.
- Note: in previous versions of FortiOS the feature  was called *DoS sensor* , so I mention it for easier reference only. In FortiOS 6.x and newer it is called DoS Policy.
- From my personal experience, to protect large networks with this DoS feature of Fortigate is more hassle than help. The false positives, especially for TCP SYN and alike protections, would block legitimate clients to the internal servers available from the Internet due to sudden surge of the client requests. You would need then to fix the thresholds, then again... For small networks, and those that do not have accessible from the outside servers, it may be a nice to have feature.

### Configuring DoS policy.

I enable just `icmp_flood` anomaly here and change the threshold to 10 packets per second sent to destination of 12.12.12.3 :  

```
config firewall DoS-policy
    edit 1
        set interface "port1"
        set srcaddr "all"
        set dstaddr "12.12.12.3"
        set service "ALL_ICMP"
        config anomaly
            edit "tcp_syn_flood"
                set threshold 2000
            next
            edit "tcp_port_scan"
                set threshold 1000
            next
            edit "tcp_src_session"
                set threshold 5000
            next
            edit "tcp_dst_session"
                set threshold 5000
            next
            edit "udp_flood"
                set threshold 2000
            next
            edit "udp_scan"
                set threshold 2000
            next
            edit "udp_src_session"
                set threshold 5000
            next
            edit "udp_dst_session"
                set threshold 5000
            next
            edit "icmp_flood"
                set status enable
                set action block
                set threshold 10
            next
            edit "icmp_sweep"
                set threshold 100
            next
            edit "icmp_src_session"
                set threshold 300
            next
            edit "icmp_dst_session"
                set threshold 1000
            next
            edit "ip_src_session"
                set threshold 5000
            next
            edit "ip_dst_session"
                set threshold 5000
            next
            edit "sctp_flood"
                set threshold 2000
            next
            edit "sctp_scan"
                set threshold 1000
            next
            edit "sctp_src_session"
                set threshold 5000
            next
            edit "sctp_dst_session"
                set threshold 5000
            next
        end
    next
end
```
### Verification

Sending 5 packets per second, traffic is NOT blocked:

```
root@ubuntu:~# ping -i 0.2   12.12.12.3
PING 12.12.12.3 (12.12.12.3) 56(84) bytes of data.
64 bytes from 12.12.12.3: icmp_seq=1 ttl=255 time=3.03 ms
64 bytes from 12.12.12.3: icmp_seq=2 ttl=255 time=1.96 ms
64 bytes from 12.12.12.3: icmp_seq=3 ttl=255 time=0.469 ms
64 bytes from 12.12.12.3: icmp_seq=4 ttl=255 time=0.318 ms
64 bytes from 12.12.12.3: icmp_seq=5 ttl=255 time=0.405 ms
64 bytes from 12.12.12.3: icmp_seq=6 ttl=255 time=0.497 ms
```
Sending roughly 10 packets per second - Fortigate starts to block excessive icmp packets.

```
root@ubuntu:~# ping -i 0.1   12.12.12.3
PING 12.12.12.3 (12.12.12.3) 56(84) bytes of data.
64 bytes from 12.12.12.3: icmp_seq=1 ttl=255 time=1.33 ms
64 bytes from 12.12.12.3: icmp_seq=2 ttl=255 time=0.712 ms
...
--- 12.12.12.3 ping statistics ---
143 packets transmitted, 115 received, 19% packet loss, time 14526ms
```
To see the active attacks/blocked anomalies (block happens when `freq` goes 10 or higher):    

### diagnose ips anomaly list

```
FG3-AS1680 # diagnose ips anomaly list
list nids meter:
id=icmp_flood         ip=12.12.12.3 dos_id=1 exp=993 pps=1 freq=14
```
Next, I add second policy with destination address `all` but also with qurantine enabled.

```
config firewall DoS-policy
    edit 2
        set interface "port1"
        set srcaddr "all"
        set dstaddr "all"
        set service "ALL"
        config anomaly
            edit "icmp_flood"
                set status enable
                set log enable
                set quarantine attacker
                set quarantine-expiry 2m  <-- to set to 2 min I entered: 000d00h02m
                set quarantine-log disable
                set threshold 10
            next
```
Exceeding the threshold:

```
root@ubuntu:~# ping -c 2000 -i 0.01  13.13.13.6
PING 13.13.13.6 (13.13.13.6) 56(84) bytes of data.
64 bytes from 13.13.13.6: icmp_seq=1 ttl=254 time=0.741 ms
64 bytes from 13.13.13.6: icmp_seq=2 ttl=254 time=1.82 ms
64 bytes from 13.13.13.6: icmp_seq=3 ttl=254 time=1.89 ms
--- 13.13.13.6 ping statistics ---
2000 packets transmitted, 11 received, 99% packet loss, time 24308ms
```
As you can see, 1st 10 packets were allowed, the 11th packet triggered the following block.

```
FG3-AS1680 # diagnose ips anomaly list
id=icmp_flood         ip=13.13.13.6 dos_id=2 exp=998 pps=38 freq=83
```
Also, because I set qurantine period for 2 minutes, even after stopping the attack traffic, the sending server is blocked from sending ANY packets to the target 13.13.13.6 for the next 2 minutes:

```
root@ubuntu:~# ping 13.13.13.6
PING 13.13.13.6 (13.13.13.6) 56(84) bytes of data.
--- 13.13.13.6 ping statistics ---
11 packets transmitted, 0 received, 100% packet loss, time 10029ms
```
### Releasing the blocked senders

Fortigate does not show us the source IPs of the blocked hosts, just the target IP, still, we can clear the blocked attackers list and allow the blocked senders to pass through. If they again send the excessive traffic, they will be blocked again, i.e the clear action is real-time and not permanent. Also, for the senders blocked with the quarantine, clearing the list will still keep them blocked until the qurantine expiration.

```
FG3-AS1680 # diagnose ips anomaly list
list nids meter:
id=icmp_flood         ip=12.12.12.3 dos_id=1 exp=999 pps=2 freq=20
```
Clear the list:

### diagnose ips anomaly clear

```
FG3-AS1680 # diagnose ips anomaly clear
FG3-AS1680 # diagnose ips anomaly list
list nids meter:
total # of nids meters: 0.
```
*Follow me on https://www.linkedin.com/in/yurislobodyanyuk/  not to miss what I publish on Linkedin, Github, blog, and more.*
