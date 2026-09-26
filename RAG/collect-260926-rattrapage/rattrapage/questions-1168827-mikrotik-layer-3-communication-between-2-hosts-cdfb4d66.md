---
id: collect-260926-rattrapage/rattrapage/questions-1168827-mikrotik-layer-3-communication-between-2-hosts-cdfb4d66
title: "questions-1168827-mikrotik-layer-3-communication-between-2-hosts-cdfb4d66"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-rattrapage/servers-reviews/questions-1168827-mikrotik-layer-3-communication-between-2-hosts-cdfb4d66.md
source_anchor: ""
source_lines: [1, 30]
sha256: d41b6c096ae683a5e52af67586039e6443017ce1681004acf8903116579be84f
---

# questions-1168827-mikrotik-layer-3-communication-between-2-hosts-cdfb4d66

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Test question - mikrotik
In MikroTik RouterOS, Layer-3 communication between 2 hosts can be achieved by using an address subnet. Which one of those "/29" "/30" "/31" "/32"?
Why?
I understand subnetting such as "x.x.x.x/n"(CIDR not. ..) but i don't understand what is it asking me for.. i mean you can connect 2 hosts even with /24, but on the internet you have to change it because of the bigger network, so, bigger networkID and that kind of stuff (at least this is what i've understood)... still can't get the difference between using /29,30,32 ... when should i use those, and why. I'd really appreciate an "in-dept" explanation.
/31 is a special case and only should be used for Point to Point and has its own RFC 3021. It is a little unclear but it might because the test did not specify that this is Point to Point but rather just a network with only two hosts.
A /30 will give you two usable hosts with one broadcast and one network ID. I am unaware of when you would not need a broadcast and network ID. however you usually also have a gateway which takes away from one of the usable hosts but it is still considered two usable hosts
You have to understand CIDR.
The simple cases are the ones where the /x is multiple of 8.
i.e 192.168.100.0/24 can be broken in
192.168.100 network address (255.255.255.0 subnet mask)
1-254 host address
So lets take an example where /x is not multiple of 24:
192.168.100.0/30.
As before, 30 corresponds the amount of bits dedicated to the network address(or subnet mask).
we know that the first three octets are subnet address, those are 24 bits.
30-24 is 6, so the last octet can be divided in 6 bits for network address and 2 bits for host address.
64 is the max decimal number that can be represented with 6 bits and 4 is the max decimal number represented with 2 bits.
So this would mean the last octet has 64 subnet addresses and each subnet would have 4 hosts
So 192.168.100.0/30 represents 4 host addresses, from 192.168.100.0 to 192.168.100.3
192.168.100.4/30 hosts from .4 to .7
192.168.100.8/30 hosts from .8 to .11
And so on.
Note that this implies that cidr of 192.168.100.6/30 would yield the same hosts as 192.168.100.4/30 as .6 is a host of that subnetwork.
So the answer would be x.x.x.x/31 as this means there are 7 bits of the last octet for network, meaning there are 128 network addresses (on last octet) and 2 host addresses on each.
