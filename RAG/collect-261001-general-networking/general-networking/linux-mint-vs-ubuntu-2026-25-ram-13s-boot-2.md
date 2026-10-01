---
id: collect-261001-general-networking/general-networking/linux-mint-vs-ubuntu-2026-25-ram-13s-boot-2
title: "Linux Mint : installer Flatpak puis une app depuis Flathub"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "Samsung"]
dates: []
keywords: ["amd", "benchmarks", "gpu", "intel", "nvidia", "sandbox"]
source: docs/RAG/collect-261001-general-networking/linux-mint-vs-ubuntu-2026-25-ram-13s-boot.md
source_anchor: ""
source_lines: [45, 108]
sha256: 7807afa9d0e0c88608f61eba49079da5cea68e952a7409cd0f2a3797cc36561e
---

# Linux Mint : installer Flatpak puis une app depuis Flathub

La mise à jour Cinnamon 6.4 livrée avec Mint 22.1 apporte plusieurs améliorations notables : un nouveau gestionnaire de notifications, un panneau supportant le redimensionnement vertical, la prise en charge des fonds d’écran animés et un mode Wayland expérimental activable via la commande `cinnamon-launcher --wayland`. Selon Clement Lefebvre, fondateur de Mint, l’objectif officiel est d’atteindre la parité X11/Wayland avec Cinnamon 7.0 prévu fin 2026. En l’état, l’expérience Wayland reste limitée : les applets tiers cassent, le support des écrans à taux de rafraîchissement mixtes est partiel, et NVIDIA EGLStreams pose encore des soucis.

### GNOME 49 : maturité Wayland confirmée

Ubuntu utilise Wayland par défaut depuis 22.04 LTS et l’écosystème GNOME a investi massivement dans ce protocole. La version 49 livrée avec Ubuntu 25.10 corrige les derniers bugs récurrents sur multi-écrans avec scaling fractionnel (125 %, 150 % ou 175 %), améliore la prise en charge des tablettes graphiques Wacom et accélère le compositeur Mutter de 18 % sur GPU AMD selon les mesures Phoronix de février 2026. Pour un utilisateur de Blender, Krita ou DaVinci Resolve sur un MacBook ou un Dell XPS avec écran 4K, l’expérience Ubuntu reste plus aboutie que celle de Mint en avril 2026.

## Benchmarks RAM et CPU : un écart de 25 % au repos [Testé]

Les chiffres officiels masquent souvent les écarts réels. Pour la comparaison **Linux Mint vs Ubuntu**, nous avons agrégé trois sources de benchmarks publiées entre janvier et mars 2026 : Phoronix Test Suite 10.8, LinuxBench Foundation et l’équipe Tech-Insider. Tous les tests ont été menés sur la même configuration de référence : Intel Core i5-13400, 16 Go de DDR5-5600, SSD NVMe Samsung 990 Pro 1 To, GPU Intel UHD 730. Voici les résultats en consommation mémoire et processeur après un démarrage à froid avec uniquement les services par défaut.

| Métrique | Linux Mint 22.1 Cinnamon | Ubuntu 25.10 GNOME | Écart | 
|---|---|---|---|
| RAM au repos (idle) | 912 Mo | 1 218 Mo | -25,1 % | 
| RAM après ouverture Firefox | 1,84 Go | 2,21 Go | -16,7 % | 
| RAM avec LibreOffice + Firefox + Thunderbird | 3,12 Go | 3,65 Go | -14,5 % | 
| CPU au repos (gnome-shell / cinnamon) | 0,8 % | 1,4 % | -42,9 % | 
| Processus actifs au démarrage | 198 | 247 | -19,8 % | 
| Services systemd actifs | 132 | 178 | -25,8 % | 
| Empreinte ISO installable | 2,89 Go | 5,87 Go | -50,8 % | 

L’écart de 25 % au repos est constant à travers les trois sources, avec des variations inférieures à 50 Mo. Sur les charges réelles (bureautique avec Firefox, LibreOffice et Thunderbird ouverts en simultané), l’écart se réduit à environ 14 % parce qu’une partie de la mémoire est dévolue aux applications, identiques entre les deux distributions. Sur une machine équipée de seulement 4 Go de RAM, cette différence est cruciale : Mint laisse un peu moins d’un gigaoctet de marge alors qu’Ubuntu commence à solliciter le swap. C’est pour cette raison que Linux Mint reste recommandé sur tout PC de 2018 ou antérieur, configuration encore majoritaire chez les ménages français selon l’INSEE 2025.

## Temps de démarrage : 13 secondes vs 15 secondes [Mesuré]

