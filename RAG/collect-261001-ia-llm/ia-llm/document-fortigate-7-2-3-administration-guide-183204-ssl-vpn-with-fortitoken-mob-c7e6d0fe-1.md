---
id: collect-261001-ia-llm/ia-llm/document-fortigate-7-2-3-administration-guide-183204-ssl-vpn-with-fortitoken-mob-c7e6d0fe-1
title: "SSL VPN with FortiToken mobile push authentication"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-ia-llm/document-fortigate-7-2-3-administration-guide-183204-ssl-vpn-with-fortitoken-mob-c7e6d0fe.md
source_anchor: ""
source_lines: [1, 180]
sha256: f0857e10538e054141478b881abc661a60bef5a09243a4836a3d6b1f3091ba15
---

# SSL VPN with FortiToken mobile push authentication

This is a sample configuration of SSL VPN that uses FortiToken mobile push two-factor authentication. If you enable push notifications, users can accept or deny the authentication request.

## Sample topology


## Sample configuration

WAN interface is the interface connected to ISP. This example shows static mode. You can also use DHCP or PPPoE mode. The SSL VPN connection is established over the WAN interface.

###### To configure SSL VPN using the GUI:

1. Configure the interface and firewall address. The port1 interface connects to the internal network.
  1. Go to *Network > Interfaces* and edit the*wan1* interface.
  2. Set *IP/Network Mask* to*172.20.120.123/255.255.255.0* .
  3. Edit *port1* interface and set*IP/Network Mask* to*192.168.1.99/255.255.255.0* .
  4. Click *OK* .
  5. Go to *Policy & Objects > Address* and create an address for internet subnet*192.168.1.0* .
2. Go to 
3. Register FortiGate for FortiCare Support:To add or download a mobile token on FortiGate, FortiGate must be registered for FortiCare Support. If your FortiGate is registered, skip this step. 
  1. Go to *Dashboard > Licenses* .
  2. Hover the pointer on *FortiCare Support* to check if FortiCare registered. If not, click it and select*Register* .
4. Go to 
5. Add FortiToken mobile to FortiGate:If your FortiGate has FortiToken installed, skip this step. 
  1. Go to *User & Authentication > FortiTokens* and click*Create New* .
  2. Select *Mobile Token* and type in*Activation Code* .
  3. Every FortiGate has two free mobile tokens. Go to *User & Authentication > FortiTokens* and click*Import Free Trial Tokens* .
6. Go to 
7. Enable FortiToken mobile push:To use FTM-push authentication, use CLI to enable FTM-Push on the FortiGate. 
  1. Ensure `server-ip` is reachable from the Internet and enter the following CLI commands:```
config system ftm-push
    set server-ip 172.20.120.123
    set status enable
end
```
  2. Go to *Network > Interfaces* .
  3. Edit the *wan1* interface.
  4. Under *Administrative Access > IPv4* , select*FTM* .
  5. Click *OK* .
8. Ensure 
9. Configure user and user group:
  1. Go to *User & Authentication > User Definition* to create a local user*sslvpnuser1* .
  2. Enter the user's *Email Address* .
  3. Enable *Two-factor Authentication* and select one mobile*Token* from the list,
  4. Enable *Send Activation Code* and select*Email* .
  5. Click *Next* and click*Submit* .
  6. Go to *User & Authentication > User Groups* to create a group*sslvpngroup* with the member*sslvpnuser1* .
10. Go to 
11. Activate the mobile token:
  1. When the user *sslvpnuser1* is created, an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.
12. When the user 
13. Configure SSL VPN web portal:
  1. Go to *VPN > SSL-VPN Portals* to edit the*full-access* portal.This portal supports both web and tunnel mode.
  2. Disable *Enable Split Tunneling* so that all SSL VPN traffic goes through the FortiGate.
14. Go to 
15. Configure SSL VPN settings:
  1. Go to *VPN > SSL-VPN Settings* .
  2. Select the *Listen on Interface(s)* , in this example,*wan1* .
  3. Set *Listen on Port* to*10443* .
  4. Set *Server Certificate* to the authentication certificate.
  5. Under *Authentication/Portal Mapping* , set default Portal*web-access* for*All Other Users/Groups* .
  6. Create new *Authentication/Portal Mapping* for group*sslvpngroup* mapping portal*full-access* .
16. Go to 
17. Configure SSL VPN firewall policy:
  1. Go to *Policy & Objects > Firewall Policy* .
  2. Fill in the firewall policy name. In this example, *sslvpn certificate auth* .
  3. Incoming interface must be *SSL-VPN tunnel interface(ssl.root)* .
  4. Set the *Source Address* to*all* and*Source User* to*sslvpngroup* .
  5. Set the *Outgoing Interface* to the local network interface so that the remote user can access the internal network. In this example,*port1* .
  6. Set *Destination Address* to the internal protected subnet*192.168.1.0* .
  7. Set *Schedule* to*always* ,*Service* to*ALL* , and*Action* to*Accept* .
  8. Enable *NAT* .
  9. Configure any remaining firewall and security options as desired.
  10. Click *OK* .
18. Go to 

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
3. Register FortiGate for FortiCare Support.To add or download a mobile token on FortiGate, FortiGate must be registered for FortiCare Support. If your FortiGate is registered, skip this step. diagnose forticare direct-registration product-registration -a "your account@xxx.com" -p "your password" -T "Your Country/Region" -R "Your Reseller" -e 1
4. Add FortiToken mobile to FortiGate:execute fortitoken-mobile import <your FTM code> If your FortiGate has FortiToken installed, skip this step. Every FortiGate has two free mobile Tokens. You can download the free token. execute fortitoken-mobile import 0000-0000-0000-0000-0000
5. Enable FortiToken mobile push:
  1. To use FTM-push authentication, ensure `server-ip` is reachable from the Internet and enable FTM-push in the FortiGate:```
config system ftm-push
    set server-ip 172.20.120.123
    set status enable
end
```
  2. Enable FTM service on WAN interface:```
config system interface 
    edit "wan1"
        append allowaccess ftm 
    next
end
```
6. To use FTM-push authentication, ensure 
7. Configure user and user group:```
config user local
    edit "sslvpnuser1"
        set type password
        set two-factor fortitoken
        set fortitoken <select mobile token for the option list>
        set email-to <user's email address>
        set passwd <user's password>
    next
end
config user group
    edit "sslvpngroup" 
        set member "sslvpnuser1"
    next 
end
```
8. Activate the mobile token.When the user *sslvpnuser1* is created, an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.
9. Configure SSL VPN web portal:```
config vpn ssl web portal
    edit "full-access"
        set tunnel-mode enable
        set web-mode enable
        set ip-pools "SSLVPN_TUNNEL_ADDR1"
        set split-tunneling disable
    next
end
```
10. Configure SSL VPN settings:```
config vpn ssl settings
    set servercert "server_certificate"
    set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1"
    set source-interface "wan1"
    set source-address "all"
    set default-portal "web-access"
    config authentication-rule
        edit 1
            set groups "sslvpngroup"
            set portal "full-access"
        next        
    end
end
```
11. Configure one SSL VPN firewall policy to allow remote user to access the internal network:```
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
        set nat enable
    next
end
```

###### To see the results of web portal:

