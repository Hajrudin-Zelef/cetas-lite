---
id: collect-261001-fortinet/fortinet/document-fortigate-7-0-9-administration-guide-207191-ssl-vpn-with-radius-and-for-82879e55-1
title: "SSL VPN with RADIUS and FortiToken mobile push on FortiAuthenticator"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-0-9-administration-guide-207191-ssl-vpn-with-radius-and-for-82879e55.md
source_anchor: ""
source_lines: [1, 169]
sha256: 658a5cd4f4f82fe76ae8a4413dcdcd029917157f9055bb81ec93f464bcc2a4bc
---

# SSL VPN with RADIUS and FortiToken mobile push on FortiAuthenticator

This is a sample configuration of SSL VPN that uses FortiAuthenticator as a RADIUS authentication server and FortiToken mobile push two-factor authentication. If you enable push notifications, users can accept or deny the authentication request.

## Sample topology


## Sample configuration

WAN interface is the interface connected to ISP. This example shows static mode. You can also use DHCP or PPPoE mode. The SSL VPN connection is established over the WAN interface.

###### To configure FortiAuthenticator using the GUI:

1. On the FortiAuthenticator, go to *System > Administration > System Access* and configure a*Public IP/FQDN for FortiToken Mobile* . If the FortiAuthenticator is behind a firewall, the public IP/FQDN will be an IP/port forwarding rule directed to one of the FortiAuthenticator interfaces. The interface that receives the approve/deny FTM push responses must have the*FortiToken Mobile API* service enabled.
2. Add a FortiToken mobile license on the FortiAuthenticator:
  1. Go to *Authentication > User Management > FortiTokens* .
  2. Click *Create New* .
  3. Set *Token type* to*FortiToken Mobile* and enter the*FortiToken Activation codes* .
3. Go to 
4. Create the RADIUS client (FortiGate) on the FortiAuthenticator:
  1. Go to *Authentication > RADIUS Service > Clients* to add the FortiGate as a RADIUS client*OfficeServer* ).
  2. Enter the FortiGate IP address and set a *Secret* .The secret is a pre-shared secure password that the FortiGate uses to authenticate to the FortiAuthenticator.
  3. Set *Authentication method* to*Enforce two-factor authentication* .
  4. Select *Enable FortiToken Mobile push notifications authentication* .
  5. Set *Realms* to*local | Local users* .
5. Go to 
6. Create a user and assign FortiToken mobile to the user on the FortiAuthenticator:
  1. Go to *Authentication > User Management > Local Users* to create a user*sslvpnuser1* .
  2. Enable *Allow RADIUS authentication* and click*OK* to access additional settings.
  3. Enable *Token-based authentication* and select to deliver the token code by*FortiToken* .
  4. Select the FortiToken added from the FortiToken Mobile dropdown menu.
  5. 
Set *Delivery method* to*Email* and fill in the*User Information* section.
  6. Go to *Authentication > User Management > User Groups* to create a group*sslvpngroup* .
  7. Add *sslvpnuser1* to the group by moving the user from*Available users* to*Selected users* .
7. Go to 
8. Install the FortiToken mobile application on your Android or iOS smartphone.The FortiAuthenticator sends the FortiToken mobile activation to the user's email address.
9. Activate the FortiToken mobile through the FortiToken mobile application by entering the activation code or scanning the QR code.

###### To configure SSL VPN using the GUI:

1. Configure the interface and firewall address. The port1 interface connects to the internal network.
  1. Go to *Network > Interfaces* and edit the*wan1* interface.
  2. Set *IP/Network Mask* to*172.20.120.123/255.255.255.0* .
  3. Edit *port1* interface and set*IP/Network Mask* to*192.168.1.99/255.255.255.0* .
  4. Click *OK* .
  5. Go to *Policy & Objects > Address* and create an address for internet subnet*192.168.1.0* .
2. Go to 
3.  Create a RADIUS user and user group:
  1. On the FortiGate, go to *User & Authentication > RADIUS Servers* to create a user to connect to the RADIUS server (FortiAuthenticator).
  2. For *Name* , use*FAC-RADIUS* .
  3. Enter the IP address of the FortiAuthenticator, and enter the *Secret* created above.
  4. Click *Test Connectivity* to ensure you can connect to the RADIUS server.
  5. Select *Test User Credentials* and enter the credentials for*sslvpnuser1* .The FortiGate can now connect to the FortiAuthenticator as the RADIUS client.
  6. Go to *User & Authentication > User Groups* and click*Create New* to map authenticated remote users to a user group on the FortiGate.
  7. For *Name* , use*SSLVPNGroup* .
  8. In *Remote Groups* , click*Add* .
  9. In the *Remote Server* dropdown list, select*FAC-RADIUS* .
  10. Leave the *Groups* field blank.
4. On the FortiGate, go to 
5. Configure SSL VPN web portal:
  1. Go to *VPN > SSL-VPN Portals* to edit the*full-access* portal.This portal supports both web and tunnel mode.
  2. Disable *Enable Split Tunneling* so that all SSL VPN traffic goes through the FortiGate.
6. Go to 
7. Configure SSL VPN settings:
  1. Go to *VPN > SSL-VPN Settings* .
  2. Select the *Listen on Interface(s)* , in this example,*wan1* .
  3. Set *Listen on Port* to*10443* .
  4. Set *Server Certificate* to the authentication certificate.
  5. Under *Authentication/Portal Mapping* , set default Portal*web-access* for*All Other Users/Groups* .
  6. Create new *Authentication/Portal Mapping* for group*sslvpngroup* mapping portal*full-access* .
8. Go to 
9. Configure SSL VPN firewall policy:
  1. Go to *Policy & Objects > Firewall Policy* .
  2. Fill in the firewall policy name. In this example, *sslvpn certificate auth* .
  3. Incoming interface must be *SSL-VPN tunnel interface(ssl.root)* .
  4. Set the *Source Address* to*all* and*Source User* to*sslvpngroup* .
  5. Set the *Outgoing Interface* to the local network interface so that the remote user can access the internal network. In this example:*port1* .
  6. Set *Destination Address* to the internal protected subnet*192.168.1.0* .
  7. Set *Schedule* to*always* ,*Service* to*ALL* , and*Action* to*Accept* .
  8. Enable *NAT* .
  9. Configure any remaining firewall and security options as desired.
  10. Click *OK* .
10. Go to 

###### To configure SSL VPN using the CLI:

1. Configure the interface and firewall address:```
config system interface 
    edit "wan1"
        set vdom "root"
        set ip 172.20.120.123 255.255.255.0
    next
end
```
2. Configure internal interface and protected subnet, then connect the port1 interface to the internal network:```
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
3. Create a RADIUS user and user group:```
config user radius
    edit "FAC-RADIUS"
        set server "172.20.120.161"
        set secret <FAC client secret>
    next
end
```
```
config user group
    edit "sslvpngroup"
        set member "FAC-RADIUS"
    next
end
```
4. Configure SSL VPN web portal:```
config vpn ssl web portal
    edit "full-access"
        set tunnel-mode enable
        set web-mode enable
        set ip-pools "SSLVPN_TUNNEL_ADDR1"
        set split-tunneling disable
     next
end
```
5. Configure SSL VPN settings:```
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
6. Configure one SSL VPN firewall policy to allow remote user to access the internal network:```
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

