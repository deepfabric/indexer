package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "net/http/pprof"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/indexer/wal"
)

var (
	pprof = flag.String("addr-pprof", "", "pprof http server address")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	if "" != *pprof {
		log.Printf("bootstrap: start pprof at: %s", *pprof)
		go func() {
			log.Fatalf("bootstrap: start pprof failed, errors:\n%+v",
				http.ListenAndServe(*pprof, nil))
		}()
	}

	// 记录时间
	t0 := time.Now()

	S := 100000
	benchmarkWriteEntry(S, 100, 8)

	// record time, and calculate performance
	t1 := time.Now()
	log.Printf("duration %v", t1.Sub(t0))
	log.Printf("wal write speed %f entries/s", float64(S)/t1.Sub(t0).Seconds())

}

func benchmarkWriteEntry(loops int, size int, batch int) {
	p, err := ioutil.TempDir(os.TempDir(), "waltest")
	if err != nil {
		log.Fatalf("err = %v, want nil", err)
	}
	defer os.RemoveAll(p)

	w, err := wal.Create(p)
	if err != nil {
		log.Fatalf("err = %v, want nil", err)
	}
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(i)
	}
	e := &raftpb.Entry{Data: data}

	for i := 0; i < loops; i++ {
		err := w.SaveEntry(e)
		if err != nil {
			log.Fatalf("err = %v, want nil", err)
		}
	}
}
