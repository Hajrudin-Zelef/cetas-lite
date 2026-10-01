---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-administration-guide-256210-example-sd-wan-configuratio-a8158f9c-6
title: "Example SD-WAN configurations using ADVPN 2.0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-administration-guide-256210-example-sd-wan-configuratio-a8158f9c.md
source_anchor: ""
source_lines: [754, 807]
sha256: c73dfe85ad7c01b6444e2c5576d41fcdc652e7085862ddea3aedd67ecc507bbc
---

# Example SD-WAN configurations using ADVPN 2.0

Branch2_FGT# diagnose sys sdwan  health-check
 Health Check(HUB):
 Seq(3 H1_T33): state(alive), packet-loss(0.000%) latency(0.120), jitter(0.007), mos(4.404), bandwidth-up(999999), bandwidth-dw(999997), bandwidth-bi(1999996) sla_map=0x1
 Seq(3 H1_T33_0): state(alive), packet-loss(0.000%) latency(0.128), jitter(0.003), mos(4.404), bandwidth-up(1000000), bandwidth-dw(1000000), bandwidth-bi(2000000) sla_map=0x1
 Seq(1 H1_T11): state(alive), packet-loss(0.000%) latency(0.180), jitter(0.008), mos(4.404), bandwidth-up(999999), bandwidth-dw(999997), bandwidth-bi(1999996) sla_map=0x1
 **Seq(1 H1_T11_0): state(alive), packet-loss(0.000%) latency(0.259), jitter(0.023), mos(4.404), bandwidth-up(1000000), bandwidth-dw(1000000), bandwidth-bi(2000000) sla_map=0x1 
 Seq(1 H1_T11_1): state(alive), packet-loss(0.000%) latency(0.257), jitter(0.014), mos(4.404), bandwidth-up(1000000), bandwidth-dw(1000000), bandwidth-bi(2000000) sla_map=0x1 
 Seq(2 H1_T22): state(alive), packet-loss(0.000%) latency(250.169), jitter(0.009), mos(3.993), bandwidth-up(999994), bandwidth-dw(999997), bandwidth-bi(1999991) sla_map=0x0** 
 Seq(2 H1_T22_1): state(alive), packet-loss(0.000%) latency(0.245), jitter(0.013), mos(4.404), bandwidth-up(1000000), bandwidth-dw(1000000), bandwidth-bi(2000000) sla_map=0x1
 Seq(2 H1_T22_0): state(alive), packet-loss(0.000%) latency(0.223), jitter(0.005), mos(4.404), bandwidth-up(1000000), bandwidth-dw(1000000), bandwidth-bi(2000000) sla_map=0x1 

At this point, PC1 connected to Spoke 1 initiated multiple ICMP pings destined for PC1 connected to Spoke 2. The packet capture diagnostic command on Spoke 1 demonstrates that these ICMP pings have been load balanced over all shortcuts:

Branch1_FGT# diagnose sniffer packet any 'host 10.0.4.2' 4
interfaces=[any]
filters=[host 10.0.4.2]
3.481994 port4 in 10.0.3.2 -> 10.0.4.2: icmp: echo request
3.482103 **H1_T11_1** out 10.0.3.2 -> 10.0.4.2: icmp: echo request
3.482799 **H1_T11_1** in 10.0.4.2 -> 10.0.3.2: icmp: echo reply
3.482928 port4 out 10.0.4.2 -> 10.0.3.2: icmp: echo reply
4.614480 port4 in 10.0.3.2 -> 10.0.4.2: icmp: echo request
4.614580 **H1_T33_0** out 10.0.3.2 -> 10.0.4.2: icmp: echo request
4.615122 **H1_T33_0** in 10.0.4.2 -> 10.0.3.2: icmp: echo reply
4.615152 port4 out 10.0.4.2 -> 10.0.3.2: icmp: echo reply
5.286394 port4 in 10.0.3.2 -> 10.0.4.2: icmp: echo request
5.286497 **H1_T22_0** out 10.0.3.2 -> 10.0.4.2: icmp: echo request
5.287129 **H1_T22_0** in 10.0.4.2 -> 10.0.3.2: icmp: echo reply
5.287155 port4 out 10.0.4.2 -> 10.0.3.2: icmp: echo reply
6.079759 port4 in 10.0.3.2 -> 10.0.4.2: icmp: echo request
6.079883 **H1_T22_1** out 10.0.3.2 -> 10.0.4.2: icmp: echo request
6.080496 **H1_T22_1** in 10.0.4.2 -> 10.0.3.2: icmp: echo reply
6.080537 port4 out 10.0.4.2 -> 10.0.3.2: icmp: echo reply
7.983357 port4 in 10.0.3.2 -> 10.0.4.2: icmp: echo request
7.983447 **H1_T11_0** out 10.0.3.2 -> 10.0.4.2: icmp: echo request
7.984078 **H1_T11_0** in 10.0.4.2 -> 10.0.3.2: icmp: echo reply
7.984120 port4 out 10.0.4.2 -> 10.0.3.2: icmp: echo reply

Without user traffic traversing the shortcut during the idle interval time, from the diagnostic command on Spoke 1, observe that all shortcuts have been removed:

```
Branch1_FGT# diagnose system sdwan service4
Service(1): Address Mode(IPV4) flags=0x24200 use-shortcut-sla use-shortcut
 Tie break: cfg
 Shortcut priority: 3
  Gen(16), TOS(0x0/0x0), Protocol(0): src(1->65535):dst(1->65535), Mode(sla  hash-mode=round-robin)
  Members(3):
    1: Seq_num(1 H1_T11 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    2: Seq_num(2 H1_T22 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    3: Seq_num(3 H1_T33 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
  Src address(1):
        10.0.0.0-10.255.255.255
  Dst address(1):
        10.0.0.0-10.255.255.255
```
