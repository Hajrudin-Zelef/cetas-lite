---
id: collect-261001-general-networking/general-networking/geforce-now-sur-tv-shield-lg-samsung-en-13-etapes-2026-2
title: "Vérifier la version Fire OS installée en SSH/ADB (si le débogage est activé)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Nvidia", "Samsung"]
dates: []
keywords: ["attention", "datacenter", "ethernet", "latency", "nvidia"]
source: docs/RAG/collect-261001-general-networking/geforce-now-sur-tv-shield-lg-samsung-en-13-etapes-2026.md
source_anchor: ""
source_lines: [46, 135]
sha256: 3a5838e74a50fbd353e2c1703eec4d9edaf0107172b3574844246e006f0ee32e
---

# Vérifier la version Fire OS installée en SSH/ADB (si le débogage est activé)

**Étape 6 — Fire TV Stick.** Sur un Fire TV Stick 4K Plus (2e génération) ou 4K Max compatible, ouvrez l’Amazon Appstore, recherchez GeForce NOW. Si l’app n’apparaît pas (certains modèles plus anciens ne sont pas listés), vérifiez d’abord que votre firmware Fire OS est à jour, puis relancez la recherche. Cette intégration a été annoncée au CES 2026 par NVIDIA et reste, à ce stade, limitée à ces deux modèles précis. Pour les Fire TV non listés officiellement, le sideload via ADB reste une option technique mais non garantie par NVIDIA — nous le détaillons en astuce avancée plus bas.

```
# Vérifier la version Fire OS installée en SSH/ADB (si le débogage est activé)
adb connect <IP_DE_LA_FIRE_TV>:5555
adb shell getprop ro.build.version.name
```
## Étape 7 et 8 : connecter le réseau et tester la latence

**Étape 7 — Privilégiez l’Ethernet.** Sur Shield TV Pro, LG et Samsung, le port Ethernet Gigabit est natif : branchez directement un câble RJ45 depuis votre box internet ou un commutateur réseau. Sur Chromecast avec Google TV et Fire TV Stick, il n’y a pas de port Ethernet natif : utilisez un adaptateur USB-C vers Ethernet (Chromecast) ou micro-USB vers Ethernet (Fire TV) pour éviter les micro-coupures du Wi-Fi. Si le filaire est impossible, passez impérativement sur la bande 5 GHz de votre routeur, jamais sur la bande 2,4 GHz, saturée et sujette aux interférences.

**Étape 8 — Lancez le test réseau intégré.** Dans l’application GeForce NOW, ouvrez les paramètres puis « Test de connexion ». L’app affiche le débit descendant mesuré, la latence vers le datacenter le plus proche et une recommandation de résolution/FPS. Voici les seuils officiels à connaître :

| Résolution / FPS cible | Débit minimum recommandé | Latence maximale acceptée | 
|---|---|---|
| 720p / 1080p 60 FPS | 15-25 Mbps | 80 ms (40 ms recommandé) | 
| 1440p / QHD 120 FPS | 35 Mbps | 40 ms recommandé | 
| 4K (3840×2160) 120 FPS | 45 Mbps | 40 ms recommandé | 
| 5K (5120×2180) 120 FPS (client PC uniquement) | 65 Mbps | 40 ms recommandé | 
| 1440p/1080p 240-360 FPS (mode compétitif) | 48-55 Mbps | 40 ms recommandé | 

Notez que le mode 5K 120 FPS reste réservé au client PC en 2026 : sur téléviseur, même une Shield TV Pro ou une LG OLED haut de gamme plafonnera à 4K, avec du 120 Hz disponible seulement sur une poignée de modèles Micro RGB et OLED récents.

## Étape 9 et 10 : appairer manette, clavier et gérer le HDMI-CEC

**Étape 9 — Appairage Bluetooth de la manette.** Sur Android TV/Google TV/Shield, allez dans Paramètres > Télécommandes et accessoires > Ajouter un accessoire, mettez votre manette en mode appairage (DualSense : maintenez PS + Create jusqu’au clignotement rapide ; manette Xbox : maintenez le bouton Xbox et le bouton de connexion). Sur LG webOS et Samsung Tizen, l’appairage passe par les réglages Bluetooth natifs du téléviseur, pas par l’application GeForce NOW elle-même : une fois la manette reconnue par la télé, GeForce NOW la détecte automatiquement au lancement d’un jeu.

**Étape 10 — Configurez le HDMI-CEC.** Sur Shield TV, Chromecast avec Google TV et la plupart des boîtiers Android TV, le HDMI-CEC permet d’allumer automatiquement la télé et de basculer sur la bonne entrée dès que vous lancez GeForce NOW. Activez cette option dans les paramètres d’affichage du téléviseur (l’intitulé varie : « Anynet+ » chez Samsung, « SIMPLINK » chez LG, « Bravia Sync » chez Sony). Pensez aussi au clavier et à la souris : LG webOS et Samsung Tizen supportent l’USB et le Bluetooth pour de nombreux titres PC, une option pratique pour les jeux de stratégie ou de gestion peu adaptés au pad.

