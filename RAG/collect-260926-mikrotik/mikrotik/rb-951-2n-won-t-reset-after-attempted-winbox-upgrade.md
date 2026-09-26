---
id: collect-260926-mikrotik/mikrotik/rb-951-2n-won-t-reset-after-attempted-winbox-upgrade
title: "rb-951-2n-won-t-reset-after-attempted-winbox-upgrade"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/tools/rb-951-2n-won-t-reset-after-attempted-winbox-upgrade.md
source_anchor: ""
source_lines: [1, 167]
sha256: cca64b0554fce55d46223cf15917737db360d34f2ec86b5f65ff465f67993a70
---

# rb-951-2n-won-t-reset-after-attempted-winbox-upgrade

I have an RB 951-2n which was set up as a station bridge, client to another one used as an AP.  The AP was upgraded to 5.20 successfully via Winbox.  So I took the station side box to the vicinity of the AP for a good connection and then Winbox would connect to it.  I dragged the 5.20 files into the box and waited until the transfer finished.  Then I rebooted.  It didn’t come up.

So I have tried to reset it, using both the front button and bottom shorting pin.  It doesn’t.  I’ve never done that procedure before so I’m not sure about the timing, but I tried a few times, and used an old old laptop as a test jig whose IP didn’t matter.  So I tried it as 192.168.88.2, and as its old address (the real subnet here).  I think I’m getting the home page with the login screen, though that *might* be the browser confused by the cache - when I used “admin” to log in it goes back to the same screen.  And the LAN on the laptop keeps telling me that it’s disconnecting and reconnecting.

Any clues about how to really hard reset this thing, or otherwise recover?  Thanks.

             
            
           
          
          
            
            
              http://routerboard.com/pdf/413/rb951-2n-qg.pdf

Read the last sentence under under “Booting process”.  Then run Netinstall.

Tom

             
            
           
          
            
            
              Netinstall is a good idea.  I tried it.  I connected the spooged 951 to an Ethernet port on the working one (the AP, now at 5.20).  I ran Netinstall on a laptop hard-wired to the same working 951 and gave it the client bridge’s desired IP address.  When I put the Ethernet cable into jack 2 (not 1) of the spooged 951, Netinstall saw it.  I then pointed Netinstall at the 5.20 files on the PC and it downloaded them all, then rebooted.

However, it still doesn’t actually seem to be working.  It doesn’t respond to its given IP address, or anything else. I wonder if I got netinstall wrong.

             
            
           
          
            
            
              Power cycle the device after netinstall and use winbox discovery button “…” to connect to the device via MAC address. If winbox does not discover it, I would probably netinstall one more time.

             
            
           
          
            
            
              I did a second Netinstall and a few more reboots and resets, and finally it came up at 192.168.88.1. So now I just have to redo the incantations to make it a station bridge.

What I really wanted to do, what started me down this path, was make it simultaneously a station bridge (to put an Ethernet-equipped desktop machine onto a wireless LAN) and to make it a WDS repeater (to allow wireless devices near it to have a stronger signal).  I have not yet found a way to make it do both at once; the documentation (Wiki, etc.) talk about WDS but not when it’s also a bridge.  Is that possible?  Thanks.

             
            
           
          
            
            
              Just found this topic by google. I almost sent back my router to the seller. After used netinstall I got an access to my router again. Everything took like 30 seconds. Thanks 

             
            
           
          
            
            
              I also find this topic using Google. I had some troubles trying to reset 951-2n board as described in the official guide without any success. Using the “Trial and error” method I come up with this solution:

1. Disconnect everything from the device
2. Press and hold the reset button, in the same time connect the power cable into device
3. Wait a few seconds until the “ACT” led start flashing, release the reset button (be aware do not hold it until the led stop flashing)
4. Wait a few minutes, just leave it in power on state for 3-4 minutes (do not connect any RJ45 cables to it)
5. Configure your PC to the network 192.168.88.*, set the IP address let say to 192.168.88.2
6. In the console start to ping into 192.168.88.1 : ping 192.168.88.1
7. Using RJ45 cable try to connect into device ports one by one starting with eth1, see the console until you’ll get replies to ping
8. Your device is ready !

 
            
           
          
            
            
              Hello

I also have a RB951-2n, and after I tried to reset, it stopped responding. I’ve tried using the netinstall, but I did not succeed in gaining access. When I press and hold the reset button until the “ACT” led start flashing, I think it  flashes slower tan I expected, in periods of more than 1 second.

Currently it shows the following behavior:

- When it starts, it makes a strange noise, diferent from the noise it made when it worked well. You can see a video I made here:  https://www.youtube.com/watch?v=eihxiLai0z0
- After starting, the wireless is never active.
- When I connect via RJ45, I can not reach it by DHCP. When I use Fixed IP (192.168.88.1 or other), my PC (running linux or windows) accepts the connection, respond to ping, but can not access it via webfig or WinBox. I also tried in all possible ports. I think I’ve followed all the tutorials available on this forum.

The product warranty is over and I don’t know what else to do. I appreciate the help you could give me on this issue.

regards, Artur

             
            
           
          
            
            
              Hi Arthur,

Instead of trying to connect by IP using winbox, try to connect by MAC.

In winbox, use the ellipsis […] button to bring down auto-discovered MT hosts on the network, then click on the MAC address in the list so it is in the ‘Connect To:’ field, then try connect.

May or may not work, but will bypass any IP firewalls that may be blocking you currently.

Given this topic is 1+ year old you should probably have started a new thread also 

             
            
           
          
            
            
              If winbox by mac didn’t work you just need to netinstall.

Sent from my SCH-I545 using Tapatalk

             
            
           
          
            
            
              I appreciate your help, but I think my RB951-2n is death.

I tried all possible ways:

- I tried to reset with jumper hole and reset button
- I tried winbox by mac, with different configurations
- tried this: http://wiki.mikrotik.com/wiki/Manual:Winbox
- and this: http://wiki.mikrotik.com/wiki/Initial_MAC_Winbox_Connection
- I tried in 3 different computers running win7 e vista
 Winbox never managed to find my RB. And what I said in last post about ping was wrong. I don’t even know what is the correct IP.
I also tried Netinstal, but I couldn’t manage to put it on. When I hold this reset button during boot, LED never turns off. It always ends up restarting.

In my opinion, there is something odd about the noise it makes when it starts:

https://www.youtube.com/watch?v=eihxiLai0z0

I am thinking of calling the warranty, but I’ll probably pay too much for postage.

Anyone else have an idea?

Artur

             
            
           
          
            
            
              That noise you are hearing - from my experience indicates to me its just done a default config load or a reset load. Do you still have the reset pins jumpered or anything like that?

Many of my RouterOS devices have made that noise it is nothing to worry about per-se. Check the buttons are not stuck in and check the reset pins on the board are not still jumpered.

Try using the netinstall utility.

             
            
           
          
            
            
              I finally managed to access the router by netinstall, but unfortunately I did not quite understand what I did differently. It was very strange. Now it works well and makes a normal noise at startup. I thank everyone who helped me with posts.

Artur
