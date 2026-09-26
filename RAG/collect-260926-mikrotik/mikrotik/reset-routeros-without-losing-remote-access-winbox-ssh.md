---
id: collect-260926-mikrotik/mikrotik/reset-routeros-without-losing-remote-access-winbox-ssh
title: "reset-routeros-without-losing-remote-access-winbox-ssh"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/reset-routeros-without-losing-remote-access-winbox-ssh.md
source_anchor: ""
source_lines: [1, 228]
sha256: 633e35f1fe25d44c6b2a5fd1afe3c4596197592445a41f91e61590f2f28175bb
---

# reset-routeros-without-losing-remote-access-winbox-ssh

How can I reset RouterOS to factory defaults without losing remote access (Winbox/SSH)?

             
            
           
          
            
            
              How can you eat an apple but keep it intact ?

You can not.

             
            
           
          
            
            
              
I disagree, a whale can swallow it whole… and then regurgitate it back whole.

             
            
           
          
            
            
              Maybe with customisation/branding tools you can insert your own “default” script and create a remote access to the device.

Like explained here https://help.mikrotik.com/docs/spaces/ROS/pages/69664784/Branding

             
            
           
          
            
            
              
As you mention “losing” access, it implies you still have it. If so, you should be able reset the machine to default configuration using a command rather than by power cycling it and pressing the reset button, which means you could prepare in advance a configuration script with the required configuration allowing remote access and use the */system/reset-configuration* command with the *run-after-reset* parameter to run the script at the first boot after the configuration reset (look also at the *keep-users* and *no-defaults* parameters). But there are some surprises (you have to delay execution of the script for some time as e.g. it can be spawned before all the network interfaces have been detected and initialized), so it is very useful to debug it first using some locally accessible Mikrotik device, like e.g. a CHR running on any virtualisation platform you have at hand. You will have just one attempt in the real case.

             
            
           
          
            
            
              
I would think a whale eating an apple is really close to shore and will most likely end up being grounded  

My underlying meaning with my response was that it can not be done out of the box without jumping through some hoops and careful planning/testing.

Which gets backed up by the other responses.

             
            
           
          
            
            
              First of all, why do you need to do this ??

Workaround:

Insert a script and choose Run After Reset.

             
            
           
          
            
            
              
would suggest this approach too.

test this on a local MT first and examine if the default configuration can be altered as far as you need to access it remotely.

e.g.

setup a firewall input allow rule on top **specific to your remote accessing location**

OR

configure a “default” wireguard tunnel for “quasi zero configured” routers

…this could even be fleshed out to your own  “SD-WAN” or better, “(semi)zero touch deploy” solution (;

             
            
           
          
            
            
              It’s quicker to do a /export show-sensitive and then delete or reset to default everything that is not needed, it’s safer…

If instead it is done to try to clean something “dirty”, netinstall should be used…

             
            
           
          
            
            
              You can’t use netinstall remote.

             
            
           
          
            
            
              Maybe the meaning is “you can’t use netinstall to install a remote device”, if so.

Otherwise it seems “You can’t use *the* netinstall remote *control*”…

FALSE: if you have the right devices available and connected properly.

It is not the first time I do netinstall remotely from other RouterBOARD…

anyway, yours is useless note, since my comment was about whether he was trying to clean the router from something “dirty”…

It is not enough to reset the RouterBOARD to factory default.

However, by not explaining the reason, he doesn’t get the right hint.

             
            
           
          
            
            
              Why does this bring back memories of this recent thread?  MikroTik RB5009 setting up remotely first time

             
            
           
          
            
            
              
Can you explain what your real goal is?

             
            
           
          
            
            
              He probably demanded an immediate answer… and he just go away…

             
            
           
          
            
            
              
In this case you can:




/system reset-configuration keep-users=yes no-defaults=yes skip-backup=yes run-after-reset=


Make sure that this file defines an IP-address for ether1 (or whatever you use for remote access) including default routing and what else thats needed.

Also since Mikrotik changed default password to one put on the sticker at the bottom of the unit you might want to keep the current users with `keep-users=yes` otherwise to completely start over I would suggest to use `keep-users=no`.

             
            
           
          
            
            
              
The real goal is probably to reset RouterOS to factory defaults without losing remote access (Winbox/SSH)…

             
            
           
          
            
            
              You getting realize what the difference between “home-grade” vs “telco-grade” hardware.

Thats the reason why professional equipment has dedicated OOB port ( OpenBMC, Redfish, IPMI, … ).

You cannot do this without OOB (Out-of-band management) interface which has separate data and control plane.

             
            
           
          
            
            
              Hogwash, of course.  When you have a dedicated OOB port, you need remote access to that anyway.

And when that is a prerequisite, you can do it with MikroTik equipment too.  Some routers have RS232, you can cross-connect that between two routers and access one router from the other, even after factory reset.

Similar options are possible when the routers are connected via an ethernet port (there is MAC-level access), you can even netinstall one router from the other when you have interconnected the correct port.

That it is possible with “professional equipment” is only true when there is a professional admin who understands what has to be prepared.

             
            
           
          
            
            
              This can be an expensive but sizeoptimized solution to do so 

https://www.flexoptix.net/en/t-c12-rs232i.html

Other options are of course to get a proper SCS (serial console server) with dual LAN interfaces so you can go from IP to serial remotely.

             
            
           
          
            
            
              If you have other mikrotiks or any the same ethernet ports as the router in question.

(And assuming when you reset the mikrotik it doesn’t disconnect your access to those other mikrotiks)

You will likely be able to login to it from those other mikrotiks using mac telnet.

It seems there is also a mac telnet client for linux.
