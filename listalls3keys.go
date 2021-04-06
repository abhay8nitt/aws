import (
  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)  

func List(bucket string, prefix string) []string {
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
