// package middlewares

// import (
// 	"snap-bank-statement/usecase"
// 	"github.com/Saucon/errcntrct"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// type ErrorHandler struct {
// 	ErrorHandlerUsecase usecase.ErrorHandlerUsecase
// }

// type ResponseErrSnap struct {
// 	ResponseCode    string `json:"responseCode"`
// 	ResponseMessage string `json:"responseMessage"`
// }

// func NewErrorHandler(r *gin.RouterGroup, ehus usecase.ErrorHandlerUsecase) {
// 	handler := &ErrorHandler{
// 		ErrorHandlerUsecase: ehus,
// 	}

// 	r.Use(handler.errorHandler)
// }

// func (eh *ErrorHandler) errorHandler(c *gin.Context) {
// 	c.Next()

// 	errorToPrint := c.Errors.Last()
// 	if errorToPrint != nil {

// 		_, v := eh.ErrorHandlerUsecase.ResponseError(errorToPrint)
// 		s := v.(errcntrct.ErrorData)

// 		var resp ResponseErrSnap
// 		switch s.Code {
// 		case "400012":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400013":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400014":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400015":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400016":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400017":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400018":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400019":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400020":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400021":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400022":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400023":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		case "400024":
// 			resp.ResponseCode = "400012"
// 			resp.ResponseMessage = s.Msg
// 			c.JSON(http.StatusBadRequest, resp)
// 			c.Abort()
// 		}

// 		//c.JSON(eh.ErrorHandlerUsecase.ResponseError(errorToPrint))
// 		//c.Abort()
// 	}
// 	return
// }
