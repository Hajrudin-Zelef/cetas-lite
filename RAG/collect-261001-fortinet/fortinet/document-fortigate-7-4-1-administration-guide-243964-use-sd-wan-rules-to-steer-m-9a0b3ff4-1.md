---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-243964-use-sd-wan-rules-to-steer-m-9a0b3ff4-1
title: "diagnose sys sdwan health-check"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-243964-use-sd-wan-rules-to-steer-m-9a0b3ff4.md
source_anchor: ""
source_lines: [1, 165]
sha256: 5810c9b603e1828b7a4e2e0f223fb5300376fef0854e06f37780102164e1dfe9
---

# diagnose sys sdwan health-check

Use SD-WAN rules to steer multicast traffic
Use SD-WAN rules to steer multicast traffic
SD-WAN rules can now steer multicast traffic. When an SD-WAN member is out of SLA, multicast traffic can fail over to another SD-WAN member, and switch back when SLA recovers.
The new pim-use-sdwan option enables or disables the use of SD-WAN for PIM (Protocol Independent Multicast) when checking  RP (Rendezvous Point) neighbors and sending packets.
config router multicast
    config pim-sm-global
        set pim-use-sdwan {enable | disable}
    end
end
                                            |  | When SD-WAN steers multicast traffic, ADVPN is not supported. Use the set shortcut option to disable shortcuts for the service: config system sdwan     config service         edit <id>             set shortcut {enable \| disable}         next     end end | 
Example 1
In this hub and spoke example, the PIM source is behind the hub FortiGate, and the RP is set to internal port (port2) of the hub firewall. Each spoke connects to the two WAN interfaces on the hub by using an overlay tunnel. The overlay tunnels are members of SD-WAN.
Receivers behind the spoke FortiGates request a stream from the source to receive traffic on tunnel1 by default. When the overlay tunnel goes out of SLA, the multicast traffic fails over to tunnel2 and continues to flow.
Following is an overview of how to configure the topology:
- Configure the hub FortiGate in front of the PIM source. The RP is configured on internal port (port2) of the hub FortiGate.
- Configure the spoke FortiGates.
- Verify traffic failover.
To configure the hub:
- 
                                                    On the hub, enable multicast routing, configure the multicast RP, and enable PIM sparse mode on each interface: config router multicast
    set multicast-routing enable
    config pim-sm-global
        config rp-address
            edit 1
                set ip-address 172.16.205.1
            next
        end
    end
    config interface
        edit "tport1"
            set pim-mode sparse-mode
        next
        edit "tagg1"
            set pim-mode sparse-mode
        next
        edit "port2"
            set pim-mode sparse-mode
        next
    end
end
To configure each spoke:
- 
                                                    Enable SD-WAN with the following settings: 
  - Configure the overlay tunnels as member of the SD-WAN zone.
  - Configure a performance SLA health-check using ping.
  - Configure a service rule for the PIM protocol with the following settings:
    - Use the lowest cost (SLA) strategy.
    - Monitor with the ping health-check.
  - Disable ADVPN shortcut.
 config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
    end
    config members
        edit 1
            set interface "tunnel1"
        next
        edit 2
            set interface "tunnel2"
        next
    end
    config health-check
        edit "ping"
            set server "172.16.205.1"
            set update-static-route disable
            set members 0
            config sla
                edit 1
                next
            end
        next
    end
    config service
        edit 1
            set mode sla
            set protocol 103
            set dst "all"
            config sla
                edit "ping"
                    set id 1
                next
            end
            set priority-members 1 2
            set use-shortcut-sla disable
            set shortcut disable
        next
        edit 2
            set mode sla
            set dst "all"
            config sla
                edit "ping"
                    set id 1
                next
            end
            set priority-members 1 2
        next
    end
