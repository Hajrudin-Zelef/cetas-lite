---
id: collect-261001-general-networking/general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026-2
title: "repérez le nom de votre interface, par exemple eth0 ou ens18"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026.md
source_anchor: ""
source_lines: [48, 167]
sha256: 15e30bdc0f829658ce9fa21344473461a18750d1c205f2e0e3aa28d1b3a1517a
---

# repérez le nom de votre interface, par exemple eth0 ou ens18

```
echo "deb http://deb.debian.org/debian bookworm-backports main" | sudo tee /etc/apt/sources.list.d/backports.list
sudo apt update
sudo apt install -y suricata -t bookworm-backports
suricata --build-info | grep "This is Suricata"
```
Sur Ubuntu, un PPA maintenu par l’OISF distribue également les paquets 8.0.x à jour :

```
sudo add-apt-repository ppa:oisf/suricata-stable
sudo apt update
sudo apt install -y suricata
suricata -V
```
La commande `suricata -V` doit retourner au minimum `8.0.6` ; la dernière version stable publiée par l’OISF est la 8.0.7, sortie en septembre 2026, et il est recommandé d’y migrer dès qu’elle est disponible dans les dépôts backports ou le PPA. Si une version antérieure s’affiche, vérifiez que le dépôt backports ou le PPA est bien prioritaire dans votre configuration APT avant de continuer.

## Étape 3 : Comprendre l’arborescence de configuration

L’installation crée plusieurs répertoires clés qu’il faut connaître avant de toucher à quoi que ce soit :

- `/etc/suricata/suricata.yaml` : fichier de configuration principal
- `/etc/suricata/rules/` : dossier des règles de détection
- `/var/log/suricata/` : logs, alertes et fichiers`eve.json`
- `/var/lib/suricata/` : données runtime, dont la base de données GeoIP si activée

Le fichier `eve.json` est le cœur du système : c’est un flux structuré au format JSON qui contient toutes les alertes, les métadonnées de flux, les logs DNS, TLS et HTTP. C’est ce fichier que vous brancherez plus tard sur un SIEM comme Wazuh ou Graylog pour la corrélation et la visualisation des événements.

## Étape 4 : Configurer l’interface de capture dans suricata.yaml

Ouvrez le fichier de configuration principal et modifiez la section `af-packet` pour pointer vers votre interface réseau réelle.

`sudo nano /etc/suricata/suricata.yaml`
Repérez la section `af-packet` et adaptez-la :

```
af-packet:
  - interface: eth0
    threads: auto
    cluster-id: 99
    cluster-type: cluster_flow
    defrag: yes
    use-mmap: yes
    tpacket-v3: yes
```
L’option `cluster-type: cluster_flow` permet à Suricata de répartir la charge entre plusieurs threads tout en gardant les paquets d’un même flux TCP sur le même thread, ce qui est essentiel pour une inspection cohérente. Modifiez également la variable `HOME_NET` en haut du fichier pour qu’elle corresponde à votre plage IP interne, sans quoi de nombreuses règles ne se déclencheront pas correctement.

```
vars:
  address-groups:
    HOME_NET: "[192.168.1.0/24,10.0.0.0/8]"
    EXTERNAL_NET: "!$HOME_NET"
```
## Étape 5 : Choisir entre mode IDS et mode IPS

Suricata peut fonctionner en mode passif (IDS, détection uniquement avec alertes) ou en mode actif en ligne (IPS, capable de bloquer le trafic malveillant). Pour un premier déploiement, commencez toujours en mode IDS afin d’observer le comportement des règles sans risquer de bloquer du trafic légitime.

| Mode | Fonctionnement | Cas d’usage recommandé | 
|---|---|---|
| IDS (af-packet) | Capture passive via mirroring, alertes uniquement | Phase de test, tuning des règles, environnements sensibles | 
| IPS (NFQUEUE) | Inline avec iptables/nftables, blocage actif | Production après validation, passerelle réseau | 
| IPS (af-packet en ligne) | Bridge transparent entre deux interfaces | Déploiement en coupure sur un lien dédié | 

Pour activer le mode IPS via NFQUEUE sur Linux, redirigez le trafic vers Suricata avec nftables avant de lancer le moteur en mode `-q` :

