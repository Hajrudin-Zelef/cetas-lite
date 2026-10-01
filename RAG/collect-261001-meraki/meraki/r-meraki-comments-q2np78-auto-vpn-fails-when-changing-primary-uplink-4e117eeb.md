---
id: collect-261001-meraki/meraki/r-meraki-comments-q2np78-auto-vpn-fails-when-changing-primary-uplink-4e117eeb
title: "r-meraki-comments-q2np78-auto-vpn-fails-when-changing-primary-uplink-4e117eeb"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-q2np78-auto-vpn-fails-when-changing-primary-uplink-4e117eeb.md
source_anchor: ""
source_lines: [1, 121]
sha256: 21d0a274232832aff2658784ebc9a124040561848b5c1ef854ace9e493a04e37
---

# r-meraki-comments-q2np78-auto-vpn-fails-when-changing-primary-uplink-4e117eeb

Auto VPN fails when changing Primary Uplink? 
        
    We have a new higher speed internet connection that we're trying to migrate to. We tried changing the primary uplink from WAN 1 to WAN 2 where the new connection is, but it drops all of our VPN connections and they never renegotiate and connect.
Short of manually changing the IP addresses and moving the physical cables is there some way to make this work?
Both WAN 1 and 2 are static IPs and work perfectly fine.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
If the VPN tunnel never forms on the new internet connection, that sounds like something upstream might be filtering traffic. Is their a restrictive firewall on the new internet connection? In Dashboard if you browse to Help > Firewall Info you will see the line stating you need to ensure UDP 9350 or UDP 9351 are allowed to specific destinations so the Meraki can contact the VPN registry.
In addition to that, the Meraki will try to do UPD Hole Punching to decide on a UDP source port. If this is not working you can set this to "Manual" and choose a specific UDP port, but you will need to make sure any upstream filtering devices forward traffic on that chosen port to your MX.
An additional suggestion is under Security & SD-WAN > SD-WAN & Traffic Shaping, there is an "Active-Active AutoVPN" setting. From there you can control if the VPN tunnel will form on both uplinks simultaneously, or if it should wait for the failure of WAN 1 to then build the tunnel on WAN 2. This setting will only play a factor once you sort out the probable firewall issues above. Also if you are trying to move to our new internet connection as the only uplink going forward then this setting won't mean much to you.
Sounds like a typical Meraki "Feature"
