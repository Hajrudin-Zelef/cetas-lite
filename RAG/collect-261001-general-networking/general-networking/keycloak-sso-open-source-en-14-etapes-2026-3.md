---
id: collect-261001-general-networking/general-networking/keycloak-sso-open-source-en-14-etapes-2026-3
title: "Base de données"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/keycloak-sso-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [168, 265]
sha256: 645cd0e75ca3edde68542a7ff96e7c73e4feb9364ff6c92cb296991617c63fb5
---

# Base de données

## Étape 8 : Créer un client OIDC et configurer les URI de redirection

Un client représente une application qui va déléguer son authentification à Keycloak. Dans **Clients → Create client**, choisissez le protocole `openid-connect`, donnez-lui un identifiant comme `mon-app-web`, puis passez à l’écran de configuration des flux.

### Client confidentiel ou client public

Un client confidentiel possède un secret partagé et convient à une application côté serveur, qui peut garder ce secret hors de portée du navigateur. Un client public, sans secret, s’utilise pour une application mobile ou un frontend JavaScript pur, en s’appuyant sur PKCE (Proof Key for Code Exchange) pour sécuriser l’échange malgré l’absence de secret. Pour une application web classique avec un backend, choisissez confidentiel. Pour une SPA React ou Vue sans backend dédié, choisissez public avec PKCE activé.

Renseignez ensuite les **Valid redirect URIs**, par exemple `https://mon-app.exemple.fr/callback`. C’est le champ le plus souvent mal configuré du tutoriel : une erreur de casse, un port manquant ou une barre oblique finale absente provoquent immédiatement une erreur au moment du login. Gardez la valeur strictement identique à celle utilisée par votre application.

## Étape 9 : Créer des utilisateurs, groupes et rôles

Dans **Users → Add user**, créez un premier compte de test avec une adresse e-mail et un nom d’utilisateur. Dans l’onglet **Credentials**, définissez un mot de passe temporaire en cochant **Temporary**, ce qui forcera l’utilisateur à le changer à sa première connexion.

Créez ensuite au moins deux rôles dans **Realm roles**, par exemple `utilisateur` et `administrateur`, puis assignez-les depuis l’onglet **Role mapping** du profil utilisateur. Pour un scénario multi-équipes, régler des groupes dans **Groups** plutôt que d’assigner les rôles individuellement facilite grandement la maintenance à mesure que le nombre de comptes augmente. Le même résultat en ligne de commande :

```
docker exec -it keycloak /opt/keycloak/bin/kcadm.sh create users \
  -r entreprise-prod \
  -s username=test.utilisateur \
  -s [email protected] \
  -s enabled=true
```
## Étape 10 : Activer le MFA, WebAuthn et les passkeys

Dans **Authentication → Required actions**, activez **Configure OTP** pour le TOTP classique (Google Authenticator, Aegis, ou équivalent), et **Webauthn Register Passwordless** pour les passkeys. Vous pouvez rendre l’une de ces actions obligatoire pour tous les nouveaux comptes en cochant **Set as default action**.

### WebAuthn et passkeys, la vraie nouveauté de ce cycle de versions

Le support des passkeys s’est nettement stabilisé depuis la version 22 et arrive à maturité dans la branche 26. Concrètement, un utilisateur peut désormais se connecter avec Face ID, Windows Hello ou une clé de sécurité physique, sans jamais saisir de mot de passe. Réglez la politique WebAuthn dans **Authentication → Policies → WebAuthn Passwordless Policy** : signature algorithm ES256, attestation non requise pour un usage interne, et `discoverable credentials` réglé sur `required` si vous voulez une connexion sans même saisir de nom d’utilisateur.

## Étape 11 : Connecter un annuaire LDAP ou Active Directory

