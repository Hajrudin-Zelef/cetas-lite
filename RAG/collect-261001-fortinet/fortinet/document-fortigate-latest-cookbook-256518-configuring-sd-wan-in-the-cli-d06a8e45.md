---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-cookbook-256518-configuring-sd-wan-in-the-cli-d06a8e45
title: "Configuring SD-WAN in the CLI"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-cookbook-256518-configuring-sd-wan-in-the-cli-d06a8e45.md
source_anchor: ""
source_lines: [1, 109]
sha256: a68d0e95f180c70b1171c47830f5858b8b7ec5f73231b7659c6b231b1e0e417e
---

# Configuring SD-WAN in the CLI

# Configuring SD-WAN in the CLI

This example can be entirely configured using the CLI.

###### To configure SD-WAN in the CLI:

1. Configure the wan1 and wan2 interfaces:```
config system interface 
    edit "wan1"
        set alias to_ISP1
        set mode dhcp
        set distance 10
    next
    edit "wan2"
        set alias to_ISP2
        set ip 10.100.20.1 255.255.255.0
    next
end
```
2. Enable SD-WAN and add the interfaces as members:```
config system virtual-wan-link
    set status enable
    config members
        edit 1
            set interface "wan1"
        next
        edit 2
            set interface "wan2"
            set gateway 10.100.20.2
        next
    end
end
```
3. Create a static route for SD-WAN:```
config router static
    edit 1
        set virtual-wan-link enable
    next
end
```
4. Select the implicit SD-WAN algorithm:```
config system virtual-wan-link
    set load-balance-mode {source-ip-based | weight-based | source-dest-ip-based | measured-volume-based}
end
```
5. Create a firewall policy for SD-WAN:```
config firewall policy
    edit <policy_id>
        set name <policy_name>
        set srcintf internal
        set dstintf virtual-wan-link
        set srcaddr all
        set dstaddr all
        set action accept
        set schedule always
        set service ALL
        set utm-status enable
        set ssl-ssh-profile <profile_name>
        set av-profile <profile_name>
        set webfilter-profile <profile_name>			
        set dnsfilter-profile <profile_name>
        set application-list <app_list>
        set logtraffic all
        set nat enable
        set status enable
    next
end
```
6. Configure a performance SLA:```
config system virtual-wan-link
    config health-check
        edit "server"
            set server "208.91.112.53"
            set update-static-route enable
            set members 1 2
        next
    end
end
```

## Results

###### To view the routing table in the CLI:

```
# get router info routing-table all
 
Routing table for VRF=0
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default
 
S*      0.0.0.0/0 [1/0] via 172.16.20.2, wan1
                  [1/0] via 10.100.20.2, wan2
C       10.100.20.0/24 is directly connected, wan2
C       172.16.20.2/24 is directly connected, wan1
C       192.168.0.0/24 is directly connected, internal
```
                                            ###### To diagnose the Performance SLA status:

FGT # **diagnose sys virtual-wan-link health-check**
Health Check(server):
Seq(1): state(alive), packet-loss(0.000%) latency(15.247), jitter(5.231) sla_map=0x0
Seq(2): state(alive), packet-loss(0.000%) latency(13.621), jitter(6.905) sla_map=0x0