```
sudo nft add rule inet filter forward queue num 0
sudo suricata -c /etc/suricata/suricata.yaml -q 0
```
## Étape 6 : Mettre à jour et gérer les règles de détection avec suricata-update

Suricata ne détecte rien sans règles. L’outil officiel `suricata-update`, installé automatiquement avec le paquet 8.0.6, gère le téléchargement et la mise à jour des rulesets depuis plusieurs sources communautaires.

```
sudo suricata-update
sudo suricata-update list-sources
```
Par défaut, la source Emerging Threats Open est activée. Vous pouvez en ajouter d’autres, comme les règles ET Pro (payantes) ou les flux communautaires spécialisés dans les malwares connus :

```
sudo suricata-update enable-source et/open
sudo suricata-update enable-source oisf/trafficid
sudo suricata-update
```
Planifiez une tâche cron pour que ces règles se mettent à jour automatiquement chaque jour, ce qui est indispensable pour rester protégé contre les menaces les plus récentes :

`echo "0 4 * * * root suricata-update && systemctl restart suricata" | sudo tee /etc/cron.d/suricata-update`
## Comprendre la syntaxe d’une règle Suricata

Avant d’aller plus loin, il est utile de comprendre comment une règle Suricata est structurée, ne serait-ce que pour interpréter correctement les alertes que vous allez recevoir. Une règle se compose de trois blocs : une action, un en-tête réseau et des options entre parenthèses.

`alert tcp $EXTERNAL_NET any -> $HOME_NET 22 (msg:"ET SCAN Potential SSH Scan"; flags:S; threshold:type both, track by_src, count 5, seconds 60; classtype:attempted-recon; sid:2001219; rev:20;)`
L’action `alert` indique que Suricata doit générer une alerte sans bloquer le paquet (contrairement à `drop`, réservé au mode IPS). Vient ensuite le protocole (`tcp`), puis l’adresse et le port source, une direction (`->` pour un sens unique, `<>` pour bidirectionnel), et enfin l’adresse et le port de destination. Le bloc d’options contient le message affiché dans les logs (`msg`), les conditions de correspondance, la classification de la menace, et surtout le `sid`, un identifiant unique qui permet de retrouver, désactiver ou surveiller cette règle précise dans vos outils de gestion.

Le mot-clé `threshold` mérite une attention particulière : il évite de recevoir une alerte à chaque paquet suspect individuel et regroupe les événements similaires sur une fenêtre de temps donnée, ici cinq tentatives en 60 secondes avant de déclencher l’alerte. C’est ce type de réglage qui distingue une configuration Suricata bien réglée d’une installation par défaut qui noie l’équipe sécurité sous des milliers d’alertes redondantes.

## Écrire une règle de détection personnalisée

Les rulesets communautaires comme Emerging Threats Open couvrent l’essentiel, mais chaque environnement a ses propres besoins. Créez un fichier dédié à vos règles maison plutôt que de modifier les fichiers téléchargés automatiquement, qui seraient écrasés à la prochaine mise à jour.

`sudo nano /etc/suricata/rules/local.rules`
Voici un exemple de règle personnalisée qui détecte une tentative de connexion à un panneau d’administration sensible depuis l’extérieur du réseau, un cas fréquent pour protéger un service interne exposé par erreur :

`alert http $EXTERNAL_NET any -> $HOME_NET any (msg:"LOCAL Acces suspect au panneau admin depuis Internet"; flow:established,to_server; http.uri; content:"/wp-admin"; classtype:web-application-attack; sid:9000001; rev:1;)`
Ajoutez ensuite ce fichier à la liste des règles chargées dans `suricata.yaml`, sous la section `rule-files`, puis rechargez la configuration :

```
rule-files:
  - suricata.rules
  - local.rules
```
```
sudo suricata -T -c /etc/suricata/suricata.yaml -v
sudo systemctl reload suricata
```
Numérotez toujours vos SID personnalisés au-delà de 9 000 000, une plage réservée par convention aux règles locales, afin d’éviter tout conflit avec les identifiants utilisés par les rulesets publics.

## Étape 7 : Valider la configuration avant le premier lancement

Ne lancez jamais Suricata en production sans avoir vérifié que la configuration et les règles sont syntaxiquement valides. Le mode `-T` (test) permet de le faire sans démarrer réellement le service.

