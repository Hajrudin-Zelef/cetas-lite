---
id: collect-261001-meraki/meraki/r-meraki-comments-t9f5s1-autovpn-up-bidirectional-communication-between-mx-c3294769
title: "r-meraki-comments-t9f5s1-autovpn-up-bidirectional-communication-between-mx-c3294769"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-t9f5s1-autovpn-up-bidirectional-communication-between-mx-c3294769.md
source_anchor: ""
source_lines: [1, 119]
sha256: 1f91387156930d5421a1416385c5b7d07d8ef78ed5257165d4bfd1b6ed9b1d59
---

# r-meraki-comments-t9f5s1-autovpn-up-bidirectional-communication-between-mx-c3294769

Auto-VPN Up Bi-Directional communication between MX Hub and spokes down 
        
    I have a MX Hub working as a mesh hub, and three remote sites with mx devices working as a mesh spokes. The auto-vpn is up, but there is no bi-directional traffic between the hub and spokes.
I have an open ticket with Meraki support, and so far they have reset one link for the vpn-registry. However we still see no change.
Any recommendations?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Worked with a higher level of support and spent most of our time troubleshooting from the spoke devices. After we verified that the spoke devices could successfully ping the MX hub inside IP, we moved onto the MX Hub itself and found the MX Hub could no longer ping the next hops (switch) default gateway and that is where the issue was. We changed the MX inside IP to another VLAN and everything came back up.
Check what IPs your WAN is using at the HUB. Have you tried rebooting your ISP router?
