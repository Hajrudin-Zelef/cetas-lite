---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-18t0ln1-can-someone-explain-tagged-vlan-management-107a7f74
title: "r-ubiquiti-comments-18t0ln1-can-someone-explain-tagged-vlan-management-107a7f74"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-18t0ln1-can-someone-explain-tagged-vlan-management-107a7f74.md
source_anchor: ""
source_lines: [1, 162]
sha256: 9a2895ca4845b6f21307717bbf1f603a27b9194d4189702a3a97d3478026e72d
---

# r-ubiquiti-comments-18t0ln1-can-someone-explain-tagged-vlan-management-107a7f74

Can someone explain tagged VLAN Management? 
        
        
        
    
    
    Can someone explain tagged VLAN Management when configuring a switch port? and how to properly use it and what it does.
and what the options do:
Allow All
Block All
Custom
Thanks,
 
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
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
This controls which 802.1Q tags are allowed on a specific switch port. Configuring this is usually done for security reasons in larger networks so that you can only use the allowed/approved/native VLAN(s) for the specific switch port.
Allow All = Any VLAN tag allowed
Block All = No tagged VLAN traffic allowed (untagged/native VLAN permitted)
Custom = Specify which VLAN tags are allowed
Then what's the native vlan field for?
For instance how would you assign a port a single vlan then? Use default for native and select the single vlan you actually want to use under tagged?
Use the vlan you want as the native and select block all under tagged?
It literally makes no sense the way its laid out and Ive been offline and missed work today. I cant find ANY documentation and the internet searches bring up last year post which do not apply anymore.
Native = untagged
Yes. This will put all the traffic on the native vlan you selected, but block any other tagged vlan traffic.
If you don't set VLAN IDs manually on clients, it's as follows:
Uplink/Downlink ports to other unifi devices: allow all
ports for single clients: block all
Set your primary network as you wish.
I've been trying for 2 days to figure out how to use this software. There no documentation anywhere.
Set primary network to whatever? Isnt that the vlan you want to use? How does one isolate a port to a single vlan? Tagged/primary/custom just set to the vlan.
what vlan does the controller have to be in because I lose connection to switches if I mess with the vlan at all and have to factory reset.
Don’t follow this if you want a secure network. Trunks should only carry the vlans needed
Now all switches will need all vlans.
Do I get it right that Block All in Tagged Vlan Management does just that - blocks all the Vlans but the assigned one?
And only the ports that need to carry all the Vlans to another switch should have Allow All or Custom?
A key word to help with your googling is vlan "trunk" port. Those settings are what affect if the port is acting as a trunk port or access port.
Google 802.1q
