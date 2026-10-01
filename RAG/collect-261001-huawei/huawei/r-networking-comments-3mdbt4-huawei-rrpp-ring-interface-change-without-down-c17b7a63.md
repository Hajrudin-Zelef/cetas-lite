---
id: collect-261001-huawei/huawei/r-networking-comments-3mdbt4-huawei-rrpp-ring-interface-change-without-down-c17b7a63
title: "r-networking-comments-3mdbt4-huawei-rrpp-ring-interface-change-without-down-c17b7a63"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-3mdbt4-huawei-rrpp-ring-interface-change-without-down-c17b7a63.md
source_anchor: ""
source_lines: [1, 124]
sha256: 5aa45aca8487016dd79fb0da9cefaae72e87ca43fb9c66b96543e6f151d480d8
---

# r-networking-comments-3mdbt4-huawei-rrpp-ring-interface-change-without-down-c17b7a63

Huawei RRPP ring interface change without down time or loop? 
        
    Hi everybody!
 I'm currently working on an RRPP ring owned by a carrier, who uses Huawei switches (S9300s and S5300s), with real client services on it.
 Well, if any of you never worked with Huawei, I envy you, because the only thing worst than their devices is their documentation.
 We have 3 switches that are part of this ring, and we need to upgrade the links between them, from a 1Gb link, to a 2Gb link using Eth-Trunk (adding one more interface to the one already in use).
 The problem with this is, we have the live RRPP ring there, and we can't stop the traffic. The ring is attached to this interfaces that have to be changed. Once we configure the Eth-Trunk and add the new interface, we have to change the RRPP attached interface to be the Eth-Trunk, and change the configuration of the old interface to be part of said Eth-Trunk, and I don't know what the correct workflow or work order is for these activities, we can't have any downtime and of course we have to avoid creating a loop.
 Did anyone here had to do something similar? I can't find even a tip list or a best-practice when regarding RRPP rings... Any advice would be tremendously appreciated.
Best regards, and may you avoid the pain of working with Huawei devices in your life.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Uhhhh! It's complicated, but equally it interests me to learn and know. Huawei device give us headaches!!! Someone tried to test and comment about your experience?
If it's a ring topology can't you just take one link at a time down, upgrade to 2Gb and bring it back up?
Well, that`s it!
