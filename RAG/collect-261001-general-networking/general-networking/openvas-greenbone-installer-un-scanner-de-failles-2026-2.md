---
id: collect-261001-general-networking/general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026-2
title: "Vérifier l'espace disque disponible (minimum 20 Go recommandé)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit", "mai"]
source: docs/RAG/collect-261001-general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026.md
source_anchor: ""
source_lines: [39, 170]
sha256: d2e551a595f1dccc761bacafb71e70cac38a00f06db9768987884f555727d956
---

# Vérifier l'espace disque disponible (minimum 20 Go recommandé)

Si vous testez d’abord sur une machine virtuelle avant un déploiement en production, prévoyez un instantané (snapshot) juste avant l’installation. Docker isole bien les services entre eux, mais une mauvaise manipulation du fichier `docker-compose.yml` ou une coupure réseau pendant la synchronisation initiale du feed peut laisser la base de données dans un état incohérent qu’il est plus rapide de restaurer que de déboguer.

## Étape 1 : préparer le serveur Debian ou Ubuntu

Commencez par mettre à jour le système et installer les paquets de base nécessaires à Docker et à la gestion du pare-feu.

```
sudo apt update && sudo apt upgrade -y
sudo apt install -y curl ca-certificates gnupg lsb-release ufw
# Vérifier l'espace disque disponible (minimum 20 Go recommandé)
df -h /
# Vérifier la RAM disponible
free -h
```
Configurez ensuite un utilisateur dédié plutôt que de tout exécuter en root, par principe de moindre privilège :

```
sudo adduser gvmadmin
sudo usermod -aG sudo gvmadmin
su - gvmadmin
```
## Étape 2 : installer Docker Engine et Docker Compose

Greenbone Community Edition se déploie officiellement via Docker Compose depuis 2023, ce qui simplifie énormément la gestion des multiples services (scanner, base de données, feed, interface web). Un guide comparatif publié par Secrails en mai 2026 confirme d’ailleurs que les Greenbone Community Containers restent la voie d’installation la plus propre, alors que les installations via apt sur Ubuntu 22.04 continuent de se heurter à des conflits de dépendances, Kali Linux demeurant de son côté l’option bare-metal la moins friction. Ajoutez le dépôt officiel Docker :

```
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
# Ajouter votre utilisateur au groupe docker
sudo usermod -aG docker $USER
newgrp docker
# Vérifier l'installation
docker --version
docker compose version
```
Sortie attendue à ce stade :

```
Docker version 27.3.1, build ce12230
Docker Compose version v2.30.3
```
## Étape 3 : télécharger et lancer Greenbone Community Container

Greenbone publie un dépôt Docker Compose officiel qui orchestre l’ensemble des microservices : gvmd (le démon de gestion), le scanner OpenVAS, la base PostgreSQL, le service de mise à jour du feed NVT et l’interface GSA.

```
mkdir -p ~/greenbone-ce && cd ~/greenbone-ce
curl -f -L https://greenbone.github.io/docs/latest/_static/docker-compose-22.4.yml -o docker-compose.yml
curl -f -L https://greenbone.github.io/docs/latest/_static/gvm.conf -o gvm.conf
# Lancer les conteneurs en arrière-plan
docker compose up -d
# Vérifier que tous les services tournent
docker compose ps
```
Le premier démarrage déclenche automatiquement la synchronisation du feed de vulnérabilités : plus de 160 000 tests NVT, les données CERT, les correspondances CVE et les scripts CVSS. Cette synchronisation initiale peut prendre entre 30 minutes et plusieurs heures selon votre bande passante. Suivez la progression avec :

`docker compose logs -f pg-gvm gvmd`
## Étape 4 : récupérer le mot de passe administrateur généré

À la création du conteneur, Greenbone génère automatiquement un compte `admin` avec un mot de passe aléatoire, visible dans les logs du conteneur `gvmd` :

```
docker compose logs gvmd | grep -i "user created"
# Exemple de sortie :
# gvmd: User created with password 'K7x-mP2q-Rv9t-Zb4w'
```
Changez immédiatement ce mot de passe une fois connecté à l’interface web. Vous pouvez aussi le réinitialiser en ligne de commande si les logs ne l’affichent plus :

