---
id: collect-261001-meraki/meraki/r-meraki-comments-7l8ih4-mx-firewall-defining-the-internet-as-a-destination-2c030c09
title: "r-meraki-comments-7l8ih4-mx-firewall-defining-the-internet-as-a-destination-2c030c09"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-7l8ih4-mx-firewall-defining-the-internet-as-a-destination-2c030c09.md
source_anchor: ""
source_lines: [1, 134]
sha256: 59f6cdcfe8204e283f1d2e5bf5652e121d17450a4c6c0cafac2f4b2ce721598f
---

# r-meraki-comments-7l8ih4-mx-firewall-defining-the-internet-as-a-destination-2c030c09

MX Firewall, defining the Internet as a destination 
        
    I have MX's, with a few VLAN's under a template, that I want to add a 'deny all' rule to the end of the ruleset, but I also need devices on the VLAN's to access the internet.
I was thinking
- 
      Rules 1 - Allow required access between VLAN's
- 
      Rules 2 - Allow required internet access for specific VLAN's
- 
      Rules 3 - Deny All
- 
      Rule 4 (default) - Allow All
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
You can't. You can try and hack it by adding RFC1918 ranges but you'd need a proper zone based firewall to do this. Limitations like this are why I don't recommend Meraki for firewalling.
They're good for AutoVPN.. that's about it.
I am in the same boat in the company, any recommendations what to replace Meraki with ? we used to use SonicWall prior to and they had zone restrictions capability.
Palo Alto if you can afford it, Fortinet if not.
I did this the other day, I took a slightly different approach to it
Rule 1 - Allow access between required VLANs
Rule 2 - Deny access to all RFC 1918 ranges
Rule 3 - Deny access to RFC 3330 ranges (obviously some of these are covered in RFC 1918 and some are no longer relevant so I have excluded some from this deny rule)
Default Rule - Allow all outbound
This will block access to all private and reserved IP ranges that shouldn't be on the internet. I'll be interested to see others opinions on how they accomplish this though.
