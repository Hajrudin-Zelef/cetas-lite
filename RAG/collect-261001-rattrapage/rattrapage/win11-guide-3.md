---
id: collect-261001-rattrapage/rattrapage/win11-guide-3
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Qualcomm"]
dates: []
keywords: ["amd", "consumer", "copilot", "intel", "packaging"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [73, 208]
sha256: e025613b5c1637156b580dc13ee63eb90e699b733e80178659d4d8e91b303163
---

# Windows 11 en entreprise — Guide technique ultra-complet

61. [Scripts utiles : boîte à outils PowerShell de l'admin](#61-scripts-utiles--boîte-à-outils-powershell-de-ladmin)
62. [Sécurité avancée : Secure Boot, measured boot, attestation](#62-sécurité-avancée--secure-boot-measured-boot-attestation)
63. [Chiffrement avancé : BitLocker réseau (Network Unlock), DRA](#63-chiffrement-avancé--bitlocker-réseau-network-unlock-dra)
64. [Conformité : stratégies de conformité Intune, base de référence](#64-conformité--stratégies-de-conformité-intune-base-de-référence)
65. [Kiosque et postes partagés : Assigned Access, invité](#65-kiosque-et-postes-partagés--assigned-access-invité)
66. [Impression en entreprise : Universal Print, Point and Print](#66-impression-en-entreprise--universal-print-point-and-print)
67. [Langues, régions, claviers : déploiement multilingue](#67-langues-régions-claviers--déploiement-multilingue)
68. [Télémétrie et confidentialité : niveaux, GPO, registre](#68-télémétrie-et-confidentialité--niveaux-gpo-registre)
69. [Microsoft 365 Apps : déploiement, canaux, licences](#69-microsoft-365-apps--déploiement-canaux-licences)
70. [Navigateurs en entreprise : Edge géré, stratégies](#70-navigateurs-en-entreprise--edge-géré-stratégies)
71. [Sauvegarde et PRA poste de travail](#71-sauvegarde-et-pra-poste-de-travail)
72. [Gestion de l'énergie : stratégies, veille moderne, batteries](#72-gestion-de-lénergie--stratégies-veille-moderne-batteries)
73. [Erreurs classiques (20) : ce qu'il ne faut pas faire](#73-erreurs-classiques-20--ce-quil-ne-faut-pas-faire)
74. [Checklist : déploiement d'un poste de A à Z](#74-checklist--déploiement-dun-poste-de-a-à-z)
75. [Checklist : audit de sécurité d'un poste Windows 11](#75-checklist--audit-de-sécurité-dun-poste-windows-11)
76. [Pense-bête de poche : commandes et raccourcis essentiels](#76-pense-bête-de-poche--commandes-et-raccourcis-essentiels)
77. [Glossaire](#77-glossaire)
78. [Quiz : 10 questions + réponses commentées](#78-quiz--10-questions--réponses-commentées)
79. [Pour aller plus loin : documentation, outils, communautés](#79-pour-aller-plus-loin--documentation-outils-communautés)
80. [Annexe A : tableau des builds et versions](#80-annexe-a--tableau-des-builds-et-versions)
81. [Annexe B : chemins de clés de registre les plus utiles](#81-annexe-b--chemins-de-clés-de-registre-les-plus-utiles)
82. [Annexe C : modèle de plan de migration Windows 10 → 11](#82-annexe-c--modèle-de-plan-de-migration-windows-10--11)

---

## 1. Vue d'ensemble et positionnement de Windows 11

Windows 11 est la version cliente actuelle de Windows pour l'entreprise. Par rapport à Windows 10, les changements qui impactent le plus l'administration :

| Domaine | Windows 10 | Windows 11 | Impact admin |
|---|---|---|---|
| Exigences matérielles | Souples | TPM 2.0 + Secure Boot + CPU listé obligatoires | Inventaire parc indispensable avant migration |
| Menu Démarrer / barre des tâches | Personnalisable (tuiles) | Centré, non déplaçable, tuiles supprimées | Repackaging des layouts, GPO à revoir |
| Mises à jour | 2 feature updates/an (puis 1) | 1 mise à jour annuelle de fonctionnalités | Anneaux WUfB simplifiés |
| Cycle de vie Entreprise | 30 mois (H2) | 36 mois | Fenêtre de migration plus large |
| Sécurité par défaut | Opt-in | HVCI, Secure Boot, TPM exigés | Durcissement « gratuit » mais compat à tester |
| Store | Consumer | Bridge Entreprise (privé) | Nouveau packaging à prévoir |
| Copilot / Widgets / Teams Chat | Absents | Intégrés | GPO de désactivation en entreprise |

**Règle d'or** : ne jamais déployer Windows 11 comme « Windows 10 avec un nouveau thème ». C'est un changement de socle matériel et de modèle de gestion (cloud-first via Intune/Autopilot) qui mérite un projet dédié.

---

## 2. Exigences matérielles : TPM 2.0, Secure Boot, CPU, RAM, stockage

### 2.1 Tableau des exigences officielles

| Composant | Exigence minimale | Recommandé en entreprise |
|---|---|---|
| Processeur | 1 GHz+, 2 cœurs+, 64 bits, **dans la liste Microsoft** | Intel 8e gén+ / AMD Ryzen 2000+ / équivalent |
| RAM | 4 Go | 8 Go (bureautique), 16 Go (dev/CAO, multitâche lourd) |
| Stockage | 64 Go | 256 Go SSD NVMe (128 Go = trop juste avec profils + cache) |
| Firmware | UEFI, **Secure Boot capable** | Secure Boot activé |
| TPM | **TPM 2.0** (pas 1.2) | TPM 2.0 activé dans le BIOS/UEFI |
| Carte graphique | DirectX 12 / WDDM 2.0 | Intégrée suffisante en bureautique |
| Écran | 9"+, 720p | — |
| Internet | Requis pour l'OOBE (édition Famille) et les MàJ | — |

> ⚠️ **Point critique** : un PC avec TPM 1.2, un CPU 7e génération Intel ou un BIOS Legacy **ne passera pas** la vérification officielle, même si Windows 11 « pourrait » tourner. Les contournements de registre (`AllowUpgradesWithUnsupportedTPMOrCPU`) existent mais **ne sont pas supportés** : pas de mises à jour garanties, pas de support Microsoft. En entreprise : à proscrire.

### 2.2 TPM 2.0 — comprendre ce que c'est

Le TPM (Trusted Platform Module) est une puce (ou firmware : Intel PTT, AMD fTPM) qui :

- stocke les clés de chiffrement BitLocker sans les exposer à l'OS ;
- mesure l'intégrité du démarrage (measured boot) ;
- sert d'ancre pour Windows Hello, Credential Guard, l'attestation.

```powershell
# État du TPM : la commande de référence
Get-Tpm

# Sortie attendue sur un poste compatible :
# TpmPresent                : True
# TpmReady                  : True
# TpmEnabled                : True
# TpmActivated              : True
# ManagedAuthLevel          : Full
# OwnerAuth                 : (vide si géré par l'OS)
# AutoProvisioning          : Enabled
# LockedOut                 : False
```

```powershell
# Version du TPM (doit être 2.0)
(Get-CimInstance -Namespace root/cimv2/security/microsofttpm -ClassName Win32_Tpm).SpecVersion
# Exemple de sortie : 2.0, 0, 1.16
```

```powershell
# TPM via WMI (utile en inventaire à distance)
Get-WmiObject -Namespace "root\CIMV2\Security\MicrosoftTpm" -Class Win32_Tpm |
    Select-Object SpecVersion, IsEnabled, IsActivated, IsOwned
```

### 2.3 Secure Boot — vérification

```powershell
# État de Secure Boot
Confirm-SecureBootUEFI
# True = activé. Erreur "privilège non détenu" = BIOS Legacy ou Secure Boot non supporté.

# Via le registre (inventaire) :
Get-ItemProperty "HKLM:\SYSTEM\CurrentControlSet\Control\SecureBoot\State" -Name UEFISecureBootEnabled
```

### 2.4 CPU supportés — où vérifier

Microsoft publie des listes par génération (Intel, AMD, Qualcomm). En pratique :

- **Intel** : 8e génération (Coffee Lake) et supérieures. Quelques 7e gén spécifiques (ex. i7-7820HQ) sont listés mais c'est l'exception.
- **AMD** : Ryzen 2000 (Zen+) et supérieures, EPYC/Threadripper récents.
- **Qualcomm** : Snapdragon 850 et supérieures.

```powershell
# Identifier le CPU pour croiser avec la liste
Get-CimInstance Win32_Processor | Select-Object Name, Manufacturer, MaxClockSpeed
```

### 2.5 RAM et stockage — dimensionnement entreprise

| Usage | RAM | Stockage | Justification |
|---|---|---|---|
| Bureautique légère (web, Office) | 8 Go | 256 Go SSD | Marge pour MàJ (Windows.old, WinSxS) |
| Bureautique standard + visio | 16 Go | 256–512 Go SSD | Teams/Zoom + Edge consomment |
| Développement / CAO | 16–32 Go | 512 Go–1 To NVMe | Conteneurs, VMs locales |
| Postes partagés / kiosque | 8 Go | 256 Go SSD | Profils multiples |

> 💡 **Le 128 Go est un piège** : entre Windows (~30 Go), WinSxS, les points de restauration, le cache des MàJ et un profil utilisateur, on sature en moins d'un an. Standardisez sur 256 Go minimum.

---

## 3. Vérifier la compatibilité : PC Health Check, Get-TPM, scripts d'inventaire

### 3.1 PC Health Check (poste par poste)

