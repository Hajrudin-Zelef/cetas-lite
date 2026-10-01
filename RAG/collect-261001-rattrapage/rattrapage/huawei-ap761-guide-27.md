---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-27
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2747, 2777]
sha256: 22bd63396a74b9bc5abab5460a7ae7792314f9a8635382fbffc4f36b7fd9c136
---

# Guide ultra-complet — Huawei eKit AP761

**Documentation officielle (à consulter en premier pour chaque valeur « à vérifier ») :**
- Datasheet **Huawei eKitEngine AP761** (réf. 02355VFB) — la source des valeurs constructeur de ce guide.
- **Country Code & Channel Compliance Table** Huawei — les canaux autorisés par pays (indispensable avant tout déploiement hors France).
- Guide de configuration WLAN de **ta version logicielle exacte** (`display version` sur l'AP) — la syntaxe CLI varie.
- Release notes de chaque firmware avant mise à jour (chap. 74).

**Outils terrain :**
- App **HUAWEI eKit** (onboarding, supervision) + **CloudCampus** (BLE).
- App d'analyse Wi-Fi sur smartphone (scan, RSSI) — pour le survey simplifié (chap. 60).
- **iPerf3** + un serveur sur le LAN — la mesure de débit qui isole le Wi-Fi (chap. 81).
- Testeur de câble qualifiant — le meilleur investissement anti-pannes « mystères » (chap. 24).

**Tes autres guides (compléments directs) :**
- `~/workspace/user/files/onduleurs_ups_guide.md` — dimensionner l'onduleur avec la charge PoE (chap. 25).
- `~/workspace/user/files/zabbix_guide.md` — supervision SNMP des AP (chap. 68).
- `~/workspace/user/files/debian_ubuntu_guide.md` — serveur RADIUS (FreeRADIUS), syslog, NTP (chap. 48, 69, 70).
- `~/workspace/user/files/proxmox_guide.md` — héberger les VM de supervision/RADIUS.

**Formations et communautés :**
- **Huawei eKit** : webinaires et certifications en ligne (portail Huawei) — la gamme évolue vite (le Wi-Fi 7 est arrivé en 2024–2025).
- **CWNA** (Certified Wireless Network Administrator) — la référence indépendante pour la radio Wi-Fi : si tu veux passer du « je dépanne » au « je conçois », c'est LA formation.
- Communautés : forums Huawei Enterprise, groupes d'intégrateurs eKit — pour les retours de terrain sur les firmwares (chap. 74).

**Veille à mettre en place :**
- S'abonner aux **avis de sécurité** Huawei (failles Wi-Fi : KRACK en 2017, FragAttacks en 2021 — il y en aura d'autres).
- Suivre l'**ouverture du 6 GHz** dans l'UE (chap. 19) : le jour où l'outdoor y aura droit, le dimensionnement changera.
- Garder un œil sur les **prix du Wi-Fi 7** : quand l'AP772E baissera, le calcul du renouvellement (chap. 105) changera.

---

*Fin du guide — Huawei eKit AP761. Rédigé le 27/09/2026 à partir des fiches constructeur vérifiées ce jour-là. Les valeurs marquées « à vérifier sur la fiche du modèle exact » doivent être re-validées avant tout appel d'offres ou déploiement. Bon chantier.* 🔧📡
