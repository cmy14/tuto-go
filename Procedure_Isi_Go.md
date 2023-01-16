# PROCEDURE D'INSTALLATION GOLANG  

## Installer GO

<https://go.dev/doc/install>

```text
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz
```

## Variable d'environnement

Ajouter dans le Path le fichier .profile

```text  
export PATH=$PATH:/usr/local/go/bin
```

Ajouter dans le Path le fichier .bashrc

```text
export GOPROXY=https://nexus.dev.isi.nc/repository/go-proxy-group/

export GONOSUMDB=isi.nc/,.isi.nc/*

export PATH=$PATH:$HOME/go/bin
```

## Demander le fichier de config pour le mettre dans .kube

Demander le fichier de config à damien

```text
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data:
    ********************************************************************************************************************************************************************************************************************************************************************************************
    server: **************
  name: *********
contexts:
- context:
    cluster: ********
    namespace:  *******
    user: ******
  name: ******
current-context: ******
kind: Config
preferences: {}
users:
- name: ******
  user:
    token: ****************
```

## Installer Modd

```bash
go install github.com/cortesi/modd/cmd/modd@latest
```

## Installer swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

puis

```bash
go generate
```

## Lancer projet

### Soit dans mle terminal

```text
 go run . ( dossier ou se trouve le main)
```

###  Soit f5 si tu as fait le launch.json

```text
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/intranet-api",
            "env": {
                "****":"****"
              }
        }
    ]
}
```

### Soit mood si on est  passé par mood

```text
@cmd="***-api -enable-swagger=true -debug=false"
#@cmd="mail-sms-notifier"

**/*.go !data/** {
    prep: [[ -f /tmp/tls.key ]] && echo "certificate exists" || openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout /tmp/tls.key -out /tmp/tls.crt -subj /CN=localhost
    prep: go test ./...
    prep: go install ./...
    daemon:  **:**
    @cmd
}
```

### Comment aller cherher les secrets pour la valeur des params

``` bash
kubectl --context <context> -n <namspace> get secrets <nom_fichier> -oyaml
```

 pour chiffer et  dechiffrer en base64
 ```bash 
  echo "testkey" |base64 -d 
  ```

  ```bash
  echo "clefSecretCrypteeb2Rvb3Bhc3N3b3Jk" |base64 -d 
  ```