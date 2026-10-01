---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-15bzda1-best-method-to-manage-multiple-small-sites-344fa723
title: "r-ubiquiti-comments-15bzda1-best-method-to-manage-multiple-small-sites-344fa723"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-15bzda1-best-method-to-manage-multiple-small-sites-344fa723.md
source_anchor: ""
source_lines: [1, 32]
sha256: 4824511f914edd75d201688e5e842d565cd67aaca4882f95062fee74b3e94fec
---

# r-ubiquiti-comments-15bzda1-best-method-to-manage-multiple-small-sites-344fa723

Best method to manage multiple small sites? 
        
        
        
    
    
    Hi,
What is the the best way to manage multiple sites that are mostly only made of up to 10 APs?
I recently set up a site with 4 APs, one MT router and a POE switch. Controller installed on my laptop, everything is working perfectly and I can go on-site and make changes if needed.
Now I have to set up another small site, 3 APs, totally different location. I need a separation between these sites and their controller. Can I do this on a single laptop with a single instance of Unifi Controller (Now renamed to Unifi Network App)?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Few options:
UniFi cloud console - monthly fee, but very reliable and first party
run a cloud key Gen 2+ at one site and connect all sites back to that
run a self hosted controller in the cloud
run a self hosted controller at one site and connect all sites back to that
work with third party to host controller for you.
I personally offer the last option, PM me if you want to know more.
I see. I do have some servers on wich I host some docker containers and some VPSes, unifi in a docker looks like a solution.
Currently I'll only have 2 sites and I am in the near location so I'm looking how to manage these two sites individually with my laptop, on premise. Right now I have site no1 in my Unifi Network app on my laptop, how would I create a new independent site in the app? Multi site?
Enable multi site in settings
I STRONGLY recommend not using a laptop to host the controller, devices do not like when the controller changes IP’s or is connected via wireless. Your better off get a raspberry pi for each site than continuing down your current path
You can but if your laptop isn't on site then you won't be able to make changes.
You'd be far better setting up a VM on a cloud service and then adopting them to that. That way you can access the VM from anywhere and make changes to the sites without having to travel to them.
If you don't fancy running that yourself then there are hosting services for Unifi controllers.
The other alternative is to put UDM's or UDM Pros into each site. You can then use the inbuilt controller and then connect them all to a ubiquiti cloud account and remote manage them and any of the network devices behind them, like Unifi switches, AP's etc.
Feel free to ping me any questions.
