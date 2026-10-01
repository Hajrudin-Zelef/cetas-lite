---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-17
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1501, 1617]
sha256: bed3632ae8afc991dc3525b5e94c4e520ee45833411c44ccd327b00f24c318c2
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

1. **Ne pas tout migrer** : NCE ne remplace eSight que sur le **campus Huawei** (LAN/WLAN/branches). Le reste (multi-vendeurs, énergie, PON, DC) **reste** sur eSight.
2. **Coexistence** d'abord (section 111) : les deux plateformes vivent côte à côte, avec des périmètres clairs.
3. **Migration par site** (section 112) : site pilote, puis vagues.
4. **Capitaliser** : templates, politiques et procédures construits pendant la migration = le patrimoine d'exploitation NCE.

Erreur à éviter : « on migre tout d'un coup pour en finir » — c'est le meilleur moyen de tout casser en même temps.

## 111. Coexistence eSight + NCE — mode recommandé

Pendant (et après) la migration, les deux cohabitent :

| Périmètre | eSight | NCE-Campus |
|---|---|---|
| Campus Huawei (S310, AP, AR, USG) | Supervision transitoire (double alarme à gérer) | Gestion cible (pilotage) |
| Équipements non-Huawei | Supervision (durable) | — |
| Onduleurs / énergie | Supervision (durable) | — |
| PON / OLT | Supervision (durable) | — |
| Politiques 802.1X / templates | — (limité) | Cible |

Règles de coexistence : **définir qui fait quoi** (matrice RACI), éviter la double configuration (un équipement = un pilote), gérer la **double alarme** transitoire (filtrer ou accepter le bruit pendant la migration), et fixer une **date de fin** de coexistence par site (sinon elle devient permanente).

## 112. Migration par site — méthode pas à pas

Pour chaque site :

1. **Préparation** : inventaire du site (équipements, VLAN, SSID, politiques eSight), site créé dans NCE, templates du site prêts et testés en maquette.
2. **Fenêtre** : planifiée, communiquée, avec rollback (section 118).
3. **Désenrôlement eSight** : retirer le site d'eSight (ou le passer en lecture seule transitoire) — **un seul pilote à la fois**.
4. **Onboarding NCE** : enregistrer les équipements (ZTP si réinitialisation possible, sinon onboarding manuel avec reprise de la configuration via template).
5. **Reprise des services** : VLAN, SSID, 802.1X, QoS — vérifiés un par un avec des tests utilisateurs réels.
6. **Vérification** : conformité 100 %, supervision nominale 72 h, aucune plainte bloquante.
7. **Clôture** : documentation à jour, eSight nettoyé du site, retour d'expérience archivé.

Ne passer au site suivant qu'après **clôture complète** du précédent.

## 113. Ce qui est repris de eSight — inventaire

Ce qui se **transfère** (manuellement, pas automatiquement) :

- **L'inventaire** : la liste des équipements (export eSight → import/déclaration NCE) — avec les ESN à relever physiquement si manquants.
- **Les plans d'adressage** : VLAN, sous-réseaux, DHCP — la conception réseau ne change pas, seul l'outil de déploiement change.
- **Les politiques** : règles 802.1X, SSID, QoS — à **retranscrire** dans les concepts NCE (politiques, profils, templates), pas à copier-coller.
- **Les seuils d'alerte** : à réexprimer dans NCE (les métriques ne sont pas identiques).
- **Les rapports** : les besoins (pas les formats) — reconstruire les rapports équivalents.
- **L'historique** : archiver les rapports/alarmes eSight (traçabilité), mais ne pas chercher à l'importer dans NCE.

## 114. Ce qui N'EST PAS repris — limites honnêtes

