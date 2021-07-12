
[![Build status](https://dev.azure.com/noon-homa/Resharmonics/_apis/build/status/resharmonics-go-client)](https://dev.azure.com/noon-homa/Resharmonics/_build/latest?definitionId=42)

# Motivation

This is a Go client to consume [Resharmonics](https://www.resharmonics.com/) API.

To the best of my knowledge there isn't any other library publicly available to consume Resharmonics API.



# Getting started

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



# API

Currently Resharmonics exposes two APIs that are in principle unrelated (from an IT point of view) but related (they mean the same data):

* Rerum API:  https://api.rerumapp.uk/swagger/index.html - Mostly "backend" based view of the data. Information lies in a data lake where you can hardly use any filter when quering the data. In practice this means that you should use this API to "scrap" data and import it into your own systems with the structure that fits your needs.

* Guest portal: https://app.swaggerhub.com/apis/resharmonics-apis/guest-portal/1.1.0# - This part isn't yet covered in this library. This API provides a "guest" point of view to access the data. So whenever you interact with this API, each request has an implicit guest associated with it. There are some caveats though, as currently the guest needs to be separately registered in the gust portal.  

&nbsp;
&nbsp;

# Collaboration

If you want to collaborate just mail me at jose@mingorance.engineer