---
id: collect-260926-mikrotik/mikrotik/capsman-can-t-get-20mhz-channels-on-2-4ghz-1
title: "2024-12-18 16:44:27 by RouterOS 7.16.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/capsman-can-t-get-20mhz-channels-on-2-4ghz.md
source_anchor: ""
source_lines: [1, 16]
sha256: 9e8568228aa377bfbb258ec4745efb7f28baedc58aca8900790385be1bcb2e4d
---

# 2024-12-18 16:44:27 by RouterOS 7.16.2

Hi,

Am running Capsman V2 on HEX Refresh (E50UG) running 7.16.2 with CAP AC also running 7.16.2 and wireless package removed, qcom-ac package added, placed in CAPS mode.

I have performed all config in Winbox 3.41.

I have a main config that has the Configuration, Security, Steering tabs with settings and the Channel tab empty.

I then have a 2G settings config with the Channel tab setup for the 2GHz N band with 1,6,11 frequencies and 20 Mhz width. This is per Tom’s “Capsman Evolved” Youtube video. I have also tried disabling Secondary Frequency. All other tabs are empty - aside from the Configuration Name.

I also have a 5G settings config which is similar in having 5GHz ac band with ch36/5180MHz - ch52/5260MHz - ch100/5500MHz - ch132/5660MHz specified and allowing 80Mhz width.

My AP setting the 2.4Ghz radio to 40Mhz, but the 5GHz is set to 80Mhz (what is limit ? GUI shows 160).

On my phone I can see 802.11k and 802.11v but no 802.11r, despite having steering configured. Not sure if that is my phone/software.

