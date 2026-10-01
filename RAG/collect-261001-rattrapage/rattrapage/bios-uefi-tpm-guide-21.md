---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-21
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [3565, 3731]
sha256: 7303dce2ff1194b31255288d0f3fbc1dbfe03b918ba14ef88b6926867df9d125
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```
→ Carte mère à remplacer (ou reprogrammation de la puce en atelier
  avec un programmateur SPI type CH341A — opération de laboratoire).
→ C'est pour cela qu'on ne flashe JAMAIS sans onduleur/secteur stable
  et sans raison valable (§52).
```

---

## 78. Cas pratique 15 : double boot Windows/Linux cassé après une mise à jour Windows

### Contexte

Après une mise à jour majeure de Windows, le menu GRUB a disparu : le PC démarre directement sur Windows. Ou l'inverse : GRUB s'affiche mais Windows ne démarre plus.

### Analyse

```
Windows a réécrit l'entrée de boot par défaut / restauré son chargeur
comme premier dans BootOrder, ou régénéré le BCD.
→ Les fichiers Linux (shim, GRUB) sont toujours sur l'ESP : rien n'est
  perdu, c'est l'ORDRE qui a changé.
```

### Résolution A : GRUB a disparu (boot direct Windows)

```powershell
# Depuis Windows (admin) : remettre le chargeur Linux en premier
# 1. Voir les entrées firmware
bcdedit /enum firmware

# 2. Identifier l'entrée Linux (ex. {xxxxxxxx-...} "ubuntu")
# 3. La passer en premier :
bcdedit /set '{fwbootmgr}' displayorder '{GUID-ubuntu}' /addfirst

# Alternative : au démarrage, touche F12 (menu de boot) → choisir ubuntu
# → puis dans Linux : sudo efibootmgr -o XXXX,YYYY (remettre l'ordre).
```

### Résolution B : GRUB s'affiche mais Windows ne boote plus

```bash
# Depuis Linux : réparer l'entrée Windows
sudo os-prober          # doit détecter Windows
sudo update-grub        # régénère grub.cfg avec Windows
# Si os-prober ne voit rien : vérifier que l'ESP est montée et que
# /EFI/Microsoft/Boot/bootmgfw.efi existe toujours.
```

```powershell
# Depuis Windows (si GRUB est le problème, pas Windows) : restaurer
# le chargeur Windows comme défaut sans effacer GRUB :
bcdboot C:\Windows /s S: /f UEFI
bcdedit /set '{fwbootmgr}' displayorder '{bootmgr}' /addfirst
```

### Prévention

```
→ Après chaque mise à jour majeure Windows sur un double-boot :
  vérifier l'ordre de boot (bcdedit /enum firmware).
→ Garder une clé USB de réparation (boot-repair ou WinPE) à portée.
→ En entreprise : éviter le double-boot sur les postes de production ;
  préférer une VM ou un poste dédié.
```

---

## 79. Erreurs classiques : les 18 pièges à éviter

| N° | Erreur | Conséquence | Bonne pratique |
|---|---|---|---|
| 1 | Flasher le BIOS sans suspendre BitLocker | Clé de récupération demandée (panique) | `Suspend-BitLocker -RebootCount 1` avant |
| 2 | `Clear-Tpm` sans séquestre de la clé BitLocker | **Données perdues** | Vérifier + sauvegarder la clé avant |
| 3 | Effacer les clés Secure Boot « pour tester » | Secure Boot désactivé, machine en Setup Mode | Ne jamais effacer sans plan de re-provisioning |
| 4 | Passer en Legacy/CSM « parce que ça bootait avant » | Windows 11 impossible, Secure Boot perdu | Rester en UEFI ; convertir avec mbr2gpt |
| 5 | Formater l'ESP « pour faire du propre » | PC non bootable | ESP intouchable ; `bcdboot` si besoin |
| 6 | Initialiser un disque GPT en MBR dans le gestionnaire de disques | Perte de la table GPT | Vérifier le style AVANT toute action |
| 7 | Désactiver Secure Boot définitivement pour un outil | Porte ouverte aux bootkits | Désactivation temporaire, tracée, réactivée |
| 8 | Mot de passe BIOS posé sans le consigner | Procédure constructeur lourde (§69) | Coffre d'équipe dès la réception |
| 9 | MAJ firmware TPM sans précaution | Équivaut à un Clear TPM → BitLocker bloqué | Même protocole que Clear TPM (§56) |
| 10 | Options DHCP 66/67 + WDS sur le même réseau | Conflit, PXE-E53 | Option 60 du WDS, pas 66/67 (§50) |
| 11 | Oublier les pilotes NIC/NVMe dans boot.wim | WinPE sans réseau/disque | Injecter les pilotes (§49) |
| 12 | `bcdedit` sans export préalable | BCD cassé sans filet | `bcdedit /export` systématique |
| 13 | Laisser `testsigning Yes` après un dépannage | Pilotes non signés acceptés durablement | Le désactiver + vérifier le filigrane |
| 14 | Contourner les exigences Windows 11 (BypassTPMCheck) | Parc non supporté, non patché | Remplacer le matériel non conforme |
| 15 | WoL qui « ne marche pas » | Démarrage rapide Windows activé | Désactiver le fast startup + config carte (§45) |
| 16 | Clé USB d'install créée en MBR pour un poste UEFI | Setup qui ne démarre pas en UEFI | Rufus en GPT/UEFI, ou Media Creation Tool |
| 17 | Ne pas re-vérifier les réglages après un flash | TPM/Secure Boot/VT-x retombés par défaut | Check-list post-flash (§52) |
| 18 | Stocker la clé BitLocker uniquement sur le poste | Clé perdue avec le poste | Séquestre AD/Entra ID systématique (§38) |

