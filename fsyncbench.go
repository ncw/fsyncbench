// fsync benchmark
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	pN  = flag.Int("n", 100, "Iterations to test - default 100")
	dir = flag.String("dir", "", "Target directory - default os.TmpDir()")
)

func main() {
	flag.Parse()
	out, err := ioutil.TempFile(*dir, "fsyncbench")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = out.Close()
		if err != nil {
			log.Print(err)
		}
		err = os.Remove(out.Name())
		if err != nil {
			log.Print(err)
		}
	}()

	buf := []byte{'A'}
	N := *pN
	duration := time.Duration(0)
	for i := 0; i < N; i++ {
		_, err = out.Write(buf)
		if err != nil {
			log.Fatal(err)
		}
		start := time.Now()
		err = out.Sync()
		end := time.Now()
		duration += end.Sub(start)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("That took %s for %d fsyncs", duration, N)
	log.Printf("That took %s per fsync", duration/time.Duration(N))
}
