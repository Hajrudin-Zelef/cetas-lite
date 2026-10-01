---
id: collect-261001-general-networking/general-networking/cyberchef-decoder-des-donnees-en-13-etapes-2026-2
title: "cyberchef-decoder-des-donnees-en-13-etapes-2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "apache"]
source: docs/RAG/collect-261001-general-networking/cyberchef-decoder-des-donnees-en-13-etapes-2026.md
source_anchor: ""
source_lines: [46, 123]
sha256: 3b07cefa35b7e3ada39b5d6ebbed5f5b3dd1e8f500ed8f2bd414f1d887ca6ac3
---

# cyberchef-decoder-des-donnees-en-13-etapes-2026

`docker run -d -p 8000:8000 --name cyberchef --restart unless-stopped mpepping/cyberchef`
Une fois le conteneur démarré, ouvrez un navigateur sur `http://localhost:8000` (ou l’adresse IP du serveur si l’installation est distante). Pour vérifier que le conteneur tourne correctement et consulter ses journaux en cas de souci :

```
docker ps --filter name=cyberchef
docker logs -f cyberchef
```
Il existe aussi des images alternatives comme **obeoneorg/cyberchef**, construites sur une base NGINX non privilégiée, qui conviennent bien à un déploiement derrière un reverse proxy d’entreprise avec TLS. Quelle que soit l’image retenue, gardez à l’esprit que le conteneur ne fait que servir l’application web statique : le calcul reste exécuté dans le navigateur du poste client, pas sur le serveur Docker lui-même.

## Étape 3 : Compiler CyberChef depuis les sources avec Node.js

Pour contribuer au projet, tester une opération personnalisée avant de la proposer en pull request, ou simplement disposer d’une version toujours à jour sans dépendre d’une image Docker tierce, la compilation depuis les sources reste la méthode recommandée par les mainteneurs eux-mêmes.

```
git clone https://github.com/gchq/CyberChef.git
cd CyberChef
npm install
npm start
```
La commande `npm start` lance un serveur de développement local (généralement accessible sur le port 8080) avec rechargement à chaud : toute modification du code source déclenche une recompilation immédiate visible dans le navigateur. Pour générer une version de production prête à être déployée sur un serveur web statique (Apache, NGINX, ou même un simple bucket S3), utilisez plutôt :

`npm run build`
Le dossier `build/prod` généré contient alors une application 100 % statique (HTML, CSS, JavaScript) qu’il suffit de copier sur n’importe quel serveur web, sans base de données ni backend applicatif à maintenir.

## Étape 4 : Décoder Base64, hexadécimal et URL encoding

Commençons par le cas le plus fréquent en analyse de logs : une chaîne encodée en Base64 planquée dans une requête HTTP suspecte ou un e-mail de phishing. Dans le panneau Operations, tapez “From Base64” dans la recherche, puis glissez-la dans la Recipe. Collez ensuite votre chaîne dans l’Input.

Exemple concret avec une charge suspecte trouvée dans un log de proxy :

```
Input : cG93ZXJzaGVsbCAtZW5jIEpBQmpBSFVBWndCbEFHNEE=
Recipe : From Base64
Output : powershell -enc JABjAHUAZwBlAG4AbwA=
```
On voit tout de suite qu’un premier décodage Base64 révèle… un second Base64 imbriqué, typique de l’obfuscation en couches employée par de nombreux droppers PowerShell. Il suffit d’empiler une deuxième opération “From Base64” dans la Recipe pour poursuivre la chaîne, sans jamais quitter l’interface. Le même principe fonctionne pour l’hexadécimal (“From Hex”), l’URL encoding (“URL Decode”) ou l’encodage HTML (“From HTML Entity”) : ce sont les briques de base que l’on combine ensuite pour des cas plus complexes.

Astuce pratique : activez l’opération “Magic” en tête de recette (voir étape 9) si vous ne savez pas quel encodage a été utilisé. Elle testera automatiquement plusieurs hypothèses et affichera un score de “lisibilité” pour chacune.

## Étape 5 : Casser un chiffrement XOR par force brute

