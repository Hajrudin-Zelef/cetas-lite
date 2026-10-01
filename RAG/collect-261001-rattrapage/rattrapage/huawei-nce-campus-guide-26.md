---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-26
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "incident", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [2341, 2482]
sha256: 08ed514c631047c292815b2ffbb5512b1cc061ffee863dc6ffad2592e62075cc
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- **Recette d'installation** : NCE installé, licencié, HA testée (bascule), sauvegarde/restauration testée, documentation livrée.
- **Recette du pilote** : onboarding ZTP fonctionnel, template appliqué et conforme à 100 %, 802.1X validé avec les 3 profils, SSID testés, supervision nominale, REX écrit.
- **Recette par vague** : critères de la section 84 + 72 h d'observation sans incident bloquant.
- **Recette globale** : tous les sites migrés, eSight nettoyé du périmètre, documentation d'exploitation complète, formation réalisée, transfert de compétences signé.

Chaque recette = **procès-verbal signé** avec réserves éventuelles et délais de levée. Ne pas payer le solde d'une prestation sans levée des réserves — c'est le seul levier contractuel qui compte.

## 178. Dossier d'exploitation — sommaire type

Le dossier d'exploitation NCE (le « classeur » de l'équipe) :

1. Architecture (schémas à jour : physique, logique, flux).
2. Plan d'adressage et nommage.
3. Inventaire (avec ESN, versions, dates de garantie).
4. Templates (versions, changelog, variables documentées).
5. Politiques (802.1X, SSID, QoS, invités).
6. Procédures (onboarding, upgrade, rollback, restauration, certificats).
7. RACI et astreinte.
8. Seuils d'alerte et dashboards.
9. Licences (contrats, compteurs, échéances).
10. Sauvegardes (politique, tests de restauration).
11. REX d'incidents.
12. Contacts (partenaire, support Huawei, opérateurs).

Un dossier d'exploitation à jour, c'est la différence entre une équipe et une personne irremplaçable.

## 179. Questions à poser au partenaire — mini-RFP

Avant de signer, exiger des réponses **écrites** :

1. Quelle **version exacte** de NCE-Campus proposez-vous, et quelle est sa date de fin de support ?
2. Fournir la **matrice de compatibilité** version NCE × nos modèles (S310, AP361, AP761, AR720, USG6000) × versions logicielles.
3. **Dimensionnement** : fournir le calcul (outil officiel) pour notre parc à 3 ans, avec hypothèses explicites.
4. **Licences** : détail device-days/an, CampusInsight (oui/non/prix), durée, cotermination, clause de non-renouvellement.
5. **Services** : détail des jours d'intégration (installation, templates, migration, formation) — que livrez-vous exactement ?
6. **Support** : SLA écrits (délais de réponse/résolution par sévérité), périmètre (NCE seul ? équipements ?), langue.
7. **Références** : 2-3 clients comparables (taille, secteur) joignables.
8. **Roadmap eSight** : position officielle sur la coexistence et le support eSight à 3-5 ans.
9. **Sortie** : comment récupère-t-on nos configurations si on arrête ? Quel est le comportement en fin de souscription ?
10. **Formation** : programme, durée, supports, certification éventuelle.

Un partenaire qui refuse l'écrit sur ces points n'est pas un partenaire.

## 180. Feuille de route 12 mois — exemple pour Zelef

Exemple de trajectoire réaliste (à adapter) :

- **Mois 1-2** : cadrage (périmètre, inventaire, TCO, mini-RFP, choix du partenaire), commande.
- **Mois 3** : installation NCE en lab + formation du référent (et suppléant).
- **Mois 4-5** : maquette (S310 + AP361 + AR720) — ZTP, templates, 802.1X en mode monitor, supervision, API vers Zabbix.
- **Mois 6-8** : site pilote (représentatif, non critique) — migration, 72 h d'observation, REX, go/no-go.
- **Mois 9-11** : vagues par site (1-2 sites/mois selon la taille), ajustements, documentation.
- **Mois 12** : clôture — recettes, nettoyage eSight du périmètre migré, bilan TCO réel vs prévisionnel, plan année 2 (CampusInsight ? nouveaux sites ?).

Et en parallèle, toute l'année : eSight/Zabbix continuent de superviser le transverse et l'énergie — la coexistence est assumée, pas subie.

