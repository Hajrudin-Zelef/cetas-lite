---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-17
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "datacenter", "gpu", "intel", "open source"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [2710, 2915]
sha256: 8d648356330343681bbca9a767153ed2c10b864fe780d1b4f41f2b91aacf620d
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```
Trimestriel : vérifier les bulletins de sécurité firmware des constructeurs.
Semestriel  : vague de MAJ firmware (lot pilote → 25 % → 100 %),
              toujours avec BitLocker suspendu + séquestre vérifié.
Annuel      : revue des réglages (nouveautés : Kernel DMA, Pluton, etc.).
```

### Documentation

- [ ] Matrice **modèle → version BIOS validée → réglages** (wiki interne).
- [ ] Procédure **mot de passe BIOS oublié** par constructeur (§69), avec contacts support et preuves d'achat.
- [ ] Registre des **exceptions** (postes en Legacy, Secure Boot désactivé...) : qui, pourquoi, jusqu'à quand. Une exception non documentée = une faille.

### Supervision

```powershell
# Exemple : remonter les postes avec Secure Boot désactivé (via le CSV d'inventaire §40)
Import-Csv .\inventaire-firmware.csv |
    Where-Object { $_.SecureBoot -ne 'True' } |
    Select-Object ComputerName, Model, BiosVersion
# → à traiter ou à justifier (exception documentée).
```

---

## 61. UEFI en datacenter : les serveurs

### Particularités des serveurs

