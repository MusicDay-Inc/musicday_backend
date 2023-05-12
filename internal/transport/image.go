package transport

import (
	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"server/internal/core"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (h *Handler) uploadAlbumCover(c *gin.Context) {
	albumId := h.parseUUIDFromParam(c)
	file, err := c.FormFile("cover")
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to get file")
		return
	}
	a, err := h.services.Album.GetById(albumId)
	if err != nil {
		s, errS := h.services.Song.GetById(albumId)
		if errS != nil {
			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect release id")
			return
		}
		if err = c.SaveUploadedFile(file, "./img/"+albumId.String()+".png"); err != nil {
			c.String(http.StatusInternalServerError, "Failed to save file")
			return
		}
		c.JSON(http.StatusOK, s)
		return
	}
	// Save the file to the server
	if err = c.SaveUploadedFile(file, "./img/"+albumId.String()+".png"); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	c.JSON(http.StatusOK, a)
}

func (h *Handler) PostAvatar(c *gin.Context) {
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}
	// PREV
	//var form Form
	//err = c.ShouldBind(&form)
	//if err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect file form")
	//	return
	//}
	// PREV 1.1 GOOD
	//file, err := c.FormFile("picture")
	//if err != nil {
	//	c.String(http.StatusBadRequest, "Failed to get file")
	//	return
	//}
	//// Save the file to the server
	//if err = c.SaveUploadedFile(file, "./img/"+clientId.String()+".png"); err != nil {
	//	c.String(http.StatusInternalServerError, "Failed to save file")
	//	return
	//}

	// Get the uploaded file from the request
	//file, err := c.FormFile("picture")
	//if err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect file form")
	//	return
	//}

	// Determine the file extension

	// Create the WebP file destination
	dest := "./img/" + clientId.String() + ".webp"

	// Open the uploaded file
	//src, err := file.Open()
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "Failed to open file")
	//	return
	//}
	//defer func(src multipart.File) {
	//	err = src.Close()
	//	if err != nil {
	//		logrus.Error(err)
	//	}
	//}(src)

	//// Decode the image file
	//img, err := webp.Decode(src)
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "Failed to decode image")
	//	return
	//}
	file, header, err := c.Request.FormFile("picture")
	ext := filepath.Ext(header.Filename)
	var img image.Image
	if ext == ".jpeg" {
		img, err = jpeg.Decode(file)
	} else if ext == ".png" {
		img, err = png.Decode(file)
	} else if ext == ".webp" {
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)
		dstF, errF := os.Create(dest)
		if errF != nil {
			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
			return
		}
		if _, errC := io.Copy(dstF, file); errC != nil {
			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
			return
		}
		user, errU := h.services.User.UploadAvatar(clientId)
		if errU != nil {
			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
			return
		}
		//c.Data(http.StatusOK, "application/octet-stream", data)
		c.JSON(http.StatusOK, user)
		return
	} else {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect file format")
		return
	}

	// Create the WebP file
	webpFile, err := os.Create(dest)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Failed to create WebP file")
		return
	}
	//defer webpFile.Close()

	// Encode and save the image as WebP
	if err = webp.Encode(webpFile, img, &webp.Options{}); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Failed to encode and save WebP file")
		return
	}
	// TODO mb  delete?
	//data, err := ioutil.ReadFile("./img/" + clientId.String() + ".png")
	//if err != nil {
	//	//c.String(http.StatusInternalServerError, "Failed to read image")
	//	newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Can't read file")
	//	return
	//}
	user, err := h.services.User.UploadAvatar(clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	//c.Data(http.StatusOK, "application/octet-stream", data)
	c.JSON(http.StatusOK, user)
	return
}

func (h *Handler) getReleaseCover(c *gin.Context) {
	srcId := h.parseUUIDFromParam(c)
	//clientId, err := h.getClientId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
	//	return
	//}
	coverId, err := h.services.Album.GetCoverId(srcId)
	if err != nil {
		s, errS := h.services.Song.GetById(srcId)
		if errS != nil {
			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect release id")
			return
		}
		coverId = s.Id
	}

	// TODO read by buffer
	imageBytes, err := ioutil.ReadFile("./img/" + coverId.String() + ".png")
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to read image")
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Can't read file")
		return
	}

	// Set the appropriate HTTP headers
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=image.png")

	// Write the image data to the response body
	c.Data(http.StatusOK, "application/octet-stream", imageBytes)
}

func (h *Handler) getAvatar(c *gin.Context) {
	userId := h.parseUUIDFromParam(c)
	//clientId, err := h.getClientId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
	//	return
	//}
	// TODO read by buffer
	imageBytes, err := ioutil.ReadFile("./img/" + userId.String() + ".webp")
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to read image")
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Can't read file")
		return
	}

	// Set the appropriate HTTP headers
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=avatar.webp")

	// Write the image data to the response body
	c.Data(http.StatusOK, "application/octet-stream", imageBytes)
}