---

# PARTIE 21 — ANNEXES

## 181. Annexe A — Mémo une page : les 20 réflexes NCE

1. NCE = source de vérité — pas de CLI sauvage en production.
2. Tout template est versionné, testé en maquette, relu par un pair.
3. Déploiement par vagues — jamais de big bang.
4. ZTP validé sur 1 équipement avant d'en brancher 50.
5. ESN relevés à la réception, liste blanche en production.
6. 802.1X toujours en mode monitor avant bascule.
7. Certificats suivis à J-90/J-60/J-30 — responsable nommé.
8. Licence suivie trimestriellement — alerte à 80 %.
9. Sauvegarde quotidienne + test de restauration annuel (chronométré).
10. Supervision du contrôleur par Zabbix (qui supervise le superviseur).
11. Compte local de secours documenté (break-glass).
12. Un équipement = une plateforme (désenrôler eKit avant NCE).
13. Fenêtre de maintenance planifiée et communiquée pour tout changement.
14. Rollback prêt avant chaque déploiement — seuil de déclenchement écrit.
15. Après incident : REX sous 72 h, fiche enrichie.
16. Dashboard « tour de contrôle » consulté chaque matin.
17. Revue des rôles et comptes chaque trimestre.
18. Rapport mensuel d'une page pour la direction (avec commentaire humain).
19. Coexistence eSight assumée : chacun son périmètre, matrice RACI.
20. En cas de doute : maquette d'abord, production ensuite.

## 182. Annexe B — Correspondance des écrans (anglais → français usuel)

Le portail NCE étant en anglais, ce lexique aide l'équipe :

| Écran / terme NCE | Traduction usuelle |
|---|---|
| Dashboard | Tableau de bord |
| Topology / Digital map | Topologie / carte réseau |
| Device Management | Gestion des équipements |
| Onboarding / Registration | Enrôlement / enregistrement |
| Template / Profile | Modèle / profil de configuration |
| Site / Region / Tenant | Site / région / tenant (organisation) |
| Policy (access, QoS) | Politique (d'accès, de QoS) |
| WLAN / SSID / Radio profile | Wi-Fi / nom du réseau / profil radio |
| Authentication component | Composant d'authentification (branches) |
| Device Maintenance | Maintenance des équipements |
| File Management | Gestion des fichiers (firmwares, licences) |
| Upgrade / Downgrade | Mise à jour / retour de version |
| Signature Database Upgrade | Mise à jour des bases de signatures |
| Activate Device License | Activation des licences équipements |
| Alarm / Event | Alarme / événement |
| SLA Management | Gestion des SLA (sondes NQA) |
| Report | Rapport |
| System Management | Administration système |
| Third-Party Service | Services tiers (SMTP, update center...) |
| Audit log | Journal d'audit |
| Grace period | Période de grâce (licences) |

## 183. Annexe C — Modèle de procès-verbal de recette

```
PROCÈS-VERBAL DE RECETTE — iMaster NCE-Campus
Projet : ...................... Site/Jalon : ......................
Date : ...................... Participants : ......................

1. PÉRIMÈTRE RECETTÉ
   [ ] Installation / [ ] Pilote / [ ] Vague n°__ / [ ] Globale

2. CRITÈRES (cocher + preuve)
   [ ] Onboarding ZTP fonctionnel (pr. : ........................)
   [ ] Conformité templates 100 % (pr. : ........................)
   [ ] 802.1X validé (profils : .................................)
   [ ] SSID testés (liste : .....................................)
   [ ] Supervision nominale 72 h (pr. : .........................)
   [ ] Sauvegarde/restauration testée (RTO mesuré : .............)
   [ ] Documentation livrée (pr. : ..............................)

3. RÉSERVES
   N° | Description | Responsable | Délai de levée
   ...| ........... | ........... | ..............

4. DÉCISION : [ ] Recette prononcée / [ ] Prononcée avec réserves / [ ] Ajournée
Signatures : ...................... ......................
```

## 184. Annexe D — Checklists imprimables (rappel)

**Installation (détaillée section 38)** : serveurs OK, image vérifiée (checksum), IP/DNS/NTP, flux pare-feu testés, certificats prêts, fenêtre planifiée, rollback écrit, équipe formée.

