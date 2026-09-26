---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-gjobiz-system-scheduler-or-system-scripts-16812cbe
title: "r-mikrotik-comments-gjobiz-system-scheduler-or-system-scripts-16812cbe"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-gjobiz-system-scheduler-or-system-scripts-16812cbe.md
source_anchor: ""
source_lines: [1, 125]
sha256: 7e1705dfa5ba1ea3808f0e5e26c798f67c00f6ecb8315b8df2cad16954147aa4
---

# r-mikrotik-comments-gjobiz-system-scheduler-or-system-scripts-16812cbe

/System scheduler or /System scripts? 
        
        
        
    
    
    Our scripts that only ever run on a schedule and we never run manually get dropped into /system scheduler whereas scripts that we only run manually or run both manually and on a schedule get dropped into /system scripts.
What is best practice and/or common methodology?
Edit: the longest script that we put directly into scheduler is three lines.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Put all your scripts in /system scripts. Use /system scheduler to call up your scripts.
I also don’t bother to create scripts for items that can be taken care of with a single short command in a scheduler item. E.G.
/system backup save name=“diskN/nightly.backup”
Yeah, but I think OP was considering putting full-blown scripts, not just one-liners, into the scheduler.
