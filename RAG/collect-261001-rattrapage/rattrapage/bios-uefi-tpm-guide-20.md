---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-20
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft", "Nvidia"]
dates: []
keywords: ["distribution", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [3340, 3564]
sha256: ecc90ddea05d7e49e590ee2a28137824f5da1a1b1421ee7f0ee853116ee2d3d2
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```
Rappel (§4) : un BIOS Legacy ne comprend PAS le GPT.
→ Il lit le Protective MBR (qui dit "disque plein, ne touche pas"),
  ne trouve pas de MBR bootable valide → "No bootable device".
→ Certains vieux outils affichent alors le disque comme "non alloué"
  ou proposent de l'initialiser. C'EST UN AFFICHAGE, pas une perte.
```

### Résolution (ne rien formater !)

```
1. NE PAS cliquer sur "Initialiser le disque" (cela ÉCRASERAIT le GPT).
2. Revenir dans le setup UEFI → repasser en mode UEFI (CSM OFF).
3. Le disque GPT réapparaît, Windows redémarre normalement.
4. Si Windows ne redémarre pas : vérifier l'ordre de boot
   (Windows Boot Manager en premier), puis bcdboot si besoin (§15).
```

### Vérification de l'intégrité (si un outil a touché au disque)

```powershell
# Depuis Windows (une fois rebooté en UEFI) :
Get-Disk -Number 0 | Select-Object PartitionStyle, OperationalStatus
Get-Partition -DiskNumber 0 | Select-Object PartitionNumber, Type, Size

# Si une partition a été endommagée : TestDisk (outil libre) peut
# reconstruire la table GPT grâce à sa copie de secours en fin de disque.
```

> 🔴 **Réflexe :** devant un disque « non alloué » après un changement de mode firmware, la cause est à 99 % le **mode firmware**, pas le disque. On ne formate jamais avant d'avoir remis le bon mode.

---

## 74. Cas pratique 11 : Secure Boot bloque une distribution Linux

### Contexte

Installation d'Ubuntu/Debian sur un poste avec Secure Boot activé : le live USB ne démarre pas, ou après installation, `shim` est rejeté.

### Diagnostic

```
1. La distribution a-t-elle un shim signé Microsoft ?
   ✅ Ubuntu, Fedora, RHEL, SUSE, Debian (Bullseye+) : oui.
   ❌ Arch (de base), Gentoo, Kali (selon version), customs : non.
2. Le live USB a-t-il été créé en mode UEFI ? (Rufus : "GPT / UEFI",
   pas "MBR / BIOS".)
3. Message exact : "Secure Boot Violation" → signature refusée ;
   "no bootable device" → clé non bootable en UEFI.
```

### Résolutions

```
Option A : utiliser une distribution avec shim signé (recommandé).
Option B : créer la clé USB en UEFI/GPT (souvent la vraie cause).
Option C : enrôler une clé MOK pour les modules propriétaires
  (ex. pilote NVIDIA/VirtualBox) :
    mokutil --import /var/lib/shim-signed/mok/MOK.der
  → redémarrer → l'écran bleu MOK Manager demande le mot de passe
    défini → Enroll MOK → le module est accepté.
Option D (temporaire) : désactiver Secure Boot pour l'installation,
  puis le réactiver après avoir installé le shim signé.
```

### Le cas de la dbx qui révoque un vieux shim

```
Après une mise à jour dbx Windows, un Linux en double-boot ne démarre
plus (shim révoqué car vulnérable).
→ Mettre à jour le Linux (le shim récent est signé avec un certificat
  non révoqué) : sudo apt update && sudo apt install shim-signed grub-efi.
→ Ne PAS bloquer les mises à jour dbx pour "protéger" le double-boot :
  c'est le shim qu'il faut mettre à jour.
```

---

## 75. Cas pratique 12 : le TPM se désactive après une mise à jour du BIOS

### Contexte

Après un flash du BIOS, `Get-Tpm` indique `TpmPresent = False` ou `TpmReady = False`. BitLocker menace de demander la clé.

### Pourquoi ça arrive

```
- Le nouveau firmware a réinitialisé les réglages → TPM repassé sur Off.
- Le firmware a changé le mode (Discrete ↔ fTPM) → Windows voit
  un "nouveau" TPM.
- La MAJ incluait une MAJ du firmware TPM → équivalent d'un Clear.
```

### Résolution

```powershell
# 1. D'abord : sécuriser BitLocker (ne pas rebooter en boucle sans clé !)
manage-bde -protectors -get C: -type RecoveryPassword
Suspend-BitLocker -MountPoint "C:" -RebootCount 1

# 2. Setup UEFI → Security → réactiver le TPM (même mode qu'avant : fTPM/PTT
#    ou Discrete — changer de mode = nouveau TPM = re-provisioning).

# 3. Dans Windows :
Get-Tpm | Select-Object TpmPresent, TpmReady
# Si TpmReady = False :
Initialize-Tpm
# → redémarrer si demandé (RestartPending).

# 4. Re-scellé BitLocker :
Resume-BitLocker -MountPoint "C:"
Get-BitLockerVolume -MountPoint "C:" | Select-Object VolumeStatus
```

### Prévention

```
→ Toujours noter les réglages TPM AVANT un flash (photo du setup ou
  export cctk/BiosConfigUtility).
→ Après tout flash : check-list de re-vérification (§52) incluant le TPM.
```

---

## 76. Cas pratique 13 : WDS ne répond pas aux clients d'un autre sous-réseau (VLAN)

### Contexte

Le PXE fonctionne sur le VLAN du serveur WDS, mais les clients d'un autre VLAN obtiennent `PXE-E51: No DHCP offers` ou `PXE-M0F`.

### Analyse

```
Le PXE démarre par un broadcast DHCP (Discover).
→ Les broadcasts ne traversent pas les routeurs.
→ Sans relais, le client n'a ni IP ni "boot filename".
```

### Résolution : relais DHCP (ip helper)

```
Sur le routeur / switch L3, sur l'interface du VLAN client :

Cisco IOS :
  interface Vlan20
    ip helper-address 10.1.0.10    ← serveur DHCP
    ip helper-address 10.1.0.20    ← serveur WDS (si différent du DHCP)

HP/Aruba (similaire), MikroTik :
  /ip dhcp-relay add interface=vlan20 dhcp-server=10.1.0.10,10.1.0.20

⚠️ Si DHCP et WDS sont sur la MÊME machine : un seul helper suffit,
   mais le WDS doit avoir l'option "Ne pas écouter le port 67" + option 60
   (§50), sinon conflit DHCP/WDS.
```

### Alternative : options DHCP 66/67 sur le DHCP du site distant

```
Si un DHCP local existe sur le VLAN distant :
  option 66 = IP du serveur WDS
  option 67 = boot\x64\wdsmgfw.efi   (UEFI)
             boot\x64\wdsnbp.com     (Legacy)
→ Le client obtient IP + serveur de boot en une seule réponse DHCP.
```

### Vérification

```powershell
# Depuis un client du VLAN distant (en OS) : le DHCP répond-il ?
ipconfig /all   # serveur DHCP visible ?

# Côté WDS : voit-on la requête arriver ?
Get-WinEvent -LogName 'Microsoft-Windows-Deployment-Services-Diagnostics/Operational' `
    -MaxEvents 10 | Select-Object TimeCreated, Message | Format-Table -Wrap
```

---

## 77. Cas pratique 14 : écran noir après flash du BIOS — procédure de récupération

### Contexte

Pendant ou après un flash : écran noir, ventilateurs qui tournent, aucun affichage, parfois des bips ou des LED qui clignotent. **Ne pas paniquer, ne pas multiplier les tentatives.**

### Premiers gestes (5 minutes)

```
1. Attendre 10 minutes : certains flashs font plusieurs redémarrages
   avec de longs écrans noirs. NORMAL.
2. Forcer l'extinction (10 s), débrancher 2 minutes, rebrancher, rallumer.
3. Tester un écran externe (portable) / une autre sortie vidéo (tour).
4. Écouter les bips, noter les codes LED (voir manuel constructeur).
```

### Récupération par constructeur

```
DELL — BIOS Recovery :
  1. Éteindre. Maintenir CTRL + ÉCHAP, brancher le secteur (toujours
     en maintenant), puis allumer.
  2. L'écran "BIOS Recovery" propose de restaurer depuis le disque
     ou une clé USB (fichier .rcv/.hdr du support Dell).

HP — Sure Start / Win+B :
  1. Éteindre. Maintenir WIN + B, brancher le secteur, allumer
     en maintenant 5-10 s.
  2. L'utilitaire HP BIOS Update cherche un BIOS (USB ou en ligne).
  3. Sur les modèles Sure Start récents : la restauration est AUTOMATIQUE
     (copie dorée) → le PC redémarre seul après quelques minutes.

LENOVO — Emergency Reset + USB :
  1. Trou "reset" (trombone) 10 s, ou bouton Novo.
  2. Sur ThinkPad : clé USB avec le BIOS renommé selon la doc du modèle,
     Fn+R au démarrage (selon génération — vérifier le manuel !).

GÉNÉRIQUE (carte mère tour) :
  - Fonction "BIOS Flashback" (ASUS, MSI, Gigabyte...) : clé USB avec
    le fichier renommé (ex. MSI.ROM) sur le port dédié, bouton Flashback
    5 s, PC éteint mais branché. La LED clignote pendant le flash.
  - Double BIOS (Gigabyte DualBIOS) : bascule automatique ou par jumper.
```

### Si rien ne fonctionne

