---
id: collect-260926-mikrotik/mikrotik/t-mikrotik-script-to-restore-backup-in-scheduler-170169-2-fce79300
title: "t-mikrotik-script-to-restore-backup-in-scheduler-170169-2-fce79300"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/t-mikrotik-script-to-restore-backup-in-scheduler-170169-2-fce79300.md
source_anchor: ""
source_lines: [1, 188]
sha256: 2b346271e93fc6c1a51bd990abab4c008f5eda9ba3e6b04c64210e35e0d9dd67
---

# t-mikrotik-script-to-restore-backup-in-scheduler-170169-2-fce79300

post by xitriyded on Oct 7, 2023
      
        
          
            
              
              
                
                  
                
                
                  
                    Hi!
I’m trying to create a script to restore backup. My script is very simple, but it’s not working from scheduler
:local backupfile "MikroTik_2011UiAS_408-20231006-0923.backup" 
/system backup load name=$backupfile
Any suggestions on how to restore backup from *.backup file?
 
                
                
                  
                
               
             
              
          
        
      
 
          
  
  post by xitriyded on Oct 7, 2023
      
        
          
            
              
              
                
                  
                
                
                  
                    Found solution
Script
:local backupfile "MikroTik_2011UiAS_408-20231006-0923.backup"
/system backup load name=$backupfile password=""
/y
In Scheduler, in On Event: just type script’s name.
 
                
                
                  
                
               
             
          
        
      
 
          
  
  post by rextended on Oct 12, 2023
      
        
          
            
              
              
                
                  
                
                
                  
                    /y at the end is wrong and useless.
 
                
                
                  
                
               
             
          
        
      
 
            
          
  
  post by xitriyded on Feb 15, 2024
      
        
          
            
              
              
                
                  
                
                
                  
                    
Not working without it though
 
                  
                  
                 
                
                
                  
                
               
             
          
        
      
 
          
  
  post by rextended on Feb 15, 2024
      
        
          
            
              
              
                
                  
                
                
                  
                    
Still wrong, becayse "/" before the y is useless, and is still useless inside the scheduler.
"y" only can be used just if the script is tested on terminal.
It's like this and there's nothing you can do about it.
 
                  
                  
                 
                
                
                  
                
               
             
          
        
      
 
          
  
  post by jaclaz on Feb 15, 2024
      
        
          
            
              
              
                
                  
                
                
                  
                    Ok, I have to ask.
Why is a static value assigned to a variable?
I.e., wouldn’t this be the same?
/system backup load name="MikroTik_2011UiAS_408-20231006-0923.backup" password=""
 
                
                
                  
                
               
             
          
        
      
 
          
  
  post by rextended on Feb 15, 2024
          
  
  post by jaclaz on Feb 16, 2024
            
          
  
  post by miku on Mar 23, 2024
