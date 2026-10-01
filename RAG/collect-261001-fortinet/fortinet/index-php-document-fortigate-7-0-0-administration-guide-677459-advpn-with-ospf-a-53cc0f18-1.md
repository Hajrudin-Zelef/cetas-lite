---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18-1
title: "index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18.md
source_anchor: ""
source_lines: [1, 236]
sha256: f56dd475a16a9193399027565ed8cc346f2470b7ccd02668eb15b9bb359e5504
---

# index-php-document-fortigate-7-0-0-administration-guide-677459-advpn-with-ospf-a-53cc0f18

ADVPN with OSPF as the routing protocol
                                                
                                            This is a sample configuration of ADVPN with OSPF as the routing protocol. The following options must be enabled for this configuration:
- 
                                                    On the hub FortiGate, IPsec phase1-interface net-device enable must be run.
- 
                                                    OSPF must be used between the hub and spoke FortiGates.
To configure ADVPN with OSPF as the routing protocol using the CLI:
- 
                                                    Configure hub FortiGate's WAN, internal interface, and static route: config system interface
    edit "port9"
        set alias "WAN"
        set ip 22.1.1.1 255.255.255.0
    next
    edit "port10"
        set alias "Internal"
        set ip 172.16.101.1 255.255.255.0
    next
end   
config router static
    edit 1
        set gateway 22.1.1.2
        set device "port9"
    next  
end  		
- 
                                                    Configure the hub FortiGate: 
  - 
                                                            Configure the hub FortiGate IPsec phase1-interface and phase2-interface: config vpn ipsec phase1-interface
    edit "advpn-hub"
        set type dynamic
        set interface "port9"
        set peertype any
        set net-device enable
        set proposal aes128-sha256 aes256-sha256 3des-sha256 aes128-sha1 aes256-sha1 3des-sha1
        set add-route disable
        set dpd on-idle
        set auto-discovery-sender enable
        set psksecret sample
        set dpd-retryinterval 5
    next
end
config vpn ipsec phase2-interface
    edit "advpn-hub"
        set phase1name "advpn-hub"
        set proposal aes128-sha1 aes256-sha1 3des-sha1 aes128-sha256 aes256-sha256 3des-sha256
    next
endWhen net-device is disabled, a tunnel ID is generated for each dynamic tunnel. This ID, in the form of an IP address, is used as the gateway in the route entry to that tunnel. Thetunnel-search option is removed in FortiOS 7.0.0 and later.
  - 
                                                            Configure the hub FortiGate firewall policy: config firewall policy
    edit 1
        set name "spoke2hub"
        set srcintf "advpn-hub"
        set dstintf "port10"
        set srcaddr "all"
        set dstaddr "172.16.101.0"
        set action accept
        set schedule "always"
        set service "ALL"
    next
    edit 2
        set name "spoke2spoke"
        set srcintf "advpn-hub"
        set dstintf "advpn-hub"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
    next
end   
  - 
                                                            Configure the hub FortiGate's IPsec tunnel interface IP address: config system interface
    edit "advpn-hub1"
        set ip 10.10.10.254 255.255.255.255
        set remote-ip 10.10.10.253 255.255.255.0
    next
end
  - 
                                                            Configure the hub FortiGate's OSPF: config router ospf
    set router-id 1.1.1.1
    config area
        edit 0.0.0.0
        next
    end
    config network
        edit 1
            set prefix 10.10.10.0 255.255.255.0
        next
        edit 2
            set prefix 172.16.101.0 255.255.255.0
        next
    end
end
- 
                                                            
- 
                                                    Configure the spoke FortiGates: 
  - 
                                                            Configure the spoke FortiGates' WAN, internal interfaces, and static routes: 
    - 
                                                                    Configure Spoke1: config system interface
    edit "wan1"
        set alias "primary_WAN"
        set ip 15.1.1.2 255.255.255.0
    next
    edit "wan2"
        set alias "secondary_WAN"
        set ip 12.1.1.2 255.255.255.0
    next
    edit "internal"
        set ip 10.1.100.1 255.255.255.0
    next
end 
config router static
    edit 1
        set gateway 12.1.1.1
        set device "wan2"
        set distance 15        
    next
    edit 2
        set gateway 15.1.1.1
        set device "wan1"
    next
end   
    - 
                                                                    Configure the Spoke2: config system interface
    edit "wan1"
        set alias "primary_WAN"
        set ip 13.1.1.2 255.255.255.0
    next
    edit "wan2"
        set alias "secondary_WAN"
        set ip 17.1.1.2 255.255.255.0
    next
    edit "internal"
        set ip 192.168.4.1 255.255.255.0
    next
end 
config router static
    edit 1
        set gateway 17.1.1.1
        set device "wan2"
        set distance 15        
    next
    edit 2
        set gateway 13.1.1.1
        set device "wan1"
    next
end     
  - 
                                                                    
  - 
                                                            Configure the spoke FortiGates' IPsec phase1-interface and phase2-interface: 
    - 
                                                                    Configure Spoke1: config vpn ipsec phase1-interface
    edit "spoke1"
        set interface "wan1"
        set peertype any
        set net-device enable
        set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1
        set add-route disable
        set dpd on-idle
        set auto-discovery-receiver enable
        set remote-gw 22.1.1.1
        set psksecret sample 
        set dpd-retryinterval 5
    next
    edit "spoke1_backup"
        set interface "wan2"
        set peertype any
        set net-device enable
        set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1
        set add-route disable
        set dpd on-idle
        set auto-discovery-receiver enable
        set remote-gw 22.1.1.1
        set monitor "spoke1"
        set psksecret sample
        set dpd-retryinterval 5
    next    
end
config vpn ipsec phase2-interface
    edit "spoke1"
        set phase1name "spoke1"
        set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305
        set auto-negotiate enable
    next
    edit "spoke1_backup"
        set phase1name "spoke1_backup"
        set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305
        set auto-negotiate enable
    next  
end
    - 
                                                                    Configure Spoke2: config vpn ipsec phase1-interface
    edit "spoke2"
        set interface "wan1"
        set peertype any
        set net-device enable
        set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1
        set add-route disable
        set dpd on-idle
        set auto-discovery-receiver enable
        set remote-gw 22.1.1.1
        set psksecret sample 
        set dpd-retryinterval 5
    next
    edit "spoke2_backup"
        set interface "wan2"
        set peertype any
        set net-device enable
        set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1
        set add-route disable
        set dpd on-idle
        set auto-discovery-receiver enable
        set remote-gw 22.1.1.1
        set monitor "spoke2"
        set psksecret sample
        set dpd-retryinterval 5
    next    
end
config vpn ipsec phase2-interface
    edit "spoke2"
        set phase1name "spoke2"
        set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305
        set auto-negotiate enable
    next
    edit "spoke2_backup"
        set phase1name "spoke2_backup"
        set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305
        set auto-negotiate enable
    next  
end
  - 
                                                                    
