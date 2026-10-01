---
id: collect-261001-rattrapage/rattrapage/win11-guide-8
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [1005, 1176]
sha256: eb00c03705f5a03d4b3366a6bcbbc2ba57e628f682c30bbdf00dcde059b6c7fb
---

# Windows 11 en entreprise — Guide technique ultra-complet

```powershell
# Vérifier l'état d'inscription Entra
dsregcmd /status
# Lignes clés :
#   AzureAdJoined : YES
#   EnterpriseJoined : NO
#   DomainJoined : NO (pur cloud) / YES (hybride)
```

### 14.4 Hybride : points de vigilance

1. **Microsoft Entra Connect** doit synchroniser les objets ordinateurs.
2. Le SCP (Service Connection Point) dans AD indique au poste où s'enregistrer.
3. Délai : l'enregistrement hybride se fait au redémarrage suivant la sync (jusqu'à 30 min).
4. Dépannage : `dsregcmd /status`, journaux `Applications and Services Logs\Microsoft\Windows\User Device Registration`.

```powershell
# Forcer une tentative d'enregistrement (dépannage)
dsregcmd /join /debug
# Quitter proprement avant de recommencer
dsregcmd /leave
```

---

## 15. BitLocker : fondamentaux, TPM seul vs TPM+PIN, XTS-AES

### 15.1 Ce que BitLocker chiffre (et ce qu'il ne chiffre pas)

- ✅ Volume OS, volumes de données, disques amovibles (BitLocker To Go).
- ❌ Ne protège pas contre un OS compromis en cours d'exécution : c'est une protection **au repos** (vol/perte) et **au démarrage** (intégrité via TPM).

### 15.2 Protecteurs : TPM seul vs TPM+PIN vs clé USB

| Protecteur | Sécurité | Ergonomie | Recommandé pour |
|---|---|---|---|
| TPM seul | Bonne (lie le déchiffrement à l'intégrité du boot) | Transparente | Postes fixes en entreprise, standard |
| TPM + PIN | Forte (PIN = 2e facteur, anti-DMA/cold boot) | PIN à chaque démarrage | Portables sensibles, dirigeants, nomades |
| TPM + clé USB | Forte | Clé à brancher | Cas spécifiques (serveurs, postes sans clavier) |
| Mot de passe seul (sans TPM) | Moyenne | Mot de passe | Postes sans TPM (rare sous Win11) |

```powershell
# État BitLocker d'un poste
Get-BitLockerVolume | Select-Object MountPoint, VolumeStatus, ProtectionStatus, EncryptionMethod

# Détail des protecteurs
(Get-BitLockerVolume -MountPoint C:).KeyProtector | Select-Object KeyProtectorType, KeyProtectorId
```

### 15.3 Algorithmes : XTS-AES 128 vs 256

| Algorithme | Vitesse | Sécurité | Recommandation |
|---|---|---|---|
| XTS-AES 128 | + rapide | Suffisante (standard) | **Défaut recommandé** |
| XTS-AES 256 | ~10-20 % plus lent | Marge théorique | Données très sensibles, exigences réglementaires |

> 💡 La différence pratique entre 128 et 256 bits est négligeable face aux autres risques. Standardisez sur **XTS-AES 128** sauf contrainte de conformité explicite. Le paramètre se règle par GPO avant le chiffrement (un changement après = déchiffrement/rechiffrement).

### 15.4 Activer BitLocker en ligne de commande

```powershell
# Chiffrement simple : TPM seul (le plus courant)
Enable-BitLocker -MountPoint "C:" -TpmProtector -UsedSpaceOnly -SkipHardwareTest

# TPM + PIN (le PIN est demandé de façon sécurisée)
Enable-BitLocker -MountPoint "C:" -TpmAndPinProtector -UsedSpaceOnly -SkipHardwareTest
# Puis définir le PIN :
# (Add-BitLockerKeyProtector -MountPoint "C:" -Pin (Read-Host -AsSecureString "PIN") -TpmAndPinProtector)

# Options :
# -UsedSpaceOnly        : ne chiffre que l'espace utilisé (rapide, suffisant sur disque neuf)
# -SkipHardwareTest     : pas de redémarrage de test (à n'utiliser qu'en déploiement maîtrisé)

# Vérifier la progression
Get-BitLockerVolume -MountPoint C: | Select-Object VolumeStatus, EncryptionPercentage
```

```powershell
# Suspendre temporairement (ex. : avant mise à jour du BIOS !)
Suspend-BitLocker -MountPoint "C:" -RebootCount 1

# Reprendre
Resume-BitLocker -MountPoint "C:"
```

> ⚠️ **Avant toute mise à jour du BIOS/UEFI ou changement matériel** : suspendez BitLocker (1 redémarrage). Sinon, le TPM détecte un changement d'intégrité → **écran de récupération** au prochain boot (section 18).

---

(Bloc 1/6 — sections 1 à 15)

---

## 16. BitLocker : clés de récupération (AD, fichier, impression, Entra)

### 16.1 Les 4 destinations de la clé de récupération (48 chiffres)

| Destination | Avantage | Inconvénient |
|---|---|---|
| **Active Directory** (attribut msFVE) | Centralisé, lié à l'objet ordinateur, ACL possibles | Nécessite le schéma étendu + GPO |
| **Microsoft Entra ID** | Visible dans Intune / myapps | Postes Entra join uniquement |
| **Fichier** (`C:\...`, clé USB) | Simple | À protéger (chiffrer le partage, ACL) |
| **Impression** | Hors-ligne | Papier = à coffre-fort |

**En entreprise : imposez AD et/ou Entra par GPO, jamais « au choix de l'utilisateur ».**

### 16.2 Sauvegarder la clé dans AD (prérequis)

1. Étendre le schéma AD (une fois par forêt) : les classes `msFVE-RecoveryInformation` existent depuis Server 2008 — vérifiez avec :

```powershell
# Vérifier que le schéma contient les attributs BitLocker
Get-ADObject -SearchBase (Get-ADRootDSE).schemaNamingContext -Filter { name -like "ms-FVE-*" } |
    Select-Object Name | Format-Table -AutoSize
```

2. GPO : `Configuration ordinateur\Stratégies\Modèles d'administration\Composants Windows\BitLocker\Lecteurs du système d'exploitation` → **« Choisir comment les lecteurs chiffrés par BitLocker peuvent être récupérés »** → cocher « Enregistrer les informations de récupération BitLocker dans les services de domaine Active Directory ».

3. Déléguer le droit d'écriture sur les objets ordinateurs (le poste écrit lui-même sa clé à l'activation).

