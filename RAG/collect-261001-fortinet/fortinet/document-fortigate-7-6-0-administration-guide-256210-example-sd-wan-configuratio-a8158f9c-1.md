---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-administration-guide-256210-example-sd-wan-configuratio-a8158f9c-1
title: "Example SD-WAN configurations using ADVPN 2.0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-administration-guide-256210-example-sd-wan-configuratio-a8158f9c.md
source_anchor: ""
source_lines: [1, 205]
sha256: 6aea654af85420a2eeab37608de9c44495771c829e7ac6021d4e5e1f3bf58e66
---

# Example SD-WAN configurations using ADVPN 2.0

## Example SD-WAN configurations using ADVPN 2.0

The configuration example illustrates the edge discovery and path management processes for a typical hub and spoke topology. This example focuses on SD-WAN configuration for steering traffic and establishing shortcuts in the direction from Spoke 1 to Spoke 2.

|  | Currently, ADVPN 2.0 only supports IPv4. | 

In this example, BGP per overlay was used for dynamic routing to distribute the LAN routes behind each spoke to the other spoke. However, this was a design choice. You can also use BGP on loopback for this example.


Spokes 1 and 2 have the following VPN overlays between themselves and the hub:

| VPN Overlays | IP address on Spoke 1 | IP address on Spoke 2 | 
|---|---|---|
| H1_T11 | 172.31.80.1/32 | 172.31.80.2/32 | 
| H1_T22 | 172.31.81.1/32 | 172.31.81.2/32 | 
| H1_T33 | 172.31.82.1/32 | 172.31.82.2/32 | 

SD-WAN Rules/Services defined on Spoke 1:

|  | SD-WAN Rule/Service 1 | SD-WAN Rule/Service 2 | SD-WAN Rule/Service 3 | 
|---|---|---|---|
|    | H1_T11 | H1_T22 | H1_T33 | 
|  | H1_T22 | H1_T11 | H1_T11 | 
|  | H1_T33 | H1_T33 | H1_T22 | 
| Strategy for choosing outgoing interfaces | Lowest cost (SLA) | Lowest cost (SLA) | Best quality, link cost factor: packet loss | 

Throughout this example, transport group 1 is used for VPN overlays over Internet links while transport group 2 is used for the VPN overlay over an MPLS link.

In this example, user traffic is initiated behind Spoke 1 and destined to Spoke 2. Because of this, Spoke 1 is considered the local spoke, and Spoke 2 is considered the remote spoke.

This section includes:

- 
                                                    SD-WAN configuration and health check status on Spoke 1:
- 
                                                    SD-WAN configuration and health check status on Spoke 2:

```
config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
        edit "overlay"
            
```
**set advpn-select enable
            set advpn-health-check "HUB"**
        next
    end
    config members
        edit 1
            set interface "H1_T11"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 2
            set interface "H1_T22"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 3
            set interface "H1_T33"
            set zone "overlay"
            **set transport-group 2** 
        next
    end
    config health-check
        edit "HUB"
            set server "172.31.100.100"
            set members 1 2 3
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                next
            end
        next
    end
    config service
        edit 1
            set name "1"
            **set mode sla
            set shortcut-priority enable**
            set dst "spoke-2_LAN-1" "Tunnel_IPs"
            set src "spoke-1_LAN-1" "Tunnel_IPs"
            config sla
                edit "HUB"
                    set id 1
                next
            end
            **set priority-members 1 2 3**
        next
        edit 2
            set name "2"
            **set mode sla
            set shortcut-priority enable**
            set dst "spoke-2_LAN-2" "Tunnel_IPs"
            set src "spoke-1_LAN-1" "Tunnel_IPs"
            config sla
                edit "HUB"
                    set id 1
                next
            end
            **set priority-members 2 1 3**
        next
        edit 3
            set name "3"
            **set mode priority**
            set dst "spoke-2_LAN-3" "Tunnel_IPs"
            set src "spoke-1_LAN-1" "Tunnel_IPs"
            set health-check "HUB"
            **set link-cost-factor packet-loss
            set priority-members 3 1 2**
        next
    end
