---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-ap-ac-pro-stuck-in-adopting-update-required-wireless-2bb53960-6a-70795b24
title: "questions-unifi-ap-ac-pro-stuck-in-adopting-update-required-wireless-2bb53960-6a-70795b24"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-ap-ac-pro-stuck-in-adopting-update-required-wireless-2bb53960-6a-70795b24.md
source_anchor: ""
source_lines: [1, 22]
sha256: c043056f04dd07fb35515498045ec0f648d0042bec42a2f261deed449280fd8a
---

# questions-unifi-ap-ac-pro-stuck-in-adopting-update-required-wireless-2bb53960-6a-70795b24

@UI-Team
I've tried everything from reseting to factory default, forgetting and re-introducing, to starting rolling upgrade and scheduling upgrade and googling, with no success. Can someone advise before I return the AP?
might be the same thing I just went through with some new uap-ac-m's...
Try forgetting them, then use a toothpick and press the reset button for more than 10 seconds (Factory reset), then connect them back to the network and see if they show up to adopt again.
@jslande01 I also tried that, unsuccessfully... cant even do it in the aforementioned state as the "Forget" fails!
@UI-Glenn no, I have not... will try now...
thanks!
@UI-Glenn my AP appears to be stuck in a state where UNMS thinks its being adopted, but waiting for upgrade (therefore UNMS is not allowing it to be forgotten) and the AP, after being reset to factory default, is waiting to be adopted... so i'm completely unble to proceed according to the process you shared. Help!
UNMS? You mean Unifi Controller? Still waiting for the version number, BTW. There may be controller specific issues we need to address. Can't really help without all the info but try a factory reset and while it is doing so, forget it from the controller. As long as it is power up and talking to the controller, you will not always be able to forget it. SSH into the AP and perform a manual upgrade using the link Glenn provided.
@UI-Glenn @jslande01 after about 15 cycles through trying things, magically (?) I was able to un-adopt, update and provision the new AP successfully... thanks so much for your pointers and advice! Great community!
I am seeing the same thing with an AP Pro (3.4.14.3413) and UDM-Pro (1.7.2 with Network Controller 5.13.30)
I have done:
What does not work:
What I see:
Any ideas?
This was an odd fix for me and I hope someone can explain:
Here is the fix:
I am thinking it really did not have an IP address and that is why I could not ping or SSH.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
