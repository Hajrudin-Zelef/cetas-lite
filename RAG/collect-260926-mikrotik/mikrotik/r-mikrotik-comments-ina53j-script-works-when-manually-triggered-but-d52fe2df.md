---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-ina53j-script-works-when-manually-triggered-but-d52fe2df
title: "r-mikrotik-comments-ina53j-script-works-when-manually-triggered-but-d52fe2df"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-ina53j-script-works-when-manually-triggered-but-d52fe2df.md
source_anchor: ""
source_lines: [1, 19]
sha256: fb34ca0749ba3d2b55ace1552ef21adf219cefe4f4f7d1bcd1edfaa09b15c22b
---

# r-mikrotik-comments-ina53j-script-works-when-manually-triggered-but-d52fe2df

Script works when manually triggered, but scheduled run doesn't work
Hi r/mikrotik
I'm planning to automatize HE Tunnel endpoint update.
      Updated this gist for the script https://gist.github.com/horzadome/8e5d99d84525ad8a8ccf
I've made few changes like commented L41,48,56,59 for make it more simple.
    
It does the job when I trigger it with `system script run number=0`.
      I've set up schedule for it, but the script stops at https://gist.github.com/horzadome/8e5d99d84525ad8a8ccf#file-heupdater-txt-L47.
I mean when scheduler triggers the script it prints out this line with he proper values, but nothing else happen.
    
I've tried to read up on policy options for script/schedule, but the mikrotik wiki didn't help much.
What do I wrong?
Section des commentaires
I think I may know what’s going on. Try replacing ‘:local’ with ‘:global’ and see if the script completes. Everything else looks solid.
Hey, thanks for your comment. Tried with global, but didn't help.
Changed back to the :local then tried u/MtHoodlum advice. It's solved the problem.
Commentaire supprimé par un membre de l’équipe de modération
Thank you for your advice.
Recreated the script and schedule then rebooted. It's working now.
