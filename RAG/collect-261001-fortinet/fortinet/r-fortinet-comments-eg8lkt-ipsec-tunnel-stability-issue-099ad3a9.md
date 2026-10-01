---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-eg8lkt-ipsec-tunnel-stability-issue-099ad3a9
title: "r-fortinet-comments-eg8lkt-ipsec-tunnel-stability-issue-099ad3a9"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-eg8lkt-ipsec-tunnel-stability-issue-099ad3a9.md
source_anchor: ""
source_lines: [1, 142]
sha256: 02ea7cbb946ccfa836c6c3554948c42ad218fdeb7ff600327f2b7d145d248895
---

# r-fortinet-comments-eg8lkt-ipsec-tunnel-stability-issue-099ad3a9

Ipsec Tunnel stability issue 
        
    
      Hi ,
I have an issue with an Ipsec S2S tunnel between FGT 500E and Forcepoint , every 3 or 4 days the tunnel becomes Down and I have to use this  2 commands everytime to make it UP again :
    
      diagnose vpn ike restart
diagnose vpn ike gateway clear
    
Does anyone have a clue of what might be causing this ?
Thanks
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
We had a costumer with same issues, FGT 90D to Cisco Router. And after the commands in your post was everything fine.
My solution was: We have add a blackhole route (same as the normal static route, with the interface Blackhole and the distance 254), for each VPN interface. Since then are the tunnels always up and I never used the commands in your post again.
I will try this , should I add the blackhole route on both FW or only in Fortigate ?
If you can on both site.
In my case I‘ve found no Blackhole on the Cisco Router site. :)
Can you explain why the black hole route resolves this issue?
Dead Peer Detection on Idle in Phase1
Autokeepalive and autonegotiate on individual Phase2s
Make sure you are running the latest code
Hi!
Normally the tunnel goes down if there is no traffic passing through. Did you try to ping the other site before running the diagnose commands?
Yes I tried pinging the other side but the ping doesn't go through since the tunnel is down
It is a route based tunnel not a policy based one
Look at Phase 2 Selectors, under Advanced.
Verify the Key lifetime is the same on both ends of the tunnel.
With no tunnel, the two sides negotiate and come up. If one times out early, it drops, tries to re-key with the other tunnel that still has a good key with life left on it, so it rejects the re-key attempt.
This is why your tunnel comes up at first, but then goes down until you clear the tunnel.
The key lifetime is the same on both sides
Not sure if this relevant in the latest firmware but I use to have this issue in 5.6 when using address objects (dst-addr-type name) for the phase 2.
