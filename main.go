package main

import(
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main(){
	app := &cli.NewApp{
		Name:"HealthChecker",
		Usage:"A tiny tool that checks whether a website is running or down",
		Flags:cli.Flag{
			&cli.StringFlag{
				Name:"domain",
				Aliases:[]string{"d"},
				Usage:"The domain to check",
				Required:true,
			},
			&cli.StringFlag{
				Name:"port",
				Aliases:[]string{"p"},
				Usage:"The port to check",
				Required:false,
			},
		},
		Action:func(c *cli.Context) error{
			port := c.String("port")
			if c.String("port") == ""{
				port = "80"
			}
			status := Check(c.String("domain"),port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}