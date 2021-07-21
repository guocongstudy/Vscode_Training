package pkg

import "fmt"

func Hello() {
	fmt.Println("hello")
	score := 88
	if score >= 60 {
		fmt.Println("A")
	}

	var (
		gender string
		age    uint
		weight uint
	)
	fmt.Println("请输入你的性别male/woman/else?")
	fmt.Scan(gender)
	age = 20
	weight = 200

	if gender == "male" {
		fmt.Println("是男的")
		if age >= 18 {
			fmt.Println("已经成年")
		} else {
			fmt.Println("是男的，但是未成年")
		}
		if weight >= 200 {
			fmt.Println("是个胖子")
		} else {
			fmt.Println("身材还行")
		}

	} else {
		fmt.Println("不是男的")
	}

	switch gender {
	case "male":
		{
			fmt.Println("男")
		}
	case "woman":
		{
			fmt.Println("女生")
		}
	//在switch中，以上条件都不满足。for用的else，switch会用到default
	//上面所有条件都不满足的情况下执行
	default:
		fmt.Println("你是一个迷你小太监")
	}

}
