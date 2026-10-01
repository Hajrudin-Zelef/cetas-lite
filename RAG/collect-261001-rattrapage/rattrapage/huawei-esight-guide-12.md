---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-12
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1799, 1969]
sha256: 7e3847ec203f1168d2bef8e4c64175f9160fdb6f82d5df903dde1e0076d07ff8
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**Solution :** mettre à jour **partout** (c'est un chantier : script
de déploiement de la nouvelle IP en trap-target via le Smart Config
Tool, section 60), mettre à jour le DNS, communiquer la nouvelle URL.
**Prévention** : utiliser un **nom DNS** partout où c'est possible
plutôt que l'IP en dur, et documenter l'IP comme paramètre critique
du dossier d'exploitation. Idéalement : ne jamais changer l'IP d'un
NMS en production.

## 94. Cas n°16 — Compte verrouillé / plus d'accès admin

**Symptômes** : personne ne peut se connecter en admin (mot de passe
perdu, compte verrouillé après tentatives, AD indisponible).

**Diagnostic :** s'agit-il du compte local ou de l'authentification
annuaire ? L'AD répond-il ?

**Solution :**
- Si AD en panne : utiliser le **compte admin local de secours**
  (celui du coffre-fort, section 68).
- Si mot de passe local perdu : procédure de réinitialisation
  constructeur (**à vérifier sur la documentation officielle** —
  généralement via accès console au serveur eSight).
- Après récupération : changer les mots de passe, vérifier les
  verrouillages, consigner au coffre.
- **Prévention** : deux détenteurs du coffre-fort, jamais un seul.

## 95. Cas n°17 — Montée de version qui se passe mal (rollback)

**Symptômes** : après upgrade, fonctionnalités cassées, services
instables, données manquantes.

**Diagnostic :** identifier ce qui a changé (notes de version lues ?
prérequis vérifiés ? backup pré-upgrade existant ?).

**Solution :**
1. Ne pas « réparer » à chaud pendant des heures : si le rollback est
   possible (snapshot VM, sauvegarde pré-upgrade), **revenir en arrière**
   proprement — c'est pour ça qu'on snapshotte.
2. Restaurer la sauvegarde pré-upgrade (section 71).
3. Analyser à froid : cause de l'échec, prérequis manqué, puis
   replanifier avec le partenaire/support.
4. **Règle** : jamais d'upgrade en production sans : snapshot + sauvegarde
   testée + procédure de rollback écrite + fenêtre de maintenance +
   binôme.

---

# 15. BONNES PRATIQUES D'EXPLOITATION

## 96. Organiser l'exploitation autour d'eSight

- **Le NOC vit dans eSight** : topologie + alarmes courantes = écran
  principal. Tout le reste (tickets, docs) gravite autour.
- **Rôles clairs** : qui surveille (NOC), qui traite (niveau 2/3),
  qui décide (chef de service). eSight reflète cette organisation
  dans ses rôles et domaines (section 11).
