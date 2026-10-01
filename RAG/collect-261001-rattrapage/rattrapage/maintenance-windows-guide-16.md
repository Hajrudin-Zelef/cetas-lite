---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-16
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [2934, 3132]
sha256: 0c1b98bee87e5947c6c8e55a5f2f9b7bae175da418f130245a795e8a3a4ec27f
---

# Maintenance et exploitation Windows en entreprise

```cmd
:: Quel DC sert ce poste ?
nltest /dsgetdc:contoso.local

:: État du canal sécurisé avec le domaine
nltest /sc_query:contoso.local
:: Doit retourner : "Trusted DC ... The command completed successfully"
:: Drapeaux attendus : 0x... (dont 0x200 = secure channel OK selon version)

:: Lister les DC du domaine
nltest /dclist:contoso.local

:: Forcer la découverte d'un DC (après correction DNS/réseau)
nltest /dsgetdc:contoso.local /force
```

Si `/dsgetdc` échoue : le poste ne trouve pas de DC → DNS (les SRV
`_ldap._tcp.contoso.local`, §81), réseau (ports 389/636/88/445 vers le DC, §79),
ou heure (§104). **Toujours** dans cet ordre.

---

## 103. Test-ComputerSecureChannel et réinitialisation

Le **canal sécurisé** (mot de passe du compte machine) peut se désynchroniser
(restauration d'une VM, double compte machine, clone mal syspreppé). Symptôme :
« La relation d'approbation entre cette station et le domaine a échoué ».

```powershell
# Tester le canal sécurisé
Test-ComputerSecureChannel
# False -> tenter la réparation (nécessite un compte autorisé)
Test-ComputerSecureChannel -Repair -Credential (Get-Credential)
# True -> canal réparé sans quitter/réintégrer le domaine
```

Si la réparation échoue : **quitter puis réintégrer** le domaine (méthode
classique, redémarrage requis entre les deux) :

```powershell
# Quitter (compte admin local requis après le reboot)
Remove-Computer -UnjoinDomainCredential (Get-Credential) -Restart -Force
# Puis réintégrer
Add-Computer -DomainName 'contoso.local' -Credential (Get-Credential) -Restart -Force
```

> Avant de réintégrer : supprimez le compte machine obsolète dans l'AD
> (utilisateurs et ordinateurs AD) pour éviter les doublons, et vérifiez qu'il
> n'existe pas de clone avec le même SID (sysprep !).

---

## 104. Temps : w32tm, la cause n°1 des échecs Kerberos

Kerberos tolère **5 minutes** d'écart (par défaut). Au-delà : échecs d'ouverture
de session, « relation d'approbation », GPO qui ne s'appliquent plus.

```cmd
:: État de la synchro : source, écart, dernière synchro réussie
w32tm /query /status
:: -> "Source : DC-contrôleur.contoso.local" attendu sur un poste du domaine
::    "Stratum", "Last Successful Sync Time" : vérifier la fraîcheur

:: Pairs configurés
w32tm /query /peers

:: Forcer une resynchronisation
w32tm /resync /rediscover
```

**Hiérarchie** : les postes se synchronisent sur le DC qui les authentifie, les
DC sur le PDC de la forêt, le PDC sur une source externe fiable (NTP du
fournisseur, `pool.ntp.org`, horloge interne). **Ne pointez jamais les postes
directement sur Internet** : tout passe par la hiérarchie AD.

```cmd
:: Sur le PDC : déclarer une source externe fiable
w32tm /config /manualpeerlist:"0.fr.pool.ntp.org 1.fr.pool.ntp.org" /syncfromflags:manual /reliable:yes /update
net stop w32time && net start w32time
```

**Cas VM** : désactivez la synchronisation d'horloge de l'hyperviseur **pour les
DC** (sinon l'hôte écrase w32tm) ; laissez-la pour les postes simples. Un DC
virtualisé qui dérive = tout le domaine qui dérive.

---

## 105. Jonction au domaine : procédure et dépannage

Checklist **avant** de joindre :

- [ ] DNS du poste = DC/DNS interne (§82), suffixe DNS correct.
- [ ] Heure à moins de 5 min du DC (§104).
- [ ] Le compte machine n'existe pas déjà (ou est réinitialisé).
- [ ] Compte autorisé à joindre (droit « Ajouter des stations de travail », 10
      par défaut pour les utilisateurs standard).
- [ ] Ports vers le DC ouverts (DNS 53, Kerberos 88, LDAP 389, SMB 445).

```powershell
# Jonction scriptée (déploiement)
Add-Computer -DomainName 'contoso.local' -OUPath 'OU=Postes,OU=Paris,DC=contoso,DC=local' `
  -Credential (Get-Credential) -Restart -Force
```

Échecs fréquents : « le domaine spécifié n'existe pas » = DNS ; « accès refusé »
= droits ou compte machine existant protégé ; erreur réseau = pare-feu/ports.
**Après** la jonction : vérifiez `nltest /dsgetdc`, `gpresult /r` (les GPO
s'appliquent au prochain redémarrage/ouverture de session), et déplacez le compte
dans la bonne OU (**jamais** laisser les postes dans `Computers` en production).

---

## 106. Monitoring : ce qu'on supervise sur Windows

Le socle minimal, par machine :

| Famille | Quoi | Source |
|---|---|---|
| **Disponibilité** | Ping, agent joignable | Superviseur |
| **Système** | CPU, mémoire, espace disque C: et volumes | Compteurs / agent |
| **Services** | Services critiques démarrés (DNS, DHCP, AD DS, Spooler…) | Agent |
| **Journaux** | Erreurs Système/Sécurité (patterns §37-38) | WEF / agent / SIEM |
| **Mises à jour** | Conformité patch, redémarrage en attente | Script / Intune |
| **Matériel** | SMART, RAID, température (via constructeur) | Agent + outils OEM |
| **Sauvegarde** | Succès du dernier job | Rapport du logiciel |
| **Certificats** | Expiration < 60 j | Script §120 |
| **AD (DC)** | Réplication, SYSVOL, temps, rôles FSMO | Scripts + `repadmin` |

**Règle** : chaque alerte doit avoir un **propriétaire** et une **procédure**
(lien vers la section du guide). Une alerte sans procédure = du bruit qu'on
finit par ignorer.

Redémarrage en attente (indicateur simple et précieux) :

```powershell
# Le serveur attend-il un redémarrage ? (patch, installation...)
$pending = @(
    'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Component Based Servicing\RebootPending',
    'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\WindowsUpdate\Auto Update\RebootRequired',
    'HKLM:\SYSTEM\CurrentControlSet\Control\Session Manager\PendingFileRenameOperations'
)
$pending | ForEach-Object {
    [pscustomobject]@{ Cle = $_; EnAttente = (Test-Path $_) }
} | Format-Table -AutoSize
```

---

## 107. Seuils d'alerte recommandés

| Métrique | Warning | Critique | Remarque |
|---|---|---|---|
| Espace libre volume système | < 20 % | < 10 % | §31 |
| CPU (15 min) | > 80 % | > 95 % | + file d'attente §55 |
| Mémoire libre | < 15 % | < 5 % | §56 |
| Latence disque lecture | > 20 ms | > 50 ms | HDD ; diviser par ~4 sur SSD §57 |
| Ping | perte > 5 % | hôte injoignable | — |
| Service critique arrêté | — | immédiat | Selon rôle |
| Erreurs journal Système (nouveau pattern) | — | immédiat | Via WEF/SIEM §36 |
| Échecs logon 4625 (pic) | ×5 vs baseline | ×20 vs baseline | §38 |
| Réplication AD en erreur | — | immédiat | §39 |
| Certificat expire dans | < 60 j | < 14 j | §120 |
| Sauvegarde non réussie | > 24 h | > 48 h | §123 |
| Redémarrage en attente | > 7 j | > 30 j | §106 |

**Évitez l'alert fatigue** : un seuil qui déclenche chaque semaine sans action =
un mauvais seuil. Revisitez les seuils trimestriellement avec l'équipe.

---

## 108. Zabbix : agent2 et templates Windows

**Zabbix agent 2** (recommandé, avec plugins) se déploie par GPO/script. Points
d'attention : autoriser le serveur Zabbix dans le pare-feu (port 10050), TLS
PSK si les flux traversent des zones non maîtrisées.

Templates à appliquer :

- **Windows by Zabbix agent** : CPU, mémoire, disques, interfaces, services,
  `system.localtime` (dérive d'horloge !), uptime.
- Complétez avec des **items personnalisés** (`UserParameter` ou clés
  `system.run[]` — à restreindre via `AllowKey`/`DenyKey` pour la sécurité) :

```ini
# zabbix_agent2.conf - exemples d'items maison
# Espace libre C: en % (déjà dans le template, exemple pédagogique)
# Redémarrage en attente (0/1) via script PowerShell
AllowKey=system.run[*]
```

```powershell
# C:\Scripts\zabbix\pending-reboot.ps1 -> retourne 1 si reboot en attente
$paths = @(
  'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Component Based Servicing\RebootPending',
  'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\WindowsUpdate\Auto Update\RebootRequired'
)
if ($paths | Where-Object { Test-Path $_ }) { 1 } else { 0 }
```

