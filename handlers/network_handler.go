package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type INetworkHandler interface {
	ListNetworks(fiber.Ctx) error
	CreateNetwork(fiber.Ctx) error
	DeleteNetwork(fiber.Ctx) error
}

type NetworkHandler struct {
	networkService services.INetworkService
}

// ListNetworks godoc
//
// @ID			listNetworks
// @Tags    network
// @Accept  json
// @Produce json
// @Param		args			query		types.ListNetworksArgs	false "List Networks"
// @Success 200				{array}	db.NetworkModel
// @Router  /network [get]
func (n *NetworkHandler) ListNetworks(c fiber.Ctx) error {
	args := types.ListNetworksArgs{OwnerId: lo.ToPtr(fiber.Locals[string](c, "userId"))}

	if err := c.Bind().Query(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	networks, err := n.networkService.ListNetworks(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(networks)
}

// CreateNetwork godoc
//
// @ID			createNetwork
// @Tags    network
// @Accept  json
// @Produce json
// @Param		args			body			types.CreateNetworkArgs	true "List Networks"
// @Success 201				{object}	db.NetworkModel
// @Router  /network [post]
func (n *NetworkHandler) CreateNetwork(c fiber.Ctx) error {
	args := types.CreateNetworkArgs{}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	network, err := n.networkService.CreateNetwork(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(network)
}

// DeleteNetwork godoc
//
// @ID			deleteNetwork
// @Tags    network
// @Accept  json
// @Produce json
// @Param		networkId							path			string	true "Network Id"
// @Success 204										{object}	nil
// @Router  /network/{networkId} 	[delete]
func (n *NetworkHandler) DeleteNetwork(c fiber.Ctx) error {
	networkId := fiber.Params[string](c, "networkId")

	err := n.networkService.DeleteNetwork(types.DeleteNetworkArgs{Id: networkId})
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func NewNetworkHandler(networkService services.INetworkService) INetworkHandler {
	return &NetworkHandler{networkService: networkService}
}
