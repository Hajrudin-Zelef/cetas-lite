---
id: collect-261001-general-networking/general-networking/keycloak-sso-open-source-en-14-etapes-2026-5
title: "Base de données"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "distribution", "incident", "open source"]
source: docs/RAG/collect-261001-general-networking/keycloak-sso-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [350, 409]
sha256: 14a14972a38ed872ebc7080f1be2f511174167079ba271ad502b020e963552b2
---

# Base de données

- **Garder la base H2 embarquée en production.** Elle est parfaite pour un test de cinq minutes, mais elle ne supporte ni la montée en charge ni des sauvegardes fiables. Basculez sur PostgreSQL dès le premier déploiement réel.
- **Laisser les identifiants administrateur par défaut.** Un compte admin avec un mot de passe prévisible est la première cible des scans automatisés dès qu’un service Keycloak répond sur un port exposé.
- **Mal configurer les URI de redirection du client.** Une casse différente, un port manquant ou une barre oblique finale absente provoquent une erreur immédiate au moment du login.
- **Oublier KC_PROXY_HEADERS et KC_HOSTNAME derrière un reverse proxy.** Sans ces réglages, Keycloak continue de générer des liens en HTTP interne, ce qui casse le flux de connexion pour l’utilisateur final.
- **Laisser start-dev tourner en production.** Ce mode désactive plusieurs protections et n’est pas conçu pour encaisser du trafic public.
- **Ne jamais exporter le royaume avant une mise à jour majeure.** Un export récent transforme une mise à jour ratée en simple restauration plutôt qu’en incident majeur.
- **Distribuer le rôle d’administration du royaume trop largement.** Réservez les rôles realm-management aux comptes qui en ont réellement besoin, et documentez qui les détient.

## Dépannage : 10 problèmes courants et leurs solutions

Voici les symptômes les plus fréquemment rencontrés en suivant ce type de déploiement, avec leur cause probable et la correction à appliquer. Gardez cette table sous la main lors de votre première mise en production, la majorité des tickets de support liés à Keycloak se résument à l’un de ces dix cas.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| “Invalid parameter: redirect_uri” | L’URI de redirection du client ne correspond pas exactement à celle enregistrée | Vérifiez la casse, le port et la barre oblique finale dans la configuration du client | 
| Erreur 404 “Realm does not exist” | Le nom du royaume dans l’URL ne correspond à aucun royaume créé | Vérifiez l’orthographe exacte du royaume, sensible à la casse | 
| Boucle de redirection infinie après connexion | KC_HOSTNAME ou KC_PROXY_HEADERS mal configuré derrière le reverse proxy | Définissez explicitement KC_HOSTNAME et activez KC_PROXY_HEADERS=xforwarded | 
| Le conteneur Keycloak redémarre en boucle | La base PostgreSQL n’est pas encore prête au démarrage | Ajoutez une condition depends_on avec healthcheck sur le service postgres | 
| “Failed to obtain JDBC connection” | Identifiants incorrects ou base de données non initialisée | Vérifiez KC_DB_URL, KC_DB_USERNAME, KC_DB_PASSWORD et l’état du conteneur postgres | 
| Console admin inaccessible après passage en mode start | Le hostname strict bloque les accès sur un autre domaine que celui configuré | Vérifiez KC_HOSTNAME et utilisez le nom de domaine exact dans l’URL | 
| “Client not found” lors de l’appel au token endpoint | Le client_id est mal orthographié ou appartient à un autre royaume | Vérifiez le royaume dans l’URL et l’exactitude du client_id | 
| Codes OTP toujours rejetés | Décalage d’horloge entre le serveur et l’appareil de l’utilisateur | Synchronisez l’horloge du serveur via NTP | 
| L’import du royaume JSON échoue silencieusement | Le fichier exporté référence des identifiants qui n’existent plus sur la cible | Réexportez avec kc.sh export –realm et réimportez sur une base propre | 
| Le endpoint /metrics renvoie une erreur 404 | L’interface de gestion n’est pas exposée sur le port 9000 | Vérifiez que le port 9000 est publié et que KC_METRICS_ENABLED=true est défini | 

