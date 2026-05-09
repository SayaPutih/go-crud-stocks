#psql -U postgres
#\l
#\c go_products
#\dt
#\q

Fields: user_id (PK), full_name, email, phone_number, id_card_number (NIK), is_verified, created_at.

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(100) NOT NULL, -- 30 mungkin terlalu sempit untuk nama lengkap
    email VARCHAR(100) UNIQUE  NOT NULL, -- Tambahkan UNIQUE agar email tidak duplikat
    phone_number VARCHAR(20) NOT NULL, -- Gunakan VARCHAR untuk nomor telepon (karena ada angka '0' di depan)
    password_hash TEXT NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--ALTER TABLE users ADD balance DECIMAL(15,2) default 0;
ALTER TABLE users ADD role VARCHAR(10) default 'user' CHECK(role IN ('user','admin')); --User --Admin --Ada ga yang Kasih Stock
--/login
--/register
--/update-user
--/delete-user


CREATE TABLE stocks (
    stock_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    symbol VARCHAR(10) UNIQUE NOT NULL,     -- contoh: AAPL, TSLA
    name VARCHAR(100) NOT NULL,
    price DECIMAL(15,2) NOT NULL CHECK(price > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--/Get-All-Stocks?Params=
--/Get-Stock-Details
--/Update-Stock (Admin)
--/Insert-Ipo (Admin)
--/Delete-Stock (Admin)

CREATE TABLE stock_price_histories (
    history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stock_id UUID NOT NULL REFERENCES stocks(stock_id) ON DELETE CASCADE,
    price DECIMAL(15,2) NOT NULL CHECK(price > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    stock_id UUID NOT NULL REFERENCES stocks(stock_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK(quantity > 0),
    price DECIMAL(15,2) NOT NULL CHECK(price > 0),
    total DECIMAL(15,2) NOT NULL CHECK(total >= 0),
    type VARCHAR(10) NOT NULL CHECK(type IN ('BUY','SELL')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--/Buy-Stock
--/Sell-Stock


CREATE TABLE portfolios (
    portfolio_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    stock_id UUID NOT NULL  REFERENCES stocks(stock_id) ON DELETE CASCADE,
    quantity INT DEFAULT 0 CHECK(quantity >= 0),
    avg_price DECIMAL(15,2) DEFAULT 0 CHECK(avg_price >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id,stock_id)
);
--/Upsert-Portofolio 

CREATE TABLE orders(
    order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL  REFERENCES users(user_id) ON DELETE CASCADE,
    stock_id UUID NOT NULL  REFERENCES stocks(stock_id) ON DELETE CASCADE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('BUY','SELL')),
    status VARCHAR(20)  NOT NULL DEFAULT 'PENDING' CHECK(status IN ('PENDING','COMPLETED','CANCELLED')),
    quantity INT NOT NULL CHECK(quantity > 0),
    price DECIMAL(15,2) NOT NULL CHECK(price > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallet(
    wallet_id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    user_id UUID NOT NULL UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    balance DECIMAL(15,2) DEFAULT 0 CHECK(balance >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallet_histories (
    wallet_history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_id UUID NOT NULL REFERENCES wallet(wallet_id) ON DELETE CASCADE,
    amount DECIMAL(15,2) NOT NULL CHECK(amount > 0),
    type VARCHAR(20) NOT NULL  CHECK(type IN ('DEPOSIT','WITHDRAW','BUY','SELL')), -- DEPOSIT / WITHDRAW / BUY / SELL
    references_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--/Insert-Balance

CREATE INDEX idx_stock_symbol ON stocks(symbol);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_stock_id ON transactions(stock_id);
CREATE INDEX idx_portfolios_user_id ON portfolios(user_id);
CREATE INDEX idx_wallet_user_id ON wallet(user_id);
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_wallet_histories_wallet_id ON wallet_histories(wallet_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_transactions_type ON transactions(type);
CREATE INDEX idx_transactions_user_created ON transactions(user_id, created_at DESC);
