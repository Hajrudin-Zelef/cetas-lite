---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-12-administration-guide-307303-ssl-vpn-split-tunnel-for-r-8a1d9cbb
title: "document-fortigate-7-4-12-administration-guide-307303-ssl-vpn-split-tunnel-for-r-8a1d9cbb"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-12-administration-guide-307303-ssl-vpn-split-tunnel-for-r-8a1d9cbb.md
source_anchor: ""
source_lines: [1, 110]
sha256: 7f5ce229af1086d994de527a663cbd1b1420674cf196c8ef3fac46e8b4cd08fe
---

# document-fortigate-7-4-12-administration-guide-307303-ssl-vpn-split-tunnel-for-r-8a1d9cbb

SSL VPN split tunnel for remote user
SSL VPN split tunnel for remote user
This is a sample configuration of remote users accessing the corporate network and internet through an SSL VPN by tunnel mode using FortiClient but accessing the Internet without going through the SSL VPN tunnel.
Sample topology
Sample configuration
WAN interface is the interface connected to ISP. This example shows static mode. You can also use DHCP or PPPoE mode. The SSL VPN connection is established over the WAN interface. Ensure that SSL VPN feature visibility is enabled before starting the configuration.
|  | The split tunneling routing address cannot explicitly use an FQDN or an address group that includes an FQDN. To use an FQDN, leave the routing address blank and apply the FQDN as the destination address of the firewall policy. | 
To configure SSL VPN using the GUI:
- Enable SSL VPN feature visibility:
  - Go to System > Feature Visibility.
  - In the Core Features section, enable SSL-VPN.
  - Click Apply.
- Configure the interface and firewall address. The port1 interface connects to the internal network.
  - Go to Network > Interfaces and edit the wan1 interface.
  - Set IP/Network Mask to 172.20.120.123/255.255.255.0.
  - Edit port1 interface and set IP/Network Mask to 192.168.1.99/255.255.255.0.
  - Click OK.
  - Go to Policy & Objects > Address and create an address for internal subnet 192.168.1.0.
- Configure user and user group.
  - Go to User & Authentication > User Definition to create a local user sslvpnuser1.
  - Go to User & Authentication > User Groups to create a group sslvpngroup with the member sslvpnuser1.
- Configure SSL VPN web portal.
  - Go to VPN > SSL-VPN Portals to create a tunnel mode only portal my-split-tunnel-portal.
  - Enable Tunnel Mode and select one of the Split tunneling settings. See Split tunneling settings for more information.
  - Select Routing Address Override to define the destination network (usually the corporate network) that will be routed through the tunnel. Leave Routing Address Override undefined to use the destination in the respective firewall policies.
  - Select Source IP Pools for users to acquire an IP address when connecting to the portal. There is always a default pool available if you do not create your own.
- Configure SSL VPN settings.
  - Go to VPN > SSL-VPN Settings.
  - For Listen on Interface(s), select wan1.
  - Set Listen on Port to 10443.
  - Choose a certificate for Server Certificate. The default is Fortinet_Factory.
  - In Authentication/Portal Mapping All Other Users/Groups, set the Portal to tunnel-access.
  - Create new Authentication/Portal Mapping for group sslvpngroup mapping portal my-split-tunnel-portal.
- Configure SSL VPN firewall policy.
  - Go to Policy & Objects > Firewall Policy.
  - Fill in the firewall policy name. In this example, sslvpn split tunnel access.
  - Incoming interface must be SSL-VPN tunnel interface(ssl.root).
  - Choose an Outgoing Interface. In this example, port1.
  - Set the Source to all and group to sslvpngroup.
  - In this example, the Destination is the internal protected subnet 192.168.1.0.
  - Set Schedule to always, Service to ALL, and Action to Accept.
  - Click OK.
 Avoid setting all as the destination address in a firewall policy when the user or group associated with that policy is using a portal with Split tunneling enabled. Setting all as the destination address will cause portal to function as a full tunnel, potentially leading to misconfigurations and complicating troubleshooting efforts.
To configure SSL VPN using the CLI:
- 
                                                    Enable SSL VPN feature visibility: config system settings
    set gui-sslvpn enable
end
- Configure the interface and firewall address.config system interface 
    edit "wan1"
        set vdom "root"
        set ip 172.20.120.123 255.255.255.0
    next
end
- Configure internal interface and protected subnet, then connect the port1 interface to the internal network.config system interface
    edit "port1"
        set vdom "root"
        set ip 192.168.1.99 255.255.255.0
    next
endconfig firewall address
    edit "192.168.1.0"
        set subnet 192.168.1.0 255.255.255.0
    next
end
- Configure user and user group.config user local
    edit "sslvpnuser1" 
        set type password
        set passwd your-password
    next 
endconfig user group
    edit "sslvpngroup" 
        set member "sslvpnuser1"
    next 
end
- Configure SSL VPN web portal.config vpn ssl web portal
    edit "my-split-tunnel-portal"
        set tunnel-mode enable
        set split-tunneling  enable
        set split-tunneling-routing-address "192.168.1.0"
        set ip-pools "SSLVPN_TUNNEL_ADDR1"
    next
end
- Configure SSL VPN settings.config vpn ssl settings
    set servercert "Fortinet_Factory"
    set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1"
    set tunnel-ipv6-pools "SSLVPN_TUNNEL_IPv6_ADDR1"
    set source-interface "wan1"
    set source-address "all"
    set source-address6 "all"
    set default-portal "full-access"
    config authentication-rule
        edit 1
            set groups "sslvpngroup"
            set portal "my-split-tunnel-portal" 
        next        
    next
end
- Configure one SSL VPN firewall policy to allow the remote user to access the internal network. Traffic is dropped from internal to remote client.config firewall policy 
    edit 1
        set name "sslvpn split tunnel access"
        set srcintf "ssl.root"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "192.168.1.0"
        set groups "sslvpngroup"
        set action accept
        set schedule "always"
        set service "ALL"
    next
endAvoid setting all as the destination address in a firewall policy when the user or group associated with that policy is using a portal with Split tunneling enabled. Setting all as the destination address will cause portal to function as a full tunnel, potentially leading to misconfigurations and complicating troubleshooting efforts.
