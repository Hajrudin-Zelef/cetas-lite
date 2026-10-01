---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-17
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [2406, 2500]
sha256: 7121519879ad19db6b8f21d72b71d01deaaac1fff6ef755517c7a5e7fda48115
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

### Cas n°3 — Départ d'un salarié : procédure complète
**Contexte** : départ d'un technicien un vendredi soir.
**Objectif** : couper les accès immédiatement, proprement, avec traçabilité.
**Solution** : script section 22 (désactivation + mot de passe aléatoire + déplacement OU + retrait des groupes + description datée). Vérifier `Search-ADAccount -LockedOut` n'est pas concerné, archiver la fiche.
**Commentaire** : ne supprimez pas le compte tout de suite (besoin RH/legal, réattribution de fichiers). La suppression intervient après le délai défini (ex. 90 jours).

### Cas n°4 — Restaurer un utilisateur supprimé par erreur
**Contexte** : le helpdesk a supprimé `j.martin` au lieu de `j.martine`.
**Objectif** : restaurer avec groupes et attributs.
```powershell
Get-ADObject -Filter { SamAccountName -eq "j.martin" } -IncludeDeletedObjects | Restore-ADObject
Enable-ADAccount j.martin
Set-ADAccountPassword j.martin -Reset -NewPassword (ConvertTo-SecureString "Temporaire!2026" -AsPlainText -Force)
Set-ADUser j.martin -ChangePasswordAtLogon $true
Get-ADUser j.martin -Properties MemberOf | Select -ExpandProperty MemberOf
```
**Commentaire** : grâce à la corbeille, les appartenances aux groupes sont revenues. Sans corbeille, il aurait fallu tout réassigner à la main.

### Cas n°5 — Déployer un nouveau DC sur un site distant
**Contexte** : l'agence de Lyon grandit (60 personnes), le RODC ne suffit plus.
**Objectif** : promouvoir un DC inscriptible à Lyon.
**Étapes** : 1) `dcdiag /e` + `repadmin /replsummary` propres. 2) Serveur 2022 joint au domaine, DNS vers Paris. 3) `Install-ADDSDomainController` avec `-SiteName "Site-Lyon"` (ou déplacer après). 4) Vérifier réplication, DNS (`_msdcs`), GC. 5) Ajouter son IP dans le DHCP de Lyon. 6) Surveiller 48h.
**Commentaire** : utilisez `-InstallationMediaPath` (IFM) si la liaison est lente. Ne transférez aucun FSMO à Lyon (reste à Paris, section 36).

### Cas n°6 — Transférer les rôles FSMO avant maintenance du DC principal
**Contexte** : remplacement des disques du DC01 (détient les 5 FSMO), intervention de 4h.
**Objectif** : zéro impact.
```powershell
# Avant : santé
dcdiag /e | Select-String "passed|failed"
# Transfert vers DC02
Move-ADDirectoryServerOperationMasterRole -Identity "DC02" -OperationMasterRole 0,1,2,3,4 -Force
netdom query fsmo
# ... maintenance ...
# Retour (optionnel, pour garder la doc cohérente)
Move-ADDirectoryServerOperationMasterRole -Identity "DC01" -OperationMasterRole 0,1,2,3,4 -Force
```
**Commentaire** : le transfert est sans coupure (quelques secondes). Le PDC est le rôle dont l'absence se sent le plus : transférez-le en priorité si vous ne déplacez pas tout.

### Cas n°7 — Mettre en place une PSO pour les administrateurs
**Contexte** : audit interne : les admins ont la même politique mot de passe que les utilisateurs.
**Objectif** : 16 caractères min, 60 jours max, verrouillage à 3 essais pour les admins.
**Solution** : script de la section 46 (`PSO-Admins`, Precedence 10, appliqué à « Admins du domaine »). Vérifier avec `Get-ADUserResultantPasswordPolicy` sur un admin.
**Commentaire** : communiquez avant d'appliquer (les admins devront changer leur mot de passe au prochain cycle). Combinez avec du MFA.

