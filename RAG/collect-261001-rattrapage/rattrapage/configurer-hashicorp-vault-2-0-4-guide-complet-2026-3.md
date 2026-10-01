---
id: collect-261001-rattrapage/rattrapage/configurer-hashicorp-vault-2-0-4-guide-complet-2026-3
title: "Vault v2.0.4, built with go1.23"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution", "open source"]
source: docs/RAG/collect-261001-rattrapage/configurer-hashicorp-vault-2-0-4-guide-complet-2026.md
source_anchor: ""
source_lines: [190, 312]
sha256: dc353e9d1b49884cb2e980005463f9714bd3fb4a2be1b3e28134d748441e63e2
---

# Vault v2.0.4, built with go1.23

```
vault secrets enable transit
vault write -f transit/keys/donnees-clients
echo -n "numero-carte-1234567890123456" | base64
vault write transit/encrypt/donnees-clients \
  plaintext="bnVtZXJvLWNhcnRlLTEyMzQ1Njc4OTAxMjM0NTY="
# ciphertext: vault:v1:8f3a1b2c...
vault write transit/decrypt/donnees-clients \
  ciphertext="vault:v1:8f3a1b2c..."
```
Cette approche est particulièrement utile pour le chiffrement au niveau applicatif de données personnelles sensibles au sens du RGPD (numéros de carte bancaire, données de santé, identifiants). La clé de chiffrement ne quitte jamais Vault, ce qui simplifie considérablement les audits de sécurité et la démonstration de conformité.

## Étape 9 : écrire des politiques d’accès (policies) précises

Sans politiques, tout utilisateur authentifié aurait potentiellement accès à tous les secrets. Les policies Vault utilisent une syntaxe HCL déclarative pour limiter les droits chemin par chemin.

```
path "secret/data/production/*" {
  capabilities = ["read", "list"]
}
path "database/creds/role-lecture-seule" {
  capabilities = ["read"]
}
path "secret/data/dev/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
```
```
vault policy write equipe-backend backend-policy.hcl
vault token create -policy="equipe-backend" -ttl="8h"
```
Le principe du moindre privilège s’applique ici de manière très granulaire : un développeur backend obtient un accès en lecture seule sur les secrets de production, mais un accès complet en environnement de développement. Le jeton généré expire après 8 heures, ce qui limite la fenêtre d’exploitation si le jeton venait à fuiter.

## Étape 10 : intégrer Vault à Kubernetes avec vault-k8s

Pour les équipes qui déploient sur Kubernetes, l’injecteur officiel vault-k8s (en version 1.7.6 depuis le 6 août 2026) permet d’injecter des secrets directement dans les pods via un sidecar, sans jamais les faire transiter par un ConfigMap ou un Secret Kubernetes natif en clair.

```
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo update
helm install vault hashicorp/vault \
  --set "injector.enabled=true" \
  --set "injector.image.tag=1.7.6" \
  --namespace vault-system --create-namespace
```
Une fois l’injecteur déployé, il suffit d’ajouter des annotations sur vos manifests de déploiement pour que Vault s’occupe automatiquement d’injecter les secrets au démarrage du pod.

```
annotations:
  vault.hashicorp.com/agent-inject: "true"
  vault.hashicorp.com/role: "api-backend"
  vault.hashicorp.com/agent-inject-secret-db-creds: "database/creds/role-lecture-seule"
```
L’authentification Kubernetes native (via le token de compte de service) évite d’avoir à distribuer un jeton Vault statique dans chaque pod : c’est l’identité du pod lui-même qui sert de preuve d’authentification.

## Étape 11 : automatiser le provisioning avec Terraform

Le flux DevOps recommandé en 2026 sépare clairement les responsabilités : Terraform pour provisionner l’infrastructure Vault elle-même (moteurs, policies, rôles), Vault pour la distribution dynamique des secrets à l’exécution.

```
terraform {
  required_providers {
    vault = {
      source  = "hashicorp/vault"
      version = "~> 4.0"
    }
  }
}
provider "vault" {
  address = "https://vault.entreprise.fr:8200"
}
resource "vault_mount" "kv" {
  path = "secret"
  type = "kv-v2"
}
resource "vault_policy" "backend" {
  name   = "equipe-backend"
  policy = file("${path.module}/backend-policy.hcl")
}
```
Cette approche « tout en code » permet de versionner vos politiques d’accès au même titre que votre infrastructure, avec revue de code obligatoire avant toute modification des droits d’accès aux secrets de production.

