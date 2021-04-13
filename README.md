
[![Build status](https://dev.azure.com/noon-homa/Resharmonics/_apis/build/status/resharmonics-go-client)](https://dev.azure.com/noon-homa/Resharmonics/_build/latest?definitionId=42)



# Basic usage

See the `test` folder for an example

```
import "github.com/JoseFMP/resharmonics"


func main(){

    creds := resharmonics.Credentials{ Username: "foo", Password: "baar" }

    rhClient, errSettingUpClient := resharmonics.Init(creds)

    if(errSettingUpClient != nil){
        panic("Something went wrong setting up Resharmonics client")
    }

    rhClient.....
    // go for gold!

}
```

# Motivation

This is a Go client to consume [Resharmonics](https://www.resharmonics.com/) API.

To the best of my knowledge there isn't any other publicly available.

# API

Some endpoints here: https://api.rerumapp.uk/swagger/index.html

Some other here: https://app.swaggerhub.com/apis/resharmonics-apis/guest-portal/1.1.0#

# Collaboration

If you want to collaborate just mail me at jose@mingorance.engineer