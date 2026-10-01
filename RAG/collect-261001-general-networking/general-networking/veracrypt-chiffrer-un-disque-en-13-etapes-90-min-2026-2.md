---
id: collect-261001-general-networking/general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026-2
title: "Windows (PowerShell)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Apple", "Intel"]
dates: []
keywords: ["amd", "asic", "distribution", "gpu", "intel", "mai"]
source: docs/RAG/collect-261001-general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [45, 104]
sha256: 443e0bd08d19d7170af58c4aa7d5580dc0124aa2cf8d9fab47ede1e3e02ba44f
---

# Windows (PowerShell)

VeraCrypt propose plusieurs algorithmes de chiffrement, combinables en cascade avec un ou deux autres pour ajouter des couches de protection supplémentaires, au prix d’une perte de performance. Le mode de fonctionnement utilisé est XTS, spécifiquement conçu pour le chiffrement de disque. La fonction de dérivation Argon2id a fait ses débuts dans la version 1.26.27 du 20 septembre 2025, pour le hachage du mot de passe et des montages plus rapides, avant d’être étendue aux volumes non-système avec la 1.26.29. Le tableau ci-dessous liste les algorithmes disponibles dans cette version.

| Type | Options disponibles | Recommandation | 
|---|---|---|
| Chiffrement (simple) | AES, Serpent, Twofish, Camellia, Kuznyechik | AES pour la majorité des usages (accélération matérielle AES-NI) | 
| Chiffrement (cascade) | AES(Twofish), Serpent(AES), Twofish(Serpent), AES(Twofish(Serpent)), et autres combinaisons | À réserver aux données très sensibles, impact sur la vitesse | 
| Fonctions de hachage | SHA-256, SHA-512, Streebog, Whirlpool, BLAKE2s-256 | SHA-512 par défaut, solide et rapide | 
| Dérivation de clé (nouveauté 1.26.29) | Argon2id (volumes non-système), PBKDF2 historique | Argon2id quand disponible : plus résistant au craquage par GPU/ASIC | 
| Mode de fonctionnement | XTS | Seul mode proposé, standard pour le chiffrement de disque | 

Pour un usage courant, l’algorithme AES normalisé par le NIST américain, combiné à SHA-512, constitue le meilleur compromis entre sécurité et performance, grâce à l’accélération matérielle AES-NI présente sur la quasi-totalité des processeurs Intel et AMD depuis plus d’une décennie. Depuis la version 1.26.20 du 3 février 2025, cette logique d’accélération matérielle s’étend aussi aux processeurs ARM64 pour la fonction de hachage SHA-256, un gain notable sur les Mac Apple Silicon et les machines Linux ARM. Les cascades comme AES(Twofish(Serpent)) ajoutent une marge de sécurité théorique contre la découverte future d’une faiblesse sur un algorithme unique, mais elles ralentissent sensiblement les opérations de lecture et d’écriture, sans bénéfice pratique pour l’utilisateur grand public.

## Étape 1 : Télécharger VeraCrypt et vérifier son authenticité

Téléchargez toujours VeraCrypt depuis le site officiel du projet, veracrypt.io (l’ancien domaine veracrypt.fr redirige désormais vers cette adresse). Un logiciel de chiffrement téléchargé depuis un site tiers ou un miroir non officiel peut avoir été modifié pour intégrer une porte dérobée : c’est précisément le scénario que l’outil est censé empêcher. Une fois le fichier récupéré, vérifiez sa signature PGP ou, a minima, son empreinte SHA-256, publiée sur la page officielle des téléchargements. Ce n’est pas une précaution superflue : en janvier 2025, le site gHacks a documenté la correction de deux failles de sécurité, référencées CVE-2024-54187 et CVE-2025-23021, avec la sortie de la version 1.26.18 pour Linux et macOS, un rappel que mieux vaut toujours installer la dernière version disponible plutôt qu’un exécutable oublié dans un dossier de téléchargements.

Sous Windows ou macOS, ouvrez un terminal (PowerShell ou Terminal.app) et calculez l’empreinte du fichier téléchargé :

