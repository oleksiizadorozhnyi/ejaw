-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items (
                             order_id UUID NOT NULL,
                             product_id UUID NOT NULL,
                             quantity INTEGER NOT NULL DEFAULT 1,
                             price NUMERIC(10,2) NOT NULL,
                             PRIMARY KEY (order_id, product_id),
                             CONSTRAINT fk_order
                                 FOREIGN KEY (order_id)
                                     REFERENCES orders (id)
                                     ON DELETE CASCADE,
                             CONSTRAINT fk_product
                                 FOREIGN KEY (product_id)
                                     REFERENCES products (id)
                                     ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
