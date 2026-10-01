---
id: collect-261001-cisco/cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026-6
title: "Génère un secret de session Graylog (obligatoire, 16+ caractères)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026.md
source_anchor: ""
source_lines: [389, 389]
sha256: 41c9f45a10c115ba5be27d221d065930931f4294a641fef36d018ce8997881ce
---

# Génère un secret de session Graylog (obligatoire, 16+ caractères)

Oui, via l’agent NXLog ou Winlogbeat déployé sur les serveurs Windows, configuré pour transmettre les journaux d’événements (sécurité, système) vers l’input GELF ou Syslog de Graylog, exactement comme pour une source Linux.
