---
id: collect-261001-rattrapage/rattrapage/win11-guide-13
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [1928, 2108]
sha256: f9306d204efd94cf1019284daca29a047ed03c9ee07bb33c3769812c70826d77
---

# Windows 11 en entreprise — Guide technique ultra-complet

```powershell
# Exemple : bloquer le contenu exécutable d'un client de messagerie (règle audit puis block)
# GUIDs des règles ASR : voir documentation Microsoft (liste officielle)
$asrRules = @(
    "BE9BA2D9-53EA-4CDC-84E5-9B1EEEE46550",  # Bloquer exécutable depuis client mail
    "D4F940AB-401B-4EFC-AADC-ADB37D02A12E",  # Bloquer exécutable depuis Adobe Reader
    "3B576869-A4EC-4529-8536-B80A7769E899"   # Bloquer Office créant du contenu exécutable
)
foreach ($guid in $asrRules) {
    Add-MpPreference -AttackSurfaceReductionRules_Ids $guid -AttackSurfaceReductionRules_Actions Enabled
}
# Actions : Disabled=0, Block=1, Audit=2, Warn=6
# TOUJOURS commencer en Audit (2), analyser, puis passer en Block (1).
```

### 30.4 Journaux Defender utiles

| Journal | Emplacement |
|---|---|
| Opérationnel | `Applications and Services Logs\Microsoft\Windows\Windows Defender\Operational` |
| Détections | Même journal, ID 1116 (malware détecté), 1117 (action) |
| Ligne de commande | `Get-MpThreat`, `Get-MpThreatDetection` |

```powershell
# Historique des menaces
Get-MpThreat | Select-Object ThreatName, SeverityID, InitialDetectionTime, Resources
```

---

(Bloc 2/6 — sections 16 à 30)

---

## 31. Sécurité : SmartScreen, protection réseau, contrôle des applications

### 31.1 Microsoft Defender SmartScreen

SmartScreen filtre les URL et les fichiers (réputation) dans Edge et l'Explorateur.

| Réglage | GPO |
|---|---|
| Activer SmartScreen pour Edge | `Microsoft Edge\SmartScreen` → Activé |
| Activer pour l'Explorateur | `Composants Windows\Explorateur Windows` → « Configurer Windows Defender SmartScreen » |
| Empêcher le contournement | « Empêcher le contournement des avertissements SmartScreen » |

```powershell
# État via registre
Get-ItemProperty "HKLM:\SOFTWARE\Policies\Microsoft\Windows\System" -Name "EnableSmartScreen" -ErrorAction SilentlyContinue
# 1 = activé (stratégie), 2 = désactivé
```

### 31.2 Contrôle des applications : les 3 niveaux

| Technologie | Granularité | Complexité | Recommandé |
|---|---|---|---|
| **Smart App Control** | Réputation IA (binaire) | Faible (on/off) | Postes standard récents (édition Famille/Pro) |
| **AppLocker** | Règles par éditeur/chemin/hash | Moyenne | Entreprise : standard historique |
| **Windows Defender Application Control (WDAC)** | Stratégies signées, noyau | Élevée | Environnements critiques |

**En entreprise : AppLocker** reste le meilleur compromis (GPO natives, audit possible).

```powershell
# AppLocker : exporter la stratégie en XML (audit)
Get-AppLockerPolicy -Effective -Xml > C:\Admin\applocker-effectif.xml

# Tester une règle
Test-AppLockerPolicy -XmlPolicy C:\Admin\applocker.xml -Path C:\Outils\outil.exe -User "ENTREPRISE\jdupont"
```

Politique AppLocker minimale viable :

1. Règles par défaut (tout le monde peut exécuter dans `C:\Windows` et `C:\Program Files`).
2. Règles **Éditeur** pour les applications métier (plutôt que chemin/hash, plus maintenables).
3. Mode **Audit** 2-4 semaines → analyser le journal → passer en **Appliquer**.

---

## 32. Sécurité : Credential Guard, LSA Protection, HVCI / Memory Integrity

### 32.1 Vue d'ensemble (virtualization-based security)

| Fonctionnalité | Protège contre | Prérequis |
|---|---|---|
| **Credential Guard** | Vol d'identifiants en mémoire (Mimikatz & co) | Entreprise, VBS, UEFI |
| **LSA Protection (PPL)** | Injection de code dans LSASS | Registre ou stratégie |
| **HVCI / Memory Integrity** | Pilotes/code non signés dans le noyau | CPU avec VBS, Secure Boot |

### 32.2 Activer Credential Guard (Entreprise)

GPO : `Configuration ordinateur\Modèles d'administration\Système\Device Guard` → **« Activer la sécurité basée sur la virtualisation »** :

- Sélectionner le niveau de sécurité de plateforme : **Démarrage sécurisé** (ou + DMA).
- Credential Guard : **Activé avec verrouillage UEFI**.

