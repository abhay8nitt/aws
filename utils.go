import (
  	"github.com/aws/aws-sdk-go/aws"
  	"github.com/aws/aws-sdk-go/aws/session"
  	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AllServices struct{
	S3svc   	*s3.S3	
	S3Downloader 	*s3manager.Downloader
	Dynamodbsvc     *dynamodb.DynamoDB
	CloudWatchsvc   *cloudwatch.CloudWatch
}

func New() *AllServices{
	allservices:= &AllServices{}
	config := &aws.Config{
		Region: aws.String(a.Args.Region),		
	}
	sess, _ := session.NewSession(config)
	
	allservices.S3svc = s3.New(sess)
	allservices.Downloader = s3manager.NewDownloader(sess)
	allservices.Dynamodbsvc = dynamodb.New(sess)
	allservices.CloudWatchsvc =cloudwatch.New(sess)
	return allservices
}

func (a *AllServices) List(bucket string, prefix string) []string {
	list := make([]string, 0)
	i := 0
	err := a.Service.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: bucket,
		Prefix: prefix,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		for _, obj := range p.Contents {
			list = append(list, *obj.Key)
		}
		i = i + 1
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return list
	}
	return list
}
