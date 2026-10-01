---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-20
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1815, 1907]
sha256: 21962efef9859cb1076c8dab876bd3b63f0122b05f47ceefb528deefa938a9d0
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

**Résolution** : rétablir le lien (opérateur/secours), vérifier la resynchronisation (les équipements repassent « normal », les configs se resync).

**Leçon** : pour les sites critiques, **composant d'authentification local + lien de secours** — le centralisé pur est un pari sur le WAN.

## 137. Cas pratique 14 — Conflit entre config manuelle CLI et NCE

**Symptôme** : NCE signale des **écarts de conformité** sur un switch ; un technicien « a juste ajouté une ligne en CLI ».

**Diagnostic** : la vérification de conformité (get-config vs attendu) montre exactement les différences — c'est fait pour ça.

**Conduite** : (1) Comprendre **pourquoi** le technicien est passé en CLI (urgence ? template incomplet ? habitude ?). (2) Si la modif est légitime : l'intégrer **au template** (versionner, déployer proprement). (3) Si non : l'écraser par le template (NCE = source de vérité). (4) Dans les deux cas : en parler en REX, pas en sanction (la première fois).

**Leçon** : tout besoin récurrent de CLI = un **template incomplet** — le corriger plutôt que de verrouiller le CLI (qui reste nécessaire en dépannage).

## 138. Cas pratique 15 — Certificat expiré : tout s'arrête

**Symptôme** : portail NCE inaccessible (certificat HTTPS expiré) et/ou EAP en échec massif (certificat du serveur RADIUS expiré).

**Réaction** : (1) Confirmer (navigateur/logs : « certificate expired »). (2) Appliquer la **procédure de renouvellement d'urgence** (régénérer/déployer — section 106). (3) Vérifier tous les services après renouvellement.

**Leçon** : le registre des certificats avec alertes J-90/J-60/J-30 (section 106) aurait évité ça. C'est l'incident le plus **évitable** de cette liste — le mettre au plan de maintenance préventive avec un responsable nommé.

## 139. Cas pratique 16 — Restauration après sinistre du serveur NCE

**Symptôme** : le serveur NCE est perdu (panne matérielle, ransomware...).

**Procédure** : (1) Reconstruire la plateforme (nouvelle VM selon la doc d'installation, **même version** NCE). (2) Restaurer le **backup applicatif** (section 108). (3) Restaurer la licence (fichier archivé). (4) Vérifier : inventaire, un onboarding test, alarmes, API, supervision. (5) Les équipements se **reconnectent** au contrôleur restauré (ils gardent leur config locale en attendant).

**Leçon** : le **RTO** (temps de restauration) se connaît parce qu'on l'a **testé** (exercice annuel), pas parce qu'on l'a estimé. Archiver hors site : le backup sur le même rack que le serveur, c'est pas un backup.

## 140. Cas pratique 17 — Montée en charge : NCE sature

**Symptôme** : portail lent, alarmes en retard, traitements qui s'accumulent — après l'ajout de 200 équipements.

**Diagnostic** : (1) Ressources de la VM (CPU/RAM/disque/IOPS — la télémétrie et les bases sont gourmandes). (2) Rétention des données (trop longue ? granularité trop fine ?). (3) Le dimensionnement initial prévoyait-il cette charge ?

**Résolution** : augmenter les ressources (si VM), ajuster la **rétention** (données fines 30 j, agrégées au-delà), répartir la télémétrie, et si structurel : re-dimensionner avec l'outil officiel (voire ajouter un nœud selon l'architecture).

**Leçon** : le dimensionnement se fait pour **l'année 3**, pas pour le jour 1 (section 23-25) — avec des paliers de croissance planifiés.

## 141. Cas pratique 18 — Intégration RADIUS/LDAP qui casse l'authentification admin

**Symptôme** : après le raccordement de NCE à l'AD, plus personne ne se connecte au portail (même les comptes qui marchaient).

**Diagnostic** : (1) Le connecteur LDAP pointe-t-il sur le bon serveur/port ? (2) Le compte de liaison (bind) est-il valide (mot de passe expiré ?). (3) Les groupes AD sont-ils bien mappés aux rôles NCE ? (4) L'AD est-il joignable depuis NCE (pare-feu) ?

