---
id: collect-261001-general-networking/general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026-2
title: "Calculer l'empreinte SHA-256 de l'installeur téléchargé"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attention", "mai", "nand"]
source: docs/RAG/collect-261001-general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026.md
source_anchor: ""
source_lines: [39, 99]
sha256: d1f515563b2237e3325d471bd400aeb461bbd85bde1537486fd9da86586546cd
---

# Calculer l'empreinte SHA-256 de l'installeur téléchargé

La règle d’or : ne téléchargez **jamais** CrystalDiskInfo depuis un portail de téléchargement tiers. Ces sites empaquettent régulièrement l’installeur avec des logiciels indésirables (barres d’outils, « optimiseurs » douteux). Rendez-vous exclusivement sur le site officiel Crystal Dew World, sur le dépôt SourceForge du projet, ou via le Microsoft Store (édition sans publicité, mise à jour automatiquement).

Sur la page de téléchargement, vous rencontrerez deux décisions. La première : quelle **édition** – choisissez « CrystalDiskInfo » (Standard), sans suffixe. La seconde : **installeur ou version portable**. L’installeur (fichier `.exe`) s’intègre au menu Démarrer et gère les mises à jour ; la version portable (archive `.zip`) s’extrait dans un dossier et se lance sans installation – idéale pour une clé USB de dépannage ou un poste verrouillé, et plébiscitée par les utilisateurs (la build Portable 9.9.0.0 du 19 mai 2026 affiche une note de 5/5 sur Uptodown, avec 1 343 téléchargements hebdomadaires sur SourceForge contre 2 376 pour l’installeur EXE équivalent de la même version). Notez que les installeurs récents incluent parfois une offre sponsorisée : la build « 9.9.1 Ads » du 5 juin 2026 pèse par exemple 15,9 Mo – contre 6,2 Mo pour l’EXE classique – et totalisait pourtant 57 963 téléchargements hebdomadaires sur SourceForge. Si vous voulez un exécutable strictement sans publicité, préférez plutôt la dernière archive ZIP portable, entièrement propre.

Vérifiez toujours l’empreinte du fichier lorsque la source la publie. Sous PowerShell, une commande suffit à confirmer que le binaire téléchargé n’a pas été altéré :

```
# Calculer l'empreinte SHA-256 de l'installeur téléchargé
Get-FileHash -Algorithm SHA256 "$env:USERPROFILE\Downloads\CrystalDiskInfo9_9_1.exe"
# Sortie attendue (exemple) :
# Algorithm  Hash                              Path
# ---------  ----                              ----
# SHA256     3F9A1C...E7B2                     C:\Users\vous\Downloads\CrystalDiskInfo9_9_1.exe
```
## Étape 2 – Installer CrystalDiskInfo (ou utiliser la version portable)

Si vous avez choisi l’installeur, double-cliquez sur le `.exe` et acceptez l’invite du Contrôle de compte d’utilisateur (UAC) – cette élévation est indispensable, car la lecture S.M.A.R.T. exige un accès matériel de bas niveau. Suivez l’assistant : acceptez la licence, conservez le dossier d’installation par défaut, et surveillez chaque écran. **Décochez toute offre optionnelle** (logiciel additionnel, page d’accueil, extension de navigateur) qui pourrait apparaître. L’installation se termine en quelques secondes.

