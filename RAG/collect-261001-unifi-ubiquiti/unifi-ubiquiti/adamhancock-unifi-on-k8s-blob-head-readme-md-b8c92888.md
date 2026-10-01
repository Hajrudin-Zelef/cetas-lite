---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/adamhancock-unifi-on-k8s-blob-head-readme-md-b8c92888
title: "Apply all manifests"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2024-01-01", "2026-01"]
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/adamhancock-unifi-on-k8s-blob-head-readme-md-b8c92888.md
source_anchor: ""
source_lines: [1, 100]
sha256: 8680bba5f4cd3e59d6f115fd958a2f38eac0b2f92a40e0b472887a52e2d8e359
---

# Apply all manifests

Deploy the Ubiquiti UniFi Network Application on Kubernetes with an external MongoDB database.
Updated January 2026: This guide has been completely rewritten to use the new linuxserver/unifi-network-application image. The previous linuxserver/unifi-controller image was deprecated on January 1, 2024.
┌─────────────────┐     ┌─────────────────────────────┐
│     MongoDB     │◄────│  UniFi Network Application  │
│   StatefulSet   │     │        StatefulSet          │
└─────────────────┘     └─────────────────────────────┘
         │                           │
         └───────────────────────────┴──► LoadBalancer / Ingress
- Kubernetes cluster (managed or bare metal)
- kubectl configured
- cert-manager (optional, for TLS on web UI)
- Persistent storage provisioner
- MetalLB or cloud load balancer (for device adoption)
git clone https://github.com/adamhancock/UniFi-on-k8s.git
cd UniFi-on-k8s/yaml
Edit unifi-secret.yaml and set secure passwords:
stringData:
  mongo-root-password: "your-secure-root-password"
  mongo-password: "your-secure-unifi-password"
Edit unifi-ingress.yaml and replace:
- unifi.yourdomain.tld with your hostname
- Update the cert-manager.io/cluster-issuer annotation if needed
For bare metal with MetalLB, edit unifi-service.yaml and uncomment/set:
loadBalancerIP: 192.168.1.100# Create namespace
kubectl create namespace unifi
# Apply all manifests
kubectl apply -f . -n unifikubectl get pods -n unifi -w
Wait for both unifi-db-0 and unifi-network-application-0 to be Running.
- Via Ingress: https://unifi.yourdomain.tld
- Via LoadBalancer: https://<EXTERNAL-IP>:8443
For UniFi to adopt devices, you need to set the Inform Host:
- Go to Settings > System > Advanced
- Set Inform Host to your LoadBalancer IP or a hostname accessible by your devices
- Check Override
ssh ubnt@<device-ip>
set-inform http://<controller-ip>:8080/inform
Default device password is ubnt.
| Port | Protocol | Purpose | 
|---|---|---|
| 8443 | TCP | Web UI (HTTPS) | 
| 8080 | TCP | Device communication | 
| 3478 | UDP | STUN | 
| 10001 | UDP | AP discovery | 
| 8843 | TCP | Guest portal HTTPS | 
| 8880 | TCP | Guest portal HTTP | 
| 6789 | TCP | Mobile throughput test | 
| 5514 | UDP | Remote syslog | 
If using Traefik instead of nginx-ingress, you need to allow insecure backend certificates:
# Traefik v2+
apiVersion: traefik.io/v1alpha1
kind: ServersTransport
metadata:
  name: unifi-transport
  namespace: unifi
spec:
  insecureSkipVerify: true
---
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: unifi
  namespace: unifi
spec:
  entryPoints:
    - websecure
  routes:
    - match: Host(`unifi.yourdomain.tld`)
      kind: Rule
      services:
        - name: unifi-network-application
          port: 8443
          serversTransport: unifi-transport
  tls:
    certResolver: letsencrypt
If migrating from linuxserver/unifi-controller:
- Take a full backup from the old UniFi web UI (Settings > System > Backup)
- Include history if you want to preserve statistics
- Shut down the old controller
- Deploy this new setup with a fresh database
- Restore from backup using the setup wizard
Check if your CPU supports AVX (required for MongoDB >4.4 on x86_64):
grep -o avx /proc/cpuinfo | head -1
If no output, use mongo:4.4 instead of mongo:7.0 in mongodb-statefulset.yaml.
- Ensure port 8080 is accessible from your devices
- Check the Inform Host setting is correct
- Verify the LoadBalancer has an external IP: kubectl get svc -n unifi
# UniFi logs
kubectl logs -f unifi-network-application-0 -n unifi
# MongoDB logs
kubectl logs -f unifi-db-0 -n unifi
| File | Description | 
|---|---|
| unifi-secret.yaml | MongoDB credentials (edit before deploying!) | 
| mongodb-configmap.yaml | MongoDB init script | 
| mongodb-service.yaml | MongoDB headless service | 
| mongodb-statefulset.yaml | MongoDB StatefulSet | 
| unifi-statefulset.yaml | UniFi Network Application | 
| unifi-service.yaml | LoadBalancer service | 
| unifi-ingress.yaml | Ingress for web UI | 
MIT
