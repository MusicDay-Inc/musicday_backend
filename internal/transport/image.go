package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"server/internal/core"
)

//	func (h *Handler) uploadAlbumCover(c *gin.Context) {
//		srcId := h.parseUUIDFromParam(c)
//		coverId, err := h.services.Album.GetCoverId(srcId)
//		if err != nil {
//			s, errS := h.services.Song.GetById(srcId)
//			if errS != nil {
//				newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect release id")
//				return
//			}
//			coverId = s.Id
//		}
//		dest := "./img/" + coverId.String() + ".webp"
//
//		file, header, err := c.Request.FormFile("cover")
//		// Determine the file extension
//		ext := filepath.Ext(header.Filename)
//		var img image.Image
//		if ext == ".jpeg" {
//			img, err = jpeg.Decode(file)
//		} else if ext == ".png" {
//			img, err = png.Decode(file)
//		} else if ext == ".webp" {
//			defer func(file multipart.File) {
//				_ = file.Close()
//			}(file)
//			dstF, errF := os.Create(dest)
//			if errF != nil {
//				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//				return
//			}
//			if _, errC := io.Copy(dstF, file); errC != nil {
//				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//				return
//			}
//			c.JSON(http.StatusOK, coverId)
//			return
//		} else {
//			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect file format")
//			return
//		}
//
//		// Create the WebP file
//		webpFile, err := os.Create(dest)
//		if err != nil {
//			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Failed to create WebP file")
//			return
//		}
//		//defer webpFile.Close()
//		// Encode and save the image as WebP
//		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
//		if err != nil {
//			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get encoder")
//			return
//		}
//		if err = webp.Encode(webpFile, img, options); err != nil {
//			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Failed to encode and save WebP file")
//			return
//		}
//		//c.Data(http.StatusOK, "application/octet-stream", data)
//		c.JSON(http.StatusOK, coverId)
//		return
//	}
//
// TODO Prev version
//func (h *Handler) PostAvatar(c *gin.Context) {
//	clientId, err := h.getClientId(c)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
//		return
//	}
//
//	file, err := c.FormFile("picture")
//	if err != nil {
//		c.String(http.StatusBadRequest, "Failed to get file")
//		return
//	}
//	// Save the file to the server
//	if err = c.SaveUploadedFile(file, "./img/"+clientId.String()+".jpeg"); err != nil {
//		c.String(http.StatusInternalServerError, "Failed to save file")
//		return
//	}
//	user, err := h.services.User.UploadAvatar(clientId)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//		return
//	}
//	//c.Data(http.StatusOK, "application/octet-stream", data)
//	c.JSON(http.StatusOK, user)
//	return
//}
//
//func (h *Handler) getReleaseCover(c *gin.Context) {
//	srcId := h.parseUUIDFromParam(c)
//	coverId, err := h.services.Album.GetCoverId(srcId)
//	if err != nil {
//		s, errS := h.services.Song.GetById(srcId)
//		if errS != nil {
//			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect release id")
//			return
//		}
//		coverId = s.Id
//	}
//	// TODO read by buffer
//	imageBytes, err := ioutil.ReadFile("./img/" + coverId.String() + ".jpeg")
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Can't read file")
//		return
//	}
//	// Set the appropriate HTTP headers
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", "attachment; filename=cover.jpeg")
//	// Write the image data to the response body
//	c.Data(http.StatusOK, "application/octet-stream", imageBytes)
//}

// TODO New
func (h *Handler) PostAvatar(c *gin.Context) {
	clientId, err := h.getClientId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
		return
	}

	//file, err := c.FormFile("picture")
	//if err != nil {
	//	c.String(http.StatusBadRequest, "Failed to get file")
	//	return
	//}
	dest := "./img/" + clientId.String() + ".jpeg"
	file, header, err := c.Request.FormFile("picture")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect couldn't get file from the form")
		return
	}
	// Determine the file extension
	ext := filepath.Ext(header.Filename)
	var img image.Image
	resFile, errCreate := os.Create(dest)
	if errCreate != nil {
		logrus.Error(errCreate)
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't create file")
		return
	}

	if ext != ".jpeg" && ext != ".jpg" {
		if ext == ".png" {
			img, err = png.Decode(file)
			if err != nil {
				logrus.Error(err)
				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't decode png file")
				return
			}
			var opt jpeg.Options
			// Качество
			opt.Quality = 95
			if err = jpeg.Encode(resFile, img, &opt); err != nil {
				logrus.Error(err)
				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't encode to jpeg")
				return
			}

			user, errUsr := h.services.User.UploadAvatar(clientId)
			if errUsr != nil {
				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
				return
			}
			//c.Data(http.StatusOK, "application/octet-stream", data)
			sAmount, err := h.services.Review.CountSongReviewsOf(clientId)
			if err != nil {
				sAmount = 0
				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
				return
			}
			aAmount, err := h.services.Review.CountAlbumReviewsOf(clientId)
			if err != nil {
				aAmount = 0
				newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
				return
			}
			bio, err := h.services.User.GetBio(clientId)
			c.JSON(http.StatusOK, map[string]interface{}{
				"user":                 user.ToPayload(),
				"bio":                  bio,
				"is_client_subscribed": false,
				"song_amount":          sAmount,
				"album_amount":         aAmount,
			})
			//c.JSON(http.StatusOK, user)
			return
			//err := jpeg.Encode(w, img, &jpeg.Options{Quality: jpegCompression})
		} else {
			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect file format")
			return
		}
	}

	// Save the file to the server
	//if err = c.SaveUploadedFile(file, "./img/"+clientId.String()+".jpeg"); err != nil {
	//	c.String(http.StatusInternalServerError, "Failed to save file")
	//	return
	//}
	if _, errC := io.Copy(resFile, file); errC != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}

	user, err := h.services.User.UploadAvatar(clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	//c.Data(http.StatusOK, "application/octet-stream", data)
	sAmount, err := h.services.Review.CountSongReviewsOf(clientId)
	if err != nil {
		sAmount = 0
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	aAmount, err := h.services.Review.CountAlbumReviewsOf(clientId)
	if err != nil {
		aAmount = 0
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
		return
	}
	bio, err := h.services.User.GetBio(clientId)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":                 user.ToPayload(),
		"bio":                  bio,
		"is_client_subscribed": false,
		"song_amount":          sAmount,
		"album_amount":         aAmount,
	})
	//c.JSON(http.StatusOK, user)
	return
}

