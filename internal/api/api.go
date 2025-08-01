package api

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/theexcelrobin/kodee-notifier/internal/email"
	"github.com/theexcelrobin/kodee-notifier/internal/telegram"
	"github.com/theexcelrobin/kodee-notifier/internal/whatsapp"
)

type Api struct {
	Address  string
	Port     string
	Router   *gin.Engine
	Email    *email.Email
	Telegram *telegram.Telegram
	Whatsapp *whatsapp.Whatsapp
}

func NewApi(e *email.Email, t *telegram.Telegram, w *whatsapp.Whatsapp) (*Api, error) {
	api := &Api{
		Address:  os.Getenv("ADDRESS"),
		Port:     os.Getenv("PORT"),
		Router:   gin.Default(),
		Email:    e,
		Telegram: t,
		Whatsapp: w,
	}

	// NOTIFY ORDER ENDPOINT
	api.Router.POST(
		"/notify/order",
		api.NotifyOrder,
	)

	// NOTIFY LISTING ENDPOINT
	api.Router.POST(
		"/notify/listing",
		api.NotifyListing,
	)

	return api, nil
}

func (a *Api) SpawnServer() error {
	if err := a.Router.Run(fmt.Sprintf("%s:%s", a.Address, a.Port)); err != nil {
		return fmt.Errorf("kodee notifier server error: %s", err.Error())
	}
	return nil
}

/*-------------------------------------------------- HANDLERS --------------------------------------------------*/

type NotifyOrderRequest struct {
	Item        string `json:"item"`
	ClientName  string `json:"name"`
	ClientEmail string `json:"email"`
	ClientPhone string `json:"phone"`
}

func (a *Api) NotifyOrder(c *gin.Context) {
	var req NotifyOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	msg := fmt.Sprintf(
		"*New Order Request*\n\n*Name*: `%s`\n\n*Item*: `%s`\n\n_© Kodee Enterprise_",
		req.Item,
		req.ClientName,
	)

	wg := sync.WaitGroup{}
	ec := make(chan error, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if a.Email != nil {
			if err := a.Email.Notify(req.ClientEmail, "New Order Request", msg); err != nil {
				ec <- err
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if a.Telegram != nil {
			if err := a.Telegram.Notify(req.ClientPhone, msg); err != nil {
				ec <- err
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if a.Whatsapp != nil {
			if err := a.Whatsapp.Notify(req.ClientPhone, msg); err != nil {
				ec <- err
			}
		}
	}()

	wg.Wait()

	c.JSON(http.StatusOK, http.StatusNoContent)
}

func (a *Api) NotifyListing(c *gin.Context) {
	//
}
