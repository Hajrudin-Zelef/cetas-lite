---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-10
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [1334, 1542]
sha256: a97c7b105f5151eea416a170252bb9148c8d0efc2a2cd532a8d1807a1d5f9106
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```powershell
# Dell : Dell Command Configure
cctk --tpm=on
cctk --tpmactivation=activate

# Lenovo : via WMI (exemple)
(Get-WmiObject -Namespace root\wmi -Class Lenovo_SetBiosSetting).SetBiosSetting("SecurityChip,Active")

# HP : HP BIOS Configuration Utility (fichier de config)
# Security/TPM Device=Available ; Security/TPM State=On
```

> ⚠️ **L'activation du TPM peut nécessiter un redémarrage + validation physique** (appuyer sur une touche pour confirmer l'activation : protection anti-malware). Prévoyez-le dans le planning de déploiement : pas de « tout à distance sans toucher les machines ».

---

## 32. Vérifier le TPM : Get-TPM et tpm.msc

### Get-TPM : la méthode scriptable

```powershell
Get-Tpm | Format-List *
```

Sortie typique d'un TPM sain :

```
TpmPresent                : True
TpmReady                  : True
TpmEnabled                : True
TpmActivated              : True
TpmOwned                  : True
RestartPending            : False
ManufacturerId            : 1229346816
ManufacturerIdTxt         : IFX
ManufacturerVersion       : 7.85.4555.0
ManagedAuthLevel          : Full
OwnerAuth                 :
OwnerClearDisabled        : False
AutoProvisioning          : Enabled
LockedOut                 : False
LockoutHealTime           : 2 hours
LockoutCount              : 0
LockoutMax                : 32
SelfTest                  : {}
```

### Interprétation des champs clés

| Champ | Signification | Valeur attendue |
|---|---|---|
| `TpmPresent` | Puce détectée par Windows | `True` |
| `TpmReady` | Prêt à l'emploi (provisionné) | `True` |
| `TpmEnabled` / `TpmActivated` | Activé dans le firmware | `True` |
| `TpmOwned` | Propriétaire défini (provisionné) | `True` |
| `AutoProvisioning` | Windows peut provisionner seul | `Enabled` |
| `LockedOut` | Verrouillage anti-bruteforce actif | `False` (normal) |
| `RestartPending` | Un redémarrage est requis | `False` (sinon : rebooter) |

### tpm.msc : la console graphique

```
Win + R → tpm.msc
```

Affiche : fabricant, version de la spécification (2.0), état (« Le TPM est prêt à être utilisé »), et les actions : **Préparer le TPM**, **Effacer le TPM**, **Activer/Désactiver**.

### Diagnostic rapide en une ligne

```powershell
$t = Get-Tpm
if ($t.TpmPresent -and $t.TpmReady) { "TPM OK" }
elseif ($t.TpmPresent) { "TPM présent mais non prêt → Initialize-Tpm puis reboot" }
else { "TPM ABSENT → vérifier le firmware (Security → TPM On)" }
```

---

## 33. Clear TPM : quand, pourquoi, précautions

**Effacer le TPM** réinitialise la puce à l'état d'usine : toutes les clés, tous les secrets scellés sont **détruits définitivement**.

### Quand c'est nécessaire

