---
id: collect-261001-general-networking/general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes-3
title: "Sur Fedora / RHEL / Rocky Linux"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes.md
source_anchor: ""
source_lines: [125, 195]
sha256: 88c83a8a974447094b0c96bfba1ed9ff9d06b42a9c9d4e3b470e3ff953b345d0
---

# Sur Fedora / RHEL / Rocky Linux

## Étape 8 : suivre un flux TCP, HTTP ou TLS

Un flux réseau se compose de dizaines, parfois de milliers de paquets individuels. Reconstituer une conversation complète paquet par paquet serait ingérable. La fonction “Follow Stream” (clic droit sur un paquet > Follow > TCP Stream, ou HTTP Stream, ou TLS Stream) réassemble automatiquement tous les paquets d’une même connexion et les affiche comme une conversation lisible, avec le trafic client en rouge et le trafic serveur en bleu.

Pour du trafic HTTP en clair, cette vue affiche directement les en-têtes de requête, le corps de la réponse, les cookies, tout en texte lisible. Pour du trafic chiffré en TLS, le contenu applicatif reste illisible par défaut : Wireshark voit les paquets, mais pas leur contenu déchiffré. Pour déboguer une application qui utilise HTTPS, la méthode standard consiste à exporter la clé de session TLS depuis le navigateur ou l’application, via la variable d’environnement `SSLKEYLOGFILE`, puis à la charger dans Wireshark.

```
# Sous Linux/macOS, avant de lancer le navigateur
export SSLKEYLOGFILE=~/tls-keys.log
firefox &
# Sous Windows PowerShell
$env:SSLKEYLOGFILE = "C:\Users\vous\tls-keys.log"
# Puis dans Wireshark :
# Édition > Préférences > Protocols > TLS
# Champ "(Pre)-Master-Secret log filename" -> pointer vers tls-keys.log
```
Une fois ce fichier chargé, Wireshark déchiffre le trafic TLS à la volée pour les sessions capturées après la création du fichier de clés, ce qui permet de voir le contenu HTTP réel derrière le HTTPS, sans jamais casser le chiffrement lui-même : la technique fonctionne uniquement parce que vous détenez déjà la clé côté client, légitimement.

## Étape 9 : lire les statistiques réseau

Le menu Statistiques regroupe des vues agrégées qui évitent de parcourir manuellement des milliers de lignes. Trois d’entre elles sont incontournables au quotidien. “Protocol Hierarchy” affiche la répartition du trafic capturé par protocole, en pourcentage et en nombre de paquets : utile pour repérer en un coup d’œil qu’un poste envoie 40 % de son trafic en DNS, ce qui est anormal et évoque potentiellement de l’exfiltration par tunneling DNS.

“Conversations” liste chaque paire source/destination avec le volume échangé, la durée et le débit, triable par colonne : la méthode la plus rapide pour identifier quelle machine consomme le plus de bande passante sur un lien saturé. “Expert Information” agrège automatiquement tous les paquets que le moteur de Wireshark juge suspects ou problématiques (retransmissions, fenêtres TCP à zéro, checksums invalides), classés par sévérité (Chat, Note, Warning, Error), ce qui donne un point de départ immédiat sur une capture inconnue de plusieurs milliers de paquets.

Le graphique “IO Graph” (Statistiques > IO Graph) trace le débit dans le temps et permet de superposer plusieurs filtres sur la même courbe, par exemple le trafic total contre le seul trafic d’erreur, pour visualiser instantanément à quel moment précis un incident a commencé.

## Étape 10 : repérer les anomalies et signes d’intrusion

Wireshark ne génère pas d’alerte automatique, mais certains motifs de trafic sont des signaux classiques une fois qu’on sait où regarder. Le tableau suivant liste les cas les plus fréquents rencontrés en analyse post-incident.

