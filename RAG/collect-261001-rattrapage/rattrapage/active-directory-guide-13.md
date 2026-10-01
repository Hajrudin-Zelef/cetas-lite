---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-13
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1989, 2155]
sha256: a708dcf38b833b2c1c24a953d13a4255b112c63037e8ba8dcb28629463ee2efd
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

1. Créer avec un **commentaire** (objectif, demandeur, date).
2. Tester sur une **OU pilote** (quelques postes).
3. **Sauvegarder** (`Backup-GPO`) avant chaque modification majeure.
4. Déployer progressivement.
5. **Désactiver** (pas supprimer) les GPO obsolètes pendant 3 mois, puis supprimer.
6. Revue annuelle : GPO non liées, vides, ou en doublon → nettoyage.

```powershell
# Trouver les GPO non liées (candidates au nettoyage)
Get-GPO -All | Where-Object {
  ($_ | Get-GPOReport -ReportType Xml) -notmatch "LinksTo"
} | Select-Object DisplayName

# Sauvegarder une GPO avant modification
Backup-GPO -Name "ORDI - Tous - Verrouillage session 10 min" `
  -Path "E:\Sauvegardes\GPO" -Comment "Avant passage à 5 min - $(Get-Date -Format 'yyyy-MM-dd')"
```

---

## 71. Santé du domaine : checklist quotidienne

À automatiser (script section 73) + coup d'œil humain de 5 minutes :

- [ ] `repadmin /replsummary` : 0 échec sur tous les DC.
- [ ] `dcdiag` rapide sur chaque DC (ou au moins les FSMO) : tout « passed ».
- [ ] Espace disque des DC (> 20 % libre sur le volume NTDS/SYSVOL).
- [ ] Verrouillages de comptes anormaux (pic = attaque ou service avec vieux mot de passe).
- [ ] Sauvegarde d'état système de la nuit : **succès** vérifié.
- [ ] Événements critiques (ID 2108, 1084, 1388, 16650, 4740 en masse).
- [ ] Le PDC est joignable et l'heure est saine (`w32tm /query /status`).

---

## 72. Santé du domaine : checklist hebdomadaire et mensuelle

**Hebdomadaire** :

- [ ] Revue des comptes créés/désactivés/supprimés de la semaine.
- [ ] Comptes verrouillés récurrents (même utilisateur 3 semaines d'affilée = problème).
- [ ] GPO modifiées (qui, quand) : `Get-GPO -All | Sort-Object ModificationTime -Descending`.
- [ ] Ordinateurs inactifs > 60 jours : à désactiver/supprimer après validation.
- [ ] Sauvegardes GPO à jour.
- [ ] Correctifs Windows sur les DC (fenêtre de maintenance planifiée, **un DC à la fois**, jamais tous ensemble).

**Mensuelle** :

- [ ] Revue des appartenances aux groupes sensibles (Admins du domaine, Opérateurs de sauvegarde...).
- [ ] Comptes avec mot de passe qui n'expire jamais : justification à jour ?
- [ ] Délégations sur les OU : toujours valides ?
- [ ] Test de restauration (labo isolé) au moins 2x/an.
- [ ] Documentation à jour (topologie, FSMO, mots de passe DSRM au coffre).
- [ ] Nettoyage : comptes/ordinateurs obsolètes, GPO orphelines.

```powershell
# Membres des groupes sensibles (revue mensuelle)
$groupes = "Admins du domaine","Admins de l’entreprise","Opérateurs de sauvegarde","Admins du schéma"
foreach ($g in $groupes) {
  Write-Host "== $g ==" -ForegroundColor Cyan
  Get-ADGroupMember -Identity $g -Recursive | Select-Object Name, objectClass
}
```

---

## 73. Scripts de rapport automatisés

**Rapport de santé quotidien** (à planifier via une tâche sur un serveur d'admin, avec un gMSA) :

```powershell
# Sante-AD-Quotidien.ps1
$rapport = "C:\Rapports\AD\Sante-$(Get-Date -Format 'yyyy-MM-dd').txt"
$dc = Get-ADDomainController -Filter * | Select-Object -ExpandProperty HostName

"=== REPLICATION ===" | Out-File $rapport
repadmin /replsummary | Out-File $rapport -Append

