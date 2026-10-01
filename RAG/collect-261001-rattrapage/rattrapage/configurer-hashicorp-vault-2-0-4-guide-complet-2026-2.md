---
id: collect-261001-rattrapage/rattrapage/configurer-hashicorp-vault-2-0-4-guide-complet-2026-2
title: "Vault v2.0.4, built with go1.23"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "aws"]
source: docs/RAG/collect-261001-rattrapage/configurer-hashicorp-vault-2-0-4-guide-complet-2026.md
source_anchor: ""
source_lines: [45, 189]
sha256: 4ecd93260c4ffaf1f0d70e26145ddb72dbf4c11b55cb8e96420fffd6dd8b54a7
---

# Vault v2.0.4, built with go1.23

Pour ce tutoriel, nous installons Vault Community 2.0.4 en local, l’option la plus courante pour découvrir l’outil avant un déploiement en production.

## Étape 2 : installer Vault sur Linux

Sur Ubuntu ou Debian, la méthode la plus simple consiste à passer par le dépôt APT officiel de HashiCorp.

```
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install vault=2.0.4-1
vault --version
# Vault v2.0.4, built with go1.23
```
Si vous préférez isoler l’installation, l’image Docker officielle fonctionne tout aussi bien et évite de polluer le système hôte.

```
docker pull hashicorp/vault:2.0.4
docker run -d --name vault-dev \
  --cap-add=IPC_LOCK \
  -p 8200:8200 \
  -e 'VAULT_DEV_ROOT_TOKEN_ID=root-token-dev' \
  -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200' \
  hashicorp/vault:2.0.4
```
Le mode dev démarre un serveur Vault en mémoire, non chiffré au repos, avec auto-unseal désactivé et un jeton racine prédéfini. C’est parfait pour apprendre les commandes, mais il ne faut jamais l’utiliser en production : toutes les données disparaissent au redémarrage du conteneur.

## Étape 3 : initialiser un serveur Vault en mode production

Pour une installation persistante, créez un fichier de configuration qui définit le backend de stockage (Raft intégré, sans dépendance à Consul depuis plusieurs versions), le listener TLS et l’API d’écoute.

```
storage "raft" {
  path    = "/opt/vault/data"
  node_id = "vault-node-1"
}
listener "tcp" {
  address       = "0.0.0.0:8200"
  tls_cert_file = "/opt/vault/tls/vault.crt"
  tls_key_file  = "/opt/vault/tls/vault.key"
}
api_addr     = "https://vault.entreprise.fr:8200"
cluster_addr = "https://vault.entreprise.fr:8201"
ui           = true
```
Lancez le serveur avec ce fichier, puis initialisez le coffre-fort. Cette étape génère les clés de descellement (unseal keys) et le jeton racine : conservez-les dans un gestionnaire de mots de passe séparé, jamais dans le même dépôt que votre code.

```
vault server -config=/etc/vault.d/vault.hcl &
export VAULT_ADDR='https://vault.entreprise.fr:8200'
vault operator init -key-shares=5 -key-threshold=3
# Unseal Key 1: 8a3f...e91c
# Unseal Key 2: 7b1d...f402
# Unseal Key 3: 4c9e...a880
# Unseal Key 4: 1f2a...cd77
# Unseal Key 5: 90bb...1234
# Initial Root Token: hvs.AbCdEf1234567890
```
Le schéma de Shamir (5 clés, seuil de 3) répartit la responsabilité du descellement entre plusieurs personnes : aucune ne peut ouvrir le coffre seule. En production, on lui préfère souvent l’auto-unseal, qui délègue le chiffrement de la clé racine à un service cloud comme AWS KMS ou Azure Key Vault, pour éviter d’avoir à taper trois clés manuellement à chaque redémarrage.

## Étape 4 : configurer l’auto-unseal avec AWS KMS

Ajoutez un bloc `seal` dans votre configuration pour déléguer le descellement à une clé KMS gérée dans le cloud.

