---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-creating-vpn-between-2-usg-on-1-unifi-controller-86a39224-529d-4e21-a5-6c8cecee
title: "questions-creating-vpn-between-2-usg-on-1-unifi-controller-86a39224-529d-4e21-a5-6c8cecee"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-creating-vpn-between-2-usg-on-1-unifi-controller-86a39224-529d-4e21-a5-6c8cecee.md
source_anchor: ""
source_lines: [1, 49]
sha256: f0ecbd979c6d1a944bc45763463b1150f2648b464c7cbf65e23395b113b1f810
---

# questions-creating-vpn-between-2-usg-on-1-unifi-controller-86a39224-529d-4e21-a5-6c8cecee

@UI-Team
Hello,
We are currently setting up a connection between 2 of our branches with USG (Unifi Security Gateway)
My question is: is it possible to connect the USG of site b to the Unifi Controller of Site A. if so, what are the steps to make this work?
Site A with Unifi Controller -> USG 1 -> INTERNET <- USG2 <- Site B
I want to make this connection so i can manage site B AP / Camera's etc. via Site 1.
Kind regards
Bram
Hi Bram,
Oh, alright. Thanks for clarifying! The USG supports L3 adoption just like other UniFi hardware. So it probably would be simplest to use SSH to achieve this. You may need to open port(s) to perform the initial adoption (not sure if you already did for the UAPs at the remote site).
Connect and issue:
mca-cli
After that you will drop to the UniFi# console. From there issue:
set-inform http://ip.of.controller:8080/inform
After the initial adoption you need to issue the set-inform again (it will say this in the shell session too). This is important as saves the remote inform IP.
Once it's adopted and the tunnel is up, you can remove the port forward(s) that you don't need any more.
Hope that's helpful.
Cheers,
Mike
Both sites need to be in the same UniFi controller. Go to Settings>Networks>Create Net Network>Site-to-Site VPN. Then choose the appropriate site from the remote site drop down.
The cameras have their own controller, and would be accessible one the VPN is up.
Hi Mike,
yes I have already done those steps, the only issue that I encounter is that I can't connect my USG which has a different IP (other location, 3miles down the road) to my Unifi controller. Any idea how this is done?
kind regards
Hi,
Thanks for this post, it has really helped.
What I have found is that if you set up the "Main" USG first. This is the one with the UniFi Controller behind it and make sure you set port forward for port 8080 to the UniFi Controller.
Then when you set up your USG not matter where you are you http to the new USG and once you have configured the internet settings there is a "info" button bottom right. If you change the settings here and replace the word "unifi" with the WAN IP address of your Main USG.
The USG will then appear in the UniFi controller and you can adopt it.
Amazingly it worked first time for me
You then set up the VPN network under Networks by adding the network as a VPN and select the site.
I'm trying to do something simular.
I am installing a campground with 5 ISP PoPs spread throughout.
Each PoP has it's own modem and WAN address.
I have a USG, a USW and a CloudKey at the Main Pop, no problem.
At PoPs 2 through 5 I have a USG and a USW.
I SSH into the USG at PoP 2 and Set:inform to the main PoP address.
it shows up on the controller and says "Pending Adoption", but there is no button anywhere to make the controller adopt it.
I have more questions, but hopefully someone can get me past this first one.
Thanks
Are you logged in as Super Administrator, and are you trying to Adopt the USG into a new site (one without a USG adopted already)?
One USG per site...then you can move onto setting up the VPN between the sites.
Yes I am logged in as the Super Administrator.
There is a USG already adopted.
Do I have to create a new site for each PoP?
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