```
# Windows (PowerShell)
Get-FileHash .\VeraCrypt_Setup_1.26.29.exe -Algorithm SHA256
# macOS / Linux
shasum -a 256 VeraCrypt_1.26.29.dmg
```
Comparez la sortie obtenue caractère par caractère avec l’empreinte publiée sur le site officiel. Si les deux ne correspondent pas, ne procédez pas à l’installation : supprimez le fichier et retéléchargez-le depuis une connexion différente pour écarter un problème de proxy d’entreprise ou de cache DNS corrompu.

## Étape 2 : Installer VeraCrypt sous Windows, macOS et Linux

Sous Windows, lancez l’exécutable téléchargé, acceptez la licence et laissez les options par défaut sauf besoin spécifique (langue de l’interface, dossier d’installation). Depuis la version 1.26.24, publiée le 30 mai 2025, l’installateur Windows embarque aussi une protection d’écran supplémentaire contre la capture de mot de passe, activable dans les options avancées. Sous macOS, montez l’image .dmg puis suivez l’assistant d’installation ; macOS demandera d’autoriser l’extension système de VeraCrypt dans Réglages Système > Confidentialité et sécurité, une étape facile à manquer et qui bloque silencieusement le montage de volumes si elle est ignorée.

### Installation sous Linux en ligne de commande

Sous Linux, VeraCrypt est distribué sous forme d’un script d’installation autonome (console ou GUI) plutôt que via les dépôts officiels des distributions. Depuis la version 1.26.24 du 30 mai 2025, un format AppImage est également proposé en alternative, pratique si vous préférez éviter le script d’installation classique. Téléchargez l’archive correspondant à votre architecture, puis exécutez le script d’installation graphique :

```
tar xjf veracrypt-1.26.29-setup.tar.bz2
./veracrypt-1.26.29-setup-gui-x64
# Vérifier l'installation
veracrypt --text --version
```
Sur les distributions qui appliquent AppArmor ou SELinux de façon stricte (Fedora, openSUSE), il arrive que le module FUSE nécessaire au montage des volumes soit bloqué par défaut. Si le montage échoue à l’étape suivante, installez le paquet `fuse` ou `fuse3` via le gestionnaire de paquets de votre distribution avant de continuer.

## Étape 3 : Créer votre premier conteneur chiffré (volume standard)

Lancez VeraCrypt et cliquez sur **Create Volume** (Créer un volume). Choisissez la première option, « Create an encrypted file container », pour créer un conteneur : un simple fichier qui se comportera comme un disque virtuel une fois monté. C’est l’option la plus flexible pour débuter, car elle ne touche à aucune partition existante et peut être supprimée à tout moment en effaçant simplement le fichier.

- Sélectionnez **Standard VeraCrypt volume** pour ce premier essai (le volume caché sera abordé à l’étape 9)
- Choisissez l’emplacement et le nom du fichier conteneur, par exemple `coffre.hc`
- Sélectionnez l’algorithme de chiffrement (AES par défaut convient à la majorité des cas) et la fonction de hachage (SHA-512)
- Indiquez la taille du volume : commencez par un petit conteneur de test, par exemple 500 Mo, avant de créer vos volumes définitifs
- Choisissez un mot de passe long et unique, idéalement une phrase de passe de 20 caractères ou plus. VeraCrypt affiche un avertissement si le mot de passe est jugé trop court
- Bougez la souris de façon aléatoire dans la fenêtre pendant la génération des clés : ce mouvement alimente le générateur de nombres aléatoires
- Choisissez un système de fichiers (exFAT pour un usage multiplateforme Windows/macOS/Linux, NTFS si le volume reste sur Windows)

Cliquez sur **Format** pour lancer la création. Pour un conteneur de 500 Mo, l’opération prend généralement moins d’une minute.

## Étape 4 : Monter et démonter un volume au quotidien

Dans la fenêtre principale de VeraCrypt, sélectionnez une lettre de lecteur libre (Windows) ou un point de montage (macOS/Linux), cliquez sur **Select File** pour pointer vers votre conteneur, puis sur **Mount**. Saisissez votre mot de passe : le volume apparaît alors comme un disque normal dans l’explorateur de fichiers, avec sa propre lettre ou son propre point de montage.

