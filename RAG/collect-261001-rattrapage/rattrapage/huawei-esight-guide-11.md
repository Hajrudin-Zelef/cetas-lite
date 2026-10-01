---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-11
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1629, 1798]
sha256: 3414cd934affca6c4ebb2337663fb570dce203de00da1351bbd5dac108238e18
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**Symptômes** : eSight déclare un équipement injoignable alors qu'il
fonctionne (les utilisateurs ne voient rien).

**Diagnostic :**
1. **Perte de paquets** entre eSight et l'équipement ? (ping prolongé,
   MTR.) Un lien WAN chargé fait perdre les requêtes SNMP.
2. **Timeout trop court** dans le profil de découverte/polling.
3. L'équipement est-il **surchargé** (CPU à 100 % → il ne répond plus
   au SNMP) ? Dans ce cas l'alarme est un symptôme, pas un faux positif.
4. **IP dupliquée** sur le réseau ?

**Solution :** augmenter timeout/retries pour les sites distants,
traiter la cause réseau (QoS du trafic de management en priorité),
corriger le duplex/IP dupliquée. Ne jamais « régler » en allongeant
les timeouts à l'infini : ça masque les vrais problèmes.

## 85. Cas n°7 — Sauvegarde des configurations qui échoue

**Symptômes** : les tâches de backup de config passent en erreur sur
tout ou partie du parc.

**Diagnostic :**
1. Le compte **SSH/Telnet** utilisé par eSight est-il toujours valide
   (mot de passe expiré ? compte verrouillé après trop de tentatives ?).
2. L'accès SSH depuis l'IP d'eSight est-il autorisé (ACL vty) ?
3. L'équipement a-t-il changé d'IP ou de hostname ?
4. Espace disque côté eSight pour stocker les configs ?

**Solution :** compte de service **dédié** avec mot de passe
n'expirant pas (ou géré par rotation documentée), ACL vty autorisant
eSight, relance des tâches en échec après correction. **Alerte sur
l'échec des backups** : un backup qui échoue en silence pendant 6 mois,
c'est la catastrophe du jour où on en a besoin.

## 86. Cas n°8 — Les rapports sont vides ou incohérents

**Symptômes** : rapport de disponibilité à 0 % ou vide, chiffres
aberrants.

**Diagnostic :**
1. La **période** du rapport couvre-t-elle des données existantes ?
   (Rétention dépassée → données purgées → rapport vide.)
2. Les **équipements sources** étaient-ils bien supervisés sur la période ?
   (Un équipement découvert le 15 ne peut pas avoir de stats du 1er au 14.)
3. Les **seuils de disponibilité** sont-ils bien définis (qu'est-ce qui
   compte comme « down ») ?
4. Fuseau horaire du serveur vs équipements : un décalage fausse les
   fenêtres.

**Solution :** aligner périodes et rétention, vérifier la supervision
effective sur la période, synchroniser les horloges (NTP partout).

## 87. Cas n°9 — Les notifications mail/SMS ne partent plus

**Symptômes** : plus aucun e-mail/SMS d'alarme, alors que les alarmes
sont bien dans la console.

**Diagnostic :**
1. Le **relais SMTP** répond-il ? (IP changée ? authentification
   modifiée ? TLS exigé désormais ?)
2. La **passerelle SMS** est-elle joignable (crédit épuisé ? APN modifié ?).
3. Les **règles de notification** sont-elles toujours actives (pas
   désactivées lors d'une maintenance puis oubliées) ?
4. Les **groupes destinataires** sont-ils à jour (départs, changements
   de numéros) ?

**Solution :** corriger le connecteur, **tester mensuellement**
l'envoi réel (alarme de test → vérifier la réception sur le téléphone),
revoir les destinataires à chaque mouvement d'équipe.

## 88. Cas n°10 — eSight lent (console qui rame)

**Symptômes** : pages qui mettent 30 s à s'afficher, timeouts navigateur.

**Diagnostic :**
1. **Ressources serveur** : CPU, RAM (swap ?), disque (plein ? I/O
   saturés ?). Voir le System Monitor Tool.
2. **Base de données** : taille, fragmentation, requêtes lentes —
   le coupable n°1 (voir cas n°13).
3. **Trop d'alarmes courantes** non soldées (des dizaines de milliers
   d'alarmes actives ralentissent l'affichage : solder/nettoyer).
