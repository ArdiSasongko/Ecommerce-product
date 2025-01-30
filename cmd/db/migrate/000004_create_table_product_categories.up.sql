CREATE TABLE IF NOT EXISTS product_categories (
    product_id INT NOT NULL,
    category_id INT NOT NULL,
    foreign key (product_id) references products(id) on delete cascade,
    foreign key (category_id) references categories(id) on delete cascade
);