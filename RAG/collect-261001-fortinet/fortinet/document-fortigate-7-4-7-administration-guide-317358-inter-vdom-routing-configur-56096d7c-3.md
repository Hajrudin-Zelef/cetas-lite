---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c-3
title: "document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c.md
source_anchor: ""
source_lines: [289, 456]
sha256: 01d459857428e54bc8977bc9c7bec6eb9968fe560e4e13c06772b8ba9efbabf1
---

# document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c

                set subnet 192.168.20.10 255.255.255.255
            next
        end
    next
end
- 
                                                    Add the virtual IP address to VDOM-B: config vdom
    edit VDOM-B
        config firewall vip 
            edit FTP-server-VIP
                set extip 172.20.10.2
                set extintf wan2
                set mappedip 192.168.20.10
            next
        end 
    next
end
- 
                                                    Add a default route to VDOM-B: config vdom
    edit VDOM-B
        config router static
            edit 0
                set gateway 172.20.10.254
                set device wan2
            next
        end
    next
end
- 
                                                    Add the firewall policy to VDOM-B: config vdom
    edit VDOM-B
        config firewall policy
            edit 1
                set name "Access-server"
                set srcintf "wan2"
                set dstintf "port2"
                set srcaddr "all"
                set dstaddr "FTP-server-VIP"
                set action accept
                set schedule "always"
                set service "FTP"
                set nat enable
            next
        end
    next
end
To configure the VDOM link:
- 
                                                    Configure the VDOM link: config global
    config system vdom-link
        edit "VDOM-link"
        next
    end    
    config system interface
        edit VDOM-link0
            set vdom VDOM-A
            set ip 11.11.11.1 255.255.255.252
            set allowaccess https ping ssh
            set description "VDOM-A side of the VDOM link"
        next
        edit VDOM-link1
            set vdom VDOM-B
            set ip 11.11.11.2 255.255.255.252
            set allowaccess https ping ssh
            set description "VDOM-A side of the VDOM link"
        next
    end
end
- 
                                                    Configure the firewall addresses on VDOM-A: config vdom
    edit VDOM-A
        config firewall address
            edit "FTP-server"
                set associated-interface "VDOM-link0"
                set allow-routing enable
                set subnet 192.168.20.10 255.255.255.255
            next
        end
    next
end
- 
                                                    Add the firewall policy to VDOM-B: config vdom
    edit VDOM-B
        config firewall policy
            edit 1
                set name "Access-server"
                set srcintf "wan2"
                set dstintf "port2"
                set srcaddr "all"
                set dstaddr "FTP-server-VIP"
                set action accept
                set schedule "always"
                set service "FTP"
                set nat enable
            next
        end
    next
end
- 
                                                    Add the static route on VDOM-A: config vdom
    edit VDOM-A
        config router static 
            edit 0
                set device VDOM-link0
                set dstaddr FTP-server
                set gateway 11.11.11.2 
            next
        end
    next
end
- 
                                                    Configure the firewall addresses on VDOM-B: config vdom
    edit VDOM-B
        config firewall address
            edit internal-network
                set associated-interface VDOM-link1
                set allow-routing enable
                set subnet 192.168.10.0 255.255.255.0
            next
        end
    next
end
- 
                                                    Add the static route on VDOM-B: config vdom
    edit VDOM-B
        config router static 
            edit 0
                set device VDOM-link1
                set dstaddr internal-network
                set gateway 11.11.11.1
            next
        end
    next
end
- 
                                                    Add the security policy on VDOM-A: config vdom
    edit VDOM-A
        config firewall policy 
            edit 0
                set name Access-FTP-server
                set srcintf port1
                set dstintf VDOM-link0
                set srcaddr internal-network
                set dstaddr FTP-server
                set action accept
                set schedule always
                set service FTP
            next
        end
    next
end
- 
                                                    Add the firewall policy on VDOM-B: config vdom
    edit VDOM-B
        config firewall policy 
            edit 0
                set name Internal-server-access
                set srcintf VDOM-link1
                set dstintf port2
                set srcaddr internal-network
                set dstaddr FTP-server
                set action accept
                set schedule always
                set service FTP
            next
        end
    next
end
