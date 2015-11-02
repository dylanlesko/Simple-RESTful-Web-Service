# A Simple RESTFUL Program
###Dependencies
	Golang
	mongodb
	
###Environment Setup
	As Setup on my Ubuntu virtual machine. Tested on iLab machines as well.
	- Path Structure
	    $HOME/go
	        -bin/
	        -pkg/
	        -src/
		        -main/
			        -main.go/
    - Environment Variables
	    export GOPATH=$HOME/go/
    - Package Installation
	    1. Option 1-
		    cd ~/go/src/main
		    go get
	    2. Option 2-
		    I noticed the mongo packages took me a long time to download, so I will 
		    include a zipped copy of downloaded packages.
    -mongoDB
	    Installed mongoDB by following these instructions[1]
	    -sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
	    -echo "deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.0 multiverse"
		     | sudo tee /etc/apt/sources.list.d/mongodb-org-3.0.list
	    -sudo apt-get update
		-sudo apt-get install -y mongodb-org

	    
### Running the Program
	navigate to $HOME/go/src/main/ 
> POST Operation
>> $ go run test.go  -url="http://localhost:1234/Student" -method=Create  -data=’{"NetID":"147001234", "Name":"Mike","Major":"Computer Science","Year":2015,"Grade":90,"Rating":"D"}’

> GET Operation
>> $ go run test.go  -url="http://localhost:1234/Student/getstuent?name=Mary" -method=list

>DELETE Operation
>>$ go run test.go –url="http://localhost:1234/Student" -method=remove  -year=2010

>Update Operation
>>$ go run test.go –url="http://localhost:1234/Student" -method=update

>LIST Operation
>>$ go run test.go -url="http://localhost:1234/Student/listall" -method=list


###External Links
	[1]: https://docs.mongodb.org/getting-started/shell/tutorial/install-mongodb-on-ubuntu/#install-mongodb
	