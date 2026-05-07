#psql -U postgres
#\l
#\c go_products
#\dt
#\q

Fields: user_id (PK), full_name, email, phone_number, id_card_number (NIK), is_verified, created_at.

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(100), -- 30 mungkin terlalu sempit untuk nama lengkap
    email VARCHAR(100) UNIQUE, -- Tambahkan UNIQUE agar email tidak duplikat
    phone_number VARCHAR(20), -- Gunakan VARCHAR untuk nomor telepon (karena ada angka '0' di depan)
    password_hash TEXT NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE users ADD balance DECIMAL(15,2) default 0;
ALTER TABLE users ADD role VARCHAR(10) default 'user' CHECK(role IN ('user','admin')); --User --Admin --Ada ga yang Kasih Stock
--/login
--/register
--/update-user
--/delete-user


CREATE TABLE stocks (
    stock_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    symbol VARCHAR(10),     -- contoh: AAPL, TSLA
    name VARCHAR(100),
    price DECIMAL(15,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--/Get-All-Stocks?Params=
--/Get-Stock-Details
--/Update-Stock (Admin)
--/Insert-Ipo (Admin)
--/Delete-Stock (Admin)

CREATE TABLE stock_price_histories (
    history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stock_id UUID REFERENCES stocks(stock_id) ON DELETE CASCADE,
    price DECIMAL(15,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    stock_id UUID REFERENCES stocks(stock_id),
    type VARCHAR(10), -- BUY / SELL
    quantity INT,
    price DECIMAL(15,2),
    total DECIMAL(15,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--/Buy-Stock
--/Sell-Stock


CREATE TABLE portfolios (
    portfolio_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    stock_id UUID REFERENCES stocks(stock_id),
    quantity INT,
    avg_price DECIMAL(15,2),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--/Upsert-Portofolio 

CREATE TABLE trades(

);

CREATE TABLE orders(
    order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(user_id),
    stock_id UUID REFERENCES stocks(stock_id),
    type VARCHAR(10) CHECK (type IN ('BUY','SELL')),
    
);

CREATE TABLE wallet(
    wallet_id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    user_id UUID UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    balance DECIMAL(15,2) DEFAULT 0 CHECK(balance >= 0),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallet_histories (
    wallet_history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_id UUID UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    amount DECIMAL(15,2),
    type VARCHAR(20) CHECK(type IN ('DEPOSIT','WITHDRAW','BUY','SELL')), -- DEPOSIT / WITHDRAW / BUY / SELL
    references_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
--/Insert-Balance