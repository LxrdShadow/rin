# 🚀 RIN – REST Interface Nexus

**RIN** (REST Interface Nexus) est une bibliothèque Go modulaire et légère conçue pour simplifier l'interaction avec des APIs REST.  
Elle fournit une interface claire et robuste pour créer des outils DevOps, des clients API, ou des automatisations d'infrastructure.

---

## 🎯 Objectif

RIN vise à :
- Simplifier l'écriture de requêtes HTTP dans Go
- Centraliser la gestion de l’authentification, des headers, et du parsing JSON
- Servir de base solide pour des **outils CLI DevOps** fiables et maintenables

---

## ⚙️ Fonctionnalités

- ✅ Support complet des méthodes HTTP : `GET`, `POST`, `PUT`, `DELETE`
- ✅ Authentification : token bearer, headers personnalisés
- ✅ Parsing automatique des réponses JSON dans des structs Go
- ✅ Gestion des erreurs HTTP centralisée
- ✅ Extensible via middlewares : retry, logs, instrumentation

---

## 🧪 Exemple d'utilisation

```go
package main

import (
    "fmt"
    "github.com/Rakotoarilala51/rin"
)

func main() {
    client := rin.NewClient("https://api.github.com", rin.WithAuthToken("your_token_here"))

    var repo GitHubRepo
    err := client.Get("/repos/RaMaitre/rin", &repo)
    if err != nil {
        panic(err)
    }

    fmt.Println("Nom du repo :", repo.Name)
}
