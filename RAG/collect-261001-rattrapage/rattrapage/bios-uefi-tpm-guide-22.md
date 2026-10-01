---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-22
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: ["4011-69-68"]
keywords: ["amd", "intel", "memory", "open source"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [3732, 3836]
sha256: 906571dc420200d840458d3439f49aa422e3ad97cedfa7c03c45ddea8862230d
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```
Modèle : _______________  N° série : _______________
Version BIOS : _______________  Date : _______________
TPM : [ ] discret [ ] fTPM/PTT   Version : _______________
Secure Boot : [ ] ON   Clé BitLocker séquestrée : [ ] AD [ ] Entra [ ] fichier
Mot de passe admin UEFI consigné au coffre : [ ] oui
```

---

## 82. Glossaire

| Terme | Définition |
|---|---|
| **ACPI** | Interface de gestion d'énergie/matériel entre firmware et OS. |
| **AHCI** | Mode du contrôleur SATA (vs IDE/RAID) ; à figer avant l'install. |
| **Attestation** | Preuve cryptographique à distance de l'intégrité du démarrage (TPM). |
| **BCD** | Boot Configuration Data : base de config du démarrage Windows. |
| **BMC** | Contrôleur de gestion de la carte mère (iDRAC, iLO...) : administration hors OS. |
| **Boot Guard** | Technologie Intel : vérification matérielle du firmware avant exécution. |
| **Bootkit** | Malware s'exécutant avant l'OS, dans la chaîne de démarrage. |
| **bcdboot** | Outil de reconstruction des fichiers de démarrage Windows. |
| **bcdedit** | Outil de lecture/écriture du BCD. |
| **CSM** | Compatibility Support Module : émulation BIOS dans un firmware UEFI. |
| **db** | Base des signatures autorisées par Secure Boot. |
| **dbx** | Base des signatures révoquées par Secure Boot. |
| **DMA** | Accès direct à la mémoire par un périphérique (risque si non isolé). |
| **ELAM** | Early Launch Anti-Malware : pilote antivirus chargé en premier au boot. |
| **ESP** | EFI System Partition : partition FAT32 contenant les chargeurs UEFI. |
| **Evil maid** | Attaque avec accès physique temporaire au poste. |
| **Fast Boot** | Initialisation firmware accélérée (saute certains tests). |
| **fTPM** | TPM implémenté en firmware (AMD) ; équivalent Intel : PTT. |
| **GPT** | GUID Partition Table : table de partitions moderne (UEFI). |
| **HVCI** | Hypervisor-protected Code Integrity (Memory Integrity) : protection du noyau. |
| **iPXE** | Firmware PXE avancé open source (HTTP, scripts). |
| **KEK** | Key Exchange Key : signe les mises à jour de db/dbx. |
| **Kernel DMA Protection** | Protection du noyau contre les attaques DMA (Thunderbolt). |
| **MBR** | Master Boot Record : secteur 0 + table de partitions historique. |
| **mbr2gpt** | Outil Microsoft de conversion MBR→GPT sans perte. |
| **Measured Boot** | Enregistrement des mesures de boot dans les PCR du TPM. |
| **MOK** | Machine Owner Key : clés Secure Boot gérées par l'utilisateur (Linux). |
| **MSR** | Partition Microsoft Reserved (16 Mo) sur disque GPT. |
| **NVRAM** | Mémoire non volatile du firmware (variables UEFI, ordre de boot). |
| **Option ROM** | Firmware d'une carte d'extension (réseau, RAID...), vérifié par Secure Boot. |
| **PCR** | Platform Configuration Registers : registres de mesures du TPM. |
| **PK** | Platform Key : clé maîtresse de Secure Boot (propriétaire plateforme). |
| **POST** | Power-On Self-Test : tests matériels au démarrage. |
| **Protective MBR** | Faux MBR protégeant un disque GPT des vieux outils. |
| **PTT** | Intel Platform Trust Technology : TPM firmware d'Intel. |
| **PXE** | Preboot Execution Environment : boot réseau (DHCP+TFTP). |
| **Redfish** | API REST standard d'administration des serveurs (via BMC). |
| **Scellage (sealing)** | Chiffrement de données liées à un état PCR du TPM. |
| **Secure Boot** | Vérification des signatures à chaque maillon du démarrage (UEFI). |
| **Setup Mode** | État sans PK : les clés Secure Boot sont modifiables librement. |
| **Shim** | Petit chargeur Linux signé Microsoft, premier maillon Secure Boot sous Linux. |
| **SMBIOS** | Tables d'inventaire matériel exposées par le firmware (dmidecode). |
| **TPM** | Trusted Platform Module : composant cryptographique, racine de confiance. |
| **UEFI** | Unified Extensible Firmware Interface : firmware moderne remplaçant le BIOS. |
| **VBS** | Virtualization-Based Security : sécurité basée sur l'hyperviseur. |
| **vTPM** | TPM virtuel (VM Hyper-V gen2). |
| **WDS** | Windows Deployment Services : déploiement réseau d'images Windows. |
| **WinPE** | Windows Preinstallation Environment : mini-Windows de déploiement/dépannage. |
| **WinRE** | Windows Recovery Environment : environnement de récupération. |
| **WoL** | Wake-on-LAN : allumage à distance par paquet magique. |

---

## 83. Quiz : 10 questions pour valider

**Q1.** Un disque MBR peut adresser au maximum : a) 128 partitions / b) 4 partitions primaires / c) 2 To par partition ?
> **R1.** **b) 4 partitions primaires** (et le disque entier est limité à ~2 To par l'adressage 32 bits). La limite « 2 To » concerne le disque, pas la partition.

**Q2.** Pourquoi une mise à jour du BIOS déclenche-t-elle la récupération BitLocker ?
> **R2.** BitLocker scelle la clé de chiffrement sur les **PCR du TPM**, dont **PCR 0 = mesure du firmware**. Le nouveau BIOS change cette mesure → le TPM refuse de desceller → BitLocker demande la clé de récupération. Prévention : `Suspend-BitLocker -RebootCount 1` avant le flash.

**Q3.** Citez les 4 bases de clés de Secure Boot et le rôle de chacune.
> **R3.** **PK** (Platform Key : clé maîtresse, signe les MAJ de KEK, définit le propriétaire) ; **KEK** (signe les MAJ de db/dbx) ; **db** (signatures autorisées à booter) ; **dbx** (signatures révoquées/interdites).

**Q4.** Quelle commande convertit un disque système MBR en GPT sans perte de données, et quelle est la condition n°1 avant de la lancer ?
> **R4.** `mbr2gpt /convert /disk:0 /allowFullOS` (après `/validate`). Condition n°1 : **suspendre BitLocker** (et vérifier le séquestre de la clé de récupération), puis basculer le firmware en UEFI après conversion.

**Q5.** `Confirm-SecureBootUEFI` retourne une erreur « cmdlet not supported ». Deux causes possibles ?
> **R5.** La machine démarre en **mode Legacy/BIOS** (pas d'UEFI), ou le firmware ne supporte pas Secure Boot (matériel trop ancien). Dans les deux cas : non conforme Windows 11.

**Q6.** Quelle est la différence entre un TPM discret et un fTPM / Intel PTT ?
> **R6.** Le TPM **discret** est une puce dédiée soudée (isolation matérielle maximale) ; le **fTPM/PTT** est implémenté en logiciel dans le firmware/CPU (gratuit, standard, suffisant pour Windows 11). Les deux sont des TPM 2.0 valides.

**Q7.** Un poste affiche `PXE-E53: No boot filename received`. Citez 3 vérifications dans l'ordre.
> **R7.** 1) Le service **WDSServer** est-il démarré ? 2) Le **pare-feu** laisse-t-il passer UDP 67/68/69/4011 ? 3) Si DHCP et WDS cohabitent : options **« Ne pas écouter le port 67 » + option DHCP 60** activées ? (Puis : stratégie de réponse, fichier `wdsmgfw.efi` présent.)

