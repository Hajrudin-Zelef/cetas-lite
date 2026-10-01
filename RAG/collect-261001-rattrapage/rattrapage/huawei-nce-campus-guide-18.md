---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-18
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1618, 1724]
sha256: 4f9f34ab4376e6abd607b634bff2593c959bda457f27c8bb5e01627aaeab3ff2
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

| Dimension | eSight | iMaster NCE-Campus |
|---|---|---|
| Nature | NMS traditionnel unifié | Manager + contrôleur SDN + analyseur |
| Protocoles | SNMP surtout, CLI | NETCONF/YANG, SNMP, CAPWAP, HTTP/2... |
| Périmètre vendeurs | **Multi-vendeurs** (atout) | Quasi **Huawei uniquement** |
| Périmètre fonctionnel | Supervision, config, rapports | + orchestration, ZTP, politiques, IA |
| Filaire/WLAN/WAN | Supervisés séparément | **Unifiés** et orchestrés |
| 802.1X | Gestion limitée | Politiques centralisées, autorisation dynamique |
| Automatisation | Faible (scripts externes) | **Forte** (ZTP, templates, API) |
| Analyse/IA | Basique | CampusInsight (ML, prédiction) — en option |
| Modèle de licence | Perpétuelle + SnS | **Souscription device-day** récurrente |
| Coût d'entrée | Modéré | **Élevé** (projet) |
| Complexité | Modérée | **Élevée** (projet + formation) |
| Idéal pour | Supervision transverse, multi-vendeurs, énergie | Campus Huawei neuf/renouvelé, multi-sites, automatisation |

## 120. Comparaison détaillée par fonction

| Fonction | eSight | NCE-Campus | Verdict |
|---|---|---|---|
| Inventaire | Oui (découverte SNMP) | Oui (enregistrement SDN) | Équivalent |
| Topologie | Oui | Oui (auto, temps réel) | NCE légèrement devant |
| Alarmes | Oui | Oui + corrélation | NCE devant |
| Configuration | CLI/SNMP, manuelle | Templates, ZTP, NETCONF | **NCE largement devant** |
| Firmware en masse | Limité | Politiques natives | NCE devant |
| 802.1X/NAC | Basique | Centralisé, dynamique | **NCE largement devant** |
| WLAN | Supervision | Orchestration (WAC+AP) | NCE devant |
| SD-WAN branches | Non | Oui (AR, orchestration) | NCE seul |
| Analyse/ML | Non | CampusInsight (option) | NCE seul |
| API | Limitée | RESTful riche (500+) | NCE devant |
| Multi-vendeurs | **Oui** | Non | **eSight seul** |
| Énergie/onduleurs (SNMP) | **Oui** | Non (pas son rôle) | **eSight seul** |
| PON/OLT | Oui | Non | eSight seul |
| Simplicité | Oui | Non (projet) | eSight devant |
| Coût | Modéré | Élevé (récurrent) | eSight devant |

## 121. Coût comparé — TCO sur 5 ans (méthode)

Ne comparez jamais les prix année 1. Méthode TCO 5 ans :

**Postes eSight (5 ans)** : licence plateforme + licences devices/AP + SnS × 5 + serveurs (amortis) + ETP d'exploitation (le poste dominant : config manuelle, déplacements) + formation ponctuelle.

**Postes NCE (5 ans)** : licences device-day × 5 + CampusInsight × 5 (si pris) + serveurs/VM + **intégration** (année 1, souvent = licences année 1) + formation (année 1 + recyclage) + ETP d'exploitation (réduits si l'automatisation tient ses promesses — à modéliser en fourchette).

**Règle de lecture** : NCE gagne le TCO quand (a) le parc est grand/volatile (beaucoup de changements), (b) les sites sont distants (ZTP évite des déplacements), (c) l'équipe est petite (l'automatisation compense). eSight gagne quand le parc est stable et mixte. **Chiffrer les deux scénarios** avec le partenaire avant de décider — c'est un livrable du cadrage, pas une intuition.

## 122. Complexité comparée

- **eSight** : installation simple (serveur Windows/Linux + package), concepts connus (SNMP, alarmes), courbe d'apprentissage de quelques jours pour un admin réseau.
- **NCE-Campus** : projet d'intégration (dimensionnement, HA, organisation, templates, politiques, ZTP, formation), concepts nouveaux (YANG, intent, device-day, RBAC multi-tenant), courbe d'apprentissage de **plusieurs semaines à mois**, besoin d'un **référent NCE** dans l'équipe.

Traduction managériale : eSight = un outil qu'on installe ; NCE = une **compétence** qu'on construit. Prévoir le turn-over (documenter !) — si le référent NCE part sans transmission, le contrôleur devient une boîte noire.

