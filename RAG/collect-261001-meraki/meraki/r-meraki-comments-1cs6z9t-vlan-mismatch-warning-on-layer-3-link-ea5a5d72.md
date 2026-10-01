---
id: collect-261001-meraki/meraki/r-meraki-comments-1cs6z9t-vlan-mismatch-warning-on-layer-3-link-ea5a5d72
title: "r-meraki-comments-1cs6z9t-vlan-mismatch-warning-on-layer-3-link-ea5a5d72"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1cs6z9t-vlan-mismatch-warning-on-layer-3-link-ea5a5d72.md
source_anchor: ""
source_lines: [1, 125]
sha256: 287ab1a2374a88d30d022e07b18050041b85ff4f4a86612c44a9cd40d1c82041
---

# r-meraki-comments-1cs6z9t-vlan-mismatch-warning-on-layer-3-link-ea5a5d72

[supprimé]
      [deleted by user] 
        
    Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
It's a software issue on the Meraki side. Have you reached out to Meraki support and asked if they have a work around to remove the warning. It will work fine but will just throw that error
Yes many times I have errors in the dashboard that I know are incorrect and they take quite awhile to clear or I have to contact Meraki support to clear them eventually.
I had the similar issue on trunk ports with all vlans allowed on the 390 series Meraki switches allow 1-4098 or whichever 4k number is max while 390 allowed me to put only 1-1000
Move the Cisco end to SVI's and make the VLAN IDs match.
Do you have the svi configured on the meraki switch?
Commentaire supprimé par le membre
As others said above, this might be a software issue - I would check with Meraki support
Meraki check all the VLANS configured both side are same or not, if not same it will generate the warning
I may be reading this wrong but it sounds like it is a vlan mismatch. You're allowing all vlans on the cisco side and specific on the meraki side?
Did you try VLAN 0 ?