`docker compose exec gvmd gvmd --user=admin --new-password='VotreMotDePasseFort123!'`
## Étape 5 : accéder à l’interface Greenbone Security Assistant (GSA)

Ouvrez un navigateur et rendez-vous sur `https://IP_DU_SERVEUR:9392`. Le certificat TLS auto-signé par défaut déclenchera un avertissement de sécurité, normal pour un premier accès local. Connectez-vous avec le compte `admin` et le mot de passe récupéré à l’étape précédente.

Dans le tableau de bord, vérifiez trois indicateurs avant de lancer le moindre scan : le statut du feed NVT (menu Administration → Feed Status), qui doit afficher « Current », la date de dernière synchronisation, et le nombre de tests chargés (autour de 160 000 en cas de feed communautaire complet).

## Étape 6 : configurer un premier scan cible

Créez d’abord une cible (Target) dans le menu Configuration → Targets. Renseignez une plage IP restreinte pour ce premier test, par exemple un seul serveur non critique :

- Nom : `Test-Serveur-Web-01`
- Hosts : `192.168.1.50`
- Méthode de découverte de port : `OpenVAS Default`
- Alive Test : `ICMP, TCP-ACK Service Ping, ARP Ping`

Rendez-vous ensuite dans Scans → Tasks, cliquez sur « New Task », associez la cible créée et sélectionnez une configuration de scan. Pour un premier passage, privilégiez le profil `Full and fast`, qui équilibre couverture et durée sans saturer la cible.

## Étape 7 : automatiser les scans via l’API GMP (Greenbone Management Protocol)

Pour intégrer OpenVAS dans un pipeline d’automatisation ou un cron, utilisez le protocole GMP, disponible en version 22.8 et 22.7 dans la documentation Greenbone mise à jour le 12 août 2026. Le client Python `python-gvm` reste la méthode la plus fiable pour scripter des scans récurrents.

```
pip install python-gvm
python3 - <<'EOF'
from gvm.connections import TLSConnection
from gvm.protocols.gmp import Gmp
from gvm.transforms import EtreeCheckCommandTransform
connection = TLSConnection(hostname="192.168.1.10", port=9390)
transform = EtreeCheckCommandTransform()
with Gmp(connection=connection, transform=transform) as gmp:
    gmp.authenticate("admin", "VotreMotDePasseFort123!")
    version = gmp.get_version()
    print(version)
    # Lancer une tâche existante par son UUID
    gmp.start_task("11111111-2222-3333-4444-555555555555")
EOF
```
Planifiez ensuite un cron hebdomadaire qui déclenche ce script chaque lundi à 2h du matin, hors heures de bureau, pour limiter l'impact sur la production :

```
crontab -e
# Ajouter la ligne suivante :
0 2 * * 1 /usr/bin/python3 /home/gvmadmin/scripts/scan_hebdo.py >> /var/log/openvas_scan.log 2>&1
```
## Étape 8 : lire et prioriser un rapport de vulnérabilités

Une fois le scan terminé, Greenbone classe chaque résultat selon le score CVSS v3.1, calculé via la même méthodologie que le calculateur officiel du FIRST. Ne traitez pas toutes les failles de la même façon : concentrez-vous d'abord sur les scores critiques et sur les vulnérabilités exploitées activement.

| Score CVSS | Niveau | Délai de correction recommandé | 
|---|---|---|
| 9.0 – 10.0 | Critique | 24 à 72 heures | 
| 7.0 – 8.9 | Élevé | 7 jours | 
| 4.0 – 6.9 | Moyen | 30 jours | 
| 0.1 – 3.9 | Faible | 90 jours ou cycle de patch normal | 

Croisez systématiquement les identifiants CVE remontés avec la base officielle du CVE Program pour confirmer la description exacte de la faille et vérifier si un correctif existe déjà côté éditeur.

## Étape 9 : exporter les rapports en PDF et CSV pour l'audit NIS2

Dans le menu Scans → Reports, sélectionnez un rapport terminé puis « Export » pour choisir un format. Pour un dossier de conformité NIS2 ou un rapport à destination d'un RSSI, le format PDF avec le modèle « Executive Summary » reste le plus lisible. Pour une intégration dans un tableau de bord ou un outil tiers, exportez en CSV ou en XML :

