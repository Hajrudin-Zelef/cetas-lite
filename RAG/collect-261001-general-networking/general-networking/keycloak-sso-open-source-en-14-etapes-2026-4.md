---
id: collect-261001-general-networking/general-networking/keycloak-sso-open-source-en-14-etapes-2026-4
title: "Base de données"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "apache", "decode", "incident", "mcp", "model context protocol", "open source"]
source: docs/RAG/collect-261001-general-networking/keycloak-sso-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [266, 349]
sha256: 7ed02d883e0cbf2b7c20af5cffa96fa71dd5ed65e9591499b61b7ddd546be332
---

# Base de données

```
import jwt
from jwt import PyJWKClient
realm_url = "https://auth.votre-domaine.fr/realms/entreprise-prod"
jwks_client = PyJWKClient(f"{realm_url}/protocol/openid-connect/certs")
def verifier_jeton(token: str):
    signing_key = jwks_client.get_signing_key_from_jwt(token)
    return jwt.decode(
        token,
        signing_key.key,
        algorithms=["RS256"],
        audience="mon-app-web",
        issuer=realm_url,
    )
```
Si `verifier_jeton` renvoie le contenu décodé sans lever d’exception, votre flux OIDC fonctionne de bout en bout, de la connexion de l’utilisateur jusqu’à la validation côté application.

## Étape 14 : Sauvegarder, exporter et passer en production

Une fois la configuration validée, exportez le royaume complet pour pouvoir le reconstruire ailleurs de façon identique, un principe d’infrastructure as code appliqué à l’IAM :

```
docker exec -it keycloak /opt/keycloak/bin/kc.sh export \
  --dir /opt/keycloak/data/export \
  --realm entreprise-prod \
  --users realm_file
```
Programmez aussi une sauvegarde régulière de la base PostgreSQL, la vraie source de vérité de votre déploiement :

```
docker exec keycloak-postgres pg_dump -U keycloak keycloak \
  | gzip > backup-keycloak-$(date +%F).sql.gz
```
Pour la mise en production, remplacez la commande `start-dev` par un vrai build optimisé. Créez un Dockerfile minimal qui exécute `kc build` avec les fonctionnalités dont vous avez réellement besoin, puis démarrez avec `start --optimized` plutôt qu’en mode développement, qui désactive plusieurs protections et expose des endpoints de debug à ne jamais laisser accessibles publiquement.

## Aller plus loin : protéger une API avec l’introspection de jeton

Une fois le SSO en place pour vos interfaces web, la question suivante arrive presque toujours : comment un microservice backend vérifie-t-il qu’un jeton présenté par un appelant est encore valide ? Deux approches coexistent. La première, déjà couverte à l’étape 13, valide la signature du jeton localement via les clés publiques JWKS, sans appeler Keycloak à chaque requête : c’est la méthode la plus rapide et la plus scalable. La seconde, l’introspection, interroge Keycloak en direct pour chaque jeton, ce qui coûte une requête réseau supplémentaire mais permet de révoquer un accès immédiatement, même avant l’expiration naturelle du jeton.

L’introspection convient particulièrement aux API qui manipulent des données sensibles, où le délai entre une révocation d’accès et sa prise en compte doit être quasi nul. Depuis janvier 2026, Keycloak documente d’ailleurs officiellement ce même mécanisme pour un cas d’usage émergent : servir de serveur d’autorisation pour des serveurs MCP (Model Context Protocol), la même logique de jetons et de scopes s’appliquant désormais aux agents IA qui appellent vos API. Voici l’appel type, à exécuter depuis votre service backend plutôt que depuis le navigateur :

```
curl -X POST \
  "https://auth.votre-domaine.fr/realms/entreprise-prod/protocol/openid-connect/token/introspect" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -u "mon-app-api:le-secret-du-client-api" \
  -d "token=eyJhbGciOiJSUzI1NiIsInR5cCIgOiAi..."
```
Keycloak répond avec un objet JSON qui indique si le jeton est actif, ses scopes et son propriétaire :

```
{
  "active": true,
  "scope": "openid email profile",
  "client_id": "mon-app-api",
  "username": "test.utilisateur",
  "exp": 1752345600,
  "realm_access": {
    "roles": ["utilisateur"]
  }
}
```
Si `active` vaut `false`, rejetez la requête avec un code 401, quelle que soit la raison (jeton expiré, révoqué ou signé par un autre royaume). Pour une API à fort trafic, réservez l’introspection aux endpoints les plus sensibles et gardez la validation locale JWKS pour le reste : le compromis entre latence et réactivité de révocation se règle service par service, pas au niveau de toute l’architecture.

## Keycloak face à Auth0, Okta et Authentik : le comparatif des coûts

La question du prix revient systématiquement dès qu’une équipe évalue Keycloak face à une alternative SaaS. Les chiffres ci-dessous correspondent aux tarifs publics affichés par les éditeurs, en dollars, à titre indicatif : ils évoluent régulièrement et méritent d’être vérifiés au moment de votre décision.

| Solution | Licence | Hébergement | Tarif indicatif | 
|---|---|---|---|
| Keycloak | Apache 2.0 (open source) | Auto-hébergé | Gratuit, coût d’infrastructure uniquement | 
| Red Hat build of Keycloak | Apache 2.0 + support commercial | Auto-hébergé | Abonnement de support sur devis | 
| Auth0 (Okta) | Propriétaire | SaaS | Gratuit en développement, puis environ 25 à 50 $ par utilisateur actif et par mois en entreprise | 
| Okta Identity Engine | Propriétaire | SaaS | Environ 6 à 10 $ par utilisateur et par mois, jusqu’à 20 $ ou plus pour les options avancées | 
| Authentik | Open source, ou SaaS géré | Auto-hébergé ou SaaS | Gratuit en auto-hébergement, environ 15 $ par utilisateur et par mois en SaaS | 

Le calcul penche presque toujours en faveur de Keycloak dès que le nombre de comptes dépasse quelques centaines, à condition d’avoir en interne les compétences pour l’exploiter. C’est justement le compromis : Auth0 et Okta vendent de la tranquillité d’esprit et un support contractuel, Keycloak vend la maîtrise complète de la pile au prix du temps d’administration système. Une équipe de deux ou trois développeurs sans administrateur système dédié mettra probablement plus de temps à durcir correctement son instance Keycloak que ne coûterait un abonnement Auth0 pendant la même période, ce qui remet le calcul en perspective : le logiciel est gratuit, l’exploitation ne l’est jamais complètement.

Authentik occupe une position intermédiaire intéressante pour les équipes qui veulent l’esprit open source de Keycloak avec une interface jugée plus moderne par certains administrateurs, au prix d’un écosystème et d’une communauté nettement plus restreints. Le choix final dépend surtout de la taille de votre parc applicatif et de votre tolérance à opérer vous-même un service critique.

## Les erreurs les plus fréquentes lors du déploiement de Keycloak

Sept erreurs reviennent systématiquement dans les déploiements Keycloak, y compris chez des équipes expérimentées qui découvrent l’outil. Aucune n’est propre à Keycloak en particulier, ce sont globalement les mêmes pièges que pour n’importe quel service critique déployé en conteneur, mais leurs conséquences sont plus visibles ici puisqu’un incident touche immédiatement la capacité de tous vos utilisateurs à se connecter.

