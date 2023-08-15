package algoritm

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Data struct {
	Numb int
}

func GetData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err.Error())
	}

	var data Data

	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		logrus.Error(err.Error())
	}

	w.Write([]byte(algoritm(data.Numb)))
}

func algoritm(numb1 int) string {
	b := []int{}

	for numb1 > 0 {
		numb := numb1 % 10
		b = append(b, numb)
		numb1 /= 10
	}

	for i := 0; i < len(b); i++ {
		if i%2 != 0 {
			b[i] *= 2
			if b[i] > 9 {
				b[i] -= 9
			}
		}
	}

	sum := 0
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}

	if sum%10 == 0 {
		return "validate"
	} else {
		return "non validate"
	}
}
