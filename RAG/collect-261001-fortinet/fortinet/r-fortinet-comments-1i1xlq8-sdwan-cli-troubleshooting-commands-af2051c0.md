---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1i1xlq8-sdwan-cli-troubleshooting-commands-af2051c0
title: "r-fortinet-comments-1i1xlq8-sdwan-cli-troubleshooting-commands-af2051c0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1i1xlq8-sdwan-cli-troubleshooting-commands-af2051c0.md
source_anchor: ""
source_lines: [1, 133]
sha256: 8a8d6f8d993eed856a544eaec409662e34c394aa86f9c16c29035e773b16f61a
---

# r-fortinet-comments-1i1xlq8-sdwan-cli-troubleshooting-commands-af2051c0

SD-WAN cli troubleshooting commands 
        
    When troubleshooting via the GUI it's possible to see which SD-WAN member / path is selected by the black tick box that is shown in the SD-WAN rules section. Which is the CLI command which I can use to see the same info?
I've looked at various commands and it's not obvious.
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
diag sys sdwan service
will be similar to what you see in the gui
sla(0x1) == SLA is good
sla(0x0) == SLA is bad
Cheers appreciate that. I thought that was it but the problem is it's possible for a link to be out of SLA and still be used. That shows if the SLA is healthy or not, but it doesn't show if the link is "selected" e.g. the black tick in the gui.
diagnose firewall proute list
will show you the active path
When using SDWAN, a flow is not especially tied to a single outgoing member. It can change over time.
diag sniffer packet is the only way to know precisely where each packets are routed 1 by one. You'll maybe notice asymetric routing, multiples changes in the selected members in a small timerange...
Cheers thank you.
Googling your title + FortiGate literally gives you the answer.
So you don't know the answer?
A quick check I use is go to the forwarding log, add the destination interface column and look on there. Even if the swan interface changes then an existing session will still forward using the session. Otherwise as the others have said above
Thank you
https://community.fortinet.com/t5/FortiGate/Technical-Tip-Configure-and-Diagnostic-commands-to-check-the/ta-p/194246
That was hard to find
