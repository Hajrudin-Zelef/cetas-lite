---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-166krcq-fortigate-creating-rules-in-cli-5515e2f8
title: "r-fortinet-comments-166krcq-fortigate-creating-rules-in-cli-5515e2f8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-166krcq-fortigate-creating-rules-in-cli-5515e2f8.md
source_anchor: ""
source_lines: [1, 124]
sha256: 0abb0f30a7364db1c93bc4cf7f7029c4f2c54f2b54d972a1729f632bafd875d4
---

# r-fortinet-comments-166krcq-fortigate-creating-rules-in-cli-5515e2f8

Fortigate creating rules in cli 
        
    Would it be possible when creating firewall rules in CLI to automatically place them at the top?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
There’s a move command under firewall policy in cli, enter the command to move the new policy before the very first one
I have same basic security policy's that i would just like to past in cli and automaticly put them on top.
Not possible. You would need to at the very least get the ID of the first policy and then you could do it, or you purge all policies and then add your policies plus the previously existing ones.
Not with the CLI but you can do it with the API.
If you have FortiManager, you can use header/footer policies
yes there is a move command where you could tell it to move the rule before rule 1 once it is created
example:
That doesn't necessarily put it at the top, it just puts it ahead of 1.
hm yea you are right, if 1 isn't at the top then that doesn't help at all does it? Didn't consider that
