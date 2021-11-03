package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mauwahid/kafman/cmd/http"
	"github.com/mauwahid/kafman/cmd/queue"
	"github.com/mauwahid/kafman/internal/platform/config"
)

const banner = `
.##....##....###....########.##.....##....###....##....##
.##...##....##.##...##.......###...###...##.##...###...##
.##..##....##...##..##.......####.####..##...##..####..##
.#####....##.....##.######...##.###.##.##.....##.##.##.##
.##..##...#########.##.......##.....##.#########.##..####
.##...##..##.....##.##.......##.....##.##.....##.##...###
.##....##.##.....##.##.......##.....##.##.....##.##....##

		version %v
`

var version = "1.0"

func main() {

	var (
		cfg *viper.Viper
		err error
	)

	var rootCmd = &cobra.Command{
		Use:   "kafman",
		Short: "kafman is http kafka proxy",
		Run:   cobraRunner,
	}

	if cfg, err = createViper("config.json"); err != nil {
		panic("config not found")
	}

	config.InjectConfig(cfg)

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf(banner, version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createViper(file string) (v *viper.Viper, err error) {

	v = viper.New()
	v.SetConfigFile(file)
	if err = v.ReadInConfig(); err != nil {
		return
	}

	return
}

func cobraRunner(cobraCmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("run publisher and subscriber")
		queue.RunSubscriber()
		http.RunHttp()
		return
	}

	arg := args[0]

	switch arg {
	case "publisher":
		fmt.Println("run http only")
		http.RunHttp()
	case "subscriber":
		fmt.Println("run subscriber only")
		queue.RunSubscriber()
	default:
		fmt.Println("run publisher and subscriber")
		queue.RunSubscriber()
		http.RunHttp()
	}

}
