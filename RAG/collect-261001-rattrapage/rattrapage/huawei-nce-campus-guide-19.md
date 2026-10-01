---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-19
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1725, 1814]
sha256: 39fafaf43b99813c2bf7b226e00cde3aa9d55c2a610501030f1d2890129278d6
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

**Résolution** : mettre à jour le firmware **manuellement** (USB/console) vers la version minimale, régler l'horloge, lever le doublon. **Leçon** : l'audit de versions **avant** la vague d'onboarding évite ce cas (checklist section 50).

## 128. Cas pratique 5 — Un AR720 qui ne télécharge pas sa config

**Symptôme** : l'AR720 s'enregistre mais le template ne s'applique pas (ou partiellement).

**Diagnostic** : (1) Le template est-il bien **affecté** au site/à l'équipement ? (2) Les **variables** sont-elles renseignées ? (un template avec une variable vide = échec silencieux ou partiel). (3) La taille du template vs les capacités (un AR d'entrée de gamme n'accepte pas tout). (4) Les écarts affichés par la vérification de conformité (get-config vs attendu).

**Résolution** : renseigner les variables manquantes, simplifier le template, appliquer par étapes. **Leçon** : un template se **teste variable par variable** en maquette — jamais directement sur l'AR de production d'un site distant.

## 129. Cas pratique 6 — Un template qui casse la prod

**Symptôme** : après application d'une nouvelle version de template, un site perd son accès (VLAN de management écrasé, uplink en erreur...).

**Réaction immédiate** : (1) Ne pas paniquer, ne pas multiplier les changements. (2) **Rollback** vers la version précédente du template (section 87) — c'est pour ça qu'elle existe. (3) Si le management est coupé : intervention locale (console) avec la procédure d'urgence.

**Analyse** : rejouer le template en maquette, identifier la ligne fautive (souvent : une variable non substituée, un VLAN oublié sur un trunk, une ACL trop restrictive). Corriger, **re-tester**, re-versionner.

**Leçons** : vagues de déploiement (section 84), jamais de big bang ; fenêtre de maintenance ; template relu par un pair ; rollback testé. Ce cas arrive à tout le monde — ce qui distingue les équipes, c'est le **temps de retour**, pas l'absence d'erreur.

## 130. Cas pratique 7 — Licence dépassée : que se passe-t-il ?

**Symptôme** : le compteur device-days passe en négatif ; alerte « license exceeded ».

**Ce qui se passe** : période de **grâce de 30 jours** (documentée) — le système continue de fonctionner. C'est le délai pour régulariser, pas pour ignorer.

**Conduite** : (1) Vérifier le compteur : croissance normale ou anomalie (équipements en double ? site oublié ?). (2) Commander le complément (procédure d'achat d'urgence — section 29). (3) Importer et valider la licence. (4) Post-mortem : pourquoi l'alerte à 80 % n'a-t-elle pas été vue ? (mettre en place le suivi trimestriel).

**Leçon** : la licence est un **poste d'exploitation**, pas un achat ponctuel — elle a sa ligne au budget et son responsable.

## 131. Cas pratique 8 — Mise à jour firmware massive qui échoue à mi-parcours

**Symptôme** : sur 60 AP, 40 sont à jour, 20 sont restés à l'ancienne version (ou en échec).

