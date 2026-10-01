---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-13l8jyo-firewall-rules-not-working-a1fa2236
title: "Firewall Rules not working"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google", "Oracle"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-13l8jyo-firewall-rules-not-working-a1fa2236.md
source_anchor: ""
source_lines: [1, 61]
sha256: def6743ce55472fdcbc0502e58366e89454a664918352e117f0215d22db713e1
---

# Firewall Rules not working

I am new to OPNsense. Just received Fiber in house and installed on a j6313 i226-v.

Very impressed by Sunny Valley documentation is really good.

I cannot get any firewall rule to work that I try to make. I attached screenshot of a simple icmp rule to allow ping on the wan interface - whats wrong?

Also, I find it strange that Googling my IP shows a different IP address than the one provided in the Dashboard of OPNsense. I dont know if this is related?

If your behind cgnat. Have you turned off block bogus networks on wan?

Are you behind cgnat as well? Sounds like you are

Seems I am behind cgnat because wan int says 100.98.xxx.xx which is a space reserved for that purpose I see.

Would turning off block bogus networks do anything? I dont understand how it would.

So cgnat is actually preventing me reaching my router from wan and hit the rules I set?

I have had bogus rules stop me reaching the internet behind cgnat.

If your behind A cgnat, your not exposed too the internet,

Your technically in your ISP's local network with 30-60 other people sharing one public ipv4. It shouldn't stop ipv6 from working correctly however

Isp confirmed im behind cgnat and thats the reason its not working. Suggestion was to buy static ipv4 or use ipv6.

That would do it every day of the week!

Another option you have is to make a wireguard tunnel to something like Oracle cloud free, and port forward your ports though that to the services inside your network.

Did you get this working? I'm in a similar situation.

I was behind cgnat so that was the reason. I ordered a static IP from my internet provider and the issue was resolved

Thanks for confirming. I'll have to check with my ISP as well.

I would check your NAT, and the rule for the destination object should say a specific destination (or any as its icmp only)

The rule looks fine to allow PING. The different IP addresses definitely may be related. Is the IP address reported for your WAN in OPNsense in the range 100.64.0.0 and 100.127.255.255 ? If so, your ISP may be using CGNAT.

Interface WAN says: 100.98.xxx.xx Gateway WAN says: 100.98.x.x Google says: 87.52.xxx.xx

But most importantly, would CGNAT prevent my rules working? I dont think it should

CGNAT prevents inbound connections from every reaching your WAN. So it's not a question of the rules working or not, they have no inbound packets to evaluate.

What are the first two octets of your WAN IP address? IE, "10.10.x.y". And what does Google show?

Interface WAN says: 100.98.xxx.xx Gateway WAN says: 100.98.x.x Google says: 87.52.xxx.xx

That's CGNAT.

Try changing the Destination from This Firewall to WAN Address.

I tried but didn't work.

I'm not sure why it isn't working for you. I just performed a quick test adding the same rule as yours. I was able to ping my public IP from my PC at my office. Without the rule, the ping requests got blocked. The rule I created worked with the Destination set to This Firewall, WAN Address, as well as Any.

The only other thing that I can think of is to check the Reply-to setting under Advanced Options. You can try setting it to Disable or choose WAN_DHCP.