Le XOR à clé unique reste l’une des techniques d’obfuscation les plus utilisées par les malwares pour masquer une configuration, une adresse de commande et contrôle ou une chaîne de caractères sensible dans un binaire. CyberChef propose une opération dédiée, “XOR Brute Force”, qui teste toutes les clés possibles sur un octet (0x00 à 0xFF) et affiche les résultats filtrés par un critère de “printable characters”.

- Glissez l’opération “XOR Brute Force” dans la Recipe.
- Réglez le paramètre “Crib (known plaintext string)” avec un mot que vous pensez présent dans le résultat déchiffré, par exemple “http” pour repérer une URL ou “cmd” pour une commande.
- Lancez l’analyse : CyberChef affiche uniquement les clés dont le résultat contient la chaîne recherchée, ce qui réduit immédiatement des centaines de résultats à quelques lignes exploitables.

Une fois la bonne clé identifiée grâce au brute force, remplacez l’opération par un simple “XOR” avec la clé exacte en paramètre : c’est plus rapide à exécuter et cela documente clairement, dans la recette sauvegardée, quelle clé a permis le déchiffrement. Cette démarche recoupe directement le travail habituel avec Wireshark lorsqu’on extrait un flux binaire suspect d’une capture réseau avant de le désobfusquer.

## Étape 6 : Décoder et vérifier un JSON Web Token (JWT) suspect

Les jetons JWT circulent partout dans les applications web modernes, et un jeton volé ou forgé figure souvent dans les preuves d’une compromission de session. Un JWT se compose de trois parties séparées par des points : un en-tête, une charge utile (payload) et une signature, chacune encodée en Base64 URL-safe.

```
Input : eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJhZG1pbiIsImlhdCI6MTc1ODAwMDAwMH0.4f9c3a1e2b8d7f6a5c4b3a2918f7e6d5c4b3a291
Recipe : JWT Decode
Output : {"alg":"HS256"}
{"sub":"admin","iat":1758000000}
```
L’opération “JWT Decode” affiche instantanément l’en-tête et la charge utile en clair, sans nécessiter la clé secrète (la partie payload n’est jamais chiffrée, seulement encodée et signée). C’est suffisant pour repérer un champ “sub” ou “role” anormal indiquant une escalade de privilèges. Pour aller plus loin et vérifier la validité de la signature elle-même, CyberChef propose également “JWT Verify” et “JWT Sign”, à condition de disposer de la clé ou du certificat correspondant. Ce type d’analyse s’articule naturellement avec un travail de threat intelligence structuré dans une plateforme comme MISP, où l’IOC extrait (ici, la valeur du champ “sub” compromis) peut être partagé avec d’autres équipes.

## Étape 7 : Désobfusquer un script PowerShell malveillant

Les scripts PowerShell obfusqués restent un vecteur d’attaque courant, souvent combinés à plusieurs couches d’encodage pour échapper aux antivirus signature-based. Un enchaînement typique observé dans les campagnes de 2026 combine Base64, compression Gzip et concaténation de chaînes.

Pour un payload encodé avec le paramètre `-EncodedCommand` de PowerShell (souvent abrégé `-enc`), la recette CyberChef standard est :

```
Recipe :
1. From Base64
2. Decode text (UTF-16LE)
```
PowerShell encode en effet ses commandes en UTF-16LE avant de les convertir en Base64, une subtilité qui piège souvent les analystes débutants quand le résultat du “From Base64” seul ressemble à du texte entrecoupé de caractères nuls illisibles. Si le script est en plus compressé (signature typique : les octets commencent par `H4sI` une fois en Base64), ajoutez une opération “Gunzip” ou “Raw Inflate” entre les deux étapes ci-dessus, selon l’algorithme de compression utilisé. Chaque nouvelle couche découverte s’ajoute simplement en bas de la Recipe, sans jamais retaper l’input : c’est tout l’intérêt du pipeline visuel face à une succession de commandes en ligne dans un terminal.

## Étape 8 : Extraire des indicateurs de compromission (IOC) depuis des logs bruts

Face à un fichier de logs volumineux (souvent plusieurs milliers de lignes), rechercher manuellement chaque adresse IP, domaine ou hash suspect prend un temps considérable. CyberChef propose une famille d’opérations “Extract” (Extract IP addresses, Extract domains, Extract email addresses, Extract hashes, Extract URLs) qui utilisent des expressions régulières optimisées pour repérer ces motifs automatiquement.

