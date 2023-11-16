package try

import (
	"context"
	"fmt"
	"time"
)

func TrySomething() {

	ctx := context.Background()
	// hashSet := make(map[string]struct{}, 1)
	// interface{} is better for SET/HashSet usage
	hashSet := make(map[string]interface{}, 1)

	myValue, ok := hashSet["iamcod3r"].(int)
	if !ok {
		fmt.Println("not found")
		// return
	}

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("5 second done")
	case <-ctx.Done():
		fmt.Println("context done")
	}

	fmt.Println(myValue)
}
