package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/cbrgm/clickbaiter/clickbaiter"
)

var hashtagsFlag = flag.Bool("hashtags", false, "use hashtags")

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	b := *hashtagsFlag

	cbg := clickbaiter.NewGenerator()
	cbg.UseHashtags(b)

	fmt.Println(cbg.RandomSentence())
}
