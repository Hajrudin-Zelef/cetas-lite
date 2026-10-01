---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-7
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [935, 1105]
sha256: 7f931f2438d136414c569e7d9a278bcba5ce7a6a956f8b98def5afd93f79e131
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

- Définir les règles d'agrégation sur les alarmes connues pour être
  bavardes (flapping, seuils limites).
- L'alarme agrégée doit rester **visible et acquittable** comme les autres.

## 43. Masquage : taire ce qu'on sait déjà

Les **règles de masquage** (masking rules) empêchent la remontée
d'alarmes pendant des situations connues :

- Fenêtre de **maintenance** sur un équipement : on masque ses alarmes
  (sinon le NOC est noyé et rate les vraies pannes ailleurs).
- Équipement en **dérangement connu** en attente de pièce : masqué avec
  date de fin, pas « oublié ».

**Discipline** : tout masquage a une **date de fin** et un **motif**.
Un masquage permanent = une panne invisible = à interdire.
Revue hebdomadaire des masquages actifs (section 16).

## 44. Redéfinition des sévérités et des noms

eSight permet de **redéfinir la sévérité et le type** des alarmes
(severity and type redefinition rules) et même leur **nom**
(name redefinition rules).

Usages :
- Descendre en Minor une alarme remontée Critical par défaut mais sans
  impact chez vous.
- Renommer une alarme au jargon de l'entreprise (« Alarme carte
  d'alimentation » → « Défaut alim — switch cœur ») pour que le NOC
  comprenne sans dictionnaire MIB.
- Aligner les types pour le reporting (toutes les alarmes « énergie »
  dans la même catégorie, par exemple — utile pour un chef de service
  systèmes & énergies qui veut suivre les alimentations/onduleurs).

## 45. Alarmes intermittentes (toggling) : les traiter à part

Les **alarmes intermittentes** (qui apparaissent/disparaissent en boucle)
polluent et masquent les vrais problèmes. La documentation prévoit des
**règles de traitement des alarmes intermittentes/toggling**.

- Identifier les équipements « toggleurs » chroniques (rapport
  d'occurrences).
- Traiter la **cause racine** (souvent : lien défectueux, SFP fatigué,
  seuil trop sensible) plutôt que de masquer.
- En attendant la réparation : règle de toggling pour contenir le bruit,
  avec date de fin (comme tout masquage).

## 46. Tempête d'alarmes (alarm storm) : le scénario catastrophe

**Symptômes** : des centaines/milliers d'alarmes en quelques minutes
(panne du cœur de réseau, boucle, flap massif).

**Conduite à tenir :**
1. **Ne pas tout acquitter en panique** : on perd la chronologie.
2. Utiliser la **corrélation** (section 47) : eSight peut rattacher les
   alarmes « conséquences » à l'alarme « cause racine » (ex. : tous les
   « équipement injoignable » rattachés au « lien backbone down »).
3. Traiter la cause racine ; les alarmes conséquences se solderont
   (clear) en cascade.
4. Après l'incident : analyser la tempête (quelles règles de corrélation
   manquaient ?) et les ajouter — chaque tempête doit rendre le système
   plus intelligent.

**Prévention** : seuils de tempête (si X alarmes en Y minutes → alarme
méta « storm suspected » + notification renforcée) — **selon les
possibilités de votre version, à vérifier sur la documentation officielle**.

## 47. Corrélation : trouver la cause racine

Les **règles de corrélation** (correlation rules) lient des alarmes entre
elles selon des critères (temps, topologie, type) pour désigner une
**alarme racine** et des alarmes **corrélées/conséquences**.

Exemple : à 14:02, « Lien inter-sites down » (racine) ; à 14:02:05,
« 12 équipements injoignables » (conséquences). Sans corrélation : 13
alarmes critiques à traiter. Avec : 1 alarme racine à traiter + 12
rattachées automatiquement.

- Construire les règles de corrélation sur la **topologie réelle**
  (dépendances : tel site dépend de tel lien).
