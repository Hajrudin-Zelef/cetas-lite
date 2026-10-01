---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-179703-ssl-vpn-tunnel-mode-host-ch-1fd59a6e
title: "document-fortigate-7-4-8-administration-guide-179703-ssl-vpn-tunnel-mode-host-ch-1fd59a6e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-179703-ssl-vpn-tunnel-mode-host-ch-1fd59a6e.md
source_anchor: ""
source_lines: [1, 124]
sha256: e81e76c7e4b945c89a6d9c79e05aa1ae69c41bac7f33ac757c7c3f41fceaf13d
---

# document-fortigate-7-4-8-administration-guide-179703-ssl-vpn-tunnel-mode-host-ch-1fd59a6e

SSL VPN tunnel mode host check
SSL VPN tunnel mode host check
This is a sample configuration of remote users accessing the corporate network through an SSL VPN by tunnel mode using FortiClient with AV host check.
Sample topology
Sample configuration
WAN interface is the interface connected to ISP. This example shows static mode. You can also use DHCP or PPPoE mode. The SSL VPN connection is established over the WAN interface.
|  | The split tunneling routing address cannot explicitly use an FQDN or an address group that includes an FQDN. To use an FQDN, leave the routing address blank and apply the FQDN as the destination address of the firewall policy. | 
To configure SSL VPN using the GUI:
- Configure the interface and firewall address. The port1 interface connects to the internal network.
  - Go to Network > Interfaces and edit the wan1 interface.
  - Set IP/Network Mask to 172.20.120.123/255.255.255.0.
  - Edit port1 interface and set IP/Network Mask to 192.168.1.99/255.255.255.0.
  - Click OK.
  - Go to Policy & Objects > Address and create an address for internet subnet 192.168.1.0.
- Configure user and user group.
  - Go to User & Authentication > User Definition to create a local user sslvpnuser1.
  - Go to User & Authentication > User Groups to create a group sslvpngroup with the member sslvpnuser1.
- Configure SSL VPN web portal.
  - Go to VPN > SSL-VPN Portals to create a tunnel mode only portal my-split-tunnel-portal.
  - Enable Tunnel Mode and select one of the Split tunneling settings. See Split tunneling settings for more information.
  - Select Routing Address Override.
  - Select Source IP Pools for users to acquire an IP address when connecting to the portal. There is always a default pool available if you do not create your own.
- Configure SSL VPN settings.
  - Go to VPN > SSL-VPN Settings.
  - For Listen on Interface(s), select wan1.
  - Set Listen on Port to 10443.
  - Choose a certificate for Server Certificate.
  - In Authentication/Portal Mapping All Other Users/Groups, set the Portal to tunnel-access.
  - Create new Authentication/Portal Mapping for group sslvpngroup mapping portal my-split-tunnel-portal.
 It is HIGHLY recommended that you acquire a signed certificate for your installation. Please review the SSL VPN best practices and learn how to Procuring and importing a signed SSL certificate.
- Configure SSL VPN firewall policy.
  - Go to Policy & Objects > Firewall Policy.
  - Fill in the firewall policy name. In this example, sslvpn tunnel access with av check.
  - Incoming interface must be SSL-VPN tunnel interface(ssl.root).
  - Choose an Outgoing Interface. In this example, port1.
  - Set the Source to all and group to sslvpngroup.
  - In this example, the Destination is all.
  - Set Schedule to always, Service to ALL, and Action to Accept.
  - Click OK.
- Use CLI to configure SSL VPN web portal to enable the host to check for compliant antivirus software on the user's computer.config vpn ssl web portal
    edit my-split-tunnel-access
        set host-check av
    next
end
To configure SSL VPN using the CLI:
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
        set member "vpnuser1"
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
    end
end
- Configure one SSL VPN firewall policy to allow remote user to access the internal network. Traffic is dropped from internal to remote client.config firewall policy 
    edit 1
        set name "sslvpn web mode access"
        set srcintf "ssl.root"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "192.168.1.0"
        set groups "sslvpngroup"
        set action accept
        set schedule "always"
        set service "ALL"
    next
end
- Configure SSL VPN web portal to enable the host to check for compliant antivirus software on the user's computer:config vpn ssl web portal
    edit my-split-tunnel-access
        set host-check av
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
- Use the credentials you've set up to connect to the SSL VPN tunnel.If the user's computer has antivirus software, a connection is established; otherwise FortiClient shows a compliance warning.
- After connection, traffic to 192.168.1.0 goes through the tunnel. Other traffic goes through local gateway.
- On the FortiGate, go to VPN > Monitor > SSL-VPN Monitor to verify the list of SSL users.
- On the FortiGate, go to Log & Report > Forward Traffic and view the details for the SSL entry.
