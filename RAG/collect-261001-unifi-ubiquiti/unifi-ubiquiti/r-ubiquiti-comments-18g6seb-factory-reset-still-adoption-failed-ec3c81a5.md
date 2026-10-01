---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-18g6seb-factory-reset-still-adoption-failed-ec3c81a5
title: "r-ubiquiti-comments-18g6seb-factory-reset-still-adoption-failed-ec3c81a5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-18g6seb-factory-reset-still-adoption-failed-ec3c81a5.md
source_anchor: ""
source_lines: [1, 23]
sha256: 1501cff6f1788334ec2d7b5ee5f710c360bce0e4aa1eb2ee1fc321c99f6eb563
---

# r-ubiquiti-comments-18g6seb-factory-reset-still-adoption-failed-ec3c81a5

factory reset still adoption failed 
        
        
        
    
    
    I purchased the Unifi Express and am swapping APs over to it from Cloud Key Gen 2. Could not export and import the config cleanly. Cool. Factory reset my two APs. Still failed to adopt on the Unifi Express. Multiple factory resets and still adoption fails. Why doesnt a simple factory reset allow the AP to be adopted by the new controller?
Just for the record I also removed the APs from the cloud key. I see a lot of different suggestions but a factory reset should allow the devices to be adopted, right?
UPDATE ---
If someone stumbles upon this - it seems that i was being impatient? After the factory reset and multiple failed adoptions, I unplugged both APs from the POE power and plugged them back in after 5 to 10 minutes. Both were adopted easily.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
If someone stumbles upon this - it seems that i was being impatient? After the factory reset and the failed adoptions, I unplugged both APs from the POE power and plugged them back in after 5 to 10 minutes. Both were adopted easily.
You must have held the reset button for more that 10 seconds, so the APs went into DFU mode. Yep, rebooting them would get them back to a state where they could be recognised by the controller.
ok cool, thanks for the clarification.
Are the APs firmware up to date? What model are they?
If they are not up to date with the latest firmware, you may have to log into the web interface for each and update them manually. Once updated, try to adopt them again.
so logging into the web interface would entail me hard coding an ip on the 192.168.1.x range and sshing into them?
I appreciate you giving me assistance but damn man this is kinda corny. i didnt think i was gonna have to spend my evening swapping APs and it shouldn't be this difficult. Especially after factory reset.
