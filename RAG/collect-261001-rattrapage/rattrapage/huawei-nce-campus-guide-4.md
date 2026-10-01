---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-4
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [234, 330]
sha256: 1edd25d6992d73afff4d4619eb462ca05a913571c173ee768334c45e67bfb204
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- eSight = **NMS traditionnel** : il supervise (SNMP), configure (CLI/SNMP), génère des rapports. Il ne fait pas d'orchestration SDN, pas de ZTP avancé, pas d'analyse IA poussée.
- NCE-Campus = **manager + contrôleur + analyseur** : il supervise ET pilote (NETCONF/YANG), automatise le déploiement (ZTP, templates), analyse (big data/ML via CampusInsight).
- Huawei vend NCE-Campus comme la plateforme cible pour les nouveaux projets campus, eSight restant la plateforme installée chez des milliers de clients.

Il n'y a pas, dans la documentation publique consultée, d'outil de migration automatique eSight → NCE : la migration est un **re-projet** (ré-onboarding des équipements, re-création des politiques), pas un upgrade. C'est un point structurant de la stratégie (voir Partie 11).

## 6. eSight reste-t-il valable ? — réponse honnête

**Oui.** Sans ambiguïté :

- eSight reste **commercialisé** (brochure eSight 23.1 datée de 2024, références S-Part actives : 88034GED plateforme, 88034GEE licence network/device, 88034GEF licence WLAN/AP, etc., avec souscriptions SnS annuelles).
- eSight couvre des périmètres que NCE-Campus ne couvre pas ou moins bien : supervision **multi-vendeurs** (NCE est quasi exclusivement Huawei), gestion des onduleurs et équipements d'énergie via SNMP générique (point crucial pour Zelef !), PON/OLT/ONT, data center, vidéosurveillance.
- Pour un parc mixte ou une supervision d'infrastructure énergie (UPS, climatisation, groupes électrogènes via SNMP), eSight reste souvent **plus pertinent** que NCE.
- Le coût d'eSight (licence plateforme + licences par device/AP) est sans commune mesure avec un projet NCE complet (serveurs, licences device-day, services d'intégration).

La bonne question n'est donc pas « eSight ou NCE » mais « **pour quel périmètre** » : NCE pour l'automatisation du campus Huawei neuf, eSight pour la supervision transverse/multi-vendeurs/énergie.

## 7. Fin de vie eSight — ce qui est connu et ce qui ne l'est pas (à vérifier)

Honnêteté requise sur ce point sensible :

- **Aucune date officielle de fin de vie (EOM/EOS) d'eSight n'a été trouvée** dans les sources publiques consultées pour ce guide (septembre 2026). Huawei met à disposition un outil « EOM, EOFS & EOS Date Query » sur e.huawei.com pour vérifier par produit — c'est la source faisant foi.
- Les rumeurs de « mort d'eSight » circulent depuis des années dans la communauté ; elles sont **non confirmées** par de la documentation officielle accessible.
- Ce qui est observable : l'effort R&D et marketing de Huawei va très majoritairement vers NCE (et vers eKit pour les PME). eSight évolue lentement (versions 23.x).
- **Recommandation** : avant tout investissement eSight significatif, interroger le partenaire Huawei sur la roadmap officielle et exiger un écrit. Marquer ce point **à vérifier** à chaque comité de pilotage annuel.

En résumé : eSight n'est pas mort, mais ce n'est plus le produit d'avenir de Huawei. On investit dessus pour **maintenir**, on investit sur NCE pour **transformer**.

## 8. iMaster NCE-Campus vs Huawei eKit — la question du petit site

Huawei eKit est la plateforme cloud légère pour PME (petits switches, AP eKit, gestion via app/cloud). La question se pose pour les petits sites de Zelef :

| Critère | eKit | iMaster NCE-Campus |
|---|---|---|
| Cible | PME, petits sites (< ~50 équipements, ordre de grandeur) | Campus moyens/grands, multi-sites |
| Modèle | Cloud eKit, simple, app mobile | Contrôleur on-prem/cloud, complet |
| Coût | Faible / inclus | Licences device-day + infra + intégration |
| Fonctionnalités | Gestion de base, SSID, VLAN simples | ZTP avancé, 802.1X, CampusInsight, API, multi-tenant |
| Compétences requises | Faibles | Équipe formée, projet d'intégration |

