package TestGolangLog

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	"github.com/icrowley/fake"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
)

var (
	ProjectId = os.Getenv("GCP_PROJECT")
)

type AppError struct {
	Err  string
	Code int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error code: %d, error: %s", e.Code, e.Err)
}

type User struct {
	Name, Email string
}

func TestGolangLog(w http.ResponseWriter, r *http.Request) {
	err := errors.New("custom error using New function")
	fmt.Printf("%v", err)

	err = &AppError{Err: "custom error using struct type and fields", Code: 100}
	fmt.Printf("%v", err)

	user := User{
		Name:  fake.UserName(),
		Email: fake.EmailAddress(),
	}
	fmt.Printf("print struct type: %#v", user)

	if _, err := fmt.Fprint(w, "ok"); err != nil {
		http.Error(w, "fmt.Fprint", http.StatusInternalServerError)
	}
}

func TestGolangLogWithNewlineSymbol(w http.ResponseWriter, r *http.Request) {
	err := errors.New("custom error using New function")
	fmt.Printf("%v\n", err)

	err = &AppError{Err: "custom error using struct type and fields", Code: 100}
	fmt.Printf("%v\n", err)

	user := User{
		Name:  fake.UserName(),
		Email: fake.EmailAddress(),
	}
	fmt.Printf("print struct type: %#v\n", user)

	if _, err := fmt.Fprint(w, "ok"); err != nil {
		http.Error(w, "fmt.Fprint", http.StatusInternalServerError)
	}
}

func TestGolangLogUsingLog(w http.ResponseWriter, r *http.Request) {
	err := errors.New("custom error using New function")
	log.Printf("%v\n", err)

	err = &AppError{Err: "custom error using struct type and fields", Code: 100}
	log.Printf("%v\n", err)

	user := User{
		Name:  fake.UserName(),
		Email: fake.EmailAddress(),
	}
	log.Printf("print struct type: %#v\n", user)

	if _, err := fmt.Fprint(w, "ok"); err != nil {
		http.Error(w, "fmt.Fprint", http.StatusInternalServerError)
	}
}

func createLogger(r *http.Request) (*logging.Logger, *logging.Client, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, ProjectId)
	if err != nil {
		return nil, nil, err
	}

	logName := "cloudfunctions.googleapis.com/cloud-functions"
	logger := client.Logger(
		logName,
		logging.CommonResource(&mrpb.MonitoredResource{
			Labels: map[string]string{
				"function_name": os.Getenv("FUNCTION_NAME"),
				"project_id":    os.Getenv("GCP_PROJECT"),
				"region":        os.Getenv("FUNCTION_REGION"),
			},
			Type: "cloud_function",
		}),
		logging.CommonLabels(map[string]string{
			"execution_id": r.Header.Get("Function-Execution-Id"),
		}),
	)
	return logger, client, nil
}

func TestGolangLogUsingCloudLogging(w http.ResponseWriter, r *http.Request) {

	logger, client, err := createLogger(r)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, "create logger error", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("closing logging client error: %#v", err)
			http.Error(w, "closing logging client error", http.StatusInternalServerError)
			return
		}
	}()

	err = errors.New("custom error using New function")
	logger.Log(logging.Entry{Payload: err.Error(), Severity: logging.Error})

	err = &AppError{Err: "custom error using struct type and fields", Code: 100}
	logger.Log(logging.Entry{Payload: err, Severity: logging.Error})

	user := User{
		Name:  fake.UserName(),
		Email: fake.EmailAddress(),
	}
	logger.Log(logging.Entry{Payload: user, Severity: logging.Debug})

	if _, err := fmt.Fprint(w, "ok"); err != nil {
		http.Error(w, "fmt.Fprint", http.StatusInternalServerError)
	}
}
