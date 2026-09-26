---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-uc4yzx-capsman-proper-configuration-bdc3f6f8
title: "r-mikrotik-comments-uc4yzx-capsman-proper-configuration-bdc3f6f8"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["sol"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/r-mikrotik-comments-uc4yzx-capsman-proper-configuration-bdc3f6f8.md
source_anchor: ""
source_lines: [1, 133]
sha256: 388c6b4546d125e544b23104d5b379d386bde468745c23190cb95fe5544e9cf2
---

# r-mikrotik-comments-uc4yzx-capsman-proper-configuration-bdc3f6f8

[supprimé]
      CapsMan proper configuration 
        
     
     Désolé, cette publication a été supprimée par son auteur.  
         
        
        
        
        
          
        
        
        
        
         Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
CAPsman certainly can. I have three networks on three vlans on my system. When you set up CAPsman, you set the vlans in the datapath section. This is used later in the configuration to set the SSID and other wireless info, this in turn goes into the provisioning where you create the main wifi and any other virtual ones as the slaves.
CAPsman makes a lot more sense if you open CAPsman open Winbox, and in the dialog box, work the configurations step by step going from the rightmost tab and working to the left.
Anything to the left, overwrites settings to the right.
Edit: oh, and if it has issues with provisioning slaves, that is usually due to channel or feature configuration issues (such as assigning ccXXcccc to a 2.4ghz radio....
What do you mean virtual APs? In CAPSman's provisioning settings you can set master and slave configs - there's no difference between them, you can have as many slaves as you need. This is mine:
/caps-man provisioningadd action=create-dynamic-enabled identity-regexp=ap-outdoor master-configuration=cfg-LAN name-format=identity slave-configurations=cfg-16,cfg-IOT,cfg-guestadd action=create-dynamic-enabled identity-regexp=ap-landing master-configuration=cfg-LAN name-format=identity slave-configurations=cfg-IOT,cfg-guest,cfg-vivoadd action=create-dynamic-enabled master-configuration=cfg-LAN name-format=identity slave-configurations=cfg-IOT,cfg-guest
