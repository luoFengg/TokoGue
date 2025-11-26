package exceptions

import (
	"net/http"
	"tokogue-api/models/web"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // jalankan handler dulu
        
        // Check jika ada error
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            
            // Handle based on error type
            switch e := err.(type) {
            case NotFoundError:
                c.JSON(http.StatusNotFound, web.WebResponse{
                    Success:   false,
                    Message: e.ErrorMessage,
                    Data:  nil,
                })
            case DuplicateError:
                c.JSON(http.StatusConflict, web.WebResponse{
                    Success:   false,
                    Message: e.ErrorMessage,
                    Data:  nil,
                })
            case UnauthorizedError:
                c.JSON(http.StatusUnauthorized, web.WebResponse{
                    Success:   false,
                    Message: e.ErrorMessage,
                    Data:  nil,
                })
            case BadRequestError:
                c.JSON(http.StatusBadRequest, web.WebResponse{
                    Success:   false,
                    Message: e.ErrorMessage,
                    Data:  nil,
                })
            default:
                c.JSON(http.StatusInternalServerError, web.WebResponse{
                    Success:   false,
                    Message: "INTERNAL SERVER ERROR",
                    Data:   nil,
                })
            }
            return
        }
    }
}