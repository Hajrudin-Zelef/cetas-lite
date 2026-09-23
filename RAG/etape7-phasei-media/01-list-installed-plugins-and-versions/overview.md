---
id: etape7-phasei-media/01-list-installed-plugins-and-versions/overview
title: "List installed plugins and versions"
domain: list-installed-plugins-and-versions
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape7_phaseI_media.md
source_anchor: ""
source_lines: [557, 559]
section: "List installed plugins and versions"
sha256: d69bd0eb148779159d0ef289a3775d94b1c26e638258341ebb33b82ac31288d6
---

# List installed plugins and versions
curl -s -H "X-Emby-Token: $K" "$J/Plugins" | \
  python3 -c "import json,sys;[print(p['Name'],p['Version'],p['Status']) for p in json.load(sys.stdin)]"
