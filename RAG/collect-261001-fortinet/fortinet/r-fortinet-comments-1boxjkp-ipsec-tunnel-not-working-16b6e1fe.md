---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1boxjkp-ipsec-tunnel-not-working-16b6e1fe
title: "r-fortinet-comments-1boxjkp-ipsec-tunnel-not-working-16b6e1fe"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1boxjkp-ipsec-tunnel-not-working-16b6e1fe.md
source_anchor: ""
source_lines: [1, 135]
sha256: c8f5ca40496ffea9079086db0f723b4b54e068bf039b791624398a0c7bbc6e01
---

# r-fortinet-comments-1boxjkp-ipsec-tunnel-not-working-16b6e1fe

IPSec Tunnel Not Working 
        
    Hello!! I have an IPsec tunnel that works between 100D and 40F. The problem is that the tunnel appears to be active but the connection is not working. Anyway, when I restart one of the FGs, everything works! Anyone having the same problem?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Check your route table:
get router info details x.x.x.x
get router info routing-table all
It may be the case that your private traffic is routing via your default route (0.0.0.0)?
If that's the case, add a black hole route to null with a high AD so that the FGT doesn't create a session that will never go anywhere.
https://community.fortinet.com/t5/FortiGate/Technical-Tip-Use-of-Black-hole-route-in-site-to-site-IPsec-VPN/ta-p/192526
Otherwise, you'll need to take an IKE debug to see what's going on
https://community.fortinet.com/t5/FortiGate/Troubleshooting-Tip-IPsec-VPNs-tunnels/ta-p/195955
I would say whether or not this is the problem, black holeing your ipsec tunnels is good practice.
To add to this - also use "diag sniff packet any 'host <IP>' 4 a " - to check if
a) traffic is forwarded out from FGT
b) if the proper interface has been used
do he same on the remote side to see if the traffic has been received on the ipsec tunnel?
Since the traffic is working occasionally it is then either routing or some tunnel SA lifetime issue
Check if Phase 2 is up. It should give you a warning and tell you the reason it’s down.
Sounds like Phase 2 is going down and is not configured with keep alive.
check the logs:
No one will be able to help you like this. What exactly is not working? How did you configure the tunnel? Do the firewall policies allow the desired traffic?
There is a known issue on the latest and greatest firmware that can cause this. I had to disable npu-offload on my ipsec interface to get it to come up.
First, is phase 1 up? That will show that your configuration matches on both ends. If phase 2 is down it could be a firewall policy or routing issue.
