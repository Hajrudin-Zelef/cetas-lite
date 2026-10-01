---
id: collect-261001-meraki/meraki/r-meraki-comments-51cdjp-easiest-way-to-troubleshoot-what-is-hitting-a-4f6cb5b0
title: "r-meraki-comments-51cdjp-easiest-way-to-troubleshoot-what-is-hitting-a-4f6cb5b0"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-51cdjp-easiest-way-to-troubleshoot-what-is-hitting-a-4f6cb5b0.md
source_anchor: ""
source_lines: [1, 10]
sha256: 9a1dc4d65097ae57b5ea8b5bf91f44cabf019b1af634f9079f6cf15eab0a9da4
---

# r-meraki-comments-51cdjp-easiest-way-to-troubleshoot-what-is-hitting-a-4f6cb5b0

Easiest way to troubleshoot what is hitting a firewall rule
Like any good network person, my firewall rule ends with a deny all at the end. I noticed that something(s) are hitting that rule and im just curious on what is exactly hitting that deny all. Is my only option to see what is hitting that rule through a syslog server?
https://meraki.cisco.com/blog/2012/06/splunk-your-network-data-with-the-meraki-mx/
There seriously has to be an easier way (I know with Cisco, sophos, and Palo Alto I can view denys all from the system itself)
Section des commentaires
Deny all is fairly uncommon in SMB from what I've seen. Why restrict outbound access?
Because its good network security for any size business that cares? Now I do understand some smbs dont have the hardware or a person to do that, but we do.
Also I'll give you a great example why, we have a guest wireless network that I locked down to only allow 80,443,53 with firewall rules. Everything else is dropped because guests on our internet connection dont need to be doing anything else on that network but surf the internet and not access our internal assets.
Everything is flowing correctly traffic wise because I have taken the time to learn what is usual traffic for our network. Again its more of a curiosity thing, we moved from a Sophos device to this because Meraki was giving it out for free. Once in a blue moon I have to troubleshoot some firewall rules and like I said having to dump them to a third party solution is asinine
So add the proper block rules to deny guest to internal and allow everything else. Allowing https means you can tunnel any service out so you've added no security.
