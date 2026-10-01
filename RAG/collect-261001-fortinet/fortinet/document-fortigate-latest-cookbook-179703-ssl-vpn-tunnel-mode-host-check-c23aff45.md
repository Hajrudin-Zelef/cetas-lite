---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-cookbook-179703-ssl-vpn-tunnel-mode-host-check-c23aff45
title: "SSL VPN tunnel mode host check"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-cookbook-179703-ssl-vpn-tunnel-mode-host-check-c23aff45.md
source_anchor: ""
source_lines: [1, 163]
sha256: 9ad5302f7116be45193846d1233f52ce9b40770836093dc5803249acd1679b30
---

# SSL VPN tunnel mode host check

# SSL VPN tunnel mode host check

This is a sample configuration of remote users accessing the corporate network through an SSL VPN by tunnel mode using FortiClient with AV host check.

## Sample topology


## Sample configuration

WAN interface is the interface connected to ISP. This example shows static mode. You can also use DHCP or PPPoE mode. The SSL VPN connection is established over the WAN interface.

|  | The split tunneling routing address cannot use an FQDN or an address group that includes an FQDN. | 

###### To configure SSL VPN using the GUI:

1. Configure the interface and firewall address. The port1 interface connects to the internal network.
  1. Go to *Network > Interfaces* and edit the*wan1* interface.
  2. Set *IP/Network Mask* to*172.20.120.123/255.255.255.0* .
  3. Edit *port1* interface and set*IP/Network Mask* to*192.168.1.99/255.255.255.0* .
  4. Click *OK* .
  5. Go to *Policy & Objects > Address* and create an address for internet subnet*192.168.1.0* .
2. Go to 
3. Configure user and user group.
  1. Go to *User & Device > User Definition* to create a local user*sslvpnuser1* .
  2. Go to *User & Device > User Groups* to create a group*sslvpngroup* with the member*sslvpnuser1* .
4. Go to 
5. Configure SSL VPN web portal.
  1. Go to *VPN > SSL-VPN Portals* to create a tunnel mode only portal*my-split-tunnel-portal* .
  2. Enable *Tunnel Mode* and*Enable Split Tunneling* .
  3. Select *Routing Address* .
6. Go to 
7. Configure SSL VPN settings.
  1. Go to *VPN > SSL-VPN Settings* .
  2. For *Listen on Interface(s)* , select*wan1* .
  3. Set *Listen on Port* to*10443* .
  4. Choose a certificate for *Server Certificate* . The default is*Fortinet_Factory* .
  5. In *Authentication/Portal Mapping**All Other Users/Groups* , set the*Portal* to*tunnel-access* .
  6. Create new *Authentication/Portal Mapping* for group*sslvpngroup* mapping portal*my-split-tunnel-portal* .
8. Go to 
9. Configure SSL VPN firewall policy.
  1. Go to *Policy & Objects > IPv4 Policy* .
  2. Fill in the firewall policy name. In this example, *sslvpn tunnel access with av check* .
  3. Incoming interface must be *SSL-VPN tunnel interface(ssl.root)* .
  4. Choose an *Outgoing Interface* . In this example,*port1* .
  5. Set the *Source* to*all* and group to*sslvpngroup* .
  6. In this example, the *Destination* is*all* .
  7. Set *Schedule* to*always* ,*Service* to*ALL* , and*Action* to*Accept* .
  8. Click *OK* .
10. Go to 
11. Use CLI to configure SSL VPN web portal to enable the host to check for compliant antivirus software on the user's computer.```
config vpn ssl web portal
    edit my-split-tunnel-access
        set host-check av
    next
end
```

###### To configure SSL VPN using the CLI:

1. Configure the interface and firewall address.```
config system interface 
    edit "wan1"
        set vdom "root"
        set ip 172.20.120.123 255.255.255.0
    next
end
```
2. Configure internal interface and protected subnet, then connect the port1 interface to the internal network.```
config system interface
    edit "port1"
        set vdom "root"
        set ip 192.168.1.99 255.255.255.0
    next
end
```
```
config firewall address
    edit "192.168.1.0"
        set subnet 192.168.1.0 255.255.255.0
    next
end
```
3. Configure user and user group.```
config user local
    edit "sslvpnuser1" 
        set type password
        set passwd your-password
    next 
end
```
```
config user group
    edit "sslvpngroup" 
        set member "vpnuser1"
    next 
end
```
4. Configure SSL VPN web portal.```
config vpn ssl web portal
    edit "my-split-tunnel-portal"
        set tunnel-mode enable
        set split-tunneling  enable
        set split-tunneling-routing-address "192.168.1.0"
        set ip-pools "SSLVPN_TUNNEL_ADDR1"
    next
end
```
5. Configure SSL VPN settings.```
config vpn ssl settings
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
```
6. Configure one SSL VPN firewall policy to allow remote user to access the internal network. Traffic is dropped from internal to remote client.```
config firewall policy 
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
```
7. Configure SSL VPN web portal to enable the host to check for compliant antivirus software on the user's computer:```
config vpn ssl web portal
    edit my-split-tunnel-access
        set host-check av
    next
end
```

###### To see the results:

1. Download FortiClient from www.forticlient.com.
2. Open the FortiClient Console and go to *Remote Access* .
3. Add a new connection:
  - Set *VPN Type* to*SSL VPN* .
  - Set *Remote Gateway* to the IP of the listening FortiGate interface, in this example,*172.20.120.123* .
4. Set 
5. Select *Customize Port* and set it to*10443* .
6. Save your settings.
7. Use the credentials you've set up to connect to the SSL VPN tunnel.If the user's computer has antivirus software, a connection is established; otherwise FortiClient shows a compliance warning.
8. After connection, traffic to *192.168.1.0* goes through the tunnel. Other traffic goes through local gateway.
9. On the FortiGate, go to *VPN > Monitor > SSL-VPN Monitor* to verify the list of SSL users.
10. On the FortiGate, go to *Log & Report > Forward Traffic* and view the details for the SSL entry.
