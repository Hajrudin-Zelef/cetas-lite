---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-19
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [3133, 3339]
sha256: 1d4187bb854c074af49aec857b88ba22328da7e9915ef7ee1a2689f86c2fee2b
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```
DELL :
  - Code à 8 caractères affiché après 3 échecs → le support Dell peut
    fournir un code de déverrouillage (preuve d'achat requise).
  - Outil : Dell Command Configure avec l'ancien mot de passe
    (si connu d'un coffre) : cctk --setuppwd= --valsetuppwd=<ancien>.

HP :
  - Après 3 échecs : code "System Disabled" → utilitaire HP BIOS
    Configuration Utility ou support HP (preuve d'achat).
  - Sur ProBook/EliteBook : la procédure SMC.bin (fichier de déblocage
    généré par HP) via clé USB.

LENOVO (ThinkPad) :
  - Mot de passe supervisor : AUCUNE procédure officielle de contournement.
    Carte mère à remplacer (c'est volontaire : sécurité).
  - Mot de passe "power-on" (au boot) : parfois réinitialisable par
    déconnexion batterie + CMOS (selon modèle), mais pas le supervisor.

GÉNÉRIQUE (tour/CM) :
  - Jumper "Clear CMOS" / "Password Clear" sur la carte mère
    (voir manuel ; souvent noté CLR_CMOS, JBAT1, ou PSWD).
  - Retrait pile + débrancher 15 min (uniquement vieilles cartes).
```

### Prévention (leçon managériale)

```
→ TOUS les mots de passe firmware dans le coffre d'équipe, dès la
  réception du matériel.
→ Procédure écrite de transmission lors des départs.
→ Ne jamais laisser un prestataire poser un mot de passe sans le consigner.
```

---

## 70. Cas pratique 7 : BitLocker demande la clé de récupération à chaque démarrage

### Contexte

Après un changement (ou sans cause apparente), Windows demande la clé de récupération à **chaque** démarrage. Le scellé TPM ne fonctionne plus.

### Causes possibles

| Cause | Indice |
|---|---|
| Firmware mis à jour sans re-scellé | PCR 0 changé, `Suspend/Resume` non fait |
| Périphérique USB branché au boot | Certains BIOS mesurent les périphériques (PCR) |
| Option firmware modifiée | Secure Boot on/off, CSM, ordre de boot → PCR 1/7 |
| TPM en erreur / re-provisionné | `Get-Tpm` → vérifier l'état |
| Protecteur TPM supprimé | `manage-bde -protectors -get C:` → plus de protecteur TPM |

### Résolution

```powershell
# 1. Débloquer avec la clé de récupération (une fois).
# 2. Une fois dans Windows, diagnostiquer :
Get-Tpm | Select-Object TpmPresent, TpmReady
manage-bde -protectors -get C:
manage-bde -status C:

# 3. Si le protecteur TPM a disparu → le recréer :
#    (nécessite la clé de récupération sous la main)
manage-bde -protectors -add C: -tpm

# 4. Forcer le re-scellé sur les PCR actuels :
Suspend-BitLocker -MountPoint "C:" -RebootCount 0
Resume-BitLocker -MountPoint "C:"

# 5. Redémarrer 2 fois et vérifier : plus de demande ?
# 6. Si le coupable est un périphérique USB : le débrancher au boot
#    ou figer la config (ne plus le brancher).
```

### Durcissement du profil PCR (optionnel, avancé)

```powershell
# Restreindre les PCR surveillés (ex. retirer le PCR 2 sensible aux cartes)
# Via GPO : "Configurer le profil de validation de plateforme TPM"
# ou en CLI (à manier avec précaution) :
manage-bde -protectors -add C: -tpm -pcr 0,4,11
# ⚠️ Réduire les PCR = réduire la sécurité. Documenter le choix.
```

---

## 71. Cas pratique 8 : Clear CMOS — quand et comment

### Contexte

Le **Clear CMOS** réinitialise les réglages du firmware aux valeurs d'usine (pas le mot de passe supervisor sur les machines modernes, pas le firmware lui-même).

### Quand l'utiliser

| Situation | Clear CMOS utile ? |
|---|---|
| Réglages incohérents après des bidouilles (overclocking, ordre de boot exotique) | ✅ Oui |
| PC qui ne démarre plus après un mauvais réglage (fréquence RAM...) | ✅ Oui |
| Mot de passe **utilisateur** (power-on) oublié sur vieille machine | ⚠️ Parfois |
| Mot de passe **supervisor/admin** oublié (machine moderne) | ❌ Non (voir §69) |
| Firmware corrompu (flash raté) | ❌ Non (voir §77) |
| TPM non détecté après MAJ | ⚠️ En dernier recours |