"`n=== DCDIAG (erreurs) ===" | Out-File $rapport -Append
foreach ($s in $dc) {
  "--- $s ---" | Out-File $rapport -Append
  dcdiag /s:$s | Select-String "failed|error" | Out-File $rapport -Append
}

"`n=== FSMO ===" | Out-File $rapport -Append
netdom query fsmo | Out-File $rapport -Append

"`n=== COMPTES VERROUILLES ===" | Out-File $rapport -Append
Search-ADAccount -LockedOut | Select-Object Name, SamAccountName | Out-File $rapport -Append

"`n=== DISQUE DC ===" | Out-File $rapport -Append
foreach ($s in $dc) {
  Get-WmiObject Win32_LogicalDisk -ComputerName $s -Filter "DriveType=3" |
    Select-Object @{N="Serveur";E={$s}}, DeviceID,
      @{N="Libre %";E={[math]::Round($_.FreeSpace/$_.Size*100,1)}} |
    Out-File $rapport -Append
}
# Envoi par mail (serveur SMTP interne)
Send-MailMessage -To "exploitation@entreprise.fr" -From "ad-rapport@entreprise.fr" `
  -Subject "[AD] Rapport santé $(Get-Date -Format 'yyyy-MM-dd')" `
  -Body (Get-Content $rapport | Out-String) -SmtpServer "smtp.entreprise.fr"
```

**Rapport mouvements de comptes (hebdo)** : interrogez le journal de sécurité du PDC (ID 4720 création, 4726 suppression, 4722 activation, 4725 désactivation) ou activez l'audit et centralisez (SIEM/Wazuh — voir votre guide Wazuh).

---

## 74. Supervision : compteurs et alertes critiques

| Source | ID | Gravité | Action |
|---|---|---|---|
| NTDS Replication | 2108, 1084 | 🔴 | Réplication en échec durable |
| NTDS Replication | 1388 | 🟠 | Lingering object potentiel |
| Directory Service | 16650 | 🔴 | Pool RID épuisé |
| Directory Service | 2162 | 🔴 | USN rollback détecté (VM-GenerationID) |
| Sécurité (PDC) | 4740 | 🟠 | Verrouillage de compte (pic = alerte) |
| Sécurité | 4728/4732/4756 | 🟠 | Ajout à un groupe sensible |
| Sécurité | 4720/4726 | 🟢 | Création/suppression de compte (revue) |
| DNS Server | 4015 | 🟠 | Problème DNS sur DC |
| Microsoft-Windows-Backup | 14, 8 | 🔴 | Échec sauvegarde |

**Seuils** : espace disque < 20 % = alerte, < 10 % = critique ; réplication > 60 min sans succès = alerte ; un DC injoignable > 15 min = critique (sauf maintenance déclarée).

---

## 75. Pense-bête de poche : commandes essentielles

```powershell
# --- Info ---
netdom query fsmo
Get-ADDomainController -Filter * | Select Name,Site,IsGlobalCatalog
repadmin /replsummary
dcdiag /e
nltest /dsgetsite          # site du poste
nltest /dclist:domaine     # liste des DC

# --- Utilisateurs / Groupes ---
New-ADUser / Get-ADUser / Set-ADUser / Disable-ADAccount / Unlock-ADAccount
New-ADGroup / Add-ADGroupMember / Get-ADGroupMember
Search-ADAccount -LockedOut | Unlock-ADAccount
Search-ADAccount -AccountInactive -TimeSpan 90 -UsersOnly

# --- GPO ---
Get-GPO -All
Backup-GPO -All -Path E:\Sauvegardes\GPO
gpresult /r ; gpresult /h rapport.html
gpupdate /force
Invoke-GPUpdate -Computer PC-01 -Force

# --- Réplication / Dépannage ---
repadmin /showrepl
repadmin /syncall DC01 /AdeP
repadmin /kcc
dcdiag /test:dns /v

# --- FSMO ---
Move-ADDirectoryServerOperationMasterRole -Identity DC02 -OperationMasterRole 0,1,2,3,4 -Force

# --- Divers ---
w32tm /query /status
redircmp "OU=Postes,DC=ad,DC=entreprise,DC=fr"
Enable-ADOptionalFeature -Identity "Recycle Bin Feature" -Scope ForestOrConfigurationSet -Target "ad.entreprise.fr"
```

---

## 76. Glossaire

