---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-23
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "intel"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [3837, 3877]
sha256: bec1d9b8712a07d1404ee95cc9e96d1ea0ba9322bd85cfc4f2f7eec421f812fd
---

# BIOS / UEFI — Secure Boot — TPM 2.0

- **UEFI Forum** — Spécification UEFI (chapitre 27 : Secure Boot) : uefi.org/specifications
- **TCG** — Spécifications TPM 2.0 (Trusted Computing Group) : trustedcomputinggroup.org
- **Microsoft Learn** — Secure Boot, TPM, BitLocker, WDS : learn.microsoft.com
- **NIST SP 800-147** — Protection du BIOS ; **NIST SP 800-155** — Intégrité du firmware

### Outils à maîtriser

| Outil | Usage |
|---|---|
| `Dell Command Configure` / `Update` | Gestion BIOS Dell en masse |
| `HP BIOS Configuration Utility` / `Image Assistant` | Gestion BIOS HP en masse |
| `Lenovo System Update` / Thin Installer | Gestion BIOS Lenovo en masse |
| **MDT** (Microsoft Deployment Toolkit) | Industrialiser le déploiement au-dessus de WDS |
| **Rufus** | Créer des clés USB UEFI/GPT propres |
| **TestDisk** | Reconstruire une table GPT endommagée |
| `mokutil` / `efibootmgr` | Gérer Secure Boot et les entrées UEFI sous Linux |

### Sujets connexes à creuser

- **Microsoft Pluton** : le processeur de sécurité intégré aux CPU récents (au-delà du TPM).
- **DRTM** (Dynamic Root of Trust for Measurement) : Intel TXT / AMD SKINIT — mesurer *après* le firmware.
- **Supply chain firmware** : attaques de la chaîne d'approvisionnement (firmwares pré-infectés).
- **Confidential Computing** : AMD SEV-SNP, Intel TDX — chiffrer la mémoire des VM dans le cloud.
- **Passkeys / FIDO2** : le TPM comme ancrage des clés d'authentification sans mot de passe.

### Feuille de route suggérée pour votre parc (6 mois)

```
Mois 1-2 : inventaire complet (§40), matrice modèles/firmwares, coffre
           mots de passe, séquestre BitLocker vérifié à 100 %.
Mois 3    : vague pilote MAJ firmware + activation TPM/Secure Boot (10 %).
Mois 4    : généralisation par vagues, conversion mbr2gpt des Legacy.
Mois 5    : durcissement (profils BIOS standard, USB boot, check-list §59),
           migration Windows 11 des postes conformes.
Mois 6    : audit final, documentation, plan de maintien (trimestriel).
```

---

*Fin du guide — BIOS / UEFI / Secure Boot / TPM 2.0 en entreprise.*
*Document de travail : adaptez les chemins et options à vos modèles exacts, et testez toute procédure sur un lot pilote avant généralisation.*
