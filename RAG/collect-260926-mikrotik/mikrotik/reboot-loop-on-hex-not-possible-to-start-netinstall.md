---
id: collect-260926-mikrotik/mikrotik/reboot-loop-on-hex-not-possible-to-start-netinstall
title: "reboot-loop-on-hex-not-possible-to-start-netinstall"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/tools/reboot-loop-on-hex-not-possible-to-start-netinstall.md
source_anchor: ""
source_lines: [1, 76]
sha256: caa7b520f2a25cca969b5bcec0838e5b6fe454620018336d2cd8f31aa76a4783
---

# reboot-loop-on-hex-not-possible-to-start-netinstall

I have one weird hEX which failed during upgrade, however I am not able to put it netinstall mode.

What happens is that the device just restarts at one point while I am holding the reset (at about 15-20s?).

I presume the device is irepairable, is there anything else I could try?

             
            
           
          
            
            
              The hope is not over.

There is lot you can do.

If you using windows, there lot’s of problems being using that.

If you are knowing how to use your comporter with another OS, you should try:

To write usb flash stick with some Linux Live cd.

Then booting this up and then download netinstall for linux and try that.

In the best worlds even Mikrotik should provide a ISO file, with like say Bootable Alpine Linux with the latest netinstall.

My self struggling a lot, to get netinstall to work in linux, because of I setting the device to boot via dhcp instead of bootp.

You can read about my journey here and how i solved the problem:

http://forum.mikrotik.com/t/trying-to-unbrick-my-rg750gl/165209/1

             
            
           
          
            
            
              There are two (BTW confusing) ways to boot in Routerboot mode:

1. press the reset button and while it is pressed power the device, then keep pressed until the device appears in netinstall ← this will use the backup bootloader
2. power the device and immediately after press the reset button, then keep pressed until the device appears in netinstall ← this will use the normal bootloader

Can you try with another power supply? It could also be that a peak of needed current is needed at a given point of booting and a (defective) power supply just cannot provide it.

             
            
           
          
            
            
              
I didn’t know this… I will take the device apart, maybe the button is damaged. Maybe I couls short the pins together. Will try.





You mean boot PC with linux or the rotuerOS? I would actually like to see linux running on routerOS… 

Windows are fine, did many installs with this machine. Thanks though.

             
            
           
          
          
            
            
              Well, I fixed it. The router had the button stuck. Managed to put it in netinstall mode and it worked. HOWEVER - the router behaves in a weird way. The normal bridging didn’t work after I got it up, and I just needed it as a simple switch. So I put it away for now and ordered a 260GS. Too much trouble for an old router.

Anyways, thanks for help guys.

As for Linux - I alway used netinstall over Windows. The issue is if the Win Firewall is active. That thing blocks communications.

The serial console stuff worked best on the Mac though.
