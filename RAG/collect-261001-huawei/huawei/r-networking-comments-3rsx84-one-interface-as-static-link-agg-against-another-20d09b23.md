---
id: collect-261001-huawei/huawei/r-networking-comments-3rsx84-one-interface-as-static-link-agg-against-another-20d09b23
title: "r-networking-comments-3rsx84-one-interface-as-static-link-agg-against-another-20d09b23"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-3rsx84-one-interface-as-static-link-agg-against-another-20d09b23.md
source_anchor: ""
source_lines: [1, 132]
sha256: bca2ade4d21b7dfa8c801a46f430ddcbe612495e86b249c832d36e2b0835e260
---

# r-networking-comments-3rsx84-one-interface-as-static-link-agg-against-another-20d09b23

One interface as static Link Agg against another one as dynamic Link Agg? 
        
    Hello everyone,
Quick question here. I am working with two Huawei switches. One of them supports Eth-Trunks (LACP), while the other one only supports static Link Aggregation. I need to create an Eth-Trunk between them. Would it work if I configure each device differently, one with an Eth-Trunk and the other one with an static LA?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
They must be part of the same group and use the same settings.
Are you sure one of them lacks support for LACP and the other lacks support for static trunk?
Which models are those?
Models of the devices are quidway huawei s3900 and s5700
The s3900 got Manual Aggregation Group, Static LACP Aggregation Group and Dynamic LACP Aggregation Group.
Manual Aggregation Group:
Static LACP Aggregation Group:
Dynamic LACP Aggregation Group:
The s5700 has eth-trunk (which is the manual method described above) along with static LACP. I couldnt locate if it supports dynamic LACP but if you load both devices with latest firmware it should be easy to find out.
Stuff to google for:
huawei s3900 configuration guide
huawei s5700 configuration guide
also google for:
huawei s5700 Typical Configuration Examples
You're gonna have to go with lowest common denominator - static. LACP wont form unless it's on both sides.
I'm sure both switches at least support static link aggregation. But I'm sure both supports LACP too. Dig into the documentation more.
