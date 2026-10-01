---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-13
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [1937, 2121]
sha256: 10bfc65ab483c31779cc275493871e7d191c09e465272690f0b74d7e96ed9749
---

# BIOS / UEFI — Secure Boot — TPM 2.0

> ⚠️ Sur certains BIOS, VT-d est caché dans un sous-menu « Chipset » ou « Security ». Et sur les portables d'entrée de gamme anciens, l'option peut ne pas exister (VT-x présent mais non exposé).

---

## 43. Mot de passe BIOS/UEFI (admin / supervisor)

Le **mot de passe administrateur du firmware** verrouille l'accès au setup et (selon les modèles) le changement de l'ordre de boot.

### Les 3 niveaux (selon constructeurs)

| Niveau | Nom courant | Protège |
|---|---|---|
| **Admin / Supervisor** | `Admin Password`, `Supervisor Password` | L'accès au setup UEFI lui-même |
| **System / User** | `System Password`, `User Password` | Le **démarrage** de la machine (demandé avant le boot) |
| **HDD** | `HDD Password` | Le **disque** (chiffrement ATA, lié au disque, suit le disque) |

### Recommandations

- [ ] **Toujours** définir un mot de passe **admin** sur le parc (déploiement via outils constructeur).
- [ ] Mot de passe **système** (au boot) : uniquement sur postes sensibles — il bloque aussi le Wake-on-LAN et les redémarrages à distance non supervisés.
- [ ] Stocker les mots de passe dans un **coffre** (Keepass d'équipe, solution PAM) : un mot de passe admin perdu = procédure constructeur lourde (section 69).
- [ ] Ne **jamais** utiliser le même mot de passe admin sur tout le parc *sans* plan de rotation ; à défaut, au minimum un mot de passe par site/service.

### Déploiement en masse

```powershell
# Dell (Dell Command Configure) — définit le mot de passe admin
cctk --setuppwd=<nouveau> --valsetuppwd=<ancien si existant>

# HP (BiosConfigUtility) — via fichier config :
#   Setup Password
#       <nouveau encodé>

# Lenovo — via WMI :
(Get-WmiObject -Namespace root\wmi -Class Lenovo_SetBiosPassword).SetBiosPassword(
    "pap,$null,<nouveau>,ascii,us")   # pap = password admin present
```

> 🔴 **Mot de passe BIOS oublié ?** Voir le cas pratique 6 (section 69) : procédures par constructeur, jamais de « master password » universel fiable.

---

## 44. Désactivation des périphériques : USB boot et ports

### L'objectif

Empêcher le démarrage sur un support externe (clé USB, disque USB) et limiter l'exfiltration : c'est la contre-mesure directe à l'**evil maid** de base (section 58).

### Options typiques du setup

| Option | Effet | Recommandation |
|---|---|---|
| `USB Boot` / `Boot from USB` | Autorise/interdit le boot USB | **Disabled** (sauf exceptions documentées) |
| `USB Ports` (avant/arrière) | Désactive physiquement des ports | Désactiver les ports avant sur postes exposés |
| `Thunderbolt Boot` / `Pre-boot DMA` | Boot et DMA via Thunderbolt | **Disabled** sauf besoin (risque DMA — voir VT-d §42) |
| `SD Card Boot`, `CD/DVD Boot` | Boot sur ces supports | Disabled si inutilisés |
| `PXE Boot` | Boot réseau | Enabled uniquement si WDS/déploiement utilisé |

### En complément (couche Windows)

Le verrouillage firmware ne suffit pas : un OS démarré normalement voit toujours les clés USB. Compléter avec :

```powershell
# Restreindre le stockage USB amovible via stratégie (exemple registre)
# HKLM\SYSTEM\CurrentControlSet\Services\USBSTOR → Start = 4 (désactivé)
Set-ItemProperty 'HKLM:\SYSTEM\CurrentControlSet\Services\USBSTOR' -Name Start -Value 4

# Via GPO (recommandé) : Modèles d'administration → Système →
# Accès au stockage amovible → "Disques amovibles : refuser l'accès en écriture"
```

> ⚠️ **Attention au support :** si vous désactivez totalement l'USB boot, prévoyez une procédure d'exception (mot de passe admin + réactivation temporaire tracée) pour les réinstallations sur site.

---

## 45. Wake-on-LAN (WoL)

Le **WoL** permet d'allumer une machine éteinte (mais branchée) en lui envoyant un **paquet magique** (*magic packet*) sur le réseau.

### Prérequis (les 3 couches)

```
1. FIRMWARE : Wake-on-LAN → Enabled (souvent dans Power Management).
   ⚠️ Sur portable : ne fonctionne généralement que sur secteur + parfois
   uniquement en Ethernet (pas en Wi-Fi, ou "Wake on WLAN" séparé).
2. WINDOWS : carte réseau → Propriétés → Gestion de l'alimentation →
   ✅ "Autoriser cet appareil à sortir l'ordinateur du mode veille"
   ✅ "Autoriser uniquement un paquet magique à sortir l'ordinateur du mode veille"
   + désactiver le "démarrage rapide" (fast startup) qui met en hybride.
3. RÉSEAU : le paquet magique est en broadcast → ne traverse pas les routeurs.
   Prévoir un relais par sous-réseau (agent, switch avec directed broadcast,
   ou outil de déploiement type SCCM).
```

### Le paquet magique

```
FF FF FF FF FF FF  + 16 × adresse MAC de la cible
(ex. pour AA:BB:CC:DD:EE:FF : 6 octets FF puis 16 répétitions de la MAC)
Envoyé en UDP, ports 7 ou 9, en broadcast.
```

### Envoyer un WoL en PowerShell

```powershell
function Send-WakeOnLan {
    param([Parameter(Mandatory)][string]$MacAddress)
    $mac = ($MacAddress -replace '[^0-9A-Fa-f]','')
    $packet = [byte[]]::new(102)
    for ($i = 0; $i -lt 6; $i++) { $packet[$i] = 0xFF }
    $macBytes = for ($i = 0; $i -lt 12; $i += 2) {
        [Convert]::ToByte($mac.Substring($i,2), 16) }
    for ($i = 0; $i -lt 16; $i++) {
        [Array]::Copy($macBytes, 0, $packet, 6 + $i*6, 6) }
    $udp = New-Object Net.Sockets.UdpClient
    $udp.Connect([Net.IPAddress]::Broadcast, 9)
    $udp.Send($packet, $packet.Length) | Out-Null
    $udp.Close()
    "Paquet magique envoyé à $MacAddress"
}

Send-WakeOnLan -MacAddress 'AA-BB-CC-DD-EE-FF'
```

### Vérifier la config côté Windows

```powershell
Get-NetAdapter | Select-Object Name, MacAddress, Status
# Puis, pour chaque carte :
Get-NetAdapterPowerManagement -Name 'Ethernet' |
    Select-Object Name, WakeOnMagicPacket, WakeOnPattern, DeviceSleepOnDisconnect
```

### En entreprise

- **Inventaire nocturne / patching :** allumer les postes à 3h, patcher, éteindre (GPO + tâche planifiée).
- **Limites :** ne traverse pas les VLANs sans relais ; incompatible avec le mot de passe système au boot (§43) ; sur les portables, préférer la sortie de veille moderne.

---

## 46. PXE boot : principe

Le **PXE** (*Preboot Execution Environment*) permet à une machine **sans OS** de démarrer sur une image fournie par le réseau. C'est le fondement du déploiement automatisé (WDS, SCCM, MDT).

### Séquence PXE (simplifiée)

```
① Client : firmware avec ROM PXE → DHCP Discover (avec options PXE)
② Serveur DHCP : répond avec IP + option 66 (serveur TFTP) / 67 (fichier boot)
   └─ ou : le serveur WDS répond lui-même (DHCP option 60 = "PXEClient")
③ Client : télécharge le chargeur via TFTP (wdsmgfw.efi en UEFI, wdsnbp.com en Legacy)
④ Chargeur : affiche le menu WDS (F12 pour boot réseau)
⑤ Téléchargement de l'image de démarrage (boot.wim) via TFTP/multicast
⑥ WinPE démarre → assistant d'installation / déploiement
```

### Composants réseau requis

| Composant | Rôle |
|---|---|
| **DHCP** | Attribue l'IP au client PXE |
| **TFTP** | Transfère le chargeur initial (petit, UDP) |
| **Serveur de déploiement** | WDS / SCCM : fournit les images |
| **DNS/AD** (WDS) | WDS s'intègre à l'AD en mode natif |

### Options DHCP critiques

| Option | Nom | Valeur |
|---|---|---|
| 60 | Class ID | `PXEClient` (posée par le serveur WDS lui-même) |
| 66 | Boot Server Host Name | Nom/IP du serveur TFTP |
| 67 | Bootfile Name | `boot\x64\wdsmgfw.efi` (UEFI) ou `boot\x64\wdsnbp.com` (Legacy) |

> ⚠️ **Ne pas poser les options 66/67 sur le DHCP si le WDS est sur le même réseau** : le WDS répond directement via l'option 60. Les options 66/67 ne servent que si DHCP et WDS sont séparés (et dans ce cas, ne pas cocher « ne pas écouter le port 67 » par erreur — voir §50).

### UEFI vs Legacy en PXE

| | Legacy PXE | UEFI PXE |
|---|---|---|
| Fichier de boot | `wdsnbp.com` | `wdsmgfw.efi` |
| Protocole | TFTP | TFTP (WDS classique) |
| Secure Boot | N/A | Le chargeur PXE doit être signé |

---

