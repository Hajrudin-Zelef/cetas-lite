---
id: collect-261001-general-networking/general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes-5
title: "Sur Fedora / RHEL / Rocky Linux"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes.md
source_anchor: ""
source_lines: [258, 292]
sha256: 621b2f1d8a0eace39fc7fed0d0b68e27c338b03d546902028a2a7e8124c89bb4
---

# Sur Fedora / RHEL / Rocky Linux

Ces outils ne sont pas concurrents mais complémentaires. Sur un poste de travail, Nmap découvre ce qui existe sur le réseau, Wireshark montre ce qui y circule réellement, et Suricata ou Wazuh alertent en continu pendant que personne ne regarde l’écran. La combinaison des trois couvre l’essentiel des besoins d’un administrateur système ou d’un analyste SOC débutant, sans nécessiter de budget logiciel : les cinq outils du tableau sont open source et gratuits.

## Questions fréquentes

### Wireshark est-il légal à utiliser en France ?

Oui, le logiciel lui-même est parfaitement légal et téléchargeable librement. Ce qui peut poser problème, c’est l’usage : capturer du trafic sur un réseau qui ne vous appartient pas, sans autorisation, expose à des poursuites au titre de l’atteinte au secret des correspondances (article 226-15 du Code pénal). Sur votre propre réseau domestique ou professionnel, avec l’accord de l’employeur le cas échéant, aucun problème.

### Wireshark peut-il déchiffrer n’importe quel trafic HTTPS ?

Non. Wireshark ne casse jamais le chiffrement TLS. Il peut afficher le contenu déchiffré uniquement si vous fournissez la clé de session, généralement via la variable `SSLKEYLOGFILE` exportée avant la capture depuis une machine où vous avez déjà un accès légitime. Sans cette clé, le trafic TLS reste illisible, ce qui est le comportement attendu et voulu.

### Quelle est la différence entre Wireshark et Suricata ?

Wireshark est un outil d’analyse manuelle : vous ouvrez une capture et vous cherchez vous-même les anomalies. Suricata est un système de détection et de prévention d’intrusion (IDS/IPS) qui tourne en continu et déclenche des alertes automatiques selon des règles prédéfinies. En pratique, Suricata alerte, Wireshark permet d’investiguer en détail ce qui a déclenché l’alerte.

### Faut-il des connaissances réseau avancées pour démarrer ?

Des bases sur le modèle TCP/IP (adresses IP, ports, notion de paquet) suffisent pour commencer à utiliser l’interface et les filtres simples. La compréhension fine des handshakes TCP, des en-têtes TLS ou des mécanismes DNS s’acquiert progressivement, en pratiquant directement sur des captures réelles plutôt qu’en lisant uniquement de la théorie.

### Pourquoi ma capture ne montre-t-elle pas le trafic d’un autre appareil sur mon réseau Wi-Fi ?

Sur un réseau moderne, filaire ou Wi-Fi, chaque appareil ne reçoit par défaut que son propre trafic. Pour voir le trafic d’un autre appareil, il faut soit un accès physique à un port mirroring sur le switch, soit une carte Wi-Fi compatible avec le mode moniteur, soit une capture directe sur l’appareil cible avec son autorisation.

### Wireshark ralentit-il la machine ou le réseau pendant une capture ?

L’impact reste minime sur un poste ou un serveur avec des ressources normales : la capture consomme surtout du CPU pour le décodage en temps réel si l’interface graphique reste ouverte, et de l’espace disque pour l’enregistrement. Sur un lien très chargé, il est préférable d’utiliser TShark plutôt que l’interface graphique, et de filtrer à la capture pour limiter la charge.

### Quelle est la différence entre le format .pcap et .pcapng ?

Le format .pcap est l’ancien format historique, plus simple mais limité : il ne conserve pas de métadonnées comme les commentaires ou les informations d’interface multiples. Le format .pcapng, adopté par défaut depuis plusieurs versions de Wireshark, permet de stocker plusieurs interfaces dans un seul fichier, d’ajouter des commentaires et de mieux gérer les captures de grande taille. Sauf contrainte de compatibilité avec un outil tiers ancien, utilisez toujours .pcapng.

### Peut-on utiliser Wireshark sur un Raspberry Pi ou un NAS ?

TShark s’installe sans difficulté sur un Raspberry Pi ou un NAS sous Linux, pour de la capture légère en ligne de commande. L’interface graphique complète de Wireshark fonctionne aussi mais reste plus confortable sur un poste avec davantage de RAM, surtout pour l’analyse de captures volumineuses. Une approche courante consiste à capturer sur le petit appareil avec TShark, puis à transférer le fichier .pcapng vers un poste plus puissant pour l’analyse détaillée dans l’interface graphique.
