---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-5
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [683, 843]
sha256: f90d294b4069e7a1858757e7620ea41f6e2742218e22944c02d2902ca72aa3d3
---

# Guide Zabbix complet — Supervision d'infrastructure en production

### 20.1. PSK (Pre-Shared Key) — simple et efficace

Générez une clé par hôte (ou par groupe) :

```bash
openssl rand -hex 32 > /etc/zabbix/psk_srv-fichiers-01.key
chmod 640 /etc/zabbix/psk_srv-fichiers-01.key
chown root:zabbix /etc/zabbix/psk_srv-fichiers-01.key
cat /etc/zabbix/psk_srv-fichiers-01.key
# exemple : 9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e1f0a9b8c7d6e5f4a3b2c1d0e
```

Côté agent (`zabbix_agent2.conf`) :

```ini
TLSConnect=psk
TLSAccept=psk
TLSPSKFile=/etc/zabbix/psk_srv-fichiers-01.key
TLSPSKIdentity=psk-srv-fichiers-01
```

Côté frontend, sur l'hôte : *Configuration → Hosts → srv-fichiers-01 → Encryption* : cocher PSK pour les connexions, renseigner **la même** identité PSK et la clé.

Pour le proxy : paramètres `TLSConnect`, `TLSAccept`, `TLSPSKFile`, `TLSPSKIdentity` dans `zabbix_proxy.conf`, et déclaration PSK dans *Administration → Proxies*.

### 20.2. Certificats — pour les environnements exigeants

```ini
# zabbix_agent2.conf (extrait)
TLSConnect=cert
TLSAccept=cert
TLSCAFile=/etc/zabbix/ca.crt
TLSCertFile=/etc/zabbix/agent.crt
TLSKeyFile=/etc/zabbix/agent.key
```

> 💡 **Recommandation pragmatique** : PSK suffit dans 90 % des cas internes. Réservez les certificats aux flux traversant des réseaux non maîtrisés, avec une CA interne (section "Pour aller plus loin").

## 21. Premier démarrage : assistant de configuration du frontend

