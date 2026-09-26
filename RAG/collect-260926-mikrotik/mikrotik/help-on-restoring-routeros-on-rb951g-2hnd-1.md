---
id: collect-260926-mikrotik/mikrotik/help-on-restoring-routeros-on-rb951g-2hnd-1
title: "help-on-restoring-routeros-on-rb951g-2hnd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/help-on-restoring-routeros-on-rb951g-2hnd.md
source_anchor: ""
source_lines: [1, 131]
sha256: 6f5babef2f6b42b1222bd6eb0e182e2b617421ec6994f7db33d987950f488669
---

# help-on-restoring-routeros-on-rb951g-2hnd

Hello Every Body,

I have “RB951G-2HnD” that I have been using for many years. Couple of months ago I though it’ll be a good idea to test a custom firmware on it “OpenWRT”. So I put my mind on it and said why not give it a try and check how it goes. So I dug up for a while on the OpenWRT and the compatibility with my RouterBoard and I found out that it’s compatible, so I started reading on how to install the OpenWRT on my router board and following the instructions on their website I managed to install it on my RouterBoard (After Backing up my license). Started to work on the OpenWRT and configuring my device and give it a shot, saying that I can always revert back to RouterOS once I’m done exploring the OpenWRT, had no intention to stick with OpenWRT, but just curious about it. So I’ve done the configuration and the walk around and put it to test as my home’s main router. Sadly the router started to act out and rebooting every now and then, I told myself there must be something wrong with configuration. Hence, I did a reset to the router and reconfigured it from the beginning hoping that the issue will be resolved and I can do more work on OpenWRT and gain some experience along the way, but the rebooting issue remained. Therefore, I decided to try the OpenWRT on the Raspberry Pi and put back RouterOS to my RouterBoard. I never had any issues/complains with the RouterBoard before and It was working fine until I tried out OpenWRT. I’ve used several MikroTik products for years and still are. Every time I do something that damages the router I’d just do a netinstall and start fresh with no issues (Done it “NetInstall” via serial console cable and Ethernet cable if the device doesn’t have a serial console interface on it plenty of times with no issues that can’t be managed like OS, firewall or antivirus restrictions).

Back to my “RB951G-2HnD”, I tried to restore the RouterOS on it via Netinstall but the device won’t show up. So I tried again and again and again with no success. I thought I’m doing something wrong here let me read the Wiki and check I might be missing something, but all what i did was just right and as the instruction said in the Wiki but the device won’t show up on NetInstall. So I started finding for a solution:

1. Disabling the firewall and the antivirus … No Success.
2. Trying different OS (Win XP, Win 7, Win 8.1, Win 10) … No Success.
3. Enabling/Disabling the Netboot in NetInstall … No Success.
4. Trying different NIC … No Success.
5. Use a switch in between the router and the PC … No Success.
6. Pushing the Reset button for different amount of times before and after powering up the device … No Success.
7. Followed all the possible solutions mentioned in the threads posted by others on different and similar device as the one I have … No Success.
8. Trying different versions of NetInstall going back to NetInstall V3 … No Success.
9. Flash OpenWRT back and tried NetInstall again … No Success.
10. Trying to change the boot order within OpenWRT … No Success.
11. Trying some packages on OpenWRT to help me change the boot order … No Success.
12. Searching for an Image of the device that I can put back on the RouterBoard via TFTP … No Success.
and many many more with no luck

The device is currently running OpenWRT but it’s not stable, even after upgrading and downgrading, I can’t get it to stop acting out on OpenWRT and stop it from rebooting every now and then.

I’ve ran out of things to test in order for me to put back RouterOS into my RouterBoard again. Is there a way that I can try and connect to it via console cable ???

Any help/suggestions/ideas from anyone ?

Sorry for the long elaboration, just wanted to be clear on how things went from the beginning till now.

Thank you all

             
            
           
          
            
            
              Anybody, Ideas, suggestions, thoughts… I’m stuck here   

             
            
           
          
            
            
              You must using ether1 when netinstall, see all description https://wiki.mikrotik.com/wiki/Manual:Netinstall

             
            
           
          
            
            
              Thank you for your reply.

Yes, I’m aware of that. Most if not all RouterBoards that doesn’t have serial console port uses ether1 port for the NetInstall. Tried it many times but with no success.

             
            
           
          
            
            
              Please help I’m open to all and any suggestions

             
            
           
          
            
            
              Do you perhaps know if OpenWRT replaces RouterBoot firmware (the BIOS-like firmware which wakes up device and boots actual OS) when installed? If it does, then you’ll have to take care about that before netinstalling the unit.

             
            
           
          
            
            
              
I believe it did replace it “The Bootloader”, since it’s part of the installation process https://openwrt.org/docs/techref/bootloader. But I haven’t modify or done any sort of adjustment to it, so I assume by default and following the installation procedure of OpenWRT that it’ll change the bootloader in order for OpenWRT to be installed. I had no idea how to access the bootloader or modify it if it’s even possible.

Following the tutorial on OpenWRT website they said it’s possible to revert back to RouterOS via NetInstall but following the instructions on how to use NetInstall on their “OpenWRT” website and MikroTik website didn’t help at all and I the device never showed up in NetInstall no matter what I did.

This is the link for MikroTik products common procedure on installing OpenWRT and for the requirements for the revert back process:

https://openwrt.org/toh/mikrotik/common

             
            
           
          
            
            
              
I’m going to take a look at it, I have a problem with them too, thanks.

             
            
           
          
            
            
              Read the manual https://wiki.mikrotik.com/wiki/Manual:RouterBOOT and try activate backup boot loader as:




RouterBOARD reset button

RouterBOOT reset button has three functions:

Hold this button during boot time until LED light starts flashing, release the button to reset RouterOS configuration (total 5 seconds)

Keep holding for 5 more seconds, LED turns solid, release now to turn on CAPs mode (total 10 seconds)

Or Keep holding the button for 5 more seconds until until LED turns off, then release it to make the RouterBOARD look for Netinstall servers (total 15 seconds)

Note: If you hold the button before applying power, backup RouterBOOT will be used in addition to all the above actions. To do the above actions without loading the backup loader, push the button right after applying power to the device.


Try using USB-TTL UART adapter for connect to UART pad on board https://s00.yaplakal.com/pics/pics_original/8/1/3/3969318.jpg, find unsoldered contact pads GND,RX,TX near the bottom right corner of the picture with UART title. After connect pads to USB-TTL adapter, run terminal program like “putty”, select “Serial” and desired port for USB-TTL adapter, then see activities on console during boot RB951G-2HnD (see https://www.trishtech.com/2015/01/testing-usb-to-ttl-adapter-using-putty)

             
            
           
          
            
            
              
I have USB to Serial Converter, Can I use that to connect to the pads in the board ? Shall I connect to the 4 “Square shaped” pads on the board or what ? Following the manual on using the serial console cable there are 9 pins based on the cable and as shown on the wiki here https://wiki.mikrotik.com/wiki/Manual:System/Serial_Console so which is it that I need to use

