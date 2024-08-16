package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud/service"
	"go-crud/types"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}
		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})
	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateUserResquest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Create 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다.")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserResquest

	u.userService.Update(nil, nil)

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteUserResquest

	u.userService.Delete(nil)

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
	})
}
