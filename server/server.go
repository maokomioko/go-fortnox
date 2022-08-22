package server

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"

	fortnox "github.com/thats4fun/go-fortnox-sdk/client"
)

type APIServer struct {
	s *httptest.Server
}

func NewTestAPIServer() *APIServer {
	return &APIServer{
		s: httptest.NewServer(
			addRoutes(),
		),
	}

}

func (s *APIServer) Close() {
	s.s.Close()
}

func (s *APIServer) GetURL() string {
	return s.s.URL
}

func addRoutes() http.Handler {
	r := gin.Default()

	addProjects(r)
	addPrintTemplates(r)
	addLabels(r)
	addCostCenters(r)
	addWayOfDeliveries(r)
	addTermsOfPayments(r)
	addTermsOfDelivery(r)
	addPriceLists(r)
	addTaxReductions(r)

	addInvoices(r)
	addCustomers(r)

	return r
}

func addProjects(r *gin.Engine) {
	projects := []fortnox.Project{}

	r.GET("/projects/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range projects {
			if v.ProjectNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Project": v,
				})
				return
			}
		}
	})
	r.PUT("/projects/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range projects {
			if v.ProjectNumber == id {
				p := fortnox.UpdateProjectReq{}
				err := c.ShouldBind(&p)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"Error": err,
					})
				}
				v = p.Project
				c.JSON(http.StatusOK, gin.H{
					"Project": v,
				})
				return
			}
		}
	})
	r.DELETE("/projects/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, v := range projects {
			if v.ProjectNumber == id {
				projects = append(projects[:i], projects[i+1:]...)
				return
			}
		}
	})
	r.GET("/projects", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Projects": projects,
		})
	})
	r.POST("/projects", func(c *gin.Context) {
		p := fortnox.CreateProjectReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		projects = append(projects, p.Project)
		c.JSON(http.StatusOK, gin.H{
			"Project": p.Project,
		})
	})
}

func addPrintTemplates(r *gin.Engine) {
	printTemplates := []fortnox.PrintTemplate{
		{
			Template: "#1",
			Name:     "someName",
		},
	}
	r.GET("/printtemplates", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"PrintTemplates": printTemplates,
		})
	})
}

func addLabels(r *gin.Engine) {
	labels := []fortnox.Label{
		{
			ID:          1,
			Description: "Cool Label",
		},
	}
	r.GET("/labels", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Labels": labels,
		})
	})
	r.POST("/labels", func(c *gin.Context) {
		req := fortnox.CreateLabelReq{}
		err := c.Bind(&req)
		fmt.Println("p:", req.Label)
		p := req.Label
		p.ID = 2
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		labels = append(labels, p)
		c.JSON(http.StatusOK, gin.H{
			"Label": p,
		})
	})
	r.PUT("/labels/:id", func(c *gin.Context) {
		id := c.Param("id")
		idI, _ := strconv.Atoi(id)
		p := fortnox.UpdateLabelReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range labels {
			if v.ID == idI {
				v = p.Label
				c.JSON(http.StatusOK, gin.H{
					"Label": p.Label,
				})
				return
			}
		}
	})
	r.DELETE("/labels/:id", func(c *gin.Context) {
		id := c.Param("id")
		idI, _ := strconv.Atoi(id)
		for i, v := range labels {
			if v.ID == idI {
				labels = append(labels[:i], labels[i+1:]...)
				return
			}
		}
	})
}

func addCostCenters(r *gin.Engine) {
	costCenter := []fortnox.CostCenter{}

	r.GET("/costcenters", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"CostCenters": costCenter,
		})
	})
	r.POST("/costcenters", func(c *gin.Context) {
		p := fortnox.CreateCostCenterReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		costCenter = append(costCenter, p.CostCenter)
		c.JSON(http.StatusOK, gin.H{
			"CostCenter": p.CostCenter,
		})
	})
	r.GET("/costcenters/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range costCenter {
			if v.Code == id {
				c.JSON(http.StatusOK, gin.H{
					"CostCenter": v,
				})
				return
			}
		}
	})
	r.PUT("/costcenters/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdateCostCenterReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range costCenter {
			if v.Code == id {
				v = p.CostCenter
				c.JSON(http.StatusOK, gin.H{
					"CostCenter": p.CostCenter,
				})
				return
			}
		}
	})
	r.DELETE("/costcenters/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, v := range costCenter {
			if v.Code == id {
				costCenter = append(costCenter[:i], costCenter[i+1:]...)
				return
			}
		}
	})
}