| Situation | Pourquoi |
|---|---|
| **Cession / reconditionnement** d'un poste | Supprimer les clés de l'ancien utilisateur (BitLocker, Hello, certificats) |
| **Changement de propriétaire** (sortie d'un collaborateur sensible) | Hygiène cryptographique |
| **TPM corrompu / incohérent** | Dernier recours après échec de provisioning |
| **Recyclage vers un autre usage** (poste → serveur de test) | Repartir sur une base saine |

### Quand c'est INTERDIT (sans précautions)

| Situation | Risque |
|---|---|
| BitLocker actif avec protecteur TPM | **Disque illisible** → clé de récupération obligatoire |
| Windows Hello for Business déployé | L'utilisateur devra ré-enrôler ses identifiants |
| Certificats machine dans le TPM (EAP-TLS, VPN) | Perte des certificats → réémission PKI |

### Procédure sécurisée

```powershell
# 1. VÉRIFIER BitLocker (si actif → ne PAS continuer sans la clé de récupération)
manage-bde -status C:

# 2. Sauvegarder la clé de récupération (elle doit être dans l'AD/Entra ID)
manage-bde -protectors -get C: -type RecoveryPassword

# 3. Suspendre BitLocker (le temps de l'opération)
Suspend-BitLocker -MountPoint "C:" -RebootCount 1

# 4. Effacer le TPM (demande un redémarrage + validation physique)
Clear-Tpm
# ou via tpm.msc → "Effacer le TPM" → redémarrer → confirmer (F12 ou touche indiquée)

# 5. Après redémarrage : re-provisionner
Initialize-Tpm
# (Windows le fait aussi automatiquement si AutoProvisioning = Enabled)

# 6. Réactiver BitLocker
Resume-BitLocker -MountPoint "C:"
```

> 🔴 **Règle absolue :** jamais de `Clear-Tpm` sans avoir **vérifié et sauvegardé la clé de récupération BitLocker**. C'est l'erreur qui transforme un poste en presse-papier.

---

## 34. Provisioning et attestation du TPM

### Le provisioning (prise de possession)

Le **provisioning** prépare le TPM : création de la hiérarchie de stockage (SRK), définition du propriétaire. Sur Windows moderne, c'est **automatique** :

```powershell
# Vérifier l'auto-provisioning
Get-Tpm | Select-Object AutoProvisioning, TpmReady, TpmOwned

# Forcer le provisioning manuel (si TpmReady = False)
Initialize-Tpm
# → redémarrage souvent nécessaire

# Désactiver l'auto-provisioning (durcissement avancé, rare)
Disable-TpmAutoProvisioning
# Réactiver :
Enable-TpmAutoProvisioning
```

### L'attestation : prouver l'intégrité à distance

L'**attestation** permet à un serveur de vérifier *à distance* qu'une machine a démarré sainement :

```
1. Au boot, chaque composant est "mesuré" (hash SHA-256) dans les PCR.
2. Le TPM signe les valeurs PCR avec une clé d'attestation (AIK).
3. Le serveur de vérification compare les mesures à des valeurs de référence.
4. Conforme → accès accordé (ex. Conditional Access, accès VPN Zero Trust).
   Non conforme → accès refusé / quarantaine.
```

### En entreprise : où on la rencontre

| Solution | Usage de l'attestation |
|---|---|
| **Microsoft Entra Conditional Access** | Conformité de l'appareil (Intune + attestation) |
| **Windows Autopilot** | Vérification matérielle à l'enrôlement |
| **Solutions Zero Trust / NAC** | Preuve d'intégrité avant accès réseau |
| **Credential Guard** | Isolation basée sur l'intégrité mesurée |

---

## 35. PCR : Platform Configuration Registers

Les **PCR** sont des registres du TPM qui accumulent les **mesures** (hashes) des composants de démarrage. On ne les écrit pas : on les **étend** (`PCR_new = Hash(PCR_old || mesure)`), ce qui rend toute falsification détectable.

### Les 24 PCR et leur usage (TPM 2.0, profil PC Client)

| PCR | Mesuré par | Contenu typique |
|---|---|---|
| **0** | Firmware (SRTM) | Code du firmware / BIOS |
| **1** | Firmware | Configuration du firmware |
| **2** | Firmware | Code des ROMs d'extension (cartes) |
| **3** | Firmware | Données des ROMs d'extension |
| **4** | Chargeur (MBR/GPT) | **MBR / gestionnaire de boot** (bootmgfw.efi) |
| **5** | Chargeur | Configuration du gestionnaire (BCD) |
| **6** | — | Données spécifiques plateforme |
| **7** | Firmware | **État Secure Boot** (clés, mode) |
| **8–15** | OS | Composants OS (winload, noyau...) — usage OS |
| **16–23** | — | Usage libre / débogage |

### Les PCR critiques pour BitLocker

Par défaut, BitLocker scelle la clé sur les PCR **0, 2, 4, 11** (profil Windows standard) :

```powershell
# Voir le profil PCR utilisé par BitLocker
manage-bde -protectors -get C: -type TPM
# Affiche : "PCR Validation Profile: 0, 2, 4, 11"
```

| PCR | Pourquoi il compte pour BitLocker |
|---|---|
| 0 | **Firmware** → toute MAJ BIOS change PCR 0 → récupération BitLocker ! |
| 2 | ROMs d'extension → ajout d'une carte = changement |
| 4 | Chargeur de boot → modification du boot = changement |
| 7 | Secure Boot (non inclus par défaut, mais recommandé en durcissement) |
| 11 | Code du gestionnaire de démarrage |

### Lire les PCR (diagnostic)

