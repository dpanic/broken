package main

import (
	"fmt"
	"os"

	_ "github.com/anthonynsimon/bild/imgio"
	_ "github.com/fatih/color"
	_ "github.com/gabriel-vasile/mimetype"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-errors/errors"
	_ "github.com/go-redis/redis/v8"
	_ "github.com/golang-jwt/jwt"
	_ "github.com/google/uuid"
	_ "github.com/leekchan/accounting"
	_ "github.com/nyaruka/phonenumbers"
	_ "github.com/patrickmn/go-cache"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	_ "github.com/teris-io/shortid"
	_ "go.uber.org/automaxprocs"
	_ "go.uber.org/zap"
	_ "go.uber.org/zap/buffer"
	_ "go.uber.org/zap/zapcore"
	_ "golang.org/x/text"
	_ "golang.org/x/text/cases"
	_ "golang.org/x/text/language"
	_ "golang.org/x/xerrors"

	"broken/io"
	"broken/resolution"
)

func main() {
	fmt.Println("1")

	stat, _ := os.Stat("/")
	io.FileInfo(stat)
	resolution.GetPrimary()
}
