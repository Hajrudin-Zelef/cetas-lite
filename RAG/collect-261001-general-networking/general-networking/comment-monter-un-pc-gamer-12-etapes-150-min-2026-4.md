---
id: collect-261001-general-networking/general-networking/comment-monter-un-pc-gamer-12-etapes-150-min-2026-4
title: "Rechercher les pilotes GPU disponibles via winget"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["gpu", "amd", "benchmark", "intel", "nvidia"]
source: docs/RAG/collect-261001-general-networking/comment-monter-un-pc-gamer-12-etapes-150-min-2026.md
source_anchor: ""
source_lines: [131, 209]
sha256: 2d8fd74fc982da5ec87b358ed72a74f4fd3c83d7984d02b14c021f6acce3acaa
---

# Rechercher les pilotes GPU disponibles via winget

Terminez par les ventilateurs de boîtier (SYS_FAN) et, pour un AIO, le connecteur de pompe (souvent marqué PUMP ou AIO_PUMP, à ne pas relier à un en-tête de ventilateur classique qui couperait la pompe en cas de mise en veille agressive du contrôle thermique).

## Étape 10 : premier démarrage et configuration du BIOS/UEFI

Branchez clavier, souris et écran (sur la carte graphique, donc), puis mettez l’interrupteur de l’alimentation sur « I » et appuyez sur le bouton d’allumage du boîtier. Si tout est correctement enclenché, l’écran affiche le logo du fabricant puis le BIOS/UEFI. Profitez-en pour vérifier trois choses : le processeur est bien reconnu avec son nom exact, la quantité totale de RAM installée s’affiche, et le SSD apparaît dans la liste des périphériques de stockage.

```
Touches d'accès au BIOS/UEFI au démarrage (par fabricant) :
- ASUS ........ Suppr (Del) ou F2
- MSI ......... Suppr (Del)
- Gigabyte .... Suppr (Del)
- ASRock ...... Suppr (Del) ou F2
- Menu de démarrage ponctuel : F8, F11 ou F12 selon la carte
```
Activez ensuite le profil XMP (Intel) ou EXPO (AMD) dans la section liée à la mémoire, pour que la RAM fonctionne à sa vitesse annoncée plutôt qu’à sa fréquence de base. Terminez en plaçant la clé USB d’installation de Windows 11 en tête de l’ordre de démarrage, puis enregistrez et redémarrez.

## Étape 11 : installer Windows 11 et les pilotes

Depuis un autre ordinateur, créez une clé USB bootable avec l’outil de création de média de Microsoft ou avec Rufus, sur une clé de 32 Go minimum. Démarrez ensuite sur cette clé, suivez l’assistant d’installation, formatez la partition cible, et laissez Windows 11 copier ses fichiers puis redémarrer plusieurs fois automatiquement.

Un point de compatibilité à vérifier si l’installation refuse de démarrer : Windows 11 exige un module TPM 2.0 actif. Sur toute carte mère AM5 ou LGA1851 récente, ce module existe déjà sous forme logicielle (fTPM côté AMD, PTT côté Intel), mais reste parfois désactivé par défaut dans le BIOS. Un passage rapide dans les paramètres de sécurité du BIOS, avant de relancer l’installation, suffit à débloquer la situation.

Une fois sur le bureau, installez d’abord les pilotes chipset de la carte mère, puis les pilotes de la carte graphique (application Nvidia ou logiciel AMD selon le modèle). Windows Update propose parfois une version générique fonctionnelle mais rarement optimale pour le jeu : mieux vaut toujours passer par le site du fabricant du GPU pour la dernière version certifiée. Si vous changez de marque de carte graphique par rapport à une précédente installation, un nettoyage préalable avec DDU (Display Driver Uninstaller) évite les conflits entre anciens et nouveaux pilotes.

```
# Rechercher les pilotes GPU disponibles via winget
# (PowerShell lancé en tant qu'administrateur)
winget search nvidia
winget search "AMD Software"
winget install ID-retourne-par-la-recherche
```
## Étape 12 : activer le XMP/EXPO et valider le système

Retournez brièvement dans le BIOS pour confirmer que le profil XMP/EXPO activé à l’étape 10 est toujours actif après l’installation de Windows : il arrive qu’une mise à jour du BIOS ou un reset CMOS accidentel le désactive. Depuis Windows, une commande PowerShell permet de vérifier la vitesse effective de la RAM sans installer d’utilitaire tiers.

