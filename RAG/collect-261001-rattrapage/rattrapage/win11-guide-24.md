---
id: collect-261001-rattrapage/rattrapage/win11-guide-24
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["agent", "amd", "intel", "license"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [4019, 4130]
sha256: 3e3077cefce5278bef7426ca5d98b8cc6465722e39b87ff15844e7e1f4cd0cb7
---

# Windows 11 en entreprise — Guide technique ultra-complet

| Terme | Définition |
|---|---|
| **ADBA** | Active Directory-Based Activation : activation en volume via AD, sans serveur KMS dédié |
| **ADK** | Assessment and Deployment Kit : outils Microsoft (WSIM, USMT, WinPE) pour le déploiement |
| **Anneau (ring)** | Groupe de postes recevant les MàJ avec un délai donné (WUfB) |
| **AppLocker** | Contrôle des applications par règles (éditeur, chemin, hash) — Entreprise |
| **ASR** | Attack Surface Reduction : règles Defender bloquant les comportements malveillants |
| **Autopilot** | Service cloud transformant un PC OEM en poste d'entreprise sans image |
| **BitLocker** | Chiffrement de volume natif Windows (XTS-AES) |
| **CSP** | Configuration Service Provider : équivalent MDM des GPO pour Intune |
| **DRA** | Data Recovery Agent : certificat permettant de déchiffrer tout volume (passe-partout) |
| **ESP** | Enrollment Status Page : écran bloquant l'OOBE tant que la config n'est pas prête |
| **ESU** | Extended Security Updates : correctifs payants après la fin de support |
| **FSLogix** | Conteneurs VHDX pour les profils (VDI/RDS) |
| **GVLK** | Generic Volume License Key : clé générique pour activation KMS/ADBA |
| **HVCI** | Hypervisor-protected Code Integrity (« Intégrité de la mémoire ») |
| **KFM** | Known Folder Move : redirection Bureau/Documents/Images vers OneDrive |
| **LAPS** | Local Administrator Password Solution : mots de passe admin locaux uniques et rotatifs |
| **LTSC** | Long-Term Servicing Channel : édition à support long, sans feature updates |
| **MDT** | Microsoft Deployment Toolkit : déploiement d'images « Lite Touch » |
| **OOBE** | Out-Of-Box Experience : écrans de premier démarrage |
| **PCR** | Platform Configuration Registers : registres du TPM stockant les mesures de démarrage |
| **PXE** | Preboot Execution Environment : démarrage réseau (WDS/MDT) |
| **Quality update** | Mise à jour mensuelle cumulative (sécurité + correctifs) |
| **Feature update** | Mise à jour annuelle de version (ex. 24H2) |
| **Sysprep** | Généralisation d'une installation avant capture d'image |
| **TPM** | Trusted Platform Module : puce/firmware de sécurité (clés, mesures de démarrage) |
| **UBR** | Update Build Revision : numéro de révision après le build (ex. 26100.XXXX) |
| **USMT** | User State Migration Tool : migration de profils (ScanState/LoadState) |
| **VBS** | Virtualization-Based Security : socle (Credential Guard, HVCI) |
| **WDAC** | Windows Defender Application Control : contrôle d'applications par stratégie signée |
| **WEF** | Windows Event Forwarding : collecte centralisée native des journaux |
| **WinRE** | Windows Recovery Environment : environnement de récupération |
| **WUfB** | Windows Update for Business : gestion cloud des MàJ (anneaux, délais) |
| **XTS-AES** | Mode de chiffrement par blocs utilisé par BitLocker |

---

## 78. Quiz : 10 questions + réponses commentées

### Q1 — Un poste affiche « TPM 2.0 » dans le BIOS mais `Get-Tpm` retourne `TpmPresent : False`. Quelle est la première chose à vérifier ?
- [ ] A. Réinstaller Windows
- [ ] B. Activer Intel PTT / AMD fTPM (ou le TPM discret) dans le firmware UEFI
- [ ] C. Acheter une puce TPM externe
- [ ] D. Désactiver Secure Boot

**Réponse : B.** Sur beaucoup de machines, le TPM firmware est simplement désactivé dans l'UEFI. `Get-Tpm` ne voit que ce qui est exposé par le firmware. La réinstallation (A) ne change rien, et Secure Boot (D) est un sujet distinct.

### Q2 — Quelle commande suspend BitLocker pour exactement un redémarrage (avant une MàJ du BIOS) ?
- [ ] A. `Disable-BitLocker -MountPoint "C:"`
- [ ] B. `Suspend-BitLocker -MountPoint "C:" -RebootCount 1`
- [ ] C. `Manage-bde -off C:`
- [ ] D. `Remove-BitLockerKeyProtector -MountPoint "C:"`

**Réponse : B.** A et C **déchiffrent** le volume (long et inutile ici), D supprime un protecteur. `Suspend-BitLocker -RebootCount 1` expose temporairement la clé en clair pour un seul boot — c'est la procédure standard avant MàJ BIOS.

### Q3 — Vous déployez Windows 11 via MDT. Où placez-vous les pilotes pour éviter les conflits entre modèles ?
- [ ] A. Tous dans un seul dossier « Drivers »
- [ ] B. Un dossier par modèle + profil de sélection dans la task sequence
- [ ] C. Dans l'image WIM directement
- [ ] D. Via Windows Update après déploiement

**Réponse : B.** Un dossier par modèle (`Out-of-Box Drivers\Win11-x64\<Modèle>`) avec un profil de sélection = la bonne pratique (section 11.2). A provoque des conflits, C alourdit l'image, D n'est pas maîtrisé.

### Q4 — Un utilisateur se connecte avec un « profil temporaire ». Dans le registre `ProfileList`, vous voyez sa clé SID et une clé `SID.bak`. Que faites-vous ?
- [ ] A. Vous renommez `.bak` en supprimant l'autre clé
- [ ] B. Vous supprimez les deux clés après sauvegarde des données, puis l'utilisateur se reconnecte sur un profil neuf
- [ ] C. Vous restaurez le système
- [ ] D. Vous copiez `C:\Users\TEMP` vers son dossier

**Réponse : B.** La méthode propre : sauvegarde des données, suppression des clés + du profil (via `Win32_UserProfile`), reconnexion = profil sain (cas pratique 8). A peut replanter si `NTUSER.DAT` est corrompu.

### Q5 — Quelle différence entre une mise à jour « qualité » et « fonctionnalité » dans WUfB ?
- [ ] A. Aucune, ce sont des synonymes
- [ ] B. Qualité = correctifs mensuels cumulatifs ; Fonctionnalité = nouvelle version annuelle (24H2…)
- [ ] C. Qualité = pilotes ; Fonctionnalité = sécurité
- [ ] D. Qualité = payante ; Fonctionnalité = gratuite

**Réponse : B.** Les délais WUfB se règlent séparément : court pour la qualité (7-14 j), long pour les fonctionnalités (60-120 j) — section 26.1.

### Q6 — `Confirm-SecureBootUEFI` retourne une erreur « Un privilège requis n'est pas détenu par le client ». Qu'est-ce que cela signifie ?
- [ ] A. Il faut lancer PowerShell en administrateur
- [ ] B. La machine est en BIOS Legacy ou Secure Boot n'est pas supporté/activé
- [ ] C. Le TPM est verrouillé
- [ ] D. BitLocker est désactivé

**Réponse : B.** Cette erreur (même en admin) indique que l'API UEFI Secure Boot n'est pas disponible : firmware Legacy ou Secure Boot non supporté. C'est un poste non compatible Windows 11 en l'état.

### Q7 — En GPO, vous filtrez sur un groupe de sécurité mais la GPO ne s'applique pas. Cause la plus probable depuis 2016 ?
- [ ] A. Le contrôleur de domaine est en panne
- [ ] B. Il manque les droits « Lire » (+ « Appliquer ») pour « Ordinateurs authentifiés » sur la GPO
- [ ] C. Il faut redémarrer le poste 3 fois
- [ ] D. Le filtre WMI est obligatoire

**Réponse : B.** Depuis la KB3163622, le filtrage de sécurité exige que « Ordinateurs authentifiés » (ou « Utilisateurs authentifiés ») ait au minimum « Lire » sur la GPO — l'oubli classique (cas pratique 10, erreur n°7).

### Q8 — Quelle méthode de déploiement ne nécessite AUCUNE image maître à maintenir ?
- [ ] A. MDT
- [ ] B. Windows Autopilot
- [ ] C. Clonezilla
- [ ] D. DISM /Capture-Image

**Réponse : B.** Autopilot transforme l'installation OEM via Intune : pas de WIM à maintenir, pas de PXE (section 12). Les autres reposent sur une image.

### Q9 — Un poste BitLocker redemande la clé à chaque démarrage après un changement de carte mère. Quelle est la bonne séquence ?
- [ ] A. Saisir la clé, `Suspend-BitLocker -RebootCount 1`, redémarrer, recréer le protecteur TPM si besoin, `Resume-BitLocker`, vérifier la sauvegarde AD
- [ ] B. Désactiver BitLocker définitivement
- [ ] C. Effacer le TPM sans la clé
- [ ] D. Réinstaller Windows

**Réponse : A.** C'est la procédure complète (sections 18.1-18.3), avec l'étape critique souvent oubliée : **vérifier que la clé est sauvegardée en AD après recréation du protecteur**.

