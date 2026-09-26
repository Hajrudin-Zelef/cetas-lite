---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-ggi7qh-dual-dhcp-wan-recursive-failover-w-pcc-load-aa911c9c
title: "r-mikrotik-comments-ggi7qh-dual-dhcp-wan-recursive-failover-w-pcc-load-aa911c9c"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-ggi7qh-dual-dhcp-wan-recursive-failover-w-pcc-load-aa911c9c.md
source_anchor: ""
source_lines: [1, 129]
sha256: d51d563c24149f38b0fdb3c97ea78297d1438f8f15c6c6c4593ef662b6306c93
---

# r-mikrotik-comments-ggi7qh-dual-dhcp-wan-recursive-failover-w-pcc-load-aa911c9c

Dual dhcp wan recursive failover w/ pcc load balancing
So I finally made this working and felt it's worth sharing. Found a few examples in the internet but WANs are on static IP. Appreciate comments and suggestions.
WANs are on eth4 and eth5.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
What do you believe the distance rfs are doing? Why for?
Also might want to censor out your publics.
Otherwise it looks fine good use of checkgw
I actually have no idea how distance works. From what I observed, routes w/ same gateway must have different distances. The route filters are just for adding comments since the WAN routes are automatically added by the DHCP client.
I'm not sure which publics are you referring to but the IPs in checkgw are public DNS server. I didn't bother replacing them with Google's and Cloudflare
So this is essentially what I need too, but I can't get it to work. (Main wan for devices and some specific devices on my second wan, but both need to be able to fail over to one another)
I'm using an RB4011.
WAN1 = ether9
WAN2 = ether10
(My WAN can be DHCP or static, does not matter, in the future one may become PPPoE though, but that is not for now).
Then I have the main bridge, and a few vlans all in an interface list.
I keep getting a connection mark error when entering them into mangle and I can't figure it out.
Sorry, I'm coming from edgeOS to where I'd just simply create two load-balance groups, and then make a firewall rule with the list of IPs I'd like to port through my second wan.
Maybe the rule is on the wrong chain. Also, you probably don't need the pcc part but instead mark the connection based on src address or src interface. Duplicate the rule per route and uncheck passthrough.
