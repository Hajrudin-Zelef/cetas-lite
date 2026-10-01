---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-7v9had-cli-command-to-view-only-denied-packets-62239228
title: "r-fortinet-comments-7v9had-cli-command-to-view-only-denied-packets-62239228"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-7v9had-cli-command-to-view-only-denied-packets-62239228.md
source_anchor: ""
source_lines: [1, 126]
sha256: 931d6eafb52d508f20272e6f8d8ebe6016f8705763b440dc3d3531f67837c150
---

# r-fortinet-comments-7v9had-cli-command-to-view-only-denied-packets-62239228

CLI command to view only denied packets? 
        
    I'm coming from iptables/netfilter and tail -f | grep used to be my best friend.
Now I'm having a hard time troubleshooting some connectivity problems inside my network. Is there a CLI command I can only look at the real time denied packets? I'm messing with diagnose debug flow but not having good results.
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
I would just check the logs for the deny rule in real time on the gui.
I was going to suggest the same thing cause that's what I do now, but I am actually hoping someone has a way to quickly show/filter on denied logs from the CLI!
If you know header info for the specific traffic that you suspect is getting denied you can use the follow commands:
diagnose debug flow filter <filter info>
diagnose debug flow trace start
diagnose debug enable
If you know its the implicit deny dropping the traffic then enabling logging on policy 0 is easier, but if you're not sure doing the debug flow will tell you what policy the traffic is matching. This is also useful if traffic is getting blocked by a non-policy reason, such as failing reverse path forwarding.
You may want to turn "on" logging of the implicit deny policy on your list of policies, so the denied traffic is logged. Additionally, this policy is known as Policy ID 0.
Logs at the cli may be viewed with execute log display, and execute log filter to set what logs you want to see (category, starting line, number of lines, etc). Kinda ugly imo, but haven’t found another CLI method.