**Diagnostic** : (1) Les 20 en échec ont-ils téléchargé le package ? (serveur de fichiers joignable ? espace disque sur l'AP ?). (2) Ont-ils redémarré ? (3) La version est-elle compatible avec leur modèle exact ? (4) Les logs de la tâche d'upgrade dans NCE.

**Résolution** : relancer la tâche **uniquement** sur les en-échec (pas sur tout le monde), après correction de la cause (espace, connectivité, compatibilité). Vérifier la conformité post-upgrade.

**Leçons** : vagues (section 85), ne jamais upgrader 100 % d'un coup ; garder la version précédente validée pour downgrade ; planifier en fenêtre de maintenance (les AP redémarrent = coupure Wi-Fi).

## 132. Cas pratique 9 — Roaming Wi-Fi dégradé après migration NCE

**Symptôme** : après la migration sous NCE, les appels Wi-Fi coupent en se déplaçant (alors que ça marchait avant).

**Diagnostic** : (1) Le profil SSID migré a-t-il bien **802.11k/v/r** activés ? (souvent oubliés dans la retranscription — section 76). (2) Les puissances/canaux sont-ils les mêmes qu'avant ? (le RRM a peut-être tout rebattu). (3) Le recouvrement de couverture est-il suffisant ? (mesure terrain). (4) Les logs de roaming par client dans NCE.

**Résolution** : réactiver k/v/r, figer temporairement le RRM le temps de stabiliser, ajuster les puissances, **tester en marchant avec un appel réel**.

**Leçon** : la migration WLAN = **retranscription + re-validation radio**, pas un simple transfert de SSID. Le test d'acceptation, c'est l'utilisateur qui marche, pas le ping.

## 133. Cas pratique 10 — 802.1X qui rejette tous les utilisateurs un lundi matin

**Symptôme** : à 8 h, personne ne se connecte en filaire ; les ports restent bloqués.

**Diagnostic express** : (1) Le **RADIUS** répond-il ? (service NCE/authentification, connectivité, certificats). (2) L'**AD/LDAP** est-il joignable ? (un AD en panne = tout le monde rejeté). (3) Y a-t-il eu un **changement** ce week-end ? (certificat EAP expiré ? mise à jour ?). (4) Les logs RADIUS : la cause exacte du rejet (identifiants ? certificat client ?).

**Réaction** : si la cause n'est pas trouvée en < 30 min, basculer les ports critiques en **mode ouvert/monitor** temporaire (procédure d'urgence écrite à l'avance) pour rétablir le service, puis diagnostiquer à tête froide.

**Leçons** : le 802.1X a un **mode dégradé** prévu et testé ; les certificats EAP sont suivis (section 106) ; tout changement le week-end = vérification du 802.1X le lundi à 7 h 30 par l'astreinte.

## 134. Cas pratique 11 — Le contrôleur NCE ne répond plus

**Symptôme** : le portail est inaccessible, les API ne répondent pas.

**Rappel crucial** : le réseau **continue de fonctionner** avec la dernière configuration (le plan de données est local) — ce qui s'arrête : supervision, nouveaux déploiements, authentification centralisée (sauf composant local).

**Diagnostic** : (1) La VM tourne-t-elle ? (hyperviseur). (2) Espace disque plein ? (cause n°1 des plantages applicatifs — logs/bases). (3) Services NCE démarrés ? (4) Réseau (la VM a-t-elle perdu son IP ?).

**Résolution** : selon la cause (libérer l'espace, redémarrer les services, basculer sur le nœud HA, restaurer depuis le backup en dernier recours — section 108).

**Leçons** : supervision **du** contrôleur par Zabbix (qui supervise le superviseur), alerte disque à 75 %, exercice de bascule HA annuel, backup testé.

## 135. Cas pratique 12 — Faux positifs d'alarmes après un orage

**Symptôme** : après un orage, NCE affiche des dizaines d'alarmes (ports, AP, télémétrie).

**Diagnostic** : trier — (1) Alarmes **physiques réelles** (équipement vraiment down → intervention, voir cas 3). (2) Alarmes **transitoires** (micro-coupures, flaps — acquitter après vérification). (3) **Faux positifs ML** (l'anomalie détectée n'en est pas une — le modèle apprend).

**Conduite** : traiter le réel d'abord (site par site), acquitter le transitoire avec commentaire, **marquer** les faux positifs pour affiner les seuils (le ML a besoin de retours).

**Leçon** : après chaque événement météo majeur, prévoir une **ronde d'alarmes** (30 min) avant de crier au désastre — et affiner les seuils au fil des événements.

## 136. Cas pratique 13 — Un site distant perd le lien vers NCE

**Symptôme** : tout un site passe « hors ligne » dans NCE.

**Diagnostic** : (1) Le lien WAN est-il coupé ? (opérateur, AR720, 4G de secours). (2) Si le lien est OK : routage ? pare-feu ? (3) L'AR720 du site répond-il ?

**Comportement** : le site **continue de fonctionner** localement (dernière config) ; l'authentification continue si le **composant d'authentification local** est déployé (section 21) — sinon, les nouvelles authentifications 802.1X peuvent échouer (prévoir le mode dégradé).

