---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-11
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [1866, 2078]
sha256: 9f7ba66e12885fcedff232373539b68ad815dc8c257ecab7089a8908b6628e61
---

# Maintenance et exploitation Windows en entreprise

- Ne supprimez jamais la partition de récupération pour « gagner 500 Mo ».
- Sur les serveurs, testez l'accès WinRE **avant** d'en avoir besoin (au moins une
  fois par an, en fenêtre de maintenance).
- Gardez une **clé USB d'installation** de la même version à portée de main en
  salle serveur : c'est le WinRE de secours (§62).
- Ajoutez vos outils à WinRE via `winrecfg` / image personnalisée seulement si
  vous maîtrisez la chaîne (sinon : clé USB bootable avec vos scripts).

---

## 64. Réparation du démarrage : diagnostic

Symptômes : « Bootmgr is missing », logo Windows qui boucle, écran noir après le
logo, `0xc000000e` / `0xc0000225` (périphérique inaccessible).

Méthode (depuis WinRE > Dépannage > Invite de commandes) :

1. **Identifier les volumes** : les lettres changent sous WinRE ! `diskpart` >
   `list volume` — notez quelle lettre porte `\Windows`.
2. **Réparation automatique** d'abord : Dépannage > Réparation du démarrage
   (règle les cas simples : BCD corrompu léger).
3. Si échec : **manuel** avec `bcdedit` (§65) puis `bootrec` (§66).
4. En UEFI : la partition EFI (FAT32, ~100 Mo) doit être montée pour `bcdboot`.

```cmd
:: Sous WinRE : lister les volumes et trouver Windows + partition EFI
diskpart
list volume
:: (repérez : volume Windows = ex. C:, volume EFI = FAT32 "EFI")
exit
```

---

## 65. BCD : bcdedit en pratique

Le **BCD** (Boot Configuration Data) remplace `boot.ini`. Lecture :

```cmd
:: Magasin BCD courant (depuis Windows démarré)
bcdedit /enum

:: Depuis WinRE : préciser le magasin
bcdedit /store C:\Boot\BCD /enum
```

Opérations courantes :

```cmd
:: Sauvegarder le BCD avant toute modification (OBLIGATOIRE)
bcdedit /export C:\Temp\bcd-backup

:: Restaurer
bcdedit /import C:\Temp\bcd-backup

:: Définir l'OS par défaut / le délai du menu
bcdedit /default {current}
bcdedit /timeout 10

:: Recréer les entrées de démarrage (UEFI) : S: = lettre de la partition EFI montée
bcdboot C:\Windows /s S: /f UEFI
:: Variante BIOS/MBR
bcdboot C:\Windows /s C: /f BIOS
```

> `bcdboot` **recrée** des fichiers de démarrage sains à partir du Windows
> installé : c'est souvent plus efficace que de rafistoler un BCD corrompu.

---

## 66. bootrec : /fixmbr, /fixboot, /rebuildbcd

Depuis l'invite WinRE (surtout utile en **BIOS/MBR** ; en UEFI préférez `bcdboot`
§65) :

```cmd
:: Réécrire le MBR (sans toucher à la table de partitions)
bootrec /fixmbr

:: Réécrire le secteur de démarrage de la partition système
bootrec /fixboot
:: Si "Accès refusé" sous WinRE : monter la partition EFI puis bcdboot (§65),
:: ou : diskpart > sel vol EFI > assign letter=S: > exit, puis bcdboot C:\Windows /s S: /f UEFI

:: Rechercher les installations Windows et reconstruire le BCD
bootrec /scanos
bootrec /rebuildbcd
```

Séquence type MBR corrompu : `bootrec /fixmbr` → `bootrec /fixboot` →
`bootrec /rebuildbcd` → redémarrage. Si le disque n'est pas vu du tout
(`scanos` ne trouve rien) : problème **matériel/BIOS** (disque, câble, contrôleur),
pas logiciel — ne perdez pas 2 h sur bootrec.

---

## 67. Restauration du système et clichés instantanés

La **restauration du système** (points de restauration) annule les changements
système (pilotes, registre, correctifs) sans toucher aux données utilisateurs.

```powershell
# Vérifier / activer la protection sur C:
vssadmin list shadows
Get-ComputerRestorePoint
Enable-ComputerRestore -Drive 'C:\'

# Créer un point manuel avant une opération à risque
Checkpoint-Computer -Description 'Avant MAJ pilote RAID' -RestorePointType 'MODIFY_SETTINGS'
```

Restauration : WinRE > Dépannage > Restauration du système, ou `rstrui.exe` depuis
Windows.

