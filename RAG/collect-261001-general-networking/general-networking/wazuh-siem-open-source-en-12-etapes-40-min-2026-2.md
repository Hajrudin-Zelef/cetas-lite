---
id: collect-261001-general-networking/general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026-2
title: "Redémarre l'agent et relance les scans SCA"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [48, 150]
sha256: 6fec778f0e2957cf616c9bea035b8efe56e5c467306aa61411f2b7923d5a74d8
---

# Redémarre l'agent et relance les scans SCA

Le principal écueil d’une installation de Wazuh est le sous-dimensionnement. L’indexeur, basé sur OpenSearch (donc sur la JVM), a besoin de mémoire : en dessous de **8 Gio de RAM**, l’assistant d’installation refuse de continuer. Voici les besoins matériels officiels pour un déploiement tout-en-un, en fonction du nombre d’agents et d’une rétention de 90 jours.

| Composant | Agents surveillés | CPU | RAM | Stockage (90 j) | 
|---|---|---|---|---|
| Serveur tout-en-un | 1 à 25 | 4 vCPU | 8 Gio | 50 Go | 
| Serveur tout-en-un | 25 à 50 | 8 vCPU | 8 Gio | 100 Go | 
| Serveur tout-en-un | 50 à 100 | 8 vCPU | 8 Gio | 200 Go | 
| Agent Wazuh | – | négligeable | ~35 Mo | quelques Mo | 

Côté système d’exploitation, les composants centraux exigent un Linux 64 bits (x86_64 ou ARM64). Sont officiellement pris en charge : **Ubuntu 16.04 à 24.04**, **Red Hat Enterprise Linux 7 à 10**, **CentOS Stream 10**, **Amazon Linux 2 et 2023**. Pour ce tutoriel, nous utilisons Ubuntu 24.04 LTS. Les agents Wazuh, eux, tournent sur Linux, Windows, macOS, Solaris, AIX et HP-UX : la couverture est quasi universelle.

### Ports réseau à ouvrir

Avant de commencer, planifiez votre pare-feu. Ces ports doivent être joignables entre agents, serveur et administrateurs (jamais exposés à Internet sans filtrage).

| Port | Protocole | Usage | 
|---|---|---|
| 1514 | TCP | Communication agent → manager (flux chiffré des événements) | 
| 1515 | TCP | Enregistrement des agents (service `authd` ) | 
| 1516 | TCP | Communication interne du cluster Wazuh | 
| 514 | UDP | Réception Syslog (optionnel, équipements réseau) | 
| 55000 | TCP | API RESTful du manager Wazuh | 
| 9200 | TCP | API de l’indexeur Wazuh (OpenSearch) | 
| 443 | TCP | Tableau de bord Wazuh (HTTPS) | 

Vous aurez enfin besoin d’un accès `root` ou `sudo` sur les deux machines, d’une connexion Internet pour télécharger les paquets, et d’un navigateur récent. Vérifiez aussi que l’heure système est synchronisée (NTP) : un décalage d’horloge fausse l’horodatage des alertes et désynchronise les agents.

## Étape 1 – Préparer le serveur hôte

Connectez-vous à la machine qui hébergera Wazuh (notre `10.0.0.10`). On commence par mettre le système à jour, fixer le nom d’hôte, régler le fuseau horaire sur celui de la France et installer les quelques outils nécessaires. Un fuseau et une horloge corrects sont essentiels : c’est le socle d’une corrélation fiable des événements.

```
sudo hostnamectl set-hostname wazuh-server
sudo timedatectl set-timezone Europe/Paris
sudo apt update && sudo apt upgrade -y
sudo apt install -y curl tar apt-transport-https gnupg
sudo systemctl enable --now systemd-timesyncd
timedatectl status | grep -E "Time zone|synchronized"
```
La dernière commande doit confirmer que l’horloge est synchronisée. Exemple de sortie attendue :

```
System clock synchronized: yes
                Time zone: Europe/Paris (CEST, +0200)
```
Assurez-vous enfin de disposer d’au moins 8 Gio de RAM libres et de 50 Go d’espace disque sur la partition qui accueillera `/var/lib/wazuh-indexer`. Sur une machine virtuelle sous-dimensionnée, l’installation échouera à l’étape suivante.

## Étape 2 – Installer les composants centraux de Wazuh

Wazuh fournit un assistant d’installation « tout-en-un » qui déploie le manager, l’indexeur et le tableau de bord, génère les certificats SSL et crée les mots de passe. C’est la méthode recommandée pour découvrir Wazuh ou couvrir jusqu’à une centaine d’agents. Une seule commande suffit :