## 123. Cas d'usage : quand choisir quoi — arbre de décision

```
Le parc à gérer est-il majoritairement Huawei campus ?
├─ NON (multi-vendeurs, énergie, PON...) → eSight (ou NMS tiers) reste la base
└─ OUI
   ├─ < ~20 équipements, 1 site, besoins simples → eKit (+ eSight si supervision transverse)
   ├─ 20-50 équipements OU multi-sites OU 802.1X requis
   │   ├─ Budget projet + équipe formable → NCE-Campus (pilote d'abord)
   │   └─ Pas de budget projet → eSight + gestion locale (assumer les limites)
   └─ > 50 équipements / campus multi-sites → NCE-Campus (avec TCO et pilote)
```

Et dans tous les cas : **la supervision énergie/onduleurs reste hors NCE** (eSight/Zabbix/GTE) — ce n'est pas un choix, c'est une architecture.

---

# PARTIE 14 — CAS PRATIQUES (25)

## 124. Cas pratique 1 — L'onboarding ZTP qui échoue (DHCP)

**Symptôme** : les switches neufs restent en « non enregistré », ils ont une IP mais ne contactent jamais NCE.

**Diagnostic** : (1) Vérifier que le DHCP distribue bien l'option contrôleur (`ipconfig`/`dhcp-lease` côté équipement ou capture). (2) Tester la connectivité IP équipement → NCE (ping). (3) Vérifier que l'ESN est pré-déclaré / que la politique d'auto-enregistrement l'autorise. (4) Contrôler les flux pare-feu (le ping passe mais pas le port d'enregistrement ?).

**Causes fréquentes** : option DHCP mal formatée ou sur le mauvais scope, ESN non déclaré, route manquante vers NCE, pare-feu.

**Résolution** : corriger l'option DHCP, déclarer les ESN, ouvrir les flux — puis **réinitialiser** un équipement test et rejouer la séquence complète. **Leçon** : toujours valider le ZTP sur 1 équipement avant d'en brancher 50.

## 125. Cas pratique 2 — Un AP361 qui ne remonte pas dans NCE

**Symptôme** : l'AP est branché, mais n'apparaît ni dans NCE ni sur le WAC.

**Diagnostic** : (1) L'AP est-il **alimenté** ? (LED, test PoE du port du S310 — l'AP361 a besoin de PoE 802.3af ; un port non-PoE = AP mort). (2) A-t-il une IP ? (DHCP du site fonctionnel ?). (3) Voit-il le WAC ? (option DHCP/DNS du WAC, connectivité CAPWAP UDP/5246-5247). (4) Le WAC lui-même est-il géré par NCE ?

**Causes fréquentes** : PoE absent/désactivé, DHCP en panne, WAC non joignable, AP déjà enrôlé ailleurs (eKit ! — voir cas 19).

**Résolution** : traiter dans l'ordre **énergie → IP → WAC → NCE** (toujours du physique vers le logique). **Leçon** : 80 % des « problèmes NCE » sont des problèmes d'alimentation ou de DHCP.

## 126. Cas pratique 3 — Un AP761 extérieur qui reste hors ligne

**Symptôme** : l'AP761 fonctionnait, il est passé « hors ligne » après un orage.

**Diagnostic** : (1) Sécurité d'abord : ne pas intervenir sous orage. (2) Vérifier l'alimentation PoE (injecteur/switch — l'injecteur a-t-il pris la foudre ?). (3) État du câble extérieur et des connecteurs (infiltration ?). (4) Terre du mât et parafoudre (visuel). (5) Si l'AP répond au ping mais pas au contrôleur : vérifier les flux.

**Causes fréquentes** : surtension (injecteur grillé), infiltration d'eau, câble rongé/cisaillé, terre défectueuse.

**Résolution** : remplacer l'injecteur/contrôler le câble, refaire l'étanchéité, vérifier la terre au mégohmmètre. **Leçon** : un AP extérieur = un **actif exposé** — l'inspection annuelle (section 58) n'est pas optionnelle, et le stock de pièces (injecteur, câble) doit exister.

## 127. Cas pratique 4 — Un S310 qui boucle en enregistrement

**Symptôme** : le S310 apparaît « en enregistrement » puis disparaît, en boucle.

**Diagnostic** : (1) Version logicielle du S310 vs version minimale requise par NCE (matrice de compatibilité) — un firmware trop ancien = échec de négociation. (2) Date/heure du switch (NTP) — un certificat rejeté pour cause d'horloge fausse. (3) Doublon d'ESN (équipement déjà déclaré ailleurs). (4) Logs NCE côté onboarding.

