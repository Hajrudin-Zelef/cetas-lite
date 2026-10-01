---
id: collect-261001-general-networking/general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026-5
title: "Mettre à jour la liste des paquets et le système"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["acquisition", "agents", "cyber", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026.md
source_anchor: ""
source_lines: [364, 432]
sha256: 89740e2b874d13810a8baebd02b665e716a93b86be0ca022ce73097aaa886541
---

# Mettre à jour la liste des paquets et le système

- **Le service ne démarre pas.** Lancez`sudo journalctl -u crowdsec -n 50 --no-pager` pour lire l’erreur exacte ; une faute de syntaxe YAML dans une acquisition ou un profil est la cause la plus fréquente. Validez avec`sudo cscli config show` .
- **Aucune IP n’est bloquée malgré des alertes.** Le moteur détecte mais le bouncer ne reçoit rien : vérifiez`cscli bouncers list` et la validité de la clé API du bouncer.
- **« unable to read API key » côté bouncer.** La clé du bouncer est manquante ou erronée. Régénérez-la avec`sudo cscli bouncers add mon-bouncer` et reportez-la dans le fichier de configuration du bouncer.
- **Zéro ligne lue (Acquisition).** Le chemin de logs est faux ou la distro journalise via systemd. Ajustez`acquis.yaml` ou ajoutez une source de type`journalctl` .
- **Les règles nftables ne s’appliquent pas.** Vous avez installé la variante iptables sur un système nftables. Désinstallez et installez`crowdsec-firewall-bouncer-nftables` .
- **La Console reste « déconnectée ».** Vérifiez`cscli capi status` et que le port sortant 443 n’est pas filtré par votre pare-feu sortant.
- **Trop de faux positifs.** Identifiez l’IP fautive dans`cscli alerts list` et ajoutez-la en liste blanche, puis rechargez le moteur.
- **La mise à jour du Hub échoue.** Exécutez`sudo cscli hub update` avant`upgrade` ; un index obsolète empêche la résolution des dépendances.

## Astuces avancées pour la production

Une fois les bases en place, ces optimisations font passer votre installation de CrowdSec au niveau professionnel.

- **Souscrivez à des blocklists supplémentaires.** En plus de la blocklist communautaire gratuite, la Console permet d’abonner gratuitement votre moteur à des listes thématiques (botnets, scanners, etc.) et, en option payante, à des listes premium à très faible taux de faux positifs.
- **Automatisez les mises à jour.** Une tâche cron hebdomadaire`cscli hub update && cscli hub upgrade && systemctl reload crowdsec` maintient votre détection alignée sur les dernières menaces.
- **Surveillez avec Prometheus.** Le moteur expose des métriques Prometheus (port 6060) ; couplées à Grafana, elles offrent une supervision fine des alertes et décisions.
- **Passez à PostgreSQL pour le multi-serveur.** Au-delà de quelques agents, remplacez SQLite par PostgreSQL pour gérer la concurrence d’écriture sur la LAPI centrale.
- **Combinez plusieurs bouncers.** Pare-feu + Nginx + Cloudflare offrent une défense en profondeur : le trafic malveillant est filtré au plus tôt, soulageant votre infrastructure.
- **Durcissez l’OS en parallèle.** CrowdSec complète, sans les remplacer, les bonnes pratiques : clés SSH plutôt que mots de passe, mises à jour automatiques de sécurité et pare-feu restrictif par défaut.

## Recevoir des alertes : notifications Slack, Discord et e-mail

Surveiller un tableau de bord en permanence n’est pas réaliste. CrowdSec intègre un système de *notifications* qui pousse les alertes vers vos canaux habituels – Slack, Discord, e-mail (SMTP), Microsoft Teams ou un webhook HTTP générique. Les plugins sont fournis d’origine ; il suffit de les configurer puis de les associer à un profil de décision.

```
# Lister les plugins de notification disponibles
ls /etc/crowdsec/notifications/
# Configurer le webhook HTTP (Slack, Discord, Teams...)
sudo nano /etc/crowdsec/notifications/http.yaml
# Activer la notification dans le profil de decision
sudo nano /etc/crowdsec/profiles.yaml
#   notifications:
#     - http_default
sudo systemctl reload crowdsec
```
Dans `profiles.yaml`, la clé `notifications` liste les canaux déclenchés à chaque décision. Vous pouvez filtrer pour n’être alerté que sur les bannissements de longue durée ou les scénarios critiques, afin d’éviter la fatigue d’alerte. Couplé à la Console et à un export Prometheus/Grafana, ce mécanisme ferme la boucle d’observabilité : détection, remédiation et, désormais, notification proactive de votre posture de sécurité. C’est l’un des chaînons qui font passer CrowdSec du simple outil défensif à une véritable plateforme de supervision.

## CrowdSec, RGPD et souveraineté numérique européenne

Pour les organisations européennes, CrowdSec présente un argument rarement mis en avant : c’est une solution de cybersécurité **européenne**, open source et transparente sur le traitement des données. À l’heure où la dépendance aux fournisseurs extra-européens est scrutée de près, déployer un moteur de détection français édité sous licence MIT s’inscrit pleinement dans une démarche de souveraineté numérique.

La question du RGPD se pose légitimement, puisqu’une adresse IP est une donnée à caractère personnel au sens de la CNIL. CrowdSec a conçu son modèle de partage en conséquence : lorsque vous contribuez à la communauté, seuls des signaux d’attaque réduits et pseudonymisés sont transmis (l’IP de l’attaquant, le scénario déclenché et un horodatage), et non le contenu de vos logs ni les données de vos utilisateurs légitimes. Le partage reste par ailleurs configurable. Cette approche défensive collective fait écho aux orientations de l’ANSSI et au renforcement de la posture cyber nationale.

Concrètement, pour une mise en conformité propre : documentez l’usage de CrowdSec dans votre registre de traitements, mentionnez la finalité (sécurité du système d’information, base légale de l’intérêt légitime), et configurez vos durées de rétention des alertes. L’open source facilite ici l’audit : le code étant public sur GitHub, rien n’est opaque. Pour aller plus loin sur la gestion centralisée des secrets et identités, notre guide Vaultwarden : auto-héberger Bitwarden complète idéalement cette démarche d’autonomie.

## Récapitulatif : votre projet CrowdSec complet

Voici la check-list du projet fonctionnel que vous venez de bâtir. Si chaque case est cochée, votre serveur dispose d’une défense collaborative de bout en bout, des couches réseau et applicative jusqu’à la supervision.

- **Moteur de sécurité installé** et actif (`cscli metrics` affiche des lignes lues).
- **Bouncer pare-feu** connecté à la LAPI (`cscli bouncers list` ).
- **SSH protégé** via la collection`crowdsecurity/sshd` .
- **Nginx protégé** en détection (collection) et en remédiation (bouncer Nginx).
- **AppSec / WAF** à l’écoute sur 127.0.0.1:7422 pour le virtual patching.
- **Console enrôlée** et CTI communautaire actif (`cscli capi status` ).
- **Listes blanches** en place pour vos IP d’administration.
- **Notifications** branchées sur Slack, Discord ou e-mail.
- **Mises à jour du Hub** automatisées par une tâche cron.

Ce socle couvre l’essentiel d’une protection sérieuse. Pour une infrastructure plus large, dupliquez le moteur sur chaque hôte et centralisez la LAPI comme décrit à l’étape 12, puis pilotez l’ensemble depuis une seule Console.

## Foire aux questions (FAQ)

### CrowdSec est-il vraiment gratuit ?

Oui. Le moteur de sécurité, les bouncers, le Hub, la blocklist communautaire et la Console (formule de base) sont gratuits et open source sous licence MIT. Seules certaines blocklists premium et des fonctionnalités d’entreprise sur la Console relèvent d’offres payantes, totalement optionnelles pour un usage personnel ou PME : depuis mai 2026, la grille tarifaire de CrowdSec propose par exemple un accès API au CTI à partir de **49 $/mois** pour 5 000 requêtes, des Platinum Blocklists à **900 $/mois** pour un usage individuel ou **3 900 $/mois** en usage illimité, une réplication locale du CTI (Local CTI Replication) à **9 000 $/mois**, ainsi qu’une offre Console Premium à **1 000 $/mois** – soit six formules payantes actives recensées en septembre 2026, aux côtés de l’offre gratuite qui reste, elle, inchangée.

### CrowdSec remplace-t-il mon pare-feu ?

