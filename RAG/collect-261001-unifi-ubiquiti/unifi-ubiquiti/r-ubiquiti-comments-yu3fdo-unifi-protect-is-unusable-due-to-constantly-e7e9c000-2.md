---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000-2
title: "r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000.md
source_anchor: ""
source_lines: [7, 54]
sha256: e576ac444a8cb9e796eac5514dc9e04bf762148a292f871084de33e5a27ae875
---

# r-ubiquiti-comments-yu3fdo-unifi-protect-is-unusable-due-to-constantly-e7e9c000

    I recently purchased a Network Video Recorder to pair with my Dream Machine Pro. For now I have the two connected via ethernet (until I can get a 10g cable). I have configured my UDMP with a separate VLAN and Wifi SSID just for security cameras. The UNVR has four Seagate Skyhawk AI 16TB drives. Those took nearly a day to format. After they finished formatting I hooked up a single G4 Instant via Wifi to start playing with Unifi Protect. That's when all the issues started to happen.
Clips sporadically won't load or play. Footage is being triggered and recorded just fine, but when I go to playback clips from either the web or my iPhone, they load/play sporadically. There's no rhyme or reason to it. Sometimes clips won't load at all, they just have the 3 loading dots continuously. Other times they play right away.
Next, I can't seem to export any clips from Protect iOS. I just see "Calculating..." and an error message pops up "Could not download selected clip".
Lastly, exported clips from the web interface play choppy on my laptop. This one is odd, b/c they play smooth on the web interface, but when I export/download them and play them in Quicktime or Finder, they're choppy.
Here's my setup...
- 
      UNVR drives: 4qty Seagate Skyhawk AI 16TB (ST16000VE002)
- 
      UNVR has Unifi OS: UNVR v2.5.11 and Unifi Protect 2.2.6
- 
      UDMP has Unifi OS: UDM Pro v1.12.30 and Network 7.2.95
- 
      iPhone 14 Pro Max: iOS 16.1.1 and Protect 1.8.0 (590)
- 
      MacBook Pro 2019: Chrome v107.0.5304.110 (Official Build) (x86_64) on Monterey 12.6
I'm extremely disappointed in the entire Unifi Protect experience, especially for the cost that I've incurred. My Wyze and Ring cameras offer a much better experience in terms of just being usable. There's no way I can trust this system to be reliable if I can't even retrieve clips. This can't be normal, yet I've read lots of you all having similar issues.
Any tips or advice before I return all this stuff?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Intermittent/performance issues like this can be tough to troubleshoot. It could be anything, like a bad cable, poorly mated connector, dust, or anything else. You can use tools like iperf to try to isolate problem areas in your network.
Regarding the choppy playback on the web interface, I had the same problem. From what I've read, the Protect web interface is not designed for regular camera monitoring. It doesn't delay/buffer the stream for seamless playback, and the result is the stuttering effects that you see. I don't think this is a good excuse but that seems to be the way it is. A skeptic might say this is to create demand for their dedicated device to display the camera feeds. Regardless, you have other options like what are described here: https://www.reddit.com/r/Ubiquiti/comments/uu3ov6/how_to_locally_stream_rtsp_on_vlc_guide/
For me, the stuttering problem was localized to the Protect web interface. I had no issues using the Android app, which is coded to correctly buffer the livestreams from the cameras. So you may have multiple issues at play. Try running the instructions at the link above, to be able to isolate whether this is a problem with your local network vs a problem with your ISP connection. Keep in mind that when viewing camera feeds on your phone app, it's sending those camera streams to the cloud, then back to your device, so you need an ISP with a reasonable uplink speed (often the pain point for ISP offerings)
Thanks, I will check that thread and try out the RTSP stream. I'm not currently doing that, b/c mostly what I'm trying to do is review and download existing footage (not that interested in live viewing). I'm simply trying to get the super basics to work. Segregated VLAN, single wifi camera, viewing and saving clips either on the web or my iPhone.
FWIW, I have Google Fiber 1gbps up/down. Just did a speed test and it's damn near that.
Are you using raid 5 or 10? I would try raid 10 in “half of disks”.
I did RAID 5 "One Disk". I'd prefer RAID 5 so I have more storage. Would the RAID 10 striping really make a difference on 1qty 2k camera?
Sorry I miss read, I thought you added a camera. Yeah you should not be having these issues with just a single camera. Is the camera disconnecting from wifi or has a bad signal? My g4 doorbell has really terrible playback compared to all my wired cameras. I cant really do much to fix the signal other than replace it with the wired g4 pro due to its location/brick exterior wall.
I have the same issue after a somewhat recent update.
At least there's hope that the system isn't complete trash. I only opened the G4 Instant and the UNVR to test the whole Protect thing out. So far I am REALLY not impressed. I want to like it, but if I can't easily review and download/export videos, what's the point!?
I had choppy video, which wouldn’t show at full resolution, when I first setup.
It appears the issue was (likely) caused by me.
When you install the HDD you should do this with the power off, then power on, and allow NVR to format the disk. Interrupting this process can cause corruption, such as turning it off whilst still formatting.
Not only did I install with it turned on, but I then turned it off shortly after, which clearly buggered up the disk. Nothing I did could resolve it.
I got a replacement disk, installed correctly, allowed it a few hours to format, then started recording. Not had an issue since.
Obviously this is just one of multiple possible candidates though. It could be a bad cable too. One of my cameras would only show FE connection, when it was meant to be GBE. I chopped off the RJ45’s at each end and redid it, fixed the issue.
Interesting! I definitely installed my HDDs with power off. I didn't let them fully format until I started recording, b/c that was slated to take 32 hrs. I got impatient and hooked up a single G4 Instant. That being said, they're formatted now and I could easily reformat as I don't care what footage is on there now.
I did swap out the ethernet cable and that didn't seem to affect anything. Thanks for the tips!
Use the iOS app on your phone and see if you’re getting a direct connection. If you’re not, then something is screwy with your VLAN settings that is messing up your direct connection to the UNVR. If you intentionally aren’t allowing a direct connection to the UNVR, I’d ask why?
I wasn’t intentionally at first. My FW rules had blocked it, but it’s fixed. Protect iOS app is showing the green lightning direct connect icon now.
Cool… any improvement to performance when viewing recordings from the iOS app? I’d start there, and if everything seems good then move to the web via direct IP and then unifi.ui.com.
Please share more information and any existing support ticket numbers here: https://community.ui.com/social-feedback so we can properly escalate, review, and assist. Thank you.
Done!
The primary issue now is that when I export videos from the web interface or iOS, the downloaded clip either plays back extremely choppy, or it's the wrong clip, or it's an hour of footage vs 10s.
All of this seems to point to an issue with the database. Is there a way to check it, repair it?
