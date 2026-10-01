---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-8
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [1057, 1202]
sha256: 63ac278971cab5854d1a7a819a39f60ff3a6b2c7af224d026f4ea60d000e24ff
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

For inbound traffic: this is where things gets confusing to me

I tried this:

Object service DNS-DST

Service udp destination eq 53

Object service Telnet-DST

Service tcp destination eq 23

exit

Nat (outside,inside) source static Global-IP MULTI-SWITCH service Telnet-DST Telnet-DST

Nat (outside,inside) source static Global-IP DNS-server service DNS-DST DNS-DST

And itâs not working and I donât know why ?

When I try the same approach using twice NAT for INBOUND traffic it works perfectly.

So from what I understand for OUTBOUND traffic we must use AUTO NAT (source only) or manual NAT.

For INBOUND TRAFFIC we must use TWICE NAT?

Correct me if am wrong

Very informative .

I should have got this information at the start of my career , would have helped me a lot to crack the interviews.

This is the best and most detailed ASA NAT explanation I have come across. Great job!

You are awesome!! Saved me loads of worry

Thanks heaps for this, very informative. Have you considered putting this tutorial into a .pdf file for printing, I would love to have a hard copy for reference. Cheers, Matt.

Thank you, Ed, for the very clear explanations!

Thank you – best explanation of NAT I have seen – most helpful

This just saved me thank you so much. Keep up the good work!!

This is the best explanation i have ever read for nat on ASA , it’s well organized , supported with examples with all possible show cli outputs commands and each section was fully clarified and explained in simple method , difference between each nat type also was mentioned , this article is too helpful , special thanks for this awsome article.

Great explanation!

Buddy i don’t have word to explain my feeling.This document helping me a lot.I wanna say thank you from the bottom of my heart.

If possible would you please share some more documents on Cisco ASA 8.4 and above version traffic flow & Failover etc.

Best NAT tutorial for ASA, thanks so much!

Wow! Best explanation of ASA NAT I’ve come across. Cisco Press should use this in their books. Thanks for taking the time to create this.

Thank you for the kind words, Sean. =)

I hardly comment online, but I got to say this is one of the best explanation for a technical subject ‘ve seen so far, your illustration was on point. Thanks

Thanks Benny, glad you enjoyed it =)

Best NAT tutorial, i ever seen. Great work!!!!

Best NAT resource on the Web, been to many sites… AWESOME!!

Thanks a bunch!

I have been configuring NAT for many years to date (OFF/ON – even transitioning to 8.3+ NAT) – and gets confused time after time without ASDM GUI interface (even sometimes with that too!). Your tutorial is simply THE BEST I have seen! It clarify everything about NAT I have been struggling with for years, even simply “reading” the statements meaningfully. Your human-readable method and examples given are superb! I can’t thank you enough for your (Practical Networking) many excellent tutorials. May you and your team be abundantly blessed! ð

Hi BigLad =) Thank you for the kind words. Glad you found this content useful =)

I usually do not comment on web pages, but you deserve to be thanked for the way you broke down NAT. There is so much useless information out there, but your page helped me to understand NAT totally. I will continue to use your site and if there is somewhere to donate to you for your hard work and effort please let me know.

Hi Mr. Robot =) I appreciate the kind words. I don’t have a mechanism in place for donations, so the best way to thank me for my efforts is to share the content on this blog across your social networks. Any additional exposure to this content is greatly appreciated! In the future, I’ll have printable PDF versions of a lot of these webpages that people can buy for a few dollars to show their support as Patrons of my content. Sign up for the newsletter to get updates on that.

Hello Ed, really good explanation appreciate your efforts,

I have concern, If i wan to have static nat for inbound connect and dynamic pat for outbound connection; can i achieve this.

If you want one IP address to have a Static NAT, and the rest of the IP addresses in a particular network to have Dynamic PAT, then yes. That is very possible.

If you want the

sameIP address Static NAT’ed inbound, and Dynamically PAT’ed outbound? Then no, that is not possible.
You could configure an additional IP address on the server, and have *that* IP being Static NAT’ed. Then the server’s primary IP could be Dynamic PAT’ed outbound. That would mimic what you are looking for. But you can’t have the same address translated different ways.

I guess, possibly you could if you make each of those methods translate differently based upon the IP address you are speaking with. Then it is simply a Policy NAT.

That’s Great explanation about NAT, I was searching the NAT topics and this is the best one.

Great Explanation so far ever on the internet on ASA NAT.

Thanks alot Ed,

I have got one more concern, can i NAT one public IP address to to different Private ip address for inbound connection with port forwarding.

Example:

10.1.1.2 (80,443) ======== 100.1.1.10

10.1.1.3 (1200,1300) ===== 100.1.1.10

Hi Owais. Yes. As long as only one translation exists for every combination of IP:Port. So, your example would work. =)

In fact, that is the exact example provided in the Static PAT section.

Hi Ed,

Thank you VERY much for this article and indeed all the content you have posted on this site.

As others have said already – it’s the best explanation of ASA OS NAT in a single article that I have yet seen on the Web, and I have been a ‘techie’ for over 25 years!

Simply excellent!

Hi Stewart =) Thank you for the kind words. I am flattered. I’m glad you enjoyed the article.

The best explanation on the internet I have found so far after an intensive research ð

Glad you liked it. Thank you for the kind words =)

Great presentation. Do these NAT commands also carry over to the Router syntax, or is that completely different??

I do agree with Al earlier, that Fortigate and even Palo Alto appear to have a more straight forward syntax when it comes to NAT cli. But I guess like you said its just a matter of getting used to it.

Excellent documentation, thanks for taking the time to do it.

Glad you enjoyed it. Yes, I think the platform you’re exposed to earlier ends up being the one you end up most comfortable with.

The Router syntax is very different. I’ve outlined the Cisco IOS Router commands in this article:

https://www.practicalnetworking.net/stand-alone/cisco-nat-configurations-ios-router/

I wish a Cisco documentation would be presented this way and live would be easier.

Good job! Thanks!

Dzony

I wish Cisco documents looked like this as well! Then I wouldn’t have to go through the trouble of writing these ;). Thanks for the kind words. Glad you enjoyed it.

Brilliant.. Well written and certainly easy to understand

Wowww…. Wonderful explanation in all types of NAT examples… Really useful for everyone…

Glad you enjoyed it, Maheshwaran =)