## Conseils avancés pour un déploiement en production

Une fois le socle stable, plusieurs réglages font la différence entre une instance qui tient une charge de test et une instance qui encaisse le trafic réel d’une entreprise. Ces réglages ne sont pas nécessaires le premier jour, mais mieux vaut les connaître avant qu’un pic de trafic ne les impose dans l’urgence.

### Cache distribué et haute disponibilité

Pour faire tourner plusieurs nœuds Keycloak derrière un répartiteur de charge, activez le cache distribué Infinispan en mode cluster plutôt que le cache local par défaut. Chaque nœud partage alors les sessions actives, ce qui évite qu’un utilisateur soit déconnecté simplement parce que sa requête suivante atterrit sur un autre nœud. Depuis la version 26.5.0 de janvier 2026, Keycloak expose également un support natif d’OpenTelemetry pour les métriques et les journaux, ce qui facilite son intégration dans une chaîne d’observabilité existante sans dépendre uniquement du endpoint Prometheus du port 9000. C’est le seuil à partir duquel il devient pertinent d’envisager l’Operator Keycloak officiel pour Kubernetes plutôt que Docker Compose.

Au-delà de la haute disponibilité, quatre autres réglages valent le détour pour une instance qui tourne en continu :

- **Supervision Prometheus.** Le endpoint /metrics du port 9000 s’intègre directement dans un tableau de bord Grafana pour suivre le taux d’échec de connexion et le temps de réponse du token endpoint.
- **Thèmes personnalisés.** Un thème custom déposé dans /opt/keycloak/themes permet d’habiller l’écran de connexion aux couleurs de votre marque, sans toucher au cœur de Keycloak.
- **Extensions SPI.** Les Service Provider Interfaces en Java permettent d’ajouter une règle métier, par exemple une validation de mot de passe spécifique à votre politique interne.
- **Limitation de débit en amont.** Un rate limiting configuré au niveau de Traefik ou d’un WAF freine une tentative de force brute avant même qu’elle n’atteigne Keycloak, en complément de la protection native du royaume.

## Le projet complet : ce que vous venez de construire

À ce stade, votre dossier `keycloak-stack` contient une stack Docker Compose à trois services (PostgreSQL, Keycloak, Traefik), un fichier `.env` avec vos secrets, un export JSON du royaume reproductible sur une autre machine, et un script de sauvegarde de base de données prêt à être placé dans une tâche cron. Ce socle gère l’authentification centralisée, le MFA avec passkeys, la fédération LDAP optionnelle et le TLS automatique, sans dépendre d’un fournisseur tiers.

Pour aller plus loin, committez le docker-compose.yml et l’export du royaume dans un dépôt git (jamais le fichier .env), ce qui transforme votre configuration Keycloak en infrastructure as code que n’importe quel membre de l’équipe peut reconstruire à l’identique en quelques minutes.

## Foire aux questions

### Keycloak est-il vraiment gratuit ?

Oui. Keycloak est publié sous licence Apache 2.0 par la Fondation Eclipse, sans coût de licence ni limite de comptes utilisateurs. Vous ne payez que l’infrastructure sur laquelle vous l’hébergez : serveur, base de données et certificat TLS.

### Quelle est la différence entre Keycloak et le Red Hat Build of Keycloak ?

Le Red Hat Build of Keycloak (RHBK) est la distribution commerciale construite sur le projet open source. Le code reste le même, mais Red Hat ajoute un support contractuel, des correctifs certifiés et une compatibilité garantie avec ses propres produits, notamment OpenShift.

### Keycloak peut-il remplacer Active Directory ?

Pas totalement. Keycloak gère l’authentification et les accès applicatifs, tandis qu’Active Directory gère aussi les postes de travail, les stratégies de groupe (GPO) et l’annuaire réseau Windows au sens large. Dans la pratique, la plupart des entreprises connectent Keycloak à leur Active Directory existant via la fédération LDAP plutôt que de le remplacer entièrement : Keycloak devient la couche SSO moderne posée par-dessus un annuaire qui continue de gérer le reste du parc informatique.