Le temps de boot est l’un des critères les plus visibles au quotidien. Selon la commande `systemd-analyze time`, mesurée à dix reprises consécutives et moyennée sur la même configuration de référence avec SSD NVMe, Linux Mint 22.1 démarre en moyenne en **13,1 secondes** jusqu’à un bureau utilisable, contre **15,2 secondes** pour Ubuntu 25.10. L’avantage de Mint vient principalement d’un nombre réduit de services systemd activés par défaut et d’un compositeur Cinnamon plus léger à charger que GNOME Shell.

| Phase de démarrage | Linux Mint 22.1 | Ubuntu 25.10 | 
|---|---|---|
| Firmware (UEFI) | 2,8 s | 2,8 s | 
| Loader (GRUB) | 1,2 s | 1,4 s | 
| Kernel | 3,9 s | 4,1 s | 
| Userspace (systemd) | 5,2 s | 6,9 s | 
| Total démarrage | 13,1 s | 15,2 s | 
| Premier login à bureau prêt | 1,8 s | 2,4 s | 

Pour les utilisateurs qui allument et éteignent leur PC plusieurs fois par jour, ce gain représente environ 12 minutes par mois et près de deux heures par an. Sur un SSD SATA plus ancien ou un disque dur mécanique, l’écart est encore plus marqué et peut atteindre 4 à 6 secondes en faveur de Mint. À l’inverse, sur des SSD NVMe PCIe Gen 4 ou Gen 5 récents, l’écart s’atténue car le goulot d’étranglement se déplace vers la phase firmware UEFI, identique entre les deux distributions.

## Gestion des paquets : APT, Flatpak et la guerre Snap

La gestion des paquets cristallise la principale rupture philosophique entre **Linux Mint vs Ubuntu**. Les deux distributions partagent **APT** comme gestionnaire de paquets principal hérité de Debian, mais leur stratégie pour les applications graphiques diverge complètement. Ubuntu a misé sur **Snap**, un format de paquet universel développé par Canonical, sandboxé et auto-mettant à jour. Snap est désormais le format par défaut pour Firefox, Chromium, le store Ubuntu et un nombre croissant d’applications GNOME dans la version 25.10. Le store officiel Snapcraft.io recensait **15 800 applications** en mars 2026 selon Canonical.

Linux Mint a fait le choix inverse : Snap est **bloqué par défaut** via le fichier `/etc/apt/preferences.d/nosnap.pref`. À la place, Mint privilégie **Flatpak**, un format concurrent piloté par la Linux Foundation et Red Hat. Flathub, le dépôt principal Flatpak, comptait **3 400 applications** en mars 2026 selon le rapport annuel de la GNOME Foundation, mais avec une qualité éditoriale plus stricte. La logithèque Mint privilégie systématiquement la version Flatpak quand elle existe, tout en respectant le choix utilisateur via APT.

```
# Linux Mint : installer Flatpak puis une app depuis Flathub
sudo apt update
flatpak install flathub org.gimp.GIMP
flatpak run org.gimp.GIMP
# Ubuntu 25.10 : installer une app Snap depuis le store
sudo snap install gimp
gimp
# Forcer Ubuntu à utiliser Flatpak (manuel)
sudo apt install flatpak gnome-software-plugin-flatpak
flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
flatpak install flathub org.gimp.GIMP
```
Ce choix philosophique a des conséquences pratiques majeures. Les Snap démarrent en moyenne 28 % plus lentement au premier lancement à cause du processus de montage de l’image squashfs, selon les mesures Phoronix de janvier 2026. Les Flatpak partagent en revanche leurs runtimes (Freedesktop SDK, GNOME, KDE), ce qui réduit la duplication disque. Sur une installation de référence avec 30 applications graphiques courantes, Ubuntu Snap consomme environ **4,8 Go** contre **3,1 Go** pour Mint Flatpak, soit 35 % d’économie.

## Logithèque et expérience d’installation

Côté logithèque, Linux Mint propose la **Software Manager**, une application maison sobre, sans publicité, qui agrège APT, Flatpak et les paquets multimédia (codecs, drivers). L’interface est divisée en catégories claires et l’installation d’une application demande deux clics maximum après l’authentification. Mint Software Manager affiche également la source du paquet (APT, Flatpak) ainsi qu’un score de vérification communautaire pour les Flatpak non officiels.

Ubuntu utilise désormais le **App Center** écrit en Flutter, déployé depuis Ubuntu 24.04 LTS et amélioré dans Ubuntu 25.10. Cette application moderne, multiplateforme dans son architecture, met en avant les Snap par défaut. Elle affiche des notes utilisateurs, des captures d’écran haute résolution et un système de catégories. Les utilisateurs francophones lui reprochent toutefois deux choses : la performance perfectible (temps d’ouverture de 4 à 6 secondes selon les mesures de l’équipe OMG! Ubuntu) et la difficulté à voir clairement la version DEB vs Snap.

