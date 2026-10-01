---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-serffe-check-policy-id-using-command-line-38f4df70
title: "r-fortinet-comments-serffe-check-policy-id-using-command-line-38f4df70"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-serffe-check-policy-id-using-command-line-38f4df70.md
source_anchor: ""
source_lines: [1, 126]
sha256: 57f83d09274500d4552c8f68ab1e30f12b357caa14f309070011050b06c54aed
---

# r-fortinet-comments-serffe-check-policy-id-using-command-line-38f4df70

Check policy ID using command line. 
        
    Hi everyone,
I have this scenario where a fortigate is connecting a workstation and a server and the fortigate has various number of policies. I want to know which command can I use to identify the policy ID that allows communication between the PC and the server using the backend of Fortigate firewall server without accessing the GUI?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Thank you Golle!
If the traffic is already being allowed, you can also find the existing session in the session list and get the policy ID from there:
diag sys session filter clear
diag sys session filter src <IP>
diag sys session filter dst <IP>
diag sys session filter dport <port>
diag sys session list
One of the resulting lines should contain an item saying
policy_id=<ID>.
Oh! I will try that. Thank you Pabechan.
