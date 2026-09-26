---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-yv3iew-capsman-and-wifiwave2-tutorial-needed-041975e5
title: "r-mikrotik-comments-yv3iew-capsman-and-wifiwave2-tutorial-needed-041975e5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/r-mikrotik-comments-yv3iew-capsman-and-wifiwave2-tutorial-needed-041975e5.md
source_anchor: ""
source_lines: [1, 137]
sha256: 626c81fa2366ca94cc6c7304535952d64b600a8a2f3eb63502c2d19088abdf36
---

# r-mikrotik-comments-yv3iew-capsman-and-wifiwave2-tutorial-needed-041975e5

CAPsMAN and WifiWave2 - Tutorial needed 
        
    Hey world,
      Could somebody be kind enough to spare some time and do a tutorial on how to do a CAPsMAN configuration for the WifiWave2 7.7 beta 6?
I've tried following the short guide from here but with no functional results.
    
I'm trying to setup a pair of Audiences, running 7.7 beta 6 with wifiwave2 but I have limited knowledge on RouterOS and setting up CAPsMAN on CLI is proving too hard for me.
Cheers!
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
7.7 Beta 6 adds initial support for CapsMan w/ WifiWave 2, as they stated here
However, from what I've read on the official thread, it is really in early dev. stages as CapsMan can't control local interface yet - see Guntis's answer
At least happy to see they are working on it!
With 7.9rc1 I managed to get the roaming to work.
My config (I own 2 Audience units but below I left the config for Wifi1+2):
And config for CAP (check it manually):
I got it working by moving all the security details from security profiles to the security fields within each configuration. Others have succeeded the same with security profiles listed as sources in each configuration.
Hi, it looks like a very basic configuration. Have you delved into anything more complex? I mean, like VLANs, radius, or specific frequency settings?
No, I'm using it for a simple home setup, no need for VLANs...yet. However I encourage you to use the forum - those guys have all the answers 🙂
Commentaire supprimé par le membre
I was using Ethernet backhaul but meanwhile I’ve switched to an Omada setup.
Has there been any updates or improvements in this topic?
This might be outdated info… but I didn’t think capsman was compatible with the wave2 package yet?
It is on the 7.7 beta now although CAPsMAN needs to be configured in the terminal
No capsman for wifi2 package at this time.
It is on the 7.7 beta now although CAPsMAN needs to be configured in the terminal
Great to hear they are working on it though
