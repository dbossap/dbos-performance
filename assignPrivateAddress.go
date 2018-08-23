package main

import (
      "fmt"

    	"github.com/aws/aws-sdk-go/aws"
    	"github.com/aws/aws-sdk-go/aws/awserr"
    	"github.com/aws/aws-sdk-go/aws/session"
      "github.com/aws/aws-sdk-go/aws/credentials"
      "github.com/aws/aws-sdk-go/service/ec2"
      "github.com/coreos/go-iptables/iptables"
  )


/*
This function is used to create a session with custom credentials.
returns, session.session
Note : Rather then passing credntials here check if these can be picked from envirnment or shared config
*/
func createSessionWithCustomCredentials() *session.Session {

  sess := session.Must(session.NewSession(&aws.Config{
	   Region: aws.String("us-east-1"),
     Credentials: credentials.NewStaticCredentials("AKIAIPG6CFN3LWQAGXQQ", "BLmbP/czQq0u73arva6EnoIqr/C3tlHzup1z5um9", ""),
  }))

  return sess
}


/*
This function is used call the AssignPrivateIpAddress function with custom session.
returns, nothing may change retunr type later.
*/
func callAssignPrivateIpAddress(customSession *session.Session) {

  serviceClient := ec2.New(customSession)

  input := &ec2.AssignPrivateIpAddressesInput{
      NetworkInterfaceId: aws.String("eni-a29a6c91"),
      PrivateIpAddresses: []*string{
        aws.String("10.11.49.101"),
      },
      AllowReassignment: aws.Bool(true),
  }

  result, err := serviceClient.AssignPrivateIpAddresses(input)

  if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
          switch aerr.Code() {
          default:
              fmt.Println(aerr.Error())
          }
      } else {
          // Print the error, cast err to awserr.Error to get the Code and
          // Message from an error.
          fmt.Println(err.Error())
      }
      return
  }

  fmt.Println("result is ", result)
}

func main() {

  /*sess, err := session.NewSession(&aws.Config{
    Region:      aws.String("us-east-1"),
    Credentials: credentials.NewStaticCredentials("AKIAIPG6CFN3LWQAGXQQ", "BLmbP/czQq0u73arva6EnoIqr/C3tlHzup1z5um9", ""),
  })*/

  ipt, err := iptables.New()
	if err != nil {
		panic(fmt.Sprintf("New failed: %v", err))
	}

	ipts := []*iptables.IPTables{ipt}

  fmt.Println("Can connect : ", ipts)

  customSession := createSessionWithCustomCredentials()
  callAssignPrivateIpAddress(customSession)

}
