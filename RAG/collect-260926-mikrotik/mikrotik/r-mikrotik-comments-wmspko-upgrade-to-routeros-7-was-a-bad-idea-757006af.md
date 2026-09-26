---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-wmspko-upgrade-to-routeros-7-was-a-bad-idea-757006af
title: "Upgrade to RouterOS 7 - Was a bad idea"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-wmspko-upgrade-to-routeros-7-was-a-bad-idea-757006af.md
source_anchor: ""
source_lines: [1, 115]
sha256: 836d2b3c98b045ff3f6f90b76db5abbbce62520bb01036ddbe755dccd24fef51
---

# Upgrade to RouterOS 7 - Was a bad idea

A community-contributed subreddit for all things Mikrotik. General ISP and network discussion also permitted. Please ensure if you're asking a question you have checked the Wiki First: https://help.mikrotik.com

# Upgrade to RouterOS 7 - Was a bad idea

Today I tried to update my LHGG to the latest versions (RouterOS and LTE modem)

I was version 6.48.4 and a modem version of 25 so I said it was time to update to the latest versions that are available.

I worked according to Mikrotik's docs, where it says to update the RouterOS to the newest version before updating the modem, so I did it that way.

In the beginning, when I checked an available update it appeared to me that 6.49.6 is the most recent version. I thought maybe it should be on the most recent version of 6 to be able to update to 7. But even after updating the latest update of 6, the update to 7 was not available.

After searching on Google I found this guide:

Then I was able to update to 7 successfully.

After that, I updated the version of the modem to version 34 took a few minutes, and then everything seems to work.

BUT THEN!!!

I had to change the APN to another APN because my main APN is with Passthrough Interface but in 7 no APNs appeared at all.

I tried to restore from a backup with the "force-v6-to-v7-configuration-upgrade=yes". (It was an older restore because I'm stupid and didn't make a backup before starting the process - but that backup is still fine. At least there was something)

did not help.

So I tried to add a new APN

Then this message appeared to me:

"can not add new LTE apn - feature not implemented" - or something like that.

At this point, I was very angry!! How do you let an LTE device be updated and remove something so important as APN?! How come in their Docs, in the update section to version 7 there is no huge message in red that warns those who use LTE to not update!??!

I still didn't give up and tried to add and edit the APN through the CLI, but all the commands changed from 6, I couldn't even get to the info for the LTE interface.

Now I know that the APNs were still in the system because somehow I was still able to see them through the CLI. But through the webfig, it was no longer possible! WTF.

Oh, and I also noticed that all the packages disappeared after the update to 7. I only have RouterOS left after the update, I don't know if it's supposed to be like that or not.

In the end, I went back to the latest version of 6, thank God it's possible.

What's more, throughout the entire process and restarts, I never once worried that the antenna wouldn't go back up and I'd have to go up on the roof to physically connect to it.

Now if I did something wrong and APN does exist in 7 I would love to know what I did wrong and what needs to be done.

A process that should have taken 10-15 minutes ended up taking over an hour.

Ok, good to know. so want to do? Factory restore or net install? And if I do factory restore can I after restore from backup? And how to do net install?

if i were you i'd just netinstall back to v6 and restore the old backup and wait.

if you really want to use v7, first make sure that feature you need is actually present (email support@mikrotik.com asking about it). then netinstall to v7 and try again.

Like I said in the post and the comments I I’m back to 6 and will stay in 6 for the time being

I have wAP with LTE which has been upgraded to v7 few months ago. LTE works, including old APN which has been configured in v6 and i just tried it - i can add new APN and it works too (7.4, did not install 7.4.1 yet).

And no i, am not using winbox (and not going to) either, it works in web ui.

So something must have gone wrong... may be a clean install would help?

Do you remember how you upgrade to v7 from 6? Maybe the guide I follow was missing something. And what packages appears in your web ui in v7? And how do you recommend to do a clean install? From what I understand netinstall isn’t an option for me because I don’t have direct access to the LHGG.

I selected "upgrade" update channel in packages and hit "download and install". It rebooted and everything came right up as if nothing have changed.

In packages i see single "routeros" package, but that's expected - they made v7 a single integrated package with everything apart from very few optional packages none of which is installed by default. Good or bad not sure - but it is what it is.

And yeah, if netinstall is not and option... i'd probably just leave it at v6 for now. It works, it'll still get security updates and there is not much benefit in upgrading to v7 unless you need some specific feature. If anything v6 is probably still more stable...

Ok thanks you! I will leave it on v6 for now. At least I was able upgrade my modem version.😁

Webfig is basically deprecated now.

I know I can edit LTE APNs in Winbox.

Also yes, routeros is now a single package, with just some extra packages.

What?! Why? Login to a WebUI is much easier and flexible then to install software on a PC. If they remove the WebUI in favor of Winbox it will be stupid move!

Speak for yourself, I much prefer well made tools like Winbox rather than clunky webuis.

same here, I hate webfig with passion

They can improve the WebUI instead of wasting time on Winbox. And Winbox (as the name indicates) is a Windows software! I use MacOS and Linux and running a VM to manage my LHGG is stupid!

And just to make sure, you can edit and add APN in Winbox on RouterOS 7?

Winbox runs perfectly fine on Wine, so you don't need a VM.

Did you at least clear your browser cache after upgrading?

I have run Winbox on wine all day for a year. i remember editing my APN on winbox on rosv7

Oh come on. Why so much hate towards webfig? I think it's amazing. Also very consistent with winbox and Shell.

Those downvotes are not justified.

You know nothing Jon snow

Netinstall would require holding the reset button until netinstall shows up. https://help.mikrotik.com/docs/display/ROS/Netinstall

Is there any reason you want to be on v7? Are you looking at replacing the current modem with a new modem?

Oh ok so netinstall isn’t an option for me - the antenna is on the roof. No, just want to be on the most recent version.

I also have SXT LTE on the roof - and upgrading to ROS7 was one of the most uneventful upgrades I've done. It was just like minor update, download & reboot, update firmware, another reboot, done.

I did it with winbox, though.

Stay on v6, unless some of the new LTE changes are something you need.

It's not a bad idea. You just did it wrong.

https://reddit.com/r/mikrotik/comments/wix2ys/_/ijg5jnb/?context=1
