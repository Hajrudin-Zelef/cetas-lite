---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates-cd95a124-2
title: "document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124.md
source_anchor: ""
source_lines: [140, 205]
sha256: d094f2270885d2c99e16acc97ac6721724dc709ece06dde0af0fb7f9f8e199c6
---

# document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124

                                                            Edit the firewall policy settings: config firewall policy
    edit 1
        set name "to_server_policy"
        set srcintf "port5"
        set dstintf "port6"
        set action accept
        set srcaddr "all"
        set dstaddr "all"
        set schedule "always"
        set service "ALL"
        set logtraffic-start enable
    next
end
- 
                                                            
- 
                                                    On the secondary FortiGate (FG-1800F), verify that the settings were synchronized. 
  - 
                                                            Verify the interface settings: show system interface
    config system interface
        ...
        edit "port5"
            set vdom "root"
            set ip 10.1.100.1 255.255.255.0
            set allowaccess ping https ssh http telnet
            set type physical
            set alias "To_Client_PC"
            set snmp-index 9
            config ipv6
                set ip6-address 2000:10:1:100::1/64
                set ip6-allowaccess ping https ssh http
            end
        next
        edit "port6"
            set vdom "root"
            set ip 172.16.200.1 255.255.255.0
            set allowaccess ping https ssh http fgfm
            set type physical
            set alias "To_Server"
            set snmp-index 10
            config ipv6
                set ip6-address 2000:172:16:200::1/64
                set ip6-allowaccess ping https ssh http
            end
        next
    end
  - 
                                                            Verify the firewall policy settings: show firewall policy
    config firewall policy
            edit 1
            set name "to_server_policy"
            set uuid 82a05e78-fe90-51ed-eb16-ee7bdea60de0
            set srcintf "port5"
            set dstintf "port6"
            set action accept
            set srcaddr "all"
            set dstaddr "all"
            set schedule "always"
            set service "ALL"
            set logtraffic-start enable
        next
    end
  - 
                                                            Verify the HA checksum: # diagnose sys ha checksum show is_manage_primary()=0, is_root_primary()=0 debugzone global: 4e 15 af c3 c6 87 32 f5 69 5c b7 33 b1 8b 27 12 root: 4a 52 e4 f1 6a 2b eb 7d 84 7d f1 48 50 93 fe d9 all: 95 4e 92 c3 39 75 8e 0e db 83 8d b7 b2 b1 9f 04 checksum global: 4e 15 af c3 c6 87 32 f5 69 5c b7 33 b1 8b 27 12 root: 4a 52 e4 f1 6a 2b eb 7d 84 7d f1 48 50 93 fe d9 all: 95 4e 92 c3 39 75 8e 0e db 83 8d b7 b2 b1 9f 04
- 
                                                            