### Méthodes (par ordre de préférence)

```
Méthode 1 — Jumper (tour / carte mère) :
  1. Éteindre, DÉBRANCHER le cordon d'alimentation.
  2. Localiser le jumper CLR_CMOS / JBAT1 / CLRTC (voir manuel carte mère).
  3. Déplacer le cavalier sur la position Clear 10 secondes,
     puis le remettre en position normale.
  4. Rebrancher, démarrer, reconfigurer le setup.

Méthode 2 — Pile (si pas de jumper accessible) :
  1. Éteindre, débrancher, retirer la pile CR2032 15 minutes
     (appuyer sur le bouton power 10 s pour vider les condensateurs).
  2. Remettre la pile, rebrancher, démarrer.

Méthode 3 — Bouton (portables / cartes récentes) :
  Certains portables ont un trou "reset" (trombone 10 s) ou
  une combinaison (bouton power 30 s sans batterie ni secteur).

Méthode 4 — Depuis le setup :
  "Load Optimized Defaults" / "Restore Defaults" = Clear CMOS logiciel
  (ne touche pas aux mots de passe).
```

### Après un Clear CMOS : reconfiguration obligatoire

```powershell
# Le Clear CMOS remet TOUT par défaut. Repasser en revue :
#  [ ] Mode UEFI (pas Legacy)
#  [ ] Secure Boot → Enabled
#  [ ] TPM → On
#  [ ] Ordre de boot → disque en premier
#  [ ] Virtualisation → Enabled (si utilisée)
#  [ ] Mot de passe admin → REPOSER (il a pu être effacé !)
#  [ ] Date/heure → vérifier (la pile retirée = horloge perdue)
#
# Puis :
Confirm-SecureBootUEFI
Get-Tpm | Select-Object TpmPresent, TpmReady
# Si BitLocker demande la clé : normal (PCR changés) → la saisir,
# puis Suspend/Resume pour re-sceller (§37).
```

---

## 72. Cas pratique 9 : Windows 11 refuse de s'installer — « ce PC ne peut pas exécuter Windows 11 »

### Contexte

L'assistant d'installation ou `setup.exe` bloque : « Ce PC ne répond pas à la configuration système minimale requise ». Pourtant le PC semble correct.

### Diagnostic précis (ne pas deviner !)

```powershell
# 1. Les 4 piliers
Confirm-SecureBootUEFI                        # doit être True
(Get-Tpm).TpmReady                            # doit être True
(Get-Disk | Where-Object IsSystem).PartitionStyle  # doit être GPT

# 2. Détail du blocage (registre d'évaluation)
Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\TargetVersionUpgradeExperienceIndicators\NI23H2' |
    Format-List UpgEx, UpgExU, RedReason
# RedReason indique la cause exacte (ex. "Tpm", "SecureBoot", "Cpu", "Ram", "Storage")

# 3. Log du setup
Get-Content "$env:SystemRoot\Panther\setupact.log" -Tail 50 |
    Select-String -Pattern 'Compat|TPM|SecureBoot|UEFI' -CaseSensitive:$false
```

### Résolutions par cause

```
Cause = TPM (le plus fréquent) :
  → Setup UEFI → activer le TPM / fTPM / PTT (§31).
  → Parfois nommé "Security Device Support" (ASUS) ou "Trusted Computing".

Cause = SecureBoot :
  → Passer en UEFI natif (CSM OFF), activer Secure Boot (§23).
  → Si le disque est en MBR : mbr2gpt d'abord (§8).

Cause = Disque MBR :
  → mbr2gpt /convert /disk:0 /allowFullOS (§8).

Cause = CPU non supporté :
  → Vérifier la liste Microsoft (PC trop ancien) → remplacement.
  → ⚠️ Pas de contournement registre en entreprise.

Cause = "Tout est OK mais ça bloque quand même" :
  → Mettre à jour le BIOS (vieux firmware = TPM mal exposé).
  → Vérifier que le TPM est en 2.0 (pas 1.2).
  → Lancer l'installation depuis une clé USB créée avec l'outil
    officiel (Media Creation Tool), pas une ISO bricolée.
```

---

## 73. Cas pratique 10 : disque GPT vu comme « non alloué » après un passage en mode Legacy

### Contexte

Après avoir basculé le firmware en mode Legacy/CSM « pour tester », le disque système GPT n'apparaît plus comme bootable ; dans un outil de partitionnement, il semble vide ou « non alloué ». Panique.

### Analyse

