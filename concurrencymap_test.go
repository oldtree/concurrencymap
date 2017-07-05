package ConcurrencyMap

import (
	"math/rand"
	"testing"
)

var length = len(answers)
var r *rand.Rand
var answers = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

var lm = &ConcurrencyMapWithLock{
	data: make(map[interface{}]interface{}),
}

var lr = &ConcurrencyMapWithRWLock{
	data: make(map[interface{}]interface{}),
}

var lc = &ConcurrencyMapWithChan{
	data: make(map[interface{}]interface{}),
}

func init() {
	go lc.OperationLoop()
	r = rand.New(rand.NewSource(99))
}

func TestConcurrencyMapWithLock_Add(t *testing.T) {
	lm.Add("one", 1)
	lm.Add("two", 2)
	lm.Add("three", 3)
	lm.Add("four", 4)
}

func TestConcurrencyMapWithLock_Get(t *testing.T) {

	if 1 != lm.Get("one") {
		t.Fail()
	}
	if 2 != lm.Get("two") {
		t.Fail()
	}
	if 3 != lm.Get("three") {
		t.Fail()
	}
	if 4 != lm.Get("four") {
		t.Fail()
	}
}
func TestConcurrencyMapWithLock_Remove(t *testing.T) {
	lm.Remove("one")
	lm.Remove("two")
	lm.Remove("three")
	lm.Remove("four")
}

func Benchmark_ConcurrencyMapWithLock_Add(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		lm.Add(num, answers[num%length])
	}
}

func Benchmark_ConcurrencyMapWithLock_Get(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		if lm.Get(num) != answers[num%length] {
			tb.Failed()
		}
	}
}

func Benchmark_ConcurrencyMapWithLock_Remove(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		lm.Remove(num)
	}
}

func TestConcurrencyMapWithChan_Add(t *testing.T) {
	lc.Add("one", 1)
	lc.Add("two", 2)
	lc.Add("three", 3)
	lc.Add("four", 4)

}

func TestConcurrencyMapWithChan_Get(t *testing.T) {
	if 1 != lc.Get("one") {
		t.Fail()
	}
	if 2 != lc.Get("two") {
		t.Fail()
	}
	if 3 != lc.Get("three") {
		t.Fail()
	}
	if 4 != lc.Get("four") {
		t.Fail()
	}
}

func TestConcurrencyMapWithChan_Remove(t *testing.T) {
	lc.Remove("one")
	lc.Remove("two")
	lc.Remove("three")
	lc.Remove("four")
}

func Benchmark_ConcurrencyMapWithChan_Add(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		lc.Add(num, answers[num%length])
	}
}

func Benchmark_ConcurrencyMapWithChan_Get(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		if lc.Get(num) != answers[num%length] {
			tb.Failed()
		}
	}
}

func Benchmark_ConcurrencyMapWithChan_Remove(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		lc.Remove(num)
	}
}

func TestConcurrencyMapWithRWLock_Add(t *testing.T) {
	lr.Add("one", 1)
	lr.Add("two", 2)
	lr.Add("three", 3)
	lr.Add("four", 4)
}

func TestConcurrencyMapWithRWLock_Get(t *testing.T) {
	if 1 != lr.Get("one") {
		t.Fail()
	}
	if 2 != lr.Get("two") {
		t.Fail()
	}
	if 3 != lr.Get("three") {
		t.Fail()
	}
	if 4 != lr.Get("four") {
		t.Fail()
	}
}

func TestConcurrencyMapWithRWLock_Remove(t *testing.T) {
	lr.Remove("one")
	lr.Remove("two")
	lr.Remove("three")
	lr.Remove("four")
}

func Benchmark_ConcurrencyMapWithRWLock_Add(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		lr.Add(num, answers[num%length])
	}
}

func Benchmark_ConcurrencyMapWithRWLock_Get(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		if lr.Get(num) != answers[num%length] {
			tb.Failed()
		}
	}
}

func Benchmark_ConcurrencyMapWithRWLock_Remove(tb *testing.B) {
	for index := 0; index < tb.N; index++ {
		num := r.Int()
		lr.Remove(num)
	}
}