end
- 
                                                    Enable multicast routing and configure the multicast RP. Enable PIM sparse-mode on each interface: config router multicast
    set multicast-routing enable
    config pim-sm-global
        set spt-threshold disable
        set pim-use-sdwan enable
        config rp-address
            edit 1
                set ip-address 172.16.205.1
            next
        end
    end
    config interface
        edit "tunnel1"
            set pim-mode sparse-mode
        next
        edit "tunnel2"
            set pim-mode sparse-mode
        next
        edit "port4"
            set pim-mode sparse-mode
        next
    end
end
To verify traffic failover:
With this configuration, multicast traffic starts on tunnel1. When tunnel1 becomes out of SLA, traffic switches to tunnel2. When tunnel1 is in SLA again, the traffic switches back to tunnel1.
The following health-check capture on the spokes shows tunnel1 in SLA with packet-loss (1.000%):
# diagnose sys sdwan health-check
Health Check(ping):
Seq(1 tunnel1): state(alive), packet-loss(0.000%) latency(0.056), jitter(0.002), mos(4.404), bandwidth-up(999999), bandwidth-dw(1000000), bandwidth-bi(1999999) sla_map=0x1
Seq(2 tunnel2): state(alive), packet-loss(0.000%) latency(0.100), jitter(0.002), mos(4.404), bandwidth-up(0), bandwidth-dw(0), bandwidth-bi(0) sla_map=0x1
# diagnose sys sdwan health-check
Health Check(ping):
Seq(1 tunnel1): state(alive), packet-loss(1.000%) latency(0.056), jitter(0.002), mos(4.404), bandwidth-up(999999), bandwidth-dw(1000000), bandwidth-bi(1999999) sla_map=0x1
Seq(2 tunnel2): state(alive), packet-loss(0.000%) latency(0.100), jitter(0.002), mos(4.404), bandwidth-up(0), bandwidth-dw(0), bandwidth-bi(0) sla_map=0x1
The following example shows tunnel1 out of SLA with packet-loss (3.000%):
# diagnose sys sdwan health-check
Health Check(ping):
Seq(1 tunnel1): state(alive), packet-loss(3.000%) latency(0.057), jitter(0.003), mos(4.403), bandwidth-up(999999), bandwidth-dw(1000000), bandwidth-bi(1999999) sla_map=0x0
Seq(2 tunnel2): state(alive), packet-loss(0.000%) latency(0.101), jitter(0.002), mos(4.404), bandwidth-up(0), bandwidth-dw(0), bandwidth-bi(0) sla_map=0x1
The following example shows tunnel1 back in SLA again:
# diagnose sys sdwan health-check
Health Check(ping):
Seq(1 tunnel1): state(alive), packet-loss(1.000%) latency(0.061), jitter(0.004), mos(4.404), bandwidth-up(999999), bandwidth-dw(1000000), bandwidth-bi(1999999) sla_map=0x0
Seq(2 tunnel2): state(alive), packet-loss(0.000%) latency(0.102), jitter(0.002), mos(4.404), bandwidth-up(0), bandwidth-dw(0), bandwidth-bi(0) sla_map=0x1
# diagnose sys sdwan health-check
Health Check(ping):
Seq(1 tunnel1): state(alive), packet-loss(0.000%) latency(0.061), jitter(0.004), mos(4.404), bandwidth-up(999999), bandwidth-dw(1000000), bandwidth-bi(1999999) sla_map=0x0
Seq(2 tunnel2): state(alive), packet-loss(0.000%) latency(0.102), jitter(0.002), mos(4.404), bandwidth-up(0), bandwidth-dw(0), bandwidth-bi(0) sla_map=0x1
The following example how traffic switches to tunnel2 while tunnel1 health-check is out of SLA. Source (172.16.205.11) sends traffic to the multicast group. Later the traffic switches back to tunnel1 once SLA returns to normal:
195.060797 tunnel1 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
195.060805 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
196.060744 tunnel1 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
196.060752 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
197.060728 tunnel1 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
197.060740 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
198.060720 tunnel2 in 172.16.205.11 -> 225.1.1.1: icmp: echo request        
198.060736 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
199.060647 tunnel2 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
199.060655 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
200.060598 tunnel2 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
200.060604 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
... ...
... ...
264.060974 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
