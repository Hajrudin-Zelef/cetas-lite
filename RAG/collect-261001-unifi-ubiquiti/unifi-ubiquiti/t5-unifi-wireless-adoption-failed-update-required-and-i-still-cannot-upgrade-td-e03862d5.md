---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/t5-unifi-wireless-adoption-failed-update-required-and-i-still-cannot-upgrade-td-e03862d5
title: "t5-unifi-wireless-adoption-failed-update-required-and-i-still-cannot-upgrade-td--e03862d5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/t5-unifi-wireless-adoption-failed-update-required-and-i-still-cannot-upgrade-td--e03862d5.md
source_anchor: ""
source_lines: [1, 52]
sha256: 1530deee82cabc9890ba82df2dbc4a622427041a82125cc38f90816e72ba0e20
---

# t5-unifi-wireless-adoption-failed-update-required-and-i-still-cannot-upgrade-td--e03862d5

@UI-Team
Hello everyone,
I recently was on business trip for a while and when I came back, my parents were asking me why their Internet did not work!
I took a look at the Controller and they were all out of date. So I update the firmware of the Cloud key (gen1) and I now have the latest Unifi Controller (version 5.10.19-11646-1).
I have 3 UAP-AC-PRO at my parents' house and when I login into the Unifi Controller interface, in the "Device", it is showing that 3 of the devices are "Adoption Failed (Update required)"
So I "forgot" the device and try to adopt it again. Got the pop-up windows saying that once the adoption is done, my UAP-AC-PRO will have the latest firmware, etc.. etc..
But that never come... The 3 UAP-AC-PRO still remained "Adoption Failed (Update Required)".
I came across another article here that suggest we should SSH directly into the UAP-AC-PRO and upload (for force update) the firmware from here. My issue with this method is, I tried the username/password ubnt/ubnt but it did not work, saying "Access denied"
Even if I put my own login credential into the UAP-AC-PRO, it still says invalid.
Anyone knows any other way to upgrade the firmware for this UAP-AC-PRO?
Thank you very much
James
Hello @jamesnb ,
Yeah, then someone has to press the reset button for 10 seconds..
You can also try caching the firmware on the controller.
Settings > Maintenance > Firmware
Regards,
Glenn R.
The UAP AC ( v2 ) is only supported on 5.6.x
What firmware version is it on?
Are you sure its a UAP AC? ( squared unit )
Hello Glenn,
Sorry for the typo, it is actually UAP-AC-PRO and I have editted all in my original post.
Still, could not get those UAP updated...
Anyway that I can SSH into those and force the update?
Thank you
Yes you can update them via SSH, I recommend following the upgrade path below:
Firmware Before 3.7.58 > 3.7.58 > 3.9.54 > 4.0.21
Use the following commands to upgrade the UAP via CLI.
wget <firmware_binary_link_location> -O /tmp/fwupdate.bin
syswrapper.sh upgrade2 &
or
curl <firmware_binary_link_location> -o /tmp/fwupdate.bin
syswrapper.sh upgrade2 &
Hey Glenn,
Thank you for your advice.
My issue is that, I could not even SSH into the UAP-AC-PRO. I can ping all of them but when I tried to ssh with ubnt/ubnt or root/ubnt login credential, it did not work. It keeps saying access denied or permission denied
I even use the login credential of the Unifi Controller, it still did not work... Hence I cannot SSH into it
Do I have to ask someone to press the "reset" button on the device itself? I am currently not at the site of the UAP-AC-Pro though.
Thank you for your advice anyway
I have updated firmware with using Putty.
Unifi controller still shows version 3.7.49 and last seen: never
And status is adoption failed (update reguired)
sudo syswrapper.sh restore-default
Now it works. IP have to change back. (Static IP)
I installed this version (UniFi AP-AC 5.6.42), and it worked for me just fine. Here's the link: https://dl.ui.com/unifi/5.6.42/UniFi-installer.exe
@jamesnb wrote:
You may already have the answer to this, idk, but the way I ssh into something is if you go to Site in settings and enable advanced features (think you need a USG) then go onto device authentication you can ssh into the device you need to with the credentials it gives you. :)
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
