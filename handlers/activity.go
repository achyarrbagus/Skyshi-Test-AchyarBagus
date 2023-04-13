package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	activitydto "todo-app/dto/activity"
	dto "todo-app/dto/result"
	"todo-app/models"
	"todo-app/repository"

	"github.com/labstack/echo/v4"
)

type handlerActivity struct {
	ActivityRepository repository.ActivityRepository
}

func HandlerActivity(ActivityRepository repository.ActivityRepository) *handlerActivity {
	return &handlerActivity{ActivityRepository}
}

func (h *handlerActivity) FindActivity(c echo.Context) error {
	activity, err := h.ActivityRepository.FindActivity()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertAllActivity(activity)})

}

func (h *handlerActivity) UpdateActivity(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	request := new(activitydto.UpdateActivityRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	activity, err := h.ActivityRepository.GetActivity(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}
	if request.Title != "" {
		activity.Title = request.Title
	}

	activity.UpdatedAt = time.Now()

	activityGroup := models.ActivityGroup{
		ID:        activity.ID,
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}

	data, err := h.ActivityRepository.UpdateActivity(activityGroup)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertActivityRes(data)})

}

func (h *handlerActivity) DeleteActivity(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	activity, err := h.ActivityRepository.GetActivity(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}

	data, err := h.ActivityRepository.DeleteActivity(activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertActivityRes(data)})

}

func (h *handlerActivity) GetActivity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	activity, err := h.ActivityRepository.GetActivity(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}
	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertActivityRes(activity)})

}

func (h *handlerActivity) CreateActivity(c echo.Context) error {

	request := new(activitydto.CreateActivityRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ActivityGroup := models.ActivityGroup{
		Title:     request.Title,
		Email:     request.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newActivity, err := h.ActivityRepository.CreateActivity(ActivityGroup)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	data, _ := h.ActivityRepository.GetActivity(newActivity.ID)

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertActivityRes(data)})
}

func convertAllActivity(u []models.ActivityGroup) []activitydto.ActivityResponse {
	responses := make([]activitydto.ActivityResponse, 0)
	for _, activity := range u {
		response := activitydto.ActivityResponse{
			ID:        activity.ID,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: activity.CreatedAt,
			UpdatedAt: activity.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses
}

func convertActivityRes(u models.ActivityGroup) activitydto.ActivityResponse {
	return activitydto.ActivityResponse{
		ID:        u.ID,
		Title:     u.Title,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ResultMessage(err error, id int) string {
	ids := strconv.Itoa(id)
	if err != nil {
		return fmt.Sprintf("Activity with ID %s Not Found", ids)
	}
	return fmt.Sprintf("Activity with ID %s Updated Successfully", ids)
}

func ResultStatus(code int) string {

	if code == 200 {
		return "Success"
	} else if code == 400 {
		return "Not Found"
	} else {
		return "Pending"
	}
}
