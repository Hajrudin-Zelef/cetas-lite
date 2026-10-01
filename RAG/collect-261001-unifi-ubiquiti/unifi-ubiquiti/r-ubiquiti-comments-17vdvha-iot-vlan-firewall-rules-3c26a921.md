---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-17vdvha-iot-vlan-firewall-rules-3c26a921
title: "r-ubiquiti-comments-17vdvha-iot-vlan-firewall-rules-3c26a921"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["sol", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-17vdvha-iot-vlan-firewall-rules-3c26a921.md
source_anchor: ""
source_lines: [1, 167]
sha256: f7988215641b5093df0872daed498939a6e10feef4ccf4e683a38fcbd819d876
---

# r-ubiquiti-comments-17vdvha-iot-vlan-firewall-rules-3c26a921

[supprimé]
      IoT VLAN Firewall Rules 
        
        
        
    
    
    
  
        
        
         
     Désolé, cette publication a été retirée par les filtres de Reddit.  
         
         
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
Devices on the same VLAN can talk directly to each other (Layer 2 switching) whereas devices on separate VLANs will need a L3 router/firewall to route between the 2 VLANs. This allows the router/firewall to inspect that traffic and deny/allow only specific ports and/or devices to talk to other ports/devices so in your example you could allow your laptop to access the IoT camera but the camera would not be allowed to access your laptop.
BTW, you should never overlap the same subnet on different VLANs.
rejig the subnets so they don’t overlap (as previously mentioned)
enable isolated network / guest on the IoT vlan, if you haven’t already (settings > network > select network > tick box for isolate network)
Why? By default ubiquity enable inter-vlan routing, unless you isolate the vlan.
3) traffic from default to IoT is the correct way to do this (should be guest out FW rule or did you set a traffic rule? App Fw > traffic rules) Either way, if possible I would lock it down further, use the profiles and create an IP group of users you want to have access and ‘lock’ / fix those ips to those devices on the router, or specify a single device / ip, simplified when creating a traffic rule.
4) does this defeat the purpose? not really, you’re putting less secure devices on a vlan and allowing traffic one way; from main / default to less secure, not the other way around, preventing these devices from having access to more sensitive areas of your network.
Commentaire supprimé par un membre de l’équipe de modération
I'm currently going through this myself as the moment. Would you mind sharing what you ended up with as far as rules?
I think I understand VLANs and I’ve tried to follow the guides but I just have this mental block where it doesn’t all add up to me. The firewall rules are what really f my head up. Like so many things in tech I know I’ll eventually click but I get frustrated I can’t ‘get it’ quicker. Does sound like you’re doing it right though from what I remember
First of all, what exactly is your purpose for using vlans? Performance, security, management? I ask because there are different methods to achieve this. Layer 2 or layer 3 vlan, vrf, acl. What's your goal?
If you want to use multiple vlans, I've found it best to set the switchport to trunk mode with native vlan on default vlan 1 and then set the allowed vlans on your trunk port to your respective vlans, including the mgmt vlan
I've had mixed success when setting the unifi mgmt vlan to the native vlan. Better to have the Unifi send everything tagged if you will be using vlans
Commentaire supprimé par un membre de l’équipe de modération
That's understandable. There's a big jump now with ai automation, but you just need some old-fashioned access lists rather than the complexity of an additional vlan (unless you have cameras that are bottlenecking your switch throughput)
Can i ask what type of access rules more specifically
You have two major types - simple and extended
Simple acl just checks the incoming ip against a list at the destination, meaning trying to access the device will still fully traverse your network before getting blocked so it's not an optimal use of resources but easier to setup
Extended access list let's you check against both the source and destination so you can block said traffic at the first hop in order not to waste switch resources.
What is your switch and network equipment besides the ap/dreamrouter?
So basically, don't put your default vlan in your Unifi networks (which would be set as vlan only networks)
I think your issue is using the default vlan. Set a new vlan for mgmt and set the AP Mgmt on that vlan only network
It just works better that way when you are using multiple vlans with the unifi.
I also exclude vlan 1 from my trunk port to the unifi for good measure and have no devices on default vlan (that's more for security than anything)
Hey OP, just watch this video: https://www.youtube.com/watch?v=UGBobTInIBc
Explains things pretty well. This is what I did, and works like a charm. It's actually not that difficult once you do it once.