func (h *Handler) getReleaseCover(c *gin.Context) {
	srcId := h.parseUUIDFromParam(c)
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
	imageBytes, err := ioutil.ReadFile("./img/" + coverId.String() + ".jpeg")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Can't read file")
		return
	}
	// Set the appropriate HTTP headers
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=cover.jpeg")
	// Write the image data to the response body
	c.Data(http.StatusOK, "application/octet-stream", imageBytes)
}

func (h *Handler) getAvatar(c *gin.Context) {
	userId := h.parseUUIDFromParam(c)
	// TODO read by buffer
	imageBytes, err := ioutil.ReadFile("./img/" + userId.String() + ".jpeg")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Can't read file")
		return
	}
	// Set the appropriate HTTP headers
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=avatar.jpeg")
	// Write the image data to the response body
	c.Data(http.StatusOK, "application/octet-stream", imageBytes)
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
		if err = c.SaveUploadedFile(file, "./img/"+albumId.String()+".jpeg"); err != nil {
			c.String(http.StatusInternalServerError, "Failed to save file")
			return
		}
		c.JSON(http.StatusOK, s)
		return
	}
	// Save the file to the server
	if err = c.SaveUploadedFile(file, "./img/"+albumId.String()+".jpeg"); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	c.JSON(http.StatusOK, a)
}

//func (h *Handler) PostAvatar(c *gin.Context) {
//	clientId, err := h.getClientId(c)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get clientId from context")
//		return
//	}
//	// Create the WebP file destination
//	dest := "./img/" + clientId.String() + ".webp"
//
//	file, header, err := c.Request.FormFile("picture")
//	// Determine the file extension
//
//	ext := filepath.Ext(header.Filename)
//	var img image.Image
//	if ext == ".jpeg" {
//		img, err = jpeg.Decode(file)
//	} else if ext == ".png" {
//		img, err = png.Decode(file)
//	} else if ext == ".webp" {
//		defer func(file multipart.File) {
//			_ = file.Close()
//		}(file)
//		dstF, errF := os.Create(dest)
//		if errF != nil {
//			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//			return
//		}
//		if _, errC := io.Copy(dstF, file); errC != nil {
//			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//			return
//		}
//		user, errU := h.services.User.UploadAvatar(clientId)
//		if errU != nil {
//			newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//			return
//		}
//		//c.Data(http.StatusOK, "application/octet-stream", data)
//		c.JSON(http.StatusOK, user)
//		return
//	} else {
//		newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect file format")
//		return
//	}
//
//	// Create the WebP file
//	webpFile, err := os.Create(dest)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Failed to create WebP file")
//		return
//	}
//	//defer webpFile.Close()
//	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "couldn't get encoder")
//		return
//	}
//	// Encode and save the image as WebP
//	if err = webp.Encode(webpFile, img, options); err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, "Failed to encode and save WebP file")
//		return
//	}
//	user, err := h.services.User.UploadAvatar(clientId)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, core.CodeInternalError, core.ErrInternal.Error())
//		return
//	}
//	//c.Data(http.StatusOK, "application/octet-stream", data)
//	c.JSON(http.StatusOK, user)
//	return
//}
//func (h *Handler) uploadAlbumCover(c *gin.Context) {
//	albumId := h.parseUUIDFromParam(c)
//	file, err := c.FormFile("cover")
//	if err != nil {
//		c.String(http.StatusBadRequest, "Failed to get file")
//		return
//	}
//	a, err := h.services.Album.GetById(albumId)
//	if err != nil {
//		s, errS := h.services.Song.GetById(albumId)
//		if errS != nil {
//			newErrorResponse(c, http.StatusBadRequest, core.CodeIncorrectBody, "incorrect release id")
//			return
//		}
//		if err = c.SaveUploadedFile(file, "./img/"+albumId.String()+".jpeg"); err != nil {
//			c.String(http.StatusInternalServerError, "Failed to save file")
//			return
//		}
//		c.JSON(http.StatusOK, s)
//		return
//	}
//	// Save the file to the server
//	if err = c.SaveUploadedFile(file, "./img/"+albumId.String()+".jpeg"); err != nil {
//		c.String(http.StatusInternalServerError, "Failed to save file")
//		return
//	}
//
//	c.JSON(http.StatusOK, a)
//}
