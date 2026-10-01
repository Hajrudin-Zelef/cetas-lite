---
id: collect-261001-general-networking/general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026-4
title: "Mettre à jour la liste des paquets et le système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "agent", "agents", "datacenter", "distribution", "mai"]
source: docs/RAG/collect-261001-general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026.md
source_anchor: ""
source_lines: [272, 363]
sha256: 175d9162a71dcab26c4b65f4a98bc3535008a80ac7cb09ad87c17f4eca0a48f7
---

# Mettre à jour la liste des paquets et le système

```
# Rejouer le journal d'authentification a travers le moteur
sudo cscli explain --file /var/log/auth.log --type syslog
# Tester une ligne precise de log Nginx
sudo cscli explain --log '...ligne de log nginx...' --type nginx
```
Pour un test « grandeur nature » du blocage réseau, bannissez votre propre IP de test (depuis une connexion secondaire, jamais votre seule session SSH !) et vérifiez que le pare-feu la rejette. Côté web, une requête simulant une traversée de répertoire vers une URL inexistante doit générer une alerte visible dans `cscli alerts list`. Si l’alerte apparaît puis qu’une décision en découle et que l’IP est bloquée, votre projet CrowdSec est pleinement fonctionnel.

```
# Observer en direct les nouvelles decisions
watch -n 2 "sudo cscli decisions list"
# Confirmer la regle appliquee par le bouncer pare-feu (iptables)
sudo iptables -L -n | grep -i crowdsec
```
## Anatomie d’une alerte : lire une détection CrowdSec

Savoir interpréter une alerte est essentiel pour distinguer une vraie attaque d’un faux positif. La commande `cscli alerts inspect` détaille un événement : l’IP source, son pays et son AS (l’opérateur réseau), le scénario déclenché, le nombre d’événements corrélés et la décision appliquée. Voici un exemple de sortie typique pour une attaque par force brute SSH.

```
# Inspecter une alerte en detail (ID visible via cscli alerts list)
sudo cscli alerts inspect 42 -d
# Sortie (extrait) :
#  - ID            : 42
#  - Scenario      : crowdsecurity/ssh-bf
#  - Source IP     : 203.0.113.42
#  - Country       : RU
#  - AS            : 12345 EXAMPLE-AS
#  - Events count  : 6
#  - Decision      : ban (4h)
```
Chaque champ raconte une histoire. « Events count : 6 » signifie que le scénario a corrélé six échecs d’authentification avant de déclencher, ce qui écarte un simple oubli de mot de passe d’un utilisateur légitime. Le couple pays/AS aide à juger la légitimité : un AS d’hébergeur (datacenter) qui tente de se connecter en SSH est presque toujours malveillant. Aller plus loin est immédiat : `cscli alerts list --scope ip --value 203.0.113.42` agrège tout l’historique d’une adresse, et la Console affiche sa réputation communautaire – signalée par combien d’autres moteurs, depuis quand et pour quels scénarios. C’est cette mise en contexte qui transforme une simple ligne de log en renseignement réellement actionnable, et c’est précisément ce que `fail2ban` ne sait pas faire.

## Étape 12 – Docker et déploiement multi-serveur

CrowdSec brille dans les environnements conteneurisés et distribués. Pour un déploiement Docker, l’image officielle `crowdsecurity/crowdsec` permet de tout configurer via variables d’environnement – le projet publie des versions à un rythme soutenu, comme en témoigne la cadence 2025-2026 des paquets Windows référencés sur SourceForge : la 1.7.0 (75,7 Mo, 1er septembre 2025, avec une refonte de l’auto-détection et des métriques), puis la 1.7.1 (79,6 Mo, 21 octobre 2025), la 1.7.8 (90,6 Mo, 11 mai 2026), puis la **1.8.0** en août 2026, qui a justement introduit une source de logs native pour **Kubernetes** – un atout direct pour ce type de déploiement conteneurisé – avant que la **1.8.1** ne prenne le relais comme dernière version publiée sur GitHub en septembre 2026. Montez vos logs en lecture seule, persistez la base de données et la configuration, et préférez épingler un tag de version précis (idéalement 1.8.x) plutôt que `latest` pour maîtriser vos mises à jour.