Ouvrez `http://zabbix.votre-domaine.local` (ou l'IP du server). L'assistant en 6 étapes :

1. **Welcome** → Next.
2. **Check of pre-requisites** : tout doit être vert (version PHP, extensions, `max_execution_time`, fuseau horaire). ⚠️ Si `Time zone` est rouge : réglez `date.timezone` dans `/etc/php/8.2/fpm/php.ini` (ex. `Africa/Abidjan`, `Africa/Dakar`… selon votre pays) puis `systemctl restart php8.2-fpm`.
3. **Configure DB connection** : type PostgreSQL (ou MySQL), host, port, nom, user, `mot_de_passe_ici`.
4. **Settings** : nom du server Zabbix (ex. `Zabbix-Production`), fuseau horaire par défaut, thème.
5. **Pre-installation summary** → Next.
6. **Install** → Finish.

Identifiants par défaut : **Admin / zabbix**. ⚠️ **Changez ce mot de passe immédiatement** (*Administration → Users → Admin → Password*) et créez un compte nominatif par exploitant (section 23).

🖨️ **Checklist post-installation immédiate :**

- [ ] Mot de passe `Admin` changé
- [ ] Fuseau horaire correct (`Africa/...`)
- [ ] `zabbix_server.log` sans erreur
- [ ] Frontend accessible, page *Monitoring → Dashboard* affichée
- [ ] Snapshot VM / sauvegarde de la base faite (base vide = restauration triviale)

## 22. Configuration initiale : réglages globaux essentiels

*Administration → General* — les réglages à passer en revue dès le premier jour :

### 22.1. GUI et fuseau horaire

- *Administration → General → GUI* : thème par défaut, limite de recherche, périodes de graphes.
- Chaque utilisateur peut surcharger son fuseau horaire (*User settings → Time zone*). Imposez le fuseau du site principal par défaut.

### 22.2. Housekeeping (rétention) — à régler AVANT de collecter

*Administration → General → Housekeeping* :

| Donnée | Réglage conseillé (départ) | Commentaire |
|---|---|---|
| History (données brutes) | 90 jours | Au-delà, les trends suffisent |
| Trends (agrégats horaires) | 730 jours (2 ans) | Pour les rapports annuels |
| Events / alertes | 365 jours | Traçabilité |
| Audit | 365 jours | Exigences internes |

⚠️ **Ne mettez jamais History à 365 jours "au cas où"** : c'est le premier facteur d'explosion de la base. Les trends (min/max/avg/count par heure) suffisent pour l'analyse long terme.

### 22.3. Autres réglages globaux importants

- *Administration → General → Other* : `Refresh unsupported items` (défaut 10 min — mettez 1 h pour limiter le bruit), groupes d'hôtes par défaut pour la découverte.
- *Administration → Media types* : vérifiez que les médias Email/Telegram sont configurés (sections 41-43).
- *Administration → General → Images / Icon mapping* : pour de belles cartes réseau (section 65).

## 23. Utilisateurs, groupes, rôles et permissions

> 💡 **Règle d'or** : jamais de compte partagé `admin` au quotidien. Un compte nominatif par personne, avec le rôle minimal.

### 23.1. Rôles (⚠️ Zabbix 5.4+/6.x : le système de rôles remplace les anciens "user types")

| Rôle | Usage type |
|---|---|
| Super admin | Vous, et votre second. Configuration complète. |
| Admin | Chefs d'équipe : gèrent hôtes/templates de leur périmètre. |
| User (opérateur) | Exploitation : dashboards, acquittements, maintenance. |
| Guest (désactivé par défaut) | À laisser désactivé. |

Créez des rôles sur mesure (*Administration → User roles*) : par exemple un rôle "Opérateur N1" qui peut acquitter et mettre en maintenance mais pas modifier les templates.

### 23.2. Groupes d'utilisateurs et permissions par groupe d'hôtes

1. *Administration → User groups → Create* : `Exploitation`, `Energie`, `Direction` (lecture seule).
2. Pour chaque groupe : onglet *Permissions*, ajoutez les **groupes d'hôtes** avec le niveau *Read* ou *Read-write*.
3. Onglet *Users* : rattachez les comptes nominatifs.

Exemple de matrice :

| Groupe d'hôtes | Exploitation | Energie | Direction |
|---|---|---|---|
| `Serveurs` | Read-write | Read | Read |
| `Onduleurs` | Read | Read-write | Read |
| `Réseau` | Read-write | Deny | Read |

### 23.3. Authentification

- *Administration → Authentication* : interne par défaut. Activez **LDAP/AD** si disponible (moins de mots de passe à gérer), et imposez le **2FA (TOTP)** pour les Super admins (⚠️ natif depuis 6.4 ; en 6.0 utilisez un plugin MFA).
- Verrouillez les tentatives : *Administration → Authentication → Login attempts* (ex. 5 essais → blocage 30 s).

## 24. Audit et traçabilité

Tout est journalisé : *Administration → Audit* (qui a modifié quel trigger, quand). Le *Audit log* frontend et la table `auditlog` en base sont vos alliés en cas d'incident de configuration ("qui a désactivé cette alerte ?").

> 💡 En production, conservez l'audit **365 jours minimum** et exportez-le avant toute mise à jour (section 72).

## 25. Concepts clés : hôtes, groupes d'hôtes, interfaces

- **Host (hôte)** : tout équipement supervisé (serveur, switch, onduleur, imprimante…).
- **Host group** : regroupement logique (`Serveurs/Linux`, `Onduleurs`, `Réseau/Switchs`). Un hôte appartient à **au moins un** groupe. Les permissions et les actions s'appuient dessus.
- **Interface** : comment joindre l'hôte — Agent (10050), SNMP (161), IPMI (623), JMX. Un hôte peut avoir plusieurs interfaces.
- **Template** : paquet réutilisable d'items/triggers/graphes (sections 27-29).

**Convention de nommage conseillée** (à adapter) :

```
<type>-<fonction>-<numéro>   ex. : srv-fichiers-01, sw-coeur-01, ups-salle-01
```

> 💡 Un nom d'hôte **stable et sans espace** évite 80 % des problèmes : c'est la clé de jointure entre l'agent actif (`Hostname=`) et le frontend.

## 26. Ajouter son premier hôte pas à pas

Exemple : `srv-fichiers-01` (192.168.10.61, Agent 2 en mode actif).

1. *Configuration → Hosts → Create host*.
2. Onglet **Host** :
   - *Host name* : `srv-fichiers-01` (exactement le `Hostname=` de l'agent !)
   - *Visible name* : `Serveur fichiers 01 (bureaux)` (libellé lisible)
   - *Groups* : `Serveurs/Linux`
   - *Interfaces* : *Add → Agent*, IP `192.168.10.61`, port `10050`
   - *Monitored by* : `Server` (ou le proxy du site)
3. Onglet **Templates** : *Link new templates* → `Linux by Zabbix agent` → *Add*.
4. *Add* en bas de page.
5. Vérifiez : la colonne **Availability** doit passer au vert (icône ZBX) sous 2-3 minutes.

Dépannage express si ça reste gris/rouge : section 74, erreur n°2.

### Ajout en masse : l'import CSV et l'API

