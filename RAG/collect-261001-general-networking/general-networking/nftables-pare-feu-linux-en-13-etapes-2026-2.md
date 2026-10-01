---
id: collect-261001-general-networking/general-networking/nftables-pare-feu-linux-en-13-etapes-2026-2
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/nftables-pare-feu-linux-en-13-etapes-2026.md
source_anchor: ""
source_lines: [77, 183]
sha256: 0cb76cb5c9907bdaff0cced8f8ed3048f19a350e5e93452b8d9e0d2b47e8c2f9
---

# Debian / Ubuntu

Contrairement à iptables qui imposait des tables figées (filter, nat, mangle), nftables laisse l’administrateur nommer ses propres tables et organiser sa logique comme il l’entend. Cette flexibilité est puissante mais demande de la rigueur : sans convention de nommage claire, une configuration complexe devient vite illisible pour la personne qui la reprendra six mois plus tard.

- **Table** : conteneur logique, par exemple`inet filtre` pour IPv4 et IPv6 combinés
- **Chaîne de base** : rattachée à un hook noyau (input, output, forward, prerouting, postrouting)
- **Chaîne régulière** : appelée depuis une autre chaîne via`jump` ou`goto` , utile pour modulariser
- **Priorité** : détermine l’ordre d’exécution quand plusieurs chaînes de base partagent le même hook
- **Politique** : verdict par défaut appliqué si aucune règle ne correspond (accept ou drop)

## Étape 4 : Créer une politique de base par défaut

La règle d’or en sécurité réseau reste la même depuis des décennies : tout refuser par défaut, puis autoriser explicitement ce qui est nécessaire. Créez une table et une chaîne d’entrée avec une politique drop.

```
sudo nft add table inet filtre
sudo nft add chain inet filtre entree { type filter hook input priority 0 \; policy drop \; }
sudo nft add chain inet filtre sortie { type filter hook output priority 0 \; policy accept \; }
sudo nft add chain inet filtre transit { type filter hook forward priority 0 \; policy drop \; }
```
À ce stade, si vous appliquez cette politique sans autoriser au préalable les connexions établies et le trafic de loopback, vous risquez de couper votre propre session SSH immédiatement. Ajoutez donc systématiquement ces deux règles en premier, avant toute politique restrictive.

```
sudo nft add rule inet filtre entree iif lo accept
sudo nft add rule inet filtre entree ct state established,related accept
sudo nft add rule inet filtre entree ct state invalid drop
```
## Étape 5 : Autoriser SSH, HTTP et HTTPS en entrée

Une fois la base posée, ouvrez uniquement les ports que votre serveur doit réellement exposer. Sur un serveur web classique, cela se limite souvent à trois ports.

```
sudo nft add rule inet filtre entree tcp dport 22 accept
sudo nft add rule inet filtre entree tcp dport { 80, 443 } accept
sudo nft add rule inet filtre entree ip protocol icmp accept
sudo nft add rule inet filtre entree ip6 nexthdr icmpv6 accept
```
Restreindre le port SSH à une plage d’adresses IP connues réduit encore davantage la surface d’attaque. Si votre équipe travaille depuis une IP fixe ou un VPN d’entreprise, verrouillez l’accès dès maintenant plutôt que d’ouvrir SSH au monde entier.

`sudo nft add rule inet filtre entree ip saddr 203.0.113.0/24 tcp dport 22 accept`
## Étape 6 : Limiter le débit et contrer le brute-force SSH

SSH exposé sur Internet reçoit des tentatives de connexion automatisées en continu, avec des scanners qui testent des milliers de combinaisons de mots de passe par heure. nftables propose un mécanisme de limitation de débit natif, sans dépendre d’un outil externe, même si le combiner avec Fail2ban ou CrowdSec reste recommandé pour une défense en profondeur.

```
sudo nft add rule inet filtre entree tcp dport 22 ct state new limit rate 4/minute accept
sudo nft add rule inet filtre entree tcp dport 22 ct state new log prefix "SSH-DROP: " drop
```
Cette paire de règles autorise au maximum quatre nouvelles connexions SSH par minute et journalise, avant de rejeter, toute tentative supplémentaire. Un utilisateur légitime qui tape mal son mot de passe deux ou trois fois ne sera jamais bloqué. Un script de brute-force qui tente des centaines de combinaisons par minute se heurte immédiatement au mur.

