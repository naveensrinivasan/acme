#acme mock server

stub that simply returns the string as a certificate when a HTTP GET request to
 `/cert/{domain}` is made.

A mock code base that:
* properly handles certificates for different domains
* allows certificates to live for 10 minutes before expiring
* sleeps for 10 seconds when a new certificate is generated
* generates its own certificate and keeps its certificate up-to-date when it expires
* can generate multiple certificates at the same time

##Bootstrap
Bootstrapping is done by running the `./bootstrap.sh`.

##Building and Running
To build the app `./build.sh`

**Tested it on a mac** should work on other platforms. 

##Running the application
* Build the app which should generate an executable `acme`.
* Run the `acme` executable which is webserver.
* Hit the endpoint with something like  `http://localhost:8080/cert/naveen`. This should return a 
`json` something like
```json
{
"certificate":"277686f3-5634-4e43-9f25-be03f71f360f",
"domain":"naveen"
}
``` 