- Les tester sur incidents passés (rejouer l'historique si possible).
- C'est un chantier **continu** : chaque incident majeur enseigne une
  nouvelle règle.

## 48. Synchronisation des alarmes

La **synchronisation** (synchronizing alarms) consiste à réaligner eSight
avec l'état réel des équipements (re-polling) : utile après une coupure
de communication entre eSight et les équipements (les traps perdus en
route laissent des alarmes « fantômes »).

- Après toute coupure réseau entre eSight et un site : **synchroniser**
  avant de traiter (sinon on traite des alarmes obsolètes).
- À intégrer au runbook « retour de coupure » (section 15).

## 49. Filtrage northbound : ne remonter que l'utile

Quand eSight remonte ses alarmes vers un NMS supérieur (manager-of-managers),
le **filtrage northbound** (northbound filtering rules) choisit **quelles**
alarmes sont transmises : par sévérité, par alarmes ciblées, par source
(équipements), avec des conditions avancées.

Objectif : le NMS central ne reçoit que le significatif (pas les warnings
d'un site distant). Sans filtrage : congestion et perte des vraies alertes
— la documentation le dit explicitement.

## 50. Mettre en place une politique d'alarmes : le plan en 5 étapes

1. **Recenser** : extraire les 50 alarmes les plus fréquentes sur 3 mois.
2. **Calibrer** : atelier équipe — sévérité juste, nom parlant, action
   attendue pour chacune (runbook, section 15).
3. **Régler** : redéfinitions, agrégations, masquages types, corrélations
   de base, notifications.
4. **Tester** : générer des alarmes de test, vérifier notifications et
   corrélation.
5. **Revoir** : comité mensuel « alarmes » — bruit résiduel, nouvelles
   règles, masquages à lever. Une politique d'alarmes est un organisme
   vivant.

---

# 8. SUPERVISION DES PERFORMANCES

## 51. Les indicateurs suivis

eSight collecte en polling SNMP les indicateurs standards :

**Équipement (santé) :**
- CPU (utilisation %, par processeur/carte sur les châssis).
- Mémoire (utilisation %, disponible).
- Température, état des ventilateurs et alimentations (selon MIB).
- Temps de fonctionnement (uptime) — un uptime qui retombe à zéro =
  reboot à investiguer.

**Interfaces (trafic) :**
- Débit entrant/sortant (bps), utilisation en % de la capacité.
- Erreurs, discards, paquets unicast/multicast/broadcast.
- État opérationnel et administratif.

**Selon modules** : performances serveurs (CPU/RAM/disque), stockage
(IOPS, latence, capacité), WLAN (utilisateurs par AP, bruit radio),
applications (temps de réponse d'URL, état des services).

## 52. Seuils : quand un chiffre devient une alarme

Un **seuil** (threshold) = une valeur limite sur un indicateur, dont le
dépassement génère une alarme.

Seuils de départ raisonnables (à affiner par équipement) :

| Indicateur | Warning | Critical | Remarque |
|---|---|---|---|
| CPU équipement | > 70 % pendant 15 min | > 90 % pendant 5 min | Toujours avec durée pour éviter les pics |
| Mémoire | > 80 % | > 95 % | Une fuite mémoire se voit en tendance |
| Utilisation lien | > 70 % (tendance) | > 90 % | Le critical sur un lien = risque de saturation |
| Erreurs interfaces | > 0,1 % des paquets | > 1 % | Zéro erreur = normal sur cuivre/fibre sain |
| Température | Selon spec constructeur - 10 °C | Selon spec | À adapter par modèle |
| Espace disque (serveurs) | > 80 % | > 95 % | Classique mais vital |

**Principes :**
- Toujours un **délai** (le dépassement doit durer) : un pic CPU de
  30 secondes n'est pas une alarme.
- Seuils **différenciés** : le seuil du lien datacenter n'est pas celui
  du lien d'une agence.
- Revoir les seuils **tous les 6 mois** : un seuil qui ne déclenche
  jamais est inutile ; un qui déclenche tout le temps est du bruit.

## 53. Tableaux de bord : l'écran du matin

Construire des **dashboards** par public :

