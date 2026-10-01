---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-6zyqod-fast-roaming-using-unifi-ap-ac-pros-edgerouter-x-8250c010
title: "r-ubiquiti-comments-6zyqod-fast-roaming-using-unifi-ap-ac-pros-edgerouter-x-8250c010"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-6zyqod-fast-roaming-using-unifi-ap-ac-pros-edgerouter-x-8250c010.md
source_anchor: ""
source_lines: [1, 137]
sha256: e496f6fffd0800acd8278156d243912bb9bd2c19df164d17cb091be13cf88786
---

# r-ubiquiti-comments-6zyqod-fast-roaming-using-unifi-ap-ac-pros-edgerouter-x-8250c010

[supprimé]
      Fast Roaming using UniFi AP AC PRO's, EdgeRouter X and EdgeSwitch 24 
        
     
     Publication archivée. Impossible de voter et de publier de nouveaux commentaires.  
        
        
        
          
        
        
        
        
         Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
https://help.ubnt.com/hc/en-us/articles/115004662107-UniFi-Fast-Roaming
The tldr of it at the moment is that it works and you don't need to do anything. In the future when they implement more industry standards (802.11r/k/v) then that may change.
deleted 0.9665 What is ^^^this?
Really not sure, they've been cagey about what exactly fast roaming does and how, as it's not quite industry standard and they want to keep the details secret as much as possible. I would guess it requires the controller for configuration even if it's not full time in the loop...
Either way, if you're going to use unifi APs, IMHO you really want a controller running. Can be on a cloud key, a raspberry pi, a PC that's running most of the time, or even a free AWS instance. If you're not going to run the controller, you're probably better off with a different AP.
Edit: To answer the other part of your question, any switch or router you choose is fine and won't affect fast roaming one way or another. The APs need to be able to talk to each other and perhaps the controller, but nothing else in the system matters from a fast roaming perspective.
deleted 0.8245 What is ^^^this?
Just curious, why do you feel it is necessary to have the controller running 24/7? We have two AC Pro APs running through an Edgerouter Lite and they appear to work the same /including roaming) regardless of whether or not the controller is active.
For most things, doesn't have to be 24/7, though you do get useful statistics and logging if you do. What I was talking about is mostly the fact that it is possible to set the APs up with the phone app and never use a controller at all, and IMHO if you're going to do that, you're probably better served with a different solution.
Don't all the APs need to be on the same channel too?
No, that was the zero handoff setup, which is depreciated. Zero handoff pretended to be one virtual AP, while fast roaming just reduces the round trips a client needs to make to authenticate to a new AP on any channel.
deleted 0.5959 What is ^^^this?
