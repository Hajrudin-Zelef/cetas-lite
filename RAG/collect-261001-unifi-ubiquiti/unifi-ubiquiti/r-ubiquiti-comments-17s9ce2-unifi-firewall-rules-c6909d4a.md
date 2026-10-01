---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-17s9ce2-unifi-firewall-rules-c6909d4a
title: "r-ubiquiti-comments-17s9ce2-unifi-firewall-rules-c6909d4a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-17s9ce2-unifi-firewall-rules-c6909d4a.md
source_anchor: ""
source_lines: [1, 39]
sha256: 3b2bf49befe9f24a63c6379b70de8f08cf6d54049c38689dec53d7d4c18f6f64
---

# r-ubiquiti-comments-17s9ce2-unifi-firewall-rules-c6909d4a

UniFi firewall rules 
        
        
        
    
    
    Ok so I have a UDM Pro and id like to start using the firewall rules. However I'm very amateur to this topic.
I have 4 Vlans set up.
- 
      main
- 
      iot
- 
      cameras
- 
      Plex server
The rules I'd like to establish for each.
Main needs to connect to everything
Iot
Internet in access Internet out no access Local in access Local out no access
Cameras
Internet in no access Internet out no access Local in access Local out access
Server
Internet in access Internet out access Local in access Local out access
Does that look about right or do I have something backwards?
If I do have something messed up could someone help me through this? I'd like to actually learn how this works vs just regurgitating what's in a video. Or if someone has a video that steps through this completely I'd love it but the ones I've found just say do this, this, this and that.
Thanks
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
It really depends on the iot and camera devices you have. For iot do you have Alexa or Google home that requires outbound internet access to work? Will your cameras ever need updates?
In my world my cameras are blocked from the internet in both directions, but they are allowed to talk to the NVR. The NVR is allowed to talk to the internet so that I can access my cameras remotely.
It all depends on the devices you have and your use case.
I'm figuring out the iot stuff as I go. The cameras are Unifi.
Check here. The articles on "Introduction to firewall rules" and "Traffic rules" will be the most relevant.
You can watch this video for most of what you are wanting to do. I did help me when I got my UDM Pro that same year. The rules are still relevant today.
