---
id: collect-261001-rattrapage/rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026-4
title: "Calculer l'empreinte SHA-256 de l'archive (PowerShell)"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026.md
source_anchor: ""
source_lines: [225, 306]
sha256: 1bd24ee6385661ef53ca88d4986bc3eaf0bfaa956905d97db1b4e89c037bbf57
---

# Calculer l'empreinte SHA-256 de l'archive (PowerShell)

Une fois le pilote en place, validez la stabilité avant de conclure. Lancez un jeu exigeant ou un test de charge (3DMark, Superposition, ou simplement une boucle de benchmark intégré) pendant une dizaine de minutes et surveillez les températures, la fréquence et l’absence d’artefacts. Réactivez ensuite les fonctionnalités que vous utilisez : G-SYNC/FreeSync, résolution et taux de rafraîchissement natifs de l’écran, profils de couleur. Si tout est stable, n’oubliez pas de rétablir la recherche de pilotes via Windows Update si vous comptez laisser le système se mettre à jour à l’avenir – en gardant à l’esprit qu’un pilote installé manuellement reste prioritaire tant qu’il est plus récent.

## Étape 12 – Automatiser DDU en ligne de commande

Pour les techniciens qui gèrent plusieurs machines, ou simplement pour gagner du temps, DDU s’exécute en mode silencieux via des commutateurs en ligne de commande. C’est notre « projet complet » : un script batch qui nettoie les pilotes NVIDIA et redémarre, sans la moindre interaction. À lancer depuis le mode sans échec, en administrateur.

```
@echo off
:: ddu-auto.bat – Nettoyage automatise des pilotes GPU avec DDU 18.1.5.4
:: A executer EN ADMINISTRATEUR, de preference en mode sans echec
set "DDU=C:\DDU\Display Driver Uninstaller.exe"
if not exist "%DDU%" (
  echo Erreur : DDU introuvable dans C:\DDU
  pause
  exit /b 1
)
echo Nettoyage des pilotes NVIDIA en cours...
"%DDU%" -silent -cleannvidia -removemonitors -restart
:: Variantes :
::   AMD   -^>  "%DDU%" -silent -cleanamd -restart
::   Intel -^>  "%DDU%" -silent -cleanintel -restart
::   Audio -^>  "%DDU%" -silent -cleanrealtek
```
Le tableau suivant récapitule les principaux commutateurs reconnus par DDU 18.1.5.4. Notez que l’option « supprimer PhysX » n’est disponible que dans l’interface graphique, pas en ligne de commande.

| Commutateur | Effet | 
|---|---|
| `-Silent` | Exécution silencieuse, sans interface graphique | 
| `-CleanNvidia` | Nettoie les pilotes NVIDIA | 
| `-CleanAMD` | Nettoie les pilotes AMD | 
| `-CleanIntel` | Nettoie les pilotes Intel (y compris Arc) | 
| `-RemoveMonitors` | Supprime les données EDID / moniteurs enregistrés | 
| `-CleanRealtek` | Nettoie les pilotes audio Realtek (ajout 18.1.5.x) | 
| `-CleanSoundblaster` | Nettoie les pilotes audio Sound Blaster | 
| `-Restart` | Redémarre le PC après le nettoyage | 

Enregistrez le script sous `ddu-auto.bat`, faites un clic droit puis *Exécuter en tant qu’administrateur*. Pour un déploiement de masse, vous pouvez combiner ce batch avec un démarrage forcé en mode sans échec (`bcdedit`) orchestré par GPO ou par un outil de gestion de parc. Pensez toujours à rétablir le mode normal en fin de script.

## Windows 11 24H2 : le piège du code PIN en mode sans échec

Voici le problème le plus signalé en 2026. Sur certaines configurations **Windows 11 24H2**, la connexion par **code PIN (Windows Hello) ne fonctionne pas en mode sans échec**. Vous arrivez à l’écran de connexion, vous saisissez votre PIN… et rien ne se passe, ou un message indique que l’option est indisponible. Comme DDU s’utilise précisément en mode sans échec, l’utilisateur se retrouve bloqué hors de sa session, parfois sans se souvenir de son mot de passe de compte Microsoft.

La cause : en mode sans échec, le service de gestion des identités (et certains composants de Windows Hello) n’est pas toujours chargé. Trois parades fiables :

