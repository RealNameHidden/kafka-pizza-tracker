
package common

type PizzaOrder struct {
    OrderID string `json:"orderId"`
    User    string `json:"user"`
    Pizza   string `json:"pizza"`
    Size    string `json:"size"`
}