### 16.3 Récupérer une clé depuis AD

```powershell
# Trouver la clé de récupération d'un poste (depuis une session admin du domaine)
$computer = "PC-COMPTA-042"
Get-ADObject -Filter "objectClass -eq 'msFVE-RecoveryInformation'" `
    -SearchBase (Get-ADComputer $computer).DistinguishedName `
    -Properties msFVE-RecoveryPassword, whenCreated |
    Sort-Object whenCreated -Descending |
    Select-Object -First 1 msFVE-RecoveryPassword, whenCreated
```

### 16.4 Sauvegarder / exporter manuellement

```powershell
# Sauvegarder les protecteurs vers AD (si la GPO ne l'a pas fait)
Backup-BitLockerKeyProtector -MountPoint "C:" -KeyProtectorId "{ID-du-protecteur}"

# Lister les ID de protecteurs
(Get-BitLockerVolume -MountPoint C:).KeyProtector

# Exporter la clé de récupération dans un fichier (à stocker de façon sécurisée !)
(Get-BitLockerVolume -MountPoint C:).KeyProtector |
    Where-Object { $_.KeyProtectorType -eq 'RecoveryPassword' } |
    Select-Object -ExpandProperty RecoveryPassword |
    Out-File "\\SRV-FICHIERS\BitLocker$\PC-042-recovery.txt"
```

### 16.5 Clé dans Entra ID / Intune

- Intune → Appareils → sélectionner l'appareil → **Clés de récupération** → afficher/copier.
- L'utilisateur peut aussi retrouver sa clé sur `myaccount.microsoft.com` (si la stratégie l'autorise).

> 🔐 **Procédure d'urgence** : quand un utilisateur appelle avec l'écran bleu de récupération, demandez-lui l'**ID de clé** affiché (8 premiers caractères), retrouvez la clé correspondante dans AD/Intune, dictez les 48 chiffres par blocs de 6. Ne jamais envoyer une clé de récupération par e-mail non chiffré.

---

## 17. BitLocker : stratégies GPO, MBAM / BitLocker Management, déploiement silencieux

### 17.1 GPO BitLocker — les réglages à standardiser

Emplacement : `Configuration ordinateur\Stratégies\Modèles d'administration\Composants Windows\BitLocker\`

| Stratégie | Réglage recommandé | Pourquoi |
|---|---|---|
| Choisir la méthode de chiffrement (lecteurs OS) | XTS-AES 128 | Standard, rapide (section 15) |
| Exiger un chiffrement de l'espace utilisé uniquement | Activé | Déploiement rapide sur disques neufs |
| Choisir comment les lecteurs peuvent être récupérés | Sauvegarde AD obligatoire | Pas de clé = pas de chiffrement |
| Configurer l'utilisation de mots de passe (lecteurs OS) | Selon politique PIN | — |
| Refuser l'écriture sur lecteurs amovibles non chiffrés | Activé (optionnel) | DLP basique via BitLocker To Go |
| Activer BitLocker sans TPM compatible | Selon parc | À éviter sous Win11 (TPM requis de toute façon) |

