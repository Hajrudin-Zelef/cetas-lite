---
id: collect-261001-rattrapage/rattrapage/netbox-guide-15
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "distribution"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [2864, 3019]
sha256: 4307ceaf45649b81a151f9b128b23804e7a0af90a5f62f4b06b692d152488ae7
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

**Contexte** : le RSSI demande « qu'est-ce qui est exposé sur le site de Lyon ? ».
Réponse en 30 minutes, sans scanner réseau :

1. **Périmètre** : filtre `site=AGENCE-Lyon, statut=Actif` sur les équipements →
   export CSV.
2. **Services exposés** : onglet Services (section 39) → listez `https`, `ssh`,
   `telnet` (⚠️ si vous en trouvez, c'est un finding d'audit à lui tout seul).
3. **Management** : rapport « équipements sans IP de management » (section 59) →
   tout équipement administrable sans IP tracée = zone d'ombre.
4. **Comptes** : Admin > Utilisateurs → derniers accès, comptes inactifs depuis
   > 90 jours → à désactiver (section 85).
5. **Accès API** : tokens sans expiration, tokens `write` inutilisés → révoquer.
6. **Livrable** : un export CSV + 10 lignes de synthèse. Le RSSI adore les tableaux
   courts avec des chiffres.

---

## 91. Pense-bête de poche — commandes essentielles

### Services et logs

```bash
systemctl status netbox netbox-rq postgresql redis-server nginx --no-pager
systemctl restart netbox netbox-rq          # après modif configuration.py
journalctl -u netbox -n 50 --no-pager       # logs applicatifs
tail -50 /var/log/netbox/gunicorn-error.log # erreurs Django/gunicorn
tail -50 /var/log/nginx/error.log           # erreurs frontal
```

### Django / NetBox (utilisateur netbox, venv activé)

```bash
cd /opt/netbox/netbox
python netbox/manage.py check               # santé de la configuration
python netbox/manage.py migrate             # appliquer les migrations
python netbox/manage.py collectstatic --no-input
python netbox/manage.py createsuperuser     # (une seule fois)
python netbox/manage.py changepassword admin
python netbox/manage.py clearsessions       # ménage hebdomadaire
python netbox/manage.py rqworker high default low   # worker manuel (debug)
```

### Base de données

```bash
sudo -u postgres pg_dump -Fc netbox > /var/backups/netbox/netbox-$(date +%Y%m%d).dump
psql -h localhost -U netbox -d netbox -c \
  "SELECT pg_size_pretty(pg_database_size('netbox'));"
```

### API (rappel)

```bash
export NETBOX_TOKEN="..."
curl -s -H "Authorization: Token $NETBOX_TOKEN" \
  https://netbox.lan-entreprise.fr/api/dcim/devices/?status=active \
  | python3 -m json.tool | head -30
```

---

## 92. Pense-bête de poche — où cliquer pour quoi

| Je veux… | Aller à… |
|---|---|
| Trouver une IP / un équipement | Barre de recherche globale (en haut) |
| Voir l'utilisation d'un sous-réseau | IPAM > Préfixes > cliquer le préfixe |
| Savoir ce qu'il y a dans une baie | DCIM > Baies > cliquer la baie (vue U) |
| Tracer un câble de bout en bout | Fiche interface > onglet Câble > chemin |
| Voir qui a modifié quoi | Fiche objet > onglet Journal |
| Créer un token API | Profil (avatar) > API Tokens |
| Lancer un script | Admin > Scripts |
| Voir les rapports | Admin > Rapports |
| Importer en masse | Liste d'objets > Importer |
| Exporter en CSV | Liste d'objets > Exporter |
| Gérer les utilisateurs/permissions | Admin > Utilisateurs / Groupes / Permissions |
| Voir les files de tâches | Admin > Files d'attente RQ |
| Configurer un webhook | Admin > Webhooks |

---

## 93. Glossaire

