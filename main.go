package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

func main() {
	ak, exist := os.LookupEnv("AWS_ACCESSKEY_ID")
	if !exist {
		log.Fatal("AWS_ACCESS_KEY_ID not defined")
	}
	sak, exist := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !exist {
		log.Fatal("AWS_SECRET_ACCESSKEY not defined")
	}
	cred := credentials.NewStaticCredentialsProvider(ak, sak, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
		config.WithCredentialsProvider(cred))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := route53.NewFromConfig(cfg)

	input := route53.ListHostedZonesInput{}
	resp, err := svc.ListHostedZones(context.TODO(), &input)
	if err != nil {
		log.Fatalf("unable to list hosted zones, %v", err)
	}

	for _, zone := range resp.HostedZones {
		fmt.Printf("%+v\n", zone)
		fmt.Print(*zone.Name)
	}
}
