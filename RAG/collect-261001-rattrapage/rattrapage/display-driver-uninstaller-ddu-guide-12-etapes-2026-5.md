---
id: collect-261001-rattrapage/rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026-5
title: "Calculer l'empreinte SHA-256 de l'archive (PowerShell)"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "intel", "mai", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026.md
source_anchor: ""
source_lines: [307, 349]
sha256: 103c08560ff61dc50c143f4dec6a1b0f7c1561ae37e1d22818614876f1a765c5
---

# Calculer l'empreinte SHA-256 de l'archive (PowerShell)

Ensuite, l’option **« Nettoyer le cache des shaders » seule**. Inutile de tout effacer pour régler un problème de saccades dans un jeu précis : DDU peut purger uniquement le cache de shaders DirectX/Vulkan, souvent responsable de stutters après une mise à jour de pilote. C’est un nettoyage non destructif, sans redémarrage.

Enfin, pensez à **conserver une copie de DDU sur une clé USB** avec votre dernier pilote stable. En cas d’écran noir total empêchant tout accès au réseau, vous pourrez nettoyer et réinstaller hors ligne depuis le mode sans échec. Pour les passionnés de réglage fin, l’enchaînement logique reste : pilotes propres avec DDU, puis overclocking et undervolting – un terrain que nous couvrons dans notre guide dédié à MSI Afterburner.

Pour les configurations à double GPU (iGPU Intel ou AMD intégré + carte dédiée), un dernier conseil : nettoyez et réinstallez les pilotes dans l’ordre, du plus simple au plus complexe. Traitez d’abord le GPU intégré, redémarrez, puis la carte dédiée. Des outils tiers comme GPU-Z ou HWiNFO confirment, après coup, que seul le bon pilote est chargé et que la carte fonctionne à pleine vitesse PCIe (et non en mode dégradé x1). Cette rigueur évite les conflits de pilotes hybrides qui provoquent des micro-saccades difficiles à diagnostiquer, en particulier sur les ordinateurs portables gamer où le commutateur graphique (Optimus, par exemple) ajoute une couche de complexité.

### Related Coverage

## FAQ : Display Driver Uninstaller (DDU)

### DDU est-il sûr et gratuit ?

Oui. Display Driver Uninstaller est entièrement gratuit, développé par Wagnard depuis plus de dix ans, et largement recommandé par les communautés matérielles. Téléchargé depuis le site officiel Wagnardsoft ou le miroir Guru3D, il ne contient aucun logiciel indésirable. Comme il modifie le registre, créez un point de restauration au préalable, par simple prudence.

### Faut-il obligatoirement utiliser le mode sans échec ?

Ce n’est pas obligatoire, mais c’est fortement recommandé. En mode normal, Windows verrouille des fichiers du pilote, ce qui empêche une suppression à 100 %. Le mode sans échec garantit un nettoyage complet en un seul passage. En mode normal, il faudrait répéter le cycle nettoyer/redémarrer deux fois pour un résultat comparable.

### DDU fonctionne-t-il avec les cartes Intel Arc ?

Oui. Display Driver Uninstaller prend en charge les pilotes NVIDIA, AMD et Intel, y compris les cartes graphiques dédiées Intel Arc et les iGPU Intel. Sélectionnez simplement « Intel » comme fabricant dans le menu déroulant, puis installez ensuite le pilote Arc & Graphics le plus récent.

### Pourquoi mon code PIN ne fonctionne-t-il pas en mode sans échec ?

Sur Windows 11 24H2, Windows Hello (et donc le code PIN) peut être indisponible en mode sans échec, car certains services d’identité ne sont pas chargés. Connectez-vous avec votre mot de passe de compte Microsoft via « Options de connexion », ou démarrez en mode sans échec avec prise en charge réseau. Préparez toujours votre mot de passe avant l’opération.

### DDU supprime-t-il aussi les pilotes audio ?

Partiellement. DDU peut nettoyer certains pilotes audio liés au GPU (audio HDMI/DisplayPort) et, depuis les versions 18.1.5.x, les pilotes Realtek et Sound Blaster via les commutateurs `-cleanrealtek` et `-cleansoundblaster`. Il ne remplace toutefois pas un désinstalleur audio dédié pour les cartes son complexes.

### Dois-je débrancher Internet pendant l’opération ?

C’est la méthode la plus sûre. Tant que vous n’avez pas réinstallé votre pilote, Windows Update tentera de télécharger un pilote générique automatiquement. Couper le réseau (ou appliquer les clés de registre de l’étape 6) empêche cette réinstallation parasite qui annulerait votre nettoyage.

### Quelle est la dernière version de DDU en 2026 ?

Au 17 août 2026, la version stable est DDU 18.1.5.6, publiée le 17 juillet 2026 par Wagnardsoft. Elle succède à un rythme soutenu de mises à jour depuis la 18.1.3.7 d’octobre 2025 (couverte par iTechGuides), en passant par la 18.1.3.9 de décembre 2025 signalée par Guru3D, la 18.1.4.0 puis la 18.1.4.1 de janvier 2026, la 18.1.4.2 de février 2026, la 18.1.5.2 d’avril 2026 (qui a introduit l’exigence de .NET Framework 4.8) et la 18.1.5.3 de mai 2026. Il s’agit d’une mise à jour de maintenance améliorant la suppression des services, le programme d’installation et la prise en charge audio en ligne de commande. Vérifiez toujours la page Wagnardsoft ou Guru3D pour la toute dernière révision.

### À quelle fréquence faut-il utiliser DDU ?

Pas systématiquement. La mise à jour « par-dessus » convient pour les pilotes courants. Réservez DDU aux situations problématiques : changement de fabricant, pilote corrompu, plantages persistants, retour à une version antérieure, ou avant une session sérieuse d’overclocking. L’utiliser à chaque mise à jour est inutile et fait perdre du temps.

*Tutoriel publié le 29 juin 2026. Versions logicielles et procédures vérifiées à cette date ; reportez-vous aux sites officiels des éditeurs pour les dernières révisions.*