1. **Utilisez votre mot de passe** de compte Microsoft (ou local) plutôt que le PIN. Sur l’écran de connexion, cliquez sur « Options de connexion » puis choisissez l’icône de mot de passe. C’est la solution la plus simple – d’où l’importance de le connaître avant de commencer.
2. **Démarrez en « Mode sans échec avec prise en charge réseau »** (touche 5 ou F5 dans les Paramètres de démarrage) : davantage de services sont chargés, ce qui rétablit parfois l’authentification.
3. **Basculez temporairement sur un compte local** avant l’opération (Paramètres > Comptes), avec un mot de passe simple, puis rebasculez sur votre compte Microsoft ensuite.

Par sécurité, réinitialisez ou mémorisez votre mot de passe de compte Microsoft *avant* de redémarrer en mode sans échec. C’est l’unique précaution qui vous évitera une mauvaise surprise sur 24H2.

## Pièges courants à éviter avec DDU

Même bien conçu, **Display Driver Uninstaller** punit l’imprudence. Voici les erreurs qui reviennent le plus souvent et comment les contourner :

- **Laisser Windows Update réinstaller le pilote.** C’est l’erreur n°1. Sans le blocage de l’étape 6 ni la déconnexion réseau, Windows réinjecte un pilote générique dans les minutes qui suivent le nettoyage, et tout est à refaire.
- **Oublier de télécharger le nouveau pilote avant.** Se retrouver sur l’affichage basique Microsoft sans paquet d’installation sous la main transforme une tâche de 30 minutes en galère.
- **Lancer DDU en mode normal.** Cela fonctionne, mais laisse des résidus. Le mode sans échec reste la recommandation officielle de Wagnard.
- **Sélectionner le mauvais fabricant.** Choisir AMD alors qu’on a une carte NVIDIA ne supprime rien d’utile et fait perdre un cycle complet de redémarrage.
- **Exécuter DDU depuis un lecteur réseau ou une clé USB lente.** L’outil exige un disque local ; sur un support réseau, il refuse de démarrer ou se comporte de façon erratique.
- **Télécharger DDU depuis une source douteuse.** Les versions repackagées embarquent parfois des logiciels indésirables. Tenez-vous-en à Wagnardsoft et Guru3D.

## Dépannage : 8 problèmes fréquents et leurs solutions

Voici les incidents les plus rapportés après un passage de **DDU**, avec leur cause probable et la marche à suivre.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Écran noir après nettoyage | Aucun pilote chargé, sortie vidéo sur le mauvais port | Branchez l’écran sur la sortie de la carte dédiée, réinstallez le pilote | 
| Résolution bloquée en 1024×768 | Pilote d’affichage basique Microsoft actif | Installez le pilote GPU téléchargé à l’étape 4 | 
| Windows réinstalle un vieux pilote | Recherche Windows Update non bloquée | Appliquez les clés de registre de l’étape 6, restez hors ligne | 
| Code PIN refusé en mode sans échec | Windows Hello indisponible (Win11 24H2) | Connectez-vous par mot de passe, ou mode sans échec avec réseau | 
| DDU refuse de se lancer | Exécution depuis un lecteur réseau, ou pas d’admin | Copiez DDU sur C:, lancez en administrateur | 
| Carte non détectée après réinstallation | Connecteur PCIe ou alimentation mal branchés | Vérifiez le branchement, testez un autre port PCIe | 
| Application NVIDIA/Adrenalin ne s’ouvre pas | Composant logiciel non installé | Réinstallez le pilote complet (option personnalisée) | 
| Paquet pilote résiduel après nettoyage | INF tiers verrouillé par le système | Supprimez-le via `pnputil /delete-driver` | 

Si malgré tout l’instabilité persiste, revenez au point de restauration créé à l’étape 5, puis recommencez le nettoyage en respectant scrupuleusement le mode sans échec et le blocage de Windows Update.

## Astuces avancées pour aller plus loin

Quelques techniques pour les utilisateurs aguerris. D’abord, le **journal de DDU** : chaque session génère un fichier texte dans le sous-dossier `Logs`. En cas de problème récurrent, il documente précisément ce qui a été supprimé et les éventuelles erreurs – indispensable pour un diagnostic sérieux ou un signalement sur le forum Wagnardsoft.

