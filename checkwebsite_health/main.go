package main 

import (
    "os"
	"fmt"
	"log"
	"github.com/urfave/cli/v2"
	
)
	


func main(){
	app := &cli.App{
		Name: "Health Check Of Application",
		Usage : "A Tool To Checks The Given Domain is Down Or Up.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name to Check.",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port" ,
				Aliases: []string{"p"},
				Usage: "Port Number To Check",
				Required : false,
			},
		},
		Action : func(c *cli.Context) error{
           port := c.String("port")
		   if c.String("port") == "" {
			port = "80"
		   }
		   status := Check(c.String("domain"), port)
		   fmt.Println(status)
		   return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	
}