---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-wr4h30-cant-upgrade-to-routeros-v7-6cdc95ef
title: "Can't upgrade to RouterOS v7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-wr4h30-cant-upgrade-to-routeros-v7-6cdc95ef.md
source_anchor: ""
source_lines: [1, 32]
sha256: 8a36abcf7f17e7d504eb696472d703bb146f1d96934fc126d64003bf2a915c01
---

# Can't upgrade to RouterOS v7

*Source : https://www.reddit.com/r/mikrotik/comments/wr4h30/cant_upgrade_to_routeros_v7/*
*Auteur : u/toot4noot | Score : 10 | r/mikrotik*

I have a **2011UAS-2HnD**, and when trying to upgrade from 6.49.6 to 7.4.1 through the Upgrade channel i get an error `openflow-7.4.1-mipsbe.npk,multicast-7.4.1-mipsbe.npk,lcd-7.4.1-mipsbe.npk missing, use ignore-missing or disable package(s)`. 

Not sure how can i upgrade as i'm more of a newbie. I have the latest RouterBoard firmware for it; 6.49.6.

If i'm not mistaken i should somehow upgrade manually, and **RB2011UiAS-2HnD-IN** is the closest router, so do the files match and can those be used for manual upgrading ?

EDIT: i managed to upgrade it by simply dragging the routeros7 file into the Files menu, then it asked me to reboot. after rebooting i had RouterOS v7 and in the RouterBoard menu it automatically found the v7 upgrade, so i upgraded and rebooted. now my router is on v7.

---

## Commentaires

**yottabit42** (score 4):

These are extra packages that are either not available in 7, or are included now in the base. Disable and remove these packages first, or use the ignore-missing flag from the CLI as instructed.

  **toot4noot** (score 4):

  found rhe solution by manually dragging the file in the webfig, so it automatically installed it

**Milenium_s** (score 2):

\>If i'm not mistaken i should somehow upgrade manually

i have ac3. Should i upgrade  one manually ? or waiting upgrade on the stable channel ?

ver 6.49.6