- **La console n'est pas un mur d'alarmes** : si le NOC voit 200 alarmes
  non acquittées en permanence, le système est mal réglé (revoir la
  politique d'alarmes, section 50) — pas l'équipe qui est négligente.
- **eSight ne remplace pas le terrain** : une alarme « température haute »
  dans un local technique se termine par quelqu'un qui va vérifier la
  clim. Le NMS détecte, l'humain constate et agit.

## 97. Astreinte : que l'outil serve l'astreinte, pas l'inverse

- L'astreinte ne reçoit que le **critical** (SMS) + le **major** (mail).
  Tout le reste attend le lendemain 8h.
- **Procédure d'astreinte écrite** : réception du SMS → connexion VPN →
  console eSight → acquitter avec commentaire → appliquer le runbook →
  escalader si besoin. Un astreintier qui improvise à 3h du matin,
  c'est un incident qui dure 3 fois plus longtemps.
- **Test mensuel** du SMS d'astreinte (section 41).
- **Rotation** : personne ne doit être le seul à savoir lire eSight.
  La formation des nouveaux inclut « lire une alarme et l'acquitter »
  dès la première semaine.

## 98. Runbooks : un par alarme critique

Un **runbook** = la fiche réflexe pour une alarme donnée :

```markdown
# RUNBOOK — Alarme : Lien backbone SITE-A ↔ SITE-B down (Critical)

## Symptômes dans eSight
- Alarme Critical « linkDown » sur l'interface concernée + alarmes
  corrélées « équipements injoignables » sur SITE-B.

## Vérifications (dans l'ordre)
1. Topologie : le lien est-il vraiment down ? (pas un flap : voir
   l'historique de l'interface sur 1h)
2. Équipements d'extrémité : joignables en SSH ? (si oui : problème
   de lien ; si non : problème d'équipement ou d'énergie)
3. Fournisseur : incident connu sur la liaison ? (numéro du support
   dans le PRA)

## Actions
- Si flap : masquage temporaire + ticket + intervention planifiée.
- Si down franc : bascule sur le lien de secours (procédure PRA),
  ticket fournisseur, information des utilisateurs.
- Acquitter l'alarme avec le n° de ticket.

## Escalade
- Si non résolu en 1h → chef de service. Si impact utilisateurs
  majeur → cellule de crise.

## Retour d'expérience
- Compléter le journal d'expérience eSight après clôture.
```

**Objectif** : runbook pour les **10 alarmes les plus fréquentes**
d'abord, puis élargir. Un runbook se teste (exercice à blanc) et se
met à jour après chaque incident réel.

## 99. KPI d'exploitation : piloter avec eSight

Tableau de bord mensuel du chef de service (alimenté par les rapports
eSight, section 55) :

| KPI | Cible indicative | Source |
|---|---|---|
| Disponibilité par site | ≥ 99,5 % (à définir par contrat) | Rapport SLA |
| Alarmes critical/major par mois | Tendance à la baisse | Rapport alarmes |
| Délai moyen d'acquittement (critical) | ≤ 15 min | Logs d'alarmes |
| % d'alarmes avec runbook associé | 100 % pour critical | Revue manuelle |
| Équipements non supervisés (hors parc) | 0 | Inventaire vs parc réel |
| Succès des sauvegardes de config | 100 % | Rapport tâches |
| Succès des sauvegardes eSight | 100 % | Contrôle backup |

**Utilisation** : ces KPI vont au COPIL, justifient les investissements
(« les liens à 85 % : il faut upgrader ») et mesurent le progrès de
l'équipe. Un KPI qui ne déclenche aucune décision est un KPI à supprimer.

## 100. Rituels d'équipe autour de l'outil

- **Quotidien (5 min)** : revue des alarmes de la nuit par le NOC,
  point à la prise de poste.
- **Hebdomadaire (30 min)** : revue des masquages actifs, des alarmes
  récurrentes, des backups en échec.
- **Mensuel (1 h)** : comité alarmes (politique d'alarmes, section 50),
  revue des KPI, plan d'action (seuils à ajuster, équipements à traiter).
- **Trimestriel** : nettoyage topologie + inventaire, revue des comptes,
  test de restauration (annuel pour la restauration complète).
- **Annuel** : test PRA, revue du dimensionnement, renouvellement SnS,
  plan de formation.

---

# 16. MAINTENANCE

## 101. Plan de maintenance annuel type

| Période | Actions |
|---|---|
| Quotidien | Revue alarmes nuit ; contrôle succès sauvegardes (eSight + configs) ; espace disque eSight |
| Hebdomadaire | Revue masquages ; alarmes récurrentes ; patchs OS critiques en attente |
| Mensuel | Comité alarmes ; test notifications SMS/mail ; Environmental Health Check ; revue logs d'audit (échantillon) ; exports inventaire |
| Trimestriel | Nettoyage topologie/inventaire ; revue comptes et rôles ; vérification rétention/purge ; test d'un runbook à blanc |
| Semestriel | Revue des seuils ; revue du dimensionnement ; exercice PRA partiel |
| Annuel | **Test de restauration complet** ; test PRA grandeur nature ; renouvellement licences/SnS ; plan de formation ; revue d'architecture |

## 102. Check-list quotidienne (NOC — 10 minutes)

- [ ] Alarmes critical/major non acquittées : 0 (ou chacune a un ticket).
- [ ] Équipements injoignables : liste connue et justifiée ?
- [ ] Sauvegarde eSight de la nuit : succès ?
- [ ] Sauvegardes de configs : 0 échec (ou échecs traités) ?
- [ ] Espace disque du serveur eSight : < 80 % ?
- [ ] Services eSight : tous démarrés ?

## 103. Check-list mensuelle (admin eSight — 1 heure)

