---
id: collect-260926-mikrotik/mikrotik/questions-670042-mikrotik-server-networks-and-cain-abel-e6ed9366
title: "questions-670042-mikrotik-server-networks-and-cain-abel-e6ed9366"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-670042-mikrotik-server-networks-and-cain-abel-e6ed9366.md
source_anchor: ""
source_lines: [1, 11]
sha256: e0a23a7ca6a0770bda474c412f92e6e6d7cdc23ba962346702f4069c5b8e991b
---

# questions-670042-mikrotik-server-networks-and-cain-abel-e6ed9366

Cain and Abel is essentially a auditing and attack suite, that when installed within your network, is capable of gathering password and other security specific information. 
It is a utility, so you (or a malicious user who has already broken into your network) have to install it and ask it to audit whatever it is you want. its more about attacking your internal network. see a list of capabilities here:
http://en.wikipedia.org/wiki/Cain_and_Abel_%28software%29
there is no real defense per se, other than preventing people from getting the kind of access necessary to install it on your network in the first place. 
it does have some wifi penetration utilities so someone can use it to break into your wifi. your best bet there is to use WPA2 in AES only mode (no TKIP) with a long strong password, and perhaps even a mac filter (weak protection by itself).
you would only know that its being used if you saw the process running or intercepted suspicious traffic, perhaps with an IDPS. 
no tor would have no affect on C&A because its not an internet-bourne thing. be more afraid of your neighbor than a hacker in another country. 
no it would probably not affect SSL connections. 
yes, skype is probably 'safe' to use, but your conversations may be recorded. 
the utility does have a network sniffer built in, so  it can capture any kind of traffic that it wittnesses, but that in itself may not be a huge risk for you.
just keep it off your local network, and your wifi, and you have nothing to worry about, unless someone is already inside your net.
