package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"textgrpc/send"
)

type send_struct struct {
	Timestamp  int64
	Metric     string
	Dimensions map[string]string
	Value      float64
	AlertType  string
}

const (
	address = ":8081" //定义端口

)

var in send_struct

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	sendClient := send.NewSendServiceClient(conn)
	//定义数据测试
	in.Dimensions = make(map[string]string)
	in.Dimensions["ip"] = "1.1.1.1"
	in.Timestamp = 1642125600
	in.Metric = "cpu_rate"
	in.Value = 0.9
	in.AlertType = "WARN"

	sendRes, err := sendClient.Send(context.Background(), &send.SendReq{
		Timestamp:  in.Timestamp,
		Metric:     in.Metric,
		Value:      in.Value,
		Dimensions: in.Dimensions,
		AlertType:  in.AlertType})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Code: %d, msg: %s", sendRes.Code, sendRes.Msg)
}