func addWayOfDeliveries(r *gin.Engine) {
	wayOfDeliveries := []fortnox.WayOfDelivery{}

	r.GET("/wayofdeliveries", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"WayOfDeliveries": wayOfDeliveries,
		})
	})
	r.POST("/wayofdeliveries", func(c *gin.Context) {
		p := fortnox.CreateWayOfDeliveriesReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		wayOfDeliveries = append(wayOfDeliveries, p.WayOfDelivery)
		c.JSON(http.StatusOK, gin.H{
			"WayOfDelivery": p.WayOfDelivery,
		})
	})
	r.GET("/wayofdeliveries/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range wayOfDeliveries {
			if v.Code == id {
				c.JSON(http.StatusOK, gin.H{
					"WayOfDelivery": v,
				})
				return
			}
		}
	})
	r.PUT("/wayofdeliveries/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdateWayOfDeliveryReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range wayOfDeliveries {
			if v.Code == id {
				v = p.WayOfDelivery
				c.JSON(http.StatusOK, gin.H{
					"WayOfDelivery": p.WayOfDelivery,
				})
				return
			}
		}
	})
	r.DELETE("/wayofdeliveries/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, v := range wayOfDeliveries {
			if v.Code == id {
				wayOfDeliveries = append(wayOfDeliveries[:i], wayOfDeliveries[i+1:]...)
				return
			}
		}
	})
}

func addTermsOfPayments(r *gin.Engine) {
	termsOfPayments := []fortnox.TermsOfPayment{}

	r.GET("/termsofpayments", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"TermsOfPayments": termsOfPayments,
		})
	})
	r.POST("/termsofpayments", func(c *gin.Context) {
		p := fortnox.CreateTermOfPaymentReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		termsOfPayments = append(termsOfPayments, p.TermsOfPayment)
		c.JSON(http.StatusOK, gin.H{
			"TermsOfPayment": p.TermsOfPayment,
		})
	})
	r.GET("/termsofpayments/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range termsOfPayments {
			if v.Code == id {
				c.JSON(http.StatusOK, gin.H{
					"TermsOfPayment": v,
				})
				return
			}
		}
	})
	r.PUT("/termsofpayments/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdateTermOfPaymentReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range termsOfPayments {
			if v.Code == id {
				v = p.TermsOfPayment
				c.JSON(http.StatusOK, gin.H{
					"TermsOfPayment": p.TermsOfPayment,
				})
				return
			}
		}
	})
	r.DELETE("/termsofpayments/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, v := range termsOfPayments {
			if v.Code == id {
				termsOfPayments = append(termsOfPayments[:i], termsOfPayments[i+1:]...)
				return
			}
		}
	})
}

func addTermsOfDelivery(r *gin.Engine) {
	termsOfDeliveries := []fortnox.TermsOfDelivery{}

	r.GET("/termsofdeliveries", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"TermsOfDeliveries": termsOfDeliveries,
		})
	})
	r.POST("/termsofdeliveries", func(c *gin.Context) {
		p := fortnox.CreateTermsOfDeliveriesReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		termsOfDeliveries = append(termsOfDeliveries, p.TermsOfDelivery)
		c.JSON(http.StatusOK, gin.H{
			"TermsOfDelivery": p.TermsOfDelivery,
		})
	})
	r.GET("/termsofdeliveries/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range termsOfDeliveries {
			if v.Code == id {
				c.JSON(http.StatusOK, gin.H{
					"TermsOfDelivery": v,
				})
				return
			}
		}
	})
	r.PUT("/termsofdeliveries/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdateTermOfDeliveryReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range termsOfDeliveries {
			if v.Code == id {
				v = p.TermsOfDelivery
				c.JSON(http.StatusOK, gin.H{
					"TermsOfDelivery": p.TermsOfDelivery,
				})
				return
			}
		}
	})
}

func addPriceLists(r *gin.Engine) {
	priceLists := []fortnox.PriceList{}

	r.GET("/pricelists", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"PriceLists": priceLists,
		})
	})
	r.POST("/pricelists", func(c *gin.Context) {
		p := fortnox.CreatePriceListReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		priceLists = append(priceLists, p.PriceList)
		c.JSON(http.StatusOK, gin.H{
			"PriceList": p.PriceList,
		})
	})
	r.GET("/pricelists/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range priceLists {
			if v.Code == id {
				c.JSON(http.StatusOK, gin.H{
					"PriceList": v,
				})
				return
			}
		}
	})
	r.PUT("/pricelists/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdatePriceListReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range priceLists {
			if v.Code == id {
				v = p.PriceList
				c.JSON(http.StatusOK, gin.H{
					"PriceList": p.PriceList,
				})
				return
			}
		}
	})
}

