---
id: collect-261001-general-networking/general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes-2
title: "Sur Fedora / RHEL / Rocky Linux"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "ethernet"]
source: docs/RAG/collect-261001-general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes.md
source_anchor: ""
source_lines: [40, 124]
sha256: 8413264c2d5009bb32a8ff8376d9b762fa575e2ef89c87ecdf8c29d808e1ad44
---

# Sur Fedora / RHEL / Rocky Linux

```
sudo apt update
sudo apt install wireshark tshark
# Sur Fedora / RHEL / Rocky Linux
sudo dnf install wireshark wireshark-cli
# Sur Arch Linux
sudo pacman -S wireshark-qt wireshark-cli
```
Sur Debian et Ubuntu, l’installateur demande si les utilisateurs non root doivent pouvoir capturer des paquets. Répondez “Oui” uniquement si vous comprenez l’implication : cela crée le groupe `wireshark` avec des droits élevés sur l’interface de capture. Vous configurerez ce point précisément à l’étape suivante.

### Étape 2 : installer Npcap ou configurer libpcap sans lancer Wireshark en root

Sous Windows, Npcap 1.88 s’installe généralement en même temps que Wireshark. Si ce n’est pas le cas, téléchargez-le séparément depuis npcap.com, le projet maintenu par l’équipe Nmap qui a repris le flambeau de WinPcap. Cochez l’option “Support raw 802.11 traffic” uniquement si vous comptez analyser du Wi-Fi en mode moniteur, et laissez l’option “Restrict Npcap driver’s access to Administrators only” décochée pour permettre à un utilisateur standard membre du bon groupe de capturer sans élévation permanente.

Sous Linux, la bonne pratique consiste à ajouter votre utilisateur au groupe `wireshark` plutôt que de lancer l’application avec `sudo` à chaque fois, ce qui exposerait l’intégralité de l’interface graphique (et ses bibliothèques de décodage) avec les droits root :

```
sudo usermod -aG wireshark $USER
# Déconnexion/reconnexion nécessaire pour appliquer le groupe
# Vérifier les permissions sur dumpcap
getcap /usr/bin/dumpcap
# Résultat attendu : cap_net_admin,cap_net_raw+eip
```
Si `getcap` ne renvoie rien, les capacités n’ont pas été correctement attribuées lors de l’installation. Corrigez-le manuellement avec `sudo setcap cap_net_raw,cap_net_admin=eip /usr/bin/dumpcap`. Cette approche limite l’élévation de privilèges au seul binaire `dumpcap`, celui qui capture réellement les paquets, pendant que l’interface graphique tourne avec vos droits normaux.

### Étape 3 : choisir la bonne interface réseau

Au lancement, l’écran d’accueil de Wireshark liste toutes les interfaces disponibles avec un mini-graphique d’activité en temps réel : Ethernet, Wi-Fi, interfaces virtuelles VPN, boucle locale (loopback), et parfois des interfaces Bluetooth ou USB. Sur un poste connecté en filaire, choisissez l’interface Ethernet active (celle qui affiche du trafic). Sur un serveur avec plusieurs cartes réseau, identifiez la bonne interface via `ip addr` (Linux) ou `ipconfig` (Windows) avant de lancer la capture, pour éviter de scruter une interface inactive pendant dix minutes sans rien voir passer.

## Étapes 4 et 5 : première capture et prise en main de l’interface

### Étape 4 : lancer votre première capture de paquets

Double-cliquez sur l’interface choisie, ou sélectionnez-la puis cliquez sur l’icône en forme de requin bleu. Le trafic commence à défiler immédiatement. Pour un premier test, ouvrez un navigateur et visitez un site quelconque : vous verrez apparaître des requêtes DNS, un handshake TCP, puis du trafic TLS. Cliquez sur le carré rouge pour arrêter la capture une fois que vous avez quelques centaines de paquets, largement suffisant pour s’exercer.

Sauvegardez immédiatement cette première capture via Fichier > Enregistrer sous, au format .pcapng (le format natif de Wireshark, qui conserve les métadonnées et les commentaires, contrairement à l’ancien .pcap). Prenez l’habitude de nommer vos fichiers avec la date et le contexte, par exemple `2026-09-01_test-dns.pcapng`, cela évite de se perdre dans un dossier rempli de fichiers nommés `capture1`, `capture2`.

