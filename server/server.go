package server

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v3"

	"github.com/maokomioko/go-fortnox"
)

type APIServer struct {
	s *httptest.Server
}

// NewTestAPIServer creates a Test API Server
//
// don't forget to call Close method
func NewTestAPIServer() *APIServer {
	return &APIServer{
		s: httptest.NewServer(
			addRoutes(),
		),
	}

}

// Close closes internal http server
func (s *APIServer) Close() {
	s.s.Close()
}

// GetURL returns url like http://localhost:randomPortNumber
func (s *APIServer) GetURL() string {
	return s.s.URL
}

// GetHost returns the actual address for using without the schema
func (s *APIServer) GetHost() string {
	return strings.Split(s.s.URL, "//")[1]
}

func addRoutes() http.Handler {
	r := gin.Default()

	addInvoices(r)
	addCustomers(r)

	return r
}

func addInvoices(r *gin.Engine) {
	invoices := []fortnox.Invoice{}

	r.GET("/invoices/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.PUT("/invoices/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := UpdateInvoiceReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrorInformation": gin.H{
					"error":   http.StatusBadRequest,
					"message": err.Error(),
					"code":    -100,
				},
			})
			return
		}

		for i := range invoices {
			if invoices[i].DocumentNumber == id {
				invoices[i] = p.Invoice
				c.JSON(http.StatusOK, gin.H{
					"Invoice": p.Invoice,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.GET("/invoices", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Invoices": invoices,
		})
	})
	r.POST("/invoices", func(c *gin.Context) {
		p := CreateInvoiceReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrorInformation": gin.H{
					"error":   http.StatusBadRequest,
					"message": err.Error(),
					"code":    -100,
				},
			})
			return
		}
		invoices = append(invoices, p.Invoice)
		c.JSON(http.StatusOK, gin.H{
			"Invoice": p.Invoice,
		})
	})
	r.PUT("/invoices/:id/externalprint", func(c *gin.Context) {
		id := c.Param("id")
		for i := range invoices {
			if invoices[i].DocumentNumber == id {
				invoices[i].Sent = null.BoolFrom(true)
				c.JSON(http.StatusOK, gin.H{
					"Invoice": invoices[i],
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.PUT("/invoices/:id/bookkeep", func(c *gin.Context) {
		id := c.Param("id")
		for i := range invoices {
			if invoices[i].DocumentNumber == id {
				invoices[i].Booked = null.BoolFrom(true)
				c.JSON(http.StatusOK, gin.H{
					"Invoice": invoices[i],
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.PUT("/invoices/:id/cancel", func(c *gin.Context) {
		id := c.Param("id")
		for i := range invoices {
			if invoices[i].DocumentNumber == id {
				invoices[i].Cancelled = null.BoolFrom(true)
				c.JSON(http.StatusOK, gin.H{
					"Invoice": invoices[i],
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.PUT("/invoices/:id/credit", func(c *gin.Context) {
		id := c.Param("id")
		for i := range invoices {
			if invoices[i].DocumentNumber == id {
				invoices[i].Credit = "1"
				invoices[i].CreditInvoiceReference = fortnox.IntIsh(2)
				c.JSON(http.StatusOK, gin.H{
					"Invoice": invoices[i],
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.PUT("/invoices/:id/warehouseready", func(c *gin.Context) {
		id := c.Param("id")
		for i := range invoices {
			if invoices[i].DocumentNumber == id {
				invoices[i].NotCompleted = false
				c.JSON(http.StatusOK, gin.H{
					"Invoice": invoices[i],
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	// TODO: do we need this one?
	r.PUT("/invoices/:id/print", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				pdfBts := []byte(v.WayOfDelivery)
				c.String(http.StatusOK, "%s", pdfBts)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.GET("/invoices/:id/einvoice", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.GET("/invoices/:id/eprint", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
}

func addCustomers(r *gin.Engine) {
	customers := []fortnox.Customer{}

	r.GET("/customers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Customers": customers,
		})
	})
	r.POST("/customers", func(c *gin.Context) {
		p := CreateCustomerReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrorInformation": gin.H{
					"error":   http.StatusBadRequest,
					"message": err.Error(),
					"code":    -100,
				},
			})
			return
		}
		customers = append(customers, p.Customer)
		c.JSON(http.StatusOK, gin.H{
			"Customer": p.Customer,
		})
	})
	r.GET("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range customers {
			if v.CustomerNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Customer": v,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.PUT("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := UpdateCustomerReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrorInformation": gin.H{
					"error":   http.StatusBadRequest,
					"message": err.Error(),
					"code":    -100,
				},
			})
			return
		}

		for i := range customers {
			if customers[i].CustomerNumber == id {
				customers[i] = p.Customer
				c.JSON(http.StatusOK, gin.H{
					"Customer": p.Customer,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
	r.DELETE("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, v := range customers {
			if v.CustomerNumber == id {
				customers = append(customers[:i], customers[i+1:]...)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"ErrorInformation": gin.H{
				"error":   http.StatusNotFound,
				"message": "Kan inte hitta kunden.",
				"code":    2000433,
			},
		})
	})
}

// Invoices

type GetInvoiceResp struct {
	Invoice fortnox.Invoice `json:"Invoice"`
}

type UpdateInvoiceReq struct {
	Invoice fortnox.Invoice `json:"Invoice"`
}

type UpdateInvoiceResp struct {
	Invoice fortnox.Invoice `json:"Invoice"`
}

type CreateInvoiceReq struct {
	Invoice fortnox.Invoice `json:"Invoice"`
}

type CreateInvoiceResp struct {
	Invoice fortnox.Invoice `json:"Invoice"`
}

// Customers

type CreateCustomerReq struct {
	Customer fortnox.Customer `json:"Customer"`
}

type CreateCustomerResp struct {
	Customer fortnox.Customer `json:"Customer"`
}

type GetCustomerResp struct {
	Customer fortnox.Customer `json:"Customer"`
}

type UpdateCustomerReq struct {
	Customer fortnox.Customer `json:"Customer"`
}

type UpdateCustomerResp struct {
	Customer fortnox.Customer `json:"Customer"`
}
