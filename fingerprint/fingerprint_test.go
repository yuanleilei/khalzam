package fingerprint

import (
	// "math"
	"encoding/binary"
	"bufio"
	"bytes"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"io"
	"os"
	"testing"
)

// func TestDecodeMpg123(t *testing.T) {
// 	Decode("../resources/Metallica - Master Of Puppets.mp3")
// }

func TestDecodeGo(t *testing.T) {
	filename := "../resources/modjo_lady_sample.mp3"
	file, _ := os.Open(filename)
	defer file.Close()
	d, _ := mp3.NewDecoder(file)
	defer d.Close()

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	written, _ := io.Copy(w, d)
	fmt.Println(written)
	w.Flush()
	r := bufio.NewReader(&buf)

	tmp := make([]int16, written / 2)
	var rawData []int16
	binary.Read(r, binary.LittleEndian, tmp)
	for i := 0; i < len(tmp); i += 2 {
		x := (tmp[i] + tmp[i+1]) / 2
		fmt.Println(x)
		rawData = append(rawData, (x))
	}

	f, _ := os.Create(filename + "raw.raw")
	binary.Write(f, binary.LittleEndian, rawData)
	f.Close()
}

// func TestConvertStereoToMono(t *testing.T) {
// 	StereoToMono("/Users/chingachgook/dev/gocode/src/github.com/kisasexypantera94/khalzam/resources/journeydontstop.mp3")
// }