```
docker run -d \
  --name crowdsec \
  -e COLLECTIONS="crowdsecurity/sshd crowdsecurity/nginx" \
  -v /var/log:/var/log:ro \
  -v crowdsec-config:/etc/crowdsec \
  -v crowdsec-data:/var/lib/crowdsec/data \
  -p 8080:8080 \
  crowdsecurity/crowdsec:v1.7.8
```
Pour une architecture **multi-serveur**, le principe est d’avoir une LAPI centrale (sur un serveur dédié ou le plus stable) et plusieurs moteurs « agents » distribués qui lui remontent leurs alertes. Chaque agent distant s’enregistre auprès de la LAPI centrale, et les bouncers de chaque machine interrogent cette même API. On obtient une vision consolidée et une remédiation cohérente sur tout le parc.

```
# Sur la LAPI centrale : creer un identifiant pour un agent distant
sudo cscli machines add agent-web-01 --auto
# Sur l'agent distant : pointer vers la LAPI centrale
# (editer les identifiants dans /etc/crowdsec/local_api_credentials.yaml)
sudo cscli lapi register --machine agent-web-01 -u http://LAPI_CENTRALE:8080
# Valider l'enregistrement cote central
sudo cscli machines list
```
Cette topologie est idéale pour les hébergeurs, les agences web et les administrateurs gérant plusieurs VPS. Elle s’articule très bien avec un réseau privé maillé : voyez notre tutoriel Tailscale VPN Mesh pour relier vos moteurs et leur LAPI centrale sur un réseau chiffré privé, sans exposer le port 8080 sur Internet.

## Aide-mémoire : les commandes cscli essentielles

Gardez ce tableau sous la main : il couvre 95 % de l’exploitation quotidienne de CrowdSec en production.

| Commande | Rôle | 
|---|---|
| `cscli metrics` | Statistiques du moteur (acquisition, parseurs, scénarios) | 
| `cscli collections list` | Lister les collections installées | 
| `cscli hub update && cscli hub upgrade` | Mettre à jour le Hub et ses éléments | 
| `cscli decisions list` | Voir les IP actuellement bannies | 
| `cscli decisions add --ip X --duration 24h` | Bannir manuellement une IP | 
| `cscli decisions delete --ip X` | Lever un bannissement | 
| `cscli alerts list` | Historique des alertes déclenchées | 
| `cscli bouncers list` | Lister les bouncers connectés | 
| `cscli console enroll CLE` | Lier le moteur à la Console | 
| `cscli explain --file ... --type ...` | Déboguer parseurs et scénarios | 
| `cscli capi status` | Tester la connexion à la blocklist communautaire | 

## Pièges courants à éviter

Au fil des déploiements, les mêmes erreurs reviennent. Les anticiper vous fera gagner des heures de débogage.

- **Oublier le bouncer.** Installer le moteur seul ne protège rien : il détecte mais ne bloque pas. Sans au moins un bouncer (pare-feu et/ou Nginx), CrowdSec n’est qu’un système d’alerte. C’est de loin l’erreur n°1.
- **Se bannir soi-même.** Lors des tests SSH, ne bannissez jamais l’IP de votre unique session ouverte. Mettez toujours votre IP d’administration en liste blanche*avant* toute simulation d’attaque.
- **Confondre iptables et nftables.** Installer le bouncer iptables sur un système qui utilise nftables (ou l’inverse) aboutit à des règles non appliquées. Vérifiez votre pare-feu avant de choisir le paquet.
- **Mauvais chemin de logs.** Si`acquis.yaml` pointe vers un fichier inexistant ou si la distribution journalise via systemd plutôt qu’un fichier plat, le moteur lit zéro ligne. Contrôlez toujours la section Acquisition de`cscli metrics` .
- **Ne pas recharger après modification.** Toute modification d’acquisition, de liste blanche ou de profil nécessite un`systemctl reload crowdsec` . Sans cela, vos changements restent lettre morte.
- **Exposer la LAPI sur Internet.** Le port 8080 doit rester en localhost ou sur un réseau privé. Ne le publiez jamais directement ; passez par un VPN ou un tunnel chiffré pour l’architecture multi-serveur.

## Dépannage : 8 problèmes fréquents et solutions

Voici les incidents les plus courants rencontrés lors d’une installation et configuration de CrowdSec, avec leur résolution.

