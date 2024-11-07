// json 데이터를 반환하는 http 웹 서버
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   // student 구조체를 json 형식으로 변환한다.
	w.Header().Add("content-type", "application/json") // json 형식으로 응답한다.
	w.WriteHeader(http.StatusOK)                       // 상태 코드를 200으로 설정한다.
	fmt.Fprint(w, string(data))                        // json 형식으로 변환된 데이터를 응답 바디에 쓴다.
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", StudentsHandler)
	return mux
}

func main() {
	http.ListenAndServe(":8080", MakeWebHandler())
}
