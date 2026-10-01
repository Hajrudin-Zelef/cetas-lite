---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-15
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [2252, 2321]
sha256: ac4d25e1acdc6b0cf3552f1581ef9832039e2f6ceea2ac9906b317742f205e08
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

**Q7. À quoi sert un RODC dans une agence sans local sécurisé ?**
<details><summary>Réponse</summary>Il authentifie localement sans exposer une base inscriptible : en cas de vol du serveur, l'attaquant ne récupère qu'une copie en lecture seule, et par défaut aucun mot de passe n'est mis en cache (sauf stratégie de réplication autorisée). Les écritures sont relayées vers un DC inscriptible.</details>

**Q8. Vous supprimez par erreur l'OU « Compta » (200 utilisateurs). La corbeille AD est activée. Que faites-vous ?**
<details><summary>Réponse</summary>`Get-ADObject -Filter {Name -eq "Compta"} -IncludeDeletedObjects | Restore-ADObject`, puis restaurer les objets enfants, réactiver les comptes, vérifier les appartenances aux groupes (conservées grâce à la corbeille). Pas besoin de restauration faisant autorité.</details>

**Q9. Quelle GPO définit la longueur minimale des mots de passe du domaine, et peut-on en mettre une différente par service ?**
<details><summary>Réponse</summary>La **Default Domain Policy** (liée à la racine du domaine) définit la stratégie par défaut. Oui pour différencier : les **FGPP/PSO** permettent des politiques par groupe (ex. 16 caractères pour les admins), avec un ordre de précédence.</details>

**Q10. `repadmin /replsummary` montre des échecs croissants entre deux sites. Citez 4 causes possibles et les commandes de diagnostic.**
<details><summary>Réponse</summary>Causes : liaison réseau/VPN coupée, DNS défaillant, décalage d'heure > 5 min (Kerberos), pare-feu bloquant les ports AD. Diagnostics : `dcdiag /test:dns /v`, `repadmin /showrepl`, `w32tm /query /status`, `Test-NetConnection DC-Distant -Port 389`, vérification des liaisons de sites et coûts.</details>

---

## 78. Erreurs classiques (18)

### Erreur n°1 — Suppression d'une OU « protégée » impossible (et inversement, OU non protégée supprimée par erreur)
**Symptôme** : « Accès refusé » à la suppression d'une OU alors qu'on est admin du domaine.
**Cause** : la case « Protéger l'objet contre une suppression accidentelle » est cochée (c'est le comportement par défaut à la création via les consoles).
**Remède** : décocher la protection (ou `Set-ADOrganizationalUnit -ProtectedFromAccidentalDeletion $false`), puis supprimer.
**Prévention** : activez la protection sur **toutes** les OU de production via script (section 17) ; l'erreur inverse (suppression d'une OU non protégée) se rattrape avec la corbeille AD (section 49).

### Erreur n°2 — GPO liée mais jamais appliquée : l'objet est dans le conteneur par défaut
**Symptôme** : `gpresult /r` ne montre pas la GPO ; pourtant elle est liée « au domaine ».
**Cause** : le PC/l'utilisateur est dans `CN=Computers` / `CN=Users`, pas dans une OU. Les GPO liées à des **OU** ne touchent pas les conteneurs par défaut, et on ne peut pas lier de GPO à un conteneur `CN=`.
**Remède** : déplacer l'objet dans la bonne OU ; exécuter `redircmp`/`redirusr` une fois pour toutes (section 26).
**Prévention** : `redircmp` dès la création du domaine ; procédure de jonction incluant le pré-staging.

### Erreur n°3 — Filtrage de sécurité : GPO invisible à cause de MS16-072
**Symptôme** : après avoir remplacé « Utilisateurs authentifiés » par un groupe restreint, plus personne ne reçoit la GPO.
**Cause** : depuis 2016, le **compte ordinateur** doit avoir le droit **Lecture** sur la GPO pour la traiter. En retirant « Utilisateurs authentifiés », on a retiré la lecture aux ordinateurs.
**Remède** : ajouter « Ordinateurs du domaine » (ou « Utilisateurs authentifiés ») en `GpoRead`, et le groupe cible en `GpoApply` (section 59).
**Prévention** : ne jamais mettre `None` sans ajouter la lecture ; tester sur un poste pilote.