**Limites** : désactivée par défaut sur beaucoup de serveurs ; ne remplace pas une
sauvegarde (§123) ; inutile contre une corruption de données applicatives.
**Politique** : activez-la sur les postes (avec espace réservé 5-10 %) et créez un
point avant chaque changement à risque (scriptable, ci-dessus).

---

## 68. Réinitialisation et réinstallation propre

Quand le système est trop dégradé (ou pour recycler un poste) :

| Option | Effet | Quand |
|---|---|---|
| **Réinitialisation** (« Réinitialiser ce PC ») | Réinstalle Windows, conserve ou non les fichiers | Poste très instable, avant réaffectation |
| **Nouvelle installation** (USB/ISO) | Système neuf | Serveur corrompu sans sauvegarde système, ou standardisation |
| **Réinstallation via image maître** (MDT/Intune Autopilot) | Poste conforme au standard | **À privilégier** en entreprise |

```powershell
# Lancer la réinitialisation (interactive)
systemreset.exe
# Réinitialisation "tout supprimer" en ligne de commande (poste à recycler)
systemreset.exe --cleanpc   # selon version : ouvre l'assistant
```

> En entreprise, un poste qui cumule les problèmes système récurrents ne se
> « répare » pas indéfiniment : on le **réinstalle depuis l'image standard**.
> C'est plus rapide et plus fiable que 3 h de DISM/SFC/chkdsk en cascade.
> Pour les serveurs : même logique — un serveur « rafistolé » reste fragile ;
> reconstruisez-le proprement puis restaurez données et rôles (d'où l'importance
> de la documentation §128 et des sauvegardes §123).

---

## 69. BSOD : lire un écran bleu

Un **BSOD** (écran bleu) = le noyau a rencontré une erreur irrécupérable et a
stoppé la machine pour protéger les données. L'écran affiche :

- Le **code d'arrêt** (ex. `IRQL_NOT_LESS_OR_EQUAL`) + sa valeur hexadécimale
  (ex. `0x0000000A`) ;
- Parfois le **pilote fautif** (ex. `ntoskrnl.exe`, `nvlddmkm.sys`) ;
- Un QR code (peu utile en pratique).

**Premiers réflexes :**

1. **Photographiez l'écran** (code + pilote) : c'est souvent la seule trace si
   aucun dump n'est configuré.
2. Notez le contexte : que faisait la machine ? (mise à jour récente ? nouveau
   matériel/pilote ? charge particulière ?)
3. Redémarrez : BSOD **isolé** après un patch = à surveiller ; BSOD **répété** =
   diagnostic (§72-74).
4. Récupérez le dump (`C:\Windows\Minidump\`, §70) **avant** tout nettoyage.

---

## 70. Configurer les fichiers de vidage (dump)

Sans dump configuré, impossible d'analyser un BSOD après coup. Vérifiez :

```powershell
# Configuration actuelle des vidages
Get-ItemProperty 'HKLM:\SYSTEM\CurrentControlSet\Control\CrashControl' |
  Select-Object CrashDumpEnabled, DumpFile, MinidumpDir, AutoReboot, LogEvent

# Valeurs de CrashDumpEnabled : 1 = complet, 2 = noyau, 3 = petit (minidump 256 Ko),
# 7 = automatique (recommandé : gère la taille selon la RAM)
```

Réglage recommandé **serveurs** : **vidage automatique** (7) — ni le complet
(taille = RAM, trop lourd) ni le seul minidump (parfois insuffisant). Vérifiez
aussi :

- Espace libre ≥ taille de la RAM + 1 Go sur le volume du dump ;
- Fichier d'échange sur le volume système ≥ 800 Mo (requis pour écrire le dump) ;
- `DumpFile` = `%SystemRoot%\MEMORY.DMP`, `MinidumpDir` = `%SystemRoot%\Minidump`.

> Après un BSOD, **copiez** `MEMORY.DMP` / le `.dmp` ailleurs avant toute
> manipulation : un nouveau BSOD écrase le précédent, et certains nettoyages
> suppriment les dumps (§28).

---

## 71. WinDbg : installation et prise en main

**WinDbg** (Windows Debugger) est l'outil d'analyse des dumps. Installation :

1. **Microsoft Store** : « WinDbg » (version moderne, recommandée) ; ou via le
   **SDK Windows** (option Debugging Tools).
2. Au premier lancement : configurez le **chemin des symboles** (indispensable
   pour une analyse lisible) :

```text
.sympath srv*C:\Symbols*https://msdl.microsoft.com/download/symbols
```

3. Ouvrez le dump : Fichier > Open dump file > `C:\Windows\Minidump\*.dmp`.

Commandes de base une fois le dump chargé :

