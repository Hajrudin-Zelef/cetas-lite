---
id: collect-260926-mikrotik/mikrotik/wifi-qcom-capsman-cap-joined-but-no-ssid-broadcast-no-connection-to-capsman-error-2
title: "NAME    CONFIGURATION.MODE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-02-05"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/wifi-qcom-capsman-cap-joined-but-no-ssid-broadcast-no-connection-to-capsman-error.md
source_anchor: ""
source_lines: [216, 420]
sha256: 634ee1c5545ea7100b61ea64025603839adf677427084577399601ea926a57b8
---

# NAME    CONFIGURATION.MODE

```
2026-02-05 12:34:13 by RouterOS 7.21.2
software id = CKNL-JXW2
model = C52iG-5HaxD2HaxD
serial number = HJR0ARECNZ1
/interface bridge
add admin-mac=04:F4:1C:34:81:DD auto-mac=no comment=defconf name=bridge
/interface ethernet switch
set 0 cpu-flow-control=yes
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wifi datapath
add bridge=bridge disabled=no name=dp1
/interface wifi security
add authentication-types=wpa2-psk,wpa3-psk disabled=no name=sec1 passphrase=
***
/interface wifi steering
add name=st1 neighbor-group=dynamic
/interface wifi configuration
add country=Hungary datapath=dp1 disabled=no manager=capsman name=conf1 
security=sec1 ssid=wMw steering=st1
/interface wifi
set [ find default-name=wifi1 ] configuration=conf1 configuration.manager=
local .mode=ap disabled=no
set [ find default-name=wifi2 ] configuration=conf1 configuration.manager=
local .mode=ap disabled=no
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1 list=WAN
/interface wifi cap
set certificate=request discovery-interfaces=bridge
/interface wifi capsman
set ca-certificate=auto certificate=auto enabled=yes interfaces=bridge 
require-peer-certificate=no upgrade-policy=none
/interface wifi provisioning
add action=create-dynamic-enabled disabled=no master-configuration=conf1
```

I hope this helps to identify any bridge or tagging issues.******

             
            
           
          
            
            
              Doesn't make sense as everything is set correct (as far as I can see).

Please remove your passphrase as it is of no use for this problem.

             
            
           
          
            
            
              Hi,

I have edited the post and replaced the requested information with '***'. I turned to this forum specifically to overcome this issue with the help of more experienced users. I have truly examined everything that I could think of. I am looking forward to any advice from the pros.

Thank you.

             
            
           
          
            
            
              Just to be sure, what RouterOS and firmware version are you running on the CAPS and CAPsMAN? Make sure to have all versions equal.

Before resetting CAPsMAN, please give this a try on the **CAP**:

```
# reset the wifi interface
/interface wifi reset wifi1
# this could give an error when the datapath is still available
/interface wifi datapath
add bridge=bridgeLocal comment=defconf disabled=no name=capdp
# proceed even if the above statement gives an error
/interface wifi
set [ find default-name=wifi1 ] configuration.manager=capsman datapath=capdp
/interface wifi cap
set discovery-interfaces=bridgeLocal enabled=yes slaves-datapath=capdp
```

             
            
           
          
            
            
              Hi,

I finally found the misconfiguration! Although I don't 100% understand why it affected the remote CAP this way (as I thought it would only cause issues for the CAPsMAN's local radios), changing this one setting fixed everything.

The culprit: In my /interface wifi configuration, I had the manager set strictly to capsman: /interface wifi configuration add ... manager=capsman

The fix: I changed the manager mode to capsman-or-local. After this, cap disable and enable again,  and the connection established immediately on both sides.

CAPsMAN (ax2) status now:

```
Flags: M - MASTER; D - DYNAMIC; B - BOUND; R - RUNNING
NAME       CONFIGURATION.MODE  CONFIGURATION.SSID
;;; operated by CAP 48:A9:8A:A8:59:8C%bridge, traffic processing on CAP
0 MDB  cap-wifi1                      wMw
1 M BR wifi1      ap                  wMw
2 M BR wifi2      ap                  wMw
```

CAP (ax lite) status now:

```
Flags: M - MASTER; B - BOUND; R - RUNNING
NAME
;;; managed by CAPsMAN 04:F4:1C:34:81:DD%bridgeLocal, traffic processing on CAP
;;; mode: AP, SSID: wMw, channel: 2472/ax/eC
0 MBR wifi1
```

It seems that when using manager=capsman, the system might be too restrictive for certain CAP-side operations, even if provisioning seems to succeed at first. Changing it to capsman-or-local solved the "no connection" loop.

But anyway one could tell precise explanation, please let us know.

I hope this helps someone else struggling with the same issue on RouterOS 7.21.2!

Thank you.

             
            
           
          
            
            
              Huh, that's a funny one. IMO it's intended for configuring local wifi interfaces but it seems that it interferes with capsman ...

CAP interfaces, run by CAPsMAN, are "managed by local" on CAPsMAN device. And actual local interfaces should be managed by "local" anyway. Surely, the same setting on CAP device should be set to "capsman".

             
            
           
          
            
            
              
Exactly on which device did you change this mode ?

If on capsman controller which also has local radios, it should be local. Not even capsman-or-local.

If on APs, it should be capsman (unless you also foresee some local config in case communication to capsman controller becomes unavailable, it's a fallback option).

If this was on controller and this resulted in all remote caps to not behave as they should, then it might be worthwhile to file a bug report.

             
            
           
          
            
            
              I agree, it's definitely a 'funny' one. Logically, you are right: the CAP device should be 'capsman' and the local should be 'local'.

             
            
           
          
            
            
              Hi,

I changed it on the **CAPsMAN controller. That is why I couldn’t understand exactly the situation and would be happy for an explanation, how it could work.**

To clarify: the 'manager=capsman' setting was part of the configuration profile that the controller (hapac2) was provisioning to the remote CAPs (hapax lite). As soon as I updated that profile on the controller to 'capsman-or-local', and remote CAP connected.

Before this change, the remote CAPs would see the CAPsMAN, attempt to bind, but then immediately drop with a 'no connection to CAPsMAN' message, even though the configuration seemed to be pushed successfully.

             
            
           
          
            
            
              But still, on the hapAC2, it should be only "local".

As holvoeth explained, the setting for capsman-or-local is intended on the CAPS devices, NOT on the CAPSMAN device, as a sort of failover/fallback configuration in case of connection problems with the CAPSMAN.

What should be happening right now should be that the hapAC2 (the CAPSMAN) looks for a CAPSMAN on the LAN, cannot find it (as there is none) and resorts to use the local.

Setting it as local should avoid this initial search for CAPSMAN (though probably it has no particular side effect).

             
            
           
          
            
            
              To clarify the topology and address the recent comments: In my setup, the ax2 is the sole CAPsMAN controller. Both the ax lite and the hAP ac2 are configured strictly as CAP devices.

Following the logic mentioned by holvoeth and jaclaz, setting the manager to 'capsman' on these CAP devices should be the 'correct' way. However, in practice with ROS 7.21.2, that setting caused a continuous connection loop—the CAPs would bind and then immediately drop.

