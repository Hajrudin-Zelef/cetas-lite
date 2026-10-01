---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-9
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1281, 1453]
sha256: 2da6996e37ec3a2be171506129a6ccaa6b457812e6acd7a1d2f3d6fd07a036bf
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**Limite honnête** : eSight reste un NMS généraliste ; pour du
troubleshooting radio très fin (spectrum analysis poussé, planification
prédictive), des outils spécialisés complètent. Pour l'exploitation
quotidienne d'un Wi-Fi de campus, le module suffit.

## 64. WLAN et énergie : le clin d'œil au chef de service

Les AP sont souvent alimentés en **PoE** par les switches d'accès :
- Superviser la **consommation PoE** par switch (budget PoE utilisé vs
  disponible) : un switch dont le budget PoE est à 95 % ne peut plus
  accueillir d'AP supplémentaires.
- En cas de coupure secteur sur onduleur, savoir quels switches portent
  le Wi-Fi critique permet de **prioriser** (éteindre le Wi-Fi invité
  pour préserver l'autonomie du Wi-Fi production, par exemple).
- Corréler « AP hors ligne » avec « switch en défaut d'alimentation » :
  la cause n'est pas radio, elle est électrique — et c'est votre domaine.

---

# 11. COMPTES, RÔLES ET AUDIT

## 65. Rôles : le moindre privilège

eSight gère des **rôles** avec des droits différenciés. Découpage type
à mettre en place :

| Rôle | Périmètre | Exemple |
|---|---|---|
| Administrators | Tout, y compris système et comptes | 1-2 personnes max |
| Network operators | Alarmes, topo, perfs, acquittement | Équipe NOC |
| Config operators | + gestion des configurations | Réseau niveau 2/3 |
| Read-only / Guests | Consultation seule | Managers, auditeurs |
| Site-restricted | Périmètre géographique limité | Technicien d'un site |

**Règles :**
- Comptes **nominatifs** : jamais de `reseau1` partagé. En cas d'incident,
  on doit savoir *qui* a fait *quoi*.
- Le compte admin générique ne sert qu'aux opérations d'administration
  système, jamais à l'exploitation courante.
- Revue des comptes **tous les 6 mois** : départs, changements de poste.

## 66. Domaines : cloisonner par organisation

La gestion par **domaines** permet de restreindre la visibilité et les
actions à un périmètre (site, filiale, client) :

- Le technicien du site B ne voit que le site B (moins de bruit, moins
  de risque d'action sur le mauvais équipement).
- En environnement multi-entités, chaque entité a son domaine avec ses
  exploitants, et la DSI centrale garde la vision globale.

À combiner avec les vues de topologie par site (section 34) pour une
expérience cohérente.

## 67. Audit : qui a fait quoi

eSight journalise les **logs** d'administration et d'exploitation :
connexions, acquittements d'alarmes, modifications de configuration,
changements de seuils, imports de licences…

- **Vérifier** que l'audit est actif et que les logs sont conservés
  assez longtemps (ex. : 1 an — contrainte réglementaire possible).
- En cas d'incident, les logs d'audit sont la **première pièce** de
  l'enquête (« qui a modifié la config du firewall vendredi à 18h ? »).
- Protéger les logs : un admin ne doit pas pouvoir effacer ses propres
  traces (droits différenciés, export vers un syslog externe si possible).

## 68. Intégration à l'annuaire d'entreprise

La brochure 23.1 mentionne l'**intégration à un serveur
d'authentification** (authentication server integration).

- Raccorder eSight à l'**AD/LDAP** de l'entreprise : un seul mot de
  passe par exploitant, désactivation automatique au départ (plus de
  compte orphelin).
- Mapper les **groupes AD** vers les rôles eSight (ex. : groupe
  `G-reseau-NOC` → rôle Network operators).
- Garder **un compte admin local de secours** (coffre-fort) pour accéder
  à eSight si l'AD est indisponible — sinon, panne AD = plus de supervision.
- Détails de configuration (LDAPS, attributs, mapping) → **à vérifier
  sur la documentation officielle** de votre version.

---

# 12. SAUVEGARDE D'ESIGHT ET PRA

## 69. Ce qu'il faut sauvegarder (et pourquoi)

eSight concentre des années de paramétrage : profils de découverte,
règles d'alarmes, seuils, rapports, comptes. Perdre eSight sans sauvegarde
= des semaines de reconfiguration. À sauvegarder :

1. **La base de données** (inventaire, alarmes, perfs, configs
   d'équipements sauvegardées, paramétrage) — via le **Database Backup &
   Restore Tool** fourni.
2. **Le système eSight** (binaires, fichiers de configuration applicative).
3. **Les licences** (fichiers/actes — sans eux, pas de réinstallation
   fonctionnelle).
4. **Le dossier d'exploitation** (schémas, mots de passe au coffre,
   procédures) — hors eSight, évidemment.

## 70. Stratégie de sauvegarde : le plan type

| Quoi | Fréquence | Rétention | Destination |
|---|---|---|---|
| Base eSight (complète) | Quotidienne (nuit) | 30 jours glissants | Stockage externe (NAS, pas le serveur eSight lui-même !) |
| Base eSight | Avant chaque changement majeur | Jusqu'à validation du changement + 30 j | Stockage externe |
| Système eSight | Hebdomadaire | 4 semaines | Stockage externe |
| Licences + dossier d'exploitation | À chaque changement | Permanent (versionné) | Coffre + PRA |

**Automatiser** : sauvegarde planifiée (le Backup Tool le permet en
périodique d'après le constructeur), avec **contrôle du succès**
(e-mail de confirmation ou supervision du fichier généré — une sauvegarde
qui échoue en silence ne sert à rien).

## 71. Restauration : la procédure à tester

1. **Environnement de test** : restaurer au moins une fois par an sur
   une VM isolée pour valider que la sauvegarde est exploitable.
2. **Procédure écrite** : pas à pas, avec les commandes exactes, les
   comptes à utiliser, les durées constatées lors du test.
3. **Ordre de restauration** : OS → base → applicatif eSight → licences →
   vérifications (login, découverte d'un équipement test, alarme de test).
4. **Ne restaurez jamais « par-dessus » un eSight en production qui
   fonctionne** sans avoir figé une sauvegarde juste avant.

**Test annuel obligatoire** : une sauvegarde non testée = pas de
sauvegarde. Inscrivez le test au plan de maintenance (section 16) et
faites-le signer.

## 72. PRA : aller au-delà de la sauvegarde

Le PRA (Plan de Reprise d'Activité) d'eSight comprend :

- **RPO/RTO cibles** : ex. RPO 24 h (perte max d'une journée de données),
  RTO 4 h (console à nouveau opérationnelle). À valider avec la direction.
- **Scénarios** : panne serveur simple (restauration VM/snapshot),
  corruption de base (restauration Backup Tool), perte du site
  (réinstallation sur site de secours + restauration).
- **Cluster hot standby** (édition Professional, Linux) : c'est de la
  **haute disponibilité**, pas un PRA — ça ne remplace pas les
  sauvegardes externalisées (section 12 du chapitre architecture).
- **Dépendance circulaire** : eSight supervise le réseau… dont il a
  besoin pour être supervisé et sauvegardé. Prévoyez un chemin
  **hors bande** (console, réseau de management indépendant) pour
  administrer eSight quand le réseau de production est en panne.
- **Annuaire des contacts** : support Huawei/partenaire, contrats,
  numéros — imprimé et dans le PRA, pas seulement dans eSight.

---

# 13. INTÉGRATION NORTHBOUND ET ÉCOSYSTÈME

## 73. Interface SNMP northbound : eSight vu d'en haut

En éditions Standard et Professional, eSight expose une **interface
northbound SNMP** : il peut **transférer ses alarmes** (traps) vers un
NMS de niveau supérieur (manager-of-managers du groupe, hyperviseur
national…).

- Configurer le **filtrage northbound** (section 49) pour ne remonter
  que l'utile.
- Déclarer le NMS supérieur comme destinataire des traps northbound
  d'eSight.
- Tester la chaîne complète : défaut sur équipement → alarme eSight →
  trap northbound → alarme visible dans le NMS supérieur.

## 74. API et développements spécifiques

La brochure 23.1 décrit eSight comme une **plateforme ouverte** pour une
exploitation personnalisée (*« flexible open platform for enterprises to
customize their own network management systems »*).