```powershell
# Vérifier que VBS / Credential Guard sont actifs
Get-CimInstance -ClassName Win32_DeviceGuard -Namespace root\Microsoft\Windows\DeviceGuard |
    Select-Object VirtualizationBasedSecurityStatus, SecurityServicesConfigured, SecurityServicesRunning
# SecurityServicesRunning = 1 (Credential Guard) / 2 (HVCI)

# Alternative lisible :
(Get-ComputerInfo).DeviceGuardSecurityServicesConfigured
```

> ⚠️ Credential Guard **casse** l'authentification NTLMv1 et certains fournisseurs d'identifiants tiers. Tester sur l'anneau pilote avant généralisation. Wi-Fi 802.1X avec certificats machine : OK.

### 32.3 LSA Protection

```powershell
# Activer LSA Protection (PPL pour LSASS)
$lsa = "HKLM:\SYSTEM\CurrentControlSet\Control\Lsa"
Set-ItemProperty -Path $lsa -Name "RunAsPPL" -Value 1 -Type DWord
# Redémarrage requis. Vérifier dans le journal :
# Applications and Services Logs\Microsoft\Windows\CodeIntegrity\Operational, ID 3065/3066
```

### 32.4 HVCI / Intégrité de la mémoire

Paramètres → Confidentialité et sécurité → Sécurité Windows → Isolation du noyau → **Intégrité de la mémoire** : Activé (défaut sur installations récentes de Windows 11).

```powershell
# Forcer par stratégie (recommandé en entreprise)
$dg = "HKLM:\SYSTEM\CurrentControlSet\Control\DeviceGuard"
New-Item -Path "$dg\Scenarios\HypervisorEnforcedCodeIntegrity" -Force | Out-Null
Set-ItemProperty -Path "$dg\Scenarios\HypervisorEnforcedCodeIntegrity" -Name "Enabled" -Value 1 -Type DWord
```

> ⚠️ HVCI peut bloquer des **pilotes anciens non signés** (imprimantes, dongles, logiciels industriels). Symptôme : périphérique avec point d'exclamation + ID d'événement 3089. Solution : pilote à jour signé, ou exclusion ciblée (documentée).

---

## 33. Sécurité : durcissement — baselines Microsoft, ASR, Attack Surface Reduction

### 33.1 Security Baselines Microsoft

Microsoft publie des **baselines** (GPO packagées) pour chaque version de Windows : ~200 réglages prêts (Defender, Edge, Credential Guard, SMB, etc.).

Procédure :

```
1. Télécharger la "Security Compliance Toolkit" correspondant à la version (ex. Windows 11 24H2)
2. Importer les GPO via les scripts fournis (Baseline-LocalInstall.ps1 pour tester en local)
3. Comparer avec Policy Analyzer avant/après
4. Adapter : TOUTE baseline se personnalise (ex. : exceptions métier)
```

```powershell
# Appliquer la baseline en local (lab/test uniquement)
.\Baseline-LocalInstall.ps1 -Win11
# Puis vérifier avec Policy Analyzer (même toolkit)
```

### 33.2 Points de durcissement prioritaires (top 15)

| # | Réglage | Effet |
|---|---|---|
| 1 | Désactiver SMBv1 | Protocole obsolète, vecteur WannaCry |
| 2 | Désactiver LLMNR / NetBIOS | Anti-spoofing local |
| 3 | Signer SMB obligatoire | Anti-relay |
| 4 | UAC au maximum | Élévation systématique |
| 5 | Désactiver AutoRun/AutoPlay | Anti-clé USB |
| 6 | Règles ASR en Block (après audit) | Anti-ransomware/phishing |
| 7 | Bloquer les macros Office non signées | Vecteur n°1 |
| 8 | PowerShell Constrained Language (WDAC) | Anti-living-off-the-land |
| 9 | Désactiver PowerShell v2 | Moteur obsolète |
| 10 | LAPS Windows actif | Mots de passe admin locaux uniques |
| 11 | Credential Guard | Anti-vol d'identifiants |
| 12 | HVCI | Anti-pilotes malveillants |
| 13 | BitLocker imposé | Vol/perte |
| 14 | Pare-feu : bloquer entrant par défaut | Réduction de surface |
| 15 | Compte Administrator intégré désactivé | Moins de cible |

```powershell
# Désactiver SMBv1 (si encore présent)
Disable-WindowsOptionalFeature -Online -FeatureName SMB1Protocol -NoRestart
Get-SmbServerConfiguration | Select-Object EnableSMB1Protocol, EnableSMB2Protocol

# Désactiver PowerShell v2
Disable-WindowsOptionalFeature -Online -FeatureName MicrosoftWindowsPowerShellV2Root -NoRestart

