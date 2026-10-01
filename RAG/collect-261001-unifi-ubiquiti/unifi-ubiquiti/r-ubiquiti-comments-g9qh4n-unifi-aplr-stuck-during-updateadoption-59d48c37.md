---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-g9qh4n-unifi-aplr-stuck-during-updateadoption-59d48c37
title: "r-ubiquiti-comments-g9qh4n-unifi-aplr-stuck-during-updateadoption-59d48c37"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-g9qh4n-unifi-aplr-stuck-during-updateadoption-59d48c37.md
source_anchor: ""
source_lines: [1, 11]
sha256: 36523b9e2f1572093aca28f63c8a4ecd28ebdfb75f8ec6d504da81992fd72ae5
---

# r-ubiquiti-comments-g9qh4n-unifi-aplr-stuck-during-updateadoption-59d48c37

UniFi AP-LR Stuck during update/adoption 
        
    Hi Everyone,
      I have an UniFi AP-LR that has been in storage for a few months during a move to a new office building. I brought it out today, connected it onto my network and was going to adopt it to add to the AC Lite I purchased for my new office. First of all, it was managed under an old controller at my old office, so I had to hold the reset button down to ready it for adoption. Afterwards, it was identified by the new controller. I clicked to adopt/update firmware and it has been sitting close to an hour with the status of: "UPDATING: (DOWNLOADING)"
Any suggestions on how to proceed? I don't want to brick it by disconnecting it in the middle of a firmware upgrade, but it doesn't seem to be doing anything. There is a solid green ring on the AP.
    
Section des commentaires
SSH into it. set-default would probably do a decent job.
Thanks! I hadn't thought of that. Will give it a try!
It turns out after a bit more searching, this is a fairly widely experienced issue - seemingly due to the controller versions being bundled with specific firmware. So, I think what the issue was is that the AP had a firmware version (3.7.5.4969) that was so old it could not be directly upgraded to whichever version the controller was trying to upgrade it to. So, it kept failing, and then trying the upgrade again. But, I never got any error. It would just go offline for a moment, and then begin the loop again of upgrading.
I finally ended up manually upgrading the firmware via SSH to 3 versions newer ( 3.7.5 > 3.7.58 > 3.9.54). I then rebooted the controller and it auto-detected the AP and successfully upgrade it to 4.0.10.9653. Thanks again for the SSH suggestion!
