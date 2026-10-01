---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-8
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft", "Nvidia"]
dates: []
keywords: ["distribution", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [955, 1159]
sha256: c14f35f9d88d73a79c459cb0ad5c155cd2aa6fcd2412c30d1c8109bab225e549
---

# BIOS / UEFI — Secure Boot — TPM 2.0

| | Mode Standard | Mode Custom |
|---|---|---|
| Clés | Clés d'usine (Microsoft + OEM) | Modifiables par l'administrateur |
| Usage | **99 % des entreprises** | Cas spéciaux : OS maison, chargeurs auto-signés, R&D |
| Sécurité | Maximale par défaut | Dépend de la rigueur de l'admin |
| Gestion | Automatique (Windows Update) | Manuelle (import/export de clés) |

### Quand utiliser le mode Custom

- Déployer un **chargeur auto-signé** (outil de déploiement maison, noyau Linux custom) sans passer par le shim Microsoft.
- **Enrôler sa propre PK** d'entreprise (modèle « je suis le propriétaire de la plateforme », très avancé, rarement justifié).
- Signer des binaires EFI internes avec une CA d'entreprise ajoutée à `db`.

### Procédure type (mode Custom)

```
1. Passer le firmware en "Custom Mode".
2. Exporter/sauvegarder les clés d'usine (fichiers .auth / .esl selon firmware).
3. Enrôler la clé : importer le certificat (.cer/.der) dans db.
   (Certains firmwares exigent un fichier signé .auth ; d'autres acceptent
   l'import direct depuis une clé USB en FAT32.)
4. Signer le binaire EFI avec sbsign (Linux) ou signtool (Windows).
5. Tester le boot, puis verrouiller (repasser en User Mode si PK posée).
```

> 💡 **Recommandation :** restez en **mode Standard** sauf besoin métier démontré. Le mode Custom multiplie les risques d'erreur (machine qui ne boote plus après un import raté) et la charge de gestion des clés.

---

## 23. Activer Secure Boot : procédure

### Préconditions

- [ ] Firmware en **UEFI natif** (CSM **désactivé**).
- [ ] Disque système en **GPT**.
- [ ] OS compatible : Windows 8+ (64 bits), Windows 11 (requis), Linux avec shim signé.

### Procédure générique

```
1. Redémarrer → entrer dans le setup UEFI (F2 / DEL / F10 / F12 selon marque).
2. Aller dans l'onglet Boot ou Security.
3. Secure Boot → Enabled.
   - Si l'option est grisée : vérifier que le mode est UEFI (pas Legacy/CSM),
     et que les clés sont provisionnées (Secure Boot Mode = Standard,
     état ≠ "Setup Mode").
4. Si "Setup Mode" : choisir "Provision Factory Keys" / "Restore Factory Keys"
   (selon firmware) pour installer les clés d'usine.
5. Sauvegarder (F10) et redémarrer.
6. Vérifier dans Windows (section suivante).
```

### Cas particulier : activer Secure Boot sur un Windows installé en Legacy

Impossible directement. Ordre des opérations :

```
1. mbr2gpt /convert (section 8) — BitLocker suspendu !
2. Basculer le firmware en UEFI (CSM off).
3. Activer Secure Boot.
4. Vérifier : Confirm-SecureBootUEFI → True.
```

### Déploiement en masse

- **Dell :** Dell Command Configure (`cctk --secureboot=enable`), ou profils BIOS via SCCM/Intune.
- **HP :** HP BIOS Configuration Utility (`BiosConfigUtility.exe`), ou HP Manageability Kit.
- **Lenovo :** ThinkPad Setup Settings Capture/Deploy, ou WMI (`Lenovo_SetBiosSetting`).
- **Intune :** profils de configuration (Endpoint Security) pour exiger Secure Boot.

---

## 24. Vérifier Secure Boot : Confirm-SecureBootUEFI et msinfo32

```powershell
# Méthode 1 : cmdlet officiel (retourne True / False)
Confirm-SecureBootUEFI

# Méthode 2 : msinfo32 → ligne "État du démarrage sécurisé"
msinfo32

# Méthode 3 : registre (utile en inventaire, sans cmdlet)
Get-ItemProperty 'HKLM:\SYSTEM\CurrentControlSet\Control\SecureBoot\State' |
    Select-Object UEFISecureBootEnabled
# 1 = activé, 0 = désactivé

# Méthode 4 : System Information via CIM
Get-CimInstance -ClassName Win32_ComputerSystem |
    Select-Object Name, Domain, Manufacturer, Model
```

### Interpréter les résultats

| `Confirm-SecureBootUEFI` | Registre | Signification |
|---|---|---|
| `True` | 1 | Secure Boot actif ✅ |
| `False` | 0 | Secure Boot désactivé ou firmware en Legacy ⚠️ |
| Erreur « cmdlet not supported » | — | Machine en **Legacy BIOS** ou firmware sans Secure Boot ❌ |

### Script de vérification rapide (une machine)

```powershell
#requires -RunAsAdministrator
$sb = Confirm-SecureBootUEFI
$tpm = Get-Tpm
[PSCustomObject]@{
    ComputerName      = $env:COMPUTERNAME
    SecureBoot        = $sb
    TpmPresent        = $tpm.TpmPresent
    TpmReady          = $tpm.TpmReady
    Conforme_Win11    = ($sb -and $tpm.TpmPresent -and $tpm.TpmReady)
} | Format-List
```

---

## 25. Mise à jour des clés et certificats Secure Boot

### Pourquoi mettre à jour

- **Révocations (dbx)** : bloquer les chargeurs vulnérables (patchs de sécurité).
- **Nouveaux certificats (db)** : remplacer les certificats expirés (échéance 2026 des certificats 2011).
- **Nouveaux OS/chargeurs** : ajouter de nouvelles autorités (rare en Standard).

### Canaux de distribution

| Canal | Ce qu'il met à jour |
|---|---|
| **Windows Update** | `dbx` (KB de révocation), nouveaux certificats `db`/`KEK` Microsoft |
| **Mise à jour firmware OEM** | Clés OEM, `dbx` embarquée, nouveaux certificats |
| **Manuel (mode Custom)** | Import depuis clé USB dans le setup |

### Vérifier l'état des mises à jour dbx

```powershell
# Voir les KB de mise à jour dbx installées
Get-HotFix | Where-Object { $_.HotFixID -match 'KB5012170|KB5025885|KB5036212' } |
    Select-Object HotFixID, InstalledOn, Description

# Journal des mises à jour Secure Boot (observateur d'événements)
# → Journaux Windows → Système, source "SecureBoot" / ID liés au firmware
Get-WinEvent -LogName System -MaxEvents 50 |
    Where-Object { $_.ProviderName -match 'SecureBoot|Kernel-Boot' } |
    Select-Object TimeCreated, Id, LevelDisplayName, Message |
    Format-Table -AutoSize -Wrap
```

### Plan d'action entreprise (renouvellement des certificats)

```
1. INVENTAIRE : lister les modèles du parc et leurs versions de firmware
   (script section 40 + Get-CimInstance Win32_BIOS).
2. PILOTE : appliquer firmware + KB sur 5-10 machines représentatives,
   vérifier Confirm-SecureBootUEFI et le boot (y compris double-boot).
3. DÉPLOIEMENT : vagues de 25 %, avec point de contrôle BitLocker
   (suspendre BitLocker avant flash, section 37).
4. CONTRÔLE : re-vérifier Secure Boot + dbx à jour sur 100 % du parc.
5. DOCUMENTATION : noter les versions de firmware validées par modèle.
```

> ⚠️ **Ne jamais bloquer les mises à jour dbx** par « prudence » : une dbx obsolète laisse passer des chargeurs vulnérables connus. Le risque d'une dbx à jour (rare blocage d'un vieux chargeur) est très inférieur au risque inverse.

