---
id: collect-260926-mikrotik/mikrotik/netinstall-cli-hanging
title: "netinstall-cli-hanging"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2025-02-24"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/tools/netinstall-cli-hanging.md
source_anchor: ""
source_lines: [1, 125]
sha256: 9bd238b9b812163750a7d1749158587b89d74a133d5ff95632ad4540df312e84
---

# netinstall-cli-hanging

I’m new to MikroTik products, but I’m trying to do a clean install of routerOS via the netinstall-cli on my E50UG (hEX refresh) because I bought it second hand, but it keeps hanging. I’ve been following this guide from the mikrotik docs and also referencing this youtube video by mikrotik.

Im using the right port on the router for etherBOOT (port 1), Im setting the IP correctly on my laptop (verified via `ip -br -c a` and WinBox that they’re in the same subnet), Im pretty sure Im using the right routerOS architecture (I checked via `/system/resource/print` before downloading the npk) and there are no switches or ethernet to USB adapters between the router and my laptop. I’m able to connect to the router via the netinstall-cli, but it hangs near the end and I cant figure out why:

```
$ sudo ./netinstall-cli -e -a 192.168.88.3 routeros-7.18-arm.npk
Version: 7.18(2025-02-24 09:55:03)
Will apply empty config
Using interface enp0s25
Using interface enp0s25
Waiting for Link-UP on enp0s25
Waiting for RouterBOARD...
Assigned 192.168.88.3 to F4:1E:57:9D:E7:98
Booting device F4:1E:57:9D:E7:98 into setup mode
Formatting device F4:1E:57:9D:E7:98
Sending packages to device F4:1E:57:9D:E7:98
Packages sent to device F4:1E:57:9D:E7:98
Rebooting device F4:1E:57:9D:E7:98
Successfully finished installing the device with MAC address F4:1E:57:9D:E7:98
Unknown BOOTP architecture option Flashboot from F4:1E:57:9D:E7:98
... hangs here forever
```

I assume this has something to do with the System->RouterBOARD->Settings->Boot Device, but Im not certain. It also appears like the install is successful because when I boot up WinBox, it lists the rOS version I installed via the netinstall-cli, I’m just not certain that it’s a completely clean install due to the cli hanging at the end, which was important to me because I bought this router second hand. I’ve tried using rOS 7.18, 7.17.2, 7.17 and 7.16.2 (which is the version it came with) and all hang with the same error message. My laptop is running Debian bookworm with all other network interfaces disabled.

I’m having a hard time finding any information on the error message `Unknown BOOTP architecture option Flashboot` and would love any help. Thanks so much in advance

             
                
            
           
          
            
            
              When you getting “Rebooting device F4:1E:57:9D:E7:98”, your device should be flashed ok.

And you can stop netinstall-cli with keys .

Maybe you forget to switch to Ethernet port to 2-5, and try to connect to it with winbox via mac address

or with the default ip address of 192.168.88.1(if you delete -e option with netinstall-cli) .

Or you could try to just flash it again without the -e option that flash the device with the default config.

With -e option you can only connect to the device with winbox. The device has no ip address at all.

             
            
           
          
            
            
              I tried flashing it with -e, with -r, without either of them, using -a  or-i , and regardless of the options I was using, it would always hang at the end. That said, each install seemed like it was successful and booted into the correct version of rOS that was flashed in that run.

It is reassuring to know that after the Rebooting and Successfully finished lines, everything should be done, but it is still weird to me that no matter what the command would hang, even though that wasn’t the case on all guides I saw online.

             
            
           
          
            
            
              Are you really sure it hang ? I used

```
./netinstall-cli -r -a 192.168.88.3 routeros-7.17.2-arm.npk
```

and it take very long time, over 30 minuts on a crs-326.

I aready gave up and have done some other works, just when i open back the terminal it was done.

Maybe you connect the console to the serial of switch, then you can see what is happening on client site.

             
            
           
          
            
            
              The new version of netinstall-cli doesn’t stop running after flashing a device, it’s just waiting to flash more devices.

With the new version on netinstall-cli you can provide option -o to avoid to flash a device twice i think.

Maybe you have to make a suggestion to the dev’s to add option like -1 to just flash one device and then stop.

             
            
           
          
            
            
              
That almost feels like the process died after that amount of time; its hard to believe the install process would take that long.

You’re right, though, I haven’t broken out wireshark or something yet to verify whether or not there are still packets going over the wire. I was hoping to avoid doing that but it might be my next step.

             
            
           
          
            
            
              
If that’s the case, I had no idea that was the default behavior. Do you happen to have a link to the docs, changelog or even commit that references that? I couldn’t seem to find anything mentioning it. But you’re right, I do think a `-1` flag could be a good idea, especially since I think that’s probably what the majority of home users would use; at the very least, a callout on the netinstall-cli docs referencing at what point flashing has finished and when its safe to .

             
            
           
          
            
            
              
I generally agree that clearer description in docs would be welcome. However, it’s (?) common knowledge that after device being netinstalled reboots, the process is over and done. Netinstall only takes care of flashing ROS anew, no initial configuration is done by netinstall. After reboot, ROS should be up&running. After that, one has to reconnect using winbox … or webfig/CLI if default config is applied which includes IP address.

             
            
           
          
            
            
              
Yea, realistically this might just be my ignorance on the matter showing through. I’m new to MikroTik, and it was my first time using the netinstall-cli; every reference I found showed it gracefully shutting down the server and returning cleanly after a single flash. I was expecting the netinstall-cli to behave more like a standard linux command-line program - runs, returns some result, and exits - and less like a server - stays online to serve requests until the process is killed - but that was probably a poor assumption.

Hopefully someone else finds this helpful and soothes their worries about the board not flashing completely
