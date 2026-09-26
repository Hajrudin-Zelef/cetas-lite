---
id: collect-260926-mikrotik/mikrotik/is-my-routerboard-rb750r2-bricked-no-response-from-router-for-netinstall
title: "is-my-routerboard-rb750r2-bricked-no-response-from-router-for-netinstall"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2024-11-26"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/tools/is-my-routerboard-rb750r2-bricked-no-response-from-router-for-netinstall.md
source_anchor: ""
source_lines: [1, 70]
sha256: 15f85035b4930bc0d34705d557c68d7cd7f18bb05daefce788e6f49b942766a1
---

# is-my-routerboard-rb750r2-bricked-no-response-from-router-for-netinstall

I have tried doing a netinstall as follows:

1. Press and hold reset button
2. Insert power cable.
3. Wait for flashing act light, continue to wait for On act light, continue to wait act light off.  then release reset button.

I Already have netinstall running on linux machine which has an up (once power on) ethernet connection configured as 192.168.88.10/24 DFG as .1

The netinstall just waits for RouterBoard.

Running wireshark on the used ethernet I can see the routerboard sending Spanning packets and CDP packets I can also see Linux host sending IPv6 and other solicitain messages but no communicaiton with netinstall.

I have tried all interfaces from the Internet one all the way along the board.

```
netinstall-7.16.2$sudo ./netinstall-cli -r -i enp198s0f4u1 routeros-mipsbe-6.49.13.npk 
Version: 7.16.2(2024-11-26 13:15:20)
Will reset to default config
Waiting for Link-UP on enp198s0f4u1
Using client IP 192.168.88.11
Waiting for RouterBOARD...
```

 
            
           
          
            
            
              You did use ether1 to connect your RB ?

Sometimes it helps to put a dumb switch in between.

             
            
           
          
            
            
              Yes.

eth1 being the one with the label interent on the box. As the others are then numbered 2 to 5.

I’ve also tried all others !

             
            
           
          
            
            
              
Try earlier versions of net install.

I think best choice is to use version of netinstall that corresponds to rb “current firmware”.

Helped me many times.

Do not release button until you see device in netinstall gui.

             
            
           
          
            
            
              
If you’re running wireshark, then it’s pretty safe to keep pressing the reset button on router *until you see BOOTP packets* arriving from router. Then you release the reset button. Sometimes observing light flashing pattern doesn’t give you the right information.

As far as the docs go, the only problem of holding button this long is that backup routerboot gets executed instead of the main one. But in my (limited) experience this is not really a problem, netinstall doesn’t touch routerboot anyways, it’s only necessary to get netinstall (or ROS) going. When device reboots after netinstall is done, it’ll start main routerboot (unless you play with reset button while rebooting).
