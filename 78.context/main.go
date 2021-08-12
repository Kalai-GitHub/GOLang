package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context: \t", ctx)             //context.Background
	fmt.Println("context error: \t", ctx.Err()) //<nil>
	fmt.Printf("context type: \t%T\n", ctx)     //*context.emptyCtx
	fmt.Println("============")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context: \t", ctx)             //context.Background.WithCancel
	fmt.Println("context error: \t", ctx.Err()) //<nil>
	fmt.Printf("context type: \t%T\n", ctx)     //*context.cancelCtx
	fmt.Println("cancel: \t\t", cancel)         //0x107c240
	fmt.Printf("cancel type: \t%T\n", cancel)   //context.CancelFunc
	fmt.Println("============")

	cancel()

	fmt.Println("context: \t", ctx)             //context.Background.WithCancel
	fmt.Println("context error: \t", ctx.Err()) //context canceled
	fmt.Printf("context type: \t%T\n", ctx)     //*context.cancelCtx
	fmt.Println("cancel: \t\t", cancel)         //0x107c240
	fmt.Printf("cancel type: \t%T\n", cancel)   //context.CancelFunc
	fmt.Println("============")

}
