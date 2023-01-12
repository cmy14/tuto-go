# tuto-go
tuto  go 
Golang

## Initialisation du projet 
go mod init <nom_projet>

Création du main de fichier main.go

package main 

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

## Pour  lancer le run 

go run . Ou <nom fichier>

 recupere les dependance 

go get golang.org/x/example


 definir le workspace
go work init <nom_dossier>

go 1.18

use <nom du programme a utiliser>


 ## les types en go   
  string   int uint bool , float 

## saisie  clavier 
   scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
   scanner.Scan()                      // lancement du scanner
    entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable


##  convertion string en int

"strconv" 
strconv.Atoi("22")

##  temps 

"time"
time.Now() 

##  switch

  pas obliger de mettre ube varaible après 
## boucle 
 pas  boucle while

 for condition {}

 for initialisation ; condition ; incrementation 

  break & continue 


  for index, j := range jours 

  for _, j := range jours  si je  ne veux pas  l'index 

##  les function 

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

   ### function anomyme 
   
   func(x float64) float64 { return math.Sqrt(x) }
   
   ## Tableau 
var nomTableau [taille] type
len(tableau1)