Si votre organisation dispose déjà d’un annuaire, inutile de recréer tous les comptes à la main, et c’est justement l’un des cas d’usage où Keycloak brille par rapport à une solution SaaS qui exige souvent un connecteur payant pour la fédération LDAP. Dans **User federation → Add provider → ldap**, renseignez l’URL de connexion (`ldaps://votre-serveur:636`), le DN de liaison, le mot de passe du compte de service et le DN de base de recherche des utilisateurs. Réglez **Edit mode** sur `READ_ONLY` si Keycloak ne doit jamais écrire dans votre annuaire existant, ce qui est le choix le plus sûr pour une première intégration. Réservez le mode `WRITABLE` aux cas où vous voulez explicitement gérer la création de comptes depuis Keycloak plutôt que depuis votre annuaire d’origine.

Lancez une synchronisation manuelle avec le bouton **Sync all users** pour valider la configuration avant de programmer une synchronisation périodique. Keycloak ne remplace pas Active Directory : il vient s’y brancher pour offrir le SSO et les protocoles modernes (OIDC, SAML) à des applications qui ne savent pas parler LDAP nativement.

## Étape 12 : Déployer un reverse proxy TLS avec Traefik

Keycloak ne doit jamais servir de trafic HTTP en clair sur un domaine public. Les jetons d’accès, les cookies de session et les mots de passe transiteraient alors en clair sur le réseau, ce qui annule d’un coup tout le travail fait sur le MFA et la politique de mots de passe des étapes précédentes. Ajoutez Traefik à votre docker-compose.yml pour gérer automatiquement le certificat Let’s Encrypt et rediriger tout le trafic vers HTTPS. Nginx ou Caddy fonctionnent tout aussi bien si votre équipe les maîtrise déjà mieux.

```
  traefik:
    image: traefik:v3.2
    container_name: traefik
    restart: unless-stopped
    command:
      - "--providers.docker=true"
      - "--providers.docker.exposedbydefault=false"
      - "--entrypoints.websecure.address=:443"
      - "--certificatesresolvers.le.acme.httpchallenge=true"
      - "--certificatesresolvers.le.acme.httpchallenge.entrypoint=web"
      - "[email protected]"
      - "--certificatesresolvers.le.acme.storage=/certs/acme.json"
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./certs:/certs
    networks:
      - keycloak-net
    labels:
      - "traefik.http.routers.keycloak.rule=Host(`auth.votre-domaine.fr`)"
      - "traefik.http.routers.keycloak.tls.certresolver=le"
      - "traefik.http.services.keycloak.loadbalancer.server.port=8080"
```
Ajoutez ensuite deux variables à la section `environment` du service `keycloak` : `KC_PROXY_HEADERS=xforwarded` et `KC_HOSTNAME=${KC_HOSTNAME}`. Sans ces deux réglages, Keycloak continue de générer des liens en `http://` et sur le port interne 8080 même derrière un proxy HTTPS, ce qui casse silencieusement le flux de connexion pour vos utilisateurs.

## Étape 13 : Tester le flux d’authentification OIDC de bout en bout

Avant de brancher une vraie application, validez le flux directement en ligne de commande. Récupérez un jeton d’accès avec le grant `password`, pratique pour tester mais à réserver aux environnements de développement :

```
curl -X POST \
  "https://auth.votre-domaine.fr/realms/entreprise-prod/protocol/openid-connect/token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "client_id=mon-app-web" \
  -d "client_secret=le-secret-du-client" \
  -d "grant_type=password" \
  -d "username=test.utilisateur" \
  -d "password=le-mot-de-passe"
```
Une réponse réussie ressemble à ceci (jeton tronqué par souci de lisibilité) :

```
{
  "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAi...",
  "expires_in": 300,
  "refresh_expires_in": 1800,
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCIgOi...",
  "token_type": "Bearer",
  "not-before-policy": 0,
  "scope": "openid email profile"
}
```
Côté application, la validation d’un jeton se fait en vérifiant sa signature contre les clés publiques exposées par Keycloak, sans jamais interroger le serveur à chaque requête. Un exemple minimal en Python avec la bibliothèque `PyJWT` :

