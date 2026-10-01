---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-243964-use-sd-wan-rules-to-steer-m-9a0b3ff4-2
title: "diagnose sys sdwan health-check"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-243964-use-sd-wan-rules-to-steer-m-9a0b3ff4.md
source_anchor: ""
source_lines: [166, 318]
sha256: 6c0209e5bfe7a723a4f11066c5c838a86f859f6d8519c782f65c2fa2cece4c37
---

# diagnose sys sdwan health-check

265.060950 tunnel2 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
265.060958 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
266.060867 tunnel2 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
266.060877 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
267.060828 tunnel2 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
267.060835 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
268.060836 tunnel1 in 172.16.205.11 -> 225.1.1.1: icmp: echo request          
268.060854 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
269.060757 tunnel1 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
269.060767 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
270.060645 tunnel1 in 172.16.205.11 -> 225.1.1.1: icmp: echo request
270.060653 port4 out 172.16.205.11 -> 225.1.1.1: icmp: echo request
Example 2
In this hub and spoke example, the PIM source is behind spoke 1, and the RP is configured on the hub FortiGate. BGP is used for routing. The hub uses embedded SLA in ICMP probes to determine the health of each tunnel, allowing it to prioritize healthy IKE routes.
The receiver is on another spoke. Upon requesting a stream, source passes the traffic to the RP on the hub FortiGate, and routes the traffic to the receiver over tunnel1. If a tunnel falls out of SLA, the multicast traffic fails over to the other tunnel.
In this configuration, SD-WAN steers multicast traffic by using embedded SLA information in ICMP probes. See also Embedded SD-WAN SLA information in ICMP probes. With this feature, the hub FortiGate can use the SLA information of the spoke's health-check to control BGP and IKE routes over tunnels.
Following is an overview of how to configure the topology:
- Configure the hub FortiGate. The RP is configured on the hub FortiGate.
- Configure the spoke FortiGate in front of the traffic receiver.
- Configure the spoke FortiGate in front of the PIM source.
To configure the hub:
- 
                                                    Configure loopbacks hub-lo1 172.31.0.1 for BGP and hub-lo100 172.31.100.100 for health-check: config system interface
    edit "hub-lo1"
        set vdom "hub"
        set ip 172.31.0.1 255.255.255.255
        set allowaccess ping
        set type loopback
        set snmp-index 82
    next
    edit "hub-lo100"
        set vdom "hub"
        set ip 172.31.100.100 255.255.255.255
        set allowaccess ping
        set type loopback
        set snmp-index 81
    next
end
- 
                                                    Enable multicast routing with the following settings: 
  - Configure internal interface p25-v90 as RP.
  - Enable interfaces for PIM sparse-mode.
 config router multicast
    set multicast-routing enable
    config pim-sm-global
        config rp-address
            edit 1
                set ip-address 192.90.1.11
            next
        end
    end
    config interface
        edit "p11"
            set pim-mode sparse-mode
        next
        edit "p101"
            set pim-mode sparse-mode
        next
        edit "p25-v90"
            set pim-mode sparse-mode
        next
    end
end
- 
                                                    Enable SD-WAN with the following settings: 
  - Add interfaces p11 and p101 as members.
  - Configure embedded SLA health-checks to detect ICMP probes from each overlay tunnel. Prioritize based on the health of each tunnel.
 config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
    end
    config members
        edit 1
            set interface "p11"
        next
        edit 2
            set interface "p101"
        next
    end
    config health-check
        edit "1"
            set detect-mode remote
            set probe-timeout 60000
            set recoverytime 1
            set sla-id-redistribute 1
            set members 1
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                    set priority-in-sla 10
                    set priority-out-sla 20
                next
            end
        next
        edit "2"
            set detect-mode remote
            set probe-timeout 60000
            set recoverytime 1
            set sla-id-redistribute 1
            set members 2
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                    set priority-in-sla 15
                    set priority-out-sla 25
                next
            end
        next
    end
end
- 
                                                    Configure BGP to peer with neighbors. Neighbor group is configured for tunnel interface IP addresses: config router bgp set as 65505 set router-id 172.31.0.1 set ibgp-multipath enable set additional-path enable set recursive-inherit-priority enable config neighbor-group edit "gr1" set remote-as 65505 set update-source "hub-lo1" set additional-path both set route-reflector-client enable next end config neighbor-range edit 1 set prefix 10.10.0.0 255.255.0.0 set neighbor-group "gr1" next edit 66 set prefix 172.31.0.66 255.255.255.255 set neighbor-group "gr1" next end config network .... edit 90 set prefix 192.90.0.0 255.255.0.0 next end end
To configure the spoke (in front of the receiver):
- 
                                                    Enable multicast routing to use SD-WAN. Configure the RP address. Enable interfaces for PIM sparse-mode. config router multicast set multicast-routing enable config pim-sm-global set spt-threshold disable set pim-use-sdwan enable config rp-address edit 1 set ip-address 192.90.1.11 next end end config interface edit "p195" set pim-mode sparse-mode next edit "p196" set pim-mode sparse-mode next edit "internal4" set pim-mode sparse-mode set static-group "225-1-1-122" next end end
- 
                                                    Configure SD-WAN with the following settings: 
  - Add overlay tunnel interfaces as members.
  - Configure a performance SLA health-check to send ping probes to the hub.
  - Configure a service rule for the PIM protocol. Use the lowest cost (SLA) strategy, and monitor with the ping health-check.
  - Disable ADVPN shortcuts.
 config system sdwan set status enable config zone edit "virtual-wan-link" next end config members edit 6 set interface "p196" next edit 5 set interface "p195" next end config health-check edit "ping" set server "172.31.100.100" set update-static-route disable set members 0 config sla edit 1 set link-cost-factor latency set latency-threshold 100 next end next end config service edit 1 set mode sla set protocol 103 set dst "all" config sla edit "ping" set id 1 next end set priority-members 5 6 set use-shortcut-sla disable set shortcut disable next edit 2 set mode sla set dst "all" config sla edit "ping" set id 1 next end set priority-members 5 6 next end end
- 
                                                    Configure BGP and set neighbors to the overlay gateway IP address on the hub: config router bgp
    set as 65505
    set router-id 122.1.1.122
    set ibgp-multipath enable
    set additional-path enable
    config neighbor
        edit "10.10.100.254"
            set soft-reconfiguration enable
            set remote-as 65505
            set connect-timer 10
            set additional-path both
        next
        edit "10.10.101.254"
            set soft-reconfiguration enable
            set remote-as 65505
            set connect-timer 10
            set additional-path both
        next
    end
    config network
        edit 3
            set prefix 192.84.0.0 255.255.0.0
        next
    end
end
- 
