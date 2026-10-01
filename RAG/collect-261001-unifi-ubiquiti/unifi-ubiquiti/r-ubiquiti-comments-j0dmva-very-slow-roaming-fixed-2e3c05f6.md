---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-j0dmva-very-slow-roaming-fixed-2e3c05f6
title: "r-ubiquiti-comments-j0dmva-very-slow-roaming-fixed-2e3c05f6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS", "Apple", "Samsung"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-j0dmva-very-slow-roaming-fixed-2e3c05f6.md
source_anchor: ""
source_lines: [1, 32]
sha256: 8ed89127e5d006edfe5fabd10f6971b0c33e791fa102c6afe3702cefb8a81d61
---

# r-ubiquiti-comments-j0dmva-very-slow-roaming-fixed-2e3c05f6

Very Slow Roaming (fixed) 
        
    2 months ago I started receiving complaints of dropped video calls and failed downloads while moving from one side of the house to the other.
Did some testing and roaming between my two Ubiquiti access points (UAP-AC-LR and UAP-AC-M) was really slow. I tried every solution I could find online... minimum RSSI, different channels, stopped using the fast roaming option, etc, without success.
Today, after another complaint, I decided to start fresh. I deleted all wifi networks, created them again, and voilà, roaming is fast again (even with the "fast roaming" setting disabled).
I don't know what was causing this weird issue, but this fixed it for me. Sharing it here in case someone is experiencing the same problem.
Section des commentaires
I thought roaming was a function of the wireless client. Someone correct me if I'm wrong here?
We can speed roaming up by using the "Fast Roaming" option (it doesn't work well with some devices), but yes, it's the client that takes care of that*.
In my case, I'm inclined to think that it wasn't a client side issue because it was affecting all devices I tested (1 iPhone, 3 Asus/Samsung/Android phones, a Macbook laptop). Recreating the WiFi networks fixed the issue without any client-side change and also lowered the "retry rate" from around 40-50% to 20-30%.
It was a weird problem and after trying everything I could think of, this was what fixed it.
[* The old Zero-Handoff that Gen1 APs used was the exception.]
I personally disagree with people giving client is entirely client based. Because those people probably do understand the detail but not giving the detail so misinformation/lost in translation happens.
I call client DRIVEN AP dependent process. I love analogy though this may be a bad one. Client is a driver AP is a car. In the end driver decides when to drive, how to drive so on. But car matters and if car runs out of gas, if you don’t have the right key or car decides to lock you out, you can’t drive the car.
Most clients are slow to roam since early WiFi specs usually called for them to try to step down the data rate to the basic rate in a series of steps; when encountering retries. Finally they roam (at some point). It’s often tied to the client programming deciding it’s had enough with its associated station.
Spectrum interference can cause havoc too.
Once the client has had enough; it can try associating with other stations in its existing BSS list or it can issue a probe request and just start over. I have seen both and they work differently depending on the situation.
Round a corner outside a building and you are wasting time trying to associate with every station in your BSS list because they are....now gone. Atheros I’m looking at your rubbish baseband LOL.
In a home situation you are way better trying a BSS list alternate. Even then, most client vendors set a bias where the new station has to be so many dBm stronger or they won’t roam (retries be damned). It stops ping pong effects in boundary areas which can be a problem. Like killing a fly with a hammer.
The station (AP) can also do its own individual or coordinated load balancing and deny associations as it can be setup to deny roaming if it’s busy. I’ll just call this steering the clients.
Some vendors for clients (Laird) keep the initial data rate low after an association; unless there is a data burst requiring more speed. It allows for better range with less errors. Better user experience - unless you are just looking at link speed wondering why a idle client is slow and you feel it shouldn’t be because other clients go for max link speed then have retries.
In short, it’s really very complicated, every situation is unique. I have been supporting this kind of "fun" for a long time (25 years) at multiple enterprise accounts. All the way back to crystal based 900 MHz with a site wide speed of 128kbps. Character based emulation using UDP controllers that held open TCP connections for wireless clients. Woot woot woot. Fun times.
I hope this explanation actually helps and doesn’t further muddy things.
What is a “BSS list alternate” and how do I try it?
With UniFi in a small environment when in doubt reset everything and setup a new controller 60 percent of the time it works every time
I don't remember who it was (Amazon AWS? not sure), but they noticed that some of the weird issues with their machines could be fixed with a reboot... so instead of finding the root issue, that's what they did.
I guess a reboot or a clean install/reset from time to time helps fixing these weird issues :-)
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic and picture posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
