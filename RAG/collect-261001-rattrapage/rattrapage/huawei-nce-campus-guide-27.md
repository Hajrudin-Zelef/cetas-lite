---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-27
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["license", "training"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [2483, 2507]
sha256: 6b4a2ecb9d468ab452acc875fa4439d93d98b12553debd821b0760265f09027a
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

**Onboarding d'un site (détaillée sections 56-60)** : site créé, ESN pré-déclarés, DHCP/option prêts, templates validés, firmware cible déposé, fiche technicien prête, tests utilisateurs, conformité 100 %, documentation à jour.

**Migration d'un site (détaillée section 115)** : périmètre défini, compatibilité validée, NCE prêt, templates testés, pilote choisi, fenêtre communiquée, rollback écrit, équipe formée, critères go/no-go écrits, REX prévu.

**Quotidien NOC** : dashboard tour de contrôle, alarmes critiques acquittées ou escaladées, équipements hors ligne > 15 min = ticket, conformité vérifiée, sauvegarde de la nuit OK.

**Mensuel** : rapport direction, revue des licences, revue des certificats (échéances), REX des incidents, mise à jour du dossier d'exploitation.

**Annuel** : test de restauration (chronométré), exercice de bascule HA, revue des rôles, audit de configuration, plan de formation, révision du TCO.

## 185. Annexe E — Bibliographie commentée

- **iMaster NCE-Campus Product Overview (V300R022C10)** — la référence d'architecture : composants (manager/controller/analyzer), authentification distribuée (20 composants), southbound (NETCONF/SNMP/HTTP/2/HTTPS/TCP), restrictions par génération d'équipements. À lire en premier.
- **iMaster NCE-Campus Monitoring and O&M (V300R020C10)** — l'exploitation au quotidien : politiques d'upgrade, signature database, activation des licences devices, SLA/NQA. Le manuel de l'admin.
- **MSP Training Manual — License Usage Guide (V300R022C00 / V300R024C00)** — le modèle device-day, la grâce de 30 jours, le mode MSP et la licence commune de 30 000 device-days. Indispensable avant tout chiffrage.
- **iMaster NCE-CampusInsight Datasheet (V100R025C00)** — les souscriptions d'analyse (package de base obligatoire + options). Pour chiffrer l'« IA ».
- **iMaster NCE-Campus Product Brochure / Datasheet (R24C00)** — vision produit et arguments (OPEX, Tolly +50 % Wi-Fi) — à lire avec recul commercial.
- **Huawei SD-WAN Solution (manuel)** — le rôle WAN de NCE (NCE-Campus comme futur contrôleur SD-WAN mainstream selon Huawei).
- **eSight 23.1 Brochure + S-Part list** — pour comparer les modèles de licence (88034GED/88034GEE/88034GEF...) et garder eSight à sa juste place.
- **Outil EOM/EOFS/EOS Query (e.huawei.com)** — la seule source fiable pour le statut de support d'eSight et des équipements.
- **Guides compagnons de Zelef** : AP361/AP761/S310/AR720/USG6000/eKit (les équipements pilotés), onduleurs/UPS (l'énergie — hors périmètre NCE, à superviser en transverse), eSight (la coexistence).

---

*Fin du guide — v1.0, septembre 2026. Document de travail : recouper chaque valeur dimensionnante avec la documentation officielle de la version NCE cible et le partenaire avant engagement.*
