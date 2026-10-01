---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-9
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "attention", "intel"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [1160, 1333]
sha256: 0257849fb17ae4c6ef64157555f0d830e3deb14a544ecf9ea92de5d5f56d2543
---

# BIOS / UEFI — Secure Boot — TPM 2.0

| Message / écran | Cause probable |
|---|---|
| `Secure Boot Violation — Invalid signature detected` | Chargeur non signé ou signature inconnue |
| Écran rouge du firmware (certains OEM) | Échec de vérification |
| Boot direct sur le setup / « No bootable device » | Entrée de boot pointant vers un `.efi` refusé |
| Windows démarre, puis écran bleu `0xc0000428` | Pilote non signé bloqué par Code Integrity |

### Arbre de décision

```
Secure Boot bloque le démarrage
│
├─► C'est un OS/outil que JE veux booter (Linux, utilitaire, WinPE maison) ?
│   ├─ OUI → 1) Vérifier qu'il existe une version signée (shim, WinPE signé).
│   │         2) Sinon : signer le binaire (mode Custom) OU désactiver
│   │            temporairement Secure Boot (tracer l'action !).
│   └─ NON (c'est Windows qui ne boote plus) → continuer
│
├─► Changement récent ? (MAJ firmware, dbx, nouveau disque, clonage)
│   ├─ OUI → Restaurer l'état antérieur / réparer le boot (bcdboot, §15).
│   └─ NON → continuer
│
├─► Vérifier l'intégrité : Secure Boot est-il en "Deployed Mode" avec les
│   clés d'usine ? (setup → Secure Boot → Key Management)
│   ├─ Clés effacées / Setup Mode → restaurer les clés d'usine.
│   └─ Clés OK → continuer
│
└─► Piste malware/bootkit ? (chargeur modifié)
    → Analyser depuis un support sain, comparer les hashes des .efi
      avec une machine de référence.
```

### Commandes de diagnostic

```powershell
# Depuis WinRE : vérifier que les fichiers de boot existent et sont signés
# (comparer avec une machine saine)
dir S:\EFI\Microsoft\Boot\bootmgfw.efi
dir S:\EFI\Boot\bootx64.efi

# Vérifier l'état Secure Boot vu par Windows (si Windows démarre encore)
Confirm-SecureBootUEFI
Get-ItemProperty 'HKLM:\SYSTEM\CurrentControlSet\Control\SecureBoot\State'
```

> 🔒 **Règle d'or :** désactiver Secure Boot pour « faire passer » un problème, c'est comme enlever la porte parce que la serrure coince. Autorisé en dépannage, **toujours temporaire**, toujours tracé, toujours réactivé après.

---

## 28. TPM 2.0 : principe et racine de confiance matérielle

Le **TPM** (*Trusted Platform Module*) est un composant cryptographique matériel (puce dédiée ou firmware) qui fournit une **racine de confiance** : un point d'ancrage inviolable pour la sécurité du système.

### Ce que fait un TPM

| Fonction | Détail |
|---|---|
| **Génération de clés** | Crée des clés RSA/ECC **à l'intérieur** de la puce ; la clé privée ne sort jamais. |
| **Stockage protégé** | Clés, certificats, secrets scellés dans la mémoire interne. |
| **Chiffrement / signature** | Opérations RSA, ECC, SHA-1/SHA-256, HMAC réalisées par la puce. |
| **Mesures d'intégrité (PCR)** | Enregistre les hashes des composants de démarrage (voir section 35). |
| **Scellage (sealing)** | Chiffre des données liées à un état PCR : elles ne se déchiffrent que si le boot est identique. |
| **Générateur aléatoire** | Source d'aléa matérielle (TRNG) pour les clés. |
| **Compteurs monotones** | Compteurs anti-rejeu (anti-rollback). |

### Hiérarchie des clés TPM 2.0

```
Endorsement Key (EK)      ← clé RSA/ECC unique, gravée à la fabrication,
                            certifiée par le fabricant. Identité du TPM.
   └─ Storage Root Key (SRK)
         └─ Clés applicatives : BitLocker, Windows Hello, certificats,
            clés d'attestation, Credential Guard...
```

### À quoi sert concrètement le TPM en entreprise

- **BitLocker** : la clé de chiffrement du disque est scellée au TPM (+ PCR) → le disque est illisible sur une autre machine.
- **Windows Hello for Business** : clés d'authentification protégées par le TPM.
- **Measured Boot / attestation** : prouver à distance qu'une machine a démarré sainement (scénarios Zero Trust, Conditional Access).
- **Credential Guard / HVCI** : isolation des secrets basée sur la virtualisation + TPM.
- **VPN / Wi-Fi EAP-TLS** : certificats machine stockés dans le TPM.

