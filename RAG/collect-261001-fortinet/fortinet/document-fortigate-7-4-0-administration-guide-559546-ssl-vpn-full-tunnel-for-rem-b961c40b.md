---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-559546-ssl-vpn-full-tunnel-for-rem-b961c40b
title: "document-fortigate-7-4-0-administration-guide-559546-ssl-vpn-full-tunnel-for-rem-b961c40b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-559546-ssl-vpn-full-tunnel-for-rem-b961c40b.md
source_anchor: ""
source_lines: [1, 118]
sha256: 21f7e1103ac7c0b080d48f428d7877dbd5b3659fd35b99011e4cd3d4f6ef65ed
---

# document-fortigate-7-4-0-administration-guide-559546-ssl-vpn-full-tunnel-for-rem-b961c40b

SSL VPN full tunnel for remote user
SSL VPN full tunnel for remote user
This is a sample configuration of remote users accessing the corporate network and internet through an SSL VPN by tunnel mode using FortiClient.
Sample topology
Sample configuration
WAN interface is the interface connected to ISP. This example shows static mode. You can also use DHCP or PPPoE mode. The SSL VPN connection is established over the WAN interface.
To configure SSL VPN using the GUI:
- Configure the interface and firewall address:
  - Go to Network > Interfaces and edit the wan1 interface.
  - Set IP/Network Mask to 172.20.120.123/255.255.255.0.
  - Edit port1 interface and set IP/Network Mask to 192.168.1.99/255.255.255.0.
  - Click OK.
- Configure user and user group:
  - Go to User & Authentication > User Definition to create a local user sslvpnuser1.
  - Go to User & Authentication > User Groups to create a group sslvpngroup with the member sslvpnuser1.
- Configure SSL VPN web portal:
  - Go to VPN > SSL-VPN Portals to create a tunnel mode only portal my-full-tunnel-portal.
  - Disable Split Tunneling.
- Configure SSL VPN settings:
  - Go to VPN > SSL-VPN Settings.
  - For Listen on Interface(s), select wan1.
  - Set Listen on Port to 10443.
  - Choose a certificate for Server Certificate. The default is Fortinet_Factory.
  - In Authentication/Portal Mapping All Other Users/Groups, set the Portal to tunnel-access.
  - Create new Authentication/Portal Mapping for group sslvpngroup mapping portal my-full-tunnel-portal.
- Configure SSL VPN firewall policies to allow remote user to access the internal network:
  - Go to Policy & Objects > Firewall Policy and click Create New.
  - Set Name to sslvpn tunnel mode access.
  - Set Incoming Interface to SSL-VPN tunnel interface(ssl.root).
  - Set Outgoing Interface to port1.
  - Set the Source Address to all and User to sslvpngroup.
  - Set Destination to all, Schedule to always, Service to ALL, and Action to Accept.
  - Click OK.
  - Click Create New.
  - Set Name to sslvpn tunnel mode outgoing.
  - Configure the same settings as the previous policy, except set Outgoing Interface to wan1.
  - Click OK.
To configure SSL VPN using the CLI:
- Configure the interface and firewall address.config system interface 
    edit "wan1"
        set vdom "root"
        set ip 172.20.120.123 255.255.255.0
    next
end
- Configure the internal interface and protected subnet, then connect the port1 interface to the internal network.config system interface
    edit "port1"
        set vdom "root"
        set ip 192.168.1.99 255.255.255.0
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
- Configure SSL VPN web portal and predefine RDP bookmark for windows server.config vpn ssl web portal
    edit "my-full-tunnel-portal"
        set tunnel-mode enable
        set split-tunneling disable
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
            set portal "my-full-tunnel-portal"
        next        
    end
end
- Configure SSL VPN firewall policies to allow remote user to access the internal network. Traffic is dropped from internal to remote client.config firewall policy 
    edit 1
        set name "sslvpn tunnel mode access"
        set srcintf "ssl.root"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "all"
        set groups "sslvpngroup"
        set action accept
        set schedule "always"
        set service "ALL"
    next
    edit 2
        set name "sslvpn tunnel mode outgoing"
        set srcintf "ssl.root"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set groups "sslvpngroup"
        set action accept
        set schedule "always"
        set service "ALL"
    next
end
To see the results:
- Download FortiClient from www.forticlient.com.
- Open the FortiClient Console and go to Remote Access.
- Add a new connection:
  - Set VPN Type to SSL VPN.
  - Set Remote Gateway to the IP of the listening FortiGate interface, in this example, 172.20.120.123.
- Select Customize Port and set it to 10443.
- Save your settings.
- Use the credentials you've set up to connect to the SSL VPN tunnel.
- After connection, all traffic except the local subnet will go through the tunnel FGT.
- Go to VPN > Monitor > SSL-VPN Monitor to verify the list of SSL users.
- On the FortiGate, go to Log & Report > Forward Traffic and view the details for the SSL entry.
