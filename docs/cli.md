## cmd cli

add github token

## download and install an app cli

```
cd cmd
```

first add a store, the store is used to store how an app is to be installed and its name

```
go build main.go && sudo ./main apps --store-add=true --app=flow-framework  --version=latest --download-path=/home/aidan/downloads 
```

Once a store is added you can install a new app, you need to pass in the

```
go build main.go && sudo ./main apps --store-add=false --install=true  --app=flow-framework  --version=latest --download-path=/home/aidan/downloads --token= TOKEN
```