| Terme | Définition |
|---|---|
| Agrégat | Bloc d'adresses de haut niveau (ex. `10.0.0.0/8`), rattaché à un RIR |
| API REST | Interface HTTP JSON permettant de lire/écrire dans NetBox par programme |
| ASN | Autonomous System Number — numéro de système autonome (BGP) |
| Baie (rack) | Armoire 19 pouces recevant les équipements, mesurée en U |
| Câble (cable) | Liaison physique entre deux terminaisons (interfaces, ports, alimentations) |
| Changement (changelog) | Journal d'audit : qui a modifié quoi et quand |
| Champ personnalisé | Attribut supplémentaire typé ajouté à un modèle NetBox |
| Circuit | Liaison opérateur/FAI documentée (débit, identifiants, terminaisons) |
| Container (statut) | Préfixe parent découpé en sous-préfixes, sans adresses directes |
| DCIM | Data Center Infrastructure Management — le module « physique » |
| Équipement (device) | Instance physique d'un type d'équipement, installée dans une baie/un site |
| FQDN | Fully Qualified Domain Name — nom d'hôte complet (`netbox.lan-entreprise.fr`) |
| Gunicorn | Serveur d'application Python (WSGI) exécutant NetBox |
| IPAM | IP Address Management — le module « adressage logique » |
| OOB | Out-Of-Band — réseau d'administration séparé du trafic de production |
| PDU | Power Distribution Unit — multiprise de baie alimentée par l'onduleur |
| Plage IP (IP range) | Intervalle réservé dans un préfixe (DHCP, VIP…) |
| Préfixe | Sous-réseau en notation CIDR (`10.10.20.0/24`) |
| RIR | Regional Internet Registry — registre régional d'adresses (ex. RIPE-NCC) |
| Rôle | Qualification d'usage (rôle d'équipement, de préfixe, de baie…) |
| RQ (django-rq) | Files de tâches asynchrones (webhooks, scripts, rapports) |
| Script (custom script) | Formulaire exécutable dans NetBox pour automatiser une action |
| Site | Lieu physique documenté (bureaux, agence, datacenter) |
| Source de vérité | Référentiel unique faisant foi sur l'état de l'infrastructure |
| Tag | Étiquette libre et transverse apposée sur les objets |
| Tenant | Entité propriétaire d'objets (client, département, projet) |
| Type d'équipement | Modèle physique (fabricant + modèle + composants types) |
| U | Unité de hauteur en baie (1U = 44,45 mm) |
| VID | VLAN Identifier — numéro du VLAN (1–4094) |
| VLAN | Virtual LAN — domaine de broadcast de niveau 2 |
| VRF | Virtual Routing and Forwarding — table de routage isolée |
| Webhook | Notification HTTP envoyée par NetBox à chaque événement objet |
| WSGI | Interface Python entre serveur web et application (ici : gunicorn) |

---

## 94. Quiz — 10 questions pour valider vos acquis

**Q1.** Que se passe-t-il si vous créez deux fois la même adresse IP dans la même VRF ?
**Q2.** Quelle est la différence entre un préfixe au statut `Container` et un préfixe `Actif` ?
**Q3.** Citez les 4 briques techniques indispensables pour faire tourner NetBox 4.x en production.
**Q4.** Pourquoi faut-il lancer `migrate` AVANT le premier démarrage de gunicorn ?
**Q5.** Un technicien doit modifier uniquement les câbles du site de Lyon : quel mécanisme
de NetBox utilisez-vous, et à quel niveau l'appliquez-vous ?
**Q6.** Où sont stockées les pièces jointes (photos de baies), et que faut-il sauvegarder
en plus de la base PostgreSQL ?
**Q7.** Votre inventaire Ansible doit refléter NetBox en temps réel : quelle solution
recommandée par ce guide, et quel est le principal bénéfice ?
**Q8.** Lors d'une maintenance onduleur, comment NetBox vous aide-t-il à établir la liste
exacte des équipements impactés ?
**Q9.** Un webhook n'est jamais reçu par votre application : quels sont les deux premiers
points à vérifier côté NetBox ?
**Q10.** Citez trois éléments de la check-list de durcissement (section 46).

### Réponses

**R1.** NetBox refuse le doublon : l'unicité d'une adresse IP est garantie par VRF.
C'est une protection volontaire, pas un bug (voir section 75). Si le doublon semble
légitime, c'est probablement qu'il faut deux VRF distinctes.

**R2.** `Container` = bloc parent uniquement découpé en sous-préfixes (ex. `10.10.0.0/16`),
on n'y met jamais d'adresses IP directement. `Actif` = préfixe en production, qui
contient les adresses (voir section 33).

**R3.** PostgreSQL (données), Redis (cache + files de tâches), gunicorn (serveur
d'application Python), nginx (reverse proxy HTTPS). Le tout supervisé par systemd
(`netbox.service` + `netbox-rq.service`) — voir sections 3, 13, 14.

**R4.** Parce que `migrate` crée les tables de la base. Sans elles, gunicorn démarre
sur une base vide et chaque page renvoie une erreur 500 (voir section 12).

