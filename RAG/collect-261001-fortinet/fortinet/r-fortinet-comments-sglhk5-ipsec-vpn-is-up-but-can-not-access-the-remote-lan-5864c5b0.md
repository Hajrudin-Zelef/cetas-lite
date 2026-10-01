---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-sglhk5-ipsec-vpn-is-up-but-can-not-access-the-remote-lan-5864c5b0
title: "r-fortinet-comments-sglhk5-ipsec-vpn-is-up-but-can-not-access-the-remote-lan-5864c5b0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-sglhk5-ipsec-vpn-is-up-but-can-not-access-the-remote-lan-5864c5b0.md
source_anchor: ""
source_lines: [1, 122]
sha256: 55ec3926a47a54d2263de3eb4c8a93abb64052554bc7451f153b9d3affb868bb
---

# r-fortinet-comments-sglhk5-ipsec-vpn-is-up-but-can-not-access-the-remote-lan-5864c5b0

Ipsec VPN is up , but can not access the remote LAN 
        
    IPsec VPN tunnel between FortiGate and Checkpoint is up, but no traffic .
FortiGate can not ping the remote LAN of the Checkpoint .
SSL VPN users also can not access the remote Lan!
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Had the same issue between Fortinet and Sophos. Tunnel was up but not passing traffic, had to change the encryption algorithm and then it worked. What’s your IPSEC config? Try changing the encryption and give it another try. Also make sure you have a policy that allows the traffic…
Do a diag debug packet any "host x.x.x.x" 4 on the FGT and a fw monitor -F "x.x.x.x,0,0,0,0" on the Checkpoint to see if the packets are being sent across the tunnel, could be an SA issue and the packets are only be sent one way but not the other.
Have you created phase2 for sslvpn and policies between ssl interface and ipsec interface?
You are trying to ping in FTG CLI? Did set ping-options source “LAN-IP”
Trying ping source and if that doesn’t work, look at route table + try bouncing tunnel interface itself. Had issue where tunnel was up but IPs of next hood weren’t showing up in routing table as next hop, had to bounce tunnel interface (admin interface down, then back up) and it started passing traffic with no changes. Im on 6.4.4
