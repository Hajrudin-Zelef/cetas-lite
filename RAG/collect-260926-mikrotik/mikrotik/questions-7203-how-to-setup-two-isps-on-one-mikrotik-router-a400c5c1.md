---
id: collect-260926-mikrotik/mikrotik/questions-7203-how-to-setup-two-isps-on-one-mikrotik-router-a400c5c1
title: "questions-7203-how-to-setup-two-isps-on-one-mikrotik-router-a400c5c1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-7203-how-to-setup-two-isps-on-one-mikrotik-router-a400c5c1.md
source_anchor: ""
source_lines: [1, 24]
sha256: 3396493246e016b502a22e6e31c27be55a516ff60d4c2f48c712c290f9ee418a
---

# questions-7203-how-to-setup-two-isps-on-one-mikrotik-router-a400c5c1

I have One Mikrotik router running router OS v6.1. I have One ISP connected to it but i recently got another ISP as backup. I want to connect the second ISP to the router in such a way that if the first ISP fails, the second one will kick in immediately. The link from both ISPs are ethernet links> Does anybody have any idea how to do this ?.
- 
        5Have you sen this mikrotik wiki entry?Mike Pennington– Mike Pennington2014-04-04 17:44:01 +00:00Commented Apr 4, 2014 at 17:44
- 
            
            
- 
        Are you hosting any sites? Do you own any public IP space?Avery Abbott– Avery Abbott2014-04-10 21:01:29 +00:00Commented Apr 10, 2014 at 21:01
- 
        @Avery, No, I just have an application server that branches connect to. Yes i have a /29 class c address from both ISPs. I already solved this problem. It wasnt that difficult to figure out once i went into the Mikrotik router. I just configured 2 static routes with different admin distances. The main ISP with a lower admin distance and a ping check on it so that once it goes down, the second ISP kicks in. Thanks for your concern though.dennix2014– dennix20142014-04-12 08:52:13 +00:00Commented Apr 12, 2014 at 8:52
2 Answers 2
It wasnt that difficult to figure out once i went into the Mikrotik router. I just configured 2 static routes pointing towards the ISPs with different administrative distances. The main ISP with a lower administrative distance and a ping check on it so that once it goes down, the second ISP kicks in. To this, i logged into the Mikrotik router with winbox and went to IP tab. I chose routes from the drop down menu. The rest was self explanatory. After configuring it, i pulled the cable from ISP1 and PCs on the LAN were still able to surf the internet. I ran a tracert 4.2.2.2 which returned the backup ISP's router. I plugged the cable from ISP1 back into the router and immediately ran a tracer 4.2.2.2 which returned ISP1's router. Thanks everybody for your help.
- 
        It's working, but if primary ISP lost internet it can keep link and active gateway, so soute over secondary ISP not activated. For smarter routes switching is better to use script and sheduler or netwatch.mmv-ru– mmv-ru2014-07-15 15:29:22 +00:00Commented Jul 15, 2014 at 15:29
- 
            
            
- 
        @mmv-ru, it happens like u said - at times if the default gateway is still reachable but internet is down at the isp end, router wont switch to the second gateway. I really dont know much about script, scheduler and netwatch.dennix2014– dennix20142014-07-17 08:25:15 +00:00Commented Jul 17, 2014 at 8:25
- 
        its too long to make soficticated reply, because there no one perfect solution. Try start from wiki.mikrotik.com/wiki/Failover_via_Netwatch_III_%28English%29 and wiki.mikrotik.com/wiki/Failover_Scriptingmmv-ru– mmv-ru2014-07-20 16:04:13 +00:00Commented Jul 20, 2014 at 16:04
it was my problem too. you should create tow static route: one of them referred to ISP1 and another to ISP2 , but ISP1's distance is 1 and ISP2's distance is 2. in this situation all client go out through ISP1 and if ISP1 is failed it is transferred to ISP2 . I wish you good luck.
- 
        depends on how it fails. if the link is still up, it hasn't "failed". You'd have to setup some sort of health monitor to detect when the gateway is unreachable.Ricky– Ricky2015-01-28 20:06:52 +00:00Commented Jan 28, 2015 at 20:06