### Erreur n°4 — Heure non synchronisée : authentifications Kerberos en échec
**Symptôme** : ouvertures de session impossibles, erreurs « l'horloge n'est pas synchronisée », réplication en échec, `dcdiag` rouge.
**Cause** : décalage > 5 minutes entre client et DC (Kerberos), souvent après une panne NTP, une VM dont l'horloge dérive, ou un PDC mal configuré.
**Remède** : `w32tm /resync` sur le poste ; vérifier la hiérarchie (`w32tm /query /status`), reconfigurer le PDC sur des sources externes fiables (section 32) ; désactiver la synchro d'heure de l'hyperviseur vers les DC (le PDC fait foi).
**Prévention** : supervision du décalage NTP ; le PDC synchronisé sur 2-3 sources externes ; GPO éventuelle pour forcer `w32time`.

### Erreur n°5 — Rétrograder un DC en l'éteignant simplement
**Symptôme** : erreurs de réplication persistantes, objets serveur fantômes, `dcdiag` en échec, lenteurs d'ouverture de session.
**Cause** : un DC éteint sans rétrogradation laisse ses métadonnées, ses rôles FSMO éventuels et ses enregistrements DNS.
**Remède** : `Remove-ADDomainController` / suppression via Sites et services (metadata cleanup, section 12), saisie des FSMO si besoin (section 35), nettoyage DNS.
**Prévention** : procédure écrite de décommission : transférer FSMO → `Uninstall-ADDSDomainController` → vérifier → désinstaller le rôle.

### Erreur n°6 — Saisir un rôle FSMO alors que le détenteur est juste injoignable
**Symptôme** : deux DC se croient maîtres RID ; SID en double ; objets impossibles à créer ; schéma divergent.
**Cause** : `seize` lancé pendant une simple coupure réseau ou une maintenance prolongée, puis l'ancien détenteur revient en ligne.
**Remède** : il n'y en a pas de bon — c'est une situation de crise : isoler l'ancien détenteur, le réinstaller from scratch, nettoyer les métadonnées. D'où la règle : **seize = mort définitive**.
**Prévention** : ne saisir qu'après confirmation que le serveur est irrécupérable (disque mort, VM détruite) ; toujours préférer le transfert.

### Erreur n°7 — Restore de snapshot d'un DC virtualisé (USN rollback)
**Symptôme** : mots de passe « qui ne passent plus », comptes créés qui disparaissent selon le DC interrogé, événement 2108/1084.
**Cause** : retour en arrière de l'USN (section 43).
**Remède** : rétrogradation forcée + metadata cleanup + repromotion du DC fautif. Pas de réparation « douce ».
**Prévention** : sauvegardes d'état système uniquement ; VM-GenerationID actif ; interdire les snapshots sur les DC dans la politique de virtualisation ; former l'équipe.

### Erreur n°8 — Schéma étendu sans sauvegarde (irréversible)
**Symptôme** : après une extension de schéma ratée (mauvais OID, attribut mal défini), impossible de revenir en arrière proprement.
**Cause** : les suppressions de classes/attributs du schéma sont **définitives dans leur effet** (on peut désactiver, pas effacer).
**Remède** : restauration faisant autorité de **toute la forêt** depuis une sauvegarde antérieure — opération lourde, à éviter.
**Prévention** : sauvegarde état système avant ; test en labo ; ne jamais étendre le schéma « pour essayer » en production.

### Erreur n°9 — Mot de passe DSRM inconnu le jour où on en a besoin
**Symptôme** : impossible de démarrer en DSRM pour restaurer ; restauration bloquée en pleine crise.
**Cause** : mot de passe défini à la promotion puis oublié, jamais documenté, jamais changé au départ des admins.
**Remède** : préventif uniquement — le changer régulièrement via `ntdsutil`/`Set-DSRMPassword` et le stocker au coffre (section 52).
**Prévention** : entrée « mots de passe DSRM » dans le coffre d'entreprise, rotation annuelle, vérification semestrielle.

