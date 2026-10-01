---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1clebj8-how-is-ubiquity-with-multiple-sites-f62a347b
title: "r-ubiquiti-comments-1clebj8-how-is-ubiquity-with-multiple-sites-f62a347b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1clebj8-how-is-ubiquity-with-multiple-sites-f62a347b.md
source_anchor: ""
source_lines: [1, 150]
sha256: 9bd8e4288c1234014eefa54941491a3b75505fee5402cae8f31f90a720b12407
---

# r-ubiquiti-comments-1clebj8-how-is-ubiquity-with-multiple-sites-f62a347b

How is Ubiquity with multiple sites 
        
        
        
    
    
    I'm working with a company with quite a few sites 20+
I am thinking of replacing their network devices switches and AP with ubiquity, but when I previously used it I was quite disappointed with the multi site support.
They already have a firewall IIDS/IPS and I don't want to use cloudkey.
 
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
To answer your question can it be done, yes, multi site support is easy either using Unifi.ui.com with the cloud console lineup or using a cloud key with multi site support built in to the controller.
To answer your question about will it work for them, there are so many questions and you provided so little information. Questions like how many users? How big are the spaces? What are there network bandwidth needs? Do they need PoE? Just to name a few
I think it needs to be said along with my answer, but why would you deploy something, especially to that many sites, without enough experience with the solution to know if it will work or not?
I'm asking this because when I last used it they did not support it at all, so my question is quite simple I think does it work without the cloudkey, I'm fine using a management "console" installed at each site but the cloudkey is in my opinion to expensive and I've had too many issues with it.
As for the amount of client's that's hard to answer, as some sites might have 10 users and the next a 1000, bandwidth is at minimal 100Mbit but there are multiple sites that have 800-1000Mbit, mostly fibre.
Some spaces are office buildings while others are factories, so it's very diverse and most sites still need to be indexed.
The cloud key is $200. Not really that expensive if you are talking about 20 sites. They just work in my experience. Had one in a rack mount adapter for years and it’s still running. Maybe you looked at the enterprise version which is overkill and expensive for your needs? If you are concerned about cost you can always self host the controller for free.
However, the cloud key cannot manage any of the cloud consoles like UDM pro or UDM SE. You would have to use an independent gateway if you wanted Unifi routing. Gateways and Cloud Keys
The cloud consoles will work and you would use Unifi.ui.com to manage your fleet. You would have to setup each individual console though and can’t take advantage of the multiple site function like a cloud key would have to quickly deploy wifi or switching config at multiple sites.
People complain about Unifi being not enterprise grade but 90% of the time its because they didn’t know what they were doing when they deployed it. It is perfectly capable of this type of deployment.
Multiple Sites as in Site-to-Site-VPN or one controller with remote unifi devices?
With the integrated "Site Magic" you have a maximum of 15 sites to be interconnected. Q3 in the FAQ: https://community.ui.com/questions/Introducing-Site-Magic/4861b026-5374-404b-9b1b-d2b142acae62
Not sure if IPSec or OpenVPN would handle stuff better or more sites than WireGuard (which is used by "Site Magic") when manually set-up. IRC IPSec also has 15 tunnels max.
What about multi site did you not like?