---

## 26. Secure Boot et Linux en entreprise

### Le mécanisme : le shim

La plupart des distributions « Secure Boot compatibles » (Ubuntu, Red Hat/RHEL, SUSE, Debian depuis Bullseye...) utilisent un **shim** :

```
Firmware → shimx64.efi (signé Microsoft, clé "UEFI CA 2011" dans db)
           └─ vérifie grubx64.efi avec la clé de la distribution (MOK)
              └─ GRUB vérifie le noyau signé
```

**MOK** (*Machine Owner Key*) : base de clés gérée par l'utilisateur/la distribution, distincte de `db`. Permet d'ajouter ses propres clés (ex. modules DKMS VirtualBox/NVIDIA auto-signés) via `mokutil`.

### Commandes utiles (Linux)

```bash
# Vérifier Secure Boot
mokutil --sb-state
# ou
bootctl status | grep -i secure

# Lister les clés MOK enrôlées
mokutil --list-enrolled

# Enrôler une nouvelle clé MOK (demande un mot de passe, validé au reboot)
mokutil --import ma-cle.der
```

### En entreprise

- **Postes Linux :** choisir une distribution avec shim signé ; documenter l'enrôlement MOK pour les modules propriétaires.
- **Double-boot :** Windows Update peut mettre à jour `dbx` et révoquer un vieux shim → prévoir la procédure de mise à jour du Linux (section 78).
- **Serveurs Linux :** Secure Boot activable sur la plupart des serveurs récents ; vérifier la compatibilité de la distribution avant.

---

## 27. Dépannage Secure Boot : boot bloqué

### Symptômes typiques

