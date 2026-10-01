---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-20
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["intel", "valuation"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2820, 2855]
sha256: 6ced2fe37503b88df6414fb7fabea351ef377d9bc5e9cece85ca2ccffaf57cef
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

**R1.** local (100) — le firewall lui-même ; trust (85) — réseau de confiance ; dmz (50) — serveurs exposés ; untrust (5) — Internet/non fiable. (Section 20.)
**R2.** Non : c'est la **première règle qui matche** qui s'applique (évaluation de haut en bas). Le `permit any` placé avant intercepte tout. Il faut ordonner du plus spécifique au plus général. (Section 31.)
**R3.** Parce que la règle NAT easy-ip traduirait aussi les adresses du trafic destiné au tunnel, ce qui casse l'IPSec. L'exemption `no-nat` se place **avant** la règle NAT. (Section 52.)
**R4.** Phase 1 OK, phase 2 KO : ACL (miroir), transform-set ou PFS incohérents entre les deux côtés — ou NAT/politique qui bloque le trafic clair. (Sections 61, 148.)
**R5.** (1) Modèle sous-dimensionné pour l'UTM (débit UTM ≪ débit firewall), (2) profils appliqués sur des flux volumineux non à risque, (3) inspection SSL sur tout le trafic. (Sections 73, 152.)
**R6.** HRP synchronise : sessions, tables NAT, tables de routage, SA VPN. Il NE synchronise PAS : certificats, licences, certains états applicatifs — à refaire sur chaque nœud. (Section 103.)
**R7.** Identifier la signature en cause dans les logs, créer une **exception ciblée** (whitelist/exclusion de signature), documenter, revue trimestrielle — jamais désactiver tout le profil. (Sections 86, 153.)
**R8.** Juridique : information des salariés/représentants du personnel, exclusions (banque, santé). Technique : déployer la CA interne sur les postes, gérer le certificate pinning, coût CPU très élevé. (Section 82.)
**R9.** `display zone` : l'interface est-elle bien **membre d'une zone** ? Une interface sans zone ne fait passer aucun trafic. (Section 24.)
**R10.** (1) Sauvegarde config vérifiée, (2) firmware exact + checksum, (3) console branchée + fenêtre de maintenance, (4) upgrade (passif d'abord en HA), (5) vérification complète + garder l'ancien firmware 1 semaine + plan de rollback écrit. (Sections 117, 123.)

</details>

## 190. Pour aller plus loin : les chantiers d'après

1. **Firewalls virtuels (vsys)** : découper un USG6680 en plusieurs firewalls logiques indépendants (multi-clients, séparation prod/hors-prod) — à étudier quand un seul boîtier doit servir plusieurs entités.
2. **Pilotage centralisé** : NCE / SecoManager pour gérer un parc d'USG (politiques centralisées, logs, conformité) — indispensable au-delà de 5–10 boîtiers.
3. **SD-WAN sécurisé** : les gammes récentes (USG6000F) intègrent le SD-WAN chiffré — piste pour remplacer des MPLS coûteux entre sites.
4. **Automatisation** : l'USG expose une **API REST** (selon version) + NETCONF — scriptage des audits de config, détection de dérive (comparaison périodique avec la config de référence).
5. **Threat intelligence** : abonner l'USG aux flux de threat intel (listes d'IP malveillantes dynamiques) en complément des signatures.
6. **Zero Trust** : coupler l'USG (segmentation, authentification) avec du 802.1X/NAC sur les switchs pour une vraie architecture Zero Trust d'entreprise.
7. **Maquette permanente** : garder un USG de labo (ou USG6000V virtualisé) pour tester **chaque** changement avant la prod — le meilleur investissement sécurité qui soit.

## 191. Sources et références

- Fiche technique « Huawei USG6000 Series NGFW Overview » (débits par modèle : 500 Mbit/s → 40 Gbit/s, IPS+AV jusqu'à ~15 Gbit/s, 6000+ applications, base URL 130+ catégories / 500M+ URL).
- « HUAWEI USG6000, USG6000E, USG9500, and NGFW Module Quick Configuration Guide (With New Web UI) » — basé sur USG6000E V600R007C00 (management GE0/0/0, 192.168.0.1/24 par défaut, premier login web avec création d'administrateur sur V600R007C20+).
- White paper technique « HUAWEI Secospace USG6000 Series » (4 zones par défaut : Trust, Untrust, DMZ, Local ; disques SAS RAID1 hot-swap pour les logs).
- Spécifications matérielles par modèle (ports, alimentations AC/DC redondantes 1+1 sur 3U, console RJ45/Mini-USB).
- Documentation CLI VRP : exemples `security-policy`, `nat-policy` (easy-ip), zones, IKE/IPSec issus des guides de configuration Huawei.

> ⚠️ **Note de l'auteur** : les valeurs constructeur citées datent de septembre 2026. Avant tout achat ou dimensionnement, **vérifie sur la fiche du modèle exact** (les gammes E et F ont leurs propres fiches). Les commandes CLI sont syntaxiquement plausibles pour VRP V500/V600 mais **certaines variantes existent selon la version** — les points marqués « à vérifier selon version » doivent être validés avec `?` en maquette avant la prod. Aucun mot de passe réel ne figure dans ce guide : tous les secrets des exemples sont fictifs.

---

*Fin du guide — Huawei USG6000. Bon courage en baie, et que tes bascules HA soient toujours en heures creuses.* 🔧
