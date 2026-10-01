---
id: collect-261001-meraki/meraki/r-meraki-comments-pk0l2c-load-balancing-and-activeactive-auto-vpn-aa070d3c
title: "r-meraki-comments-pk0l2c-load-balancing-and-activeactive-auto-vpn-aa070d3c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-pk0l2c-load-balancing-and-activeactive-auto-vpn-aa070d3c.md
source_anchor: ""
source_lines: [1, 122]
sha256: 203f36b0e6c3797156b7677aa3c944a88b47c4d93066cf17e60bbc6fc2db601b
---

# r-meraki-comments-pk0l2c-load-balancing-and-activeactive-auto-vpn-aa070d3c

Load balancing and active-active auto VPN. 
        
    I have enabled active-active autoVPN on my hub site but not load balancing since its not actually handling any of the "internet" traffic. The remote sites use their own circuits for that. Is there any benefit to enabling load balancing on my hub device?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
All my experiences with load balancing on meraki were a huge nightmare and most voip products completely fall apart.
What? I have a large org with autoVPN load balancing enabled everywhere and everything works just fine. We have a combination of internal SIP trunking, Cisco Skinny phones, MS Teams voice, Webex Teams voice and Skype for Business voice. It all works everywhere. 40+ sites, 12,000 employes, 4 geo-diverse Call Manager clusters and no problems.
Not sure what to tell you. We were told to turn it off by our voice partner because calls would take 10 to 15 seconds to establish the media stream and it was pissing our customers off. Audio would pick up half way through our customer service reps giving their greeting. Tried lots of stuff to make it work right over a few months. As soon as we disabled load balancing all the problems went away.
I also have lots of failover issues with our MX67c devices. Hardline to cellular can take up to 2 minutes to happen. We lose our POS, phones, etc. Currently piloting fortigates to stop the problems.
active-active VPN just means you are bringing up autoVPN tunnels using WAN 1 and WAN 2 interfaces at the same time. If you don't turn that on, the MX will only use one interface for autoVPN tunnels, then failover to WAN 2 if WAN 1 can no longer reach Dashboard.
Load balancing has to do with Internet egress through the MX. If you have users sitting behind the MX, load balancing will send Internet-bound user traffic across both WAN interfaces. This has nothing to do with autoVPN traffic.
Depends which setting you're referring to. Some settings talk to load balancing the vpn traffic over the 2 internet circuits and some speak to internet. The one that's specifically talking about internet load balancing is useless if you're not sending internet through the appliance. Just remember that setting talks to any internet being sent through it not just over the vpn. If you're routing internet for that location through the mx it is effected by that rule too.