Pour la version portable, extrayez simplement l’archive ZIP dans un dossier de votre choix, par exemple `C:\Outils\CrystalDiskInfo\`. Le dossier contient plusieurs exécutables : choisissez celui adapté à votre système. La convention de nommage est la suivante :

```
CrystalDiskInfo\
├── DiskInfo.exe        # Version 32 bits (x86)
├── DiskInfo64.exe      # Version 64 bits (x64) – a privilegier
├── DiskInfoA64.exe     # Version ARM64 (Windows on ARM)
├── DiskInfo.ini        # Fichier de configuration (cree au 1er lancement)
└── Language\           # Fichiers de langue (dont Francais.lang)
```
Lancez `DiskInfo64.exe` sur un PC 64 bits moderne. Astuce pour les administrateurs : la version portable est parfaite pour un déploiement en masse via GPO ou script, puisqu’elle ne laisse aucune trace dans le registre et se pilote entièrement par son fichier `DiskInfo.ini` – nous exploiterons cette propriété à l’étape 10.

## Étape 3 – Premier lancement et vue d’ensemble de l’interface

Au premier démarrage, CrystalDiskInfo interroge tous les disques détectés et affiche sa fenêtre principale. Prenez le temps de la décoder, car chaque zone a un rôle précis :

- **La barre d’onglets supérieure** : un onglet par disque détecté. Cliquez pour basculer d’un disque à l’autre. La couleur de l’onglet reflète déjà l’état de santé.
- **Le cartouche « État de santé »** (à gauche) : le verdict global – Bon, Attention ou Mauvais – accompagné d’un pourcentage (surtout pertinent pour les SSD, où il reflète l’usure restante).
- **Le cartouche « Température »** : la température instantanée du disque, en °C.
- **Le bandeau d’informations** : modèle, micrologiciel (firmware), numéro de série, capacité, interface (Serial ATA, NVM Express…), mode de transfert négocié, lettre de lecteur et fonctionnalités supportées (S.M.A.R.T., APM, AAM, TRIM…).
- **Le tableau S.M.A.R.T.** (partie basse) : la liste complète des attributs, cœur du diagnostic, que nous détaillons à l’étape 5.

Première action recommandée : passer l’interface en français. Ouvrez le menu `Language` (ou `言語`) et sélectionnez **Français**. Toute l’interface, y compris les libellés d’attributs, bascule alors dans la langue de Molière. Ensuite, si vous gérez plusieurs disques, notez que l’ordre des onglets suit l’énumération matérielle, pas les lettres de lecteur : le disque « (1) » n’est pas forcément votre lecteur C:.

## Étape 4 – Lire et interpréter l’état de santé (Bon / Attention / Mauvais)

Le cartouche « État de santé » est la synthèse que 90 % des utilisateurs consulteront. CrystalDiskInfo agrège l’ensemble des attributs S.M.A.R.T. et les compare à leurs seuils constructeur pour produire un verdict codé par couleur. Voici la grille de lecture complète :

| État | Couleur | Signification | Action recommandée | 
|---|---|---|---|
| Bon | Bleu (ou vert) | Tous les attributs sont dans les limites normales | Aucune ; surveillance périodique | 
| Attention | Jaune | Un attribut a franchi un seuil d’alerte (secteurs réalloués, en attente, usure élevée…) | Sauvegarder immédiatement, planifier le remplacement | 
| Mauvais | Rouge | Un attribut critique a dépassé le seuil de panne | Cloner d’urgence, cesser tout usage important | 
| Inconnu | Gris | Le S.M.A.R.T. n’est pas lisible (pont USB, RAID, disque virtuel) | Brancher en SATA/NVMe direct pour diagnostiquer | 

Le point le plus mal compris concerne l’état **Attention** (jaune). Il ne signifie pas « panne imminente dans l’heure », mais « un compteur a commencé à se dégrader ». Sur un disque dur, quelques secteurs réalloués peuvent rester stables des années ; en revanche, si le compteur *augmente* à chaque analyse, la spirale est enclenchée. La règle pratique : ce n’est pas la valeur absolue qui compte, mais sa **tendance**. Un état Attention stable se surveille ; un état Attention qui progresse impose le remplacement.

Sur les SSD, le pourcentage de santé mérite une lecture nuancée. Il reflète l’usure des cellules NAND (endurance d’écriture), pas un défaut mécanique. Un SSD à 90 % après trois ans est parfaitement sain ; il ne tombera « Attention » que lorsque l’usure approchera de la limite garantie (le fameux TBW, *Total Bytes Written*). Nous y revenons en détail à l’étape 11.

## Étape 5 – Décrypter les attributs S.M.A.R.T. essentiels

Le tableau S.M.A.R.T. affiche pour chaque attribut : un identifiant (ID en hexadécimal), les valeurs *Actuelle* / *Pire* / *Seuil* (normalisées de 1 à 253, où plus haut = mieux) et une colonne **Valeurs brutes** (RawValues) – c’est cette dernière, en clair, qui vous intéresse vraiment. La norme S.M.A.R.T. définit des dizaines d’attributs, mais une poignée suffit à diagnostiquer 95 % des cas. Pour un disque dur mécanique :

