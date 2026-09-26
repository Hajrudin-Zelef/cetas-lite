---
id: collect-260926-mikrotik/mikrotik/new-rb951g-2hnd-bricked-solved-with-netinstall-2
title: "new-rb951g-2hnd-bricked-solved-with-netinstall"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/tools/new-rb951g-2hnd-bricked-solved-with-netinstall.md
source_anchor: ""
source_lines: [190, 304]
sha256: 0627c7123899e8486ec5526d2fd90d716da2862ac9139e40b8a9a38bf240f84d
---

# new-rb951g-2hnd-bricked-solved-with-netinstall

I had one variation from Pradeep’s list that I wanted to note in case someone else runs across this thread. I got two new units yesterday, one of which was bricked out of the box.

After following steps E-Q above multiple times, I had varying levels of “almost” success. Eventually, the router started acting like it was successfully rebooting, but instead of the single beep noted in Q, I was getting a quick double-beep. This happened both after the netinstall procedure and if just rebooting without anything connected. The wireless light never turned on.

I noted that step L didn’t mention the “Default configuration” option, which was unchecked, so I enabled that. (None of the screenshots I’ve looked up show that option, so I’m guessing it’s a recent addition). The router rebooted successfully, the wireless light came on, and I was able to see the configuration web page.

Thanks to everyone in this thread for the help!

             
            
           
          
            
            
              Just to add - NEVER select “ALL” packages on RouterBOARD devices. Many of them have little space available, better check only the packages you know you need. Usually for basic operation you need system, advanced tools, wireless - maybe something specific, but that’s it. routing package is only for dynamic routing, and ntp is only for server, so you mostly don’t need to install these.

             
            
           
          
            
            
              my sxt 2hnd is bricked…after powering the user led and power button glows non stop along with two other led’s…is there a way to fix it or its damaged?

             
            
           
          
            
            
              
I had to do this process too. Took me about 2-3 attempts to get the timing right and have the unit show up. It also helps to know which interface you need to be hooked upto depending on which model you might be doing netinstall on. But process certainly works well.

             
            
           
          
            
            
              
I believe you have a wiki article that talks about this. I should probably go find and link it here.

             
            
           
          
            
            
              I Guess you were referring to this link?

http://wiki.mikrotik.com/wiki/Manual:System/Packages

Thanks for the advice.

It worked for me wonderfully.

The only difference was that the unit got bricked by installing a backup that had features enabled not support by the RouterOS version, as the OS was not upgraded first. The initial backup was created in RouterOS 6.29.1 and the device had 6.23 on at the time.

I selected the standard 3, advanced-tools, system & wireless, as prescribed by Normis.

Also I had to select -keep old configuration.

For some reason it refused to install with deleting the old config.

This was on a Mikrotik BaseBox

             
            
           
          
            
            
              I just received MikroTik RB951Ui-2HnD from one of my friend to setup the hotspot with it.

Well during my experiments (which included firmware upgrade to 6.33.5) somehow I bricked it.

Google search landed me on this page. I tried everything mentioned here but my modem never picked up by netinstall.exe

After some more search on net I finally succeeded.

I thought it would be interesting to share the whole procedure in following simple steps:

**How to unbrick MikroTik RB951Ui-2HnD**

You must first have downloaded these:

1. netinstall.exe
2. and the required firmware (eg. routeros-mipsbe-6.33.5.npk) from MicroTik website.

Of course your modem must be hooked up with PC through LAN cable.

Use ‘port 1’ (labelled PoE In) on your modem.

1. Disable your PC’s Firewall and Antivirus.
2. Disable all ethernet adapters (eg. VMWare, Wifi etc)
3. But left enable only ONE ethernet adapter which we are using to connect the modem with the PC.
4. Set the adapter’s IP to: 10.0.0.1 and subnet mask to 255.255.255.0  (Leave rest of the settings blank)
5. Run netinstall.exe
6. In Net Booting option: Check Enable Boot Server & set client IP as: 10.0.0.2
7. Now boot your modem while pressing the reset button for at least 15 seconds (until activity indicator turns off) then let it go. This will reboot your modem (with a single beep after a few seconds) with netinstall search.
8. The MAC address will appear in netinstall. Select it.
9. Select your firmware (which you must have already downloaded).
10. Check Apply default config.
11. Hit install and after a while (a few seconds) press reboot (if appears). Done!
12. Don’t forget to revert all of the adapter changes you made in previous steps (2-4)

This is how I got my modem back to work. 

Hope it will help you.

Thanks

Dr. Zeeshan

25 Jan 2016
