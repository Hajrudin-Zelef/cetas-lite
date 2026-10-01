---
id: collect-261001-rattrapage/rattrapage/win11-guide-19
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-10-14", "2026-09-27"]
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [3139, 3305]
sha256: aa0646814523b036b982ab7129ec54c0daa7ffd5ad93e3803313d80ba384c12b
---

# Windows 11 en entreprise — Guide technique ultra-complet

| Cause | Indice | Solution |
|---|---|---|
| SMBv1 désactivé côté client, serveur ancien | Dialect vide / erreur | Mettre à jour le serveur (pas réactiver SMBv1 !) |
| Signature SMB obligatoire d'un côté | Négociation échouée | Aligner les GPO des deux côtés |
| Pare-feu (client ou serveur) | Port 445 filtré | Règle pare-feu (section 35) |
| Point and Print restreint (imprimantes) | Pilote refusé | GPO Point and Print (section 66) |
| Horloge désynchronisée | Kerberos échoue | w32tm /resync |

---

## 58. Fin de vie de Windows 10 (14/10/2025) : stratégie de migration

> 📌 Au 27/09/2026, Windows 10 est **hors support depuis près d'un an**. Les postes Windows 10 restants ne reçoivent plus de correctifs de sécurité (hors programme ESU payant). Chaque poste Windows 10 du parc est désormais un **risque documenté**.

### 58.1 Les 3 options pour les postes Windows 10 restants

| Option | Quand | Coût |
|---|---|---|
| **Migrer vers Windows 11** | Poste compatible (TPM 2.0, CPU listé) | Temps IT + validation |
| **Remplacer le poste** | Poste incompatible (< 8e gén, pas de TPM 2.0) | Achat |
| **ESU (Extended Security Updates)** | Temporisation < 12 mois, contrainte métier bloquante | Payant / poste / an, **prix croissant** |

> ⛔ L'ESU n'est **pas** une stratégie : c'est un pansement payant pour gagner du temps. Tout poste sous ESU doit avoir une date de sortie écrite.

### 58.2 Plan de migration en 5 phases

```
Phase 1 — INVENTAIRE (2-4 semaines)
  → Script de compatibilité (section 3.2) sur 100 % du parc
  → Classer : Compatible / Compatible après action (TPM, MBR2GPT) / Incompatible
  → Lier avec GLPI (section 59)

Phase 2 — PILOTE (4-6 semaines)
  → 5-10 % du parc, volontaires + IT
  → Valider : image/MDT ou Autopilot, GPO, applications métier, pilotes, BitLocker
  → Documenter les problèmes et leurs solutions

Phase 3 — VAGUES (8-16 semaines)
  → Par service / site, 15-25 % du parc par vague
  → Communication : 2 semaines avant, rappel J-3, support renforcé J+3
  → Fenêtre de rollback de 10 jours par poste (ne pas forcer /ResetBase trop tôt)

Phase 4 — RÉSIDUEL (4 semaines)
  → Postes problématiques : remplacement ou ESU avec date de sortie
  → Zéro poste Windows 10 "oublié"

Phase 5 — CLÔTURE
  → Rapport : 100 % migré ou plan daté pour chaque exception
  → Archiver les images/MDT Windows 10, couper les GPO spécifiques
```

### 58.3 Communication type (à adapter)

> « Votre poste sera migré vers Windows 11 le [date]. Sauvegardez vos fichiers sur OneDrive (déjà actif). La migration dure ~1 h, vos fichiers et applications sont conservés. En cas de problème : [numéro du support]. »

---

## 59. Inventaire du parc : PowerShell, GLPI, Intune, rapports

### 59.1 Inventaire PowerShell (sans agent)

