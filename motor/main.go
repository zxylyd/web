package main

import (
	"flag"
	"fmt"
	"os"
)

// var motorType string
// var T, L, M, R float64

// func init() {
// 	flag.StringVar(&motorType, "motorType", "", "指定电机型号:motorType=\"4810\"")
// 	flag.Float64Var(&T, "T", 0, "指定匝数：T=2000")
// 	flag.Float64Var(&L, "L", 0, "指定线径：L=0.17")
// 	flag.Float64Var(&M, "M", 0, "指定重量：M=45")
// 	flag.Float64Var(&R, "R", 0, "指定电阻：T=80")
// 	flag.Parse()

// }
// func main() {

//		fmt.Println("hello go", motorType)
//	}

type MyFlagSet struct {
	*flag.FlagSet
	cmdComment string // 二级子命令本身的注释
}

func main() {
	// 匝数计算
	TCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("T", flag.ExitOnError),
		cmdComment: "计算匝数：",
	}
	TCmd.Bool("a", false, "Show all containers (default shows just running)")
	TCmd.Bool("s", false, "Display total file sizes")

	// 电阻计算
	RCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("R", flag.ExitOnError),
		cmdComment: "计算电阻：",
	}
	RCmd.Int("c", 1, "CPU shares (relative weight)")
	RCmd.String("name", "", "Assign a name to the container")

	// 线径计算
	LCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("L", flag.ExitOnError),
		cmdComment: "计算线径：",
	}
	LCmd.Int("c", 1, "CPU shares (relative weight)")
	LCmd.String("name", "", "Assign a name to the container")

	// 重量计算
	MCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("M", flag.ExitOnError),
		cmdComment: "计算重量：",
	}
	MCmd.Int("c", 1, "CPU shares (relative weight)")
	MCmd.String("name", "", "Assign a name to the container")

	// 用 map 保存所有的二级子命令，方便快速查找
	subcommands := map[string]*MyFlagSet{
		TCmd.Name(): TCmd,
		RCmd.Name(): RCmd,
		LCmd.Name(): LCmd,
		MCmd.Name(): MCmd,
	}

	useage := func() { // 整个命令行的帮助信息
		fmt.Printf("使用说明: \n\n")
		for _, v := range subcommands {
			fmt.Printf("%s %s\n", v.Name(), v.cmdComment)
			v.PrintDefaults() // 使用 flag 库自带的格式输出子命令的选项帮助信息
			fmt.Println()
		}
		os.Exit(2)
	}

	if len(os.Args) < 2 { // 即没有输入子命令
		useage()
	}

	cmd := subcommands[os.Args[1]] // 第二个参数必须是我们支持的子命令
	if cmd == nil {
		useage()
	}

	cmd.Parse(os.Args[2:]) // 注意这里是 cmd.Parse 不是 flag.Parse，且值是 Args[2:]

	// 输出解析后的结果
	fmt.Println("command name is:", cmd.Name())
	cmd.Visit(func(f *flag.Flag) {
		fmt.Printf("option %s, value is %s\n", f.Name, f.Value)
	})
}
