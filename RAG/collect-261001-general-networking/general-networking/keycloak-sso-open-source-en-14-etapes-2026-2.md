---
id: collect-261001-general-networking/general-networking/keycloak-sso-open-source-en-14-etapes-2026-2
title: "Base de données"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/keycloak-sso-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [56, 167]
sha256: bb0f71b9862c2e6ec8348c743a6d8bb7201ada34995b1092f7bf5f3726853684
---

# Base de données

Le fichier `.env` centralise tous les secrets et paramètres que docker-compose injectera dans les conteneurs. Générez des mots de passe robustes, pas les valeurs d’exemple ci-dessous.

```
# Base de données
POSTGRES_DB=keycloak
POSTGRES_USER=keycloak
POSTGRES_PASSWORD=changez-ce-mot-de-passe-maintenant
# Administration Keycloak (remplace l'ancien KEYCLOAK_ADMIN, deprecie depuis la v26)
KC_ADMIN_USER=admin
KC_ADMIN_PASSWORD=un-autre-mot-de-passe-solide
# Nom d'hote public utilise en production
KC_HOSTNAME=auth.votre-domaine.fr
```
C’est le premier piège classique du déploiement Keycloak : les tutoriels rédigés avant fin 2024 utilisent encore les variables `KEYCLOAK_ADMIN` et `KEYCLOAK_ADMIN_PASSWORD`. Elles fonctionnent toujours en 26.7.1 mais affichent un avertissement de dépréciation dans les journaux. Les variables recommandées aujourd’hui sont `KC_BOOTSTRAP_ADMIN_USERNAME` et `KC_BOOTSTRAP_ADMIN_PASSWORD`, utilisées dans le fichier docker-compose ci-dessous. Gardez aussi un œil sur le rythme des correctifs de sécurité : la 26.5.7, publiée en avril 2026, corrige à elle seule sept CVE d’après le projet Keycloak, et Clever Cloud a répercuté cette même mise à jour sur son offre Keycloak hébergée dès avril 2026, signe que le projet publie des correctifs à un rythme soutenu qu’il vaut mieux suivre plutôt qu’ignorer. Même les bibliothèques clientes suivent ce tempo : la version 26.0.10 a été annoncée le 30 juin 2026 par Peter Skopek sur le groupe Google Keycloak Dev, et la 26.0.12 figure comme version des sources téléchargeables sur la page Downloads officielle de keycloak.org au 1er septembre 2026 — un rappel qu’il vaut mieux épingler une version précise plutôt que suivre une branche à l’aveugle dans vos dépendances Maven ou npm.

## Étape 3 : Écrire le fichier docker-compose.yml

Voici la configuration complète pour démarrer Keycloak en mode développement, adossé à PostgreSQL plutôt qu’à la base H2 embarquée par défaut. Utiliser PostgreSQL dès le premier test évite d’avoir à migrer une base plus tard.

```
services:
  postgres:
    image: postgres:16
    container_name: keycloak-postgres
    restart: unless-stopped
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    volumes:
      - ./data/postgres:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER}"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - keycloak-net
  keycloak:
    image: quay.io/keycloak/keycloak:26.7.0
    container_name: keycloak
    command: start-dev
    restart: unless-stopped
    depends_on:
      postgres:
        condition: service_healthy
    environment:
      KC_DB: postgres
      KC_DB_URL: jdbc:postgresql://postgres:5432/${POSTGRES_DB}
      KC_DB_USERNAME: ${POSTGRES_USER}
      KC_DB_PASSWORD: ${POSTGRES_PASSWORD}
      KC_BOOTSTRAP_ADMIN_USERNAME: ${KC_ADMIN_USER}
      KC_BOOTSTRAP_ADMIN_PASSWORD: ${KC_ADMIN_PASSWORD}
      KC_METRICS_ENABLED: "true"
      KC_HEALTH_ENABLED: "true"
    ports:
      - "8080:8080"
      - "9000:9000"
    volumes:
      - ./data/keycloak:/opt/keycloak/data
    networks:
      - keycloak-net
networks:
  keycloak-net:
    driver: bridge
```
Le port 9000 mérite une explication : depuis la branche 26, Keycloak sépare son interface de gestion (santé, métriques Prometheus) du port applicatif 8080. C’est une bonne pratique de sécurité, elle évite d’exposer les endpoints de supervision sur la même interface que le trafic public. Vous ne l’ouvrirez d’ailleurs jamais publiquement en production, uniquement à votre système de monitoring interne.

