-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        buyer_id UUID NOT NULL,
                        order_date TIMESTAMP NOT NULL DEFAULT NOW(),
                        status VARCHAR(50) NOT NULL DEFAULT 'pending',
                        total NUMERIC(10,2) DEFAULT 0,
                        CONSTRAINT fk_buyer
                            FOREIGN KEY (buyer_id)
                                REFERENCES buyers (id)
                                ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
