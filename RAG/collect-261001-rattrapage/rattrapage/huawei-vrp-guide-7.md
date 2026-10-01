---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-7
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [1309, 1489]
sha256: 310dadc7becc846f007d23e082e8b84b20e6d76ca9b0111a581615fc39b140ae
---

# VRP — Le système d'exploitation transversal Huawei

```vrp
<AR720>startup system-software flash:/AR720-V300R022C00SPC500.cc
Info: Succeeded in setting the software for booting system.
<AR720>display startup
MainBoard:
  Startup system software:        flash:/AR720-V300R019C00SPC200.cc
  Next startup system software:   flash:/AR720-V300R022C00SPC500.cc
  ...
<AR720>save                                     # fige la config actuelle
<AR720>reboot
  Warning: The system will reboot. Continue? [y/n]:y
```

`display startup` doit montrer l'ancien en « Startup » et le nouveau en
« Next startup ». Si ce n'est pas le cas, NE REBOOTEZ PAS.

## 78. Vérifications post-upgrade

```vrp
<AR720>display version                          # la nouvelle version est active ?
<AR720>display startup                          # Startup == Next startup == nouveau .cc
<AR720>display device                           # tout le matériel reconnu ?
<AR720>display interface brief                  # interfaces up ?
<AR720>display ip routing-table                 # routes présentes ?
<AR720>display logbuffer | include -i error     # erreurs au boot ?
<AR720>display patch-information                # patchs réappliqués si besoin
```

Checklist post-upgrade :

```text
[ ] display version = version cible
[ ] Uptime reparti de zéro (normal)
[ ] Config présente (display current-configuration | include sysname)
[ ] Interfaces et protocoles montés (ping passerelle, display ospf peer)
[ ] Supervision OK (SNMP, syslog, Zabbix)
[ ] Ancien .cc conservé en flash pour rollback (ou archivé sur serveur)
[ ] Nettoyage : suppression de l'ancien .cc après 1-2 semaines de stabilité
```

## 79. Patchs logiciels (`.pat`)

Les patchs corrigent des bugs sans changer de version majeure :

```vrp
<AR720>tftp 192.168.100.50 get AR720-V300R022SPH010.pat
<AR720>startup patch flash:/AR720-V300R022SPH010.pat
<AR720>display patch-information
<AR720>reboot                                   # certains patchs exigent un reboot
```

Rollback de patch :

```vrp
<AR720>undo startup patch                       # (syntaxe à vérifier selon version)
<AR720>reboot
```

## 80. Rollback en cas d'échec d'upgrade

Scénario : le nouveau `.cc` ne boote pas ou le comportement est anormal.

**Cas A — l'équipement boote encore (accès CLI OK) :**

```vrp
<AR720>startup system-software flash:/AR720-V300R019C00SPC200.cc   # ancien .cc
<AR720>display startup                          # vérifier
<AR720>reboot
```

**Cas B — l'équipement ne boote plus (BootROM/BootLoad) :**
1. Console, reboot, `Ctrl+B` au prompt « Press Ctrl+B to break auto
   startup ».
2. Mot de passe BootROM (défaut selon version : voir section 83).
3. Menu « Enter ethernet submenu » ou « Enter serial submenu » pour
   recharger un `.cc` via TFTP/FTP en BootROM.
4. « Boot with default mode » avec l'ancien fichier.

**Cas C — la config est incompatible :** restaurez la sauvegarde externe
(`copy tftp: ... vrpcfg.zip` + `startup saved-configuration` + `reboot`).

> Règle : ne supprimez JAMAIS l'ancien `.cc` avant d'avoir validé le
> nouveau en production pendant au moins une semaine.

## 81. Upgrade d'un stack (S310) — principes

1. Uploadez le `.cc` sur le **master** ; il se propage aux membres
   (selon modèle/version — à vérifier : `display stack`).
2. `startup system-software` sur le master.
3. `reboot` du stack (tous les membres redémarrent).
4. Vérifiez que tous les membres sont en `Ready` avec la bonne version :
   `display stack`, `display version` (détail par slot).

Si un membre ne remonte pas : console sur le membre, BootROM, rechargement
manuel du `.cc`.

## 82. Password recovery — principes généraux

Deux situations :
- **On a encore un accès admin** (SSH/Telnet/console avec un autre compte)
  → on change le mot de passe en ligne (méthode douce, sans coupure).
- **On a perdu TOUS les accès** → BootROM/BootLoad/Uboot via la console
  (coupure de service, fenêtre de maintenance).

Méthode douce (accès admin existant) :

```vrp
[AR720]aaa
[AR720-aaa]local-user admin password irreversible-cipher NouveauMotDePasse123!
[AR720-aaa]quit
[AR720]quit
<AR720>save
```

Ou pour la console :

```vrp
[AR720]user-interface console 0
[AR720-ui-console0]set authentication password cipher NouveauMotDePasse123!
[AR720-ui-console0]quit
```

## 83. Password recovery via BootROM/BootLoad (AR720, S310)

1. Reliez-vous en **console** (9600 8N1).
2. `reboot` (ou power cycle).
3. Au message `Press Ctrl+B to break auto startup ...`, pressez **Ctrl+B**
   (certains switchs : **Ctrl+E**).
4. Saisissez le mot de passe BootROM/BootLoad. Valeurs par défaut selon
   version (à vérifier dans le manuel de votre version exacte) :
   - V200R003 et antérieur : `huawei`
   - V200R005 et ultérieur : `Admin@huawei`
5. Dans le menu, choisissez **« Clear password for console user »**.
6. **Important** : choisissez ensuite **« Boot with default mode »**
   (option 1) — NE choisissez PAS « Reboot », sinon l'effacement ne
   s'applique pas.
7. Au démarrage, connectez-vous en console **sans mot de passe**.
8. Redéfinissez immédiatement un mot de passe console + `save` :

```vrp
<AR720>system-view
[AR720]user-interface console 0
[AR720-ui-console0]authentication-mode password
[AR720-ui-console0]set authentication password cipher NouveauMotDePasse123!
[AR720-ui-console0]quit
[AR720]quit
<AR720>save
```

> ⚠️ Tant que vous n'avez pas redéfini de mot de passe, un `reboot`
> referme la porte : refaites la manip. Et changez le mot de passe
> BootROM par défaut (`Modify BootROM password` dans le menu) !

## 84. Password recovery sur USG6000 — principes

Même logique que l'AR720 : console + BootROM/BootLoad + « Clear password
for console user ». Particularités USG :
- L'USG peut avoir un **mot de passe BootLoad** différent (voir la plaque
  / documentation livrée avec l'appareil).
- Si l'USG est en cluster HRP, faites l'opération sur les deux membres
  (ou cassez le cluster temporairement).
- Après récupération, pensez à re-synchroniser : `hrp` + `save` sur les
  deux membres.

> Les options exactes du menu BootLoad varient selon les versions USG ;
> suivez les libellés affichés à l'écran (« Clear password... », « Boot
> with default mode »).

## 85. Sécuriser le BootROM après récupération

```text
Menu BootROM/BootLoad -> "Modify BootROM password" (option 7)
Entrez un mot de passe fort, notez-le dans le coffre de l'équipe.
```

Un BootROM avec mot de passe par défaut = quiconque a accès à la console
peut effacer les mots de passe. C'est un vecteur classique d'escalade
physique.

## 86. Tableau Cisco IOS ↔ Huawei VRP — correspondances essentielles

