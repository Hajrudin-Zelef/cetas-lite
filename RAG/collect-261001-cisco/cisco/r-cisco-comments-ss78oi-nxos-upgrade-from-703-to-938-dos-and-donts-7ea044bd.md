---
id: collect-261001-cisco/cisco/r-cisco-comments-ss78oi-nxos-upgrade-from-703-to-938-dos-and-donts-7ea044bd
title: "r-cisco-comments-ss78oi-nxos-upgrade-from-703-to-938-dos-and-donts-7ea044bd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "research"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-ss78oi-nxos-upgrade-from-703-to-938-dos-and-donts-7ea044bd.md
source_anchor: ""
source_lines: [1, 143]
sha256: cb5ce0158488e00152093cf521771789717de9b9ba1f4a6792933a38292e0ade
---

# r-cisco-comments-ss78oi-nxos-upgrade-from-703-to-938-dos-and-donts-7ea044bd

NX-OS Upgrade from 7.0(3) to 9.3.8 - do's and don'ts 
        
        
        
    
    
    Hi r/Cisco!
We´re planning an upgrade of our 4x Nexus 93108TC-FX switch vPC cluster from NX-OS 7.0(3) to the recommended version, 9.3.8 at the moment. I´ll check the release notes, upgrade guides and do some research but I just wanted to check with you to see if anyone had any hints or tips they wish they knew when they upgraded?
Thank you!
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Those are all good things to do.
Those are all things that you should do.
Asking your social media friends for input & gotchas is also wise.
But, if this is your first time making an upgrade this large, considering how much money you are paying for support on those devices, why not also open a ticket with TAC and ask them for guidance, and if you want it, their real-time assistance too?
There is no reason to go into this cold and alone.
Thank you for your input! I updated several Nexus 6000 switches in the past and all went kind of well. I heard some things about bricked Nexus 3000 switches after an upgrade so I thought it would be an idea to ask here.
I´ll open a ticket with TAC, you´re totally right. Got to change my attitude that we are paying money for support and that we can actually use this possibility in a proactive way :-)
I also assume you've already seen this:
https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus9000/sw/recommended_release/b_Minimum_and_Recommended_Cisco_NX-OS_Releases_for_Cisco_Nexus_9000_Series_Switches.html
Some of the N3k were an oddity. Cisco mentions in the support notes that upgrading some N3k models using the reload command is not supported. I can easily see that some people didn't know this and simply copied the new NX-OS to the flash:, set the boot statement and then rebooted the switch.
Only thing I can say is follow the upgrade path tool that Sk1tza linked, and use the install-all command, we ran into some nasty bugs that cisco wouldn't help us troubleshoot until we installed the latest version with that method.
It was our first time ever using NX-OS and it was a massive headache.
Thank you!
https://www.cisco.com/c/dam/en/us/td/docs/dcn/tools/nexus-9k3k-issu-matrix/index.html
Have you had any positive experience with issu? I have tried multiple times in our Nexus platform from the 7k, 7700, 9Ks (NX-os mode) and never had a positive experience. Always resorted to changing the boot command and reloading.
Was more to highlight the upgrade paths you need to take. You don't say which 7.x you're running but youre going to need a couple of jumps id say.
super - useful tool!
Funny you post this. I'm planning on doing the same jump on my 93108TC-EX soon.
I just put a 9.3.3 into my GNS3 to upgrade one of my nodes there and see what happens with VXLAN.
Our 9ks were on the ISSU compatible list, but running a show install proved that to be false. Each one had a BIOS upgrade or something else that required a reboot. All in all no problems with the code in different data centers handling different workloads.
we got strange issue upgrading from NX-OS 7.0(3)I4(7) to 9.3.9.
The OSPF network area command started to not work. We needed to workaround this by putting ip router ospf command under the interface.
