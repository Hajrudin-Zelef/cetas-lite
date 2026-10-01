---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-848299-example-cli-configuration-40bf8d20-2
title: "document-fortigate-7-4-7-administration-guide-848299-example-cli-configuration-40bf8d20"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-848299-example-cli-configuration-40bf8d20.md
source_anchor: ""
source_lines: [157, 233]
sha256: af1fe567c2171806bc6d92b30f181bf93e5059081cc3c0c38cc4995905cab7e5
---

# document-fortigate-7-4-7-administration-guide-848299-example-cli-configuration-40bf8d20

To configure the LAN extension interface and firewall policy on the FortiGate Controller:
- 
                                                    After the IPsec tunnel is setup and the VXLAN is created over the tunnel, the LAN extension interface is automatically created on the Controller: config system interface
    edit "FGT60E0000000001"
        set vdom "root"
        set ip 192.168.0.254 255.255.255.0
        set allowaccess ping ssh
        set type lan-extension
        set role lan
        set snmp-index 27
        set ip-managed-by-fortiipam enable
        set interface "fg-ipsec-XdSpij"
    next
endDevices on the remote LAN network will use this IP address as their gateway.
- 
                                                    Observe that with IPAM enabled on the Controller that the DHCP server settings have been automatically configured: config system dhcp server
    edit 3
        set dns-service default
        set default-gateway 9.9.9.99
        set netmask 255.255.255.0
        set interface "FGT60E0000000001"
        config ip-range
            edit 1
                set start-ip 9.9.9.100
                set end-ip 9.9.9.254
            next
        end
        set dhcp-settings-from-fortiipam enable
        config exclude-range
            edit 1
                set start-ip 9.9.9.254
                set end-ip 9.9.9.254
            next
        end
    next
end
- 
                                                    Configure the firewall policy to allow traffic from the LAN extension interface to the WAN interface (port1): config firewall policy
    edit "2"
        set name "lan-ext"
        set srcintf "FGT60E0000000001"
        set dstintf "port1"
        set action accept
        set srcaddr "all"
        set dstaddr "all"
        set schedule "always"
        set service "ALL"
        set nat enable
    next
endOptionally, security profiles and other settings can be configured. The policy allows remote LAN clients to access the internet through the backhaul channel. Clients in the remote LAN behind the Connector receive an IP address over DHCP and access the internet securely through the Controller.
To verify the FortiGate LAN extension configuration:
- 
                                                    Verify the IPsec tunnels' phase 1 and phase 2 negotiations on the Controller and Connector: # diagnose ike vpn gateway list # diagnose vpn tunnel list
- 
                                                    Verify the VXLAN tunnel forwarding database list on the Controller and Connector: # diagnose sys vxlan fdb list
- 
                                                    Verify the DHCP server lease list on the Controller: # execute dhcp lease-list
- 
                                                    Verify the LAN extension session information on the Controller: Controller-FGT # get extender session-info Total 1 WS sessions, 0 AS sessions: fg connector sessions: FGT60E0000000001 : 1.1.1.10:5246 (dport 65535)lan-extension, running, install, data-enable, refcnt 6, miss_echos -1, up-time 1554 secs, change 1 extender sessions: In this example, the Connector is in a working state.
- 
                                                    Verify the LAN extension status on the Connector:* Connector-FGT (lan-ext) # get extender lanextension-vdom-status
Control-Channel:
        controller ip: 1.1.1.0
        controller port: 5246
        controller name: FG5H1E0000000001
        missed echo: 0
        up time(seconds): 29483
        status: EXTWS_RUN
Data-Channel:
uplink [0]: port1
        IPsec tunnel ul-port1
        VxLAN interface vx-port1
uplink [1]: port2
        IPsec tunnel ul-port2
        VxLAN interface vx-port2
downlink [0]: lan
In this example, the Connector is in a working state.
