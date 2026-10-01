---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1dbe83r-udm-pro-stuck-at-boot-793204ad
title: "r-ubiquiti-comments-1dbe83r-udm-pro-stuck-at-boot-793204ad"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1dbe83r-udm-pro-stuck-at-boot-793204ad.md
source_anchor: ""
source_lines: [1, 153]
sha256: 201848929b44ecfa5e95851eb824a2e842f7583fa3f7c04d2fca7ec8ccc9ca51
---

# r-ubiquiti-comments-1dbe83r-udm-pro-stuck-at-boot-793204ad

Udm pro stuck at boot
We had a storm about a week ago and after rebooting the unit, this is as far as it gets. Holding down the reset button doesn’t do any good and there doesn’t appear to be anything on the inside of the unit I can do either. Anybody have any ideas? Needless to say needless to say restarting doesn’t help, restarting as the prompt suggested doesn’t help.
 
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
Had this happen the other week.
Sounds insane but - unplug it for 30 minutes, and plug a spare power cable in (that isn’t plugged into an AC receptacle*) - then short the prongs on the plugs to a ground with a wire. Shorting to ground won’t hurt anything - and apparently helps drain the leftover power in the caps.
10 minutes and 20 minutes didn’t work for me - needed 30, and the power plug thing.
When that’s done - go ahead and plug your regular power cable back in - and just reboot or reboot with the reset pin held in while plugging in (to get to recovery mode).
Not sure if there is a bit of near volatile memory somewhere that holds a flag preventing it from booting properly - but in my case it happened from accidentally hitting the power switch on my pdu.
I just found that as well, seems it was posted a while back thank you for sharing it!
I figured I’d just use a key puller. Left it this way for about an hour then return to the unit to plug it in and hold the reset button down and I don’t think it’s helping. Photo
No problem. It’s a frustrating bit of inanity. Let us know how it goes - some people have reported needing 8 hours to get back to operational - and some never get it back (though I think this is rare?)
Did you try to leave it powered down for a while, maybe like an hour? Any activity on the network ports? Does it respond to pings? Tried SSH into it, if it’s enabled?
It’s not starting Unifi software I assume?
Ports light up when you plug stuff in. But I don’t think it responded to the default IP address, nor the one it was assigned. Yeah I don’t think it’s booting all the way up.
This is a known problem if there is an unexpected power outage. some of these devices refused to restart..
they will replace this unit under warranty if you make them aware.
Yep, a batch from a year or two ago was affected, I had one replaced.
OP if you don't replace it now everytime there's an unexpected power outage you'll come across this.
Oh whoa, had no idea this qualified for a warrenty! I installed a UPS on mine and it still does it even when I safe shut down. That's excellent to know.
Boot it into recovery mode, default it, upload new firmware, and restore it from backup. Done this for quite a few people.
Happened to me as well, didn’t respond to dropping into recovery mode. Went and bought a new one. Hope you’re able to figure it out
So far, it’s just sitting on my bench with a lid open, but I’m about to junk it
If you decide to junk it, I’d be interested in getting the chassis if that’s possible. Happy to send a shipping label.
Dude RMA it. There is a know issue with a batch or two of these.
Take out the hard drive and try to boot it.
No hard driving in it
Are you on EA firmware? Reports of 4.0.5 having random lockups that look like that. Need to ssh downgrade to 4.0.3 if so
