# Go Template walkthrough

[![Build Status][travis-svg]][travis-link]
[![Coverage Status][coverall-svg]][coverall-io]

This walkthrough will explain you how to correctly create a microservice based on our Go Template from the DevOps Console.

## Create a microservice

In order to do so, access to [Mia-Platform DevOps Console](https://console.cloud.mia-platform.eu/login), create a new project and go to the **Design** area.

From the Design area of your project select _Microservices_ and then create a new one, you have now reached [Mia-Platform Marketplace](https://docs.mia-platform.eu/development_suite/api-console/api-design/marketplace/)!
In the marketplace you will see a set of Examples and Templates that can be used to set-up microservices with a predefined and tested function.

For this walkthrough select the following template: **Go Template**.
Give your microservice the name you prefer, in this walkthrough we'll refer to it with the following name: **my-go-service-name**. Then, fill the other required fields and confirm that you want to create a microservice.  
A more detailed description on how to create a Microservice can be found in [Microservice from template - Get started](https://docs.mia-platform.eu/development_suite/api-console/api-design/custom_microservice_get_started/#2-service-creation) section of Mia-Platform documentation.

## Look inside your repository

After having created your first microservice (based on this template) you will be able to access to its git repository from the DevOps Console. Inside this repository you will find an [router.go](https://github.com/mia-platform-marketplace/Go-Template/blob/master/router.go) file with the following lines of code:

```go
package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

func setupRouter(router *mux.Router) {
    // Setup your routes here.
    router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    })
}
```

Wonderful! You are now ready to start customizing your service! Read next section to learn how.

## Add a Welcome route

Now that you have successfully created a microservice from our Go template you will add an *welcome* route to it.

In order to do so, you should add the following line inside of the import file:

```go
"encoding/json"
```

Then, you should create a *Welcome* struct that will contain your welcoming message:

```go
type Welcome struct {
    Msg string `json:"msg"`
}
```

Lastly, you should add a *welcoming* route to your service. Below you can see how the *router.go* file will look like after having defined this new route and having applied all previous modifications:

```go
package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Welcome struct {
    Msg string `json:"msg"`
}

func setupRouter(router *mux.Router) {
    // Setup your routes here.
    router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    })
    router.HandleFunc("/welcome", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        welcome := Welcome{
            Msg: "Welcome!",
        }
        body, err := json.Marshal(&welcome)
        if err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            w.Write(nil)
        }
        w.WriteHeader(http.StatusOK)
        w.Write(body)
    })
}
```

After committing these changes to your repository, you can go back to Mia Platform DevOps Console.

## Expose an endpoint to your microservice

In order to access to your new microservice it is necessary to create an endpoint that targets it.  
In particular, in this walkthrough you will create an endpoint to your microservice *my-go-service-name*. To do so, from the Design area of your project select _Endpoints_ and then create a new endpoint.
Now you need to choose a path for your endpoint and to connect this endpoint to your microservice. Give to your endpoint the following path: **/go-template**. Then, specify that you want to connect your endpoint to a microservice and, finally, select *my-go-service-name*.  
Step 3 of [Microservice from template - Get started](https://docs.mia-platform.eu/development_suite/api-console/api-design/custom_microservice_get_started/#3-creating-the-endpoint) section of Mia-Platform documentation will explain in detail how to create an endpoint from the DevOps Console.

## Save your changes

After having created an endpoint to your microservice you should save the changes that you have done to your project in the DevOps console.  
Remember to choose a meaningful title for your commit (e.g 'created service my_go_service_name'). After some seconds you will be prompted with a popup message which confirms that you have successfully saved all your changes.  
Step 4 of [Microservice from template - Get started](https://docs.mia-platform.eu/development_suite/api-console/api-design/custom_microservice_get_started/#4-save-the-project) section of Mia-Platform documentation will explain how to correctly save the changes you have made on your project in the DevOps console.

## Deploy

Once all the changes that you have made are saved, you should deploy your project through the DevOps Console. Go to the **Deploy** area of the DevOps Console.  
Once here select the environment and the branch you have worked on and confirm your choices clicking on the *deploy* button. When the deploy process is finished you will receveive a pop-up message that will inform you.  
Step 5 of [Microservice from template - Get started](https://docs.mia-platform.eu/development_suite/api-console/api-design/custom_microservice_get_started/#5-deploy-the-project-through-the-api-console) section of Mia-Platform documentation will explain in detail how to correctly deploy your project.

## Try it

Now, if you launch the following command on your terminal (remember to replace `<YOUR_PROJECT_HOST>` with the real host of your project):  

```shell
curl <YOUR_PROJECT_HOST>/go-template/welcome
```

you should see the following message:

```json
{"msg":"Welcome!"}
```

Congratulations! You have successfully learnt how to modify a blank Go template into a _Welcome_ microservice!

[travis-svg]: https://travis-ci.org/mia-platform/api-explorer.svg?branch=master
[travis-link]: https://travis-ci.org/mia-platform/api-explorer
[coverall-svg]: https://coveralls.io/repos/github/mia-platform/api-explorer/badge.svg
[coverall-io]: https://coveralls.io/github/mia-platform/api-explorer