## Étape 11 à 13 : régler l’audio, choisir le datacenter et finaliser la session

**Étape 11 — Réglez la sortie audio.** Privilégiez une sortie HDMI ARC ou eARC vers une barre de son ou un ampli plutôt que le Bluetooth de la télé, qui ajoute une latence audio perceptible dans les jeux de rythme ou les FPS compétitifs. Si vous devez utiliser un casque Bluetooth, appairez-le directement à la manette ou au téléphone plutôt qu’au téléviseur quand l’appareil le permet.

**Étape 12 — Vérifiez la sélection du datacenter.** GeForce NOW choisit automatiquement le datacenter le plus proche selon votre latence mesurée. Dans de rares cas (VPN actif, DNS tiers), l’app peut router vers un serveur plus éloigné : désactivez temporairement tout VPN avant de lancer une session pour laisser l’algorithme sélectionner le point de présence européen le plus proche (Paris ou Francfort pour la majorité des utilisateurs français).

**Étape 13 — Lancez une session de calibration.** Ouvrez un jeu léger (un jeu de plateforme 2D par exemple) pour valider que l’image, le son et la manette répondent correctement avant de vous lancer dans un titre exigeant. Ajustez manuellement la résolution et le bitrate dans les paramètres de streaming si le test automatique s’est montré trop conservateur par rapport à votre débit réel.

## Mini-projet complet : script de vérification pré-session pour TV

Si votre boîtier Android TV/Shield accepte le débogage ADB (activé dans Paramètres > Système > À propos > appuyez 7 fois sur le numéro de build), vous pouvez automatiser un contrôle réseau et matériel avant chaque session, directement depuis un PC ou un Raspberry Pi sur le même réseau local. Ce script Bash teste la latence vers un point de référence, mesure le débit local et vérifie que l’appareil TV répond bien sur le réseau avant de lancer la session.

```
#!/bin/bash
# pre-session-check.sh — vérification réseau avant session GeForce NOW sur TV
# Usage : ./pre-session-check.sh <IP_DU_BOITIER_TV>
TV_IP="$1"
MAX_LATENCY_MS=40
MIN_BANDWIDTH_MBPS=35
if [ -z "$TV_IP" ]; then
  echo "Usage : $0 <IP_DU_BOITIER_TV>"
  exit 1
fi
echo "=== Vérification pré-session GeForce NOW ==="
# 1. Vérifie que le boîtier TV répond sur le réseau local
if ping -c 2 -W 2 "$TV_IP" > /dev/null 2>&1; then
  echo "[OK] Boîtier TV joignable à $TV_IP"
else
  echo "[ERREUR] Boîtier TV injoignable — vérifiez le câble Ethernet ou le Wi-Fi"
  exit 1
fi
# 2. Mesure la latence moyenne vers un point de test
LATENCY=$(ping -c 10 8.8.8.8 | tail -1 | awk -F '/' '{print $5}')
LATENCY_INT=${LATENCY%.*}
if [ "$LATENCY_INT" -le "$MAX_LATENCY_MS" ]; then
  echo "[OK] Latence moyenne : ${LATENCY} ms (sous le seuil de ${MAX_LATENCY_MS} ms)"
else
  echo "[ATTENTION] Latence élevée : ${LATENCY} ms — réduisez le nombre d'appareils actifs sur le réseau"
fi
# 3. Test de débit avec speedtest-cli (pip install speedtest-cli)
if command -v speedtest-cli > /dev/null; then
  DOWNLOAD=$(speedtest-cli --simple | grep "Download" | awk '{print $2}')
  echo "[INFO] Débit descendant mesuré : ${DOWNLOAD} Mbps"
else
  echo "[INFO] speedtest-cli non installé — test de débit ignoré"
fi
echo "=== Vérification terminée ==="
```
Résultat attendu à l’exécution sur un réseau domestique en bon état :

```
=== Vérification pré-session GeForce NOW ===
[OK] Boîtier TV joignable à 192.168.1.42
[OK] Latence moyenne : 18.421 ms (sous le seuil de 40 ms)
[INFO] Débit descendant mesuré : 187.34 Mbps
=== Vérification terminée ===
```
Adaptez le script à votre configuration : remplacez l’adresse de test 8.8.8.8 par l’adresse IP de votre routeur si vous voulez isoler la latence du réseau local, ou par un serveur de test que vous jugez pertinent pour votre région.

## 5 erreurs fréquentes qui plombent l’expérience GeForce NOW sur TV

