# Weather-Monster


Dependencies

Golang 1.16 or higher

Docker for Mac/Windows

github.com/shyam81992/Site-Information-job // Which is used to notify the subscribers on temperature record creation

Steps to run the project 

    1. Create a base folder and clone the projects github.com/shyam81992/Site-Information.
    2. Run go mod tidy to install the dependencies.
    
    3. To run the testcases run the command . ../config/dev.sh && go test ./... -v

Note: If you are running on windows please install git bash so that you can able to run the shell script  ". config/dev.sh"
    


Weather-Monster-Job
Now it only listens to the temperature creation event and notify the subscriber which can be improved to have filter (ex : filter option can be provided at the webhook level). If notification rate per second is high we can improve the performance of it by scaling through region or by any other factor.(Kinesis, sqs, sns, rabbitmq)

 
