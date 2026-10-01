---
id: collect-261001-huawei/huawei/r-networking-comments-d5wcnh-same-vlan-on-two-trunk-ports-5b3a69b5
title: "r-networking-comments-d5wcnh-same-vlan-on-two-trunk-ports-5b3a69b5"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "pruning"]
source: docs/RAG/collect-261001-huawei/r-networking-comments-d5wcnh-same-vlan-on-two-trunk-ports-5b3a69b5.md
source_anchor: ""
source_lines: [1, 134]
sha256: 083956b53109e1acb263bbb61629280d5c58ca2bcac726294da0aaf89663240b
---

# r-networking-comments-d5wcnh-same-vlan-on-two-trunk-ports-5b3a69b5

Same VLAN on two Trunk ports 
        
    I have a pair of stacked Huawei S6720 switches. Can I add the same VLAN(s) to two link aggregation groups (EthTrunks)? i.o.w. eth-trunk 1 contains vlans 101-104 and eth-trunk 2 also contains vlans 101-104.
EDIT: Yup, a stressful situation with crazy time pressures is the breeding ground of finger trouble. A day later and a second bite at the apple proved successful. Thanks to all that contributed.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
STP will block one of those.
The trunks will be connecting to separate devices in a high-availability pair of sonicwalls.
Why wouldn't you be able to do this?
If both LAGs are going to the same downstream device it may need some more thinking about, but multiple ports in the same VLAN is standard operating for switches. Almost the whole point of them.
Combine the 2 links into etherchannel to make better use of the link.
There is no issue with putting 2 vlans on a trunk link. I would certainly advise allowing all vlans on the link if it is going to another distribution switch otherwise your topology might get real messy down the line. You can turn on vlan pruning so broadcasts are limited to devices that have corresponding vlans.
Of cause you can, a LAG is logically a normal port.
LACP
Sure you can.
Flowwise it will be like this:
somedevice <-> VLAN 101-104 <-> your switch <-> VLAN 101-104 <-> someotherdevice
eth-trunk1 is one static link aggregation group which contains 1 or more physical interfaces.
And eth-trunk2 is another static link aggregation group which contains 1 or more (other than above) physical interfaces.
So when you configure stuff you configure it against the lag interfaces (eth-trunk1 and eth-trunk2) and not individual interfaces.
Only thing to configure individual interfaces is like description if needed.
I'm sorry.
Yes. LAGs work, even on Huawei.
i dont see why not