```
# Vérifier la vitesse effective de chaque module de RAM (PowerShell)
Get-CimInstance -ClassName Win32_PhysicalMemory | Select-Object BankLabel, Manufacturer, Capacity, Speed
```
Si la valeur retournée pour « Speed » correspond à la fréquence de base (autour de 4 800-5 200 MT/s pour de la DDR5) plutôt qu’à la fréquence annoncée sur le kit, le profil XMP/EXPO n’est pas actif. Retournez dans le BIOS et vérifiez qu’il est bien sélectionné et sauvegardé.

## Valider son montage : stress test, benchmark et températures cibles

Un PC qui démarre et affiche Windows n’est pas nécessairement un PC stable. Avant de considérer le montage terminé, un passage par un test de stress s’impose : Cinebench R24 pendant 10 à 15 minutes pour le processeur, puis 3DMark (Time Spy ou Steel Nova) pour la carte graphique, en surveillant les températures avec un utilitaire comme HWiNFO64 en tâche de fond.

Le guide de montage 2026 de Newegg recommande de rester sous 85 °C côté GPU et sous 80 °C côté CPU en pleine charge pour une configuration correctement ventilée. Le tableau suivant reprend ces seuils, complétés par les autres composants à surveiller.

| Composant | Au repos (idle) | En charge (stress test) | Seuil d’alerte | 
|---|---|---|---|
| Processeur (CPU) | Moins de 45 °C | Moins de 80 °C | Plus de 90 °C | 
| Carte graphique (GPU) | Moins de 40 °C | Moins de 85 °C | Plus de 90 °C | 
| VRM / étages d’alimentation | Moins de 50 °C | Moins de 90 °C | Plus de 100 °C | 
| SSD NVMe | Moins de 40 °C | Moins de 70 °C | Plus de 80 °C | 
| Mémoire RAM (DDR5) | Moins de 40 °C | Moins de 55 °C | Plus de 65 °C | 

```
Exemple de relevé pendant un test de stress de 20 minutes (HWiNFO64) :
CPU Package .......... repos 38 °C | charge 76 °C
GPU Hot Spot .......... repos 34 °C | charge 81 °C
Vitesse RAM détectée .. DDR5-6000 CL30 (profil EXPO actif)
Fréquence CPU en charge ... environ 5,3 GHz en boost multi-cœurs
```
Ces chiffres restent un exemple illustratif : la ventilation du boîtier, la qualité de la pâte thermique et la température ambiante font varier les résultats d’une configuration à l’autre, même à composants identiques.

## 5 erreurs fréquentes à éviter lors du montage d’un PC gamer

La plupart des montages qui tournent mal ne viennent pas d’une pièce défectueuse, mais d’un geste précipité. Voici les erreurs les plus courantes, dans l’ordre où elles surviennent généralement.

- **Oublier le connecteur d’alimentation du processeur (EPS 4+4 ou 8 broches).** Le PC s’allume, les ventilateurs tournent, mais aucun affichage n’apparaît et aucun bip ne se fait entendre : c’est le symptôme classique.
- **Forcer l’installation de la RAM ou du processeur.** Un composant qui résiste n’est presque jamais un composant à pousser plus fort, c’est un composant mal aligné.
- **Ignorer la décharge électrostatique.** Un carrelage sec en hiver suffit à accumuler assez d’électricité statique pour endommager une carte mère au moindre contact.
- **Brancher le connecteur USB 3.0 façade à l’envers.** Ce connecteur 19 ou 20 broches n’est pas toujours aussi bien détrompé que les autres. Vérifiez toujours l’orientation indiquée dans le manuel avant de forcer.
- **Négliger la gestion des câbles.** Des câbles qui obstruent le flux d’air peuvent faire grimper les températures internes de 5 à 10 °C, sans qu’aucune pièce ne soit en cause.
- **Oublier d’activer le XMP ou l’EXPO après le premier démarrage.** Sans cette activation, la RAM tourne à sa fréquence JEDEC de base, souvent bien inférieure à la vitesse annoncée sur la boîte.

## Dépannage : 8 problèmes courants et leurs solutions

Même en suivant chaque étape avec soin, un premier démarrage qui ne se passe pas comme prévu reste fréquent. Voici les huit situations les plus signalées après un montage, et comment les résoudre sans tout démonter.

