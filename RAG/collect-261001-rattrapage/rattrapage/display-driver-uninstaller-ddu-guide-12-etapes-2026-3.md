---
id: collect-261001-rattrapage/rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026-3
title: "Calculer l'empreinte SHA-256 de l'archive (PowerShell)"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: ["2026-12-06"]
keywords: ["amd", "arr", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026.md
source_anchor: ""
source_lines: [142, 224]
sha256: a480a83ac8246eb22d5ce865a287e49760691db3a022a32a7c4ae5429f3beadc
---

# Calculer l'empreinte SHA-256 de l'archive (PowerShell)

Wagnard recommande explicitement d’utiliser **Display Driver Uninstaller** en mode sans échec. La raison est simple : en mode normal, Windows maintient le pilote graphique chargé en mémoire et verrouille certains fichiers, ce qui empêche une suppression totale. En mode sans échec, seul le pilote d’affichage minimal est actif, et DDU peut tout effacer sans interférence.

La méthode la plus fiable consiste à forcer le démarrage en mode sans échec via `bcdedit`, puis à annuler ce réglage après l’opération :

```
:: Invite de commandes EN ADMINISTRATEUR
:: Forcer le prochain demarrage en mode sans echec minimal
bcdedit /set {current} safeboot minimal
:: Redemarrer immediatement
shutdown /r /t 0
:: --- APRES le nettoyage et la reinstallation, ANNULER le mode sans echec ---
bcdedit /deletevalue {current} safeboot
```
Vous pouvez aussi laisser DDU gérer le redémarrage : son interface propose un bouton « Lancer en mode sans échec » qui prépare le boot automatiquement. Autre voie, plus visuelle, via l’environnement de récupération : exécutez `shutdown /r /o /f /t 0`, puis suivez *Dépannage > Options avancées > Paramètres de démarrage > Redémarrer*, et appuyez sur 4 ou F4 pour le mode sans échec. Quelle que soit la méthode, n’oubliez pas votre mot de passe de connexion.

## Étape 8 – Lancer DDU et régler les options

Une fois en mode sans échec, ouvrez le dossier `C:\DDU` et lancez `Display Driver Uninstaller.exe` par un clic droit puis *Exécuter en tant qu’administrateur*. Au premier démarrage, une fenêtre d’options s’affiche. Les réglages par défaut conviennent à la plupart des utilisateurs, mais vérifiez les points suivants :

- **Empêcher les téléchargements de pilotes via Windows Update** : à cocher, en complément des clés de registre de l’étape 6.
- **Supprimer les dossiers de pilotes** (C:\AMD, C:\NVIDIA, C:\Intel) : utile pour un nettoyage total.
- **Supprimer les moniteurs enregistrés** : à activer si vous changez d’écran ou si vous rencontrez des soucis d’EDID/résolution.
- **Créer un point de restauration** : laissez coché si vous n’en avez pas créé à l’étape 5.

Dans le menu déroulant en haut à droite, sélectionnez d’abord le **type de périphérique « GPU »**, puis le **fabricant correspondant à votre carte** : NVIDIA, AMD ou Intel. C’est l’étape critique : choisir le mauvais fabricant ne supprimera rien d’utile. Si vous avez à la fois un iGPU Intel et une carte dédiée, traitez chaque fabricant l’un après l’autre. L’interface est traduite et indique en clair la carte détectée.

## Étape 9 – Nettoyer les pilotes (Clean and Restart)

DDU propose trois actions principales. Le tableau ci-dessous les compare pour vous aider à choisir. Dans la grande majorité des cas, **« Nettoyer et redémarrer »** est le bon choix : il efface les pilotes puis relance Windows proprement.

| Bouton DDU | Action | Quand l’utiliser | 
|---|---|---|
| Nettoyer et redémarrer | Supprime les pilotes puis redémarre le PC | Cas standard : réinstallation ou changement de carte | 
| Nettoyer sans redémarrer | Supprime les pilotes, reste en session | Pour enchaîner sur un second fabricant avant de redémarrer | 
| Nettoyer et éteindre | Supprime les pilotes puis arrête le PC | Lorsque vous remplacez physiquement la carte graphique | 

Cliquez sur le bouton voulu et laissez DDU travailler : une barre de progression et un journal détaillé défilent. L’opération dure de 30 secondes à 2 minutes. L’écran peut clignoter ou changer de résolution : c’est normal, le pilote est en train d’être retiré. Ne touchez à rien jusqu’au redémarrage automatique.

Si vous nettoyez en mode normal plutôt qu’en mode sans échec (déconseillé), répétez le cycle **Nettoyer → Redémarrer → Nettoyer → Redémarrer** deux fois pour un résultat équivalent. En mode sans échec, un seul passage suffit. À ce stade, votre système redémarre sans pilote graphique propriétaire : place à la vérification.

## Étape 10 – Vérifier la suppression avec pnputil

Comment être sûr que DDU n’a rien oublié ? L’utilitaire Windows **pnputil** liste tous les paquets de pilotes encore enregistrés dans le magasin de pilotes. Après un nettoyage réussi, la classe « Affichage » ne doit plus contenir que l’adaptateur de base de Microsoft (ou rien du tout). Ouvrez une invite de commandes en administrateur :

```
:: Lister les paquets de pilotes d'affichage encore presents
pnputil /enum-drivers /class Display
:: Variante compatible toutes versions (filtre la sortie)
pnputil /enum-drivers | findstr /i "Display Affichage nv_disp amd ig"
```
Avant nettoyage, la sortie ressemble à ceci (un pilote NVIDIA tiers encore enregistré) :

```
Utilitaire Microsoft PnP
Nom publie :        oem42.inf
Nom d'origine :     nv_dispi.inf
Fournisseur :       NVIDIA
Classe :            Cartes graphiques
Version :           06/12/2026 32.0.15.7652
Nom signataire :    Microsoft Windows Hardware Compatibility Publisher
```
Après un nettoyage DDU réussi, la même commande renvoie « Aucun élément trouvé » pour la classe Affichage, ou uniquement `Microsoft Basic Display Adapter`. S’il subsiste un paquet tiers récalcitrant, supprimez-le manuellement (remplacez `oem42.inf` par le nom réel) :

```
:: Supprimer un paquet pilote residuel et le desinstaller du systeme
pnputil /delete-driver oem42.inf /uninstall /force
```
La documentation complète des commutateurs est disponible dans la référence officielle pnputil de Microsoft. Cet outil est précieux pour diagnostiquer les pilotes fantômes que même DDU peut, rarement, laisser passer.

## Étape 11 – Réinstaller des pilotes propres (NVIDIA, AMD, Intel Arc)

Votre système est désormais vierge de tout pilote graphique propriétaire. Lancez le paquet d’installation téléchargé à l’étape 4. Quelques recommandations par fabricant :

### NVIDIA

Choisissez l’option **« Personnalisée »** puis cochez **« Effectuer une installation propre »**. Depuis 2024, l’installeur NVIDIA est modulaire : vous pouvez décocher GeForce Experience / l’application NVIDIA si vous ne l’utilisez pas. Pour la majorité des joueurs, le pilote Game Ready suffit ; les créateurs préféreront la branche Studio, plus stable pour les applications professionnelles.

### AMD et Intel

Côté AMD, l’installeur Adrenalin propose lui aussi une **« Installation usine »** qui repart d’une base propre – pratique, mais inutile ici puisque DDU a déjà fait le ménage ; une installation standard convient. Pour les cartes **Intel Arc**, installez le pilote Arc & Graphics le plus récent : les gains de performances pilote après pilote restent significatifs sur cette gamme. Après installation, redémarrez et n’oubliez pas d’annuler le mode sans échec (`bcdedit /deletevalue {current} safeboot`) si vous l’aviez forcé.

Vérifiez enfin que la bonne version est active en relançant la commande `Get-CimInstance Win32_VideoController` de l’étape 3 : le numéro de pilote doit correspondre à celui que vous venez d’installer. Si vous suivez l’actualité des GPU, notre comparatif RX 9070 XT vs RTX 5070 Ti et notre analyse de la stratégie NVIDIA DSX complètent ce tutoriel.

