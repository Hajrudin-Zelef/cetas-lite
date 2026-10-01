---
id: collect-261001-meraki/meraki/r-meraki-comments-tyfwbu-ms-switch-stack-ip-issues-95c28fe8
title: "r-meraki-comments-tyfwbu-ms-switch-stack-ip-issues-95c28fe8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-tyfwbu-ms-switch-stack-ip-issues-95c28fe8.md
source_anchor: ""
source_lines: [1, 128]
sha256: 7754d00ea176478ecb281ced04f0f0e92922d561bfa85e234842fdfd6ef620c3
---

# r-meraki-comments-tyfwbu-ms-switch-stack-ip-issues-95c28fe8

[supprimé]
      MS Switch Stack IP Issues 
        
        
        
    
    
    Hi guys
We have two MS390 switches in a stack. We try changing their IPs on each individual switch but they both change to this IP instead of changing themselves to their respective one. I.e. if we change switch 1 in the stack to 10.1.1.10 the other will also change to that same IP. We’ve been this this behaviour before on other sites.
Any advise?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
MS390s are different than 'regular' Meraki switches as its actually a Catalyst 9300 running Meraki software. The stacking is the same as the Catalyst style though, so once they're stacked they're sharing a control plane and using the same IP
Great. Thanks for that info
Also, seeing very weird behaviour when we introduce devices to the network like a Meraki access point for example. All the network drops and then comes back up
We've found for those to just leave the default port config on first time boot. After they get to the internet and update/receive their config, then we change to the correct vlan, etc.
Otherwise we have issues with the initial config of say a switch or AP.
Thanks for the reply. We plug the switches straight into a FortiGate, some SVIs exist on the switch, we have one SVI / subnet that obtains it’s DHCP addresses on the fortigate. Out of our control, we would’ve had all interfaces on the fortigate but due to design and implement, we had to set the core Meraki stack up in with interfaces