**Règle pratique** : un site avec quelques S310 et 2-3 AP361/AP761 en gestion simple → **eKit suffit** et NCE serait du surdimensionnement coûteux. NCE se justifie quand il faut du 802.1X, du multi-site centralisé, de l'automatisation ou de l'analyse avancée. Attention au piège : un AP géré par eKit doit être **désenrôlé** avant de passer sous NCE (voir cas pratique 19).

## 9. Qui a besoin de NCE-Campus — profils types

NCE-Campus est pertinent quand **au moins deux** de ces conditions sont réunies :

1. **Parc Huawei campus significatif** : > 30-50 équipements administrables (switches, AP, AR) — ordre de grandeur indicatif.
2. **Multi-sites** : plusieurs agences/branches à administrer depuis un point central, avec des intervenants locaux peu qualifiés réseau (le ZTP brille ici).
3. **Exigence d'authentification** : 802.1X filaire/Wi-Fi, invités, BYOD, segmentation IoT — la gestion centralisée des politiques devient vite indispensable.
4. **Équipe réduite** : peu d'administrateurs pour un parc qui grandit — l'automatisation compense.
5. **Renouvellement** : projet de refresh du campus = fenêtre idéale pour introduire NCE (greenfield partiel).
6. **Exigence de visibilité** : besoin de SLA, de rapports d'expérience utilisateur, d'analyse de cause racine rapide (contexte exigeant : industrie, santé, éducation).

Pour Zelef : si son entreprise multi-sites déploie du S310/AP361/AP761/AR720 de façon récurrente avec une petite équipe, le profil correspond.

## 10. Qui n'en a PAS besoin — quand eSight ou eKit suffisent

À l'inverse, NCE-Campus est **déconseillé** (ou prématuré) si :

- Le parc est **petit et stable** (< 20 équipements, peu de mouvements) : eKit ou la gestion locale + eSight suffisent.
- Le parc est **multi-vendeurs** : NCE ne gère pratiquement que du Huawei ; eSight ou un NMS tiers (Zabbix, etc.) restent nécessaires en transverse.
- Le besoin principal est la **supervision d'énergie** (onduleurs, etc.) : eSight/SNMP ou une GTE dédiée sont plus adaptés — NCE n'est pas un superviseur d'infrastructure technique.
- **Aucune ressource projet** : NCE exige un vrai projet (dimensionnement, licences, formation, intégration). Sans sponsor ni budget, c'est un échec programmé.
- L'équipe est **débordée par l'exploitation courante** : ajouter un contrôleur SDN sans dégager du temps de montée en compétence = usine à gaz.

Le courage de dire non à NCE sur un petit site est une qualité d'architecte, pas un aveu de faiblesse.

---

# PARTIE 2 — ARCHITECTURE

## 11. Architecture globale — vue à 10 000 mètres

L'architecture logique d'iMaster NCE-Campus s'organise en couches, du terrain vers l'utilisateur :

```
+--------------------------------------------------------------+
|  COUCHE APPLICATION (portails, API NBI, applis tierces)      |
|  Portail admin / Portail tenant / Self-service / VAS store   |
+--------------------------------------------------------------+
|  iMaster NCE-Campus                                          |
|  +------------+  +------------+  +------------------------+  |
|  |  Manager   |  | Controller |  | Analyzer (base)        |  |
|  | (gestion)  |  | (SDN)      |  | + CampusInsight (opt.) |  |
|  +------------+  +------------+  +------------------------+  |
+--------------------------------------------------------------+
|  Southbound : NETCONF/YANG | SNMP | HTTP/2 | HTTPS | TCP ...  |
+--------------------------------------------------------------+
|  COUCHE RÉSEAU : Switches | Routeurs AR | AP/WAC | USG | ...  |
+--------------------------------------------------------------+
```

Deux compléments distribués existent pour les architectures multi-sites :

- **Composant d'authentification** : déployable (jusqu'à 20 instances selon la documentation V300R022C10) sur des branches distantes pour assurer l'authentification locale même si le lien vers le contrôleur central est dégradé.
- **iMaster NCE-CampusInsight** : l'analyseur avancé, composant indépendant connecté à NCE-Campus (redirection one-click via proxy de service depuis NCE).