func addTaxReductions(r *gin.Engine) {
	taxReductions := []fortnox.TaxReduction{}

	r.GET("/taxreductions", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"TaxReductions": taxReductions,
		})
	})
	r.POST("/taxreductions", func(c *gin.Context) {
		p := fortnox.CreateTaxReductionReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		taxReductions = append(taxReductions, p.TaxReduction)
		c.JSON(http.StatusOK, gin.H{
			"TaxReduction": p.TaxReduction,
		})
	})
	r.GET("/taxreductions/:id", func(c *gin.Context) {
		id := c.Param("id")
		idS, _ := strconv.Atoi(id)
		for _, v := range taxReductions {
			if v.Id == idS {
				c.JSON(http.StatusOK, gin.H{
					"TaxReduction": v,
				})
				return
			}
		}
	})
	r.PUT("/taxreductions/:id", func(c *gin.Context) {
		id := c.Param("id")
		idS, _ := strconv.Atoi(id)
		p := fortnox.UpdateTaxReductionReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range taxReductions {
			if v.Id == idS {
				v = p.TaxReduction
				c.JSON(http.StatusOK, gin.H{
					"TaxReduction": p.TaxReduction,
				})
				return
			}
		}
	})
	r.DELETE("/taxreductions/:id", func(c *gin.Context) {
		id := c.Param("id")
		idS, _ := strconv.Atoi(id)
		for i, v := range taxReductions {
			if v.Id == idS {
				taxReductions = append(taxReductions[:i], taxReductions[i+1:]...)
				return
			}
		}
	})
}

func addInvoices(r *gin.Engine) {
	invoices := []fortnox.Invoice{}

	r.GET("/invoices/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Customer": v,
				})
				return
			}
		}
	})
	r.PUT("/invoices/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdateInvoiceReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		for _, v := range invoices {
			if v.DocumentNumber == id {
				v = p.Invoice
				c.JSON(http.StatusOK, gin.H{
					"Invoice": p.Invoice,
				})
				return
			}
		}
	})
	r.GET("/invoices", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Invoices": invoices,
		})
	})
	r.POST("/invoices", func(c *gin.Context) {
		p := fortnox.CreateInvoiceReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		invoices = append(invoices, p.Invoice)
		c.JSON(http.StatusOK, gin.H{
			"Invoice": p.Invoice,
		})
	})
	r.PUT("/invoices/:id/bookkeep", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				v.Booked = true
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
	})
	r.PUT("/invoices/:id/cancel", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				v.Cancelled = true
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
	})
	r.PUT("/invoices/:id/credit", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				v.Credit = "1"
				v.CreditInvoiceReference = "2"
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
	})
	r.PUT("/invoices/:id/externalprint", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				v.Sent = true
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
	})
	r.PUT("/invoices/:id/warehouseready", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				v.NotCompleted = false
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
	})
	r.PUT("/invoices/:id/print", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				pdfBts := []byte(v.WayOfDelivery)
				c.String(http.StatusOK, "%s", pdfBts)
				return
			}
		}
	})
	r.PUT("/invoices/:id/email", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				if v.EmailInformation.EmailBody != "" {
					c.JSON(http.StatusOK, gin.H{
						"Invoice": v,
					})
					return
				}
			}
		}
	})
	r.GET("/invoices/:id/printreminder", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				pdfBts := []byte(v.WayOfDelivery)
				c.String(http.StatusOK, "%s", pdfBts)
				return
			}
		}
	})
	r.GET("/invoices/:id/preview", func(c *gin.Context) {
		id := c.Param("id")
		for _, v := range invoices {
			if v.DocumentNumber == id {
				c.JSON(http.StatusOK, gin.H{
					"Invoice": v,
				})
				return
			}
		}
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
		p := fortnox.CreateCustomerReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
		}
		customers = append(customers, p.Customer)
		c.JSON(http.StatusOK, gin.H{
			"Customer": p.Customer,
		})
	})
	r.GET("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		for _, v := range customers {
			if v.CustomerNumber == int64(idInt) {
				c.JSON(http.StatusOK, gin.H{
					"Customer": v,
				})
				return
			}
		}
	})
	r.PUT("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")
		p := fortnox.UpdateCustomerReq{}
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}
		idInt, _ := strconv.Atoi(id)
		for _, v := range customers {
			if v.CustomerNumber == int64(idInt) {
				v = p.Customer
				c.JSON(http.StatusOK, gin.H{
					"Customer": p.Customer,
				})
				return
			}
		}
	})
	r.DELETE("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		for i, v := range customers {
			if v.CustomerNumber == int64(idInt) {
				customers = append(customers[:i], customers[i+1:]...)
				return
			}
		}
	})
}
