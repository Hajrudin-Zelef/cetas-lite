---
id: collect-261001-rattrapage/rattrapage/win11-guide-23
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "copilot"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [3861, 4018]
sha256: 2b291ab3b6669934121bf5f9f7c005296ce4405deb489a4071befe597d526b1c
---

# Windows 11 en entreprise — Guide technique ultra-complet

### Erreur n°16 — Laisser le compte Administrator intégré activé avec un mot de passe faible
Cible n°1 des attaques par force brute (RID 500, impossible à renommer efficacement). Le **désactiver**, gérer l'admin local via LAPS.

### Erreur n°17 — Configurer l'accès conditionnel sans compte de secours
Une stratégie trop stricte vous enferme dehors vous aussi. **Compte break glass** exclu de toutes les stratégies, avec authentification forte, dont les accès sont alertés (section 64.2).

### Erreur n°18 — Croire que « réinitialiser ce PC » suffit pour un départ
Sans Autopilot Reset / effacement Intune, l'ancien profil, les certificats et parfois les données restent accessibles. Procédure de réaffectation écrite (sections 13, 71).

### Erreur n°19 — Ne pas surveiller l'espace disque des postes
C: plein = MàJ impossibles, profils corrompus, tickets en cascade. Storage Sense par GPO + alerte Zabbix/Intune sous 15 % libre.

### Erreur n°20 — Faire du poste un « animal de compagnie » au lieu du « bétail »
Passer 3 heures à « réparer » un poste quand le redéploiement prend 45 minutes. **Doctrine** : poste jetable, données dans OneDrive, redéploiement standardisé (sections 40.3, 71).

---

## 74. Checklist : déploiement d'un poste de A à Z

### Phase préparation
- [ ] Poste compatible vérifié (TPM 2.0, Secure Boot, CPU, 256 Go SSD min.)
- [ ] Méthode de déploiement choisie : MDT / Autopilot / manuelle
- [ ] Média à jour (ISO du mois, 24H2)
- [ ] Pilotes du modèle validés (dossier dédié MDT ou package constructeur)
- [ ] Compte de jointure / profil Autopilot prêt
- [ ] Nom de poste selon la convention (`SITE-FONCTION-###`)

