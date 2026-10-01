---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-16
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["exploit"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [2701, 2730]
sha256: 2f9c566b8320156789de6a7e37531cba46b25454a90481800ff52651831c4013
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

- Datasheet « Huawei NetEngine AR700 Series Enterprise Router Data Sheet » (source des
  valeurs hardware de ce guide).
- Guide d'installation hardware AR720 (câblage, rack, mise à la terre).
- « VRP Configuration Guide » de la version exacte (`display version`) : la référence
  pour chaque fonction (IPSec, OSPF, QoS…).
- Notes de version (release notes) avant chaque upgrade : comportements modifiés,
  failles corrigées.

## 133. Sujets à approfondir après ce guide

1. IPv6 sur l'AR720 (adressage, DHCPv6, OSPFv3) — le dual-stack arrive vite.
2. BGP pour les sites multi-opérateurs ou les interconnexions partenaires.
3. MPLS L3VPN si le cœur du réseau opérateur l'exige.
4. PKI complète et renouvellement automatisé des certificats VPN.
5. Scripts d'exploitation : backup automatisé via expect/Ansible (module `community.network`
   ou scripts SSH), parsing de `display` en Python.
6. Intégration au superviseur (Zabbix/Prometheus SNMP) avec templates Huawei.

## 134. Dernier mot : la discipline qui fait la différence

Le routeur n'est jamais le problème le plus fréquent : c'est l'exploitation qui fait la
différence. `save` systématique, sauvegardes externes datées, dossier de site à jour,
checklist mensuelle faite pour de vrai, firmware suivi, mots de passe uniques. Un AR720
bien exploité tourne des années sans un ticket. Bon courage sur le terrain.

---

*Fin du guide — Huawei NetEngine AR720. Rédigé le 2026-09-27. Valeurs hardware issues de
la datasheet AR700 Series ; vérifier la version VRP exacte (`display version`) et la fiche
du modèle livré pour les points marqués « à vérifier ».*
