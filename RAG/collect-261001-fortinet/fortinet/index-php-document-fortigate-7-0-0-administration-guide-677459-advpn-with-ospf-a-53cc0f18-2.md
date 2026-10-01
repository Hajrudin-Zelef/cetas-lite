---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18-2
title: "index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18.md
source_anchor: ""
source_lines: [237, 350]
sha256: cc3d5628a410d3b78ed72f34612e8cdde1786f65c5f5e2cbc2ecfa5916ae2e01
---

# index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18

  - 
                                                            Configure the spoke FortiGates' firewall policies: 
    - 
                                                                    Configure Spoke1: config firewall policy
    edit 1
        set name "outbound_advpn"
        set srcintf "internal"
        set dstintf "spoke1" "spoke1_backup"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
    next
    edit 2
        set name "inbound_advpn"
        set srcintf "spoke1" "spoke1_backup"
        set dstintf "internal"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
    next
end
    - 
                                                                    Configure Spoke2: config firewall policy
    edit 1
        set name "outbound_advpn"
        set srcintf "internal"
        set dstintf "spoke2" "spoke2_backup"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
    next
    edit 2
        set name "inbound_advpn"
        set srcintf "spoke2" "spoke2_backup"
        set dstintf "internal"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
    next
end
  - 
                                                                    
  - 
                                                            Configure the spoke FortiGates' tunnel interface IP addresses: 
    - 
                                                                    Configure Spoke1: config system interface
    edit "spoke1"
        set ip 10.10.10.1 255.255.255.255
        set remote-ip 10.10.10.254 255.255.255.0
    next
    edit "spoke1_backup"
        set ip 10.10.10.2 255.255.255.255
        set remote-ip 10.10.10.254 255.255.255.0
    next    
end
    - 
                                                                    Configure Spoke2: config system interface
    edit "spoke2"
        set ip 10.10.10.3 255.255.255.255
        set remote-ip 10.10.10.254 255.255.255.0
    next
    edit "spoke2_backup"
        set ip 10.10.10.4 255.255.255.255
        set remote-ip 10.10.10.254 255.255.255.0
    next    
end
  - 
                                                                    
  - 
                                                            Configure the spoke FortiGates' OSPF: 
    - 
                                                                    Configure Spoke1: config router ospf
    set router-id 7.7.7.7
    config area
        edit 0.0.0.0
        next
    end
    config network
        edit 1
            set prefix 10.10.10.0 255.255.255.0
        next
        edit 2
            set prefix 10.1.100.0 255.255.255.0
        next
    end
end
    - 
                                                                    Configure Spoke2: config router ospf
    set router-id 8.8.8.8
    config area
        edit 0.0.0.0
        next
    end
    config network
        edit 1
            set prefix 10.10.10.0 255.255.255.0
        next
        edit 2
            set prefix 192.168.4.0 255.255.255.0
        next
    end
end
  - 
                                                                    
- 
                                                            
