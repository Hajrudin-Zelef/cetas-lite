---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-need-to-incorporate-udm-pro-unifi-dream-machine-into-existing-network-348a0fc7
title: "questions-need-to-incorporate-udm-pro-unifi-dream-machine-into-existing-network--348a0fc7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-need-to-incorporate-udm-pro-unifi-dream-machine-into-existing-network--348a0fc7.md
source_anchor: ""
source_lines: [1, 37]
sha256: 0639921a8689dd5e61a0a7e48cb1f7a7d9df3f2a831f4110545dee5c56af532d
---

# questions-need-to-incorporate-udm-pro-unifi-dream-machine-into-existing-network--348a0fc7

@UI-Team
Hi All,
I have a unique configuration that I need some serious help with.
I have a school where I have fourty two (42) nanoHD access points installed connected together by two USW-Pro-24 switches and a US-16-XG switch used like a fiber backbone all running on a base IP range of 192.168.100.x. The Internet is provided through a large SonicWall firewall that performs substantial content filtering. The SonicWall provides DHCP and the level of content filtering is based on VLAN IP ranges that are tied to SSIDs associated with different classes of wireless users. The system has an external Unifi controller running on a dedicated PC. The system seems to be running well. I want to incorporate a UDM Pro Unifi Dream Machine in order to start getting more detailed information about the traffic usage. I do not need the UDM Pro to do any filtering or DHCP services. I just want the reporting.
It is my understanding that the UDM Pro has a built-in UniFi controller and there cannot be two different local controllers. Correct?
Do I need to setup the UDM Pro, , disable DHCP, change it's IP address and do a restore from the backup of the original UniFi controller to get it moved over?
How do I make the UDM Pro not do any of the filtering and just let the SonicWall to handle all that?
I don't want to have to return the UDM Pro but will do so if I can't get the granular reporting that it is supposed to provide.
Thanks
Dave
With mixed sonic wall and UniFi setups, I always add the vlans off of the x0 interface as sonic walls wlan interface is awful and breaks so much stuff. The wlan subnet is really geared for their awful sonicAP devices
I'll see if I can pull up some screenshots of how I've got my customers with a mixed sonicwall & unifi setup. But pretty certain just doing the following will fix all your issues:
X0 - create new virtual interface. Create your vlan range. Assign a DHCP scope to it. (No relay needed)
In unifi - just make sure your main corporate network matches x0 range. Then create as many vlan only networks as you need and allocate accordingly.
I don't think this is a good idea. The UDM Pro is not designed to handle that many APs with switches as well. I feel like you are going away from something that you said is working to something that will just give you nightmares constantly.
The reporting isn't even close to accurate anyway...so you won't be getting too much out of it, in my opinion.
Hmmm. Your observation about the reporting not being accurate is interesting to me. Can you (or anyone else for that matter) elaborate on that? This is my first appliance of this nature from Ubiquiti. It's a trial run kinda thing. I have a lot of experience with SonicWall. Their flexibility is great but to get reporting from them is expensive and a pain to setup.
What kind of reporting are you looking for exactly? What is it you want to see?
The only extra reporting you'd get from a UDM-Pro is based on traffic going through it as the Gateway which isn't going to be the case if you are using a SonicWall without introducing double NAT (It doesn't support any sort of Disable NAT or Stealth Mode).
@verisarioc Ummm i have over 50 devices being used on my UDM pro and it works fine. Please provide where you have read this before. @StLukeChurch I would do exactly like you have already stated. Set up the udm pro from the backup config, then disable dhcp, then start to plug everything in. Leave everything disconnected from the UDM pro while setting it up so it doesnt start to give IPs out
Also you dont need to enable the firewall settings on the udm pro, that will allow the sonic wall to do everything
@JD34 wrote:
Its just what I have seen around this community site. It seems like people that want to use more than 40 APs and some switches the UDM Pro is having issues with. UI doesn't give a definitive answer so my guess is as good as anyone's.
Glad you are one of the few having good luck with that many devices.
There are many users in this one thread that back that number up. devil IS in the details here
number of AP's and switches AND users matter.
@_Space So reading your link, it does state it depends on user usage. did the OP state he had a lot of users? After i contacted support they said they havent released a number yet but 50+ hasnt been an issue. Also he wouldnt being using the UDM Pro to full untilization. AGAIN he is not using DHCP and the fire wall. He already have a controller in use that is handling it currently, i no see how udm pro will not work
I still think its a terrible idea to introduce the UDM Pro. But thats just me. If he does it and it works..fantastic. But I would hate for him to come back saying its junk as a lot of people have noticed. (Sorry, not a fan of it. I have the SE model and it seems "better" but thats not saying much.
I just pasted a link to people in the real world using the product. as I stated devil IS in the details. I did not say anything would or would not work. Just a link to a thread where this topic was gone over.
you dont need to put AGAIN in all caps for my benefit.
key item may YOU MISSED (hehe)
so ALL traffic WILL pass through it. so while the UDMP may not be doing DHCP, it will be tracking all of MAC's etc. it WILL add to the overall utilization. it's NOT just a UDMP for controller only...
@StLukeChurch it may work for you. but based on some peoples usage you are possibly nearing a limit.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
