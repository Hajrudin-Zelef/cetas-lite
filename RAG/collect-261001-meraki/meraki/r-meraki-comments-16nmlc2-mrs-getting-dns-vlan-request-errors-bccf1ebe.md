---
id: collect-261001-meraki/meraki/r-meraki-comments-16nmlc2-mrs-getting-dns-vlan-request-errors-bccf1ebe
title: "r-meraki-comments-16nmlc2-mrs-getting-dns-vlan-request-errors-bccf1ebe"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-16nmlc2-mrs-getting-dns-vlan-request-errors-bccf1ebe.md
source_anchor: ""
source_lines: [1, 138]
sha256: 8ba9343169fa996923359033fb3015706b34ee9ffb0d4bea84bd5882bf429d71
---

# r-meraki-comments-16nmlc2-mrs-getting-dns-vlan-request-errors-bccf1ebe

MRs getting DNS VLAN request errors 
        
        
        
    
    
    I'm seeing MR access points in our network getting "red" indicators that they're experiencing 2 or more VLAN request errors, and they pertain to VLAN 0 (untagged). Then occasionally the errors will go away and the indicator will turn green. Is this anything to get concerned over, and how do I fix it if so? The SSID has tagging turned off, and the AP isn't tagged with any VLAN, the field is blank. Thanks!
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
What does your switchport configuration What does your switchport config look like for the APs?
All of them set to Access, not Trunk, and not tagged.
I would set the APs as follows:
Port status Enabled
Type Trunk
Native VLAN 1
Allowed VLANs 1,3-7,33
Access policy Open
Link negotiation Auto negotiate (1 Gbps)
RSTP Enabled (Forwarding)
Port schedule Unscheduled
Port isolation Disabled
Trusted DAI Disabled
UDLD Alert only
Tags
PoE Enabled
Port mirroring Not mirroring traffic
I use Trunk ports too for my MR's but we have different SSID's on different VLANs. If you're only using 1 VLAN for Management + SSID's (or Meraki DHCP) then access should be fine.
Wouldn't hurt to try though? Change to trunk, make native trunk vlan the one you use normally and leave it at that. See if it makes a difference.
