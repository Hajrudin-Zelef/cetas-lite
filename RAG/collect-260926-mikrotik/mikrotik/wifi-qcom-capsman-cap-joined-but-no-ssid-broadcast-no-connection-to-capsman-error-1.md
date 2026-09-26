---
id: collect-260926-mikrotik/mikrotik/wifi-qcom-capsman-cap-joined-but-no-ssid-broadcast-no-connection-to-capsman-error-1
title: "NAME    CONFIGURATION.MODE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-02-04"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/wifi-qcom-capsman-cap-joined-but-no-ssid-broadcast-no-connection-to-capsman-error.md
source_anchor: ""
source_lines: [1, 215]
sha256: d4f8d72231f0ec2062aa63212520b358cf5c612f696664a1a2141d3d0f2a316a
---

# NAME    CONFIGURATION.MODE

Hello Guys,

I've been trying to set up my managed wifi network for a week now, but I'm struggling even with the first CAP. It seems like a simple configuration, but I'm stuck.

My Setup:

- CAPsMAN: hAP ax2
- CAP: hAP ax lite
- Firmware:All devices are running RouterOS 7.21.2 with the wifi-qcom package installed.

The Issue:

The CAP appears to be under CAPsMAN control, but it doesn't seem to broadcast the SSID. On the floor where the ax lite is located, the signal is non-existent, as if the radio were off.

Strange symptom: On the CAP side, the interface shows as "Bound", but on the CAPsMAN side, the dynamic interface shows a comment: "no connection to CAPsMAN"

I lack the experience to spot the error here. Any help would be greatly appreciated.

---**1. CAP Side - Interface Status:**

The MBI flags are present, which is concerning.

````
Flags: M - MASTER; B - BOUND; I - INACTIVE
#   NAME    CONFIGURATION.MODE
;;; managed by CAPsMAN 04:F4:1C:34:81:DD%bridgeLocal
0 MBI wifi1   ap
```**2. CAP Configuration:**
```/interface wifi
set [ find default-name=wifi1 ] configuration.manager=capsman .mode=ap disabled=no
/interface wifi cap
set certificate=request discovery-interfaces=bridgeLocal enabled=yes slaves-datapath=*1
```**3. CAP Logs (after toggle):**
```2026-02-04 09:57:07 caps,info selected CAPsMAN ax2@04:F4:1C:34:81:DD%*7
2026-02-04 09:57:07 caps,info connected to ax2@04:F4:1C:34:81:DD%*7
**1. Interface Status on Manager:**
Note the "no connection" comment.
[netadm@ax2] /interface/wifi> print
Flags: M - MASTER; B - BOUND; I - INACTIVE, R - RUNNING
#      NAME         CONFIGURATION.MODE   CONFIGURATION.SSID
;;; no connection to CAPsMAN
0 MBI  cap-wifi1                         wMw
**2. CAPsMAN Configuration:**
/interface wifi capsman
set ca-certificate=auto certificate=auto enabled=yes interfaces=bridge require-peer-certificate=no upgrade-policy=none
/interface wifi configuration
add country=Hungary datapath=dp1 disabled=no manager=capsman name=conf1 security=sec1 ssid=wMw steering=st1
/interface wifi provisioning
add action=create-dynamic-enabled disabled=no master-configuration=conf1
**3. CAPsMAN Logs:**
I see a lot of "no suitable CAPsMAN" messages.
2026-02-04 09:57:07 caps,info axl@48:A9:8A:A8:59:8C%*9 joined
2026-02-04 09:57:17 caps,debug no suitable CAPsMAN
2026-02-04 09:57:30 caps,debug no suitable CAPsMAN
Thank you in advance for your help!
````

             
            
           
          
            
            
              This worries me:

Can you reset the CAP to CAPS Mode:

```
/system reset capsmode
```

And please provide `/interface wifi export` for both CAP and CAPsMAN and post between Preformatted text tags by using the </> button.

             
            
           
          
            
            
              Hi,

I have performed the full system reset on the CAP as requested (/system reset-configuration capsmode=yes). Unfortunately, the situation remains exactly the same:

On CAPsMAN side: I still see the "No connection to CAPsMAN" error on the dynamic interface.

On CAP side: The interface says "managed by CAPsMAN 04:F4:1C:34:81:DD%bridgeLocal", but remains Inactive.