```
seal "awskms" {
  region     = "eu-west-3"
  kms_key_id = "arn:aws:kms:eu-west-3:123456789012:key/abcd-1234-efgh-5678"
}
```
Notez le choix de la région `eu-west-3` (Paris) : pour rester conforme aux exigences de résidence des données de nombreuses entreprises françaises, gardez la clé KMS et le cluster Vault dans la même région européenne. Une fois ce bloc en place, redémarrez le serveur : le descellement devient automatique au démarrage, sans intervention humaine, tout en conservant le chiffrement de la clé racine côté cloud.

## Étape 5 : activer le moteur de secrets KV version 2

Le moteur KV (key-value) version 2 est le point d’entrée le plus courant : il stocke des paires clé-valeur versionnées, avec un historique complet des modifications.

```
vault secrets enable -path=secret kv-v2
vault kv put secret/production/base-donnees \
  utilisateur="app_prod" \
  mot_de_passe="Xk9#mP2vL8qR"
vault kv get secret/production/base-donnees
# ====== Metadata ======
# Key         Value
# ---         -----
# version     1
#
# ====== Data ======
# Key             Value
# ---             -----
# mot_de_passe    Xk9#mP2vL8qR
# utilisateur     app_prod
```
Chaque écriture crée une nouvelle version. Vous pouvez consulter l’historique avec `vault kv get -version=1 secret/production/base-donnees`, ou restaurer une ancienne version en cas d’erreur de configuration. C’est déjà un progrès énorme par rapport à un fichier `.env` versionné dans Git, où toute l’historique des secrets reste visible pour quiconque a accès au dépôt.

## Étape 6 : passer aux secrets dynamiques pour vos bases de données

Le KV statique reste un secret qui ne change pas tout seul. L’intérêt majeur de Vault se révèle avec les secrets dynamiques : au lieu de stocker un mot de passe fixe, Vault génère un identifiant de base de données temporaire, avec un bail (lease) et une expiration automatique.

```
vault secrets enable database
vault write database/config/postgres-prod \
  plugin_name=postgresql-database-plugin \
  connection_url="postgresql://{{username}}:{{password}}@db.entreprise.fr:5432/prod" \
  allowed_roles="role-lecture-seule" \
  username="vault_admin" \
  password="mot-de-passe-admin-initial"
vault write database/roles/role-lecture-seule \
  db_name=postgres-prod \
  creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
  default_ttl="1h" \
  max_ttl="24h"
vault read database/creds/role-lecture-seule
# Key                Value
# ---                -----
# lease_id           database/creds/role-lecture-seule/8f3a...
# lease_duration     1h
# username            v-token-role-lect-a1b2c3
# password            A1b2C3d4E5f6!
```
Chaque appel de cette commande crée un nouveau compte PostgreSQL, avec un mot de passe unique, une durée de vie d’une heure maximum, et une révocation automatique à l’expiration du bail. Un attaquant qui intercepterait ce mot de passe ne pourrait s’en servir que pendant une fenêtre très réduite.

## Étape 7 : mettre en place le moteur PKI pour vos certificats internes

Le moteur PKI transforme Vault en autorité de certification interne, capable d’émettre des certificats TLS à courte durée de vie pour vos microservices.

```
vault secrets enable pki
vault secrets tune -max-lease-ttl=87600h pki
vault write pki/root/generate/internal \
  common_name="entreprise.fr" \
  ttl=87600h
vault write pki/roles/service-interne \
  allowed_domains="entreprise.fr" \
  allow_subdomains=true \
  max_ttl="720h"
vault write pki/issue/service-interne \
  common_name="api.entreprise.fr" \
  ttl="24h"
```
Chaque service peut ainsi demander un certificat valable 24 heures, renouvelé automatiquement par un agent, plutôt que de gérer manuellement des certificats valides plusieurs années qui deviennent des cibles de choix en cas de compromission.

## Étape 8 : chiffrer des données applicatives avec le moteur Transit

Le moteur Transit permet à vos applications de chiffrer et déchiffrer des données sans jamais manipuler directement la clé de chiffrement, qui reste dans Vault (modèle « encryption as a service »).

