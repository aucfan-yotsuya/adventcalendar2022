package main

import (
	"log"
	"math"
	"os"

	gowav "github.com/youpy/go-wav"
)

type (
	Score struct {
		K float64 // 鍵
		V float64 // 強さ
		L float64 // 音符
	}
)

// ピアノ88鍵盤
const (
	k06  float64 = 27.500
	k06s float64 = 29.135
	k07  float64 = 30.868
	k11  float64 = 32.703
	k11s float64 = 34.648
	k12  float64 = 36.708
	k12s float64 = 38.891
	k13  float64 = 41.203
	k14  float64 = 43.654
	k14s float64 = 46.249
	k15  float64 = 48.999
	k15s float64 = 51.913
	k16  float64 = 55.000
	k16s float64 = 58.270
	k17  float64 = 61.735
	k21  float64 = 65.406
	k21s float64 = 69.296
	k22  float64 = 73.416
	k22s float64 = 77.782
	k23  float64 = 82.407
	k24  float64 = 87.307
	k24s float64 = 92.499
	k25  float64 = 97.999
	k25s float64 = 103.826
	k26  float64 = 110.000
	k26s float64 = 116.541
	k27  float64 = 123.471
	k31  float64 = 130.813
	k31s float64 = 138.591
	k32  float64 = 146.832
	k32s float64 = 155.563
	k33  float64 = 164.814
	k34  float64 = 174.614
	k34s float64 = 184.997
	k35  float64 = 195.998
	k35s float64 = 207.652
	k36  float64 = 220.000
	k36s float64 = 233.082
	k37  float64 = 246.942
	k41  float64 = 261.626
	k41s float64 = 277.183
	k42  float64 = 293.665
	k42s float64 = 311.127
	k43  float64 = 329.628
	k44  float64 = 349.228
	k44s float64 = 369.994
	k45  float64 = 391.995
	k45s float64 = 415.305
	k46  float64 = 440.000
	k46s float64 = 466.164
	k47  float64 = 493.883
	k51  float64 = 523.251
	k51s float64 = 554.365
	k52  float64 = 587.330
	k52s float64 = 622.254
	k53  float64 = 659.255
	k54  float64 = 698.456
	k54s float64 = 739.989
	k55  float64 = 783.991
	k55s float64 = 830.609
	k56  float64 = 880.000
	k56s float64 = 932.328
	k57  float64 = 987.767
	k61  float64 = 1046.502
	k61s float64 = 1108.731
	k62  float64 = 1174.659
	k62s float64 = 1244.508
	k63  float64 = 1318.510
	k64  float64 = 1396.913
	k64s float64 = 1479.978
	k65  float64 = 1567.982
	k65s float64 = 1661.219
	k66  float64 = 1760.000
	k66s float64 = 1864.655
	k67  float64 = 1975.533
	k71  float64 = 2093.005
	k71s float64 = 2217.461
	k72  float64 = 2349.318
	k72s float64 = 2489.016
	k73  float64 = 2637.020
	k74  float64 = 2793.826
	k74s float64 = 2959.955
	k75  float64 = 3135.963
	k75s float64 = 3322.438
	k76  float64 = 3520.000
	k76s float64 = 3729.310
	k77  float64 = 3951.066
	k81  float64 = 4186.009
)

// 音の強さ
const (
	F  = 0.3   // フォルテ
	MF = F / 2 // メゾフォルテ
	N  = F / 5 // 普通
	MP = N / 2 // メゾピアノ
	P  = N / 5 // ピアノ
	M  = 0.000 // 無音
)

// サンプリングレート
const Rate = 48000

// テンポ
const Tempo = 160

// 音符の種類
const (
	L4   = float64(Rate * 60 / Tempo) // 4分音符
	L1   = L2 * 2                     // 全音符
	L2   = L4 * 2                     // 2分音符
	L2H  = L2 + L4                    // 符点2分音符
	L8   = L4 / 2                     // 8分音符
	L4H  = L4 + L8                    // 符点4分音符
	L16  = L8 / 2                     // 16分音符
	L8H  = L8 + L16                   // 符点8分音符
	L32  = L16 / 2                    // 32分音符
	L16H = L16 + L32                  // 符点16分音符
)

const (
	MAX = 120000
	MIN = -40000
)

func main() {
	if err := Run(); err != nil {
		log.Fatalln(err)
	}
}
func Run() error {
	//xmas := Sin(Xmas())
	//xmasM := Sin(XmasM())
	//xmasL := Sin(XmasL())
	//Mix(&xmasM, &xmasL)
	//Mix(&xmas, &xmasM)
	melody := Sin(Aquarion())
	melodyM := Sin(AquarionM())
	melodyL := Sin(AquarionL())
	Mix(&melody, &melodyL)
	Mix(&melody, &melodyM)
	if err := Write(melody); err != nil {
		return err
	}
	return nil
}
func Max(a, b int) int {
	if a+b > MAX {
		return MAX
	} else {
		return a + b
	}
}
func Min(a, b int) int {
	if a+b < MIN {
		return MIN
	} else {
		return a + b
	}
}
func Sin(score []Score) []gowav.Sample {
	var samples []gowav.Sample
	for _, s := range score {
		for i := 0; i < int(s.L); i++ {
			si := math.Sin((float64(i) / Rate) * s.K * 2.0 * math.Pi)
			gw := gowav.Sample{}
			gw.Values[0] = int(si * 0x7fff * s.V)
			samples = append(samples, gw)
		}
	}
	// 余白
	for i := 0; i < Rate; i++ {
		samples = append(samples, gowav.Sample{})
	}
	return samples
}
func Mix(wav1, wav2 *[]gowav.Sample) {
	for i, _ := range *wav1 {
		if len(*wav2) < i {
			break
		}
		if (*wav1)[i].Values[0] < 0 {
			(*wav1)[i].Values[0] = Min((*wav1)[i].Values[0], (*wav2)[i].Values[0])
		} else {
			(*wav1)[i].Values[0] = Max((*wav1)[i].Values[0], (*wav2)[i].Values[0])
		}
	}
}
func Write(samples []gowav.Sample) error {
	fp, err := os.OpenFile("a.wav", os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	writer := gowav.NewWriter(fp, uint32(len(samples)), 2, Rate, 16)
	writer.WriteSamples(samples)
	return nil
}
