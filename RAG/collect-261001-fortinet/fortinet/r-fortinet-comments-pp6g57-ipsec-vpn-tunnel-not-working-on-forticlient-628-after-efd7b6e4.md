---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-pp6g57-ipsec-vpn-tunnel-not-working-on-forticlient-628-after-efd7b6e4
title: "r-fortinet-comments-pp6g57-ipsec-vpn-tunnel-not-working-on-forticlient-628-after-efd7b6e4"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2021-15-09"]
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-pp6g57-ipsec-vpn-tunnel-not-working-on-forticlient-628-after-efd7b6e4.md
source_anchor: ""
source_lines: [1, 36]
sha256: 4f65a83f8c21abe4f79bf5a9e2ecf116e5e386d0f1a3ee2a1966fc042ba08092
---

# r-fortinet-comments-pp6g57-ipsec-vpn-tunnel-not-working-on-forticlient-628-after-efd7b6e4

IPSEC VPN Tunnel not working on FortiClient 6.2.8 after upgrading EMS to 6.4.4 - SOLVED 
        
    After upgrading our EMS Server from 6.2 to 6.4.4 build 1658, the IPSEC VPN Tunnels on FortiClients version 6.2.8.1012 stopped working.
Fortinet Support's answer was :
      
      This is known issue reported here #0723465 with summary "EMS 6.4.4 profiles do not sync IPSEC Phase 2 configuration to FortiClient 6.2.8.1012"
- To fix this, you will upgrade forticlients to version 6.4.6
    
However, it will take us some time to get all users updated and we need this working now, so I continued troubleshooting.
Note, our profile has 2 tunnels, the first one is SSL-VPN and the second one is IPSEC, the SSL-VPN Tunnel was working fine but I think this was only because it was the 1st one on the list, so maybe if you only have one VPN tunnel on your profile you will not have this issue, here's why:
After enabling debug logging on the profile, I was able to find the following error:
 9/15/2021 11:20:06 PM    Debug    Config    ImportConfig: tag <\forticlient_configuration\vpn\ipsecvpn\connections\connection\traffic_control\mode> has an invalid value 1. Was imported as main.
9/15/2021 11:20:06 PM    Error    Config    program=FCConfig msg=ImportConfig: VPN connection name is empty
It appears that Traffic Control XML Tag was introduced on 6.4.1:
 The VPN <traffic_control> XML tag contains global information controlling application-based split tunnel.
I removed the traffic control elements from the XML Config on EMS and the IPSEC Tunnel started working. But you have to be careful, any change using the EMS GUI to the profile will add them back and breaks the IPSEC VPN again. To fix it, simple delete the following from Traffic Control sections (one on each tunnel, SSL/IPSEC, must remove from all of them).
 BEFORE - NOT COMPATIBLE WITH 6.2:
                  <traffic_control>
                     <enabled>0</enabled>
                     <mode>1</mode>
                     <apps>
                     </apps>
                     <fqdns>
                     </fqdns>
                     <isdb_objects>
                     </isdb_objects>
                 </traffic_control>
AFTER - COMPATIBLE WITH 6.2:
                  <traffic_control>
                     <enabled>0</enabled>
                  </traffic_control>
I believe that the FCConfig on FC 6.2 starts importing the profile and when it gets to traffic_control\mode it doesn't know what that is and stops there, since traffic control is at the end of the connection, the first connection was almost fully imported (enough for it to work) before halting, causing the 2nd vpn connection to fail to load properly.
Note: after removing the traffic control elements, I'm still seeing the Error message "VPN connection name is empty" but my VPN tunnels are working so not sure what that is about.
I hope this saves someone a ton of time!
Cheers!
E.
