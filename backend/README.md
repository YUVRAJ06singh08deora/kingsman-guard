## Lock Management System
 
This system provides management of smart locks, enabling users to configure locks, set up temporary guest codes, and enable/disable lock automation.
 
## Service details:
 - Local AWS credentials need to be provided for the deploy to AWS
 - Get URL from AWS API Gateway to run the application
 
### Build
To build this micro-service, `cd` into the scripts folder and run the appropriate script based on your OS:

- Linux : `bash ./compileBinariesOnLinux.sh`
- Windows : `bash ./compileBinariesOnWindows.sh`
 
### Deploy
To deploy this micro-service in AWS, `cd` into the scripts folder and run `bash ./createInfra.sh`
 
### Destroy
To destroy this micro-service in AWS, `cd` into the scripts folder and run `bash ./destroyInfra.sh`
 
## Tools
- `go mod tidy`
  - Builds/cleans/updates the go.mod file with the dependencies required in the source files