---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-6
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [767, 934]
sha256: c33592ecc46f32facb313cdd1dfc29132f814bc929566e5a3a40edea8ae54d98
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

Vues disponibles (selon édition et version) :
- **Topologie physique** : équipements et liens réels (découverts via
  LLDP/CDP, tables ARP/forwarding).
- **Topologie IP** : vision par sous-réseaux/adresses.
- **Vues personnalisées** : par site, par service, par client — construites
  à la main.

**Réflexe terrain** : la topologie ne vaut que si elle est **à jour**.
Une carte avec des équipements décommissionnés depuis 6 mois = une carte
qu'on ne regarde plus. Nettoyage trimestriel obligatoire.

## 34. Personnaliser les vues

- Créer une vue **par site** (bâtiment A, bâtiment B, datacenter) et une
  vue **« backbone »** (cœur de réseau uniquement).
- Positionner les équipements manuellement pour refléter la géographie
  réelle (pas en vrac automatique).
- Colorer / étiqueter par criticité : le lien du datacenter n'a pas la
  même importance visuelle que celui de l'imprimante du 3ᵉ étage.
- Définir la **vue par défaut** affichée à la connexion selon le profil
  (le NOC voit la vue globale, le technicien site voit son site).

## 35. Fonds de plan : rendre la carte parlante

eSight permet d'utiliser des **fonds de plan** (images) derrière la
topologie : plan de bâtiment, carte de la ville/région pour du multi-sites.

