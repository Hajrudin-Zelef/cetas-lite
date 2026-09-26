---
id: collect-260926-mikrotik/mikrotik/cannot-see-wifi-with-capsman-solved-2
title: "2025-02-19 09:44:55 by RouterOS 7.17.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2025-02-19"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/cannot-see-wifi-with-capsman-solved.md
source_anchor: ""
source_lines: [294, 575]
sha256: 6b6f9bd359da15dbb12eac6e5964f3df921f1538bf851607fc71d52e10bf5e8d
---

# 2025-02-19 09:44:55 by RouterOS 7.17.2

2.>> Setup the cAP AX to retrieve the CAPsMAN settings from the router

Somewhere I'm making a mistake using the Winbox user interface to setup CAPsMAN as the cAP AX is receiving some CAPsMAN config but not everything. Below are the scripts for the router and cAP AX

**Router Script:**

# 2025-02-19 13:36:40 by RouterOS 7.17.2

# software id = FC9Z-1KWF


# model = RB5009UPr+S+

# serial number = *****

/interface bridge

add name=bridge1

/interface ethernet

set [ find default-name=ether1 ] comment=WAN name="ether1[internet]"

/interface wifi channel

add band=2ghz-ax disabled=no frequency=2412,2432,2472 name="2.4Ghz AX" width=

20mhz

add band=5ghz-ax disabled=no frequency=2300-7300 name="5Ghz AX" width=

20/40/80mhz

/interface wifi datapath

add bridge=bridge1 disabled=no name=datapath1

/interface wifi security

add authentication-types=wpa2-psk,wpa3-psk disabled=no name=secKoekWiFi

/interface wifi configuration

add channel="2.4Ghz AX" country="South Africa" datapath=datapath1 disabled=no 

manager=capsman mode=ap name=cfg_WiFi2GHz_AX security=secKoekWiFi ssid=

KoekiesWiFi

add channel="5Ghz AX" country="South Africa" datapath=datapath1 disabled=no 

mode=ap name=cfg_WiFi5GHz_AX security=secKoekWiFi security.connect-group=

"" ssid=KoekiesWiFi

/ip pool

add name=dhcp_pool0 ranges=192.168.11.100-192.168.11.199

/ip dhcp-server

add address-pool=dhcp_pool0 interface=bridge1 lease-time=8h name=dhcp1

/interface bridge port

add bridge=bridge1 interface=ether2

add bridge=bridge1 interface=ether3

add bridge=bridge1 interface=ether4

add bridge=bridge1 interface=ether5

add bridge=bridge1 interface=ether6

add bridge=bridge1 interface=ether7

add bridge=bridge1 interface=ether8

add bridge=bridge1 interface=sfp-sfpplus1

/ip firewall connection tracking

set udp-timeout=10s

/ip neighbor discovery-settings

set discover-interface-list=!dynamic

/interface ovpn-server server

add mac-address=**:**:**:**:**:** name=ovpn-server1

/interface wifi capsman

set ca-certificate=auto certificate=auto enabled=yes interfaces=all 

package-path="" require-peer-certificate=no upgrade-policy=none

/interface wifi provisioning

add action=create-dynamic-enabled disabled=no

/ip address

add address=192.168.11.1/24 comment="LAN IP" interface=bridge1 network=

192.168.11.0

/ip dhcp-client

add comment="Internet WAN" interface="ether1[internet]"

/ip dhcp-server network

add address=192.168.11.0/24 gateway=192.168.11.1

/ip firewall filter

add action=fasttrack-connection chain=forward connection-state=

established,related hw-offload=yes

add action=accept chain=forward connection-state=established,related

add action=drop chain=forward connection-nat-state="" connection-state=

invalid

add action=drop chain=forward connection-mark="" connection-nat-state=!dstnat 

connection-state=new in-interface="ether1[internet]"

/ip firewall nat

add action=masquerade chain=srcnat comment="Internet Connection Rule" 

out-interface="ether1[internet]"

/ip ipsec profile

set [ find default=yes ] dpd-interval=2m dpd-maximum-failures=5

/ip service

set telnet disabled=yes

set ftp disabled=yes

set www disabled=yes

set api disabled=yes

set api-ssl disabled=yes

/system clock

set time-zone-name=Africa/Johannesburg

/system identity

set name="MikroTik Router"

/system note

set show-at-login=no

**cAP AX Script:**

# 2025-02-19 13:37:26 by RouterOS 7.17.2

# software id = WWE0-FFJ0


# model = cAPGi-5HaxD2HaxD

# serial number = ******

/interface bridge

add name=bridge1

/interface wifi

# managed by CAPsMAN 192.168.11.1, traffic processing on CAP

set [ find default-name=wifi1 ] configuration.manager=capsman .mode=ap 

disabled=no

# managed by CAPsMAN 192.168.11.1, traffic processing on CAP

set [ find default-name=wifi2 ] configuration.manager=capsman .mode=ap 

disabled=no

/interface bridge port

add bridge=bridge1 interface=ether1

add bridge=bridge1 interface=ether2

add bridge=bridge1 interface=wifi1

add bridge=bridge1 interface=wifi2

/interface wifi cap

set caps-man-addresses=192.168.11.1 enabled=yes

/interface wifi capsman

set enabled=yes interfaces=bridge1 package-path="" require-peer-certificate=

no upgrade-policy=none

/interface wifi provisioning

add action=create-dynamic-enabled disabled=no

/ip address

add address=192.168.11.41/24 interface=bridge1 network=192.168.11.0

/ip dhcp-client

add interface=bridge1

/system clock

set time-zone-name=Africa/Johannesburg

/system identity

set name="WiFi Kantoor"

/system note

set show-at-login=no

Belwo is the screenshots from the Router and cAP AX:

**Router:**


**cAP AX:**
             
            
           
          
            
            
              manager=capsman needs to removed from configuration you are passing via CAPsMAN to CAP, and you need to two separate provisioning rules, it would be best to add one rule for each, using supported-bands as a matcher.

             
            
           
          
            
            
              ```
/interface wifi provisioning
add action=create-dynamic-enabled disabled=no
```

What should be provisioned? You are missing the configuration (“master-configuration”) here.

```
/interface wifi capsman
set enabled=yes interfaces=bridge1 package-path="" require-peer-certificate=\
no upgrade-policy=none
```

You should disable CAPsMAN on the cAP AX, as it has no purpose.

             
            
           
          
            
            
              Thank you, I was setting the provision but as you mentioned no “master” was selected (I did not know that this was required - BIG mistake!!). I selected the master and now everythig is working.
