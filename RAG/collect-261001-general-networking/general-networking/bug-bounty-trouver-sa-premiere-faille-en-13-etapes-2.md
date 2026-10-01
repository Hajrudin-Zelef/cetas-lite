---
id: collect-261001-general-networking/general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes-2
title: "Installer Go si nécessaire"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes.md
source_anchor: ""
source_lines: [39, 107]
sha256: 7f6828da8b88049ba3769cc18d460b832edf252b0c00e903ae05fd2de4e2e399
---

# Installer Go si nécessaire

Le choix de la plateforme conditionne le type de programmes accessibles, la langue de support et les modalités de paiement. Pour un débutant francophone, YesWeHack présente l’avantage d’une communauté et d’un support en français, ainsi que d’un accès à des programmes publics français (administrations, collectivités) qui imposent des exigences de conformité claires et un scope souvent bien documenté. HackerOne et Bugcrowd offrent un volume de programmes bien plus large à l’échelle mondiale, avec une compétition proportionnellement plus forte. Intigriti, basée en Belgique, se positionne comme une alternative européenne avec une forte présence de marques du e-commerce et de la fintech.

La recommandation pratique : créer un compte sur au moins deux plateformes dès le départ, mais concentrer son énergie sur une seule au début. Passer d’un programme à l’autre sans approfondir dilue l’apprentissage. Une fois le compte créé, il faut généralement passer par une phase de vérification d’identité (KYC) avant de pouvoir recevoir des paiements, ce qui peut prendre plusieurs jours. Autant l’anticiper avant d’avoir trouvé sa première faille.

## Étape 2 : Sélectionner son premier programme

Toutes les plateformes distinguent les programmes publics (ouverts à tous les chercheurs inscrits) des programmes privés (accessibles sur invitation, généralement après avoir fait ses preuves). Pour débuter, il faut filtrer les programmes publics selon trois critères : un scope large incluant des applications web classiques, une table de récompenses publiée (les programmes “reconnaissance uniquement”, sans paiement, sont utiles pour le portfolio mais ne doivent pas être la priorité), et un historique de réponse rapide de l’entreprise, visible dans les statistiques du programme (temps de première réponse, temps de résolution).

Il faut absolument lire la politique du programme dans son intégralité avant de commencer, en particulier trois sections : le scope exact (domaines, sous-domaines, applications mobiles inclus ou exclus), les types de tests interdits (déni de service, ingénierie sociale, spam, tests automatisés agressifs), et les catégories de vulnérabilités explicitement hors périmètre (souvent le clickjacking sur des pages sans action sensible, l’absence de rate limiting sur des endpoints non critiques, ou les en-têtes de sécurité manquants sans impact démontré). Soumettre un rapport sur un élément explicitement exclu est la première cause de rejet chez les débutants.

## Étape 3 : Installer la boîte à outils de reconnaissance

La suite ProjectDiscovery (Subfinder, httpx, Nuclei, dnsx) constitue la base de reconnaissance la plus utilisée par la communauté en 2026, en complément de Burp Suite pour l’analyse manuelle. Voici l’installation via Go, qui garantit d’obtenir les dernières versions directement depuis les dépôts officiels.

```
# Installer Go si nécessaire
sudo apt update && sudo apt install golang-go -y
# Installer les outils ProjectDiscovery
go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest
go install -v github.com/projectdiscovery/dnsx/cmd/dnsx@latest
# Ajouter le binPath Go au PATH
export PATH=$PATH:$(go env GOPATH)/bin
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
# Vérifier les installations
subfinder -version
httpx -version
nuclei -version
```
Nuclei nécessite ensuite une mise à jour de ses modèles de détection, qui sont maintenus par la communauté et mis à jour quasi quotidiennement pour intégrer les nouvelles CVE et les nouveaux patterns de mauvaise configuration.

```
nuclei -update-templates
nuclei -update
```
Pour Burp Suite, la version Community Edition, gratuite, suffit largement à débuter : elle inclut le proxy d’interception, Repeater pour rejouer des requêtes modifiées, et Intruder en version limitée. La version Professional de PortSwigger ajoute le scanner actif automatisé et Intruder sans restriction de débit, un investissement à envisager une fois les premiers rapports validés et les premières primes encaissées, pas avant.

## Étape 4 : Cartographier la surface d’attaque (reconnaissance passive)

La reconnaissance passive consiste à identifier tous les sous-domaines et actifs numériques d’une cible sans envoyer la moindre requête directe vers ses serveurs, en s’appuyant uniquement sur des sources publiques (certificats TLS, journaux de transparence, moteurs de recherche spécialisés). Cette étape est cruciale car un sous-domaine oublié, souvent un environnement de préproduction ou de démonstration mal désindexé, concentre une part disproportionnée des vulnérabilités trouvées par les débutants.

```
# Énumération passive des sous-domaines
subfinder -d exemple-cible.com -all -recursive -o subdomains.txt
# Résolution DNS et filtrage des domaines actifs
cat subdomains.txt | dnsx -silent -o resolved.txt
# Vérification HTTP/HTTPS avec titres de page et technologies
cat resolved.txt | httpx -silent -title -tech-detect -status-code -o live_hosts.txt
cat live_hosts.txt
```
Le fichier `live_hosts.txt` obtenu constitue la carte de la surface d’attaque. Il faut le confronter méthodiquement au scope publié du programme : tout sous-domaine hors périmètre doit être écarté, même s’il paraît vulnérable. Tester en dehors du scope, même par curiosité, viole les conditions du programme et peut entraîner une exclusion définitive de la plateforme.

## Étape 5 : Scanner les vulnérabilités connues avec Nuclei

Une fois la liste des hôtes actifs et dans le scope établie, Nuclei permet de détecter rapidement les mauvaises configurations courantes et les vulnérabilités déjà documentées dans ses modèles communautaires : panneaux d’administration exposés, fichiers de configuration accessibles, versions de logiciels obsolètes avec CVE connue, absence d’en-têtes de sécurité critiques.

```
# Scan avec limitation du débit pour respecter les règles du programme
nuclei -l live_hosts.txt -severity medium,high,critical -rate-limit 10 -o nuclei_results.txt
# Scan ciblé sur les expositions de fichiers sensibles
nuclei -l live_hosts.txt -tags exposure,misconfig -rate-limit 10 -o exposures.txt
```
Le paramètre `-rate-limit` n’est pas optionnel. La quasi-totalité des programmes imposent une limite de requêtes par seconde, et un scan trop agressif peut être interprété comme une tentative de déni de service, motif d’exclusion immédiate. Un résultat positif de Nuclei n’est jamais une preuve suffisante : il s’agit d’un indice à valider manuellement avant toute soumission, car le taux de faux positifs sur les scans automatisés reste significatif.

## Sécuriser son propre environnement de test

Un point souvent oublié par les débutants concerne la sécurité de leur propre poste de travail. Un chercheur qui teste des applications potentiellement compromises ou qui télécharge des outils communautaires depuis des dépôts non vérifiés s’expose lui-même à des risques : payloads malveillants déguisés en outils de recon, extensions de navigateur compromises, ou scripts partagés sur des forums qui embarquent des portes dérobées. La pratique recommandée consiste à isoler l’environnement de test dans une machine virtuelle dédiée, sans accès aux identifiants personnels ni aux données sensibles du chercheur, et à ne jamais exécuter un outil téléchargé sans en avoir vérifié la provenance sur le dépôt officiel du projet.