| Signe observé | Filtre d’affichage utile | Interprétation possible | 
|---|---|---|
| Rafale de paquets SYN vers des ports variés depuis une même source | `tcp.flags.syn==1 and tcp.flags.ack==0` | Scan de ports (type Nmap) | 
| Réponses ARP multiples pour une même IP avec des adresses MAC différentes | `arp.duplicate-address-detected` | ARP spoofing / attaque de l’homme du milieu | 
| Volume anormal de requêtes DNS TXT ou de sous-domaines aléatoires longs | `dns.qry.type==16` | Tunneling DNS / exfiltration de données | 
| Multiples tentatives d’authentification SSH depuis une même IP en quelques secondes | `tcp.port==22 and tcp.flags.syn==1` | Brute-force SSH | 
| Connexions sortantes régulières vers une IP externe à intervalles fixes | `ip.addr==X.X.X.X` + tri par temps | Balise de commande et contrôle (C2 beaconing) | 
| Certificat TLS auto-signé ou domaine récemment enregistré | `tls.handshake.type==11` | Serveur de phishing ou infrastructure malveillante | 

Aucun de ces motifs n’est une preuve à lui seul : un pic de requêtes SSH peut être un script légitime de supervision, un certificat auto-signé peut appartenir à un service interne. L’intérêt de Wireshark est de fournir la preuve matérielle qui confirme ou infirme une alerte déjà remontée par un IDS comme Suricata, un SIEM comme Wazuh, ou un outil de blocage comportemental comme CrowdSec.

## Étapes 11 à 13 : automatiser avec TShark, exporter et respecter le cadre légal

### Étape 11 : automatiser l’analyse avec TShark

TShark est la version ligne de commande de Wireshark, installée avec le même paquet. Elle partage le même moteur de décodage et la même syntaxe de filtres, ce qui la rend idéale pour scripter des captures sur un serveur sans interface graphique, ou pour intégrer une extraction de données dans un pipeline d’automatisation.

```
# Capturer 500 paquets sur l'interface eth0 et sauvegarder
tshark -i eth0 -c 500 -w capture.pcapng
# Lire un fichier existant et n'afficher que le trafic HTTP
tshark -r capture.pcapng -Y "http"
# Extraire uniquement les adresses IP source et destination, format CSV
tshark -r capture.pcapng -T fields -e ip.src -e ip.dst -E separator=,
# Capturer en continu avec rotation de fichiers (utile en production)
tshark -i eth0 -b filesize:100000 -b files:10 -w /var/log/pcap/rotation.pcapng
```
La rotation de fichiers (`-b filesize:100000 -b files:10`) évite de saturer le disque sur un serveur en capture permanente : ici, dix fichiers de 100 Mo maximum, les plus anciens étant écrasés automatiquement. C’est la configuration de base pour un poste de capture longue durée destiné à l’investigation a posteriori.

### Étape 12 : exporter et documenter une capture

Pour un rapport d’incident ou une preuve destinée à un client, exportez toujours au format .pcapng plutôt qu’en copie d’écran : le fichier brut permet à un tiers de rejouer votre analyse et de vérifier vos conclusions. Fichier > Exporter des objets permet aussi d’extraire directement les fichiers transférés en clair (images, documents, exécutables) détectés dans une capture HTTP ou SMB, une fonction régulièrement utilisée en réponse à incident pour récupérer un malware capturé en transit.

Ajoutez systématiquement un commentaire de capture (Édition > Commentaire de paquet, ou Statistiques > Résumé de capture) précisant le contexte : date, machine, objectif de la capture, autorisation associée. Un fichier .pcapng sans contexte perd toute valeur probante six mois plus tard.

### Étape 13 : respecter le cadre légal en France

Trois règles simples évitent les ennuis. D’abord, ne capturez jamais de trafic sur un réseau qui ne vous appartient pas sans autorisation écrite, qu’il s’agisse d’un client, d’un employeur ou d’un réseau public. Ensuite, si la capture peut contenir des données personnelles (identifiants, contenus de messages, cookies de session), traitez le fichier .pcapng comme une donnée sensible au sens du RGPD : chiffrement au repos, durée de conservation limitée, accès restreint. Enfin, sur un poste professionnel, informez généralement les utilisateurs concernés ou appuyez-vous sur la politique de sécurité de l’entreprise, en lien avec le délégué à la protection des données. Ces principes rejoignent ceux déjà détaillés dans les guides consacrés au scan réseau avec Nmap ou à la mise en conformité NIS2.

## Projet complet : détecter un brute-force SSH de bout en bout