> 💡 **En résumé :** le TPM est le coffre-fort matériel du PC. Sans lui, BitLocker n'est qu'un cadenas logiciel ; avec lui, la clé de chiffrement est liée au matériel *et* à l'intégrité du démarrage.

---

## 29. TPM 1.2 vs TPM 2.0

| Critère | TPM 1.2 | TPM 2.0 |
|---|---|---|
| Standard | TCG, 2011 | TCG, 2014 (ISO/IEC 11889) |
| Algorithmes | RSA, SHA-1 **imposés** | **Agile** : RSA, ECC, SHA-256, etc. (négociables) |
| Hiérarchies de clés | 1 (Storage) | 3 (Endorsement, Storage, Platform) |
| PCR | 24, SHA-1 uniquement | 24, **banques SHA-1 + SHA-256** |
| Autorisation | HMAC simple | Politiques d'autorisation riches |
| Support Windows | Jusqu'à Windows 10 | **Requis pour Windows 11** |
| Support constructeurs | Abandonné | Standard depuis ~2018 |

### Migration 1.2 → 2.0

Certains TPM (notamment Infineon sur des Dell/Lenovo 2016-2018) permettent une **conversion firmware 1.2 ↔ 2.0** via un utilitaire du constructeur. Points d'attention :

```
1. La conversion EFFACE le TPM (toutes les clés perdues).
2. Suspendre/désactiver BitLocker AVANT (sinon : demande de clé de récupération).
3. Sauvegarder les clés de récupération BitLocker (AD/Entra ID, §38).
4. Après conversion : réactiver, re-provisionner (Initialize-Tpm), réactiver BitLocker.
```

> ⚠️ Sur du matériel récent, le TPM est nativement 2.0 : rien à convertir. La migration 1.2→2.0 ne concerne que le parc 2016-2019.

---

## 30. TPM discret vs TPM firmware (fTPM / Intel PTT)

| | TPM discret (dTPM) | TPM firmware (fTPM / Intel PTT) |
|---|---|---|
| Implémentation | **Puce dédiée** soudée (ex. Infineon, STMicro) | Code s'exécutant dans le **firmware** (AMD fTPM / Intel PTT) |
| Sécurité | Isolation matérielle maximale | Isolation via extensions CPU (moins « dur » qu'une puce) |
| Coût | + composant | Gratuit (inclus au CPU/chipset) |
| Activation | Souvent activé par défaut (pro) | Parfois **désactivé par défaut** → à activer dans le setup |
| Vulnérabilités connues | Failles Infineon RSA (2017, CVE-2017-15361) → MAJ firmware | Failles de contournement sur certaines plateformes (faible impact pratique) |
| Cas d'usage | Postes sensibles, serveurs | **Standard sur le parc bureautique** |

### Identifier le type de TPM

```powershell
Get-Tpm | Select-Object TpmPresent, TpmReady, ManufacturerId, ManufacturerVersion

# ManufacturerId (hex → ASCII) :
#  0x49465800 = "IFX" (Infineon, souvent discret)
#  0x414D4400 = "AMD" (souvent fTPM)
#  0x494E5443 = "INTC" (Intel PTT)
```

```powershell
# Décoder le ManufacturerId en lettres
$tpm = Get-Tpm
$bytes = [BitConverter]::GetBytes($tpm.ManufacturerId)
[System.Text.Encoding]::ASCII.GetString($bytes).Trim([char]0)
```

### Recommandation

- **Parc standard :** le fTPM/PTT suffit (c'est ce que Microsoft valide pour Windows 11).
- **Postes sensibles** (direction, finance, R&D) : préférer le TPM discret si l'option existe.
- Dans tous les cas : **mettre à jour le firmware du TPM** (section 56).

---

## 31. Activer le TPM dans le firmware

### Chemins typiques par constructeur

| Constructeur | Chemin dans le setup | Nom de l'option |
|---|---|---|
| **Dell** | Security → TPM 2.0 Security | `TPM On`, `TPM 2.0 Security On` |
| **HP** | Security → TPM Embedded Security | `TPM Device` = Available, `TPM State` = On |
| **Lenovo** | Security → Security Chip | `Security Chip` = Enabled (Discrete / Intel PTT / AMD fTPM) |

### Procédure générique

```
1. Entrer dans le setup UEFI.
2. Security → activer le TPM (On / Enabled / Available).
3. Vérifier la version : TPM 2.0 (si choix 1.2/2.0 → 2.0).
4. Sauvegarder et redémarrer.
5. Dans Windows : Get-Tpm → TpmPresent = True, TpmReady = True.
   Si TpmReady = False : initialiser (Initialize-Tpm ou tpm.msc → "Préparer le TPM").
```

### Activation à distance / en masse