---

## 80. Checklist : configuration d'un poste neuf en entreprise

### À la réception (atelier)

- [ ] Vérifier le modèle, le numéro de série, la garantie (enregistrer dans l'inventaire)
- [ ] Flasher le dernier BIOS **validé** (matrice §60)
- [ ] Restaurer les réglages par défaut, puis appliquer le **profil standard** :
  - [ ] Mode UEFI natif (CSM désactivé)
  - [ ] Secure Boot → Enabled (Standard)
  - [ ] TPM 2.0 → On, provisionné (`Get-Tpm` → TpmReady)
  - [ ] Virtualisation → Enabled (VT-x/AMD-V + VT-d/AMD-Vi)
  - [ ] Ordre de boot → disque en premier
  - [ ] USB boot → Disabled (sauf exception)
  - [ ] PXE → selon politique (Enabled si déploiement réseau)
  - [ ] WoL → selon politique
  - [ ] Mot de passe admin UEFI → défini, consigné au coffre
- [ ] Vérifier le disque en GPT (ou convertir si image Legacy)

### Déploiement de l'OS

- [ ] Déployer l'image master (WDS/MDT/SCCM/Intune)
- [ ] Vérifier : `Confirm-SecureBootUEFI` = True, `Get-Tpm` OK
- [ ] Activer BitLocker (TPM ou TPM+PIN selon profil), **vérifier le séquestre**
- [ ] Jonction au domaine / Entra ID, GPO/Intune appliquées
- [ ] Inventaire : ajouter la machine au CSV (§40)

### Avant remise à l'utilisateur

- [ ] Test de redémarrage à froid + sortie de veille
- [ ] Vérifier qu'aucun mot de passe ne bloque l'utilisateur (PIN BitLocker communiqué si TPM+PIN)
- [ ] Fiche de remise : consigner le modèle, le n° de série, la clé BitLocker séquestrée

---

## 81. Pense-bête de poche

### Les 10 commandes à connaître par cœur

```powershell
Confirm-SecureBootUEFI                                  # Secure Boot ? True/False
Get-Tpm | Select-Object TpmPresent,TpmReady             # TPM OK ?
Get-Disk | Select-Object Number,PartitionStyle,Size     # GPT ou MBR ?
bcdedit /enum                                           # Config de boot
bcdedit /export C:\BCD-backup                          # Sauvegarder le BCD
bcdboot C:\Windows /s S: /f UEFI                        # Réparer le boot UEFI
mbr2gpt /validate /disk:0 /allowFullOS                  # Tester la conversion
mbr2gpt /convert /disk:0 /allowFullOS                   # Convertir en GPT
manage-bde -status C:                                   # État BitLocker
Suspend-BitLocker -MountPoint "C:" -RebootCount 1       # Avant flash BIOS
```

### Touches d'accès au firmware (indicatif)

| Constructeur | Setup | Menu de boot |
|---|---|---|
| Dell | F2 | F12 |
| HP | F10 (ou Échap puis F10) | F9 (ou Échap puis F9) |
| Lenovo ThinkPad | Entrée puis F1 (ou F2) | F12 |
| ASUS | DEL ou F2 | F8 |
| MSI | DEL | F11 |
| Gigabyte | DEL | F12 |

### Les 5 réflexes

```
1. AVANT de toucher au BCD/firmware : exporter/sauvegarder.
2. AVANT un flash : BitLocker suspendu + séquestre vérifié + secteur.
3. DEVANT un disque "non alloué" : vérifier le MODE FIRMWARE d'abord.
4. DEVANT une demande BitLocker "spontanée" : chercher la CAUSE (MAJ ? intrusion ?).
5. Secure Boot désactivé = TEMPORAIRE + TRACÉ + RÉACTIVÉ.
```

### Numéros utiles à noter (par poste)