## Étape 7 : Configurer le NAT et le masquerading pour un routeur Linux

Si votre machine Linux fait office de routeur ou de passerelle pour un réseau local, nftables gère aussi la traduction d’adresses réseau (NAT). Cela concerne typiquement les box maison sous Linux, les hyperviseurs avec des VM en réseau privé, ou les serveurs qui partagent une connexion Internet.

```
sudo nft add table ip nat
sudo nft add chain ip nat postrouting { type nat hook postrouting priority 100 \; }
sudo nft add rule ip nat postrouting oif "eth0" masquerade
echo 'net.ipv4.ip_forward=1' | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```
Sans activer `ip_forward` au niveau du noyau, aucune règle de NAT ne fonctionnera, même parfaitement écrite. C’est l’une des causes d’échec les plus fréquentes rencontrées par les administrateurs qui configurent leur premier routeur Linux.

## Étape 8 : Journaliser les paquets rejetés pour l’analyse forensique

Un pare-feu qui bloque silencieusement ne raconte aucune histoire. Journaliser les paquets rejetés permet de repérer des tentatives de scan, des attaques ciblées, ou simplement une mauvaise configuration côté client qui envoie du trafic inattendu.

`sudo nft add rule inet filtre entree log prefix "NFT-DROP: " flags all counter drop`
Placez cette règle en toute dernière position de la chaîne, juste avant la fin implicite. Les logs atterrissent dans `journalctl -k` ou `/var/log/kern.log` selon votre configuration syslog. Sur un serveur à fort trafic, limitez le volume de logs avec `limit rate` pour éviter de saturer le disque.

`sudo journalctl -k -f | grep NFT-DROP`
Exemple de sortie typique après une tentative de scan de port :

```
kernel: NFT-DROP: IN=eth0 OUT= MAC=... SRC=198.51.100.42 DST=203.0.113.10
LEN=40 TOS=0x00 PREC=0x00 TTL=48 ID=54321 PROTO=TCP SPT=51422 DPT=3389
WINDOW=1024 RES=0x00 SYN URGP=0
```
Cette ligne montre une tentative de connexion sur le port 3389 (RDP) depuis une adresse externe, un port que ce serveur Linux n’expose jamais. C’est exactement le type de bruit de fond qu’on observe sur n’importe quelle IP publique.

## Étape 9 : Simplifier les règles avec les sets et les verdict maps

Quand le nombre de règles grandit, les gérer une par une devient vite ingérable. nftables introduit les sets (ensembles) et les verdict maps, deux structures qui remplacent des dizaines de règles individuelles par une seule ligne dynamique.

```
# Un set nommé de ports autorisés, modifiable sans réécrire la règle
sudo nft add set inet filtre ports_autorises { type inet_service \; }
sudo nft add element inet filtre ports_autorises { 22, 80, 443, 8443 }
sudo nft add rule inet filtre entree tcp dport @ports_autorises accept
# Un set d'IP bannies, alimenté automatiquement par un script ou Fail2ban
sudo nft add set inet filtre ip_bannies { type ipv4_addr \; flags timeout \; }
sudo nft add rule inet filtre entree ip saddr @ip_bannies drop
```
Le second exemple, avec le flag `timeout`, permet d’ajouter dynamiquement une IP bannie pour une durée limitée sans jamais toucher aux règles elles-mêmes. C’est exactement le mécanisme qu’utilisent les outils comme CrowdSec ou Fail2ban en coulisses pour interagir avec nftables.

## Étape 10 : Rendre la configuration persistante avec systemd

Toutes les commandes tapées jusqu’ici disparaissent au redémarrage si elles ne sont pas sauvegardées dans un fichier de configuration. nftables charge automatiquement `/etc/nftables.conf` au démarrage via le service systemd activé à l’étape 2.

```
# Exporter la configuration active vers le fichier persistant
sudo nft list ruleset > /etc/nftables.conf
# Vérifier la syntaxe avant de redémarrer le service
sudo nft -c -f /etc/nftables.conf
# Recharger sans redémarrer le serveur
sudo systemctl restart nftables
```
L’option `-c` (check) valide la syntaxe du fichier sans l’appliquer. Prenez l’habitude de toujours passer par cette vérification avant un redémarrage du service : une erreur de syntaxe dans `/etc/nftables.conf` peut empêcher le service de démarrer, laissant potentiellement le serveur sans aucun filtrage actif.

