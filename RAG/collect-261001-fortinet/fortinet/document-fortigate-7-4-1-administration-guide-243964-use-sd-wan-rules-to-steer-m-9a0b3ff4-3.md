---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-243964-use-sd-wan-rules-to-steer-m-9a0b3ff4-3
title: "diagnose sys sdwan health-check"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-243964-use-sd-wan-rules-to-steer-m-9a0b3ff4.md
source_anchor: ""
source_lines: [319, 457]
sha256: 0ba0e3e41b71680b4a4b58398e4432532945799aab03ce75ebada29382b28bdd
---

# diagnose sys sdwan health-check

                                                    Configure the default gateway to use the SD-WAN zone. Other routes are for the underlay to route traffic to the hub's WAN interfaces: config router static edit 10 set distance 1 set sdwan-zone "virtual-wan-link" next .... next end
To configure the spoke (in front of the source):
- 
                                                    Enable multicast routing to use SD-WAN. Configure the RP address. Enable interfaces for PIM sparse-mode: config router multicast
    set multicast-routing enable
    config pim-sm-global
        set pim-use-sdwan enable
        config rp-address
            edit 1
                set ip-address 192.90.1.11
            next
        end
    end
    config interface
        edit "p198"
            set pim-mode sparse-mode
        next
        edit "p200"
            set pim-mode sparse-mode
        next
        edit "npu0_vlink0"
            set pim-mode sparse-mode
        next
    end
end
- 
                                                    Configure loopback interface lo66 for BGP and sourcing SD-WAN traffic: config system interface
    edit "lo66"
        set vdom "root"
        set ip 172.31.0.66 255.255.255.255
        set allowaccess ping
        set type loopback
        set snmp-index 21
    next
end
- 
                                                    Configure SD-WAN: 
  - Add overlay tunnel interfaces as members.
  - Configure a performance SLA health-check to send ping probes to the hub.
  - Configure a service rule for the PIM protocol. Use the lowest cost (SLA) strategy, and monitor with the ping health-check.
  - Disable the use of an ADVPN shortcut.
 In the following example, 11.11.11.11 is the underlay address for one of the WAN links on the hub, and 172.31.100.100 is the loopback address on the server. config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
        edit "overlay"
        next
    end
    config members
        edit 1
            set interface "p198"
            set zone "overlay"
            set source 172.31.0.66
        next
        edit 2
            set interface "p200"
            set zone "overlay"
            set source 172.31.0.66
        next
    end
    config health-check
        edit "ping"
            set server "11.11.11.11"            
            set members 0
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                next
            end
        next
        edit "HUB"
            set server "172.31.100.100"        
            set embed-measured-health enable
            set members 0
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
                                                    Configure BGP: config router bgp
    set as 65505
    set router-id 123.1.1.123
    set ibgp-multipath enable
    set additional-path enable
    config neighbor
        edit "172.31.0.1"
            set next-hop-self enable
            set soft-reconfiguration enable
            set remote-as 65505
            set update-source "lo66"
        next
    end
    config network
        edit 3
            set prefix 192.87.0.0 255.255.0.0
        next
    end
end
- 
                                                    Configure the default gateway to use the SD-WAN zone. Other routes are for the underlay to route to the hub's WAN interfaces: config router static
    edit 10
        set distance 1
        set sdwan-zone "virtual-wan-link" "overlay"
    next
    ...
    next
end
