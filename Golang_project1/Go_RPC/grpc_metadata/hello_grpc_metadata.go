package grpcmetadata_test

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"google.golang.org/grpc/metadata"
)

func HelloGrpcMetadata() {
	//创建 metadata的第一种方式
	md := metadata.New(map[string]string{
		"key1": "value1",
		"key2": "value2",
	})

	//创建 metadata的第二种方式, 使用Pairs函数
	// Pairs函数的参数是可变参数, 需要成对出现
	// 例如: "key1", "value1", "key2", "value2"
	// 也可以使用Pairs函数创建一个空的metadata
	// 例如: metadata.Pairs()
	// key不区分大小写，会统一转换为小写
	md2 := metadata.Pairs(
		"key1", "value1", //即使是 key 和 value 之间也使用逗号分隔
		// 也可以使用空格分隔
		// "key1 value1",
		// "key2 value2",
		"key2", "value2",
	)

	fmt.Println(md)
	fmt.Println(md2)

	//新建一个有 metadata 的 context
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reponse, err := client.SomeRPC(ctx, someRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

}
