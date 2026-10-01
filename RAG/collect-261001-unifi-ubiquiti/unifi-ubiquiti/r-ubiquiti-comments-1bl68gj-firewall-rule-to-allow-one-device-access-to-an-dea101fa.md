---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1bl68gj-firewall-rule-to-allow-one-device-access-to-an-dea101fa
title: "r-ubiquiti-comments-1bl68gj-firewall-rule-to-allow-one-device-access-to-an-dea101fa"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1bl68gj-firewall-rule-to-allow-one-device-access-to-an-dea101fa.md
source_anchor: ""
source_lines: [1, 154]
sha256: facbb581f5c78177b018c9ac8546766aabb66682c90a18415ed90f14b756ed91
---

# r-ubiquiti-comments-1bl68gj-firewall-rule-to-allow-one-device-access-to-an-dea101fa

Firewall rule to allow one device access to an isolated network? 
        
        
        
    
    
    This seems like an easy question but I have been trying to get this to work with simple and advanced firewall rules and wondering what I am doing wrong.
I have a device on Default network (vlan 1) that I would like to be able to connect to a device on Isolated network (vlan 3). Any tips on how to set that up?
 
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
The rule will prevent the isolated VLANs from initiating traffic to the Default network.
Make sure it is placed below your allow rules.
Isolated VLANs cant initiate traffic to any other LAN network by default so I don’t understand what this rule is doing?
It is my understanding that VLANs can talk to the Default network by default unless you put rules in place to prevent this or you've enabled network isolation. I assumed you wanted to isolate the network with rules.
I don't think you can accomplish what you are asking with network isolation enabled. It's kind of the point of that setting.
This rule will allow any isolated VLANs to reply to traffic initiated by a device on your default network.
This should be the very first firewall rule.
Thank you! This was driving me nuts.
Is it possible that this is currently bugged? I have an isolated network IoT and I can’t get a rule working that allows a different network to access IoT. I have to manually allow specific IP addresses.
Works: IoT (isolated) IoT traffic to IP of my PC
Doesn’t work: IoT (isolated) IoT to Network where my PC is located
Here is the simple traffic rule that lets my HomeAssistant into other isolated networks. I am not a firewall expert but this seems to work. I can see in the detailed firewall rules that Unifi put this ahead of the isolation rules.
I have a similar rule that lets these networks also connect to my home assistant based on it's IP address. That seems to be required for HA.
If someone can suggest a better/easier way to do that please do!
This doesn't work for me :/
This rule will allow all traffic from your Default network to any VLAN.
You will need to create an address group for all addresses reserved for the local space (RFC1918 Subnets):
192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8
If you only require one specific address, use it instead of the RFC1918 Subnets.
