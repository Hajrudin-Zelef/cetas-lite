---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-1
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [1, 11]
sha256: 8e3964d8c5a763ab1b64701428fe96e4e2f048ee309a7b4c70574019da680c1e
---

# Guide Wazuh — SIEM & XDR Open Source en production

> **Public :** Zelef, chef de service systèmes & énergies, sysadmin avec des responsabilités d'équipe.
> **Angle :** pratique, production, petites équipes SOC. Chaque section contient des commandes et des fichiers prêts à copier.
> **Version :** Wazuh **4.x** (4.7 → 4.12). Les chemins, ports et paramètres ci-dessous sont ceux de la 4.x. En cas de doute, vérifiez le changelog de votre version mineure : certaines options changent entre versions (ex. `/etc/filebeat`, Indexer basé sur OpenSearch).
> **Avertissement :** ce guide contient des exemples avec des valeurs factices (`motdepasse`, IP `203.0.113.x`). Ne copiez jamais ces valeurs en production ; générez vos propres secrets et adaptez les réseaux.

---

## Table des matières