4. **Sessions** : trop de consoles ouvertes simultanément ?
5. Antivirus qui scanne les répertoires eSight/base en temps réel.

**Solution :** selon la cause — ajouter RAM, purger les alarmes soldées,
exclure les répertoires de l'antivirus temps réel, redimensionner la VM.
Si ça persiste : **Fault Information Collection Tool** → package pour
le support.

## 89. Cas n°11 — Un service eSight ne démarre plus

**Symptômes** : après reboot ou mise à jour, un service reste en erreur,
la console est partiellement ou totalement inaccessible.

**Diagnostic :**
1. Consulter les **logs** du service (répertoire de logs eSight).
2. La **base de données** est-elle démarrée et accessible ? (La plupart
   des services eSight en dépendent : base down = tout down.)
3. **Port déjà utilisé** par un autre processus ?
4. **Espace disque** plein (les logs eux-mêmes peuvent remplir le disque).
5. Changement récent (patch OS, mise à jour) : quoi, quand ?

**Solution :** libérer le disque, démarrer la base d'abord puis les
services dans l'ordre, restaurer la sauvegarde pré-changement si
nécessaire. **Toujours snapshoter avant une mise à jour** (section 16).

## 90. Cas n°12 — Base de données pleine / croissance anormale

**Symptômes** : alertes d'espace disque, eSight qui se fige, sauvegardes
qui échouent.

**Diagnostic :**
1. Qu'est-ce qui grossit ? Données de **performance** (pas de collecte
   trop fin sur trop d'indicateurs), **alarmes historiques** jamais
   purgées, **logs** applicatifs, **fichiers de collecte** oubliés.
2. Vérifier les **politiques de rétention/purge** : sont-elles actives ?
3. Un équipement qui **flap** génère des milliers d'alarmes/jour :
   traiter la cause (section 45), pas seulement purger.

**Solution :** purger selon la politique (jamais de DELETE sauvage en
SQL direct), ajuster la rétention et le pas de collecte, archiver les
données anciennes avant purge si besoin réglementaire, étendre le disque.
**Prévention** : supervision de l'espace disque d'eSight lui-même
(voir section 16 — eSight doit se superviser).

## 91. Cas n°13 — Performances dégradées : la base rame

**Symptômes** : lenteurs générales malgré des ressources serveur correctes.

**Diagnostic :**
1. Taille de la base vs RAM : si la base active ne tient plus en mémoire,
   chaque requête va au disque.
2. **Index/fragmentation** : une base qui n'a jamais été maintenue
   (rebuild d'index, statistiques) se dégrade avec le temps.
3. Rétention trop longue pour le dimensionnement (voir cas n°12).
4. Requêtes concurrentes : rapports lourds lancés en pleine journée +
   polling + consoles utilisateurs.

**Solution :** maintenance de la base (selon le SGBD : rebuild index,
mise à jour des stats — **via les outils fournis ou avec l'aval du
support**, jamais d'opération exotique), déplacer les gros rapports la
nuit, réduire la rétention ou augmenter les ressources. Si récurrent :
revoir le dimensionnement (section 3) avec le partenaire.

## 92. Cas n°14 — Après une coupure électrique : tout est rouge

**Symptômes** : au retour du courant, des centaines d'alarmes, des
équipements qui ne reviennent pas dans eSight.

**Diagnostic :** distinguer le **vrai** du **fantôme** :
1. Quels équipements sont réellement joignables (ping) ?
2. Lesquels ne sont pas repartis (alim, boot bloqué) ?
3. Les alarmes « équipement down » d'équipements qui sont en fait UP =
   traps perdus pendant la coupure → **synchroniser** (section 48).

**Solution :** traiter les vrais down (intervention terrain), synchroniser
les alarmes pour solder les fantômes, vérifier que les sauvegardes de
config planifiées n'ont pas tourné dans le vide pendant la coupure.
**Leçon** : ce scénario se prépare — runbook « retour de coupure »
(section 15) et onduleurs supervisés (votre domaine !).

## 93. Cas n°15 — Changement d'IP du serveur eSight : tout casse

**Symptômes** : après migration/changement d'IP du serveur, plus de
traps, SNMP en échec, console injoignable à l'ancienne adresse.

**Diagnostic :** recenser tout ce qui référençait l'ancienne IP :
trap-targets des équipements, ACL SNMP, règles firewall, favoris des
exploitants, relais, supervision d'eSight lui-même, DNS.

