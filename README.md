<h1 align="center"> Tuto-go </h1>

Tuto Golang

## Initialisation du projet

```bash
go mod init <nom_projet>
```

Création du main de fichier main.go

``` go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
} 
```

## Pour lancer le run

```bash
go run . 
```

ou

```bash
go run < Nom programme > 
```

Récupérer les dépendences

```bash
go get golang.org/x/example
```

## Definir le workspace

``` bash  
 go work init <nom_dossier>
```

Contenu du fichier

```go
go 1.18

use <nom_du_programme_a_utiliser>
```

## Les types en go

```text
string 
uint
byte
uint8
uint16
uint32
uint64
int
int8
int16
int32
int64
rune  
float32
float64
bool 
```

## Saisie clavier

```go
scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
scanner.Scan()                      // lancement du scanner
entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
```

## Convertion string en int

```go
"strconv"
strconv.Atoi("22")
```

## Temps

```go
"time"
time.Now()
```

## Switch

  Le switch est equivalent dans les autres langages
  Pas obligé de mettre une variable après le switch
  possibilité de mettre des conditions à la place des case

## Les boucles

 pas boucle while

```go
 for <condition > {}

 for initialisation ; condition ; incrementation {}
  
 for index, j := range jours {}

 for _, j := range jours  {}// si je  ne veux pas l'index mettre _
```

Pour sortir d'un boucle

```go
break 
```

Passer à la table suivantes

```go
continue
```

## Les fonctions

```go
  func nom_function(param ...int ) ( v1 int )
  {
 total := 0
    for _, value := range param { //j'ai mis un underscore "_" car je ne souhaite pas récupérer l'index de la range
        total += value
    }
    return total
  }  
  func nom_function(v1I int , v2I) ( v1o int, v2o int  )
  {
  }  
```

### Fonction anomyme

```go
func(x float64) float64 { return math.Sqrt(x) }
```

## Tableau

### Statique

```go
var nomTableau [taille] type
len(tableau1)
```

### Dynamique ou Slice

#### Déclaration

```go
var nombres = make([]int, 5) 
```

ou

```go
var nombres = []int{0, 0, 0, 0, 0} 
```

#### Ajout élément

Utiliser la fonction append()

```go
nombres = append( nombres ,  12)
```

#### Suppression d'élément

Il n'existe pas réelement de suppression
Cependant on peut combiner les tableaux pour pallier ce probleme.

```go
mois = append(mois[:indexASupprimer], mois[(indexASupprimer+1):]...)
```

#### Copie d'éléménts

Attention slice doivent avoir la même taille

```go
copy(animaux2, animaux1)
```

## Maps

### Déclaration map

```go
var notes map[string]int
```

ou

```go
var notes = make(map[string]int)
```

### Instanciation

```go
 var notes = make(map[string]int)
notes["Hatim"] = 20
notes["Alex"] = 18
notes["Kevin"] = 15
```

### Boucle sur map

```go
notes := map[string]int{"Hatim": 20, "Alex": 18, "Kevin": 15, "Robert": 17}

for eleve := range notes {
fmt.Println("La note de ", eleve, "est", notes[eleve])
}
```

### Suppression sur map

fonction => delete()  

```go
delete(notes, "Hatim")
```

## Les pointeurs

 Définition

```go
  var point *int
  var vartest int = 12
  point = &vartest
```

Pour obtenir le contenu du pointeur mettre *

```go
 fmt.Printf( " %d" ,*point)
```

Pour obtenir l'adresse de la variable

```go
 fmt.Printf( "%p" ,point)
```

## Les Structures

 si les variables ou méthodes ne commence pas par une majuscule elles sont considérées comme privées.

### Définition

```go
type Nom_de_ma_structure struct {
    nom_var_1 type;
    nom_var_2 type;
    nom_var_3 type;
}
```

### les Constructeurs

```go
func New(nom_var_1 type,nom_var_2 type, nom_var_3 type) Nom_de_ma_structure {
    ma_stucture:= Nom_de_ma_structure {nom_var_1, nom_var_2, nom_var_3 }
    return ma_stucture
}

func main(){

  test := nom_de_ma_structure.New("test", "test", "test")
}
```

### Affectation

```go
// déclaration de la structure nommée Personnage
type Personnage struct {
    nom string
    age int
}
perso := Personnage{"Hatim", 20} // surchargement des valeurs par défaut 
```

### Modifications

```go
perso.age = 50
```

### Affectations des methodes

```go
func (p Personnage) affichage() {}
```

### Utilisation des pointeurs

exemple création de constructeur

```go
func (p *Personnage) Init(nom string, vie int, puissance int, mort bool, inventaire [3]string) {
    p.nom = nom
    p.vie = vie
    p.puissance = puissance
    p.mort = mort
    p.inventaire = inventaire
}
```

## Les interfaces

### Définition des interfaces

```go
type Interface1 interface { // création de L'interface Interface1
    Methode1() float64 // signature de la Methode1()
    Methode2() float64 // signature de la Methode2()
}
```

Pour implémenter une interface dans Go, il suffit
d'implémente toutes les méthodes de l'interface.

```go
type Interface1 interface {
    Methode1() float64
    Methode2() float64
}

type Reel struct {
    largeur  float64
    longueur float64
}

/* 
Pour implémenter une interface dans Go, il suffit
d'implémente toutes les méthodes de l'interface. Ici on
implémente la méthode Air() de l'interface Forme.
 */
func (r Reel) Methode1() float64 {
 return r.largeur * r.longueur
}

func (r Reel) Methode2() float64 {
 return (r.largeur * r.longueur)*2
}
```

## Les erreurs

Utilisation des erreurs

```go
errors.New("message d'erreur")
```

exemple

```go
 func testErreur(nbr int) (int ,  error){
      if nbr > 0 {
        return  nbr+2 , nil  
      } else {
        return   0  , errors.New("Erreur: nombre inferieur à 0")
      } 
 }

 func main(){
  nbr:= 12
  nbr , err := testErreur(nbr)

  if err == nil {
    panic(err) // arret le programme et affiche le statut
  }else 
  {
    ftm.Println(nbr)
  }
 }
```

## les modules

Ajouter des module

Si on fait des dossiers mettre dans le import  du main  le  chemin relatif en fonction du main

## les goroutines

### les channels

## Lectures de fichiers

## Sources

<https://devopssec.fr/article/lire-et-ecrire-dans-un-fichier-golang#begin-article-section>