**Filet** : le **compte local de secours** (break-glass, section 103) permet de se connecter et de réparer — c'est exactement pour ça qu'il existe.

**Leçon** : toute intégration d'authentification se fait **en maintenant un accès local de secours**, et se teste avec un compte de test avant de basculer les admins.

## 142. Cas pratique 19 — eKit et NCE se disputent un AP

**Symptôme** : un AP enrôlé autrefois dans eKit refuse de s'enregistrer dans NCE (ou disparaît de NCE pour réapparaître dans eKit).

**Cause** : l'AP reste **lié** à la plateforme cloud eKit (il cherche à s'y enregistrer).

**Résolution** : **désenrôler proprement** l'AP d'eKit (retirer de la plateforme), le **réinitialiser** en configuration d'usine, puis l'onboarder dans NCE (ZTP ou manuel).

**Leçon** : un équipement = **une** plateforme de gestion. L'inventaire de migration doit tracer la plateforme d'origine de chaque équipement (eKit, eSight, local) pour prévoir le désenrôlement.

## 143. Cas pratique 20 — Audit de sécurité : durcir en urgence

**Symptôme** : un audit révèle : comptes partagés, SNMP v2c en `public`, portail exposé sans restriction, pas de revue des rôles depuis 2 ans.

**Plan d'action (30-60-90 jours)** :
- **J+30** : changer les mots de passe, supprimer les comptes partagés (comptes nominatifs), passer en SNMPv3, restreindre l'accès au portail (VPN/allowlist).
- **J+60** : revoir tous les rôles (moindre privilège), activer/renforcer l'audit, exporter les logs vers le SIEM, inventorier les certificats.
- **J+90** : test d'intrusion ciblé ou audit de configuration, exercice de restauration, documentation à jour.

**Leçon** : appliquer la checklist de durcissement (section 105) **dès l'installation** — le rattrapage coûte toujours plus cher que la prévention, et l'audit n'est pas une surprise si on s'auto-audite chaque année.

## 144. Cas pratique 21 — Migration eSight → NCE d'un site pilote

**Contexte** : premier site migré (12 switches S310, 18 AP361, 1 AR720).

**Déroulé** : (1) Inventaire et ESN relevés depuis eSight + terrain. (2) Templates du site créés à partir des configs eSight **retranscrites** (pas copiées). (3) Maquette : 1 switch + 2 AP validés. (4) Fenêtre un samedi : désenrôlement eSight, onboarding NCE (ZTP pour les AP, manuel assisté pour les switches — pas de réinitialisation massive sans filet). (5) Tests : 802.1X avec 3 profils, SSID, débit, roaming. (6) 72 h d'observation avant clôture.

**Incident rencontré** : 2 AP restés en « non enregistré » (ESN mal recopiés — erreur de saisie). Corrigé en 20 min grâce à la liste blanche vérifiée à la réception.

**Leçon** : le pilote sert à **découvrir** les problèmes — le planifier comme tel (temps, astreinte, droit à l'échec), et en faire un REX écrit avant la vague suivante.

## 145. Cas pratique 22 — Rapport mensuel pour la direction

**Besoin** : chaque mois, la direction veut une page sur l'état du réseau.

**Mise en œuvre** : (1) Définir les **4-5 indicateurs** qui parlent à la direction : disponibilité par site (%), incidents majeurs (nombre, durée), expérience Wi-Fi (taux de satisfaction ou plaintes), projets en cours, risques à venir. (2) Construire le rapport **automatique** dans NCE (planifié, envoyé par mail). (3) Le compléter d'un **commentaire humain** (3 lignes : ce qui s'est bien passé, ce qui a coincé, ce qu'on fait).

**Leçon** : un rapport automatique sans commentaire = un PDF que personne ne lit. Le commentaire du chef de service, c'est la valeur ajoutée — et c'est ce qui justifie les budgets.

## 146. Cas pratique 23 — Supervision croisée NCE + Zabbix

**Contexte** : Zelef supervise déjà ses onduleurs dans Zabbix ; il veut corréler avec NCE.

