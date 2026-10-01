---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-7
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "attention", "intel"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [805, 954]
sha256: 941f44bc82368f02672834738b0c87f88ab72ec7a8f4582435cd7af94eb1cf8e
---

# BIOS / UEFI — Secure Boot — TPM 2.0

- ❌ L'OS une fois démarré (c'est le rôle de l'antivirus, de HVCI/VBS...)
- ❌ Les attaques avec accès physique + mot de passe firmware inconnu (voir evil maid, section 58)
- ❌ Les vulnérabilités dans un chargeur **légitimement signé** (ex. failles dans des bootloaders révoqués ensuite via dbx)
- ❌ Le vol de données sur disque non chiffré (Secure Boot ≠ chiffrement → coupler avec BitLocker)

### Prérequis

1. Firmware en mode **UEFI natif** (CSM désactivé).
2. Disque système en **GPT**.
3. Clés Secure Boot provisionnées (mode Standard : clés Microsoft d'usine).

---

## 17. La chaîne de confiance du démarrage

```
                    ┌──────────────┐
                    │  Firmware   │── vérifié par la racine matérielle
                    │  (signé OEM)│    (Boot Guard Intel / PSB AMD, si actif)
                    └──────┬───────┘
                           │ Secure Boot vérifie la signature
                    ┌──────▼───────┐
                    │ bootmgfw.efi │── signé Microsoft (clé dans db)
                    │ (Boot Mgr)   │
                    └──────┬───────┘
                           │ vérifie la signature
                    ┌──────▼───────┐
                    │ winload.efi  │── signé Microsoft
                    └──────┬───────┘
                           │ vérifie les signatures (CI = Code Integrity)
              ┌────────────▼────────────┐
              │ ntoskrnl.exe, HAL,      │── signés Microsoft
              │ drivers boot (ELAM)     │
              └─────────────────────────┘
```

### Maillons et signatures (Windows)

| Maillon | Fichier | Signé par | Vérifié par |
|---|---|---|---|
| Boot Manager | `bootmgfw.efi` | Microsoft Windows Production PCA | Firmware (clé MS dans `db`) |
| Chargeur OS | `winload.efi` | Microsoft | bootmgfw.efi |
| Noyau | `ntoskrnl.exe` | Microsoft | winload.efi (Code Integrity) |
| Pilotes kernel | `*.sys` | Microsoft / WHQL | Code Integrity du noyau |

### ELAM (Early Launch Anti-Malware)

Windows charge en premier un pilote anti-malware spécial (**ELAM**) : c'est lui qui valide les autres pilotes de démarrage *avant* qu'ils ne s'exécutent. Votre solution EDR/antivirus d'entreprise fournit ce pilote. Si l'ELAM est absent ou non signé → Secure Boot peut bloquer.

---

## 18. Les clés Secure Boot : PK (Platform Key)

La **PK** est la clé maîtresse : **elle contrôle la plateforme**. Celui qui détient la clé privée de la PK est le « propriétaire » du Secure Boot de la machine.

| Propriété | Détail |
|---|---|
| Rôle | Signer les mises à jour de la **KEK** (et donc indirectement tout le reste) |
| Nombre | **Une seule** par machine |
| Détenteur (mode Standard) | Le **constructeur** (Dell, HP, Lenovo...) |
| Mode Setup | Si aucune PK n'est installée, la machine est en **Setup Mode** : n'importe qui peut écrire les clés (état d'usine avant provisioning) |
| Mode User | PK installée → **User Mode** : les clés sont protégées, modifiables uniquement avec signature PK |

### Opérations liées à la PK

| Action | Effet |
|---|---|
| Installer une PK | Passe en User Mode, verrouille la configuration |
| **Effacer la PK** (`Clear Secure Boot Keys`) | Repasse en **Setup Mode** → Secure Boot **désactivé** de fait |
| Remplacer la PK | Possible en mode Custom (signature avec l'ancienne PK requise) |

> ⚠️ **Ne jamais effacer les clés Secure Boot** (« Clear Secure Boot Keys ») sans raison valable : cela désactive Secure Boot jusqu'à re-provisioning complet. En entreprise, c'est une action à tracer.

---

## 19. Les clés Secure Boot : KEK (Key Exchange Key)

La **KEK** fait le lien entre la PK (propriétaire plateforme) et les bases de signatures (db/dbx).

| Propriété | Détail |
|---|---|
| Rôle | Signer les mises à jour de **`db`** et **`dbx`** |
| Nombre | **Plusieurs** possibles |
| Détenteurs typiques (mode Standard) | **Microsoft** (`Microsoft Corporation KEK CA 2011`) + le **constructeur** (KEK OEM) |
| Mise à jour | Via Windows Update / capsule firmware, signée par la PK |

### Hiérarchie complète

```
PK (constructeur)
 └─ signe les mises à jour de la KEK
     KEK (Microsoft + OEM)
      └─ signe les mises à jour de db et dbx
          db  : chargeurs autorisés  (ex. "Microsoft Windows Production PCA 2011",
                                      "Microsoft Corporation UEFI CA 2011")
          dbx : chargeurs révoqués   (hashes interdits)
```

Concrètement : quand Microsoft révoque un chargeur vulnérable, il publie une mise à jour de la **dbx** signée par sa KEK → le firmware l'accepte car il fait confiance à cette KEK.

---

## 20. Les clés Secure Boot : db (base de signatures autorisées)

La **`db`** (*authorized signature database*) contient les certificats et hashes **autorisés** à démarrer.

### Contenu typique en mode Standard

| Entrée | Usage |
|---|---|
| `Microsoft Windows Production PCA 2011` | Signe `bootmgfw.efi`, `winload.efi` (Windows 8 → 11) |
| `Microsoft Corporation UEFI CA 2011` | Signe le **shim** Linux (Canonical, Red Hat, SUSE...) |
| `Microsoft Option ROM UEFI CA` | Signe les ROMs d'extension (cartes réseau, RAID) |
| Certificats OEM | Chargeurs/outils du constructeur (diagnostics Dell/HP/Lenovo) |

### Comment le firmware vérifie

1. Le firmware calcule le hash du binaire `.efi` à exécuter.
2. Il vérifie la **signature** du binaire avec les certificats de `db`.
3. Si la signature est valide **et** que le hash n'est pas dans `dbx` → exécution autorisée.
4. Sinon → **refus de démarrage** (écran rouge / message « Secure Boot Violation »).

---

## 21. Les clés Secure Boot : dbx (base de révocation)

La **`dbx`** (*forbidden signature database*) contient les certificats et hashes **interdits**, même s'ils sont signés par une clé de `db`.

### À quoi ça sert

Révoquer des chargeurs **légitimement signés mais vulnérables**. Exemple historique : des failles dans le chargeur GRUB ou dans des utilitaires de boot signés Microsoft ont permis de contourner Secure Boot → Microsoft a publié leurs hashes en `dbx` via Windows Update.

### Points d'attention en entreprise

- Les mises à jour `dbx` arrivent via **Windows Update** (KB spécifiques) ou les **mises à jour firmware** du constructeur.
- Une `dbx` obsolète = des chargeurs vulnérables toujours acceptés → **vérifier que les mises à jour `dbx` sont bien appliquées sur le parc** (surtout après les correctifs majeurs).
- ⚠️ **Effet de bord connu :** une mise à jour `dbx` trop agressive peut bloquer un **double-boot Linux** dont le shim a été révoqué → prévoir une procédure (section 74).

### L'échéance des certificats 2026

Les certificats Secure Boot Microsoft de 2011 (`...2011`) arrivent à expiration : Microsoft a publié de nouveaux certificats (`...2023`) et pousse leur déploiement via Windows Update et les firmwares OEM. **En entreprise :**

- [ ] Vérifier que les postes reçoivent les mises à jour de clés (Windows Update actif ou WSUS avec la catégorie).
- [ ] Sur les images master (MDT/SCCM), intégrer un firmware récent incluant les nouvelles clés.
- [ ] Tester le déploiement sur un lot pilote avant généralisation (un raté = parc qui ne boote plus).

---

## 22. Mode Standard vs mode Custom

