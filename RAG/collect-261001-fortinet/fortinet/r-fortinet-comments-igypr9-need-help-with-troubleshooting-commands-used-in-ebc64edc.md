---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-igypr9-need-help-with-troubleshooting-commands-used-in-ebc64edc
title: "r-fortinet-comments-igypr9-need-help-with-troubleshooting-commands-used-in-ebc64edc"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-igypr9-need-help-with-troubleshooting-commands-used-in-ebc64edc.md
source_anchor: ""
source_lines: [1, 131]
sha256: e411477add88f4520a958504ad58e1a2de5166711955e673518909ceceac6e9a
---

# r-fortinet-comments-igypr9-need-help-with-troubleshooting-commands-used-in-ebc64edc

Need help with troubleshooting commands used in the CLI on Fortigate v6.2.4(GA). 
        
        
        
    
    
    I have 2 sites. On site A, RDP to an outside server works. On site B, it doesn't. Both firewalls are nearly identical with a few VIP exceptions. I am looking for some commands that will help me figure out maybe what policy (or something else) is stopping the RDP session from working on site B?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
diagnose debug disable
diagnose debug flow trace stop
diagnose debug flow filter clear
diagnose debug reset
diagnose debug flow filter addr x.x.x.x <-- source public ip
diagnose debug console timestamp enable
diagnose debug flow trace start 100
diagnose debug enable
Thanks. I ran the command, and now I am trying to decipher the results.
In addition to the above, sniffing might help too.
Diag sniffer packet any 'host x.x.x.x and port x' 4
Typing a '?' will also show you the syntax and examples.
