---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-uyqy9m-scripting-question-scripts-work-running-from-cli-c563c5be
title: "Scripting question: Scripts work running from CLI but not in 'Scripts' - how to troubleshoot - point me in the right direction"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-uyqy9m-scripting-question-scripts-work-running-from-cli-c563c5be.md
source_anchor: ""
source_lines: [1, 73]
sha256: 7e976a9968039c77ed92c8572d254d76bf6240b6050af8aff58f5b6052c29db8
---

# Scripting question: Scripts work running from CLI but not in 'Scripts' - how to troubleshoot - point me in the right direction

*Source : https://www.reddit.com/r/mikrotik/comments/uyqy9m/scripting_question_scripts_work_running_from_cli/*
*Auteur : u/XenSid | Score : 8 | r/mikrotik*

Hello, 

I keep encountering issues with scripting where I can run :if :for :foreach etc code encased in nested {} and be able to run it from the command line but am unable to get the same scripts to run as scripts.

They just don't run or partially run and I have no way to know how or why they aren't succeeding. 

Can anyone point me in the right direction on what is going on here or how I troubleshoot these scripts correctly?

I assume (hope) this is one of those things where there is some feature or best practice or something similar that I don't know about/I am missing and then when I know what it is it will be something I can  
 mostly work out on my own and with Googles help.

Thanks in advance.

---

## Commentaires

**RoutingMonkey** (score 2):

There is mikrotik scripting which uses commands and then there is Mikrotik Scripting which uses their script language

  **XenSid** (score 1):

  Aargh, I'll trawl the wiki, this will be my problem. I didn't realise it was different. I assumed I was making a lot of formatting mistakes so the scripts are failing because of that.

    **RoutingMonkey** (score 1):

    You can run text file command scripts by clicking and dragging them and then typing “import filename.txt” into the terminal, or by copy and pasting the contents directly into the terminal

    **RoutingMonkey** (score 2):

    All these other commenters missed the forest for the trees lol

**ZivH08ioBbXQ2PGI** (score 3):

The permissions on the script have to match the permissions on the scheduled task. 

Not sure that that’s your problem, but don’t forget it!

  **XenSid** (score 1):

  I've actually never changed the permissions on a script... good to know I should be setting them though :D 
  
  😊 thanks

**yabdali** (score 3):

You can go to system logging and enable the topics and choose script.  Alternatively,  your script can output your error catching lines using :log parameter.

  **XenSid** (score 2):

  😮 script logging. Crap. That will do it. Thanks a lot. That's amazing. 
  
  With log on-error can you catch the output to log or is it just a message you write for yourself?

    **yabdali** (score 1):

    You can do both, I presume.  You need to ro check the docs or look at some code examples for script error logging.

    **yabdali** (score 2):

    Check some good discussion about how to do it here 
    
    [https://forum.mikrotik.com/viewtopic.php?t=113197](https://forum.mikrotik.com/viewtopic.php?t=113197)

**jigglym** (score 1):

If script is running in cli but not in scripts then you likely have syntax errors like semicolon missing somewhere or bracket not closed or combination of all. If you’re running parts of your script in cli, try running it all, error will get highlighted in terminal. If full script runs, you can just drag it to your winbox and then run from system>scheduler at intervals by importing it - /import yourscript.rsc or .txt
