---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c-2
title: "document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c.md
source_anchor: ""
source_lines: [123, 288]
sha256: 03305fc02d33557a0a8c1eeaaf620f9d9921c7deb0601d8e6844c4cd4e4c7495
---

# document-fortigate-7-4-7-administration-guide-317358-inter-vdom-routing-configur-56096d7c

                                                    Enter the following information: Destination Subnet IP address 0.0.0.0/0.0.0.0 Gateway 172.20.201.254 Interface wan2 Administrative Distance 10
- 
                                                    Click OK.
To add the firewall policy in the GUI:
- 
                                                    Go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Access-server Incoming Interface wan2 Outgoing Interface port2 Source all Destination FTP-server-VIP Schedule always Service FTP Action ACCEPT NAT enabled
- 
                                                    Click OK.
Configure the VDOM link
The VDOM link allows connections from VDOM-A to VDOM-B. The VDOM link interface configured in this step will be used for inter-VDOM routing.
This step requires you to connect to the global VDOM using a global administrator account.
To add the VDOM link in the GUI:
- 
                                                    In the Global VDOM, go to Network > Interfaces.
- 
                                                    Create New > VDOM link.
- 
                                                    Enter the following information: Name VDOM-link Interface 0 Virtual Domain VDOM-A IP/Netmask 11.11.11.1/255.255.255.252 Interface 1 Virtual Domain VDOM-B IP/Netmask 11.11.11.2/255.255.255.252
- 
                                                    Click OK.
Configure inter-VDOM routing
Inter-VDOM routing allows users on the internal network to route traffic to the FTP server through the FortiGate.
The configuration of inter-VDOM routing includes the following:
- 
                                                    Firewall addresses for the FTP server on VDOM-A and for the internal network on VDOM-B
- 
                                                    Inter-VDOM routing using static routes for the FTP server on VDOM-A and for the internal network on VDOM-B
- 
                                                    Policies allowing traffic using the VDOM link
The procedures described above require you to connect to both VDOM-A and VDOM-B, either using a global or per-VDOM administrator account.
To add the firewall address on VDOM-A in the GUI:
- 
                                                    In the VDOM-A VDOM, go to Policy & Objects > Addresses and select Address.
- 
                                                    Click Create new.
- 
                                                    Enter the following information: Name FTP-server Type Subnet IP/Netmask 192.168.20.10/32 Interface VDOM-link2 Static route configuration enabled
- 
                                                    Click OK.
To add the static route on VDOM-A in the GUI:
- 
                                                    Connect to VDOM-A.
- 
                                                    Go to Network > Static Routes and create a new route.
- 
                                                    Enter the following information: Destination Named Address Named Address FTP-server Gateway 11.11.11.2 Interface VDOM-link0
- 
                                                    Click OK.
To add the firewall address on VDOM-B in the GUI:
- 
                                                    In the VDOM-B VDOM, go to Policy & Objects > Addresses and select Address.
- 
                                                    Click Create new.
- 
                                                    Enter the following information: Name internal-network Type Subnet IP/Netmask 192.168.10.0/24 Interface VDOM-link1 Static route configuration enabled
- 
                                                    Click OK.
To add the static route on VDOM-B in the GUI:
- 
                                                    In the VDOM-B VDOM, go to Network > Static Routes and create a new route.
- 
                                                    Enter the following information: Destination Named Address Named Address internal-network Gateway 11.11.11.1 Interface VDOM-link1
- 
                                                    Click OK.
Configure firewall policies using the VDOM link
Firewall policies using the VDOM link allows users on the internal network to access the FTP server through the FortiGate.
Configuring policies allowing traffic using the VDOM link require you to connect to both VDOM-A and VDOM-B, respectively, either using a global or per-VDOM administrator account.
To add the firewall policy on VDOM-A in the GUI:
- 
                                                    In the VDOM-A VDOM, go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Access-FTP-server Incoming Interface port1 Outgoing Interface VDOM-link0 Source internal-network Destination FTP-server Schedule always Service FTP Action ACCEPT NAT disabled
- 
                                                    Click OK.
To add the firewall policy on VDOM-B in the GUI:
- 
                                                    In the VDOM-B VDOM, go to Policy & Objects > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following information: Name Internal-server-access Incoming Interface VDOM-link1 Outgoing Interface port2 Source internal-network Destination FTP-server Schedule always Service FTP Action ACCEPT NAT disabled
- 
                                                    Click OK.
Configuration with the CLI
The example can also be configured in the CLI.
To configure the two VDOMs:
- 
                                                    Enable multi-VDOM mode: config system global
    set vdom-mode multi-vdom
endYou will be logged out of the device when VDOM mode is enabled.
- 
                                                    Create the VDOMs: config vdom
    edit VDOM-A
    next
    edit VDOM-B
    next
end
- 
                                                    Assign interfaces to the VDOMs: config global
    config system interface
        edit port1
            set vdom VDOM-A
        next
        edit port2
            set vdom VDOM-B
        next
        edit wan1
            set vdom VDOM-A
        next
        edit wan2
            set vdom VDOM-B
        next
    end
end
- 
                                                    Add the firewall addresses to VDOM-A: config vdom
    edit VDOM-A
        config firewall address
            edit internal-network
                set associated-interface port1
                set subnet 192.168.10.0 255.255.255.0
            next
        end
    next
end
- 
                                                    Add a default route to VDOM-A: config vdom
    edit VDOM-A
        config router static
            edit 0
                set gateway 172.20.201.254
                set device wan1
            next
        end
    next
end
- 
                                                    Add the firewall policy to VDOM-A: config vdom
    edit VDOM-A
        config firewall policy
            edit 1
                set name "VDOM-A-Internet"
                set srcintf "port1"
                set dstintf "wan1"
                set srcaddr "internal-network"
                set dstaddr "all"
                set action accept
                set schedule "always"
                set service "ALL"
                set nat enable
            next
        end
    next
end
- 
                                                    Add the firewall addresses to VDOM-B: config vdom
    edit VDOM-B
        config firewall address
            edit FTP-server
                set associated-interface port2