- **L'historique fin** (courbes, alarmes anciennes) : pas d'import — archiver eSight en lecture seule le temps nécessaire.
- **Les scripts/customisations eSight** : à réécrire (API différentes).
- **Les équipements non-Huawei** : ils restent sur eSight (ou un autre NMS) — NCE ne les prendra pas.
- **Les habitudes** : l'équipe doit **réapprendre** (templates, politiques, ZTP) — budgéter la formation et accepter la phase « on est moins efficaces qu'avant » (2-3 mois typiques).
- **Les licences eSight** : non transférables vers NCE (modèles différents) — à intégrer au TCO.

## 115. Checklist de migration complète

- [ ] Périmètre de migration défini (quels sites, quels équipements — et ce qui reste sur eSight)
- [ ] Matrice de compatibilité NCE × modèles × versions validée
- [ ] NCE installé, licencié, sauvegardé (Partie 5)
- [ ] Templates du site standard conçus, testés en maquette, versionnés
- [ ] Politiques 802.1X/SSID/QoS retranscrites et validées
- [ ] Site pilote choisi (représentatif mais non critique)
- [ ] Fenêtre de migration planifiée et communiquée
- [ ] Procédure de rollback écrite et testée
- [ ] Équipe formée (au minimum : onboarding, templates, supervision, rollback)
- [ ] Astreinte renforcée pendant la migration
- [ ] Plan de communication utilisateurs (ce qui change pour eux : rien, idéalement)
- [ ] Critères de succès/échec écrits (quand on rollback ?)
- [ ] Documentation cible (dossier d'exploitation NCE) initiée
- [ ] Retour d'expérience prévu (post-mortem même en cas de succès)

## 116. Risques de la migration et mitigations

| Risque | Impact | Mitigation |
|---|---|---|
| Template erroné poussé en production | Panne de site | Maquette + vagues + rollback (cas pratique 6) |
| Équipement incompatible (version) | Onboarding partiel | Matrice de compatibilité + audit préalable |
| Perte de supervision transitoire | Angle mort | Coexistence eSight en lecture seule |
| 802.1X qui bloque les utilisateurs | Paralysie du site | Mode monitor/open d'abord, bascule progressive (cas pratique 10) |
| Équipe insuffisamment formée | Erreurs, lenteurs | Formation + pilote + astreinte experte |
| Dépassement budgétaire (licences) | Arrêt du projet | TCO préalable, marge, jalons go/no-go |
| Résistance au changement | Contournement (CLI sauvage) | Communication, formation, détection des écarts |

## 117. Planning type de migration (exemple 3 sites)

Exemple indicatif pour un parc modeste (à adapter) :

| Phase | Durée indicative | Contenu |
|---|---|---|
| Cadrage | 2-3 semaines | Périmètre, TCO, compatibilité, jalons go/no-go |
| Installation NCE | 2-4 semaines | Serveurs, licences, organisation, sauvegarde |
| Maquette + templates | 3-6 semaines | Templates, politiques, tests, formation |
| Site pilote | 2-3 semaines | Migration, vérification 72 h, REX |
| Vague 2 (1 site) | 2 semaines | Généralisation, ajustements |
| Vague 3 (1 site) | 2 semaines | Généralisation |
| Clôture | 1-2 semaines | Documentation, nettoyage eSight, bilan |

Total indicatif : **3 à 5 mois** pour 3 sites avec une petite équipe — sans compter les impondérables. Ne pas promettre moins sans maquette validée.

## 118. Retour en arrière — plan B

Chaque migration de site a un **plan B écrit** :

- **Seuil de déclenchement** : ex. « si le site n'est pas nominal à H+4, on rollback ».
- **Procédure** : ré-enrôler les équipements dans eSight (configurations sauvegardées avant migration !), rouvrir les accès, vérifier les services.
- **Prérequis** : **sauvegarde des configurations eSight/équipements avant** chaque migration (non négociable).
- **Communication** : message prêt (« retour à la situation antérieure, investigations en cours »).
- **Post-mortem** : analyser avant de retenter (jamais de « on recommence demain pareil »).

Un plan B non écrit n'existe pas.

---

# PARTIE 13 — COMPARAISON ESIGHT vs NCE-CAMPUS

## 119. Comparaison eSight vs NCE-Campus — tableau général

