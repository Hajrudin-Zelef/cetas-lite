---
id: collect-260926-mikrotik/mikrotik/questions-692578-how-to-install-mikrotik-routeros-on-my-computer-9933952d
title: "questions-692578-how-to-install-mikrotik-routeros-on-my-computer-9933952d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2015-05-17", "2015-05-18"]
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-692578-how-to-install-mikrotik-routeros-on-my-computer-9933952d.md
source_anchor: ""
source_lines: [1, 23]
sha256: c546be9d40d1fca2039fc65ef04acfce5ad45bd07412d66b78081df4531e1897
---

# questions-692578-how-to-install-mikrotik-routeros-on-my-computer-9933952d

I got a mikrotik routerboard 950G and i downloaded netinstall v6.8 but i cannot install the routerOS into my computer. Can anybody help me out?
1 Answer 1
Routerboards come with MikroTik pre-installed.
Netinstall is a Windows utility that allows you to 'format' a routerboard and do a clean install of Mikrotik RouterOS on it ( Netinstall ).
You would only need to Netinstall Mikrotik onto Routerboard if the installation is broken. Otherwise you could just reset its configuration (Password reset).
If you want to install Mikrotik RouterOS to a PC you don't need the Netinstall utility.
You need the ISO image to burn to a CD, boot from and install (eg: http://download2.mikrotik.com/routeros/6.28/mikrotik-6.28.iso )
Since you mention that you already got a Routerboard I can assume that you don't want to install Mikrotik on your PC, rather connect to Mikrotik from your PC.
To connect to Mikrotik you need Winbox. http://download2.mikrotik.com/routeros/winbox/3.0rc9/winbox.exe
Mikrotik has excellent documentation. I suggest you take a look. You will find almost anything you need main page of wiki
- 
        thans a bunch. I cant express my gratitude in wordsToye_Brainz– Toye_Brainz2015-05-17 12:46:55 +00:00Commented May 17, 2015 at 12:46
- 
            
            
- 
        Just for clarification: You don't need Winbox to connect to the Mikrotik router, you can also use the webinterface.Frederik– Frederik2015-05-17 17:38:50 +00:00Commented May 17, 2015 at 17:38
- 
        Correct. Also there is a pretty nice CLI via telnet or SSH. I just find Winbox much more usable and intuitive :)Cha0s– Cha0s2015-05-17 17:41:19 +00:00Commented May 17, 2015 at 17:41
- 
        I want to connect to webfig n i need my routerboard's ip address but i dont know the ip address of my routerboard ... Is there anyway i can get the information from the routerboardToye_Brainz– Toye_Brainz2015-05-18 06:46:04 +00:00Commented May 18, 2015 at 6:46
- 
        Yes. Using winbox. On the field where you type the address, on the right there is a square little button. You press it and it will search your LAN for any mikrotik device. Once found it will list its MAC address and its IP (if assigned any).Cha0s– Cha0s2015-05-18 08:49:16 +00:00Commented May 18, 2015 at 8:49
