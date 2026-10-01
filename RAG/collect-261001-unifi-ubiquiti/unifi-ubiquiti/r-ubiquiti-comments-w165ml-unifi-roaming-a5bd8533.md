---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-w165ml-unifi-roaming-a5bd8533
title: "r-ubiquiti-comments-w165ml-unifi-roaming-a5bd8533"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-w165ml-unifi-roaming-a5bd8533.md
source_anchor: ""
source_lines: [1, 143]
sha256: 48dcec04ec5549a87762fbbfe662410e2e4a6bd966c420c3bbf81f4284871e1e
---

# r-ubiquiti-comments-w165ml-unifi-roaming-a5bd8533

Unifi Roaming 
        
        
        
    
    
    If roaming from WAP to WAP is controlled by the device (iPhone), what does unifi do to provide seamless coverage that any other collection of WAPs doesn’t do?
 
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
The unifi controller doesnt do anything to help wireless client roam outside the radio settings you set on your access points.
So why is unifi advertised as a solution for seamless wifi connectivity?
The controller allows you to control the transmit power and minimum RSSI of clients before they get kicked which can help you fine-tune a large WiFi area. But WiFi roaming at the end of the day is a standard handled by the CLIENT, so there are lots of companies that can do this. You could accomplish the same with Tp Link Omada, Aruba Instant, Ruckus, Merkai, among others.
The one thing a controller WILL help with is fast roaming. If you are using wpa-enterprise with a Radius server that can improve roaming speed. (It helps with regular WPA too but the advantage is small)
People like ubiquiti for good WiFi coverage and roaming because they produce a lot of different Poe-powered APs in different form factors at reasonable prices. (In-wall with switch ports, ceiling mounted small and large ‘saucers,’ table top or outdoor mounted (FlexHD/mesh 6) etc. That is usually a better looking installation than using a few consumer routers switched to AP mode.
Because it sounds cooler. Its all marketing, unifi controller just manages the access points. It actually doesnt do anything like move clients over to access points that would be "better for the wireless client" like you see in other wireless enterprise gear that will actually attempt to move clients to better access points
What do you think happens in all of wireless world? It’s no different than a cellular handoff. 5G handoff at high speed (think driving down a freeway) … that is en engineer challenge. Don’t put your APs too close to each other, try to use their coverage tool, and you should be fine. APs too close from one another screw up the handoff process as both candidates are too alike.
Setting a minimum RSSI value in the AP is what sends the signal to the client that encourages it to find another AP. So, the actual roaming is always done by the client. I can imagine that different clients give this "received direction" to find another AP different weights in their own algorithms, but I've always been able get acceptable roaming results in multiple AP setups by doing the tuning the channel selection, transmit power and RSSI values carefully.
Nothing.