### Phase installation
- [ ] Installation / task sequence / Autopilot exécutée sans erreur
- [ ] Mises à jour Windows complètes (plus de MàJ en attente)
- [ ] Pilotes tous installés (aucun point d'exclamation dans le Gestionnaire de périphériques)
- [ ] Fuseau horaire, langue, clavier : fr-FR

### Phase configuration
- [ ] Jointure domaine / Entra vérifiée (`dsregcmd /status`, `Get-CimInstance Win32_ComputerSystem`)
- [ ] GPO appliquées (`gpresult /r` : toutes les GPO attendues présentes)
- [ ] BitLocker actif + clé en AD/Entra vérifiée
- [ ] LAPS fonctionnel (mot de passe récupérable)
- [ ] Defender à jour, protection temps réel active
- [ ] Applications du socle installées (M365, navigateur géré, VPN, agent supervision, agent GLPI)
- [ ] Wi-Fi d'entreprise / VPN testé
- [ ] Imprimantes déployées et testées

### Phase remise
- [ ] Session utilisateur testée (ouverture < 1 min, profil propre)
- [ ] OneDrive KFM actif et synchronisé
- [ ] Fiche de remise signée (utilisateur informé du support)
- [ ] Ticket GLPI de déploiement clôturé avec n° de série et date

---

## 75. Checklist : audit de sécurité d'un poste Windows 11

- [ ] Build ≥ 26100 (24H2) et dernière cumulative installée
- [ ] TPM 2.0 prêt (`Get-Tpm`), Secure Boot activé (`Confirm-SecureBootUEFI`)
- [ ] BitLocker : `ProtectionStatus = On`, clé en AD/Entra
- [ ] HVCI / Intégrité de la mémoire : activé
- [ ] Credential Guard : actif (si Entreprise)
- [ ] LSA Protection : activée
- [ ] Defender : temps réel OK, signatures < 7 jours, Tamper Protection active
- [ ] Règles ASR : en mode Block (post-audit)
- [ ] Pare-feu : profils actifs, règles locales fusionnées ou non selon politique
- [ ] LAPS : mot de passe géré et récupérable
- [ ] Compte Administrator intégré : désactivé
- [ ] Utilisateur quotidien : **non administrateur**
- [ ] SMBv1 : désactivé ; LLMNR : désactivé
- [ ] PowerShell v2 : désactivé
- [ ] UAC : niveau maximum
- [ ] Mises à jour : dans l'anneau attendu, pas de pause anormale
- [ ] Télémétrie : niveau Sécurité (Entreprise)
- [ ] Copilot / Widgets : désactivés (si politique)
- [ ] Store : selon politique (bloqué ou catalogue approuvé)
- [ ] Exclusions Defender : documentées et justifiées
- [ ] Dernière analyse complète < 30 jours

```powershell
# Score rapide : compte les points validés (extrait — à compléter selon votre politique)
$score = 0; $total = 0
function Test-Point($nom, $condition) {
    $script:total++
    if (& $condition) { $script:score++; Write-Host "[OK] $nom" -ForegroundColor Green }
    else { Write-Host "[KO] $nom" -ForegroundColor Red }
}
Test-Point "TPM 2.0 prêt" { (Get-Tpm).TpmReady }
Test-Point "Secure Boot" { try { Confirm-SecureBootUEFI } catch { $false } }
Test-Point "BitLocker actif" { (Get-BitLockerVolume -MountPoint "C:").ProtectionStatus -eq 'On' }
Test-Point "Defender temps réel" { (Get-MpComputerStatus).RealTimeProtectionEnabled }
Test-Point "Admin intégré désactivé" { (Get-LocalUser -Name "Administrateur").Enabled -eq $false }
Write-Host "`nScore : $score / $total" -ForegroundColor Cyan
```

---

## 76. Pense-bête de poche : commandes et raccourcis essentiels

### PowerShell — diagnostic express

```powershell
Get-ComputerInfo | Select-Object WindowsProductName, WindowsVersion, OsBuildNumber
Get-Tpm | Select-Object TpmPresent, TpmReady, TpmEnabled
Confirm-SecureBootUEFI
Get-BitLockerVolume | Select-Object MountPoint, ProtectionStatus, EncryptionPercentage
Get-MpComputerStatus | Select-Object AntivirusEnabled, RealTimeProtectionEnabled, AntivirusSignatureLastUpdated
gpresult /r
dsregcmd /status
systeminfo | Select-String "Version du système","Type du système"
```

### Réseau

```powershell
Test-NetConnection srv.entreprise.local -Port 445
Resolve-DnsName intranet.entreprise.local
Get-NetIPAddress -AddressFamily IPv4
netsh wlan show interfaces
ipconfig /flushdns
netsh winsock reset        # + redémarrage
```

### Dépannage système

```powershell
sfc /scannow
dism /Online /Cleanup-Image /RestoreHealth
chkdsk C: /f
shutdown /r /o /f /t 0     # redémarrer vers WinRE
reagentc /info
wusa /uninstall /kb:XXXXXXX /quiet /norestart
```

### AD / GPO / postes distants

```powershell
gpupdate /force
gpresult /h C:\Admin\gpresult.html
Reset-ComputerMachinePassword -Server DC01 -Credential (Get-Credential)
Enter-PSSession -ComputerName PC-042
Invoke-Command -ComputerName PC-042 -ScriptBlock { Get-BitLockerVolume }
```

### Raccourcis clavier (support)

| Touches | Action |
|---|---|
| `Win + X` | Menu admin (Terminal, Gestion du disque…) |
| `Win + R` → `msconfig` | Configuration système |
| `Win + R` → `gpedit.msc` | Stratégie de groupe locale |
| `Win + R` → `eventvwr` | Observateur d'événements |
| `Win + R` → `tpm.msc` | Console TPM |
| `Win + R` → `secpol.msc` | Stratégie de sécurité locale |
| `Maj + Redémarrer` | Démarrage avancé (WinRE) |
| `Ctrl + Maj + F3` (OOBE) | Mode audit |
| `Win + Ctrl + Maj + B` | Redémarrer le pilote graphique (écran figé) |

---

## 77. Glossaire

