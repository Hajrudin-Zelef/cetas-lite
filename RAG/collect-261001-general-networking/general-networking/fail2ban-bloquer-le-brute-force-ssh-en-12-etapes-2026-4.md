---
id: collect-261001-general-networking/general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026-4
title: "Vérifier la version installée"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026.md
source_anchor: ""
source_lines: [315, 419]
sha256: de0df0e4b78e7a88e1f394e57cc0d9739086af311e2bdc6aa13c098c8cf53cc4
---

# Vérifier la version installée

Le seul scénario où l’empreinte grimpe sensiblement est celui d’un backend systemd mal filtré, sans `journalmatch` ciblé, qui force Fail2ban à parcourir l’intégralité du journal système au lieu de se limiter aux unités pertinentes. C’est une raison de plus pour toujours définir `journalmatch` précisément plutôt que de laisser un jail lire le journal complet par défaut. Sur un VPS d’entrée de gamme avec 1 Go de RAM, Fail2ban ne représente généralement qu’une fraction marginale de la mémoire totale disponible, largement compensée par la réduction du bruit dans les journaux et la charge évitée sur le service SSH lui-même face aux scans automatisés.

## Fail2ban vs CrowdSec en 2026 : lequel choisir

La question revient systématiquement dès qu’on évoque le bannissement d’IP sur Linux : faut-il rester sur Fail2ban ou migrer vers CrowdSec, l’alternative plus récente qui mutualise les données de menaces entre utilisateurs via sa plateforme communautaire ? Les deux outils répondent à des besoins différents.

| Critère | Fail2ban | CrowdSec | 
|---|---|---|
| Version actuelle (été 2026) | 1.1.0 (avril 2024) | 1.7.8 (mai 2026) | 
| Étoiles GitHub | ~18 007 | ~13 941 | 
| Couche de protection | Réseau (L3/L4) : IP et ports | Applicatif (L7) : motifs HTTP, scénarios avancés | 
| Intelligence des menaces | Locale uniquement, à gérer manuellement | Communautaire, blocklists partagées en temps réel | 
| Cas d’usage idéal | Serveur unique, VPS, configuration simple et éprouvée | Infrastructures multi-serveurs, environnements distribués | 
| Courbe d’apprentissage | Faible, syntaxe INI classique | Modérée, concept de scénarios et bouncers | 

En résumé, pour un serveur unique ou un petit parc de VPS, Fail2ban reste un choix pragmatique et parfaitement suffisant en 2026 : sa maturité et sa simplicité de configuration en font une valeur sûre. CrowdSec devient pertinent à partir du moment où vous gérez plusieurs serveurs et que vous voulez bénéficier d’une intelligence de menaces mutualisée entre la communauté d’utilisateurs, avec une couverture qui va au-delà du simple filtrage réseau jusqu’aux motifs applicatifs de couche 7. Rien n’empêche non plus de faire tourner les deux en parallèle sur des périmètres différents, par exemple Fail2ban pour SSH et les services système historiques, et CrowdSec en frontal des applications web les plus exposées.

## Projet complet : configuration Fail2ban prête pour un VPS de production

Voici l’arborescence complète et fonctionnelle que vous pouvez reproduire telle quelle sur un VPS Ubuntu 24.04 fraîchement provisionné, combinant tout ce qui a été vu dans ce tutoriel.

```
# /etc/fail2ban/jail.local
[DEFAULT]
ignoreip = 127.0.0.1/8 ::1 VOTRE_IP_FIXE/32
bantime = 3600
bantime.increment = true
bantime.factor = 2
bantime.maxtime = 604800
findtime = 600
maxretry = 5
banaction = nftables
banaction_allports = nftables[type=allports]
destemail = [email protected]
sender = [email protected]
action = %(action_mwl)s
# /etc/fail2ban/jail.d/sshd.local.conf
[sshd]
enabled = true
port = ssh
filter = sshd
backend = systemd
journalmatch = _SYSTEMD_UNIT=ssh.service
maxretry = 4
bantime = 7200
# /etc/fail2ban/jail.d/web.local.conf
[nginx-http-auth]
enabled = true
filter = nginx-http-auth
logpath = /var/log/nginx/error.log
maxretry = 5
[nginx-botsearch]
enabled = true
filter = nginx-botsearch
logpath = /var/log/nginx/access.log
maxretry = 10
bantime = 86400
# /etc/fail2ban/jail.d/recidive.local.conf
[recidive]
enabled = true
filter = recidive
logpath = /var/log/fail2ban.log
bantime = 604800
findtime = 86400
maxretry = 3
```
Appliquez cette configuration puis validez qu’aucune erreur de syntaxe ne bloque le démarrage :

