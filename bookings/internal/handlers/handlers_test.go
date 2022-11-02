package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/yazuka000/bookings/internal/driver"
	"github.com/yazuka000/bookings/internal/models"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomId: 1,
		Room: models.Room{
			Id:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code : got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code : got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test with non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomId = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code : got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}
}

func TestRepository_PostReservation(t *testing.T) {
	// reqBody := "start_date=2050-01-01"
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	postedData := url.Values{}
	postedData.Add("start_date", "2050-01-01")
	postedData.Add("end_date", "2050-01-02")
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "555-555-5555")
	postedData.Add("room_id", "1")

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code : got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for missing post body
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code : got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid start date
	postedData = url.Values{}
	postedData.Add("start_date", "invalid")
	postedData.Add("end_date", "2050-01-02")
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "1234567890")
	postedData.Add("room_id", "1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid start date: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid end date
	postedData = url.Values{}
	postedData.Add("start_date", "2050-01-01")
	postedData.Add("end_date", "invalid")
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "1234567890")
	postedData.Add("room_id", "1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid end date: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid room id
	postedData = url.Values{}
	postedData.Add("start_date", "2050-01-01")
	postedData.Add("end_date", "2050-01-02")
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "1234567890")
	postedData.Add("room_id", "invalid")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid room id: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid data
	postedData = url.Values{}
	postedData.Add("start_date", "2050-01-01")
	postedData.Add("end_date", "2050-01-02")
	postedData.Add("first_name", "J")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "1234567890")
	postedData.Add("room_id", "1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code for invalid data: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for failure to insert reservation into database
	postedData = url.Values{}
	postedData.Add("start_date", "2050-01-01")
	postedData.Add("end_date", "2050-01-02")
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "1234567890")
	postedData.Add("room_id", "2")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for failure to insert restriction into database
	postedData = url.Values{}
	postedData.Add("start_date", "2050-01-01")
	postedData.Add("end_date", "2050-01-02")
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "1234567890")
	postedData.Add("room_id", "1000")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestNewRepo(t *testing.T) {
	var db driver.DB
	testRepo := NewRepo(&app, &db)

	if reflect.TypeOf(testRepo).String() != "*handlers.Repository" {
		t.Errorf("Did not get correct type from NewRepo: got %s, wanted *Repository", reflect.TypeOf(testRepo).String())
	}
}

func TestRepository_PostAvailability(t *testing.T) {
	// first case -- rooms are not available

	// create our request body
	postedData := url.Values{}
	postedData.Add("start", "2050-01-01")
	postedData.Add("end", "2050-01-02")

	// create request
	req, _ := http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make handler handlerfunc
	handler := http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr := httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post Availability when no rooms available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// second case -- rooms are available

	// this time, we specify a start date before 2040-01-01, which will give us
	// a non-empty slice, indicating that rooms are available
	postedData = url.Values{}
	postedData.Add("start", "2040-01-01")
	postedData.Add("end", "2040-01-02")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusOK {
		t.Errorf("Post Availability when rooms are available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// third case -- empty post body

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", nil)

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability with empty request body (nil) gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// fourth case -- start date in wrong format

	postedData = url.Values{}
	postedData.Add("start", "invalid")
	postedData.Add("end", "2040-01-02")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability with invalid start date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// fifth case -- end date in wrong format

	postedData = url.Values{}
	postedData.Add("start", "2040-01-01")
	postedData.Add("end", "invalid")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability with invalid end date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// sixth case -- database query fails
	// this time, we specify a start date of 2060-01-01, which will cause
	// our testdb repo to return an error

	postedData = url.Values{}
	postedData.Add("start", "2060-01-01")
	postedData.Add("end", "2060-01-02")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability when database query fails gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_AvailabilityJson(t *testing.T) {
	// first case - rooms are not available

	// create our request body
	postedData := url.Values{}
	postedData.Add("start", "2050-01-01")
	postedData.Add("end", "2050-01-02")
	postedData.Add("room_id", "1")

	// create request
	req, _ := http.NewRequest("POST", "/search-availability-json", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "x-www-urlencoded")

	// make handler handlerfunc
	handler := http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr := httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	var j jsonResponse
	err := json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	// first case -- rooms are not available

	// create our request body
	postedData := url.Values{}
	postedData.Add("start", "2050-01-01")
	postedData.Add("end", "2050-01-02")

	// create request
	req, _ := http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make handler handlerfunc
	handler := http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr := httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post Availability when no rooms available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// second case -- rooms are available

	// this time, we specify a start date before 2040-01-01, which will give us
	// a non-empty slice, indicating that rooms are available
	postedData = url.Values{}
	postedData.Add("start", "2040-01-01")
	postedData.Add("end", "2040-01-02")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusOK {
		t.Errorf("Post Availability when rooms are available gave wrong status code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// third case -- empty post body

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", nil)

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability with empty request body (nil) gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// fourth case -- start date in wrong format

	postedData = url.Values{}
	postedData.Add("start", "invalid")
	postedData.Add("end", "2040-01-02")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability with invalid start date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// fifth case -- end date in wrong format

	postedData = url.Values{}
	postedData.Add("start", "2040-01-01")
	postedData.Add("end", "invalid")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability with invalid end date gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// sixth case -- database query fails
	// this time, we specify a start date of 2060-01-01, which will cause
	// our testdb repo to return an error

	postedData = url.Values{}
	postedData.Add("start", "2060-01-01")
	postedData.Add("end", "2060-01-02")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set request header
	req.Header.Set("Content-Type", "application/x-www-urlencoded")

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.PostAvailability)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	// since we have no rooms available, we expect to get status http.StatusSeeOther
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Post Availability when database query fails gave wrong status code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
