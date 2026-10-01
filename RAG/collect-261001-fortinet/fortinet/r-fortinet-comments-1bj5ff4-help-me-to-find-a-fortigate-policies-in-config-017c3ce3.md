---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1bj5ff4-help-me-to-find-a-fortigate-policies-in-config-017c3ce3
title: "r-fortinet-comments-1bj5ff4-help-me-to-find-a-fortigate-policies-in-config-017c3ce3"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1bj5ff4-help-me-to-find-a-fortigate-policies-in-config-017c3ce3.md
source_anchor: ""
source_lines: [1, 138]
sha256: 4351ab79c6cd08c9f6682fd0c33eb4c7bf2d85c7725e87253b432b61b21c31e1
---

# r-fortinet-comments-1bj5ff4-help-me-to-find-a-fortigate-policies-in-config-017c3ce3

Help me to find a fortigate policies in config file 
        
    hi, im still new in this field.. i ve migrate one of firewall, but seem like the policies of the firewall quite not the same as the firewall before. oh btw i used copy and paste from config file method. now i have no idea how to find those policies.. help me guys. thanks :>
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
The firewall policy section starts with:
hi, thankyou for the comments.. ive tried to find those..but still failed. btw this is the config file.. is this firewall policy?
"config firewall address
edit "FIREWALL_AUTH_PORTAL_ADDRESS"
set uuid 936103d0-3d75-51e7-782e-314711bd4064
set visibility disable
next
edit "all"
set uuid 78252252-f1e6-51e6-174b-4cb68f6c3147
next
edit "SSLVPN_TUNNEL_ADDR1"
set uuid 782527c0-f1e6-51e6-75ba-bd572fd89ea5
set type iprange
set start-ip xx.xxx.xxx.xxx
set end-ip xx.xxx.xxx.xxx
next
edit "Peplink_LAN"
set uuid 78252ce8-f1e6-51e6-df70-5310a54b509f"
GUI: Firewall -> Policy
CLI:
config firewall policy
Ref: https://docs.fortinet.com/document/fortigate/7.2.8/cli-reference/84566/fortios-cli-reference
thankyou !
