---
id: collect-261001-general-networking/general-networking/serveur-wireguard-auto-heberge-12-etapes-2026-2
title: "Les blocs [Peer] des clients seront ajoutés ici à l'étape 7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/serveur-wireguard-auto-heberge-12-etapes-2026.md
source_anchor: ""
source_lines: [51, 160]
sha256: 92b1c8a2b375392303ba646cc94aa8b6b67920a4a9f7665763582c05477022ff
---

# Les blocs [Peer] des clients seront ajoutés ici à l'étape 7

```
cd /etc/wireguard
umask 077
wg genkey | tee server_private.key | wg pubkey > server_public.key
cat server_private.key
cat server_public.key
```
La commande `umask 077` garantit que les fichiers créés ensuite ne sont lisibles que par root, ce qui est indispensable : quiconque récupère votre clé privée peut usurper votre serveur. Gardez les deux fichiers de côté, vous allez en avoir besoin dans le fichier de configuration à l’étape suivante.

## Étape 4 : créer le fichier de configuration du serveur

Créez le fichier `/etc/wireguard/wg0.conf`. C’est ce fichier qui définit l’adresse IP virtuelle du serveur sur le tunnel, le port d’écoute et les règles de routage.

```
[Interface]
PrivateKey = <contenu_de_server_private.key>
Address = 10.10.0.1/24
ListenPort = 51820
PostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
# Les blocs [Peer] des clients seront ajoutés ici à l'étape 7
```
Le port 51820/UDP est le port par défaut de WireGuard, mais rien ne vous empêche de le changer pour un port moins prévisible. Remplacez `eth0` par le nom réel de votre interface réseau sortante (vérifiable avec `ip a`). Les lignes `PostUp` et `PostDown` activent le masquerading NAT pour que le trafic des clients puisse sortir vers internet via le serveur.

## Étape 5 : activer le routage IP (IP forwarding)

Par défaut, un serveur Linux ne route pas le trafic entre interfaces. Il faut l’activer explicitement pour que les paquets des clients VPN puissent transiter vers internet.

```
sudo sed -i 's/#net.ipv4.ip_forward=1/net.ipv4.ip_forward=1/' /etc/sysctl.conf
sudo sysctl -p
```
Si la ligne `net.ipv4.ip_forward=1` n’existe pas dans `/etc/sysctl.conf`, ajoutez-la manuellement à la fin du fichier. La commande `sysctl -p` recharge la configuration sans redémarrer le serveur. Une erreur fréquente à cette étape : oublier cette activation, ce qui fait que le tunnel se connecte mais qu’aucun trafic internet ne passe côté client.

## Étape 6 : démarrer le service WireGuard

WireGuard s’intègre à systemd via `wg-quick`. Démarrez l’interface et activez-la au démarrage du serveur.

```
sudo systemctl start wg-quick@wg0
sudo systemctl enable wg-quick@wg0
sudo wg show
```
La commande `wg show` affiche l’état de l’interface : adresse IP, port d’écoute et clé publique. Si rien ne s’affiche ou si le service refuse de démarrer, consultez immédiatement les journaux avec `journalctl -u wg-quick@wg0 -n 50` pour identifier l’erreur de syntaxe dans le fichier de configuration.

## Étape 7 : générer les clés client et ajouter un pair (peer)

Chaque appareil qui se connecte au VPN (ordinateur portable, téléphone, tablette) a besoin de sa propre paire de clés. Sur le serveur, ou directement sur l’appareil client, générez une nouvelle paire.

`wg genkey | tee client1_private.key | wg pubkey > client1_public.key`
Ajoutez ensuite un bloc `[Peer]` dans le fichier `/etc/wireguard/wg0.conf` du serveur, un bloc par appareil autorisé.

```
[Peer]
PublicKey = <contenu_de_client1_public.key>
AllowedIPs = 10.10.0.2/32
```
Chaque appareil reçoit une adresse IP unique dans la plage virtuelle (10.10.0.2, 10.10.0.3, et ainsi de suite). Rechargez la configuration sans couper les connexions existantes avec `wg syncconf wg0 <(wg-quick strip wg0)`, ou plus simplement redémarrez l’interface avec `systemctl restart wg-quick@wg0`.

## Étape 8 : configurer le fichier client

