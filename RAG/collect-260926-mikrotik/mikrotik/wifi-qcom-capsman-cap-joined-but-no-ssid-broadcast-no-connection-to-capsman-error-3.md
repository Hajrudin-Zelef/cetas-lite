---
id: collect-260926-mikrotik/mikrotik/wifi-qcom-capsman-cap-joined-but-no-ssid-broadcast-no-connection-to-capsman-error-3
title: "NAME    CONFIGURATION.MODE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/wifi-qcom-capsman-cap-joined-but-no-ssid-broadcast-no-connection-to-capsman-error.md
source_anchor: ""
source_lines: [421, 480]
sha256: 9bf87f9c07d77ca47a5c00996b29dc95896a56d80a559be9fb35e47cbaf7df4d
---

# NAME    CONFIGURATION.MODE

I just finished integrating the hAP ac2 (after upgrading it to 7.21.2 and installing the wifi-qcom-ac package). Just like with the ax lite, using 'capsman-or-local' in the configuration profile was the only way to achieve a stable connection. Once the wifi interfaces were added to the local bridge on the ac2, provisioning completed successfully.

Here is the result from the hAP ac2 acting as a CAP:

```
Flags: M - MASTER; B - BOUND; R - RUNNING
Columns: NAME
NAME
;;; managed by CAPsMAN 04:F4:1C:34:81:DD%bridgeLocal, traffic processing on CAP
;;; mode: AP, SSID: wMw, channel: 2417/n/Ce
0 MBR wifi1
```

Whether the CAPs are 'resorting to local' or not, this is the only configuration that results in a stable MBR state. It appears that in this ROS version, the 'pure capsman' manager setting interferes with the provisioning handshake for remote CAPs.

             
            
           
          
            
            
              Maybe it is a bug specific on 7.21.2?

But it is more probable that there is *something else* escaping us.

Can you post your latest exports?

In your original ones I noticed (maybe relevant, maybe not) that you had the setting in two (actually three) places (not specifically this one, but I remember people having issues when the same setting was in more than one  places).

In:

/interface wifi

set [ find default-name=wifi1 ] configuration=conf1 configuration.**manager=local** .mode=ap disabled=no

set [ find default-name=wifi2 ] configuration=conf1 configuration.**manager=local** .mode=ap disabled=no

And in:

/interface wifi configuration

add country=Hungary datapath=dp1 disabled=no **manager=capsman** name=conf1 security=sec1 ssid=wMw steering=st1

             
            
           
          
            
            
              
May be it is time to load 7.19.6 or some similar trusted version to test this hypothesis?

             
            
           
          
            
            
              
I don't remember exactly where, but the developer wrote that in the configuration settings, you don't need to set the manager field; leave it blank by default.
