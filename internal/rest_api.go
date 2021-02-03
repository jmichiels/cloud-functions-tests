package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type restApi struct {
	srv *service
}

// NewServer creates a new server for REST API requests.
func newRestApi(srv *service) *restApi {
	return &restApi{
		srv: srv,
	}
}

// registerRoutes registers the REST API routes in the provided gin.IRouter (a gin.Engine or gin.RouterGroup).
func (api *restApi) registerRoutes(router gin.IRouter) {
	accountsGroup := router.Group("/accounts")
	{
		// GET "/accounts"
		accountsGroup.GET("/", func(c *gin.Context) {
			accounts, err := api.srv.getAllAccounts()
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			dto := make([]*AccountDto, len(accounts))
			for i, account := range accounts {
				dto[i] = newAccountDtoFromDomain(account)
			}
			c.JSON(http.StatusOK, dto)
		})

		// POST "/accounts"
		accountsGroup.POST("/", func(c *gin.Context) {
			var dto AccountDto
			if err := c.BindJSON(&dto); err != nil {
				return
			}
			account, err := dto.ToDomain()
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			if err := api.srv.storeAccount(account); err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.Status(http.StatusOK)
		})

		//})
		//accountsGroup.GET("/:id", func(c *gin.Context) {
		//	//accountId := c.Param("id")
		//
		//})
	}
	clientsGroup := router.Group("/clients")
	{
		// GET "/clients"
		clientsGroup.GET("/", func(c *gin.Context) {
			clients, err := api.srv.getAllClients()
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			dto := make([]*ClientDto, len(clients))
			for i, client := range clients {
				dto[i] = newClientDtoFromDomain(client)
			}
			c.JSON(http.StatusOK, dto)
		})

		// POST "/clients"
		clientsGroup.POST("/", func(c *gin.Context) {
			var dto ClientDto
			if err := c.BindJSON(&dto); err != nil {
				return
			}
			client, err := dto.ToDomain()
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			if err := api.srv.storeClient(client); err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.Status(http.StatusOK)
		})
	}
}