```powershell
<#
.SYNOPSIS
    Inventaire complet d'un poste : OS, matériel, logiciels, BitLocker, MàJ.
#>
$inv = [ordered]@{}
$os  = Get-CimInstance Win32_OperatingSystem
$cs  = Get-CimInstance Win32_ComputerSystem
$inv.Hostname     = $env:COMPUTERNAME
$inv.OS           = $os.Caption
$inv.Build        = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").CurrentBuild
$inv.UBR          = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").UBR
$inv.InstallDate  = $os.InstallDate
$inv.Manufacturer = $cs.Manufacturer
$inv.Model        = $cs.Model.Trim()
$inv.Serial       = (Get-CimInstance Win32_BIOS).SerialNumber.Trim()
$inv.CPU          = (Get-CimInstance Win32_Processor).Name.Trim()
$inv.RAM_GB       = [math]::Round($cs.TotalPhysicalMemory / 1GB, 1)
$inv.DiskFree_GB  = [math]::Round((Get-CimInstance Win32_LogicalDisk -Filter "DeviceID='C:'").FreeSpace / 1GB, 1)
$inv.TPM2         = ((Get-CimInstance -Namespace root/cimv2/security/microsofttpm -ClassName Win32_Tpm -ErrorAction SilentlyContinue).SpecVersion -split ',')[0]
$inv.SecureBoot   = try { Confirm-SecureBootUEFI } catch { 'N/A' }
$inv.BitLocker    = (Get-BitLockerVolume -MountPoint "C:" -ErrorAction SilentlyContinue).ProtectionStatus
$inv.Domain       = $cs.Domain
$inv.LastBoot     = $os.LastBootUpTime

# Logiciels installés (registre Uninstall — rapide, sans WMI lent)
$inv.AppsCount = (Get-ChildItem "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall" |
                  Where-Object { (Get-ItemProperty $_.PSPath -ErrorAction SilentlyContinue).DisplayName }).Count

[pscustomobject]$inv | Export-Csv "\\SRV-FICHIERS\Inventaire$\INV_$($env:COMPUTERNAME).csv" -NoTypeInformation -Encoding UTF8
```

### 59.2 Lier avec GLPI

GLPI + **agent GLPI** (ou FusionInventory) sur chaque poste = inventaire automatique (matériel, logiciels, licences) remonté au serveur GLPI.

Points d'intégration Windows 11 :

- Déployer l'agent via **GPO / MDT / Intune** (MSI silencieux) :

```powershell
# Installation silencieuse de l'agent GLPI (exemple)
msiexec /i glpi-agent-1.x-x64.msi /quiet SERVER=https://glpi.entreprise.local `
    RUNNOW=1 INSTALLTYPE=from-scratch
```

- Planifier l'inventaire (tâche planifiée quotidienne ou au démarrage).
- Croiser avec le script de compatibilité (section 3.2) : champ personnalisé « Compatible Win11 » dans GLPI.
- Utiliser les **rapports GLPI** pour piloter les vagues de migration (section 58.2).

### 59.3 Intune comme source d'inventaire (parc cloud)

Intune → Appareils → exporter (CSV) : modèle, OS, version, conformité, chiffrement, dernier check-in. Compléter avec des **requêtes de conformité** personnalisées (ex. : build minimale 26100).

---

## 60. Supervision et journaux : Event Viewer, journaux clés, collecte

### 60.1 Collecte centralisée : les options

| Solution | Coût | Remarque |
|---|---|---|
| Abonnement aux journaux (Event Forwarding natif) | Gratuit | WEF : collecteur + GPO d'abonnement, parfait pour démarrer |
| Zabbix / agent | Gratuit | Surveillance + inventaire, pas d'analyse de logs poussée |
| Wazuh | Gratuit | SIEM léger, excellent pour les logs Windows |
| Microsoft Sentinel | Payant (ingestion) | SIEM cloud, intégration Entra/Defender native |

### 60.2 Windows Event Forwarding (WEF) — mise en place express

```powershell
# Sur le COLLECTEUR :
# 1. Configurer le service :
winrm quickconfig -q
# 2. Créer l'abonnement (console eventvwr → Abonnements → Créer) :
#    - Source : ordinateurs du domaine (via GPO)
#    - Journaux : System, Application, Security (filtré), BitLocker-Operational

# Sur les POSTES (via GPO) :
# Configuration ordinateur\Modèles d'administration\Composants Windows\Transfert d'événements
# → "Configurer l'adresse du serveur cible" : FQDN du collecteur
# + winrm : Set-Item WSMan:\localhost\Client\TrustedHosts -Value "collecteur.entreprise.local"
```

### 60.3 Événements à surveiller en priorité (alertes)

| ID | Journal | Signification |
|---|---|---|
| 4625 | Security | Échec d'ouverture de session (bruteforce ?) |
| 4672 | Security | Privilèges spéciaux attribués |
| 4720/4728 | Security | Création de compte / ajout à un groupe sensible |
| 1001 | System | BugCheck (écran bleu) |
| 6008 | System | Arrêt inattendu |
| 1116 | Defender/Operational | Malware détecté |
| 1014 | System (DNS Client) | Échec de résolution (réseau/DNS) |
| 8198/8200 | BitLocker-Operational | Échec de chiffrement |

---

(Bloc 4/6 — sections 46 à 60)

---

## 61. Scripts utiles : boîte à outils PowerShell de l'admin

### 61.1 Audit express d'un poste (à lancer en admin)

