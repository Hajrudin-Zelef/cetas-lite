---
id: etape7-phasei-media/04-copy-chapters-from-source-when-remuxing/overview
title: "Copy chapters from source when remuxing"
domain: copy-chapters-from-source-when-remuxing
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [605, 606]
section: "Copy chapters from source when remuxing"
sha256: fe0014b9adbafc8f2d0f96d103eac01a7e01188e705296166db0b1960df2da54
---

# Copy chapters from source when remuxing
ffmpeg -i in.mkv -map 0 -c copy -map_chapters 0 out.mp4
