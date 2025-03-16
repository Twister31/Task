package handlers

import (
	"Task/internal/userService"
	"Task/internal/web/users"
	"context"
)

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

type UserHandler struct {
	Service *userService.UserService
}

func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	responce := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		responce = append(responce, users.User{
			Id:       &user.ID,
			Email:    &user.Email,
			Password: &user.Password})
	}
	return responce, nil
}

func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	createNewUser := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	newUser, err := u.Service.CreateUser(createNewUser)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &newUser.ID,
		Email:    &newUser.Email,
		Password: &newUser.Password,
	}

	return response, nil
}

func (u *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := u.Service.DeleteUserByID(uint(request.Id))
	if err != nil {
		return nil, err
	}

	deleteUser := users.DeleteUsersId204Response{}

	return &deleteUser, nil
}

func (u *UserHandler) PutUsersId(_ context.Context, request users.PutUsersIdRequestObject) (users.PutUsersIdResponseObject, error) {

	updateUser := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	updateUserN, err := u.Service.UpdateUserByID(uint(request.Id), updateUser)
	if err != nil {
		return nil, err
	}

	response := users.PutUsersId200JSONResponse{
		Id:       &updateUserN.ID,
		Email:    &updateUserN.Email,
		Password: &updateUserN.Password,
	}
	return response, nil
}