`curl -sO https://packages.wazuh.com/4.14/wazuh-install.sh && sudo bash ./wazuh-install.sh -a`
L’option `-a` (pour *all-in-one*) enchaîne l’installation des trois composants. Comptez 8 à 15 minutes selon votre bande passante et votre CPU. L’assistant télécharge les paquets, configure OpenSearch, applique les certificats et démarre les services. À la fin, il affiche une ligne décisive :

```
INFO: --- Summary ---
INFO: You can access the web interface https://10.0.0.10:443
    User: admin
    Password: <MOT_DE_PASSE_GENERE>
INFO: Installation finished.
```
Notez soigneusement ce mot de passe. L’assistant crée également une archive `wazuh-install-files.tar` contenant tous les identifiants et certificats. Conservez-la précieusement et hors ligne. Vérifions maintenant que les trois services tournent :

`sudo systemctl status wazuh-manager wazuh-indexer wazuh-dashboard --no-pager`
Les trois doivent afficher `active (running)`. Si l’indexeur reste en `activating` ou tombe en échec, c’est presque toujours un manque de RAM : consultez la section dépannage. La documentation de référence pour cette étape est le guide de démarrage rapide officiel.

## Étape 3 – Accéder au tableau de bord Wazuh

Ouvrez votre navigateur sur `https://10.0.0.10:443` (remplacez par l’IP de votre serveur). Le certificat étant auto-signé par défaut, votre navigateur affichera un avertissement de sécurité : acceptez l’exception pour continuer. Connectez-vous avec l’utilisateur `admin` et le mot de passe relevé à l’étape 2.

Si vous avez égaré le mot de passe, extrayez-le de l’archive générée par l’installateur :

`sudo tar -O -xvf wazuh-install-files.tar wazuh-install-files/wazuh-passwords.txt`
Cette commande affiche l’ensemble des identifiants (utilisateurs `admin`, `kibanaserver`, `wazuh-wui`, etc.). Une fois connecté, vous arrivez sur la vue d’ensemble de Wazuh : modules de sécurité, agents, alertes récentes et état de conformité. Le tableau de bord est encore vide car aucun agent n’est enrôlé – c’est l’objet de l’étape suivante.

Bonne pratique : changez immédiatement le mot de passe `admin` par défaut. Le script dédié se trouve dans le répertoire de l’indexeur (`/usr/share/wazuh-indexer/plugins/opensearch-security/tools/wazuh-passwords-tool.sh`). Nous détaillerons le durcissement complet à l’étape 11.

## Étape 4 – Déployer votre premier agent Wazuh

Un SIEM sans agent ne voit rien. Connectez-vous à la machine à surveiller (`web01`, `10.0.0.21`) et installez l’agent Wazuh. La méthode la plus propre passe par le dépôt APT officiel, avec une variable d’environnement qui pointe l’agent vers votre manager et lui donne un nom lisible.

```
curl -s https://packages.wazuh.com/key/GPG-KEY-WAZUH | sudo gpg --no-default-keyring \
  --keyring gnupg-ring:/usr/share/keyrings/wazuh.gpg --import
sudo chmod 644 /usr/share/keyrings/wazuh.gpg
echo "deb [signed-by=/usr/share/keyrings/wazuh.gpg] https://packages.wazuh.com/4.x/apt/ stable main" \
  | sudo tee /etc/apt/sources.list.d/wazuh.list
sudo apt update
sudo WAZUH_MANAGER="10.0.0.10" WAZUH_AGENT_NAME="web01" apt install -y wazuh-agent
sudo systemctl daemon-reload
sudo systemctl enable --now wazuh-agent
```
L’agent s’enregistre automatiquement auprès du manager via le port 1515, puis bascule sur le port 1514 pour transmettre ses événements. Retournez sur le serveur Wazuh et listez les agents connectés :

`sudo /var/ossec/bin/agent_control -l````
Wazuh agent_control. List of available agents:
   ID: 000, Name: wazuh-server (server), IP: 127.0.0.1, Active/Local
   ID: 001, Name: web01, IP: 10.0.0.21, Active
```
Le statut `Active` confirme que la liaison fonctionne. Dans le tableau de bord, le nouvel agent apparaît sous *Agents management*. Pour un poste Windows, la démarche est identique via le programme d’installation MSI, en passant les mêmes paramètres `WAZUH_MANAGER` et `WAZUH_AGENT_NAME`. Répétez l’opération pour chaque machine de votre parc.

## Étape 5 – Surveiller l’intégrité des fichiers (FIM)