Sur l’appareil client, créez un fichier `client1.conf` avec les informations suivantes. Ce fichier peut aussi être converti en QR code pour une importation rapide sur mobile via l’application officielle WireGuard.

```
[Interface]
PrivateKey = <contenu_de_client1_private.key>
Address = 10.10.0.2/24
DNS = 1.1.1.1
[Peer]
PublicKey = <contenu_de_server_public.key>
Endpoint = votre-ip-publique-ou-ddns:51820
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
```
`AllowedIPs = 0.0.0.0/0` route tout le trafic de l’appareil à travers le VPN (mode tunnel complet). Si vous voulez uniquement accéder à votre réseau domestique sans rediriger la navigation internet générale, remplacez cette valeur par la plage de votre réseau local, par exemple `192.168.1.0/24`. Le paramètre `PersistentKeepalive = 25` envoie un paquet toutes les 25 secondes pour maintenir le tunnel ouvert à travers un NAT, ce qui est indispensable sur mobile en 4G/5G.

## Étape 9 : rediriger le port UDP sur une Freebox

C’est l’étape où la majorité des lecteurs français bloquent. Si le serveur WireGuard tourne derrière une Freebox, connectez-vous à l’interface de gestion sur `mafreebox.freebox.fr`, puis allez dans **Paramètres de la Freebox > Mode avancé > Redirections de ports**. Créez une nouvelle redirection avec le protocole UDP, le port externe 51820, et l’IP locale de votre serveur (visible avec `ip a` sur le serveur lui-même). Pensez aussi à attribuer un bail DHCP fixe à votre serveur dans les paramètres réseau de la Freebox, sinon son adresse locale peut changer après un redémarrage et casser la redirection.

## Étape 10 : rediriger le port UDP sur une Livebox

Sur une Livebox 5 ou 6, la procédure est similaire mais l’interface diffère. Connectez-vous sur `http://livebox.home` ou `192.168.1.1`, ouvrez le menu **Réseau domestique > NAT/PAT**, puis ajoutez une règle personnalisée pour le port UDP 51820 vers l’adresse IP locale du serveur. Comme pour la Freebox, réservez une IP fixe pour le serveur afin d’éviter que la redirection ne se rompe après un changement d’adresse DHCP. Si votre FAI applique du CGNAT (fréquent sur certaines offres fibre ou 4G/5G box), la redirection de port ne fonctionnera pas du tout : dans ce cas, un VPS avec IP publique fixe reste la seule solution fiable.

## Étape 11 : tester la connexion depuis un réseau externe

Ne testez jamais uniquement depuis le réseau local : une connexion qui fonctionne en Wi-Fi domestique peut échouer totalement depuis l’extérieur si la redirection de port est mal configurée. Basculez votre téléphone en 4G/5G (Wi-Fi désactivé), importez le fichier `client1.conf` dans l’application WireGuard officielle, et activez le tunnel.

```
# Depuis le client, une fois connecté
curl ifconfig.me
ping 10.10.0.1
```
La commande `curl ifconfig.me` doit retourner l’IP publique de votre serveur, pas celle de votre opérateur mobile : c’est la preuve que tout votre trafic sort bien par le tunnel. Le `ping` vers `10.10.0.1` confirme que le tunnel lui-même répond. Sur le serveur, la commande `wg show` doit afficher un `latest handshake` récent (moins d’une minute) pour ce pair.

## Étape 12 : durcir la sécurité du serveur

Une fois le tunnel fonctionnel, quelques réglages supplémentaires réduisent la surface d’attaque. Limitez d’abord l’accès SSH à une authentification par clé uniquement (désactivez `PasswordAuthentication` dans `/etc/ssh/sshd_config`), puis restreignez le pare-feu pour n’autoriser que le port SSH et le port WireGuard en entrée.

```
sudo ufw default deny incoming
sudo ufw allow 22/tcp
sudo ufw allow 51820/udp
sudo ufw enable
sudo ufw status verbose
```
Pensez aussi à faire tourner vos clés régulièrement (une fois par an suffit pour un usage personnel), à retirer immédiatement un bloc `[Peer]` dès qu’un appareil est perdu ou volé, et à surveiller les tentatives de connexion SSH avec un outil comme Fail2ban pour bloquer automatiquement le brute-force. Consultez notre tutoriel Fail2ban si ce n’est pas déjà en place sur votre serveur.

## Installer l’application client sur chaque plateforme

