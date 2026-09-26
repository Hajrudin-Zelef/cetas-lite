---
id: collect-260926-mikrotik/mikrotik/questions-70578-cisco-and-mikrotik-as-console-server-3ac3e6eb
title: "questions-70578-cisco-and-mikrotik-as-console-server-3ac3e6eb"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-70578-cisco-and-mikrotik-as-console-server-3ac3e6eb.md
source_anchor: ""
source_lines: [1, 16]
sha256: 5567356d5ee3d7aef7ed24a150dd7dc33f268a0c9ef6c1167a7ba6172e6166c5
---

# questions-70578-cisco-and-mikrotik-as-console-server-3ac3e6eb

we have 2x Mikrotik CCR 1072 and CCR 1036 8G 2S+ and we have cisco nexus switches, so I want to connect the console port from one of the nexuses to the 1072 and connect other nexuses to 1036, so should i use a normal copper cable? or I should use a cross cable ? and just connect and login to the cisco console from Mikrotik? thank you.
1 Answer 1
If you want to connect from Console port (not managment) to normal ethernet port in Mikrotik then it won't work unless Mikrotik have some option to make ethernet port a console one then maybe. If you want to connect Managment port on nexus to mikrotik then you can use normal copper cable and it should work without logging into mikrotik, just straight to Nexus (unless mikrotik have to act as jump server).
- 
        mikrotik has an console portBlackmetal– Blackmetal2020-10-20 19:36:18 +00:00Commented Oct 20, 2020 at 19:36
- 
            
            
- 
        ok but this console port is used to get an access to mikrotik it cannot be used as a some sort of bridge to another console. What Nexus model do you have? Because most of them have management port which is out of band so it's better than logging into nexus via normal port which is in bound port. Ofc. management port have to be configured so in case of configuration loss then you need to connect via console. If you really want control over Console port then maybe consider setting up raspberry pi and connect to it console cables of the devices that you need to manage via console?kubn2– kubn22020-10-20 19:41:47 +00:00Commented Oct 20, 2020 at 19:41
- 
        i want to know if one day my cpu usages will be %100 i can not login to my switch by ssh but am i able to login with console port? i need console port for that day, or when my CPU is %100 there is no difference and i can not login to switch with console or mgmt port ? i have 3064pqBlackmetal– Blackmetal2020-10-20 19:46:29 +00:00Commented Oct 20, 2020 at 19:46
- 
        So when your switch is overloaded the most chance you will have by logging via console then management port which is out of band and then normal port. If you are happy with not the biggest but also not the smallest chance to login which is management port then I would do that. If you really need this extra feeling that you will be able get into switch during hard times and raspberry pi or some other old unused device isn't much hassle then you can use that as some sort of jump server for the console accesskubn2– kubn22020-10-20 19:56:38 +00:00Commented Oct 20, 2020 at 19:56
- 
        hm.. so you suggest me use raspery as console server ?Blackmetal– Blackmetal2020-10-20 20:07:07 +00:00Commented Oct 20, 2020 at 20:07