| Aspect | Serveur vs poste |
|---|---|
| Firmware | UEFI avec **Redfish/BMC** (administration à distance) |
| Boot | Souvent **réseau** (iPXE/HTTP Boot) ou SAN |
| Secure Boot | Supporté (à activer, avec les clés de l'OS serveur) |
| TPM | TPM 2.0 sur carte mère serveur (pour BitLocker, attestation) |
| Mises à jour | Via le BMC (montage d'ISO virtuel) ou outils OEM (iDRAC, iLO, XCC) |

### Ordre de boot type d'un hyperviseur

```
1. Disque local / RAID (hyperviseur)
2. Réseau (iPXE) — pour re-provisionnement
3. (Désactiver l'USB boot)
```

### Bonnes pratiques

- [ ] **Mot de passe du BIOS serveur** défini (distinct du BMC !).
- [ ] **Secure Boot activé** (vérifier la compatibilité de l'hyperviseur : ESXi, Hyper-V, Proxmox le supportent).
- [ ] **Boot réseau sécurisé** : iPXE avec HTTPS, pas de TFTP en clair sur un réseau non fiable.
- [ ] **Firmware à jour** : BIOS + BMC + cartes (planifié, jamais « à l'arrache »).
- [ ] **TPM activé** pour le chiffrement des VM / vTPM.

---

## 62. Redfish et BMC : administrer le firmware à distance

Le **BMC** (*Baseboard Management Controller* : iDRAC Dell, iLO HP, XCC Lenovo, IMM...) permet d'administrer le serveur **hors OS**, via **Redfish** (API REST standard).

### Cas d'usage firmware

- Monter une ISO à distance (flash du BIOS sans se déplacer).
- Lire/modifier les réglages UEFI (ordre de boot, Secure Boot).
- Forcer un boot PXE au prochain redémarrage.
- Inventorier les versions de firmware du parc serveurs.

### Exemple Redfish (PowerShell)

```powershell
# 1. Lire les réglages BIOS actuels
$bmc = 'https://idrac-01.lan.lan'
$cred = Get-Credential  # compte BMC
Invoke-RestMethod -Uri "$bmc/redfish/v1/Systems/System.Embedded.1/Bios" `
    -Credential $cred -SkipCertificateCheck |
    Select-Object -ExpandProperty Attributes |
    Select-Object BootMode, SecureBoot, TpmSecurity

# 2. Forcer un boot PXE au prochain redémarrage
$body = @{ Boot = @{ BootSourceOverrideTarget = 'Pxe';
                     BootSourceOverrideEnabled = 'Once' } } |
        ConvertTo-Json
Invoke-RestMethod -Uri "$bmc/redfish/v1/Systems/System.Embedded.1" `
    -Method Patch -Body $body -ContentType 'application/json' `
    -Credential $cred -SkipCertificateCheck

# 3. Redémarrer
$reset = @{ ResetType = 'ForceRestart' } | ConvertTo-Json
Invoke-RestMethod -Uri "$bmc/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset" `
    -Method Post -Body $reset -ContentType 'application/json' `
    -Credential $cred -SkipCertificateCheck
```

> 🔒 **Sécurité BMC :** changer les mots de passe par défaut (root/calvin...), dédier un VLAN d'administration, HTTPS uniquement, désactiver les vieux protocoles (Telnet !). Un BMC compromis = contrôle total du serveur, **en dessous de l'OS**.

---

## 63. Secure Boot sur Hyper-V : les VM de génération 2

Hyper-V propose deux générations de VM :

| | Génération 1 | Génération 2 |
|---|---|---|
| Firmware | BIOS émulé | **UEFI émulé** |
| Disque système | MBR (IDE) | **GPT** (SCSI) |
| **Secure Boot** | ❌ | ✅ (modèles de stratégie) |
| PXE | Legacy | UEFI |
| Usage | Vieux OS | **Tout OS moderne** |

### Activer Secure Boot sur une VM gen2

```powershell
# À la création
New-VM -Name 'SRV-WEB-01' -Generation 2 -MemoryStartupBytes 4GB `
    -NewVHDPath 'D:\VMs\SRV-WEB-01.vhdx' -NewVHDSizeBytes 80GB

# Sur une VM existante (éteinte)
Set-VMFirmware -VMName 'SRV-WEB-01' -EnableSecureBoot On

# Choisir le modèle (template) de clés
Set-VMFirmware -VMName 'SRV-WEB-01' -SecureBootTemplate 'MicrosoftWindows'
# Modèles disponibles :
#   MicrosoftWindows      → Windows uniquement
#   MicrosoftUEFICertificateAuthority → Windows + Linux (shim)
#   OpenSourceShieldedVM  → distributions open source
```

### vTPM (TPM virtuel)

```powershell
# Ajouter un vTPM à une VM gen2 (prérequis : hôte avec TPM ou HGS pour les shielded VM)
# 1. Chiffrer la VM avec un protecteur par certificat (key protector)
$cert = New-SelfSignedCertificate -Subject "ShieldedVM"
$kp = New-HgsKeyProtector -Owner $cert -AllowUntrustedRoot
Set-VMKeyProtector -VMName 'SRV-WEB-01' -KeyProtector $kp.RawData

# 2. Activer le vTPM
Enable-VMTPM -VMName 'SRV-WEB-01'

# 3. Vérifier
Get-VMSecurity -VMName 'SRV-WEB-01' |
    Select-Object VMName, TpmEnabled, Shielded
```

> 💡 **Cas d'usage :** BitLocker **dans** la VM (vTPM), Windows 11 en VM (exige vTPM + Secure Boot), Shielded VM en datacenter.

---

## 64. Cas pratique 1 : TPM non détecté par Windows

### Contexte

Un poste neuf (ou après une MAJ BIOS) : `Get-Tpm` retourne `TpmPresent = False`, `tpm.msc` affiche « TPM introuvable ». BitLocker refuse de s'activer.

### Diagnostic pas à pas

```powershell
# Étape 1 : confirmer l'absence
Get-Tpm | Select-Object TpmPresent, TpmReady
# → False / False

# Étape 2 : le firmware voit-il la puce ? (pas de cmdlet direct :
# c'est le setup UEFI qui tranche — mais on peut déjà vérifier le pilote)
Get-PnpDevice -Class 'SecurityDevices' -ErrorAction SilentlyContinue |
    Select-Object FriendlyName, Status
# → vide = Windows ne voit aucun périphérique TPM du tout
```

### Résolution

```
Étape 3 : redémarrer → setup UEFI → Security :
  - TPM / Security Chip : est-il sur "Off" ou "Hidden" ?
    → passer sur On / Enabled.
  - Version : choisir TPM 2.0 (pas 1.2).
  - Sur AMD : "AMD fTPM switch" → AMD CPU fTPM.
    Sur Intel : "PTT" → Enabled.
  - Sauvegarder (F10), redémarrer.

Étape 4 : dans Windows
  Get-Tpm → TpmPresent = True ?
  Si TpmReady = False → Initialize-Tpm (ou tpm.msc → "Préparer le TPM"),
  puis redémarrer.

Étape 5 : si toujours absent
  - Re-flasher le BIOS (un flash corrompu peut masquer le TPM).
  - Clear CMOS (§71) en dernier recours.
  - Vérifier la compatibilité : certains CPU très anciens n'ont ni
    fTPM ni PTT et la carte mère n'a pas de header TPM → pas de TPM
    possible → poste à remplacer (non conforme Windows 11).
```

### Points de vigilance

- Sur **HP**, l'option s'appelle parfois « TPM Embedded Security » avec *deux* réglages : `TPM Device` (Available/Hidden) **et** `TPM State` (On/Off). Il faut les deux.
- Sur **Lenovo**, après un Clear CMOS, le Security Chip peut repasser en « Inactive » : le réactiver + `Initialize-Tpm`.

---

## 65. Cas pratique 2 : Secure Boot bloque le démarrage après un changement matériel

### Contexte

Après l'ajout d'une carte réseau 10 Gb/s (ou d'un GPU) sur une station, le PC affiche au démarrage : `Secure Boot Violation — Invalid signature detected`, puis retourne au setup.

### Analyse