**Q8.** Pourquoi ne faut-il JAMAIS faire un `Clear-Tpm` « pour voir » sur un poste chiffré ?
> **R8.** Le Clear TPM **détruit définitivement** toutes les clés du TPM, y compris celle qui descelle BitLocker. Sans la **clé de récupération à 48 chiffres** (séquestrée), les données sont perdues. Règle : séquestre vérifié d'abord, Clear ensuite.

**Q9.** Secure Boot est-il une protection contre le vol de données en cas de vol du PC ? Justifiez.
> **R9.** **Non, pas seul.** Secure Boot garantit l'intégrité du démarrage (pas de chargeur pirate), mais **ne chiffre rien** : le disque reste lisible si on le monte ailleurs. Contre le vol : **BitLocker** (chiffrement) + TPM + PIN.

**Q10.** Citez 3 réglages firmware à vérifier après chaque mise à jour du BIOS, et pourquoi.
> **R10.** **Secure Boot** (peut être retombé sur Disabled/Setup), **TPM** (peut être désactivé → BitLocker/Windows 11 en panne), **ordre de boot et mode UEFI/CSM** (un retour en Legacy rend le disque GPT non bootable). Bonus : virtualisation, mot de passe admin.

---

## 84. Pour aller plus loin

### Documentation officielle

