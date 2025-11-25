-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    platform VARCHAR(100) NOT NULL,
    duration INT NOT NULL,
    image_url VARCHAR(500),
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by VARCHAR(100) NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pilot (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    experience INT NOT NULL,
    team VARCHAR(100) NOT NULL,
    iracing_id VARCHAR(100) NOT NULL,
    image_url VARCHAR(500),
    created_by VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
DROP TABLE pilot;
-- +goose StatementEnd
