---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-ayfepm-unifi-ap-stuck-in-upgradingdownloading-0fc55ec7
title: "r-ubiquiti-comments-ayfepm-unifi-ap-stuck-in-upgradingdownloading-0fc55ec7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-ayfepm-unifi-ap-stuck-in-upgradingdownloading-0fc55ec7.md
source_anchor: ""
source_lines: [1, 136]
sha256: 6beee0c4d728254181ddd9c4ee25974fe4e1310c3e7f518239cc6cd086f2ab6d
---

# r-ubiquiti-comments-ayfepm-unifi-ap-stuck-in-upgradingdownloading-0fc55ec7

Unifi AP stuck in Upgrading(Downloading) 
        
    I have 14 of these things, and all but this one have successfully upgraded. I can't forget the device, I can't SSH into it, either. I've already tried the factory reset procedure, with no success. Any ideas on how to proceed before I replace it?
 
     Publication archivée. Impossible de voter et de publier de nouveaux commentaires.  
        
        
        
          
        
        
        
        
         Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
You will have to do an ssh update with the direct link to the firmware on ubnt website.
Except I can't get an SSH connection, unless I don't have the correct port.
TFTP Recovery, the following link has all the details. https://help.ubnt.com/hc/en-us/articles/204910124-UniFi-TFTP-Recovery-for-Bricked-Access-Points
This worked!
Yes, this procedure work fine.
Not sure if it is the same but I had similar issues in the past. Before I now update I make sure I have the firmware cached. Never had problems after that.
You might have the correct port. Do you have the ability to plug into a network jack?
Do you have access to a network jack?
I have the same issue. Just stuck. I've also tried resetting. The controller keeps adopting and forcing adoption even after reset. Were you able to figure it out. It must be something to do with Controller.
So I fixed my issue by downloading an earlier version of the Controller software. I hate Unifi's.
What version? And where did you download it?