## Étape 12 : mettre en place la rotation et l’audit des accès

Activez systématiquement les logs d’audit dès le déploiement, avant même de créer vos premiers secrets. Vault enregistre chaque requête d’accès, chaque lecture, chaque création de jeton.

```
vault audit enable file file_path=/var/log/vault/audit.log
vault audit list
# Path     Type    Description
# ----     ----    -----------
# file/    file     n/a
```
Renvoyez ensuite ces logs vers votre SIEM (nous avions détaillé la mise en place de Wazuh en tant que SIEM open source dans un précédent tutoriel) pour détecter les schémas d’accès anormaux : un jeton qui lit soudainement des centaines de secrets en quelques secondes, ou un accès depuis une plage IP inhabituelle, sont des signaux à investiguer immédiatement.

## Étape 13 : appliquer les correctifs de sécurité récents

La discipline de patching n’est pas optionnelle. Le 10 août 2026, HashiCorp a publié l’avis de sécurité HCSEC-2026-26, décrivant une faille de contournement d’autorisation LIST via une suppression de barre oblique finale (trailing-slash strip), qui affectait Vault et Vault Enterprise jusqu’à la version 2.0.2. Le correctif est arrivé dans la 2.0.3, avec des rétroportages sur les branches 1.21.8, 1.20.13 et 1.19.19, avant que la 2.0.4 ne consolide l’ensemble début août 2026.

```
# Vérifier la version installée
vault --version
# Comparer avec la dernière version publiée
curl -s https://api.releases.hashicorp.com/v1/releases/vault/latest | grep version
```
Consultez régulièrement les notes de version officielles et abonnez-vous au flux d’annonces de sécurité de HashiCorp. Un coffre-fort de secrets non patché devient lui-même une cible de choix : c’est le paradoxe d’un outil censé réduire la surface d’attaque.

## 5 pièges courants lors de la mise en place de Vault

**1. Laisser le mode dev tourner en production.** Le serveur de développement stocke tout en mémoire et désactive le TLS. Il est tentant de le garder « juste le temps de finir le sprint », mais chaque redémarrage efface l’ensemble des secrets, et l’absence de chiffrement expose le trafic en clair sur le réseau.

**2. Laisser le jeton racine dans un fichier accessible.** Le jeton racine généré à l’initialisation donne un accès total et permanent. Il doit être révoqué (`vault token revoke`) juste après la configuration initiale des politiques, puis remplacé par des jetons à privilèges limités et à durée de vie courte.

**3. Oublier de configurer le TTL maximal des baux.** Un secret dynamique sans `max_ttl` explicite peut être renouvelé indéfiniment par une application mal configurée, ce qui annule l’intérêt de la rotation automatique. Fixez toujours une durée de vie maximale raisonnable, généralement entre 24 et 72 heures pour des identifiants de base de données.

**4. Confondre Ansible Vault et HashiCorp Vault.** Ansible propose son propre mécanisme appelé « Ansible Vault » pour chiffrer des fichiers de variables, un outil totalement distinct du produit HashiCorp Vault présenté ici (nous avions couvert Ansible et son Vault intégré dans un autre tutoriel). Les deux peuvent d’ailleurs coexister : Ansible Vault pour chiffrer des fichiers statiques, HashiCorp Vault pour la distribution dynamique de secrets à l’exécution.

**5. Négliger le coût réel de HCP Vault Dedicated.** Le tarif d’appel à 450 dollars par mois pour un cluster de développement extra-small paraît abordable, mais un cluster « Essentials Small » de production démarre autour de 1 150 dollars par mois, auxquels s’ajoutent des frais par client de 72,92 dollars mensuels sur les paliers Essentials et Standard. Un cluster standard avec 50 clients atteint facilement 5 000 dollars par mois. Simulez le coût total avant de vous engager sur l’offre managée.

## Vault face à la concurrence : comparatif 2026

Le marché de la gestion de secrets s’est diversifié ces deux dernières années, notamment avec la montée d’alternatives open source plus légères à opérer que Vault.