### Cas n°8 — Diagnostiquer une réplication en échec entre deux sites
**Contexte** : `repadmin /replsummary` montre 12 échecs vers le site de Lyon, en augmentation.
**Objectif** : identifier la cause racine.
**Méthode** :
```powershell
repadmin /showrepl DC-Lyon          # détail des erreurs (ex. 1722 = RPC inaccessible)
dcdiag /test:dns /v                 # le DNS est coupable dans 50% des cas
Test-NetConnection DC-Lyon -Port 389
w32tm /query /status                # décalage d'heure ?
Get-WinEvent -LogName "Directory Service" | Where-Object Id -in 2108,1084 | Select -First 5
```
**Résolution typique** : VPN tombé, règle pare-feu, ou DNS. Après réparation : `repadmin /syncall /AdeP` puis surveillance.
**Commentaire** : notez l'heure de début des échecs et croisez avec les changements réseau — la cause est presque toujours un changement récent.

### Cas n°9 — Nettoyer un DC mort après crash disque
**Contexte** : DC03 (site distant) : disque mort, pas de sauvegarde récente, les 2 autres DC sont sains.
**Objectif** : nettoyer l'annuaire.
**Étapes** : 1) Vérifier qu'il ne détenait aucun FSMO (`netdom query fsmo`). 2) `Remove-ADDomainController -Identity "DC03" -Force` ou suppression via Sites et services. 3) Nettoyer le DNS (enregistrements A, NS, `_msdcs`). 4) `repadmin /replsummary`, `dcdiag /e`. 5) Reconstruire un serveur neuf et le promouvoir (jamais de « réparation » de l'ancien).
**Commentaire** : si DC03 détenait des FSMO, les saisir (section 35) **avant** le nettoyage.

### Cas n°10 — Déployer une imprimante par GPP avec ciblage par site
**Contexte** : chaque site a son imprimante ; les utilisateurs nomades doivent avoir la bonne par défaut.
**Objectif** : une seule GPO « Imprimantes », 3 éléments ciblés.
**Solution** : GPO `USER - Tous - Imprimantes` liée à la racine Utilisateurs ; 3 éléments GPP « Imprimante partagée » (`\\srv-print\Paris`, `\\srv-print\Lyon`...) chacun avec ciblage « le site AD est ... ». Cocher « Définir comme imprimante par défaut ».
**Commentaire** : le ciblage par site suit l'utilisateur quand il se déplace — mieux qu'une GPO par OU de site pour les nomades.

### Cas n°11 — Verrouiller les postes d'atelier partagés (bouclage)
**Contexte** : 10 PC d'atelier partagés par 40 opérateurs en 3x8 ; il faut un bureau identique et verrouillé.
**Objectif** : config utilisateur imposée par le poste.
**Solution** : OU `Atelier` avec les 10 PC ; GPO `ORDI - Atelier - Bouclage` activant le loopback en **Remplacement** ; GPO `USER - Atelier - Bureau verrouillé` (même OU) configurant le bureau : fond imposé, pas de panneau de config, verrouillage 5 min.
**Commentaire** : en mode Remplacement, les GPO utilisateur de l'opérateur sont ignorées sur ces PC — documentez-le pour éviter des tickets « mes lecteurs ont disparu à l'atelier ».

### Cas n°12 — Migrer les comptes de service vers des gMSA
**Contexte** : 8 applications utilisent `svc_appli` (mot de passe qui n'expire jamais, connu de 3 personnes).
**Objectif** : passer au gMSA, un par application.
**Étapes** (par application) : 1) `Add-KdsRootKey` (une fois). 2) Créer le gMSA, autoriser les serveurs. 3) `Install-ADServiceAccount` sur chaque serveur. 4) Changer le compte du service (services.msc → `ad\gmsa_appli$`, mot de passe vide). 5) Redémarrer le service, vérifier les logs. 6) Désactiver l'ancien `svc_appli` après 1 mois de validation.
**Commentaire** : faites-le application par application, jamais en big-bang. Certaines vieilles applis ne supportent pas les gMSA : PSO dédié en repli.

### Cas n°13 — Auditer et réduire les membres des groupes sensibles
**Contexte** : préparation d'un audit de sécurité.
**Objectif** : liste justifiée des admins.
```powershell
"Admins du domaine","Admins de l’entreprise","Opérateurs de sauvegarde" | ForEach-Object {
  Write-Host "== $_ ==" -ForegroundColor Cyan
  Get-ADGroupMember -Identity $_ -Recursive | Format-Table Name, objectClass
}
```
**Suite** : pour chaque membre, justification écrite (fonction, besoin). Retirer les comptes de service et les anciens prestataires. Mettre en place la revue mensuelle (section 72).
**Commentaire** : le groupe « Admins du domaine » doit tenir sur une demi-page. Au-delà, c'est un symptôme.