## Étape 4 : Démarrer les conteneurs et vérifier les journaux

Lancez la stack en arrière-plan, puis suivez les journaux de Keycloak pour confirmer que le démarrage se termine sans erreur.

```
docker compose up -d
docker compose logs -f keycloak
```
Un démarrage réussi affiche une sortie proche de celle-ci, avec un temps de démarrage généralement compris entre 5 et 15 secondes sur une machine de développement :

```
keycloak  | Keycloak 26.7.0 on JVM (powered by Quarkus 3.x) started in 8.412s
keycloak  | Listening on: http://0.0.0.0:8080
keycloak  | Management interface listening on http://0.0.0.0:9000
keycloak  | Profile prod activated.
keycloak  | Config source PropertiesConfigSource loaded
```
Si le conteneur redémarre en boucle à la place, la cause la plus fréquente est une base PostgreSQL pas encore prête. Le `healthcheck` et la condition `service_healthy` définis à l’étape précédente sont justement là pour éviter ce cas, mais vérifiez tout de même les journaux du service postgres avec `docker compose logs postgres` si le problème persiste.

## Étape 5 : Première connexion à la console d’administration

Ouvrez `http://localhost:8080/admin/master/console/` dans votre navigateur. Connectez-vous avec les identifiants définis dans `KC_ADMIN_USER` et `KC_ADMIN_PASSWORD`. Vous atterrissez dans le royaume `master`, celui qui gère Keycloak lui-même : ne créez jamais vos utilisateurs applicatifs ici, il sert uniquement à l’administration de la plateforme.

Première chose à faire une fois connecté : si vous comptez exposer cette instance au-delà de votre machine locale, changez immédiatement le mot de passe administrateur depuis l’interface, même s’il provient d’un fichier `.env` que vous jugez robuste. Un compte admin par défaut est la cible numéro un des scans automatisés dès qu’un service Keycloak est détecté sur un port public.

## Étape 6 : Créer un royaume dédié à votre application

Un royaume (realm) est un espace isolé : ses utilisateurs, ses rôles et sa configuration ne se mélangent jamais avec un autre royaume. C’est la brique qui permet, par exemple, de séparer proprement un environnement de recette et un environnement de production sur la même instance Keycloak, ou d’isoler chaque client si vous construisez une plateforme multi-tenant — pour ce dernier cas, la fonctionnalité Organizations, passée en support complet dès la version 26.0.0 d’octobre 2024 et dont le statut a été reconfirmé par la documentation Keycloak en janvier 2025, offre désormais une alternative plus légère à la multiplication des royaumes. Depuis le menu déroulant en haut à gauche de la console, cliquez sur **Create realm**, nommez-le par exemple `entreprise-prod`, puis validez.

Vous pouvez aussi automatiser cette étape avec `kcadm.sh`, le client en ligne de commande embarqué dans l’image Keycloak, particulièrement utile si vous scriptez un déploiement reproductible :

```
docker exec -it keycloak /opt/keycloak/bin/kcadm.sh config credentials \
  --server http://localhost:8080 --realm master \
  --user admin --password 'un-autre-mot-de-passe-solide'
docker exec -it keycloak /opt/keycloak/bin/kcadm.sh create realms \
  -s realm=entreprise-prod -s enabled=true
```
## Étape 7 : Durcir la politique de sécurité du royaume

Avant de créer le moindre utilisateur, réglez la politique de mots de passe et la protection contre les attaques par force brute. Dans **Realm settings → Security defenses**, activez la détection de force brute et fixez un seuil raisonnable, par exemple un verrouillage temporaire après 5 tentatives échouées. Dans **Authentication → Policies**, définissez une longueur minimale de mot de passe d’au moins 12 caractères et exigez au moins un chiffre et une casse mixte.

Réglez également la durée de vie des sessions dans **Realm settings → Sessions**. La valeur par défaut convient pour une démonstration, mais en production, un timeout d’inactivité de 30 minutes et une durée de vie maximale de session de 10 heures constituent un point de départ raisonnable pour une application interne d’entreprise.

