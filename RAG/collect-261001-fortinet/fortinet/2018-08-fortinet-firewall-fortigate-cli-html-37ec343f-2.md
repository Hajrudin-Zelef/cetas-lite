---
id: collect-261001-fortinet/fortinet/2018-08-fortinet-firewall-fortigate-cli-html-37ec343f-2
title: "2018-08-fortinet-firewall-fortigate-cli-html-37ec343f"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/2018-08-fortinet-firewall-fortigate-cli-html-37ec343f.md
source_anchor: ""
source_lines: [271, 365]
sha256: 72928257f5d42b58d1f42339ac7a9a9e6f8ac61bfae432d9ac39ed13039e46e2
---

# 2018-08-fortinet-firewall-fortigate-cli-html-37ec343f

| FWF60D # tree execute ping-option | 
| `-- ping-options -- data-size -- <integer>  (0)` | 
|                 `\|- df-bit -- <string>  (0)` | 
|                 `\|- pattern -- <string>  (0)` | 
|                 `\|- repeat-count -- <string>  (0)` | 
|                 `\|- source -- <string>  (0)` | 
|                 `\|- interface -- <string>  (0)` | 
|                 `\|- timeout -- <integer>  (0)` | 
|                 `\|- adaptive-ping -- <string>  (0)` | 
|                 `\|- interval -- <integer>  (0)` | 
|                 `\|- tos -- <string>  (0)` | 
|                 `\|- ttl -- <integer>  (0)` | 
|                 `\|- validate-reply -- <string>  (0)` | 
|                 `\|- view-settings`  | 
|                 `+- reset` ```   ``` ``` FWF60D # tree system ddns -- [ddns] --*ddnsid  (0,4294967295)           \|- ddns-server            \|- ddns-server-ip            \|- ddns-zone  (65)           \|- ddns-ttl  (60,86400)           \|- ddns-auth            \|- ddns-keyname  (65)           \|- ddns-key            \|- ddns-domain  (65)           \|- ddns-username  (65)           \|- ddns-sn  (65)           \|- ddns-password            \|- use-public-ip            \|- update-interval  (60,2592000)           \|- clear-text            \|- ssl-certificate  (36)           \|- bound-ip            +- [monitor-interface] --*interface-name  (65) FWF60D #   ``` ```   ``` ```   ``` ```   ``` ```   ``` ```   ``` ```   ``` ```   ``` ```   ```  | 

## 7. Disable Local Logs

If you donot want to log your local logs, here is a way to disable it. But you may lose some view of local broadcasting noise.

FWF60D # config log memory filter

```
FWF60D (filter) # set local-traffic disable
FWF60D (filter) # end
FWF60D #  
```
## **8. Configuration Example**

Fortigate Device Information:

LAN : 192.168.200.1/24

WAN : 85.86.87.2/29

Default Gateway to ISP: 85.86.87.1

|  config system global# Set the http admin port to 80/tcp set admin-port 80 # Set the https admin port to 443/tcp set admin-sport 443 # Set the ssh admin port to 22/tcp set admin-ssh-port 22 # Set the telnet admin port to 23/tcp set admin-telnet-port 23 # Set the hostname set hostname “FW-Office-1” # Set the ntp server to “0.ca.pool.ntp.org” and enable it set ntpserver “0.ca.pool.ntp.org” set ntpsync enable # Set to 43200 seconds the tcp-halfclose timer set tcp-halfclose-timer 43200 end # Set the telnet 23/tcp port timeout to 43200 seconds. config system session-ttl set default 43200 config port edit 23 set timeout 43200 next end # Set the IP address and administrative access options (ping https http) for lan interface. config system interface edit “lan” set ip 192.168.200.1 255.255.255.0 set allowaccess ping https http set type physical next # Set the IP address and administrative access options (ping https) for wan interface. # Set “gateway Detect” option enable and set the “Ping Server” destination. # Set the interface speed to 10 Mb/s Half Duplex, this is useful for some connections like radio bridge. edit “wan1″ set ip 85.86.87.2 255.255.255.248 set allowaccess ping https set gwdetect enable set detectserver “85.86.87.23″ set type physical set speed 10half next end # Set DNS Servers and DNS options config system dns set primary 192.168.200.3 set secondary 8.8.8.8 set domain ” set autosvr disable set dns-cache-limit 5000 set cache-notfound-responses disable end # Set a firewall policy to enable traffic from lan TO WAN using NAT # Set a protection profile (a default one) called “scan” config firewall policy edit 1 set srcintf “lan” set dstintf “wan″ set srcaddr “all” set dstaddr “all” set action accept set schedule “always” set service “ANY” set profile-status enable set profile “scan” set nat enable next end # Set a default gateway on the WAN interface config router static edit 1 set device “wan″ set gateway 85.86.87.1 end | 

## 9. Proper shut down of FortiGate unit

## 10. BGP

## 11. Restrick Access to Web and SSL VPN

### 1. For FortiGate Web Portal

**Command line only**: local-in-policy

```
sh full firewall local-in-policy 
config firewall local-in-policy
    edit 1
        set uuid d7f13336-22a0-51ef-ee79-444132007804
        set intf "port1"
        set srcaddr "Allowed_Countries"
        set srcaddr-negate enable
        set dstaddr "wan-10.254.3.5/32"
        set internet-service-src disable
        set dstaddr-negate disable
        set action deny
        set service "ALL"
        set service-negate disable
        set schedule "always"
        set status enable
        set comments ''
    next
end
```
Notes:https://docs.fortinet.com/document/fortigate/7.4.4/administration-guide/363127/local-in-policy

If a local-in-policy is not functioning correctly and traffic that should be blocked is being allowed through, the issue may be that the implicit deny local-in-policy has not been created. Unlike IPv4 policies, there is no default implicit deny policy. The implicit deny policy should be placed at the bottom of the list of local-in-policies. Local-in-policies are created for each interface, but if you want to create a general implicit deny rule for all interfaces for a specific service, source, address, or destination address, use the `any` interface.

View default Local In Policy from Web GUI:

### 2. For SSL VPN

- Using Geograph Limitation

- Extend Lock time from 60 seconds to 600 seconds when user failed log in using the combination of username and password.

This indicates if user enters incorrect username/password combinations continuously twice, the firewall will block attempts and prompt with message as 'Too many bad attempts. Please try again in few minutes'.

WAN : 85.86.87.2/30



ReplyDelete
Subnetmask : 255.255.255.252

Default Gateway to ISP: 85.86.87.1

LAN POOL : 117.197.102.180 to 117.197.102.187

Subnet :- 255.255.255.248 how to configure in cisco router