```
sudo fail2ban-client -t
sudo systemctl restart fail2ban
sudo fail2ban-client status
```
Résultat attendu si tout fonctionne correctement :

```
Status
|- Number of jail:      4
`- Jail list:   sshd, nginx-http-auth, nginx-botsearch, recidive
```
## Cinq pièges fréquents à éviter

La plupart des problèmes rencontrés avec Fail2ban proviennent d’un petit nombre d’erreurs récurrentes, largement documentées par la communauté et confirmées par les manuels officiels du projet.

1. **Oublier ignoreip et se bannir soi-même** : le classique absolu. Toujours ajouter son IP fixe ou son VPN avant d’activer un jail SSH strict.
2. **Mélanger logpath et journalmatch** : avec le backend systemd, spécifier`logpath` n’a aucun effet et peut empêcher le jail de démarrer. Utilisez exclusivement`journalmatch` dans ce cas.
3. **Modifier jail.conf au lieu de jail.local** : vos changements seront écrasés à la prochaine mise à jour du paquet, et sur Debian/Ubuntu ils sont de toute façon surchargés par`defaults-debian.conf` .
4. **Backend systemd sans python3-systemd installé** : le service ne pourra pas lire le journal et échouera silencieusement, sans message d’erreur explicite au premier abord.
5. **Ignorer IPv6** : historiquement, certaines actions n’étaient orientées que IPv4. Avec nftables, la gestion IPv6 est bien meilleure nativement, mais vérifiez tout de même que votre action de bannissement couvre les deux familles d’adresses, surtout si votre serveur a une adresse IPv6 publique active.

## Astuces avancées pour aller plus loin

Une fois la base solide en place, plusieurs pistes permettent d’affiner encore la protection sur des environnements plus exigeants.

- **Centraliser les bannissements sur un parc de serveurs** : Fail2ban ne partage rien entre machines par défaut. Pour synchroniser les listes de bannissement entre plusieurs serveurs, un script cron qui interroge`fail2ban-client status` sur chaque hôte et pousse les IP vers un pare-feu périmétrique commun reste la méthode la plus simple sans changer d’outil.
- **Combiner Fail2ban et un WAF** : Fail2ban agit sur la couche réseau, mais ne comprend pas le contenu applicatif HTTP en profondeur. Sur un site à fort trafic, l’associer à un pare-feu applicatif web en amont couvre les scénarios que Fail2ban ne voit pas.
- **Auditer les jails inutiles** : un jail activé sur un service que vous n’exposez pas consomme des ressources pour rien. Passez en revue`fail2ban-client status` régulièrement et désactivez ce qui ne correspond à aucun service réellement exposé.
- **Journaliser vers un SIEM** : pour les environnements avec obligations de conformité, exportez les logs Fail2ban vers un outil de corrélation comme Wazuh, qui permet de croiser les bannissements avec d’autres signaux de sécurité.
- **Créer des filtres personnalisés pour des applications maison** : si vous exposez une API interne avec son propre système d’authentification, écrivez un filtre dédié dans`filter.d/` plutôt que de vous limiter aux filtres livrés en standard. La syntaxe reste une expression régulière classique qui capture le champ`<HOST>` depuis votre format de log applicatif.
- **Tester avant un déploiement à grande échelle** : sur un parc de serveurs identiques, validez votre configuration`jail.local` sur une seule machine pendant quelques jours avant de la propager, afin de calibrer`maxretry` et`findtime` sur du trafic réel plutôt que sur des valeurs génériques.

## Dépannage : 8 problèmes courants et leurs solutions

Voici les incidents les plus fréquemment rencontrés lors du déploiement de Fail2ban, avec la cause probable et la solution associée.