Here are the fresh, clean exports from both devices:

## CAPsMAN (hAP ax2) /interface/wifi export:

```
/interface wifi security
add authentication-types=wpa2-psk,wpa3-psk disabled=no name=sec1
/interface wifi steering
add name=st1 neighbor-group=dynamic
/interface wifi
set [ find default-name=wifi1 ] configuration=conf1 configuration.manager=local .mode=ap disabled=no
set [ find default-name=wifi2 ] configuration=conf1 configuration.manager=local .mode=ap disabled=no
/interface wifi cap
set certificate=request discovery-interfaces=bridge enabled=yes
/interface wifi capsman
set ca-certificate=auto certificate=auto enabled=yes interfaces=all require-peer-certificate=no upgrade-policy=none
/interface wifi configuration
add country=Hungary datapath=dp1 disabled=no manager=capsman name=conf1 security=sec1 ssid=wMw steering=st1
/interface wifi datapath
add bridge=bridge disabled=no name=dp1
/interface wifi provisioning
add action=create-dynamic-enabled disabled=no master-configuration=conf1
```

## CAP (hAP ax lite) / interface/wifi export:

```
/interface wifi
managed by CAPsMAN 04:F4:1C:34:81:DD%bridgeLocal
set [ find default-name=wifi1 ] configuration.manager=capsman datapath=capdp
/interface wifi cap
set discovery-interfaces=bridgeLocal enabled=yes slaves-datapath=capdp
/interface wifi datapath
```

add bridge=bridgeLocal comment=defconf disabled=no name=capdp

Thank you for your continued help!

             
            
           
          
            
            
              The only thing I notice is that you set interfaces to all:

```
/interface wifi capsman
set ca-certificate=auto certificate=auto enabled=yes interfaces=all require-peer-certificate=no upgrade-policy=none
```

Better change the interfaces to the bridge on the CAPsMAN.

In regards to the "no connection". Is this the radio of the CAP or is it the radio on the CAPsMAN? You can figure that out by checking the MAC Address.

If that doesn't provide information, I would reset the wifi settings on the CAPsMAN and reconfigure it (this can be done by executing the config commands).

             
            
           
          
            
            
              Hi erlinden,

I have changed the CAPsMAN interfaces back to bridge as suggested. Regarding the previous interfaces=all setting, I only used it as a "brute-force" test to eliminate any potential interface-level blocking while troubleshooting; now it's back to the standard bridge configuration.

Regarding the "no connection" error: It appears on the interface belonging to the CAP (hAP ax lite). Its radio MAC is 48:A9:8A:A8:59:90. The local radios of the CAPsMAN (ax2) are working fine in local mode.

Even after changing the interface to bridge, the CAP still shows "Inactive" on the manager side with the same error. Here is the detailed status of the CAP interface from the CAPsMAN side:

```
Flags: M - master; D - dynamic; B - bound; X - disabled, I - inactive, R - running
0 MDBI ;;; no connection to CAPsMAN
cap="MikroTik@48:A9:8A:A8:59:8C%*9" name="cap-wifi1" mac-address=48:A9:8A:A8:59:90 arp-timeout=auto
radio-mac=48:A9:8A:A8:59:90 configuration=conf1
configuration.ssid="wMw" .country=Hungary .manager=capsman
security.authentication-types=wpa2-psk,wpa3-psk
datapath.bridge=bridge
steering.neighbor-group=dynamic
```

I am now considering your suggestion to reset the wifi configuration on the CAPsMAN side and re-apply the settings.

Before I do that, is there any specific log topic I should enable to see why the manager is rejecting the connection after provisioning the interface?

Regards

             
            
           
          
            
            
              
Not AFAIK.

Make sure to do an export with show sensitive:

```
/interface wifi
export show-sensitive
```

But befor doing so...can you delete the dynamic interface?

Is the interface removed when you disable CAP on the CAPS?

Can you do an `/interface export` on the CAPsMAN and share it?

             
            
           
          
            
            
              Hi erlinden,

Here are the results of the tests you requested:

Manual removal on CAPsMAN: I could not manually delete the dynamic interface while the CAP was active.

Interface removal via CAP: When I disable CAP mode on the CAP device, the dynamic interface disappears from the CAPsMAN list immediately.

Full interface export from CAPsMAN (Passphrases obfuscated):