- Utilisez des plans **simplifiés et à jour** (pas le plan DWG de
  l'architecte de 2012).
- Placez les équipements à leur position approximative réelle :
  en panne, on localise en 10 secondes (« le switch du local technique
  R+2, bâtiment B »).
- Pour le multi-sites : une vue « France » avec un point par site,
  chaque point ouvrant la vue détaillée du site (navigation en 2 clics).

## 36. Les liens : ce qu'ils montrent (et ce qu'ils cachent)

- Les liens affichent l'état (up/down/dégradé) et peuvent afficher
  l'utilisation (épaisseur ou couleur selon la charge — **selon version**).
- Un lien « down » dans eSight doit correspondre à une alarme :
  si ce n'est pas le cas, la découverte de topologie ou les traps
  sont mal configurés (voir cas n°5).
- **Limite connue** : les liens agrégés (LACP), les tunnels et les
  chemins MPLS ne se représentent pas toujours fidèlement — complétez
  par des vues logiques si vous en avez beaucoup (édition Standard
  requise pour MPLS VPN/tunnels).

## 37. Bonnes pratiques topologie (le minimum vital)

1. Une vue « NOC » globale + une vue par site : pas 47 vues.
2. Nettoyage trimestriel : équipements sortis, liens morts.
3. Fonds de plan à jour après chaque déménagement/travaux.
4. Former les nouveaux exploitants à **lire** la topologie (codes couleur,
   navigation) dès leur arrivée.
5. Captures d'écran de la topologie de référence archivées : en cas
   d'incident majeur, comparer « avant/après » fait gagner un temps fou.

---

# 7. GESTION DES ALARMES

## 38. Le cycle de vie d'une alarme

Comprendre le cycle de vie, c'est comprendre tout le fault management :

```
Événement équipement (panne, seuil dépassé)
        │ trap SNMP / syslog / polling
        ▼
┌──────────────────┐
│ Alarme COURANTE  │ ← visible, à traiter
│ (active)         │
└────────┬─────────┘
         │ acquittement (prise en charge)
         ▼
┌──────────────────┐
│ Alarme ACQUITTÉE │ ← quelqu'un s'en occupe
└────────┬─────────┘
         │ fin du défaut (clear) ou suppression manuelle
         ▼
┌──────────────────┐
│ Alarme HISTORIQUE│ ← archivée, consultable
└──────────────────┘
```

- **Courante** : le défaut est présent (ou pas encore soldé).
- **Acquittée** : un exploitant a pris en charge (ça n'éteint PAS le
  défaut, ça dit « je m'en occupe »).
- **Historique** : soldée, conservée pour analyse et reporting.

**Règle d'or du NOC** : aucune alarme critique ne reste « courante non
acquittée » plus de X minutes (X défini dans vos SLA internes, ex. 15 min).
C'est LE KPI du fault management (section 15).

## 39. Les sévérités : parler le même langage

Échelle usuelle (libellés exacts **à vérifier sur la documentation
officielle** de votre version) :

| Sévérité | Signification | Exemple |
|---|---|---|
| Critical | Service interrompu ou risque imminent | Lien backbone down, équipement injoignable |
| Major | Dégradation forte | Perte d'un membre d'un stack, alim redondante HS |
| Minor | Dégradation faible / anomalie | Seuil CPU dépassé brièvement, ventilateur en défaut (redondé) |
| Warning | Signal préventif | Certificat expirant, espace disque > 80 % |
| Info / Cleared | Informatif / fin d'alarme | Fin de défaut, acquittement |

**Travail de fond indispensable** : la sévérité « sortie d'usine » des
traps n'est pas toujours pertinente pour VOTRE contexte. Une alarme
« ventilateur » sur un switch d'accès redondé n'est pas critique ;
la même sur le seul switch d'un site isolé, si. D'où la **redéfinition
des sévérités** (section 44) : prévoyez un atelier de 2 heures avec
l'équipe pour calibrer les 30 alarmes les plus fréquentes.

## 40. Acquitter : la discipline de base

- **Acquitter = « je prends en charge »**. Toute alarme critique/majeure
  acquittée doit avoir un **commentaire** (ticket, action en cours).
- Interdire l'acquittement « pour faire du propre » sans action :
  c'est la porte ouverte aux pannes ignorées.
- L'**acquittement automatique** (règles) existe : eSight peut
  automatiquement acquitter et basculer en historique les alarmes
  traitées qui correspondent à des règles (d'après la documentation :
  *« after automatic acknowledgement rules are set, eSight automatically
  confirms processed and cleared alarms and moves them to the historical
  alarm list »*). À utiliser avec parcimonie, sur des alarmes
  bien connues et sans impact.
- **Journal d'expérience** : la documentation mentionne l'enregistrement
  de l'expérience de traitement des alarmes (*Recording Alarm Handling
  Experience*) : capitalisez les résolutions (« alarme X → cause Y →
  action Z ») pour les nouveaux arrivants.

## 41. Notifications : mail et SMS

eSight permet de définir des **règles de notification** : pour une source
d'alarme + une sévérité données, notifier un **groupe d'utilisateurs**
par e-mail et/ou SMS.

Mise en place type :
1. Configurer le **serveur SMTP** (relais de l'entreprise) et tester.
2. Configurer la **passerelle SMS** (interface SMS server — selon
   l'opérateur/équipement, **à vérifier sur la documentation officielle**).
3. Créer les **groupes** : `astreinte-reseau`, `astreinte-systeme`,
   `managers` (qui ne reçoit que le critical).
4. Créer les règles : ex. *Critical sur backbone → SMS + mail à
   astreinte-reseau ; Major → mail ; Minor/Warning → pas de notification
   (consultation en console)*.
5. **Tester** chaque règle (alarme de test) et vérifier la réception
   réelle sur les téléphones.

**Règles d'hygiène :**
- Le SMS est réservé au **critical** (sinon on ne lit plus les SMS).
- Inclure dans le message : sévérité, équipement, libellé, heure —
  pas un « alarme eSight » générique qui oblige à se connecter pour
  comprendre.
- Tester les notifications **tous les mois** (un relais SMTP qui change
  d'IP et plus personne n'est prévenu — grand classique, voir cas n°9).

## 42. Agrégation : lutter contre le bruit

L'**agrégation** (aggregation rules) regroupe les alarmes répétitives
identiques en une seule alarme « agrégée » avec compteur, au lieu
d'inonder la console.

Cas d'école : un lien qui flappe (up/down toutes les 2 minutes) génère
des dizaines d'alarmes. Avec agrégation : **une** alarme
« Interface flapping — 47 occurrences ».

