---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-17qlimv-why-would-i-roam-from-a-perfectly-good-ap-to-a-bb7cf1fc
title: "r-ubiquiti-comments-17qlimv-why-would-i-roam-from-a-perfectly-good-ap-to-a-bb7cf1fc"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Samsung"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-17qlimv-why-would-i-roam-from-a-perfectly-good-ap-to-a-bb7cf1fc.md
source_anchor: ""
source_lines: [1, 161]
sha256: 0b804a844e86372f880d458bd7435e9892c44456300190ce66f0492b413cc9d5
---

# r-ubiquiti-comments-17qlimv-why-would-i-roam-from-a-perfectly-good-ap-to-a-bb7cf1fc

Why would I roam from a perfectly good AP to a terrible one?
I wasn’t moving at all. This seems to happen a lot. 2 APs, one on second floor one on the first floor almost at opposite ends of the house. House is about 3500 sq ft.
 
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
I figured it out. It was a bad AP apparently. Whenever I put any amount of traffic through it from speedtest.net I noticed speedtest would freeze part way through, the AP would stop broadcasting, and I’d switch to the Carter AP. After about a minute the bad AP would come back online and since the signal was stronger the client would connect to it. This cycle would continue indefinitely. I replaced the cable first, same behavior. Grabbed another AP and it appears to be working.
Now to figure out what’s wrong with the bad AP. Posting this in case it helps someone in the future.
Also thanks to everyone with suggestions. I’ve tried basically every combination of transmit power, min RSSI, every AP configuration possible.
Normally when this happens you have a POE issue, not enough power for the AP.
This is the longest of the runs in the house but it’s well under the max length. It’s plugged into a UniFi 24 port switch. I’m going to try it with a shorter run in the basement after I factory reset it.
That used to happen to me when I used an underpowered PoE+ injector to power a U6 Enterprise: the AP was more or less working fine but as soon as I tried a speed test it would reboot. I fixed that by switching to a Unifi U-POE-at injector instead.
Thank you for the update!
Sounds like you need to turn down the power on both APs a bit. You might be in the middle of the signals and they are competing for the device.
Id definitely try this, have lots of times seen where APs are too close and turning down the power helps the roaming select a closer AP.
I didn't see it until after my initial post but his current wifi indicator is nearly zero, which tells me he's right on the edge of that AP. Adjusting the power should definitely help here. Plus who knows what's in-between him and the ap, concrete walls, electrical, etc
Will that just move the border between the APs?
Roaming is done by clients. This support document outlines how iOS devices roam. i have not found a comparable guide for any android device.
https://support.apple.com/en-gb/HT203068
Samsung documents its roaming behavior here. Other Android devices may behave differently.
https://docs.samsungknox.com/admin/knox-platform-for-enterprise/kbas/kba-115013403768/
So if I set my minimum rssi to 75db a Samsung phone would be forced to find a better AP?
Looks like OPs device went for ax over ac.
Despite what Apple has in their documents, I've found the trigger to be closer to -80. Setting a minimum RSSI has helped keep my iOS devices on the better AP.
Where is this in the app?
It's the WiFiman app.
Have to ask your device manufacturer.. lol. This is one of life's mysteries to me tbh. I have some devices that constantly bounce around, others find the worst one and just stick with it or constantly seem to hunt down the worst one to connect to, and some that do what they are supposed to... They behave like my 3 kids sometimes. Drives me nuts and the only way i could stop it was to specify connection rules for the AP to no allow it to connect if the signal was below a certain threahold or speed was under a certain limit.
Don’t listen to him, you’re doing a great job AP!
Have you enabled BSS transition and Fast roaming in the wifi network settings? With BSS transition on the AP can send the client a signal to disconnect from it and "suggests" another AP to connect to. Fast roaming is 802.11r and it enables the client to switch APs more quickly without annoying disconnections
100%
802.11k is on regardless.
802.11r = Fast Roaming in the settings.
802.11v = BSS Transition in the settings.
Had similar issue to OP. I think the combination of tuning the power a bit and turning off BSS Transition has helped somewhat.... But in an apartment setting, it's hard to prevent overlap of AP coverage....
Might try tuning your minimum RSSI settings.
What are the AP models involved ? Apple devices, by design, always prefer Wifi6 over WiFi 5. You should avoid mixing wifi6 and wifi5 APs for a seamless roaming experience with Apple devices.
This is why managed APs were such a hit thing for so long.
I set minimum strength to connect to AP sometimes at like -70 db to ensure this never happens. Not saying it is a good solution but works sometimes.
