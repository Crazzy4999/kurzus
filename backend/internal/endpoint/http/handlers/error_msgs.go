package handler

const (
	JSON_TRANSFORM_FAILED    = "transforming request from json failed"
	INVALID_CREDENTIALS      = "invalid credentials"
	ACCESS_TOKEN_EXPIRED     = "access token expired"
	REFRESH_TOKEN_EXPIRED    = "refresh token expired"
	STRING_CONVERSION_FAILED = "converting string to integer failed"

	CREATING_ADDRESS_FAILED = "creating address failed"
	GET_ALL_ADDRESS_FAILED  = "getting all addresses failed"
	UPDATING_ADDRESS_FAILED = "updating address failed"
	DELETING_ADDRESS_FAILED = "deleting address failed"

	PASSWORD_MISMATCH          = "passwords doesn't match"
	GET_ALL_USERS_FAILED       = "getting all users failed"
	EMAIL_TAKEN                = "email already in use"
	PASSWORD_GENERATION_FAILED = "couldn't generate password"
	CREATING_USER_FAILED       = "creating user failed"
	TOKEN_GENERATION_FAILED    = "generating token failed"
	NO_SUCH_USER               = "no user is found with this email"
	SEND_RESET_FAILED          = "sending reset email failed"
	GET_USER_FAILED            = "getting user failed"
	MISSING_RESET_KEY          = "bad url, missing reset key"
	PASSWORD_RESET_FAILED      = "reseting password failed"

	CREATING_CATEGORY_FAILED  = "creating category failed"
	GET_ALL_CATEGORIES_FAILED = "getting all categories failed"
	DELETING_CATEGORY_FAILED  = "deleting category failed"

	CREATING_DRIVER_FAILED = "creating driver failed"
	GET_ALL_DRIVER_FAILED  = "getting all drivers failed"
	UPDATING_DRIVER_FAILED = "updating driver failed"
	DELETING_DRIVER_FAILED = "deleting driver failed"

	CREATING_ITEM_FAILED              = "creating item failed"
	GET_ALL_ITEM_FAILED               = "getting all items failed"
	DELETING_ITEM_FAILED              = "deleting item failed"
	DELETING_MENU_ITEM_BY_ITEM_FAILED = "deleting menu_item by item id failed"

	CREATING_MENU_FAILED              = "creating menu failed"
	GET_ALL_MENU_FAILED               = "getting all menus failed"
	CREATING_MENU_ITEM_FAILED         = "creating menu_item failed"
	DELETING_MENU_FAILED              = "deleting menu failed"
	DELETING_MENU_ITEM_BY_MENU_FAILED = "deleting menu_item by menu id failed"

	GET_ALL_ITEM_MENU_FAILED = "getting all items_menus failed"

	CREATING_ORDER_FAILED = "creating order failed"
	GET_ALL_ORDER_FAILED  = "getting all orders failed"
	UPDATING_ORDER_FAILED = "updaing order failed"

	CREATING_ORDER_MENU_FAILED            = "creating order_menu failed"
	GET_ALL_MENU_BY_ORDER_FAILED          = "getting all menus by order id failed"
	GET_ALL_ORDER_MENU_BY_ORDER_ID_FAILED = "getting all order_menus by order id failed"
	UPDATING_ORDER_MENU_BY_ORDER_FAILED   = "updating order_menu by order id failed"
	DELETING_ORDER_MENU_BY_ORDER_FAILED   = "deleting order_menu by order id failed"

	GET_ALL_SUPPLIER_FAILED  = "getting all suppliers failed"
	SUPPLIER_EXISTS          = "supplier already exists"
	CREATING_SUPPLIER_FAILED = "creating supplier failed"
	GET_SUPPLIER_TYPE_FAILED = "getting supplier type by id failed"
	UPDATING_SUPPLIER_FAILED = "updating supplier failed"
	DELETING_SUPPLIER_FAILED = "deleting supplier failed"

	UPDATING_USER_FAILED = "updating user failed"
	DELETING_USER_FAILED = "deleting user failed"
)