### Étape 5 : comprendre l’interface en trois volets

L’écran principal se divise en trois zones empilées. En haut, la liste des paquets (Packet List) affiche une ligne par paquet avec le numéro, l’horodatage, les adresses source et destination, le protocole détecté et une description résumée. Au milieu, les détails du paquet (Packet Details) montrent la structure en couches : trame Ethernet, en-tête IP, en-tête TCP ou UDP, puis le protocole applicatif, chaque couche pouvant être dépliée pour voir chaque champ individuellement. En bas, les octets bruts (Packet Bytes) affichent l’hexadécimal et sa représentation ASCII, avec surbrillance synchronisée : cliquez sur un champ dans les détails, et l’octet correspondant s’illumine en bas.

Les couleurs de fond dans la liste de paquets ne sont pas décoratives : le vert clair signale généralement du TCP standard, le bleu du DNS, le noir avec du rouge des paquets marqués comme erreurs ou retransmissions par le moteur d’analyse d’expert (Expert Info). Ces règles de coloration, personnalisables via Affichage > Règles de coloration, permettent de repérer un problème d’un coup d’œil avant même de lire le détail.

## Étape 6 : appliquer un filtre de capture (BPF)

Sur un serveur ou un lien saturé, capturer tout le trafic sans filtre génère des fichiers énormes et noie l’information utile. Les filtres de capture s’appliquent avant l’enregistrement, au niveau du pilote (BPF, Berkeley Packet Filter), et rejettent les paquets non voulus avant même qu’ils touchent le disque. Ils se saisissent dans le champ “Capture Filter” de l’écran de démarrage, avec une syntaxe différente de celle des filtres d’affichage.

```
# Ne capturer que le trafic HTTP et HTTPS
port 80 or port 443
# Ne capturer que le trafic vers ou depuis une machine précise
host 192.168.1.50
# Exclure tout le trafic SSH pour ne pas polluer une capture de debug distant
not port 22
# Ne capturer que le trafic d'un sous-réseau donné
net 10.0.0.0/24
# Combiner : trafic HTTP uniquement depuis un hôte donné
host 192.168.1.50 and port 80
```
Un filtre de capture mal écrit ne provoque pas d’erreur visible : Wireshark capture simplement zéro paquet, ou tout le trafic si la syntaxe est invalide et silencieusement ignorée selon la plateforme. Testez toujours un filtre de capture sur quelques secondes avant de lancer un enregistrement de plusieurs heures.

## Étape 7 : maîtriser les filtres d’affichage

Les filtres d’affichage, eux, s’appliquent après la capture, sur des paquets déjà enregistrés. Ils utilisent une syntaxe propre à Wireshark, bien plus riche que celle des filtres de capture, et se saisissent dans la barre verte en haut de la fenêtre principale (elle devient rouge si la syntaxe est invalide). C’est l’outil que vous utiliserez 90 % du temps.

```
# Isoler tout le trafic HTTP
http
# Isoler les requêtes DNS uniquement (pas les réponses)
dns.flags.response == 0
# Trouver les paquets marqués comme problématiques par Wireshark
tcp.analysis.retransmission or tcp.analysis.duplicate_ack
# Filtrer par adresse IP source
ip.src == 10.0.0.15
# Filtrer les échecs de connexion TCP (paquets RST)
tcp.flags.reset == 1
# Repérer un scan de ports classique (paquets SYN sans réponse)
tcp.flags.syn == 1 and tcp.flags.ack == 0
# Filtrer le trafic TLS et afficher la version négociée
tls.handshake.type == 1
```
Exemple de sortie attendue après avoir tapé `tcp.analysis.retransmission` sur une capture avec des soucis réseau : la liste se réduit d’un coup à une poignée de lignes surlignées en noir, chacune indiquant “[TCP Retransmission]” dans la colonne Info, avec le numéro de séquence dupliqué visible dans les détails du paquet. C’est exactement ce type de filtre qui permet de prouver, en quelques secondes, qu’un problème de lenteur applicative vient du réseau et non du code.

Astuce pour les débutants : clic droit sur n’importe quel champ dans le volet des détails, puis “Apply as Filter” > “Selected”, génère automatiquement la syntaxe correcte. C’est la méthode la plus rapide pour apprendre le langage des filtres sans mémoriser la documentation.

