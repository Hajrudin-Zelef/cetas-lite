---
id: collect-260926-mikrotik/mikrotik/mikrotik-openvpn-client-over-udp-stopped-forwarding-traffic-on-routeros-7-10-1
title: "mikrotik-openvpn-client-over-udp-stopped-forwarding-traffic-on-routeros-7-10-1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-openvpn-client-over-udp-stopped-forwarding-traffic-on-routeros-7-10-1.md
source_anchor: ""
source_lines: [1, 196]
sha256: c302eedbce74a8b9004c85455f9e92f1dd0231a183f4f969786d1c7f840a25db
---

# mikrotik-openvpn-client-over-udp-stopped-forwarding-traffic-on-routeros-7-10-1

Hi team,

I just upgraded from 7.9.x to 7.10.1 and noticed that my OpenVPN connection is coming up (IP is properly received by the client), but there is no connectivity over tunnel. No configuration change was made, nothing is blocked by firewall, MTU is set to 1300 bytes and was working perfectly fine before upgrade.

I could reproduce this issue on Mikrotik CAP AC XL and HAP AC3.

If I change on server and client from UDP to TCP and keep previous config - everything seems to be working fine, so most likely a bug in UDP client ( or network stack ) in RouterOS 10.x

I added ‘accept’ rules to firewall and can see packet counters incrementing in both directions, but neither site responds to ping nor can open any connection.

             
                
            
           
          
            
            
              Facing same issue.

Definitely it is a firmware issue.

             
            
           
          
            
            
              In version  7.9.x.

After the OVPN client connects, the IP4 address given to client appears in the list of ipv4 addresses

ip/address/print

In version 7.10.x It’s does not. And no one packets pass through the router.

My solution is: don’t use version 7.10.x

             
            
           
          
            
            
              Same issue on 7.10. Cannot rollback to prev firmware.

Temporary solution for me was to assign permament binding for ovpn connection on server side and assign static ip address for that connection. Since then ovpn tunnel works but this solution is some kind of cucumbersome.

Hope Mikrotik fix this in next fw releases.

             
            
           
          
            
            
              I got static IP set for client, client gets the IP, firewall accept rule counters are incrementing but data is not forwarded nor replied over tunnel. Only solution was to switch to TCP based tunnel.

             
            
           
          
            
            
              Hi,

is possible to make some statement about this bug from Mikrotik dev and support team? Is this issue in progress? I read changelog 7.11.rc1, but there is no information about this.

Many thanks for replies…

             
            
           
          
            
            
              Same problem here. This solution works fine for me but changing udp to tcp or tcp to udp does not solve the problem in my case when client connect, ip address is not dynamically added to address table.

             
            
           
          
            
            
              I had the same problem but now it’s working with 7.11rc3 (2023-Aug-09 17:41)

             
            
           
          
            
            
              I confirm upgrading firmware to 7.11rc3 solves this issue.

             
            
           
          
            
            
              No,

I have tested yet and issue persists in RouterOS 7.11 and 7.11rc4.

tcp mode works fine, udp not

rolling back to 7.9.2, this version works in both modes :-/

running on arm hAP ac^3 with hw crypto CBC

             
            
           
          
            
            
              On a RB4011 with arm architecture

For me version 7.11 did’nt help either. OpenVPN UDP is broken, TCP still works.

Workaround

revert to 7.9

I mailed support@mikrotik.com, hope they pick it up.

Regards,

Tom

             
            
           
          
            
            
              Had the same issue on my CCR1009-8G-1S-1S+ after upgrading from 7.3.1 to 7.11.2. Switching to TCP or downgrading to 7.9.2 helped.

             
            
           
          
            
            
              Same here. With 7.11.2 the issue persist.

             
            
           
          
            
            
              Same here…Lost 3 remote locations by updating my Mikrotiks… MMIPS and ARM devices

Can you guys investigate this issue, please?

             
            
           
          
            
            
              In 7.12RC4, for me it is working again, even there is nothing in changelog about this. Hap AC3

             
            
           
          
            
            
              RC5 broke it again, reverter to RC4

             
            
           
          
            
            
              Switch to wireguard or ipsec,  live longer, less gray hair or ulcers.

OVPN = https://media.giphy.com/media/xxCNsOokj8Rxp8EtUL/giphy.gif

             
            
           
          
            
            
              I know what do you mean , unfortunately I can not switch … some devices that i have to monitor are not able to install wireguard…

             
            
           
          
            
            
              The current stable version **7.12 Stable** works for me. OpenVPN UDP came back to life