end
# diagnose sys sdwan health-check
Health Check(HUB):
Seq(1 H1_T11): state(alive), packet-loss(0.000%) latency(0.231), jitter(0.029), mos(4.404), bandwidth-up(999999), bandwidth-dw(999997), bandwidth-bi(1999996) **sla_map=0x1**
Seq(2 H1_T22): state(alive), packet-loss(0.000%) latency(0.193), jitter(0.010), mos(4.404), bandwidth-up(999994), bandwidth-dw(999997), bandwidth-bi(1999991) **sla_map=0x1**
Seq(3 H1_T33): state(alive), packet-loss(0.000%) latency(0.144), jitter(0.007), mos(4.404), bandwidth-up(999999), bandwidth-dw(999997), bandwidth-bi(1999996) **sla_map=0x1**

```
config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
        edit "overlay"
            
```
**set advpn-select enable
            set advpn-health-check "HUB"**
        next
    end
    config members
        edit 1
            set interface "H1_T11"
            set zone "overlay"
            **set cost 100
            set transport-group 1** 
        next
        edit 2
            set interface "H1_T22"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 3
            set interface "H1_T33"
            set zone "overlay"
            **set transport-group 2** 
        next
    end
    config health-check
        edit "HUB"
            set server "172.31.100.100"
            set members 3 1 2
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                next
            end
        next
    end
end
# diagnose sys sdwan health-check
Health Check(HUB):
Seq(3 H1_T33): state(alive), packet-loss(0.000%) latency(0.124), jitter(0.009), mos(4.404), bandwidth-up(999999), bandwidth-dw(999998), bandwidth-bi(1999997) **sla_map=0x1**
Seq(1 H1_T11): state(alive), packet-loss(0.000%) latency(0.216), jitter(0.043), mos(4.404), bandwidth-up(999999), bandwidth-dw(999998), bandwidth-bi(1999997) **sla_map=0x1**
Seq(2 H1_T22): state(alive), packet-loss(0.000%) latency(0.184), jitter(0.012), mos(4.404), bandwidth-up(999994), bandwidth-dw(999998), bandwidth-bi(1999992) **sla_map=0x1**

In this scenario, PC1 connected to Spoke 1 initiates an ICMP ping destined for PC1 connected to Spoke 2. Therefore, this user traffic matches SD-WAN rule 1 and triggers shortcut path selection and establishment.

The Path Manager of Spoke 1 calculates the best shortcut path by comparing transport group, link quality (for SLA mode), link cost, and member configuration order between Spoke 1 and Spoke 2.

For an SLA mode service, the following algorithm is used to consider endpoints of the best shortcut path:

1. 
                                                    Overlays with the same transport group
2. 
                                                    In-SLA overlays
3. 
                                                    Lowest link-cost overlays
4. 
                                                    Member configuration order as a final tiebreaker

Based on this algorithm, the Path Manager on Spoke 1 selects Spoke 1 H1_T11 because it is first in the priority-members order for SD-WAN rule 1, it has the lowest link cost, and it is within SLA. Likewise, the Path Manager on Spoke 1 selects Spoke 2 H1_T22 since it has the lowest link cost compared to Spoke 2 H1_T11 (which has a cost of 100), it is within SLA, and has the same transport group as Spoke 1 H1_T11. Therefore, the Path Manager of Spoke 1 calculates the best shortcut path as Spoke 1 H1_T11 to Spoke 2 H1_T22.

The Path Manager will advise IKE to establish the best shortcut and add it to SD-WAN rule 1 as follows:

```
Branch1_FGT# diagnose sys sdwan service4
 Service(1): Address Mode(IPV4) flags=0x4200 use-shortcut-sla use-shortcut
  Tie break: cfg
  Shortcut priority: 1
   Gen(11), TOS(0x0/0x0), Protocol(0): src(1->65535):dst(1->65535), Mode(sla), sla-compare-order
   Member sub interface(4):
     2: seq_num(1), interface(H1_T11):
        1: H1_T11_0(71)
   Members(4):
     
