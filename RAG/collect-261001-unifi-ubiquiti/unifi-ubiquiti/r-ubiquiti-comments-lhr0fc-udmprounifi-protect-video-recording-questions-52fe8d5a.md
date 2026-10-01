---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-lhr0fc-udmprounifi-protect-video-recording-questions-52fe8d5a
title: "r-ubiquiti-comments-lhr0fc-udmprounifi-protect-video-recording-questions-52fe8d5a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-lhr0fc-udmprounifi-protect-video-recording-questions-52fe8d5a.md
source_anchor: ""
source_lines: [1, 41]
sha256: 898169bfdacec9002013fe9b6023f29deb9ae2860d9d2e987fb1b631043c4ccf
---

# r-ubiquiti-comments-lhr0fc-udmprounifi-protect-video-recording-questions-52fe8d5a

UDM-PRO/Unifi Protect - Video Recording Questions 
        
        
        
    
    
    Hello,
I apologize in advance if any of these questions have been answered. I'm new to Ubiquity equipment and NVR's.
I recently made a huge upgrade from a few Nest cams to a UDM-PRO for my home. I purchased a 16TB Seagate Skyhawk AI Surveillance HDD to go along with it. I currently have (3) G3 Instants and (2) G4 Doorbells and have plans to add 5 more G3 Instants.
My questions are about ideal recording for this setup. I have played with the record on motion settings and found that it doesn't seem to capture everything even when settings are adjusted to -Minimum seconds 0 - Before motion 10 - After motion 10. On my nest cams I would often go back and save clips when my toddler did something that I didn't have time to capture on my phone. So here are my questions
- 
      Are there other settings that might help with motion only recording?
- 
      Is it ok to Always Record with that many Wi-Fi cameras? (I have a 1Gb connection)
- 
      If I Always Record, will I be able to retain a minimum of a rolling 30 days of footage?
- 
      I'm trying to understand the Recording Quality setting more in depth. I have it set at 50% for all cameras right now, so I'm assuming it streams in 1080p and records in 720p?
Thanks again for any help offered.
Section des commentaires
Scamming the system with multiple accounts to get so many G3 instants or are you grabbing them off the black market? :)
You can toy with increasing the record before motion setting as well as after, but most find recording 24/7 to be much more reliable. 16TB should easily get you 30 days, specially for relatively little movement scenes. I get more than 45 days with a similar mix of g3s on a 6TB drive. If not satisfied with retention, you can drop the frame rate to 15, you probably not be able to tell the difference.
The quality setting changes how aggressive the cammera compresses the video. Doesn't affect resolution, although using low quality settings you probably will see nighttime compression artifacts and anytime with fast movement.
There are so many variables that accurately predicting drive space for a period of time for a set of cameras is impossible. You can certainly calculate worse case as you know the maximum bit rate each camera can generate. But the bit rate varies with scene. So anything from the minimum steady state bit rate and max bit rate is a subjective estimate. So buy as much storage as you can afford, and adapt.
Haha my one account and 2 overpriced ones on EBay!
Thanks for the long response that helps me a lot!
none of your streams hit the internet, the only thing is you are eating up is more HD space
Depends on how many cameras, the resolution, etc
see this post
https://community.ui.com/questions/Unifi-Protect-calculator/e97f22d9-ab5f-44f7-9443-bdb212ea7cc7
Some kind of Protect calculator would be pretty nice thing to have when it comes to figuring out storage needs
Duh! IDK what I was thinking there! I guess I was thinking of reprocussions on internal bandwidth
10 Total Cameras (8 G3 Instants [1080p] & 2 G4 Doorbells [2mp]) set at 50% recording quality
I understand that this will vary, I'm just looking for a ballpark estimate since I have no baseline being new to NVR systems.
While I use the "Unifi-Video" controller installed on an Ubuntu box, we have 220 Cameras (98% are G3 bullet cams), all set to record on motion 24/7 (and we get a ton of motion, these are at a manufacturing plant) on 'low' quality. We get 4 weeks on a 16tb drive (iSCSI to a NAS).
Your 10 cams on 16Tb should get you a lot of footage history.
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic and picture posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
